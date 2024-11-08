---
createTime: 2024-07-29 10:25
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 08 - Policies

## 0x01 Overview

Mihomo core 的 Policies 整体和 Clash core 的相同

## 0x02 Built-in Policies

5 种内置的 Policies

1. DIRECT
	直连，数据直接出站
2. REJECT
	拒绝，拦截数据出站
3. REJECT-DROP
	拒绝，与`REJECT`不同的是，该策略将静默抛弃请求
4. PASS
	绕过，会使匹配规则时跳过此规则
5. COMPATIBLE
	兼容，在策略组筛选不出节点时出现，等效 `DIRECT`

## 0x03 Proxy

Mihomo core 和 Clash core 一样，也支持直接使用 Proxy 做为 Policy

```yaml
- DOMAIN-SUFFIX,ipinfo.io,🇯🇵日本 01 | 专线
```

但是在 Clash verge rev 中的 Edit Rules 不能直接将 Proxy 作为 Policy。如果想要直接使用 Proxy 可以通过 Global Extend Script 实现

```ts
// Define main function (script entry)
function main(config) {
  let rules = [
    "DOMAIN-SUFFIX,ipinfo.io,🇯🇵日本 01 | 专线",
    ...config.rules
  ]
  config.rules = rules 
  return config;
}
```

## 0x04 Proxy Groups

可以将多个 proxies 定义成一个 proxy group，从而让某些 rules 使用这个 proxy group 实现定向分流

例如
```yaml
proxy-groups:
	- name: 国外媒体
	  type: select
	  proxies:
	  - DIRECT
	  - 香港A1
	  - 香港A2
	  - 台湾A1
	  - 台湾A2
	  - 美国A1
	  - 美国A2
	  - 新加坡A1
	  - 新加坡A2
	  - 日本A1
	  - 日本A2
```

具体通用字段配置参考 [策略组配置 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-groups/)

### 0x04a Types

Proxy groups 支持 5 种类型

#### select[^1]

手动选择 proxy group 使用的 proxy

```yaml
proxy-groups:
- name: Proxy
  type: select
  proxies:
  - ss
  - vmess
  - auto
  #disable-udp: true
```

#### url-test[^2]

Proxies 向指定 URL 发送 HEAD 请求，自动选择 RTT 最小的 Proxy，一般用作自动选择

```yaml
proxy-groups:
- name: "自动选择"
  type: url-test
  proxies:
  - ss
  - vmess
  url: 'https://www.gstatic.com/generate_204'
  interval: 300
  #tolerance: 50
  #lazy: true
```

#### fallback[^3]

Proxies 向指定 URL 发送 HEAD 请求，自动选择第一个可用的 Proxy(返回报文的)，一般用作故障转移

```yaml
proxy-groups:
- name: "fallback"
  type: fallback
  proxies:
  - ss
  - vmess
  url: 'https://www.gstatic.com/generate_204'
  interval: 300
  #lazy: true
```

#### load-balance[^4]

负载均衡，支持 2 种策略

1. `consistent-hashing` 将会把相同顶级域名的请求分配给策略组内的同一个代理节点
2. `round-robin` 将会把所有的请求分配给策略组内不同的代理节点

```yaml
proxy-groups:
- name: "load-balance"
  type: load-balance
  proxies:
  - ss1
  - ss2
  - vmess1
  url: 'https://www.gstatic.com/generate_204'
  interval: 300
  #lazy: true
  #strategy: consistent-hashing # or round-robin
```

#### relay[^5]

按照 proxies 的先后顺序，链式代理

```yaml
proxy-groups:
# Traffic: Clash <-> http <-> vmess <-> ss1 <-> ss2 <-> Internet
- name: "relay"
  type: relay
  proxies:
    - http
    - vmess
    - ss1
    - ss2
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[手动选择 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-groups/select/)
[^2]:[自动选择 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-groups/url-test/)
[^3]:[自动回退 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-groups/fallback/)
[^4]:[负载均衡 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-groups/load-balance/)
[^5]:[链式代理 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-groups/relay/)