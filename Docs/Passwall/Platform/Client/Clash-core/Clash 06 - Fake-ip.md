---
createTime: 2024-07-16 12:44
tags:
- "#Passwall"
- "#Clash"
---

# Clash 06 - Fake-ip

## 0x01 Overview

fake-ip 是 Clash 中防止 DNS Pollution 的一种手段。这一的概念来自 [RFC3089](https://tools.ietf.org/rfc/rfc3089)

## 0x01 Originate

在介绍 fake-ip 前，需要先明白为什么需要 fake-ip，以及 fake-ip 是为了解决什么

### Socks-based IPv4/IPv6 Gateway Mechanism

我们都知道 IPv4 的地址可以和 IPv4 的地址互相通信，IPv6 的地址可以和 IPv6 的地址互相通信。现在想要 IPv4 的地址和 IPv6 的地址互相通信，那么就需要借助 Gateway 即

A      IPv4     IPv4       homogeneous
B      IPv4     IPv6       heterogeneous
C      IPv6     IPv4       heterogeneous
D      IPv6     IPv6       homogeneous

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

## 0x02 Clash Fake IP

Clash Fake IP 和 RFC3089 大体上相同

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

