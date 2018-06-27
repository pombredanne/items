package handlers

import (
	"zonstfe_api/common/my_context"
	"net/http"
	"zonstfe_api/advertiser/models"
	"zonstfe_api/common/utils/mydate"
)

// 基本报表
func (c *Handler) ReportBase(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ReportBaseReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Sdate == nil && req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	}
	report_bases, report_sum := &models.ReportBases{}, &models.ReportSum{}

	var total int64
	if err := models.GetReportBases(user["id"], &total, &req, report_bases); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if err := models.GetReportBaseSum(user["id"], &req, report_sum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonSumPage(w, len(*report_bases), total, report_bases, report_sum)
}

// 首页基本报表
func (c *Handler) ReportBaseHour(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ReportBaseHourReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate() + " 00:00")
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate() + " 23:59")
	}
	report_bases, report_sum := &models.ReportBases{}, &models.ReportSum{}
	if err := models.GetReportBaseHour(user["id"], &req, report_bases); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if err := models.GetReportBaseHourSum(user["id"], &req, report_sum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonSum(w, report_bases, report_sum)
}

// 地域报表
//func (c *Handler) ReportGeoCountry(w http.ResponseWriter, r *http.Request) {
//	current_user := c.GetUser(r)
//	var req struct {
//		Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
//		Edate      *string `form:"edate" db:"report_date" query:"lte"`
//		CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
//		VendorId   *int    `form:"vendor_id" db:"vendor_id" query:"eq"`
//	}
//	if err := c.Bind(r, &req); err != nil {
//		c.JsonError(w, my_context.ErrBadValid, err)
//		return
//	}
//	if req.Sdate == nil || req.Edate == nil {
//		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
//		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
//	}
//	report_geos := &models.ReportGeos{}
//
//	sql := `select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,COALESCE(SUM(clk),0) clk,
//	COALESCE(SUM(cost),0) "cost",country_code,report_date
//		from report_geo {{sql_where}} group by report_date,
//	country_code order by report_date,country_code`
//	if err := corm.Select(sql).Req(&req).Where(map[string]interface{}{
//		"user_id": current_user["id"],
//	}).Loads(report_geos); err != nil {
//		c.JsonError(w, my_context.ErrDefault, err)
//		return
//	}
//	report_sum := &models.ReportSum{}
//	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
//	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
//	from report_geo {{sql_where}}`).Req(&req).Where(map[string]interface{}{
//		"user_id": current_user["id"],
//	}).Load(report_sum); err != nil {
//		c.JsonError(w, my_context.ErrDefault, err)
//		return
//	}
//	c.JsonSum(w, report_geos, report_sum)
//}
//
//// 省份报表
//func (c *Handler) ReportGeoProvince(w http.ResponseWriter, r *http.Request) {
//	current_user := c.GetUser(r)
//	var req struct {
//		Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
//		Edate      *string `form:"edate" db:"report_date" query:"lte"`
//		CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
//		VendorId   *int    `form:"vendor_id" db:"vendor_id" query:"eq"`
//	}
//	if err := c.Bind(r, &req); err != nil {
//		c.JsonError(w, my_context.ErrBadValid, err)
//		return
//	}
//	if req.Sdate == nil || req.Edate == nil {
//		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
//		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
//	}
//	report_geos := &models.ReportGeos{}
//
//	sql := `select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,COALESCE(SUM(clk),0) clk,
//	COALESCE(SUM(cost),0) "cost",province_code,report_date
//		from report_geo {{sql_where}} group by report_date,
//	province_code order by report_date,province_code`
//	if err := corm.Select(sql).Req(&req).Where(map[string]interface{}{
//		"user_id": current_user["id"],
//	}).Loads(report_geos); err != nil {
//		c.JsonError(w, my_context.ErrDefault, err)
//		return
//	}
//	report_sum := &models.ReportSum{}
//	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
//	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
//	from report_geo {{sql_where}}`).Req(&req).Where(map[string]interface{}{
//		"user_id": current_user["id"],
//	}).Load(report_sum); err != nil {
//		c.JsonError(w, my_context.ErrDefault, err)
//		return
//	}
//	c.JsonSum(w, report_geos, report_sum)
//}
//
//// 城市报表
//func (c *Handler) ReportGeoCity(w http.ResponseWriter, r *http.Request) {
//	current_user := c.GetUser(r)
//	var req struct {
//		Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
//		Edate      *string `form:"edate" db:"report_date" query:"lte"`
//		CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
//		VendorId   *int    `form:"vendor_id" db:"vendor_id" query:"eq"`
//	}
//	if err := c.Bind(r, &req); err != nil {
//		c.JsonError(w, my_context.ErrBadValid, err)
//		return
//	}
//	if req.Sdate == nil || req.Edate == nil {
//		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
//		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
//	}
//	report_geos := &models.ReportGeos{}
//	sql := `select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,COALESCE(SUM(clk),0) clk,
//	COALESCE(SUM(cost),0) "cost",city_code,report_date
//		from report_geo {{sql_where}} group by report_date,
//	city_code order by report_date,city_code`
//	if err := corm.Select(sql).Req(&req).Where(map[string]interface{}{
//		"user_id": current_user["id"],
//	}).Loads(report_geos); err != nil {
//		c.JsonError(w, my_context.ErrDefault, err)
//		return
//	}
//	report_sum := &models.ReportSum{}
//	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
//	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
//	from report_geo {{sql_where}}`).Req(&req).Where(map[string]interface{}{
//		"user_id": current_user["id"],
//	}).Load(report_sum); err != nil {
//		c.JsonError(w, my_context.ErrDefault, err)
//		return
//	}
//	c.JsonSum(w, report_geos, report_sum)
//}

// 地域报表
func (c *Handler) ReportGeo(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ReportGeoReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	}
	report_geos, report_sum := &models.ReportGeos{}, &models.ReportSum{}

	var total int64
	if err := models.GetReportGeos(user["id"], &total, &req, report_geos); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if err := models.GetReportGeoSum(user["id"], &req, report_sum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonSumPage(w, len(*report_geos), total, report_geos, report_sum)
}

// APP 报表
func (c *Handler) ReportApp(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.ReportAppReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	}
	report_apps, report_sum := &models.ReportApps{}, &models.ReportSum{}
	var total int64
	if err := models.GetReportApps(user["id"], &total, &req, report_apps); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if err := models.GetReportAppSum(user["id"], &req, report_sum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonSumPage(w, len(*report_apps), total, report_apps, report_sum)
}
