## 广告创意列表
### 请求URL:
- GET host/api/v1/campaign/creative/list

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "creative_id": 1,
            "name": "视频广告",
            "ad_type": "video",
            "ad_size": "video",
            "material": "570x320视频",
            "description": "横向视频广告",
            "ol": 2
        },
        {
            "creative_id": 2,
            "name": "视频广告",
            "ad_type": "video",
            "ad_size": "video",
            "material": "320x570视频",
            "description": "纵向视频广告",
            "ol": 1
        },
        {
            "creative_id": 3,
            "name": "开屏广告",
            "ad_type": "graphic",
            "ad_size": "interstitial",
            "material": "768x1024单图",
            "description": "开屏广告",
            "ol": 1
        },
        {
            "creative_id": 4,
            "name": "开屏广告",
            "ad_type": "graphic",
            "ad_size": "interstitial",
            "material": "320x512单图",
            "description": "开屏广告",
            "ol": 1
        },
        {
            "creative_id": 5,
            "name": "开屏广告",
            "ad_type": "graphic",
            "ad_size": "interstitial",
            "material": "320x480单图",
            "description": "开屏广告",
            "ol": 1
        },
        {
            "creative_id": 6,
            "name": "Banner广告",
            "ad_type": "graphic",
            "ad_size": "banner",
            "material": "320x50单图",
            "description": "Banner广告",
            "ol": 2
        }
    ]
}
```

### 返回参数
| 字段 | 类型 | 说明 |
| --- |--- | --- |
| name  | string | 创意名称 |
| material  | string | 所需素材 |
| creative  | string | 素材上传 所需对象 |


## 活动列表
### 请求URL:
- GET host/api/v1/campaign/list

### 查询参数
| 字段 | 约束 | 类型 | 说明 |
| --- | --- | --- | --- |
| name | 可选 | string | 活动名称 |
| sdate | 可选 | string | 创建时间开始 |
| edate | 可选 | string | 创建时间结束 |
| bundle_id | 可选 | string | 唯一标识 |
| app_platform | 可选 | string | 平台('android','ios') |
| status | 可选 | int | 状态((0,暂停),(1,活跃)) |
| page | 可选 | string | 页码 |
| page_size | 可选 | string | 页数 |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "campaign_id": 1,
            "name": "TEST",
            "bundle_id": "2312",
            "budget_day": 100,
            "bidding_type": "CPM",
            "create_date": "2017-08-29",
            "status": 1
        }
    ],
    "count": 1,
    "total": 1
}
```

## 广告活动状态修改
### 简要描述:
- 当前余额不足100时 开启会失败

### 请求URL:
- PUT host/api/v1/campaign/switch/{campaign_id}

### 返回示例:
```json
{
    "status":0,
    "msg":"success",
    "data":[]
}
```

## 单活动查询
### 请求URL:
- GET host/v1/campaign/{campaign_id}

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "campaign_id": 3,
            "name": "TEST",
            "bundle_id": "com.xq5.ncmj",
            "category": "1018",
            "sub_category": "101801",
            "app_platform": "android",
            "budget_day": 100,
            "bidding_max": 1000,
            "bidding_min": 100,
            "bidding_type": "CPM",
            "freq": {
                "open": true,
                "type": "day",
                "num": 100
            },
            "speed": 1,
            "targeting": {
                "vendor": {
                    "list": [
                        "zonst"
                    ],
                    "open": true
                },
                "carrier": {
                    "list": [],
                    "open": false
                },
                "network": {
                    "list": [],
                    "open": false
                },
                "geo_code": {
                    "list": [
                        "156650100",
                        "156650200"
                    ],
                    "open": true
                },
                "os_version": {
                    "list": [],
                    "open": false
                },
                "day_parting": {
                    "list": [
                        "100001",
                        "100002"
                    ],
                    "open": true
                },
                "device_type": {
                    "list": [],
                    "open": false
                },
                "app_category": {
                    "list": [
                        "100101"
                    ],
                    "open": true
                },
                "segment":{
                "list":[],
                 "open": false
                }
               
            },
            "url": {
                "jump_url": "http://www.baidu.com",
                "deep_link_url": "http://www.baidu.com",
                "tracking_clk_url": "http://www.baidu.com",
                "tracking_imp_url": "http://www.baidu.com"
            }
        }
    ]
}
```

## 活动创建
### 简要描述:
-  tracking_imp_url tracking_clk_url deep_link_url 非必填
     
### 请求URL:
- POST host/api/v1/campaign

### 活动信息
```json
{
    "name": "TEST",
    "bundle_id": "com.xq5.ncmj",
    "app_platform": "ios",
    "category": "1018",
    "sub_category":"101801",
    "budget_day": 100,
    "bidding_max": 1000,
    "bidding_min": 100,
    "bidding_type": "CPM",
    "freq": {
        "open": true,
        "type": "day",
        "num": 100
    },
    "speed": 1
}
```

### 定向信息
```json
{
    "vendor": {
          "open": true,
          "list": [
              "zonst"
          ]
    },
    "geo_code": {
        "open": true,
        "list": [
               "130100000000"
        ]
    },
    "app_category": {
        "open": true,
        "list": [
           "1001"
        ]
    },
    "day_parting": {
        "open": true,
        "list": [
            "10001"
        ]
    },
    "device_type": {
        "open": false,
        "list": []
    },
    "device_brand": {
        "open": false,
        "list": []
    },
    "os_version": {
        "open": false,
        "list": []
    },
    "carrier": {
        "open": false,
        "list": []
    },
    "network": {
        "open": false,
        "list": []
    },"segment":{
     "list":[],
      "open": false
    }
}
```

### URL信息
```json
{
    "jump_url": "",
    "tracking_imp_url": "",
    "tracking_clk_url": "",
    "deep_link_url": ""
}
```

## 最终上报结构 成功后跳转素材上传
```json
{
    "name": "TEST",
    "bundle_id": "com.xq5.ncmj",
    "app_platform": "ios",
    "category": "1018",
    "sub_category":"101801",
    "budget_day": 100,
    "bidding_max": 1000,
    "bidding_min": 100,
    "bidding_type": "CPM",
    "freq": {
        "open": true,
        "type": "day",
        "num": 100
    },
    "speed": 1,
    "targeting": {
        "vendor": {
            "open": true,
            "list": [
                "zonst"
            ]
        },
        "geo_code": {
            "open": true,
            "list": [
                "156650100","156650200"
            ]
        },
        "app_category": {
            "open": true,
            "list": ["1018"]
        },
        "day_parting": {
            "open": true,
            "list": ["0","1"]
        },
        "device_type": {
            "open": false,
            "list": []
        },
        "device_brand": {
            "open": false,
            "list": []
        },
        "os_version": {
            "open": false,
            "list": []
        },
        "carrier": {
            "open": false,
            "list": []
        },
        "network": {
            "open": false,
            "list": []
        },"segment":{
              "list":[],
               "open": false
             }
    },
    "url": {
       "tracking_imp_url": "http://www.baidu.com",
       "tracking_clk_url": "http://www.baidu.com",
       "jump_url": "http://www.baidu.com",
       "deep_link_url": "http://www.baidu.com"
    }
}
```

## 活动修改
### 请求URL:
- PUT host/api/v1/campaign/{campaign_id}

### 请求示例:
```json
{
    "campaign_id":1,
    "name": "TEST",
    "bundle_id": "111",
    "app_platform": "ios",
    "category": "1018",
    "sub_category":"101801",
    "budget_day": 100,
    "bidding_max": 1000,
    "bidding_min": 100,
    "bidding_type": "CPM",
    "freq": {
        "open": true,
        "type": "day",
        "num": 100
    },
    "speed": 1,
    "targeting": {
        "vendor": {
            "open": true,
            "list": [
                "zonst"
            ]
        },
        "geo_code": {
            "open": true,
            "list": ["156650100","156650200"]
        },
        "app_category": {
            "open": true,
            "list": ["1018"]
        },
        "day_parting": {
            "open": true,
           "list": ["0","1"]
        },
        "device_type": {
            "open": false,
            "list": []
        },
        "device_brand": {
            "open": false,
            "list": []
        },
        "os_version": {
            "open": false,
            "list": []
        },
        "carrier": {
            "open": false,
            "list": []
        },
        "network": {
            "open": false,
            "list": []
        },"segment":{
              "list":[],
               "open": false
             }
    },
    "url": {
        "tracking_imp_url": "http://www.baidu.com",
        "tracking_clk_url": "http://www.baidu.com",
        "jump_url": "http://www.baidu.com",
        "deep_link_url": "http://www.baidu.com"
    }
}
```
### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data":{
            "campaign_id": 1
        }
}
```

## 创意上传
### 简要描述
- 需要form 提交 当前素材约束尺寸  width height 
### 请求URL:
- POST host/api/v1/campaign/creative/upload

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "width":320,
            "height":512,
            "path": "https://static.greenclick.cn/p/d41d8cd98f00b204e9800998ecf8427e.mp4",
            "duration":9
        }
    ]
}
```

## 广告新增
### 请求URL:
- POST host/api/v1/campaign/ad/{creative_id}

### 视频广告请求示例
```json
{
    "width": 320,
    "height": 570,
    "video": "http://",
    "image": "http://",
    "name": "ad",
    "duration": 10
}
```

### 图片广告请求示例
```json
{
    "width": 320,
    "height": 570,
    "image": "http://",
    "name": "ad"
}
```


## 广告活动AD列表
### 请求URL:
- GET host/api/v1/campaign/ads/{campaign_id}

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "name": "dasd_横幅",
            "ad_type": "graphic",
            "ad_size": "banner",
            "ol": 2,
            "duration": 0,
            "creative": {
                "image": "http://www.baidu.com"
            },
            "campaign_id": 1,
            "width":320,
            "height":50,
            "creative_id": 6,
            "status": 0
        },
        {
            "name": "dasd_插屏",
            "ad_type": "graphic",
            "ad_size": "interstitial",
            "ol": 1,
            "duration": 0,
            "width":320,
            "height":512,
            "creative": {
                "image": "http://www.baidu.com"
            },
            "campaign_id": 1,
            "creative_id": 4,
            "status": 0
        },
        {
            "name": "dasd_纵向_视频",
            "ad_type": "video",
            "ad_size": "video",
            "ol": 1,
            "duration": 9,
            "width":320,
            "height":570,
            "creative": {
                "video": "http://www.baidu.com",
                "image": "http://www.baidu.com"
            },
            "campaign_id": 1,
            "creative_id": 1,
            "status": 0
        },
        {
            "name": "dasd_横向_视频",
            "ad_type": "video",
            "ad_size": "video",
             "width":570,
            "height":320,
            "ol": 2,
            "duration": 9,
            "creative": {
                "video": "http://www.baidu.com",
                "image": "http://www.baidu.com"
            },
            "campaign_id": 1,
            "creative_id": 2,
            "status": 0
        }
    ],
    "count": 4,
    "total": 4
}
```

## 人群包上传
### 请求URL:
- POST host/api/v1/upload

### 返回示例:
```json
{
    "duration": "9.039000",
    "height": 320,
    "path": "https://static.qinglong365.com/p/d41d8cd98f00b204e9800998ecf8427e.mp4",
    "size": "234120",
    "file_name": "output.mp4",
    "width": 570
}
```




## 人群包列表
### 请求URL:
- GET host/api/v1/campaign/segment/list

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "name": "dds",
            "user_id": 1,
            "user_email": "1020300659@qq.com",
            "uv": 1000,
            "create_date": "2017-09-14 15:05:56",
            "update_date": "2017-09-14 15:05:56"
        }
    ],
    "count": 1,
    "total": 1
}
```

## 人群包创建
### 请求URL:
- POST host/api/v1/campaign/segment
`需要提供一个上传界面 包括名称填写 类型勾选 以及上传包组件 然后提交到此接口`

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| name | 是 | string | 人群包名称 |
| type | 是 | int | 当前操作类型 1 增量 2 全量 |
| pkg_path | 是 | string | 人群包地址 |


### 请求示例:
```json
{
    "name": "dasdas",
    "type": 1,
    "pkg_path": "http://www.baidu.com"
}
```

### 返回示例:
```json
{
    "status":0,
    "msg":"success",
    "data":[]
}
```






























