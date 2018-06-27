package handlers

import (
	"zonstfe_api/common/my_context"
	"net/http"
	"zonstfe_api/developer_admin/models"
	"github.com/go-chi/chi"
)

// app列表
func (c *Handler) Apps(w http.ResponseWriter, r *http.Request) {
	var req models.AppReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	apps := &models.Apps{}
	var total int64
	if err := models.GetApps(&total, &req, apps); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonPage(w, len(*apps), total, apps)
}

// app审核
func (c *Handler) AppReview(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	appId := chi.URLParam(r, "app_id")
	appReview, app := &models.AppReview{}, &models.App{}
	if err := c.Bind(r, appReview); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := models.GetApp(appId, app); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	status := 1
	if *appReview.ReviewType == -1 {
		status = -1
	}
	if err := models.ReviewApp(appId, status); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, nil)
	c.AppCache(appId, "APP审核成功缓存更新")
	c.LogAction(user["id"], *app.UserId, appId, my_context.ActionModule.App, "审核", r)

}

// 单个APP
func (c *Handler) AppOne(w http.ResponseWriter, r *http.Request) {
	app := &models.App{}
	appId := chi.URLParam(r, "app_id")
	if err := models.GetApp(appId, app); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, app)
}
