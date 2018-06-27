package models

import (
	"zonstfe_api/common/utils"
	"encoding/json"
	"zonstfe_api/corm"
)

type AppReq struct {
	Name     *string `form:"name" db:"name" query:"eq"`
	BundleId *string `form:"bundle_id" db:"bundle_id" query:"eq"`
	Os       *int    `form:"os" db:"os" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
}

type App struct {
	Id          *int    `json:"app_id" db:"id"`
	Name        *string `json:"name" db:"name" validate:"required,gt=0,lte=40" label:"媒体名称"`
	BundleId    *string `json:"bundle_id" db:"bundle_id" validate:"required,gt=0,regexp=bundle" label:"程序主包名"`
	Os          *string `json:"os" db:"os" validate:"required,gt=0" label:"应用操作系统"`
	Category    *string `json:"category" db:"category" validate:"required,len=4" label:"应用分类"`
	SubCategory *string `json:"sub_category" db:"sub_category" validate:"required,len=6" label:"应用分类"`
	KeyWords    *string `json:"keywords,omitempty" db:"keywords" validate:"required"  label:"应用关键词"`
	//KeyWordArray  *[]string        `json:"keyword_array,omitempty" validate:"required"`
	StoreName     *string            `json:"store_name,omitempty" db:"store_name" validate:"required,gt=0" label:"应用商店"`
	StoreUrl      *string            `json:"store_url,omitempty" db:"store_url" validate:"required,regexp=link,lte=300" label:"应用详情页"`
	Describtion   *string            `json:"describtion,omitempty" db:"describtion" validate:"required" label:"应用简介"`
	UserId        *int               `json:"user_id,omitempty" db:"user_id"`
	ZonstUserId   *int               `json:"zonst_user_id,omitempty" db:"zonst_user_id"`
	CategoryLimit *json.RawMessage   `json:"category_limit,omitempty" db:"category_limit" validate:"required,gt=0" label:"分类限制"`
	Reward        *json.RawMessage   `json:"reward,omitempty"  db:"reward" validate:"required,gt=0" label:"激励配置"`
	Slots         *json.RawMessage   `json:"slots,omitempty" db:"slots" label:"广告位配置"`
	SlotMap       map[string]float32 `json:"slot_map,omitempty" validate:"required,gt=0"`
	CreateDate    *utils.JSONTime    `json:"create_date,omitempty" db:"create_date"`
	Status        *int               `json:"status" db:"status"`
}
type Apps []App

type CategoryLimit struct {
	Enable *int     `json:"enable" validate:"required,gte=0,lte=1"`
	List   []string `json:"list" validate:"required"`
}

type AppEdit struct {
	Name        *string `json:"name" db:"name" validate:"required,gt=0,lte=40" label:"应用名称"`
	Category    *string `json:"category" db:"category" validate:"required,len=4" label:"应用分类"`
	SubCategory *string `json:"sub_category" db:"sub_category" validate:"required,len=6" label:"应用分类"`
	KeyWords    *string `json:"keywords,omitempty" db:"keywords"  validate:"required" label:"应用关键词"`

	//KeyWordArray *[]string `json:"keyword_array,omitempty" validate:"required"`
	//KeyWords      *json.RawMessage `json:"keywords" db:"keywords" validate:"required" label:"应用关键词"`
	Describtion   *string          `json:"describtion,omitempty" db:"describtion" validate:"required,gte=40"`
	CategoryLimit *json.RawMessage `json:"category_limit,omitempty" db:"category_limit" validate:"required,gt=0" label:"分类限制"`
	Reward        *json.RawMessage `json:"reward,omitempty" db:"reward" validate:"required,gt=0" label:"激励配置"`
	//Slots         *json.RawMessage `json:"slots,omitempty" db:"slots" validate:"nonzero" label:"广告位配置"`
	SlotMap    map[string]float32 `json:"slot_map,omitempty" validate:"required,gt=0"`
	CreateDate *utils.JSONTime    `json:"create_date" db:"create_date"`
	Status     *int               `json:"status" db:"status"`
}

//type Slot struct {
//	SlotId      *string  `json:"slot_id"`
//	Bidding     *float32 `json:"bidding" validate:"required"`
//	BiddingType *string  `json:"bidding_type" validate:"required"`
//}

//type BiddingSlots []Slot

type Slots []string

type Reward struct {
	Enable      *int    `json:"enable" validate:"required,gte=0,lte=1"`
	Currency    *string `json:"currency" validate:"lte=300"`
	Amount      *int    `json:"amount" validate:"gte=0"`
	CallBack    *int    `json:"callback" validate:"gte=0,lte=1"`
	CallBackUrl *string `json:"callback_url" validate:"omitempty,regexp=link"`
}

func GetRewardJson() string {
	info_map := map[string]interface{}{
		"enable":       0,
		"currency":     "",
		"amount":       0,
		"callback":     0,
		"callback_url": "",
	}
	info_json, _ := json.Marshal(info_map)
	return string(info_json)
}

// app列表
func GetApps(userId interface{}, total *int64, req *AppReq, apps *Apps) error {
	orm := corm.Select(`select id,name,bundle_id,os,category,create_date,status
	from app_app {{sql_where}} order by create_date desc`).Req(&req).Where(map[string]interface{}{
		"user_id": userId,
	}).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(apps); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 单个APP
func GetApp(userId, appId interface{}, app *App) error {
	if err := corm.Db.Get(app, `select id,name,os,category,keywords,describtion,bundle_id,
	slots,reward,category_limit,status from app_app where user_id=$1 and id=$2`, userId, appId); err != nil {
		return err
	}
	return nil
}

// 添加App
func AddApp(userId interface{}, appId *int64, reward, categoryLimit, slotJson string, app *App) error {
	if err := corm.Db.QueryRow(`insert into app_app(name,bundle_id,
		os,category,sub_category,keywords,store_name,store_url,describtion,
		user_id,reward,category_limit,slots) values() `, *app.Name, *app.BundleId,
		*app.Os, *app.Category, *app.SubCategory, *app.KeyWords, *app.StoreName, *app.StoreUrl,
		*app.Describtion, userId, reward, categoryLimit, slotJson).Scan(appId); err != nil {
		return err
	}
	return nil
}

// 修改App
func UpdateApp(userId, appId interface{}, reward, categoryLimit, slotJson string, app *AppEdit) error {
	if _, err := corm.Db.Exec(`update app_app set keywords=$1,category=$2,
	sub_category=$3,describtion=$4,name=$5,reward=$6,category_limit=$6,
	slots=$7 where id=$8 and user_id=$9`, *app.KeyWords, *app.Category, *app.SubCategory,
		*app.Describtion, *app.Name, reward, categoryLimit, slotJson, appId, userId); err != nil {
		return err
	}
	return nil
}

// 验证APP 是否存在
func CheckAppExist(userId interface{}, bundleId, os string, apps *Apps) error {
	if err := corm.Db.Select(apps, `select id from app_app where user_id=$1  and bundle_id=$2 and os=$3`,
		userId, bundleId, os); err != nil {
		return err
	}
	return nil
}
