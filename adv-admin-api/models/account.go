package models

import (
	"corm"
	"api-libs/my-date"
	"api-libs/my-field"
	"api-libs/rsp"
	"encoding/json"
	"api-libs/password"
	"github.com/satori/go.uuid"
)

type AccountReq struct {
	UserId   *int    `form:"user_id" db:"user_id" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Sdate    *string `form:"sdate" db:"DATE(create_date)" query:"gte"`
	Edate    *string `form:"edate" db:"DATE(create_date)" query:"lte"`
	Status   *int    `form:"status" db:"status" query:"eq"`
}
type RechargeReq struct {
	UserId   *int    `form:"user_id" db:"user_id" query:"eq"`
	Status   *int    `form:"status" db:"status" query:"eq"`
	Page     *uint   `form:"page"`
	PageSize *uint   `form:"page_size"`
	Sdate    *string `form:"sdate" db:"DATE(create_date)" query:"gte"`
	Edate    *string `form:"edate" db:"DATE(create_date)" query:"lte"`
}

type Account struct {
	UserId        *int    `json:"user_id" db:"user_id" label:"用户ID"`
	UserType      *string `json:"user_type" db:"user_type" label:"用户类型"`
	Email         *string `json:"email" db:"email" label:"邮箱"`
	CompanyName   *string `json:"company_name" db:"company_name"  validate:"omitempty,regexp=company_name" label:"企业名称" error:"请输入正确的企业名称"`
	RealName      *string `json:"real_name" db:"real_name" validate:"required,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	CampaignCount *int    `json:"campaign_count,omitempty" db:"campaign_count"`

	//AppKey      *string `json:"app_key" db:"app_key" label:"APP_KEY"`
	//DealType   *string         `json:"deal_type,omitempty" db:"deal_type" label:"盈利模式"`
	//DealScale  *float32        `json:"deal_scale,omitempty" db:"deal_scale" label:"分成比例"`
	Balance     *float64        `json:"balance" db:"balance"`
	Qq          *string         `json:"qq" db:"qq" validate:"required,regexp=qq"`
	Phone       *string         `json:"phone" db:"phone" validate:"required,regexp=phone"`
	CreateDate  *my_field.JSONTime `json:"create_date" db:"create_date"`
	Status      *int            `json:"status" db:"status"`
	AccountType *string         `json:"account_type" db:"account_type"` //账户类型 {"prepay":"预付","virtual":"虚拟"}
	BusinessId  *int            `json:"business_id" db:"business_id"`   //商务ID
}

type Accounts []Account

type Reg struct {
	Email       *string `json:"email" validate:"required,gt=0,regexp=email" label:"邮箱" error:"请输入正确的邮箱"`
	PassWord    *string `json:"password" validate:"required,gte=6,lte=20" label:"密码" error:"请输入正确的密码"`
	DpassWord   *string `json:"dpassword" validate:"required,gte=6,lte=20" label:"确认密码" error:"两次密码不相同"`
	CompanyName *string `json:"company_name" validate:"required" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" validate:"required,gt=0,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	UserType    *string `json:"user_type" validate:"required,gt=0,option=user_type" label:"用户类型"`
	Qq          *string `json:"qq" validate:"required,gt=0,regexp=qq" label:"QQ" error:"请输入正确的QQ号"`
	Phone       *string `json:"phone" validate:"required,gt=0,regexp=phone" label:"手机" error:"请输入正确的手机号"`
	AccountType *string `json:"account_type" validate:"required,gt=0"` //账户类型 {"prepay":"预付","virtual":"虚拟"}
	BusinessId  *int    `json:"business_id" validate:"required,gt=0"`  //商务ID
}
type AccountUpdate struct {
	UserType    *string `json:"user_type" validate:"required,gt=0,option=user_type" label:"用户类型"`
	CompanyName *string `json:"company_name" validate:"required" label:"企业名称" error:"请输入正确的企业名称"`
	RealName    *string `json:"real_name" validate:"required,gt=0,regexp=real_name" label:"姓名" error:"请输入正确的联系人"`
	Qq          *string `json:"qq" validate:"required,gt=0,regexp=qq" label:"QQ" error:"请输入正确的QQ号"`
	Phone       *string `json:"phone" validate:"required,gt=0,regexp=phone" label:"手机" error:"请输入正确的手机号"`
	AccountType *string `json:"account_type" validate:"required,gt=0"` //账户类型 {"prepay":"预付","virtual":"虚拟"}
	BusinessId  *int    `json:"business_id" validate:"required,gt=0"`  //商务ID
}

// 如果为个人 tax 可不填
type AccountOpen struct {
	Account       *json.RawMessage `json:"account" validate:"required"`
	Tax           *json.RawMessage `json:"tax"`
	Finance       *json.RawMessage `json:"finance"`
	Deliver       *json.RawMessage `json:"deliver"`
	Qualification *json.RawMessage `json:"qualification"`
	Industry      *json.RawMessage `json:"industry"`
}

type Deliver struct {
	Id        int     `json:"deliver_id" db:"id"`
	UserId    *int    `json:"user_id" db:"user_id"`
	UserRole  *int    `json:"user_role" db:"user_role"`
	Address   *string `json:"address" db:"address" validate:"required,lte=300"`
	Receiver  *string `json:"receiver" db:"receiver" validate:"required,regexp=real_name"`
	Telephone *string `json:"telephone" db:"telephone" validate:"required,regexp=telephone"`
	Email     *string `json:"email" db:"email" validate:"required,regexp=email"`
}

type Delivers []Deliver

type Tax struct {
	Id          int     `json:"tax_id" db:"id"`
	UserId      *int    `json:"user_id" db:"user_id"`
	UserRole    *int    `json:"user_role" db:"user_role"`
	TaxNo       *string `json:"tax_no" db:"tax_no" validate:"required,gte=15,lte=20"`
	CompanyName *string `json:"company_name" db:"company_name"  validate:"required,gt=0,lte=300,regexp=company_name"`
	Address     *string `json:"address" db:"address" validate:"required,gt=0,lte=300"`
	Telephone   *string `json:"telephone" db:"telephone" validate:"required,regexp=telephone"`
	BankName    *string `json:"bank_name" db:"bank_name" validate:"required,gt=0,lte=100"`
	BankNo      *string `json:"bank_no" db:"bank_no" validate:"required,gte=15,lte=19"`
}

type Taxs []Tax

type Finance struct {
	Id           *int    `json:"finance_id" db:"id"`
	AccountName  *string `json:"account_name" db:"account_name" validate:"required,gt=0"`
	BankName     *string `json:"bank_name" db:"bank_name" validate:"required,gt=0,lte=100"` //开户行
	BankNo       *string `json:"bank_no" db:"bank_no" validate:"required,gte=15,lte=19"`    //银行账号
	BankProvince *string `json:"bank_province" db:"bank_province" validate:"required,gt=0"` //开户省
	BankCity     *string `json:"bank_city" db:"bank_city" validate:"required,gt=0"`
	SubBranch    *string `json:"sub_branch" db:"sub_branch" validate:"required,gt=0,lte=300"` //开户地支行地址
}

type Finances []Finance

type AccountRecharge struct {
	Id         *int `json:"recharge_id" db:"id"`
	UserId     *int `json:"user_id" db:"user_id" validate:"required"`
	BusinessId *int `json:"business_id"  db:"business_id"`
	//OrderNo     *string          `json:"order_no" db:"order_no" validate:"required"`
	OrderMoney  *float64        `json:"order_money" db:"order_money" validate:"required"`
	OrderType   *int            `json:"order_type" db:"order_type" validate:"required"`
	OrderDate   *my_field.JSONTime `json:"order_date" db:"order_date" validate:"required"`
	AccountNo   *string         `json:"account_no" db:"account_no" validate:"required"`
	AccountName *string         `json:"account_name" db:"account_name" validate:"required"`
	Description *string         `json:"description" db:"description" validate:"required"`
	UpdateDate  *my_field.JSONTime `json:"update_date" db:"update_date"`
	CreateDate  *my_field.JSONTime `json:"create_date" db:"create_date"`
	Status      *int            `json:"status" db:"status"`
}
type AccountRecharges []AccountRecharge

type RechargeCost struct {
	ReportDate  *my_field.JSONTime `json:"report_date" db:"report_date"`
	UserId      *int            `json:"user_id" db:"user_id"`
	CostMoney   *float64        `json:"cost_money" db:"cost_money"`
	ReportMoney *float64        `json:"report_money" db:"report_money"`
	RechargeId  *int            `json:"recharge_id" db:"recharge_id"`
}

type RechargeCosts []RechargeCost

type AccountBalance struct {
	//UserId  *int `json:"user_id" db:"user_id"`
	Balance *float64 `json:"balance" db:"balance"`
}

type Review struct {
	ReviewType *int `json:"review_type" validate:"required"`
}
type ReviewBatch struct {
	ReviewType *int   `json:"review_type" validate:"required"`
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

// 获取财务信息
func GetFinance(userId interface{}, data *Finances) error {
	if err := corm.Db.Select(data, `select * from account_finance where user_id=$1`, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 账户信息
func GetAccount(userId interface{}, roleId int, account *Account) error {
	if err := corm.Db.Get(account, `select user_id,email,user_type,real_name,
	company_name,qq,phone,to_date(create_date::text, 'yyyy-mm-dd') create_date,status,(select balance
	from account_balance where user_id=t1.user_id) as balance,account_type,business_id
	from account_account as t1 where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 联系信息
func GetDeliverByUserId(userId interface{}, data *Delivers) error {
	if err := corm.Db.Select(data, `select * from account_deliver
		where user_id=$1`, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 税务信息
func GetDeliverById(deliverId interface{}, roleID int, deliver *Deliver) error {
	if err := corm.Db.Get(deliver, `select * from account_deliver
		where id=$1 and user_role=$2`, deliverId, roleID); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 税务信息
func GetTaxByUserId(userId interface{}, tax *Taxs) error {
	if err := corm.Db.Select(tax, `select * from account_tax
		where user_id=$1`, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 税务信息
func GetTax(taxId interface{}, roleId int, tax *Tax) error {
	if err := corm.Db.Get(tax, `select * from account_tax
		where id=$1 and user_role=$2`, taxId, roleId); err != nil {
		return rsp.HandlerError(err)
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
	from campaign_campaign where user_id=t1.user_id) as campaign_count,
(select balance
	from account_balance where user_id=t1.user_id) as balance
	from account_account as t1 {{sql_where}} order by create_date desc`).Req(req).Where(map[string]interface{}{
		"user_role": roleId,
	}).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

type Qualification struct {
	IdCardPositive   *string `json:"id_card_positive" db:"id_card_positive"`
	IdCardNegative   *string `json:"id_card_negative" db:"id_card_negative"`
	BusinessLicense  *string `json:"business_license" db:"business_license"`
	TaxRegistration  *string `json:"tax_registration" db:"tax_registration"`
	OrganizationCode *string `json:"organization_code" db:"organization_code"`
	UserType         *string `json:"user_type" db:"user_type"`
}
type Qualifications []Qualification

type PersonQualification struct {
	IdCardPositive *string `json:"id_card_positive" validate:"required,gt=0,regexp=link"`
	IdCardNegative *string `json:"id_card_negative" validate:"required,gt=0,regexp=link"`
}
type CompanyQualification struct {
	BusinessLicense  *string `json:"business_license" validate:"required,gt=0,regexp=link"`
	TaxRegistration  *string `json:"tax_registration" validate:"required"`
	OrganizationCode *string `json:"organization_code" validate:"required"`
}

// 开户
func OpenAccount(userId *int64, roleId int, accountOpen *AccountOpen, reg *Reg, companyName, ip string) error {
	tax := &Tax{}
	finance := &Finance{}
	deliver := &Deliver{}
	person_qf := &PersonQualification{}
	company_qf := &CompanyQualification{}

	tx := corm.Db.MustBegin()
	// 创建账号信息
	if err := tx.QueryRow(`insert into user_user(email,password,type,role,
     real_name,company_name,ip_address,status)
	 values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`, *reg.Email, password.SetPassword(*reg.PassWord), *reg.UserType, roleId,
		*reg.RealName, companyName, ip, 1).Scan(userId); err != nil {
		return rsp.HandlerError(err)
	}
	appKey := uuid.NewV4().String()
	// 创建账户信息
	if _, err := tx.Exec(`insert into account_account(user_id,user_type,user_role,email,real_name,company_name,
	app_key,qq,phone,status,account_type,business_id) values($1,$2,$3,$4,$5,
	$6,$7,$8,$9,$10,$11,$12)`, *userId, *reg.UserType, roleId, *reg.Email,
		*reg.RealName, companyName, appKey, *reg.Qq,
		*reg.Phone, 1, *reg.AccountType, *reg.BusinessId); err != nil {
		tx.Rollback()
		return rsp.HandlerError(err)
	}
	// 创建税务信息
	if accountOpen.Tax != nil {
		if err := rsp.BindJson(string(*accountOpen.Tax), tax); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
		if _, err := tx.Exec(`insert into account_tax(user_id,user_role,tax_no,
			company_name,address,telephone,bank_name,bank_no)
			values($1,$2,$3,$4,$5,$6,$7,$8)`, *userId, roleId, *tax.TaxNo, *tax.CompanyName,
			*tax.Address, *tax.Telephone, *tax.BankName, *tax.BankNo); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	} else {
		if _, err := tx.Exec(`insert into account_tax(user_id,
		user_role,company_name) values($1,$2,$3)`, *userId, roleId, companyName); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	}
	//if *reg.UserType == "company" {
	//	if accountOpen.Tax == nil {
	//		tx.Rollback()
	//		return errors.New("")
	//	}
	//	if err := rsp.BindJson(string(*accountOpen.Tax), tax); err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//	if _, err := tx.Exec(`insert into account_tax(user_id,user_role,tax_no,
	//		company_name,address,telephone,bank_name,bank_no)
	//		values($1,$2,$3,$4,$5,$6,$7,$8)`, *userId, roleId, *tax.TaxNo, *tax.CompanyName,
	//		*tax.Address, *tax.Telephone, *tax.BankName, *tax.BankNo); err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//} else {
	//	if _, err := tx.Exec(`insert into account_tax(user_id,
	//	user_role,company_name) values($1,$2,$3)`, *userId, roleId, companyName); err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//}
	// 创建邮寄信息
	if accountOpen.Deliver != nil {
		if err := rsp.BindJson(string(*accountOpen.Deliver), deliver); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
		if _, err := tx.Exec(`insert into account_deliver(user_id,telephone,user_role,email,address,receiver)
		values($1,$2,$3,$4,$5,$6)`, *userId, roleId, *deliver.Email,
			*deliver.Address, *deliver.Receiver); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	} else {
		if _, err := tx.Exec(`insert into account_deliver(user_id,user_role)
		values($1,$2)`, *userId, roleId); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	}
	// 创建账户余额
	if _, err := tx.Exec(`insert into account_balance(user_id,user_role,balance,app_key) values($1,$2,$3,$4)`,
		*userId, roleId, 0, appKey); err != nil {
		tx.Rollback()
		return rsp.HandlerError(err)
	}
	// 创建财务信息
	if accountOpen.Finance != nil {
		if err := rsp.BindJson(string(*accountOpen.Finance), finance); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
		if _, err := tx.Exec(`insert into account_finance(user_id,user_role,
			account_name,bank_name,bank_no,bank_province,bank_city,sub_branch) values($1,$2,$3,$4,
			$5,$6,$7,$8)`, *userId, roleId, *finance.AccountName,
			*finance.BankName, *finance.BankNo,
			*finance.BankProvince, *finance.BankCity, *finance.SubBranch); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	} else {

		if _, err := tx.Exec(`insert into account_finance(user_id,user_role) values($1,$2)`, *userId, roleId); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	}
	if accountOpen.Qualification != nil {
		if *reg.UserType == "person" {
			if err := rsp.BindJson(string(*accountOpen.Qualification), person_qf); err != nil {
				tx.Rollback()
				return rsp.HandlerError(err)
			}
		} else {
			if err := rsp.BindJson(string(*accountOpen.Qualification), company_qf); err != nil {
				tx.Rollback()
				return rsp.HandlerError(err)
			}
		}
	} else {
		if _, err := tx.Exec(`insert into account_qualification(user_id) values($1)`, *userId); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	}
	if accountOpen.Industry != nil {

	} else {

	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return rsp.HandlerError(err)
	}
	return nil
}

// 账号信息
func GetUser(email string, roleID int, user *User) error {
	if err := corm.Db.Get(user, `select * from user_user where email=$1 and role=$2`, email, roleID); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改密码
func UpdatePassword(password, email string, roleId int) error {
	if _, err := corm.Db.Exec(`update user_user set password=$1 where email=$2 and role=$3`,
		password, email, roleId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改税务信息
//func UpdateTax(userId, roleId int, tax *Tax) error {
//	if _, err := corm.Db.Exec(`update account_tax set tax_no=$1,
//	company_name=$2,bank_no=$3,bank_name=$4,
//	address=$5,telephone=$6 where user_id=$7 and user_role=$8`, *tax.TaxNo,
//		*tax.CompanyName, *tax.BankNo, *tax.BankName,
//		*tax.Address, *tax.Telephone, userId, roleId); err != nil {
//		return err
//	}
//	return nil
//}

// 修改税务信息
func UpdateTax(userId, roleId interface{}, taxId *int64, tax *Tax) error {
	if err := corm.Db.QueryRow(`insert into account_tax(user_id,user_role,tax_no,
    company_name,bank_no,bank_name,address,telephone) values($1,$2,$3,$4,$5,$6,$7,$8)
ON CONFLICT (user_id,user_role) DO UPDATE SET tax_no=$3,
company_name=$4,bank_no=$5,bank_name=$6,address=$7,telephone=$8 RETURNING id;`, userId, roleId, *tax.TaxNo,
		*tax.CompanyName, *tax.BankNo, *tax.BankName,
		*tax.Address, *tax.Telephone).Scan(taxId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改财务信息
func UpdateFinance(userId, roleId interface{}, financeId *int64, finance *Finance) error {
	if err := corm.Db.QueryRow(`
insert into account_finance(user_id,user_role,account_name,bank_name,
    bank_no,bank_province,bank_city,sub_branch) values($1,$2,$3,$4,$5,$6,$7,$8)
ON CONFLICT (user_id,user_role) DO UPDATE SET SET account_name=$3,bank_name=$4,
bank_no=$5,bank_province=$6,bank_city=$7,sub_branch=$8 RETURNING id;`, userId, roleId, *finance.AccountName,
		*finance.BankName, *finance.BankNo, *finance.BankProvince,
		*finance.BankCity, *finance.SubBranch).Scan(financeId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改资质信息
func UpdateCompanyQualification(userId interface{}, qfId *int64, company *CompanyQualification) error {
	if err := corm.Db.QueryRow(`insert into account_qualification(user_id,business_license,
	tax_registration,organization_code) values($1,$2,$3,$4) ON CONFLICT (user_id) DO UPDATE SET id_card_positive=$2,
	id_card_negative=$3;`, userId, *company.BusinessLicense,
		*company.TaxRegistration, *company.OrganizationCode).Scan(qfId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

func UpdatePersonQualification(userId interface{}, qfId *int64, person *PersonQualification) error {
	if err := corm.Db.QueryRow(`insert into account_qualification(user_id,id_card_positive,
	id_card_negative) values($1,$2,$3) ON CONFLICT (user_id) DO UPDATE SET id_card_positive=$2,
	id_card_negative=$3;`, userId, *person.IdCardPositive, *person.IdCardNegative).Scan(qfId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改联系人信息
func UpdateDeliver(userId, roleId int, deliver *Deliver) error {
	if _, err := corm.Db.Exec(`update account_deliver set receiver=$1,email=$2,address=$3,
	telephone=$4 where user_id=$5 and user_role=$6`, *deliver.Receiver,
		*deliver.Email, *deliver.Address, *deliver.Telephone, userId, roleId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 新增联系人信息
func CreateDeliver(userId, roleId interface{}, deliverId *int64, deliver *Deliver) error {
	sql := `
	insert into account_deliver(user_id,user_role,receiver,email,address,telephone) values($1,$2,$3,$4,$5,$6) 
	ON CONFLICT (user_id,user_role) DO UPDATE SET receiver=$3,email=$4,address=$5,telephone=$6 RETURNING id`
	if err := corm.Db.QueryRow(sql, userId, roleId, *deliver.Receiver,
		*deliver.Email, *deliver.Address, *deliver.Telephone).Scan(deliverId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 账户信息修改
func UpdateAccount(userId interface{}, roleId int, accountId *int64, companyName string, account *AccountUpdate) error {
	if err := corm.Db.QueryRow(`update account_account set real_name=$1,
	company_name=$2,qq=$3,phone=$4,user_type=$5,account_type=$6,business_id=$7 
	where user_id=$8 and user_role=$9 RETURNING id`, *account.RealName, companyName,
		*account.Qq, *account.Phone, *account.UserType,
		*account.AccountType, *account.BusinessId, userId, roleId).Scan(accountId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 验证当前充值是否已经处理
func CheckRechargeStatus(rechargeId interface{}, status int) error {
	recharge := &AccountRecharge{}
	if err := corm.Db.Get(recharge, `select * from
	account_recharge where id=$1 and status=$2`, rechargeId, status); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 获取充值记录
func GetRecharge(rechargeId interface{}, roleId int, recharge *AccountRecharge) error {
	if err := corm.Db.Get(recharge, `select * from account_recharge 
		where id=$1 and user_role=$2`, rechargeId, roleId); err != nil {
		return rsp.HandlerError(err)
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
		return rsp.HandlerError(err)
	}
	return nil
}

// 充值列表
func GetRecharges(roleId int, total *int64, req *RechargeReq, recharges *AccountRecharges) error {
	orm := corm.Select(`select * from account_recharge {{sql_where}}
	ORDER BY create_date desc`).Req(req).Where("user_role", roleId).Paginate(
		req.Page, req.PageSize)
	if err := orm.Loads(recharges); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 充值审核
func ReviewRecharge(rechargeID interface{}, roleID, status int, recharge *AccountRecharge) error {
	tx := corm.Db.MustBegin()
	if _, err := tx.Exec(`update account_recharge set 
		status=$1,update_date=$2 where id=$3 and user_role=$4 and status!=1`, status,
		my_date.CurrentDateTime(), rechargeID, roleID); err != nil {
		tx.Rollback()
		return rsp.HandlerError(err)
	}
	if *recharge.Status == 0 {
		if _, err := tx.Exec(`update account_balance set 
			balance=balance+$1 where user_id=$2 and user_role=$3`, *recharge.OrderMoney,
			*recharge.UserId, roleID); err != nil {
			tx.Rollback()
			return rsp.HandlerError(err)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return rsp.HandlerError(err)
	}
	// 通知 更新缓存
	return nil

}

// 账户审核
func ReviewAccount(userId interface{}, roleId, status int) error {
	if _, err := corm.Db.Exec(`update account_account set 
		status=$1 where user_id=$2 and user_role=$3`, status, userId, roleId); err != nil {
		return rsp.HandlerError(err)

	}
	return nil
}

// 余额
func GetAccountBalance(userId interface{}, roleId int, balance *AccountBalance) error {
	if err := corm.Db.Get(balance, `select * from account_balance 
		where user_id=$1 and user_role=$2`, userId, roleId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

type AccountContractReq struct {
	UserId       *int    `form:"user_id" db:"user_id" query:"eq"`
	Sdate        *string `form:"sdate" db:"sdate" query:"gte"`
	Edate        *string `form:"edate" db:"edate" query:"lte"`
	ContractName *string `form:"contract_name" db:"contract_name" query:"like"`
	Page         *uint   `form:"page"`
	PageSize     *uint   `form:"page_size"`
}

// 获取合同
func GetAccountContract(total *int64, req *AccountContractReq, data *AccountContracts) error {
	orm := corm.Select(`select * from account_contract {{sql_where}}`).Req(req).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 合同
type AccountContract struct {
	Sdate        *string `json:"sdate" db:"sdate" validate:"required,gt=0"`                           //开始时间
	Edate        *string `json:"edate" db:"edate" validate:"required,gt=0"`                           //结束时间
	ContractName *string `json:"contract_name" db:"contract_name" validate:"required,gt=0,lte=300"`   //合同名称
	UserId       *int    `json:"user_id" db:"user_id" validate:"required,gt=0"`                       // 用户ID
	ContractUrl  *string `json:"contract_url" db:"contract_url" validate:"required,gt=0,regexp=link"` //合同地址
	Status       *int    `json:"status" db:"status" `                                                 //合同状态 {-1:"过期",1:"有效"}
}
type AccountContracts []AccountContract

// 创建合同
func CreateAccountContract(contractId *int64, contract *AccountContract) error {
	if err := corm.Db.QueryRow(`insert into account_contract(contract_name,user_id,sdate,
		edate,contract_url) values($1,$2,$3,$4,$5) RETURNING id;`, *contract.ContractName,
		*contract.UserId, *contract.Sdate, *contract.Edate, *contract.ContractUrl).Scan(contractId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 获取资质
func GetQualification(userId interface{}, data *Qualifications) error {
	if err := corm.Db.Select(data, `select * from account_qualification where user_id=$1`, userId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

type IndustryReq struct {
	UserId          *int    `form:"user_id" db:"user_id" query:"eq"`
	IndustryCode    *string `form:"industry_code" db:"industry_code" query:"eq"`
	SubIndustryCode *string `json:"sub_industry_code" db:"sub_industry_code" query:"eq"`
	Page            *uint   `form:"page"`
	PageSize        *uint   `form:"page_size"`
}

type Industry struct {
	UserId          *int    `json:"user_id" db:"user_id" validate:"required,gt=0"`
	IndustryCode    *string `json:"industry_code" db:"industry_code" validate:"required,gt=0"`
	SubIndustryCode *string `json:"sub_industry_code" db:"sub_industry_code" validate:"required,gt=0"`
	Industry        *string `json:"industry" db:"industry" validate:"required,gt=0"`
}
type Industrys []Industry

// 获取行业
func GetIndustry(total *int64, req *IndustryReq, data *Industrys) error {
	orm := corm.Select(`select * from account_industry {{sql_where}}`).Req(req).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 创建行业信息
func CreateIndustry(industryId *int64, data *Industry) error {
	if err := corm.Db.QueryRow(`insert into account_industry(user_id,
	industry_code,sub_industry_code,industry) values($1,$2,$3,$4) RETURNING id;`, *data.UserId,
		*data.IndustryCode, *data.SubIndustryCode, *data.Industry).Scan(industryId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

// 修改行业信息
func UpdateIndustry(industryId interface{}, data *Industry) error {
	if _, err := corm.Db.Exec(`update set account_industry industry_code=$1,
		sub_industry_code=$2,industry=$3 where id=$4 and user_id=$5`, *data.IndustryCode,
		*data.SubIndustryCode, *data.Industry, industryId, *data.UserId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil

}

type RechargeCostReq struct {
	RechargeId *int    `form:"recharge_id" db:"recharge_id" query:"eq"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
}

// 获取充值消耗
func GetRechargeCost(total *int64, req *RechargeCostReq, data *RechargeCosts) error {
	orm := corm.Select(`select * from log_recharge_cost {{sql_where}} 
	order by user_id,report_date desc`).Req(req).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}

type AccountCost struct {
	UserId *int     `json:"user_id" db:"user_id"`
	Cost   *float64 `json:"cost" db:"cost"`
}

type AccountCosts []AccountCost

type AccountCostReq struct {
	UserId     *int    `form:"user_id" db:"user_id" query:"eq"`
	BusinessId *int    `form:"business_id" db:"business_id" query:"user_id in (select user_id from account_account WHERE business_id=:business_id)"`
	Sdate      *string `form:"sdate" db:"report_date" query:"gte"`
	Edate      *string `form:"edate" db:"report_date" query:"lte"`
	Page       *uint   `form:"page"`
	PageSize   *uint   `form:"page_size"`
}

// 获取用户流水
func GetAccountCost(total *int64, req *AccountCostReq, data *AccountCosts) error {
	orm := corm.Select(`select COALESCE(SUM(cost),0) as cost,user_id from 
	report_base {{sql_where}} GROUP BY user_id`).Req(req).Paginate(req.Page, req.PageSize)
	if err := orm.Loads(data); err != nil {
		return rsp.HandlerError(err)
	}
	if err := orm.Total(total); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}
