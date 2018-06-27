package handlers

import (
	"github.com/gin-gonic/gin"
	"api-libs/rsp"
	"adv-admin-api/models"
)

// 人群包列表
func Segments(c *gin.Context) {
	var req models.SegmentReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	segments := &models.Segments{}
	var total int64
	if err := models.GetSegments(&total, &req, segments); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*segments), total, segments))

}
