package models

import (
	"zonstfe_api/corm"
	"zonstfe_api/common/utils/password"
	"zonstfe_api/common/utils"
	"strconv"
	"encoding/base64"
)

type UserModel struct {
}

type User struct {
	Id        *int    `db:"id"`
	Email     *string `json:"email" db:"email" validate:"required,gt=0,regexp=email"`
	Password  *string `json:"password" db:"password" validate:"required,gte=6,lte=20"`
	Role      *int    `json:"role" db:"role"`
	IpAddress *string `json:"ip_address" db:"ip_address"`
	Status    *int    `json:"status" db:"status"`
	RegDate   *string `json:"reg_date" db:"reg_date"`
}

type Users []User
type LoginResult struct {
	Id        *int    `db:"id"`
	Email     *string `json:"email" db:"email" validate:"required,gt=0,regexp=email"`
	Password  *string `json:"password" db:"password" validate:"required,gte=6,lte=20"`
	Role      *int    `json:"role" db:"role"`
	IpAddress *string `json:"ip_address" db:"ip_address"`
	Status    *int    `json:"status" db:"status"`
	RegDate   *string `json:"reg_date" db:"reg_date"`
	UserType  *string `json:"user_type" db:"user_type"`
	AppKey    *string `json:"app_key" db:"app_key"`
	DealType  *string `json:"deal_type" db:"deal_type"`
}
type Reg struct {
	Email       *string `json:"email" validate:"required,gt=0,regexp=email" label:"邮箱" error:"请输入正确的邮箱"`
	PassWord    *string `json:"password" validate:"required,gte=6,lte=20" label:"密码" error:"请输入正确的密码"`
	DpassWord   *string `json:"dpassword" validate:"required,gte=6,lte=20" label:"确认密码" error:"两次密码不相同"`
	CompanyName *string `json:"company_name" validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" validate:"required,gt=0,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	UserType    *string `json:"user_type" validate:"required,gt=0" label:"用户类型"`
	Qq          *string `json:"qq" validate:"required,gt=0,regexp=qq" label:"QQ" error:"请输入正确的QQ号"`
	Phone       *string `json:"phone" validate:"required,gt=0,regexp=phone" label:"手机" error:"请输入正确的手机号"`
	Role        *int    `json:"role" label:"角色"`
}

// 验证用户是否存在
func CheckUserExist(email string, roleId int, result *LoginResult) error {
	if err := corm.Db.Get(result, `select * from (select id,password,email from user_user where
	email=$1 AND status!=0 AND role=$2) as t1 LEFT JOIN
	(select user_type,app_key,deal_type,user_id,status from account_account
	) as t2 ON t1."id"=t2.user_id`, email, roleId); err != nil {
		return err
	}
	return nil
}

// 验证用户是否存在
func CheckEmailExist(email string, roleId int) bool {
	userEmail := ""
	if err := corm.Db.Get(userEmail, `select email from user_user where 
	email=$1 and role=$2`, email, roleId); err != nil {
		return true
	}
	return false
}

// 注册
func OpenAccount(userId *int64, roleId int, reg *Reg, companyName, ip, appSecretKey string) error {
	tx := corm.Db.MustBegin()
	// 创建账号信息
	if err := tx.QueryRow(`insert into user_user(email,password,
     type,role,real_name,company_name,ip_address,status)
	 values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`, *reg.Email,
		password.SetPassword(*reg.PassWord), *reg.UserType, roleId,
		*reg.RealName, companyName, ip, 1, ).Scan(userId); err != nil {
		tx.Rollback()
		return err
	}
	appKey, err := utils.NewCBCEncrypter([]byte(strconv.Itoa(int(*userId))), []byte(appSecretKey))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 创建账户信息
	if _, err := tx.Exec(`insert into account_account(user_id,user_type,
	user_role,email,real_name,company_name,
	app_key,qq,phone) values($1,$2,$3,$4,$5,$6,$7,$8,$9)`, *userId, *reg.UserType, roleId, *reg.Email,
		*reg.RealName, companyName, base64.RawStdEncoding.EncodeToString(appKey), *reg.Qq, *reg.Phone); err != nil {
		tx.Rollback()
		return err
	}
	// 创建邮寄信息
	if _, err := tx.Exec(`insert into account_deliver(user_id,telephone,
		user_role,email) values($1,$2,$3,$4)`, *userId, *reg.Phone, roleId, *reg.Email); err != nil {
		tx.Rollback()
		return err
	}
	// 创建账户余额
	if _, err := tx.Exec(`insert into 
		account_balance(user_id,user_role,balance,app_key) values($1,$2,$3,$4)`,
		*userId, roleId, 0, base64.RawStdEncoding.EncodeToString(appKey)); err != nil {
		tx.Rollback()
		return err
	}
	// 创建财务信息
	if _, err := tx.Exec(`insert into account_finance(user_id,user_role) values($1,$2)`, *userId, roleId); err != nil {
		tx.Rollback()
		return err
	}
	// 创建税务信息
	if _, err := tx.Exec(`insert into account_tax(user_id,user_role,
		company_name) values($1,$2,$3)`, *userId, roleId, companyName); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
