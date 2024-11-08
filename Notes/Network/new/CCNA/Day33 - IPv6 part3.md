# Day33 - IPv6 part3

## IPv6 Header

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-01.4ovyzn98kr5s.webp)

IPv4 header 长度可以在 20 - 60 bytes 之间，而 IPv6 header 的长度固定在 40 bytes，所以在 IPv6 header 中就没有 header length 只一个字段用于表示 header 的长度。因为报文头都是固定的，对 router 而言处理效率就比 IPv4 高

- Version

  4 bits

  Indicates the version of IP that is used

  Fixed value of 6(0b0110) to indicate IPv6

- Traffic Class

  8 bits

  Used for Qos(Quality of Service), to indicate high priority traffic

  For example IP phone traffic, live video calls,etc, will have a Traffic Class value which gives them priority over other traffic

- Flow Label

  20 bits

  Used to identify specific traffic ‘flows’(communications between as a specific source and destination)

- Payload Length

  16 bits

  Indicates the length of the payload(the encapsulated Layer 4 segment) in bytes. The length of the IPv6 header itself isn’t included, because it’s always 40 bytes

- Next Header

  8 bits

  Indicates the type of the ‘next header’(header of the encapsulated sgement), for example TCP or UDP

  Same function as the IPv4 header’s ‘Protocol’ filed

- Hop limit

  8 bits

  The value in this filed is decremented by 1 by each router that forwards it. If it reaches 0, the packet is discarded

  Same function as the IPv4 header’s ‘TTL’ field

- Source Address/Destination Address

  128 bits each

  These fields contain the IPv6 addresses of the packet’s source and the packet’s intended destination

## Solicited-Node Multicast Address

Solicited-Node Multicast Address 是从对应接口的 Unicast Address 计算而来的

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-16.32716m6qpmio.webp)

通过取 unicast address 的最后 6 digits，拼接 ff02::1:ff 而来

例如

如果我们使用 `show ipv6 int <interface-name>` 可以看到不仅只有 ff02::1(all nodes) 和 ff02::2(all router) 加入到 multicast address group 中外，还有一个地址，就是 solicited-node multicast address

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-18.3mn1ui4f00zk.webp)

R1 G0/0 Unicast address 是 fe80::ef8:22ff:fe36:8500，所以取后 6 digits 为 36:8500，拼接后为 ff02::1:ff:36:8500

## Neighbor Discovery Protocol

IPv4 中通过 ARP 来学习 MAC address， 而在 IPv6 中通过 Neighbor Discovery Protocol(NDP) 来学习 MAC address，主要使用 ICMPv6 和 Solicicted-node multicast addresses 机制来实现

> ARP 通过 broadcast 来实现，而 NDP 通过 multicast 来实现，所以效率比 ARP 更高

通过发送两种 messages 来实现学习 MAC address 的功能

1. Neighbor Solicitation(NS) = ICMPv6 type 135

   Solicitation 中文为征集

2. Neighbor Advertisement(NA) = ICMPv6 type 136

例如

R1 想要 ping R2

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-27.5h2xstiurgqo.webp)

因为 R1 想要访问 R2, 就必须要知道对方的 MAC

1. 首先 R1 会发送 NS，到端口互联的链路上

   ![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-29.23jlbrnm2fi8.webp)

   这里的目的 IP 为 R2 的 solicited-node multicast address, 但是 R1 是怎么知道 R2 的 solicited-node multicast address 的呢？因为只要知道对端的 IPv6 地址就可以计算出来，这里 R1 访问 R2 是知道 R2 的 IPv6 地址的

   目的 MAC 是 Multicast MAC 基于 R2’s solicited-node address 计算出来，因为不在 CCNA 的考试范围内，这里只用于对比 ARP

2. R2 在收到 R1 发送过来的 NS，如果目的 IP 匹配就会从收到 NS 的端口链路上回送 NA

   ![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-35.qc83kkymdcg.webp)

   回送 NA 的逻辑和 ARP reply 类似

NDP 除了类似 ARP 学习 MAC address 的功能外，还可以自动的发现 routers on the local network

通过发送两种 messages 来实现这个功能

1. Router Solicitation(RS) = ICMPv6 type 133

   - sent to multicast address FF02::02(all routers)
   - ask all routers on the local link to identify themselves
   - sent when an interface is enable/host is connected to the network

2. Router Advertisement(RA) = ICMPv6 type

   - sent to multicast address FF02::1(all nodes)

   - the router announces its presence, as well as other information about the link(the prefix of the network)

   - these messages are sent in reponse to RS messages

     如果发送了 RA 就会回送 RS

   - they are also sent periodically, even if the router hasn’t received an RS

例如

R2 G0/0 配置了 IPv6 并 enable 

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-49.pprwn1j0cr4.webp)

1. R2 会自动发送 RS 到 G0/0 互联的链路上，询问链路上是否有其他的 routers
2. R1 收到 R2 发送过来的 RS，会回送 RA 到收到 RS 端口的链路上

> 通过这个功能，hosts 可以实现自动学习 default gateway 地址

### DAD

NDP 除支持上面两个功能外，还有一个 duplicate address detection(DAD) 的功能，用于自动检查 local link 中是否有一样的 IPv6 地址

*Any time an IPv6-enabled interface initializes(`no shutdown` command), or an IPv6 address is configured on an interface(by any method:manual, SLAAC,etc), it performs DAD*

DAD 通过发送 NS/NA 来实现

1. the host will send an NS to its own IPv6 address. If it doesn’t get a reply, it knows the address is unique
2. If it gets a reply, it means another host on the network is already using the address

如果在 Cisco 的设备上检查到了 duplicate address 会显示如下信息

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_21-11.1fusn6xrhjz4.webp) 

## IPv6 neighbor table

因为 IPv6 不使用 ARP，也就没有 arp table 这一说法。在 IPv6 中被称为 neighbor table

可以通过 `show ipv6 neighbor` 来查看 neighbor table

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-36.6cxjvozx8d1c.webp)

这里可以看到 R1 不仅学习了互联的 R2 g0/0 IPv6，还学习了 R2 link-local address

> 这里并不需要通过类似 ping 的命令让 R1 学习 R2 link-local address，在 IPv6 中会自动学习互联端口对应的 link-local address

- Age

  和 ARP 中一样，表示老化的时间，单位 minute

- Link-layer Addr

  表示对应的 IPv6 address 的 MAC address

- State

  表示状态和 ARP 中的一样

- Interface

  是从那个端口学来的

## SLAAC

Stateless Address Auto-configuration(SLAAC) 是另外一种配置 IPv6 的方式，可以让 Hosts ==通过 RS/RA== 自动学习 IPv6 prefix 并自动生成 IPv6 address

在没有 SLAAC 之前，我们可以使用 `ipv6 address <prefix/prefix-length> eui-64`  或者手动指定地址来配置 IPv6

但是在有 SLAAC 之后，我们可以使用 `ipv6 address autoconfig` 来自动学习 prefix，之后设备会使用 EUI-64 来生成 host portion 或者直接随机生成

例如

R1 和 R2 互联，R1 配置了 IPv6 address，但是 R2 没有配置 IPv6 address

如果 R2 使用了 `ipv6 address autoconfig` 就会自动生成一个 IPv6 address，当然也包括 link-local address

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_20-57.10p3w6esz1hc.webp)

> SLAAC 并不是 Cisco 独有，在 end host 上同样也可以使用 SLAAC 来自动配置 IPv6 address

## IPv6 Static Routing

IPv6 和 IPv4 routing 逻辑上相同，但是有几点细节有区别

1. IPv6 和 IPv4 使用单独的 routing table，可以使用 `show ipv6 route` 来查看

2. 默认启用 IPv4 routing，但是默认关闭 IPv6 routing，需要使用 `ipv6 unicast-routing` 来启用 IPv6 routing 的功能。如果没有启用 IPv6 routing, router 可以收到和发送 IPv6 traffic，但是不能在不同 network 之间转发流量

   > 只有记住使用 IPv6 就要使用 `ipv6 unicast-routing`

例如有如下拓扑

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_21-17.2egd5goahw8w.webp)

先看一下 R1 的 routing table

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_21-18.1juy3a3a5d1c.webp)

和 IPv4 一样，只要配置了 IPv6，就会自动添加两条路由，connected route 和 local route

- A connected network route is automatically added for each connected network
- A local network route is automatically added for each configured on the router

> FF00::/8 是 IPv6 multicast range, 关于这条路由不在 CCNA 考试的范围内，所以不过多介绍

==还需要注意的一点是。虽然会自动配置 link-local address，但是并不会有对应 link-local 的 connected route 或者是 local route==

### Configure IPv6 static route

如果需要配置 static route 可以使用 `ipv6 route destination/prefix-length {next-hop | exit-interface [next-hop]} [ad]`, 根据使用的参数不同可以分为几种

> 这里不分 IPv4 或者是 IPv6， IPv4 中同样有这几种类型的 route

1. Directly attached static route

   Only the exit interface is specified

   > 逻辑上是和当前 router 直联的，但是实际并不是

   对应的命令为 `ipv6 route destination/prefix-length exit-intereface`

   例如 `ipv6 route 2001:db8:0:3::/64 g0/0`

   *In IPv6, you can’t use directly attached static route if the interface is an Ethernet interface*

   > 如果是 serial interface 就可以使用 directly attached static route，但是 ethernet interface 不支持，但是并不影响 IPv4

   即 `ipv6 route 2001:db8:0:3::/64 g0/0` 实际是无需的，但是这个命名并没有语法错误，所以 router 会让你输入，==但是对应的 route 并不会生效==

2. Recursive static route

   Only the next hop is specified

   对应的命令为 `ipv6 route destination/prefix-length next-hop`

   例如 `ipv6 route 2001:db8:0:3::/64 2001:db8:0:12::2`

   > 因为 traffic 是通过端口转发的，知道 next-hop address 并不能直接知道本设备对应的端口是那个，所以同样需要通过查询 routing table

   ![](https://github.com/dhay3/image-repo/raw/master/20230625/2023-06-25_23-47.16jbu50idpz4.png)

   先找红框中的，然后按照 nexthop 找黄框中的

3. Fully specified static route

   Bothe the exit interface and next hop are specified

   对应的命令为 `ipv6 route destination/prefix-length exit-interface next-hop`

   例如 `ipv6 route 2001:db8:0:3::/64 g0/0 2001:db8:0:12::2`

除了按照使用的参数分外还可以按照目的地址分，同样在 IPv4 中也适用

1. Network route

   目的地址是一个段的

   例如 `ipv6 route 2001:db8:0:3::/64 2001:db8:0:12::2`

2. Host route

   目的地址是一个 unicast address

   例如 `ipv6 route 2001:db8:0:3::100/128 2001:db8:0:12::1`

   > 在 IPv6 中使用 128 prefix，在 IPv4 中使用 32 prefix

3. Default route

   目的地址是 wildcard address

   例如 `ipv6 route ::/0 2001:db8:0:23::1`

和 IPv4 floating static route 一样，IPv6 也有，通过 AD 来实现

### Link-Local Next-hops

在 day32 中的 lab，如果直接使用 link-local address 作为 recursive static route 中的 next-hop 会报错

![](https://github.com/dhay3/image-repo/raw/master/20230625/2023-06-26_00-07.3eof7oqt324g.webp)

需要使用 fully specified static route

如果使用 link-local address 为什么需要指定 exit-interface 才可以呢？

因为 link-local address 并不会加入 routing table 中，所以 router 不知道对应的路由，所以必须要通过指定 exit-interface 的方式告诉 router 对应的 link-local address 是那个接口互联的

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230625/2023-06-26_00-26.2uig3yc9yxc0.webp)

### 0x01

Enable IPv6 routing on each router

```
R1(config)#ipv6 unicast-routing 
R2(config)#ipv6 unicast-routing 
R3(config)#ipv6 unicast-routing 
```

### 0x02

Use SLAAC to configure IPv6 addresses on the PCs

只需要将 PC gateway 配置成 automatic

PC1

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::20A:41FF:FE4D:1BBC
   IPv6 Address....................: 2001:DB8:0:1:20A:41FF:FE4D:1BBC
   IPv4 Address....................: 0.0.0.0
   Subnet Mask.....................: 0.0.0.0
   Default Gateway.................: FE80::202:4AFF:FE23:E201
                                     0.0.0.0
```

PC2

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::240:BFF:FE69:9B18
   IPv6 Address....................: 2001:DB8:0:3:240:BFF:FE69:9B18
   IPv4 Address....................: 0.0.0.0
   Subnet Mask.....................: 0.0.0.0
   Default Gateway.................: FE80::290:2BFF:FECC:A101
                                     0.0.0.0
```

### 0x03

Configure static routes on the routers to allow PC1 and PC2 to ping each  other. The path via R2 should be used only as backup path

先用 `sh ipv6 route` 看一下 routing table

```
ND  ::/0 [2/0]
     via FE80::290:2BFF:FECC:A102, GigabitEthernet0/1
C   2001:DB8:0:1::/64 [0/0]
     via GigabitEthernet0/0, directly connected
L   2001:DB8:0:1::1/128 [0/0]
     via GigabitEthernet0/0, receive
C   2001:DB8:0:13::/64 [0/0]
     via GigabitEthernet0/1, directly connected
L   2001:DB8:0:13::1/128 [0/0]
     via GigabitEthernet0/1, receive
L   2001:DB8:0:13:202:4AFF:FE23:E202/128 [0/0]
     via GigabitEthernet0/1, receive
L   FF00::/8 [0/0]
     via Null0, receive
```

配置 R1

这里大写的地址是 R2 s0/0/0 对应的 link-local address

```
R1(config)#ipv6 route 2001:db8:0:3::/64 2001:db8:0:13::2
R1(config)#ipv6 route 2001:db8:0:3::/64 s0/0/0 FE80::20B:BEFF:FED7:4901 200
```

配置 R3

这里大写的地址是 R2 s0/0/1 对应的 link-local address

```
R3(config)#ipv6 route 2001:db8:0:1::/64 2001:db8:0:13::1
R3(config)#ipv6 route 2001:db8:0:1::/64 s0/0/0 FE80::20B:BEFF:FED7:4901 200
```

> 这里还可以发现 R2 s0/0/0 s0/0/1 link-local address 是相同的
>
> 即 link-local address 可以在不通接口间共用

配置 R2

这里第一个大写的地址是 R1 s0/0/0 对应的 link-local address

这里第二个大写的地址是 R2 s0/0/0 对应的 link-local address

```
R2(config)#ipv6 route 2001:db8:0:3::/64 s0/0/1 FE80::290:2BFF:FECC:A101
R2(config)#ipv6 route 2001:db8:0:1::/64 s0/0/0 FE80::202:4AFF:FE23:E201
```

然后查看 R1 routing table

```
ND  ::/0 [2/0]
     via FE80::290:2BFF:FECC:A102, GigabitEthernet0/1
C   2001:DB8:0:1::/64 [0/0]
     via GigabitEthernet0/0, directly connected
L   2001:DB8:0:1::1/128 [0/0]
     via GigabitEthernet0/0, receive
S   2001:DB8:0:3::/64 [1/0]
     via 2001:DB8:0:13::2
C   2001:DB8:0:13::/64 [0/0]
     via GigabitEthernet0/1, directly connected
L   2001:DB8:0:13::1/128 [0/0]
     via GigabitEthernet0/1, receive
L   2001:DB8:0:13:202:4AFF:FE23:E202/128 [0/0]
     via GigabitEthernet0/1, receive
```

只能看见一条 static route

可以看一下 tracert 的结果

```
C:\>tracert 2001:DB8:0:3:240:BFF:FE69:9B18

Tracing route to 2001:DB8:0:3:240:BFF:FE69:9B18 over a maximum of 30 hops: 

  1   0 ms      0 ms      0 ms      2001:DB8:0:1::1
  2   0 ms      0 ms      0 ms      2001:DB8:0:13::2
  3   0 ms      0 ms      0 ms      2001:DB8:0:3:240:BFF:FE69:9B18
```

这里可以看到回包的分别是 R1 G0/0 和 R3 G0/0

如果想要验证 via R2 是 backup route，可以关闭 R1 G0/1 或者直接删除 R1 R3 互联的 link 来测试，这是再查看 R1 routing table 并测试

```

C   2001:DB8:0:1::/64 [0/0]
     via GigabitEthernet0/0, directly connected
L   2001:DB8:0:1::1/128 [0/0]
     via GigabitEthernet0/0, receive
S   2001:DB8:0:3::/64 [200/0]
     via FE80::20B:BEFF:FED7:4901, Serial0/0/0
L   FF00::/8 [0/0]
     via Null0, receive
```

同样的也可以使用 tracert 来观察

```
C:\>tracert 2001:DB8:0:3:240:BFF:FE69:9B18

Tracing route to 2001:DB8:0:3:240:BFF:FE69:9B18 over a maximum of 30 hops: 

  1   0 ms      0 ms      0 ms      2001:DB8:0:1::1
  2   *         *         *         Request timed out.
  3   37 ms     1 ms      6 ms      2001:DB8:0:3::1
  4   0 ms      1 ms      0 ms      2001:DB8:0:3:240:BFF:FE69:9B18
```

这里可以发现第 2 条超时了，因为第二跳对应 R2 s0/0/0 并没有配置 IPv6 address，而是使用了 link-local address，因为 PC1 的 IP 不在一个 subnet 中所以不通

**referneces**

1. [jeremy’s IT Lab]: https://www.youtube.com/watch?v=rwkHfsWQwy8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=63