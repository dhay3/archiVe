---
createTime: 2024-07-11 15:03
tags:
  - "#hash1"
  - "#hash2"
---

# Clash 03 - Clash Configuration

> [!important] 
> 一些机场会有 Clash 的单独配置，有一些配置可能并不是我们想要的。我们可以通过 GUI client 对配置进行覆写，例如在 Clash verge rev 中就是 Global Extend Config/Global Extend Script/Extend Config/Extend Script

## 0x01 Overview

Clash 的配置文件和 V2ray 逻辑上大差不差

## 0x02 Inbound[^1]

Clash 支持多种入站协议，包括

- Socks5
- HTTP(S)
- Redirect TCP
- TProxy TCP
- TProxy UDP
- Linux TUN 设备 (仅 Premium 版本)

所有协议共有一个端口以及规则集

配置如下

```yaml
# HTTP(S) 代理服务端口
# port: 7890

# SOCKS5 代理服务端口
socks-port: 7891

# HTTP(S) 和 SOCKS4(A)/SOCKS5 代理服务共用一个端口
mixed-port: 7890

# Linux 和 macOS 的透明代理服务端口 (TCP 和 TProxy UDP 重定向)
# redir-port: 7892

# Linux 的透明代理服务端口 (TProxy TCP 和 TProxy UDP)
# tproxy-port: 7893

# 设置为 true 以允许来自其他 LAN IP 地址的连接
# allow-lan: false
```

## 0x03 Outbound[^2]



### 0x03a Proxies

代理节点，根据协议不同，配置不同


## 0x04 Rules

## 0x05 DNS

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Inbound 入站 | Clash 知识库](https://clash.wiki/configuration/inbound.html)
[^2]:[Outbound 出站 | Clash 知识库](https://clash.wiki/configuration/outbound.html)