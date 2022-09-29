package starter

import (
	"gin-quick-starter/util/jsonx"
	"github.com/gin-gonic/gin"
)

var funcPool = make(map[string]gin.HandlerFunc)
var engine *gin.Engine

func BindStarter(obj *jsonx.JObj) []*RouterNode {
	ret := make([]*RouterNode, 0)
	for path, value := range *obj {
		if rootNode, hit := jsonx.GetJObjFromInterface(value); hit {
			ret = append(ret, makeTreeFromJObj(path, rootNode))
		} else {
			panic("wrong struct json for gin-quick-starter")
		}
	}
	return ret
}
func GinStart(rootNodes []*RouterNode) *gin.Engine {
	if !checkEngine() {
		panic("engine haven't init")
	}
	for _, rootNode := range rootNodes {
		rootNode.Init(engine)
	}
	return engine
}
func AddHandler(name string, addFunc gin.HandlerFunc) bool {
	if _, hit := funcPool[name]; hit {
		return false
	}
	funcPool[name] = addFunc
	return true
}
func getHandlerByName(name string) gin.HandlerFunc {
	if handlerFunc, hit := funcPool[name]; hit {
		return handlerFunc
	}
	panic("func pool not exists " + name + ",maybe you didn't register it")
}
func NewFromExist(exist *gin.Engine) bool {
	if checkEngine() {
		return false
	}
	engine = exist
	return true
}
func New() bool {
	if checkEngine() {
		return false
	}
	engine = gin.New()
	return true
}
func Default() bool {
	if checkEngine() {
		return false
	}
	engine = gin.New()
	return true
}
func checkEngine() bool {
	if engine == nil {
		return false
	}
	return true
}
