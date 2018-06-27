package models

import (
	"net/http"
	"api-libs/my-field"
	"api-libs/rsp"
	"corm"
	"adv-admin-api/config"
)

var (
	ActionModule = actionModule{
		Account:  "账户",
		Campaign: "广告系列",
		App:      "媒体",
		Ad:       "广告",
		//TarGet:      "定向",
		CreativeSet: "创意",
		Finance:     "财务",
		Segment:     "人群包",
	}
)

type actionModule struct {
	Account  string
	Campaign string
	Ad       string
	//TarGet      string
	CreativeSet string
	Finance     string
	Segment     string
	App         string
}

func LogAction(userId, actionUserId, actionModuleId interface{}, actionModule, actionType, ip string, r *http.Request) {
	//sqls := strings.Join(sql_list, ";")
	query := `insert into log_actions (user_id,action_user_id,action_sql,request_path,
	request_method,platform_id,ip_address,
	action_module,action_type,action_id) values ($1, $2, $3,$4,$5,$6,$7,$8,$9,$10)`
	args := []interface{}{userId, actionUserId, "", r.RequestURI, r.Method, config.Conf.RoleId, ip,
		actionModule, actionType, actionModuleId}
	if _, err := corm.Db.Exec(query, args...); err != nil {
		rsp.HandlerError(err)
	}
}

type ActionReq struct {
	UserId       *int    `form:"user_id" db:"user_id" query:"eq"`
	ActionUserId *string `form:"action_user_id" db:"action_user_id" query:"action_user_id"`
	ActionModule *string `form:"action_module" db:"action_module" query:"eq"`
	PlatformId   *string `form:"user_role" db:"platform_id" query:"eq"`
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
	CreateDate *my_field.JSONTime `json:"create_date" db:"create_date"`
}

type Actions []Action

// 操作记录
func GetActions(platformIds []int, total *int64, req *ActionReq, actions *Actions) error {
	sql := `select * from (select *,(select email from user_user where id=user_id) as user_email,
	(select email from user_user where id=action_user_id) as user_action_email from log_actions
	ORDER BY create_date desc) as t200 {{sql_where}}`
	orm := corm.Select(sql).Req(&req).Where("platform_id", "in", platformIds)
	if err := orm.Paginate(req.Page, req.PageSize).Loads(actions); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}
