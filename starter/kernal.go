package starter

import "github.com/gin-gonic/gin"

var funcPool = make(map[string]gin.HandlerFunc)
var engine *gin.Engine

func AddHandler(name string, addFunc gin.HandlerFunc) bool {
	if _, hit := funcPool[name]; hit {
		return false
	}
	funcPool[name] = addFunc
	return true
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
