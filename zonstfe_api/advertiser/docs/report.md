## 基本报表
### 请求URL:
- GET host/api/v1/report/base
### 返回示例:

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位  cpc clk/cost*1000 cpm imp/cost*1000

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) 需要传递 hour |
| edate | 否 | string | 结束时间(默认当天）需要传递 hour |
| campaign_id | 否 | int | 计划 ID  select 列表 通过 GET host/api/v1/option/campaign |
| ad_id | 否 | int | ad ID 通过 campaign_id select 级联 |
| vendor_id | 否 | int | 流量平台 ID 通过 options.vendor |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "campaign_id": 1,
            "ad_id": 75,
            "report_date": "2018-01-05",
            "campaign_name": "喜马拉雅FM「听书社区」电台有声小说相声评书",
            "ad_name": "asd",
            "eimp": 3980,
            "imp": 8271,
            "clk": 78,
            "cost": 290
        }
    ],
    "sum": {
        "imp": 8271,
        "eimp": 3980,
        "clk": 78,
        "cost": 290
    },
    "count": 1,
    "total": 1
}

```


## 基本小时报表
### 请求URL:
- GET host/api/v1/report/base/hour

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位  cpc clk/cost*1000 cpm imp/cost*1000

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) 需要传递 hour |
| edate | 否 | string | 结束时间(默认当天）需要传递 hour |
| campaign_id | 否 | int | 计划 ID  select 列表 通过 GET host/api/v1/option/campaign |
| ad_id | 否 | int | ad ID 通过 campaign_id select 级联 |
| vendor_id | 否 | int | 流量平台 ID 通过 options.vendor |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "report_date": "2018-01-05",
            "hour": 0,
            "hour_date": "2018-01-05 00:00",
            "eimp": 5840,
            "imp": 4371,
            "clk": 351,
            "cost": 460
        }
    ],
    "sum": {
        "imp": 4371,
        "eimp": 5840,
        "clk": 351,
        "cost": 460
    }
}
```

## 地域国家报表
### 请求URL:
- GET host/api/v1/report/geo

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位  cpc clk/cost*1000 cpm imp/cost*1000
- 需要映射 地域code 中文 根据options.geo_name

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) |
| edate | 否 | string | 结束时间(默认当天）|
| campaign_id | 否 | int | 计划 ID  select 列表 通过 GET host/api/v1/option/campaign |
| ad_id | 否 | int | ad ID 通过 campaign_id select 级联 |
| province_code | 否 | string |  province_code city_code 级联 options.province_city_code |
| city_code | 否 | string |  通过 province_code 级联 |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        
        {
            "province_code": "156140000",
            "city_code": "156140600",
            "campaign_id": 1,
            "ad_id": 31,
            "report_date": "2018-01-05",
            "imp": 7815,
            "eimp": 3217,
            "clk": 820,
            "cost": 751,
            "campaign_name": "喜马拉雅FM「听书社区」电台有声小说相声评书",
            "ad_name": "asd"
        }
    ],
    "sum": {
        "imp": 7815,
        "eimp": 3217,
        "clk": 820,
        "cost": 751
    },
    "count": 1,
    "total": 1
}
```

## 地域国家报表
### 请求URL:
- GET host/api/v1/report/geo/country

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位  cpc clk/cost*1000 cpm imp/cost*1000
- 需要映射 地域code 中文 根据options.geo_name

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) |
| edate | 否 | string | 结束时间(默认当天）|
| campaign_id | 否 | int | 计划 ID  select 列表 通过 GET host/api/v1/option/campaign |
| vendor_id | 否 | int | 流量平台 ID 通过 options.vendor |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "country_code": "156000000",
            "report_date": "2018-01-05",
            "imp": 1710978,
            "eimp": 1512083,
            "clk": 176357,
            "cost": 169442
        }
    ],
    "sum": {
        "imp": 1710978,
        "eimp": 1512083,
        "clk": 176357,
        "cost": 169442
    }
}
```

## 地域省份报表
### 请求URL:
- GET host/api/v1/report/geo/province

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位  cpc clk/cost*1000 cpm imp/cost*1000
- 需要映射 地域code 中文 根据options.geo_name

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) |
| edate | 否 | string | 结束时间(默认当天）|
| campaign_id | 否 | int | 计划 ID  select 列表 通过 GET host/api/v1/option/campaign |
| vendor_id | 否 | int | 流量平台 ID 通过 options.vendor |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "province_code": "156110000",
            "report_date": "2018-01-05",
            "imp": 334,
            "eimp": 5006,
            "clk": 713,
            "cost": 944
        },
        {
            "province_code": "156120000",
            "report_date": "2018-01-05",
            "imp": 5856,
            "eimp": 5540,
            "clk": 860,
            "cost": 62
        },
        {
            "province_code": "156130000",
            "report_date": "2018-01-05",
            "imp": 71674,
            "eimp": 61375,
            "clk": 7558,
            "cost": 4857
        }
    ],
    "sum": {
        "imp": 1710978,
        "eimp": 1512083,
        "clk": 176357,
        "cost": 169442
    }
}
```

## 投放媒体报表
### 请求URL:
- GET host/api/v1/report/app

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位  cpc clk/cost*1000 cpm imp/cost*1000
- 通过 os+"_"+bundle_id 去 options.app_name 映射  app_name 

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) |
| edate | 否 | string | 结束时间(默认当天）|
| os | 否 | string | 操作系统 通过 options.os |
| campaign_id | 否 | int | 计划 ID  select 列表 通过 GET host/api/v1/option/campaign |
| vendor_id | 否 | int | 流量平台 ID 通过 options.vendor |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "os": "android",
            "campaign_id": 1,
            "ad_id": 56,
            "bundle_id": "com.example.libzadsdk_demo",
            "report_date": "2018-01-05",
            "campaign_name": "喜马拉雅FM「听书社区」电台有声小说相声评书",
            "ad_name": "ads",
            "eimp": 7366,
            "imp": 3902,
            "clk": 92,
            "cost": 446
        }
    ],
    "sum": {
        "imp": 3902,
        "eimp": 7366,
        "clk": 92,
        "cost": 446
    },
    "count": 1,
    "total": 1
}
```







