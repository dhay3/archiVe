# Day11 - Routing Fundamentals

## What is routing

- Routing is the process that routers use to determine the path that IP packets should take over a network to reach their destination

  - Router store routes to all of their known destinations in a **routing table**
  - When routers receive packets, they look in the routing table to find the best route to forward that packet

- There are two main routing methods (methods that routers use to learn routes)

  - **Dynamic Routing**

    Routers use dynamic  routing protocols(ie. OSPF) to share routing information with each other automatically and build their routing tables

  - **Static Routing**

    A network engineer/admin manully configures routes on the router

- A router tells the router：to send a packet to destination X，you should send the packet to **next-hop**(the next router in the path to the destination) Y first

  - or, if the destination is directly connected to the router, send the packet directly to the destination.
  - or, if the destination is the router’s own IP address receive the packet for itselft(don’t forwrd it)

以下图为例子说明，主要关注 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-24_15-55.4xlds643ghkw.webp)

先配置 R1 接口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-24_15-58.1af5bdoaf5gg.webp)

## show ip route

配置完 R1 接口后，可以使用 `show ip route` 来查看路由

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-24_16-00.67ycin57c6ww.webp)

Codes 部分是 legend (图例)

- L - local

  A route to the actual IP address configured on the interface(with a /32 netmask)

- C - connected

  A route to the network the interface is connected to(with the actual netmask configured on the interface)

只要你对一个接口配置了 IP 并使用了 `no shutdown`，就会有 2 条 routes (connected route, local route)自动加入到 routing table

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-24_16-06.1vj6vy4vyb9c.webp)

### Connected and Local routes

> 如果端口配置的是 192.168.1.1/24，那么对应的 local route 就会是 192.168.1.1/32，对应的 connected route 就会是 192.168.1.0/24

- 蓝色部分 Connected route

  A Connected route is a route to the network the interface is connected to. It provides a route to all hosts in that network(192.168.1.0 -  192.168.1.255)

  *R1 knowns “If I need to send a packet to any host in the 192.168.1.0/24 network, I should send it out of G0/2”*

  例如

  192.168.1.2 match， send packet out of G0/2

  192.168.1.89 match, send packet out of G0/2

  ==192.168.2.1 not match, send the a different route, or drop the packet if there is no matching route==

- 绿色部分 Local route

  A local route is a route to the exact IP address configured on the interface. A /32 netmask is used to specify the exact IP address of the interface

  *R1 knowns “If I receive a packet destined for this IP address, the message is for me”*

  192.168.1.1/32 matches onlly 192.168.1.1

### Route selection

假设现在 R3 发送了一个 Dst IP 192.168.1.1 的报文到 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-24_16-18.77wd0pog4d1c.webp)

R1 上匹配 192.168.1.1 的路由有两条

- 192.168.1.0/24
- 192.168.1.1/32

那条路由会被采用？Router 会采用 the most specific matching route

192.168.1.0/24 中包含 256 不同的 IP address, 192.168.1.0 - 192.168.1.255

192.168.1.1/32 中只包含 1 个 IP address, 192.168.1.1

所以 192.168.1.1/32 对应的路由 more specific，所以 Router 会使用该条路由。即 R1 会自己接受这个报文，而不是通过 G0/2 转发

> Most specfic mathing route = the matching route with the longest prefix length

你可能注意到除 connected route 和 local route 外还有一部分，这些并不是路由

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-24_16-30.5pii4eq16wlc.webp)

例如 

```
192.168.1.0/24 is vairably subnetted, 2 subnets, 2 masks
```

表示

there are two routes to subnets that fit within the 192.168.1.0/24 Class C network, with two different netmasks(/24 and /32)



**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=aHwAm8GYbn8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=19