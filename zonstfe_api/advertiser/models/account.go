package models

import (
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
	comm "zonstfe_api/common/models"
	"github.com/satori/go.uuid"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

type Account struct {
	//UserId      *int    `json:"user_id" db:"user_id" label:"用户ID"`
	UserType    *string `json:"user_type" db:"user_type" label:"用户类型"`
	Email       *string `json:"email" db:"email" label:"邮箱"`
	CompanyName *string `json:"company_name" db:"company_name"  validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" db:"real_name" validate:"required,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	AppKey      *string `json:"app_key" db:"app_key" label:"APP_KEY"`
	//DealType    *string `json:"deal_type" db:"deal_type" label:"盈利模式"`
	Qq     *string `json:"qq" db:"qq" validate:"required,gt=0,regexp=qq"`
	Phone  *string `json:"phone" db:"phone" validate:"required,gt=0,regexp=phone"`
	Status *int    `json:"status" db:"status"`
}

type Accounts []Account

type AccountRecharge struct {
	Id *int `json:"recharge_id" db:"id"`
	//UserId      *int `json:"user_id" db:"user_id" validate:"required"`
	//ZonstUserId *int `json:"zonst_user_id"  db:"zonst_user_id" validate:"required"`
	//OrderNo     *string `json:"order_no" db:"order_no" validate:"required"`
	OrderMoney *utils.JSONFloat `json:"order_money" db:"order_money" validate:"required"`
	OrderType  *int             `json:"order_type" db:"order_type" validate:"required"`
	//OrderDate   *utils.JSONTime `json:"order_date" db:"order_date" validate:"required"`
	AccountNo   *string `json:"account_no" db:"account_no" validate:"required"`
	AccountName *string `json:"account_name" db:"account_name" validate:"required"`
	Description *string `json:"description" db:"description" validate:"required"`
	//UpdateDate  *utils.JSONTime `json:"update_date" db:"update_date"`
	CreateDate *utils.JSONTime `json:"create_date" db:"create_date"`
	Status     *int            `json:"status" db:"status"`
}
type AccountRecharges []AccountRecharge

type RechargeReq struct {
	Status   *int  `form:"status" db:"status" query:"eq"`
	Page     *uint `form:"page"`
	PageSize *uint `form:"page_size"`
}

type AccountBalance struct {
	//UserId  *int `json:"user_id" db:"user_id"`
	Balance *utils.JSONFloat `json:"balance" db:"balance"`
}

type Deliver struct {
	Id        int     `json:"deliver_id"`
	UserId    *int    `json:"user_id"`
	UserRole  *int    `json:"user_role"`
	Address   *string `json:"address" validate:"required,gt=0,lte=300"`
	Receiver  *string `json:"receiver" validate:"required,regexp=real_name"`
	Telephone *string `json:"telephone" validate:"required,regexp=telephone"`
	Email     *string `json:"email" validate:"required,regexp=email"`
}

type Tax struct {
	Id          int     `json:"tax_id"`
	UserId      *int    `json:"user_id"`
	UserRole    *int    `json:"user_role"`
	TaxNo       *string `json:"tax_no" validate:"required,gte=15,lte=20"`
	CompanyName *string `json:"company_name"  validate:"required,regexp=company_name"`
	Address     *string `json:"address" validate:"required,gt=0,lte=300"`
	Telephone   *string `json:"telephone" validate:"required,regexp=telephone"`
	BankName    *string `json:"bank_name" validate:"required,gt=0"`
	BankNo      *string `json:"bankno" validate:"required,gte=16,lte=19"`
}

// 余额
func GetAccountBalance(userId interface{}, roleId int, balance *AccountBalance) error {
	if err := corm.Db.Get(balance, `select * from account_balance 
		where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 账户信息
func GetAccount(userId interface{}, roleId int, account *Account) error {
	if err := corm.Db.Get(account, `select * from account_account where user_id=$1 
	and user_role=$2`, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 修改账户
func UpdateAccount(userId interface{}, roleId int, account *Account) error {
	if _, err := corm.Db.Exec(`update account_account set real_name=$1,
		company_name=$2,phone=$3,qq=$4  where user_id=$5 and user_role=$2`, *account.RealName, *account.CompanyName,
		*account.Phone, *account.Qq, userId, roleId); err != nil {
		return err
	}
	return nil
}

// 判断用户是否存在
func CheckUserExist(userId interface{}, user *comm.User) error {
	if err := corm.Db.Get(user, `select * from user_user where id=$1`, userId); err != nil {
		return err
	}
	return nil
}

// 修改密码
func UpdatePassWord(userId interface{}, password string) error {
	if _, err := corm.Db.Exec(`update user_user set password=$1 where id=$2`, password, userId); err != nil {
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

// 查询是否还有待审核数据
func GetRechargeReview(userId interface{}, recharges *AccountRecharges) error {
	if err := corm.Db.Select(recharges, `select * from account_recharge 
		where user_id=$1 and status=0`, userId); err != nil {
		return err
	}
	return nil
}

// 新增充值记录
func AddRecharge(userId interface{}, roleId int, rechargeId *int64, recharge *AccountRecharge) error {
	if err := corm.Db.QueryRow(`insert into
	account_recharge(user_id,user_role,order_money,order_type,account_no,account_name,description)
	values($1,$2,$3,$4,$5,$6,$7)
	 RETURNING id`, userId, roleId, *recharge.OrderMoney, *recharge.OrderType,
		*recharge.AccountNo, *recharge.AccountName, *recharge.Description, ).Scan(rechargeId); err != nil {
		return err
	}
	return nil
}

// 获取充值记录列表
func GetRecharges(userId interface{}, total *int64, req *RechargeReq, recharges *AccountRecharges) error {
	orm := corm.Select(`select * from account_recharge {{sql_where}} 
	ORDER BY create_date desc`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(recharges); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}
