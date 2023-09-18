# Day27 - OSPF Part 2

## OSPF metric

在 OSPF 中 metric 通常也被称为 OSPF cost，是路由路径中所有出向端口的 interface cost 计算总和

*The OSPF cost to a destination is the total cost of the ‘outgoing/exit’ interfaces*

> 这点和 STP 中的 root cost 类似

interface cost 基于端口的 bandwidth(speed)，公式如下

$reference\ bandwidth \div interface\ bandwidth$

==reference bandwidth 默认为 100 mbps==，例如

端口是 10mbps 的，interface cost 就会为 100/10 = 10

端口是 100mbps 的，interface cost 就会为 100/100 = 1

那么 1Gbps 的端口 interface cost 是多少？

*All values less than 1 will be converted to 1* 

所以 1Gbps 的端口 interface cost 为 1，10Gbps 的端口 interface cost 还是为 1 。以此类推

可以使用 `show ip ospf interface <interface-id>` 来查看端口对应的 interface cost

例如

R3 f2/0 cost 为 1，因为 100/100 = 1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_17-46.32dxb99mlts0.webp)

R3 G0/0 cost 也为 1，因为 100/1000 < 1，向上取整为 1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_17-49.4su1zx9f096o.webp)

显然的这种方式来计算 OSFP cost 显然是不合理的，所以 reference bandwidth 也是可以通过命令 `R1(config-router)# auto-cost reference-bandwidth <megabits-per-sencond>` 来修改的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_17-51.4n9s2i9accxs.webp)

*You  should configure a reference bandwidth greater than the fastest links in your network(to allow for future upgrades)*

*You should configure the same reference bandwidth on all OSPF routers in the network* 

> 这里就是为了统一 OSPF network 中的 cost

### Example

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_18-00.6nbtatgh8em8.webp)

例如上图，使用了 `auto-cost reference-bandwidth 100000`，现在 R1 需要访问 192.168.4.0/24 这个段。那么 OSPF cost 如下
$$
100(R1\ G0/0) + 100(R2\ G1/0) + 100(R4\ G1/0) = 300
$$
那现在 R1 访问 R2 loopback 0 呢？

> *Loopback interfaces have a cost of 1(Not matter the reference-bandwidth)*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_18-07.jq1gjijcs74.webp)

那么 OSPF cost 如下
$$
100(R1\ G0/0) + 100(R2\ Looback0) = 101
$$
可以看一下 R1 在未修改 reference bandwidth 之前的路由表

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_18-09.16q0vk7rzjb4.webp)

这里可以观察到 OSPF 会引入 2 条 ECMP 路由，via R2 和 via R3，因为对应 reference  bandwidth 100 来说，两条 route cost 都为 3 

对比一下修改了 reference bandwidth 到 100000 后，R1 的路由表

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_18-14.4d374wdd9vnk.webp)

这里就可以看到只会引入一条 via R2 的路由，因为对应 via R3 的路由 cost 为 $100(R1\ G1/0)+1000(R3\ F2/0)+100(R4\ G1/0)=1200$

而 via R2 的路由，对应 cost 为 $100(R1\ G0/0)+100(R2\ G1/0)+100(R4\ G1/0)=300$

同样的还可以看到 R1 到 2.2.2.2(R2 loopback0)的 cost 为 101，因为 $100(R1\ G0/0)+1(R2\ lo0)=101$ 

### ip ospf cost

除了 `auto-cost reference-bandwidth <megabits-per-sencond>` 这种间接的方式来修改 interface cost 外，还可以通过 `ip ospf cost <cost>` 的方式来直接修改 interface cost

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_18-23.5cgj90824oe8.webp)

### bandwidth

除了 `ip ospf cost <cost>` 之外，还可以通过直接修改端口的 bandwidth 来达到修改 interface cost 的目的（原因查看 interface cost 计算公式）

*Although the bandwidth matches the interface speed by default, changing the interface bandwidth <u>doesn’t actually change the speed at which the interface operates</u>*

> 如果想要修改端口的 speed，可以使用 `speed` 命令

*The bandwidth is just a value that is used to calculate OSPF cost, EIGRP metric,etc*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-13_18-33.2czvtnzvuvnk.webp)

> 注意 bandwidth 命令 number 使用的 unit 是 kilobits，而 reference-bandwidth 命令 number 使用的 unit 是 megabits

但是十分不推荐通过 `bandwidth <bandwidth>` 命令来修改 interface cost，因为不仅仅只有 OSPF 使用 bandwidth 作为 cost 的计算因子，像 EIGRP metric 也会使用 bandwidth 作为计算因子。所以尽量使用 `auto-cost reference-bandwidth <megabits-per-sencond>` 或者 `ip ospf cost <cost>` 来修改 interface cost

### Summary

有 3 种方式可以来修改 OSPF cost

1. change the reference bandwidth

   `R1(config-router)# auto-cost reference-bandwidth <megabits-per-second>`

2. Manual configuration

   `R1(config-if)# ip ospf cost <cost>`

3. Change the interface bandwidth

   `R1(config-if)# bandwidth kilobits-per-second`

如果需要查看每个接口的 OSPF cost，除了使用 `show ip ospf interface <interface-id>` 外，还可以使用 `show ip ospf interface br` 来查看

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_10-18.3lohnlxh8ikg.webp)

## OSPF Neighbors

neighbors 是 OSPF 中重要的一个概念。只有 routers 成为 neighbors，OSPF 才能在 routers 之间正常运行

*When OSPF is activated on an interface, the router starts sending OSPF **hello** messages out of the interface at regular intervals(determined by the **hello timer**). These are used to introduce the router to potential OSPF neighbors*

> 只要端口加入到了 OSPF network 中，就会发送 hello message

*The default hello timer is 10 seconds on an Ethernet connection*

*Hello messages are multicast to 224.0.0.5(multicast address for all OSPF routers)*

> hello messages 3 层报文头中的 Protocol 字段值为 89

## OSPF States

在两台 router 之间，使用 OSPF，neighbors 需要经过如下几种状态

1. Down
2. Init
3. 2-way
4. Exstart
5. Exchange
6. Loading
7. Full

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_11-45.1ue4qvepmqo0.webp)

- 前 3 种状态为构建两台 router 之间的 adjacency 关系，即 becomes neighbors
- 后 4 种状态为交换两台 router 之间的 LSDB 中的 LSAs 信息，即 Exchange LSAs

### Down State

假设 R2 G0/0 已经加入到 OSPF network，现在 R1 G0/0 加入到了 OSPF network 中

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_10-33.7iolsefrmfls.webp)

首先 R1 会发送 OSPF hello message 到 OSPF multicast address 224.0.0.5

hello message 中主要包含几个字段

- My RID(router ID)
- Neighbor RID(router ID)

因为 R1 现在还不知道对端互联的 R2 是 OSPF neighbor，所以 Neighbor RID 字段的值为 0.0.0.0(理解成类似 unkown unicast)，所以逻辑上就认为 R2 的状态为 down state

> 这时 R1 的 OSPF neighbor table 中是空的，并不会记录 R2 down state

### Init State

现在 R2 收到了从 R1 发来的 OSPF hello message

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_10-51.4hc7ltudzczk.webp)

R2 会将 R1 对应的信息加入到自己的 OSPF neighbor table 中，同时在 R2 OSPF neigbor table 中 R1 的状态标识为 init

> 注意因为当前 R2 并没有回送任何消息给 R1，所以 R1 OSPF neighbor table 还是为空 

*Basically, the init state means that a Hello packet was received, but R2’s own router ID is not the Hello packet*

### 2-way State

在 R2 收到 R1 发来的 hello message 进入 init state 之后

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_10-59.2obbgu8e046.webp)

R2 会发送一个 hello messge，包含 the RID of both routers(R1 和 R2)

然后 R1 会把 R2 对应的信息加入到自己的 OSPF neighbor table 中，同时在 R1 OSPF neighbor table 中 R2 的状态为 2-way state

同时 R1 也会发送另外一个 hello message，和第一个发送的 hello message 不同，这时会携带上 R2 RID，R2 在收到 R1 发来的这个 hello message 后会将自己的 OSPF neighbor table 中，R1 对应的状态改为 2-way state

*The 2-way state means the router has received a hello packet with own RID in it*

如果两台 routers 都进入了 2-way state, 意味着他们直接已经准备成为 OSPF neighbors 了，准备发送 LSAs 到对端的 LSDB

*In some network types, a DR(Designated Router)and BDR(Backup Designated Router) will be elected at this point.* 在第 28 天中的内容解释

### Exstart State

两台 routers 都进入 2-way state 后

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_11-14.6l13qyrbq3k0.webp)

会开始准备交换对方的 LSDB 信息，在此之前，双方会协商那台 router 是 master，那台是 slave

> 通过 DBD 来决定

*They decide which will be the Master and which will be the slave in the Exstart state*

RID 值大的那台 router 会成为 Master，RID 值小的那台 router 会成为 Slave。所以例子 R1 是 Slave，R2 是 Master

### Exchange State

在决定那台 Router 是 Master 还是 Slave 后

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_11-16.k66sh54k5v4.webp)

Routers 通过含有各自 LSDB 中的 LSAs 基本信息的 DBDs 交换信息，对比决定需要对方的那条 LSA 的详细信息

### Loading state

在确定需要对方那条 LSA 详细信息之后

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_11-23.7h7ttsfkho8w.webp)

> 这里只展示了 R1 做  LSR 的过程，实际 R2 也会做一样的

两台 Router 都会发送 Link State Request(LSR)，请求对方发送 Exchange state 中决定的 LSA

两台 Routers 会将需要发送给对方的 LSA 通过 LSU 发送给对方

当对方收到 LSU 后，会发送 LSAck 给发送 LSU 的哪一方

> 这一点就和 TCP three-way handshake 类似

### Full State

在两台 Routers 都收到了对方的 LSAck 后

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_11-34.6sfqt30nk9s0.webp)

两台 Router 会将自己的 OSPF neighbor table 对方的条目，变为 full state

> 在 full state 的状态下，两台 router 构成了完整的 OSPF adjaccency 在各自的 LSDB 中有一样的 LSAs

但是这没有完全结束，双方会继续发送 hello message(every 10 seconds by default)以维护 neighbor adjacency

因为网络拓扑会变更，OSPF 对应的信息也会变更。所以和 STP 中的 Max age timer 类似，OSPF 也需要一个 timer 来表示已经完全收不到对端 OSPF 信息(hello messge)的条件，在 OSPF 中被称为 Dead timer(40 seconds by default)

当 timer 从 40 到 0 时，就会将对端从自己的 OSPF neighbor table 中移除，如果在 40 内收到对端发送过来的 hello message 就会将 timer 重置为 40 

## OSPF message

在 OSPF states 之间的变化过程中，需要发送的 OSPF message 一共有 5 种

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_12-05.6dhjgdo81mkg.webp)

## OSPF show command

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_13-41.1uq577o9s2w0.webp)

可以使用 `show ip ospf neighbor` 来查看当前 router OSPF neighbor 相关的信息

- `Neighbor ID`

  表明 neighbor 的 Router ID

- `State`

  表明 neighbor 的状态

- `Dead Time`

  dead timer 计时的时间，默认从 40 seconds 开始

- `Address`

  neighbor 接口的 IP address

可以使用 `show ip ospf interface <interface-id>` 来查看对应接口关于 OSPF 的信息

- `Hello 10,Dead 40`

  对应 hello timer 和 dead timer 的值

- `Hello due in 00:00:07`

  表示 R1 会在 7 秒内发送一次 hello message

- `Neighbor Count is 1,Adjacent neighbor count is 1`

  表示通过 R1 G0/0 互联的 OSPF neighbor 只有 1 个

  > Adjacent Neighbor 指的是 full state 的 neighbors
  >
  > 而 Neighbor 仅仅指的是 OSPF neighbor 
  >
  > Aajacent Neighbor 是 Neighbor 的真子集

## Additional OSPF Configuration

OSPF 可以使用 network 命令，将在指定 IP 范围的 interface 加入 OSPF network 中外，还可以通过 `ip ospf <process-id> area <number>` 的方式，将端口显示声明加入到 OSPF 中的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_14-53.3z13cudu6ncw.webp)

另外除了使用 `passive-interface <interface-id>` 来让指定的端口变为 passive 外，还可以使用 `passive-interface default` 来让所有的端口变为 passive，然后使用 `no passive-interface <interface-id>` 来让指定端口不再 passive

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_14-55.24fapysecmo0.webp)

如果是通过上述的方法直接将端口加入到 OSPF 中，在 `show ip protocols` 中输出的内容会和使用 network 命令加入到 OSPF 中的有区别

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_14-57.6co74skriruo.webp)

Routing for Network 部分是空的，Routing on interfaces Configured Explicitly 部分会显示非通过 network 命令直接加入到 OSPF 的接口

## LAB

### 0x01

Configure the appropriate hostname and IP addresses on each device. Enable router interfaces(You don’t have to configure ISPR1)

**R1**

```
Router(config)#hostname R1
R1(config)#int g3/0
R1(config-if)#ip add 203.0.113.1 255.255.255.252
R1(config-if)#no shutdown
R1(config)#int f1/0
R1(config-if)#ip add 10.0.13.1 255.255.255.252
R1(config-if)#no shutdown
R1(config)#int g0/0
R1(config-if)#ip add 10.0.12.1 255.255.255.252
R1(config-if)#no shutdown
```

**R2**

```
Router(config)#hostname R2
R2(config)#int g0/0
R2(config-if)#ip add 10.0.12.2 255.255.255.252
R2(config-if)#no shutdown
R2(config)#int f1/0
R2(config-if)#ip add 10.0.24.1 255.255.255.252
R2(config-if)#no shutdown
```

**R3**

```
Router(config)#hostname R3
R3(config)#int f1/0
R3(config-if)#ip add 10.0.13.2 255.255.255.252
R3(config-if)#no shutdown
R3(config)#int f2/0
R3(config-if)#ip add 10.0.34.1 255.255.255.252
R3(config-if)#no shutdown
```

**R4**

```
Router(config)#hostname R4
R4(config)#int f2/0
R4(config-if)#ip add 10.0.34.2 255.255.255.252
R4(config-if)#no shutdown
R4(config)#int f1/0
R4(config-if)#ip add 10.0.24.2 255.255.255.252
R4(config-if)#no shutdown
R4(config)#int g0/0
R4(config-if)#ip add 192.168.4.254 255.255.255.252
R4(config-if)#no shutdown
```

### 0x02

Configure a loopback interface on each router(1.1.1.1/32 for R1,2.2.2.2/32 for R2，etc)

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

Enable OSPF directly on each interface of the routers.

```
R1(config-if)#int f1/0
R1(config-if)#ip ospf 1 area 0
R1(config-if)#int g0/0
R1(config-if)#ip ospf 1 area 0
R1(config-if)#int lo0
R1(config-if)#ip ospf 1 area 0

R2(config-if)#int g0/0
R2(config-if)#ip ospf 1 area 0
R2(config-if)#int f1/0
R2(config-if)#ip ospf 1 area 0
R2(config-if)#int lo0
R2(config-if)#ip ospf 1 area 0

R3(config-if)#int f1/0
R3(config-if)#ip ospf 1 area 0
R3(config-if)#int f2/0
R3(config-if)#ip ospf 1 area 0
R3(config-if)#int lo0
R3(config-if)#ip ospf 1 area 0

R4(config-if)#int f2/0
R4(config-if)#ip ospf 1 area 0
R4(config-if)#int f1/0
R4(config-if)#ip ospf 1 area 0
R4(config-if)#int g0/0
R4(config-if)#ip ospf 1 area 0
R4(config-if)#int lo0
R4(config-if)#ip ospf 1 area 0
```

Configure passive interfaces as appropriate

```
R1(config)#router ospf 1
R1(config-router)#passive-interface default 
R1(config-router)#no passive-interface f1/0
R1(config-router)#no passive-interface g0/0

R2(config)#router ospf 1
R2(config-router)#passive-interface default 
R2(config-router)#no passive-interface f1/0
R2(config-router)#no passive-interface g0/0

R3(config)#router ospf 1
R3(config-router)#passive-interface default 
R3(config-router)#no passive-interface f1/0
R3(config-router)#no passive-interface f2/0

R4(config)#router ospf 1
R4(config-router)#passive-interface default 
R4(config-router)#no passive-interface f1/0
R4(config-router)#no passive-interface f2/0
```

### 0x04

Configure the reference bandwidth on each router so a FastEthernet interface has a cost of 100

```
R1(config-router)#auto-cost reference-bandwidth 10000
R2(config-router)#auto-cost reference-bandwidth 10000
R3(config-router)#auto-cost reference-bandwidth 10000
R4(config-router)#auto-cost reference-bandwidth 10000
```

可以使用 `show ip ospf interface <interface-id>` 来查看是否实际生效

### 0x05

Configure R1 as an ASBR that advertise a default route in the OSPF domain

```
R1(config)#ip route 0.0.0.0 0.0.0.0 203.0.113.2
R1(config-router)#default-information originate 
```

这里只会导入 via R2 的 default route

```
O*E2 0.0.0.0/0 [110/1] via 10.0.24.1, 4294967273:4294967262:4294967247, FastEthernet1/0
```

### 0x06

Check th routing tables of R4.

```
O*E2 0.0.0.0/0 [110/1] via 10.0.24.1, 00:00:13, FastEthernet1/0
```

What default route(s) were added

这里只会将 via R2 的默认路由加入到 routing table 中，因为 via R2 OSPF cost 更小 $10(R1\ G0/0)+100(R2\ F1/0)=110$，而 via R3 OSPF cost 为 $100(R1\ F1/0)+100(R3\ F2/0) = 200$

> 在 OSPF 中没有类似 EIGRP 中的 `show ip eirgp topology` 的命令，来查看所有从 OSPF 学来的路由

### 0x07

Use simmulation mode to view the OSPF hello messages being sent by routers.

What fields are included in the hello message

具体看 Packettracer OSPF PDU details 

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=VtzfTA21ht0&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=51