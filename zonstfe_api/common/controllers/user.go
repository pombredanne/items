package controllers

import (
	"net/http"
	"zonstfe_api/common/models"
	"zonstfe_api/common/my_context"
	"zonstfe_api/common/utils/password"
	"regexp"
	"github.com/garyburd/redigo/redis"
	"github.com/go-chi/jwtauth"
	"strconv"
	"time"
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/go-chi/chi"
)

type UserController struct {
	*my_context.Context
}

func NewUserController(content *my_context.Context) *UserController {
	return &UserController{content}
}

// 登录
func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	user, result := &models.User{Role: &c.RoleID}, &models.LoginResult{Role: &c.RoleID}
	if err := c.Bind(r, user); err != nil {
		c.JsonError(w, "数据异常", err)
		return
	}
	// 验证用户是否存在
	if err := models.CheckUserExist(*user.Email, c.RoleID, result); err != nil {
		c.JsonError(w, "当前用户不存在", err)
		return
	}
	// 验证密码是否相同
	if user.Password == nil || result.Password == nil || !password.CheckPassword(*user.Password, *result.Password) {
		c.JsonError(w, "密码错误", nil)
		return
	}
	// 当前账号未审核
	//if *result.Status == 0 {
	//	c.JsonError(w, "当前账号未审核", nil)
	//	return
	//}
	//生成token
	client := c.Rd.Get()
	defer client.Close()
	version, _ := redis.String(client.Do("GET", "login_user:"+strconv.Itoa(*result.Id)))
	if version == "" {
		version = uuid.NewV4().String()
		if _, err := client.Do("SET", "login_user:"+strconv.Itoa(*result.Id), version); err != nil {
			c.JsonError(w, "服务错误", err)
		}

	}
	_, tokenString, err := c.TokenAuth.Encode(jwtauth.Claims{"id": *result.Id, "app_key": *result.AppKey,
		"timestamp": time.Now().Unix(),
		"deal_type": *result.DealType,
		"exp": 24 * time.Hour,
		"version": version})
	if err != nil {
		c.JsonError(w, "服务错误", err)
	}
	//if err := c.CreateCookieToken(w, r, result); err != nil {
	//	c.JsonError(w, -1, err)
	//	return
	//}

	data := map[string]interface{}{
		"email":     *result.Email,
		"app_key":   *result.AppKey,
		"deal_type": *result.DealType,
		"token":     tokenString,
	}
	c.JsonBase(w, data)

}
func (c UserController) SignupEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if m, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email); !m {
		c.JsonError(w, "邮箱格式错误", nil)
		return
	}
	//验证邮箱是否存在
	if !models.CheckEmailExist(email, c.RoleID) {
		c.JsonError(w, "该登录邮箱已存在", nil)
		return
	}
	c.JsonBase(w, nil)
}

//注册
func (c UserController) Signup(w http.ResponseWriter, r *http.Request) {
	reg := &models.Reg{}
	if err := c.Bind(r, reg); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if !models.CheckEmailExist(*reg.Email, c.RoleID) {
		c.JsonError(w, "注册邮箱已存在", nil)
		return
	}
	var userId int64
	companyName := ""
	if *reg.PassWord != *reg.DpassWord {
		c.JsonError(w, "两次密码不相同", nil)
		return

	}
	if reg.CompanyName == nil {
		companyName = ""
	} else {
		companyName = *reg.CompanyName
	}
	if err := models.OpenAccount(&userId, c.RoleID, reg, companyName, c.GetIPAdress(r), c.AppSecretKey); err != nil {
		c.JsonError(w, "注册失败", err)
		return
	}
	c.JsonBase(w, nil)
	c.Context.LogAction(userId, userId, userId, my_context.ActionModule.Account, "注册", r)

}

func (c UserController) Logout(w http.ResponseWriter, r *http.Request) {
	current_user := c.GetUser(r)
	client := c.Rd.Get()
	defer client.Close()
	// 注销更新版本
	if _, err := client.Do("SET", fmt.Sprintf("login_user:%v", current_user["id"]), uuid.NewV4().String()); err != nil {
		c.JsonError(w, "服务错误", err)
		return
	}
	c.JsonBase(w, nil)

}

//func (c UserController) CreateCookieToken(w http.ResponseWriter, r *http.Request, user *models.LoginResult) error {
//	str_json, err := json.Marshal(user)
//	if err != nil {
//		return err
//	}
//	max_age := 3 * 24 * 60 * 60
//	//expire := time.Now().AddDate(20, 0, 0)
//	token, err := utils.NewCBCEncrypter([]byte(strconv.Itoa(*user.Id)+"$"+strconv.Itoa(rand.Intn(999999))), []byte(c.SecretKey))
//	if err != nil {
//		return err
//	}
//	cookie := http.Cookie{
//		Name:  "session",
//		Value: base64.RawStdEncoding.EncodeToString(token),
//		//Domain:   r.Host, //设置域名
//		Path:     "/",
//		Expires:  time.Now().Add(2 * 24 * time.Hour),
//		HttpOnly: true,
//	}
//	http.SetCookie(w, &cookie)
//	// 通过redis 存储相关信息
//	client := c.Rd.Get()
//	key := "login_user:" + strconv.Itoa(*user.Id)
//	client.Do("EXPIRE", key, max_age)
//	client.Do("SET", key, str_json)
//	defer client.Close()
//	return nil
//}
