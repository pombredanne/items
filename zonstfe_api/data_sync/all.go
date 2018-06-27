package main

import (
	"time"
	"math/rand"
	"fmt"
	"zonstfe_api/common/utils/myfile"
	"zonstfe_api/common/options"
)

//import (
//	"fmt"
//	"time"
//	"encoding/json"
//	"zonstfe_api/common/options"
//	"math/rand"
//
//)
//
//
//// 为每个campaign 创建素材集合
//func TestAd()  {
//	campaigns := &Campaigns{}
//	err:=Pgx.Select(campaigns, `select id,name from campaign_campaign`)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	for _,item:=range *campaigns{
//		_,err:=Pgx.Exec(`insert into campaign_ad(name,ad_type,ad_size,
//	creative_set_id,creative,campaign_id,status) values($1,$2,$3,$4,$5,$6,$7)`,*item.Name,"graphic","banner",1,
//		"{\"banner\": \"https://res1.applovin.com/o6a2f144/a73a4aa289b2aec5437fadaab3a71f3aa4cb0ca2_raw11.gif\"}",*item.Id,1)
//		if err!=nil{
//			fmt.Println(err)
//		}
//		_,err2:=Pgx.Exec(`insert into campaign_ad(name,ad_type,ad_size,
//	creative_set_id,creative,campaign_id,status) values($1,$2,$3,$4,$5,$6,$7)`,*item.Name,"graphic","interstitial",1,
//			"{\"inter_l\": \"https://res1.applovin.com/o6a2f144/507e0af6100b39df85a9cfc989d92a3d4132b4a3_v21_phone.jpg\",\"inter_p\": \"https://res1.applovin.com/o6a2f144/507e0af6100b39df85a9cfc989d92a3d4132b4a3_v21_phone.jpg\"}",*item.Id,1)
//		if err2!=nil{
//			fmt.Println(err2)
//		}
//		_,err3:=Pgx.Exec(`insert into campaign_ad(name,ad_type,ad_size,
//	creative_set_id,creative,campaign_id,status) values($1,$2,$3,$4,$5,$6,$7)`,*item.Name,"video","video_l",1,
//			"{\"video\": \"https://res1.applovin.com/o6a2f144/76f2badb8bd8fa7ba6dcd738edbbf30f329eca6e_v23_phone.mp4\", \"inter_l\": \"https://res1.applovin.com/o6a2f144/507e0af6100b39df85a9cfc989d92a3d4132b4a3_v21_phone.jpg\",\"inter_p\": \"https://res1.applovin.com/o6a2f144/507e0af6100b39df85a9cfc989d92a3d4132b4a3_v21_phone.jpg\"}",*item.Id,1)
//		if err3!=nil{
//			fmt.Println(err3)
//		}
//		_,err4:=Pgx.Exec(`insert into campaign_ad(name,ad_type,ad_size,
//	creative_set_id,creative,campaign_id,status) values($1,$2,$3,$4,$5,$6,$7)`,*item.Name,"video","video_p",1,
//			"{\"video\": \"https://res1.applovin.com/o6a2f144/76f2badb8bd8fa7ba6dcd738edbbf30f329eca6e_v23_phone.mp4\", \"inter_l\": \"https://res1.applovin.com/o6a2f144/507e0af6100b39df85a9cfc989d92a3d4132b4a3_v21_phone.jpg\",\"inter_p\": \"https://res1.applovin.com/o6a2f144/507e0af6100b39df85a9cfc989d92a3d4132b4a3_v21_phone.jpg\"}",*item.Id,1)
//		if err4!=nil{
//			fmt.Println(err4)
//		}
//	}
//
//}
//
////查询spider_app 获取 os name
//func TestData() {
//	fmt.Println("start")
//	genres := map[string]string{
//		"图书":    "1018",
//		"商务":    "1009",
//		"商品指南":  "1005",
//		"教育":    "1007",
//		"娱乐":    "1012",
//		"财务":    "1009",
//		"美食佳饮":  "1011",
//		"游戏":    "1017",
//		"健康健美":  "1006",
//		"生活":    "1004",
//		"报刊杂志":  "1018",
//		"医疗":    "1006",
//		"音乐":    "1013",
//		"导航":    "1010",
//		"新闻":    "1018",
//		"摄影与录像": "1003",
//		"效率":    "1001",
//		"参考":    "1004",
//		"购物":    "1005",
//		"社交":    "1016",
//		"体育":    "1006",
//		"贴纸":    "1017",
//		"旅游":    "1010",
//		"工具":    "1004",
//		"天气":    "1004",
//	}
//	my_genres := map[string][]string{
//		"1001": {"100101", "100102", "100103", "100104", "100105"},
//		"1002": {"100201", "100202", "100203", "100204", "100205"},
//		"1003": {"100301", "100302", "100303"},
//		"1004": {"100401", "100402", "100403", "100404"},
//		"1005":{"100501", "100502", "100503", "100504", "100505"},
//		"1006": {"100601", "100602", "100603", "100604", "100605"},
//		"1007": {"100701", "100702", "100703", "100704", "100705", "100706"},
//		"1008": {"100801", "100802", "100803"},
//		"1009": {"100901", "100902", "100903", "100904"},
//		"1010": {"101001", "101002", "101003", "101004", "101005", "101006", "101007"},
//		"1011": {"101101", "101102"},
//		"1012": {"101201", "101202"},
//		"1013": {"101301", "101302", "101303", "101304"},
//		"1014": {"101401", "101402", "101403", "101404", "101405"},
//		"1015": {"101501", "101502"},
//		"1016": {"101601", "101602", "101603", "101604", "101605", "101606", "101607"},
//		"1017": {"101701", "101702", "101703", "101704", "101705", "101706", "101707", "101708", "101709", "101710", "101711", "101712"},
//		"1018": {"101801", "101802", "101803", "101804", "101805"},
//	}
//	type app struct {
//		Os    *string `db:"os"`
//		Name  *string `db:"name"`
//		AppId *string `db:"app_id"`
//		Genre *string `db:"genre"`
//		Url   *string `db:"url"`
//	}
//	type apps []app
//
//	my_app := &apps{}
//	err:=Pgx.Select(my_app, `select os,name,app_id,genre,info->'trackViewUrl' as url from spider_app limit 100`)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	for _, item := range *my_app {
//		rand.Seed(time.Now().UTC().UnixNano())
//		n := rand.Int() % len(my_genres[genres[*item.Genre]])
//		c_genre := my_genres[genres[*item.Genre]][n]
//		user_id := GenerateRangeNum(1000, 1100)
//		budget_day := GenerateRangeNum(1000, 10000)
//		bidding_min := GenerateRangeNum(1, 3)
//		bidding_max := GenerateRangeNum(3, 10)
//		speed_list := []int{1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2}
//		speed_value := speed_list[rand.Int()%len(speed_list)]
//		rand.Seed(time.Now().UTC().UnixNano())
//		freq_open_value := rand.Intn(10)
//		freq_map := make(map[string]interface{}, 0)
//		if freq_open_value>5 {
//			freq_type_value := rand.Intn(10)
//
//			if freq_type_value>5 {
//				freq_map = map[string]interface{}{
//					"open": true,
//					"num":  100000,
//					"type": "day",
//				}
//			} else {
//				freq_map = map[string]interface{}{
//					"open": true,
//					"num":  10000,
//					"type": "hour",
//				}
//			}
//
//		} else {
//			freq_map = map[string]interface{}{
//				"open": false,
//				"num":  1000,
//				"type": "day",
//			}
//		}
//		target_map := make(map[string]interface{}, 0)
//		target_map["vendor"] = map[string]interface{}{
//			"open": true,
//			"list": []string{"zonst"},
//		}
//		//
//		carrier_open_value := rand.Intn(10)
//
//		if carrier_open_value>5 {
//			target_map["carrier"] = map[string]interface{}{
//				"open": true,
//				"list": []string{"1", "2", "3"},
//			}
//		} else {
//			target_map["carrier"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//		//
//		rand.Seed(time.Now().UTC().UnixNano())
//		network_open_value := rand.Intn(10)
//		if network_open_value >5 {
//			target_map["network"] = map[string]interface{}{
//				"open": true,
//				"list": []string{"1", "2", "3", "4"},
//			}
//		} else {
//			target_map["network"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//		//
//		rand.Seed(time.Now().UTC().UnixNano())
//		os_version_open_value := rand.Intn(10)
//
//		if os_version_open_value >5 {
//			if *item.Os == "ios" {
//				target_map["os_version"] = map[string]interface{}{
//					"open": true,
//					"list": []string{"7", "8", "9", "10", "11"},
//				}
//			} else {
//				target_map["os_version"] = map[string]interface{}{
//					"open": true,
//					"list": []string{"4", "5", "6", "7"},
//				}
//			}
//
//		} else {
//			target_map["os_version"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//		//
//		rand.Seed(time.Now().UTC().UnixNano())
//		day_parting_open_value := rand.Intn(10)
//
//		if day_parting_open_value>5 {
//			target_map["day_parting"] = map[string]interface{}{
//				"open": true,
//				"list": []string{"0", "1", "2", "3", "4", "5", "6"},
//			}
//		} else {
//			target_map["day_parting"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//		//
//		rand.Seed(time.Now().UTC().UnixNano())
//
//		device_type_open_value := rand.Intn(10)
//
//		if device_type_open_value >5 {
//			target_map["device_type"] = map[string]interface{}{
//				"open": true,
//				"list": []string{"1", "2"},
//			}
//		} else {
//			target_map["device_type"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//		//
//		rand.Seed(time.Now().UTC().UnixNano())
//
//		geo_code_open_value := rand.Intn(10)
//
//		if geo_code_open_value >5 {
//			geocode_list := make([]string, 0)
//			for _, counrty := range options.CountryList {
//				rand.Seed(time.Now().UTC().UnixNano())
//
//				counrty_select_value := rand.Intn(10)
//				if counrty_select_value >5 {
//					geocode_list = append(geocode_list, counrty)
//				}else {
//					if counrty != "000000000" {
//						for _, province := range options.ProvinceList {
//							rand.Seed(time.Now().UTC().UnixNano())
//
//							province_select_value := rand.Intn(10)
//							// 模拟全选概率
//							if province_select_value >5 {
//								geocode_list = append(geocode_list, province)
//							} else {
//								if len(options.ProvinceCityMap[province])>5{
//									n2 := rand.Int() % (len(options.ProvinceCityMap[province]) - 1)
//									for n2 > 0 {
//										geocode_list = append(geocode_list, options.ProvinceCityMap[province][n2-1])
//										n2 -= n2
//									}
//								}else {
//									geocode_list = append(geocode_list, province)
//								}
//
//							}
//						}
//
//
//					}
//
//
//				}
//
//			}
//			target_map["geo_code"] =  map[string]interface{}{
//				"open": true,
//				"list": geocode_list,
//			}
//		} else {
//			target_map["geo_code"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//
//		//
//		rand.Seed(time.Now().UTC().UnixNano())
//
//		app_category_open_value := rand.Intn(10)
//
//		if app_category_open_value >5 {
//			category_list := make([]string, 0)
//			// 先模拟每个分类是否选
//			for key, value := range my_genres {
//				rand.Seed(time.Now().UTC().UnixNano())
//
//				app_category_select_value :=rand.Intn(10)
//
//
//				if app_category_select_value>5 {
//					// 再模拟分类是否是全选的可能性 预设随机Index 等于最后一个元素
//					if len(value) > 0 {
//						n1 := rand.Int() % len(value)
//						if n1+1 == len(value) {
//							category_list = append(category_list, key)
//						} else {
//							n2 := rand.Int() % (len(value) - 1)
//							for n2 > 0 {
//								category_list = append(category_list, value[n2-1])
//								n2 -= n2
//							}
//						}
//					}
//
//				}
//
//			}
//			target_map["app_category"] = map[string]interface{}{
//				"open": true,
//				"list": category_list,
//			}
//		} else {
//			target_map["app_category"] = map[string]interface{}{
//				"open": false,
//				"list": []string{},
//			}
//		}
//		target_map["segment"] = map[string]interface{}{
//			"open": false,
//			"list": []string{},
//		}
//		url_map := map[string]string{
//			"jump_url":         *item.Url,
//			"deep_link_url":    "",
//			"tracking_clk_url": "",
//			"tracking_imp_url": "",
//		}
//		url_map_byte, _ := json.Marshal(url_map)
//		freq_map_byte, _ := json.Marshal(freq_map)
//		target_map_byte, _ := json.Marshal(target_map)
//		_,err:=Pgx.Exec(`
//		insert into campaign_campaign(name,user_id,bundle_id,app_platform,app_category,
//		budget_day,bidding_max,bidding_min,bidding_type,freq,targeting,url,speed,status) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
//		`, *item.Name, user_id, *item.AppId, *item.Os, c_genre, budget_day, bidding_max, bidding_min, "cpm", string(freq_map_byte), string(target_map_byte), string(url_map_byte), speed_value, 1)
//		if err!=nil{
//			fmt.Println(err)
//		}
//
//	}
//}
//
//func GenerateRangeNum(min, max int) int {
//	rand.Seed(time.Now().Unix())
//	randNum := rand.Intn(max - min)
//	randNum = randNum + min
//	return randNum
//}

func GenerateReport() {
	//base
	base := make([]string, 0)
	user_list := []string{
		"10000", "10001", "10002",
	}
	app_key_list := []string{
		"5aBPdptyRBfFoHGXHR70lJmRwiaKXmlmPfEoLvscEoM=", "ycCZTDTqtcRAsxCm2WOXRDtQGQPPnJ1tB8tG4L2FD/c=",
		"tYlilPq/Cbcs/faRhUdpqidi2NtKYC+MPAVcyfjssTE=",
	}
	os_list := []string{
		"ios", "android",
	}

	bundle_id_list := []string{
		"com.xqw.ncmj", "com.example.libzadsdk_demo", "com.zonst.libzadsdk_gen", "com.zonst.libzadsdk.appoc", "com.zonst.libzadsdk-demo-oc",
		"com.zonst.libzadsdk-demo-swift",
	}
	slot_id_list := []string{
		"1001", "1002", "1003", "1004", "1005",
	}
	report_date_list := []string{"2018-01-01", "2018-01-02", "2018-01-03", "2018-01-04", "2018-01-05"}
	//columns := []string{"vendor_id", "user_id", "campaign_id",
	//	"ad_id", "report_date", "hour", "win", "imp", "clk", "eimp", "cost"}
	//COPY report_base(vendor_id,user_id,campaign_id,
	//	ad_id,report_date,hour,win,imp,clk,eimp,cost)
	//FROM '/home/tonnn/ad/base.txt' DELIMITER E'\t';
	// 每个用户
	for u := 0; u < len(user_list); u++ {
		user_id := user_list[u]
		vendor_id := 1
		campaign_id := 1
		//每天
		for i := 0; i < len(report_date_list); i++ {
			report_date := report_date_list[i]
			// 每小时
			for k := 0; k < 24; k++ {
				hour := k
				rand.Seed(time.Now().UTC().UnixNano())
				ad_id := rand.Intn(100)
				win := rand.Intn(8000)
				imp := rand.Intn(10000)
				clk := rand.Intn(1000)
				eimp := rand.Intn(9000)
				cost := rand.Intn(1000)
				base = append(base, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v", vendor_id, user_id, campaign_id, ad_id,
					report_date, hour, win, imp, clk, eimp, cost))
			}

		}
	}
	// 写入文件
	if err := myfile.WriteLines(base, "/home/tonnn/ad/base.txt"); err != nil {
		fmt.Println(err)
	}
	//geo
	//columns := []string{"vendor_id", "user_id", "campaign_id",
	//	"ad_id", "country_code", "province_code", "city_code", "report_date", "win", "imp", "clk", "eimp", "cost"}
	//COPY report_geo(vendor_id,user_id,campaign_id,
	//	ad_id,country_code,province_code,city_code,report_date,win,imp,clk,eimp,cost)
	//FROM '/home/tonnn/geo.txt' DELIMITER E'\t';
	geo := make([]string, 0)
	for u := 0; u < len(user_list); u++ {
		user_id := user_list[u]
		vendor_id := 1
		campaign_id := 1
		country := "156000000"
		//每天
		for i := 0; i < len(report_date_list); i++ {
			report_date := report_date_list[i]

			//每个省
			for _, province := range options.ProvinceList {

				for _, city := range options.ProvinceCityMap[province] {
					rand.Seed(time.Now().UTC().UnixNano())
					ad_id := rand.Intn(100)
					win := rand.Intn(8000)
					imp := rand.Intn(10000)
					clk := rand.Intn(1000)
					eimp := rand.Intn(9000)
					cost := rand.Intn(1000)
					geo = append(geo, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v", vendor_id, user_id, campaign_id, ad_id,
						country, province,city, report_date, win, imp, clk, eimp, cost))
				}
			}
			//每个市
		}
	}
	fmt.Println(len(geo))
	// 写入文件
	if err := myfile.WriteLines(geo, "/home/tonnn/ad/geo.txt"); err != nil {
		fmt.Println(err)
	}
	//app
	//columns := []string{"vendor_id", "user_id", "campaign_id",
	//	"ad_id", "os", "bundle_id", "report_date", "win", "imp", "clk", "eimp", "cost"}
	//COPY report_app(vendor_id,user_id,campaign_id,
	//	ad_id,os,bundle_id,report_date,win,imp,clk,eimp,cost)
	//FROM '/tmp/app.txt' DELIMITER E'\t';
	app := make([]string, 0)
	for u := 0; u < len(user_list); u++ {
		//每天
		for i := 0; i < len(report_date_list); i++ {
			report_date := report_date_list[i]
			user_id := user_list[u]
			vendor_id := 1
			campaign_id := 1
			//每个os
			for _, os := range os_list {
				//每个bundle_id
				for _, bundle_id := range bundle_id_list {
					rand.Seed(time.Now().UTC().UnixNano())
					ad_id := rand.Intn(100)
					win := rand.Intn(8000)
					imp := rand.Intn(10000)
					clk := rand.Intn(1000)
					eimp := rand.Intn(9000)
					cost := rand.Intn(1000)
					app = append(app, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v", vendor_id, user_id, campaign_id, ad_id,
						os, bundle_id, report_date, win, imp, clk, eimp, cost))
				}
			}

		}

	}
	// 写入文件
	if err := myfile.WriteLines(app, "/home/tonnn/ad/app.txt"); err != nil {
		fmt.Println(err)
	}
	//app_slot
	//columns := []string{"os", "bundle_id", "app_key",
	//	"slot_id", "report_date", "imp", "clk"}
	//COPY report_app_slot(os,bundle_id, app_key,slot_id,report_date,imp,clk)
	//FROM '/tmp/app_slot.txt' DELIMITER E'\t';
	app_slot := make([]string, 0)
	for u := 0; u < len(app_key_list); u++ {
		app_key := app_key_list[u]

		//每天
		for i := 0; i < len(report_date_list); i++ {
			report_date := report_date_list[i]
			//每个os
			for _, os := range os_list {
				//每个bundle_id
				for _, bundle_id := range bundle_id_list {
					rand.Seed(time.Now().UTC().UnixNano())
					//每个slot_id
					imp := rand.Intn(10000)
					clk := rand.Intn(1000)
					for _, slot_id := range slot_id_list {
						app_slot = append(app_slot, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v", os, bundle_id, app_key,
							slot_id, report_date, imp, clk))
					}
				}

			}
		}
	}
	// 写入文件
	if err := myfile.WriteLines(app_slot, "/home/tonnn/ad/app_slot.txt"); err != nil {
		fmt.Println(err)
	}
	////app_reward
	//columns := []string{"os", "bundle_id", "app_key",
	//	"report_date", "imp", "amount", "reward", "uv"}
	//COPY report_app_reward(os,bundle_id,app_key,report_date,imp,amount,reward,uv)
	//FROM '/tmp/app_reward.txt' DELIMITER E'\t';
	app_reward := make([]string, 0)
	for u := 0; u < len(app_key_list); u++ {
		app_key := app_key_list[u]

		//每天
		for i := 0; i < len(report_date_list); i++ {
			report_date := report_date_list[i]
			//每个os
			for _, os := range os_list {
				//每个bundle_id
				for _, bundle_id := range bundle_id_list {
					rand.Seed(time.Now().UTC().UnixNano())
					imp := rand.Intn(10000)
					amount := rand.Intn(10000)
					reward :=  rand.Intn(10000)
					uv := rand.Intn(1000)
					app_reward = append(app_reward, fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v", os, bundle_id, app_key,
						report_date, imp, amount, reward,uv))
				}

			}
		}
	}
	// 写入文件
	if err := myfile.WriteLines(app_reward, "/home/tonnn/ad/app_reward.txt"); err != nil {
		fmt.Println(err)
	}

}
