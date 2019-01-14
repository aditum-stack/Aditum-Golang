# Aditum Golang

> Golang service for Aditum. Golang语言实现聊天室服务, 邮件服务, 搜索引擎服务. 通过Sidecar集成Eureka实现微服务注册.

-------------------------------------------------------------------------------

# Sidecar 

1. 通过Sidecar框架将Golang服务注册到SpringCloud Eureka中，从而能够由其他微服务进行调用。

2. 故为每个Go服务添加对应的Sidecar(Java)服务，通过实现REST接口和监听Go服务接口，实现注册。

3. Springboot应用使用apollo配置中心进行配置获取。

4. sidecar与go服务必须通过localhost通信。

远程启动参数: 

-Dapollo.configService=http://47.106.11.84:8080

### Eureka API:

需要实现 "/health" 接口，返回JSON {"status":"UP"}

## Chatroom 在线聊天室服务 30021

This is a simple chat web app written in Go

Just run the following

```
cd ./src
go get github.com/gorilla/websocket
go run main.go
```

Then point your browser to http://localhost:8000

## Email 邮件发送服务 30022

## Searcher 搜索引擎服务 30023