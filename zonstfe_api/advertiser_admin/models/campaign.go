package models

import (
	"encoding/json"
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
	"github.com/jmoiron/sqlx"
)

type CampaignReq struct {
	CampaignId *int    `form:"campaign_id" db:"id" query:"eq"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
}

type AdReq struct {
	AdSize       *string `form:"ad_size" db:"ad_size" query:"eq"`
	AdType       *string `form:"ad_type" db:"ad_type" query:"eq"`
	Status       *int    `form:"status" db:"status" query:"eq"`
	Sdate        *string `form:"sdate" db:"create_date" query:"gte"`
	Edate        *string `form:"edate" db:"create_date" query:"lte"`
	CampaignNmae *string `form:"campaign_name" db:"campaign_name" query:"like"`
	UserEmail    *string `form:"user_email" db:"user_email" query:"like"`
	Name         *string `form:"name" db:"name" query:"like"`
	Page         *uint   `form:"page"`
	PageSize     *uint   `form:"page_size"`
}

type CampaignAd struct {
	Id           *int             `json:"ad_id" db:"id"`
	CampaignName *string          `json:"campaign_name,omitempty" db:"campaign_name"`
	UserEmail    *string          `json:"user_email,omitempty" db:"user_email"`
	Name         *string          `json:"name" db:"name" validate:"required,gt=0"`
	AdType       *string          `json:"ad_type" db:"ad_type"`
	AdSize       *string          `json:"ad_size" db:"ad_size"`
	Width        *string          `json:"width" db:"width"`
	Height       *string          `json:"height" db:"height"`
	Duration     *float64         `json:"duration" db:"duration"`
	CampaignId   *int             `json:"campaign_id" db:"campaign_id"`
	CreativeId   *int             `json:"creative_id" db:"creative_id"`
	UserId       *int             `json:"user_id" db:"user_id"`
	Creative     *json.RawMessage `json:"creative" db:"creative" validate:"required,gt=0"`
	Status       *int             `json:"status" db:"status"`
}
type CampaignAds []CampaignAd

type Campaign struct {
	Id         int             `json:"campaign_id" db:"id"`
	Name       *string         `json:"campaign_name" db:"name"`
	BudgetDay  *int            `json:"budget_day" db:"budget_day"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
	Status     *int            `json:"status" db:"status"`
	Imp        *int            `json:"imp" db:"imp"`
	Clk        *int            `json:"clk" db:"clk"`
	Eimp       *int            `json:"eimp" db:"eimp"`
	Cost       *float64        `json:"cost" db:"cost"`
}

type Campaigns [] Campaign

type ReportSum struct {
	Win  *int             `json:"win,omitempty" db:"win"`
	Imp  *int             `json:"imp,omitempty" db:"imp"`
	Eimp *int             `json:"eimp,omitempty" db:"eimp"`
	Clk  *int             `json:"clk,omitempty" db:"clk"`
	Cost *utils.JSONFloat `json:"cost,omitempty" db:"cost"`
}

// 活动列表
func GetCampaigns(total *int64, req *CampaignReq, campaigns *Campaigns) error {
	orm := corm.Select(`select t1.*,COALESCE(t2.clk,0) clk,COALESCE(t2.imp,0) imp,
     COALESCE(t2.eimp,0) eimp,COALESCE(t2.mycost,0) as 
    cost from (select id,name,budget_day,to_date(create_date::text, 'yyyy-mm-dd') 
	create_date,status 
    from campaign_campaign {{t1_where}}) as t1 
    LEFT  JOIN (select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) mycost,campaign_id 
    FROM report_base {{t2_where}} GROUP BY campaign_id) as t2 ON t1.id=t2.campaign_id`).Req(req).EqualWhere("t2_where",
		[][]interface{}{
			{"campaign_id", "=", req.CampaignId},
			{"sdate", ">=", req.Sdate}, {"edate", "<=", req.Edate},
		}).EqualWhere("t1_where", "id", "=", req.CampaignId).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(campaigns); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 基本报表汇总
func GetReportCampaignSum(req *CampaignReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(eimp),0) eimp,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_base {{t2_where}}`).Req(req).EqualWhere("t2_where", [][]interface{}{
		{"campaign_id", "=", req.CampaignId},
		{"sdate", ">=", req.Sdate}, {"edate", "<=", req.Edate},
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}

// 广告列表
func GetAds(total *int64, req *AdReq, ads *CampaignAds) error {
	orm := corm.Select(`select *,(select name from campaign_campaign where id=t1.campaign_id) as
		campaign_name,(select email from account_account where user_id=t1.user_id) as
		user_email from campaign_ad as t1 {{sql_where}}`).Req(req)
	if err := orm.Paginate(req.Page, req.PageSize).Loads(ads); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil

}

func GetCampaignAd(adId interface{}, campaignAd *CampaignAd) error {
	if err := corm.Db.Get(campaignAd, `select * from campaign_ad 
	where id=$1`, adId); err != nil {
		return err
	}
	return nil
}

// 广告批量审核
func ReviewBatchAd(adList []int, status int) error {
	query, args, err := sqlx.In("update campaign_ad set status=$1 where id in ($2);", status, adList)
	if err != nil {
		panic(err)
	}
	query = corm.Db.Rebind(query)
	if _, err := corm.Db.Exec(query, args...); err != nil {
		return err
	}
	return nil

}
