package starter

import (
	"github.com/MH2608/gin-quick-starter/util/jsonx"
	"github.com/gin-gonic/gin"
)

var hookerPool = make(map[string]Hooker)
var funcPool = make(map[string]gin.HandlerFunc)
var engine *gin.Engine

func AddHooker(name string, hooker Hooker) bool {
	if _, hit := hookerPool[name]; hit {
		return false
	}
	hookerPool[name] = hooker
	return true
}

func getHookerByName(name string) Hooker {
	if hooker, hit := hookerPool[name]; hit {
		return hooker
	}
	panic("hooker pool not exists " + name + ",maybe you didn't register it")
}

// BindStarter construct all RouterNode by your given json Map
// Notice: this function do not bind your all handler in your gin-router-tree
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

// GinStart call this function after you bound all routerNode by func BindStarter
// of course you should init engine before call it
func GinStart(rootNodes []*RouterNode) *gin.Engine {
	if !checkEngine() {
		panic("engine haven't init")
	}
	for _, rootNode := range rootNodes {
		rootNode.init(engine)
	}
	return engine
}

// AddHandler for add a handler into handler pool
// actually you should call this function to register all your handler ,so you can register your handler/middleware to the gin-route-tree
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

// NewFromExist init gin-quick-starter engine by your exists gin.Engine
func NewFromExist(exist *gin.Engine) bool {
	if checkEngine() {
		return false
	}
	engine = exist
	return true
}

// New init gin-quick-starter engine by gin.New
func New() bool {
	if checkEngine() {
		return false
	}
	engine = gin.New()
	return true
}

// Default init gin-quick-starter engine by gin.Default
func Default() bool {
	if checkEngine() {
		return false
	}
	engine = gin.Default()
	return true
}
func checkEngine() bool {
	if engine == nil {
		return false
	}
	return true
}
