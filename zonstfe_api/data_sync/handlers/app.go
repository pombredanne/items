package handlers

import (
	"net/http"
	"zonstfe_api/data_sync/models"
	"fmt"
	"strings"
	"encoding/json"
)

func (c *Handler) AppCacheAll(w http.ResponseWriter, r *http.Request) {
	redis_clent := c.SspRd.Get()
	defer redis_clent.Close()
	redis_clent.Send("MULTI")
	apps := &models.Apps{}
	if err := c.Pgx.Select(apps, `select t1.*,t2.app_key,t2.deal_scale,t2.deal_type from (select * from app_app) as t1
	LEFT JOIN (select app_key,deal_type,deal_scale,user_id from account_account) as t2 ON t1.user_id=t2.user_id`); err != nil {
		c.Logger.Println(err)
		w.Write([]byte(""))
		return
	}
	for _, app := range *apps {
		slots := map[string]float32{}
		if err := json.Unmarshal([]byte(*app.Slots), &slots); err != nil {
			c.Logger.Println(err)
			w.Write([]byte(""))
			return
		}
		if *app.Status == 1 {
			category_limit := &models.CategoryLimit{}
			if err := json.Unmarshal([]byte(string(*app.CategoryLimit)), category_limit); err != nil {
				c.Logger.Println(err)
				w.Write([]byte(""))
				return
			}

			app_map := map[string]interface{}{
				"limit_category": strings.Join(category_limit.List, ","),
				"category1":      *app.Category,
				"category2":      *app.SubCategory,
			}
			app_map_json, _ := json.Marshal(app_map)
			redis_clent.Send("HSET", "app_list", fmt.Sprintf("%s:%s:%s", *app.AppKey, *app.Os, *app.BundleId), string(app_map_json))
			redis_clent.Send("HSET", "app_reward", fmt.Sprintf("%s:%s:%s", *app.AppKey, *app.Os, *app.BundleId), string(*app.Reward))
			for key, value := range slots {
				bidding := map[string]interface{}{
					"deal_type": *app.DealType,
				}
				if *app.DealType == "share" {
					bidding["deal_value"] = int(*app.DealScale * 10000)
				}
				if *app.DealType == "bidding" {
					bidding["deal_value"] = int(value * 1000)
				}
				bidding_json, _ := json.Marshal(bidding)
				redis_clent.Send("HSET", "app_slot",
					fmt.Sprintf("%s:%s:%s:%s", *app.AppKey, *app.Os, *app.BundleId, key), string(bidding_json))

			}
		} else {
			redis_clent.Send("HDEL", "app_list", fmt.Sprintf("%s:%s:%s", *app.AppKey, *app.Os, *app.BundleId))
			redis_clent.Send("HDEL", "app_reward", fmt.Sprintf("%s:%s:%s", *app.AppKey, *app.Os, *app.BundleId))
			for key, _ := range slots {
				redis_clent.Send("HDEL", "app_slot",
					fmt.Sprintf("%s:%s:%s:%s", *app.AppKey, *app.Os, *app.BundleId, key))
			}
		}

	}
	_, err := redis_clent.Do("EXEC")
	if err != nil {
		c.Logger.Println(err)
		w.Write([]byte(""))
		return
	}
	w.Write([]byte(""))
}

func (c *Handler) AppIdentifier(w http.ResponseWriter, r *http.Request) {
	redis_clent := c.SspRd.Get()
	defer redis_clent.Close()
	redis_clent.Send("MULTI")
	identifiers := &models.Identifiers{}
	if err := c.Pgx.Select(identifiers, `select device,identifier from spider_identifier`); err != nil {
		c.Logger.Println(err)
		w.Write([]byte(""))
		return
	}
	for _, item := range *identifiers {
		redis_clent.Send("HSET", "app_device", fmt.Sprintf("%s:%s", "ios", strings.ToLower(*item.Identifier)), strings.ToLower(*item.Device))
	}
	_, err := redis_clent.Do("EXEC")
	if err != nil {
		c.Logger.Println(err)
		w.Write([]byte(""))
		return
	}
	w.Write([]byte(""))
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
