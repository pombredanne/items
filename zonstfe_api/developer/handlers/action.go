package handlers

import (
	"net/http"
	"zonstfe_api/advertiser/models"
	"zonstfe_api/common/my_context"
)

func (c *Handler) Actions(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ActionReq
	actions := &models.Actions{}
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	var total int64
	if err := models.GetActions(user["id"], &total, &req, actions); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonPage(w, len(*actions), total, actions)
}
