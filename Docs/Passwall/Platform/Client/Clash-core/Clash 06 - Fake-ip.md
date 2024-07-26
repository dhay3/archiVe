---
createTime: 2024-07-16 12:44
tags:
- "#Passwall"
- "#Clash"
---

# Clash 06 - Fake-ip

## 0x01 Overview

> [!NOTE]
> Prequestion: [Clash 05 - Tun](Clash%2005%20-%20Tun.md)

fake-ip 是 Clash 中防止 DNS Pollution 的一种手段。这一的概念来自 [RFC3089](https://tools.ietf.org/rfc/rfc3089)

## 0x01 Originate

在介绍 fake-ip 前，需要先明白为什么需要 fake-ip，以及 fake-ip 是为了解决什么

### Socks-based IPv4/IPv6 Gateway Mechanism

我们都知道 IPv4 的地址可以和 IPv4 的地址互相通信，IPv6 的地址可以和 IPv6 的地址互相通信。现在想要 IPv4 的地址和 IPv6 的地址互相通信，那么就需要借助 Gateway 即

- A      IPv4     IPv4       homogeneous
- B      IPv4     IPv6       heterogeneous
- C      IPv6     IPv4       heterogeneous
- D      IPv6     IPv6       homogeneous

Gateway 就是为了解决 heterogeneous 的通信问题

假设 Client C IPv X 想要访问 Destination D IPvY

```
                  Client C       Gateway G     Destination D
               +-----------+     (Server)
               |Application|
           +-->+===========+  +-------------+  +-----------+
      same-+   |*SOCKS Lib*|  |  *Gateway*  |  |Application|
       API +-->+===========+  +=====---=====+  +-----------+
               | Socket DNS|  | Socket  DNS |  | Socket DNS|
               +-----------+  +-------------+  +-----------+
               | [ IPv X ] |  |[IPvX]|(IPvY)|  | ( IPv Y ) |
               +-----------+  +-------------+  +-----------+
               |Network I/F|  | Network I/F |  |Network I/F|
               +-----+-----+  +---+-----+---+  +-----+-----+
                     |            |     |            |
                     +============+     +------------+
                       socksified           normal
                       connection         connection
                      (ctrl)+data          data only
```

1. Client 会先和 Gateway 建立 IPv X Socks 连接。通告 Client 想要访问 Destination ，访问 Gateway IPv X 即可
3. Client 通过 Socks 连接向 Gateway 发送请求，Gateway 在收到 Client 发过来的请求就会转发到 Destination，并将 Destination 回送的响应转发到 Client

这一机制也被称为 **Socks-based IPv4/IPv6 Gateway Mechanism**


### DNS Name Resolving Procedure

在网络通信的过程中，我们必须要先获取 IP 地址，才能通信或者转发数据包。如果数据包中是一个域名，就会触发 DNS 解析的机制(不考虑 PTR)。
现在我们要让 Client C 访问 Destination D。如果 DNS 解析是在 Client C 上发生的，因为在没有 IPv Y 的情况下，DNS Nameserver 就不会返回 IPv Y 的记录值。那么 Client C 就不能和 Destination D 通过 Gateway G 建立连接，同样的如果 Destination 想要访问 Client C，DNS Nameserver 也不会返回 IPv X 的记录值。所有显然 DNS 解析的过程需要发生在 Gateway G 上（必须是 Dual Stack）

但是 DNS 的逻辑中，系统默认只会使用返回的第一条 DNS 记录。而系统通常也会有一个 Local DNS Nameserver，这个 Local DNS 和 Gateway G 同时做 DNS 解析。通过网络传输，系统有可能会使用 Local DNS Nameserver 的记录值，也可能会使用 Gateway G 返回的记录值。所以引入一个 **Fake IP** 的逻辑

完整的过程如下

The detailed internal procedure of the "DNS name resolving delegation" and address mapping management related issues are described as follows.

1. An application on the source node (Client C) tries to get the IP address information of the destination node (Destination D) by calling the DNS name resolving function (e.g., gethostbyname()). At this time, the logical host name ("FQDN") information of the Destination D is passed to the application's *Socks Lib* as an argument of called APIs.
2. Since the *Socks Lib* has replaced such DNS name resolving APIs, the real DNS name resolving APIs is not called here.  The argued "FQDN" information is merely registered into a mapping table in *Socks Lib*, and a "fake IP" address is selected as information that is replied to the application from a reserved special IP address space that is never used in real communications (e.g. 0.0.0.x).  The address family type of the "fake IP" address must be suitable for requests called by the applications.  Namely, it must belong to the same address family of the Client C, even if the address family of the Destination D is different from it.  After the selected "fake IP" address is registered into the mapping table as a pair with the "FQDN", it is replied to the application.
3. The application receives the "fake IP" address, and prepares a "socket".  The "fake IP" address information is used as an element of the "socket".  The application calls socket APIs (e.g. connect()) to start a communication.  The "socket" is used as an argument of the APIs.
4. Since the *Socks Lib* has replaced such socket APIs, the real socket function is not called.  The IP address information of the argued socket is checked.  If the address belongs to the special address space for the fake address, the matched registered "FQDN" information of the "fake IP" address is obtained from the mapping table.
5. The "FQDN" information is transferred to the *Gateway* on the relay server (Gateway G) by using the SOCKS command that is matched to the called socket APIs.  (e.g., for connect(), the CONNECT command is used.)
6. Finally, the real DNS name resolving API (e.g., getaddrinfo()) is called at the *Gateway*.  At this time, the received "FQDN" information via the SOCKS protocol is used as an argument of the called APIs.
7. The *Gateway* obtains the "real IP" address from a DNS server, and creates a "socket".  The "real IP" address information is used as an element of the "socket".
8. The *Gateway* calls socket APIs (e.g., connect()) to communicate with the Destination D.  The "socket" is used as an argument of the APIs.

### Why Clash Needs Fake IP

那么为什么 Clash 需要 Fake IP 呢？ 可以参考一下 [关于 Clash 科学上网的最佳实践](https://www.pupboss.com/post/2024/clash-tun-fake-ip-best-practice/#topic-2) 这篇博文

假设在没有开启 Fake IP 的情况下，会同时向 Local DNS Nameservers 以及 Clash DNS NameServers 发送 DNS qry requests。系统会使用返回的第一条 DNS qry reponse，即使 Clash DNS Nameservers 返回的结果没有被 GFW DNS 污染，但是也不能保证 Local DNS Nameservers 返回的结果没有被 GFW DNS 污染。且通常 Local DNS Nameservers 响应的时间会比 Clash DNS Nameservers 的快。所以这个返回的 DNS qry reponse 的中地址，可能根本无法使用。

例如使用本地 DNS 解析 `www.google.com`

```sh
$ dig +nocookie www.google.com

; <<>> DiG 9.18.27 <<>> +nocookie www.google.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 51120
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4000
;; QUESTION SECTION:
;www.google.com.                        IN      A

;; ANSWER SECTION:
www.google.com.         187     IN      A       31.13.94.49

;; Query time: 0 msec
;; SERVER: 172.18.10.11#53(172.18.10.11) (UDP)
;; WHEN: Fri Jul 26 10:08:16 CST 2024
;; MSG SIZE  rcvd: 59

```

这里可以看到使用 Local DNS Nameserver 解析的记录 31.13.94.49 其实是 Facebook 的。你当然不能用 `HOST: www.google.com` 去访问人家 Facebook 的服务器啦

```sh
$ curl ipinfo.io/31.13.94.49
{
  "ip": "31.13.94.49",
  "hostname": "edge-z-p3-shv-01-eze1.facebook.com",
  "city": "Buenos Aires",
  "region": "Buenos Aires F.D.",
  "country": "AR",
  "loc": "-34.4696,-58.6713",
  "org": "AS32934 Facebook, Inc.",
  "postal": "1612",
  "timezone": "America/Argentina/Buenos_Aires",
  "readme": "https://ipinfo.io/missingauth"
}% 
```

在 United State Ashburn 的服务器上解析

```sh
dig www.google.com

; <<>> DiG 9.16.48-Ubuntu <<>> www.google.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 1592
;; flags: qr rd ra; QUERY: 1, ANSWER: 6, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 512
;; QUESTION SECTION:
;www.google.com.                        IN      A

;; ANSWER SECTION:
www.google.com.         195     IN      A       172.253.63.147
www.google.com.         195     IN      A       172.253.63.105
www.google.com.         195     IN      A       172.253.63.103
www.google.com.         195     IN      A       172.253.63.106
www.google.com.         195     IN      A       172.253.63.104
www.google.com.         195     IN      A       172.253.63.99

;; Query time: 3 msec
;; SERVER: 8.8.8.8#53(8.8.8.8)
;; WHEN: Fri Jul 26 03:16:57 BST 2024
;; MSG SIZE  rcvd: 139
```

这里可以发现 172.253.63.68/25 这个段才是 google 的

```sh
curl ipinfo.io/172.253.63.147
{
  "ip": "172.253.63.147",
  "hostname": "bi-in-f147.1e100.net",
  "city": "Reston",
  "region": "Virginia",
  "country": "US",
  "loc": "38.9687,-77.3411",
  "org": "AS15169 Google LLC",
  "postal": "20190",
  "timezone": "America/New_York",
  "readme": "https://ipinfo.io/missingauth"
}% 
```

所以 DNS 解析的结果不一定能用，而 Fake IP 就是解决这个问题的一种方案

## 0x02 Clash Fake IP

### Clash Tun Disabled

Clash Fake IP 和 RFC3089 逻辑上大体相同。开启 Fake IP 非常简单，只需要将 `dns.enhanced-mode` 置为 `fake-ip` 即可。但是想要完全启用 Fake IP 还需要开启 Clash tun

例如 mihomo core 配置如下

```yaml
mode: rule
mixed-port: 37897
socks-port: 37898
port: 37899
allow-lan: false
log-level: info
external-controller: 127.0.0.1:9097
secret: ''
dns:
  enable: true
  enhanced-mode: fake-ip
  fake-ip-range: 198.40.0.1/16
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
bind-address: '*'
```

使用 curl 通过 socks 访问 `www.google.com`

```sh
$ curl -vLsSo /dev/null -x socks5://127.0.0.1:37897 www.google.com
*   Trying 127.0.0.1:37897...
* Connected to 127.0.0.1 (127.0.0.1) port 37897
* Host www.google.com:80 was resolved.
* IPv6: (none)
* IPv4: 31.13.73.169
* SOCKS5 connect to 31.13.73.169:80 (locally resolved)
* SOCKS5 request granted.
* Connected to 127.0.0.1 (127.0.0.1) port 37897
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/8.8.0
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Thu, 25 Jul 2024 03:31:38 GMT
< Expires: -1
< Cache-Control: private, max-age=0
< Content-Type: text/html; charset=ISO-8859-1
< Content-Security-Policy-Report-Only: object-src 'none';base-uri 'self';script-src 'nonce-m4Y0bdd9BwL2AX7CIw6q3Q' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
< P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
< Server: gws
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
< Set-Cookie: AEC=AVYB7cqOXO_mrTGWc_3pUald4f1gT5b1KCq2oYmIMZVNC0y2obyr2INuxA; expires=Tue, 21-Jan-2025 03:31:38 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax
< Set-Cookie: NID=516=Bf95C5aBH20kbKfkftjNd2yt7PfS62HvDJe6NE1wcs-WpTS2mPTM1A24ahiQgBLgIoMKTAMon5GH_i_OZAEezmgB6SLBepMfVLI4dtgNGaIHuELab7JIitre3IEj6rYhu30NWbOReruW7fsx9zopZ6JDmL7xsgKUHSetgc5vVxA; expires=Fri, 24-Jan-2025 03:31:38 GMT; path=/; domain=.google.com; HttpOnly
< Accept-Ranges: none
< Vary: Accept-Encoding
< Transfer-Encoding: chunked
<
{ [2581 bytes data]
* Connection #0 to host 127.0.0.1 left intact
```

这里可以发现 curl 并没有使用 `dns.fake-ip-range` 定义的 198.40.0.1/16 中的任一地址，作为解析记录值，即使 mihomo core 指定 `mixed-port` 监听了 37897

```
$ sudo ss -lnap sport 37897
Netid            State              Recv-Q             Send-Q                         Local Address:Port                          Peer Address:Port             Process
udp              UNCONN             0                  0                                  127.0.0.1:37897                              0.0.0.0:*                 users:(("verge-mihomo",pid=2849,fd=11))
tcp              LISTEN             0                  8192                               127.0.0.1:37897                              0.0.0.0:*                 users:(("verge-mihomo",pid=2849,fd=10))
tcp              ESTAB              0                  0                                  127.0.0.1:37897                            127.0.0.1:57662             users:(("verge-mihomo",pid=2849,fd=39))
tcp              ESTAB              0                  0                                  127.0.0.1:37897                            127.0.0.1:36478             users:(("verge-mihomo",pid=2849,fd=24))
```

那是因为 `-x socks://` 逻辑会让 DNS 解析发生在 Socks Command Request 连接之前，这点可以在 curl 输出的结果中得出

```
*   Trying 127.0.0.1:37897...
* Connected to 127.0.0.1 (127.0.0.1) port 37897 -----> 建立 Socks 连接
* Host www.google.com:80 was resolved.
* IPv6: (none)
* IPv4: 31.13.73.169 -----> DNS 解析
* SOCKS5 connect to 31.13.73.169:80 (locally resolved)
* SOCKS5 request granted. -----> Socks command request
```

在 wireshark 中的表现为

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240725/2024-07-25_12-08-54.26ldmizhv6.webp)

如果想要 curl 使用 Socks Server 做 DNS 解析，要使用 `--sock5s-hostname`

```sh
 curl -vLsSo /dev/null --socks5-hostname 127.0.0.1:37897 www.google.com
*   Trying 127.0.0.1:37897...
* Connected to 127.0.0.1 (127.0.0.1) port 37897
* SOCKS5 connect to www.google.com:80 (remotely resolved)
* SOCKS5 request granted.
* Connected to 127.0.0.1 (127.0.0.1) port 37897
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/8.8.0
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Thu, 25 Jul 2024 04:02:22 GMT
< Expires: -1
< Cache-Control: private, max-age=0
< Content-Type: text/html; charset=ISO-8859-1
< Content-Security-Policy-Report-Only: object-src 'none';base-uri 'self';script-src 'nonce-Xj2dX_jwI50IueuOVwxDVg' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
< P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
< Server: gws
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
< Set-Cookie: AEC=AVYB7coT6qizYH6pgI4ACSj07kAU734zWhwiOjc2eGH5hVuHgL44EVzaTB0; expires=Tue, 21-Jan-2025 04:02:22 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax
< Set-Cookie: NID=516=qXtUzrvNULYA3AU7cNdTWBD6wATpOxkHAJ-qUnO1Ri3VM_gPS08lgOg4KXDZtN6tdgT4bQ3VNYwR2P0v-k0cPX4GApknhEJC-7zMIM9FGTY2sms_tLbhRtQMjGZ2YF4AErKc2lleV1AT6iV5UqDnLKyD6k6za0k352uBobSwZL4; expires=Fri, 24-Jan-2025 04:02:22 GMT; path=/; domain=.google.com; HttpOnly
< Accept-Ranges: none
< Vary: Accept-Encoding
< Transfer-Encoding: chunked
<
{ [11036 bytes data]
* Connection #0 to host 127.0.0.1 left intact

```

在 wireshark 中的表现为 

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240726/2024-07-26_09-05-42.839xusmunu.webp)
Clash 在收到 frame 39th 和 frame 41th 的数据包后，会将这两个数据包直接转发到代理服务器。由代理服务器完成 DNS 解析以及 HTTP 请求，并将 HTTP 响应回传

所以可以得出要想完全开启 Fake IP 的功能就需要，将 `enhanced-mode` 置为 `fake-ip` 且开启 Clash Tun

### Clash Tun Enabled

将 `enhanced-mode` 置为 `fake-ip` 且开启 Clash Tun，mihomo core 配置如下

```yaml
mode: rule
mixed-port: 37897
socks-port: 37898
port: 37899
allow-lan: false
log-level: info
external-controller: 127.0.0.1:9097
secret: ''
dns:
  enable: true
  enhanced-mode: fake-ip
  fake-ip-range: 198.40.0.1/16
  nameserver:
  - 114.114.114.114
  - 223.5.5.5
  - 8.8.8.8
  fallback: []
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
bind-address: '*'
```

使用 curl 访问 `www.google.com`

```sh
$ curl -vLsSo /dev/null  www.google.com
* Host www.google.com:80 was resolved.
* IPv6: (none)
* IPv4: 198.40.0.12
*   Trying 198.40.0.12:80...
* Connected to www.google.com (198.40.0.12) port 80
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/8.8.0
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Fri, 26 Jul 2024 01:36:50 GMT
< Expires: -1
< Cache-Control: private, max-age=0
< Content-Type: text/html; charset=ISO-8859-1
< Content-Security-Policy-Report-Only: object-src 'none';base-uri 'self';script-src 'nonce-SYfXV-dHFRN5YJUt38PsIQ' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
< P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
< Server: gws
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
< Set-Cookie: AEC=AVYB7cr6ICExwnkEWULeN2YoYFjIAd9dglrTFJ-86vMElpkjG4llnDFjQ_A; expires=Wed, 22-Jan-2025 01:36:50 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax
< Set-Cookie: NID=516=BCDVVzCyc2mr9Vo4T02H-6dL7EEbBhrPEb4D4C7g4yxDscr1TeqSlC_Ei99zvRYuVz0Oy74RmDMAb0YbENPYy8rhdY34KyTkIVBYZ8XA-GdcXpbQgPJQFdK4Jvn8yp5Ucz6aDu-uA9N3p6ygyRmRA5p7c3Em5rK84fGPjsdLgDXeaSsK-DmZs50oBQ; expires=Sat, 25-Jan-2025 01:36:50 GMT; path=/; domain=.google.com; HttpOnly
< Accept-Ranges: none
< Vary: Accept-Encoding
< Transfer-Encoding: chunked
<
{ [4981 bytes data]
* Connection #0 to host www.google.com left intact
```

这里看到 DNS 解析会返回一个 198.40.0.12 Fake IP，然后和 198.40.0.12 建立 TCP 连接

在 wireshark 中的表现为

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240726/2024-07-26_09-48-05.99t93fjpb8.webp)



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**