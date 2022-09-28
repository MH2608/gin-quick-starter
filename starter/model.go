package starter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
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
	SonNodes       []RouterNode
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
