package models

import (
	"corm"
	"api-libs/rsp"
)

type AdCreative struct {
	Id *int `json:"creative_id" db:"id"`
	//Name        *string `json:"name" db:"name"`
	AdType *string `json:"ad_type" db:"ad_type"`
	//AdSize      *string `json:"ad_size" db:"ad_size"`
	//Material    *string `json:"material" db:"material"`
	//Description *string `json:"description" db:"description"`
	//OL          *int    `json:"ol" db:"ol"`
	Width  *int `json:"width" db:"width"`
	Height *int `json:"height" db:"height"`
}
type AdCreativeList []AdCreative

// 创意列表
func GetCreatives(creatives *AdCreativeList) error {
	if err := corm.Db.Select(creatives, `select * from ad_creative where status=1`); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 单个创意
func GetCreative(creativeId interface{}, creative *AdCreative) error {
	if err := corm.Db.Get(creative, `select * from ad_creative where id=$1`, creativeId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

