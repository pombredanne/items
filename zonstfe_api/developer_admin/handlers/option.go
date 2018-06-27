package handlers

import (
	"net/http"
	"zonstfe_api/common/options"
	"zonstfe_api/corm"
	"zonstfe_api/common/my_context"
)

type User struct {
	id    *int    `json:"user_id",db:"id"`
	Email *string `json:"email" db:"email"`
}
type Users []User

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

type OptionUser struct {
	Id    *int    `json:"user_id" db:"id"`
	Email *string `json:"email" db:"email"`
}
type OptionUserList []OptionUser

func (c *Handler) OptionUser(w http.ResponseWriter, r *http.Request) {
	users := &OptionUserList{}
	if err := corm.Select(`select id,email from user_user where role in (:role_id)`).Where(map[string]interface{}{
		"role_id": []int{1, 2},
	}).Loads(users); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, users)
}
