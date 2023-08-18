# Day29 - First Hop Redundancy Protocols

## Why need/What FHRP

在 STP 中，已经说明了网络 Redundancy 的重要性

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-23.74nvmvmb4lfk.webp)

如果 R1 到 互联网的连接有问题，end hosts 同样还是可以通过 R2 访问互联网

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-25.21ic2eh7mlmo.webp)

现在假设所有的 end hosts 都指定使用 R1 .254 作为 default gateway，当 R1 挂掉后，这些 end hosts 需要怎么访问公网?

因为 end hosts 并不知道 R1 挂了，default gateway 任然会使用 R1 .254，所以 end hosts 就不能通过 R2 正常的这台 Router 访问公网

所以我们就需要一个协议，来让当 R1 出现问题时，会自动切换到 R2。这也就是 First Hop Redundancy Protocol(FHRP)

*A first hop redundancy protocl(FHRP) is a computer networking protocol which is designed to protect the default gateway used on a subnetwork by allowing two or more routers to provide backup for that address; int the event of failure of an active router, the backup router will take over the address, usually within a few seconds*

如果使用了 FHRP，两台 router 会共享一个 VIP，可以将 end hosts 的 default gateway 配置成这个 IP，而不是 R1/2 任意的一个物理接口 IP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-34.5juz7wq67uo0.webp)

那 routers 之间怎么来共享这个 VIP 呢？

他们之间会通过 multicast 发送 Hello message

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-36.601sbncdrmrk.webp)

通过 Hello message 来协商主备，主 router 被称为 Active router，备 router 被称为 Standby router

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-38.6bmz8gmsu800.webp)

Active router 作为首选，即如果 R1 没有问题 end hosts 会先走 R1； Standby router 作为次选，即如果 R1 有问题 end hosts 会走 R2

假设现在 PC1 想要访问公网的机器

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-41.14t01esi3p8g.webp)

因为 PC2 访问公网的机器不在同 LAN，所以需要通过 default route，对应的 default gateway 为 172.16.0.252，但是目前还不知道对应的 MAC，所以会先做 ARP request 学习 default gateway MAC

ARP request 会做 2 层的广播，所以 R1/2 都会收到 ARP request(这里因为 STP，所以 Switches 有些端口并不会转发 ARP request)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-43.41xntz46wozk.webp)

这时只有 R1 会回送 unicast ARP reply 到 PC1，因为 R1 是 Active router

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-48.jthivew8x4g.webp)

> 注意这里回送的 MAC address 是 virtual 的，并不是实际物理端口对应有的 MAC address

PC1 收到 R1 回送的 ARP reply，假设需要 PC1 实际需要访问的是 8.8.8.8，那么对应的请求报文如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-50.6us4mar91q80.webp)

现在 R1 突然挂了，那么 R2 就收不到从 R1 来的 hello message，R2 就会认为自己就是 Active router

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-52.2hm3a92814hs.webp)

现在 R2 是 Active router，那么 R2 需要通过什么方式向其他设备宣告，将流量发送到 R2 呢？

因为 R1/2 共享一个 VIP，end hosts 会使用这个 VIP 作为 default gateway

在 PC1 中 ARP table 如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_17-54.2jpwpg880pq8.webp)

end hosts 并不需要修改什么东西，因为对应的 MAC address 都是虚拟的。那么什么设备需要修改呢？

Switches 需要，因为 R1 挂掉了那么就不能从 R1 对应的 .254 端口学到 MAC，所以需要更新 MAC address table

R2 会发送 source MAC address of the virtual MAC address 的 ARP 报文，告诉网络内的所有设备 Virtual IP address 现在的 Virtual MAC 在 R2 —— 被称为 gratuitous(免费的) ARP

*Gratuitous ARP: ARP replies send without being requested(no ARP request message was received)*

> 即发送 gratuitous ARP 不需要像正常在收到 ARP request 后才发送，而是可以直接发送 ARP reply
>
> ==如果 R1 挂了 R2 会发送 gratuitous ARP，如果 R1 重新上线了 R1 也会发送 gratuitous ARP==

*The destination MAC address of Gratuitous ARP if FFFF.FFFF.FFFF(normal ARP replies are unicast)*

> 但是和普通的 ARP reply 不一样，gratuitous ARP 的目的 MAC address 是 2 层广播地址，所以会被所有的 Switches 广播(==Switches 也会更新自己的 MAC address table==)；而普通 ARP reply 则是发送 ARP request 的机器 MAC，即 unicast

R2 发送 gratuitous ARP 过程如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_18-05.1idvhtwtyxhc.webp)

现在 PC1 还要访问 8.8.8.8，那么就会通过 R2

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_18-08.3fzhky29d1fk.webp)

假设现在 R1 重新上线，并正常了；R1 会因为之前是 Active router 还会保持是 Active router 吗？

答案是不会

*FHRPs are ‘non-preemptive’. The current active router will not automatically give up its role, even if the former active router return*

> 这点和 OSPF 选择 DR/DBR 的逻辑一样，都是非抢占式的

*You can change this setting to make R1 ‘preempt’ R2 and take back its active role automatically*

> 和 OSPF 不同的是，FHRP 是可以手动设置是否抢占

实现 FHRP 的协议，一共有 3 种

1. HSRP
2. VRRP
3. GLBP

## HSRP

Hot Standby Router Protocol

是 Cisco 独有的协议，有两个版本

1. version 1

   - version 1 uses 224.0.0.2 for multicast

   - version 2 uses 0000.0c07.acXX(XX = HSRP group number) for Virtual MAC address

     例如 group 1，那么对应的 MAC address 就为 0000.0c07.ac01

2. version 2

   - version 2 uses 224.0.0.102 for multicast

   - version 2 uses 0000.0c9f.fXXX(XXX = HSRP group number) for Virtual MAC address

     例如 group 1，那么对应的 MAC address 就为 0000.0c9f.f001

   - version 2 adds IPv6 support and increases the number of groups that can be configured

*In a situation with multiple subnets/VLANs, you can configure a different active router in each subnet/VLAN to load balance*

> 和 STP 类似，HSRP 在不同的 VLAN 中也可以配置不同的 active router

例如

PC1/3 在 VLAN 1,  R1/2 HSRP virtual IP 1.252, .253 为 Router on stick 需要的 3 层地址(sub-interface)，在 VLAN 1 中 R1 是 Active Router 而 R2 Standby Router

PC2/4 在 VLAN 2,  R1/2 HSRP virtual IP 2.252, .253 为 Router on stick 需要的 3 层地址(sub-interface)，在 VLAN 2 中 R2 是 Active Router 而 R1 Standby Router

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_18-30.1g706qpmwujk.webp)

> VLAN1 用 R1 作为 default gateway
>
> VLAN2 用 R2 作为 default gateway
>
> 先到 virtual IP 然后内部转发到 sub-interface IP

## VRRP

Virtual Router Redundancy Protocol

是 Open standard

- 和 HSRP 不同，在 VRRP 中 Active Router = Master Router, Standby Router = Backup Router
- 只有一个多播地址 224.0.0.18
- Virtual MAC address: 0000.5e00.01XX(XX = VRRP group number)

*In a situation with multiple subnets/VLANs, you can configure a different active router in each subnet/VLAN to load balance*

> 这点和 HSRP 一样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_13-53.44l671pmehog.webp)

和 HSRP 中例子的拓扑一样，只是替换了 Active 为 Master，Standby 为 Backup

## GLBP

Gateway Load Balancing 和 HSRP 一样也是 Cisco 独有协议

- Load balances among multiple routers with a single subnet

  例如 PC1/2 都在 VLAN1， PC1 可以使用 R1 作为 default gateway，而 PC2 可以使用 R2 作为 default gateway

  > 这点和 HSRP 或者是 VRRP 不一样，在这两种协议中，一个 VLAN 或者 subnet 中只有一台 Router 可以做为 default gateway 

- Multicast address: 224.0.0.102
- An **AVG(Active Virtual Gateway)** is elected
- Up to four **AVFs(Active Virtual Forwarders)** are assigned by the AVG(the AVG can be an AVF, too). Each AVF acts as the default gateway for a portion of the hosts in the subnet
- Virtual MAC address: 0007.b400.XXYY(XX = GLBP group number, YY = AVF number)

## Comparing HSRP 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-07.55kwbhyfe328.webp)

## Configuring HSRP

配置如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-24.6iaji0g29jpc.webp)

先看 R1 的配置，需要确认 Router 那个互联的端口作为 default gateway

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-25.klcobkuwhzk.webp)

> 默认会使用 HSRP version1，如果使用 HSRP version2 group number 范围在 0-4095

可以使用 `standby version 2` 来使用 version2 HSRP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-27.1kcrokd1zx6o.webp)

这里只有一个 subnet(VLAN1)，所以值需要配置一个 HSRP group

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-31.2i5ai2t3onuo.webp)

> 为了方便记忆，和 Router on a stick 配置 router 的 sub-interface 一样，group number 可以使用 VLAN number
>
> ==如果不指定 group number 默认会使用 group 0==

可以使用 `standby <group number> ip <ip address>` 来配置 HSRP virtual IP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-33.6qeagij6o2rk.webp)

这里还使用了 `standby <group number> priority <number>`设置了 priority

priority 用于决定 Active router

The active router is determined in the order

1. Highest priority(==default 100==)
2. Highest IP address

代码的逻辑如下, a b 分别为两台 router 和 subnet 互联的接口

```
if a.priority > b.priority then:
	a = active
else if a.priority < b.priority then:
	b = active
else if a.priority == b.priority then:
	if a.ip > b.ip then:
		a = active
	else if a.ip < b.ip then:
		b = active
```

上面还使用了 `standy <group number> preempt` 来让 R1 做抢占式成为 active router

例如

正常情况下 R1 是 Active router，R2 是 Standby router；当 R1 挂掉后，R2 会成为 Active router。当 R1 重新上线后，在不使用 preempt 命令(即默认 non-preempt)，R2 任然会是 Active router，而 R1 是 Standy router，即使 R1 有更大的 priority 或者是 IP；如果使用了 preempt 命令，R1 重新上线后就会抢占 R2 Active router 的位置，成为 Active router(因为 priority 和 IP 都比 R2 大)，R2 变回 Standby router

*Preempt cause the router to take the role of the active router, even if another router already has the role*

> 可以简单的理解成，当使用 preempt 命令后，有新的 router 加入到 HSRP group 中，就会触发一次判断那台是 Active router 的逻辑

看一下 R2 的配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-49.6hxxs0kbrxts.webp)

大部分和 R1 类似，但是需要注意一点

*HSRP version 1 and version 2 are not compatiable. If R1 uses version 2, R2 must use version 2 also*

> 构成 HSRP group 的 router，HSRP version 必须相同

在配置完 R1/2 后，可以使用 `show standy` 来查看 HSRP 相关信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_14-52.43nlg9c846io.webp)

- `Group 1(version 2)`

  接口所在的 HSRP group 是 group 1，使用 HSRP version 2，Active Standby 均相同

- `State is Active/Standby`

  当前 router 是 Active/Standby

- `Virtual IP address is`

  HSRP 使用的 Virtual IP，Active Standby 均相同

- `Virtual/Local MAC address is`

  HSRP 使用的 Virtual MAC，Active Standby 均相同

- `preemption enabled`

  是否使用强制式

- `Active router is local`

  表明当前 router 是 active

- `Active router is 172.16.0.253，priority 200`

  Active router 互联的地址，Active router 使用的 priority

- `Standby router is 172.16.0.252, priority 50`

  Standby router 互联的地址，Standby router 使用的 priority

- `priority 200(configured 200)`

  当前 router 使用的 priority

## LAB

### 0x01

#### Ping exxternal server 8.8.8.8 from PC1/PC2.What is the default gateway configured as?

PC1 10.0.1.253 R1

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::201:42FF:FEB2:53DD
   IPv6 Address....................: ::
   IPv4 Address....................: 10.0.1.1
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: ::
                                     10.0.1.253
```

PC2 10.0.1.253 R1

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::260:3EFF:FE30:BB5
   IPv6 Address....................: ::
   IPv4 Address....................: 10.0.1.2
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: ::
                                     10.0.1.253

```

### 0x02

#### Configure HSRPv2 on R1/R2.Rasie R1’s priority above the default, lower R2’s priority below the default.Enable preemption

```
R1(config)#int g0/0
R1(config-if)#standby version 2
R1(config-if)#standby 1 ip 10.0.1.254
R1(config-if)#standby 1 priority 200
R1(config-if)#standby 1 preempt 

R1(config)#int g0/0
R1(config-if)#standby version 2
R1(config-if)#standby 1 ip 10.0.1.254
R1(config-if)#standby 1 priority 50
R1(config-if)#standby 1 preempt 
```

可以使用 `show standby` 来查看是否配置成功

```
R1(config-if)#do show standby
GigabitEthernet0/0 - Group 1 (version 2)
  State is Active
    6 state changes, last state change 00:10:21
  Virtual IP address is 10.0.1.254
  Active virtual MAC address is 0000.0C9F.F001
    Local virtual MAC address is 0000.0C9F.F001 (v2 default)
  Hello time 3 sec, hold time 10 sec
    Next hello sent in 1.276 secs
  Preemption enabled
  Active router is local
  Standby router is 10.0.1.252
  Priority 200 (configured 200)
  Group name is hsrp-Gig0/0-1 (default)
```

### 0x03

Cofnigure the VIP as the default gateway of PC1/PC2. Ping 8.8.8.8 from the PCs. Check the PCs’ ARP table. What MAC address is mapped to the VIP?

PC1

```
C:\>arp -a
  Internet Address      Physical Address      Type
  10.0.1.253            00d0.585b.7501        dynamic
  10.0.1.254            0000.0c9f.f001        dynamic
```

PC2

```
C:\>arp -a
  Internet Address      Physical Address      Type
  10.0.1.253            00d0.585b.7501        dynamic
  10.0.1.254            0000.0c9f.f001        dynamic
```

如果在 PC1/2 上使用 `tracert` 可以发现一个有有趣的现象

```
C:\>tracert 8.8.8.8

Tracing route to 8.8.8.8 over a maximum of 30 hops: 

  1   0 ms      0 ms      0 ms      10.0.1.253
  2   0 ms      0 ms      0 ms      8.8.8.8

Trace complete.
```

即使使用了 10.0.1.254 HSRP virtual IP，但是 tracert 显示回包的是 default gateway 实际的物理接口 IP

> 我们可以通过 tracert 和 default gateway 来简单的判断 router 是否使用了 FHRP

### 0x04

Turn of R1(save the config first!). After it restarts, ping from PC1 to 8.8.8.8 again. IS R2 used as the default gateway

会使用 R2 作为 default gateway 使用 `show standby` 查看

```
R2#show standby
GigabitEthernet0/0 - Group 1 (version 2)
  State is Active
    6 state changes, last state change 00:18:44
  Virtual IP address is 10.0.1.254
  Active virtual MAC address is 0000.0C9F.F001
    Local virtual MAC address is 0000.0C9F.F001 (v2 default)
  Hello time 3 sec, hold time 10 sec
    Next hello sent in 1.246 secs
  Preemption enabled
  Active router is local
  Standby router is unknown
  Priority 50 (configured 50)
  Group name is hsrp-Gig0/0-1 (default)
```

这里可以看见 R2 已经是 Active router，如果使用了 `tracert` 还可以看见回包的是 R2 互联的接口 IP

```
C:\>tracert 8.8.8.8

Tracing route to 8.8.8.8 over a maximum of 30 hops: 

  1   0 ms      0 ms      0 ms      10.0.1.252
  2   0 ms      0 ms      0 ms      8.8.8.8

Trace complete.
```

### 0x05

Trun on R1 again. Does it become the active router again?

R1 会重新变为 active router，因为使用了 preempt，同样的可以使用 `show standby` 来查看

```
R1#show standby
GigabitEthernet0/0 - Group 1 (version 2)
  State is Active
    7 state changes, last state change 00:00:27
  Virtual IP address is 10.0.1.254
  Active virtual MAC address is 0000.0C9F.F001
    Local virtual MAC address is 0000.0C9F.F001 (v2 default)
  Hello time 3 sec, hold time 10 sec
    Next hello sent in 0.653 secs
  Preemption enabled
  Active router is local
  Standby router is unknown
  Priority 200 (configured 200)
  Group name is hsrp-Gig0/0-1 (default)
```

同样也可以使用 `tracert` 来检查

```
C:\>tracert 8.8.8.8

Tracing route to 8.8.8.8 over a maximum of 30 hops: 

  1   0 ms      0 ms      0 ms      10.0.1.253
  2   0 ms      0 ms      0 ms      8.8.8.8

Trace complete.
```

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=43WnpwQMolo&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=55

