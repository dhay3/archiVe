# Day18 - VLANs part3

## Native VLAN on a router

Native VLAN 有一个好处就是因为不需要 VLAN tag，所以 frame 更小网络发送和接受的速度更快,在 part2 中的例子，将 Native VLAN 改成实际需要使用的 VLAN10

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_19-24.1iz8p1h6k1wg.webp)

现在配置 Router Native VLAN, 有两种方法

### method one

Use the command `encapsulation dot1q <vlan-id> native` on the router subinterface

```
R1(config)#int g0/0.10
R1(config-subif)#encapsualtion dot1q 10 native
```

VLAN20 PC .65 ping VLAN10 PC .1，在 R1 和 SW2 之间抓包

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_19-30.12ymwov55hc0.webp)

ICMP request 在 wireshark 中显示如下

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_19-32.45kpu0169agw.webp)

可以看到 Ethernet II Source 和  Type:IPv4 间多出了一部分，即 VLAN tag。Type 对应 0x8100 就表示是 VLAN(TPID);802.1Q Virtual LAN 的部分，其中主要观察 ID 字段，即 VLAN ID

现在看从 R1 发出到 SW2 的 ICMP request

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_19-37.5gc96nodo0zk.webp)

因为 VLAN10 是 R1 和 SW2 配置的 Native VLAN, 所以发出的报文不会包含 VLAN tag,一直到 VLAN10 PC .1 回送 ICMP reply 到 R1，R1 查路由重新打上 VLAN20 tag 到 VLAN20 PC .65

### method two

Configure the IP address for the native VLAN on the router’s physical interface(the `encapsulation dot1q <vlan-id>` command is not necessary)

```
R1(config)#no int g0/0.10
R1(config)#interface g0/0
R1(config-if)#ip add 192.168.1.62 255.255.255.192
```

使用方法二需要先删除想要配置 Native VLAN 对应的 subinteface,然后配物理口的 IP，就表示该对应段的内都是 Native VLAN。所有接口的配置如下，红框中的部分是默认自带的

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_19-50.59o772an29vk.webp)

> subinterface G0/0.20 G0/0.30 同样需要配置，只是去掉了 G0/0.10 的配置，表示 VLAN10 为 Native VLAN

## Layer3(Multilayer) Switches

在之前的课程中使用的都是 Layer2 Switches,但是现在还有一种 Switches 叫做 Layer3 Switches，用以下图例表示

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_19-53.38vofu5mcccg.webp)

Layer3 Switch 有如下功能

1. A multilplayer switch is capable of both switching and routing

2. It is ‘Layer 3 aware’

   传统的 Switch 只会识别到 Layer2，上层的报文头都不会被 encapsulation 或者是 de-encapsulation，但是 Layer3 Switch 可以对 Layer3 的报文头进行 encapsulation 或者是 de-encapsulation

3. You can assign IP addresses to its interfaces, like a router

4. You can create virtual interfaces for each VLAN, and assign IP addresses to those interfaces

   Switch 通过虚拟的接口来分配 IP(通过 software 来实现)

5. You can configure routes on it, just like a router

6. It can be used for inter-VLAN routing

> 在大型的网络中，如果 Router 和 Switch 直接频繁的传输大量的流量，会导致 network congestion。所以一般使用 Layer3 Switch 来替代 Router 和 Switch

## Inter-VLAN routing via SVI

现在使用 Layer3 Switch 替代 Layer2 Switch(SW2)

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-04.3lpknanij7eo.webp)

流量就不需要像上图一样走了，而是可以通过 SVIs(Switch Virtual Interfaces)

> 将 SVIs 想象成原来 Router 的 subinterfaces，是 Switch 中虚拟的一部分

- SVIs(Switch Virtual Interfaces) are the virtual interfaces you can assign IP addresses to in a multilayer switch
- Configure each PC to use the SVI(Not the router) as their gateway address
- To send traffic to different subnets/VLANs, the PCs will send the traffic to the switch, and the switch will route the traffic

这样流量就不需要走到 R1，SW2 做查路由并转发

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-15.5mnljrkx780.webp)

如果需要访问互联网，可以通过如下拓扑，在 R1 和 SW2 之间建立 P2P(30位 subnet mask) 连接

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-17.3cwf5kxghmps.webp)

可以通过入向配置来实现

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-19.umm703v0o80.webp)

需要注意的是即使 subinterface 使用了 `no` 命令来取消，在 `show ip int br` 中仍会显示，以 deleted 标识，只有 reload router 才会消失

现在看一下 SW2 的配置

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-23.37pmun8jkdts.webp)

1. 首先需要通过 `default interface g0/1` 将配置了 subinterfaces 的端口还原（这里假设 SW2 前一天中的拓扑结构，实际是一个 Layer3 Switch）

2. 使用 `ip routing` 声明当前的 Switch 可以使用 layer3 routing

3. 对需要使用 layer3 routing 的端口使用 `no switchport` 来标识当前接口属于 Layer3，可以配置 IP

   可以使用 `show interface status` 来查看，如果是用做 routing 的接口会在 Vlan 列有 routed 的标识

接下来配置 default route

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-28.4y0916s6tb40.webp)

这里就可以看到 SW2 有了 route table

下面配置 SVIs 

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-38.5b9i7pdwv1c0.webp)

值需要使用 `vlan<id>` 接口即可，需要注意的一点是因为 SVIs 默认是 shutdown 的，所以需要使用 `no shutdown` 命令来开启端口

假设现在配置一个 VLAN40 SVI，并使用了 `no shutdown` 命令，状态任然是 down/down

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-41.1pc64s104u1s.webp)

因为 VLAN40 并不存在 SW2 上，如果需要 up/up 要满足如下几点

- The VLAN must exist on the switch

  需要注意的是，在 Switch 上当 VLAN 不存在的时候，会自动创建 VLAN;但是在 Layer3 Switch 上创建 SVI，是默认不会创建 VLAN 的

- The switch must have at least one access port in the VLAN in an up/up state, AND/OR one trunk port that allows the VLAN that is in a up/up state

  ![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-49.6keun8ig4glc.webp)

  以这个拓扑为例，因为 G1/0 是 VLAN 20 access port，所以 SVI vlan20 up/up; 因为 G0/2 和 G0/3 VLAN10 access port，所以 SVI vlan10 up/up; 因为 G0/0 是 VLAN 10/30 trunk port，所以 SVI vlan30 up/up;因为没有 VLAN40 access port 或者是 trunk port 所以 VLAN40 down/down

- The VLAN must not be shutdown(you can use the `shutdown` command to disable a VLAN)

  这里并不是指的是 SVI，而是指的是 VLAN(access port),在 packet tracer 中不能测试

最后在 SW2 上的路由显示如下，Connected route 和 Local route 和之前的配置 subinterfaces 的 R1 一样

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_20-53.58ewmrgxt2ps.webp)

## Lab

> 基于 Day17 中的配置

![](https://github.com/dhay3/image-repo/raw/master/20230530/2023-05-30_21-40.3e4c3cn9nbpc.webp)

### R1

```
R1>en
R1#>conf t
R1(config)#no interface g0/0.10
R1(config)#no interface g0/0.20
R1(config)#no interface g0/0.30
R1(config)#int g0/0
R1(config-if)#ip add 10.0.0.194 255.255.255.252
R1(config-if)#no shutdown
```

### SW2

```
SW2>en
SW2#conf t
SW2(config)#ip routing
SW2(config)#default int g1/0/2
SW2(config-if)#int g1/0/2
SW2(config-if)#no switchport
SW2(config-if)#ip add 10.0.0.193 255.255.255.252
SW2(config-if)#no shutdown
SW2(config-if)#int vlan10
SW2(config-if)#ip add 10.0.0.62 255.255.255.192
SW2(config-if)#no shutdown
SW2(config-if)#int vlan20
SW2(config-if)#ip add 10.0.0.126 255.255.255.192
SW2(config-if)#no shutdown
SW2(config-if)#int vlan30
SW2(config-if)#ip add 10.0.0.190 255.255.255.192
SW2(config-if)#no shutdown
SW2(config-if)#ip route 0.0.0.0 0.0.0.0 10.0.0.194
```

这里需要注意的一点是，`sh ip route` 只会显示 Connected route 但是不会显示 Local route，是因为 packet tracer 的原因，正常应该都会显示

```
SW2#sh ip route
Codes: C - connected, S - static, I - IGRP, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default, U - per-user static route, o - ODR
       P - periodic downloaded static route

Gateway of last resort is 10.0.0.194 to network 0.0.0.0

     10.0.0.0/8 is variably subnetted, 4 subnets, 2 masks
C       10.0.0.0/26 is directly connected, Vlan10
C       10.0.0.64/26 is directly connected, Vlan20
C       10.0.0.128/26 is directly connected, Vlan30
C       10.0.0.192/30 is directly connected, GigabitEthernet1/0/2
S*   0.0.0.0/0 [1/0] via 10.0.0.194
```

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=OkPB028l2eE&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=34