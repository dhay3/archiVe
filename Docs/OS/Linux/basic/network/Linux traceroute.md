# Linux traceroute

ref:

https://www.imperva.com/learn/ddos/ip-fragmentation-attack-teardrop/

https://networkengineering.stackexchange.com/questions/54851/why-there-is-only-one-hop-in-the-tracert

https://community.cisco.com/t5/routing/tracert-show-same-hop-twice/td-p/1502358

https://community.cisco.com/t5/network-security/tracert-same-ip-in-multiple-hops/td-p/2282189

https://networkengineering.stackexchange.com/questions/7135/traceroute-many-hops-with-the-same-ip

https://www.reddit.com/r/networking/comments/9yhl46/trace_route_hop_repeating_the_same_ip_address/

https://forum.netgate.com/topic/121406/traceroute-shows-the-same-address-for-each-hop

https://supportportal.juniper.net/s/article/ScreenOS-Tracert-shows-same-IP-address-for-each-hop?language=en_US#:~:text=When%20a%20tracert%20is%20initiated,have%20the%20same%20IP%20address

## 0x1Digest

syntax:`traceroute [options] host`

traceroute 是 LInux 上的一个网络工具，显示从源到目的包的路径

### principle

TTL 详情查看：https://github.com/dhay3/archive/blob/master/Docs/Net/Grocery/TTL.md

traceroute 使用 IP 协议中 TTL 字段来实现，traceroute 开始探测时会发一个 ttl 值为 1 的，然后监听 nexthop 发回的 ICMP “time exceeded” 包（到达 nexthop 后仍不是目的 IP 包就会被丢弃，然后回送ICMP type 11），然后 源将 ttl 值加 1 继续发包往目的IP ，按照上述递归，直到回包是 ICMP “port unreachable”（默认使用UDP通常是30000以上的端口，如果端口没开就会回送 port unreachable） 或者是 TCP rest 或者是 hit max（ttl的值到了最大值，默认 30 hops）

### probe mode

traceroute 默认支持 3 种探测方式，ICMP，UDP（默认探测方式），TCP

每针对一跳会发 3 个探测包

#### UDP

默认探测方式，为了不让目的端处理UDP包，探测端口默认33434（通常是不会使用的端口），每探测一次然后依次 + 1

#### TCP

TPC使用 half-open technique（半连接）

如果TCP探测的端口在目的端没有开放，会回送TCP reset

#### ICMP

按照正常ICMP报文回送

### method

- default

- icmp

- tcp

- tcpconn

  使用全连接

- raw

  只使用IP协议

### output

traceroute 默认会打印 3 个字段 ttl ，address of the gateway and round trip time（rtt）

#### asterisk

address of gateway 显示的是 gateway 回包路由的源接口（下图中的f0/1）， 一般发包路由和回包路由都相同，但是也有可能发包路由和回包路由不同

![2022-04-12_21-28](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220412/2022-04-12_21-28.1j5fzz3c2l9c.webp)

如果address of gateway 显示的是 asterisk（*），表示在指定时间内（默认5sec）没有从 gateway 收到回包，造成这种的原因通常有

1. 回包的链路中路由缺失，可以是回包链路中的任何一个节点（来回路由通常一样，出现这种情况概率在是来回路由不一致）
2. 回包的源IP是一个私网IP，到达运营商后被丢弃
3. 回包的链路中有ACL，可以是回包链路中的任何一个节点
4. 当前大多数 firewall 都会过滤 UDP 端口，甚至是ICMP，碰到这种情况可以使用其他协议(TCP)来绕过 firewall

==但是如果显示 * 并不表示 gateway 不可达，因为回包路由和入路由可能不一样，通过入路由能到达目的，但是节点因为路由或ACL原因没回包（在云主机中通常会出现这种情况）==

#### annotation

都是ICMP的回包类型

- `!H | !N | !P`

  分别表示host，network，protocol unreachable

- `!S`

  source route failed

- `!F`

   fragmentation needed（包没传完，需要分片），这时候可以加大MTU的值

  如果数据包太大(大于MTU，通常是1500byte)，就会被分片（fragment），通过IP协议中 fragmentflag 来表示是否是分片包

- `!X`

  communication administratively prohibted

  通常是主机的ACL或防火墙（iptables，firewalld等等）未放通

- `!V`

  host precedence violation

- `!C`

  precedence cutoff in effect

- `!<num>`

  ICMP unreachable code num

## 0x2 Optional args

1. `-n`，`-w`，`-t` 在一定程度上会加快traceroute的探测速度
2. `--sport`，`--source`，`-zq` 可以提供绕过防火墙或ACL

3. `--back`，`-d` 可以提供以下debug信息

### probe args

- `-4 | -6`

  traceroute 使用 IPv4 还是 IPv6，默认traceroute会去解析 host （==即使是IP也会去解析DNS，和windows一样==）

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

  ==如果来回路由不一致会显示出来==

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

  ```
  cpl in ~ λ traceroute -nf 5 220.181.38.251
  traceroute to 220.181.38.251 (220.181.38.251), 30 hops max, 60 byte packets
   5  115.233.18.33  9.366 ms * 61.164.24.101  18.436 ms
  ```

  上述表示直接从ttl 5开始探测

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

## 0x3 Cautions

### proxy

当目的 IP 在 LAN，通常一跳就会到达，即目的IP。但是有一种情况很特殊——主机开启了代理，例如：使用了 v2ray 或 软路由，TCP 流量会被引流到 proxy，由 proxy 完成TCP连接 ，并将数据回送源。而 proxy 的 IP 通常会设置成内网 IP （192.168.80.1），这样就会导致源认为 one hop 就到了

```
cpl in ~ λ sudo traceroute -T baidu.com
traceroute to baidu.com (220.181.38.148), 30 hops max, 60 byte packets
 1  220.181.38.148 (220.181.38.148)  36.867 ms  36.694 ms  36.656 ms
```

### same ip from different hops



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

第 5 跳是一个防火墙，第 5 跳回送包时做了 SNAT

补图

第二种：

```
 1  74.117.154.4 (74.117.154.4)  82.222 ms  81.909 ms  80.970 ms
 2  74.117.154.1 (74.117.154.1)  83.212 ms  83.725 ms  81.852 ms
 3  74.117.154.4 (74.117.154.4)  97.692 ms  81.136 ms *
 4  74.117.154.1 (74.117.154.1)  83.025 ms  82.698 ms  88.137 ms
```

路由指向错误，形成循环路由



### not end

```
Sample: tracert 10.77.87.1

Tracing route to 10.77.87.1 over a maximum of 30 hops:

1    <1 ms    <1 ms    <1 ms  pfsense.firewall.intern.org
  2    38 ms    44 ms    47 ms  10.77.87.1
  3    52 ms    31 ms    30 ms  10.77.87.1
  4    26 ms    43 ms    47 ms  10.77.87.1
  5    44 ms    55 ms    50 ms  10.77.87.1
  6    45 ms    36 ms    40 ms  10.77.87.1
  7    56 ms    56 ms    65 ms  10.77.87.1
  8    54 ms    42 ms    46 ms  10.77.87.1
```

在使用 traceroute 时，可能会出现在路径中已经出现了目的 IP ，但是还在发包

TODO

## 0x4 Packet analyze

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



