---
createTime: 2025-04-09 22:15
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Zerotier 01 - Overview

## 0x01 Preface

> ZeroTier is a way to connect devices over your own private network anywhere in the world. You do this by creating a network and then joining _two or more devices_ to that network. You can use ZeroTier to play games, connect to remote business resources or even as a cloud backplane for your enterprise.[^1]

zerotier 是一个分布式的 VPN 网络方案，通过 zerotier 自有的协议加密传输

## 0x02 Installation

> 具体查看 [Download Page](https://docs.zerotier.com/releases/)

Arch 可以按照 [ZeroTier - ArchWiki](https://wiki.archlinux.org/title/ZeroTier) 安装

```
pacman -S zerotier-one
```

## 0x03 Configuration[^2]

假设有 2 台不同段的主机，Linux/Windows 要加入到 zerotier network

### 0x03a Create an Account if Not Exist

如果没有账号的话需要创建一个 zerotier 账号

https://my.zerotier.com/

### 0x03b Create an Virtual Network

直接在 my.zerotier.com 创建一个虚拟网络

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-09_23-43-04.83a82vwgk6.png)

这里点击 NETWORK ID 并拷贝，后面需要用到

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-10_00-05-43.60ufeup0h7.webp)

### 0x03c Setup Zerotier Node

#### Linux

启动 zerotier-one daemon，可以按需使用 `--enable-now`

```
systemctl start zerotier-one
```

检查 zerotier-one daemon 工作是否正常

```
sudo zerotier-cli info
200 info 3f7285bd35 1.14.2 ONLINE
```

其中 `3f7285bd35` 为 zerotier node ID(必须在整个虚拟网络唯一)

将 zerotier node 加入刚才创建好的虚拟网络

```
cc in ~ λ sudo zerotier-cli join db64858fed92e0c4
200 join OK
```

但是这时还不能直接使用 Zerotier 创建的网络，还需要虚拟网络对 node 授权。可以使用 `zerotier-cli listnetworks` 来查看 node 授权状态

```
sudo zerotier-cli listnetworks
200 listnetworks <nwid> <name> <mac> <status> <type> <dev> <ZT assigned ips>
200 listnetworks db64858fed92e0c4  c6:df:e0:68:32:b0 ACCESS_DENIED PRIVATE ztksero3zj -
```

这里可以看到 `ACCESS_DENIED` 表示没有授权

回到刚刚创建的虚拟网络管理界面，按照箭头示意操作授权

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-10_00-03-47.8l09rhmqq6.webp)

授权后 Auth column 状态会发生改变，同时 zerotier 会分配一个 IP 给 Node

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-10_00-10-45.3nrsxnhnhd.webp)

可以观察到这个 IP 会被分配到 zerotier 创建的 VNIC 上

```
ip -br a
ztksero3zj       UNKNOWN        172.28.20.203/16 fe80::c4df:e0ff:fe68:32b0/64
```

#### Windows

Windows 没有命令行，通过 GUI 操作

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-04-10_16-30-59.5q7lmp717s.webp)

输入对应的 network ID 即可。加入对应的网络后同样需要对 Node 授权，授权完成后 zerotier 会分配一个 IP 给 Node

```
C:>ipconfig
...
以太网适配器 ZeroTier One [db64858fed92e0c4]:

   连接特定的 DNS 后缀 . . . . . . . :
   本地链接 IPv6 地址. . . . . . . . : fe80::62bb:c8d0:f948:4dfc%53
   IPv4 地址 . . . . . . . . . . . . : 172.28.142.252
   子网掩码  . . . . . . . . . . . . : 255.255.0.0
   默认网关. . . . . . . . . . . . . : 25.255.255.254 
...
```

#### NAS

#### Mobile


### 0x03d Testing

测试 Linux 访问 Windows

```
sudo nping --tcp -p 3389 -c 3 172.28.142.252

Starting Nping 0.7.95 ( https://nmap.org/nping ) at 2025-04-10 17:00 CST
SENT (0.0066s) TCP 172.28.20.203:13589 > 172.28.142.252:3389 S ttl=64 id=27553 iplen=40  seq=1175287429 win=1480
RCVD (0.0114s) TCP 172.28.142.252:3389 > 172.28.20.203:13589 SA ttl=200 id=51991 iplen=44  seq=768404595 win=64000 <mss 2760>
SENT (1.0067s) TCP 172.28.20.203:13589 > 172.28.142.252:3389 S ttl=64 id=27553 iplen=40  seq=1175287429 win=1480
RCVD (1.0188s) TCP 172.28.142.252:3389 > 172.28.20.203:13589 SA ttl=200 id=51992 iplen=44  seq=768682473 win=64000 <mss 2760>
SENT (2.0078s) TCP 172.28.20.203:13589 > 172.28.142.252:3389 S ttl=64 id=27553 iplen=40  seq=1175287429 win=1480
RCVD (2.0143s) TCP 172.28.142.252:3389 > 172.28.20.203:13589 SA ttl=200 id=51993 iplen=44  seq=768951644 win=64000 <mss 2760>

Max rtt: 12.033ms | Min rtt: 4.794ms | Avg rtt: 7.786ms
Raw packets sent: 3 (120B) | Rcvd: 3 (132B) | Lost: 0 (0.00%)
Nping done: 1 IP address pinged in 2.03 seconds
```

测试 Windows 访问 Linux

```
C:>nping --tcp -p 22 -c 3 172.28.20.203

Starting Nping 0.7.95 ( https://nmap.org/nping ) at 2025-04-10 17:02 中国标准时间
SENT (0.2340s) TCP 172.28.142.252:9270 > 172.28.20.203:22 S ttl=64 id=15353 iplen=40  seq=3234238849 win=1480
RCVD (0.2410s) TCP 172.28.20.203:22 > 172.28.142.252:9270 SA ttl=64 id=0 iplen=44  seq=3440602201 win=49680 <mss 2760>
SENT (1.2360s) TCP 172.28.142.252:9270 > 172.28.20.203:22 S ttl=64 id=15353 iplen=40  seq=3234238849 win=1480
RCVD (1.2450s) TCP 172.28.20.203:22 > 172.28.142.252:9270 SA ttl=64 id=0 iplen=44  seq=3440602201 win=49680 <mss 2760>
SENT (2.2360s) TCP 172.28.142.252:9270 > 172.28.20.203:22 S ttl=64 id=15353 iplen=40  seq=3234238849 win=1480
RCVD (2.2870s) TCP 172.28.20.203:22 > 172.28.142.252:9270 SA ttl=64 id=0 iplen=44  seq=3440602201 win=49680 <mss 2760>

Max rtt: 51.000ms | Min rtt: 9.000ms | Avg rtt: 23.666ms
Raw packets sent: 3 (162B) | Rcvd: 3 (132B) | Lost: 0 (0.00%)
Nping done: 1 IP address pinged in 2.29 seconds 
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [What is ZeroTier? \| ZeroTier Documentation](https://docs.zerotier.com/wat/)
- [Create a Network \| ZeroTier Documentation](https://docs.zerotier.com/start/)
- [ZeroTier - ArchWiki](https://wiki.archlinux.org/title/ZeroTier)

***References***

[^1]:[What is ZeroTier? \| ZeroTier Documentation](https://docs.zerotier.com/wat/)
[^2]:[ZeroTier - ArchWiki](https://wiki.archlinux.org/title/ZeroTier#Configuration)
