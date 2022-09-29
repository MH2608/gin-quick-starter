package starter

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	init()
}
type Handler struct {
	SpecificPath string
	FuncType     string
	HandlerFunc  gin.HandlerFunc
}
type RouterNode struct {
	MiddleWares    []gin.HandlerFunc
	SonNodes       []*RouterNode
	SpecificPath   string
	RouterHandlers []Handler
}

func (node *RouterNode) init(router gin.IRouter) {
	if node == nil {
		return
	}
	var thisGroup *gin.RouterGroup
	if node.MiddleWares != nil {
		thisGroup = router.Group(node.SpecificPath, node.MiddleWares...)
	}
	if node.RouterHandlers != nil {
		for _, routerFunc := range node.RouterHandlers {
			if ginFunc, err := restfulToGin(routerFunc.FuncType, routerFunc.SpecificPath, thisGroup); err == nil {
				ginFunc(routerFunc.HandlerFunc)
			}
		}
	}
	if node.SonNodes != nil {
		for _, routerNode := range node.SonNodes {
			routerNode.init(thisGroup)
		}
	}

}
