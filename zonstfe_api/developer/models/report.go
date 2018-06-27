package models

import (
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
)

type ReportAppSlotReq struct {
	Sdate    *string `form:"sdate" db:"report_date"  query:"gte"`
	Edate    *string `form:"edate" db:"report_date" query:"lte"`
	BundleId *string `form:"bundle_id" db:"bundle_id" query:"eq"`
	SlotId   *string `form:"slot_id" db:"slot_id" query:"eq"`
	Os       *string `form:"os" db:"os" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
}

type ReportAppRewardReq struct {
	Sdate    *string `form:"sdate" db:"report_date" query:"gte"`
	Edate    *string `form:"edate" db:"report_date" query:"lte"`
	BundleId *string `form:"bundle_id" db:"bundle_id" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Os       *string `form:"os" db:"os" query:"eq"`
}

type ReportAppReward struct {
	BundleId   *string         `json:"bundle_id,omitempty" db:"bundle_id"`
	ReportDate *utils.JSONTime `json:"report_date,omitempty" db:"report_date"`
	Os         *string         `json:"os" db:"os"`
	Amount     *int            `json:"amount" db:"amount"`
	Reward     *int            `json:"reward" db:"reward"`
	Uv         *int            `json:"uv" db:"uv"`
	Imp        *int            `json:"imp" db:"imp"`
}
type ReportAppRewards []ReportAppReward

type ReportAppSlot struct {
	BundleId   *string         `json:"bundle_id,omitempty" db:"bundle_id"`
	ReportDate *utils.JSONTime `json:"report_date,omitempty" db:"report_date"`
	//Detail     *json.RawMessage `json:"detail" db:"detail"`
	Imp    *int    `json:"imp" db:"imp"`
	Clk    *int    `json:"clk" db:"clk"`
	SlotId *string `json:"slot_id,omitempty" db:"slot_id"`
	Os     *string `json:"os,omitempty" db:"os"`
}

type ReportAppSlots []ReportAppSlot

type ReportSum struct {
	Win    *int             `json:"win,omitempty" db:"win"`
	Imp    *int             `json:"imp,omitempty" db:"imp"`
	Eimp   *int             `json:"eimp,omitempty" db:"eimp"`
	Clk    *int             `json:"clk,omitempty" db:"clk"`
	Cost   *utils.JSONFloat `json:"cost,omitempty" db:"cost"`
	Amount *int             `json:"amount,omitempty" db:"amount"`
	Reward *int             `json:"reward,omitempty" db:"reward"`
	Uv     *int             `json:"uv,omitempty" db:"uv"`
}

// APP广告位报表
func GetReportAppSlot(appKey interface{}, total *int64, req *ReportAppSlotReq, data *ReportAppSlots) error {
	sql := `select COALESCE(SUM(imp),0) imp,COALESCE(SUM(clk),0) clk,os,bundle_id,slot_id,report_date
	from report_app_slot {{sql_where}} group by os,bundle_id,slot_id,
	report_date order by os,bundle_id,slot_id,report_date`
	orm := corm.Select(sql).Req(req).Where(map[string]interface{}{
		"app_key": appKey,
	})
	if err := orm.Paginate(req.Page, req.PageSize).Loads(data); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// APP广告位报表汇总
func GetReportAppSlotSum(appKey interface{}, req *ReportAppSlotReq, reportSum *ReportSum) error {
	if err := corm.Select(`select COALESCE(SUM(imp),0) imp,COALESCE(SUM(clk),0) clk
		from report_app_slot {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"app_key": appKey,
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}


// APP激励报表
func GetReportAppReward(appKey interface{}, total *int64, req *ReportAppRewardReq, data *ReportAppRewards) error {
	sql := `select COALESCE(SUM(imp),0) imp,
	COALESCE(SUM(amount),0) amount,COALESCE(SUM(reward),0) reward,COALESCE(SUM(uv),0) uv,os,bundle_id,report_date
	from report_app_reward {{sql_where}} group by os,bundle_id,report_date order by os,bundle_id,report_date`
	orm := corm.Select(sql).Req(req).Where(map[string]interface{}{
		"app_key": appKey,
	})
	if err := orm.Paginate(req.Page, req.PageSize).Loads(data); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// APP激励报表汇总
func GetReportAppRewardSum(appKey interface{}, req *ReportAppRewardReq, reportSum *ReportSum) error {
	if err := corm.Select(`select  COALESCE(SUM(imp),0) imp,
	COALESCE(SUM(amount),0) amount,COALESCE(SUM(reward),0) reward,COALESCE(SUM(uv),0) uv
		from report_app_reward {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"app_key": appKey,
	}).Load(reportSum); err != nil {
		return err
	}
	return nil
}