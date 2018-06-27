package handlers

import (
	"github.com/gin-gonic/gin"
	"adv-admin-api/models"
	"api-libs/rsp"
	"api-libs/password"
	"adv-admin-api/config"
)

// 登录
func Login(c *gin.Context) {
	user, result := &models.User{}, &models.User{}
	if err := rsp.Bind(c.Request, user); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadFormat, err))
		return
	}
	// 验证用户是否存在
	if err := models.CheckUserExist(*user.Email, config.Conf.RoleId, result); err != nil {
		c.JSON(rsp.Error("当前用户不存在", err))
		return
	}
	// 验证密码是否相同
	if user.Password == nil || result.Password == nil || !password.CheckPassword(*user.Password, *result.Password) {
		c.JSON(rsp.Error("密码错误"))
		return
	}

	// 如果当前环境为测试环境只允许测试账号登录
	if config.Conf.EnvModel == "testing" && !*result.Test {
		c.JSON(rsp.Error("登录异常,当前环境为测试"))
		return
	}

	// 当前账号未审核
	//if *result.Status == 0 {
	//	c.JsonError(w, "当前账号未审核")
	//	return
	//}
	//生成token
	version, err := models.GetTokenVersion(*result.Id)
	if err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	tokenString, err := models.CreateToken(*result.Id, version, config.Conf.SigningKey)
	if err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
		return
	}
	data := map[string]interface{}{
		"user_id": *result.Id,
		"email":   *result.Email,
		"token":   tokenString,
	}
	c.JSON(rsp.Base(data))

}

// 注销
func Logout(c *gin.Context) {
	user := c.MustGet("user").(map[string]interface{})
	// 注销更新版本
	if err := models.UpdateTokenVersion(user["id"]); err != nil {
		c.JSON(rsp.Error(rsp.ErrDefault, err))
	}

	c.JSON(rsp.Base())

}
