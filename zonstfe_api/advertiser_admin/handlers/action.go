package handlers

import (
	"net/http"
	"zonstfe_api/advertiser_admin/models"
	"zonstfe_api/common/my_context"
)

func (c *Handler) Actions(w http.ResponseWriter, r *http.Request) {
	var req models.ActionReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	actions := &models.Actions{}
	var total int64
	if err := models.GetActions([]int{3, 4}, &total, &req, actions); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonPage(w, len(*actions), total, actions)
}
