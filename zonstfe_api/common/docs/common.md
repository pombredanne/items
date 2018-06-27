## 用户注册
### 简要描述:
- 用户注册接口:
    - 需要先验证邮箱是否存在
    - 根据 user_type 判断是否要填写 company_name(默认user_type person) 

### 请求URL:
- POST host/api/signup/{email} # 验证是否存在
- POST host/api/signup # 注册提交

### 参数:
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| user_type | 是 | string | 用户类型 {"company":"企业账户","person":"个人账户"} |
| real_name | 是 | string | 姓名 |
| company_name | 否 | string | 企业名称 当 user_type 为 company 时必填 |
| email | 是 | string | 邮箱 |
| password | 是 | string | 密码 (长度 6-20) |
| dpassword | 是 | string | 确认密码 (长度 6-20) |
| qq | 是 | string | QQ|
| phone | 是 | string | 手机 |

### 相关正则:
- email:```^([0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*@([0-9a-zA-Z][-\w]*[0-9a-zA-Z]\.)+[a-zA-Z]{2,9})$```
- qq:```[1-9]\d{4,12}$```
- phone:```^1[34578]\d{9}$```
- company_name:```^[a-zA-Z\(\（\)\）\u4e00-\u9fa5]{1,100}$```
- real_name:```^[a-zA-Z\u4e00-\u9fa5\-\.\·]{1,20}$```

### 请求示例:
```user_type 为person 时 company_name 字段不传```

```json
{
	"user_type":"comapny",
	"real_name":"张三",
	"company_name":"中至", 
	"email":"1020300659@qq.com",
	"password":"123456",
	"dpassword":"123456",
	"qq":"1020300659",
	"phone":"18679958179"
	
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
## 用户登录
### 请求URL:
- POST host/api/login

### 参数:
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| email | 是 | string | 邮箱 |
| password | 是 | string | 密码 (长度 6-20) |
### 返回示例:
```json
{
    "status":0,
    "data":[{"email":"1020300659@qq.com","app_key":"dasd"}]
}
```
## 验证邮箱是否存在
### 请求URL:
- POST host/api/signup/{email}

### 请求示例:
```json
{
    "email":"1020300659@qq.com",
    "password":"123456"
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

## 用户注销
### 简要描述:
- 请求后会删除服务上的用户信息 所有登录过的设备都需要重新登录
### 请求URL:
- POST host/api/logout

### 返回示例:
```json
{
    "status":0,
    "msg":"success",
    "data":[]
}
```

## 文件上传
### 请求URL:
- POST host/api/upload

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




















