package handlers

import (
	"strconv"
	"strings"
	"corm"
	"github.com/gin-gonic/gin"
	"adv-admin-api/models"
	"api-libs/rsp"
	"encoding/json"
	"fmt"
)

// 广告列表
func AdList(c *gin.Context) {
	var req models.AdReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	ads := &models.CampaignAds{}

	var total int64
	if err := models.GetAds(&total, &req, ads); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Page(len(*ads), total, ads))

}

// 广告创建
func AdCreate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	adCreative := &models.AdCreative{}
	creativeId := c.Param("creativeId")
	if err := models.GetCreative(creativeId, adCreative); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	var adId int64
	data := make(map[string]interface{}, 0)
	switch *adCreative.AdType {
	case "video":
		cretive_obj := &models.VideoCreative{}
		if err := rsp.Bind(c.Request, cretive_obj); err != nil {
			c.JSON(rsp.Error(rsp.ErrBadFormat, err))
			return
		}
		// 获取campaign 信息
		campaign := &models.Campaign{}
		if err := models.GetCampaign(*cretive_obj.CampaignId, campaign); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}

		if err := models.AddVideoAd(*campaign.UserId, &adId, adCreative, cretive_obj); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
		data = map[string]interface{}{
			"campaign_id": *cretive_obj.CampaignId,
			"ad_id":       adId,
		}
	case "graphic":
		cretive_obj := &models.ImageCreative{}
		if err := rsp.Bind(c.Request, cretive_obj); err != nil {
			c.JSON(rsp.Error(rsp.ErrBadFormat, err))
			return
		}
		// 获取campaign 信息
		campaign := &models.Campaign{}
		if err := models.GetCampaign(*cretive_obj.CampaignId, campaign); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
		if err := models.AddImageAd(*campaign.UserId, &adId, adCreative, cretive_obj); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
		data = map[string]interface{}{
			"campaign_id": *cretive_obj.CampaignId,
			"ad_id":       adId,
		}

	}
	c.JSON(rsp.Base(data))
	models.CampaignCache(data["campaign_id"], "Campaign AD 修改缓存更新")
	models.LogAction(user["id"], user["id"], adId, models.ActionModule.Ad,
		"创建", c.ClientIP(), c.Request)

}

// 广告修改
func AdUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	ad_id := c.Param("adId")
	old_campaign_ad := &models.CampaignAd{}
	if err := models.GetCampaignAd(ad_id, old_campaign_ad); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	ad_name := ""
	switch *old_campaign_ad.AdType {
	case "video":
		cretive_obj := &models.VideoCreative{}
		if err := rsp.Bind(c.Request, cretive_obj); err != nil {
			c.JSON(rsp.Error(rsp.ErrBadFormat, err))
			return
		}
		creativeStr, _ := json.Marshal(map[string]string{
			"image": *cretive_obj.Image,
			"video": *cretive_obj.Video,
		})
		ad_name = *cretive_obj.Name
		// 判断素材是否有变更
		if strings.Replace(string(*old_campaign_ad.Creative), " ", "", -1) != strings.Replace(string(creativeStr), " ", "", -1) {
			if err := models.UpdateVideoAdWithStatus(*old_campaign_ad.UserId,
				ad_id, 1, 1, string(creativeStr), cretive_obj); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		} else {
			if err := models.UpdateVideoAd(*old_campaign_ad.UserId, ad_id, string(creativeStr), cretive_obj); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		}
	case "graphic":
		cretive_obj := &models.ImageCreative{}
		if err := rsp.Bind(c.Request, cretive_obj); err != nil {
			c.JSON(rsp.Error(rsp.ErrBadFormat, err))
			return
		}
		ad_name = *cretive_obj.Name
		creativeStr, _ := json.Marshal(map[string]string{
			"image": *cretive_obj.Image,
		})
		// 判断素材是否有变更
		if strings.Replace(string(*old_campaign_ad.Creative), " ", "", -1) != strings.Replace(string(creativeStr), " ", "", -1) {
			if err := models.UpdateImageAdWithStatus(*old_campaign_ad.UserId,
				ad_id, 1, 1, string(creativeStr), cretive_obj); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		} else {
			if err := models.UpdateImageAd(*old_campaign_ad.UserId, ad_id, string(creativeStr), cretive_obj); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		}
	}

	data := map[string]interface{}{
		"campaign_id": *old_campaign_ad.CampaignId,
		"ad_id":       ad_id,
	}
	c.JSON(rsp.Base(data))
	models.CampaignCache(*old_campaign_ad.CampaignId, "Campaign AD 修改缓存更新")
	models.LogAction(user["id"], *old_campaign_ad.UserId, ad_id, models.ActionModule.Ad,
		ad_name+"修改", c.ClientIP(), c.Request)

}

//
func AdOne(c *gin.Context) {
	adId := c.Param("adId")
	campaign_ad := &models.CampaignAd{}
	if err := models.GetCampaignAd(adId, campaign_ad); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base(campaign_ad))
}

// 广告批量审核
func CampaignAdReviewBatch(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	review_batch, campaign_ads := models.ReviewBatch{}, &models.CampaignAds{}
	if err := rsp.Bind(c.Request, &review_batch); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := corm.Select(`select * from campaign_ad where id=:id)
	`).Where("id", "in", *review_batch.List).Load(campaign_ads); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}

	if len(*review_batch.List) > 0 {
		if *review_batch.ReviewType == -1 {
			if err := models.ReviewBatchAd(*review_batch.List, -1, 0); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		} else {
			if err := models.ReviewBatchAd(*review_batch.List, 1, 1); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		}
	}
	c.JSON(rsp.Base())
	for _, item := range *campaign_ads {
		models.CampaignCache(*item.CampaignId, "后台批量审核AD缓存更新")
	}
	models.LogAction(user["id"], 0, 0, models.ActionModule.Ad, "审核", c.ClientIP(), c.Request)

}

// 广告审核
func CampaignAdReview(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	adId := c.Param("adId")
	review, campaignAd := &models.Review{}, &models.CampaignAd{}
	if err := rsp.Bind(c.Request, review); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := models.GetCampaignAd(adId, campaignAd); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	adID, _ := strconv.Atoi(adId)
	if *review.ReviewType == -1 {
		if err := models.ReviewBatchAd([]int{adID}, -1, 0); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}

	} else {
		if err := models.ReviewBatchAd([]int{adID}, 1, 1); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
	}
	c.JSON(rsp.Base())
	models.CampaignCache(*campaignAd.CampaignId, "Campaign缓存更新")
	models.LogAction(user["id"], *campaignAd.UserId, adId, models.ActionModule.Ad, "审核", c.ClientIP(), c.Request)

}

// 广告开关
func AdSwitch(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	adId := c.Param("adId")
	ad := &models.CampaignAd{}
	if err := models.GetCampaignAd(adId, ad); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	switch_status := 0
	if *ad.Status == 1 {
		if *ad.SwitchStatus == 1 {
			switch_status = 0
			//开启状态则关闭
			if err := models.UpdateAdStatus(*ad.UserId, adId, switch_status); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		} else {
			switch_status = 1
			//开启状态则关闭
			if err := models.UpdateAdStatus(*ad.UserId, adId, switch_status); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
		}
	} else {
		c.JSON(rsp.Error("操作失败 当前广告处于待审核状态"))
		return
	}

	data := map[string]interface{}{
		"switch_status": switch_status,
	}
	c.JSON(rsp.Base(data))
	models.CampaignCache(*ad.CampaignId, "Campaign修改缓存更新")
	models.LogAction(user["id"], *ad.UserId, adId, models.ActionModule.Ad, "上线状态更改", c.ClientIP(), c.Request)
}

type AdIncormImport struct {
	AdId *int    `json:"ad_id"`
	Data *string `json:"data"`
}

// 查询广告 收益
func AdCost(c *gin.Context) {
	var req models.AdIncomeReq
	data := &models.AdIncomes{}
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	var total int64
	if err := models.GetAdIncomes(&total, &req, data); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*data), total, data))
}

// 查询广告收益 明细
func AdCostDetail(c *gin.Context) {
	var req models.AdIncomeDetailReq
	data := &models.AdIncomeDetails{}
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	var total int64
	if err := models.GetAdIncomeDetail(&total, &req, data); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*data), total, data))
}

// 收益导入
func AdCostImport(c *gin.Context) {
	incomeImport, ad := &models.IncomeImport{}, &models.CampaignAd{}
	if err := rsp.Bind(c.Request, incomeImport); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := models.GetCampaignAd(*incomeImport.Adid, ad); err != nil {
		c.JSON(rsp.Error(rsp.ErrObject, err))
		return
	}
	lines := strings.Split(*incomeImport.Data, "\n")
	dataMap := make(map[string]float64)
	for _, line := range lines {
		data := strings.Fields(line)
		if len(data) != 2 {
			c.JSON(rsp.Error("填写数据格式错误"))
			return
		}
		i, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			c.JSON(rsp.Error("填写数据格式错误"))
			return
		}
		dataMap[data[0]] = i
	}

	if err := models.AdIncomeImport(ad, dataMap); err != nil {
		if strings.Contains(fmt.Sprintf("%v", err), "duplicate key value violates unique constraint") {
			c.JSON(rsp.Error("导入失败 存在重复数据", err))
		} else {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
		}
		return
	}
	c.JSON(rsp.Base())
}
