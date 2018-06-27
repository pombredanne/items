package models

import (
	"zonstfe_api/common/utils"
	"zonstfe_api/corm"
	comm "zonstfe_api/common/models"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/satori/go.uuid"
)

type PaymentReq struct {
	Status   *int  `form:"status" db:"status" query:"eq"`
	Page     *uint `form:"page"`
	PageSize *uint `form:"page_size"`
}

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

type InvoiceInfo struct {
	Name *string `json:"name"`                                         //开票项目名称
	Type *int    `json:"type" validate:"required,option=invoice_type"` //发票类型
	//EntityName  *string `json:"entity_name"`                              //开票主体
	CompanyName *string `json:"company_name" validate:"required,gt=0,regexp=company_name"` //开票企业名称
	TaxNo       *string `json:"tax_no" validate:"required,gt=0"`                           //企业税号
	Address     *string `json:"address" validate:"required,gt=0"`                          //注册地址
	Telephone   *string `json:"telephone" validate:"required,gt=0,regexp=telephone"`       //注册电话
	BankName    *string `json:"bank_name" validate:"required,gt=0"`                        //开户行
	BankNo      *string `json:"bank_no" validate:"required,gt=0"`                          //银行账号
}
type Deliver struct {
	Receiver  *string `json:"receiver"`  //收件人
	Address   *string `json:"address"`   //邮寄地址
	Telephone *string `json:"telephone"` //邮寄电话
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

type AccountBalance struct {
	//UserId  *int `json:"user_id" db:"user_id"`
	Balance *utils.JSONFloat `json:"balance" db:"balance"`
}

type Payment struct {
	Id         *int            `json:"payment_id" db:"id"`
	Balance    *float32        `json:"balance" db:"balance"`
	ApplyDate  *utils.JSONTime `json:"apply_date" db:"apply_date" validate="required,gt=0"`
	OrderMoney *float32        `json:"order_money" db:"order_money" validate="required,gte=1000"`
	Status     *int            `json:"status" db:"status"`
}

type Payments []Payment

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

// 余额
func GetAccountBalance(userId interface{}, roleId int, balance *AccountBalance) error {
	if err := corm.Db.Get(balance, `select * from account_balance 
		where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
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

// 财务信息
func GetFinance(userId interface{}, finance *Finance) error {
	if err := corm.Db.Get(finance, `select * from account_finance where user_id=$1`, userId); err != nil {
		return err
	}
	return nil
}

// 修改财务信息
func UpdateFinance(userId, financeId interface{}, finance *Finance) error {
	if _, err := corm.Db.Exec(`update account_finance set account_name=$1,
		bank_name=$2,bank_province=$3,bank_city=$4,bank_no=$5,sub_branch=$6 where id=$7 and user_id=$8`,
		*finance.AccountName, *finance.BankName, *finance.BankProvince,
		*finance.BankCity, *finance.BankNo, *finance.SubBranch, financeId, userId); err != nil {
		return err
	}
	return nil
}

// 提现列表
func GetPayments(userId interface{}, total *int64, req *PaymentReq, payments *Payments) error {
	orm := corm.Select(`select order_money,apply_date,status
	from account_payment {{sql_where}} order by apply_date desc`).Req(req).Where("user_id", userId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(payments); err != nil {
		return err
	}
	if err := orm.Total(total); err != nil {
		return err
	}
	return nil
}

// 提现申请
func AddPayment(paymentId *int64, args []interface{}) error {
	if err := corm.Db.QueryRow(`insert into 
        account_payment(user_id,user_role,balance,order_money,status) 
		values($1,$2,$3,$4,$5) RETURNING id`, args...).Scan(paymentId); err != nil {
		return err
	}
	return nil
}

// 获取提现列表 通过status
func GetPaymentByStatus(userId interface{}, status int, payments *Payments) error {
	if err := corm.Db.Select(payments, `select * from account_payment
	where user_id=$1 and status=$2`, userId, status); err != nil {

	}
	return nil
}
