package handlers

import (
	"zonstfe_api/common/options"
	"net/http"

)

func (c *Handler) Options(w http.ResponseWriter, r *http.Request) {
	all_map := make(map[string]map[string]interface{}, 0)
	all_map["app_category"] = options.AppCategory
	all_map["app_status"] = options.AppStatus
	all_map["os"] = options.OsMap
	all_map["bank_map"] = options.BankMap
	all_map["app_store"] = options.AppStore
	all_map["province_city_name_map"] = options.ProvinceCityNameMap
	c.JsonBase(w, all_map)
}


