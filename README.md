# Srv_gateway
这是一个简单业务网关监听项目。目前已经完成的功能：
- [x] 反向代理和请求路由
- [x] 基于jwt进行令牌（token）鉴权
- [ ] 支持路由策略动态配置
- [ ] 错误处理
- [ ] 单元测试

目前已经实现的功能是：  
1 对 /login 和 /register 请求路径 进行放行，目的是让未登录用户拿到令牌  
2 对其 请求路径 进行鉴权，通过对请求的 Header.Auth 携带 token 进行 加密方式（HMAC）、密钥和是否过期 进行鉴定，鉴定通过则 路由到网关配置的 目标host:port 和 请求路径中，鉴定失败直接由网关拦截。

## Background
就微服务网关方面，nginx已经做得很好，但是基本是系统级别的，当然nginx也支持插件，但是始终不是自己的技术栈，并且我还是希望能有一个定制化的 业务网关，所以有了本项目。

## Install
下载源码，准备好 Go 1.18 环境

## Usage
在项目根目录：
```
$ go run .
```

## Contributing
提个Issues说出你的需求 或 提个PR实现你的想法，我会测试并整合到本项目。

## License
MIT © heytheww