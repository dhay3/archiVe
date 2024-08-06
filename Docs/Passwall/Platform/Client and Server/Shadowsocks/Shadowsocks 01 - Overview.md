---
createTime: 2024-08-06 11:11
tags:
  - "#hash1"
  - "#hash2"
---

# Shadowsocks 01 - Overview

## 0x01 Overview

> Shadowsocks - A fast tunnel proxy that helps you bypass firewalls[^1]

Shadowsocks[^4] 是一个 CS 模式的翻墙代理工具，其使用的加密协议也称为 Shadowsocks(基于 Socks5，而 shadow 也意为加密)

## 0x02 History 

Shadowsocks 最早在 V2EX 上被 clowwindy 向众人公开[^3]

而在 2015 年 8 月的时候 Clowwindy 在 Github [#issue124](https://web.archive.org/web/20150822042959/https://github.com/shadowsocks/shadowsocks-iOS/issues/124#issuecomment-133630294) 中提到

> Two days ago the police came to me and wanted me to stop working on this. Today they asked me to delete all the code from GitHub. I have no choice but to obey.
> 
> I hope one day I'll live in a country where I have freedom to write any code I like without fearing.
> 
> I believe you guys will make great stuff with Network Extensions.
> 
> Cheers!

大陆蓝皮狗要求作者删除 Github 上所有关于 Shadowsocks 的代码，至此 Shadowsocks 也进入停更。好在有各种各样的 fork 以及 implementation

## 0x03 Implementation

Shadowsocks 这一个概念最早以 Python 的方式实现，具体代码看 master 分支[^3](由于 clowwindy 本人被请去喝茶 [0x02 History](#0x02%20History)，项目目前停止维护)

除了 Python 的实现方式外，也有各种各样语言的实现

- C - [shadowsocks-libev](https://github.com/shadowsocks/shadowsocks-libev)
- Go - [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)
- Rust - [shadowsocks-rust](https://github.com/shadowsocks/shadowsocks-rust)

等等，其中 Rust 社区的活跃度最高(目前仍在更新)

## 0x04 GUI Client

Shadowsocks 本体和 V2ray-core 一样，就是一个二进制的包以及一些配置文件(可以说 V2ray-core 就是借鉴了这种模式)。为此开源社区也贡献了各种 GUI 客户端

- Android - [shadowsocks-android](https://github.com/shadowsocks/shadowsocks-android)
- Windows - [shadowsocks-windows](https://github.com/shadowsocks/shadowsocks-csharp)
- MacOS - [shadowsocksX-NG](https://github.com/shadowsocks/ShadowsocksX-NG)
- Windows/MacOS/Linux - [shadowsocks-qt5](https://github.com/shadowsocks/shadowsocks-qt5)

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Shadowsocks - 维基百科，自由的百科全书](https://zh.wikipedia.org/wiki/Shadowsocks)
[^2]:[Shadowsocks | A fast tunnel proxy that helps you bypass firewalls.](https://shadowsocks.org/)
[^3]:[发一个自用了一年多的翻墙工具 shadowsocks - V2EX](https://web.archive.org/web/20140719051716/https://www.v2ex.com/t/32777)
[^4]:[GitHub - shadowsocks/shadowsocks](https://github.com/shadowsocks/shadowsocks)
[^5]:[Shadowsocks - 维基百科，自由的百科全书](https://zh.wikipedia.org/wiki/Shadowsocks)