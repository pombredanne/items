package handlers

import (
	"zonstfe_api/common/my_context"
	"net/http"
	comm "zonstfe_api/common/models"
	"zonstfe_api/developer_admin/models"
	"zonstfe_api/common/utils/password"
	"github.com/go-chi/chi"
	"regexp"
)

// 账户列表
func (c *Handler) Accounts(w http.ResponseWriter, r *http.Request) {
	var req models.AccountReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	accounts := &models.Accounts{}
	var total int64
	if err := models.GetAccounts(my_context.DeveloperRoleId, &total, &req, accounts); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonPage(w, len(*accounts), total, accounts)
}

// 单个账户
func (c *Handler) AccountOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	account := &models.Account{}
	finance := &models.Finance{}
	if err := models.GetAccount(account, userId); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	if err := models.GetFinanceByUserId(finance, userId); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	data := map[string]interface{}{
		"account": account,
		"finance": finance,
	}
	c.JsonBase(w, data)

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

// 提现列表
func (c *Handler) AccountPaymentList(w http.ResponseWriter, r *http.Request) {
	var req models.PaymentReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	payments := &models.Payments{}
	var total int64
	if err := models.GetPayments(&total, &req, payments); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonPage(w, len(*payments), total, payments)

}

// 单个提现记录
func (c *Handler) AccountPaymentOne(w http.ResponseWriter, r *http.Request) {
	paymentId := chi.URLParam(r, "payment_id")
	payment := &models.Payment{}

	if err := models.GetPayment(paymentId, payment); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, payment)

}

// 提现审核
func (c *Handler) AccountPaymentReview(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	paymentId := chi.URLParam(r, "payment_id")
	review, oldPayment := &models.Review{}, &models.Payment{}
	if err := c.Bind(r, review); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}

	if err := models.GetPayment(paymentId, oldPayment); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if *oldPayment.Status == 1 {
		c.JsonError(w, "当前记录已处理", nil)
		return
	}
	status := 0
	if *review.ReviewType == -1 {
		status = -1
	} else {
		status = 1
	}
	if err := models.ReviewPayment(paymentId, my_context.DeveloperRoleId, status, oldPayment); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonBase(w, nil)
	c.LogAction(user["id"], *oldPayment.UserId, paymentId, my_context.ActionModule.Finance, "提现审核", r)

}

// 更改密码
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
	if err := models.GetUser(*req.Email, my_context.DeveloperRoleId, result); err != nil {
		c.JsonError(w, "当前用户不存在", nil)
		return
	}
	if err := models.UpdatePassword(password.SetPassword(*req.PassWord),
		*req.Email, my_context.DeveloperRoleId); err != nil {
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

// 开户
func (c *Handler) AccountOpen(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	accountOpen, reg, finance := &models.AccountOpen{}, &models.Reg{}, &models.Finance{}
	if err := c.Bind(r, accountOpen); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*accountOpen.Account), reg); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*accountOpen.Finance), finance); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	companyName := ""
	dealScale := 0.0
	if models.IfEmailExist(*reg.Email, my_context.DeveloperRoleId) {
		c.JsonError(w, "当前邮箱已存在", nil)
		return
	}

	if *reg.PassWord != *reg.DpassWord {
		c.JsonError(w, "两次密码不相同", nil)
		return
	}
	if reg.CompanyName != nil {
		companyName = *reg.CompanyName
	}
	if *reg.DealType == "share" {
		if reg.DealScale == nil {
			c.JsonError(w, "分成模式必须填写分成比例", nil)
			return
		}
		dealScale = *reg.DealScale
	}
	var userId int64
	if err := models.OpenAccount(&userId, my_context.DeveloperRoleId, dealScale,
		reg, finance, companyName, c.GetIPAdress(r), c.AppSecretKey); err != nil {
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
	if !models.CheckEmailExist(email, my_context.DeveloperRoleId) {
		c.JsonError(w, "该登录邮箱已存在", nil)
		return
	}
	c.JsonBase(w, nil)
}
