package corm

import (
	"testing"
	"fmt"
	"encoding/json"
)

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

type Req struct {
	Status   *int  `form:"status" db:"status" query:"eq"`
	Page     *uint `form:"page"`
	PageSize *uint `form:"page_size"`
}
type CampaignAds []CampaignAd

//func init() {
//	Db = mydb.GetPgx("postgresql://fe:qinglong@111.231.137.127:5432/test_fe?sslmode=disable")
//}

func TestSelectLoads(test *testing.T) {
	campaign_ads := &CampaignAds{}
	if err := Select(`select * from campaign_ad`).Loads(campaign_ads); err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(*campaign_ads))
}

func TestSelectLoad(test *testing.T) {
	campaign_ad := &CampaignAd{}
	if err := Select(`select * from campaign_ad where id=:id and user_id=:user_id`).Where(map[string]interface{}{
		"id":      864140,
		"user_id": 10002,
	}).Load(campaign_ad); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*campaign_ad.Id)
}

func TestSelectLoad2(test *testing.T) {
	campaign_ad := &CampaignAd{}
	if err := Select(`select * from campaign_ad where id=:id and user_id=:user_id`).Where([][]interface{}{
		{"id", "=", 864140},
		{"user_id", "=", 10002},
	}).Load(campaign_ad); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*campaign_ad.Id, "2")
}
func TestSelectLoad3(test *testing.T) {
	campaign_ad := &CampaignAd{}
	if err := Select(`select * from campaign_ad where id=:id`).Where("id", "=", 864140).Load(campaign_ad); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*campaign_ad.Id, "3")
}

func TestSelectLoad4(test *testing.T) {
	campaign_ad := &CampaignAd{}
	if err := Select(`select * from campaign_ad {{sql_where}}`).Where("id=864140").Load(campaign_ad); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*campaign_ad.Id, "4")
}

func TestSelectByReq(test *testing.T) {
	campaign_ads := &CampaignAds{}

	req := &Req{
		Status:   func(i int) *int { return &i }(1),
		Page:     func(i uint) *uint { return &i }(1),
		PageSize: func(i uint) *uint { return &i }(20),
	}
	if err := Select("select * from campaign_ad {{sql_where}}").Paginate(req.Page, req.PageSize).Loads(campaign_ads); err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(*campaign_ads))
}
