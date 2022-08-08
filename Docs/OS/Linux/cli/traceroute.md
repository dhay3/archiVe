# Linux traceroute

ref:

[https://www.imperva.com/learn/ddos/ip-fragmentation-attack-teardrop/](https://www.imperva.com/learn/ddos/ip-fragmentation-attack-teardrop/)

[https://networkengineering.stackexchange.com/questions/54851/why-there-is-only-one-hop-in-the-tracert](https://networkengineering.stackexchange.com/questions/54851/why-there-is-only-one-hop-in-the-tracert)

[https://community.cisco.com/t5/routing/tracert-show-same-hop-twice/td-p/1502358](https://community.cisco.com/t5/routing/tracert-show-same-hop-twice/td-p/1502358)

[https://community.cisco.com/t5/network-security/tracert-same-ip-in-multiple-hops/td-p/2282189](https://community.cisco.com/t5/network-security/tracert-same-ip-in-multiple-hops/td-p/2282189)

[https://networkengineering.stackexchange.com/questions/7135/traceroute-many-hops-with-the-same-ip](https://networkengineering.stackexchange.com/questions/7135/traceroute-many-hops-with-the-same-ip)

[https://www.reddit.com/r/networking/comments/9yhl46/trace_route_hop_repeating_the_same_ip_address/](https://www.reddit.com/r/networking/comments/9yhl46/trace_route_hop_repeating_the_same_ip_address/)

[https://forum.netgate.com/topic/121406/traceroute-shows-the-same-address-for-each-hop](https://forum.netgate.com/topic/121406/traceroute-shows-the-same-address-for-each-hop)

[https://supportportal.juniper.net/s/article/ScreenOS-Tracert-shows-same-IP-address-for-each-hop?language=en_US#:~:text=When a tracert is initiated,have the same IP address](https://supportportal.juniper.net/s/article/ScreenOS-Tracert-shows-same-IP-address-for-each-hop?language=en_US#:~:text=When%20a%20tracert%20is%20initiated,have%20the%20same%20IP%20address)

https://www.hackingarticles.in/working-of-traceroute-using-wireshark/

https://stackoverflow.com/questions/73251815/traceroute-always-shows-the-packets-arrive-to-the-socket-address-even-the-port-w?noredirect=1#comment129367207_73251815

## Digest

syntax:`traceroute [options] host`

traceroute 是一个网络工具，显示从源到目的包的路径。不同 OS 实现的方式各有不同

==如果包能到达目的主机且有回包，traceroute 就认为是正常的。所以对于 4 层的 probing mode, 即使目的主机没有打开端口只要回包了( 不管他回了什么 RST 还是 SYN-ACK ) 也是显示正常的（换言之 traceroute 并不能用于 port scanning）==

提一嘴 ACL 通常直接会将包丢掉所以也就不会回包，traceroute 就认为是有问题的

### Principle

TTL 详情查看：[https://github.com/dhay3/archive/blob/master/Docs/Net/Grocery/TTL.md](https://github.com/dhay3/archive/blob/master/Docs/Net/Grocery/TTL.md)

traceroute 使用 IP 协议中 TTL 字段来实现，traceroute 开始探测时会发一个 ttl 值为 1 的，然后监听 nexthop 发回的 ICMP “time exceeded” 包（到达 nexthop 后仍不是目的 IP 包就会被丢弃，然后回送ICMP type 11），然后源将 ttl 值加 1 继续发包往目的IP（==根据 traceroute 实现的逻辑和方式不同，可能会出现异步下发的情况，即实际 3 跳到达目的机，但是traceroute 发包的 ttl 能到 9 甚至更大的情况==）

按照上述递归，根据 probing mode 探测的协议，直到回包是 ICMP type 3 port unreachable（默认使用UDP通常是30000以上的端口，如果端口没开就会回送 port unreachable） 或者是 TCP RST 或者是 hit max（ttl的值到了最大值，默认 30 hops）。

### Probing mode

> 为了方便讨论以下均以 来回路由 ACL 一致的情况讨论，即不考虑回包路由 ACL 导致 traceroute 显示 asterisk 的场景，只要数据包能到目的就一定能回包。
>
> 实际情况下有很多即使数据包到了目的主机，且目的主机回包了，但是由于回包路由ACL不一致，导致丢包，traceroute 显示 asterisk

traceroute 默认支持 3 种探测方式，ICMP，UDP（默认探测方式），TCP。默认针对每一跳会发 3 个探测包

数据包没有到达主机，意味着去方向路由不可达或者ACL(包含中间链路ACL和目的主机ACL)

数据包到达目的主机，意味着去方向路由可达

- UDP

  默认探测方式，为了不让目的端处理 UDP 包，探测端口默认 33434（通常是不会使用的端口），==每探测一次(不是每hop)==然后依次 + 1

  如果数据包没到达目的主机，会回送 ICMP type 11 exceed, 会显示 asterisk

  如果数据包到达目的主机，但是 UDP 端口没有开放，且端口没有对应进程，会回送 ICMP type 3 port unreachable

  如果数据包达到目的主机，且 UDP 端口开放，但是没有对应的进程，数据包会被丢弃，显示 asterisk

  如果数据包到达目的主机，且 UDP 端口开放，但是对应进程，没有回包的动作，显示 asterisk

  如果数据包到达目的主机，且 UDP 端口开放，对应进程回包(不管是什么类型的包)，显示正常

- TCP

  TPC 使用 half-open technique（半连接）

  如果数据包没有到达目的主机，会显示 asterisk, ==不会重传==

  如果数据包到达目的主机，如果对应的端口没有开放，会回送 TCP ACK-RST

  如果数据包到达目的主机，如果对应的端口开放，会回送 TCP ACK-SYN, 然后 traceroute 会 TCP RST

- ICMP

  按照正常 ICMP 报文回送

  如果数据包没到达目的主机，会回送 ICMP type 11 exceed，会显示 asterisk

  如果数据包到达主机，且主机没有禁ICMP，会回送 TCMP type 8 reply

### Method

traceroute 支持的所有探测模式

-  default（UDP 30000+）
-  icmp 
-  tcp 
-  tcpconn
使用全连接 
-  raw
只使用IP协议 

### Output

traceroute 默认会打印 3 个字段 TTL, address of the gateway, round trip time（RTT）

#### Asterisk

traceroute 如果没有收到 gateway 的回包就会显示 asterisk( * ), 如果正常会显示 address of gateway，==即回包路由的源接口IP==。

例如，这里显示 first hop 回包路由的源接口 IP 是 192.168.2.1

```
cpl in ~ λ traceroute -nI baidu.com
traceroute to baidu.com (39.156.66.10), 30 hops max, 60 byte packets
 1  192.168.2.1  19.017 ms  18.988 ms  18.984 ms
```

==一般发包路由和回包路由都相当==，但是由于逻辑接口的出现，==现在也会有发包路由和回包路由不同的情况(场景还挺多)==，如下图

first hop 显示的回包地址是 f1/1 配置的 IP

![](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220412/2022-04-12_21-28.1j5fzz3c2l9c.webp#crop=0&crop=0&crop=1&crop=1&id=trttn&originHeight=296&originWidth=987&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

如果 address of gateway 显示的是 asterisk（*），表示在指定时间内（默认5sec）没有从 gateway 收到回包，造成这种的原因通常有：

1. 发包没有到达节点，可能是没有路由或者ACL deny
1. 发包到了节点，但是回包没有路由或者ACL deny （一般出现在来回路由不一致的场景）
1. 回包的源IP是一个私网IP，到达运营商后被丢弃（如果私网IP到源有路由同样会回包给源）
1. 发的包是 UDP，大多数 router 或者 firewall 都会过滤 UDP 的包，因为 UDP 可能会打垮链路导致网络拥塞
1. 发的包是 ICMP，大多数 router 或者 firewall 都会对 ICMP 限流

==所以如果显示 * 并不一定表示 gateway 不可达，因为回包路由和入路由可能不一样，通过入路由能到达目的，但是节点因为路由或ACL原因没回包（在云主机中通常会出现这种情况）==

==为了排除这种现象，一般需要 both-direction traceroute==

#### Annotation

都是ICMP的回包类型

-  `!H | !N | !P`
分别表示host，network，protocol unreachable 
-  `!S`
source route failed 
-  `!F`
fragmentation needed（包没传完，需要分片），这时候可以加大MTU的值
如果数据包太大(大于MTU，通常是1500byte)，就会被分片（fragment），通过IP协议中 fragmentflag 来表示是否是分片包 
-  `!X`
communication administratively prohibted
通常是主机的ACL或防火墙（iptables，firewalld等等）未放通 
-  `!V`
host precedence violation 
-  `!C`
precedence cutoff in effect 
-  `!<num>`
ICMP unreachable code num 

## Optional args

1.  `-n`，`-w`，`-t` 在一定程度上会加快traceroute的探测速度 
1.  `--sport`，`--source`，`-zq` 可以提供绕过防火墙或ACL 
1.  `--back`，`-d` 可以提供以下debug信息 

### Probe args

- `-4 | -6`
  traceroute 使用 IPv4 还是 IPv6，默认traceroute会去解析 host （即使是IP也会去解析DNS，和windows一样） 

- `-I | --icmp`
  使用 ICMP 探测 

- `-T | --tcp`
  使用 TCP SYN 探测 

- `-U | --udp`
  使用UDP探测，但是区别与默认的UDP探测方式，使用固定的 53 端口 

- `-P protocol | --protocol=protocol`
  使用指定的协议探测，这样我们就可以使用http或是smtp来探测 

- `-M method | --moudule=name`
  是用指定的 method，可以配合使用`-O`来指定参数 

- `-O option | --options options`
  指定 method 使用的参数 

- `-m n | --max-hops`
  指定最大的 ttl，默认 30 

- `--back`
  如果来回路由不一致会显示出来 

- `-n`
  不会将 IP 解析成 hostname，在一定程度上能快处理速度 

- `-p port | --port=port`
  指定探测的端口，如果是UDP的每探测一次就会+1，如果是TCP会使用固定值 

- `-d | --debug`
  debug 

- `-F | --dont-fragement`
  对数据包不分片 

- `-f n | --first=n`
  what ttl to start，第一次探测的ttl 
  上述表示直接从ttl 5开始探测 

  ```
  cpl in ~ λ traceroute -nf 5 220.181.38.251
  traceroute to 220.181.38.251 (220.181.38.251), 30 hops max, 60 byte packets
   5  115.233.18.33  9.366 ms * 61.164.24.101  18.436 ms
  ```

- `-g gateway | --gateway=gateway`
  指定traceroute 第一跳探测使用的g ateway 

- `-i iface | --interface=iface`
  指定 traceroute 第一跳探测使用device，默认会自动选择 

- `-t tos | --tos=tos`
  设置tos的值，可以是 8 - 16 表示优先级从高到第 

- `-w max | --wait=max`
  等待回报的最长时间，默认 5 sec，如果在 5 sec 内没有回包会显示 asterisk 

- `-q n | --queries = n`
  每一跳探测几次，默认 3 次 

- `-s source_addr | --source=source_addr`
  指定发包的源地址，默认自动选择 

- `-z n | --sendwait n`
  每探测一次等待多长时间，默认0，如果firewall设置了 ICMP rate-limit 可以使用该参数。如果该值大于 10 就表示毫秒，小于10表示秒 

- `-A | --as-path-lookups`
  每探测一次都会打印出 AS path 

- `--sport=port`
  指定使用的源端口，同时也默认暗示使用`-N1 -w 5` 

## Cautions

### 0x1 Proxy

当目的 IP 在 LAN，通常一跳就会到达，即目的IP。但是有一种情况很特殊——主机开启了代理，例如：使用了 v2ray 或 软路由，TCP 流量会被引流到 proxy，由 proxy 完成 TCP 连接 ，并将数据回送源。而 proxy 的 IP 通常会设置成内网 IP （例如 192.168.80.1），这样就会导致源认为 one hop 就到了

```
cpl in ~ λ sudo traceroute -T baidu.com
traceroute to baidu.com (220.181.38.148), 30 hops max, 60 byte packets
 1  220.181.38.148 (220.181.38.148)  36.867 ms  36.694 ms  36.656 ms
```

### 0x2 Same IP from different hops

针对不同跳，出现同一个IP，可能有如下几种情况，以上述为例

第一种：

```
traceroute 10.10.50.5
1     10.10.10.1
2     10.10.20.2
3     10.10.30.3
4     10.10.40.4
5     10.10.40.4
6     10.10.50.5
```

第 4 跳是一个防火墙，第 5 跳回送包的第4 跳时做了 SNAT

![](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220223/lQLPDhtSJheUL73NAXrNAs6wfXzG7p6VCSoCV0QKgYA_AQ_718_378.1oofuh151pls.webp#crop=0&crop=0&crop=1&crop=1&id=UBL4l&originHeight=231&originWidth=737&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

第二种：

```
 1  74.117.154.4 (74.117.154.4)  82.222 ms  81.909 ms  80.970 ms
 2  74.117.154.1 (74.117.154.1)  83.212 ms  83.725 ms  81.852 ms
 3  74.117.154.4 (74.117.154.4)  97.692 ms  81.136 ms *
 4  74.117.154.1 (74.117.154.1)  83.025 ms  82.698 ms  88.137 ms
```

路由指向错误，形成循环路由，这个比较好理解

第三种：

```
1 61.3.8.10 2ms
2 74.117.154.4 2ms 76.117.154.4 1ms 3ms
3 74.117.154.4 3ms 2ms 3ms
4 110.1.1.1 3ms 
```

走公网运营商路径不同。有的路径比较优，先到了74.117.154.4

![2022-06-30_12-45](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220630/2022-06-30_12-45.6eh32yxdkls0.webp)

### 0x3 Not end
在使用 traceroute 时，可能会出现在路径中已经出现了目的 IP ，但是还在发包。
下列中 106.11.209.136  是 LVS vip 

```
traceroute to 106.11.209.136 (106.11.209.136), 30 hops max, 60 byte packets
 1  * * *
 2  11.73.1.105 (11.73.1.105)  2.316 ms 11.73.1.85 (11.73.1.85)  2.278 ms 11.73.1.105 (11.73.1.105)  1.928 ms
 3  10.102.225.157 (10.102.225.157)  1.576 ms 10.102.225.217 (10.102.225.217)  38.096 ms 10.102.225.173 (10.102.225.173)  1.543 ms
 4  11.94.128.166 (11.94.128.166)  4.295 ms * *
 5  10.102.41.70 (10.102.41.70)  32.268 ms 10.102.50.174 (10.102.50.174)  35.230 ms 10.102.50.138 (10.102.50.138)  33.652 ms
 6  10.102.43.81 (10.102.43.81)  30.350 ms 10.102.43.61 (10.102.43.61)  29.687 ms 117.49.45.117 (117.49.45.117)  29.743 ms
 7  * 10.54.254.14 (10.54.254.14)  30.357 ms *
 8  * 119.38.212.25 (119.38.212.25)  36.998 ms 103.52.73.193 (103.52.73.193)  36.619 ms
 9  * 106.11.209.136 (106.11.209.136)  36.941 ms *
10  11.17.124.133 (11.17.124.133)  40.483 ms 11.17.124.21 (11.17.124.21)  35.675 ms 11.129.153.73 (11.129.153.73)  35.450 ms
11  * 106.11.209.136 (106.11.209.136)  35.064 ms  36.415 ms
12  106.11.209.136 (106.11.209.136)  31.246 ms * *
13  106.11.209.136 (106.11.209.136)  35.705 ms  35.579 ms  37.568 ms
```
在回 119.38.212.25 的机器上 ping 包
```bash
PING 106.11.206.136: 56  data bytes, press CTRL_C to break
    Reply from 106.11.206.136: bytes=56 Sequence=1 ttl=245 time=44 ms
```

## Packet analyze

> 下面的例子均未使用`-q`或者`-w`修改 traceroute 默认行为的参数，除 probe mode 外
>
> 网络设备上的 traceroute 和 主机上的 traceroute 实现的原理和逻辑都有细微的差别

### Cisco

#### UDP

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

![](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220328/2022-03-28_20-48.2mfdy7s8us40.webp#crop=0&crop=0&crop=1&crop=1&id=f4fDK&originHeight=493&originWidth=2230&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

traceroute 探测一次会发 3 个包，前 3 个包的 ttl 值为 1（由traceroute设置），到达 192.168.80.1 时 ttl - 1 值为 0 回送给源 ICMP type 11 (ttl 255 表示还未到达目的端，不可达，由路由器设置可以知道路由器ttl默认为255)，第二次探测的 3 个包的 ttl 值会设置为 2，但是到达了目的了，所以就没有第三次探测了，同时回送给源 ICMP type 11(ttl 254 ，会减掉 1 跳)。如果 ttl 的值到达了 30 就会终止

### Linux

![2022-08-04_21-09](https://git.poker/dhay3/image-repo/blob/master/20220804/2022-08-04_21-09.2nejewalwv40.webp?raw=true)

大家都知道 TTL 在到路由设备时会 minus 1

为了对比在 centos-1 -> R1 和 R1 -> R2 中间链路同时抓包，观察 TTL 值的变化

#### Config

centos-1

```
[root@centos-1 /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
14: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether 6e:b0:66:0a:17:30 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.1/24 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::6cb0:66ff:fe0a:1730/64 scope link 
       valid_lft forever preferred_lft forever
[root@centos-1 /]# ip r
default via 192.168.1.2 dev eth0 
192.168.1.0/24 dev eth0 proto kernel scope link src 192.168.1.1
```

R1

```
R1#show ip int brief 
Interface              IP-Address      OK? Method Status                Protocol
FastEthernet0/0        192.168.1.2     YES manual up                    up      
FastEthernet1/0        192.168.2.1     YES manual up                    up

Gateway of last resort is not set

      192.168.1.0/24 is variably subnetted, 2 subnets, 2 masks
C        192.168.1.0/24 is directly connected, FastEthernet0/0
L        192.168.1.2/32 is directly connected, FastEthernet0/0
      192.168.2.0/24 is variably subnetted, 2 subnets, 2 masks
C        192.168.2.0/24 is directly connected, FastEthernet1/0
L        192.168.2.1/32 is directly connected, FastEthernet1/0
S     192.168.3.0/24 [1/0] via 192.168.2.2
```

R2

```
R2#show ip int br                                               
Interface              IP-Address      OK? Method Status                Protocol
FastEthernet0/0        192.168.3.2     YES manual up                    up      
FastEthernet1/0        192.168.2.2     YES manual up                    up

Gateway of last resort is not set

S     192.168.1.0/24 [1/0] via 192.168.2.1
      192.168.2.0/24 is variably subnetted, 2 subnets, 2 masks
C        192.168.2.0/24 is directly connected, FastEthernet1/0
L        192.168.2.2/32 is directly connected, FastEthernet1/0
      192.168.3.0/24 is variably subnetted, 2 subnets, 2 masks
C        192.168.3.0/24 is directly connected, FastEthernet0/0
L        192.168.3.2/32 is directly connected, FastEthernet0/0
```

centos-2

```
[root@centos-2 /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
13: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether fe:9e:c6:d7:36:7b brd ff:ff:ff:ff:ff:ff
    inet 192.168.3.1/24 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::fc9e:c6ff:fed7:367b/64 scope link 
       valid_lft forever preferred_lft forever
[root@centos-2 /]# ip r
default via 192.168.3.2 dev eth0 
192.168.3.0/24 dev eth0 proto kernel scope link src 192.168.3.1
```

#### ICMP

```
[root@centos-1 /]# traceroute -nI 192.168.3.1
traceroute to 192.168.3.1 (192.168.3.1), 30 hops max, 60 byte packets
 1  192.168.1.2  7.294 ms  17.587 ms  18.090 ms
 2  192.168.2.2  27.767 ms  29.566 ms  29.937 ms
 3  192.168.3.1  38.133 ms  39.843 ms  40.255 ms
```

**centos-1 <-> R1**

这里能发现 `maxttl(echo request) == 7`，但是实际 3 跳就应该到了，逻辑上应该在 `ttl==3`的时候就应该停止发包了。所以说明 traceroute 概率是==异步发包==的。

![2022-08-05_15-51](https://git.poker/dhay3/image-repo/blob/master/20220805/2022-08-05_15-51.4zly7xs6z9mo.webp?raw=true)

观察 19th 和 35th, 是“一对” （注意 ICMP 不是对称的协议，这里只是为了逻辑上方便理解所以才称为一对）

```
19	149.189845	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=1/256, ttl=1 (no response found!)
35	149.200114	192.168.1.2	192.168.1.1	ICMP	70	Time-to-live exceeded (Time to live exceeded in transit)
```

traceroute 将 ttl 设置为 1，当数据包到达 R1 时，做 ttl minus 1，发现 `ttl == 0`，R1 查找路由通过 fa0/0 回送 ICMP type 11, ttl 255

观察 22th 和 41th, 是“一对”

```
22	149.189907	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=4/1024, ttl=2 (no response found!)
41	149.220830	192.168.2.2	192.168.1.1	ICMP	70	Time-to-live exceeded (Time to live exceeded in transit)
```

traceroute 将 ttl 设置为 2，因为经过 R1 已经做了 minus 1，到 R2 再做 minus 1 时，发现 `ttl == 0` ， R2 查看路由通过 fa0/0 回送 ICMP type 11, ttl 255，数据包经过 R1 minus 1 所以是 254

观察 25th 和 47th, 是“一对”

```
25	149.189917	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=7/1792, ttl=3 (reply in 47)
47	149.231361	192.168.3.1	192.168.1.1	ICMP	74	Echo (ping) reply    id=0x005d, seq=7/1792, ttl=62 (request in 25)
```

traceroute 将 ttl 设置为 3，经过 R1 和 R2 minus 2，ttl 的值不为 0，且在 centos-2 上发现对应的目的 IP。所以 centos-2 回送 ICMP reply，这时的 ttl 使用 OS 默认的 `ttl == 64`(对应`net.ipv4.ip_default_ttl`，这时还没有过路由器)。过了 R2 和 R1 所以 ttl 值 minus 2

观察 34th 和 56th, 是“一对”

```
34	149.189945	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=16/4096, ttl=6 (reply in 56)
56	149.233419	192.168.3.1	192.168.1.1	ICMP	74	Echo (ping) reply    id=0x005d, seq=16/4096, ttl=62 (request in 34)
```

这里可以看到 traceroute 将 ttl 设置成了 6，经过 R1 和 R2 minus 2，ttl 的值不为 0，且在 centos-2 上发现对应的目的 IP。所以 centos-2 回送 ICMP reply，这时的 ttl 使用 OS 默认的 `ttl == 64`(对应`net.ipv4.ip_default_ttl`，这时还没有过路由器)。过了 R2 和 R1 所以 ttl 值 minus 2

**R1 <-> R2**

![2022-08-05_15-12](https://git.poker/dhay3/image-repo/blob/master/20220805/2022-08-05_15-12.5fo89nbeif40.webp?raw=true)

traceroute 首先将 ttl 设置为1，数据包根据路由达到了 R1，R1 将 ttl minus 1，发现 `ttl == 0`，R1 查找路由通过 fa0/0 回送 ICMP type 11, ttl 255 。==这时数据包并没有通过 R1 转发到 R2 的链路上，所以上面并不会显示关联的数据包==

观察 46th 和 62th，是“一对”

```
46	187.337566	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=4/1024, ttl=1 (no response found!)
62	187.348383	192.168.2.2	192.168.1.1	ICMP	70	Time-to-live exceeded (Time to live exceeded in transit)
```

traceroute 将 ttl 设置为 2，当数据包到达 R1 时，做 ttl minus 1，R1 查找路由将数据包转发到 R2，当数据包到达 R2 时，这时 ttl 是 1 即对应 46th，R2 对 ttl minus 1，发现 `ttl == 0`。然后 R2 根据路由从 fa0/0 回送 ICMP type 11, ttl 255

观察 49th 和 65th，是一对收发包

```
49	187.337602	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=7/1792, ttl=2 (reply in 65)
65	187.359015	192.168.3.1	192.168.1.1	ICMP	74	Echo (ping) reply    id=0x005d, seq=7/1792, ttl=63 (request in 49)
```

traceroute 将 ttl 设置为 3，当数据包到达 R1 时，做 ttl minus 1 然后转发数据包到 R2，所以这时 ttl 为 2 对应 49th。数据包到达 R2 时，R2 做 ttl minus 1 然后根据路由转发到 centos-2。数据包到了 centos-2 发现 IP 匹配，所以回送 ICMP reply ttl 64，数据包经过 R2 minus 1，所以 65th 的 ttl 值为 63

观察 55th 和 71th，是“一对”

```
55	187.337661	192.168.1.1	192.168.3.1	ICMP	74	Echo (ping) request  id=0x005d, seq=13/3328, ttl=4 (reply in 71)
71	187.359038	192.168.3.1	192.168.1.1	ICMP	74	Echo (ping) reply    id=0x005d, seq=13/3328, ttl=63 (request in 55)
```

traceroute 将 ttl 设置为 5，经过 R1 minus 1，对应 55th。centos-2 回包，经过 R2 minus 1 ，对应 71th

#### TCP

**centos-1 <-> R1 4层TCP通**

在 centos-2 上监听了 80 port

```
[root@centos-2 /]# nc -lkv 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Listening on :::80
Ncat: Listening on 0.0.0.0:80

[root@centos-1 network-scripts]# traceroute -Tp 80 192.168.3.1
traceroute to 192.168.3.1 (192.168.3.1), 30 hops max, 60 byte packets
 1  _gateway (192.168.1.2)  15.762 ms  18.357 ms  18.367 ms
 2  192.168.2.2 (192.168.2.2)  25.906 ms  25.916 ms  25.914 ms
 3  192.168.3.1 (192.168.3.1)  35.998 ms  36.012 ms  36.009 ms
```

从 TCP SYN 包的数量和包内3层 ttl 值看，和 ICMP probe mode 一样会发大于实际 hops，ttl 值的包。每个到达目的主机的包，源都会主动对 socket pair 发 RST 包

![2022-08-05_21-12](https://git.poker/dhay3/image-repo/blob/master/20220805/2022-08-05_21-12.sz1n0ne17jk.webp?raw=true)

观察 8th 和 51th

```
8	51.957253	192.168.1.1	192.168.3.1	TCP	74	55433 → 80 [SYN] Seq=0 Win=5840 Len=0 MSS=1460 SACK_PERM=1 TSval=20262822 TSecr=0 WS=4
Frame 8: 74 bytes on wire (592 bits), 74 bytes captured (592 bits) on interface -, id 0
Ethernet II, Src: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98), Dst: ca:01:09:a1:00:00 (ca:01:09:a1:00:00)
Internet Protocol Version 4, Src: 192.168.1.1, Dst: 192.168.3.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 60
    Identification: 0xd253 (53843)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 1
        [Expert Info (Note/Sequence): "Time To Live" only 1]
    Protocol: TCP (6)
    Header Checksum: 0x6216 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.1
    Destination Address: 192.168.3.1
Transmission Control Protocol, Src Port: 55433, Dst Port: 80, Seq: 0, Len: 0


24	51.972239	192.168.1.2	192.168.1.1	ICMP	70	Time-to-live exceeded (Time to live exceeded in transit)
Frame 24: 70 bytes on wire (560 bits), 70 bytes captured (560 bits) on interface -, id 0
Ethernet II, Src: ca:01:09:a1:00:00 (ca:01:09:a1:00:00), Dst: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98)
Internet Protocol Version 4, Src: 192.168.1.2, Dst: 192.168.1.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0xc0 (DSCP: CS6, ECN: Not-ECT)
    Total Length: 56
    Identification: 0x007a (122)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 255
    Protocol: ICMP (1)
    Header Checksum: 0x3737 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.2
    Destination Address: 192.168.1.1
Internet Control Message Protocol
```

traceroute 将数据包 ttl 置为 1，TCP flag 置为 SYN，对应 8th。当数据包到达 R1 时，minus 1，根据路由回送 ICMP type 11, ttl 255 对应 24 th

观察 11th 和 27th

```
11	51.957445	192.168.1.1	192.168.3.1	TCP	74	36599 → 80 [SYN] Seq=0 Win=5840 Len=0 MSS=1460 SACK_PERM=1 TSval=20262822 TSecr=0 WS=4
Frame 11: 74 bytes on wire (592 bits), 74 bytes captured (592 bits) on interface -, id 0
Ethernet II, Src: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98), Dst: ca:01:09:a1:00:00 (ca:01:09:a1:00:00)
Internet Protocol Version 4, Src: 192.168.1.1, Dst: 192.168.3.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 60
    Identification: 0xd256 (53846)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 2
        [Expert Info (Note/Sequence): "Time To Live" only 2]
    Protocol: TCP (6)
    Header Checksum: 0x6113 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.1
    Destination Address: 192.168.3.1
Transmission Control Protocol, Src Port: 36599, Dst Port: 80, Seq: 0, Len: 0


27	51.982450	192.168.2.2	192.168.1.1	ICMP	70	Time-to-live exceeded (Time to live exceeded in transit)
Frame 27: 70 bytes on wire (560 bits), 70 bytes captured (560 bits) on interface -, id 0
Ethernet II, Src: ca:01:09:a1:00:00 (ca:01:09:a1:00:00), Dst: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98)
Internet Protocol Version 4, Src: 192.168.2.2, Dst: 192.168.1.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0xc0 (DSCP: CS6, ECN: Not-ECT)
    Total Length: 56
    Identification: 0x0077 (119)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 254
    Protocol: ICMP (1)
    Header Checksum: 0x373a [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.2.2
    Destination Address: 192.168.1.1
Internet Control Message Protocol
```

traceroute 将数据包 ttl 置为 2，TCP flag 置为 SYN，对应 11th。当数据包到达 R1 时，minus 1 值为 1，然后根据路由转发到 R2。R2 收到数据包后同样做 ttl minus 1，发现 ttl 的值为 0，会送 ICMP type 11, ttl 255。数据包经过 R1，minus 1，ttl 254 对应 27th

观察 14th 和 30th 以及 40th（==以 socket pair 为一组，直接 follow tcp stream 即可==），==wireshark 表示错误请忽略==

```
14	51.957458	192.168.1.1	192.168.3.1	TCP	74	46083 → 80 [SYN] Seq=0 Win=5840 Len=0 MSS=1460 SACK_PERM=1 TSval=20262822 TSecr=0 WS=4
Frame 14: 74 bytes on wire (592 bits), 74 bytes captured (592 bits) on interface -, id 0
Ethernet II, Src: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98), Dst: ca:01:09:a1:00:00 (ca:01:09:a1:00:00)
Internet Protocol Version 4, Src: 192.168.1.1, Dst: 192.168.3.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 60
    Identification: 0xd259 (53849)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 3
        [Expert Info (Note/Sequence): "Time To Live" only 3]
    Protocol: TCP (6)
    Header Checksum: 0x6010 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.1
    Destination Address: 192.168.3.1
Transmission Control Protocol, Src Port: 46083, Dst Port: 80, Seq: 0, Len: 0


30	51.992561	192.168.3.1	192.168.1.1	TCP	74	80 → 46083 [SYN, ACK] Seq=0 Ack=1 Win=65160 Len=0 MSS=1460 SACK_PERM=1 TSval=4191785085 TSecr=20262822 WS=128
Frame 30: 74 bytes on wire (592 bits), 74 bytes captured (592 bits) on interface -, id 0
Ethernet II, Src: ca:01:09:a1:00:00 (ca:01:09:a1:00:00), Dst: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98)
Internet Protocol Version 4, Src: 192.168.3.1, Dst: 192.168.1.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 60
    Identification: 0x0000 (0)
    Flags: 0x40, Don't fragment
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 62
    Protocol: TCP (6)
    Header Checksum: 0xb769 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.3.1
    Destination Address: 192.168.1.1
Transmission Control Protocol, Src Port: 80, Dst Port: 46083, Seq: 0, Ack: 1, Len: 0

40	51.992647	192.168.1.1	192.168.3.1	TCP	54	46083 → 80 [RST] Seq=1 Win=0 Len=0
Frame 40: 54 bytes on wire (432 bits), 54 bytes captured (432 bits) on interface -, id 0
Ethernet II, Src: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98), Dst: ca:01:09:a1:00:00 (ca:01:09:a1:00:00)
Internet Protocol Version 4, Src: 192.168.1.1, Dst: 192.168.3.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 40
    Identification: 0x0000 (0)
    Flags: 0x40, Don't fragment
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 64
    Protocol: TCP (6)
    Header Checksum: 0xb57d [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.1
    Destination Address: 192.168.3.1
Transmission Control Protocol, Src Port: 46083, Dst Port: 80, Seq: 1, Len: 0
```

traceroute 将数据包 ttl 置为 3，TCP flag 置为 SYN，对应 14th。当数据包到达 R1 时，minus 1 值为 1，然后根据路由转发到 R2。R2 收到数据包后同样做 ttl minus 1，然后根据路由转发到 centos-2 发现 ttl 值不为 0，且对应端口开放无ACL，会送 TCP flag SYN-ACK, ttl 64，经过 R2 minus 1，经过 R1 minus 1，所以最后 ttl 为 62，对应 30th。这是 centos-2 创建了 half-open socket，为了不消耗 centos-2 的资源，centos-1 发了 TCP flag RST， ttl 64（这里是系统默认的 ttl 值）

**centos-1 <-> R1 4层TCP不通**

==如果包能到达目的主机，traceroute 就认为是正常的，所以对于 4 层的 probing mode, 即使目的主机没有打开端口只要回包了(不管他回了什么)也是显示正常的（换言之 traceroute 并不能用于 port scanning）==

https://stackoverflow.com/questions/73251815/traceroute-always-shows-the-packets-arrive-to-the-socket-address-even-the-port-w?noredirect=1#comment129367481_73251815

```
[root@centos-1 network-scripts]# nc -nvz 192.168.3.1 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connection refused.

[root@centos-1 network-scripts]# traceroute -nTp 80 192.168.3.1
traceroute to 192.168.3.1 (192.168.3.1), 30 hops max, 60 byte packets
 1  192.168.1.2  5.181 ms  16.214 ms  16.528 ms
 2  192.168.2.2  26.723 ms  27.379 ms  27.498 ms
 3  192.168.3.1  37.517 ms  38.446 ms  39.343 ms
 
[root@centos-2 /]# ss -lnpt
State             Recv-Q             Send-Q                           Local Address:Port                           Peer Address:Port             Process
```

![2022-08-05_23-34](https://git.poker/dhay3/image-repo/blob/master/20220805/2022-08-05_23-34.5bzzqai490g0.webp?raw=true)

ttl 值未到 3 之前的逻辑不再赘述，这里只讨论 ttl 值为 3 的情况，即包达到了目的主机

观察 523th 和 545th

```
523	3509.113807	192.168.1.1	192.168.3.1	TCP	74	47731 → 80 [SYN] Seq=0 Win=5840 Len=0 MSS=1460 SACK_PERM=1 TSval=3518895646 TSecr=0 WS=4
Frame 523: 74 bytes on wire (592 bits), 74 bytes captured (592 bits) on interface -, id 0
Ethernet II, Src: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98), Dst: ca:01:09:a1:00:00 (ca:01:09:a1:00:00)
Internet Protocol Version 4, Src: 192.168.1.1, Dst: 192.168.3.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 60
    Identification: 0x7684 (30340)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 3
    Protocol: TCP (6)
    Header Checksum: 0xbbe5 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.1
    Destination Address: 192.168.3.1
Transmission Control Protocol, Src Port: 47731, Dst Port: 80, Seq: 0, Len: 0


545	3509.148458	192.168.3.1	192.168.1.1	TCP	54	80 → 47731 [RST, ACK] Seq=1 Ack=1 Win=0 Len=0
Frame 545: 54 bytes on wire (432 bits), 54 bytes captured (432 bits) on interface -, id 0
Ethernet II, Src: ca:01:09:a1:00:00 (ca:01:09:a1:00:00), Dst: 0a:3b:66:03:18:98 (0a:3b:66:03:18:98)
Internet Protocol Version 4, Src: 192.168.3.1, Dst: 192.168.1.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 40
    Identification: 0x0000 (0)
    Flags: 0x40, Don't fragment
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 62
    Protocol: TCP (6)
    Header Checksum: 0xb77d [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.3.1
    Destination Address: 192.168.1.1
Transmission Control Protocol, Src Port: 80, Dst Port: 47731, Seq: 1, Ack: 1, Len: 0
```

traceroute 将 ttl 的值设置成 3，TCP flag SYN，对应 523th。数据包到 R1 ttl minus 1, 转发到 R2。数据包到 R2 ttl minus 1，转发到 centos-2。centos-2 发现没有对应的socket address（没有打开80端口），回送 TCP flag ACK-RST，ttl 64。数据包到 R2 ttl minus 1，转发到 R1。数据包到 R1 ttl minus 1，转发到 centos-1，对应 545th

**centos-1 <-> R1 4层 TCP ACL不通**

使用 iptables 制造 4 层 ACL，方通 3 层

```
[root@netos-2 /]# iptables -vnL 
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       192.168.1.1          0.0.0.0/0           

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

[root@netos-2 /]# iptables-save 
# Generated by iptables-save v1.8.4 on Sun Aug  7 09:15:14 2022
*filter
:INPUT ACCEPT [0:0]
:FORWARD ACCEPT [0:0]
:OUTPUT ACCEPT [0:0]
-A INPUT -s 192.168.1.1/32 -j DROP
COMMIT
# Completed on Sun Aug  7 09:15:14 2022

[root@netos-1 /]# traceroute -Tp 80 192.168.3.1
traceroute to 192.168.3.1 (192.168.3.1), 30 hops max, 60 byte packets
 1  _gateway (192.168.1.2)  6.906 ms  8.908 ms  8.996 ms
 2  192.168.2.2 (192.168.2.2)  27.245 ms  28.067 ms  28.375 ms
 3  * * *
 4  * * *
 5  * * *
 6  * * *
 7  * * *
 8  *^C
```

![2022-08-08_12-23](https://git.poker/dhay3/image-repo/blob/master/20220807/2022-08-08_12-23.6fcw0tzmnfk0.webp?raw=true)

从 569th 开始 ttl 值为 3，可以从上图看到 centos-2 并没有回包给 centos-1，也没有出现重传

#### UDP

**centos-1 <-> R1 默认**

默认 traceroute 会使用 33434 作为探测端口，之后每探测一次（==不是每hop==）port 就会自增。如果对应的端口没有开放，按照 RFC 默认会回送 ICMP type 3 port unreachable

```
[root@netos-1 /]# traceroute 192.168.3.1
traceroute to 192.168.3.1 (192.168.3.1), 30 hops max, 60 byte packets
 1  _gateway (192.168.1.2)  24.079 ms  24.806 ms  24.805 ms
 2  192.168.2.2 (192.168.2.2)  24.766 ms  24.783 ms  24.781 ms
 3  192.168.3.1 (192.168.3.1)  24.785 ms  24.783 ms  25.889 ms
```

![2022-08-07_16-08](https://git.poker/dhay3/image-repo/blob/master/20220807/2022-08-07_16-08.4k5fjwb9mmtc.webp?raw=true)

观察 40th 和 56th

```
40	273.584313	192.168.1.1	192.168.3.1	UDP	74	36621 → 33440 Len=32
Frame 40: 74 bytes on wire (592 bits), 74 bytes captured (592 bits) on interface -, id 0
Ethernet II, Src: 5e:e2:e3:c4:66:9c (5e:e2:e3:c4:66:9c), Dst: ca:01:09:a1:00:00 (ca:01:09:a1:00:00)
Internet Protocol Version 4, Src: 192.168.1.1, Dst: 192.168.3.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT)
    Total Length: 60
    Identification: 0x709d (28829)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 3
    Protocol: UDP (17)
    Header Checksum: 0xc1c1 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.1.1
    Destination Address: 192.168.3.1
User Datagram Protocol, Src Port: 36621, Dst Port: 33440
Data (32 bytes)


56	273.607823	192.168.3.1	192.168.1.1	ICMP	102	Destination unreachable (Port unreachable)
Frame 56: 102 bytes on wire (816 bits), 102 bytes captured (816 bits) on interface -, id 0
Ethernet II, Src: ca:01:09:a1:00:00 (ca:01:09:a1:00:00), Dst: 5e:e2:e3:c4:66:9c (5e:e2:e3:c4:66:9c)
Internet Protocol Version 4, Src: 192.168.3.1, Dst: 192.168.1.1
    0100 .... = Version: 4
    .... 0101 = Header Length: 20 bytes (5)
    Differentiated Services Field: 0xc0 (DSCP: CS6, ECN: Not-ECT)
    Total Length: 88
    Identification: 0x9150 (37200)
    Flags: 0x00
    ...0 0000 0000 0000 = Fragment Offset: 0
    Time to Live: 62
    Protocol: ICMP (1)
    Header Checksum: 0x6542 [validation disabled]
    [Header checksum status: Unverified]
    Source Address: 192.168.3.1
    Destination Address: 192.168.1.1
Internet Control Message Protocol
Data (32 bytes)
```

traceroute 将 ttl 的值设置成 3，UDP，对应 40th。数据包到了 R1 ttl minus1，转发到 R2。数据包到 R2 ttl minus 1, 转发到 centos-2。centos-2 发现没有对应的socket address，回送 ICMP port unreachable, ttl 64。数据包到了 R2 ttl minus 1, 转发到 R1。数据包到了 R2 ttl minus 1，转发到 centos-1, 对应 56th

**centos-1 <-> R1 4层 UDP 监听指定端口不回包**

因为 centos-2 udp 80 打开的进程是 cat。当 traceroute 发包时，是没有 4 层 segment 的，cat 也就不能正常执行将回显返回给 centos-1（换言之就是 centos-1 是收不到 centos-2 回的包的），所以后面会显示 asterisks

netcat 之所以是正常的，因为 UDP 没有确认机制，只要包发出去就可以了(==但是不能回 ICMP type 3 port unreachable，如果回了 netcat 就会认为是异常的==)

```
[root@centos-2 network-scripts]# nc -lkvu 80 --exec cat
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Listening on :::80
Ncat: Listening on 0.0.0.0:80
...
exec: No such file or directory
Ncat: Connection from 192.168.1.1.
...

[root@centos-1 network-scripts]# nc -nvzu 192.168.3.1 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connected to 192.168.3.1:80.
Ncat: UDP packet sent successfully
Ncat: 1 bytes sent, 0 bytes received in 2.01 seconds.

[root@centos-1 network-scripts]# nc -nvzu 192.168.3.1 1900
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connected to 192.168.3.1:1900.
Ncat: Connection refused.

[root@centos-1 network-scripts]# traceroute -nUp 80 192.168.3.1
traceroute to 192.168.3.1 (192.168.3.1), 30 hops max, 60 byte packets
 1  192.168.1.2  24.949 ms  27.535 ms  27.536 ms
 2  192.168.2.2  63.526 ms  63.851 ms  63.960 ms
 3  * * *
 4  * * *
 ...
```

逻辑上和 TCP 4层不回包一样，这里不做分析 

**centos-1 <-> R1 4 层 UDP 监听指定端口回包**

需要使用代码实现

```
//TODO
```

回包 traceroute 就会认为是正常的

#### Inclusion

1. ==即使源到目的中间一共需要经过 3 跳，实际 traceroute 发的包的 ttl 值不一定等于 3，一般会大于 3。甚至 ttl 能达到 9==
2. ==只要包到了目的主机，目的主机回送任意包，traceroute 都会认为正常。所以 traceroute 并不能做 port scanning==
3. TCP probing mode , 即使端口是关闭的，回了 RST 包，traceroute 也会认为是正常的
4. UDP probing mode, 因为 UDP 默认没有确认机制，所以可以选择不回包。即使目的端口是打开的，traceroute 也会认为是异常的
5. traceroute 默认还是严格按照`-q`参数预设的值，对每 hop 进行探测

## Source code Patching

https://github.com/openbsd/src/blob/master/usr.sbin/traceroute/traceroute.c

TODO

