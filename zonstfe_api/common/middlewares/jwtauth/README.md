# 使用go进行 JWT 验证
对于使用负载均衡的服务器来说,使用 JWT(JSON WEB TOKEN) 是一个更优的选择,session受到单台服务器的限制,一个用户登录过后就只能分配到
这一台服务器上,这和负载均衡的初衷不一致啊,而 jwt 就解决了这类的痛点

# 使用 JWT 的场景
身份验证 用户在登录过后服务器会用 jwt 返回用户可访问的资源,比如权限什么的
传递信息 通过 jwt 的header和signature可以保证payload没有被篡改,保证信息的安全
# JWT 的结构
JWT 是由header,payload,signature三部分组成的,咱们先用例子说话
- header
```
{
  "alg": "HS256",
  "typ": "JWT"
}
// base64编码的字符串`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`
```
这里规定了加密算法,hash256
- payload
```
{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true
}
// base64编码的字符串`eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9`
```
```
这里的内容没有强制要求,因为 paylaod 就是为了承载内容而存在的,不过想用规范的话也可以参考下面的
* iss: jwt签发者
* sub: jwt所面向的用户
* aud: 接收jwt的一方
* exp: jwt的过期时间，这个过期时间必须要大于签发时间
* nbf: 定义在什么时间之前，该jwt都是不可用的.
* iat: jwt的签发时间
* jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
```

- signature
是用 header + payload + secret组合起来加密的,公式是:
```
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

这里 secret就是自己定义的一个随机字符串,这一个过程只能发生在 server 端,会随机生成一个 hash 值

这样组合起来之后就是一个完整的 jwt 了:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.4c9540f793ab33b13670169bdf444c1eb1c37047f18e861981e14e34587b1e04

```
这里有一个用 go 加密和验证 jwt 的 demo

# 总结
选择 jwt 最大的理由:

内容有公钥私钥,可以保证内容的合法性
token 中可以包含很多信息
不过 jwt 不保证的安全问题:

因为header,paylaod是 base64编码,相当于明文可见的,因此不能在payload中放入敏感信息
并不能保证数据传输时会不会被盗用,这一点和 sessionID 一样,因此不要迷信它有多高的安全性..