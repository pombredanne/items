package handlers

import (
	"net/http"
	"zonstfe_api/common/options"
	"fmt"
	"zonstfe_api/corm"
	"zonstfe_api/common/my_context"
)

func (c *Handler) Options(w http.ResponseWriter, r *http.Request) {
	all_map := make(map[string]map[string]interface{}, 0)
	for k, v := range options.All {
		sub_map := make(map[string]interface{}, 0)
		for kk, vv := range v {
			sub_map[fmt.Sprintf("%v", kk)] = vv
		}
		all_map[k] = sub_map
	}

}

type OptionUser struct {
	Id    *int    `json:"user_id" db:"id"`
	Email *string `json:"email" db:"email"`
}
type OptionUserList []OptionUser

func (c *Handler) OptionUser(w http.ResponseWriter, r *http.Request) {
	users := &OptionUserList{}
	if err := corm.Select(`select id,email from user_user where role in (:role_id)`).Where(map[string]interface{}{
		"role_id": []int{3, 4},
	}).Loads(users); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, users)
}
