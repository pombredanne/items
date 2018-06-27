package handlers

import (
	"github.com/gin-gonic/gin"
	"api-libs/rsp"
	"adv-admin-api/models"
)

func Actions(c *gin.Context) {
	var req models.ActionReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	actions := &models.Actions{}
	var total int64
	if err := models.GetActions([]int{3,4}, &total, &req, actions); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Page(len(*actions), total, actions))
}
