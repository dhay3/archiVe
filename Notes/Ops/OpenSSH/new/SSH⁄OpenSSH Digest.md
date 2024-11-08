# SSH/OpenSSH Digest

ref

https://en.wikipedia.org/wiki/Secure_Shell

https://www.openssh.com/

https://www.openssh.com/manual.html

https://wangdoc.com/ssh/basic.html

https://juejin.cn/post/6844903685047189512

http://walkerdu.com/2019/10/24/ssh/

https://serverfault.com/questions/586638/understand-wireshark-capture-for-ssh-key-exchange

## Digest

> 本系列文章以 OpenSSH SSH2 为基础

Secure Shell protocol ( SSH ) 是一个加密的网络协议，和 HTTP 一样都是以 C/S 模式工作

SSH 支持两种变体协议 SSH1 和 SSH2 两者互不兼容 。SSH 1 在 OpenSSH version 7.6 后被移除 

## Why SSH comes

早的时候的一个典型的例子就是服务器登录。登录远程服务器的时候，需要将用户输入的密码传给服务器，如果这个过程是明文通信，就意味着传递过程中，线路经过的中间计算机都能看到密码，这是很可怕的。SSH 就是为了解决这个问题而诞生的，它能够加密计算机之间的通信，保证不被窃听或篡改。它还能对操作者进行认证（authentication）和授权（authorization）。明文的网络协议可以套用在它里面，从而实现加密。

## SSH structure

SSH 是一个协议簇，为了方便理解将其拆解成 3 层

==注意这里的 layer 和 OSI 模型中的 layer 没有直接关系==

- transport layer

  provides algorithm negotiation and a key exchange.The key exchange includes server authentication and results in a cryptographically secured connection: it provides integrity, confidentiality and optional compression.

  处理加密算法协商和 DH 密钥交换，设置 encryption，compression 以及 integrity verfification

- user authentication layer

  uses the established connection and relies on the services provided by the transport layer. It provides several mechanisms for user authentication. These include traditional password authentication as well as public-key or host-based authentication mechanisms.

  处理 client authentication

- connection layer

  multiplexes many different concurrent channels over the authenticated connection and allows tunneling of login sessions and TCP-forwarding. It provides a flow control service for these channels. Additionally, various channel-specific options can be negotiated.

  没看明白。。

![2023-01-31_17-29](https://github.com/dhay3/image-repo/raw/master/20230131/2023-01-31_17-29.2zk7mni15f7k.webp)

## Uses

SSH 支持如下功能

- For login to a shell on a remote host (replacing [Telnet](https://en.wikipedia.org/wiki/Telnet) and [rlogin](https://en.wikipedia.org/wiki/Rlogin))
- For executing a single command on a remote host (replacing [rsh](https://en.wikipedia.org/wiki/Remote_shell))
- For setting up automatic (passwordless) login to a remote server (for example, using [OpenSSH](https://en.wikipedia.org/wiki/OpenSSH)[[26\]](https://en.wikipedia.org/wiki/Secure_Shell#cite_note-26))
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

## OpenSSH

OpenSSH 是 SSH 的 implementations 之一, 主要用于远程加密登录。由以下几个部分组成

- Remote Operation Components

  `ssh`, `scp`, `sftp`

- Key Management Components

  `ssh-add`, `ssh-keysign`, `ssh-keygen`

- Serve Side Components

  `sshd`, `sftp-server`, `ssh-agent`

## SSH packet analyze

以 30.131.92.34 SSH password authentication 到 43.142.57.183 为例( public authentication 结果可能不尽相同 )

![2023-01-31_16-03](https://github.com/dhay3/image-repo/raw/master/20230131/2023-01-31_16-03.6r9i4iwwp4w0.webp)

### Connection Request/Connection Ackownledged

正常的 TCP three-way handshakes 建立 TCP 连接

### Client Protocol

client 宣告自己使用的 SSH 协议，格式为

`SSH-protoversion-softwareversion`

### Server Protocol

server 宣告自己使用的 SSH 协议以及系统，格式为

`SSH-protoversion-softwareversion OSversion`

==如果 client 和 server SSH 协议不匹配，就会断开 4 层连接( SSH2 和 SSH1 不兼容)==

### Client Key Exchange Init

client 宣告默认使用的 key exchange algo，例如下图，表示默认使用 `curve25519-sha256` 作为 key exchange algo  。以及支持的 key exchange algo, key algo( server 支持的对称加密 algo ), encryption algo, compression algo

![2023-01-31_16-20](https://github.com/dhay3/image-repo/raw/master/20230131/2023-01-31_16-20.5y9vbjwfi6o.webp)

### Server Key Exchange Init

server 宣告( 这里也可以理解成 reply )默认使用的 key exchange algo。以及支持的 key exchange algo,key algo( server 支持的对称加密 ), encryption algo, compression algo

和 client key exchange init 差不多, 这里就不贴图了

### Client algo Key Exchange Init

client 宣告 key exchange ( 一般使用 DH, 这里使用 ECDH ) 的 公钥, encryption algo, compression algo

![2023-01-31_16-54](https://github.com/dhay3/image-repo/raw/master/20230131/2023-01-31_16-54.1tb8l15d3obk.webp)

### Server algo Key Exchange Reply & New Keys

这里也可以拆成 2 阶段

server 宣告 key exhcnage ( 一般使用 DH, 这里使用 ECDH ) 的 公钥

![2023-01-31_16-57](https://github.com/dhay3/image-repo/raw/master/20230131/2023-01-31_16-57.6l8ooqke5wg0.webp)

### (Server New Keys)Client New Keys

校验 DH 算法中生成的对称密钥, 对比参考 TLS Encrypted handshake message 