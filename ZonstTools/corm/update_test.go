package corm

import "testing"

//更新tUser中的name，age列，条件是id>2
func TestUpdate(t *testing.T) {

	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	updateBuilder :=Update("public.user").Set(map[string]interface{}{
		"name":"FT",
		"class_id":2,
	}).Where("id","=","5")   //"id","between","fff"挤了
	  								      // id,like,fff 占位符不对
	  								      //id,not in ,fff 不行 //args大于3个没处理where
	updateBuilder.Build()
	sql,err:=updateBuilder.Exec()
	t.Log(sql,err)
}