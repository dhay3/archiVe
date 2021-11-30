# Route

reference：

https://www.jannet.hk/routing-decision-zh-hant/

当一个packet从一个incoming interface 到达 router，router是如何处理的呢？当然它希望尽快吧packet放到适当的outgoing interface送出去。不过router 可能同时接驳多个interface，所以router需要查看自己的route table去决定把packet放到哪一个interface送走。

## Route Type

- 静态路由：不能自动适应网络拓扑的变化，需要手动配置
- 动态路由：能够自动适应网络拓扑结构，占用一定的网络资源

### dynamic route

#### 根据范围可以分为

- 内部网关协议(Interior Gateway Protocol, IGP)：在一个自治系统(Autonomous system, AS)内部运行，常见的IGP协议包括RIP，OSPF和IS-IS
- 外部网关协议(Exterior Gateway Protocol, EGP)：运行于不同AS之间，BGP是目前最常用的EGP协议

#### 根据使用的算法可以分为

- 距离矢量协议(Distance-Vector)：包括RIP和BGP。其中BGP也称为路径矢量协议(Path-Vector)
- 链路状态协议(Link-State)：包括OSPF和IS-IS

#### 根据目的地址类型可以分为

- 单播路由协议(Unicast Routing Protocol)：包括RIP、OSPF、BGP和IS-IS
- 组播路由协议(Multicast Routing Protocol)：包括PIM-SM

## Route Table

当一个packet到达，router会立刻检查packet的L3 destination address（即Destination IP）, 然后在route table中找到一条最适合的记录，根据记录，把packet放到对应的interface送走，这就是router每分每秒的工作。要看route table，可以用`show ip route`

```
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     2.0.0.0/32 is subnetted, 1 subnets
S       2.2.2.2 [1/0] via 192.168.30.1, Ethernet0/0
                is directly connected, Ethernet0/0
     192.168.1.0/24 is variably subnetted, 4 subnets, 4 masks
S       192.168.1.32/29 [1/0] via 192.168.12.12
S       192.168.1.32/28 [1/0] via 192.168.12.12
S       192.168.1.0/26 [1/0] via 192.168.12.11
S       192.168.1.0/24 [1/0] via 192.168.12.2
D    192.168.2.0/23 [90/409600] via 192.168.12.2, 01:21:55, Ethernet0/0
```

## Route priority

不同的协议具有不同的优先级，值越小优先级越高

| 路由协议或路由种类 | 相应优先级 |
| ------------------ | ---------- |
| DIRECT             | 0          |
| OSPF               | 10         |
| IS-IS              | 15         |
| STATIC             | 60         |
| RIP                | 100        |
| OSPF ASE           | 150        |
| OSPF NSSA          | 150        |
| IBGP               | 255        |
| EBGP               | 255        |

除直联路由DIRECT外，各种路由协议的优先级都可以有用户手工配置

## Route Table 的来源

Route Table基本会用以下三个来源去建立自己的route table

### interface network

router会收集每一个使用中的interface，如果interface有配置IP的话，这个network就会被放进route table。

```
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set
```

然后我们在interface设定IP address 并把interface启动，route table就会多一条route

```
R1#configure terminal
Enter configuration commands, one per line.  End with CNTL/Z.
R1(config)#int ethernet 0/0
R1(config-if)#ip address 192.168.12.1 255.255.255.0
R1(config-if)#no shutdown
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set
C    192.168.12.0/24
```

==其中C就表示直连在当前router上的network route==

### static route

如果Router configuration 中有配置static route的话，也会被放进route table

```
R1#configure terminal
Enter configuration commands, one per line.  End with CNTL/Z.
R1(config)#ip route 192.168.1.0 255.255.255.0 192.168.12.2
R1(config)#end
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.1.0/24 [1/0] via 192.168.12.2
```

==其中S就表示配置的static route==

### dynamic routing protocol

如果router有配置dynamic routing protocol，router与router之间就会交换route信息，dynamic routing protocol 会对这些自选进行运算，最后把结果放进route table。

现在我们在router上设定EIGRP，假设interface另一端的router已经设定好了EIGRP。完成设定后便可看到有dynamic routing protocol产生的route

```
R1(config)#router eigrp 1
R1(config-router)#network 192.168.12.0
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.1.0/24 [1/0] via 192.168.12.2
D    192.168.2.0/24 [90/409600] via 192.168.12.2
```

## route table 输出

以EIGRP所产生的router做例子

```
R1(config)#router eigrp 1
R1(config-router)#network 192.168.12.0
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.1.0/24 [1/0] via 192.168.12.2
D    192.168.2.0/24 [90/409600] via 192.168.12.2
```

- D - Codes

  route的来源，自行查看注释

- 192.168.1.0/24 - Network

  就是不同network的network id 和 prefix，router就是在这里查看目的地IP符合那一条route。如果IP是落在某一个Network的，就使用该条Route的资料去处理packet

- [90/409600] - AD/Metric

  administrative distance(AD) 和 Metric

- 192.168.12.2 - Next Hop

  next hop 就是 packet要到的下一站，可以是ip address或是interface。只要是这条route的packet都会被送到next hop。所以一个packet到达R1，R1就会先查看这个packet的目的地IP。假设目的地是192.168.1.150，router就会到route table里做route lookup，找到`S    192.168.1.0/24 [1/0] via 192.168.12.2`这条route符合条件。于是route便会把packet送到next hop 192.168.12.2。==但是router不知道192.168.12.2在哪里，可以再做一次route lookup 查看route table，根据192.168.12.2找到`C    192.168.12.0/24 is directly connected, Ethernet0/0`这条route 符合条件==于是就把packet从Ethernet0/0哪里送出去了。第二次的route lookup又被称为recursive lookup

## route 选路

> 匹配prefix最长，再比较AD，然后比较Metric

### AD/Metric

从不同来源收集回来的route有可能会出现相同的network，这时router就只会把administrative distance （AD）较小的Route放进route table。如果AD相同的话，Route Table 通常只会把metric最小的route放进route table，除非routing protocol有设定load balance

加入一条AD小于90的192.168.2.0/24 Static Route 会比EIGRP Route优先。例如加入一条AD 50 的static route

```
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.1.0/24 [1/0] via 192.168.12.2
D    192.168.2.0/24 [90/409600] via 192.168.12.2, 00:01:31, Ethernet0/0
R1#configure terminal
R1(config)#ip route 192.168.2.0 255.255.255.0 192.168.12.10 50
R1(config)#end
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.1.0/24 [1/0] via 192.168.12.2
S    192.168.2.0/24 [50/0] via 192.168.12.10
```

==Metric需要结合，不同的routing protocol==

### Subnetted Route 与 Longest Match

有时候network会被切分成subnet，如果route table收到这些subnet route的话，会用以下方式表示：

```
R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     192.168.1.0/24 is variably subnetted, 4 subnets, 4 masks
S       192.168.1.32/29 [1/0] via 192.168.12.12
S       192.168.1.32/28 [1/0] via 192.168.12.12
S       192.168.1.0/26 [1/0] via 192.168.12.11
S       192.168.1.0/24 [1/0] via 192.168.12.2
```

意思是收到了4个subnet都存在于192.168.1.0/24这个class c network中，分别有：192.168.1.32/29、192.168.1.32/28、192.168.1.0/26、192.168.1.0/24。==每条的subnet route的next hop可以相同，也可以不一样。==

现在有destination ip 是192.168.1.33的packet，四条route都符合。Router会用lognest match去选择，==即找符合的route中prefix最长的route来使用==。所以192.168.1.33会去往192.168.1.32/29，router会把packet送到next hop 192.168.12.12

## static rotue/next hop 设定

static route可以使用以下三种next hop设定

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210823/Snipaste_2021-08-11_20-17-25.25wgue9okexs.png)

### next hop IP

```
R1#show run | i route
ip route 192.168.101.0 255.255.255.0 192.168.12.2
ip route 192.168.102.0 255.255.255.0 192.168.12.2
ip route 192.168.103.0 255.255.255.0 192.168.12.2
ip route 192.168.104.0 255.255.255.0 192.168.12.2

R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.101.0/24 [1/0] via 192.168.12.2
S    192.168.102.0/24 [1/0] via 192.168.12.2
S    192.168.103.0/24 [1/0] via 192.168.12.2
S    192.168.104.0/24 [1/0] via 192.168.12.2
```

假设R1要送packet到192.168.101.1，根据route table，他会送到next hop 192.168.12.2。但他还不知道192.168.12.2在哪里，所以他要做一次recursive lookup，然后找到192.168.12.0/24的next hop interface 是 E0/0。

然后E0/0发一个arp request 问 192.167.12.2的MAC地址，于是R2回应并发arp reply，使R1得知MAC地址并记录在arp table中

之后不管destination是192.168.101.1、192.168.102.1、192.168.103.1、192.168.104.1，都不需要发arp了，因为next hop 192.168.12.2已经在arp table中了

```
R1#ping 192.168.101.1

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.101.1, timeout is 2 seconds:
.!!!!
Success rate is 80 percent (4/5), round-trip min/avg/max = 8/35/117 ms
R1#show arp
Protocol  Address          Age (min)  Hardware Addr   Type   Interface
Internet  192.168.12.1            -   cc01.064c.0000  ARPA   Ethernet0/0
Internet  192.168.12.2          101   cc02.065b.0000  ARPA   Ethernet0/0
```

cisco有些型号使用`show ip arp`来查看

### next hop interface

用interface来做next hop，有点不同。假设R1要送packet到192.168.101.1，根据route table，发现next hop 是 interface E0/0，没有next hop IP。router会用arp request在E0/0 问192.168.101.1的MAC address，虽然R2的E0/0并非192.168.101.1，==但R2的E0/0会覆盖自己的MAC地址==，原因是interface的proxy arp 预设是enable的，如果是R2的route table有192.168.101.1的话，proxy arp 设定会使R2的E0/0回复arp reply

```
R1#show run | i route
ip route 192.168.101.0 255.255.255.0 Ethernet0/0
ip route 192.168.102.0 255.255.255.0 Ethernet0/0
ip route 192.168.103.0 255.255.255.0 Ethernet0/0
ip route 192.168.104.0 255.255.255.0 Ethernet0/0


R1#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
S    192.168.101.0/24 is directly connected, Ethernet0/0
S    192.168.102.0/24 is directly connected, Ethernet0/0
S    192.168.103.0/24 is directly connected, Ethernet0/0
S    192.168.104.0/24 is directly connected, Ethernet0/0

R1#ping 192.168.101.1

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.101.1, timeout is 2 seconds:
.!!!!
Success rate is 80 percent (4/5), round-trip min/avg/max = 8/36/116 ms
R1#show arp
Protocol  Address          Age (min)  Hardware Addr   Type   Interface
```

==按照这种机制，也就是说R1需要为每一个destination IP 查询arp，这样arp record就会迅速增加，大大影响router的效能，因此避免使用这种static route设定==

```
R1#show arp
Protocol  Address          Age (min)  Hardware Addr   Type   InterfaceInternet  192.168.104.1           0   cc02.065b.0000  ARPA   Ethernet0/0
Internet  192.168.101.1          12   cc02.065b.0000  ARPA   Ethernet0/0
Internet  192.168.103.1           0   cc02.065b.0000  ARPA   Ethernet0/0
Internet  192.168.102.1           0   cc02.065b.0000  ARPA   Ethernet0/0
Internet  192.168.12.1            -   cc01.064c.0000  ARPA   Ethernet0/0
Internet  192.168.12.2          124   cc02.065b.0000  ARPA   Ethernet0/0
```

### next hop interface and IP

如果必须使用next hop interace来限制packet送到某个outgoing interface的话，最好把next hop ip也写进去，目的是避免arp record记录多个destination IP

```
R1#show run | i route
ip route 192.168.101.0 255.255.255.0 Ethernet0/0 192.168.12.2
ip route 192.168.102.0 255.255.255.0 Ethernet0/0 192.168.12.2
ip route 192.168.103.0 255.255.255.0 Ethernet0/0 192.168.12.2
ip route 192.168.104.0 255.255.255.0 Ethernet0/0 192.168.12.2
```

packet必定会使用E0/0作为outgoing interface 而next hop ip则使用192.168.12.2。==但是如果E0/0 Down了的话，所有以此interface作为出口的route就会消失==

## Cisco Express Forwarding(CEF)

上面这些方法有什么问题呢？就是不够快，如果每次有packet进来，都要从route table找route，而且可能不止一次（recursive lookup），但如果route table没有更新的话，基本上每次结果都一样，为了省掉rotue lookup的程序，就出现了CEF

CEF可以把route lookup结果保存在CEF table，当router再次收到同一个destination IP的packet，不用找route table了，按CEF table的记录直接转走（switching）。所以CEF所做的是所谓的route once switch many。当route table有更新是，CEF table也会被更新。

用`show ip cef`可以查看cef table

```
R1#show ip cef
Prefix              Next Hop             Interface
0.0.0.0/0           drop                 Null0 (default route handler entry)
0.0.0.0/32          receive
192.168.2.0/23      192.168.12.2         Ethernet0/0
192.168.12.0/24     attached             Ethernet0/0
192.168.12.0/32     receive
192.168.12.1/32     receive
192.168.12.2/32     192.168.12.2         Ethernet0/0
192.168.12.255/32   receive
192.168.101.0/24    192.168.12.2         Ethernet0/0
192.168.102.0/24    192.168.12.2         Ethernet0/0
192.168.103.0/24    192.168.12.2         Ethernet0/0
192.168.104.0/24    192.168.12.2         Ethernet0/0
224.0.0.0/4         drop
224.0.0.0/24        receive
255.255.255.255/32  receive
```

## Policy Based Routing(PBR)

基于策略的路由

