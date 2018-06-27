package handlers

import (
	"zonstfe_api/common/my_context"
	"zonstfe_api/developer/models"
	"zonstfe_api/common/utils/password"
	comm "zonstfe_api/common/models"
	"net/http"
	"github.com/go-chi/chi"
)

// 账户信息
func (c *Handler) Account(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	account := &models.Account{}
	if err := models.GetAccount(user["id"], c.RoleID, account); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonBase(w, account)
}

// 账户修改
func (c *Handler) AccountEdit(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	account := &models.Account{}
	if err := c.Bind(r, account); err != nil {
		c.JsonError(w, "数据格式错误", err)
		return
	}
	if err := models.UpdateAccount(user["id"], c.RoleID, account); err != nil {
		c.JsonError(w, "修改失败", err)
		return
	}
	c.Context.LogAction(user["id"], user["id"], user["id"], my_context.ActionModule.Account, "修改", r)

}

// 账户余额
func (c *Handler) AccountBalance(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	balance := &models.AccountBalance{}
	if err := models.GetAccountBalance(user["id"], c.RoleID, balance); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonBase(w, balance)
}

// 密码修改
func (c *Handler) AccountPassWordEdit(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var reg struct {
		OldPassWord *string `json:"old_password" validate:"required,gte=6,lte=20"`
		PassWord    *string `json:"password" validate:"required,gte=6,lte=20"`
		DpassWord   *string `json:"dpassword" validate:"required,gte=6,lte=20"`
	}

	if err := c.Bind(r, &reg); err != nil {
		c.JsonError(w, "数据格式错误", err)
		return
	}

	result := &comm.User{}
	if err := models.CheckUserExist(user["id"], result); err != nil {
		c.JsonError(w, "当前用户不存在", err)
		return
	}
	// 验证密码是否相同
	if !password.CheckPassword(*reg.OldPassWord, *result.Password) {
		c.JsonError(w, "原始密码错误", nil)
		return
	}
	if *reg.PassWord != *reg.DpassWord {
		c.JsonError(w, "两次密码不相同", nil)
		return
	}
	if err := models.UpdatePassWord(user["id"], password.SetPassword(*reg.PassWord)); err != nil {
		c.JsonError(w, "修改失败", err)
		return
	}
	if err := models.UpdateTokenVersion(user["id"], c.Rd); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], user["id"], user["id"], my_context.ActionModule.Account, "密码修改", r)

}

// 财务信息
func (c *Handler) Finance(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	finance := &models.Finance{}
	if err := models.GetFinance(user["id"], finance); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, finance)
}

// 财务信息修改
func (c *Handler) FinanceEdit(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	finance := &models.Finance{}
	financeId := chi.URLParam(r, "financeId")
	if err := c.Bind(r, finance); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := models.UpdateFinance(user["id"], financeId, finance); err != nil {
		c.JsonError(w, my_context.ErrBadExec, err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], user["id"], financeId, my_context.ActionModule.Finance, "修改", r)
}

// 提现申请
func (c *Handler) PaymentCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	balance := &models.AccountBalance{}
	payment, payments := &models.Payment{}, &models.Payments{}

	if err := c.Bind(r, payment); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	// 查询余额
	if err := models.GetAccountBalance(user["id"], c.RoleID, balance); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if *balance.Balance < 1000 {
		c.JsonError(w, "当前余额不足1000 提现申请失败", nil)
		return
	}
	// 先判断是否包含未处理申请
	models.GetPaymentByStatus(user["id"], 0, payments)
	if len(*payments) > 0 {
		c.JsonError(w, "还有待处理申请,请等待处理后再申请", nil)
		return
	}
	var paymentId int64
	if err := models.AddPayment(&paymentId, []interface{}{
		user["id"], c.RoleID, *balance.Balance, *payment.OrderMoney, 0,
	}); err != nil {

	}
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], user["id"], paymentId,
		my_context.ActionModule.Finance, "提现申请", r)

}

// 提现列表
func (c *Handler) Payments(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.PaymentReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	payments := &models.Payments{}
	var total int64
	if err := models.GetPayments(user["id"], &total, &req, payments); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonPage(w, len(*payments), total, payments)

}
