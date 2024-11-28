---
createTime: 2024-11-18 09:28
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 01 - Overview

## 0x01 Preface

> OpenSSH(OpenBSD Secure Shell) is the premier connectivity tool for remote login with the SSH protocol. It encrypts all traffic to eliminate eavesdropping, connection hijacking, and other attacks. In addition, OpenSSH provides a large suite of secure tunneling capabilities, several authentication methods, and sophisticated configuration options.[^1]

OpenSSH 是一套基于 SSH(Secure Shell) 协议的 C/S 加密网络工具，大多数 Unix-like OS 都预装了该工具

## 0x02 SSH

> The Secure Shell (SSH) Protocol is a protocol for secure remote login and other secure network services over an insecure network.[^2]

在没有出现 Secure Shell(SSH) protocol 的时代，大多数管理员都使用 telnet 或者 rlogin/rexec/rsh 这些非安全的传输协议(和 HTTP 一样传输的内容为明文)用于登入 OS 或者执行 OS 命令。而 OS 通常会要求用户进行 authentication，例如输入 account/password。所以中间的网络设备或者流量嗅探都可以看到 L3/L4 的源目信息，以及明文的 account/password，这显然是不安全的

SSH 就是为了解决这个问题而诞生的，他能够加密主机之间的通信(对 L7 报文加密)，保证不被窃听或者篡改。以及以及引入了 PKI authentication 的逻辑

### 0x02a Usages

- For login to a shell on a remote host (replacing [Telnet](https://en.wikipedia.org/wiki/Telnet) and [rlogin](https://en.wikipedia.org/wiki/Rlogin))
- For executing a single command on a remote host (replacing [rsh](https://en.wikipedia.org/wiki/Remote_shell))
- For setting up automatic (passwordless) login to a remote server (for example, using [OpenSSH](https://en.wikipedia.org/wiki/OpenSSH)
- In combination with [rsync](https://en.wikipedia.org/wiki/Rsync) to back up, copy and mirror files efficiently and securely
- For [forwarding](https://en.wikipedia.org/wiki/Port_forwarding) a port
- For [tunneling](https://en.wikipedia.org/wiki/Tunneling_protocol) (not to be confused with a [VPN](https://en.wikipedia.org/wiki/VPN), which [routes](https://en.wikipedia.org/wiki/VPN#Routing) packets between different networks, or [bridges](https://en.wikipedia.org/wiki/VPN#OSI_Layer_1_services) two [broadcast domains](https://en.wikipedia.org/wiki/Broadcast_domain) into one).
- For using as a full-fledged encrypted VPN. Note that only [OpenSSH](https://en.wikipedia.org/wiki/OpenSSH) server and client supports this feature.
- For forwarding [X](https://en.wikipedia.org/wiki/X_Window_System) from a remote [host](https://en.wikipedia.org/wiki/Host_(network)) (possible through multiple intermediate hosts)
- For browsing the web through an encrypted proxy connection with SSH clients that support the [SOCKS protocol](https://en.wikipedia.org/wiki/SOCKS).
- For securely mounting a directory on a remote server as a [filesystem](https://en.wikipedia.org/wiki/File_system) on a local computer using [SSHFS](https://en.wikipedia.org/wiki/SSHFS).
- For automated remote monitoring and management of servers through one or more of the mechanisms discussed above.
- For development on a mobile or embedded device that supports SSH.
- For securing file transfer protocols.

### 0x02a Version

SSH 有 2 个 main version

- SSH-1
	由 Tatu Ylönen 在 1995 年提出，简单
- SSH-2
	在 2006 年由 secsh 提出，使用 [DH](../../OpenSSL/Diffie-Hellman%20algrothim.md) key exchange 替代了 RSA key exchange(后来演变为 EC DH)，防止 weak key exchange 被利用；使用 md5/sha1 校验传输数据的一致性，防止被 MITMA 篡改报文

> [!important]
> SSH-1 和 SSH-2 两者之间互不兼容，且在 OpenSSH Version 7.6 后移除对 SSH1 的支持(*ssh(1): delete SSH protocol version 1 support, associated configuration options and documentation.*)[^3]

### 0x02b Structure

虽然 SSH-1 和 SSH-2 两者互不兼容，但是逻辑上可以被拆成 3 个组件，亦或是阶段(别和 OSI layer 搞混)

- transport layer

	provides algorithm negotiation and a key exchange.The key exchange includes server authentication and results in a cryptographically secured connection: it provides integrity, confidentiality and optional compression.

	transport layer 协商 connection 使用的加密算法以及加密密钥

- user authentication layer

	uses the established connection and relies on the services provided by the transport layer. It provides several mechanisms for user authentication. These include traditional password authentication as well as public-key or host-based authentication mechanisms.

	user authentiation layer 处理鉴权

- connection layer

	multiplexes many different concurrent channels over the authenticated connection and allows tunneling of login sessions and TCP-forwarding. It provides a flow control service for these channels. Additionally, various channel-specific options can be negotiated.

	connection layer 处理用户的 L7 数据

![2023-01-31_17-29](https://github.com/dhay3/image-repo/raw/master/20230131/2023-01-31_17-29.2zk7mni15f7k.webp)

### 0x02c SSH Packet Analyze

在分析报文前需要回顾一下 DH key exchange 的过程

> 以 Alice 与 Bob 之前需要交换密钥为例子
> 
> 1. 首先 Alice 与 Bob 会公开确认使用两个素数 p 和 g
> 
>    p 和 g 两个素数会通过 exchange 宣告给对端
> 
> 2. Alice 选择一个 secret integer a（这里直接直接理解成 Alice 的私钥，只有 Alice 自己知道）, 并计算 A 如下：A = $g^a \bmod p$ 
> 
> 3. Alice 将 A （Alice 的公钥） 发送给 Bob
> 
> 4. Bob 选择一个 secret integer b （Bob 的私钥，只有 Bob 自己知道）, 并计算 B 如下：B = $g^b \bmod p$
> 
> 5. Bob 将 B （Bob 的公钥）发送给 Alice
> 
> 6. Alice 计算密钥 K = $B^a \bmod p$ = $g^{ab} \bmod p$
> 
> 7. Bob 计算密钥 K = $A^b \bmod p$ = $g^{ab} \bmod p$
> 
> 8. 因此 Alice 和 Bob 可以用其进行加对称解密（==但是K1 K2 是通过非对称加密计算出来的==），计算所得的 K 不会通过网络传输所以是相对安全的

以 10.100.64.1 连接 10.0.3.49 SSH password authentication 为例（public authentication 结果可能不尽相同），按照 layer 划分

![](https://github.com/dhay3/picx-images-hosting/raw/master/2024-11-19_10-14-02.6m3xeurojb.png)

#### Transport Layer

##### TCP 3-way handshakes

frame 1th to 3th

SSH 是 over TCP/IP 的协议，需要 three-way handshakes 建立 TCP 连接

> [!note]
> SSH IANA 分配的端口为 22

##### Protocol Negotiation

frame 4th to 6th

client 和 server 分别宣告自己使用的 SSH version 和 OpenSSH version(有些还会携带 OS version)

格式为 `SSH_protoversion-OpenSSH_version [OS_version]`

![](https://github.com/dhay3/picx-images-hosting/raw/master/2024-11-19_10-16-09.9nztg2vw7f.webp)

##### Key Exchange

frame 9th to 13th

标准的 DH key exchange

- Client: Key Exchange Init

	client 宣告自己使用的 key exchange algo 以及 compression algo

	对应 DH key exchange 中的 1

- Server: Key Exchange Init

	server 宣告自己使用的 key exchange 和 compression 算法

	对应 DH key exchange 中的 1

- Client: Elliptic Curve Diffie-Hellman Key Exchange Init

	client 宣告自己的公钥

	![](https://github.com/dhay3/picx-images-hosting/raw/master/2024-11-19_10-19-59.8vmxycja8a.webp)

	对应 DH key exchange 中的 2，3

- Server: Elliptic Curve Diffie-Hellman Key Exchange Init

	server 宣告自己的公钥

	![](https://github.com/dhay3/picx-images-hosting/raw/master/2024-11-19_10-21-05.3rb992knjx.webp)

	对应 DH key exchange 中的  4，5

##### New Keys

frame 14th

仅对 key exchange 做 ackownledgement

#### User Authentication Layer

user authentication layer 会使用 key exchange 阶段对端宣告的公钥加密，所以抓包显示为 Encrypted Packet

#### Connection Layer

connection layer 会使用 key exchange 阶段对端宣告的公钥加密，所以抓包显示为 Encrypted Packet

## 0x03 OpenSSH Suite

OpenSSH 根据不同的功能可以将 tools 划分为 3 类

- Remote operations
	
	ssh, scp, and sftp.
	 
- Key management
	
	ssh-add, ssh-keysign, ssh-keyscan, and ssh-keygen. 
	
- The service side
	
	sshd, sftp-server, and ssh-agent. 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [OpenSSH - Wikipedia](https://en.wikipedia.org/wiki/OpenSSH)
- [OpenSSH - ArchWiki](https://wiki.archlinux.org/title/OpenSSH)
- [OpenSSH](https://www.openssh.com/)
- [Secure Shell - Wikipedia](https://en.wikipedia.org/wiki/Secure_Shell)
- [RFC 4253: The Secure Shell (SSH) Transport Layer Protocol](https://www.rfc-editor.org/rfc/rfc4253.html)
- [RFC 4251 - The Secure Shell (SSH) Protocol Architecture](https://datatracker.ietf.org/doc/rfc4251/)
- [linux - Understand wireshark capture for ssh key exchange - Server Fault](https://serverfault.com/questions/586638/understand-wireshark-capture-for-ssh-key-exchange)

***References***

[^1]:[OpenSSH](https://www.openssh.com/)
[^2]:[RFC 4251 - The Secure Shell (SSH) Protocol Architecture](https://datatracker.ietf.org/doc/html/rfc4251)
[^3]:[release-7.6](https://www.openssh.com/txt/release-7.6)