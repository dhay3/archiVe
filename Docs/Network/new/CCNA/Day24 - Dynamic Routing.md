# Day24 - Dynamic Routing

## Why need Dynamic Routing

有如下拓扑

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_13-58.6fjs5tnpvcsg.webp)

==主要要关注 R1==, 在没有配置任何路由的情况下，R1 上的 routing table 如下

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-00.4x5x9ipal43k.webp)

只有 2 种类型的 Route

Connected Route ∈ Network Route

Local Route ∈ Hot route

现在每台 Router 都配置 Dynamic Routing 

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-03.4b5okozaclj4.webp)

那么 

R4 就会向 R2 advertise 192.168.4.0/24 的路由通过 R4 G0/0

R2 会将 R1 advertise 192.168.4.0/24 的路由通过 R2 G0/0，同时 advertise 10.0.24.0/30 的路由通过 R2 G0/0

R1 的 routing table 如上

假设现在 R4 G0/0 down 了

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-08.3hzkazcrwq2o.webp)

那么 R1 就会将 192.168.4.0/24 和 10.0.24.0/30 路由删掉

如果所有的 Router 都没有配置 Dynamic Routing，但是为 R1 单独配置了一条 192.168.4.0/24 的路由

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-12.44wzk5pdltfk.webp)

当 R4 G0/0 down 掉，R1 对应 192.168.4.0/24 的路由不会自动删除，所以仍然会将流量通过 route 转发，因为 R2 G1/0 - R4 G0/0 link down，所以就会丢包

*Dynamic Routing will remove invalid routes*

因为 Dynamic Routing 会自动删除失效的路由，所以应该需要配置 backup route

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-17.1wlhs03iaiu8.webp)

在 R4 G0/0 正常的情况下，使用了 Dynamic Routing，可以看到 R1 只增加了一条 192.168.4.0/24 via 10.0.12.2 (R2 G0/0)的路由

> 如果有相同的 network route，Dynamic Routing 只会添加 preferred route
>
> 192.168.4.0/24 via 10.0.12.2 是 preferred route
>
> 因为 R3 和 R4 之间是 fasterethernet connection，而 R1 和 R2 之间是 gigabitethernet connection
>
> 和 STP 中的 root cost 类似，Route 也是有 cost

现在手动 down 掉 R4 G0/0

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-19.1ecn2167j2qo.webp)

R1 上对应 192.168.4.0/24 via 10.0.12.2 的路由就会被自动删除，然后增加一条 192.168.4.0/24 via 10.0.13.2(R3 G0/0) 的路由

### Points of Dynamic Routing

- Routers can use dynamic routing protocols to advertise information about the routes they known to other routers

- They form ‘adjacencies’/‘neighbor relationships’/‘neighborships’ with adjacent routers to exchange this information

  例如上图中 R1 adjacencies 就是 R2 和 R3，R1 会向 R2 和 R3 advertise routes

- If multiple routes to a destination are learned, the router determines which route is superior and adds it to the routing table. It uses the ‘**metric**’ of the route to decide which is superior(lower metric = superior)

##  Types of Dynamic Routing Protocol

Dynamic Routing 协议可以分为两大类

1. IGP(Interior Gateway Protocol)

   Used to share routes within a single autonomous system(AS), which is a single organization(ie. a company)

2. EGP(Exterior Gateway Protocol)

   Used to share routes between different autonomous systems

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-32.74jn7o2b70n4.webp)

例如上图中

Company A、ISP A、ISP B、Company B 中 routers 互相通过 IGP 学习路由

Company A 和 ISP A， ISP A 和 ISP B, ISP B 和 Company B 之间互相通过 EGP 学习路由

### IGP vs EGP

> Algorithm type
>
> used by each protocol to share route information and determin the best  route to each destination 

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-09_14-39.2vihgyh7kfsw.webp)

EGP 只使用 Path Vector Algorithm，在现在的网络中只有一种就是 BGP

> CCNA 不包括 BGP，BGP 是 CCNP 中的内容

### Distance vector protocol

> RIP 和 EIGRP 是 distance vector protocol

- Distance vector protocols were invented before link state protocols

- Early examples are RIPv1 and Cisco’s proprietary protocol IGRP(which was updated to EIGRP)

- Distance vetor protocols operate by sending the following to their directly connected neighbors

  1. their known destination networks
  2. their metric to reach their known destination networks

  > this method sharing route information is often called ‘routing by rumor’
  >
  > this is because the router doesn’t know about the network beyond its neighbors. It only knows the information that its neighbors tell it

*Called ‘distance vector’ because the routers only learn the ‘distance’ (metric) and ‘vector’ (direction, the next-hop router) of each route*

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_09-54.1hd1rxgilxvk.webp)

### Link state routing protocol

> OSPF 和 IS-IS 是 Link state protocol

- When using a link state routing protocol, every router creates a ‘connectivity map’ of the network
- To allow this each router advertises information about this interfaces(connected networks(connected route)) to its neighbors. These advertisements are passed along to other routers, until all routers in the network develop the same map of the network
- Each router independently uses this map to calculate the best routes to each destination
- Link state protocols use more resources(CPU) on the router, because more information is shared
- However, link state protocols tend to be faster in reacting to changes in the network than distance vector protocols

## Dynamic Routing Protocol Metrics

router‘s route table 只会含有最优的路由，如果使用 dynamic routing protocol 出现 two different routes to the same destination，router 怎么解决那条路由是最优路由呢？

*It uses the metric value of the routes to determine which is best. A lower metric=better*

引用上面的例子，逻辑上 R1 到 192.168.4.0/24 的路由有 2 条，via R2 和 via R3，但是因为 via R3 互联的链路是 fastethernet，而 via R2 互联的链路是 gigabitethernet，所以 via R2 的路由要优于 via R3 的路由(实际还是比较 metric)

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_10-17.6g04kz8scjuo.webp)

现在将 R3 和 R4 互联的链路变成 gigabitethernet，那么那条才是最优的路由呢？

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_10-22.4ngh3xw5se0w.webp)

实际上两条路由都会被 R1 学到

*If a router learns two(or more) routes via the **same routing protocol** to the **same destination**(same network address, same subnet mask) with the **same metric**, both will be added to the routing table. Traffic will be load-balanced over both routes* 

> 要求 相同协议，相同目的地址，相同 metric

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_10-38.2mocj5i0qykg.webp)

上图中，黄框表示两条路由均使用 OSPF，红框表示两条路由 metric 均相同为 3，蓝框表示 AD(administrative distance)

在 dynamic protocol 中这种现象也被称为 **ECMP(Equal Cost Multi-Path)**

> 其实不仅在 dynamic protocol 中有这种现象，在 static route 也有
>
> 但是 static route 并不使用 metric 来衡量路由的优劣

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_10-36.14fpi56l5iow.webp)

协议不同，metric 衡量的方式也不同，具体可以参考下表

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_10-41.1o1sksebzdfk.webp)

RIP,只以 hop 来计算 metric，一跳就是 one hop，不管互联的 link 是 fastethernet 还是 gigabitethernet 或者是 ten gigabitethernet

OSPF 以链路带宽来计算 metric

例如下面这个拓扑

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_11-02.4527fgggy4hs.webp)

- RIP

  如果使用 RIP，R1 会添加两条路由 via R2 和 via R3，即使链路的带宽不一样

- OSPF

  如果使用 OSPF，R1 只会添加 1 条路由，via R2，因为 via R2 的链路总带宽要比 via R3 的大

## Administrative Distance

在大多数的公司里一般只会使用一种 IGP 通常是 OSPF 或者是 EIGRP，但是在一些特殊的场景下，需要两种协议混用。例如两家使用不同协议 IGP 的公司需要互通

*Metric is used to compare routes <u>learned via the same routing protocol. Different routing protocols use totally different metrics, so they cannot be compared</u>*

例如，在 RIP 中 192.168.4.0/24 路由 metric 5，但是在 OSPF 中 192.168.4.0/24 路由 metric 30。Router 应该选择那条路由？

> RIP 中使用 hop 作为 metric 和 OSPF 中使用带宽作为 metric 对比显然是不合适的，函数中的变量不一致

*The **administrative distance(AD)** is used to determine which routng protocol is preferred*

*A lower AD is preferred, and indicates that the routing protocol is considered more ‘trustworthy’(more likely to select good routes)*

> 如果路由目的相同，就会比较协议的 AD 值，值小的路由作为优选路由

IGP 协议的 AD 值，可以参考下表

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_11-21.31oteve0hqkg.webp)

所以前面的那个问题，显然会使用 OSPF 对应的那条协议

*If the administrative distance is 255, the router does not believe the source of that route and does not install the route in the routing table*

如果 AD 值为 255，就表示该条路由不可靠，router 并不会将对应的路由加入到 routing table

> 协议 AD 值是可以通过配置手动修改的，static route AD 也是可以修改的

例如下图中就修改 10.0.0.0/8 static route AD 值为 100

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_11-43.3ac280x391vk.webp)

*By changing the AD of a static route, you can make it less preferred than routes learned by a dynamic routing protocol to the same destination(make sure the AD is higher than the routing protocol’s AD)*

*This is called a ‘**floating static route**’*

*The route will be inactive(not in the routing table) unless the route learned by the dynamic routing protocol is removed(for example, the router stops advertising it for some reason, or an interface failure causes an adjacency with a neighbor to be lost)*

## Route selection

==这里的 Route selection 只表示那条 route 应该加入到 routing table==，并不是表示 Router 收到报文后，使用那条 route(如果需要判断使用那条 route，只有一个标准即 the longest prefix)

> 左 AD 右 metric，先比 AD 后比 metric，等价 ECMP 

```
if a.protocol.AD < b.protocol.AD then:
	add(a.route)
else if a.protocol.AD > b.protocol.AD then:
	add(b.route)
#AD 相同证明协议相同
else if a.protocol.AD == b.protocol.AD then:
	if a.protocol.metric < b.protocol.metric then:
		add(a.route)
	else if a.protocol.metric > b.protocol.metric then:
		add(b.route)
	else if a.protocol.metric == b.protocol.metric then:
		add(a.route)
		add(b.route)
```

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230609/2023-06-12_13-31.1xtwftm28bgg.webp)

### 0x01

- which dynamic routing protocol is Enterprise A using?

  ```
  O       10.0.2.0/24 [110/2] via 10.0.0.2, 00:01:01, GigabitEthernet0/2/0
  ```

  OSPF

- which route will be used if PC1 tries to access SRV1?

  ```
  O       10.0.2.0/24 [110/2] via 10.0.0.2, 00:01:01, GigabitEthernet0/2/0
  ```

- which route will be used if PC1 tries to access remote server 1.1.1.1 over the Internet

  ```
  S*   0.0.0.0/0 [1/0] via 203.0.113.9
  ```

### 0x02

- Configure floating static rotues on R1 and R2 that allow PC1 to reach SRV1 if the link between R1 and R2 fails. 

  > 如果用 32 bit mask 就会显示

  ```
  R1(config)#ip route 10.0.2.0 255.255.255.255 203.0.113.1 250
  R2(config)#ip route 10.0.1.0 255.255.255.255 203.0.113.5 250
  ```

- Do the routes enter the routing tables of R1 and R2

  No, because those are floating static routes

### 0x03

- Shut down the G0/2/0 interface of R1 or R2.

  Do the floating static routes enter the routing tables of R1 and R2?

  yes

  ```
  R1(config)#do sh ip ro
  Gateway of last resort is 203.0.113.9 to network 0.0.0.0
  
       10.0.0.0/8 is variably subnetted, 6 subnets, 3 masks
  C       10.0.0.0/30 is directly connected, GigabitEthernet0/2/0
  L       10.0.0.1/32 is directly connected, GigabitEthernet0/2/0
  C       10.0.1.0/24 is directly connected, GigabitEthernet0/1
  L       10.0.1.254/32 is directly connected, GigabitEthernet0/1
  O       10.0.2.0/24 [110/2] via 10.0.0.2, 00:15:32, GigabitEthernet0/2/0
  S       10.0.2.1/32 [250/0] via 10.0.0.2
       203.0.113.0/24 is variably subnetted, 4 subnets, 2 masks
  C       203.0.113.0/30 is directly connected, GigabitEthernet0/0/0
  L       203.0.113.2/32 is directly connected, GigabitEthernet0/0/0
  C       203.0.113.8/30 is directly connected, GigabitEthernet0/1/0
  L       203.0.113.10/32 is directly connected, GigabitEthernet0/1/0
  S*   0.0.0.0/0 [1/0] via 203.0.113.9
  ```

  shutdown R1 G0/2/0

  ```
  R1(config)#int g0/2/0
  R1(config-if)#shutdown
  ```

  show routing table again

  ```
  R1(config-if)#do sh ip ro
  
  Gateway of last resort is 203.0.113.9 to network 0.0.0.0
  
       10.0.0.0/8 is variably subnetted, 3 subnets, 2 masks
  C       10.0.1.0/24 is directly connected, GigabitEthernet0/1
  L       10.0.1.254/32 is directly connected, GigabitEthernet0/1
  S       10.0.2.0/24 [250/0] via 203.0.113.1
       203.0.113.0/24 is variably subnetted, 4 subnets, 2 masks
  C       203.0.113.0/30 is directly connected, GigabitEthernet0/0/0
  L       203.0.113.2/32 is directly connected, GigabitEthernet0/0/0
  C       203.0.113.8/30 is directly connected, GigabitEthernet0/1/0
  L       203.0.113.10/32 is directly connected, GigabitEthernet0/1/0
  S*   0.0.0.0/0 [1/0] via 203.0.113.9
  
  R2(config)#do sh ip ro
  
  Gateway of last resort is 203.0.113.13 to network 0.0.0.0
  
       10.0.0.0/8 is variably subnetted, 3 subnets, 2 masks
  S       10.0.1.0/24 [250/0] via 203.0.113.5
  C       10.0.2.0/24 is directly connected, GigabitEthernet0/1
  L       10.0.2.254/32 is directly connected, GigabitEthernet0/1
       203.0.113.0/24 is variably subnetted, 4 subnets, 2 masks
  C       203.0.113.4/30 is directly connected, GigabitEthernet0/0/0
  L       203.0.113.6/32 is directly connected, GigabitEthernet0/0/0
  C       203.0.113.12/30 is directly connected, GigabitEthernet0/1/0
  L       203.0.113.14/32 is directly connected, GigabitEthernet0/1/0
  S*   0.0.0.0/0 [1/0] via 203.0.113.13
  ```

  

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=xSTgb8JLkvs&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=46