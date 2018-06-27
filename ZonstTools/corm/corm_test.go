package corm

import (
	"testing"
	"fmt"
)

type User struct{
	Id int
	Name string
	ClassId int `db:"class_id"`
	Description string
}

func TestSelectLoads(t *testing.T) {
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	users := make([]User,0)
	if err := Select(`select * from public.user`).Loads(&users); err != nil {
		fmt.Println(err)
	}
	t.Log(users[0].ClassId)
	t.Log(len(users))
}
//原生的sql 语句不支持 select * from tUser where xx like xx
func TestSelectLoad(t *testing.T){
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	user := make([]User,0)
	if err := Select(`select * from public.user where class_id = :class_id and id>:id and name like :name`).Where(map[string]interface{}{
		"class_id":1,
		"id":1000,
		"name":"t4",
	}).Loads(&user); err != nil {
		fmt.Println(err)
	}
	t.Log(len(user))
}
//func TestSelectLoad(test *testing.T) {
//	campaign_ad := &CampaignAd{}
//	if err := Select(`select * from campaign_ad where id=:id and user_id=:user_id`).Where(map[string]interface{}{
//		"id":      864140,
//		"user_id": 10002,
//	}).Load(campaign_ad); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(*campaign_ad.Id)
//}
//
//func TestSelectLoad2(test *testing.T) {
//	campaign_ad := &CampaignAd{}
//	if err := Select(`select * from campaign_ad where id=:id and user_id=:user_id`).Where([][]interface{}{
//		{"id", "=", 864140},
//		{"user_id", "=", 10002},
//	}).Load(campaign_ad); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(*campaign_ad.Id, "2")
//}
//func TestSelectLoad3(test *testing.T) {
//	campaign_ad := &CampaignAd{}
//	if err := Select(`select * from campaign_ad where id=:id`).Where("id", "=", 864140).Load(campaign_ad); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(*campaign_ad.Id, "3")
//}
//
//func TestSelectLoad4(test *testing.T) {
//	campaign_ad := &CampaignAd{}
//	if err := Select(`select * from campaign_ad {{sql_where}}`).Where("id=864140").Load(campaign_ad); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(*campaign_ad.Id, "4")
//}
//
//func TestSelectByReq(test *testing.T) {
//	campaign_ads := &CampaignAds{}
//
//	req := &Req{
//		Status:   func(i int) *int { return &i }(1),
//		Page:     func(i uint) *uint { return &i }(1),
//		PageSize: func(i uint) *uint { return &i }(20),
//	}
//	if err := Select("select * from campaign_ad {{sql_where}}").Paginate(req.Page,
//		req.PageSize).Req(req).Loads(campaign_ads); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(len(*campaign_ads))
//}

//RAW CRUD
func TestExec(t *testing.T) {
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	er:=Exec("update public.user set name=$1,class_id=$2 where id=$3","ft",1,5)
	if er!=nil{
		t.Fatal(er)
	}
}
func TestExecWithId(t *testing.T) {
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	id,err:=ExecWithId("update public.user set name=$1,class_id=$2 where id=$3 returning id","ft",1,5)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
	id2,err2:=ExecWithId("insert into public.user(name,class_id) values($1,$2) returning id","ft",1)
	if err2!=nil{
		t.Fatal(err2)
	}
	t.Log(id2)
}

func TestQuery(t *testing.T) {
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	users :=make([]User,0)
	err:=Query(&users,"select * from public.user where name=$1 order by created desc","ft4")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(users)
}

func TestQueryOne(t *testing.T) {
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	user :=User{}
	err:=QueryOne(&user,"select * from public.user where id=$1","1035")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(user)
}
func TestQueryCount(t *testing.T){
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	count :=0
	err:=QueryOne(&count,"select count(*) from public.user where id>$1","1035")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(count)
}


func TestQueryOneDynamic(t *testing.T){
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	count :=0
	orderBy :=make([]string,0)
	orderBy = append(orderBy,"name","created")
	whereMap := make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","name","=",
	})
	whereMap = append(whereMap,[]string{
		"and","id","=",
	})
	t.Log(whereMap)
	err:=QueryOneDynamic(&count,"select count(*) from public.user",whereMap,nil,"",-1,-1,"ft4",1035)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(count)
}

func TestQueryDynamic(t *testing.T){
	personalDb:=GetDb("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SetDb(personalDb)
	users :=make([]User,0)
	orderBy :=make([]string,0)
	orderBy = append(orderBy,"id")
	whereMap := make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","name","=",
	})
	whereMap = append(whereMap,[]string{
		"and","id",">",
	})
	whereMap = append(whereMap,[]string{
		"and","id","<>",
	})
	t.Log(whereMap)
	err:=QueryDynamic(&users,"select * from public.user",whereMap,nil,"desc",2,0,"ft4",1035,1)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(users)
}