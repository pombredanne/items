package models

import (
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
)

type ReportBaseReq struct {
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	AdId       *string `form:"ad_id" db:"ad_id" query:"eq"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
}

type ReportBaseHourReq struct {
	Sdate *string `form:"sdate" db:"report_date + interval '1 hour' * hour" query:"gte"`
	Edate *string `form:"edate" db:"report_date + interval '1 hour' * hour" query:"lte"`
}

type ReportGeoReq struct {
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	//Country    *string `form:"country" db:"country_code" query:"in"`
	Province *string `form:"province_code" db:"province_code" query:"eq"`
	City     *string `form:"city_code" db:"city_code" query:"eq"`
	AdId     *string `form:"ad_id" db:"ad_id" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
}

type ReportAppReq struct {
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	Os         *string `form:"os" db:"os" query:"eq"`
	CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	VendorId   *int    `form:"vendor_id" db:"vendor_id" query:"eq"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
}

type ReportBase struct {
	Id           *int            `json:"report_id,omitempty" db:"id"`
	UserId       *int            `json:"user_id,omitempty" db:"user_id"`
	CampaignId   *int            `json:"campaign_id,omitempty" db:"campaign_id"`
	AdId         *int            `json:"ad_id,omitempty" db:"ad_id"`
	ReportDate   *utils.JSONTime `json:"report_date,omitempty" db:"report_date"`
	CampaignName *string         `json:"campaign_name" db:"campaign_name"`
	AdName       *string         `json:"ad_name" db:"ad_name"`
	VendorId     *int            `json:"vendor_id,omitempty" db:"vendor_id"`
	Hour         *int            `json:"hour,omitempty" db:"hour"`
	HourDate     *string         `json:"hour_date,omitempty" db:"hour_date"`
	//Win  *int `json:"win" db:"win"`
	Eimp *int             `json:"eimp" db:"eimp"`
	Imp  *int             `json:"imp" db:"imp"`
	Clk  *int             `json:"clk" db:"clk"`
	Cost *utils.JSONFloat `json:"cost" db:"cost"`
}
type ReportBases []ReportBase

type ReportGeo struct {
	Id           *int            `json:"report_id,omitempty" db:"id"`
	UserId       *int            `json:"user_id,omitempty" db:"user_id"`
	VendorId     *int            `json:"vendor_id,omitempty" db:"vendor_id"`
	CountryCode  *string         `json:"country_code,omitempty" db:"country_code"`
	ProvinceCode *string         `json:"province_code,omitempty" db:"province_code"`
	CityCode     *string         `json:"city_code,omitempty" db:"city_code"`
	CampaignId   *int            `json:"campaign_id,omitempty" db:"campaign_id"`
	AdId         *int            `json:"ad_id,omitempty" db:"ad_id,omitempty"`
	ReportDate   *utils.JSONTime `json:"report_date,omitempty" db:"report_date"`
	CampaignName *string         `json:"campaign_name" db:"campaign_name"`
	AdName       *string         `json:"ad_name" db:"ad_name"`
	//Win    *int `json:"win" db:"win"`
	Imp  *int             `json:"imp" db:"imp"`
	Eimp *int             `json:"eimp" db:"eimp"`
	Clk  *int             `json:"clk" db:"clk"`
	Cost *utils.JSONFloat `json:"cost" db:"cost"`
}
type ReportGeos []ReportGeo

type ReportApp struct {
	Id           *int            `json:"report_id,omitempty" db:"id"`
	UserId       *int            `json:"user_id,omitempty" db:"user_id"`
	Os           *string         `json:"os,omitempty" db:"os"`
	VendorId     *int            `json:"vendor_id,omitempty" db:"vendor_id"`
	CampaignId   *int            `json:"campaign_id,omitempty" db:"campaign_id"`
	AdId         *int            `json:"ad_id,omitempty" db:"ad_id"`
	BundleId     *string         `json:"bundle_id,omitempty" db:"bundle_id"`
	ReportDate   *utils.JSONTime `json:"report_date,omitempty" db:"report_date"`
	CampaignName *string         `json:"campaign_name" db:"campaign_name"`
	AdName       *string         `json:"ad_name" db:"ad_name"`
	Eimp         *int            `json:"eimp" db:"eimp"`
	//Win        *int `json:"win" db:"win"`
	Imp  *int             `json:"imp" db:"imp"`
	Clk  *int             `json:"clk" db:"clk"`
	Cost *utils.JSONFloat `json:"cost" db:"cost"`
}
type ReportApps []ReportApp

type ReportSum struct {
	Win  *int             `json:"win,omitempty" db:"win"`
	Imp  *int             `json:"imp,omitempty" db:"imp"`
	Eimp *int             `json:"eimp,omitempty" db:"eimp"`
	Clk  *int             `json:"clk,omitempty" db:"clk"`
	Cost *utils.JSONFloat `json:"cost,omitempty" db:"cost"`
}

// 基本报表
func GetReportBases(userId interface{}, total *int64, req *ReportBaseReq, bases *ReportBases) error {
	orm := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost",ad_id,campaign_id,report_date,(select name
	from campaign_campaign where id=t1.campaign_id) campaign_name,(select name from campaign_ad where id=t1.ad_id)
    ad_name from report_base as t1 {{sql_where}} group by  
	ad_id,campaign_id,report_date`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(bases); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 基本报表汇总
func GetReportBaseSum(userId interface{}, req *ReportBaseReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_base {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"user_id": userId,
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}

// 地域报表
func GetReportGeos(userId interface{}, total *int64, req *ReportGeoReq, geos *ReportGeos) error {
	orm := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,COALESCE(SUM(clk),0) clk,
	COALESCE(SUM(cost),0) "cost",province_code,city_code,campaign_id,report_date,ad_id,(select name
	from campaign_campaign where id=t1.campaign_id) campaign_name,(select name from campaign_ad where id=t1.ad_id)
  ad_name from report_geo as t1 {{sql_where}} group by province_code,city_code,report_date,
	campaign_id,ad_id ORDER BY province_code,city_code,report_date,campaign_id,ad_id`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(geos); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 地域报表汇总
func GetReportGeoSum(userId interface{}, req *ReportGeoReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_geo {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"user_id": userId,
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}

// App报表
func GetReportApps(userId interface{}, total *int64, req *ReportAppReq, apps *ReportApps) error {
	orm := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,COALESCE(SUM(clk),0) clk,
	COALESCE(SUM(cost),0) "cost",os,bundle_id,report_date,ad_id,campaign_id,(select name
	from campaign_campaign where id=t1.campaign_id) campaign_name,(select name from campaign_ad where id=t1.ad_id)
    ad_name from report_app as t1 {{sql_where}}
	group by os,bundle_id,report_date,ad_id,campaign_id 
	order by os,bundle_id,report_date,ad_id,campaign_id`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(apps); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// App报表汇总
func GetReportAppSum(userId interface{}, req *ReportAppReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_geo {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"user_id": userId,
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}

// 基本小时报表(首页)
func GetReportBaseHour(userId interface{}, req *ReportBaseHourReq, bases *ReportBases) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost",hour,report_date,
	substring(report_date ||' '|| interval '1 hour' * hour from 0 for 17)  as hour_date from
	report_base {{sql_where}} GROUP BY hour_date,hour,
	report_date ORDER BY hour_date`).Req(req).Where(map[string]interface{}{
		"user_id": userId,
	}).Loads(bases); err != nil {
		return err
	}
	return nil
}

// 基本小时报表汇总(首页)
func GetReportBaseHourSum(userId interface{}, req *ReportBaseHourReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_base {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"user_id": userId,
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}
