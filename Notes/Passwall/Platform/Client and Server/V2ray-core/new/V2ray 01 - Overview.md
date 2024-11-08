---
createTime: 2024-08-05 18:27
tags:
  - "#hash1"
  - "#hash2"
---

# V2ray 01 - Overview

## 0x01 Overview


在 Shadowsocks 的作者 clowwindy 被请去喝茶后[^1]，开源社区致力于让大家更好更快的科学上网，V2ray 应运而生

## 0x02 V2ray VS Shadowsocks

1. Shadowsocks 只支持 Shadowsocks 协议。而 V2ray 支持多种协议例如 Blackhole,Trojan,Shadowsocks 以及原创的 VMess(可以说 V2ray 是一个平台)
2. Shadowsocks 本身不支持 PAC（proxy-auto-cofig）分流的功能，需要借助第三方客户端。而 V2ray 自身实现了 PAC

## 0x02 V2ray-core

V2ray-core 是 V2ray 的内核文件，你在 github 搜索可能会得到两个结果

- [v2ray/v2ray-core](https://github.com/v2ray/v2ray-core)
- [v2fly/v2ray-core](https://github.com/v2fly/v2ray-core)

v2ray/v2ray-core 是官方推出的内核，现在已经停止维护了(自从 Victoria Raymond 消失后)。而 v2fly/v2ray-core 是社区维护的核心，现在仍在更新，但是频率相对 singbox 或者是 xray 而言较慢

> [!NOTE] 
> 更推荐使用 xray-core 是 v2ray-core 的超集。虽然在的 v2ray-core 已经支持 VLESS，但是社区的活跃度明显不同

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don[V2Ray 配置指南 | 新 V2Ray 白话文指南](https://guide.v2fly.org/#%E5%B8%B8%E8%A7%81%E9%97%AE%E9%A2%98-q-a)'t want to learn.*

**references**

[^1]:[Shadowsocks - 维基百科，自由的百科全书](https://zh.wikipedia.org/wiki/Shadowsocks)
[^2]:[V2Ray 配置指南 | 新 V2Ray 白话文指南](https://guide.v2fly.org/#%E5%B8%B8%E8%A7%81%E9%97%AE%E9%A2%98-q-a)