---
createTime: 2024-07-11 14:49
tags:
  - "#hash1"
  - "#hash2"
---

# Clash 02 - How Clash Works

## 0x01 Overview

Clash 工作原理和 V2ray 类似，分为 client 和 server

![](https://clash.wiki/assets/connection-flow.a72146ab.png)

- Inbound
	也被称为 入站，是本地监听的部分，它通过打开一个本地端口并监听传入的连接来工作. 当连接进来时, Clash 会查询配置文件中配置的规则, 并决定连接应该去哪个 Outbound 出站.
- Outbound
	Outbound 出站是连接到远程端的部分. 根据配置的不同, 它可以是一个特定的网络接口、一个代理服务器或一个[策略组](https://clash.wiki/configuration/outbound.html#proxy-groups-策略组).

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

