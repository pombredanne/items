package models

import (
	"corm"
	"time"
	"github.com/satori/go.uuid"
	"encoding/json"
	"golang.org/x/net/context"
	"zonstfe_api/data_sync/proto"
	"fmt"
	"adv-admin-api/config"
	"api-libs/my-field"
	"adv-admin-api/client"
	"api-libs/rsp"
)

type CampaignReq struct {
	CampaignId *int    `form:"campaign_id" db:"id" query:"eq"`
	UserId     *int    `form:"user_id" db:"user_id" query:"eq"`
	Status     *int    `form:"status" db:"status" query:"eq"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
	Order      *string `form:"order" db:"order" query:"order"`
	OrderDesc  *string `form:"orderdesc" db:"orderdesc" query:"orderdesc"`
	Sdate      *string `form:"sdate" db:"DATE(create_date)" query:"gte"`
	Edate      *string `form:"edate" db:"DATE(create_date)" query:"lte"`
}

type CampaignAd struct {
	Id           *int    `json:"id" db:"id"`
	Name         *string `json:"name" db:"name" validate:"required,gt=0"`
	AdType       *string `json:"ad_type" db:"ad_type"`
	CampaignName *string `json:"campaign_name,omitempty" db:"campaign_name"`
	UserEmail    *string `json:"user_email,omitempty" db:"user_email"`
	//AdSize      *string          `json:"ad_size" db:"ad_size"`
	Width       *int     `json:"width"   db:"width"`
	Height      *int     `json:"height"  db:"height"`
	Title       *string  `json:"title"   db:"title" validate:"required,gt=0,lte=20"`
	BiddingMax  *float32 `json:"bidding_max" db:"bidding_max" validate:"required,gt=0,lte=100"`
	BiddingMin  *float32 `json:"bidding_min" db:"bidding_min" validate:"required,gt=0,lte=100"`
	BiddingType *string  `json:"bidding_type" db:"bidding_type" validate:"required,gt=0"`
	//Duration    *float64         `json:"duration" db:"duration"`
	CampaignId   *int             `json:"campaign_id" db:"campaign_id"`
	CreativeId   *int             `json:"creative_id" db:"creative_id"`
	UserId       *int             `json:"user_id" db:"user_id"`
	Creative     *json.RawMessage `json:"creative" db:"creative" validate:"required,gt=0"`
	Status       *int             `json:"status" db:"status"`
	SwitchStatus *int             `json:"switch_status" db:"switch_status"`
}
type CampaignAds []CampaignAd

type Campaign struct {
	Id          *int    `json:"campaign_id" db:"id"`
	Name        *string `json:"name" db:"name" validate:"required,gt=0,lte=40"`
	UserId      *int    `json:"user_id" db:"user_id" validate:"required,gt=0"`
	BundleId    *string `json:"bundle_id" db:"bundle_id" validate:"required"`
	Category    *string `json:"category,omitempty" db:"category" validate:"required,gt=0"`
	SubCategory *string `json:"sub_category,omitempty" db:"sub_category" validate:"required,gt=0"`
	//AppPlatform *string `json:"app_platform,omitempty" db:"app_platform" validate:"required,gt=0"`
	//Budget      *float32 `json:"budget" db:"budget" validate:"required,len>=100"`
	BudgetDay *int             `json:"budget_day" db:"budget_day" validate:"required,gte=100"`
	Freq      *json.RawMessage `json:"freq,omitempty" db:"freq" validate:"required"`
	Targeting *json.RawMessage `json:"targeting,omitempty" validate:"required,gt=0"`
	Url       *json.RawMessage `json:"url,omitempty" validate:"required,gt=0"`
	Speed     *int             `json:"speed" db:"speed,omitempty" validate:"required"`
	//StartDate   *utils.JSONTime `json:"start_date" db:"start_date" validate:"required"`
	//EndDate     *utils.JSONTime  `json:"end_date" db:"end_date" validate:"required"`
	CreateDate *my_field.JSONTime `json:"create_date" db:"create_date"`
	Status     *int            `json:"status" db:"status"`
	//Imp        *int            `json:"imp" db:"imp"`
	//Clk        *int            `json:"clk" db:"clk"`
	//Uclk       *int            `json:"uclk" db:"uclk"`
	//Cost *float64 `json:"cost" db:"cost"`
}

type Campaigns [] Campaign

type TargetingParser struct {
	Open *bool     `json:"open" validate:"required"`
	List *[]string `json:"list" validate:"required"`
}
type Targeting struct {
	Vendor      *TargetingParser `json:"vendor" validate:"required,gt=0"`       //{"open":false,list:["zonst"]}
	GeoCode     *TargetingParser `json:"geo_code" validate:"required,gt=0"`     //{"open":false,list:{"province_code":[]}}
	AppCategory *TargetingParser `json:"app_category" validate:"required,gt=0"` //{"open":false,list:{"monday":[]}}
	DayParting  *TargetingParser `json:"day_parting" validate:"required,gt=0"`  //{"open":false,list:{"周一":[]}}
	DeviceType  *TargetingParser `json:"device_type" validate:"required,gt=0"`  //{"open":false,list:[]}
	//DeviceBrand *TargetingParser `json:"device_brand" validate:"required,gt=0"` //{"open":false,list:[]}
	Os      *TargetingParser `json:"os" validate:"required,gt=0"`      //{"open":false,list:[]}
	Carrier *TargetingParser `json:"carrier" validate:"required,gt=0"` //{"open":false,list:[]}
	NetWork *TargetingParser `json:"network" validate:"required,gt=0"` //{"open":false,list:[]}
	Segment *TargetingParser `json:"segment" validate:"required,gt=0"`
}
type CampaignFreq struct {
	Open *bool   `json:"open" db:"freq->'open'"  validate:"required"`
	Type *string `json:"type" db:"freq->'type'" validate:"required,gt=0"`
	Num  *int    `json:"num" db:"freq->'num'"   validate:"required,gte=0"`
}

type Url struct {
	//TrackingImpUrl *string `json:"tracking_imp_url" validate:"required,regexp=link"`
	//TrackingClkUrl *string `json:"tracking_clk_url" validate:"required,regexp=link"`
	JumpUrl *string `json:"jump_url" validate:"required,gt=0,regexp=link"`
	//DeepLinkUrl    *string `json:"deep_link_url" validate:"required,regexp=link"`
}

// 活动列表
func GetCampaigns(total *int64, req *CampaignReq, campaigns *Campaigns) error {
	//orm := corm.Select(`select t1.*,COALESCE(t2.clk,0) clk,COALESCE(t2.imp,0) imp,COALESCE(t2.uclk,0) uclk,
	// COALESCE(t2.mycost,0) as
	//cost from (select id,name,user_id,bundle_id,speed,budget_day,to_date(create_date::text, 'yyyy-mm-dd')
	//create_date,status
	//from campaign_campaign {{t1_where}}) as t1
	//LEFT  JOIN (select COALESCE(SUM(imp),0) imp,COALESCE(SUM(uclk),0) uclk,
	//COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) mycost,campaign_id
	//FROM report_base {{t2_where}} GROUP BY campaign_id) as t2 ON t1.id=t2.campaign_id`).Req(req).EqualWhere("t2_where",
	//[][]interface{}{
	//	{"campaign_id", "=", req.CampaignId},
	//	{"sdate", ">=", req.Sdate}, {"edate", "<=", req.Edate},
	//	{"user_id","=",req.UserId},
	//}).EqualWhere("t1_where",
	//[][]interface{}{
	//	{"campaign_id", "=", req.CampaignId},
	//	{"user_id","=",req.UserId},
	//	{"status","=",req.Status},
	//}).Paginate(req.Page, req.PageSize)
	orm := corm.Select(`select * from campaign_campaign  {{sql_where}} order by status DESC,create_date desc`).Req(req).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(campaigns); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 基本报表汇总
func GetReportCampaignSum(req *CampaignReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(uclk),0) uclk,
	COALESCE(SUM(clk),0) clk,COALESCE(SUM(cost),0) "cost"
	from report_base {{t2_where}}`).Req(req).EqualWhere("t2_where", [][]interface{}{
		{"campaign_id", "=", req.CampaignId},
		{"sdate", ">=", req.Sdate}, {"edate", "<=", req.Edate},
		{"user_id", "=", req.UserId},
	}).Load(reportSum); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 广告列表
func GetAds(total *int64, req *AdReq, ads *CampaignAds) error {
	orm := corm.Select(`select *,(select name from campaign_campaign where id=t1.campaign_id) as
		campaign_name,(select email from account_account where user_id=t1.user_id) as
		user_email from campaign_ad as t1 {{sql_where}} order by create_date desc`).Req(req)
	if err := orm.Paginate(req.Page, req.PageSize).Loads(ads); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil

}

// 单个广告活动
func GetCampaign(campaignId interface{}, campaign *Campaign) error {
	if err := corm.Db.Get(campaign, `select * from 
		campaign_campaign where id=$1`, campaignId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改广告活动
func UpdateCampaign(userId, campaignId interface{}, campaign *Campaign) error {
	if _, err := corm.Db.Exec(`update campaign_campaign set name=$1,
	bundle_id=$2,category=$3,sub_category=$4,
	budget_day=$5,freq=$6,speed=$7,targeting=$8,url=$9  where id=$10 and user_id=$11`, *campaign.Name,
		*campaign.BundleId, *campaign.Category,
		*campaign.SubCategory, *campaign.BudgetDay, string(*campaign.Freq),
		*campaign.Speed, string(*campaign.Targeting), string(*campaign.Url), campaignId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改广告活动状态
func UpdateCampaignStatus(userId, campaignId interface{}, status int) error {
	if _, err := corm.Db.Exec(`update campaign_campaign set 
	status=$1 where id=$2 and user_id=$3`, status, campaignId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

func GetCampaignAd(adId interface{}, campaignAd *CampaignAd) error {
	if err := corm.Db.Get(campaignAd, `select * from campaign_ad 
	where id=$1`, adId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 添加广告活动
func AddCampaign(userId interface{}, campaignId *int64, campaign *Campaign) error {
	if err := corm.Db.QueryRow(`insert into campaign_campaign(name,user_id,bundle_id,
		category,sub_category, budget_day, freq, speed, targeting, url) values($1,$2,$3,$4,$5,$6,$7,$8,
		$9,$10) RETURNING id`, *campaign.Name, userId,
		*campaign.BundleId, *campaign.Category,
		*campaign.SubCategory, *campaign.BudgetDay,
		string(*campaign.Freq), *campaign.Speed, string(*campaign.Targeting),
		string(*campaign.Url)).Scan(campaignId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

func CampaignCache(campaignId interface{}, name string) {
	if config.Conf.EnvModel == "production" {
		start_time := time.Now().Unix()
		event_id := uuid.NewV4().String()
		if _, err := corm.Db.Exec(`insert into log_event(name,event_id,event_obj,start_time) values($1,$2,$3,$4)`,
			name, event_id, campaignId, start_time); err != nil {
			rsp.HandlerError(err)
		}
		client := proto.NewDataSyncClient(client.DataSync)
		_, err := client.CampaignCache(context.Background(), &proto.CampaignCacheRequest{
			CampaignId: fmt.Sprintf("%v", campaignId),
			EventId:    event_id,
		})

		if err != nil {
			rsp.HandlerError(err)
		}
	}

}
