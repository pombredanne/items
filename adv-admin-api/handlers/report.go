package handlers

import (
	"api-libs/my-date"
	"github.com/gin-gonic/gin"
	"api-libs/rsp"
	"api-libs/option"
	"adv-admin-api/models"
	"fmt"
)

// 基本报表
func ReportBase(c *gin.Context) {
	var req models.ReportBaseReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil && req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
	}
	report_bases, report_sum := &models.ReportBases{}, &models.ReportSum{}
	var total int64
	if err := models.GetReportBases(&total, true, &req, report_bases); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportBaseSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.SumPage(len(*report_bases), total, report_bases, report_sum))
}
func ReportBaseExport(c *gin.Context) {
	var req models.ReportBaseReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil && req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
	}
	report_bases, report_sum := &models.ReportBases{}, &models.ReportSum{}
	var total int64
	if err := models.GetReportBases(&total, false, &req, report_bases); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportBaseSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	fileName := fmt.Sprintf("基本报表_%s_%s.csv", *req.Sdate, *req.Edate)
	datas := [][]string{
		{"日期", "用户", "广告系列", "广告", "点击", "展示", "点击率", "去重点击", "花费", "CPC", "CPM", "UCPC"},
	}
	datas = append(datas, []string{
		"汇总", "", "", "", fmt.Sprintf("%v", *report_sum.Clk), fmt.Sprintf("%v", *report_sum.Imp),
		rsp.Division(float64(*report_sum.Clk), float64(*report_sum.Imp)),
		fmt.Sprintf("%v", *report_sum.Uclk), fmt.Sprintf("%v", *report_sum.Cost),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Clk), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Imp), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Uclk), 1000),
	})
	for _, item := range *report_bases {
		datas = append(datas, []string{
			fmt.Sprintf("%v", *item.ReportDate), *item.UserEmail, *item.CampaignName, *item.AdName, fmt.Sprintf("%v", *item.Clk),
			fmt.Sprintf("%v", *item.Imp), rsp.Division(float64(*item.Clk), float64(*item.Imp)),
			fmt.Sprintf("%v", *item.Uclk), fmt.Sprintf("%v", *item.Cost),
			rsp.Division(float64(*item.Cost), float64(*item.Clk), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Imp), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Uclk), 1000),
		})
	}
	c.String(rsp.CSV(c.Writer, fileName, datas))

}

// 首页基本报表
func ReportBaseHour(c *gin.Context) {
	var req models.ReportBaseHourReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDate() + " 00:00")
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDate() + " 23:59")
	}
	report_bases, report_sum := &models.ReportBases{}, &models.ReportSum{}
	if err := models.GetReportBaseHour(&req, report_bases); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportBaseHourSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Sum(report_bases, report_sum))
}

func ReportBaseHourExport(c *gin.Context) {
	var req models.ReportBaseHourReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDate() + " 00:00")
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDate() + " 23:59")
	}
	report_bases, report_sum := &models.ReportBases{}, &models.ReportSum{}
	if err := models.GetReportBaseHour(&req, report_bases); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportBaseHourSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	fileName := fmt.Sprintf("基本小时报表_%s_%s.csv", *req.Sdate, *req.Edate)
	datas := [][]string{
		{"日期", "时", "点击", "展示", "点击率", "花费", "CPC", "CPM"},
	}
	datas = append(datas, []string{
		"汇总", "", fmt.Sprintf("%v", *report_sum.Clk), fmt.Sprintf("%v", *report_sum.Imp),
		rsp.Division(float64(*report_sum.Clk), float64(*report_sum.Imp)), fmt.Sprintf("%v", *report_sum.Cost),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Clk), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Imp), 1000),
	})
	for _, item := range *report_bases {
		datas = append(datas, []string{
			fmt.Sprintf("%v", *item.ReportDate),
			fmt.Sprintf("%v", *item.Hour), fmt.Sprintf("%v", *item.Clk),
			fmt.Sprintf("%v", *item.Imp), rsp.Division(float64(*item.Clk), float64(*item.Imp)),
			fmt.Sprintf("%v", *item.Uclk), fmt.Sprintf("%v", *item.Cost),
			rsp.Division(float64(*item.Cost), float64(*item.Clk), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Imp), 1000),
		})
	}
	c.String(rsp.CSV(c.Writer, fileName, datas))
}

// 地域报表
func ReportGeo(c *gin.Context) {
	var req models.ReportGeoReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
	}
	report_geos, report_sum := &models.ReportGeos{}, &models.ReportSum{}

	var total int64
	if err := models.GetReportGeos(&total, true, &req, report_geos); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportGeoSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.SumPage(len(*report_geos), total, report_geos, report_sum))
}

func ReportGeoExport(c *gin.Context) {
	var req models.ReportGeoReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
	}
	report_geos, report_sum := &models.ReportGeos{}, &models.ReportSum{}

	var total int64
	if err := models.GetReportGeos(&total, false, &req, report_geos); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportGeoSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	fileName := fmt.Sprintf("地域报表_%s_%s.csv", *req.Sdate, *req.Edate)
	datas := [][]string{
		{"日期", "用户", "广告系列", "广告", "省", "市", "点击", "展示", "点击率", "去重点击", "花费", "CPC", "CPM", "UCPC"},
	}
	datas = append(datas, []string{
		"汇总", "", "", "", "", "", fmt.Sprintf("%v", *report_sum.Clk), fmt.Sprintf("%v", *report_sum.Imp),
		rsp.Division(float64(*report_sum.Clk), float64(*report_sum.Imp)),
		fmt.Sprintf("%v", *report_sum.Uclk), fmt.Sprintf("%v", *report_sum.Cost),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Clk), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Imp), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Uclk), 1000),
	})

	for _, item := range *report_geos {
		datas = append(datas, []string{
			fmt.Sprintf("%v", *item.ReportDate), *item.UserEmail, *item.CampaignName, *item.AdName,
			fmt.Sprintf("%v", option.GeoCodeName[*item.ProvinceCode]),
			fmt.Sprintf("%v", option.GeoCodeName[*item.CityCode]), fmt.Sprintf("%v", *item.Clk),
			fmt.Sprintf("%v", *item.Imp), rsp.Division(float64(*item.Clk), float64(*item.Imp)),
			fmt.Sprintf("%v", *item.Uclk), fmt.Sprintf("%v", *item.Cost),
			rsp.Division(float64(*item.Cost), float64(*item.Clk), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Imp), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Uclk), 1000),
		})
	}
	c.String(rsp.CSV(c.Writer, fileName, datas))
}

// APP 报表
func ReportApp(c *gin.Context) {
	var req models.ReportAppReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
	}
	report_apps, report_sum := &models.ReportApps{}, &models.ReportSum{}
	var total int64
	if err := models.GetReportApps(&total, true, &req, report_apps); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportAppSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.SumPage(len(*report_apps), total, report_apps, report_sum))
}

func ReportAppExport(c *gin.Context) {
	var req models.ReportAppReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if req.Sdate == nil || req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
		req.Edate = func(i string) *string { return &i }(my_date.CurrentDateAdd(-1))
	}
	report_apps, report_sum := &models.ReportApps{}, &models.ReportSum{}
	var total int64
	if err := models.GetReportApps(&total, false, &req, report_apps); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	if err := models.GetReportAppSum(&req, report_sum); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	fileName := fmt.Sprintf("媒体报表_%s_%s.csv", *req.Sdate, *req.Edate)
	datas := [][]string{
		{"日期", "用户", "广告系列", "广告", "系统", "媒体ID", "点击", "展示", "点击率", "去重点击", "花费", "CPC", "CPM", "UCPC"},
	}
	datas = append(datas, []string{
		"汇总", "", "", "", "", "", fmt.Sprintf("%v", *report_sum.Clk), fmt.Sprintf("%v", *report_sum.Imp),
		rsp.Division(float64(*report_sum.Clk), float64(*report_sum.Imp)),
		fmt.Sprintf("%v", *report_sum.Uclk), fmt.Sprintf("%v", *report_sum.Cost),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Clk), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Imp), 1000),
		rsp.Division(float64(*report_sum.Cost), float64(*report_sum.Uclk), 1000),
	})

	for _, item := range *report_apps {
		datas = append(datas, []string{
			fmt.Sprintf("%v", *item.ReportDate), *item.UserEmail, *item.CampaignName, *item.AdName,
			fmt.Sprintf("%v", *item.Os),
			fmt.Sprintf("%v", *item.BundleId), fmt.Sprintf("%v", *item.Clk),
			fmt.Sprintf("%v", *item.Imp), rsp.Division(float64(*item.Clk), float64(*item.Imp)),
			fmt.Sprintf("%v", *item.Uclk), fmt.Sprintf("%v", *item.Cost),
			rsp.Division(float64(*item.Cost), float64(*item.Clk), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Imp), 1000),
			rsp.Division(float64(*item.Cost), float64(*item.Uclk), 1000),
		})
	}
	c.String(rsp.CSV(c.Writer, fileName, datas))
}
