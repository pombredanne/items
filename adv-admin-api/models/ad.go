package models

import (
	"corm"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"api-libs/rsp"
)

type AdReq struct {
	AdSize       *string `form:"ad_size" db:"width::text || 'x' || height" query:"eq"`
	AdType       *string `form:"ad_type" db:"ad_type" query:"eq"`
	Status       *int    `form:"status" db:"status" query:"eq"`
	Sdate        *string `form:"sdate" db:"create_date" query:"gte"`
	Edate        *string `form:"edate" db:"create_date" query:"lte"`
	CampaignId   *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	CampaignNmae *string `form:"campaign_name" db:"campaign_name" query:"like"`
	UserEmail    *string `form:"user_email" db:"user_email" query:"like"`
	UserId       *int    `form:"user_id" db:"user_id" query:"eq"`
	Name         *string `form:"name" db:"name" query:"like"`
	Page         *uint   `form:"page"`
	PageSize     *uint   `form:"page_size"`
}

type VideoCreative struct {
	Video       *string  `json:"video" validate:"required,gt=0,regexp=link"`
	CampaignId  *int     `json:"campaign_id" db:"campaign_id" validate:"required,gt=0"`
	Title       *string  `json:"title" db:"title" validate:"required,gt=0,lte=18"`
	BiddingMax  *float32 `json:"bidding_max" db:"bidding_max" validate:"required,gt=0,lte=100"`
	BiddingMin  *float32 `json:"bidding_min" db:"bidding_min" validate:"required,gt=0,lte=100"`
	BiddingType *string  `json:"bidding_type" db:"bidding_type" validate:"required,gt=0"`
	Image       *string  `json:"image" db:"image" validate:"required,regexp=link"`
	Name        *string  `json:"name" db:"name" validate:"required,gt=0,lte=18"`
	Duration    *float64 `json:"duration" validate:"required,gt=0"`
	//UserId      *int     `json:"user_id" validate:"required,gt=0"`
}

type ImageCreative struct {
	CampaignId  *int     `json:"campaign_id" db:"campaign_id" validate:"required,gt=0"`
	Title       *string  `json:"title" db:"title" validate:"required,gt=0,lte=18"`
	BiddingMax  *float32 `json:"bidding_max" db:"bidding_max" validate:"required,gt=0,lte=100"`
	BiddingMin  *float32 `json:"bidding_min" db:"bidding_min" validate:"required,gt=0,lte=100"`
	BiddingType *string  `json:"bidding_type" db:"bidding_type" validate:"required,gt=0"`
	Image       *string  `json:"image" db:"image" validate:"required,regexp=link"`
	Name        *string  `json:"name" db:"name" validate:"required,gt=0,lte=18"`
	//UserId      *int     `json:"user_id" validate:"required,gt=0"`
}

// 添加视频广告
func AddVideoAd(userId interface{}, adId *int64, creative *AdCreative, video *VideoCreative) error {
	creativeStr, _ := json.Marshal(map[string]string{
		"image": *video.Image,
		"video": *video.Video,
	})
	if err := corm.Db.QueryRow(`insert into campaign_ad(name,ad_type,ad_size,
		ol,duration,campaign_id,width,height,user_id,creative_id,creative,bidding_max,bidding_min,bidding_type,title,
	status,switch_status) 
		values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17) RETURNING id`, *video.Name,
		*creative.AdType, "", 0,
		*video.Duration, *video.CampaignId, *creative.Width,
		*creative.Height, userId, *creative.Id, creativeStr,
		*video.BiddingMax, *video.BiddingMin, *video.BiddingType, *video.Title, 1, 1).Scan(adId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 添加图片广告
func AddImageAd(userId interface{}, adId *int64, creative *AdCreative, image *ImageCreative) error {
	creativeStr, _ := json.Marshal(map[string]string{
		"image": *image.Image,
	})
	if err := corm.Db.QueryRow(`insert into campaign_ad(name,ad_type,ad_size,
		ol,duration,campaign_id,width,height,user_id,creative_id,
		creative,bidding_max,bidding_min,bidding_type,title,status,switch_status) 
		values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17) RETURNING id`, *image.Name,
		*creative.AdType, "", 0,
		0, *image.CampaignId, *creative.Width,
		*creative.Height, userId, *creative.Id, creativeStr,
		*image.BiddingMax, *image.BiddingMin, *image.BiddingType, *image.Title, 1, 1).Scan(adId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 广告修改
func UpdateImageAdWithStatus(userId, AdId interface{}, status, switchStatus int, creative string, image *ImageCreative) error {
	if _, err := corm.Db.Exec(`update campaign_ad set status=$1,creative=$2,
	name=$3,title=$4,bidding_max=$5,bidding_min=$6,switch_status=$7
	 where id=$8 and user_id=$9`, status, creative,
		*image.Name, *image.Title, *image.BiddingMax,
		*image.BiddingMin, switchStatus, AdId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

func UpdateVideoAdWithStatus(userId, AdId interface{}, status, switchStatus int, creative string, video *VideoCreative) error {
	if _, err := corm.Db.Exec(`update campaign_ad set status=$1,creative=$2,
	name=$3,title=$4,bidding_max=$5,bidding_min=$6,duration=$7,switch_status=$8 where id=$9 and user_id=$10`, status, creative,
		*video.Name, *video.Title, *video.BiddingMax,
		*video.BiddingMin, *video.Duration,
		switchStatus, AdId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 广告修改
func UpdateImageAd(userId, AdId interface{}, creative string, image *ImageCreative) error {
	if _, err := corm.Db.Exec(`update campaign_ad set creative=$1,
	name=$2,title=$3,bidding_max=$4,bidding_min=$5,where id=$6 and user_id=$7`, creative,
		*image.Name, *image.Title, *image.BiddingMax,
		*image.BiddingMin, AdId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

func UpdateVideoAd(userId, AdId interface{}, creative string, video *VideoCreative) error {
	if _, err := corm.Db.Exec(`update campaign_ad set creative=$1,
	name=$2,title=$3,bidding_max=$4,bidding_min=$5,
	duration=$6 where id=$7 and user_id=$8`, creative,
		*video.Name, *video.Title, *video.BiddingMax,
		*video.BiddingMin, *video.Duration, AdId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改广告活动状态
func UpdateAdStatus(userId, adId interface{}, status int) error {
	if _, err := corm.Db.Exec(`update campaign_ad set 
	switch_status=$1 where id=$2 and user_id=$3`, status, adId, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 广告批量审核
func ReviewBatchAd(adList []int, status, switch_status int) error {
	query, args, err := sqlx.In("update campaign_ad set status=?,switch_status=? where id in (?)",
		status, switch_status, adList)
	if err != nil {
		return rsp.HandlerError(err)
	}
	query = corm.Db.Rebind(query)
	if _, err := corm.Db.Exec(query, args...); err != nil {
		return rsp.HandlerError(err)
	}
	return nil

}

type AdIncomeReq struct {
	UserId      *int    `form:"user_id" db:"user_id" query:"eq"`
	BiddingType *string `form:"bidding_type" db:"bidding_type" query:"eq"`
	CampaignId  *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	Sdate       *string `form:"sdate" db:"report_date" query:"gte"`
	Edate       *string `form:"edate" db:"report_date" query:"lte"`
	Page        *uint   `form:"page"`
	PageSize    *uint   `form:"page_size"`
}

type AdIncome struct {
	Clk          *int64           `json:"clk" db:"clk"`                     // 点击
	Imp          *int64           `json:"imp" db:"imp"`                     // 展示
	Cost         *float64         `json:"cost" db:"cost"`                   // 消耗
	Money        *float64         `json:"money" db:"money"`                 // 导入消耗
	Adid         *int             `json:"ad_id" db:"ad_id"`                 // 广告ID
	CampaignId   *int             `json:"campaign_id" db:"campaign_id"`     //广告系列ID
	Creative     *json.RawMessage `json:"creative" db:"creative"`           //素材
	UserId       *int             `json:"user_id" db:"user_id"`             //用户ID
	BiddingType  *string          `json:"bidding_type" db:"bidding_type"`   //计费类型
	Name         *string          `json:"name" db:"name"`                   //广告名称
	CampaignName *string          `json:"campaign_name" db:"campaign_name"` //广告系列名称
	BiddingMin   *float32         `json:"bidding_min" db:"bidding_min"`     //最低价
	BiddingMax   *float32         `json:"bidding_max" db:"bidding_max"`     //最高价
}

type AdIncomes [] AdIncome

type AdIncomeDetailReq struct {
	UserId     *int    `form:"user_id" db:"user_id" query:"eq"`
	CampaignId *int    `form:"campaign_id" db:"campaign_id" query:"eq"`
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
}

type AdIncomeDetail struct {
	Clk        *int64   `json:"clk" db:"clk"` // 点击
	Imp        *int64   `json:"imp" db:"imp"` // 展示
	Uclk       *int64   `json:"uclk" db:"uclk"`
	Cost       *float64 `json:"cost" db:"cost"`               // 消耗
	Money      *float64 `json:"money" db:"money"`             // 导入消耗
	Adid       *int     `json:"ad_id" db:"ad_id"`             // 广告ID
	ReportDate *string  `json:"report_date" db:"report_date"` // 时间
}

type AdIncomeDetails []AdIncomeDetail

type IncomeImport struct {
	Adid *int    `json:"ad_id" validate:"required,gt=0"`
	Data *string `json:"data" validate:"required"`
}

func GetAdIncomes(total *int64, req *AdIncomeReq, data *AdIncomes) error {
	orm := corm.Select(`select * from (select  t1.*,COALESCE(t2.money,0) money 
		from (select  SUM(clk) as clk,SUM(imp) as imp,
		 SUM("cost") as "cost",ad_id from report_base {{t1_where}} GROUP BY ad_id) as t1 LEFT JOIN 
		(select SUM(money) as money,ad_id  FROM 
		ad_income {{t2_where}} GROUP BY ad_id) as t2 ON t1.ad_id=t2.ad_id) as t3
		LEFT JOIN (SELECT id,name,campaign_id,creative,user_id,bidding_type,
		bidding_min,bidding_max,(select name from campaign_campaign 
		WHERE id=campaign_id) as campaign_name FROM campaign_ad ) as 
		t4 ON t3.ad_id=t4.id {{t3_where}} ORDER BY "cost" DESC,money DESC`).EqualWhere("t1_where",
		[][]interface{}{
			{"ad_id", "!=", 0},
			{"user_id", "=", req.UserId},
			{"campaign_id", "=", req.CampaignId},
			{"report_date", ">=", req.Sdate},
			{"report_date", "<=", req.Edate},
		}).EqualWhere("t2_where",
		[][]interface{}{
			{"user_id", "=", req.UserId},
			{"campaign_id", "=", req.CampaignId},
			{"report_date", ">=", req.Sdate},
			{"report_date", "<=", req.Edate},
		}).EqualWhere("t3_where", [][]interface{}{
		{"bidding_type", "=", req.BiddingType},
	})
	if err := orm.Paginate(req.Page, req.PageSize).Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

func GetAdIncomeDetail(total *int64, req *AdIncomeDetailReq, data *AdIncomeDetails) error {
	orm := corm.Select(`SELECT t1.*,COALESCE(t2.money,0) as money from 
	(select ad_id,report_date,"cost",clk,imp,uclk FROM 
	report_base {{sql_where}}) as t1 
	LEFT JOIN (SELECT ad_id,money,report_date from ad_income) 
	as t2 ON t1.ad_id=t2.ad_id and t1.report_date=t2.report_date ORDER BY report_date DESC`).Req(req)
	if err := orm.Paginate(req.Page, req.PageSize).Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}
func AdIncomeImport(ad *CampaignAd, data map[string]float64) error {
	tx := corm.Db.MustBegin()
	for k, v := range data {
		if _, err := tx.Exec(`insert into ad_income(ad_id,campaign_id,
			user_id,report_date,money) values($1,$2,$3,$4,$5)`, *ad.Id, *ad.CampaignId, *ad.UserId, k, v); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return rsp.HandlerError(err)
	}
	return nil
}
