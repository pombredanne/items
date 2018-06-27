package models

import (
	"corm"
	"api-libs/my-field"
	"api-libs/rsp"
)

type ReportSum struct {
	Win  *int             `json:"win,omitempty" db:"win"`
	Imp  *int             `json:"imp,omitempty" db:"imp"`
	Uclk *int             `json:"uclk,omitempty" db:"uclk"`
	Clk  *int             `json:"clk,omitempty" db:"clk"`
	Cost *float64 `json:"cost,omitempty" db:"cost"`
}
type ReportBaseReq struct {
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	AdId       *string `form:"ad_id" db:"ad_id" query:"eq"`
	UserId     *string `form:"user_id" db:"user_id" query:"eq"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
}

type ReportBaseHourReq struct {
	Sdate      *string `form:"sdate" db:"report_date + interval '1 hour' * hour" query:"gte"`
	Edate      *string `form:"edate" db:"report_date + interval '1 hour' * hour" query:"lte"`
	CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
}

type ReportGeoReq struct {
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	UserId     *string `form:"user_id" db:"user_id" query:"eq"`
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
	UserId     *string `form:"user_id" db:"user_id" query:"eq"`
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
	ReportDate   *my_field.JSONTime `json:"report_date,omitempty" db:"report_date"`
	CampaignName *string         `json:"campaign_name" db:"campaign_name"`
	UserEmail    *string         `json:"user_email" db:"user_email"`
	AdName       *string         `json:"ad_name" db:"ad_name"`
	VendorId     *int            `json:"vendor_id,omitempty" db:"vendor_id"`
	Hour         *int            `json:"hour,omitempty" db:"hour"`
	HourDate     *string         `json:"hour_date,omitempty" db:"hour_date"`
	//Win  *int `json:"win" db:"win"`
	Uclk *int             `json:"uclk" db:"uclk"`
	Imp  *int             `json:"imp" db:"imp"`
	Clk  *int             `json:"clk" db:"clk"`
	Cost *float64 `json:"cost" db:"cost"`
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
	ReportDate   *my_field.JSONTime `json:"report_date,omitempty" db:"report_date"`
	CampaignName *string         `json:"campaign_name" db:"campaign_name"`
	AdName       *string         `json:"ad_name" db:"ad_name"`
	UserEmail    *string         `json:"user_email" db:"user_email"`
	//Win    *int `json:"win" db:"win"`
	Imp  *int             `json:"imp" db:"imp"`
	Uclk *int             `json:"uclk" db:"uclk"`
	Clk  *int             `json:"clk" db:"clk"`
	Cost *float64 `json:"cost" db:"cost"`
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
	ReportDate   *my_field.JSONTime `json:"report_date,omitempty" db:"report_date"`
	CampaignName *string         `json:"campaign_name" db:"campaign_name"`
	AdName       *string         `json:"ad_name" db:"ad_name"`
	UserEmail    *string         `json:"user_email" db:"user_email"`
	Uclk         *int            `json:"uclk" db:"uclk"`
	//Win        *int `json:"win" db:"win"`
	Imp  *int             `json:"imp" db:"imp"`
	Clk  *int             `json:"clk" db:"clk"`
	Cost *float64 `json:"cost" db:"cost"`
}
type ReportApps []ReportApp

// 基本报表
func GetReportBases(total *int64, needPage bool, req *ReportBaseReq, bases *ReportBases) error {
	orm := corm.Select(`select * from (select COALESCE(imp,0) imp,COALESCE(uclk,0) uclk,
	COALESCE(clk,0) clk,COALESCE(cost,0) "cost",ad_id,campaign_id,report_date,COALESCE((select name
	from campaign_campaign where id=t1.campaign_id),'') campaign_name,COALESCE((select name from campaign_ad where id=t1.ad_id),'')
    ad_name,COALESCE((select email from account_account
	where user_id=t1.user_id),'') user_email,user_id from report_base as t1 ORDER BY report_date desc,campaign_id,ad_id) as t100  {{sql_where}}`).Req(req)
	if needPage {
		orm.Paginate(req.Page, req.PageSize)
	}
	if err := orm.Loads(bases); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 基本报表汇总
func GetReportBaseSum(req *ReportBaseReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(uclk),0) uclk,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_base {{sql_where}}`).Req(req).Load(reportSum); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 地域报表
func GetReportGeos(total *int64, needPage bool, req *ReportGeoReq, geos *ReportGeos) error {
	orm := corm.Select(`select COALESCE(imp,0) imp,COALESCE(uclk,0) uclk,COALESCE(clk,0) clk,
	COALESCE(cost,0) "cost",province_code,city_code,campaign_id,report_date,ad_id,COALESCE((select name
	from campaign_campaign where id=t1.campaign_id),'') campaign_name,COALESCE((select name from campaign_ad where id=t1.ad_id),'')
    ad_name,COALESCE((select email from account_account 
	where user_id=t1.user_id),'') user_email,user_id from report_geo as t1 {{sql_where}} ORDER BY 
	province_code,city_code,report_date desc,campaign_id,ad_id`).Req(req)
	if needPage {
		orm.Paginate(req.Page, req.PageSize)
	}

	if err := orm.Loads(geos); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 地域报表汇总
func GetReportGeoSum(req *ReportGeoReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(uclk),0) uclk,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_geo {{sql_where}}`).Req(req).Load(reportSum); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// App报表
func GetReportApps(total *int64,needPage bool, req *ReportAppReq, apps *ReportApps) error {
	orm := corm.Select(`select COALESCE(imp,0) imp,COALESCE(uclk,0) uclk,COALESCE(clk,0) clk,
	COALESCE(cost,0) "cost",os,bundle_id,report_date,ad_id,campaign_id,COALESCE((select name
	from campaign_campaign where id=t1.campaign_id),'') campaign_name,COALESCE((select name from campaign_ad where id=t1.ad_id),'')
    ad_name,COALESCE((select email from account_account 
	where user_id=t1.user_id),'') user_email,user_id from report_app as t1 {{sql_where}}
	order by os,bundle_id,report_date desc,ad_id,campaign_id`).Req(req)
	if needPage{
		orm.Paginate(req.Page, req.PageSize)
	}
	if err := orm.Loads(apps); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// App报表汇总
func GetReportAppSum(req *ReportAppReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(uclk),0) uclk,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_app {{sql_where}}`).Req(req).Load(reportSum); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 基本小时报表(首页)
func GetReportBaseHour(req *ReportBaseHourReq, bases *ReportBases) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost",hour,report_date,
	substring(report_date ||' '|| interval '1 hour' * hour from 0 for 17)  as hour_date from
	report_base_hour {{sql_where}} GROUP BY hour_date,hour,
	report_date ORDER BY hour_date`).Req(req).Loads(bases); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 基本小时报表汇总(首页)
func GetReportBaseHourSum(req *ReportBaseHourReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_base_hour {{sql_where}}`).Req(req).Load(reportSum); err != nil {
		return err
	}
	return nil
}



