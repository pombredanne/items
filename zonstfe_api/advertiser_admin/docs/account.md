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
            "campaign_count":1,
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

## 查询单个账户
### 请求URL:
- GET host/api/v1/account/{user_id}

### 返回示例
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "user_id": 16,
        "user_type": "comapny",
        "email": "1020300659@qq.com",
        "company_name": "中至",
        "real_name": "张三",
        "qq": "1020300659",
        "phone": "1020300659",
        "create_date": "2017-12-11",
        "status": 0
    }
}
```

## 查询账户税信息
### 请求URL:
- GET host/api/v1/account/tax/{user_id}

### 返回示例
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "tax_id": 1,
        "tax_no": "10101101011010110101",
        "company_name": "中至",
        "address": "新建区望城镇",
        "telephone": "18679958179",
        "bank_name": "招商银行",
        "bank_no": "6214857908143814"
    }
}  
```

## 修改账户税信息
### 请求URL:
- PUT host/api/v1/account/tax/{user_id}

### 请求示例
```json
{
    "tax_no": "10101101011010110101",
    "company_name": "中至",
    "address": "新建区望城镇",
    "telephone": "18679958179",
    "bank_name": "招商银行",
    "bank_no": "6214857908143814"
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



## 查询账户联系人信息
### 请求URL:
- GET host/api/v1/account/deliver/{user_id}

### 返回示例
```json
{
    "status": 0,
    "msg": "success",
    "data": {
        "deliver_id": 1,
        "email": "1020300659@qq.com",
        "address": "新建区望城镇",
        "telephone": "18679958179",
        "receiver": "蔡家昌"
    }
}  
```

## 修改账户税信息
### 请求URL:
- PUT host/api/v1/account/deliver/{deliver_id}

### 请求示例
```json
{
    "email": "1020300659@qq.com",
    "address": "新建区望城镇",
    "telephone": "18679958179",
    "receiver": "蔡家昌"
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

## 账户新增
### 简要描述:
- user_type 为 company时 account.company_name 字段 tax 字段  不用填写和传输 反之需要

### 请求URL:
- POST host/api/v1/account/email/{email} # 验证是否存在
- POST host/api/v1/account

### 相关正则:
- email:```^([0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*@([0-9a-zA-Z][-\w]*[0-9a-zA-Z]\.)+[a-zA-Z]{2,9})$```
- qq:```[1-9]\d{4,12}$```
- phone:```^1[34578]\d{9}$```
- company_name:```^[a-zA-Z\(\（\)\）\u4e00-\u9fa5]{1,100}$```
- real_name:```^[a-zA-Z\u4e00-\u9fa5\-\.\·]{1,20}$```
- telephone:```(1[3|4|5|8][0-9]\d{4,8})|0\d{2,3}-\d{7,8}$```

### account参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| email | 是 | string | 账户邮箱 |
| password | 是 | string | 密码 len>=6 len<=20 |
| dpassword | 是 | string | 确认密码 len>=6 len<=20 |
| user_type | 是 | string | 账户类型 根据options.user_type |
| real_name | 是 | string | 真实姓名 |
| company_name | 否 | string | 当user_type 为 company 时需要填写 |
| qq | 是 | string | QQ |
| phone | 是 | string | phone |

### tax参数(当user_type 为 company 时需要填写)
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| tax_no | 是 | string | 纳税人识别号 len>=15 len<=20 |
| company_name | 是 | string | 公司名称 |
| address | 是 | string | 公司地址 |
| telephone | 是 | string | 电话 |
| bank_name | 是 | string | 电话 |
| bank_no | 是 | string | 电话 |






### 请求示例:
```json
{
    "account": {
        "email": "1020300659@qq.com",
        "password": "123456",
        "dpassword": "123456",
        "user_type": "company",
        "real_name": "蔡家昌",
        "company_name": "中至",
        "qq": "1020300659",
        "phone": "18679958179"
    },
    "tax": {
        "tax_no": "10101101011010110101",
        "company_name": "中至",
        "address": "新建区望城镇",
        "telephone": "18679958179",
        "bank_name": "招商银行",
        "bank_no": "6214857908143814"
    },
    "deliver": {
        "email": "1020300659@qq.com",
        "address": "新建区望城镇",
        "telephone": "18679958179",
        "receiver": "蔡家昌"
    }
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

## 账户修改
### 请求URL:
- PUT host/api/v1/account/{user_id}

### 请求示例:
```json
{
    "user_type": "company_name",
    "real_name": "蔡家昌",
    "company_name": "中至",
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

## 查询充值记录 默认查询 待处理
### 请求URL:
- GET /api/v1/account/recharge/list

### 查询参数
| 字段 | 约束 | 类型 | 说明 |
| --- | --- | --- | --- |
| user_id | 可选 | int | 用户ID |
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
            "recharge_id": 1,
            "user_id": 1,
            "zonst_user_id": 2,
            "order_no": "123123",
            "order_money": 123,
            "order_type": 1,
            "order_date": "2017-08-02 12:24:20",
            "account_no": "13123213",
            "account_name": "张三",
            "description": "充值",
            "update_date": "2017-08-02 12:24:20",
            "create_date": "2017-08-02 12:24:20",
            "status": 0
        }
    ],
    "count": 1,
    "total": 1
}
```

## 查询单条充值记录
### 请求URL:
- GET host/api/v1/account/recharge/{recharge_id}

### 返回示例:
```json
{
    "status": 0,
    "msg": "success",
    "data": [
        {
            "user_id": 1,
            "zonst_user_id": 2,
            "order_no": "123123",
            "order_money": 123,
            "order_type": 1,
            "order_date": "2017-08-02 12:24:20",
            "account_no": "13123213",
            "account_name": "张三",
            "description": "充值",
            "update_date": "2017-08-02 12:24:20",
            "create_date": "2017-08-02 12:24:20",
            "status": 0
        }
    ]
}
```

## 新增充值记录
### 请求URL:
- POST host/api/v1/account/recharge

### 请求示例:
```json
{
    "user_id": 11,
    "order_no": "123123",
    "order_money": 123,
    "order_type": 1,
    "order_date": "2017-08-02 12:24:20",
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


## 充值记录审核
### 请求URL:
- PUT host/api/v1/account/recharge/review/{recharge_id}

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

## 查询操作记录
### 请求URL:
- GET host/api/v1/action/list

### 查询参数
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