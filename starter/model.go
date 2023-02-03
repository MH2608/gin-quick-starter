package starter

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	init()
}
type Hooker func(r *gin.RouterGroup)
type Handler struct {
	FuncType     string
	HandlerFuncs []gin.HandlerFunc
}
type RouterNode struct {
	MiddleWares    []gin.HandlerFunc
	SonNodes       []*RouterNode
	SpecificPath   string
	RouterHandlers []Handler
	Hookers        []Hooker
}

func (node *RouterNode) init(router gin.IRouter) {
	if node == nil {
		return
	}
	var thisGroup = router.Group(node.SpecificPath)
	if node.Hookers != nil {
		for _, hooker := range node.Hookers {
			hooker(thisGroup)
		}
	}
	if node.MiddleWares != nil {
		thisGroup.Use(node.MiddleWares...)
	}
	if node.RouterHandlers != nil {
		for _, routerFunc := range node.RouterHandlers {
			if ginFunc, err := restfulToGin(routerFunc.FuncType, thisGroup); err == nil {
				ginFunc(routerFunc.HandlerFuncs...)
			}
		}
	}
	if node.SonNodes != nil {
		for _, routerNode := range node.SonNodes {
			routerNode.init(thisGroup)
		}
	}

}
