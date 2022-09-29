package starter

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Init()
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

func (node *RouterNode) Init(router gin.IRouter) {
	if node == nil {
		return
	}
	var thisGroup *gin.RouterGroup
	if node.MiddleWares != nil {
		thisGroup = router.Group(node.SpecificPath, node.MiddleWares...)
	}
	if node.RouterHandlers != nil {
		for _, routerFunc := range node.RouterHandlers {
			if ginFunc, err := RestfulToGin(routerFunc.FuncType, routerFunc.SpecificPath, thisGroup); err == nil {
				ginFunc(routerFunc.HandlerFunc)
			}
		}
	}
	if node.SonNodes != nil {
		for _, routerNode := range node.SonNodes {
			routerNode.Init(thisGroup)
		}
	}

}
