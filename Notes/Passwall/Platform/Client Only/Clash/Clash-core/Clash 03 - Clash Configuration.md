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

具体看 [Outbound 出站 | Clash 知识库](https://clash.wiki/configuration/outbound.html#proxies-%E4%BB%A3%E7%90%86%E8%8A%82%E7%82%B9)

### 0x03b Proxy Groups

策略组，可以让多组 Rules 使用同一代理，实现分流

具体看 [Outbound 出站 | Clash 知识库#proxy-groups-策略组](https://clash.wiki/configuration/outbound.html#proxy-groups-%E7%AD%96%E7%95%A5%E7%BB%84)

支持 5 种 Proxy Groups

#### relay

代理服务器作为 relay 中继，从而实现链式代理，不支持 UDP

#### url-test

Proxies 向指定 URL 发送 HEAD 请求，自动选择 RTT 最小的 Proxy，一般用作自动选择

```yaml
   proxy-groups:
     - name: 
       type: url-test
   	proxies:
   	  - HK01
   	  - HK02
   	  - SG01
   	  - SG02
   	  - US01
   	  - US02
   	  url: http://www.gstatic.com/generate_204
   	  interval: 86400
```

#### fall-back

Proxies 向指定 URL 发送 HEAD 请求，自动选择第一个可用的 Proxy(返回报文的)，一般用作故障转移

```yaml
   proxy-groups:
     - name: 
       type: fall-back
   	proxies:
   	  - HK01
   	  - HK02
   	  - SG01
   	  - SG02
   	  - US01
   	  - US02
   	  url: http://www.gstatic.com/generate_204
   	  interval: 86400
```

#### load-balance

随机使用 proxy，但是 eTDL + 1[^4] 的请求会使用同一个 Proxy

#### select 

手动选择 proxy

```yaml
   proxy-groups:
     - name: 
       type: select
   	proxies:
   	  - HK01
   	  - HK02
   	  - SG01
   	  - SG02
   	  - US01
   	  - US02
```

### 0x03c Proxy Providers

从外部加载代理节点，可以让配置更加简洁

支持 2 种方式加载代理节点

#### http

使用 http/https 拉取 Proxies 列表

```yaml
proxy-providers:
  provider1:
    type: http
    url: "url"
    interval: 3600
    path: ./provider1.yaml
    # filter: 'a|b' # golang regex 正则表达式
    health-check:
      enable: true
      interval: 600
      # lazy: true
      url: http://www.gstatic.com/generate_204
```

#### file

使用本地文件作为 Proxies 列表

```yaml
proxy-providers:
  test:
    type: file
    path: /test.yaml
    health-check:
      enable: true
      interval: 36000
      url: http://www.gstatic.com/generate_204
```

test.yaml 如下

```yaml
proxies:
  - name: "ss1"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"

  - name: "ss2"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    plugin: obfs
    plugin-opts:
      mode: tls
```

## 0x04 Rules

策略规则，由 3 部分组成
```
# 类型,参数,策略(,no-resolve)
TYPE,ARGUMENT,POLICY(,no-resolve)
```

no-resolve 跳过 Rule 的 DNS 解析，通常和 CIDR 一起使用，例如
```
IP-CIDR,91.108.4.0/22,policy,no-resolve
```

### 0x04a TYPE,ARGUMENT

和网络设备中的 ACL 相同，按照从上往下的优先级去匹配

#### DOMAIN

ARGUMENT: 全域名

匹配全域名

例如：只匹配 baidu.com，不匹配 news.baidu.com
```
DOMAIN,baidu.com,policy
```

#### DOMAIN-SUFFIX

AGUMENT: SUFFIX

匹配以 ARGUMENT 结尾的域名

例如：匹配任何以 youtube.com 结尾的域名
```
DOMAIN-SUFFIX,youtube.com,policy
```

#### DOMAIN-KEYWORD

ARGUMENT: KEYWORD

匹配包含 ARGUMENT 的域名

例如：匹配任何包含 google 关键字的域名 
```
DOMAIN-KEYWORD,google,policy
```

#### GEOIP

ARGUMENT: GEOIP CODE

匹配包含 GEOIP CODE 的地址

例如：匹配中国 IP 的报文
```
GEOIP,CN,policy
```

#### IP-CIDR

> 如果目标为域名，Clash 会先使用内置的 DNS 将域名解析
> 可以使用 no-resolve 跳过 DNS 解析
> 
> **IPv6 地址必须在 square bracket 中**
> 例如 \[aaaa::a8aa:ff:fe09:57d8\]

ARGUMENT: IPv4 CIDR

匹配目标地址在 ARGUMENT 中的报文

例如：匹配回环地址段
```
IP-CIDR,127.0.0.0/8,policy
```

#### IP-CIDR6

> 如果目标为域名，Clash 会先使用内置的 DNS 将域名解析
> 可以使用 no-resolve 跳过 DNS 解析

ARGUMENT：IPv6 CIDR

匹配目标地址在 ARGUMENT 中的报文
```
IP-CIDR6,2620:0:2d0:200::7/32,policy
```

#### SRC-IP-CIDR

ARGUMENT: IPv4 CIDR

匹配源地址在 AGRUMENT 中的报文
```
SRC-IP-CIDR,192.168.1.201/32,policy
```

#### SRC-PORT

ARGUMENT: Port

匹配源端口是 ARGUMENT 的报文
```
SRC-PORT,80,policy
```

#### DST-PORT

ARGUMENT: Port

匹配目的端口是 ARGUMENT 的报文
```
DST-PORT,80,policy
```

#### PROCESS-NAME

ARGUMENT: PROCESS-NAME

匹配进程是 ARGUMENT 发送的报文

例如：匹配 nc 发送的报文
```
PROCESS-NAME,nc,policy
```

#### PROCESS-PATH

ARGUMENT: PROCESS-PATH

匹配进程路径是 AGRUMENT 发送的报文

例如：匹配从 `/usr/local/bin/nc` 发送的报文 
```
PROCESS-PATH,/usr/local/bin/nc,DIRECT
```

#### IPSET-IP

从外部加载规则集，只适用于 ipset

#### RULE-SET

从外部加载规则集，让配置更加简洁

支持 3 种加载方式

- domain

	```yaml
	rule-providers:
	  apple:
	    behavior: "domain" # domain, ipcidr or classical (仅限 Clash Premium 内核)
	    type: http
	    url: "url"
	    # format: 'yaml' # or 'text'
	    interval: 3600
	    path: ./apple.yaml
	rules:
	  - RULE-SET,apple,REJECT
	```
	apple.yaml 如下
	```yaml
	payload:
	  - '+.apple.com'
	  - '+.ipad.com'
	  - '+.iphone.com'
	```

- ipcidr
	```yaml
	rule-providers:
		local:
			behavior: "ipcidr" # domain, ipcidr or classical (仅限 Clash Premium 内核)
			type: http
			url: "url"
			# format: 'yaml' # or 'text'
			interval: 3600
			path: ./local.yaml
	rules:
		- RULE-SET,local,REJECT
   	```
	local.yaml 如下
	```yaml
	payload:
	  - '192.168.1.0/24'
	  - '10.0.0.0.1/32'
	```

- classical
	只在 Premium core 中支持
	```yaml
	rule-providers:
		google:
			behavior: "classical" # domain, ipcidr or classical (仅限 Clash Premium 内核)
				​￼rule-providers:
			   		​￼google:
			   			behavior: "classical" # domain, ipcidr or classical (仅限 Clash Premium 内核)
			   			type: http
			   			url: "url"
			   			# format: 'yaml' # or 'text'
			   			interval: 3600
			   			path: ./google.yaml
			   	​￼rules:
			   		- RULE-SET,local,REJECTtype: http
			url: "url"
			# format: 'yaml' # or 'text'
			interval: 3600
			path: ./google.yaml
	rules:
		- RULE-SET,local,REJECT
   	```
	google.yaml 如下
	```yaml
	payload:
	  - DOMAIN-SUFFIX,google.com
	  - DOMAIN-KEYWORD,google
	  - DOMAIN,ad.com
	  - SRC-IP-CIDR,192.168.1.201/32
	  - IP-CIDR,127.0.0.0/8
	  - GEOIP,CN
	  - DST-PORT,80
	  - SRC-PORT,7777
	  # MATCH 在这里并不是必须的
	```

#### MATCh

匹配所有没有匹配规则的报文，必须存在，通常作为最后一条规则
```
MATCH,policy
```

### 0x04b POLICY

策略(类似于 iptables 中的 targets)，有 4 种
1. DIRECT 
	匹配的报文直连
	例如：当报文的 HOST header 匹配 baidu.com 时直连
	```
	DOMAIN,baidu.com,DIRECT
	```
	
2. REJECT
	将匹配的报文丢弃
	例如：当报文的 HOST header 中包含 tracking 关键字 时将报文丢弃
	```
	DOMAIN-KEYWORD,tracking,REJECT
	```
1. Proxy
	将匹配的报文路由到指定的 Proxy
	例如：当报文的 HOST header 
	```
	DOMAIN-SUFFIX,google.com,SG-A1
	```
1. Proxy Group
	将匹配的报文路由到指定的 [0x03b Proxy Groups](#0x03b%20Proxy%20Groups)

## 0x05 DNS[^5]

DNS 是 Clash 中最复杂也是最让人难理解的部分，同时也是 Clash 最核心的功能之一

### 0x05a DNS wildcard

> [!NOTE]
> 包含 `*`，`.`，`+` 的域名必须使用 single quote 包裹
> 同时静态域名优先级大于含有 wildcard
> 例如 $foo.example.com > *.example.com > .example.com > +.example.com$

支持 3 种 wildcard

#### asterisk(`*`)

匹配单级域名

| 表达式             | 匹配                            | 不匹配                     |
| ------------------ | ------------------------------- | -------------------------- |
| `*.google.com`     | `www.google.com`                | `google.com`               |
| `*.bar.google.com` | `foo.bar.google.com`            | `bar.google.com`           |
| `*.*.google.com`   | `thoughtful.sandbox.google.com` | `one.two.three.google.com` |

#### peroid(`.`)

匹配多级域名

| 表达式        | 匹配                            | 不匹配       |
| ------------- | ------------------------------- | ------------ |
| `.google.com` | `www.google.com`                | `google.com` |
| `.google.com` | `thoughtful.sandbox.google.com` | `google.com` |
| `.google.com` | `one.two.three.google.com`      | `google.com` |

#### plus(`+`)

类似于 DOMAIN-SUFFIX rule

匹配多级域名，最全

| 表达式         | 匹配                            |
| -------------- | ------------------------------- |
| `+.google.com` | `google.com`                    |
| `+.google.com` | `www.google.com`                |
| `+.google.com` | `thoughtful.sandbox.google.com` |
| `+.google.com` | `one.two.three.google.com`      |

### 0x05b fake-ip

> [!important] 
> Clash 和其他 Client 不一样的点在于使用了 fake-ip
> 本地不会直接解析 DNS，会转由 Gateway(在 Clash 中为配置中的 DNS server) 来解析 DNS，在一定程度上可以防止 GFW DNS pollution
> 
> 具体可以看 RFC3089 DNS Name Resolving Procedure[^6]

```yaml
dns:
  enhanced-mode: fake-ip
  fake-ip-range: 198.18.0.1/16 # Fake IP 地址池 CIDR
  # 此列表中的主机名将不会使用 Fake IP 解析
  # 即, 对这些域名的请求将始终使用其真实 IP 地址进行响应
  fake-ip-filter:
     - '*.lan'
     - localhost.ptlogin2.qq.com
```

### 0x05c default-nameserver

解析 NS records 的 nameserver

```yaml
dns:
  # 这些 名称服务器(nameservers) 用于解析下列 DNS 名称服务器主机名.
  # 仅指定 IP 地址
  default-nameserver:
    - 114.114.114.114
    - 8.8.8.8
```

### 0x05d nameserver

所有 A/AAAA/NAME record 请求使用的 nameserver

```yaml
dns:
  # 支持 UDP、TCP、DoT、DoH. 您可以指定要连接的端口.
  # 所有 DNS 查询都直接发送到名称服务器, 无需代理
  # Clash 使用第一个收到的响应作为 DNS 查询的结果.
  nameserver:
    - 114.114.114.114 # 默认值
    - 8.8.8.8 # 默认值
    - tls://dns.rubyfish.cn:853 # DNS over TLS
    - https://1.1.1.1/dns-query # DNS over HTTPS
    - dhcp://en0 # 来自 dhcp 的 dns
    # - '8.8.8.8#en0'
```

### 0x05e fallback/fallback-filter

#### fallback

所有 A/AAAA/NAME record 请求使用的备选 nameserver(和 `nameserver` 并发请求)

```yaml
dns:
  # 当 `fallback` 存在时, DNS 服务器将向此部分中的服务器
  # 与 `nameservers` 中的服务器发送并发请求
  # 当 GEOIP 国家不是 `CN` 时, 将使用 fallback 服务器的响应
  fallback:
    - tcp://1.1.1.1
    - 'tcp://1.1.1.1#en0'
```

#### fallback-filter

当请求匹配 fallback-filter 会使用 `nameserver` 中的 nameserver 用于 DNS 解析，反之使用 `fallback`

```yaml
dns:
  # 如果使用 `nameservers` 解析的 IP 地址在下面指定的子网中,
  # 则认为它们无效, 并使用 `fallback` 服务器的结果.
  #
  # 当 `fallback-filter.geoip` 为 true 且 IP 地址的 GEOIP 为 `CN` 时,
  # 将使用 `nameservers` 服务器解析的 IP 地址.
  #
  # 如果 `fallback-filter.geoip` 为 false(即不按照 geoip 过滤), 且不匹配 `fallback-filter.ipcidr`,
  # 则始终使用 `nameservers` 服务器的结果(即解析记录在这个段中的，都被认为是污染的，会重新使用 `fallback` 解析)
  #
  # 这是对抗 DNS 污染攻击的一种措施.
  fallback-filter:
    geoip: true
    geoip-code: CN
    ipcidr:
      - 240.0.0.0/4
  # 这些域名直接被认为是污染的，会直接使用 fallback
    domain:
      - '+.google.com'
      - '+.facebook.com'
      - '+.youtube.com'
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Inbound 入站 | Clash 知识库](https://clash.wiki/configuration/inbound.html)
[^2]:[Outbound 出站 | Clash 知识库](https://clash.wiki/configuration/outbound.html)
[^3]:[Rules 规则 | Clash 知识库](https://clash.wiki/configuration/rules.html)
[^4]:[eTLD - MDN Web Docs Glossary: Definitions of Web-related terms | MDN](https://developer.mozilla.org/en-US/docs/Glossary/eTLD)
[^5]:[Clash DNS | Clash 知识库](https://clash.wiki/configuration/dns.html)
[^6]:[RFC 3089:  A SOCKS-based IPv6/IPv4 Gateway Mechanism](https://www.rfc-editor.org/rfc/rfc3089)