# Day52 - LAN Architetures

## Terminologies

常见的网络拓扑有

1. Star 星型拓扑

   When several devices all connect to one central device we can draw them in a ‘star’ shap like below, so this is often called a ‘star topology’

   例如下图中所有的 end hosts 都连接到了一台 switch

   ![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_14-30.5bz0am1o8ho0.webp)

2. Full Mesh 网状拓扑

   when each device is connected to each other device

   例如下图中所有设备两两互联

   ![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_14-32.3uropjoxvoo0.webp)

3. Partial Mesh 部分网状拓扑

   when some devices are connected to each other, but not all

   ![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_14-34.5vny5xf7g5s0.webp)

## Tier

网络拓扑中的层，一般有 3 种

1. Access layer 接入层

   > 通常以 Star 和 end-hosts 互联

   - the layer that end hosts connect to(PCs, printers, cameras, etc)

   - typically access layer switches have lots of ports for end hosts to connect to
   - QoS marking is typically done here
   - Security services like port security, DAI, etc are typically performed here
   - switchports might be PoE-enabled for wireless APs, IP phones, etc

2. Distribution layer 汇聚层

   > 通常以 Partial Mesh 和 Access layer 设备互联
   >
   > 以 Full Mesh 和 其他 Distribution layer 设备互联
   >
   > 以 Star 和 其他 Core layer 设备互联

   - aggregates connections from the access layer switches

   - typically is the border between layer 2 and layer 3

     所以通常是一个 3 层交换机

   - connects to services such as Internet, WAN, etc

3. Core layer 核心层

   - connects Distribution layers together in large LAN networks
   - the focus is speed(fast transport)
   - CPU-intensive operations such as security, Qos marking/classification, etc. should be avoided at this layer
   - connections are all layer 3. No spanning-tree
   - should maintain connectivity throughtout the LAN even if devices fail

### Two-Tier topology

2 层拓扑结构通常用在 Campus 中

由 Access layer 和 Distribution layer 组成

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_14-45.m5dm3snor74.webp)

> 不过在 two-tier 架构中 Distribution layer 也承担了 Core layer 的功能，所以也被称为 Core-Distribution layer

### Three-Tier topology

当 Distribution layer 中设备多的时候，要满足 Distribution layer 设备互联的条件就会需要创建多条 connections，网络拓扑的扩展性就会非常差，所以这时候就需要 Core layer 来连接 Distribution layer

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_14-57.372w0yz2pn20.webp)

> 思科推荐当 Distribution layer 中设备超过 3 台时，就应该使用 Core 	layer

例如在下图中加入 Core layer 设备和 Distribution layer 设备构成 Star

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_15-00.7756omijerc0.webp)

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_15-08.4un8ldta4h00.webp)

## Spine-Leaf Architecture

在现在的互联网架构中通常会使用 Virtual server 来对外提供服务，当流量到了 VIP 对应的设备后，会做 East-West(横向)的转发

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_15-26.3kbcum32z0y0.webp)

*The traditional three-tier architecture led to bottlenecks in bandwidth as well as variability in the server-to-server latency depending on the path the traffic takes*

> 这里没明白为什么 three-tier 的瓶颈是带宽

为了解决 three-tier 的这个问题，通常会使用 Spine-Leaf architecture(也被称为 Clos architecture)

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_16-27.1b4juonat7og.webp)

> 和 Two-Tier 对比，就是 Spine 对应的 Distribution layer 不需要互联

## SOHO Networks

Small Office/Home Office(SOHO) Network 通常指的就是小型的网络结构

*SOHO networks don’t have complex needs, so all networking functions are typically provide by a single device, often called a ‘home router’ or ‘wireless router’*

SOHO network 通常只需要一台设备，也就是家用路由器，负责几个角色

- Router
- Switch
- Firewall
- Wireless Access Point(AP)
- Modem(光猫)

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_16-35.282ug4k9oxzw.webp)

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_16-39.7gifanay62c0.webp)

Configure HSRP on DSW1/DSW2, and ensure sychronization with STP 

### 0x01

in VLAN 10

DSW1 is HSRP active/STP root

DSW2 is HSRP standby/STP secondary root

配置前先看一下 STP 和 HSRP 的状态

```
DSW1#show spanning-tree vlan 10
VLAN0010
  Spanning tree enabled protocol ieee
  Root ID    Priority    32778
             Address     0001.C912.B090
             Cost        4
             Port        3(GigabitEthernet1/0/3)
             Hello Time  2 sec  Max Age 20 sec  Forward Delay 15 sec

  Bridge ID  Priority    32778  (priority 32768 sys-id-ext 10)
             Address     000C.856A.50BD
             Hello Time  2 sec  Max Age 20 sec  Forward Delay 15 sec
             Aging Time  20

Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Gi1/0/3          Root LSN 4         128.3    P2p
Gi1/0/1          Desg LSN 4         128.1    P2p
Gi1/0/2          Desg LSN 4         128.2    P2p

DSW1#show standby brief 
                     P indicates configured to preempt.
                     |
Interface   Grp  Pri P State    Active          Standby         Virtual IP
```

配置

```
DSW1(config)#spanning-tree vlan 10 root  primary 
DSW1(config)#int vlan 10
DSW1(config-if)#standby version 2
DSW1(config-if)#standby 10 ip 10.0.10.254
DSW1(config-if)#standby 10 priority 200
DSW1(config-if)#standby 10 preempt
DSW1(config-if)#exit
DSW1(config)#ip routing

DSW2(config)#spanning-tree vlan 10 root secondary 
DSW2(config)#int vlan 10
DSW2(config-if)#standby version 2
DSW2(config-if)#standby 10 ip 10.0.10.254
DSW2(config-if)#standby 10 preempt
DSW2(config-if)#exit
DSW2(config)#ip routing
```

### 0x02

in VLAN 20

DSW2 is HSRP active/STP root

DSW1 is HSRP standby/STP secondary root

```
DSW1(config)#spanning-tree vlan 20 root secondary 
DSW1(config-if)#standby version 2
DSW1(config-if)#standby 20 ip 10.0.20.254
DSW1(config-if)#standby 20 preempt

DSW2(config)#spanning-tree vlan 20 root primary 
DSW2(config-if)#standby version 2
DSW2(config-if)#standby 20 ip 10.0.20.254
DSW2(config-if)#standby 20 priority 200
DSW2(config-if)#standby 20 preempt
```

使用 `show spanning-tree vlan 10` 检查 STP 配置

```
DSW1(config-if)#do show spanning-tree vlan 10
VLAN0010
  Spanning tree enabled protocol ieee
  Root ID    Priority    24586
             Address     000C.856A.50BD
             This bridge is the root
             Hello Time  2 sec  Max Age 20 sec  Forward Delay 15 sec

  Bridge ID  Priority    24586  (priority 24576 sys-id-ext 10)
             Address     000C.856A.50BD
             Hello Time  2 sec  Max Age 20 sec  Forward Delay 15 sec
             Aging Time  20

Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Gi1/0/3          Desg FWD 4         128.3    P2p
Gi1/0/1          Desg FWD 4         128.1    P2p
Gi1/0/2          Desg FWD 4         128.2    P2p

DSW1(config-if)#do show spanning-tree vlan 20
VLAN0020
  Spanning tree enabled protocol ieee
  Root ID    Priority    24596
             Address     0001.C912.B090
             Cost        4
             Port        3(GigabitEthernet1/0/3)
             Hello Time  2 sec  Max Age 20 sec  Forward Delay 15 sec

  Bridge ID  Priority    28692  (priority 28672 sys-id-ext 20)
             Address     000C.856A.50BD
             Hello Time  2 sec  Max Age 20 sec  Forward Delay 15 sec
             Aging Time  20

Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Gi1/0/1          Desg FWD 4         128.1    P2p
Gi1/0/3          Root FWD 4         128.3    P2p
Gi1/0/2          Desg FWD 4         128.2    P2p
```

使用 `show standby br` 检查 HSRP

```
DSW2(config-if)#do show standby br
                     P indicates configured to preempt.
                     |
Interface   Grp  Pri P State    Active          Standby         Virtual IP
Vl10        10   100 P Standby  10.0.10.1       local           10.0.10.254    
Vl20        20   200 P Active   local           10.0.20.1       10.0.20.254    
```

从 PC1 ping PC2

```
C:\>ping 10.0.20.10

Pinging 10.0.20.10 with 32 bytes of data:

Reply from 10.0.20.10: bytes=32 time<1ms TTL=127
Reply from 10.0.20.10: bytes=32 time<1ms TTL=127
Reply from 10.0.20.10: bytes=32 time<1ms TTL=127
Reply from 10.0.20.10: bytes=32 time<1ms TTL=127

Ping statistics for 10.0.20.10:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 0ms, Maximum = 0ms, Average = 0ms
```

**references**

1. ^https://www.youtube.com/watch?v=PvyEcLhmNBk&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=101