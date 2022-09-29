package starter

import (
	"fmt"
	"gin-quick-starter/util/jsonx"
	"github.com/gin-gonic/gin"
	"strings"
)

func makeTreeFromJObj(path string, j *jsonx.JObj) *RouterNode {
	thisNode := &RouterNode{SpecificPath: path, MiddleWares: make([]gin.HandlerFunc, 0), SonNodes: make([]*RouterNode, 0), RouterHandlers: make([]Handler, 0)}
	if middlewares := j.GetJArr("middlewares"); middlewares != nil {
		for _, funcName := range *middlewares {
			if got, ok := funcName.(string); ok {
				if function, hit := funcPool[got]; hit {
					thisNode.MiddleWares = append(thisNode.MiddleWares, function)
				} else {
					panic("func pool not exists " + got + ",maybe you didn't register it")
				}
			} else {
				panic("value of middleware must be string")
			}
		}
	}
	if routerHandlers := j.GetJArr("router_handlers"); routerHandlers != nil {
		for _, got := range *routerHandlers {
			if handler, hit := jsonx.GetJObjFromInterface(got); hit {
				thisNode.RouterHandlers = append(thisNode.RouterHandlers, Handler{SpecificPath: handler.GetString("path"), FuncType: handler.GetString("type"), HandlerFunc: getHandlerByName(handler.GetString("func"))})
			} else {
				panic("router_handlers format wrong with path " + path)
			}
		}
	}
	if sonNodes := j.GetJObj("son_nodes"); sonNodes != nil {
		for key, value := range *sonNodes {
			if node, hit := jsonx.GetJObjFromInterface(value); hit {
				thisNode.SonNodes = append(thisNode.SonNodes, makeTreeFromJObj(key, node))
			}
		}
	}
	return thisNode
}
func RestfulToGin(restfulType, path string, ginRouter *gin.RouterGroup) (func(func(c *gin.Context)), error) {
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
	return func(f func(c *gin.Context)) {
		retFunc(path, f)
	}, err
}
