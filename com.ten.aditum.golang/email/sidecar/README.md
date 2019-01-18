# cloud sidecar

> 异构微服务接入springcloud

注册到eureka中心

使用apollo配置中心进行配置获取

sidecar与异构服务必须通过localhost通信

#### 异构服务必须实现/health接口：

返回JSON数据：

{
    "status": "UP"
}

