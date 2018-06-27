package models

import (
	"zonstfe_api/common/utils"
	"encoding/json"
	"zonstfe_api/corm"
	"zonstfe_api/common/utils/password"
	"zonstfe_api/common/utils/mydate"
	"strconv"
	"encoding/base64"
	"errors"
	comm "zonstfe_api/common/models"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/satori/go.uuid"
	"zonstfe_api/common/my_context"
)

type AccountReq struct {
	Email    *string `form:"email" db:"email" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Sdate    *string `form:"sdate" db:"create_date" query:"gte"`
	Edate    *string `form:"edate" db:"create_date" query:"lte"`
	Status   *int    `form:"status" db:"status" query:"eq"`
}
type RechargeReq struct {
	UserId   *int  `form:"user_id" db:"user_id" query:"eq"`
	Status   *int  `form:"status" db:"status" query:"eq"`
	Page     *uint `form:"page"`
	PageSize *uint `form:"page_size"`
}

type Account struct {
	UserId        *int    `json:"user_id" db:"user_id" label:"用户ID"`
	UserType      *string `json:"user_type" db:"user_type" label:"用户类型"`
	Email         *string `json:"email" db:"email" label:"邮箱"`
	CompanyName   *string `json:"company_name" db:"company_name"  validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName      *string `json:"real_name" db:"real_name" validate:"required,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	CampaignCount *int    `json:"campaign_count,omitempty" db:"campaign_count"`
	//AppKey      *string `json:"app_key" db:"app_key" label:"APP_KEY"`
	DealType   *string         `json:"deal_type,omitempty" db:"deal_type" label:"盈利模式"`
	DealScale  *float32        `json:"deal_scale,omitempty" db:"deal_scale" label:"分成比例"`
	Qq         *string         `json:"qq" db:"qq" validate:"required,regexp=qq"`
	Phone      *string         `json:"phone" db:"phone" validate:"required,regexp=phone"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
	Status     *int            `json:"status" db:"status"`
}

type Accounts []Account

type Reg struct {
	Email       *string `json:"email" validate:"required,gt=0,regexp=email" label:"邮箱" error:"请输入正确的邮箱"`
	PassWord    *string `json:"password" validate:"required,gte=6,lte=20" label:"密码" error:"请输入正确的密码"`
	DpassWord   *string `json:"dpassword" validate:"required,gte=6,lte=20" label:"确认密码" error:"两次密码不相同"`
	CompanyName *string `json:"company_name" validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" validate:"required,gt=0,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	UserType    *string `json:"user_type" validate:"required,gt=0,option=user_type" label:"用户类型"`
	Qq          *string `json:"qq" validate:"required,gt=0,regexp=qq" label:"QQ" error:"请输入正确的QQ号"`
	Phone       *string `json:"phone" validate:"required,gt=0,regexp=phone" label:"手机" error:"请输入正确的手机号"`
}
type AccountUpdate struct {
	UserType    *string `json:"user_type" validate:"required,gt=0,option=user_type" label:"用户类型"`
	CompanyName *string `json:"company_name" validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" validate:"required,gt=0,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	Qq          *string `json:"qq" validate:"required,gt=0,regexp=qq" label:"QQ" error:"请输入正确的QQ号"`
	Phone       *string `json:"phone" validate:"required,gt=0,regexp=phone" label:"手机" error:"请输入正确的手机号"`
}

// 如果为个人 tax 可不填
type AccountOpen struct {
	Account *json.RawMessage `json:"account" validate:"required"`
	Tax     *json.RawMessage `json:"tax" `
	Deliver *json.RawMessage `json:"deliver" validate:"required"`
}

type Deliver struct {
	Id        int     `json:"deliver_id" db:"id"`
	UserId    *int    `json:"user_id" db:"user_id"`
	UserRole  *int    `json:"user_role" db:"user_role"`
	Address   *string `json:"address" db:"address" validate:"required,gt=0,lte=300"`
	Receiver  *string `json:"receiver" db:"receiver" validate:"required,regexp=real_name"`
	Telephone *string `json:"telephone" db:"telephone" validate:"required,regexp=telephone"`
	Email     *string `json:"email" db:"email" validate:"required,regexp=email"`
}

type Tax struct {
	Id          int     `json:"tax_id" db:"id"`
	UserId      *int    `json:"user_id" db:"user_id"`
	UserRole    *int    `json:"user_role" db:"user_role"`
	TaxNo       *string `json:"tax_no" db:"tax_no" validate:"required,gte=15,lte=20"`
	CompanyName *string `json:"company_name" db:"company_name"  validate:"required,regexp=company_name"`
	Address     *string `json:"address" db:"address" validate:"required,gt=0,lte=300"`
	Telephone   *string `json:"telephone" db:"telephone" validate:"required,regexp=telephone"`
	BankName    *string `json:"bank_name" db:"bank_name" validate:"required,gt=0"`
	BankNo      *string `json:"bank_no" db:"bank_no" validate:"required,gte=16,lte=19"`
}

type AccountRecharge struct {
	Id          *int `json:"recharge_id" db:"id"`
	UserId      *int `json:"user_id" db:"user_id" validate:"required"`
	ZonstUserId *int `json:"zonst_user_id"  db:"zonst_user_id"`
	//OrderNo     *string          `json:"order_no" db:"order_no" validate:"required"`
	OrderMoney  *utils.JSONFloat `json:"order_money" db:"order_money" validate:"required"`
	OrderType   *int             `json:"order_type" db:"order_type" validate:"required"`
	OrderDate   *utils.JSONTime  `json:"order_date" db:"order_date" validate:"required"`
	AccountNo   *string          `json:"account_no" db:"account_no" validate:"required"`
	AccountName *string          `json:"account_name" db:"account_name" validate:"required"`
	Description *string          `json:"description" db:"description" validate:"required"`
	UpdateDate  *utils.JSONTime  `json:"update_date" db:"update_date"`
	CreateDate  *utils.JSONTime  `json:"create_date" db:"create_date"`
	Status      *int             `json:"status" db:"status"`
}
type AccountRecharges []AccountRecharge

type Review struct {
	ReviewType *int `json:"review_type" validate:"required"`
}
type ReviewBatch struct {
	ReviewType *int      `json:"review_type" validate:"required"`
	List       *[]int `json:"list" validate:"required"`
}

type ReviewBad struct {
	ReviewType *int    `json:"review_type" validate:"required"`
	Title      *string `json:"title" validate:"required,gt=0,lte=100"`
	Content    *string `json:"content" validate:"required,gt=0"`
	GroupName  *string `json:"group_name" validate:"required,gt=0"`
}

type BankInfo struct {
	AccountName *string `json:"account_name"`
	BankName    *string `json:"bank_name"` //开户行
	BankNo      *string `json:"bank_no"`   //银行账号
}

// 账户信息
func GetAccount(userId interface{}, roleId int, account *Account) error {
	if err := corm.Db.Get(account, `select user_id,email,user_type,real_name,
	company_name,qq,phone,to_date(create_date::text, 'yyyy-mm-dd') create_date,status
	from account_account where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 联系信息
func GetDeliverByUserId(userId interface{}, roleId int, deliver *Deliver) error {
	if err := corm.Db.Get(deliver, `select * from account_deliver
		where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 税务信息
func GetDeliverById(deliverId interface{}, roleID int, deliver *Deliver) error {
	if err := corm.Db.Get(deliver, `select * from account_deliver
		where id=$1 and user_role=$2`, deliverId, roleID); err != nil {
		return err
	}
	return nil
}

// 税务信息
func GetTaxByUserId(userId interface{}, roleId int, tax *Tax) error {
	if err := corm.Db.Get(tax, `select * from account_tax
		where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 税务信息
func GetTax(taxId interface{}, roleId int, tax *Tax) error {
	if err := corm.Db.Get(tax, `select * from account_tax
		where id=$1 and user_role=$2`, taxId, roleId); err != nil {
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

// 广告主列表
func GetAccounts(roleId int, total *int64, req *AccountReq, data *Accounts) error {
	orm := corm.Select(`select user_id,email,user_type,real_name,
	company_name,qq,phone,create_date,status,(select count(*)
	from campaign_campaign where user_id=t1.user_id) as campaign_count
	from account_account as t1 {{sql_where}}`).Req(&req).Where(map[string]interface{}{
		"user_role": roleId,
	}).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(data); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 开户
func OpenAccount(c *my_context.Context, userId *int64, roleId int, accountOpen *AccountOpen, reg *Reg,
	deliver *Deliver, companyName, ip, appSecretKey string) error {
	tax := &Tax{}
	tx := corm.Db.MustBegin()
	// 创建账号信息
	if err := tx.QueryRow(`insert into user_user(email,password,type,role,
     real_name,company_name,ip_address,status)
	 values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`, *reg.Email, password.SetPassword(*reg.PassWord), *reg.UserType, roleId,
		*reg.RealName, companyName, ip, 1).Scan(userId); err != nil {
		return err
	}
	appKey, err := utils.NewCBCEncrypter([]byte(strconv.Itoa(int(*userId))), []byte(appSecretKey))
	if err != nil {
		tx.Rollback()
		return err
	}
	// 创建账户信息
	if _, err := tx.Exec(`insert into account_account(user_id,user_type,user_role,email,real_name,company_name,
	app_key,qq,phone,status) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`, *userId, *reg.UserType, roleId, *reg.Email,
		*reg.RealName, companyName, base64.RawStdEncoding.EncodeToString(appKey), *reg.Qq, *reg.Phone, 1); err != nil {
		tx.Rollback()
		return err
	}
	// 创建税务信息
	if *reg.UserType == "company" {
		if accountOpen.Tax == nil {
			tx.Rollback()
			return errors.New("")
		}
		if err := c.BindJson(string(*accountOpen.Tax), tax); err != nil {
			tx.Rollback()
			return err
		}
		if _, err := tx.Exec(`insert into account_tax(user_id,user_role,tax_no,
			company_name,address,telephone,bank_name,bank_no)
			values($1,$2,$3,$4,$5,$6,$7,$8)`, *userId, roleId, *tax.TaxNo, *tax.CompanyName,
			*tax.Address, *tax.Telephone, *tax.BankName, *tax.BankNo); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if _, err := tx.Exec(`insert into account_tax(user_id,
		user_role,company_name) values($1,$2,$3)`, *userId, roleId, companyName); err != nil {
			tx.Rollback()
			return err
		}
	}
	// 创建邮寄信息
	if _, err := tx.Exec(`insert into account_deliver(user_id,telephone,user_role,email,address,receiver) 
	values($1,$2,$3,$4,$5,$6)`, *userId, *deliver.Telephone, roleId, *deliver.Email,
		*deliver.Address, *deliver.Receiver); err != nil {
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
	if _, err := tx.Exec(`insert into account_finance(user_id,user_role) values($1,$2)`, *userId, roleId); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
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

// 修改税务信息
func UpdateTax(userId, roleId int, tax *Tax) error {
	if _, err := corm.Db.Exec(`update account_tax set tax_no=$1,
	company_name=$2,bank_no=$3,bank_name=$4,
	address=$5,telephone=$6 where user_id=$7 and user_role=$8`, *tax.TaxNo,
		*tax.CompanyName, *tax.BankNo, *tax.BankName,
		*tax.Address, *tax.Telephone, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 修改联系人信息
func UpdateDeliver(userId, roleId int, deliver *Deliver) error {
	if _, err := corm.Db.Exec(`update account_deliver set receiver=$1,email=$2,address=$3,
	telephone=$4 where user_id=$5 and user_role=$6`, *deliver.Receiver,
		*deliver.Email, *deliver.Address, *deliver.Telephone, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 账户信息修改
func UpdateAccount(userId interface{}, roleId int, accountId *int64, companyName string, account *AccountUpdate) error {
	if err := corm.Db.QueryRow(`update account_account set real_name=$1,
	company_name=$2,qq=$3,phone=$4,user_type=$5 
	where user_id=$6 and user_role=$7`, *account.RealName, companyName,
		*account.Qq, *account.Phone, *account.UserType, userId, roleId).Scan(accountId); err != nil {
		return err
	}
	return nil
}

// 验证当前充值是否已经处理
func CheckRechargeStatus(rechargeId interface{}, status int) error {
	recharge := &AccountRecharge{}
	if err := corm.Db.Get(recharge, `select * from
	account_recharge where id=$1 and status=$2`, rechargeId, status); err != nil {
		return err
	}

	return nil
}

// 获取充值记录
func GetRecharge(rechargeId interface{}, roleId int, recharge *AccountRecharge) error {
	if err := corm.Db.Get(recharge, `select * from account_recharge 
		where id=$1 and user_role=$2`, rechargeId, roleId); err != nil {
		return err
	}
	return nil
}

// 账户充值
func AddRecharge(rechargeId *int64, roleID int, recharge *AccountRecharge) error {
	if err := corm.Db.QueryRow(`insert into
	account_recharge(user_id,user_role,order_money,
	order_type,account_no,account_name,description)
	values($1,$2,$3,$4,$5,$6,$7)
	 RETURNING id`, *recharge.UserId,
		roleID,
		*recharge.OrderMoney, *recharge.OrderType,
		*recharge.AccountNo, *recharge.AccountName, *recharge.Description).Scan(rechargeId); err != nil {
		return err
	}
	return nil
}

// 充值列表
func GetRecharges(roleId int, total *int64, req *RechargeReq, recharges *AccountRecharges) error {
	orm := corm.Select(`select * from account_recharge {{sql_where}}
	ORDER BY create_date desc`).Req(req).Where("user_role", roleId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(recharges); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 充值审核
func ReviewRecharge(rechargeID interface{}, roleID, status int, recharge *AccountRecharge) error {
	tx := corm.Db.MustBegin()
	if _, err := tx.Exec(`update account_recharge set 
		status=$1,update_date=$2 where id=$3 and user_role=$4 and status!=1`, status,
		mydate.CurrentDateTime(), rechargeID, roleID); err != nil {
		tx.Rollback()
		return err
	}
	if *recharge.Status == 0 {
		if _, err := tx.Exec(`update account_balance set 
			balance=balance+$1 where user_id=$2 and user_role=$3`, *recharge.OrderMoney,
			*recharge.UserId, roleID); err != nil {
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

// 账户审核
func ReviewAccount(userId interface{}, roleId, status int) error {
	if _, err := corm.Db.Exec(`update account_account set 
		status=$1 where user_id=$2 and user_role=$3`, status, userId, roleId); err != nil {
		return err

	}
	return nil
}
