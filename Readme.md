# Gin-quick-starter
#### 写在前面：gin是一套非常优秀的框架，但包括我以内的很多人，对于gin的结构化比较难以掌握，一个大脑需要同时去思考controller的架构和service的实现，似乎有些不必要的麻烦？
#### gin-quick-starter框架就是基于此诞生的。
### for who？
1. 希望controller和service层完全独立的开发者，能够将实现提前，而路由的配置放在开发阶段的后面。
2. 刚刚入手golang的web开发的工程师，我们希望您能更少花精力在阅读不必要的源码上，而能专注实现service层的handler，仅需简单的starter和json即可启动一个中小型项目。
3. 项目复杂度已然上了一个层级，希望接手的同事能快速了解整个项目结构树，快速步入工作流的架构师。
### quick start with Gin-quick-starter
1. 安装依赖包
``` shell
go get github.com/MH2608/gin-quick-starter
```
2. 利用json文件一键启动
```go
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

```
```json
{
  "v1": {
    "son": {
      "test": {
        "handlers": ["get/testHandler"]
      },
      "admin": {
        "mid": ["checkAuthorize","setToken"],
        "son": {
          "login": {
            "handlers": ["post/checkLogin","post/authToken"]
          }
        }
      }
    }
  }
}
```
#### 聪明的您，看到每一个节点的mid、son、handlers和hookers(未在example中标出)，一定能猜到它们的含义，若希望详细学习，请移步开发文档：（暂未开发）
## 维护者：sheepMo&&shen
