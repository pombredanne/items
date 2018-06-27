package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"github.com/rs/cors"
	"travel/models/vo"
	"travel/DAO"
	"travel/Consts"
	"reflect"
	"strconv"
	"fmt"
	"time"
	"travel/Common/TokenManager"
	"container/list"
	"net/http"
	"bytes"
	"encoding/csv"
	"golang.org/x/text/transform"
	"strings"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
)
func main() {
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.Use(AllowAll())

	//travel
	router.GET("/travel", VisitorHandler)
	router.POST("/travel/user/create",UserHandler)

	//西瓜
	router.GET("/watermelon",WatermelonHandler)

	//获取报表
	router.POST("/login",TokenGet)
	//router.Use(Validate())
	router.GET("/download",TravelDownLoadHandler)
	http.ListenAndServe(":8087", cors.AllowAll().Handler(router))
	//router.Run(":8087")
}
func AllowAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		c.Writer.Header().Set("content-type", "application/json")             //返回数据格式是json
	}
}
func VisitorHandler(c *gin.Context) {
	visitorVO :=vo.Visitor{}
	visitorVO.CId =c.DefaultQuery("cid","")
	//ob:=c.Request.Header.Get("REMOTE_ADDR")
	//fmt.Println("真实解析的ip",ob)
	visitorVO.Ip=c.ClientIP()
	//visitorVO.Ip = GetIPAndPort(ob)
	//fmt.Println("客户端的ip:",visitorVO.Ip)
	DAO.SaveVisitor(visitorVO)
	c.JSON(200,Consts.OKInsert)
	return
}

func UserHandler(c *gin.Context){
	user :=vo.User{}
	c.Bind(&user)
	if !CheckEmpty(user){
		c.JSON(200,Consts.InfoLacking)
		return
	}
	DAO.SaveUser(user)
	c.JSON(200,Consts.OKInsert)
	return
}

func WatermelonHandler(c *gin.Context){
	cid,err :=strconv.Atoi(c.DefaultQuery("cid","0"))
	if err!=nil {
		c.JSON(200,gin.H{"msg":"请输入正确的cid，样式1,2,3"})
		return
	}
	actionType,err :=strconv.Atoi(c.DefaultQuery("actionType","0"))
	if err!=nil {
		c.JSON(200,gin.H{"msg":"请输入正确的actionType，样式1,2,其中1表示点击，2表示展示"})
		return
	}
	//ip :=c.Request.Header.Get("REMOTE_ADDR")
	ip:=c.ClientIP()
	fmt.Println("ip:",ip)
	_,err=DAO.InsertWaterMelon(ip,cid,actionType)
	if err!=nil {
		c.JSON(200,gin.H{"msg":fmt.Sprintf("存入数据库时发生了错误%v",err)})
		return
	}
	c.JSON(200,gin.H{"msg":"插入成功"})
}

func CheckEmpty(input interface{}) bool {
	uType := reflect.TypeOf(input)
	uValue := reflect.ValueOf(input)
	fieldNum := uType.NumField()

	for i:=0;i<fieldNum;i++{
		valueStr := uValue.Field(i).String()
		if  valueStr==""{
			return false
		}
	}
	return true
}
func GetIPAndPort(remoteAddr string)(string){
	return remoteAddr
}

func TokenGet(c *gin.Context){
	feedBack :=vo.FeedBack{}
	login := vo.Login{}
	jwts :=TokenManager.Jwts{}
	c.Bind(&login)
	if login.UserName=="邹梦君" && login.Password=="123456"{
		jToken,err := TokenManager.Token.BasicToken(Consts.Secret)
		exp := time.Now().Add(2*time.Hour).Unix()
		if err!=nil{
			fmt.Println(err)
			c.JSON(200,Consts.ResponseTokenGenerateError)
			return
		}
		feedBack.Msg="成功获取token"
		feedBack.Status=1
		feedBack.Token=jToken

		jwts.Token=jToken
		jwts.EndTimeUnix=exp

		go func() {
			fmt.Println("开始登记token")
			var eTemp *list.Element
			eTemp = TokenManager.GetTokenRegister().PushBack(jwts)
			//同一个用户不可能存在并发的,所以不用加锁
			//不同的用户key不同，map不存在并发写
			TokenManager.ListMap[login.UserName] = eTemp
			fmt.Println("登记token成功:", jwts)
		}()
		c.JSON(200,feedBack)
		return
	}
	c.JSON(200,Consts.ResponseWrongAcount)
}
func TravelDownLoadHandler(c *gin.Context){
	//feedBack :=vo.FeedBack{}
	startDateStr := c.DefaultQuery("startDate",time.Now().Format("2006-01-02"))
	endDateStr := c.DefaultQuery("endDate",time.Now().Format("2006-01-02"))

	users,err:=DAO.GetUserByDate(startDateStr,endDateStr)
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,fmt.Sprintf("发生了错误%v",err))
		return
	}

	pourData(users,c)
	//feedBack.Status=1
	//feedBack.Msg="success"
	//feedBack.Data=users
	//c.JSON(200,feedBack)
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		r:=c.Request
		fmt.Println(r)
		cookie,err:=c.Request.Cookie("Authorization")
		if cookie==nil{
			c.JSON(200,"未找到该cookie")
		}
		fmt.Println("收到cookie.Autorization",cookie.Value)
		if err!=nil{
			fmt.Println(err)
			c.JSON(200,"读取cookie出错")
			return
		}
		if JWTToken := c.Request.Header.Get("Authorization");JWTToken!=""{
			rollingFunc(JWTToken,c)
		}else if JWTToken := cookie.Value;JWTToken!=""{
			rollingFunc(JWTToken,c)
		}else {
			c.JSON(200, Consts.ResponseTokenNotFound)
			c.Abort()
			return
		}
	}
}
func rollingFunc(JWTToken string,c *gin.Context){
	fmt.Println("收到的token是:",JWTToken)
	tokenRigister :=TokenManager.GetTokenRegister()
	if !TokenManager.ListContains(tokenRigister,JWTToken){
		fmt.Println("token未被登记")
		c.Abort()
		c.JSON(200,Consts.ResponseTokenValidateError)
		return
	}
	legal,err:=TokenManager.Token.IsLegal(JWTToken,Consts.Secret)
	if err!=nil{
		fmt.Println(err)
		c.Abort()
		c.JSON(200,Consts.ResponseTokenValidateError)
		return
	}
	if !legal{
		c.Abort()
		c.JSON(200,Consts.ResponseTokenValidateWrong)
		return
	}
	c.Next()
	return
}

func structs2StringArr(users []DAO.User) [][]string{
	var userArr =make([][]string,0)
	var user = DAO.User{}
	for i:=0;i<len(users);i++{
		user = users[i]
		userArr = append(userArr,[]string{user.Name,user.Province,user.City,user.Address,user.Tdate,user.Phone,strconv.Itoa(user.Status)})
	}
	return userArr
}
func pourData(users []DAO.User,c *gin.Context){
	fileName := "test.csv"
	fmt.Println(users)
	usersStrArray :=structs2StringArr(users)
	categoryHeader := []string{"姓名","省份","城市","详细地址","填写日期","联系方式","是否联系"}
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	wr.Write(categoryHeader)
	for i := 0; i < len(usersStrArray); i++ {
		wr.Write(usersStrArray[i])
	}
	wr.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	//该转码不做，有可能乱码哦
	tet,_:= UTF82GBK(b.Bytes())
	c.String(200, tet)
}

func UTF82GBK(src []byte) (string, error) {
	reader:= transform.NewReader(strings.NewReader(string(src)), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}