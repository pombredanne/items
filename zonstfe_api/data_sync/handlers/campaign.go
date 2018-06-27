package handlers

import (
	"net/http"
	"zonstfe_api/data_sync/models"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"strconv"
	"zonstfe_api/common/options"
	"gopkg.in/fatih/set.v0"
)

//整体思路 添加上线AD到当前KEY 删除下线AD到当前KEY 删除所有AD到当前除当前key的同类额外key
//如果没有开启定向则删除包含自己的所有同类讲上线的AD添加到未设置的key上
func (c *Handler) CampaignCacheAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	redis_clent := c.DspRd.Get()
	defer redis_clent.Close()
	redis_clent.Send("MULTI")
	start := time.Now()
	// 查询出广告活动和AD
	campaigns := &models.Campaigns{}
	if err := c.Pgx.Select(campaigns, `select *,(select COALESCE(array_to_json(array_agg(row_to_json(row))),'[]') as ads from (
    SELECT * from
    campaign_ad where campaign_id=t1.id ) row) from campaign_campaign as t1`); err != nil {
		c.Logger.Println(err)
		w.Write([]byte(""))
		return
	}
	for _, campaign := range *campaigns {
		segments := &models.Segments{}
		segment_list := make([]string, 0)
		// 查询出当前用户人群包
		if err := c.Pgx.Select(segments, `select id from campaign_segment where user_id=$1`, *campaign.UserId); err != nil {
			c.Logger.Println(err)
			w.Write([]byte(""))
			return
		}
		for _, segment := range *segments {
			segment_list = append(segment_list, strconv.Itoa(*segment.Id))
		}
		target := &models.Targeting{}
		campaign_ads := &models.Ads{}
		campaign_freq := &models.FreqParser{}
		campaign_url := &models.UrlParser{}
		if err := json.Unmarshal([]byte(*campaign.Ads), campaign_ads); err != nil {
			c.Logger.Println(err)
			w.Write([]byte(""))
			return
		}
		if err := json.Unmarshal([]byte(*campaign.Freq), campaign_freq); err != nil {
			c.Logger.Println(err)
			w.Write([]byte(""))
			return
		}
		if err := json.Unmarshal([]byte(*campaign.Url), campaign_url); err != nil {
			c.Logger.Println(err)
			w.Write([]byte(""))
			return
		}
		adids := make([]interface{}, 0)
		online_adids := make([]interface{}, 0)
		unadids := make([]interface{}, 0)

		for _, item := range *campaign_ads {
			adids = append(adids, fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))
			if *item.Status == 1 {
				online_adids = append(online_adids, fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))
			} else {
				unadids = append(unadids, fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))
			}
		}
		//在所有定向key 里面清除掉
		if err := json.Unmarshal([]byte(*campaign.Targeting), target); err != nil {
			c.Logger.Println(err)
			w.Write([]byte(""))
			return
		}
		hour_frequency, day_frequency := 0, 0
		if *campaign_freq.Open {
			switch *campaign_freq.Type {
			case "day":
				hour_frequency = 0
				day_frequency = *campaign_freq.Num
			case "hour":
				hour_frequency = *campaign_freq.Num
				day_frequency = 0
			default:
				hour_frequency = 0
				day_frequency = 0
			}
		} else {
			hour_frequency = 0
			day_frequency = 0

		}
		// campaign 对象
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "low_price", *campaign.BiddingMin*1000)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "high_price", *campaign.BiddingMax*1000)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "budget", *campaign.BudgetDay*1000000)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "rate", *campaign.Speed)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "category", *campaign.Category)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "hour_frequency", hour_frequency)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "day_frequency", day_frequency)
		redis_clent.Send("HSET", fmt.Sprintf("campaign_%d", *campaign.Id), "account_id", *campaign.UserId)
		//排除没有AD的
		if len(adids) > 0 {
			//活动下线状态 剔除所有AD
			if *campaign.Status != 1 {
				//category
				for _, category := range options.CategoryList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category1_%s", category)}, adids...)...)

				}
				for _, sub_category := range options.SubCategoryList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category2_%s", sub_category)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_category"}, adids...)...)
				// 国家
				for _, country := range options.CountryList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_country_%s", country)}, adids...)...)
				}
				// 省份
				for _, province := range options.ProvinceList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_province_%s", province)}, adids...)...)
				}
				// 城市
				for _, city := range options.CityList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_city_%s", city)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_location"}, adids...)...)
				// 运营商
				for _, sp := range options.CarrierList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_sp_%s", sp)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_sp"}, adids...)...)
				// 网络
				for _, network := range options.NetWorkList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_network_%s", network)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_network"}, adids...)...)

				// 设备类型
				for _, device_type := range options.DeviceTypeList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_device_type_%s", device_type)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_device_type"}, adids...)...)
				// 投放平台
				for _, vendor := range options.VendorList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_vendor_%s", vendor)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_vendor"}, adids...)...)

				//系统
				//for _, os := range options.OsList {
				//	redis_clent.Send("SREM", fmt.Sprintf("limit_os_%s", os), strings.Join(adids, " "))
				//}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_os"}, adids...)...)

				//week
				for _, week := range options.DayPartingList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_week_%s", week)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_week"}, adids...)...)
				//ad_size
				for _, ad_size := range options.AdSizeList {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_ad_size_%s", ad_size)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_ad_size"}, adids...)...)
				//segment
				for _, segment := range segment_list {
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("segment_%s", segment)}, adids...)...)
				}
				redis_clent.Send("SREM", append([]interface{}{"unlimit_segment"}, adids...)...)
				//os_version  是否可以使用 srem "" ""_*?
				for _, os := range options.OsList {
					switch os {
					case "ios":
						for _, item := range options.IosVersionList {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", os, item)}, adids...)...)

						}
					case "android":
						for _, item := range options.AndroidVersionList {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", os, item)}, adids...)...)
						}
					}
					redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("unlimit_os_%s", os)}, adids...)...)
				}

			} else {
				if target.OsVersion.Open {

					os_version_list := make([]string, 0)
					switch *campaign.AppPlatform {
					case "ios":
						os_version_list = options.IosVersionList
					case "android":
						os_version_list = options.AndroidVersionList
					}
					os_version_s := set.New(StringSliceToInterFace(os_version_list)...)
					my_os_version_s := set.New(StringSliceToInterFace(target.OsVersion.List)...)
					diff_os_version_list := set.StringSlice(set.Difference(os_version_s, my_os_version_s))
					intersection_os_version_list := set.StringSlice(set.Intersection(os_version_s, my_os_version_s))
					for _, intersection_os_version := range intersection_os_version_list {
						if len(online_adids) > 0 {

							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, intersection_os_version)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, intersection_os_version)}, unadids...)...)
						}
					}
					for _, diff_os_version := range diff_os_version_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, diff_os_version)}, adids...)...)
					}
				} else {
					switch *campaign.AppPlatform {
					case "ios":
						for _, os := range options.IosVersionList {
							if len(online_adids) > 0 {
								redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, os)}, online_adids...)...)
							}
							if len(unadids) > 0 {
								redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, os)}, unadids...)...)
							}

						}
					case "android":
						for _, os := range options.AndroidVersionList {
							if len(online_adids) > 0 {
								redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, os)}, online_adids...)...)
							}
							if len(unadids) > 0 {
								redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_os_%s_%s", *campaign.AppPlatform, os)}, unadids...)...)
							}

						}
					}
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("unlimit_os_%s", *campaign.AppPlatform)}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("unlimit_os_%s", *campaign.AppPlatform)}, unadids...)...)
					}

				}

				if target.Segment.Open {

					segment_s := set.New(StringSliceToInterFace(segment_list)...)
					my_segment_s := set.New(StringSliceToInterFace(target.Segment.List)...)
					diff_segment_list := set.StringSlice(set.Difference(segment_s, my_segment_s))
					intersection_segment_list := set.StringSlice(set.Intersection(segment_s, my_segment_s))
					for _, intersection_segment := range intersection_segment_list {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("segment_%s", intersection_segment)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("segment_%s", intersection_segment)}, unadids...)...)
						}
					}
					for _, diff_segment := range diff_segment_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("segment_%s", diff_segment)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_segment"}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_segment"}, unadids...)...)
					}
					for _, segment := range segment_list {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("segment_%s", segment)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("segment_%s", segment)}, unadids...)...)
						}

					}
				}
				// category

				if target.AppCategory.Open {
					// 还需要剔除没有勾选
					categorys_s := set.New(StringSliceToInterFace(options.CategoryList)...)
					my_categorys_s := set.New(StringSliceToInterFace(target.AppCategory.List)...)
					diff_category_list := set.StringSlice(set.Difference(categorys_s, my_categorys_s))
					intersection_category_list := set.StringSlice(set.Intersection(categorys_s, my_categorys_s))
					sub_categorys_s := set.New(StringSliceToInterFace(options.SubCategoryList)...)

					diff_sub_category_list := set.StringSlice(set.Difference(sub_categorys_s, my_categorys_s))
					intersection_sub_category_list := set.StringSlice(set.Intersection(sub_categorys_s, my_categorys_s))
					for _, intersection_category := range intersection_category_list {
						// 加入上线AD
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_category1_%s", intersection_category)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							// 剔除下线AD
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category1_%s", intersection_category)}, unadids...)...)

						}

					}
					for _, diff_category := range diff_category_list {
						// 剔除上下线AD
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category1_%s", diff_category)}, adids...)...)
					}
					for _, intersection_sub_category := range intersection_sub_category_list {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_category2_%s", intersection_sub_category)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category2_%s", intersection_sub_category)}, unadids...)...)
						}

					}
					for _, diff_sub_category := range diff_sub_category_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category2_%s", diff_sub_category)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_category"}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_category"}, unadids...)...)
					}
					for _, category := range options.CategoryList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_category1_%s", category)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category1_%s", category)}, unadids...)...)
						}

					}
					for _, sub_category := range options.SubCategoryList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_category2_%s", sub_category)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_category2_%s", sub_category)}, unadids...)...)
						}
					}
				}
				// geo
				if target.GeoCode.Open {

					my_geocode_s := set.New(StringSliceToInterFace(target.GeoCode.List)...)
					countrys_s := set.New(StringSliceToInterFace(options.CountryList)...)
					provinces_s := set.New(StringSliceToInterFace(options.ProvinceList)...)
					citys_s := set.New(StringSliceToInterFace(options.CityList)...)

					diff_country_list := set.StringSlice(set.Difference(countrys_s, my_geocode_s))
					intersection_country_list := set.StringSlice(set.Intersection(countrys_s, my_geocode_s))
					diff_province_list := set.StringSlice(set.Difference(provinces_s, my_geocode_s))
					intersection_province_list := set.StringSlice(set.Intersection(provinces_s, my_geocode_s))
					diff_city_list := set.StringSlice(set.Difference(citys_s, my_geocode_s))
					intersection_city_list := set.StringSlice(set.Intersection(citys_s, my_geocode_s))
					for _, intersection_country := range intersection_country_list {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_country_%s", intersection_country)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_country_%s", intersection_country)}, unadids...)...)
						}
					}

					for _, diff_country := range diff_country_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_country_%s", diff_country)}, adids...)...)
					}

					for _, intersection_province := range intersection_province_list {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_province_%s", intersection_province)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_province_%s", intersection_province)}, unadids...)...)
						}

					}

					for _, diff_province := range diff_province_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_province_%s", diff_province)}, adids...)...)
					}

					for _, intersection_city := range intersection_city_list {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_city_%s", intersection_city)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_city_%s", intersection_city)}, unadids...)...)

						}

					}

					for _, diff_city := range diff_city_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_city_%s", diff_city)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_location"}, online_adids...)...)

					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_location"}, unadids...)...)
					}

					// 国家
					for _, country := range options.CountryList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_country_%s", country)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_country_%s", country)}, unadids...)...)
						}

					}
					// 省份
					for _, province := range options.ProvinceList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_province_%s", province)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_province_%s", province)}, unadids...)...)
						}

					}
					// 城市
					for _, city := range options.CityList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_city_%s", city)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_city_%s", city)}, unadids...)...)
						}

					}

				}
				// sp

				if target.Carrier.Open {

					sp_s := set.New(StringSliceToInterFace(options.CarrierList)...)
					my_sp_s := set.New(StringSliceToInterFace(target.Carrier.List)...)
					diff_carrier_list := set.StringSlice(set.Difference(sp_s, my_sp_s))

					for _, sp := range target.Carrier.List {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_sp_%s", sp)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_sp_%s", sp)}, unadids...)...)
						}
					}

					for _, sp := range diff_carrier_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_sp_%s", sp)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_sp"}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_sp"}, unadids...)...)
					}
					for _, sp := range options.CarrierList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_sp_%s", sp)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_sp_%s", sp)}, unadids...)...)
						}

					}

				}
				// network
				if target.NetWork.Open {
					network_s := set.New(StringSliceToInterFace(options.NetWorkList)...)
					my_network_s := set.New(StringSliceToInterFace(target.NetWork.List)...)
					diff_network_list := set.StringSlice(set.Difference(network_s, my_network_s))
					for _, network := range target.NetWork.List {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_network_%s", network)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_network_%s", network)}, unadids...)...)
						}

					}
					for _, network := range diff_network_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_network_%s", network)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_network"}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_network"}, unadids...)...)
					}

					for _, network := range options.NetWorkList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_network_%s", network)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_network_%s", network)}, unadids...)...)
						}

					}

				}

				// device_type
				if target.DeviceType.Open {

					device_type_s := set.New(StringSliceToInterFace(options.DeviceTypeList)...)
					my_device_type_s := set.New(StringSliceToInterFace(target.DeviceType.List)...)
					diff_device_type_list := set.StringSlice(set.Difference(device_type_s, my_device_type_s))
					for _, device_type := range target.DeviceType.List {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_device_type_%s", device_type)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_device_type_%s", device_type)}, unadids...)...)
						}

					}
					for _, device_type := range diff_device_type_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_device_type_%s", device_type)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_device_type"}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_device_type"}, unadids...)...)
					}
					for _, device_type := range options.DeviceTypeList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_device_type_%s", device_type)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_device_type_%s", device_type)}, unadids...)...)

						}

					}
				}
				// os
				//if len(online_adids) > 0 {
				//	redis_clent.Send("SADD", fmt.Sprintf("limit_os_%s", *campaign.AppPlatform), strings.Join(online_adids, " "))
				//}
				//if len(unadids) > 0 {
				//	redis_clent.Send("SREM", fmt.Sprintf("limit_os_%s", *campaign.AppPlatform), strings.Join(unadids, " "))
				//}

				// week

				if target.DayParting.Open {
					dayparting_s := set.New(StringSliceToInterFace(options.DayPartingList)...)
					my_dayparting_s := set.New(StringSliceToInterFace(target.DayParting.List)...)
					diff_day_list := set.StringSlice(set.Difference(dayparting_s, my_dayparting_s))
					for _, day := range target.DayParting.List {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_week_%s", day)}, online_adids...)...)
						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_week_%s", day)}, unadids...)...)
						}

					}
					for _, day := range diff_day_list {
						redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_week_%s", day)}, adids...)...)
					}

				} else {
					if len(online_adids) > 0 {
						redis_clent.Send("SADD", append([]interface{}{"unlimit_week"}, online_adids...)...)
					}
					if len(unadids) > 0 {
						redis_clent.Send("SREM", append([]interface{}{"unlimit_week"}, unadids...)...)
					}
					for _, day := range options.DayPartingList {
						if len(online_adids) > 0 {
							redis_clent.Send("SADD", append([]interface{}{fmt.Sprintf("limit_week_%s", day)}, online_adids...)...)

						}
						if len(unadids) > 0 {
							redis_clent.Send("SREM", append([]interface{}{fmt.Sprintf("limit_week_%s", day)}, unadids...)...)

						}

					}

				}
				track_host := ""
				if c.EnvModel == "production" {
					track_host = "https://t.qinglong365.com/v1/g"
				} else {
					track_host = "http://10.0.2.201/v1/g"
				}

				// ad_size
				for _, item := range *campaign_ads {

					if *item.Status == 1 {
						//生成code
						common_params := fmt.Sprintf("account_id=%s&ad_id=%s&ad_type=%s&"+
							"ad_size=%s&request_id={{.RequestID}}&vendor_id={{.VendorID}}"+
							"&app_id={{.App.AppID}}&idtype={{.IDType}}&id={{.ID}}&adslot_id={{.TmpADSlotID}}&ol={{.TmpOL}}",
							strconv.Itoa(*campaign.UserId), fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id), *item.AdType, *item.AdSize)
						ad_eimp_tracking_url := track_host + "/t_ad_eimp?" + common_params
						imp_r, clk_r := "", ""
						if *campaign_url.TrackingImpUrl != "" {
							imp_r = "&r=" + *campaign_url.TrackingImpUrl
						}
						if *campaign_url.TrackingClkUrl == "" {
							clk_r = "&r=" + *campaign_url.TrackingClkUrl
						}

						end_tracking_url := track_host + "/t_video_end?" + common_params + "&time={{.Time}}"
						ad_imp_tracking_url := track_host + "/t_ad_imp?" + common_params + imp_r
						click_tracking_url := track_host + "/t_ad_clk?" + common_params + clk_r
						close_tracking_url := track_host + "/t_video_close?" + common_params + "&time={{.Time}}"
						delay_tracking_url := track_host + "/t_video_delay?" + common_params + "&time={{.Time}}"
						code_map := make(map[string]interface{}, 0)
						switch *item.AdType {
						case "video":
							ad_creative := &models.VideoCreative{}
							err := json.Unmarshal([]byte(*item.Creative), ad_creative)
							if err != nil {
								fmt.Println(err)

							}
							html := StringFormat(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><div style="position:fixed;z-index:-1; left:0; top:0;width:100%;height:100%;display: flex;align-items: center;justify-content: center;background: #eee"><img src="{{image}}" style="width: 100%;height: auto;" onclick="window.open('{{click_sdk_dp_url}}')" /></div><img src="{{ad_eimp_tracking_url}}" style="position:fixed;top:-100%;width:0px;height:0px;" /></body></html>`, map[string]interface{}{
								"image":                *ad_creative.Image,
								"ad_eimp_tracking_url": ad_eimp_tracking_url,
								"click_sdk_dp_url":     "zonst://81d1370556/clk",
							})
							code_map = map[string]interface{}{
								"html":                html,
								"video":               *ad_creative.Video,
								"ad_imp_tracking_url": ad_imp_tracking_url,
								"end_tracking_url":    end_tracking_url,
								"click_sdk_dp_url":    "zonst://81d1370556/clk",
								"click_tracking_url":  click_tracking_url,
								"close_tracking_url":  close_tracking_url,
								"deeplink_url":        *campaign_url.DeepLinkUrl,
								"download_url":        *campaign_url.JumpUrl,
								"delay_tracking_url":  delay_tracking_url,
								"ol":                  *item.OL,
								"duration":            *item.Duration * 1000,
							}
						case "graphic":
							switch *item.AdSize {
							case "banner":
								ad_creative := &models.BannerCreative{}
								err := json.Unmarshal([]byte(*item.Creative), ad_creative)
								if err != nil {
									c.Logger.Println(err)
									w.Write([]byte(""))
									return

								}

								code_map = map[string]interface{}{
									"html": StringFormat(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><div style="position:fixed;z-index:-1; left:0; top:0;width:100%;height:100%;display: flex;align-items: center;justify-content: center;background: #eee"><img src="{{image}}" style="width: 100%;height: auto;" onclick="window.open('{{click_sdk_dp_url}}')" /></div><img src="{{ad_eimp_tracking_url}}" style="position:fixed;top:-100%;width:0px;height:0px;" /></body></html>`, map[string]interface{}{
										"image":                *ad_creative.Image,
										"ad_eimp_tracking_url": ad_imp_tracking_url,
										"click_sdk_dp_url":     "zonst://81d1370556/clk",
									}),
									"ad_imp_tracking_url": ad_imp_tracking_url,
									"click_sdk_dp_url":    "zonst://81d1370556/clk",
									"click_tracking_url":  click_tracking_url,
									"close_tracking_url":  close_tracking_url,
									"deeplink_url":        *campaign_url.DeepLinkUrl,
									"download_url":        *campaign_url.JumpUrl,
									"ol":                  *item.OL,
									"duration":            *item.Duration * 1000,
								}
							case "fullscreen":
								ad_creative := &models.InterCreative{}
								err := json.Unmarshal([]byte(*item.Creative), ad_creative)
								if err != nil {
									c.Logger.Println(err)
									w.Write([]byte(""))
									return
								}
								html := StringFormat(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><div style="position:fixed;z-index:-1; left:0; top:0;width:100%;height:100%;display: flex;align-items: flex-start;justify-content: center;background: #eee"><img src="{{image}}" style="width: 100%;height: auto;" onclick="window.open('{{click_sdk_dp_url}}')" /></div><img src="{{ad_eimp_tracking_url}}" style="position:fixed;top:-100%;width:0px;height:0px;" /></body></html>`, map[string]interface{}{
									"image":                *ad_creative.Image,
									"ad_eimp_tracking_url": ad_imp_tracking_url,
									"click_sdk_dp_url":     "zonst://81d1370556/clk",
								})
								code_map = map[string]interface{}{
									"html":                html,
									"ad_imp_tracking_url": ad_imp_tracking_url,
									"click_sdk_dp_url":    "zonst://81d1370556/clk",
									"click_tracking_url":  click_tracking_url,
									"close_tracking_url":  close_tracking_url,
									"deeplink_url":        *campaign_url.DeepLinkUrl,
									"download_url":        *campaign_url.JumpUrl,
									"ol":                  *item.OL,
									"duration":            *item.Duration * 1000,
								}
							}

						}
						buffer := &bytes.Buffer{}
						encoder := json.NewEncoder(buffer)
						encoder.SetIndent("", " ")
						encoder.SetEscapeHTML(false)
						err := encoder.Encode(code_map)
						if err != nil {
							c.Logger.Println(err)
							w.Write([]byte(""))
							return
						}
						redis_clent.Send("HSET", "ad_code", fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id), strings.Replace(string(buffer.Bytes()), "\n", "", -1))
						if *item.AdType == "video" {
							redis_clent.Send("SADD", fmt.Sprintf("limit_ad_size_%s", "video"), fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))

						} else {
							redis_clent.Send("SADD", fmt.Sprintf("limit_ad_size_%s", *item.AdSize), fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))

						}
					} else {
						//删除AD
						redis_clent.Send("HDEL", "ad_code", fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))
						if *item.AdType == "video" {
							redis_clent.Send("SREM", fmt.Sprintf("limit_ad_size_%s", "video"), fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))

						} else {
							redis_clent.Send("SREM", fmt.Sprintf("limit_ad_size_%s", *item.AdSize), fmt.Sprintf("%d", (*campaign.Id << 32) + *item.Id))

						}
					}
				}

			}
		}

	}
	_, err := redis_clent.Do("EXEC")
	if err != nil {
		c.Logger.Println(err)
		w.Write([]byte(""))
		return
	}
	fmt.Println(time.Since(start))
	w.Write([]byte(""))

}

func StringSliceToInterFace(t []string) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	return s
}

func StringFormat(format string, p map[string]interface{}) string {
	args, i := make([]string, len(p)*2), 0
	for k, v := range p {
		args[i] = "{{" + k + "}}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}
