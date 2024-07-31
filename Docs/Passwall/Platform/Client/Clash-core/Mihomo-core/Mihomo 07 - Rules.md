---
createTime: 2024-07-29 09:12
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 07 - Rules

## 0x01 Overview

Rules 是 Clash 中的路由规则，按照数据口中特定的内容选择策略进行分流。和网络设备中的 ACL 相同，按照从上往下的优先级去匹配。Mihomo core 在 Clash core 的基础上额外支持了更多的规则，例如 PROCESS-NAME,UID,NETWORK 等等

由 3 部分组成
```
# 类型,参数,策略(,no-resolve)
TYPE,ARGUMENT,POLICY(,no-resolve)
```

## 0x02 Types

Mihomo 支持如下几种 Rules

### 0x02a Domain

匹配完整域名

```
- DOMAIN,ad.com,REJECT
```

### 0x02b DOMAIN-SUFFIX

匹配域名后缀，等价于 `+.google.com` 的逻辑

```
- DOMAIN-SUFFIX,google.com,PROXY
```

### 0x02c DOMAIN-KEYWORD

匹配包含关键字的域名

```
- DOMAIN-KEYWORD,ads,REJECT
```

### 0x02d DOMAIN-REGEX

使用正则匹配域名

```
- DOMAIN-REGEX,^abc.*com,PROXY
```

### 0x02e GEOSITE

使用 geosite.data 去匹配域名，具体可以使用的值参考 [GitHub - Loyalsoldier/v2ray-rules-dat: 🦄 🎃 👻 V2Ray 路由规则文件加强版，可代替 V2Ray 官方 geoip.dat 和 geosite.dat，兼容 Shadowsocks-windows、Xray-core、Trojan-Go、leaf 和 hysteria。Enhanced edition of V2Ray rules dat files, compatible with Xray-core, Shadowsocks-windows, Trojan-Go, leaf and hysteria.](https://github.com/Loyalsoldier/v2ray-rules-dat)(具体可以使用的值参考 realse 分支)具体可以使用的值参考 realse 分支

```
- GEOSITE,apple-cn,DIRECT
```

### 0x02f IP-CIDR/IP-CIDR6

按照 IP 地址 CIDR 匹配

```
- IP-CIDR,10.0.0.0/8,DIRECT
- IP-CIDR,100.64.0.0/10,DIRECT
- IP-CIDR,127.0.0.0/8,DIRECT
- IP-CIDR,172.16.0.0/12,DIRECT
- IP-CIDR,192.168.0.0/16,DIRECT
- IP-CIDR6,::1/128,DIRECT
- IP-CIDR6,fc00::/7,DIRECT
- IP-CIDR6,fe80::/10,DIRECT
- IP-CIDR6,fd00::/8,DIRECT
```

### 0x02g IP-SUFFIX

按照 IP 后缀匹配，通常不用，直接使用 IP-CIDR 即可

```
- IP-SUFFIX,8.8.8.8/24,PROXY
```

### 0x02h IP-ASN

按照 IP 地址的 Autonomous System Number 匹配

```
- IP-ASN,15169,PROXY
```

### 0x02j GEOIP

按照 IP 地址的国家代码匹配

```
- GEOIP,CN,DIRECT
```

### 0x02k SRC-GEOIP

按照源 IP 的国家代码匹配

```
- SRC-GEOIP,cn,DIRECT
```

### 0x02l SRC-IP-ASN

按照源 IP 的 ASN 匹配

```
- SRC-IP-ASN,9808,DIRECT
```

### 0x02o SRC-IP-CIDR

按照源 IP 的 CIDR 匹配

```
- SRC-IP-CIDR,192.168.1.201/32,DIRECT
```

### 0x02p SRC-IP-SUFFIX

按照源 IP 的后缀匹配

```
- SRC-IP-SUFFIX,192.168.1.201/8,DIRECT
```

### 0x02q DST-PORT

按照目的端口匹配

```
- DST-PORT,80,DIRECT
```

### 0x02r SRC-PORT

按照源端口匹配

```
- SRC-PORT,7777,DIRECT
```

### 0x02s IN-PORT

按照入站端口匹配

```
- IN-PORT,7890,PROXY
```

### 0x02t IN-TYPE

按照入站类型(协议)匹配

```
- IN-TYPE,SOCKS/HTTP,PROXY
```

### 0x02u IN-USER

按照入站用户名匹配

```
- IN-USER,mihomo,PROXY
```

### 0x02v IN-NAME

按照入站名匹配

```
- IN-NAME,ss,PROXY
```

### 0x02w PROCESS-PATH

按照进程路径匹配

```
- PROCESS-PATH,/usr/bin/wget,PROXY
```

### 0x02x PROCESS-PATH-REGEX

按照进程路径正则匹配

```
- PROCESS-PATH-REGEX,.*bin/wget,PROXY
```

### 0x02y PROCESS-NAME

按照进程名匹配

```
- PROCESS-NAME,curl,PROXY
```

### 0x02z PROCESS-NAME-REGEX

按照进程名正则匹配

```
- PROCESS-NAME-REGEX,(?i)Telegram,PROXY
```

### 0x02a1 UID

按照 UID 匹配，只在 Linux 上生效

```
- UID,1001,DIRECT
```

### 0x02b1 NETWORK

匹配 TCP 或者是 UDP

```
- NETWORK,udp,DIRECT
```

### 0x02c1 DSCP

按照 IP DSCP 匹配

```
- DSCP,4,DIRECT
```

## 0x03 RULE-SET

按照 rule-providers 中的 name 匹配

```
- RULE-SET,google,PROXY
```

### 0x03a rule-providers

> 推荐使用 [ios\_rule\_script/rule/Clash at master · blackmatrix7/ios\_rule\_script · GitHub](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/Clash) 作为上游
> 参考 [GitHub - Loyalsoldier/clash-rules: 🦄️ 🎃 👻 Clash Premium 规则集(RULE-SET)，兼容 ClashX Pro、Clash for Windows 等基于 Clash Premium 内核的客户端。](https://github.com/Loyalsoldier/clash-rules?tab=readme-ov-file)

具体字段含义看 [规则集合 - 虚空终端 Docs](https://wiki.metacubex.one/config/rule-providers/)

```yaml
rule-providers:
  google:
    type: http
    path: ./rule1.yaml 
    url: "https://raw.githubusercontent.com/../Google.yaml"
    interval: 1440
    proxy: DIRECT
    behavior: classical
    format: yaml
```

behavior 通常使用 classical 即可，支持所有 Types of Rules

rule-providers 中的内容参考 [规则集合内容 - 虚空终端 Docs](https://wiki.metacubex.one/config/rule-providers/content/)

## 0x04 SUB-RULE

使用子规则，类似与 iptables 中的 goto

```
- SUB-RULE,(DOMAIN-SUFFIX,google.com),sub-rule
```

### 0x04a sub-rules

```
sub-rules:
  rule1:
    - DOMAIN-SUFFIX,baidu.com,DIRECT
    - MATCH,PROXY
  sub-rule2:
    - IP-CIDR,1.1.1.1/32,REJECT
    - IP-CIDR,8.8.8.8/32,ss1
    - DOMAIN,dns.alidns.com,REJECT
```

## 0x05 AND/OR/NOT

逻辑规则，可以规则进行组合或者取反

```
- AND,((DOMAIN,baidu.com),(NETWORK,UDP)),DIRECT
- OR,((NETWORK,UDP),(DOMAIN,baidu.com)),REJECT
- NOT,((DOMAIN,baidu.com)),PROXY
```

## 0x06 MATCH

匹配所有没有匹配规则的报文，必须存在，通常作为最后一条规则

```
- MATCH,auto
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

