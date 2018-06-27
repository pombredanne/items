package handlers

import (
	comm "zonstfe_api/common/models"
	"zonstfe_api/advertiser/models"
	"net/http"
	"zonstfe_api/common/utils/password"
	"zonstfe_api/common/my_context"
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
func (c *Handler) AccountUpdate(w http.ResponseWriter, r *http.Request) {
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
	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], user["id"], user["id"], my_context.ActionModule.Account, "修改", r)

}

// 账户密码修改
func (c *Handler) AccountPassWordUpdate(w http.ResponseWriter, r *http.Request) {
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

// 账户充值
func (c *Handler) AccountRechargesCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	recharge := &models.AccountRecharge{}
	if err := c.Bind(r, recharge); err != nil {
		c.JsonError(w, "数据格式错误", err)
		return
	}

	recharges := &models.AccountRecharges{}
	// 查询是否还有待审核数据
	if err := models.GetRechargeReview(user["id"], recharges); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
	}
	if len(*recharges) > 0 {
		c.JsonError(w, "您还有待审核充值请求未处理,请勿重复提交", nil)
		return
	}
	var rechargeId int64
	// 新增充值记录
	if err := models.AddRecharge(user["id"], c.RoleID, &rechargeId, recharge); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	c.JsonBase(w, nil)
	c.Context.LogAction(user["id"], user["id"], rechargeId, my_context.ActionModule.Finance, "充值申请", r)
}

// 获取充值记录列表
func (c *Handler) AccountRecharges(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.RechargeReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	recharges := &models.AccountRecharges{}
	var total int64
	if err := models.GetRecharges(user["id"], &total, &req, recharges); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonPage(w, len(*recharges), total, recharges)
}

// 账户余额
func (c *Handler) AccountBalance(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	balance := &models.AccountBalance{}
	if err := models.GetAccountBalance(user["id"], c.RoleID, balance); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	c.JsonBase(w, balance)
}
