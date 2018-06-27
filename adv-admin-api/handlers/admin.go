package handlers

import (
	"github.com/gin-gonic/gin"
	"api-libs/rsp"
	"adv-admin-api/models"
	"api-libs/password"
	"adv-admin-api/config"
)

func AdminSignup(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	accountOpen, reg := &models.AccountOpen{}, &models.Reg{}
	if err := rsp.Bind(c.Request, reg); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	if !models.CheckEmailExist(*reg.Email, config.Conf.RoleId) {
		c.JSON(rsp.Error("注册邮箱已存在"))
		return
	}
	var userId int64
	companyName := ""
	if *reg.PassWord != *reg.DpassWord {
		c.JSON(rsp.Error("两次密码不相同"))
		return

	}
	if reg.CompanyName == nil {
		companyName = ""
	} else {
		companyName = *reg.CompanyName
	}
	if err := models.OpenAccount(&userId, config.Conf.RoleId,
		accountOpen, reg, companyName, c.ClientIP()); err != nil {
		c.JSON(rsp.Error("注册失败", err))
		return
	}
	c.JSON(rsp.Base())
	models.LogAction(user["id"], userId, userId, models.ActionModule.Account, "注册", c.ClientIP(), c.Request)
}

func AdminPassWordUpdate(c *gin.Context) {
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
	if err := models.GetUser(*req.Email, config.Conf.RoleId, result); err != nil {
		c.JSON(rsp.Error("当前用户不存在", err))
		return
	}
	if err := models.UpdatePassword(password.SetPassword(*req.PassWord),
		*req.Email, config.Conf.RoleId); err != nil {
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
