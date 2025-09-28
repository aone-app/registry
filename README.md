# registry

## 📌 项目简介 / Project Description

本项目是一个轻量级、高性能的 服务注册中心，使用 Golang 编写，提供 HTTP 与 gRPC 两种访问方式，并通过 Redis 实现节点信息的持久化与共享。
它可用于微服务架构中，作为服务发现与注册的核心组件，支持多语言客户端（Go、C#、Java）快速集成。

This project is a lightweight and high-performance Service Registry built with Golang, supporting both HTTP and gRPC protocols. It leverages Redis to persist and share node information, making it suitable for microservices environments as a central component for service discovery and registration. Multi-language SDKs (Go, C#, Java) are supported for easy integration.

## ✨ 核心特性 / Key Features
	•	🚀 双协议支持：HTTP + gRPC，方便多语言客户端接入
	•	🧠 Redis 持久化：节点信息集中存储，支持高可用与共享
	•	⚡ 轻量高效：纯 Go 实现，无外部依赖，启动速度快
	•	🧭 服务注册 & 注销 & 发现：提供基础服务发现能力，可扩展心跳和订阅机制
	•	🌍 多语言 SDK：通过 .proto 文件自动生成 Go / C# / Java 客户端

## 🧱 典型架构 / Typical Architecture
```
+-------------+       +------------------+        +-------------+
|  Service A  | <---> |   Registry(gRPC) | <----> |   Redis DB  |
+-------------+       +------------------+        +-------------+
        |                     ^
        |                     |
        v                     |
+-------------+       +------------------+
|  Service B  | <---> | Registry(HTTP)   |
+-------------+       +------------------+
```

## 🧪 主要功能 / Main Functions

-	Register(service, addr, ttl)
注册一个服务节点，并设置 TTL 过期时间
-	Deregister(service, addr)
注销服务节点
-	GetNodes(service)
查询某个服务下所有可用节点
