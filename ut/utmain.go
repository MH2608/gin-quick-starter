package main

import (
	"fmt"
	"github.com/MH2608/gin-quick-starter/starter"
	"github.com/gin-gonic/gin"
)

func main() {
	InitTest()
	fmt.Printf("%v", starter.StartFromJsonFile("./ut/testKernal.json").Routes())

}
func InitTest() {
	starter.AddHandler("testHandler", testHandler)
	starter.AddHandler("checkAuthorize", checkAuthorize)
	starter.AddHandler("setToken", setToken)
	starter.AddHandler("checkLogin", checkLogin)
	starter.AddHandler("authToken", authToken)
}
func testHandler(c *gin.Context) {
	c.JSON(0, "helloworld")
}
func checkAuthorize(c *gin.Context) {
	_, err := c.Cookie("authToken")
	if err != nil {
		c.Set("auth", false)
	}
	c.Next()
}
func setToken(c *gin.Context) {
	if c.GetBool("auth") {
		c.SetCookie("authToken", "游客", 60000, "/v1/admin", "localhost", false, false)
	}
}
func checkLogin(c *gin.Context) {
	c.JSON(0, "checkLogin")
}
func authToken(c *gin.Context) {
	c.JSON(0, "authToken")
}
