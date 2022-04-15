

ref:
[https://www.cloudflare.com/zh-cn/learning/cdn/glossary/time-to-live-ttl/](https://www.cloudflare.com/zh-cn/learning/cdn/glossary/time-to-live-ttl/)
[https://www.techtarget.com/searchnetworking/definition/time-to-live](https://www.techtarget.com/searchnetworking/definition/time-to-live)
[https://datatracker.ietf.org/doc/rfc3443/](https://datatracker.ietf.org/doc/rfc3443/)

https://www.youtube.com/watch?v=56sNDuYlhPk

https://subinsb.com/default-device-ttl-values/

https://community.cisco.com/t5/routing/tracert-show-same-hop-twice/td-p/1502358

## applicatons
在应用层面 TTL 用来管理 cache（对应http报文头中的`Expires`,`Cache-Control`），常备用在 CDN，DNS 用来表示内容存在的时间
## network
在网络（IP 数据包）中 TTL 的含义区别于应用中的
一个数据包从网络中传输，如果没有设置限制，就会在网络中 pass from router to router indefinitely。为了解决这个问题，所以才产生了 Time to Live (ttl) 也被称为 hop limit，该值由数据包的发送者设置（1 - 255），==如果没有设置就会常用OS的预设值（不同的OS会有不同的ttl默认值）。==

每==到达==路由的一跳，ttl 就会减 1，然后将包发送给下一跳。如果 ttl 的值为 0 了，数据包就会被discarded并回送ICMP echo( Type 11 Time Exceed，ICMP reply 不需要对应的 ICMP request，比如你发了一个TCP包也可以回送一个ICMP作为响应)

Each packet has a place where it stores a numerical value determining how much longer it should continue to move through the network. Every time a router receives a packet, it subtracts one from the TTL count and then passes it onto the next location in the network.
If at any point the TTL count is equal to zero after the subtraction, the router will discard the packet and send an ICMP message back to the originating host.

## default value

通常每个操作系统的 ttl 值都不同，Unix 通常是 64，windows 通常 128 ，具体可以参考如下表格。可以通过`ping localhost`来获取当前OS的默认 ttl

| Device / OS    | Version               | Protocol     | TTL  |
| -------------- | --------------------- | ------------ | ---- |
| AIX            |                       | TCP          | 60   |
| AIX            |                       | UDP          | 30   |
| AIX            | 3.2, 4.1              | ICMP         | 255  |
| BSDI           | BSD/OS 3.1 and 4.0    | ICMP         | 255  |
| Compa          | Tru64 v5.0            | ICMP         | 64   |
| Cisco          |                       | ICMP         | 254  |
| DEC Pathworks  | V5                    | TCP and UDP  | 30   |
| Foundry        |                       | ICMP         | 64   |
| FreeBSD        | 2.1R                  | TCP and UDP  | 64   |
| FreeBSD        | 3.4, 4.0              | ICMP         | 255  |
| FreeBSD        | 5                     | ICMP         | 64   |
| HP-UX          | 9.0x                  | TCP and UDP  | 30   |
| HP-UX          | 10.01                 | TCP and UDP  | 64   |
| HP-UX          | 10.2                  | ICMP         | 255  |
| HP-UX          | 11                    | ICMP         | 255  |
| HP-UX          | 11                    | TCP          | 64   |
| Irix           | 5.3                   | TCP and UDP  | 60   |
| Irix           | 6.x                   | TCP and UDP  | 60   |
| Irix           | 6.5.3, 6.5.8          | ICMP         | 255  |
| juniper        |                       | ICMP         | 64   |
| MPE/IX (HP)    |                       | ICMP         | 200  |
| Linux          | 2.0.x kernel          | ICMP         | 64   |
| Linux          | 2.2.14 kernel         | ICMP         | 255  |
| Linux          | 2.4 kernel            | ICMP         | 255  |
| Linux          | Red Hat 9             | ICMP and TCP | 64   |
| MacOS/MacTCP   | 2.0.x                 | TCP and UDP  | 60   |
| MacOS/MacTCP   | X (10.5.6)            | ICMP/TCP/UDP | 64   |
| NetBSD         |                       | ICMP         | 255  |
| Netgear FVG318 |                       | ICMP and UDP | 64   |
| OpenBSD        | 2.6 & 2.7             | ICMP         | 255  |
| OpenVMS        | 07.01.2002            | ICMP         | 255  |
| OS/2           | TCP/IP 3.0            |              | 64   |
| OSF/1          | V3.2A                 | TCP          | 60   |
| OSF/1          | V3.2A                 | UDP          | 30   |
| Solaris        | 2.5.1, 2.6, 2.7, 2.8  | ICMP         | 255  |
| Solaris        | 2.8                   | TCP          | 64   |
| Stratus        | TCP_OS                | ICMP         | 255  |
| Stratus        | TCP_OS (14.2-)        | TCP and UDP  | 30   |
| Stratus        | TCP_OS (14.3+)        | TCP and UDP  | 64   |
| Stratus        | STCP                  | ICMP/TCP/UDP | 60   |
| SunOS          | 4.1.3/4.1.4           | TCP and UDP  | 60   |
| SunOS          | 5.7                   | ICMP and TCP | 255  |
| Ultrix         | V4.1/V4.2A            | TCP          | 60   |
| Ultrix         | V4.1/V4.2A            | UDP          | 30   |
| Ultrix         | V4.2 – 4.5            | ICMP         | 255  |
| VMS/Multinet   |                       | TCP and UDP  | 64   |
| VMS/TCPware    |                       | TCP          | 60   |
| VMS/TCPware    |                       | UDP          | 64   |
| VMS/Wollongong | 1.1.1.1               | TCP          | 128  |
| VMS/Wollongong | 1.1.1.1               | UDP          | 30   |
| VMS/UCX        |                       | TCP and UDP  | 128  |
| Windows        | for Workgroups        | TCP and UDP  | 32   |
| Windows        | 95                    | TCP and UDP  | 32   |
| Windows        | 98                    | ICMP         | 32   |
| Windows        | 98, 98 SE             | ICMP         | 128  |
| Windows        | 98                    | TCP          | 128  |
| Windows        | NT 3.51               | TCP and UDP  | 32   |
| Windows        | NT 4.0                | TCP and UDP  | 128  |
| Windows        | NT 4.0 SP5-           |              | 32   |
| Windows        | NT 4.0 SP6+           |              | 128  |
| Windows        | NT 4 WRKS SP 3, SP 6a | ICMP         | 128  |
| Windows        | NT 4 Server SP4       | ICMP         | 128  |
| Windows        | ME                    | ICMP         | 128  |
| Windows        | 2000 pro              | ICMP/TCP/UDP | 128  |
| Windows        | 2000 family           | ICMP         | 128  |
| Windows        | Server 2003           |              | 128  |
| Windows        | XP                    | ICMP/TCP/UDP | 128  |
| Windows        | Vista                 | ICMP/TCP/UDP | 128  |
| Windows        | 7                     | ICMP/TCP/UDP | 128  |
| Windows        | Server 2008           | ICMP/TCP/UDP | 128  |
| Windows        | 10                    | ICMP/TCP/UDP | 128  |

例如 ping baidu.com，ttl 返回值 52，按照上表概率可以推算是 linux 服务器。假设是windows，128 - 52 = 76 需要经过 76 跳明显不符合逻辑，64 - 52 = 12 经过 12 跳

```
cpl in ~ λ ping baidu.com
PING baidu.com (220.181.38.148) 56(84) bytes of data.
64 bytes from 220.181.38.148 (220.181.38.148): icmp_seq=1 ttl=52 time=37.3 ms
64 bytes from 220.181.38.148 (220.181.38.148): icmp_seq=2 ttl=52 time=37.1 ms
```



## Example 

ping 和 traceroute 都使用了 TTL 这里在R1上使用 `traceroute 192.168.81.2`，cisco 命令默认使用 UDP

链路 R1 -> R2 -> R3，只配置了静态路由

```
R1#show run int fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.80.1 255.255.255.0
 duplex auto
 speed auto
end
```

```
R2#show run int fa1/0
Building configuration...

*Mar  1 00:30:39.611: %SYS-5-CONFIG_I: Configured from console by console
Current configuration : 97 bytes
!
interface FastEthernet1/0
 ip address 192.168.80.2 255.255.255.0
 duplex auto
 speed auto
end

R2#show run int fa0/1
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/1
 ip address 192.168.81.1 255.255.255.0
 duplex auto
 speed auto
end
```

```
R3#show run int fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.81.2 255.255.255.0
 duplex auto
 speed auto
end
```

[trace.pcap](/home/cpl/note/appendix)

![2022-03-28_20-48](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220328/2022-03-28_20-48.2mfdy7s8us40.webp)

traceroute 探测一次会发 3 个包，前 3 个包的 ttl 值为 1（由traceroute设置），到达 192.168.80.1 时 ttl - 1 值为 0 回送给源 ICMP type 11 (ttl 255 表示还未到达目的端，不可达，由路由器设置可以知道路由器ttl默认为255)，第二次探测的 3 个包的 ttl 值会设置为 2，但是到达了目的了，所以就没有第三次探测了，同时回送给源 ICMP type 11(ttl 254 ，会减掉 1 跳)。如果 ttl 的值到达了 30 就会终止

