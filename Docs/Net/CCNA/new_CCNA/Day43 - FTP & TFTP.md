# Day43 - FTP & TFTP

## Overview

FTP(File Transfer Protocol) 和 TFTP(Trivial File Transfer Protocol) 都是 IEEE 标准的用于传输文件的协议

两者都使用 CS 模式

- Clients can use FTP or TFTP to copy files from a server
- Clients can use FTP or TFTP to copy files to a server

两种协议在网络设备中最常见的使用场景就是，设备的系统升级

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_15-33.5wvrk1qeoo74.webp)



## TFTP

Trivial FTP 是在 FTP 之后才出现的协议(但不是 FTP 的替代协议)，对比 FTP 配置更加简洁，只支持一个功能

*Only allows a client to copy a file to or from a server*

- No authentication(username/password), so servers will respond to all TFTP requests
- No encryption, so all data is sent in plain text

因为不需要认证，报文也不加密，所以 TFTP 一般只用在可控的网络环境中

- TFTP servers listen on **UDP port 69**
- UDP is connectinless and doesn’t provide reliability with retransmissions. However, TFTP has similar built-in features within the protocol itself

### TFTP Reliability

TFTP 通过应用层类似 TCP acknowledged 的机制来保证报文的可靠性

- If the client is transferring a file to the server, the server will send ACK messages
- If the server is transferring a file to the client, the client will send ACK messages

通过 Timer 来实现重传的机制

- Timers are used, and if an expected message isn’t received in time, the waiting device will resend its previous message

例如左边的是 TFTP server，右边的是 TFTP client

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_16-30.40n36s2bt3y8.webp)

### TFTP Connections

TFTP 传输文件需要有 3 个过程

1. Connection

   TFTP client sends a request to the server, and the server respond back

2. Data Transfer

   The client and server exchange TFTP messages. One sends data and the other sends acknowledgements

3. Connection Termination

   After the the last data message has been sent, a final acknowledgment is sent to terminate the connection

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_16-40.69mlvm74q1kw.webp)

### TFTP TID

- When the client sends the first message to the server, the destination port is UDP 69 and the source is a random ephemeral port
- This random port is called a ‘Transfer Identifier’(TID) and identifies the data transfer
- The server then also selects a random TID to use as the source port when it replies, not 69
- When the client sends the next message, the destination port will be the server’s TID, not 69

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_16-46.1cio3swcaf8g.webp)

## FTP

FTP 是一个非常古早的协议，出现在 1971 年

- FTP uses TCP ports 20 and 21

- Usernames and passwords are used for authentication, however there is no encryption

  FTP 默认以明文的方式传输报文，为了更加安全可以使用 FTPS(FTP over SSL/TLS) 或者是 SFTP（SSH FTP）

- FTP is more complex than TFTP and allows not only file transfers, but clients can also navigate file directories, add and remove directories, list files, etc

### FTP connection

FTP 需要使用 2 个不同端口，因为 FTP 需要 2 种不同的 connections

- An **FTP control connection(TCP 21)** is established and used to send FTP commands and replies
- When files or data are to be transferred, sperate **FTP data(TCP 20)  connections** are established and terminated as need

例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_16-58.2axjkmy4ln28.webp)

1. Client 和 Server 需要先建立 TCP connection
2. FTP 认证以及 Client 中输入的 FTP 指令会通过 Control Connection(21)

*The default method of establishing FTP data connections is **active mode**, which the server initiates the TCP connection*

> 由 server 侧主动发起建立 Data connection 的连接

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-03.35tw3ghpt5xc.webp)

假设现在 client 侧添加了一个防火墙

*In FTP **passive mode**, the client initiates the data connection. This is often necessary when the client is behind a firewall, which could block the incoming connection from the server*

> 由 client 侧主动发起建立 Data connection 的连接

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-06.55ad1bkwmfsw.webp)

## FTP vs TFTP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-13.5llek9gy33eo.webp)

## IOS File System

可以使用 `(R1)#show file systems` 来查看 IOS 的文件系统

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-14.7i221w05i0e8.webp)

> 已经从 CCNA 的考试范围内移除了,

## Upgrading Cisco IOS

现在需要对 R1 做系统升级，对应的系统以及下载到了 SRV1 上

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-20.5dfs96l7x0g0.webp)

首先先看一下 R1 上对应的 IOS 版本，可以使用 `R1#show version`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-20_1.7gk3jyrb1zeo.webp)

> 注意这里的 K9 就像之前 SSH 中的表示支持 crypto

你也可以使用 `R1#show flash` 来查看闪存中的文件

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-23.50i1hnp41tg.webp)

这里可以看到 R1 闪存中有对应的 IOS 镜像，现在就是需要从 File Transfer Server 上将镜像拷贝下来，替换 R1 正在使用的 IOS 镜像，然后重启 R1 让 R1 使用新的 IOS 镜像加载系统 

### TFTP

如果使用 TFTP 在 R1 上拉取 SRV1 上的 IOS 镜像按照如下方式操作

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-33.4cx6wrb6fibk.webp)

- `R1#copy <source> <destination>`

  例如 `copy tftp: flash:` 就是告诉 router 需要从 tftp server 上拷贝文件到 flash 中

- `Address or name of remote host []?`

  TFTP server 的地址

- `Source filename[]?`

  需要拷贝的文件名

  > TFTP 并不会补全或者知道你服务器上到底有什么文件，所以需要提前复制好被拷贝的文件名

- `Destination filename`

  拷贝后新文件的文件名

  > 你可以直接输入<kbd>Enter</kbd>，默认使用原文件名(Source filename)

当你从 TFTP Server 上拷贝完文件后，可以使用 `R1#show flash` 来校验

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-34.759qk6hrewsg.webp)

现在如何让 R1 使用新的这个 IOS 镜像文件来加载系统呢？

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-38.23afo457fwio.webp)

- `R1(config)#boot system flash:<IOS filename>`

  指定使用 filename IOS

  > 如果你不使用这个命令，默认会使用 flash 默认排序的一个镜像，即为 151-4.M4.bin

- `R1#wirte memory`

  写配置，也可以使用 `write`

- `R1#reload`

  重启 router

在上述配置完成后，可以使用 `show version` 来校验是否生效

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-41.18hc4eia5mu8.webp)

当然在配置完成后，还可以删除老的 IOS 镜像文件，可以使用 `R1#delete flash:<filename>`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-42.76fz8lab51mo.webp)

### FTP

如果使用 FTP 在 R1 上拉取 SRV1 上的 IOS 镜像，按照如下方式

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-48.6j3zd8bm9rwg.webp)

- `R1(config)#ip ftp username cisco`

  `R1(config)#ip ftp password cisco`

  指定 FTP server 认证的用户名和密码

- `copy <source> <destination>`

  例如 `copy ftp: flash:` 就是告诉 router 需要从 ftp server 上拷贝文件到 flash 中

- `Address or name of remote host []?`

  `Source filename[]?`

  `Destination filename`

  和 TFTP 中含义一样

接下来的配置就和 TFTP 中的完全一样

## Command Summary

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-49.1ztacmr2b49s.webp)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_17-54.p7urr6bc8qo.webp)

### 0x01

Configure the appropriate IP addresses on each device. Configure routing on the routers to allow full connectivity

```
R2(config-if)#ip add 192.168.12.2 255.255.255.252
R2(config-if)#no shutdown
R2(config)#router ospf 1
R2(config-router)#network 192.168.12.0 0.0.0.3 area 0

R1(config-if)#ip add 192.168.12.1 255.255.255.252
R1(config-if)#no shutdown
R1(config-if)#int g0/1
R1(config-if)#ip add 10.0.0.254 255.255.255.0
R1(config-if)#no shutdown
R1(config)#router ospf 1
R1(config-router)#network 192.168.12.0 0.0.0.3 area 0
R1(config-router)#network 10.0.0.0 0.0.0.255 area 0
R1(config-router)#passive-interface g0/1
```

> 这里 SRV1 还没有手动配置过 IP，所以需要为 SRV1 配置

在 R2 测试

```
R2(config-router)#do ping 10.0.0.1

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 10.0.0.1, timeout is 2 seconds:
.!!!!
Success rate is 80 percent (4/5), round-trip min/avg/max = 0/0/0 ms
```

### 0x02

Use TFTP on R1 to retrieve the following file from SRV1:c2900-universalk9-mz.SPA.155-3.M4a.bin

```
R1#copy tftp: flash: 
Address or name of remote host []? 10.0.0.1
Source filename []? c2900-universalk9-mz.SPA.155-3.M4a.bin
Destination filename [c2900-universalk9-mz.SPA.155-3.M4a.bin]? 

Accessing tftp://10.0.0.1/c2900-universalk9-mz.SPA.155-3.M4a.bin...
Loading c2900-universalk9-mz.SPA.155-3.M4a.bin from 10.0.0.1: !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
[OK - 33591768 bytes]

33591768 bytes copied in 0.695 secs (5074815 bytes/sec)
```

使用 `show flash` 来校验

```
R1#show flash

System flash directory:
File  Length   Name/status
  3   33591768 c2900-universalk9-mz.SPA.151-4.M4.bin
  4   33591768 c2900-universalk9-mz.SPA.155-3.M4a.bin
```

### 0x03

Upgrade R1’s OS and then delete the old file from flash

```
R1(config)#boot system c2900-universalk9-mz.SPA.155-3.M4a.bin
R1(config)#exit
R1#delete flash:c2900-universalk9-mz.SPA.151-4.M4.bin
Delete filename [c2900-universalk9-mz.SPA.151-4.M4.bin]?
Delete flash:/c2900-universalk9-mz.SPA.151-4.M4.bin? [confirm]
R1#write
Building configuration...
[OK]

R1#show flash

System flash directory:
File  Length   Name/status
  4   33591768 c2900-universalk9-mz.SPA.155-3.M4a.bin

R1#reload
```

> 注意这里指定 boot system 是在 global config mode 中的

使用 `show version` 来校验

```
R2#show version
Cisco IOS Software, C2900 Software (C2900-UNIVERSALK9-M), Version 15.5(3)M4a, RELEASE SOFTWARE (fc1)
```

### 0x04

Use FTP on R2 to retrieve the following file from SRV1:c2900-universalk9-mz.SPA.155-3.M4a.bin

(FTP username: jeremy, password: ccna)

```
R1#conf t
R2(config)#ip ftp username jeremy
R2(config)#ip ftp password ccna
R2(config)#exit

R2#copy ftp: flash: 
Address or name of remote host []? 10.0.0.1
Source filename []? c2900-universalk9-mz.SPA.155-3.M4a.bin
Destination filename [c2900-universalk9-mz.SPA.155-3.M4a.bin]? 
%Warning:There is a file already existing with this name
Do you want to over write? [confirm]

Accessing ftp://10.0.0.1/c2900-universalk9-mz.SPA.155-3.M4a.bin...
[OK - 33591768 bytes]

33591768 bytes copied in 112.052 secs (31476 bytes/sec)
```

### 0x05

Upgrade R2’s OS and then delete the old file from flash

```
R2(config)#boot system c2900-universalk9-mz.SPA.155-3.M4a.bin
R2(config)#do write
Building configuration...
[OK]
R2(config)#exit
R2#delete flash:c2900-universalk9-mz.SPA.155-3.M4a.bin
Delete filename [c2900-universalk9-mz.SPA.155-3.M4a.bin]?
Delete flash:/c2900-universalk9-mz.SPA.155-3.M4a.bin? [confirm]
R2#reload
```

**references**

1. [^https://community.cisco.com/t5/other-network-architecture-subjects/difference-between-enable-secret-command-and-service-password/td-p/173461]