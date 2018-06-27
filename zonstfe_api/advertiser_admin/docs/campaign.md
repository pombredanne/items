## Ad列表
### 请求URL:
- GET host/api/v1/campaign/ad/list

### 查询参数
| 字段 | 约束 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 可选 | string | 开始时间 |
| edate | 可选 | string | 结束时间 |
| campaign_name | 可选 | string | 活动名称 |
| ad_type | 可选 | string | 广告类型 |
| ad_size | 可选 | string | 广告尺寸 |
| status | 可选 | int | -1 审核失败 0 待审核,1 已审核 2 素材缺少) |
| page | 可选 | string | 页码 |
| page_size | 可选 | string | 页数 |

## response
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "creative_set_id": 2,
            "campaign_id": 1,
            "campaign_name": "ss",
            "ad_id": 1,
            "name": "dasd_横幅",
            "ad_type": "graphic",
            "ad_size": "banner",
            "creative": {
                "image": "http://www.baidu.com"
            },
            "status": 0
        }
    ],
    "count": 1,
    "total": 1
}
```

## 广告创意集合Ad列表
### 请求URL:
- GET host/api/v1/campaign/ad/list/{set_id}

### 查询参数
| 字段 | 约束 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 可选 | string | 开始时间 |
| edate | 可选 | string | 结束时间 |
| campaign_name | 可选 | string | 活动名称 |
| ad_type | 可选 | string | 广告类型 |
| ad_size | 可选 | string | 广告尺寸 |
| status | 可选 | int | -1 审核失败 0 待审核,1 已审核 2 素材缺少) |
| page | 可选 | string | 页码 |
| page_size | 可选 | string | 页数 |

## response
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "creative_set_id": 2,
            "campaign_id": 1,
            "campaign_name": "ss",
            "ad_id": 1,
            "name": "dasd_横幅",
            "ad_type": "graphic",
            "ad_size": "banner",
            "creative": {
                "image": "http://www.baidu.com"
            },
            "status": 0
        }
    ],
    "count": 1,
    "total": 1
}
```

## 单广告创意审核
### 请求URL:
- PUT host/api/v1/campaign/ad/review/{ad_id} 

### 审核成功请求示例:
```json
{
  "review_type":1
}
```

### 审核失败请求示例
```json
{
  "review_type":-1
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

## 批量广告创意审核
### 请求URL:
- PUT host/api/v1/campaign/ad/review/batch

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| review_type | 是 | int | 审核状态 -1 审核失败  1 审核成功 |
| list | 是 | []string | ad_id 列表 |


### 审核成功请求示例:
```json
{
  "review_type":1,
  "list":["1","2","3","4"]
}
```

### 审核失败请求示例
```json
{
  "review_type":-1,
  "list":["1","2","3","4"]
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


