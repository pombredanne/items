## 账户列表
### 请求URL:
- GET host/api/v1/account/list

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
            "user_type": "person",
            "email": "1020300659@qq.com",
            "company_name": "中至",
            "real_name": "蔡家昌",
            "app_count":1,
            "deal_type": "share",
            "deal_scale": 0,
            "qq": "1020300659",
            "phone": "18679958179",
            "create_date": "2017-11-22 14:59:04",
            "status": 0
        }
    ],
    "count": 1,
    "total": 1
}
```

### 返回参数说明:
| 参数名 | 说明|
| --- | --- |
| status | -1 审核失败 0 待审核 1已审核 -1 审核失败 |
| user_type | person 个人 company 公司 |
| deal_type | share 分成 bidding 竞价 |
| deal_scale | 分成比例 |


## 查询单个账户
### 请求URL:
- GET host/api/v1/account/{user_id}

### 返回示例
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "account": {
            "user_id": 16,
            "user_type": "comapny",
            "email": "1020300659@qq.com",
            "company_name": "中至",
            "real_name": "张三",
            "deal_type": "bidding",
            "deal_scale": 0,
            "qq": "1020300659",
            "phone": "1020300659",
            "create_date": "2017-12-11",
            "status": 0
        },
        "finance": {
            "finance_id": 2,
            "account_name": "",
            "bank_name": "",
            "bank_no": "",
            "bank_province": "",
            "bank_city":"",
            "bank_sub_address": ""
        }
    }
}
```

## 账户审核
### 简要描述:
- 先查询单个账户信息 通过点击 审核通过和审核失败 失败需要填写失败原因用户通知用户

### 请求URL:
- PUT host/api/v1/account/review/{user_id}

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| review_type | 是 | int | 审核状态 -1 审核失败  1 审核成功 |


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

## 账户密码修改
### 请求URL:
- PUT host/api/v1/account/password/update

### 请求示例:
```json
{
    "email":"1020300659@qq.com",
    "password":"123456",
    "dpassword":"123456"
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

## 账户开户
### 请求URL:
- POST host/api/v1/account/email/{email} # 验证是否存在
- POST host/api/v1/account

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| sub_branch | 是 | string | 支行 len<=300 |
| bank_no | 是 | string | 银行账号 len>=16 len<=19 |
| company_name | 否 | string | user_type 为 person 时 去除该字段 |



### 相关正则:
- email:```^([0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*@([0-9a-zA-Z][-\w]*[0-9a-zA-Z]\.)+[a-zA-Z]{2,9})$```
- qq:```[1-9]\d{4,12}$```
- phone:```^1[34578]\d{9}$```
- company_name:```^[a-zA-Z\(\（\)\）\u4e00-\u9fa5]{1,100}$```
- real_name:```^[a-zA-Z\u4e00-\u9fa5\-\.\·]{1,20}$```

### 请求示例:
```json
{
    "account": {
        "email": "1020300659@qq.com",
        "password": "123456",
        "dpassword": "123456",
        "real_name": "蔡家昌",
        "company_name": "",
        "user_type": "person",
        "deal_type": "share",
        "deal_scale": 0.5,
        "qq": "1020300659",
        "phone": "18679958179"
    },
    "finance": {
        "account_name": "蔡家昌",
        "bank_name": "招商银行",
        "bank_no": "6214857908143814",
        "bank_province": "江西",
        "bank_city": "南昌",
        "sub_branch": "红谷支行"
    }
}
```
## 提现列表
### 请求URL:
- GET host/api/v1/account/payment/list

### 查询参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| user_id | 否 | int | 用户ID  根据加载 用户邮箱列表模糊查询 |
| sdate | 否 | string | 时间开始 |
| edate | 否 | string | 时间开始 |
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
            "payment_id": 1,
            "user_id": 16,
            "balance":0,
            "user_email": "1020300659@qq.com",
            "apply_date": "2017-12-19",
            "order_money": 1000,
            "status": 0
        }
    ],
    "count": 1,
    "total": 1
}
```

## 提现单个
- GET host/api/v1/account/payment/{payment_id}

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "payment_id": 1,
        "user_id": 16,
        "balance":0,
        "user_email": "1020300659@qq.com",
        "apply_date": "2017-12-19 14:36:16",
        "order_money": 1000,
        "status": 0
    }
}
```

## 提现审核
### 简要描述:
- 先查询单个提现信息 通过点击 审核通过和审核失败 失败需要填写失败原因用户通知用户

### 请求URL:
- PUT host/api/v1/account/payment/review/{payment_id}

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















