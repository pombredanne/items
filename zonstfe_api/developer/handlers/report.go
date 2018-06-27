package handlers

import (
	"zonstfe_api/common/my_context"
	"zonstfe_api/common/utils/mydate"
	"zonstfe_api/developer/models"
	"net/http"
)

// 广告位报表
func (c *Handler) ReportAppSlot(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ReportAppSlotReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	}
	data := &models.ReportAppSlots{}
	var total int64
	if err := models.GetReportAppSlot(user["app_key"], &total, &req, data); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	reportSum := &models.ReportSum{}
	if err := models.GetReportAppSlotSum(user["app_key"], &req, reportSum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonSumPage(w, len(*data), total, data, reportSum)

}

// 激励报表
func (c *Handler) ReportAppReward(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ReportAppRewardReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	}
	var total int64
	data := &models.ReportAppRewards{}
	if err := models.GetReportAppReward(user["app_key"], &total, &req, data); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	reportSum := &models.ReportSum{}
	if err := models.GetReportAppRewardSum(user["app_key"], &req, reportSum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonSumPage(w, len(*data), total, data, reportSum)
}
