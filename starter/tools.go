package starter

import (
	"gin-quick-starter/util/jsonx"
	"github.com/gin-gonic/gin"
)

func MakeTreeFromJObj(path string, j *jsonx.JObj) *RouterNode {
	thisNode := RouterNode{SpecificPath: path, MiddleWares: make([]gin.HandlerFunc, 0), SonNodes: make([]RouterNode, 0), RouterHandlers: make([]Handler, 0)}
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

	}
}
