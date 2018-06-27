package handlers

import (
	"github.com/gin-gonic/gin"
	"adv-admin-api/models"
	"api-libs/rsp"
	"api-libs/option"
)

// 广告系列列表
func CampaignList(c *gin.Context) {
	var req models.CampaignReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	//if req.Sdate == nil && req.Edate == nil {
	//	req.Sdate = func(i string) *string { return &i }(mydate.CurrentDate())
	//	req.Edate = func(i string) *string { return &i }(mydate.CurrentDate())
	//}
	campaigns := &models.Campaigns{}
	var total int64
	if err := models.GetCampaigns(&total, &req, campaigns); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	//reportSum := &models.ReportSum{}
	//if err := models.GetReportCampaignSum(&req, reportSum); err != nil {
	//	c.JsonError(w, my_context.ErrDefault, err)
	//	return
	//}
	c.JSON(rsp.Page(len(*campaigns), total, campaigns))
}

// 查询单个广告活动
func CampaignOne(c *gin.Context) {
	//user := c.GetUser(r)
	campaignId := c.Param("campaignId")
	campaign := &models.Campaign{}
	if err := models.GetCampaign(campaignId, campaign); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base(campaign))
}

// 开启广告活动
func CampaignSwitch(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	campaignId := c.Param("campaignId")
	campaign, balance := &models.Campaign{}, &models.AccountBalance{}
	if err := models.GetCampaign(campaignId, campaign); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	status := 0
	if *campaign.Status == 1 {
		status = 0
		//开启状态则关闭
		if err := models.UpdateCampaignStatus(*campaign.UserId, campaignId, status); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
	} else {
		status = 1
		account := &models.Account{}
		if err := models.GetAccount(*campaign.UserId, option.UserRole["adv"], account); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
		if *account.AccountType == "prepay" {
			//查询账户余额
			if err := models.GetAccountBalance(*campaign.UserId, option.UserRole["adv"], balance); err != nil {
				c.JSON(rsp.Error(rsp.ErrDefault, err))
				return
			}
			if *balance.Balance < 100 {
				c.JSON(rsp.Error("当前余额不足100,请充值"))
				return
			}
		}
		//关闭状态则开启
		if err := models.UpdateCampaignStatus(*campaign.UserId, campaignId, status); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}


	}

	data := map[string]interface{}{
		"status": status,
	}
	c.JSON(rsp.Base(data))
	models.CampaignCache(campaignId, "Campaign修改缓存更新")
	models.LogAction(user["id"], *campaign.UserId, campaignId, models.ActionModule.Campaign, "状态更改", c.ClientIP(), c.Request)
}

// 广告活动修改
func CampaignUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	campaignId := c.Param("campaignId")
	campaign, oldCampaign, targeting, url := &models.Campaign{}, &models.Campaign{}, &models.Targeting{}, &models.Url{}
	if err := rsp.Bind(c.Request, campaign); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := models.GetCampaign(campaignId, oldCampaign); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}

	if err := rsp.BindJson(string(*campaign.Targeting), targeting); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := rsp.BindJson(string(*campaign.Url), url); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	//if err := validTargeting(targeting); err != nil {
	//	c.JsonError(w, content.ErrBadValid, err)
	//	return
	//}
	if err := models.UpdateCampaign(*oldCampaign.UserId, campaignId, campaign); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadExec, err))
		return
	}
	c.JSON(rsp.Base())
	models.CampaignCache(campaignId, "Campaign修改缓存更新")
	models.LogAction(user["id"], *oldCampaign.UserId, campaignId, models.ActionModule.Campaign, "修改", c.ClientIP(), c.Request)

}

// 添加广告活动
func CampaignCreate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	campaign, targeting, url, freq := &models.Campaign{}, &models.Targeting{}, &models.Url{}, &models.CampaignFreq{}
	if err := rsp.Bind(c.Request, campaign); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}

	if err := rsp.BindJson(string(*campaign.Freq), freq); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := rsp.BindJson(string(*campaign.Targeting), targeting); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := rsp.BindJson(string(*campaign.Url), url); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}

	//if err := validTargeting(targeting); err != nil {
	//	c.JsonError(w, content.ErrBadValid, err)
	//	return
	//}
	var campaignId int64
	if err := models.AddCampaign(*campaign.UserId, &campaignId, campaign); err != nil {
		c.JSON(rsp.Error("创建失败", err))
		return
	}
	data := map[string]interface{}{"campaign": campaignId,}
	c.JSON(rsp.Base(data))
	models.LogAction(user["id"], user["id"], campaignId, models.ActionModule.Campaign, "创建", c.ClientIP(), c.Request)

}
