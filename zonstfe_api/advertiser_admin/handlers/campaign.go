package handlers

import (
	"zonstfe_api/common/my_context"
	"zonstfe_api/corm"
	"zonstfe_api/advertiser_admin/models"
	"github.com/go-chi/chi"
	"net/http"
	"zonstfe_api/common/utils/mydate"
	"strconv"
)

// 广告系列列表
func (c *Handler) CampaignList(w http.ResponseWriter, r *http.Request) {
	var req models.CampaignReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, "请求参数错误", err)
		return
	}
	if req.Sdate == nil && req.Edate == nil {
		req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
		req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	}
	campaigns := &models.Campaigns{}
	var total int64
	if err := models.GetCampaigns(&total, &req, campaigns); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	reportSum := &models.ReportSum{}
	if err := models.GetReportCampaignSum(&req, reportSum); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonSumPage(w, len(*campaigns), total, campaigns, reportSum)
}

func (c *Handler) AdList(w http.ResponseWriter, r *http.Request) {
	var req models.AdReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, "请求参数错误", err)
		return
	}
	ads := &models.CampaignAds{}

	var total int64
	if err := models.GetAds(&total, &req, ads); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonPage(w, len(*ads), total, ads)

}

// 广告审核
func (c *Handler) CampaignAdReview(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	adId := chi.URLParam(r, "ad_id")
	review, campaignAd := &models.Review{}, &models.CampaignAd{}
	if err := c.Bind(r, review); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := models.GetCampaignAd(adId, campaignAd); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	adID, _ := strconv.Atoi(adId)
	if *review.ReviewType == -1 {
		if err := models.ReviewBatchAd([]int{adID}, -1); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}

	} else {
		if err := models.ReviewBatchAd([]int{adID}, 1); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
	}

	c.JsonBase(w, nil)
	c.CampaignCache(*campaignAd.CampaignId, "Campaign缓存更新")
	c.Context.LogAction(user["id"], *campaignAd.UserId, adId, my_context.ActionModule.Ad, "审核", r)
}

// 广告批量审核
func (c *Handler) CampaignAdReviewBatch(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	review_batch, campaign_ads := models.ReviewBatch{}, &models.CampaignAds{}
	if err := c.Bind(r, &review_batch); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := corm.Select(`select * from campaign_ad where id=:id)
	`).Where("id", "in", *review_batch.List).Load(campaign_ads); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	if len(*review_batch.List) > 0 {
		if *review_batch.ReviewType == -1 {
			if err := models.ReviewBatchAd(*review_batch.List, -1); err != nil {
				c.JsonError(w, my_context.ErrDefault, err)
				return
			}
		} else {
			if err := models.ReviewBatchAd(*review_batch.List, 1); err != nil {
				c.JsonError(w, my_context.ErrDefault, err)
				return
			}
		}
	}
	c.JsonBase(w, nil)
	for _, item := range *campaign_ads {
		c.CampaignCache(*item.CampaignId, "后台批量审核AD缓存更新")
	}
	c.Context.LogAction(user["id"], 0, 0, my_context.ActionModule.Ad, "审核", r)

}
