package models

import "encoding/json"

type App struct {
	Id            *int             `json:"id" db:"id" `
	AppKey        *string          `json:"app_key" db:"app_key"`
	BundleId      *string          `json:"bundle_id" db:"bundle_id" validate:"required,gt=0" label:"程序主包名"`
	Os            *string          `json:"os" db:"os" validate:"required,gt=0" label:"应用操作系统"`
	Category      *string          `json:"category" db:"category" validate:"required,gt=0" label:"应用分类"`
	SubCategory   *string          `json:"sub_category" db:"sub_category" validate:"required,gt=0" label:"应用分类"`
	UserId        *int             `json:"user_id" db:"user_id"`
	CategoryLimit *json.RawMessage `json:"category_limit" db:"category_limit" label:"分类限制"`
	Reward        *json.RawMessage `json:"reward" db:"reward" label:"激励配置"`
	Slots         *json.RawMessage `json:"slots" db:"slots" label:"广告位配置"`
	DealType      *string          `json:"deal_type" db:"deal_type"`
	DealScale     *float32         `json:"deal_scale" db:"deal_scale"`
	Status        *int             `json:"status" db:"status"`
}

type Apps []App

type CategoryLimit struct {
	Enable *int     `json:"enable" validate:"required,gte=0,lte=1"`
	List   []string `json:"list" validate:"required"`
}

type Identifier struct {
	Device     *string `json:"device" db:"device"`
	Identifier *string `json:"identifier" db:"identifier"`
}
type Identifiers []Identifier
