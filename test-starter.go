package main

import (
	"fmt"
	"github.com/MH2608/gin-quick-starter/starter"
	"github.com/MH2608/gin-quick-starter/util/jsonx"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	InitTest()
	content, err := os.ReadFile("starter/testKernal.json")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	gotJson := jsonx.DecodeFromJson(string(content))
	starter.New()
	starter.GinStart(starter.BindStarter(gotJson)).Run(":8080")
}
func InitTest() {
	starter.AddHandler("testHandler", testHandler)
	starter.AddHandler("check_authorize", checkAuthorize)
	starter.AddHandler("set_token", setToken)
	starter.AddHandler("check_login", checkLogin)
	starter.AddHandler("auth_token", authToken)
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
