# Day25 - RIP & EIGRP

## RIP

Routing Information Protocol(RIP) 是基于 Distance vetcor IGP(uses routing-by-rumor logic to learn/share routes) 公有的协议

- Uses hop count as its metric. One router = one hop(bandwitdth is irrelevant)

- The maximum hop count is 15(anything more than that is considered unreachable)

  > 显然的因为限制 15 hops，所以 RIP 不适用大型的网络拓扑。实际上在现实生产的环境中，以及不再使用 RIP

- Has three versions

  RIPv1 and RIPv2, used for IPv4

  RIPng(RIP Next Generation), used for IPv6

- Uses two message types

  Request: To ask RIP-enabled neighbor routers to send their routing table

  Repopnse: To send the local router’s routing table to neighboring routers

### RIPv1 VS RIPv2

#### RIPv1

- Only advertise classfull addresses(Class A/B/C)

- Doesn’t support VLSM,CIDR
- Doesn’t include subnet mask information in advertisements(Response messages)

> RIPv1 发送 Reponse 时不会携带 subnetmask，接收方会按照 IP 处于 A/B/C 类来处理路由

例如

10.1.1.0/24 会变成 10.0.0.0/8 A 类地址

172.16.192.0/18 会变成 172.16.0.0/16 B 类地址

192.168.1.4/30 会变成 192.168.1.0/24 C 类地址

- Messages are broadcast to 255.255.255.255

  因为广播到 3 层地址，所以所有的 host 都能收到 RIP messages

#### RIPv2

- Supports VLSM, CIDR
- Includes subnet mask information in advertisemtns
- Messages are multicast to 224.0.0.9

*Broadcast messages are delivered to all devices on the local network*

*Multicast messages are delivered only to devices that have joined that specific multicast group*

> multicast 并不在 CCNA 考试范围内

### RIP Configuration

假设有如下拓扑，使用 RIP 来配置路由

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_13-54.7l50xlev603k.webp)

R1 配置如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_13-55.1bx09pqg5l6o.webp)

- `router rip`

  声明 router 使用 RIP 协议

- `version 2`

  声明 router 使用 RIPv2

- `no auto-summary`

  auto-summary 是自动会advertise classfull addresses，例如 172.16.1.0/28 就会变为 172.16.0.0/16。==RIP 会默认使用 auto-smmary==，为了可以使用 VLSM 和 CIDR 要使用 no

- `network 10.0.0.0`

  `network 172.16.0.0`

  network 命令有主要如下几个功能和规则

  - look for interfaces with an IP address that is in the specified range
  - active RIP on the interfaces that fall in the range(在范围内)
  - form adjacencies with connected RIP neighbors
  - **advertise the network prefix of the interface(NOT the prefix in the network command)**

  > *network 10.0.12.0 it will be converted to network 10.0.0.0(a class A network)*
  >
  > *There is no need to enter the network mask*
  >
  > 需要注意的一点是，network 命令不会考虑 mask，如果输入的是 10.0.12.0 会自动识别成 10.0.0.0/8
  >
  > OSPF 或者是 EIGRP 中的 network 命令和 RIP 中的 network 命令逻辑相同

  例如图中的 `network 10.0.0.0` 表示

  1. 因为 network 命令是 classful 的，所以 10.0.0.0 会被认为 10.0.0.0/8
  2. R1 会寻找所有在 10.0.0.0/8 范围内的 interfaces
  3. 例如 10.0.12.1 和 10.0.13.1 都在 10.0.0.0/8 范围内，所以 R1 G0/0 和 G1/0 都会使用 RIP
  4. R1 和 R2 以及 R3 会形成 adjacencies 关系
  5. R1 会向他的 RIP neighbors(R2 和 R3) advertise 10.0.12.0/30 和 10.0.13.0/30(Not 10.0.0.0/8)

  > *The network command doesn’t tell the router which networks to advertise. It tells the router which interfaces to activate RIP on, and then the router will advertise network of those interfaces*

  同样的例如 `network 172.16.0.0` 表示

  1. 因为 network 命令是 classful 的，所以 172.16.0.0 会被认为 172.16.0.0/16
  2. R1 会寻找所以在 172.16.0.0/16 范围内的 interfaces
  3. 例如 172.16.1.14 在 172.16.0.0/16 范围内，所以 R1 G2/0 会使用 RIP
  4. 因为没有和 R1 G2/0 互联的 neighbors，所以就不会和其他 Router 形成 adjacencies 关系
  5. R1 会向他的 RIP neighbors(这里 R2 和 R3 还是 R1 的 neighbors，具体看下面解释) advertise 172.16.1.0/28(Not 172.16.0.0/16)

  > *Although there are no RIP neighbors connected to G2/0， R1 will continuously send RIP advertisements out of G2/0. This is unnecessarty traffic, so G2/0 should be configured as a passive interface*

  可以使用 `passive-interface <interface-id>` 来使接口变为 passive interface。==这个是 RIP 中的一部分，并不是端口自己本身的状态==

  1. The passive-interface command tells the router to stop sending RIP advertisements out of the specified interface(G2/0)
  2. *However, the router will continue to advertise the network prefix to the interface(172.16.1.0/28) to its RIP neighbors(R2,R3)*
  3. You should always use the command on interfaces which don’t have any RIP neighbors
  4. EIGRP and OSPF both have the same passive interface functionality, using the same command 

#### Advertise a default route into RIP

default route 比较特殊

例如下图，为 R1 添加了一条默认路由 via 203.0.113.2

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_18-29.a8v25855xsg.webp)

现在需要使用 RIP 告诉 R2/R3/R4 这条默认路由，可以使用 `default-information originate` 命令来宣告默认路由

使用上述命令后，那么 R4 的 routing table 如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_18-31.s5eoy2u53b4.webp)

这里 R4 会添加两条往 0.0.0.0 的路由，即使 R3 - R4 和 R1 - R2 中间链路的带宽不一样，因为 RIP metric 只会考虑 hop，并不考虑链路带宽

可以使用 `show ip protocols` 来查看当前使用的 dynamic routing

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-06.7bnnfzavhdhc.webp)

还可以使用 `maximum-paths <number>` 来修改，ECMP 最多拥有的路由条数

或者是使用 `distance <ad>` 来修改当前协议的 AD 值

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-06_1.4bzwsrn2q6ww.webp)

## EIGRP

Enhanced Interior Gateway Routing Protocol(EIGRP) 是思科独有的协议

- Much faster than RIP in reacting to changes in the network

- Does not have the 15 ‘hop-count’ limit of RIP

- Sends messages using multicast address 244.0.0.10

  > RIP 使用的多播地址为 244.0.0.9

- Is the only IGP that can perform unequal-cost load-balancing(by default it performs ECMP load-balancing over 4 paths like RIP)

### EIGRP Configuration

假设有如下拓扑，使用 EIGRP 来配置路由

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_13-54.7l50xlev603k.webp)

R1 配置如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-16.3cc1egkp3oxs.webp)

- `router eigrp 1`

  声明使用 eigrp，AS number 使用 1

  > *The As(autonomous system) number must match between routers, or they will not form an adjacency and share route information*
  >
  > 所以在 EIGRP 中的 router 必须使用相同的 AS number，例如上图中如果想要 EIGRP 正常运行，R1/R2/R3/R4 使用的 AS number 就必须要一样

- `no auto-summary`

  和 RIP 一样，EIGRP advertise classful addresses routes。但是 EIGRP 默认，会根据 Router 的型号来决定是否启用 auto-summary，最好是手动声明不使用 auto-summary

- `passive-interface g2/0`

  和 RIP 逻辑一样

- `network 10.0.0.0`

  `network 172.16.1.0 0.0.0.15`

  network 命令和 RIP 类似，如果没有声明 mask

  但是 EIGRP 中的 mask 并不是标准的 subnet mask, 而是 wildcard mask

  > - A wildcard mask is basically an ‘inverted’ subnet mask
  > - All 1s in the subnet mask are 0 in the equivalent wildcard mask. All 0s in the subnet mask are 1 in the equivalent wildcard mask
  >
  > 只要记住 wildcard mask 部分就是主机位即可
  >
  > 例如 mask 255.255.255.0 对应的 wildcast mask 为 0.0.0.255
  >
  > *A shortcut is to subtract each octet of the subnet mask from 255*
  >
  > - ‘0’ in the wildcard mask = must match
  > - ‘1’ in the wildcard mask = don’t have to match

  例如下图

  > 可以使用加 wildcard mask 的方式

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-32.7c095z6ns7ls.webp)

  如果接口是 172.16.1.14，使用了 `network 172.16.1.0 0.0.0.15`,因为 172.16.1.14 在 172.16.1.15 范围内那么 172.16.1.14 对应的接口会自动加入 EIGRP

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-34.p9mkoxfemz4.webp)

  如果使用了 `network 172.16.1.0 0.0.0.7` 就不会把 172.16.1.14 对应的接口加入到 EIGRP，因为 172.16.1.14 不在 172.16.1.7 内

同样的和 RIP 一样，也可以使用 `show ip protocols` 来查看 EIGRP 相关的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-46.ixs4nokxmq8.webp)

这里可以看到一个比较特殊的东西 Router ID，按照下面顺序的因素来选择具体使用那个 Router ID

1. Manual configuration

   可以使用 `eigrp router-id <route-id>` 来手动配置 EIGRP router id

2. Highest IP address on a loopback interface

3. Highest Ip address on a physical interface

如果 route 是通过 EIGRP 学来的，会以 D 来标识

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_21-52.coeauung07s.webp)

其中的 metric 部分明显是大于 RIP 的，因为 EIGRP 衡量 metric 的方式和 RIP 不同

### EIGRP metric

EIGRP 默认会使用 bandwidth 和 delay 作为因子，计算公式如下
$$
([K1*bandwidth + (K2*bandwidth)/(256 - load) + K3 * delay] * [K5/(reliability + K4)]) * 256
$$
默认 K1 = 1, K2 = 0, K3 = 1, K4 = 0, K5 = 0

> 只要记住 EIGRP 的因子为 bandwith 和 delay 即可
>
> metric = bandwidth + delay

例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_22-58.2ry6pz8wckqo.webp)

R1 想要访问 PC1, 因为 via R2 的路由是 preferred route，所以 metric 就是 R2 F1/0 - R4 F1/0 bandwidth(因为其他链路互联的接口均为 gigabit, 而 R2 f1/0 和 R4  f1/0 是 fastethernet) 加上 R1 to PC1 via R2 所有链路的时延(这里的时延并不是通过 ICMP 来计算的，实际的 delay 是根据 interface bandwidth 来决定的)

在 EIGRP 中还有两个名词，关联 metric

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_23-04.3cu966ufds1s.webp)

红色就是 Feasible Distance(FD)

蓝色就是 Reported Distance(RD)

如果通过 `show ip eigrp topology` 查看所有的 eigrp route

```
P 192.168.4.0/24, 1 successors, FD is 28672
         via 10.0.12.2 (28672/28416), GigabitEthernet0/0
         via 10.0.13.2 (30976/28416), FastEthernet1/0
```

`(28672/28416)` 左边部分的 28672 就是 feasible distance，右边部分的 28416 就是 reported distance。同理 `(30976/28416)`

了解 feasible distance 和 reported distance 是为了了解另外两个名词

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_23-09.5u6aarf584xs.webp)

> feasible successor 需要满足一个条件，就是 feasible successor 的 reported distance 要比 successor 的 feasible distance 小
>
> 主要是为了防环（不知道为啥）

例如上图因为 via 10.0.12.2 左边部分 feasible distance 比 via 10.0.13.2 的小，所以 via 10.0.12.2 的 route 是 best route，即是 successor，所以 via 10.0.13.2 就是 feasible successor

### EIGRP Unequal-Cost Load-Balancing

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_23-21.3ryd7gdx2zgg.webp)

在 EIGRP 中是通过 variance 值来改变 load balancing 方式的，如果 variance 值为 1 就表示只是用 ECMP (默认使用 variance 1)

所以 `show ip eigrp topology` 中显示的两条路由，只会有一条 route 加入到 routing table 中，因为 feasible distance 不一样

所以我们可以通过 `variance 2` 来改变 EIGRP 的 load balancing 方式，使用 unequal-cost load balancing

> 对需要添加 unequal-cost load balancing 的设备使用

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_23-26.75fqsbitw4g0.webp)

例如上图

via 10.0.12.2 是 successor 

via 10.0.13.2 是 feasible successor

如果 feasible successor 的 feasible distance 比 successor 2 倍的 feasible distance 小，就会被用为 unequal load balancing

*EIGRP will only perform unequal-cost load-balancing over feasible successor routes. If a route doesn’t meet the feasibility requirement, it will NEVER be selected for load-balancing, regardless of the variance*

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230609/2023-06-12_22-34.7c152gm6aha8.webp)

### 0x01

Configure the appropriate hostnames and IP addresses on each device. Enable router interfaces

#### R1

```
Router(config)#hostname R1
R1(config)#int g0/0
R1(config-if)#ip add 10.0.12.1 255.255.255.252
R1(config-if)#no shutdown
R1(config-if)#int f1/0
R1(config-if)#ip add 10.0.13.1 255.255.255.252
R1(config-if)#no shutdown
```

#### R2

```
Router(config)#hostname R2
R2(config)#int g0/0
R2(config-if)#ip add 10.0.12.2 255.255.255.252
R2(config-if)#no shutdown
R2(config-if)#int f1/0
R2(config-if)#ip add 10.0.24.1 255.255.255.252
R1(config-if)#no shutdown
```

#### R3

```
Router(config)#hostname R3
R3(config)#int f1/0
R3(config-if)#ip add 10.0.13.2 255.255.255.252
R3(config-if)#no shutdown
R3(config-if)#int f2/0
R3(config-if)#ip add 10.0.34.1 255.255.255.252
R3(config-if)#no shutdown
```

#### R4

```
Router(config)#hostname R4
R4(config)#int f2/0
R4(config-if)#ip add 10.0.34.2 255.255.255.252
R4(config-if)#no shutdown
R4(config-if)#int f1/0
R4(config-if)#ip add 10.0.24.2 255.255.255.252
R4(config-if)#no shutdown
R4(config-if)#int g0/0
R4(config-if)#ip add 192.168.4.254 255.255.255.0
R4(config-if)#no shutdown
```

### 0x02

Configure a loopback interface on each router(1.1.1.1/32 for R1, 2.2.2.2/32 for R2, etc)

```
R1(config-if)#int lo0
R1(config-if)#ip add 1.1.1.1 255.255.255.255

R2(config-if)#int lo0
R2(config-if)#ip add 2.2.2.2 255.255.255.255

R3(config-if)#int lo0
R3(config-if)#ip add 3.3.3.3 255.255.255.255

R4(config-if)#int lo0
R4(config-if)#ip add 4.4.4.4 255.255.255.255
```

> 这里并不需要执行 `no shutdown`, loopback 通常都是 up/up 的，除非手动使用了 `shutdown` 命令

### 0x03

Configure EIGRP on each router

> 在没有配置 EIGRP 前 `show ip protocols` 显示为空

```
R1(config)#router eigrp 100
R2(config)#router eigrp 100
R3(config)#router eigrp 100
R4(config)#router eigrp 100
```

disable auto-summary

```
R1(config-router)#no auto-summary 
R2(config-router)#no auto-summary 
R3(config-router)#no auto-summary 
R4(config-router)#no auto-summary 
```

> 如果 `show ip protocols` 中显示
>
> Automatic network summarization is not in effect  
>
> 就标识 auto-summary 关闭了

Enable EIGRP on each interface(including loopback interfaces)

```
R1(config-router)#network 0.0.0.0 255.255.255.255
R2(config-router)#network 0.0.0.0 255.255.255.255
R3(config-router)#network 0.0.0.0 255.255.255.255
R4(config-router)#network 0.0.0.0 255.255.255.255
```

Configure passive interfaces where appropriate(including loopback interface)

> loopback interfaces 也会发送 EIGRP messages，就像和其他 Router 接口互联一样，所以需要对 loopback interfaces 做 passive-interfaces 的操作

```
R1(config-router)#passive-interface lo0
R2(config-router)#passive-interface lo0
R3(config-router)#passive-interface lo0
R4(config-router)#passive-interface lo0
R4(config-router)#passive-interface g0/0
```

这里 g0/0 同样需要是 passive interface，因为没有和 router 互联，所以就没有并要发送 EIGRP messages

同样也也可以使用 `show ip prtocols` 来查看 router 当前的 passive interfaces

```
R4(config-router)#do sh ip pro

Routing Protocol is "eigrp  1 " 
  Outgoing update filter list for all interfaces is not set 
  Incoming update filter list for all interfaces is not set 
  Default networks flagged in outgoing updates  
  Default networks accepted from incoming updates 
  EIGRP metric weight K1=1, K2=0, K3=1, K4=0, K5=0
  EIGRP maximum hopcount 100
  EIGRP maximum metric variance 1
Redistributing: eigrp 1
  Automatic network summarization is not in effect  
  Maximum path: 4
  Routing for Networks:  
     0.0.0.0
  Passive Interface(s): 
    GigabitEthernet0/0
    Loopback0
  Routing Information Sources:  
    Gateway         Distance      Last Update 
    10.0.34.1       90            1515203    
    10.0.24.1       90            1515203    
  Distance: internal 90 external 170
```

还可以使用 `show ip eigrp neighbors` 来查看当前 router 的邻居

```
R4(config-router)#do sh ip eigrp ne
IP-EIGRP neighbors for process 1
H   Address         Interface      Hold Uptime    SRTT   RTO   Q   Seq
                                   (sec)          (ms)        Cnt  Num
0   10.0.34.1       Fa2/0          11   00:18:04  40     1000  0   16
1   10.0.24.1       Fa1/0          14   00:18:04  40     1000  0   16
```

还可以使用 `show ip eigrp topology` 来查看 eigrp 收到的所有路由(包括不是最优的路由，没有出现在 routing table 中)

```
R1(config-router)#do sh ip eigrp to
IP-EIGRP Topology Table for AS 1/ID(1.1.1.1)

Codes: P - Passive, A - Active, U - Update, Q - Query, R - Reply,
       r - Reply status

P 1.1.1.1/32, 1 successors, FD is 128256
         via Connected, Loopback0
P 2.2.2.2/32, 1 successors, FD is 130816
         via 10.0.12.2 (130816/128256), GigabitEthernet0/0
P 3.3.3.3/32, 1 successors, FD is 156160
         via 10.0.13.2 (156160/128256), FastEthernet1/0
P 4.4.4.4/32, 1 successors, FD is 156416
         via 10.0.12.2 (156416/156160), GigabitEthernet0/0
         via 10.0.13.2 (158720/156160), FastEthernet1/0
P 10.0.12.0/30, 1 successors, FD is 2816
         via Connected, GigabitEthernet0/0
P 10.0.13.0/30, 1 successors, FD is 28160
         via Connected, FastEthernet1/0
P 10.0.24.0/30, 1 successors, FD is 28416
         via 10.0.12.2 (28416/28160), GigabitEthernet0/0
P 10.0.34.0/30, 1 successors, FD is 30720
         via 10.0.13.2 (30720/28160), FastEthernet1/0
P 192.168.4.0/24, 1 successors, FD is 28672
         via 10.0.12.2 (28672/28416), GigabitEthernet0/0
         via 10.0.13.2 (30976/28416), FastEthernet1/0
```

### passive-interface lo00x04

Configure R1 to perform unequal-cost load-balancing when sending network traffic to 192.168.4.0/24

```
R1(config-router)#variance 2
```

可以使用 `show ip eigrp topology` 和 ff `show ip route` 来校验

```
#show ip eigrp topology
P 192.168.4.0/24, 2 successors, FD is 28672
         via 10.0.12.2 (28672/28416), GigabitEthernet0/0
         via 10.0.13.2 (30976/28416), FastEthernet1/0
         
#show ip route
D    192.168.4.0/24 [90/28672] via 10.0.12.2, 00:00:04, GigabitEthernet0/0
                    [90/30976] via 10.0.13.2, 00:00:04, FastEthernet1/0
```

虽然 via 10.0.12.2 和 via 10.0.13.2 的 feasible distance 不一样，且 via 10.0.12.2 的优，但是两条路由都添加到了 routing table 中，观察发现 metric 并不一样

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=N8PiZDld6Zc&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=47