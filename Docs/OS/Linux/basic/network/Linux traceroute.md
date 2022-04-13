# Linux traceroute

ref:

https://www.imperva.com/learn/ddos/ip-fragmentation-attack-teardrop/



## 0x1Digest

traceroute 是 LInux 上的一个网络工具，显示从源到目的包的路径

### principle

traceroute 使用 IP 协议中 TTL 字段来实现，traceroute 开始探测时会发一个 ttl 值为 1 的，然后监听 nexthop 发回的 ICMP “time exceeded” 包（到达 nexthop 后仍不是目的 IP 包就会被丢弃，然后回送ICMP type 11），然后 ttl 值加 1，按照上述递归，直到回包是 ICMP “port unreachable”（默认使用UDP通常是30000以上的端口，如果端口没开就会回送 port unreachable） 或者是 TCP rest 或者是 hit max（ttl的值到了最大值，默认 30 hops）

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

### output

traceroute 默认会打印 3 个字段 ttl ，address of the gateway and round trip time（rtt）

#### asterisk

address of gateway 显示的是 gateway 回包路由的源接口（下图中的f0/1）， 一般发包路由和回包路由都相同，但是也有可能发包路由和回包路由不同

![2022-04-12_21-28](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220412/2022-04-12_21-28.1j5fzz3c2l9c.webp)

如果address of gateway 显示的是 asterisk（*），表示在指定时间内（默认5sec）没有从 gateway 收到回包，造成这种的原因通常有

1. 回包的链路中路由缺失，可以是回包链路中的任何一个节点
2. 回包的链路中有ACL，可以是回包链路中的任何一个节点
3. 当前大多数 firewall 都会过滤 UDP 端口，甚至是ICMP，碰到这种情况可以使用其他协议(TCP)来绕过 firewall

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