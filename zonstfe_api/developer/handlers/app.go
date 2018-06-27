package handlers

import (
	"zonstfe_api/common/my_context"
	"zonstfe_api/developer/models"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
)

// app列表
func (c *Handler) Apps(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.AppReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, err, err)
		return
	}
	apps := &models.Apps{}
	var total int64
	if err := models.GetApps(user["id"], &total, &req, apps); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonPage(w, len(*apps), total, apps)

}

// 单个APP
func (c *Handler) AppOne(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	app := &models.App{}
	appId := chi.URLParam(r, "appId")
	if err := models.GetApp(user["id"], appId, app); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, app)
}

// app创建
func (c *Handler) AppCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	app, oldApps, category_limit, reward := &models.App{}, &models.Apps{}, &models.CategoryLimit{}, &models.Reward{}
	if err := c.Bind(r, app); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	// 判断是否存在
	if err := models.CheckAppExist(user["id"], *app.BundleId, *app.Os, oldApps); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	if len(*oldApps) > 0 {
		c.JsonError(w, "已经包含相同 包名和OS APP", nil)
		return
	}
	// 广告app 分类限制格式验证
	if err := c.BindJson(string(*app.CategoryLimit), category_limit); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if *category_limit.Enable == 0 {
		category_limit.List = make([]string, 0)
	}

	// 激励格式验证
	if err := c.BindJson(string(*app.Reward), reward); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if user["deal_type"].(string) == "bidding" {
		for k, _ := range app.SlotMap {
			app.SlotMap[k] = 1
		}
	}
	rewardStr := ""
	// 激励格式验证
	if *reward.Enable == 1 {
		if *reward.CallBack == 1 && reward.CallBackUrl == nil {
			c.JsonError(w, my_context.ErrBadFormat, nil)
			return
		}
		rewardStr = string(*app.Reward)
	} else {
		rewardStr = models.GetRewardJson()
	}
	slotJson, _ := json.Marshal(app.SlotMap)
	categoryLimitJson, _ := json.Marshal(category_limit)
	var appId int64
	if err := models.AddApp(user["id"], &appId, rewardStr, string(categoryLimitJson), string(slotJson), app); err != nil {
		c.JsonError(w, my_context.ErrBadExec, err)
		return
	}

	c.JsonBase(w, nil)
	c.LogAction(user["id"], user["id"], appId, my_context.ActionModule.App, "创建", r)

}

// APP修改
func (c *Handler) AppEdit(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	appId := chi.URLParam(r, "appId")

	app, oldApp, reward, category_limit := &models.AppEdit{}, &models.App{}, &models.Reward{}, &models.CategoryLimit{}
	if err := c.Bind(r, app); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*app.Reward), reward); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	rewardStr := ""
	// 激励格式验证
	if *reward.Enable == 1 {
		if *reward.CallBack == 1 && reward.CallBackUrl == nil {
			c.JsonError(w, my_context.ErrBadFormat, nil)
			return
		}
		rewardStr = string(*app.Reward)
	} else {
		rewardStr = models.GetRewardJson()
	}

	if user["deal_type"].(string) == "bidding" {
		for k, _ := range app.SlotMap {
			app.SlotMap[k] = 1
		}
	}
	// 屏蔽类型格式验证
	// 广告app 分类限制格式验证
	if err := c.BindJson(string(*app.CategoryLimit), category_limit); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if *category_limit.Enable == 0 {
		category_limit.List = make([]string, 0)
	}
	slotJson, _ := json.Marshal(app.SlotMap)
	categoryLimitJson, _ := json.Marshal(category_limit)
	// 查询当前 app 对象
	if err := models.GetApp(user["id"], appId, oldApp); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if err := models.UpdateApp(user["id"], appId, rewardStr, string(categoryLimitJson), string(slotJson), app); err != nil {
		c.JsonError(w, my_context.ErrBadExec, err)
		return
	}
	c.JsonBase(w, nil)
	c.AppCache(appId, "APP修改缓存更新")
	c.Context.LogAction(user["id"], user["id"], appId, my_context.ActionModule.App, "修改", r)

}
