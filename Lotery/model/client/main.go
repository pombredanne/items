package main

import (
	"net/http"
	"encoding/json"
	"log"
	"Lotery/model/vo"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"reflect"
	"sort"
)

var host1 = "http://clientsrv.wangzheka.cn/v1"
var host2 = "http://clientsrv.duoduocp.cn/v1"
var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozNjAwMDAwNDAsImV4cCI6MTUyOTQ2OTc3M30.awzUf1e9OrSQpHxxH7mjzEWPea5QQ6Q1ZhMpWSsE2tQ"
var client = http.Client{}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func main() {
	//1.测试对阵信息,rs:ok
	//test_match_query()

	//2.测试竞彩下注,{"errmsg":"unauthorized","errno":"3"}
	//test_order()

	//3.测试比赛结果 ok
	//test_match_result("0")

	//4.测试订单详情，{"errmsg":"unauthorized","errno":"3"}
	test_game_orders()

	//5.测试近期战绩详情,ok
	//test_match_history()

	//6.测试比赛积分详情,ok
	//test_match_point()

	//7.测试比赛赔率详情,ok
	//test_match_rate()
}

func commonHttp(method string, url string, data interface{}) {
	var buf string
	var req *http.Request
	var er error
	if data != nil {
		buf, er = toParam(data)
		if er != nil {
			log.Println(er.Error())
			return
		}
		fmt.Println("发送了:", string(buf))
		req, er = http.NewRequest(method, url, strings.NewReader(buf))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("x-cdd-jwt", token)
		if er != nil {
			log.Println(er.Error())
			return
		}
	} else {
		req, er = http.NewRequest(method, url, nil)
		if er != nil {
			log.Println(er.Error())
			return
		}
	}

	resp, er := client.Do(req)
	defer resp.Body.Close()
	rs, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		log.Println(er.Error())
		return
	}
	fmt.Println("收到了:", string(rs))
}

func test_match_query() {
	q := vo.MatchQuery{
		Game:      "FT",
		LotteryId: "FT01",
	}
	commonHttp("POST", host2+"/client/game/match_query", q)
}

func test_order() {
	var orderInfo = vo.GameOrderInfo{
		BetCodes: "20180402*1*050*0",
		BetMulti: "1",
		BetMoney: "2",
		BetType:  "110000000",
	}
	temp, er := json.Marshal(orderInfo)
	if er != nil {
		log.Println(er.Error())
		os.Exit(-1)
	}
	orderinfo := string(temp)
	fmt.Println(orderinfo + orderinfo)
	var q = vo.Order{
		Game:      "FT",
		LotteryId: "FT04",
		OrderInfo: orderinfo,
		UserId:    "360000001",
		CouponId:  "4",
	}

	commonHttp("POST", host2+"/client/game/order", q)
}

func test_match_result(all string) {
	q := vo.MatchResult{
		Game: "FT",
		All:  all,
	}
	commonHttp("POST", host2+"/client/game/match_result", q)
}

func test_game_orders() {
	q := vo.GameOrders{
		UserId:      "360000006",
		OrderStatus: "-1",
	}
	commonHttp("POST", host2+"/client/user/game_orders", q)
}

func test_match_history() {
	q := vo.MatchHistory{
		Game:        "FT",
		MatchDate:   "2018-05-11",
		MatchNumber: "20180511*5*001",
	}
	commonHttp("POST", host2+"/client/game/match_history", q)
}

func test_match_point() {
	q := vo.MatchPoint{
		Game:        "FT",
		MatchDate:   "2018-05-17",
		MatchNumber: "20180517*4*008",
	}
	commonHttp("POST", host2+"/client/game/match_point", q)
}

func test_match_rate() {
	q := vo.MatchPoint{
		Game:        "FT",
		MatchDate:   "2018-05-17",
		MatchNumber: "20180517*4*008",
	}
	commonHttp("POST", host2+"/client/game/match_rate", q)
}

func toParam(vx interface{}) (string, error) {
	var result string
	var tagName_FieldValue = make(map[string]string)
	var SortedArr = make([]string, 0)
	var tagValueTemp string
	var valueStrTemp string

	vType := reflect.TypeOf(vx)
	vValue := reflect.ValueOf(vx)
	for i := 0; i < vType.NumField(); i++ {
		tagValueTemp = filtTag(vType.Field(i).Tag.Get("json"))
		//设置过滤
		if tagValueTemp == "" || tagValueTemp == "," || tagValueTemp == "omitempty" {
			continue
		}
		valueStrTemp = vValue.Field(i).String()
		if valueStrTemp == "" {
			continue
		}
		tagName_FieldValue[tagValueTemp] = valueStrTemp
		SortedArr = append(SortedArr, tagValueTemp)
	}
	sort.Strings(SortedArr)

	for i, v := range SortedArr {
		if i == 0 {
			result = result + v + "=" + tagName_FieldValue[v]
			continue
		}
		result = result + "&" + v + "=" + tagName_FieldValue[v]
	}
	return result, nil
}

func filtTag(tag string) string {
	if strings.Contains(tag, ",") {
		return strings.Split(tag, ",")[0]
	} else {
		return tag
	}
}
