```
关闭了外网上传
需要服务上配置hosts 10.105.105.217 upload.qinglong365.com
业务代码 请参考 forwardHandler 转发请求
```

## 基本 素材上传(最大限制32m)
### 请求URL:
- POST http://upload.qinglong365.com/upload

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| file | 是 | string | name 值 |

### 返回示例:
```json
{
        "path": "https://static.qinglong365.com/d41d8cd98f00b204e9800998ecf8427e.mp4",
        "size": "234120",
        "file_name": "output.mp4"   
}
```


## ad 素材上传(最大限制32m)
### 请求URL:
- POST http://upload.qinglong365.com/ad/upload

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| file | 是 | string | name 值 |

### 返回示例:
```json
{
  
        "duration": "9.039000",
        "height": 320,
        "path": "https://static.qinglong365.com/d41d8cd98f00b204e9800998ecf8427e.mp4",
        "size": "234120",
        "file_name": "output.mp4",
        "width": 570
    
}
```


## 公共素材上传(最大限制32m)
### 请求URL:
- POST http://upload.qinglong365.cn/public/upload

### 参数
| 参数名 | 必填 | 类型 | 说明 |
| --- | --- | --- | --- |
| file | 是 | string | name 值 |

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







