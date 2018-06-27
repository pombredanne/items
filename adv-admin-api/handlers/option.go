package handlers

import (
	"api-libs/option"
	"fmt"
	"corm"
	"github.com/gin-gonic/gin"
	"api-libs/rsp"
)

func Options(c *gin.Context) {
	all_map := make(map[string]map[string]interface{}, 0)
	for k, v := range option.All {
		sub_map := make(map[string]interface{}, 0)
		for kk, vv := range v {
			sub_map[fmt.Sprintf("%v", kk)] = vv
		}
		all_map[k] = sub_map
	}
	all_map["geo_code"] = option.GeoCode
	all_map["geo_code_name"] = option.GeoCodeName
	all_map["app_category"] = option.AppCategory
	all_map["day_parting"] = option.DayParting
	all_map["vendor"] = option.Vendor
	all_map["android_version"] = option.AndroidVersion
	all_map["ios_version"] = option.IosVersion
	all_map["carrier"] = option.Carrier
	all_map["network"] = option.NetWork
	all_map["device_type"] = option.DeviceType
	all_map["campaign_status"] = option.CampaignStatus
	all_map["ad_status"] = option.AdStatus
	all_map["creative_set"] = option.CreativeSet
	all_map["os"] = option.OsMap
	all_map["province_city_code"] = option.ProvinceCityCodeMap
	all_map["bank_code"] = option.BankMap
	all_map["industry_code"] = option.IndustryInformation
	c.JSON(rsp.Base(all_map))
}

type OptionUserReq struct {
	UserRole *string `form:"user_role" db:"user_role" query:"in"`
}

type OptionUser struct {
	Id          *int    `json:"user_id" db:"user_id"`
	Email       *string `json:"email" db:"email"`
	RealName    *string `json:"real_name" db:"real_name"`
	CompanyName *string `json:"company_name" db:"company_name"`
}
type OptionUserList []OptionUser

func GetOptionUser(c *gin.Context) {
	var req OptionUserReq
	users := &OptionUserList{}
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}

	if err := corm.Select(`select user_id,email,real_name,
		company_name from 
		account_account {{sql_where}}`).Where("user_role", "in", []interface{}{3, 4}).Req(&req).Loads(users); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base(users))
}

type OptionCampaign struct {
	Id   *int    `json:"campaign_id" db:"id"`
	Name *string `json:"name" db:"name"`
}
type OptionCampaignList []OptionCampaign

type campaignReq struct {
	UserId *int `form:"user_id" db:"user_id" query:"eq"`
}

func GetOptionCampaign(c *gin.Context) {
	var req = campaignReq{}
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	campaigns := &OptionCampaignList{}
	if err := corm.Select(`select id,name from campaign_campaign {{sql_where}}`).Req(&req).Loads(campaigns); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base(campaigns))
}

type SegmentReq struct {
	Name   *string `form:"name" db:"name" query:"like"`
	UserId *int    `form:"user_id" db:"user_id" query:"eq"`
}
type OptionSegment struct {
	Id     *int    `json:"segment_id" db:"id"`
	UserId *int    `json:"user_id" db:"user_id"`
	Name   *string `json:"name" db:"name"`
}
type OptionSegmentList []OptionCampaign

func GetOptionSegment(c *gin.Context) {
	var req = SegmentReq{}
	segments := &OptionSegmentList{}
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	orm := corm.Select(`select * from campaign_segment {{sql_where}}`).Req(req)
	if err := orm.Loads(segments); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base(segments))
}
