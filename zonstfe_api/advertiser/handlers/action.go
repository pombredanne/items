package handlers

import (
	"net/http"
	"zonstfe_api/advertiser/models"
	"zonstfe_api/common/my_context"
)

// 操作记录
func (c *Handler) Actions(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ActionReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	actions := &models.Actions{}
	var total int64
	if err := models.GetActions(user["id"], &total, &req, actions); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonPage(w, len(*actions), total, actions)
}
