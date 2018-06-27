package models

import (
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
)

type ActionReq struct {
	UserId       *int    `form:"user_id" db:"user_id" query:"eq"`
	ActionUserId *string `form:"action_user_id" db:"action_user_id" query:"action_user_id"`
	ActionModule *string `form:"action_module" db:"action_module" query:"eq"`
	PlatformId   *string `form:"platform_id" db:"platform_id" query:"eq"`
	Sdate        *string `form:"sdate" db:"create_date" query:"gte"`
	Edate        *string `form:"edate" db:"create_date" query:"lte"`
	Page         *uint   `form:"page"`
	PageSize     *uint   `form:"page_size"`
}
type Action struct {
	UserId          *int    `json:"user_id" db:"user_id"`
	ActionUserId    *int    `json:"action_user_id" db:"action_user_id"`
	UserEmail       *string `json:"user_email" db:"user_email"`
	UserActionEmail *string `json:"user_action_email" db:"user_action_email"`
	//ActionSql     *string `json:"action_sql" db:"action_sql"`
	ActionType   *string `json:"action_type" db:"action_type"`
	ActionMoudle *string `json:"action_module" db:"action_module"`
	ActionId     *int    `json:"action_id" db:"action_id"`
	//RequestPath   *string `json:"request_path" db:"request_path"`
	//RequestMethod *string `json:"request_method" db:"request_method"`
	//RequestData   *string `json:"request_data" db:"request_data"`
	PlatformId *int            `json:"platform_id" db:"platform_id"`
	IpAddress  *string         `json:"ip_address" db:"ip_address"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
}

type Actions []Action

// 操作记录
// 操作记录
func GetActions(platformIds []int, total *int64, req *ActionReq, actions *Actions) error {
	orm := corm.Select(`select * from (select *,(select email from user_user where id=user_id) as user_email,
	(select email from user_user where id=action_user_id) as user_action_email from log_actions
	ORDER BY create_date desc) as t200 {{sql_where}} `).Req(req).Where("platform_id", "in", platformIds).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(actions); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}
