package models

import (
	"zonstfe_api/common/utils"
	"encoding/json"
	"zonstfe_api/corm"
	comm "zonstfe_api/common/models"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/satori/go.uuid"
	"encoding/base64"
	"zonstfe_api/common/utils/password"
	"strconv"
	"zonstfe_api/common/utils/mydate"
)

type AccountReq struct {
	Email    *string `form:"email" db:"email" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Sdate    *string `form:"sdate" db:"create_date" query:"gte"`
	Edate    *string `form:"edate" db:"create_date" query:"lte"`
	Status   *int    `form:"status" db:"status" query:"eq"`
}

type PaymentReq struct {
	UserId   *string `form:"user_id" db:"user_id" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Sdate    *string `form:"sdate" db:"apply_date" query:"gte"`
	Edate    *string `form:"edate" db:"apply_date" query:"lte"`
	Status   *int    `form:"status" db:"status" query:"eq"`
}

type Account struct {
	UserId      *int    `json:"user_id" db:"user_id" label:"用户ID"`
	UserType    *string `json:"user_type" db:"user_type" label:"用户类型"`
	Email       *string `json:"email" db:"email" label:"邮箱"`
	CompanyName *string `json:"company_name" db:"company_name"  validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" db:"real_name" validate:"required,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	AppCount    *int    `json:"app_count,omitempty" db:"app_count"`
	//AppKey      *string `json:"app_key" db:"app_key" label:"APP_KEY"`
	DealType   *string         `json:"deal_type" db:"deal_type" label:"盈利模式"`
	DealScale  *float32        `json:"deal_scale" db:"deal_scale" label:"分成比例"`
	Qq         *string         `json:"qq" db:"qq" validate:"required,regexp=qq"`
	Phone      *string         `json:"phone" db:"phone" validate:"required,regexp=phone"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
	Status     *int            `json:"status" db:"status"`
}

type Accounts []Account

type Review struct {
	ReviewType *int `json:"review_type" validate:"required"`
}

type PaymentReview struct {
	ReviewType *int `json:"review_type" validate:"required"`
}

type AccountReviewBad struct {
	ReviewType *int    `json:"review_type" validate:"required"`
	Title      *string `json:"title" validate:"required,gt=0,lte=100"`
	Content    *string `json:"content" validate:"required,gt=0"`
	GroupName  *string `json:"group_name" validate:"required,gt=0"`
}

type PaymentReviewBad struct {
	ReviewType *int    `json:"review_type" validate:"required"`
	Title      *string `json:"title" validate:"required,gt=0,lte=100"`
	Content    *string `json:"content" validate:"required,gt=0"`
	GroupName  *string `json:"group_name" validate:"required,gt=0"`
}

type AccountOpen struct {
	Account *json.RawMessage `json:"account" validate:"required"`
	Finance *json.RawMessage `json:"finance" validate:"required"`
}

type Reg struct {
	Email       *string  `json:"email" validate:"required,gt=0,regexp=email" label:"邮箱" error:"请输入正确的邮箱"`
	PassWord    *string  `json:"password" validate:"required,gte=6,lte=20" label:"密码" error:"请输入正确的密码"`
	DpassWord   *string  `json:"dpassword" validate:"required,gte=6,lte=20" label:"确认密码" error:"两次密码不相同"`
	CompanyName *string  `json:"company_name" validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string  `json:"real_name" validate:"required,gt=0,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	UserType    *string  `json:"user_type" validate:"required,gt=0,option=user_type" label:"用户类型"`
	DealType    *string  `json:"deal_type" validate:"required,gt=0,option=deal_type"`
	DealScale   *float64 `json:"deal_scale" validate:"omitempty,gte=0,lte=1"`
	Qq          *string  `json:"qq" validate:"required,gt=0,regexp=qq" label:"QQ" error:"请输入正确的QQ号"`
	Phone       *string  `json:"phone" validate:"required,gt=0,regexp=phone" label:"手机" error:"请输入正确的手机号"`
}

type Finance struct {
	Id           *int    `json:"finance_id" db:"id"`
	AccountName  *string `json:"account_name" db:"account_name" validate:"required,gt=0"`
	BankName     *string `json:"bank_name" db:"bank_name" validate:"required,gt=0"`         //开户行
	BankNo       *string `json:"bank_no" db:"bank_no" validate:"required,gte=16,lte=19"`    //银行账号
	BankProvince *string `json:"bank_province" db:"bank_province" validate:"required,gt=0"` //开户省
	BankCity     *string `json:"bank_city" db:"bank_city" validate:"required,gt=0"`
	SubBranch    *string `json:"sub_branch" db:"sub_branch" validate:"required,gt=0,lte=300"` //开户地支行地址
}

type Payment struct {
	Id         *int            `json:"payment_id" db:"id"`
	UserId     *int            `json:"user_id" db:"user_id"`
	Balance    *float32        `json:"balance" db:"balance"`
	UserEmail  *string         `json:"user_email" db:"user_email"`
	ApplyDate  *utils.JSONTime `json:"apply_date" db:"apply_date" validate="required,gt=0"`
	OrderMoney *float32        `json:"order_money" db:"order_money" validate="required,gte=1000"`
	Status     *int            `json:"status" db:"status"`
}

type Payments []Payment

// 账户信息
func GetAccount(account *Account, user_id interface{}) error {
	if err := corm.Select(`select user_id,email,user_type,deal_type,deal_scale,real_name,
	company_name,qq,phone,to_date(create_date::text, 'yyyy-mm-dd') create_date,status
	from account_account where user_id=:user_id`).Where(map[string]interface{}{
		"user_id": user_id,
	}).Load(account); err != nil {
		return err
	}
	return nil
}

// 财务信息
func GetFinanceByUserId(finance *Finance, user_id interface{}) error {
	if err := corm.Select(`select * from account_finance where user_id=:user_id`).Where(map[string]interface{}{
		"user_id": user_id,
	}).Load(finance); err != nil {
		return err
	}
	return nil
}

// 判断用户是否存在
func IfEmailExist(email string, roleId int) bool {
	user_email := ""
	corm.Select(`select email from user_user where email=:email and role=:role`).Where(map[string]interface{}{
		"email": email,
		"role":  roleId,
	}).Load(&user_email)
	if user_email != "" {
		return true
	}
	return false
}

// 账户列表
func GetAccounts(roleId int, total *int64, req *AccountReq, accounts *Accounts) error {
	orm := corm.Select(`select user_id,email,user_type,deal_type,deal_scale,real_name,
	company_name,qq,phone,create_date,status,(select count(*) from app_app where user_id=t1.user_id) as app_count
	from account_account as t1 {{sql_where}}`).Req(req).Where(map[string]interface{}{
		"user_role": roleId,
	}).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(accounts); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 账户审核
func ReviewAccount(userId interface{}, roleId, status int) error {
	if _, err := corm.Db.Exec(`update account_account set 
		status=$1 where user_id=$2 and user_role=$3`, status, userId, roleId); err != nil {
		return err

	}
	return nil
}

// 提现列表
func GetPayments(total *int64, req *PaymentReq, payments *Payments) error {
	orm := corm.Select(`select id,user_id,apply_date,
    order_money,balance,status,(select email from user_user where id=t1.user_id)
		user_email from account_payment as t1 {{sql_where}} order by apply_date desc`).Req(req).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(payments); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 获取单个提现
func GetPayment(paymentId interface{}, payment *Payment) error {
	if err := corm.Db.Get(payment, `select *,(select email from user_user where id=t1.user_id)
		user_email from account_payment as t1 where id=$1`, paymentId); err != nil {
		return err
	}
	return nil
}

// 提现审核
func ReviewPayment(paymentID interface{}, roleID, status int, payment *Payment) error {
	tx := corm.Db.MustBegin()
	if _, err := tx.Exec(`update account_payment set status=$1,
	order_date=$2  where id=$3 and status!=1 and user_role=$4`, status,
		mydate.CurrentDateTime(), paymentID, roleID); err != nil {
		tx.Rollback()
		return err
	}
	if *payment.Status == 0 {
		if _, err := tx.Exec(`update account_balance set 
		balance=balance-$1 where user_id=$2 and user_role=$3`, *payment.OrderMoney,
			*payment.UserId, roleID); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	// 通知 更新缓存
	return nil
}

// 账号信息
func GetUser(email string, roleID int, user *comm.User) error {
	if err := corm.Db.Get(user, `select * from user_user where email=$1 and role=$2`, email, roleID); err != nil {
		return err
	}
	return nil
}

// 修改密码
func UpdatePassword(password, email string, roleId int) error {
	if _, err := corm.Db.Exec(`update user_user set password=$1 where email=$2 and role=$3`,
		password, email, roleId); err != nil {
		return err
	}
	return nil
}

// 更改当前token 版本
func UpdateTokenVersion(userId interface{}, rd *redis.Pool) error {
	client := rd.Get()
	if _, err := client.Do("SET", fmt.Sprintf("login_user:%v", userId), uuid.NewV4().String()); err != nil {
		return err
	}
	client.Close()
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

// 开户
func OpenAccount(userId *int64, roleId int, dealScale float64, reg *Reg, finance *Finance,
	companyName, ip, appSecretKey string) error {
	tx := corm.Db.MustBegin()
	// 创建账号信息
	if err := tx.QueryRow(`insert into user_user(email,password,type,role,real_name,company_name,ip_address,status)
	 values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`, *reg.Email, password.SetPassword(*reg.PassWord), *reg.UserType, roleId,
		*reg.RealName, companyName, ip, 1, ).Scan(userId); err != nil {
		return err
	}
	appKey, err := utils.NewCBCEncrypter([]byte(strconv.Itoa(int(*userId))), []byte(appSecretKey))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 创建账户信息
	if _, err := tx.Exec(`insert into account_account(user_id,user_type,user_role,email,real_name,company_name,
	app_key,qq,phone,status,deal_type,deal_scale) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		*userId, *reg.UserType, roleId, *reg.Email,
		*reg.RealName, companyName, base64.RawStdEncoding.EncodeToString(appKey),
		*reg.Qq, *reg.Phone, 1, *reg.DealType, dealScale); err != nil {
		tx.Rollback()
		return err
	}
	// 创建税务信息
	if _, err := tx.Exec(`insert into account_tax(user_id,
		user_role,company_name) values($1,$2,$3)`, *userId, roleId, companyName); err != nil {
		tx.Rollback()
		return err
	}
	// 创建邮寄信息
	if _, err := tx.Exec(`insert into account_deliver(user_id,telephone,user_role,email) values($1,$2,$3,$4)`,
		*userId, *reg.Phone, roleId, *reg.Email); err != nil {
		tx.Rollback()
		return err
	}
	// 创建账户余额
	if _, err := tx.Exec(`insert into account_balance(user_id,user_role,balance,app_key) values($1,$2,$3,$4)`,
		*userId, roleId, 0, base64.RawStdEncoding.EncodeToString(appKey)); err != nil {
		tx.Rollback()
		return err
	}
	// 创建财务信息
	if _, err := tx.Exec(`insert into account_finance(user_id,account_name,bank_name,
    bank_no,bank_province,bank_city,sub_branch,user_role) values($1,$2,$3,$4,$5,$6,$7,$8)`, *userId, *finance.AccountName, *finance.BankName,
		*finance.BankNo, *finance.BankProvince, *finance.BankCity, *finance.SubBranch, roleId); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
