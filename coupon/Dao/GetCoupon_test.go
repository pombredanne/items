package Dao

import (
	db "coupon/xormTool"
	"testing"
)

func TestInsertGroup(t *testing.T) {
	db.DataSource("postgres://postgres:123@localhost:5432/xm_coupon?sslmode=disable")
	db.DefaultConfig()
	var coupons = []string{
		"test coupon_1",
	}
	em := InsertGroup(coupons)
	t.Log(em)
}


func TestGetCoupon(t *testing.T) {
	db.DataSource("postgres://postgres:123@localhost:5432/xm_coupon?sslmode=disable")
	db.DefaultConfig()
	cp :=GetCoupon()
	t.Log(cp)
}