package models

import (
	"zonstfe_api/common/utils"
	"encoding/json"
	"zonstfe_api/corm"
	"zonstfe_api/common/utils/mydate"
)

type CampaignReq struct {
	Status   *int  `form:"status" db:"status" query:"eq"`
	Page     *uint `form:"page"`
	PageSize *uint `form:"page_size"`
}

type AdReq struct {
	AdSize       *string `form:"ad_size" db:"ad_size" query:"eq"`
	AdType       *string `form:"ad_type" db:"ad_type" query:"eq"`
	Status       *int    `form:"status" db:"status" query:"eq"`
	Sdate        *string `form:"sdate" db:"create_date" query:"gte"`
	Edate        *string `form:"edate" db:"create_date" query:"lte"`
	CampaignId   *int    `form:"campaign_id" db:"campaign_id" query="eq"`
	CampaignNmae *string `form:"campaign_name" db:"campaign_name" query:"like"`
	UserEmail    *string `form:"user_email" db:"user_email" query:"like"`
	Name         *string `form:"name" db:"name" query:"like"`
	Page         *uint   `form:"page"`
	PageSize     *uint   `form:"page_size"`
}

type SegmentReq struct {
	Page     *uint `form:"page"`
	PageSize *uint `form:"page_size"`
}

type Campaign struct {
	Id          int     `json:"campaign_id" db:"id"`
	Name        *string `json:"name" db:"name" validate:"required,gt=0"`
	BundleId    *string `json:"bundle_id" db:"bundle_id" validate:"required,gt=0,regexp=bundle"`
	Category    *string `json:"category,omitempty" db:"category" validate:"required,gt=0"`
	SubCategory *string `json:"sub_category,omitempty" db:"sub_category" validate:"required,gt=0"`
	AppPlatform *string `json:"app_platform,omitempty" db:"app_platform" validate:"required,gt=0"`
	//Budget      *float32 `json:"budget" db:"budget" validate:"required,len>=100"`
	BudgetDay   *float32         `json:"budget_day" db:"budget_day" validate:"required,gte=100"`
	BiddingMax  *float32         `json:"bidding_max,omitempty" db:"bidding_max" validate:"required"`
	BiddingMin  *float32         `json:"bidding_min,omitempty" db:"bidding_min" validate:"required"`
	BiddingType *string          `json:"bidding_type" db:"bidding_type" validate:"required"`
	Freq        *json.RawMessage `json:"freq,omitempty" db:"freq" validate:"required"`
	Targeting   *json.RawMessage `json:"targeting,omitempty" validate:"required,gt=0"`
	Url         *json.RawMessage `json:"url,omitempty" validate:"required,gt=0"`
	Speed       *int             `json:"speed" db:"speed,omitempty" validate:"required"`
	//StartDate   *utils.JSONTime `json:"start_date" db:"start_date" validate:"required"`
	//EndDate     *utils.JSONTime  `json:"end_date" db:"end_date" validate:"required"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
	Status     *int            `json:"status" db:"status"`
}

type CampaignFreq struct {
	Open *bool   `json:"open" db:"freq->'open'"  validate:"required"`
	Type *string `json:"type" db:"freq->'type'" validate:"required,gt=0"`
	Num  *int    `json:"num" db:"freq->'num'"   validate:"required,gte=0"`
}

type Campaigns []Campaign

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
	OsVersion *TargetingParser `json:"os_version" validate:"required,gt=0"` //{"open":false,list:[]}
	Carrier   *TargetingParser `json:"carrier" validate:"required,gt=0"`    //{"open":false,list:[]}
	NetWork   *TargetingParser `json:"network" validate:"required,gt=0"`    //{"open":false,list:[]}
	Segment   *TargetingParser `json:"segment" validate:"required,gt=0"`
}

type Url struct {
	TrackingImpUrl *string `json:"tracking_imp_url" validate:"required,regexp=link"`
	TrackingClkUrl *string `json:"tracking_clk_url" validate:"required,regexp=link"`
	JumpUrl        *string `json:"jump_url" validate:"required,gt=0,regexp=link"`
	DeepLinkUrl    *string `json:"deep_link_url" validate:"required,regexp=link"`
}

type Segment struct {
	Id         *int            `josn:"segment_id" db:"id"`
	Name       *string         `json:"name" db:"name" validate:"required,gt=0"`
	UserId     *int            `json:"user_id" db:"user_id"`
	Type       *int            `json:"type" validate:"required,gt=0"`
	Uv         *int            `josn:"uv" db:"uv"`
	PkgPath    *string         `json:"pkg_path" db:"pkg_path" validate:"required,gt=0,regexp=link"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
	UpdateDate *utils.JSONTime `json:"update_date" db:"update_date"`
}

type Segments []Segment

type VideoCreative struct {
	CampaignId *int     `json:"campaign_id" validate:"required,gt=0"`
	Video      *string  `json:"video" validate:"required,gt=0,regexp=link"`
	Image      *string  `json:"image" validate:"required,gt=0,regexp=link"`
	Name       *string  `json:"name" validate:"required,gt=0"`
	Width      *int     `json:"width" validate:"required,gt=0"`
	Height     *int     `json:"height" validate:"required,gt=0"`
	Duration   *float64 `json:"duration" validate:"required,gt=0"`
}
type ImageCreative struct {
	CampaignId *int    `json:"campaign_id" validate:"required,gt=0"`
	Name       *string `json:"name" validate:"required,gt=0"`
	Image      *string `json:"image" validate:"required,gt=0,regexp=link"`
	Width      *int    `json:"width" validate:"required,gt=0"`
	Height     *int    `json:"height" validate:"required,gt=0"`
}

type Switch struct {
	Status *int `json:"status"`
}

type CampaignAd struct {
	Id         *int             `json:"ad_id" db:"id"`
	Name       *string          `json:"name" db:"name" validate:"required,gt=0"`
	AdType     *string          `json:"ad_type" db:"ad_type"`
	AdSize     *string          `json:"ad_size" db:"ad_size"`
	Width      *int             `json:"width" db:"width"`
	Height     *int             `json:"height" db:"height"`
	Duration   *float64         `json:"duration" db:"duration"`
	CampaignId *int             `json:"campaign_id" db:"campaign_id"`
	CreativeId *int             `json:"creative_id" db:"creative_id"`
	UserId     *int             `json:"user_id" db:"user_id"`
	Creative   *json.RawMessage `json:"creative" db:"creative" validate:"required,gt=0"`
	Status     *int             `json:"status" db:"status"`
}
type CampaignAds []CampaignAd

type AdCreative struct {
	Id          *int    `json:"creative_id" db:"id"`
	Name        *string `json:"name" db:"name"`
	AdType      *string `json:"ad_type" db:"ad_type"`
	AdSize      *string `json:"ad_size" db:"ad_size"`
	Material    *string `json:"material" db:"material"`
	Description *string `json:"description" db:"description"`
	OL          *int    `json:"ol" db:"ol"`
	Width       *int    `json:"width" db:"width"`
	Height      *int    `json:"height" db:"height"`
}
type AdCreativeList []AdCreative

// 单个广告活动
func GetCampaign(campaignId interface{}, userId interface{}, campaign *Campaign) error {
	if err := corm.Db.Get(campaign, `select * from 
		campaign_campaign where id=$1 and user_id=$2`, campaignId, userId); err != nil {
		return err
	}
	return nil
}

// 广告活动列表
func GetCampaigns(userId interface{}, total *int64, req *CampaignReq, campaigns *Campaigns) error {
	orm := corm.Select(`select id,name,bundle_id,category,sub_category,app_platform,
	budget_day,bidding_max,bidding_min,bidding_type,freq,speed,
	targeting,url,status,to_date(create_date::text, 'yyyy-mm-dd')
	create_date from campaign_campaign {{sql_where}}`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(campaigns); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 添加广告活动
func AddCampaign(userId interface{}, campaignId *int64, campaign *Campaign) error {
	if err := corm.Db.QueryRow(`insert into campaign_campaign(name,user_id,bundle_id,app_platform,
		category,sub_category, budget_day,
		bidding_max, bidding_min, bidding_type, freq, speed, targeting, url) values($1,$2,$3,$4,$5,$6,$7,$8,
		$9,$10,$11,$12,$13,$14) RETURNING id`, *campaign.Name, userId,
		*campaign.BundleId, *campaign.AppPlatform, *campaign.Category,
		*campaign.SubCategory, *campaign.BudgetDay, *campaign.BiddingMax, *campaign.BiddingMin, *campaign.BiddingType,
		string(*campaign.Freq), *campaign.Speed, string(*campaign.Targeting),
		string(*campaign.Url)).Scan(campaignId); err != nil {
		return err
	}
	return nil
}

// 修改广告活动
func UpdateCampaign(userId, campaignId interface{}, campaign *Campaign) error {
	if _, err := corm.Db.Exec(`update campaign_campaign set name=$1,
	bundle_id=$2,app_platform=$3,category=$4,sub_category=$5,
	budget_day=$6,bidding_min=$7,bidding_max=$8,bidding_type=$9,
	freq=$10,speed=$11,targeting=$12,url=$13  where id=$14 and user_id=$15`, *campaign.Name,
		*campaign.BundleId, *campaign.AppPlatform, *campaign.Category,
		*campaign.SubCategory, *campaign.BudgetDay, *campaign.BiddingMin,
		*campaign.BiddingMax, *campaign.BiddingType, string(*campaign.Freq),
		*campaign.Speed, string(*campaign.Targeting), string(*campaign.Url), campaignId, userId); err != nil {
		return err
	}
	return nil
}

// 修改广告活动状态
func UpdateCampaignStatus(userId, campaignId interface{}, status int) error {
	if _, err := corm.Db.Exec(`update campaign_campaign set 
	status=$1 where id=$2 and user_id=$3`, status, campaignId, userId); err != nil {
		return err
	}
	return nil
}

// 单个广告
func GetCampaignAd(adId interface{}, userId interface{}, campaignAd *CampaignAd) error {
	if err := corm.Db.Get(campaignAd, `select * from campaign_ad 
	where id=$1 and user_id=$2`, adId, userId); err != nil {
		return err
	}
	return nil
}

// 广告列表
func GetAds(userId interface{}, total *int64, req *AdReq, ads *CampaignAds) error {
	orm := corm.Select(`select *,(select name from campaign_campaign where id=t1.campaign_id) as
		campaign_name,(select email from account_account where user_id=t1.user_id) as
		user_email from campaign_ad as t1 {{sql_where}}`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(ads); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 添加视频广告
func AddVideoAd(userId interface{}, adId *int64, creative *AdCreative, video *VideoCreative) error {
	creativeStr, _ := json.Marshal(map[string]string{
		"image": *video.Image,
		"video": *video.Video,
	})
	if err := corm.Db.QueryRow(`insert into campaign_ad(name,ad_type,ad_size,
		ol,duration,campaign_id,width,height,user_id,creative_id,creative) 
		values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`, *video.Name,
		*creative.AdType, *creative.AdSize, *creative.OL,
		*video.Duration, *video.CampaignId, *video.Width,
		*video.Height, userId, *creative.Id, creativeStr).Scan(adId); err != nil {
		return err
	}
	return nil
}

// 添加图片广告
func AddImageAd(userId interface{}, adId *int64, creative *AdCreative, image *ImageCreative) error {
	creativeStr, _ := json.Marshal(map[string]string{
		"image": *image.Image,
	})
	if err := corm.Db.QueryRow(`insert into campaign_ad(name,ad_type,ad_size,
		ol,duration,campaign_id,width,height,user_id,creative_id,creative) 
		values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`, *image.Name,
		*creative.AdType, *creative.AdSize, *creative.OL,
		0, *image.CampaignId, *image.Width,
		*image.Height, userId, *creative.Id, creativeStr).Scan(adId); err != nil {
		return err
	}
	return nil
}

// 广告修改
func UpdateAdWithStatus(userId, AdId interface{}, status int, campaignAd *CampaignAd) error {
	if _, err := corm.Db.Exec(`update campaign_ad set status=$1,creative=$2,
	name=$3 where id=$4 and user_id=$5`, status, string(*campaignAd.Creative),
		*campaignAd.Name, AdId, userId); err != nil {
		return err
	}
	return nil
}

// 广告修改
func UpdateAd(userId, AdId interface{}, campaignAd *CampaignAd) error {
	if _, err := corm.Db.Exec(`update campaign_ad set creative=$1,
	name=$2 where id=$3 and user_id=$4`, string(*campaignAd.Creative),
		*campaignAd.Name, AdId, userId); err != nil {
		return err
	}
	return nil
}

// 创意列表
func GetCreatives(creatives *AdCreativeList) error {
	if err := corm.Db.Select(creatives, `select * from ad_creative`); err != nil {
		return err
	}
	return nil
}

// 单个创意
func GetCreative(creativeId interface{}, creative *AdCreative) error {
	if err := corm.Db.Get(creative, `select * from ad_creative where id=$1`, creativeId); err != nil {
		return err
	}
	return nil
}

// 人群包列表
func GetSegments(userId interface{}, total *int64, req *SegmentReq, segments *Segments) error {
	orm := corm.Select(`select * from campaign_segment {{sql_where}}`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(segments); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 人群包创建
func AddSegment(userId interface{}, segmentId *int64, uv int, segment *Segment) error {
	if err := corm.Db.QueryRow(`insert into campaign_segment(name,user_id,uv) values($1,$2,$3)
		on CONFLICT(name,user_id) do update set update_date=$4,uv=$5 RETURNING id`, *segment.Name,
		userId, uv, mydate.CurrentDateTime(), uv).Scan(segmentId); err != nil {
		return err
	}
	return nil
}

// 获取人群包
func GetSegment(userId interface{}, name string, segment *Segment) error {
	if err := corm.Db.Get(segment, `select * from campaign_segment
	where name=$1 and user_id=$2`, name, userId); err != nil {
		return err
	}
	return nil
}
