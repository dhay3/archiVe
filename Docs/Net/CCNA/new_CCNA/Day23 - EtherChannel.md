# Day23 - EtherChannel

## Why need EtherChannel

有如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_15-49.2qxr6ijifdz4.webp)

其中的两台交换机角色为

- ASW

  Access layer switch, a switch that end hosts connect to

- DSW

  Distribution layer switch, a switch that access layer switches connect to 

假设现在有 40 台 host 和 ASW1 互联(流量大概 3.5 Gbps)，ASW1 和 DSW1 通过一条 link(每条 link 1 Gbps) 互联

*The connection to DSW1 is congested. I should add another link to increase the bandwidth, so it can support all of the end hosts*

ASW1 和 DSW1 互联的链路一定会出现拥塞，因为链路带宽远小于 hosts 和 ASW1 互联的，所以需要增加 ASW1 和 DSW1 之间的 link

在 ASW1 和 DSW1 之间增加了一条 link

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_15-56.1veygrghgqow.webp)

*The connection to DSW1 is still congested. I’ll add more links*

> When the bandwidth of the interfaces connected to end hosts is greater than the bandwidth of the connection to the distribution switch(es), this is called oversubscription. Some oversubscription is acceptable, but too much will cause congestion

但是只增加一条 link 显然是不够的，所以增加到 4 条

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_15-59.1aycj7wik0dc.webp)

逻辑上 4 条 link 总带宽 4 Gbps 大于 3.5 Gbps，所以问题解决了吗？

*If you connect two switches together with multiple links, all except one will be disable by spanning tree*

*If all ASW1‘s interfaces were forwearding Layer 2 loops would form between ASW1 and DSW1, leading to broadcast storms*

*Other links will be unused unless the active link fails. In that case, one of the inactive links will start forwarding*

实际上还是只有 1 Gbps，因为 STP 的原因，为了防止 broadcast storm，ASW1 只会有一个端口可以转发正常的流量

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_16-08.6ys6irnigps0.webp)

## What is Etherchannel

Etherchannel 就是解决上面这个问题的

*Etherchannel can give us both redundancy and increased bandwidth*

Etherchannel 通常使用一个椭圆将 links 划为逻辑上的一条 link

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_16-17.5x5exq8zb5ds.webp)

- EtherChannel groups multiple interfaces togather to act as a single interface

- STP will treat this group as a single interface

  所以不会导致 2 层环，所以上面这个例子中就有 4 Gbps 的带宽，这点可以使用 `show spanning-tree` 来校验

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-57.1ekwt7tgaogw.webp)

  这里可以看到只有一个逻辑上的接口 Po1

- Traffic using the EtherChannel will be load balanced among the physical interfaces in the group. An algorithm is used to determine which traffic will use which physical interface 

使用了 Etherchannel 流量转发如下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_16-24.5g907k4e2gao.webp)

- Some other names for an EtherChannel are

  Port Channel

  LAG(Link Aggregation Group) 链路聚合

### Types of EtherChannel

在 Cisco 设备上有 3 种协议可以创建 EtherChannel

1.  PAgP(Port Aggregation Protocol)

   Cisco proprietary protocl

   思科独有的协议

   Dynamically negotiates the creation/maintenance of the EtherChannel.(like DTP does for trunks)

2. LACP(Link Aggregation Control Protocol)

   Industry standard protocol(IEEE 802.3ad)

   公有的标准协议

   Dynamically negotiates the creation/maintenance of the EtherChannel.(like DTP does for trunks)

3. Static EtherChannel

   A protocol isn’t used to determine if an EtherChannel should be formed.

   Interface are statically configured to form an EtherChannel

*Up to 8 interfaces can be formed into a single EtherChannel(LACP allows up to 16, but only 8 will be active, the other 8 will be standby mode, waiting for an active interface to fail)*

## Etherchannel Load-Balancing

有如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_16-42.4oe1uq4a054w.webp)

> 注意实际 end hosts(SRV1, printer) 不会和 DSW 连接，应该需要和 ASW 连接

现在 PC1 需要访问 printer，会使用 ASW1 那个端口呢？

会根据特定算法选择端口，主要有如下几个因素会影响算法选择端口

1. Source MAC
2. Destination MAC
3. Sources AND Destination MAC
4. Source IP
5. Destination IP
6. Source IP and Destination IP

> 但是有些 Switch 还支持 TCP/UDP 端口

例如

每次从 PC1 来的报文都走 ASW1 G0/0 转发到 DSW1 因为 Source MAC 固定

或者每次从 PC2 到 PR1 的报文都走 ASW G0/1 转发到 DSW1 因为 Source MAC 和 Destination 都固定

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_16-58.4p9r8xp4m29s.webp)

Etherchannel load balances 是基于 flows 的(理解成 TCP follow)，同一个 flow 只会从相同的链路发送

- Frames in the same flow will be forwarded using the same physical interface
- If frames in the same flow were forwarded using different physical interfaces, some frames may arrive at the destination out of order, which can cause problems

## Etherchannel Load-Balancing Configuration

可以使用 `show etherchannel load-balance` 来查看当前设备使用的 load-balance 决定因素

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_17-03.3jynkxj3udc0.webp)

上图就表示由 Source IP 和 Dstination IP 决定

例如从 10.0.0.1 到 10.0.0.2 的报文就会从 ASW1 同一个接口转发出去

可以使用 `port-channel load-balance <inputs>` 来修改 load balancing 的决定因素

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_17-05.25z4kglf9u3k.webp)

上图就会使用 src-dst-mac 决定 etherchannel 使用那个端口来转发报文

> 注意这里使用 port-channel 来修改 load balancing，使用 etherchannel 来查看使用的 load balancing

## Etherchannel Configuration

**查询的命令都是用 etherchannel, 增删改用 portchannel 或者 channel-group**

> 推荐使用 `int range` 来选中多个端口配置，因为使用 Etherchannel ==需要确保每个端口的 channel-group number 都是相同的==
>
> 但是对端的端口 channel-group number 可以不同
>
> 即 number 只在本机有效
>
> 例如
>
> ASW1 G0/1 和 G0/2 需要组成 Etherchannel，那么对应的 channel-group number 就必须一样，假设为 1；但是对端 DSW1 G0/1 和 G0/2 需要组成 Etherchannel 就不需要和 ASW1 Etherchannel 具有一样的 channel-group number，DSW1 可以使用任意一个 number，但是同样需要确保需要组成 Etherchannel 的端口 channel-group number 的值相同

可以使用 `show etherchannel summary` 来查看 etherchannel 的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-51.1ndyrblspgqo.webp)

SU 的含义看 code legend 部分，例如上述就表示 Port channel 1 是一个 2 层正在使用的中的逻辑口

(P) 表示是 port channel 聚合的口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-54.4a275cduudts.webp)

如果需要查看 port channel 具体使用的 mode 是什么，可以使用 `show etherchannel port-channel`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-56.6adzazb0fyps.webp)

例如上图中，port channel 中的端口均为 Active mode

### PAgP Configuration

可以使用 `channel-group <number> mode desireable` 或者 `channel-group <number> mode auto` 来配置 PAgP, 其中 mode auto 和 mode desirable 只适用于 PAgP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-13.6shpp9izs2kg.webp)

互联的两端端口需要配置成指定的状态，才会启用 Etherchannel(PAgP)，如下

```
auto + auto = no EtherChannel
desirable + auto = EtherChannel
desirable + desirable = EtherChannel
```

### LACP Configuration

LACP 大体上和 PAgP 类似，但是只能使用 

active mode = desirable mode

passive mode = auto mode

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-36.4lmsfyz40tk.webp)

互联的两端端口需要配置成指定的状态，才会启用 Etherchannel(LACP)，如下

```
passive + passive = no EtherChannel
active + passive = Etherchannel
active + active = Etherchannel
```

### Static EtherChannel Configuration

同样和 PAgP 类似，但是只能使用 on mode

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-40.3c6x2rpq7atc.webp) 

*On mode only works with on mode(on + desirable or on + active will not work)*

需要两端都配置成 on mode 才会生效

### Port-channel Configuration

配置 port-channel(channel-group 对应的逻辑接口) 本身

可以使用 `int port-channel <channel-group-number>` 来配置对应的 port-channel

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_21-46.7jbxvt6rfczk.webp)

## Cautions for EtherChannel

- Member interfaces must have matching configurations
  - Same duplex(full/half)
  - Same speed
  - Same switchport mode(access/trunk)
  - Same allowed VLANs/native VLAN(for trunk interfaces)
- If an interface’s configuration do not match the others, it will be exclueded from the EtherChannel

## Layer3 EtherChannel

有如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_22-01.77mt6mfomb9c.webp)

使用 Layer3 Switch(都使用了`no switch port`) 替代了 Layer2 Switch，那么就没有 2 层环，也就不需要使用 STP(因为 Router 并不会广播)，也就不存在 blocking port, 例子中 ASW1 和 DSW1 直接的链路带宽可以达到 4 Gbps，但是这样需要为每一个互联的 3 层口配置一个 IP，显然太麻烦了

同样我们还可以给 3 层口配置 EtherChannel

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_22-11.4yg3qaxkgqo0.webp)

> 这里需要为加入 Etherchannel 的端口声明 `no switchport` 
>
> 创建出来的 port-channel 自动会声明使用 `no switchport` 可以使用 `sh run` 来查看

如果使用了 `sh etherchannel summary` 就可以看到对应的端口显示为 3 层 port-channel

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_22-13.13x3jz12mvkw.webp)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_22-23.6gyvv2pum8ao.webp)

### 0x01

#### ASW1

```
ASW1(config)#int range g0/1-2
ASW1(config-if-range)#channel-group 1 mode active
ASW1(config-if-range)#int po1
ASW1(config-if)#switchport mode trunk
```

#### DSW1

```
DSW1(config)#int range G1/0/3-4
DSW1(config-if-range)#channel-group 1 mode active
DSW1(config-if-range)#int po1
DSW1(config-if)#switchport mode trunk
```

### 0x02

#### ASW2

```
ASW2(config)#int range g0/1-2
ASW2(config-if-range)#channel-group 1 mode desirable
ASW2(config-if-range)#int po1
ASW2(config-if)#switchport mode trunk
```

#### DSW2

```
DSW2(config)#int range G1/0/3-4
DSW2(config-if-range)#channel-group 1 mode desirable
DSW2(config-if-range)#int po1
DSW2(config-if)#switchport mode trunk
```

### 0x03

#### DSW1

```
DSW1(config-if-range)#int range G1/0/1-G1/0/2
DSW1(config-if-range)#channel-group 2 mode on
```

#### DSW2

```
DSW2(config-if-range)#int range G1/0/1-G1/0/2
DSW2(config-if-range)#channel-group 2 mode on
```

### 0x04

> 这里只需要对 port-channel 声明 no switchport 即可
>
> 一定要使用 `ip routing` 否则 route table 显示为空

#### DSW1

```
DSW1(config-if)#int po2
DSW1(config-if)#no switchport 
DSW1(config-if)#ip add 10.0.0.1 255.255.255.252
DSW1(config-if)#ip routing
DSW1(config)#ip route 172.16.2.0 255.255.255.0 10.0.0.2
```

#### DSW2

```
DSW2(config-if)#int po2
DSW2(config-if)#no switchport 
DSW2(config-if)#ip add 10.0.0.2 255.255.255.252
DSW2(config-if)#ip routing
DSW1(config)#ip route 172.16.1.0 255.255.255.0 10.0.0.1
```

### 0x05

```
show etherchannel load-balance
```

### 0x06

```
port-channel load-balance src-dst-ip
```

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=xuo69Joy_Nc&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=43