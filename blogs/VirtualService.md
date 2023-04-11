---
title: VirtualService
date: 2023-3-9 9:44:00
tags: [VirtualService,学习笔记]
category: VirtualService
---



### 服务网格是什么？

现代应用程序通常设计成微服务的分布式集合，每个服务执行一部分业务功能。服务网格是专门的基础设施曾，包含了组成这类体系结构的微服务网络。服务网络不仅描述了这个网络，而且还描述了分布式应用程序组件之间的交互。所以在服务之间传递数据都由服务网格控制和路由。

随着分布式服务的部署，比如基于Kubernetes的系统（规模和复杂性的增长）,可能会更加难以理解和管理。随着技术的迭代，需求变为了服务发现，负载平衡，故障恢复，度量和监视。微服务体系结构通常还有更复杂的操作需求，比如A/B测试,Canary部署,速率限制,访问控制,加密和端到端身份验证等。

服务网络可以通过可观测性,网络和安全策略来分隔应用程序的业务逻辑。通过服务网格可以链接微服务,让微服务更安全,可以更好的监控微服务。



### Istio 是什么？

Istio是一个开源服务网格,它透明地分层到现有的分布式应用程序上。它提供了一种统一和更有效的方式来保护,连接和监视服务。Istio是实现负载平衡，服务到服务身份验证和监视的路径（只需要更改很少的服务代码）它可以：

- 使用 TLS 加密、强身份认证和授权的集群内服务到服务的安全通信
- 自动负载均衡的 HTTP, gRPC, WebSocket，和 TCP 流量
- 通过丰富的路由规则、重试、故障转移和故障注入对流量行为进行细粒度控制
- 一个可插入的策略层和配置 API，支持访问控制、速率限制和配额
- 对集群内的所有流量(包括集群入口和出口)进行自动度量、日志和跟踪



Istio是为了可扩展性设计的，可以处理不同范围的部署需求。Istio的控制平面运行在Kubernetes上，可以将部署在该集群中的应用程序添加到您的网格中，将网格裸站到其他集群，甚至连接VM或者其他端点。




