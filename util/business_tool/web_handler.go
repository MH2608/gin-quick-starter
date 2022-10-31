package business_tool

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// MakeCommonHandler
// make a handler for normal business;just init a construct of BusinessReq with it
// for example
// type testReq struct {
// }
//
//	func (t testReq) DoService() (interface{}, error) {
//		return "OK", nil
//	}
//
//	func (t testReq) IsValid() bool {
//		return true
//	}
//
//	func test() {
//		Handler := MakeCommonHandler(testReq{}, ErrorFunc, SuccessFunc)
//		engine := gin.New()
//		engine.GET("test", Handler)
//	}
//
//	func ErrorFunc(c *gin.Context, code int, err error) {
//		c.JSON(code, struct {
//			Error string
//		}{
//			Error: err.Error(),
//		})
//	}
//
//	func SuccessFunc(c *gin.Context, data interface{}) {
//		c.JSON(0, data)
//	}
//
// /*
func MakeCommonHandler(req BusinessReq, errorFunc func(c *gin.Context, code int, err error), successFunc func(c *gin.Context, data interface{})) gin.HandlerFunc {
	ret := func(c *gin.Context) {
		if err := c.ShouldBind(&req); err != nil {
			errorFunc(c, 400, errors.New("req bind failed"))
			return
		}
		if !req.IsValid() {
			errorFunc(c, 400, errors.New("req is not Valid"))
			return
		}
		if data, err := req.DoService(); err != nil {
			errorFunc(c, 401, err)
			return
		} else {
			successFunc(c, data)
		}
	}
	return ret
}
