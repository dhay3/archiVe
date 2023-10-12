# Day26 - OSPF Part 1

## Link State Routing Protocol

OSPF 是 Link State Routing Protocl 中的一种协议

- When using a link state routing protocol. every router creates a ‘connectivity map’(就是数据结构中的图) of the network
- To allow this, each router advertises information about its interfaces(connected networks) to its neighbors. These advertisements are passed along to other routers, until all routers in the network develop the same map of the network
- Each router independently uses this map to calculate the best route to each destination
- Link state protocols use more resources(CPU) on the router, because more information is shared
- However, link state protocols tend to be faster in reacting to changes in the network than distance vector protocols

## OSPF

Open Shortest Path First(OSPF) 使用了 Shortes Path First 算法(由 Edsger Dijkstra 中文翻译为 狄杰斯特拉)

OSPF 一共有 3 个版本

1. OSPFv1(1989): OLD, not in use anymore
2. OSPFv2(1998): Used for IPv4 
3. OSPFv3(2008): Used for IPv6(can also be used for IPv4, but usually v2 is used)

*Routers store information about the network in LSAs(Link State Adertisements), which are organized in a structure called the LSDB(Link State Database)*

*Routers will flood LSAs until all routers in the OSPF area develop the same map of the network(LSDB)*

> 这里的 flood 指的是往 OSPF neighbors 发送 LSAs

以如下拓扑为例

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_12-19.y7eo1p32bg0.webp)

假设 R1/R2/R3/R4 都运行了 OSPF，互相都是 OSPF neighbors，有一样的 LSDB(Link state database)

1. OSPF is enabled on R4’s G1/0 interface

   所以 R4 需要宣告 R4 G1/0 端口所在的网段

2. R4 creates an LSA to tell its neighbors about the network on G1/0

   LSA 大概会包含如下几个字段

   - RID: 4.4.4.4 

     Router ID

   - IP: 192.168.4.0/24

   - Cost: 1

   同时 LSA 还有老化时间

   *Each LSA has an aging timer(30 min by default). The LSA will be flooded again after the timer expires*

3. The LSA is flooded throughout the network until all routers have received it

   flood 方式如下图

   ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_12-41.6clxlrzpnrwg.webp)

4. This results in all routers sharing the same LSDB

   因为所有的 Routers 都会往 neighbors foold LSAs

5. Each router then uses the SPF(Shortest Path First) alorithm to calculate its best route to 192.168.4.0/24

逻辑上可以分为 3 大步

1. Become neighbors with other routers connected to the same segment
2. Exchange LSAs with neighbor routers
3. Calculate the best routees to each destination, and insesrt them into the routing table

## OSPF Areas

OSPF 使用 areas 的逻辑概念来划分 network

> *An area is a set of routers and links that share the same LSDB*

- Small networks can be single-area without any negative effects on performance

  例如开头的例子的拓扑，就是一个小型的 network，4 台 routers 可以在一个 OSPF area 中，也不会出现 degradation(降级) of network performance

- In larger networks, a single-area design can have negative effects

  例如 一个由 500 台 router 组成一个 OSPF area，在 area 中有 1000 个 subnets。可能会有如下的一些负面影响

  1. the SPF algorithm takes more time to calculate routes

  2. the SPF algorithm requires exponentially more processing power on the router

     简单说就是耗电。。

  3. the larger LSDB takes up more memory on the routers

  4. any small change in the network caueses every router to flood LSAs and run the SPF algorithm again

     例如就是打开了一个端口，端口 IP 正好在 OSPF network 中，那个对应的 LSA 就会在所有的 Router 之间 flood

  *By dividing a large OSPF network into several smaller areas, you can avoid the above negative effects*

  > 在大型网络拓扑中，尽量多划分 OSPF areas

例如下图，将所有的 routes 都划分在了以 OSPF Area 0(也被称为 Backbone area)

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-03.409axfv6g4u8.webp)

通常不会这么干，而是会按照下图划分 OSPF Area

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-05.1yqla4lgv7q8.webp)

Each Area has unique LSDB

### Backbond area

Area0 = Backbond area

*The backbone area(area 0) is an area that all other areas must connect to*

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-05.1yqla4lgv7q8.webp)

例如上面的例子，Area1/Area2/Area3 必须要和 Area0 互联才可以组成 OSPF network

下图的例子是不会也不允许出现在 OSFP 中

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-10.6u1cgjvp1gn4.webp)

因为 Area1 没有和 Area0 互联，而是通过 Area2 和 Area0 互联

## Terms in OSPF

- Routers with all interfaces in the same area are called **internal routers**

  例如下图中红框部分都是 internal routers

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-13.5bmb0xoo3mgw.webp)

- Routers with interfaces in multiple areas are called **area border routers**(ABRs)

  例如下图中红框部分都是 ABR

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-15.56wget31yry8.webp)

  > *ABRs maintain a separate LSDB for each area they are connected to. It is recommend that you connect an ABR to a maximum of 2 areas. Connecting an ABR to 3+ areas can overburden the router*
  >
  > ABR 最好只互联 2 个 OSPF Area

- Routers connected to the backbon area(area 0) all called **backbone router**

  例如下图中红框部分就是 backbond router 同时也是 internal router

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-21.jr0wmvf74cg.webp)

  下图中红框部分(忽略马赛克部分)也是 backbond router 同时也是 board area router

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-23.8wf2cgkbdag.webp)

- An **intra-area route** is a route to a destination inside the same OSPF area

  例如下图红框中的 router 通过 OSPF 学到了红框中的网段对应的路由，这条路由就被称为 intra-area route，因为是 router 从同 area 学到的 route

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-25.2jy2x0u64ao0.webp)

- An **interarea route** is a route to destination in a differetn OSPF area

  例如下图红框中的 router 通过 OSPF 学到了红框中的网段对应的路由，这条路由就被称为 interarea route，因为是 router 从不同 area 学到的 route

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-28.zasjsbvy3g0.webp)

## Rules in OSFP

- OSPF areas should be contiguous(连续的)

  例如下图中的拓扑是不允许在 OSPF 中出现的

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-33.5htti5q06y9s.webp)

  因为左边的 Area1 和 右边的 Area1 是通过 Area0 互联的，不是连续的

- All OSPF areas must have least one ABR connected to the backbone area

  例如下图就是不允许的

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-36.6fpsdzy536kg.webp)

  因为 Area1 中的 ABR 没有和 Area0 的 ABR 互联

- OSPF interfaces in the same subnet must be in the same area

  例如下图

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-39.47s9ydc2e1mo.webp)

  红框中的四台 router 对应的接口分别为 192.168.1.1/2/3/4, 因为 192.168.1.1/3/4 都在 area0，但是 192.168.1.2 在 area1，所以这个拓扑是不行的

## Basic OSPF Configuration

例如需要配置成如下拓扑

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-43.3sszlyoq23eo.webp)

针对 R1 配置如下

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_13-46.4emxfsgutr7k.webp)

- `router ospf 1`

  声明使用 OSPF, 对应 OSPF 的进程 ID 为 1

  *The OSPF process ID is locally significant. Routers with different process IDs can became OSPF neighbors*

  > 即在不同的 routers 上 OSPF process ID 可以不同，也可以组成 OSPF network。这一点和 EIGRP 中的 As number 是有区别的

- `network 10.0.12.0 0.0.0.3 area 0`

  和 EIGRP 中的 network 命令类似，但是在 OSPF 中需要指定 area

  > CCNA 只需要配置 single-area(area0)

  1. look for any interfaces with an IP address contained in the range specified in the  network command
  2. Activeate OSPF on the interface in the specified area
  3. The router will then try to become OSPF neighbors with other OSPF-activated neighbor routers
  
  > 如果所有的接口都不在 `network` 命令范围内，router 就不会加入到 OSPF 中的，也无法从其他的 router 学到路由

和 RIP 以及 EIGRP 中的一样，OSPF 也有 `passive-interface <interface-id>` 的功能

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_14-01.2ts5o8s3p3i8.webp)

让 router 指定端口停止发送 OSPF hello 的信息，但是 neighbors 任然是可以收到该端口对应的 network route

> 因为只有 router 之间会发送 OSFP hello，所以该命令通常用在和 end host 互联的 interface 上

### Advertise a default route into OSPF

例如 R1 配置如下

![   ](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_14-03.3875t4hrwlc0.webp)

现在需要将 R1 default route advertise 到其他的 neighbors

和 RIP 一样可以通过 `default-information originate` 命令来实现

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_14-38.6ki3mcyqrp8g.webp)

可以看到上图中 R2 已经成功从 R1 学到了 default route

### show ip protocols

看一下配置了 OSPF 后 R1 `show ip protocols` 的信息

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_14-44.4vhrg7m1gkxs.webp)

- `“ospf 1”`

  对应 `router ospf 1` 命令中声明的

- `Router ID 172.16.1.14`

  逻辑和 EIGRP 中的 Router ID 一样

  1. Manual configuration first

     如果需要手动配置 OSPF router id，可以通过 `router-id <id>` 的方式来实现

     > 和 EIGRP 配置 router id 的方式不同(`eigrp router-id <id>`)

     在使用 `router-id` 命令后，还需要使用 `clear ip ospf process` 或者是重启 router 来让 OSPF route id 生效

     ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_14-52.2yzmsg13a874.webp)

     > 通过 `clear ip ospf process` 来使 OSPF router id 生效实际并不是一个好的方案，因为会 reset OSPF on the router
     >
     > *for a short time and won’t be able to forward traffic to those destinations*
     >
     > 这里必须选择 yes 否则不会生效'

  2. Highest IP address on a loopback interface second

  3. Highest IP address on a physical interface third

- `it is an autonomous system boundary router`

  An autonomous system boundary route(ASBR) is an OSPF router that connects the OSPF network to an external network

  *R1 is connected to the Internet, By using the `default-information originate` command, R1 becomes an ASBR*

  这里 R1 是 ASBR，因为 R1 在拓扑中还连接着互联网

- `Number of areas in this router is 1. 1 normal 0 stub 0 nssa`

  当前 router 含有 area 数量

- `Maximum path:4`

  EIGRP 支持 unequal load-balancing，OSPF 不支持 unequal load-balancing,但是 OSPF 支持 ECMP。这里的逻辑和 EIGRP 中的一样，表示针对相同目的的 ECMP route 最多能有几条

  和 RIP 以及 EIGRP 一样，可以使用 `maximum-paths` 来修改

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_15-05.515vme55g934.webp)

- `Routing for Networks`

  显示 `network` 命令宣告的网段信息

- `Passive Interface(s)`

  router 通过 `passive-interface <interface-id>` 配置的 passive interfacs

- `Routing Information Sources`

  R1 学到 route 的来源

- `Distance:(default is 110)`

  Administrativly Distance

  如果需要修改 OSPF AD 可以通过 `distance <number>` 命令

  ![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_15-19.76z9zzs3fmkg.webp)

## Quiz

### Quiz 1

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_16-04.n0vx0yosjxc.webp)

这里注意一下选项 B，通常都会使用 area 0 来配置 single-area OSPF，但是并不是 single-area OSPF 必须要使用 area 0

area0 是 single-area OSPF 的充分不必要条件 

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230613/2023-06-13_16-16.t1b6dkeffz4.webp)

### 0x01

Configure the appropriate hostname and IP addresses on each device. Enable router interface(You don’t have to configure ISPR1)

**R1**

```
Router(config)#hostname R1
R1(config)#int g0/0
R1(config-if)#ip add 10.0.12.1 255.255.255.252
R1(config-if)#no shutdown
R1(config-if)#int f1/0
R1(config-if)#ip add 10.0.13.1 255.255.255.252
R1(config-if)#no shutdown
R1(config)#int g3/0
R1(config-if)#ip add 203.0.113.1 255.255.255.252
R1(config-if)#no shutdown
```

**R2**

```
Router(config)#hostname R2
R2(config)#int g0/0
R2(config-if)#ip add 10.0.12.2 255.255.255.252
R1(config-if)#no shutdown
R2(config-if)#int f1/0
R2(config-if)#ip add 10.0.24.1 255.255.255.252
R1(config-if)#no shutdown
```

**R3**

```
Router(config)#hostname R3
R3(config)#int f1/0
R3(config-if)#ip add 10.0.13.2 255.255.255.252
R1(config-if)#no shutdown
R3(config-if)#int f2/0
R3(config-if)#ip add 10.0.34.1 255.255.255.252
R1(config-if)#no shutdown
```

**R4**

```
Router(config)#hostname R4
R4(config)#int f2/0
R4(config-if)#ip add 10.0.34.2 255.255.255.252
R1(config-if)#no shutdown
R4(config-if)#int f1/0
R4(config-if)#ip add 10.0.24.2 255.255.255.252
R1(config-if)#no shutdown
R4(config-if)#int g0/0
R4(config-if)#ip add 192.168.4.0 255.255.255.0
R1(config-if)#no shutdown
```

### 0x02

Configure a loopback interface on each router(1.1.1.1/32 for R1, 2.2.2.2 for R2, etc)

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

### 0x03

Configure OSPF on each router

> 这里并不需要保证 ospf process id 是一致的，就如字面意思，这只一个进程 ID

```
R1(config)#router ospf 1
R2(config)#router ospf 2
R3(config)#router ospf 3
R4(config)#router ospf 4
```

Enable OSPF on each interface(including loopback interfaces)

(Do not enable OSPF on R1’s internet link)

> 除 R1 外的 router, 我可以使用 network 0.0.0.0 255.255.255.255
>
> 但是为了练习，避免使用 0.0.0.0 255.255.255.255

```
R1(config-router)#network 1.1.1.1 0.0.0.0 area 0
R1(config-router)#network 10.0.12.0 0.0.0.3 area 0
R1(config-router)#network 10.0.13.0 0.0.0.3 area 0

R2(config-router)#network 2.2.2.2 0.0.0.0 area 0
R2(config-router)#network 10.0.12.0 0.0.0.3 area 0
R2(config-router)#network 10.0.24.0 0.0.0.3 area 0

R3(config-router)#network 3.3.3.3 0.0.0.0 area 0
R3(config-router)#network 10.0.13.0 0.0.0.3 area 0
R3(config-router)#network 10.0.34.0 0.0.0.3 area 0

R4(config-router)#network 4.4.4.4 0.0.0.0 area 0
R4(config-router)#network 10.0.34.0 0.0.0.3 area 0
R4(config-router)#network 10.0.24.0 0.0.0.3 area 0
R4(config-router)#network 192.168.4.0 0.0.0.255 area 0
```

这里并不需要使用 `network 203.0.113.0 0.0.0.3 area 0` 来宣告 R1 和 ISPR1 之间的网段，因为即使在 OSPF 中的 router 不知道这条路由，也能通过 R1 出公网，只需要学到能到 R1 的路由即可

可以使用 `show ip protocol` 来查看 R1 G3/0 是否加入到 OSPF network 中

```
R1#sh ip pro

Routing Protocol is "ospf 1"
  Outgoing update filter list for all interfaces is not set 
  Incoming update filter list for all interfaces is not set 
  Router ID 1.1.1.1
  Number of areas in this router is 1. 1 normal 0 stub 0 nssa
  Maximum path: 4
  Routing for Networks:
    1.1.1.1 0.0.0.0 area 0
    10.0.12.0 0.0.0.3 area 0
    10.0.13.0 0.0.0.3 area 0
```

Configure passive interfaces where appropriate(including loopback interface)

```
R1(config-router)#passive-interface g3/0
R1(config-router)#passive-interface lo0

R2(config-router)#passive-interface lo0

R3(config-router)#passive-interface lo0

R4(config-router)#passive-interface lo0
R4(config-router)#passive-interface g0/0
```

> 如果 router id 配置错误，可以使用下面的命令

```
R4(config)#router ospf 4
R4(config-router)#router-id 4.4.4.4
R4(config-router)#Reload or use "clear ip ospf process" command, for this to take effect
R4(config-router)#end
R4#sh ip pro
Routing Protocol is "ospf 4"
  Outgoing update filter list for all interfaces is not set 
  Incoming update filter list for all interfaces is not set 
  Router ID 192.168.4.254
  ...
R4#clear ip ospf process
Reset ALL OSPF processes? [no]: yes
R4#sh ip pro
Routing Protocol is "ospf 4"
  Outgoing update filter list for all interfaces is not set 
  Incoming update filter list for all interfaces is not set 
  Router ID 4.4.4.4
  ...
```

### 0x04

Configure R1 as an ASBR that advertises a default route into the OSPF domain

```
R1(config-if)#ip route 0.0.0.0 0.0.0.0 203.0.113.2
R1(config)#router ospf 1
R1(config-router)#default-information originate 
```

如果使用上面的命令，在 R1 上使用 `show ip protocol` 就可以发现

```
R1#show ip pro

Routing Protocol is "ospf 1"
  Outgoing update filter list for all interfaces is not set 
  Incoming update filter list for all interfaces is not set 
  Router ID 1.1.1.1
  It is an autonomous system boundary router
```

`  It is an autonomous system boundary router` 表示当前 router 是 ASBR 通过该 Route 连接公网

### 0x05

Check the routing tables of R2,R3 and R4. What default route(s) were added

```
#R2
O*E2 0.0.0.0/0 [110/1] via 10.0.12.1, 00:00:29, GigabitEthernet0/0
#R3
O*E2 0.0.0.0/0 [110/1] via 10.0.13.1, 00:01:32, FastEthernet1/0
#R4
O*E2 0.0.0.0/0 [110/1] via 10.0.34.1, 00:01:50, FastEthernet2/0
               [110/1] via 10.0.24.1, 00:01:50, FastEthernet1/0
```

> 这里需要了解 OSPF mtric，才能知道这里为什么会使用 ECMP

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=pvuaoJ9YzoI&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ