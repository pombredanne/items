## 账户信息
### 请求URL:
- GET host/api/v1/account

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "user_type": "person",
        "email": "1020300659@qq.com",
        "company_name": "中至",
        "real_name": "蔡家昌",
        "app_key": "zonst",
        "qq": "1020300659",
        "phone": "18679958179",
        "status": 0
    }
}
```

### 返回参数说明:
| 参数名 | 说明|
| --- | --- |
| status | -1 审核失败 0 待审核 1已审核 |

## 账户信息修改
### 简要描述:
- 根据 user_type 判断是否要填写 company_name 字段

### 请求URL:
- PUT host/api/v1/account

### 请求示例:
```json
{
  "company_name":"中至1",
  "real_name": "蔡家昌",
  "qq": "1020300659",
  "phone": "18679958179"
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

## 账号密码修改
### 请求URL:
- PUT host/api/v1/account/password

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| old_password | 是 | string | 原始密码 |
| password | 是 | string | 密码 |
| dpassword | 是 | string | 确认密码 |

### 请求示例:
```json
{
  "old_password":"123456",
  "password": "1234567",
  "dpassword": "1234567"
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

## 查询充值记录
### 请求URL:
- GET host/api/v1/account/recharge/list

### 查询参数
| 字段 | 约束 | 类型 | 说明 |
| --- | --- | --- | --- |
| status | 可选 | int | 处理状态 |
| page | 可选 | string | 页码 |
| page_size | 可选 | string | 页数 |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "order_money": 123,
            "order_type": 1,
            "order_date": "2017-08-02 12:24:20",
            "account_no": "13123213",
            "account_name": "张三",
            "description": "充值",
            "create_date": "2017-08-02 12:24:20",
            "status": 0
        }
    ],
    "count": 1,
    "total": 1
}
```

## 新增充值记录
### 请求URL:
- POST host/api/v1/account/recharge

### 请求示例:
```json
{
    "order_money": 123,
    "order_type": 1,
    "account_no": "123213123",
    "account_name": "张三",
    "description": "充值"
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

## 查询操作记录
### 请求URL:
- GET host/api/v1/action/list

## 查询参数
| 字段 | 约束 | 类型 | 说明 |
| --- | --- | --- | --- |
| action_module | 可选 | string | 操作模块 |
| sdate | 可选 | string | 开始时间 |
| edate | 可选 | string | 结束时间 |
| page | 可选 | string | 页码 |
| page_size | 可选 | string | 页数 |

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "action_module": "account",
            "action_id": 2,
            "action_type": "新增",
            "create_date": "2017-08-02 12:24:20",
            "ip_address": "127.0.0.1"
        }
    ],
    "count": 1,
    "total": 1
}
```

## 账户余额
### 请求URL:
- GET host/api/v1/account/balance

### 返回示例:
```json
{
    "status":0,
    "msg":"success",
    "data":{"balance":112}
}
```









