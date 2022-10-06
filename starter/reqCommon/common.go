package reqCommon

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func DoSimpleService(c *gin.Context, req BaseReq, doRes func(c *gin.Context, retData interface{}), doErr func(c *gin.Context, err error)) {
	if err := c.ShouldBind(&req); err != nil {
		doErr(c, err)
	}
	if !req.IsValid() {
		doErr(c, errors.New("bind req failed"))
	}
	if retData, err := req.DoService(); err != nil {
		doErr(c, err)
	} else {
		doRes(c, retData)
	}
}
