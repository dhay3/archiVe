---
createTime: 2024-07-29 11:24
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 04 - Inbound

## 0x01 Overview

Inbound 指定了流量进入 Clash 的方式

## 0x02 Listening Ports

```
# http(s)
port: 7890
# socks4/4a/5
socks-port: 7891
# http(s)/socks4/4a/5
mixed-port: 7892
# TCP 
redir-port: 7893
# TCP UDP
tproxy-port: 7894
```

## 0x03 Tun

具体逻辑可以看 [Clash 05 - Tun](../Clash%2005%20-%20Tun.md)

Mihomo 完整配置 [Tun - 虚空终端 Docs](https://wiki.metacubex.one/config/inbound/tun/)

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

