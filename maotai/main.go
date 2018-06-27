package main

import (
	"net/http"
	"bytes"
	"log"
	"io/ioutil"
	"fmt"
)
var client =http.Client{}
func main() {
	//1.破解配送方式
	//buf :=[]byte("Action=CalculateFreight&Weight=0.00&RegionId=0")
	//req,er := http.NewRequest("POST","https://www.emaotai.cn/SubmmitOrderHandler.aspx",bytes.NewReader(buf))
	//req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	//if er!=nil {
	//	log.Println(er.Error())
	//	return
	//}
	//resp,er:=client.Do(req)
	//if er!=nil {
	//	log.Println(er.Error())
	//	return
	//}
	//fmt.Println(resp.Header)
	//helpRead(resp)

	//2.破解支付方式
	//buf2 := []byte("Action=ProcessorPaymentMode&ModeId=5&CartTotalPrice=68&TotalPrice=68.00&Taxes=0.00")
	//req2,er := http.NewRequest("POST","https://www.emaotai.cn/SubmmitOrderHandler.aspx",bytes.NewReader(buf2))
	//req2.Header.Set("Content-Type","application/x-www-form-urlencoded")
	//if er!=nil {
	//	log.Println(er.Error())
	//	return
	//}
	//resp2,er:=client.Do(req2)
	//if er!=nil {
	//	log.Println(er.Error())
	//	return
	//}
	//fmt.Println(resp2.Header)
	//helpRead(resp2)

	//3.尝试提交
	buf3 := []byte(`__VIEWSTATE=%2FwEPDwUKLTM0MjU0NjI1Mg9kFgJmD2QWAmYPZBYCAgYPZBYYAgUPDxYCHgdWaXNpYmxlaGRkAh8PEGRkFCsBAGQCIQ9kFgJmD2QWAmYPPCsAEQMADxYEHgtfIURhdGFCb3VuZGceC18hSXRlbUNvdW50AgFkARAWABYAFgAMFCsAABYCZg9kFgQCAQ9kFghmD2QWAgIBDxAPZBYCHgV2YWx1ZQUCMTNkZGQCAQ8PFgIeBFRleHQFDOW%2Fq%2BmAkumFjemAgWRkAgIPZBYCAgEPDxYCHwQFP%2BmhuuS4sOW%2Fq%2BmAku%2B8jEVNU%2B%2B8jOW%2Bt%2BmCpueJqea1ge%2B8jOS4remTgeW%2Fq%2Bi%2FkO%2B8jOS6rOS4nOeJqea1gWRkAgMPZBYCZg8VAYEB5pys5ZWG5Z%2BO6KeG5oOF5Ya16YCJ55SoRU1T44CB6aG65Liw44CB5Lit6ZOB5b%2Br6L%2BQ44CB5b636YKm54mp5rWB44CB5Lqs5Lic54mp5rWB5Y%2BR6LSn77yM5LiN5o6l5Y%2BX5oyH5a6a5b%2Br6YCS77yM5pWs6K%2B36LCF6Kej44CCZAICDw8WAh8AaGRkAiMPZBYCZg9kFgJmDzwrABEDAA8WBB8BZx8CAgRkARAWABYAFgAMFCsAABYCZg9kFgoCAQ9kFgZmD2QWAgIBDxAPZBYCHwMFAjEyZGRkAgEPDxYCHwQFDOW7uuihjOaUr%2BS7mGRkAgIPZBYCZg8VAVw8aW1nIHNyYz0iL1N0b3JhZ2UvbWFzdGVyL2dhbGxlcnkvMjAxNzA0LzIwMTcwNDA2MjA1NzQ2Xzg4MjIuanBnIiBhbHQ9IiIgYm9yZGVyPSIwIiAvPjxiciAvPmQCAg9kFgZmD2QWAgIBDxAPZBYCHwMFAjE1ZGRkAgEPDxYCHwQFDOWGnOihjOaUr%2BS7mGRkAgIPZBYCZg8VAXE8aW1nIHNyYz0iL1N0b3JhZ2UvbWFzdGVyL2dhbGxlcnkvMjAxNzA4LzIwMTcwODI5MjAwMTEwXzAyMzkuanBnIiB0aXRsZT0i5Yac6KGMIiBhbHQ9IuWGnOihjCIgYm9yZGVyPSIwIiAvPjxiciAvPmQCAw9kFgZmD2QWAgIBDxAPZBYCHwMFAjEzZGRkAgEPDxYCHwQFDOW3peihjOaUr%2BS7mGRkAgIPZBYCZg8VAVw8aW1nIHNyYz0iL1N0b3JhZ2UvbWFzdGVyL2dhbGxlcnkvMjAxNzA0LzIwMTcwNDA2MjA1ODM4XzE3OTQuanBnIiBhbHQ9IiIgYm9yZGVyPSIwIiAvPjxiciAvPmQCBA9kFgZmD2QWAgIBDxAPZBYCHwMFATVkZGQCAQ8PFgIfBAUJ5pSv5LuY5a6dZGQCAg9kFgJmDxUB1AE8cD48YSBocmVmPSJodHRwOi8vZW1hb3RhaS5jbi9oZWxwL3Nob3ctOC5odG0iIHRhcmdldD0iX2JsYW5rIj48aW1nIHRpdGxlPSI1MjUxNSIgYm9yZGVyPSIwIiBhbHQ9IjUyNTE1IiBzcmM9Ii9TdG9yYWdlL21hc3Rlci9nYWxsZXJ5LzIwMTQxMS8yMDE0MTEwNTE1MjUxM185MzI2LmpwZyIgLz7mlK%2Fku5jlrp3lnKjnur%2FmlK%2Fku5jkvb%2FnlKjor7TmmI7vvIE8L2E%2BPC9wPmQCBQ8PFgIfAGhkZAIpD2QWAmYPZBYCZg8PFgIfAGdkFgICBw8WAh8CAgEWAmYPZBYUAgEPZBYCZg8PFgIeCEltYWdlVXJsBUgvU3RvcmFnZS9tYXN0ZXIvcHJvZHVjdC90aHVtYnM2MC82MF84OGM5YTVjNGVhNTc0YjQ2Yjg2MzM0YWJlMTZhMjI4MS5qcGdkZAIFDxYCHwRlZAIGDxUBATBkAgcPDxYCHgVNb25leSgpW1N5c3RlbS5EZWNpbWFsLCBtc2NvcmxpYiwgVmVyc2lvbj00LjAuMC4wLCBDdWx0dXJlPW5ldXRyYWwsIFB1YmxpY0tleVRva2VuPWI3N2E1YzU2MTkzNGUwODkHNjguMDAwMGRkAggPFQEAZAIJDxYCHwQFATFkAgsPFgIfBGVkAgwPFQEBMWQCDQ8PFgIfBigrBAc2OC4wMDAwZGQCDw8PFgQeC05hdmlnYXRlVXJsBRcvRmF2b3VyYWJsZS9zaG93LTAuYXNweB8EZWRkAisPDxYCHwQFBDAuMDBkZAI1D2QWAmYPZBYEZg8PFgIfAGhkFgICAQ88KwAJAGQCAg8PFgIfAGhkFgICAQ88KwAJAGQCNw8WAh8AaBYCZg9kFgICAQ9kFgICAQ8QFgYeDURhdGFUZXh0RmllbGQFC0Rpc3BsYXlUZXh0Hg5EYXRhVmFsdWVGaWVsZAUJQ2xhaW1Db2RlHwFnEBUBABUBATAUKwMBZxQrAQBkAkEPDxYCHwQFATBkZAJJDw8WAh8GKCsEBzY4LjAwMDBkZAJPDw8WAh8EBQI2OGRkAmMPDxYCHwYoKwQHNjguMDAwMGRkGAMFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYMBUhTdWJtbWl0T3JkZXIkQ29tbW9uX1NoaXBwaW5nTW9kZUxpc3QkXyRncmRTaGlwcGluZ01vZGUkY3RsMDIkcmFkaW9CdXR0b24FSFN1Ym1taXRPcmRlciRDb21tb25fU2hpcHBpbmdNb2RlTGlzdCRfJGdyZFNoaXBwaW5nTW9kZSRjdGwwMiRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwMiRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwMiRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwMyRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwMyRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwNCRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwNCRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwNSRyYWRpb0J1dHRvbgVGU3VibW1pdE9yZGVyJGdyZF9Db21tb25fUGF5bWVudE1vZGVMaXN0JF8kZ3JkUGF5bWVudCRjdGwwNSRyYWRpb0J1dHRvbgUVU3VibW1pdE9yZGVyJGNoa0FncmVlBRNTdWJtbWl0T3JkZXIkY2hrVGF4BTRTdWJtbWl0T3JkZXIkZ3JkX0NvbW1vbl9QYXltZW50TW9kZUxpc3QkXyRncmRQYXltZW50DzwrAAwDBhUBB0dhdGV3YXkHFCsABBQrAAEFB2NjYl9wYXkUKwABBQphYmNfcGF5X3BjFCsAAQUHaWNiY19wYxQrAAEFCWFsaXBheV9wYwgCAWQFNlN1Ym1taXRPcmRlciRDb21tb25fU2hpcHBpbmdNb2RlTGlzdCRfJGdyZFNoaXBwaW5nTW9kZQ88KwAMAwYVAQZNb2RlSWQHFCsAARQrAAECDQgCAWQ%3D&radaddresstype=taobao&SubmmitOrder%24txtShipTo=&ddlRegions1=&ddlRegions2=&ddlRegions3=&regionSelectorValue=0&regionSelectorNull=-%E8%AF%B7%E9%80%89%E6%8B%A9-&SubmmitOrder%24txtAddress=&SubmmitOrder%24txtZipcode=&SubmmitOrder%24txtTelPhone=&SubmmitOrder%24txtShippingId=&SubmmitOrder%24txtCellPhone=18970937628&SubmmitOrder%24tbActvityProductID=%2C627&SubmmitOrder%24tbActiviPrice=%2C68.00&SubmmitOrder%24txtMessage=&SubmmitOrder%24chkAgree=on&SubmmitOrder%24txtInvoiceTitle=&SubmmitOrder%24csessionid=&SubmmitOrder%24sig=&SubmmitOrder%24token=&SubmmitOrder%24scene=&SubmmitOrder%24btnCreateOrder=%E7%A1%AE%E8%AE%A4%E6%8F%90%E4%BA%A4&SubmmitOrder%24htmlCouponCode=&SubmmitOrder%24inputPaymentModeId=&SubmmitOrder%24inputShippingModeId=&SubmmitOrder%24hdbuytype=&SubmmitOrder%24inputInvoiceId=&productId=&qty=&logoFile=&pingshengFile=&userLogoFile=&userPingshengFile=&dzjNum=&dzjMinNum=`)
	req3,er := http.NewRequest("POST","https://www.emaotai.cn/SubmmitOrder.aspx?productId=627&isyushou=1&qty=1",bytes.NewReader(buf3))
	req3.Header.Set("Content-Type","application/x-www-form-urlencoded")
	if er!=nil {
		log.Println(er.Error())
		return
	}
	resp3,er:=client.Do(req3)
	if er!=nil {
		log.Println(er.Error())
		return
	}
	fmt.Println(resp3.Header)
	helpRead(resp3)
}

func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}