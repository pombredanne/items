package handlers

import (
	"github.com/gin-gonic/gin"
	"api-libs/rsp"
	"fmt"
)

func Upload(c *gin.Context) {
	fileInfo := &fileInfo{}
	if err := forward(fileInfo, c.Request); err != nil || fileInfo.Status == -1 {
		c.JSON(rsp.Error("上传失败", err))
		return
	}
	if fileInfo.Status != 0 {
		c.JSON(rsp.Error(fmt.Sprintf("%s(%s)", fileInfo.Msg, fileInfo.FileName)))
		return
	}
	c.JSON(rsp.Base(fileInfo))

}
