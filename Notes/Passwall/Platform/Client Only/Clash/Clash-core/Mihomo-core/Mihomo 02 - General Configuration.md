---
createTime: 2024-07-16 10:06
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 02 - General Configuration

> 只记录自己常用的全局配置
> 具体看官方文档 [全局配置 - 虚空终端 Docs](https://wiki.metacubex.one/config/general)

## 0x01 LAN Access[^1]

和 Clash 相同 Mihomo 也支持让 LAN 中的设备通过 Clash 访问公网

```yaml
#是否允许 LAN 中的设备通过 Clash 访问公网
#可选值 `true/false`
allow-lan: true
#Clash 监听的 IP address
#"*" 绑定所有 IP 地址
#192.168.31.31 绑定单个 IPV4 地址
#[aaaa::a8aa:ff:fe09:57d8] 绑定单个 IPV6 地址
bind-address: "*"
#允许连接的 IP 地址段，仅作用于 `allow-lan` 为 `true` 默认值为 `0.0.0.0/0`和 `::/0`
lan-allowed-ips:
- 0.0.0.0/0
- ::/0
#禁止连接的 IP 地址段，黑名单优先级高于白名单，默认值为空
lan-disallowed-ips:
- 192.168.0.3/32
#`http(s)`/`socks`/`mixed`代理的用户验证
authentication:
- "user1:pass1"
- "user2:pass2"
#设置允许跳过验证的 IP 段
skip-auth-prefixes:
- 127.0.0.1/8
- ::1/128
```

LAN 中的设备按照代理协议不同使用的地址也不同，通常使用 mixed-port 即可

- http(s)
	bind-address:port
- socks4/4a/5
	bind-address:socks-port
- https(s)/socks4/4a/5
	bind-address:mixed-port
- transparent
	bind-address:redir-port 仅代理 TCP
	bind-address:tproxy-port 可以代理 TCP/UDP

## 0x02 Proxy Mode

代理模式

```yaml
mode: rule
```

- `rule` 规则代理，按照 Rules 代理，默认使用该模式
- `global` 全局代理 (需要在 GLOBAL 策略组选择代理/策略)
- `direct` 全局直连

## 0x03 Log

Clash 内核输出日志的等级，仅在控制台和控制页面输出

```yaml
log-level: info
```

- `silent` 静默，不输出
- `error` 仅输出发生错误至无法使用的日志
- `warning` 输出发生错误但不影响运行的日志，以及 error 级别内容
- `info` 输出一般运行的内容，以及 error 和 warning 级别的日志
- `debug` 尽可能的输出运行中所有的信息

## 0x04 IPv6

> 为了安全已经解析问题，尽量避免使用 IPv6

是否允许内核接受 IPv6 流量
可选值 `true/false,`默认为 `true`

```yaml
ipv6: true
```

## 0x05 GEOIP/GEOSITE

```yaml
#geoip 数据模式
#更改 geoip 使用文件，mmdb 或者 dat，可选 true/false,true为 dat，此项有默认值 false
geodata-mode: true
#geoip 数据加载模式
#可选的加载模式如下
#standard：标准加载器
#memconservative：专为内存受限 (小内存) 设备优化的加载器 (默认值)
geodata-loader: memconservative
#自动更新 geo
geo-auto-update: false
#更新间隔单位 hour
geo-update-interval: 24
#自定义 GEOIP/GEOSITE 下载地址
geox-url:
  geoip: "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geoip.dat"
  geosite: "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geosite.dat"
  mmdb: "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/country.mmdb"
  asn: "https://github.com/xishang0128/geoip/releases/download/latest/GeoLite2-ASN.mmdb"
```

## 0x06 Miscellaneous

```yaml
#缓存
#在 Clash 官方中，profile 应为扩展配置，但在 Clash.meta, 仅作为缓存项使用
profile:
  # 可选值 true/false
  # 储存 API 对策略组的选择，以供下次启动时使用
  store-selected: true
  # 储存 fakeip 映射表，域名再次发生连接时，使用原有映射地址
  store-fake-ip: true
#开启统一延迟时，会进行两次延迟测试，以消除连接握手等带来的不同类型节点的延迟差异
#可选值 true/false
unified-delay: true
#可选值 true/false
tcp-concurrent: true
#全局 TLS 指纹，优先低于 proxy 内的 client-fingerprint。
#目前支持开启 TLS 传输的 TCP/grpc/WS/HTTP , 支持协议有 VLESS,Vmess 和 trojan.
#可选：chrome, firefox, safari, iOS, android, edge, 360, qq, random, 若选择 random, 则按 Cloudflare Radar 数据按概率生成一个现代浏览器指纹。
global-client-fingerprint: chrome
#自定义外部资源下载时使用的的 UA，默认为 clash.meta
global-ua: clash.meta
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[1]:[全局配置 - 虚空终端 Docs](https://wiki.metacubex.one/config/general/#_2)