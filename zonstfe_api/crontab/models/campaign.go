package models

import "encoding/json"

type Campaign struct {
	Id          *int             `json:"id" db:"id"`
	Name        *string          `json:"name" db:"name"`
	Status      *int             `json:"status" db:"status"`
	UserId      *int             `json:"user_id" db:"user_id"`
	Ads         *json.RawMessage `json:"ads" db:"ads"`
	Targeting   *json.RawMessage `json:"targeting" db:"targeting"`
	BiddingMax  *int             `json:"bidding_max" db:"bidding_max"`
	BiddingMin  *int             `json:"bidding_min" db:"bidding_min"`
	BudgetDay   *int             `json:"budget_day" db:"budget_day"`
	AppPlatform *string          `json:"app_platform" db:"app_platform"`
	Freq        *json.RawMessage `json:"freq" db:"freq"`
	Speed       *int             `json:"speed" db:"speed"`
	Category    *string          `json:"category" db:"category"`
	SubCategory *string          `json:"sub_category" db:"sub_category"`
	Url         *json.RawMessage `json:"url" db:"url"`
}

type FreqParser struct {
	Open *bool   `json:"open"`
	Type *string `json:"type"`
	Num  *int    `json:"num"`
}

type UrlParser struct {
	TrackingImpUrl *string `json:"tracking_imp_url"`
	TrackingClkUrl *string `json:"tracking_clk_url"`
	JumpUrl        *string `json:"jump_url"`
	DeepLinkUrl    *string `json:"deep_link_url"`
}

type Targeting struct {
	Vendor      *TargetingParser `json:"vendor" validate:"required"`       //{"open":false,list:["zonst"]}
	GeoCode     *TargetingParser `json:"geo_code" validate:"required"`     //{"open":false,list:{"province_code":[]}}
	AppCategory *TargetingParser `json:"app_category" validate:"required"` //{"open":false,list:{"monday":[]}}
	DayParting  *TargetingParser `json:"day_parting" validate:"required"`  //{"open":false,list:{"周一":[]}}
	DeviceType  *TargetingParser `json:"device_type" validate:"required"`  //{"open":false,list:[]}
	DeviceBrand *TargetingParser `json:"device_brand" validate:"required"` //{"open":false,list:[]}
	OsVersion   *TargetingParser `json:"os_version" validate:"required"`   //{"open":false,list:[]}
	Carrier     *TargetingParser `json:"carrier" validate:"required"`      //{"open":false,list:[]}
	NetWork     *TargetingParser `json:"network" validate:"required"`      //{"open":false,list:[]}
	Segment     *TargetingParser `json:"segment" validate:"required"`
}

type VideoCreative struct {
	Video *string `json:"video"`
	Image *string `json:"image"`
}

type BannerCreative struct {
	Image *string `json:"image"`
}
type InterCreative struct {
	Image *string `json:"image"`
}

type TargetingParser struct {
	Open bool     `json:"open"`
	List []string `json:"list"`
}

type Campaigns []Campaign
type Ad struct {
	Id       *int             `json:"id"`
	OL       *int             `json:"ol"`
	AdSize   *string          `json:"ad_size"`
	AdType   *string          `json:"ad_type"`
	Status   *int             `json:"status"`
	Creative *json.RawMessage `json:"creative"`
	Duration *float32         `json:"duration"`
}
type Ads []Ad

type Segment struct {
	Id *int `json:"id" db:"id"`
}
type Segments []Segment
