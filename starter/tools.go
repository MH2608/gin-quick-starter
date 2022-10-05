package starter

import (
	"fmt"
	"github.com/MH2608/gin-quick-starter/util/jsonx"
	"github.com/gin-gonic/gin"
	"strings"
)

func makeTreeFromJObj(path string, obj *jsonx.JObj) *RouterNode {
	thisNode := &RouterNode{SpecificPath: path, MiddleWares: make([]gin.HandlerFunc, 0), SonNodes: make([]*RouterNode, 0), RouterHandlers: make([]Handler, 0)}
	initMiddleWares(obj, thisNode)
	initHandlers(path, obj, thisNode)
	initHookers(obj, thisNode)
	if sonNodes := obj.GetJObj("son"); sonNodes != nil {
		for key, value := range *sonNodes {
			if node, hit := jsonx.GetJObjFromInterface(value); hit {
				thisNode.SonNodes = append(thisNode.SonNodes, makeTreeFromJObj(key, node))
			}
		}
	}
	return thisNode
}
func initHookers(j *jsonx.JObj, thisNode *RouterNode) {
	hookers := make([]Hooker, 0)
	if hookerNames := j.GetJArr("hookers"); hookerNames != nil {
		for _, hookerName := range *hookerNames {
			hookers = append(hookers, getHookerByName(jsonx.InterfaceToString(hookerName)))
		}
	}
	thisNode.Hookers = hookers
}
func initHandlers(path string, j *jsonx.JObj, thisNode *RouterNode) {
	if routerHandlers := j.GetJArr("handlers"); routerHandlers != nil {
		funcMapper := make(map[string][]gin.HandlerFunc)
		for _, got := range *routerHandlers {
			handler := strings.Split(jsonx.InterfaceToString(got), "/")
			if len(handler) != 2 {
				panic("wrong format of handler:" + jsonx.InterfaceToString(got) + " path:" + path)
			}
			if funcSlice, hit := funcMapper[handler[0]]; hit {
				funcMapper[handler[0]] = append(funcSlice, getHandlerByName(handler[1]))
			} else {
				newFuncSlice := make([]gin.HandlerFunc, 0)
				funcMapper[handler[0]] = append(newFuncSlice, getHandlerByName(handler[1]))
			}

		}
		for httpType, handlers := range funcMapper {
			thisNode.RouterHandlers = append(thisNode.RouterHandlers, Handler{FuncType: httpType, HandlerFuncs: handlers})
		}
	}
}

func initMiddleWares(j *jsonx.JObj, thisNode *RouterNode) {
	if middlewares := j.GetJArr("mid"); middlewares != nil {
		for _, funcName := range *middlewares {
			if got, ok := funcName.(string); ok {
				if function, hit := funcPool[got]; hit {
					thisNode.MiddleWares = append(thisNode.MiddleWares, function)
				} else {
					panic("func pool not exists " + got + ",maybe you didn't register it")
				}
			} else {
				panic("value of mid must be string")
			}
		}
	}
}
func restfulToGin(restfulType string, ginRouter *gin.RouterGroup) (func(...gin.HandlerFunc), error) {
	var retFunc func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	var err error
	switch strings.ToLower(restfulType) {
	case "get":
		retFunc = ginRouter.GET
	case "post":
		retFunc = ginRouter.POST
	case "put":
		retFunc = ginRouter.PUT
	case "delete":
		retFunc = ginRouter.DELETE
	case "head":
		retFunc = ginRouter.HEAD
	case "options":
		retFunc = ginRouter.OPTIONS
	default:
		err = fmt.Errorf("wrong http type")
	}
	return func(f ...gin.HandlerFunc) {
		retFunc("", f...)
	}, err
}
