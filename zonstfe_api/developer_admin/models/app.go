package models

import (
	"encoding/json"
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
)

type AppReq struct {
	UserId   *int    `form:"user_id" db:"user_id" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Sdate    *string `form:"sdate" db:"create_date" query:"gte"`
	Edate    *string `form:"edate" db:"create_date" query:"lte"`
	Status   *int    `form:"status" db:"status" query:"eq"`
}

type App struct {
	Id            *int               `json:"app_id" db:"id"`
	Name          *string            `json:"name" db:"name" validate:"required,gt=0,lte=40" label:"媒体名称"`
	BundleId      *string            `json:"bundle_id" db:"bundle_id" validate:"required,gt=0" label:"程序主包名"`
	Os            *string            `json:"os" db:"os" validate:"required,gt=0" label:"应用操作系统"`
	Category      *string            `json:"category" db:"category" validate:"required,len=4" label:"应用分类"`
	SubCategory   *string            `json:"sub_category" db:"sub_category" validate:"required,len=6" label:"应用分类"`
	KeyWords      *string            `json:"keywords,omitempty" db:"keywords" validate:"required,gt=0,lte=60"  label:"应用关键词"`
	StoreName     *string            `json:"store_name,omitempty" db:"store_name" validate:"required,gt=0" label:"应用商店"`
	StoreUrl      *string            `json:"store_url,omitempty" db:"store_url" validate:"required,regexp=link,lte=300" label:"应用详情页"`
	Describtion   *string            `json:"describtion,omitempty" db:"describtion" validate:"required,gte=40" label:"应用简介"`
	UserId        *int               `json:"user_id,omitempty" db:"user_id"`
	UserEmail     *string            `json:"user_email,omitempty" db:"user_email"`
	ZonstUserId   *int               `json:"zonst_user_id,omitempty" db:"zonst_user_id"`
	CategoryLimit *json.RawMessage   `json:"category_limit,omitempty" db:"category_limit" validate:"required,gt=0" label:"分类限制"`
	Reward        *json.RawMessage   `json:"reward,omitempty"  db:"reward" validate:"required,gt=0" label:"激励配置"`
	Slots         *json.RawMessage   `json:"slots,omitempty" db:"slots" label:"广告位配置"`
	SlotMap       map[string]float32 `json:"slot_map,omitempty" validate:"required,gt=0"`
	CreateDate    *utils.JSONTime    `json:"create_date,omitempty" db:"create_date"`
	Status        *int               `json:"status" db:"status"`
}
type Apps []App

type AppReview struct {
	ReviewType *int `json:"review_type" validate:"required"`
}
type AppReviewBad struct {
	ReviewType *int    `json:"review_type" validate:"required"`
	Title      *string `json:"title" validate:"required,gt=0,lte=100"`
	Content    *string `json:"content" validate:"required,gt=0"`
	GroupName  *string `json:"group_name" validate:"required,gt=0"`
}

// app列表
func GetApps(total *int64, req *AppReq, apps *Apps) error {
	orm := corm.Select(`select id,name,bundle_id,os,category,
	sub_category,create_date,user_id,status,(select email from user_user where id=t1.user_id)
		user_email from app_app as t1 {{sql_where}}`).Req(req).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(apps); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 获取单个app
func GetApp(appId interface{}, app *App) error {
	if err := corm.Db.Get(app, `select * from app_app where id=$1`, appId); err != nil {
		return err
	}
	return nil
}

// 审核app
func ReviewApp(appId interface{}, status int) error {
	if _, err := corm.Db.Exec(`update app_app set status=$1 and id=$2`, status, appId); err != nil {
		return err
	}
	return nil
}
