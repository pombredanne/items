package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

var data Data
func main(){
	router:=gin.Default()
	router.Use(Validate())  //使用validate()中间件身份验证
	router.GET("/",Service1)

	router.POST("/Post",Service2)
	router.Run(":8989")  //localhost:8989/
}

type Data struct{
	Username string `form:"username" binding:"required" json:"username" binding:"required"`
	Password string `form:"password" binding:"required" json:"password" binding:"required"`

}
func Service2(c *gin.Context){
    fmt.Println("in")
	c.Bind(&data)
	fmt.Println("data:",data.Username,data.Password)
	c.JSON(http.StatusOK,gin.H{"message":"你好，欢迎你"})
}

func Service1(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"你好，欢迎你"})
}

func Validate() gin.HandlerFunc{
	return func(c *gin.Context){
		//这一部分可以替换成从session/cookie中获取，
	     c.Bind(&data)

		fmt.Println("data:",data.Username,data.Password)
		if data.Username=="ft" && data.Password =="123"{
			c.JSON(http.StatusOK,gin.H{"message":"身份验证成功"})
			c.Next()
		}else {
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"身份验证失败"})
		}
	}
}