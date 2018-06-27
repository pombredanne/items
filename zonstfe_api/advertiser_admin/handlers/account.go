package handlers

import (
	"zonstfe_api/common/my_context"
	comm "zonstfe_api/common/models"
	"zonstfe_api/advertiser_admin/models"
	"net/http"
	"github.com/go-chi/chi"
	"regexp"
	"zonstfe_api/common/utils/password"
)

// 广告主列表
func (c *Handler) Accounts(w http.ResponseWriter, r *http.Request) {
	var req models.AccountReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	accounts := &models.Accounts{}
	var total int64
	if err := models.GetAccounts(my_context.AdvertiserRoleId, &total, &req, accounts); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonPage(w, len(*accounts), total, accounts)
}

// 单个广告主
func (c *Handler) AccountOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	account := &models.Account{}
	if err := models.GetAccount(userId, my_context.AdvertiserRoleId, account); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, account)

}

// 税务信息
func (c *Handler) AccountTaxOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	tax := &models.Tax{}
	if err := models.GetTaxByUserId(userId, my_context.AdvertiserRoleId, tax); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, tax)
}

// 联系信息
func (c *Handler) AccountDeliverOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	deliver := &models.Deliver{}
	if err := models.GetDeliverByUserId(userId, my_context.AdvertiserRoleId, deliver); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, deliver)
}

// 广告主开户
func (c *Handler) AccountOpen(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	accountOpen, reg, deliver := &models.AccountOpen{}, &models.Reg{}, &models.Deliver{}
	if err := c.Bind(r, accountOpen); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*accountOpen.Account), reg); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*accountOpen.Deliver), deliver); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if !models.CheckEmailExist(*reg.Email, my_context.AdvertiserRoleId) {
		c.JsonError(w, "当前邮箱已存在", nil)
		return
	}
	if *reg.PassWord != *reg.DpassWord {
		c.JsonError(w, "两次密码不相同", nil)
		return
	}
	companyName := ""
	if reg.CompanyName != nil {
		companyName = *reg.CompanyName
	}
	var userId int64
	if err := models.OpenAccount(c.Context, &userId, my_context.AdvertiserRoleId, accountOpen,
		reg, deliver, companyName, c.GetIPAdress(r), c.AppSecretKey); err != nil {
		c.JsonError(w, "注册失败", err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], userId, userId, my_context.ActionModule.Account, "开户", r)

}

// 判断开户邮箱是否存在
func (c *Handler) AccountEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if m, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email); !m {
		c.JsonError(w, "邮箱格式错误", nil)
		return
	}
	//验证邮箱是否存在
	if !models.CheckEmailExist(email, my_context.AdvertiserRoleId) {
		c.JsonError(w, "该登录邮箱已存在", nil)
		return
	}
	c.JsonBase(w, nil)
}

// 修改密码
func (c *Handler) AccountPassWordUpdate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req struct {
		Email     *string `json:"email" validate:"required,regexp=email"`
		PassWord  *string `json:"password" validate:"required,gte=6,lte=20"`
		DpassWord *string `json:"dpassword" validate:"required,gte=6,lte=20"`
	}
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, "数据格式错误", err)
		return
	}
	if *req.PassWord != *req.DpassWord {
		c.JsonError(w, "两次密码不相同", nil)
		return
	}
	result := &comm.User{}
	if err := models.GetUser(*req.Email, my_context.AdvertiserRoleId, result); err != nil {
		c.JsonError(w, "当前用户不存在", nil)
		return
	}
	if err := models.UpdatePassword(password.SetPassword(*req.PassWord),
		*req.Email, my_context.AdvertiserRoleId); err != nil {
		c.JsonError(w, "修改失败", err)
		return
	}
	if err := models.UpdateTokenVersion(*result.Id, c.Rd); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonBase(w, nil)
	// 操作记录
	c.Context.LogAction(user["id"], *result.Id, *result.Id, my_context.ActionModule.Account, "密码修改", r)
}

// 修改税务信息
func (c *Handler) AccountTaxUpdate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	taxID := chi.URLParam(r, "tax_id")
	tax, oldTax := &models.Tax{}, &models.Tax{}
	if err := c.Bind(r, tax); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := models.GetTax(taxID, my_context.AdvertiserRoleId, oldTax, ); err != nil {
		c.JsonError(w, my_context.ErrObject, err)
		return
	}
	if err := models.UpdateTax(*oldTax.UserId, my_context.AdvertiserRoleId, tax); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], *oldTax.UserId, oldTax.Id, my_context.ActionModule.Account, "税务信息修改", r)
}

// 联系人修改
func (c *Handler) AccountDeliverUpdate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	deliverId := chi.URLParam(r, "deliver_id")
	deliver, oldDeliver := &models.Deliver{}, &models.Deliver{}
	if err := c.Bind(r, deliver); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := models.GetDeliverById(deliverId, my_context.AdvertiserRoleId, oldDeliver); err != nil {
		c.JsonError(w, my_context.ErrObject, err)
		return
	}
	if err := models.UpdateDeliver(*oldDeliver.UserId, my_context.AdvertiserRoleId, deliver); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], *oldDeliver.UserId, deliverId, my_context.ActionModule.Account, "联系信息修改", r)
}

// 账户信息修改
func (c *Handler) AccountUpdate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	userId := chi.URLParam(r, "user_id")
	account := &models.AccountUpdate{}
	if err := c.Bind(r, account); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	companyName := ""
	if account.CompanyName != nil {
		companyName = *account.CompanyName
	}
	var accountId int64
	if err := models.UpdateAccount(userId, my_context.AdvertiserRoleId, &accountId, companyName, account); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], userId, accountId, my_context.ActionModule.Account, "修改", r)
}

// 充值审核
func (c *Handler) AccountRechargesReview(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	rechargeId := chi.URLParam(r, "recharge_id")
	review, oldRecharge := &models.Review{}, &models.AccountRecharge{}
	if err := c.Bind(r, review); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := models.GetRecharge(rechargeId, my_context.AdvertiserRoleId, oldRecharge); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	// 判断当前记录是否已处理
	if *oldRecharge.Status == 1{
		c.JsonError(w, "当前记录已处理", nil)
		return
	}
	status := 0
	if *review.ReviewType == -1 {
		status = -1
	} else {
		status = 1
	}
	if err := models.ReviewRecharge(rechargeId, my_context.AdvertiserRoleId, status, oldRecharge); err != nil {
		c.JsonError(w, "充值失败", err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], *oldRecharge.UserId, rechargeId, my_context.ActionModule.Finance, "充值审核", r)

}

// 获取充值记录
func (c *Handler) AccountRechargeOne(w http.ResponseWriter, r *http.Request) {
	rechargeId := chi.URLParam(r, "recharge_id")
	recharge := &models.AccountRecharge{}
	if err := models.GetRecharge(rechargeId, my_context.AdvertiserRoleId, recharge); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonBase(w, recharge)
}

// 账户充值
func (c *Handler) AccountRechargeCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	recharge := &models.AccountRecharge{}
	if err := c.Bind(r, recharge); err != nil {
		c.JsonError(w, "数据格式错误", err)
		return
	}
	var rechargeId int64
	if err := models.AddRecharge(&rechargeId, my_context.AdvertiserRoleId, recharge); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}

	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], *recharge.UserId, rechargeId, my_context.ActionModule.Finance, "充值申请", r)
}

// 充值列表
func (c *Handler) AccountRecharges(w http.ResponseWriter, r *http.Request) {
	var req models.RechargeReq
	err := c.Bind(r, &req)
	if err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	recharges := &models.AccountRecharges{}
	var total int64
	if err := models.GetRecharges(my_context.AdvertiserRoleId, &total, &req, recharges); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonPage(w, len(*recharges), total, recharges)

}

// 账户审核
func (c *Handler) AccountReview(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	userId := chi.URLParam(r, "user_id")
	review := &models.Review{}
	if err := c.Bind(r, review); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	status := 0
	if *review.ReviewType == -1 {
		status = -1
	} else {
		status = 1
	}
	if err := models.ReviewAccount(userId, my_context.AdvertiserRoleId, status); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], userId, userId, my_context.ActionModule.Account, "账户审核", r)

}
