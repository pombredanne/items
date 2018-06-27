## APP列表
### 请求URL:
- GET host/api/v1/app/list

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| name | 否 | string | app 名称 |
| bundle_id | 否 | string | app ID  |
| os | 否 | string | app 系统 |
| page | 否 | int | 当前页 |
| page_size | 否 | int | 显示条数 |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "app_id": 2,
            "name": "南昌麻将",
            "bundle_id": "com.xqw.ncmj",
            "os": "ios",
            "category": "1023",
            "sub_category":"102302",
            "create_date": "2017-11-20 16:46:55",
            "update_date": "2017-11-20 16:46:55",
            "status": 1
        },
        {
            "app_id": 1,
            "name": "南昌麻将",
            "bundle_id": "com.xqw.ncmj",
            "os": "android",
            "category": "102302",
            "create_date": "2017-11-20 16:46:27",
            "update_date": "2017-11-20 16:46:55",
            "status": 1
        }
    ],
    "count": 2,
    "total": 2
}
```
### 返回参数说明:
| 参数名 | 说明|
| --- | --- |
| category | 需要根据options隐射中文 |
| status | -1 审核失败 0 待审核 1审核通过 |

## APP创建
### 简要描述:
- 需要账户审核后才能创建

### 请求URL:
- POST host/api/v1/app

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| name | 是 | string | 媒体名称 len<=40 |
| os | 是 | string | 系统平台 |
| bundle_id | 是 | string | 程序主包名 |
| category | 是 | string | 媒体分类 |
| sub_category | 是 | string | 媒体子分类 |
| store_name | 是 | string | 应用商店  根据OS选择 加载不同的 store_list |
| store_url | 是 | string | 详情页地址  需要正则验证 |
| keywords | 是  | string | 关键词 len<=60 |
| describtion | 是  | string | 媒体简介 len>=40 |
| slot_map | 是 |  map[string]float | app 计费广告位 |
| category_limit | 是 | json obj | 广告分类屏蔽 |
| reward | 是 | json obj | app 激励配置 |


### 相关正则:
- store_url:```(http|https)://\w+(-\w+)*(\.\w+(-\w+)*)*```

### 请求示例:
```json
{
    "name": "中至南昌麻将",
    "os": "ios",
    "bundle_id": "com.xqw.ncmj",
    "category": "1023",
    "sub_category":"102302",
    "store_name": "App Store",
    "store_url": "https://itunes.apple.com/cn/app/%E4%B8%AD%E8%87%B3%E5%8D%97%E6%98%8C%E9%BA%BB%E5%B0%86/id1134908723?mt=8",
    "keywords": "哈哈",
    "describtion": "一款熟人约局的南昌麻将社交游戏，可创建房间与现实中好友对战畅玩，分享到朋友圈，战绩炫耀等功能。",
    "category_limit":{
        "enable":0,
        "list":[]
    },
    "reward": {
        "amount": 10,
        "enable": 1,
        "callback": 1,
        "currency": "$",
        "callback_url": "https://www.baidu.com?imei={IMEI}&idfa={IDFA}&amount={AMOUNT}&user_id={USER_ID}&currency={CURRENCY}&event_id={EVENT_ID}"
    },
    "slot_map": {"1000":1,"10001":2.5}
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

## APP修改
### 请求URL:
- PUT host/api/v1/app/{app_id}

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| name | 是 | string | 媒体名称 len<=40 |
| category | 是 | string | 媒体分类 |
| keywords | 是 | string | 关键词 ","号分割 len<=60 |
| describtion | 是 | string | 媒体简介 len>=40|
| slot_map | 是 |  map[string]float | app 计费广告位 |
| reward | 是 | json obj | app 激励配置 |
| category_limit | 是 | json obj | 广告分类屏蔽 |


### reward 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| enable | 是 | int | 是否开启 0 关闭 1 开启 |
| freq | 是 | int | 速率 {1:"均速",2:"随机"} |
| currency | 是 | string | 货币名称 |
| amount | 是 | int | 奖励个数 |
| callback | 是 | int | 是否服务回调 |
| callback_url | 是 | string | 服务回调地址 可选宏 {IDFA} {IMEI} {AMOUNT} {USER_ID} {CURRENCY} {EVENT_ID} {IP}|

### 相关正则:
- callback_url:```(http|https)://\w+(-\w+)*(\.\w+(-\w+)*)*```

### 请求示例:
```json
{
    "name": "中至南昌麻将-1",
    "category": "1023",
    "sub_category":"102302",
    "keywords": "哈哈",
    "describtion": "一款熟人约局的南昌麻将社交游戏，可创建房间与现实中好友对战畅玩，分享到朋友圈，战绩炫耀等功能。",
    "slot_map": {"1000":1,"10001":2.5},
    "reward": {
        "amount": 10,
        "enable": 1,
        "callback": 1,
        "currency": "$",
        "callback_url": "https://www.baidu.com?imei={IMEI}&idfa={IDFA}&amount={AMOUNT}&user_id={USER_ID}&currency={CURRENCY}&event_id={EVENT_ID}"
    },
    "category_limit":{
            "enable":0,
            "list":[]
    }
}
```

## APP单个:
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
        "category": "1023",
        "sub_category":"102302",
        "keywords": "哈哈",
        "describtion": "一款熟人约局的南昌麻将社交游戏，可创建房间与现实中好友对战畅玩，分享到朋友圈，战绩炫耀等功能。",
        "category_limit":{
                    "enable":0,
                    "list":[]
         },
        "reward": {
            "amount": 10,
            "enable": 1,
            "callback": 1,
            "currency": "$",
            "callback_url": "https://www.baidu.com?imei={IMEI}&idfa={IDFA}&amount={AMOUNT}&user_id={USER_ID }&currency={CURRENCY}&event_id={EVENT_ID}"
        },
        "slots": {"1000":1,"10001":2.5},
        "status": 1
    }
}
```

