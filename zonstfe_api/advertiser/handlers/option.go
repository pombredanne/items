package handlers

import (
	"net/http"
	"zonstfe_api/common/options"
	"fmt"
	"zonstfe_api/corm"
	"zonstfe_api/common/my_context"
	"strconv"
	"encoding/json"

)

// 静态映射列表
func (c *Handler) Options(w http.ResponseWriter, r *http.Request) {
	//client := pb.NewGreeterClient(c.DataSync)
	//xx, err := client.CampaignCache(context.Background(), &pb.HelloRequest{Id:"1",EventId:uuid.NewV4().String()})
	//
	//fmt.Println(err, "-------", xx.Message)

	all_map := make(map[string]map[string]interface{}, 0)
	for k, v := range options.All {
		sub_map := make(map[string]interface{}, 0)
		for kk, vv := range v {
			sub_map[fmt.Sprintf("%v", kk)] = vv
		}
		all_map[k] = sub_map
	}
	all_map["geo_code"] = options.GeoCode
	all_map["app_category"] = options.AppCategory
	all_map["day_parting"] = options.DayParting
	all_map["vendor"] = options.Vendor
	all_map["android_version"] = options.AndroidVersion
	all_map["ios_version"] = options.IosVersion
	all_map["carrier"] = options.Carrier
	all_map["network"] = options.NetWork
	all_map["device_type"] = options.DeviceType
	all_map["campaign_status"] = options.CampaignStatus
	all_map["ad_status"] = options.AdStatus
	all_map["creative_set"] = options.CreativeSet
	all_map["os"] = options.OsMap
	all_map["province_city_code"] = options.ProvinceCityCodeMap
	c.JsonBase(w, all_map)
}

type OptionCampaign struct {
	Id   *int             `json:"campaign_id" db:"id"`
	Name *string          `json:"name" db:"name"`
	Ads  *json.RawMessage `json:"ads" db:"ads"`
}
type OptionCampaignList []OptionCampaign

// 广告活动和广告 映射列表
func (c *Handler) OptionCampaign(w http.ResponseWriter, r *http.Request) {
	current_user := c.GetUser(r)
	option_campaign := &OptionCampaignList{}
	if err := corm.Select(`select id,name,
		COALESCE((select array_to_json(array_agg(row)) from
		(select id,name from campaign_ad where campaign_id=t1.id) row
			),'[]') as ads from campaign_campaign as t1 where user_id=:user_id`).Where(map[string]interface{}{
		"user_id": current_user["id"],
	}).Loads(option_campaign); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	campaign_ad := make(map[string]interface{})
	for _, item := range *option_campaign {
		campaign_ad[strconv.Itoa(*item.Id)] = map[string]interface{}{
			"name": *item.Name,
			"ads":  *item.Ads,
		}
	}
	c.JsonBase(w, campaign_ad)
}
