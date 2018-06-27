package handlers

import (
	"regexp"
	"github.com/gin-gonic/gin"
	"adv-admin-api/models"
	"api-libs/rsp"
	"api-libs/option"
	"api-libs/password"
)

// 广告主列表
func Accounts(c *gin.Context) {
	var req models.AccountReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	accounts := &models.Accounts{}
	var total int64
	if err := models.GetAccounts(option.UserRole["adv"], &total, &req, accounts); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*accounts), total, accounts))
}

// 单个广告主
func AccountOne(c *gin.Context) {
	userId := c.Param("userId")
	account := &models.Account{}
	if err := models.GetAccount(userId, option.UserRole["adv"], account); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(account))
}

// 税务信息
func GetAccountTax(c *gin.Context) {
	userId := c.Param("userId")
	taxs := &models.Taxs{}
	if err := models.GetTaxByUserId(userId, taxs); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(taxs))
}

// 联系信息
func GetAccountDeliver(c *gin.Context) {
	userId := c.Param("userId")
	delivers := &models.Delivers{}
	if err := models.GetDeliverByUserId(userId, delivers); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(delivers))
}

// 广告主开户
func AccountOpen(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	accountOpen, reg := &models.AccountOpen{}, &models.Reg{}
	if err := rsp.Bind(c.Request, accountOpen); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := rsp.BindJson(string(*accountOpen.Account), reg); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	//if err := rsp.BindJson(string(*accountOpen.Deliver), deliver); err != nil {
	//	c.JSON(rsp.Error(rsp.ErrBadFormat, err)
	//	return
	//}
	if !models.CheckEmailExist(*reg.Email, option.UserRole["adv"]) {
		c.JSON(rsp.Error("当前邮箱已存在"))
		return
	}
	if *reg.PassWord != *reg.DpassWord {
		c.JSON(rsp.Error("两次密码不相同"))
		return
	}
	companyName := ""
	if reg.CompanyName != nil {
		companyName = *reg.CompanyName
		if *reg.UserType == "person" {
			companyName = *reg.RealName
		}
	}
	var userId int64
	if err := models.OpenAccount(&userId, option.UserRole["adv"], accountOpen,
		reg, companyName, c.ClientIP()); err != nil {
		c.JSON(rsp.Error("注册失败", err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, userId,
		models.ActionModule.Account, "开户", c.ClientIP(), c.Request)

}

// 判断开户邮箱是否存在
func AccountEmail(c *gin.Context) {
	email := c.Param("email")
	if m, err := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email); !m {
		c.JSON(rsp.Error("邮箱格式错误", err))
		return
	}
	//验证邮箱是否存在
	if !models.CheckEmailExist(email, option.UserRole["adv"]) {
		c.JSON(rsp.Error("该登录邮箱已存在"))
		return
	}
	c.JSON(rsp.Base())
}

// 修改密码
func AccountPassWordUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	var req struct {
		Email     *string `json:"email" validate:"required,regexp=email"`
		PassWord  *string `json:"password" validate:"required,gte=6,lte=20"`
		DpassWord *string `json:"dpassword" validate:"required,gte=6,lte=20"`
	}
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error("数据格式错误", err))
		return
	}
	if *req.PassWord != *req.DpassWord {
		c.JSON(rsp.Error("两次密码不相同"))
		return
	}
	result := &models.User{}
	if err := models.GetUser(*req.Email, option.UserRole["adv"], result); err != nil {
		c.JSON(rsp.Error("当前用户不存在", err))
		return
	}
	if err := models.UpdatePassword(password.SetPassword(*req.PassWord),
		*req.Email, option.UserRole["adv"]); err != nil {
		c.JSON(rsp.Error("修改失败", err))
		return
	}
	if err := models.UpdateTokenVersion(*result.Id); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	// 操作记录
	models.LogAction(user["id"], *result.Id, *result.Id, models.ActionModule.Account, "密码修改", c.ClientIP(), c.Request)
}

// 修改税务信息
//func AccountTaxUpdate(c *gin.Context) {
//	user := c.MustGet("user").(map[string]interface{})
//	taxID := c.Param("taxId")
//	tax, oldTax := &models.Tax{}, &models.Tax{}
//	if err := rsp.Bind(c.Request, tax); err != nil {
//		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
//		return
//	}
//	if err := models.GetTax(taxID, option.UserRole["adv"], oldTax, ); err != nil {
//		c.JSON(rsp.Error(rsp.ErrObject, err))
//		return
//	}
//	if err := models.UpdateTax(*oldTax.UserId, option.UserRole["adv"], tax); err != nil {
//		c.JSON(rsp.Error(rsp.ErrDefault, err))
//		return
//	}
//	c.JSON(rsp.Base())
//	models.LogAction(user["id"], *oldTax.UserId, oldTax.Id, models.ActionModule.Account, "税务信息修改", c.ClientIP(), c.Request)
//}

func AccountTaxUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userId := c.Param("userId")
	tax := &models.Tax{}
	if err := rsp.Bind(c.Request, tax); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	var taxId int64
	if err := models.UpdateTax(userId, option.UserRole["adv"], &taxId, tax); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, taxId, models.ActionModule.Account, "税务信息修改", c.ClientIP(), c.Request)
}

// 获取财务信息
func GetFinance(c *gin.Context) {
	userId := c.Param("userId")
	finances := &models.Finances{}
	if err := models.GetFinance(userId, finances); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(finances))
}

// 修改财务信息
func AccountFinanceUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userId := c.Param("userId")
	finance := &models.Finance{}
	if err := rsp.Bind(c.Request, finance); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	var financeId int64
	if err := models.UpdateFinance(userId, option.UserRole["adv"], &financeId, finance); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, financeId, models.ActionModule.Account, "财务信息修改", c.ClientIP(), c.Request)
}

// 获取行业信息
func GetIndustry(c *gin.Context) {
	var req models.IndustryReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}

	industrys := &models.Industrys{}
	var total int64
	if err := models.GetIndustry(&total, &req, industrys); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*industrys), total, industrys))

}

// 修改行业信息
func IndustryUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	industryId := c.Param("industryId")
	industry := &models.Industry{}
	if err := rsp.Bind(c.Request, industry); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := models.UpdateIndustry(industryId, industry); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], *industry.UserId, industryId, models.ActionModule.Account, "行业信息修改", c.ClientIP(), c.Request)
}

// 新增行业信息
func IndustryCreate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	industry := &models.Industry{}
	if err := rsp.Bind(c.Request, industry); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}

	var industryId int64
	if err := models.CreateIndustry(&industryId, industry); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], *industry.UserId, industryId, models.ActionModule.Account, "行业信息新增", c.ClientIP(), c.Request)

}

// 获取资质信息
func GetQualification(c *gin.Context) {
	userId := c.Param("userId")
	qf := &models.Qualifications{}
	if err := models.GetQualification(userId, qf); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(qf))
}

// 资质信息修改新增
func QualificationUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userId := c.Param("userId")
	company_qf, person_qf, account := &models.CompanyQualification{}, &models.PersonQualification{}, &models.Account{}
	if err := models.GetAccount(userId, option.UserRole["adv"], account); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	var qfId int64
	if *account.UserType == "person" {
		if err := rsp.Bind(c.Request, person_qf); err != nil {
			c.JSON(rsp.Error(rsp.ErrBadFormat, err))
			return
		}
		if err := models.UpdatePersonQualification(userId, &qfId, person_qf); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}

	} else {
		if err := rsp.Bind(c.Request, company_qf); err != nil {
			c.JSON(rsp.Error(rsp.ErrBadFormat, err))
			return
		}
		if err := models.UpdateCompanyQualification(userId, &qfId, company_qf); err != nil {
			c.JSON(rsp.Error(rsp.ErrDefault, err))
			return
		}
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, qfId, models.ActionModule.Account, "资质信息修改", c.ClientIP(), c.Request)
}

func AccountDeliverUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userId := c.Param("userId")
	deliver := &models.Deliver{}
	if err := rsp.Bind(c.Request, deliver); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	var deliverId int64
	if err := models.CreateDeliver(userId, option.UserRole["adv"], &deliverId, deliver); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}

	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, deliverId, models.ActionModule.Account, "联系信息修改", c.ClientIP(), c.Request)
}

// 联系人修改
//func AccountDeliverUpdate(c *gin.Context) {
//	user := c.MustGet("user").(map[string]interface{})
//	deliverId := c.Param("deliverId")
//	deliver, oldDeliver := &models.Deliver{}, &models.Deliver{}
//	if err := rsp.Bind(c.Request, deliver); err != nil {
//		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
//		return
//	}
//	if err := models.GetDeliverById(deliverId, option.UserRole["adv"], oldDeliver); err != nil {
//		c.JSON(rsp.Error(rsp.ErrObject, err))
//		return
//	}
//	if err := models.UpdateDeliver(*oldDeliver.UserId, option.UserRole["adv"], deliver); err != nil {
//		c.JSON(rsp.Error(rsp.ErrDefault, err))
//		return
//	}
//
//	c.JSON(rsp.Base())
//	models.LogAction(user["id"], *oldDeliver.UserId, deliverId, models.ActionModule.Account, "联系信息修改", c.ClientIP(), c.Request)
//}

// 账户信息修改
func AccountUpdate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userId := c.Param("userId")
	account := &models.AccountUpdate{}
	if err := rsp.Bind(c.Request, account); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	companyName := ""
	if account.CompanyName != nil {
		companyName = *account.CompanyName
	}
	var accountId int64
	if err := models.UpdateAccount(userId, option.UserRole["adv"], &accountId, companyName, account); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, accountId, models.ActionModule.Account, "修改", c.ClientIP(), c.Request)
}

// 充值审核
func AccountRechargesReview(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	rechargeId := c.Param("rechargeId")
	review, oldRecharge := &models.Review{}, &models.AccountRecharge{}
	if err := rsp.Bind(c.Request, review); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if err := models.GetRecharge(rechargeId, option.UserRole["adv"], oldRecharge); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	// 判断当前记录是否已处理
	if *oldRecharge.Status == 1 {
		c.JSON(rsp.Error("当前记录已处理"))
		return
	}
	status := 0
	if *review.ReviewType == -1 {
		status = -1
	} else {
		status = 1
	}
	if err := models.ReviewRecharge(rechargeId, option.UserRole["adv"], status, oldRecharge); err != nil {
		c.JSON(rsp.Error("充值失败", err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], *oldRecharge.UserId, rechargeId, models.ActionModule.Finance, "充值审核", c.ClientIP(), c.Request)

}

// 获取充值记录
func AccountRechargeOne(c *gin.Context) {
	rechargeId := c.Param("rechargeId")
	recharge := &models.AccountRecharge{}
	if err := models.GetRecharge(rechargeId, option.UserRole["adv"], recharge); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(recharge))
}

// 充值记录消耗
func AccountRechargeCost(c *gin.Context) {
	var req models.RechargeCostReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	costs := &models.RechargeCosts{}
	var total int64
	if err := models.GetRechargeCost(&total, &req, costs); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*costs), total, costs))

}

// 账户充值
func AccountRechargeCreate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	recharge := &models.AccountRecharge{}
	if err := rsp.Bind(c.Request, recharge); err != nil {
		c.JSON(rsp.Error("数据格式错误", err))
		return
	}
	var rechargeId int64
	if err := models.AddRecharge(&rechargeId, option.UserRole["adv"], recharge); err != nil {
		c.JSON(rsp.Error("服务错误", err))
		return
	}

	c.JSON(rsp.Base())
	models.LogAction(user["id"], *recharge.UserId, rechargeId, models.ActionModule.Finance, "充值申请", c.ClientIP(), c.Request)
}

// 充值列表
func AccountRecharges(c *gin.Context) {
	var req models.RechargeReq
	err := rsp.Bind(c.Request, &req)
	if err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	recharges := &models.AccountRecharges{}
	var total int64
	if err := models.GetRecharges(option.UserRole["adv"], &total, &req, recharges); err != nil {
		c.JSON(rsp.Error("服务错误", err))
		return
	}
	c.JSON(rsp.Page(len(*recharges), total, recharges))
}

// 账户审核
func AccountReview(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	userId := c.Param("userId")
	review := &models.Review{}
	if err := rsp.Bind(c.Request, review); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	status := 0
	if *review.ReviewType == -1 {
		status = -1
	} else {
		status = 1
	}
	if err := models.ReviewAccount(userId, option.UserRole["adv"], status); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, userId, models.ActionModule.Account, "账户审核", c.ClientIP(), c.Request)
}

// 合同创建
func AccountContractCreate(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	contract := &models.AccountContract{}
	if err := rsp.Bind(c.Request, contract); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}

	var contractId int64

	if err := models.CreateAccountContract(&contractId, contract); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user, *contract.UserId, contractId, models.ActionModule.Account, "合同创建", c.ClientIP(), c.Request)
}

// 合同列表
func GetAccountContract(c *gin.Context) {
	var req models.AccountContractReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	datas := &models.AccountContracts{}
	var total int64
	if err := models.GetAccountContract(&total, &req, datas); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*datas), total, datas))

}

// 获取用户流水
func GetAccountCost(c *gin.Context) {
	var req models.AccountCostReq
	if err := rsp.Bind(c.Request, &req); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	costs := &models.AccountCosts{}
	var total int64
	if err := models.GetAccountCost(&total, &req, costs); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Page(len(*costs), total, costs))
}