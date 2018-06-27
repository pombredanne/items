package models

import (
	"corm"
	"api-libs/my-field"
	"api-libs/rsp"
)

type SegmentReq struct {
	Name     *string `form:"name" db:"name" query:"like"`
	UserId   *int    `form:"user_id" db:"user_id" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
}
type Segment struct {
	Id         *int            `josn:"segment_id" db:"id"`
	Name       *string         `json:"name" db:"name" validate:"required,gt=0"`
	UserId     *int            `json:"user_id" db:"user_id"`
	Type       *int            `json:"type" validate:"required,gt=0"`
	Uv         *int            `josn:"uv" db:"uv"`
	PkgPath    *string         `json:"pkg_path" db:"pkg_path" validate:"required,gt=0,regexp=link"`
	CreateDate *my_field.JSONTime `json:"create_date" db:"create_date"`
	UpdateDate *my_field.JSONTime `json:"update_date" db:"update_date"`
}

type Segments []Segment

// 人群包列表
func GetSegments(total *int64, req *SegmentReq, segments *Segments) error {
	orm := corm.Select(`select * from campaign_segment {{sql_where}}`).Req(req).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(segments); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}
