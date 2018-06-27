package main

import (
	"github.com/gin-gonic/gin"
	"flag"
	"coupon/xormTool"
	"log"
	"coupon/Dao"

	"net/http"
	"github.com/rs/cors"
)
var port string
type Result struct{
	Status int `json:"status"`
	Data string `json:"data"`
}
func (rs *Result) New(status int, coupon string) *Result{
	rs.Status = status
	rs.Data = coupon
	return  rs
}

func init() {
	flag.StringVar(&port, "port", ":8076", "[example] go run main.go -port ':8076'")
	flag.StringVar(&Dao.DataSource, "db", "postgres://postgres:123@localhost:5432/xm_coupon?sslmode=disable", "[example] go run main.go -db 'postgres://postgres:123@localhost:5432/xm_coupon?sslmode=disable'")
	flag.Parse()

	log.SetFlags(log.LstdFlags|log.Llongfile)

	//设置数据库配置
	xormTool.DataSource(Dao.DataSource)
	xormTool.Config(true,2000,1000)
}
func main(){
 router:=gin.New()
 router.Use(gin.Recovery())

 router.GET("/xmCoupon",Generate)
 http.ListenAndServe(port, cors.AllowAll().Handler(router))
}

func Generate(c *gin.Context){
	var rs = Result{}
	couponStr :=Dao.GetCoupon()
	if couponStr!=""{
		c.JSON(200,rs.New(1,couponStr))
		return
	}
	c.JSON(400,rs.New(-1,"获取失败，coupon为空"))
	return
}