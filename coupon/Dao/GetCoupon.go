package Dao

import (
	db "coupon/xormTool"
	"log"
)

var DataSource string
func GetCoupon() string{
	var couponId string
	var err error
	con:=db.GetDb()
	session :=con.NewSession()
	session.Begin()
	//获取未使用过的卡号
	_,err=session.SQL("select coupon_id from xm_coupon where status=? limit 1",1).Get(&couponId)
	if err!=nil {
		log.Println(err.Error())
		session.Rollback()
		return ""
	}
	_,err=session.SQL("update xm_coupon set status=? where coupon_id=?",2,couponId).Execute()
	if err!=nil {
		log.Println(err.Error())
		session.Rollback()
		return ""
	}
	session.Commit()

	return couponId
}


func InsertGroup(coupons []string) string{
	con:=db.GetDb()
	session :=con.NewSession()
	session.Begin()
	for _,coupon:=range coupons {
		_,err:=session.SQL("insert into xm_coupon(coupon_id) values(?)",coupon).Execute()
		if err!=nil {
			session.Rollback()
			return "出错回滚"+err.Error()
		}
	}
	er:=session.Commit()
	if er!=nil{
		return er.Error()
	}
	return ""
}