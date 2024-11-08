---
createTime: 2024-08-05 14:25
tags:
  - "#hash1"
  - "#hash2"
---

# Shadowsocks 01 - Overview

## 0x01 Overview

Shadowsocks 是一个基于 Socks5 的加密(shadow)协议(报文结构和 socks 一样)[^1]，流量会按照如下拓扑发送

```mermaid
flowchart LR
client <--> ss-local <--encrypted--> ss-remote <--> target
```

> The Shadowsocks local component (ss-local) acts like a traditional SOCKS5 server and provides proxy service to clients. It encrypts and forwards data streams and packets from the client to the Shadowsocks remote component (ss-remote), which decrypts and forwards to the target. Replies from target are similarly encrypted and relayed by ss-remote back to ss-local, which decrypts and eventually returns to the original client.

逻辑过程如下

1. ss-local 即为 socks server，为 client 提供代理服务，同时他会加密数据包并转发到 ss-remote
2. 然后由 ss-remote 解密数据包并转发到 target
3. target 回送数据包，由 ss-remote 加密并转发到 ss-local
4. ss-local 解密数据包并转发到 client

## 0x02 Simple Example

推荐使用 shadowsocks-libev[^2] 用作 shadowsocks 的实现，使用非常简单

假设 

Client Ubuntu 16.10 IP 192.168.56.3/24 国内虚拟机
Server Ubuntu 20.04 IP 23.94.117.197/24 Ashburn VPS

> [!NOTE] 
> 尽量使用同一 LTS 版本，不同 LTS 版本上 shadowsocks-libev 支持的参数值可能不同

现在需要通过 shadowsocks 访问 `ipinfo.io`

1. client/server 安装 shadowsocks-libev(以 Ubuntu 为例)

	```sh	
	#Client/Server
	sudo apt update
	sudo apt install shadowsocks-libev
	```

2. 配置 ss-local/ss-server
	
	安装 shadowscoks-libev 后，默认会自动生成一个配置文件 `/etc/shadowscoks-libev/config.json`
	
	修改成自己想要的配置，具体可用参数可以参考 [Config Format | Shadowsocks](https://shadowsocks.org/doc/configs.html)
	
	```json
	{
	    "server":["23.94.117.197"],
	    "mode":"tcp_and_udp",
	    "server_port":18388,
	    "local_port":11080,
	    "password":"ZazlSjJWw76E",
	    "timeout":60,
	    "method":"aes-128-cfb"
	}
	```
	
	那么会按照 `ss://method:password@server:server_port` 生成如下 ss 地址，这样的地址就可以被用在类似 SSR 中了(也可以使用 Base64 编码或者是 QRcode)
	
	```
	ss://chacha20-ietf-poly1305:ZazlSjJWw76Ed@23.94.117.197:18388
	```

3. 开启 client/server shadowsocks-libev
	
	shadowsocks-libev 默认会以 ss-server 运行,修改 client shadowscoks-libev.server
	```sh
	$ systemctl cat shadowsocks-libev | grep -Pv '#|^$'
	[Unit]
	Description=Shadowsocks-libev Default Server Service
	Documentation=man:shadowsocks-libev(8)
	After=network.target
	[Service]
	Type=simple
	EnvironmentFile=/etc/default/shadowsocks-libev
	User=root
	LimitNOFILE=32768
	ExecStart=/usr/bin/ss-local -a $USER -c $CONFFILE $DAEMON_ARGS
	[Install]
	WantedBy=multi-user.target
	
	$ systemctl daemon-load
	```
	
	shadowscoks-libev 会按需从 `/etc/shadowscoks-libev/config.json` 中读取参数
	
	```sh
	systemctl start shadowscoks-libev
	```

4. 验证 ss 
	
	关闭 client 上所有的代理服务
	
	```sh
	$ curl -x socks5://127.0.0.1:11080 -sSv ipinfo.io
	* Rebuilt URL to: ipinfo.io/
	*   Trying 127.0.0.1...
	* 34
	* 117
	* 59
	* 81
	* Connected to 127.0.0.1 (127.0.0.1) port 11080 (#0)
	> GET / HTTP/1.1
	> Host: ipinfo.io
	> User-Agent: curl/7.50.1
	> Accept: */*
	>
	< HTTP/1.1 200 OK
	< access-control-allow-origin: *
	< Content-Length: 302
	< content-type: application/json; charset=utf-8
	< date: Mon, 05 Aug 2024 08:21:30 GMT
	< referrer-policy: strict-origin-when-cross-origin
	< x-content-type-options: nosniff
	< x-frame-options: SAMEORIGIN
	< x-xss-protection: 1; mode=block
	< via: 1.1 google
	< strict-transport-security: max-age=2592000; includeSubDomains
	<
	{
	  "ip": "23.94.117.188",
	  "hostname": "23-94-117-188-host.colocrossing.com",
	  "city": "Ashburn",
	  "region": "Virginia",
	  "country": "US",
	  "loc": "39.0437,-77.4875",
	  "org": "AS36352 HostPapa",
	  "postal": "20147",
	  "timezone": "America/New_York",
	  "readme": "https://ipinfo.io/missingauth"
	* Connection #0 to host ipinfo.io left intact
	}
	```

## 0x03 Analyze


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[What is Shadowsocks? | Shadowsocks](https://shadowsocks.org/doc/what-is-shadowsocks.html)
[^2]:[GitHub - shadowsocks/shadowsocks-libev: Bug-fix-only libev port of shadowsocks. Future development moved to shadowsocks-rust](https://github.com/shadowsocks/shadowsocks-libev?tab=readme-ov-file#usage)