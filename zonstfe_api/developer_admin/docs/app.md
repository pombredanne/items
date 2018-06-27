## app列表
### 请求URL:
- GET host/api/v1/app/list

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| user_id | 否 | int | 用户ID  根据加载 用户邮箱列表模糊查询 |
| status | 否 | int | 状态 加载option -1 审核失败 0 待审核 1已审核 -1 审核失败  |
| page | 否 | int | 当前页 不传默认 1 |
| page_size | 否 | int | 每页个数 不传默认20 |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "app_id": 1,
            "name": "南昌麻将",
            "bundle_id": "com.xqw.ncmj",
            "os": "android",
            "category": "1017",
            "sub_category": "101704",
            "user_id": 16,
            "user_email": "1020300659@qq.com",
            "create_date": "2017-11-20 16:46:27",
            "status": 1
        }
    ],
    "count": 1,
    "total": 1
}
```

## app 单个
### 请求URL:
- GET host/api/v1/app/{app_id}

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "app_id": 1,
        "name": "南昌麻将",
        "bundle_id": "com.xqw.ncmj",
        "os": "android",
        "category": "1017",
        "sub_category": "101704",
        "keywords": "",
        "store_name": "",
        "store_url": "",
        "describtion": "\"\"",
        "user_id": 16,
        "zonst_user_id": 0,
        "category_limit": {
            "list": [],
            "enable": 0
        },
        "reward": {
            "freq": 1,
            "amount": 10,
            "enable": 1,
            "callback": 1,
            "currency": "$",
            "callback_url": "https://www.baidu.com?imei={IMEI}&idfa={IDFA}&amount={AMOUNT}&user_id={USER_ID }&currency={CURRENCY}&event_id={EVENT_ID}&ip={IP}"
        },
        "slots": {
            "1000": 1,
            "1001": 1,
            "1002": 2,
            "1003": 1.5,
            "1004": 3,
            "1005": 3.5,
            "10001": 2.5
        },
        "create_date": "2017-11-20 16:46:27",
        "status": 1
    }
}
```

## App 审核
### 简要描述:
- 先查询单个APP信息 通过点击 审核通过和审核失败 失败需要填写失败原因用户通知用户

### 请求URL:
- PUT host/api/v1/app/review/{app_id}

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| review_type | 是 | int | 审核状态 -1 审核失败  1 审核成功 |
| group_name | 否 | int | 审核消息 系统消息 账户消息 财务消息 |

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






