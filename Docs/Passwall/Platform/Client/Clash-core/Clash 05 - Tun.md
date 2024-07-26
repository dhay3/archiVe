---
createTime: 2024-07-16 13:08
tags:
  - "#Passwall"
  - "#Clash"
---

# Clash 05 - Tun

## 0x01 Overview

Tun 是 Premium core 中最为核心的功能之一，可以让网络中的设备真正得实现全局代理

## 0x02 TUN/TAP

在介绍 Clash Tun 前，需要先了解一下 TUN/TAP 是什么

> TUN/TAP provides packet reception and transmission for user space programs. It can be seen as a simple Point-to-Point or Ethernet device, which, instead of receiving packets from physical media, receives them from user space program and instead of sending packets via physical media writes them to the user space program.[^1]

TUN(TUN)/TAP 是系统上虚拟的 network devices。可以将经过虚拟 network devices 的数据包发送到**用户态的进程**，从而实现数据包的二次封装(这一过程也被称为 tunnleing encapsulation[^2]) 

可以参考这一实例，自己构建一个 TUN network device
[What Is the TUN Interface Used For? | Baeldung on Linux](https://www.baeldung.com/linux/tun-interface-purpose)

### 0x02a TUN vs TAP[^3]

TUN/TAP 最大的区别就在于

- TUN network device 工作在 OSI 3 层(即网络层)，接收 3 层数据包，不可以处理 Ethernet header
- TAP network device 工作在 OSI 2 层(即数据链路层)，接收 2 层数据包，可以处理 Ethernet header

即 TAP 是 TUN 的超集，TUN 能处理的数据包，TAP 都能处理

![](https://upload.wikimedia.org/wikipedia/commons/a/af/Tun-tap-osilayers-diagram.png)

### 0x02b How does TUN work

> [!NOTE]
> 推荐看看 [Linux Tun/Tap 介绍-赵化冰的博客 | Zhaohuabing Blog](https://www.zhaohuabing.com/post/2020-02-24-linux-taptun/) 这篇博文，写得非常浅显易懂

在使用 Physical NIC 上网时，应用或者系统通过 Socket API 调用系统的网络协议栈，直接将报文发送到 PNIC 并通过 Wire 传输到 endpoints

![width:500](../../../../../Excalidraw/Drawing%202024-07-22%2010.32.42.excalidraw)

而在使用 TUN 的过程中，报文会通过 TUN VNIC 发送到监听 TUN VNIC 的应用，由应用决定该怎么发送报文

![height:800](../../../../../Excalidraw/Drawing%202024-07-22%2012.57.58.excalidraw)

## 0x03 Clash Tun Basic Logical

### 0x03a Clash Tun PBR

> 如果想要详细了解原理，还需要看源码
> [GitHub - MetaCubeX/Mihomo at Meta](https://github.com/MetaCubeX/mihomo/tree/Meta)

未启用 Clash tun 前，可以观察到系统在收到数据包时，会按照 local/main/default 的 route table 去匹配路由

```
$ ip rule show
0:      from all lookup local
32766:  from all lookup main
32767:  from all lookup default

$ ip route show table all
default via 10.100.4.1 dev enp4s0 proto dhcp src 10.100.4.222 metric 100
default via 10.100.12.1 dev wlp5s0 proto dhcp src 10.100.13.47 metric 600
...
```

启用 Clash tun 后，会新增 一个 TUN network device

```
$ ip tuntap
Mihomo: tun
```

以及若干 route rules 和 route table 2022

```
$ ip rule show
0:      from all lookup local
9000:   from all to 198.18.0.0/30 lookup 2022
9001:   from all lookup 2022 suppress_prefixlength 0
9002:   not from all dport 53 lookup main suppress_prefixlength 0
9002:   from all ipproto icmp goto 9010
9002:   from all iif Mihomo goto 9010
9003:   not from all iif lo lookup 2022
9003:   from 0.0.0.0 iif lo lookup 2022
9003:   from 198.18.0.0/30 iif lo lookup 2022
9010:   from all nop
32766:  from all lookup main
32767:  from all lookup default

$ ip r show table all
default dev Mihomo table 2022
default via 10.100.4.1 dev enp4s0 proto dhcp src 10.100.4.222 metric 100
default via 10.100.12.1 dev wlp5s0 proto dhcp src 10.100.13.47 metric 600
...
```

详细拆解一下开启 Clash tun 后新增的 route rules

- `9000:   from all to 198.18.0.0/30 lookup 2022`
	到 198.18.0.0/30 的 3 层数据包会使用 2022 route table
- `9001:   from all lookup 2022 suppress_prefixlength 0`
	所有 3 层数据包都使用 2022 table 路由，如果匹配 default route 就会使用下一条 route rule
	*suppress_prefixlength reject routing decisions that have a prefix length of NUMBER or less.*
	prefix length 为 0 的路由，只有 default route 0.0.0.0/0
- `9002:   not from all dport 53 lookup main suppress_prefixlength 0`
	Destination port 非 53 的数据包使用 main route table，如果匹配 default route 就会使用下一条 route rule
- `9002:   from all ipproto icmp goto 9010`
	ICMP 数据包跳转 9010
- `9002:   from all iif Mihomo goto 9010`
	所有到 Mihomo iface 的数据包到 9010
- `9003:   not from all iif lo lookup 2022`
	所有到非 lo iface 的数据包使用 2022 route table
- `9003:   from 0.0.0.0 iif lo lookup 2022`
	所有到 lo iface 地址为 0.0.0.0 的数据包使用 2022 route table
- `9003:   from 198.18.0.0/30 iif lo lookup 2022`
	所有到 lo iface 地址为 198.18.0.0/30 的数据包使用 2022 route table
- `9010:   from all nop`
	不做任何操作

核心就是 9003 中的第 2 条 route rule，大部分数据包都会匹配 9003 route rule 从而使用 2022 route table

其中只有一条路由，即默认路由

```
$ ip route show table 2022
default dev Mihomo
```

当 Destination IP address 匹配 0.0.0.0/0 后就会将数据包发送到 Mihomo 这一 TUN network device

```
$ ip tuntap list
Mihomo: tun
```

剩下的逻辑就会由 Clash 来处理

***简单的说 Clash Tun 就是通过 PBR 实现真正的全局代理***

### 0x03a Tun Stack

Tun stack 指的是 clash 在开启 Tun 的情况下传输流量使用的网络协议栈，clash 默认提供 2 种
1. system
	直接调用系统协议栈，性能以及兼容性最好
2. gvisor[^5]
	为应用单独模拟一个系统协议栈，更安全，但是相对的性能和兼容性较差

> [!NOTE]
> 在 mihomo core 中还支持 mixed 指 tcp 流量使用 system，UDP 流量使用 gvisor

### 0x03b Clash Tun configuration

tun 部分相关的配置如下

```yaml
tun:
  enable: true
  stack: system # or gvisor
  # dns-hijack:
  #   - 8.8.8.8:53
  #   - tcp://8.8.8.8:53
  #   - any:53
  #   - tcp://any:53
  auto-route: true # manage `ip route` and `ip rules`
  auto-redir: true # manage nftable REDIRECT
  auto-detect-interface: true # 与 `interface-name` 冲突
```

但是需要提一嘴的是 TUN network device 需要借助 fake-ip-range 来设置 IP Address
具体实现代码可以看 [mihomo/config/config.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go) 中 `parseTun` 函数
```go
func parseTun(rawTun RawTun, general *General) error {
	tunAddressPrefix := T.FakeIPRange()
	if !tunAddressPrefix.IsValid() {
		tunAddressPrefix = netip.MustParsePrefix("198.18.0.1/16")
	}
	tunAddressPrefix = netip.PrefixFrom(tunAddressPrefix.Addr(), 30)

	if !general.IPv6 || !verifyIP6() {
		rawTun.Inet6Address = nil
	}
```

fake-ip-range 默认为 198.18.0.1/16，从而得出 TUN network device address 为 198.18.0.1/30，这一点在 `ip a s dev Mihomo` 中也可以得到证实
```sh
$ ip a s dev Mihomo
10: Mihomo: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc cake state UNKNOWN group default qlen 500
    link/none
    inet 198.18.0.1/30 brd 198.18.0.3 scope global Mihomo
       valid_lft forever preferred_lft forever
    inet6 fe80::9dcb:c7d7:93c9:affb/64 scope link stable-privacy proto kernel_ll
       valid_lft forever preferred_lft forever
```

30 位 netmask 可以推出 mihomo core 使用 198.18.0.2 作为自己的通信 IP

## 0x04 Clash Tun Analyze

> [!important]
> 下面的例子使用 Clash verge rev 作为 GUI client，使用 mihomo core
> 未在 emulator 中模拟纯净环境，以访问 `www.google.com` 为例

通过 Clash verge rev 来启动 mihomo core 默认会按照下面逻辑生成启动配置。具体代码逻辑可以看 [clash-verge-rev/src-tauri/src/config/clash.rs at main · clash-verge-rev/clash-verge-rev · GitHub](https://github.com/clash-verge-rev/clash-verge-rev/blob/main/src-tauri/src/config/clash.rs)
```ts
    pub fn template() -> Self {
        let mut map = Mapping::new();
        let mut tun = Mapping::new();
        tun.insert("stack".into(), "gvisor".into());
        tun.insert("device".into(), "Mihomo".into());
        tun.insert("auto-route".into(), true.into());
        tun.insert("strict-route".into(), false.into());
        tun.insert("auto-detect-interface".into(), true.into());
        tun.insert("dns-hijack".into(), vec!["any:53"].into());
        tun.insert("mtu".into(), 1500.into());
        #[cfg(not(target_os = "windows"))]
        map.insert("redir-port".into(), 7895.into());
        #[cfg(target_os = "linux")]
        map.insert("tproxy-port".into(), 7896.into());
        map.insert("mixed-port".into(), 7897.into());
        map.insert("socks-port".into(), 7898.into());
        map.insert("port".into(), 7899.into());
        map.insert("log-level".into(), "info".into());
        map.insert("allow-lan".into(), false.into());
        map.insert("mode".into(), "rule".into());
        map.insert("external-controller".into(), "127.0.0.1:9097".into());
        map.insert("secret".into(), "".into());
        map.insert("tun".into(), tun.into());

        Self(map)
    }
```

### 0x04a Clash Tun Disabled

在没有开启 Clash Tun 的情况下 Clash 全局配置如下(`mixed-port`，`socks-port`，`port` 被手动修改成 37897)
```yaml
mode: rule
mixed-port: 37897
socks-port: 37898
port: 37899
allow-lan: false
log-level: info
external-controller: 127.0.0.1:9097
secret: ''
bind-address: '*'
tun:
  stack: system
  device: Mihomo
  auto-route: true
  auto-detect-interface: true
  dns-hijack:
  - any:53
  strict-route: false
  mtu: 1500
  enable: false
```

数据包不会被直接发送到 Clash，需要借助 Socks Protocol。在 `curl` 中可以通过 `-x socks5://` 来指定 Socks server 地址(这里为了排除 IPv6 的影响，只使用 IPv4)
```shell
$ curl -4vLsSo /dev/null -x socks5://127.0.0.1:37897 www.google.com
*   Trying 127.0.0.1:37897...
* Connected to 127.0.0.1 (127.0.0.1) port 37897
* Host www.google.com:80 was resolved.
* IPv6: (none)
* IPv4: 142.251.43.4
* SOCKS5 connect to 142.251.43.4:80 (locally resolved)
* SOCKS5 request granted.
* Connected to 127.0.0.1 (127.0.0.1) port 37897
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/8.8.0
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Thu, 18 Jul 2024 07:13:56 GMT
< Expires: -1
< Cache-Control: private, max-age=0
< Content-Type: text/html; charset=ISO-8859-1
< Content-Security-Policy-Report-Only: object-src 'none';base-uri 'self';script-src 'nonce-9r0GuWFO6NqEF2XmRlJ8ww' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
< P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
< Server: gws
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
< Set-Cookie: AEC=AVYB7cozyHJn60uHbJmY1F0UV64qNrrmZ648EfXOpJi8xenQjUy5q1J7-qw; expires=Tue, 14-Jan-2025 07:13:56 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax
< Set-Cookie: NID=515=aeXxtqU5dmdyekcqWbKau8pwnuhEJppDCPjjs3KMj0H57Bgu5o2kL6QLehwP-L9Hmrb-XdYfR3QI7jdl0ERRPtftz9e2s9MVM81GH0dFiIWk0tSCMg1GG2DrDJOWmhXY8pWZTThtqy4RO37aU2FJQjBlRFW5btIMbd62b-3GtBQ; expires=Fri, 17-Jan-2025 07:13:56 GMT; path=/; domain=.google.com; HttpOnly
< Accept-Ranges: none
< Vary: Accept-Encoding
< Transfer-Encoding: chunked
<
{ [6204 bytes data]
* Connection #0 to host 127.0.0.1 left intact
```

从 curl 的结果中可推出 wireshark filter 应该为

`(tcp.port eq 37897 and tcp.stream eq 2) or dns.qry.name eq www.google.com or tcp.port eq 39041 or ip.addr eq 142.251.43.4`

> [!NOTE] 
> 1. 为了对比 Tun enabled 的情况，需要以混杂模式抓包，即 any
> 2. 在 wireshark 中 Socks Protocol 默认以 1080 端口标识。要想识别非标端口，需要使用 Analyze Decode as 功能 
> 3. tcp.stream eq 2 由 wireshark context 推出
> 4. tcp.port eq 39041 为 vmess 代理的入站端口

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240718/2024-07-18_14-38-09.6f0kmk3lka.webp)
#### Intial Socks Over TCP Connection

> frame 25th to frame 31th

建立 Socks over TCP 连接 

这部分逻辑比较简单(如果有不明白的，RFC1928[^4] 是你最好的朋友)，对应 curl 中的表现为
```shell
*   Trying 127.0.0.1:37897...
* Connected to 127.0.0.1 (127.0.0.1) port 37897
```

#### `www.google.com` DNS Resolution

> [!NOTE]
> Clash verge rev 在没有开启 Clash TUN 的情况下，默认不会使用 Clash DNS，这点可以从配置文件中看出

> frame 32th to frame 33th

DNS 解析 `www.google.com` 

应用通常会通过 `getaddrinfo()` 调用系统的 DNS 解析机制来获取 Domain 的解析记录值。假设 `/etc/resolv.conf` 配置如下，那么就会使用 172.18.10.11 作为 nameserver
```sh
cat /etc/resolv.conf
# Generated by NetworkManager
nameserver 172.18.10.11
nameserver 218.108.248.200
nameserver 212.101.172.35
```

从报文中可以得出 `www.google.com` A record 为 142.251.43.4
在 curl 中的表现为
   ```shell
   * Host www.google.com:80 was resolved.
   * IPv6: (none)
   * IPv4: 142.251.43.4
   ```

#### Socks Request/Response and Real Request

> frame 34th to frame 36th

Socks 求情响应以及发送真实请求(如果有不明白的，RFC1928[^4] 是你最好的朋友)

client 告诉 Socks server 想要访问 142.251.43.4:80 (对应 `www.google.com`)，对应 frame 34th
![](https://github.com/dhay3/picx-images-hosting/raw/master/20240718/2024-07-18_16-44-46.361gpyu2vu.webp)
Scoks server 回送，告诉 client 如果想要访问 142.251.43.4:80 需要通过 127.0.0.1:37897 建立 TCP 连接，对应 frame 35th
![](https://github.com/dhay3/picx-images-hosting/raw/master/20240718/2024-07-18_16-52-33.9dcuq4zx48.webp)
发送真实的请求，对应 frame 36th

整个过程在 curl 中的表现为
```shell
* SOCKS5 connect to 142.251.43.4:80 (locally resolved)
* SOCKS5 request granted.
* Connected to 127.0.0.1 (127.0.0.1) port 37897
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/8.8.0
> Accept: */*
```
Clash 在收到这个报文后就会按照 Profile Rules 去匹配，发现 142.251.43.4 没有匹配的规则，那就会走 MATCH rule
```
- 'MATCH,SG01'
```

#### Proxy Request and Close Connection

> [!NOTE] 
> 节点使用 ws + vmess，所以报文出站也必须使用 ws + vmess 加密，因此不能分辨出具体那个报文是请求

> tcp stream 4 frame 41th to frame 100th

向代理发送请求接受响应并挥手关闭连接

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240718/2024-07-18_17-36-07.syu8taejl.webp)
10.100.4.222 对应 Inside NAT(公网 NAT 前的地址)
120.232.63.24 对应节点的地址(推测可能走的是 IPLC)

### 0x04b Tun Enabled

在开启 Clash Tun 的情况下 Clash 全局配置如下(`mixed-port`，`socks-port`，`port` 被手动修改成 37897)
```yaml
mode: rule
mixed-port: 37897
socks-port: 37898
port: 37899
allow-lan: false
log-level: info
external-controller: 127.0.0.1:9097
secret: ''
bind-address: '*'
tun:
  stack: system
  device: Mihomo
  auto-route: true
  auto-detect-interface: true
  dns-hijack:
  - any:53
  strict-route: false
  mtu: 1500
  enable: true
dns:
  enable: true
  enhanced-mode: normal
  fake-ip-range: 198.18.0.1/16
  nameserver:
  - 114.114.114.114
  - 223.5.5.5
  - 8.8.8.8
  fallback: []
```

这里新增了 DNS 部分的配置，是因为在 Clash verge rev 中如果开启 tun 就会使用 Clash DNS 的功能（DNS 的流量也被会 Clash 处理）
具体代码逻辑可以看 [clash-verge-rev/src-tauri/src/enhance/tun.rs at main · clash-verge-rev/clash-verge-rev · GitHub](https://github.com/clash-verge-rev/clash-verge-rev/blob/main/src-tauri/src/enhance/tun.rs) (Clash verge rev 和 mihomo core 的默认配置不同，具体看代码)
```ts
fn use_dns_for_tun(mut config: Mapping) -> Mapping {
	 ...
    // 开启tun将同时开启dns
    revise!(dns_val, "enable", true);

    append!(dns_val, "enhanced-mode", "fake-ip");
    append!(dns_val, "fake-ip-range", "198.18.0.1/16");
    append!(
        dns_val,
        "nameserver",
        vec!["114.114.114.114", "223.5.5.5", "8.8.8.8"]
    );
    append!(dns_val, "fallback", vec![] as Vec<&str>);
    ...
}
```

`enhanced-mode: fake-ip` 是 Clash 中另外一特性，后面单独讲。为了控制变量，这里通过 Clash verge rev 的 Global Extend Config 将 `enhancemod` 置为 normal

在开启 TUN 的情况下， 系统会新增一个 TUN network device 以及 PBR(具体规则看 [0x03a Clash Tun PBR](#0x03a%20Clash%20Tun%20PBR))
```sh
$ ip a s dev Mihomo
7: Mihomo: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc cake state UNKNOWN group default qlen 500
    link/none
    inet 198.18.0.1/30 brd 198.18.0.3 scope global Mihomo
       valid_lft forever preferred_lft forever
    inet6 fe80::32a1:b927:db7:b33b/64 scope link stable-privacy proto kernel_ll
       valid_lft forever preferred_lft forever
```

访问 `www.google.com`，由于使用 PBR 接管了系统的路由，所以无需使用 `-x socks5://` 指定 Socks server 地址，也就没有 Socks 的建立和请求响应了
```sh
$ curl -4vLsSo /dev/null  www.google.com
* Host www.google.com:80 was resolved.
* IPv6: (none)
* IPv4: 31.13.73.169
*   Trying 31.13.73.169:80...
* Connected to www.google.com (31.13.73.169) port 80
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/8.8.0
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Thu, 25 Jul 2024 01:10:52 GMT
< Expires: -1
< Cache-Control: private, max-age=0
< Content-Type: text/html; charset=ISO-8859-1
< Content-Security-Policy-Report-Only: object-src 'none';base-uri 'self';script-src 'nonce-Eqk2rcdRc0i-2XfqaAyt9w' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
< P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
< Server: gws
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
< Set-Cookie: AEC=AVYB7crRQTPB7EefzI8C-293XaQb2RgE2Qjk9mtRSZKRnpQkYF8mUnwRtw; expires=Tue, 21-Jan-2025 01:10:52 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax
< Set-Cookie: NID=516=SSleAG_EVR4NiSn08-GHnza9eZIRqQryHtykMyC5DoyssWz_lKinMWhDBkwthWBa73pHNHS6yrUnDsq4xyc3p3K5whSWXYGslH8J3sBrAHGU3Gc_r4yIlyviv463otpyaH3iZctHi-OtixYtS-dofAFh5TINIsAmaI_KOF-V8oA; expires=Fri, 24-Jan-2025 01:10:52 GMT; path=/; domain=.google.com; HttpOnly
< Accept-Ranges: none
< Vary: Accept-Encoding
< Transfer-Encoding: chunked
<
{ [9093 bytes data]
* Connection #0 to host www.google.com left intact

```

从 curl 的结果中可推出 wireshark filter 应该为

`dns.qry.name eq www.google.com and not icmp or tcp.port eq 39041 or ip.addr eq 31.13.73.169`

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240725/2024-07-25_09-19-29.5c0vlas394.webp)

#### `www.google.com` DNS Resolution

在 [0x04a Tun disabled](#0x04a%20Tun%20disabled) 中我们知道 Clash 在没有启用 Tun 时，会调用系统的 DNS nameserver 来解析域名。这部分解析域名的 DNS nameserver 并不是由 Clash 决定的，而是由发起请求的应用来决定的。例如在 `curl` 中通过封装 `getaddrinfo()` 的函数 `struct Curl_addrinfo *Curl_ipv4_resolve_r(const char *hostname,int port)` 来调用系统的 DNS 解析机制，从 `/etc/resolv.conf` 中决定使用那个 nameserver

假设 `/etc/resolv.conf` 内容如下，那么系统就会使用 172.18.10.11 作为 nameserver
```sh
cat /etc/resolv.conf
# Generated by NetworkManager
nameserver 172.18.10.11
nameserver 218.108.248.200
nameserver 212.101.172.35
```

这时会构建一个 UDP 报文
```
src: 198.18.0.1
sport: random
dst: 172.18.10.11
dport: 53
dns.qry: www.google.com
type: A
class: IN
```

这个 UDP 报文就会通过 Tun network device (这个例子中就是 Mihomo) 发送给 mihomo core
mihomo core 在收到报文后按照 rules 去匹配规则。因为这个是一个私网 IP，机场下发的配置中通常会将其定义为 Direct 策略
```
- IP-CIDR,172.16.0.0/12,DIRECT
```

同时因为配置文件指定了 Clash 使用的 nameserver 如下
```yaml
  nameserver:
  - 114.114.114.114
  - 223.5.5.5
  - 8.8.8.8
```

那么 Clash 在解析配置文件的时候，会生成一个 `NameServer[]` slice
具体代码看 [mihomo/config/config.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go) `func parseNameServer(servers []string, respectRules bool, preferH3 bool) ([]dns.NameServer, error)` 和 `func parsePureDNSServer(server string) string`
```go
[
	{ 
		Net: udp
		Addr: 114.114.114.114:53
		ProxyName: ""
		Params: ""
		PerferH3: preferH3
	},
	{ 
		 Net: udp
		 Addr: 223.5.5.5:53
		 ProxyName: ""
		 Params: ""
		 PerferH3: preferH3
	},
	{ 
		Net: udp
		Addr: 8.8.8.8:53
		ProxyName: ""
		Params: ""
		PerferH3: preferH3
	}
]
```

这个 slice 会通过 `dns.NewResolver(cfg)` 被调用生成一个 resolver，然后赋值给 `mihomo.component.resolve.DefaultResolver` 类变量 
具体代码看 [mihomo/hub/executor/executor.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/hub/executor/executor.go)
```go
import {
	...
	"github.com/metacubex/mihomo/component/resolver"
	...
}

func updateDNS(c *config.DNS, generalIPv6 bool) {
	...
	cfg := dns.Config{
		Main:         c.NameServer,
		Fallback:     c.Fallback,
		IPv6:         c.IPv6 && generalIPv6,
		IPv6Timeout:  c.IPv6Timeout,
		EnhancedMode: c.EnhancedMode,
		Pool:         c.FakeIPRange,
		Hosts:        c.Hosts,
		FallbackFilter: dns.FallbackFilter{
			GeoIP:     c.FallbackFilter.GeoIP,
			GeoIPCode: c.FallbackFilter.GeoIPCode,
			IPCIDR:    c.FallbackFilter.IPCIDR,
			Domain:    c.FallbackFilter.Domain,
			GeoSite:   c.FallbackFilter.GeoSite,
		},
		Default:        c.DefaultNameserver,
		Policy:         c.NameServerPolicy,
		ProxyServer:    c.ProxyServerNameserver,
		Tunnel:         tunnel.Tunnel,
		CacheAlgorithm: c.CacheAlgorithm,
	}
	
	r := dns.NewResolver(cfg)
	...
	resolver.DefaultResolver = r
	resolver.DefaultHostMapper = m
	resolver.DefaultLocalServer = dns.NewLocalServer(r, m)
	resolver.UseSystemHosts = c.UseSystemHosts
	...
}
```

这个生成的 resolver 会使用配置中指定的 nameservers
具体代码看 [mihomo/dns/resolver.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/dns/resolver.go)
```go
func NewResolver(config Config) *Resolver {
		...
		r := &Resolver{
			ipv6:        config.IPv6,
			// 指定 DNS Nameservers
			main:        cacheTransform(config.Main),
			cache:       cache,
			hosts:       config.Hosts,
			ipv6Timeout: time.Duration(config.IPv6Timeout) * time.Millisecond,
		}
		...
	return r
}
```

当 TUN network 收到报文后，就会调用 `func (d *DNSDialer) ListenPacket(ctx context.Context, network, addr string) (net.PacketConn, error)` 对 Domain 做解析
具体代码看 [mihomo/tunnel/dns\_dialer.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/tunnel/dns_dialer.go)
```go
import (
	...
	"github.com/metacubex/mihomo/component/resolver"
	...
)
...
func (d *DNSDialer) ListenPacket(ctx context.Context, network, addr string) (net.PacketConn, error) {
	...
	if !metadata.Resolved() {
		// udp must resolve host first
		dstIP, err := resolver.ResolveIPWithResolver(ctx, metadata.Host, r)
		if err != nil {
			return nil, err
		}
		metadata.DstIP = dstIP
	}
	...
}
```

所以 Clash 除了会对 172.18.10.11 发送 DNS query request，还会向配置中的 nameserver 发送 DNS query request(10.100.4.222 为公网互联 IP address)
```
src: 10.100.4.222
sport: random
dst: {114.114.114.114,223.5.5.5,8.8.8.8}
dport: 53
dns.qry: www.google.com
type: A
class: IN
```

在 Wireshark 中的表现为
![](https://github.com/dhay3/picx-images-hosting/raw/master/20240725/2024-07-25_09-20-56.9kg2v4k8bt.webp)

#### Request Proxy and Response

在获取到解析的记录值后，就发起请求建立 TCP 连接了

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240725/2024-07-25_09-23-46.7p6w0xngv.webp)

从 frame 25th to frame 34th 可以推出如下几点
1. 会使用返回的第一个 DNS qry reponse 的记录值做 TCP 连接(也就不能保证 DNS 解析结果不被污染，所以 Clash 引入一个 [Clash 06 - Fake-ip](Clash%2006%20-%20Fake-ip.md) 的功能来防止 DNS Pollution)，这里为 223.5.5.5 返回的 31.13.73.169
2. 198.18.0.1 和 31.13.73.169 建立的连接其实是假的(SYN-ACK 由 Clash 代答)，从 `ip.ttl == 64` 可以推出，逻辑上也不必和 31.13.73.169 直接建立连接，因为需要按照 rules 分流
3. 会在发出应用层报文后和代理建立 WS TCP 连接，frame 35th 开始和代理建立连接

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Universal TUN/TAP device driver — The Linux Kernel  documentation](https://docs.kernel.org/networking/tuntap.html)
[^2]:[Tunneling protocol - Wikipedia](https://en.wikipedia.org/wiki/Tunneling_protocol)
[^3]:[TUN/TAP - Wikipedia](https://en.wikipedia.org/wiki/TUN/TAP)
[^4]:[RFC 1928:  SOCKS Protocol Version 5](https://www.rfc-editor.org/rfc/rfc1928)
[^5]:[GitHub - google/gvisor: Application Kernel for Containers](https://github.com/google/gvisor)