package TokenManager

import (
	"time"
	"container/list"
	"fmt"
	"sync"
	"github.com/fwhezfwhez/jwt"
)

//全局token对象
var Token = jwt.GetToken()

type Jwts struct {
	Token       string
	EndTimeUnix int64
}

var lock = sync.Mutex{}

//用来存放生成的token的list队列
var tokenRegister = list.New()
//限定list大小
var MaxSize = 5000
//默认过期时间
var tokenValidPeriod = 1 * time.Hour
//存放list对象类型 key:userName,value:*Element实例
//因为登出的时候，需要用这个map来定位Element
var ListMap = make(map[string]*list.Element, 5000)
func init() {
	jwts :=Jwts{}
	jwts.Token="eyJleHAiOiIxNTIyNzM3ODQwIn0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.VTrqu1rH8e2JjVCf6owrNIhqbO4rJe+89Z0iuWdv+pc="
	jwts.EndTimeUnix=time.Now().Add(2*time.Hour).Unix()
	tokenRegister.PushBack(jwts)
	//开启协程，管理tokenRegister
	go func() {
		//开启管理
		fmt.Println("开启管理Register")
		for {
			//防止抱死时间片不放
			time.Sleep(1*time.Second)
			//管理大小
			if tokenRegister.Len() ==MaxSize{
				tokenRegister = list.New()
			}
			//管理时效
			if tokenRegister.Len() != 0 {
				nowUnix := time.Now().Unix()
				frontElem := tokenRegister.Front()
				endUnix := frontElem.Value.(Jwts).EndTimeUnix
				if nowUnix > endUnix {
					//过期了
					tokenRegister.Remove(frontElem)
				}
			}
		}
	}()
}

func GetTokenRegister() *list.List {
	return tokenRegister
}

func ListContains(l *list.List,jwt string)bool{
	if l.Len()==0 {
		return false
	}
	var n *list.Element
	for e := l.Front(); e != nil; e = n {
		jTemp:= e.Value.(Jwts).Token
		if jTemp ==jwt{
			return true
		}
		n = e.Next()
	}
	return false
}