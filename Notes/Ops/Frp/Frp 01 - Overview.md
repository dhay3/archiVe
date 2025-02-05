---
createTime: 2025-01-07 13:35
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# Frp 01 - Overview

## 0x01 Preface

> frp is a fast reverse proxy that allows you to expose a local server located behind a NAT or firewall to the Internet. It currently supports TCP and UDP, as well as HTTP and HTTPS protocol, enabling requests to be forwarded to the internal services via domain name

frp 是一个 L4/7 内网穿透工具

- 支持多种协议，例如 TCP, UDP, QUIC, WebSocket, HTTP, HTTPS, SSH 等等
- 支持 TCP Multiplexer，单个 TCP 连接可以处理多个请求
- 支持负载均衡
- 支持端口复用，通过同一个端口访问不同的服务
- 支持扩展插件，如 HTTP 转 HTTPS，Socks5 代理，文件浏览
- 支持 UI 面板

## 0x02 Architecure

frp 主要由 2 部分组成

![](https://github.com/fatedier/frp/blob/master/doc/pic/architecture.png?raw=true)

- frps

	frp server 通常部署在公网服务器上，用户通过 frp server 访问 frp client 上的服务

- frpc

	frp client 通常部署在私网服务器上，frp client 将服务暴露给 frp server

用户通过访问公网的 frps 暴露的端口来访问 frpc 上服务

## 0x03 Installation[^1]

frp 安装的方式非常简单，直接从 [Github Release](https://github.com/fatedier/frp/releases) 下载 tarball 然后

1. 将 `frps`, `frps.toml` 拷贝到公网服务器上
2. 将 `frpc`, `frpc.toml` 拷贝到私网服务器上

## 0x04 Quick Start

tarball 中提供了样例配置文件，修改一下配置就可以直接运行

`frps.toml` 配置如下

```
#frps.toml
bindPort = 7000
```

- `bindPort` 表示 frps 监听的端口

`frps.toml` 配置如下

```
serverAddr = "127.0.0.1"
serverPort = 7000

[[proxies]]
name = "test-tcp"
type = "tcp"
localIP = "127.0.0.1"
localPort = 22
remotePort = 6000

```

- `serverAddr` 表示 frps 的地址
- `serverPort` 表示 frps 监听的端口
- `[[proxies]]` 表示需要代理的服务
- `name` 被代理服务的标识符需要全局唯一
- `type` 被代理服务的协议类型
- `localIP` 被代理服务监听的本地地址
- `localPort` 被代理服务监听的本地端口
- `remotePort` frps 映射被代理服务的端口

将 `ServerAddr` 修改成 frps 的地址，假设为 `192.168.56.1`，当访问 frps `192.168.56.1:6000` 时，frp 就会将流量转发到 frpc `127.0.0.1:22`。即访问 `192.168.56.1:6000` 等价于 访问 `127.0.0.1:22`

---

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [概览 \| frp](https://gofrp.org/zh-cn/docs/overview/)
- [GitHub - fatedier/frp at master](https://github.com/fatedier/frp/tree/master)

***References***

[^1]:[安装 \| frp](https://gofrp.org/zh-cn/docs/setup/)

