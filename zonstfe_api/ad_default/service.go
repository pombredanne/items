package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"fmt"
)

var allMap = map[string]interface{}{
	"graphic_banner": map[string]interface{}{
		"html": StringFormat(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><div style="position:fixed;z-index:-1; left:0; top:0;width:100%;height:100%;display: flex;align-items: center;justify-content: center;background: #eee"><img src="{{image}}" style="width: 100%;height: auto;" onclick="window.open('{{click_sdk_dp_url}}')" /></div><img src="{{ad_eimp_tracking_url}}" style="position:fixed;top:-100%;width:0px;height:0px;" /></body></html>`, map[string]interface{}{
			"image":                "https://static.qinglong365.com/4e96968e5f86995a6abb310978577af9.jpg",
			"ad_eimp_tracking_url": "https://t.qinglong365.com/v1/g/t_ad_imp?account_id=0&ad_id=0&ad_type=video&ad_size=video&app_id={{app_id}}&request_id={{request_id}}&vendor_id=1",
			"click_sdk_dp_url":     "zonst://81d1370556/clk",
		}),
		"ad_imp_tracking_url": "https://t.qinglong365.com/v1/g/t_ad_imp?account_id=0&ad_id=0&ad_type=video&ad_size=video&app_id={{app_id}}&request_id={{request_id}}&vendor_id=1",
		"ol":                  2,
		"download_url":        "https://www.qinglong365.com",
		"duration":            0,
	},
	"graphic_fullscreen": map[string]interface{}{
		"html": StringFormat(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><div style="position:fixed;z-index:-1; left:0; top:0;width:100%;height:100%;display: flex;align-items: center;justify-content: center;background: #eee"><img src="{{image}}" style="width: 100%;height: auto;" onclick="window.open('{{click_sdk_dp_url}}')" /></div><img src="{{ad_eimp_tracking_url}}" style="position:fixed;top:-100%;width:0px;height:0px;" /></body></html>`, map[string]interface{}{
			"image":                "https://static.qinglong365.com/efb1e1c43dd8264fff702f0624dc9ac8.jpg",
			"ad_eimp_tracking_url": "https://t.qinglong365.com/v1/g/t_ad_imp?account_id=0&ad_id=0&ad_type=video&ad_size=video&app_id={{app_id}}&request_id={{request_id}}&vendor_id=1",
			"click_sdk_dp_url":     "zonst://81d1370556/clk",
		}),
		"ad_imp_tracking_url": "https://t.qinglong365.com/v1/g/t_ad_imp?account_id=0&ad_id=0&ad_type=video&ad_size=video&app_id={{app_id}}&request_id={{request_id}}&vendor_id=1",
		"ol":                  1,
		"download_url":        "https://www.qinglong365.com",
		"duration":            0,
	},
	"video_video": map[string]interface{}{
		"html": StringFormat(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><div style="position:fixed;z-index:-1; left:0; top:0;width:100%;height:100%;display: flex;align-items: center;justify-content: center;background: #eee"><img src="{{image}}" style="width: 100%;height: auto;" onclick="window.open('{{click_sdk_dp_url}}')" /></div><img src="{{ad_eimp_tracking_url}}" style="position:fixed;top:-100%;width:0px;height:0px;" /></body></html>`, map[string]interface{}{
			"image":                "https://static.qinglong365.com/d73c184818c636bae7be1be68a2518c5.jpg",
			"ad_eimp_tracking_url": "https://t.qinglong365.com/v1/g/t_ad_imp?account_id=0&ad_id=0&ad_type=video&ad_size=video&app_id={{app_id}}&request_id={{request_id}}&vendor_id=1",
			"click_sdk_dp_url":     "zonst://81d1370556/clk",
		}),
		"ad_imp_tracking_url": "https://t.qinglong365.com/v1/g/t_ad_imp?account_id=0&ad_id=0&ad_type=video&ad_size=video&app_id={{app_id}}&request_id={{request_id}}&vendor_id=1",
		"ol":                  2,
		"download_url":        "https://www.qinglong365.com",
		"duration":            15,
		"video":               "https://static.qinglong365.com/d336858809eb501386b26164ddd4792b.mp4",
	},
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/default", func(c *gin.Context) {
		ad_size := c.DefaultQuery("ad_size", "")
		ad_type := c.DefaultQuery("ad_type", "")
		value, ok := allMap[ad_type+"_"+ad_size]
		if ok {
			c.JSON(200, value)
		} else {
			c.JSON(404, gin.H{})
		}
	})
	r.Run(":20002")
}

func StringFormat(format string, p map[string]interface{}) string {
	args, i := make([]string, len(p)*2), 0
	for k, v := range p {
		args[i] = "{{" + k + "}}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}
