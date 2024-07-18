---
createTime: 2024-07-16 13:08
tags:
  - "#hash1"
  - "#hash2"
---

# Clash 06 - Tun

## 0x01 Overview

Tun 是 Premium core 中最为核心的功能之一，可以让网络中的设备真正得实现==全局代理==

## 0x02 TUN/TAP

在介绍 Tun 前，需要先了解一下 TUN/TAP

> TUN/TAP provides packet reception and transmission for user space programs. It can be seen as a simple Point-to-Point or Ethernet device, which, instead of receiving packets from physical media, receives them from user space program and instead of sending packets via physical media writes them to the user space program.[^1]

TUN(TUN)/TAP 是系统上虚拟的 network devices。可以将经过虚拟 network devices 的数据包发送到**用户态的进程**，从而实现数据包的二次封装(这一过程也被称为 tunnleing encapsulation[^2]) 

可以参考这一实例
[What Is the TUN Interface Used For? | Baeldung on Linux](https://www.baeldung.com/linux/tun-interface-purpose)

### 0x02a TUN vs TAP[^3]

TUN/TAP 最大的区别就在于

- TUN network device 工作在 OSI 3 层(即网络层)，接收 3 层数据包，不可以处理 Ethernet header
- TAP network device 工作在 OSI 2 层(即数据链路层)，接收 2 层数据包，可以处理 Ethernet header

![](https://upload.wikimedia.org/wikipedia/commons/a/af/Tun-tap-osilayers-diagram.png)

## 0x03 Clash tun

> 如果想要详细了解原理，还需要看源码
> [GitHub - MetaCubeX/mihomo at Meta](https://github.com/MetaCubeX/mihomo/tree/Meta)

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
	到 198.18.0.0/30 (这个地址段通常用于 fake-ip)的 3 层数据包会使用 2022 route table
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

## 0x04 Analyze

> [!NOTE] important
> 下面的例子使用 Clash verge rev 均未开启 fake-ip，mixed-port 为 37897 即 local inbound 监听端口
> 	未在 emulator 中模拟纯净环境，以访问 `www.google.com` 为例

### 0x04a Tun disabled

在没有开启 TUN 的情况下，想要流量通过 Clash 需要借助 Socks Protocol。在 `curl` 中可以通过 `-x socks5://` 来指定入站代理的地址(这里为了排除 IPv6 的影响，只使用 IPv4)
	
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
> 1. 需要以混杂模式抓包，即 any
> 2. 在 wireshark 中 Socks Protocol 默认以 1080 端口标识，要想识别非标端口，需要使用 Analyze Decode as 功能 
> 3. tcp.stream eq 2 由 wireshark context
> 4. tcp.port eq 39041 为 vmess 代理的入站端口

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240718/2024-07-18_14-38-09.6f0kmk3lka.webp)
#### 建立 Socks over TCP 连接

frame 25th to frame 31th
这部分逻辑比较简单，对应 curl 中的表现为
```shell
*   Trying 127.0.0.1:37897...
* Connected to 127.0.0.1 (127.0.0.1) port 37897
```
	
#### DNS 解析 `www.google.com`

frame 32th to frame 33th
在 Clash verge rev 中如果没有开启 tun 就不会使用 Clash nameserver 的功能(具体代码逻辑可以看 [clash-verge-rev/src-tauri/src/enhance/tun.rs at main · clash-verge-rev/clash-verge-rev · GitHub](https://github.com/clash-verge-rev/clash-verge-rev/blob/main/src-tauri/src/enhance/tun.rs)，mihomo core 默认只配置 `mixed-port`)，这时只有 client 侧的 DNS nameserver 会处理 DNS query，这里的 nameserver 为局域网中的 172.18.10.11

这里得出 `www.google.com` A record 为 142.251.43.4
在 curl 中的表现为
```shell
* Host www.google.com:80 was resolved.
* IPv6: (none)
* IPv4: 142.251.43.4
```
	
#### Socks 求情响应以及发送真实请求

frame 34th to frame 36th
> 如果有不明白的，RFC[^4] 是你最好的朋友

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

#### 向代理发送请求并挥手关闭连接

tcp stream 4 frame 41th to frame 100th
> 节点使用 ws + vmess，所以报文出站也必须使用 ws + vmess 加密，因此不能分辨出具体那个报文是请求

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240718/2024-07-18_17-36-07.syu8taejl.webp)
10.100.4.222 对应 Inside NAT(公网 NAT 前的地址)
120.232.63.24 对应节点的地址(推测可能走的是 IPLC)


## 0x04 System proxy vs Clash tun



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Universal TUN/TAP device driver — The Linux Kernel  documentation](https://docs.kernel.org/networking/tuntap.html)
[^2]:[Tunneling protocol - Wikipedia](https://en.wikipedia.org/wiki/Tunneling_protocol)
[^3]:[TUN/TAP - Wikipedia](https://en.wikipedia.org/wiki/TUN/TAP)
[^4]:[RFC 1928:  SOCKS Protocol Version 5](https://www.rfc-editor.org/rfc/rfc1928)