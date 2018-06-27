## app 广告位报表
### 请求URL:
- GET host/api/v1/report/app/slot

### 简要描述:
- 前端需要添加额外显示字段 点击率(clk/imp) 保留2位

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) |
| edate | 否 | string | 结束时间(默认当天）|
| os | 否 | string | options.os |
| bundle_id | 否 | string | app_id |
| page | 否 | int | 当前页 |
| page_size | 否 | int | 显示条数 |
| slot_id | 否 | string | 广告位ID  |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "bundle_id": "com.example.libzadsdk_demo",
            "report_date": "2018-01-01",
            "imp": 1867,
            "clk": 27,
            "slot_id": "1001",
            "os": "android"
        }
    ],
    "sum": {
        "imp": 1867,
        "clk": 27
    },
    "count": 1,
    "total": 1
}
```


## app 奖励报表
### 请求URL:
- GET host/api/v1/report/app/reward

### 简要描述:
- 前端需要添加额外显示字段 人均奖励(total/uv) 保留2位 人均次数(number/uv) 保留2位

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sdate | 否 | string | 开始时间(默认当天) |
| edate | 否 | string | 结束时间(默认当天）|
| os | 否 | string | options.os |
| bundle_id | 否 | string | app_id |
| slot_id | 否 | string | 广告位ID |
| page | 否 | int | 当前页 |
| page_size | 否 | int | 显示条数 |

```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "bundle_id": "com.example.libzadsdk_demo",
            "report_date": "2018-01-01",
            "os": "android",
            "amount": 6373,
            "reward": 7287,
            "uv": 42,
            "imp": 9005
        }
    ],
    "sum": {
        "imp": 9005,
        "amount": 6373,
        "reward": 7287,
        "uv": 42
    },
    "count": 1,
    "total": 1
}
```

### 返回参数说明:
| 参数名 | 说明|
| --- | --- |
| imp | 展示 |
| os | 系统 |
| amount | 返奖总数(次数*每次返奖个数) |
| reward | 返奖次数 |
| uv | 返奖人数 |











