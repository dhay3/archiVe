# OSPF

reference：

https://www.jannet.hk/open-shortest-path-first-ospf-zh-hant/

https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/iproute_ospf/configuration/xe-16/iro-xe-16-book/iro-cfg.html

https://www.cisco.com/c/en/us/support/docs/ip/open-shortest-path-first-ospf/7039-1.html

## 概述

Open Shortest Path First是一个基于链路状态(Link-State Routing Protocol)的IGP协议，由于RIP的种种限制，RIP组件被OSPF取代。每个在OSPF里的Router都会向Neighbor交换自己的Link-State，当Router收到这些Link-State之后，就会运用Dijkstra Alogrith来计算出最短的路径

## Router Type

![2021-11-08_23-24](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211108/2021-11-08_23-24.18o156168r1c.png)

### Internal Router

Router上的所有interface都属于同一个area，R3和R5就是Internal Router

### Backbone Router

最少一个interface连接Backbone Area（Area 0），所以R2、R3和R4都是Backbone Router

### Area Border Router(ABR)

连接两个Area或以上的Router称为ABR，R2和R4都是ABR

### Autonomous System Border Router(ASBR)

有interface连接其他AS的Router就是ASBR，在这个网络中，除了执行OSPF之外，最右边蓝色的区域正在执行另一个Routing Protocol RIP。R6为OSPF 与 RIP的连接点，即连个AS的连接点，称为ASBR

## Link-state Advertisement (LSA)

OSPF是一个link-state routing protocal（链路状态协议），link简单的说就是router的interface(==也就是说如果interface不可用，就不会出现在`show ip ospf neighbor`的信息中，如果interface可用了，又会重新出现在`show ip ospf neighbor`中==)。OSPF的运作原理就是把自己每个interface连接的network告诉其他的router，于是router便可以计算出自己的routing table

LSA中包含几项信息：LSA由谁传送出来、连着的是什么network、以及去这个network的cost

LSA有不同的类型，用下面的一个例子说明

![2021-11-24_00-30](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211123/2021-11-24_00-30.62tmmwf8jrg0.png)

每只router都预设了loopback interface

R1为1.1.1.1 R2为2.2.2.2 以此类推

```
hostname R1
!
router ospf 1
 network 192.168.12.0 0.0.0.255 area 10
```

```
hostname R2
!
router ospf 1
 network 192.168.12.0 0.0.0.255 area 10
 network 192.168.23.0 0.0.0.255 area 0
 network 192.168.28.0 0.0.0.255 area 0
```

```
hostname R3
!
router ospf 1
 network 192.168.23.0 0.0.0.255 area 0
 network 192.168.34.0 0.0.0.255 area 0
```

```
hostname R4
!
router ospf 1
 network 192.168.34.0 0.0.0.255 area 0
 network 192.168.48.0 0.0.0.255 area 0
 network 192.168.45.0 0.0.0.255 area 20
```

```
hostname R5
!
router ospf 1
 network 192.168.45.0 0.0.0.255 area 20
 network 192.168.56.0 0.0.0.255 area 20
```

R6比较特别，因为R6是ASBR，需要使用redistribute指令把RIP的Route加入到OSPF中

```
hostname R6
!
router ospf 1
 redistribute rip subnets
 network 192.168.56.0 0.0.0.255 area 20
!
router rip
 version 2
 redistribute ospf 1 metric 1
 network 192.168.67.0
```

R7则是纯粹的RIP，并非在OSPF network中

```
hostname R7
!
router rip
 version 2
 network 192.168.67.0
```

```
hostname R8
!
router ospf 1
 network 192.168.28.0 0.0.0.255 area 0
 network 192.168.48.0 0.0.0.255 area 0
```

### router LSA(Type 1)

router LSA指在==同一个Area==里所有的Router(包括自己)送来的Link-State

要看router LSA可以在Router输入指令`show ip ospf [process-id] database router`，以下用R1收到的两条router LSA，留意Advertising Router （router-id 或 最大的loopback ip），第一条从1.1.1.1 即自己发出，第二条从2.2.2.2 发出。

集中观察第二条，看看2.2.2.2即R2告诉R1什么

1. 我是一只ABR（Area Border Router）

2. 我要告诉你连着一条link，就是192.168.12.2(R2上的iface)，我到哪里需要Cost 10（TOS 0 Metrics：10）

   ```
   R1#show ip ospf 1 database router
   
               OSPF Router with ID (1.1.1.1) (Process ID 1)
   
                   Router Link States (Area 10)
   
     LS age: 234
     Options: (No TOS-capability, DC)
     LS Type: Router Links
     Link State ID: 1.1.1.1
     Advertising Router: 1.1.1.1
     
     LS Seq Number: 80000002
     Checksum: 0x1C18
     Length: 36
     Number of Links: 1
   
       Link connected to: a Transit Network
        (Link ID) Designated Router address: 192.168.12.1
        (Link Data) Router Interface address: 192.168.12.1
         Number of TOS metrics: 0
          TOS 0 Metrics: 10
   
   
     Routing Bit Set on this LSA
     LS age: 254
     Options: (No TOS-capability, DC)
     LS Type: Router Links
     Link State ID: 2.2.2.2
     Advertising Router: 2.2.2.2
     
     LS Seq Number: 80000002
     Checksum: 0xE049
     Length: 36
     Area Border Router
     
     Number of Links: 1
   
       Link connected to: a Transit Network
        (Link ID) Designated Router address: 192.168.12.1
        
         Number of TOS metrics: 0
          TOS 0 Metrics: 10
   ```

### network LSA(Type 2)

network LSA是有==每个网段的DR==发给其他Router，告诉他们DR正连着那些Router

要看 network LSA，可以在router中输入`show ip ospf database network`，以下用R1做例子。因为R2是DR，所以R1收到了R2传来的network LSA，在network LSA可以看见R2正连着1.1.1.1和2.2.2.2(R2自己)(Attached Router)

```
R1#show ip ospf neighbor

Neighbor ID     Pri   State           Dead Time   Address         Interface
2.2.2.2            1   FULL/DR 		   00:00:35    192.168.12.2    Ethernet0/0
R1#
R1#show ip ospf 1 database network

            OSPF Router with ID (1.1.1.1) (Process ID 1)

                Net Link States (Area 10)

  Routing Bit Set on this LSA
  LS age: 26
  Options: (No TOS-capability, DC)
  LS Type: Network Links
  Link State ID: 192.168.12.2 (address of Designated Router)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000001
  Checksum: 0x8F1F
  Length: 32
  Network Mask: /24
  Attached Router: 2.2.2.2
  Attached Router: 1.1.1.1
        
  
```

### network summary LSA (Type 3)

==由ABR产生==，告诉Area内的Router从他那里可以到达那些network以及cost

要看network summary 可以在router输入`show ip ospf database summary`，以R1做例子，R1收到6条network summary LSA，全部来自R2(2.2.2.2)，R2告诉R1：

1. 从我这可以到达192.168.23.0(summary Network Number)，cost是10(Metric 10)
2. 从我这可以到达192.168.28.0，cost是10
3. 从我这可以到达192.168.34.0，cost是20
4. 从我这可以到达192.168.45.0，cost是30
5. 从我这可以到达192.168.48.0，cost是20
6. 从我这可以到达192.168.56.0，cost是40

```
R1#show ip ospf 1 database summary

            OSPF Router with ID (1.1.1.1) (Process ID 1)

                Summary Net Link States (Area 10)

  Routing Bit Set on this LSA
  LS age: 207
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(Network)
  Link State ID: 192.168.23.0 (summary Network Number)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xFFA9
  Length: 28
  Network Mask: /24
        TOS: 0  Metric: 10
  
  Routing Bit Set on this LSA
  LS age: 125
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(Network)
  Link State ID: 192.168.28.0 (summary Network Number)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xFFA9
  Length: 28
  Network Mask: /24
        TOS: 0  Metric: 10
          
  Routing Bit Set on this LSA
  LS age: 125
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(Network)
  Link State ID: 192.168.34.0 (summary Network Number)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xFFA9
  Length: 28
  Network Mask: /24
        TOS: 0  Metric: 20
          
  Routing Bit Set on this LSA
  LS age: 213
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(Network)
  Link State ID: 192.168.45.0 (summary Network Number)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xFFA9
  Length: 28
  Network Mask: /24
        TOS: 0  Metric: 30
               
  Routing Bit Set on this LSA
  LS age: 125
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(Network)
  Link State ID: 192.168.48.0 (summary Network Number)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xFFA9
  Length: 28
  Network Mask: /24
        TOS: 0  Metric: 20
          
  Routing Bit Set on this LSA
  LS age: 125
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(Network)
  Link State ID: 192.168.56.0 (summary Network Number)
  Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xFFA9
  Length: 28
  Network Mask: /24
        TOS: 0  Metric: 40
```

### ASBR summary LSA (Type 4)

ASBR summary是由ABR产生，==告诉Area内的Router==从他那里可以到达那些ASBR以及Cost

要看ASBR summary LSA可以在router输入`show ip ospf database asbr-summary`以下用R1做例子。R1只收到一条ASBR summary LSA，R2告诉R1从他那里可以到达一只叫6.6.6.6的ASBR，Cost 40

```
R1#show ip ospf 1 database asbr-summary

            OSPF Router with ID (1.1.1.1) (Process ID 1)

                Summary ASB Link States (Area 10)

  Routing Bit Set on this LSA
  LS age: 701
  Options: (No TOS-capability, DC, Upward)
  LS Type: Summary Links(AS Boundary Router)
  Link State ID: 6.6.6.6 (AS Boundary Router address)
    Advertising Router: 2.2.2.2
  LS Seq Number: 80000002
  Checksum: 0xB939
  Length: 28
  Network Mask: /0
        TOS: 0   Metric: 40
  
```

### external LSA(Type 5)

External LSA是由ASBR产生的，告诉==所有Area==(除了Stub Area)里的所有Router，从他那里可以到达那些External Network（==即不属于OSPF的==）

要看external  LSA，可以在Router输入指令`show ip ospf database external`，以下用R1做例子。R1 只收到一条来自远方R6(6.6.6.6)所发布的external LSA，告诉R1从它那里可以到达192.168.67.0这个network，cost 是20，而Metric Type是2，又称E2。

```
R1#show ip ospf 1 database external

            OSPF Router with ID (1.1.1.1) (Process ID 1)

                Type-5 AS External Link States

  LS age: 14
  Options: (No TOS-capability, DC)
  LS Type: AS External Link
  Link State ID: 192.168.67.0 (External Network Number )
  Advertising Router: 6.6.6.6
  LS Seq Number: 80000003
  Checksum: 0x9940
  Length: 36
  Network Mask: /24
  Metric Type: 2 (Larger than any link state path)
        TOS: 0
        Metric: 20
        Forward Address: 0.0.0.0
        External Route Tag: 0
        
```

于是，OSPF的router们就会根据回送的LSA来计算自己的Route table

==用O来表示是由OSPF计算出来的route，如果有IA则表示这个route需要跨域（Area）需要经过ABR。如果有E1或E2则代表这个Route是External Route，需要经过ASBR。110是AD(Administrative Distance)，而AD后面的数字这是该Router的OSPF Metric，Metric和Cost有关==

```
R3#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route

Gateway of last resort is not set

O IA 192.168.12.0/24 [110/20] via 192.168.23.2, 00:10:51, Ethernet0/1
O    192.168.28.0/24 [110/20] via 192.168.23.2, 00:10:51, Ethernet0/1
     3.0.0.0/32 is subnetted, 1 subnets
C       3.3.3.3 is directly connected, Loopback0
O IA 192.168.45.0/24 [110/20] via 192.168.34.4, 00:10:51, Ethernet0/0
O IA 192.168.56.0/24 [110/30] via 192.168.34.4, 00:10:51, Ethernet0/0
C    192.168.23.0/24 is directly connected, Ethernet0/1
O E2 192.168.67.0/24 [110/20] via 192.168.34.4, 00:10:51, Ethernet0/0
C    192.168.34.0/24 is directly connected, Ethernet0/0
O    192.168.48.0/24 [110/20] via 192.168.34.4, 00:10:56, Ethernet0/0
```

#### E1 VS E2

https://ipwithease.com/ospf-external-e1-and-e2-routes/

external route 指的是从非OSPF的路由协议学来的route，标识为E route(Type 5，E对应数字顺序5)

**Cost**

OSPF external type 1(E1)，E1达到network的Cost包含external metric 和 internal cost(到达ASBR的cost)

OSPF external type 2(E2)，E2到达network的Cost只包含external metric，不含internal cost(到达ASBR的cost)

 ![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211203/Snipaste_2021-08-11_20-17-25.1h7j96xxd44g.png)



## Area

OSPF是设计给一个大型网络使用的，为了解决管理上的问题，OSPF采用了一个hierachical system(分层系统)，把大型的OSPF分割成多个Area。

Area有两种表达方式，可以是一个16bits数字（0-65535），或者使用类似IP的方式（OID），例如：192.168.1.1，前者比较常见。Area0（或0.0.0.0）是一个特别的Area，我们称为Backbone Area（骨干），==所有其他Area必须与Backbone Area连接，这是规定==



OSPF 的 Area分为几种Backbone Area(Area 0)、Standard Area、Stub Area、Totally Stubby Area、Not-so-stubby Area、Totally Not-so-stubby Area

### StubArea

在说StubArea，引用SLA中例子，先看一下routing table和ospf database

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
O IA 192.168.28.0/24 [110/20] via 192.168.12.2, 00:26:29, Ethernet0/0
O IA 192.168.45.0/24 [110/40] via 192.168.12.2, 01:19:14, Ethernet0/0
O IA 192.168.56.0/24 [110/50] via 192.168.12.2, 01:19:14, Ethernet0/0
O IA 192.168.23.0/24 [110/20] via 192.168.12.2, 01:19:14, Ethernet0/0
O E2 192.168.67.0/24 [110/20] via 192.168.12.2, 00:26:01, Ethernet0/0
O IA 192.168.34.0/24 [110/30] via 192.168.12.2, 01:19:14, Ethernet0/0
O IA 192.168.48.0/24 [110/30] via 192.168.12.2, 00:25:48, Ethernet0/0

R1#show ip ospf 1 database

            OSPF Router with ID (1.1.1.1) (Process ID 1)

                Router Link States (Area 10)

Link ID         ADV Router      Age         Seq#       Checksum Link count
1.1.1.1         1.1.1.1         11          0x8000000A 0x009A8A 2
2.2.2.2         2.2.2.2         17          0x80000009 0x00DC45 1

                Net Link States (Area 10)

Link ID         ADV Router      Age         Seq#       Checksum
192.168.12.1    1.1.1.1         21          0x80000001 0x00C7EB
192.168.12.2    2.2.2.2         17          0x80000007 0x008325

                Summary Net Link States (Area 10)

Link ID         ADV Router      Age         Seq#       Checksum
192.168.23.0    2.2.2.2         22          0x80000006 0x00F7AD
192.168.28.0    2.2.2.2         22          0x80000004 0x00C4DD
192.168.34.0    2.2.2.2         22          0x80000006 0x00E2AD
192.168.45.0    2.2.2.2         22          0x80000006 0x00CDAD
192.168.48.0    2.2.2.2         22          0x80000005 0x004A39
192.168.56.0    2.2.2.2         22          0x80000006 0x00B8AD

                Summary ASB Link States (Area 10)

Link ID         ADV Router      Age         Seq#       Checksum
6.6.6.6         2.2.2.2         26          0x80000001 0x00BB38

                Type-5 AS External Link States

Link ID         ADV Router      Age         Seq#       Checksum Tag
192.168.67.0    6.6.6.6         1282        0x80000004 0x009741 0
```

收到了router LSA和external LSA，然后产出Route table，每一个Network都通了，还有什么问题呢？

对于R1来说，无论去哪一个Network，Next Hop都是 192.168.12.2 是不是 ？那么，它为什么还有记住这么多的route呢？所以，OSPF的设计者也想到这个问题了，所以发明了Stub Area。我们试试把Area 10设定成Stub Area。指令非常简单，只要在OSPF设定中加一句`area <area number> stub`就可以了

## Connect State

Router们称为neighbor其实中间经过很多状态（称为Adjacency），包含如下几种

1. Down：没有发放Hello Message
2. Init：刚刚向对方发放Hello Message
3. 2-way：他们开始沟通，选出DR和BDR。
4. ExStart：预备交换Link信息
5. Exchange：正在交换DBD（Database Descriptors），可以把DBD看成Link信息的一些目录，先给对方目录，让对方回复那些LSA（link state advertisements）是他需要的
6. Loading：正在交换LSA
7. Full：两只router称为adjacency，在这是，同一个area的router里面的topology table应该完全相同

## Neighbor

### requests to be neighbor

1. Area ID 必需一致

2. Area Type 必需一致

3. prefix和subnet mask必需一致，即在同一网段

4. hello interval和dead interval必需一致

5. authentication必需一致，即ospf密码相同

   开启密码

   ```
   R3(config-router)#area 0 authentication
   ```

   明文密码

   ```
   R3(config)#router ospf 1
   #允许authentication，使用no来取消
   R3(config-router)#area 0 authentication 
   R3(config-router)#int fa0/0
   #设置密码
   R3(config-if)#ip ospf authentication-key 123
   ```

   md5密码

   ```
   R3(config)#router ospf 1
   R3(config-router)#ip ospf authentication 
   #允许使用md5 authentication,使用no来取消
   R3(config-rouoter)#ip ospf authentication message-digest 
   R3(config-router)#int fa0/0
   #设置密码
   R3(config-if)#ip ospf message-digest-key 10 md5 123
   
   R3#sh run int fa0/0
   Building configuration...
   
   Current configuration : 168 bytes
   !
   interface FastEthernet0/0
    ip address 192.168.79.3 255.255.255.0
    ip ospf authentication-key 123
    ip ospf message-digest-key 10 md5 123
    duplex auto
    speed auto
   end
   ```

## Designated Router

在多播中，==同一网段（非area）==里可能连接3个以上的router。每个router需要建立n-1条连接（n表示router个数），fully-mesh的connection为n(n-1)/2

router越多connnection就越多，ospf会在这些router中选一个DR（designated router），所有的router只需要与DR建立连接。每次有routing update，router只需要把update传给dr，再有dr统一发放给其他的router，这样除了DR外，Area里每只router只需要处理一条connection。但DR是单点，如果出现问题就下线了，所以需要BDR(baking)

DR和BDR之间有竞选算法，priority较高会成为该网段的DR，第二高的会成为BDR。如果priority相同，router id较高者会成为DR

![2021-11-23_23-21](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211123/2021-11-23_23-21.6t9hahc8n040.png)

可以看下面一个实验

```
hostname R1
!
interface Ethernet0/0
 ip address 192.168.1.1 255.255.255.0
!
router ospf 1
 network 192.168.1.0 0.0.0.255 area 0
```

```
hostname R2
!
interface Ethernet0/0
 ip address 192.168.1.2 255.255.255.0
!
router ospf 1
 network 192.168.1.0 0.0.0.255 area 0
```

```
hostname R3
!
interface Ethernet0/0
 ip address 192.168.1.3 255.255.255.0
!
router ospf 1
 network 192.168.1.0 0.0.0.255 area 0
```

```
hostname R4
!
interface Ethernet0/0
 ip address 192.168.1.4 255.255.255.0
!
router ospf 1
 network 192.168.1.0 0.0.0.255 area 0
```

在R1看neighbor table，priority都是1(默认值)，由于priority都相同，所以会选router id 192.168.1.4最大的成为DR，192.168.1.3成为BDR

```
R1#show ip ospf neighbor 

Neighbor ID Pri State Dead Time Address Interface 
192.168.1.2 1 FULL/DROTHER 00:00:38 192.168.1.2 Ethernet0/0 
192.168.1.3 1 FULL/BDR 00:00:38 192.168.1.3 Ethernet0/0 
192.168.1.4 1 FULL/DR 00:00:37 192.168.1.4 Ethernet0/0 
R1#show ip ospf interface ethernet 0/0 
Ethernet0/0 is up, line protocol is up 
Internet Address 192.168.1.1/24, Area 0 Process ID 1, Router ID 192.168.1.1, Network Type BROADCAST, Cost: 10 
Transmit Delay is 1 sec, State DROTHER, Priority 1 
Designated Router (ID) 192.168.1.4, Interface address 192.168.1.4 
Backup Designated router (ID) 192.168.1.3, Interface address 192.168.1.3 
Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5 
oob-resync timeout 40 
Hello due in 00:00:01 
Supports Link-local Signaling (LLS) 
Index 1/1, flood queue length 0 
Next Ox0(0)/0x0(0) Last flood scan length is 0, maximum is 1 Last flood scan time is 0 msec, maximum is 0 msec 
Neighbor Count is 3, Adjacent neighbor count is 2 
Adjacent with neighbor 192.168.1.3 (Backup Designated Router) 
Adjacent with neighbor 192.168.1.4 (Designated Router) 
Suppress hello for 0 neighbor(s) 
```

要该改变priority，可以在interface中输入`ip ospf priority <priority>`

```
R1(config)#interface ethernet 0/0
R1(config-if)#ip ospf priority 100
```

虽然priority改好了，但是R1任然是DROTHER，没有成为DR，因为已经成为DR的router是不会被中途抢走的，所有要想R1成为DR，必须在R4 reset OSPF process

```
R1#show ip ospf interface ethernet 0/0
Ethernet0/0 is up, line protocol is up
  Internet Address 192.168.1.1/24, Area 0
  Process ID 1, Router ID 192.168.1.1, Network Type BROADCAST, Cost: 10
  Transmit Delay is 1 sec, State DROTHER, Priority 100
```

用`clear ip ospf process`来reset ospf process

```
R4#clear ip ospf process
Reset ALL OSPF processes? [no]: yes
R4#
*Mar  1 00:28:41.175: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.1 on Ethernet0/0 from 2WAY to DOWN, Neighbor Down: Interface down or detached
*Mar  1 00:28:41.179: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.2 on Ethernet0/0 from 2WAY to DOWN, Neighbor Down: Interface down or detached
*Mar  1 00:28:41.179: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.3 on Ethernet0/0 from FULL to DOWN, Neighbor Down: Interface down or detached
```

在看看R1，怎么还不是DR，而是BDR

```
R1#show ip ospf interface ethernet 0/0
Ethernet0/0 is up, line protocol is up
  Internet Address 192.168.1.1/24, Area 0
  Process ID 1, Router ID 192.168.1.1, Network Type BROADCAST, Cost: 10
  Transmit Delay is 1 sec, State BDR, Priority 100
```

因为当R4下线后，原来的BDR R3会立刻升格成DR，R1就只能由DROTHER进升成为BDR，只有再Reset R3，R1才能成为真正的DR。这是一个很好的机制，减少DR出现不稳定的情况。Reset R3之后，由于R4的router ID最大，所以成为新的BDR

```
Reset ALL OSPF processes? [no]: yes
R3#
*Mar  1 00:32:04.415: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.1 on Ethernet0/0 from FULL to DOWN, Neighbor Down: Interface down or detached
*Mar  1 00:32:04.419: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.2 on Ethernet0/0 from 2WAY to DOWN, Neighbor Down: Interface down or detached
*Mar  1 00:32:04.419: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.4 on Ethernet0/0 from EXSTART to DOWN, Neighbor Down: Interface down or detached
R3#
*Mar  1 00:32:04.515: %OSPF-4-NONEIGHBOR: Received database description from unknown neighbor 192.168.1.4
R3#
*Mar  1 00:32:12.799: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.1 on Ethernet0/0 from LOADING to FULL, Loading Done
*Mar  1 00:32:12.995: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.1.4 on Ethernet0/0 from LOADING to FULL, Loading Done
```

R1成为DR

```
R1#show ip ospf interface ethernet 0/0
Ethernet0/0 is up, line protocol is up
  Internet Address 192.168.1.1/24, Area 0
  Process ID 1, Router ID 192.168.1.1, Network Type BROADCAST, Cost: 10
  Transmit Delay is 1 sec, State DR，Priority 100
```

## Timer

ospf预设hello interval是10 sec，即每 10 sec向neighbor发送hello msg，对方收到后回应。但是如果经过40 sec没有收到 hello msg，就表示对方下线了。

Dead Time 从40 sec开始计算，通常到30 sec ，因为收到回应就会重新设为 40 sec，但是如果倒数到0 sec也收不到回应，就表示neighbor 下线了

==hello interval 和 dead interval 是可以在interfaace更改的，但是能成为neighbor的前提是两只router的 interval timer必需一致==

```
ip ospf hello-interval 15
ip ospf dead-interval 60
```

OSPF通过hello timer和dead timer来判断neighbor是否存活。但是为什么要控制timer？因为不同网络中的网络带宽和时延不同，neighbor要靠hello timer和dead timer去维持，在效能低的网络中，就要改成hello interval 30，dead interval 120

## Metric/Cost

> 只计算出方向，不计算入方向

Cost是OSPF计算Metric的依据，==由Source到Destination整条路径的cost总和就是Metric==。

bandwidth和Cost成反比，bandwidth越大Cost越小。Cost由如下公式计算：

cost = 10000 0000（100MB）/ bandwidth pbs

其中的100MB也被从未Reference Bandwidth

例如：现在需要计算从R1到3.3.3.3的cost

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211202/Snipaste_2021-08-11_20-17-25.2tlnpet7vs40.png)

可以使用`show interface eth0/0 | include BW`来查看端口的bandwidth

```
R1#show interfaces ether0/0 | include BW
  MTU 1500 bytes, BW 10000 Kbit, DLY 1000 usec,
```

这样就可以算出R1 eth0/0的Cost了，100MB/10MB=10。然后在经过R2的serial1/0，cost=100M / 1.544M=64

```
R2#show interfaces serial 1/0 | include BW
  MTU 1500 bytes, BW 1544 Kbit, DLY 20000 usec,
```

==到达R3后，还需要经过R3的Loopback 0 这个Interface（从lo到kernel），cost= 100 M/ / 8000M = 1 （cost最小为1）==

```
R3#show interfaces loopback 0 | include BW
  MTU 1514 bytes, BW 8000000 Kbit, DLY 5000 usec,
```

因此，从R1到达R3的3.3.3.3，cost=10+64+1=75，可以在R1的route table上查看

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/11] via 192.168.12.2, 00:18:58, Ethernet0/0
     3.0.0.0/32 is subnetted, 1 subnets
O       3.3.3.3 [110/75] via 192.168.12.2, 00:18:58, Ethernet0/0
O    192.168.23.0/24 [110/74] via 192.168.12.2, 00:18:58, Ethernet0/0
```

110为AD，75为Metric。相反的，从R3到R1的1.1.1.1应该也是75（前提是route相同）

```
R3#show ip route | begin Gateway
Gateway of last resort is not set

O    192.168.12.0/24 [110/74] via 192.168.23.2, 00:02:33, Serial0/0
     1.0.0.0/32 is subnetted, 1 subnets
O       1.1.1.1 [110/75] via 192.168.23.2, 00:02:33, Serial0/0
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/65] via 192.168.23.2, 00:02:33, Serial0/0
     3.0.0.0/32 is subnetted, 1 subnets
C       3.3.3.3 is directly connected, Loopback0
C    192.168.23.0/24 is directly connected, Serial0/0
```

### Modify Cost

由改公式得出的结果也是默认的ospf cost，但是可以使用`ip ospf cost <value>`强制一个interface的cost，也可以修改interface bandwidth 或者改 reference bandwidth 等

#### ip ospf cost

我们可以直接在interface指定cost值，例如想将R1的eth0/0的cost设为50。完成设定后R1到3.3.3.3的metric变成115了（50+64+1=115）

```
R1(config)#int ethernet 0/0
R1(config-if)#ip ospf cost 50
R1(config-if)#end
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/51] via 192.168.12.2, 00:00:08, Ethernet0/0
     3.0.0.0/32 is subnetted, 1 subnets
O       3.3.3.3 [110/115] via 192.168.12.2, 00:00:08, Ethernet0/0
O    192.168.23.0/24 [110/114] via 192.168.12.2, 00:00:08, Ethernet0/0
```

但需要留言的是，北更改cost的是R1的eth0/0，从相反方向的routing metric不会受到影响（因为值计算出方向）

```
R3#show ip route | begin Gateway
Gateway of last resort is not set

O    192.168.12.0/24 [110/74] via 192.168.23.2, 00:02:57, Serial0/0
     1.0.0.0/32 is subnetted, 1 subnets
O       1.1.1.1 [110/75] via 192.168.23.2, 00:02:58, Serial0/0
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/65] via 192.168.23.2, 00:02:58, Serial0/0
     3.0.0.0/32 is subnetted, 1 subnets
C       3.3.3.3 is directly connected, Loopback0
C    192.168.23.0/24 is directly connected, Serial0/0
```

#### interface bandwidth

根据计算cost的公式，我们还可以更改Interface bandwidth来设定cost。例如：需要把R1 eth0/0 cost设为20的话，可以把interface bandwidth 调到 5M。如下图，修改bandwidth后，metric变成85。（20 + 64 + 1 = 85）。==需要注意的是，修改Bandwidth只会对Routing Protocol计算Cost时造成影响，并不会真正改变interface的传输速度。==

```
R1(config)#int ethernet 0/0
R1(config-if)#bandwidth ?
  <1-10000000>  Bandwidth in kilobits
  inherit       Specify that bandwidth is inherited
  receive       Specify receive-side bandwidth

R1(config-if)#bandwidth 5000
R1(config-if)#end
R1#
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/21] via 192.168.12.2, 00:00:13, Ethernet0/0
     3.0.0.0/32 is subnetted, 1 subnets
O       3.3.3.3 [110/85] via 192.168.12.2, 00:00:13, Ethernet0/0
O    192.168.23.0/24 [110/84] via 192.168.12.2, 00:00:13, Ethernet0/0
```

#### reference bandwidth

按照公式，同样的也可以修改reference bandwidth。现在可以试着吧reference bandwidth 改成1000 bit，Metric就会变成165（100+64）。==尽量避免这种方式，会对troubleshooting造成障碍，如果需要修改就应该将OSPF router的reference bandwidth都改成一样==

```
R1(config)#router ospf 1
R1(config-router)#auto-cost reference-bandwidth 1000
R1(config-router)#end
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
     2.0.0.0/32 is subnetted, 1 subnets
O       2.2.2.2 [110/101] via 192.168.12.2, 00:01:46, Ethernet0/0
     3.0.0.0/32 is subnetted, 1 subnets
O       3.3.3.3 [110/165] via 192.168.12.2, 00:01:46, Ethernet0/0
O    192.168.23.0/24 [110/164] via 192.168.12.2, 00:01:46, Ethernet0/0
```

另外补充一点==如果有超过100M的interface，必须加大reference bandwidth==，因为在预设100M的情况下，10G、1G 和 100 M interface计算出来的cost都会是1（因为最小为1），也就无法准确判断link的快慢

#### neighbor cost

如果Network Type 是Point-to-Multipoint Non Broadcast的话，可以直接更改Neighbor Cost，这样可以设定对方传来的Route Cost，这种方法比较少用

```
R7(config)#int serial 0/0
R7(config-if)#ip ospf network point-to-multipoint non-broadcast 
R7(config-if)#end
R7(config)#router ospf 1
R7(config-router)#neighbor 192.168.23.2 cost 999
```

#### external route cost

```

```



## Network Type

根据 自动/手动Neighbor、是否DR选举、10,40/30,120 Timer设定这三大元素，Cisco router把不同的设定组合归纳为5中模式，我们把这玩意称为Network Type

![2021-11-25_22-20](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211123/2021-11-25_22-20.v60hfqvsywg.png)

### auto detect neighbor/ manually input neighbor

OSPF是你用Multicast自动发现Neighbor的，换句话说，如果网络不支持Broadcast，等于不支持Multicast，就不可以自动发现Neighbor。毫无疑问Ethernet是支持Broadcast

另外就是Point to Point Network，由于在Point to Point Network里面，虽然没有Broadcast，但建立Point to Point connection的两个Node绝对知道对方存在，因此也会自动建立Neighbor

因此结论是：

==凡是Non-Broadcast的Network（除Point to Point connnection）不会自动发现Neighbor，需要手动输入Neighbor==

可以使用一个Multipoint Non-Broadcast 的 Frame Relay Network来做示范

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211202/Snipaste_2021-08-11_20-17-25.feuzlb0cngo.png)

```
hostname R1
!
interface Serial0/0
 ip address 192.168.123.1 255.255.255.0
 encapsulation frame-relay
 frame-relay map ip 192.168.123.2 102
 frame-relay map ip 192.168.123.3 103
```

```
hostname R2
!
interface Serial0/0
 ip address 192.168.123.2 255.255.255.0
 encapsulation frame-relay
 frame-relay map ip 192.168.123.1 201
 frame-relay map ip 192.168.123.3 201
```

```
hostname R3
!
interface Serial0/0
 ip address 192.168.123.3 255.255.255.0
 encapsulation frame-relay
 frame-relay map ip 192.168.123.1 301
 frame-relay map ip 192.168.123.2 301
```

由于frame-relay指令没有Broadcast这个keyword，所以此网络不支持Broadcast。当尝试配置Router OSPF

```
hostname R1
!
router ospf 1
 network 192.168.123.0 0.0.0.255 area 0
```

```
hostname R2
!
router ospf 1
 network 192.168.123.0 0.0.0.255 area 0
```

```
hostname R3
!
router ospf 1
 network 192.168.123.0 0.0.0.255 area 0
```

虽然指令输入了，但是就算等到2046，Neighbor也不会建立连接，因为不支持Broadcast，这时候需要手动建立Neighbor

```
R1(config)#router ospf 1
R1(config-router)#neighbor 192.168.123.2
R1(config-router)#neighbor 192.168.123.3
R1(config-router)#
*Mar  1 00:28:05.083: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.123.2 on Serial0/0 from LOADING to FULL, Loading Done
*Mar  1 00:28:06.083: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.123.3 on Serial0/0 from LOADING to FULL, Loading Done
```

### choose DR/BDR 

选择DR，是因为不想建立太多的connection。如果只有两只Router，就不用选，因为两只Router就只有一条connection，根本不用担心会建立太多的connection，所以不用选。这种只有两只Router的网络我们叫做Point To Point Network。

==即只有像Ethernat这种Multi Access的Network才需要选DR==。如果需要修改Non-Broadcast中的Priority值来影响DR的选择，可以使用如下方法

```
R1(config)#router ospf 1
R1(config-router)#neighbor 192.168.123.2 priority 50
R1(config-router)#neighbor 192.168.123.3 priority 100
```

### control Timer

Broadcast Multi-access 和 Point-to-Point Network，预设 Timer为10/40，其他低速网络这使用30/120

### change Network Type

如果是Ethernet interface，预设都是使用Broadcast Multi-access，使用`show ip ospf interface`就可以看到

```
R1#show ip ospf interface
Ethernet0/0 is up, line protocol is up
  Internet Address 192.168.12.1/24, Area 0
  Process ID 1, Router ID 192.168.12.1, Network Type BROADCAST, Cost: 10
  Transmit Delay is 1 sec, State DR, Priority 1
  Designated Router (ID) 192.168.12.1, Interface address 192.168.12.1
  No backup designated router on this network
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    oob-resync timeout 40
    Hello due in 00:00:05
  Supports Link-local Signaling (LLS)
  Index 1/1, flood queue length 0
  Next 0x0(0)/0x0(0)
  Last flood scan length is 0, maximum is 0
  Last flood scan time is 0 msec, maximum is 0 msec
  Neighbor Count is 0, Adjacent neighbor count is 0
  Suppress hello for 0 neighbor(s)
```

在Interface 输入`ip ospf network <network type>`可以更改network type

```
R1(config)#interface ethernet 0/0
R1(config-if)#ip ospf network ?
  broadcast            Specify OSPF broadcast multi-access network
  non-broadcast        Specify OSPF NBMA network
  point-to-multipoint  Specify OSPF point-to-multipoint network
  point-to-point       Specify OSPF point-to-point network

R1(config-if)#ip ospf network non-broadcast
R1(config-if)#end
R1#show ip ospf interface
Ethernet0/0 is up, line protocol is up
  Internet Address 192.168.12.1/24, Area 0
  Process ID 1, Router ID 192.168.12.1, Network Type NON_BROADCAST, Cost: 10
  Transmit Delay is 1 sec, State WAITING, Priority 1
  No designated router on this network
  No backup designated router on this network
  Timer intervals configured, Hello 30, Dead 120, Wait 120, Retransmit 5
    oob-resync timeout 120
    Hello due in 00:00:17
    Wait time before Designated router selection 00:01:17
  Supports Link-local Signaling (LLS)
  Index 1/1, flood queue length 0
  Next 0x0(0)/0x0(0)
  Last flood scan length is 0, maximum is 0
  Last flood scan time is 0 msec, maximum is 0 msec
  Neighbor Count is 0, Adjacent neighbor count is 0
  Suppress hello for 0 neighbor(s)
```

## Route Decision

和其他的Routing Protocol一样，OSPF并不会把全部路径放进Route Table，只会选择当中一条或多条路径。当有多条路径能抵达目的地，OSPF会按一下顺序来选择：

1. 先按Route Type来选择：Intra-area(O) > Inter-area(O IA) > Type 1 External(O E1/N1) > Type 2 External(O E2/N2)
2. 如果Route Type相同，比较Metric，较小者获胜
3. 如果Metric 相同，选择n条path一并加进Route Table，进行load balancing(n = maximum - paths参数)

### Route Type

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211203/Snipaste_2021-08-11_20-17-25.66adryfxyz00.png)

为R1同时提供四条能够到达6.6.6.6而Metric相同的Route

1. 经R5的Intra-area Route(O)
2. 经R4的Inter-area Route(O IA)
3. 经R3的External Route(O E1)
4. 经R2的External Route(O E2)

各Router设定如下

```
hostname R1
!
interface Loopback0
 ip address 1.1.1.1 255.255.255.255
!
interface Ethernet0/0
 ip address 192.168.12.1 255.255.255.0
!
interface Ethernet0/1
 ip address 192.168.13.1 255.255.255.0
!
interface Ethernet0/2
 ip address 192.168.14.1 255.255.255.0
!
interface Ethernet0/3
 ip address 192.168.15.1 255.255.255.0
!
router ospf 1
 network 192.168.12.0 0.0.0.255 area 0
 network 192.168.13.0 0.0.0.255 area 0
 network 192.168.14.0 0.0.0.255 area 0
 network 192.168.15.0 0.0.0.255 area 10
```

```
hostname R2
!
interface Loopback0
 ip address 2.2.2.2 255.255.255.255
!
interface Ethernet0/0
 ip address 192.168.12.2 255.255.255.0
!
interface Ethernet0/1
 ip address 192.168.26.2 255.255.255.0
!
router eigrp 100
 network 192.168.26.0
 no auto-summary
!
router ospf 1
 redistribute eigrp 100 metric 21 subnets
 network 192.168.12.0 0.0.0.255 area 0
```

```
hostname R3
!
interface Loopback0
 ip address 3.3.3.3 255.255.255.255
!
interface Ethernet0/0
 ip address 192.168.13.3 255.255.255.0
!
interface Ethernet0/1
 ip address 192.168.36.3 255.255.255.0
!
router eigrp 200
 network 192.168.36.0
 no auto-summary
!
router ospf 1
 redistribute eigrp 200 metric 11 metric-type 1 subnets
 network 192.168.13.0 0.0.0.255 area 0
```

```
hostname R4
!
interface Loopback0
 ip address 4.4.4.4 255.255.255.255
!
interface Ethernet0/0
 ip address 192.168.14.4 255.255.255.0
!
interface Ethernet0/1
 ip address 192.168.46.4 255.255.255.0
!
router ospf 1
 network 192.168.14.0 0.0.0.255 area 0
 network 192.168.46.0 0.0.0.255 area 10
```

```
hostname R5
!
interface Loopback0
 ip address 5.5.5.5 255.255.255.255
!
interface Ethernet0/0
 ip address 192.168.15.5 255.255.255.0
!
interface Ethernet0/1
 ip address 192.168.56.5 255.255.255.0
!
router ospf 1
 network 192.168.15.0 0.0.0.255 area 10
 network 192.168.56.0 0.0.0.255 area 10
```

```
hostname R6
!
interface Loopback0
 ip address 6.6.6.6 255.255.255.255
!
interface Ethernet0/0
 ip address 192.168.26.6 255.255.255.0
!
interface Ethernet0/1
 ip address 192.168.36.6 255.255.255.0
!
interface Ethernet0/2
 ip address 192.168.46.6 255.255.255.0
 half-duplex
!
interface Ethernet0/3
 ip address 192.168.56.6 255.255.255.0
!
router eigrp 100
 network 6.6.6.6 0.0.0.0
 network 192.168.26.0
 no auto-summary
!
router eigrp 200
 network 6.6.6.6 0.0.0.0
 network 192.168.36.0
 no auto-summary
!
router ospf 1
 network 6.6.6.6 0.0.0.0 area 10
 network 192.168.46.0 0.0.0.255 area 10
 network 192.168.56.0 0.0.0.255 area 10
```

实验开始，先把R1的eth0/1、eth0/2、eth0/3关掉，这样R1就只有next hop R2(192.168.12.2)这条路径到达6.6.6.6

```
R1(config)#interface range ethernet 0/1 - 3
R1(config-if-range)#shutdown
R1(config-if-range)#end
R1#
R1#show ip route | include 6.6.6.6
O E2    6.6.6.6 [110/21] via 192.168.12.2, 00:14:35, Ethernet0/0
```

现在把eth0/1打开，由于next hop R3(192.168.13.3)这条Route是E1，OSPF认为E1比E2好，所以就选E1放入Route Table，代替E2

```
, 00:14:35, Ethernet0/0E1    6.6.6.6 [110/21] via 192.168.13.3, 00:00:46, Ethernet0/1
```

再打开eth0/2，next hop R4(192.168.14.4)是O IA Route，比E1好，所以OSPF选择了这条Route

```
R1(config)#interface ethernet 0/2
R1(config-if)#no shutdown
R1(config-if)#end
R1#
R1#show ip route | include 6.6.6.6
O IA    6.6.6.6 [110/21] via 192.168.14.4, 00:00:05, Ethernet0/2
```

最后，把eth0/3打开，当然R1会选择next hop为R5(192.168.15.5)，因为Intra-area route最好的

```
R1(config)#interface ethernet 0/3
R1(config-if)#no shutdown
R1(config-if)#end
R1#
R1#show ip route | include 6.6.6.6
O       6.6.6.6 [110/21] via 192.168.15.5, 00:00:00, Ethernet0/3
```

### Metric

在同一个Route Type下，OSPF会选择Metric较少的放进Route Table。只需要把network内的所有Intrfaces都放进一个Area中，这样，全部路径都是Intra-area Route，然后在R1把Interface设定成不同Cost

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211203/Snipaste_2021-08-11_20-17-25.4d9liw4my8g0.png)

```
R1(config)#interface ethernet 0/0
R1(config-if)#ip ospf cost 30
R1(config-if)#interface ethernet 0/1
R1(config-if)#ip ospf cost 20
R1(config-if)#interface ethernet 0/2
R1(config-if)#ip ospf cost 10
R1
R1#show ip route | include 5.5.5.5
O       5.5.5.5 [110/21] via 192.168.14.4, 00:00:03, Ethernet0/2
```

R2 Metric = 30 + 10 + 1 = 41

R3 Metric = 20 + 10 + 1 = 31

R4 Metric = 10 + 10 + 1 = 21

所以会选择Metric最少的R4路径

### Load Balancing

如果Route Type和Metric 都不会分出优劣，在Metric相同的情况下，OSPF预设最多会选4条Route放进Route Table，进行Load Balancing。再用上一个例子的网络为例，如果R1的三个Interface的Cost全部相同，三条Route都会放进Route Table

```
R1#show ip ospf interface ethernet 0/0 | include Cost
  Process ID 1, Router ID 192.168.114.1, Network Type BROADCAST, Cost: 10
R1#show ip ospf interface ethernet 0/1 | include Cost
  Process ID 1, Router ID 192.168.114.1, Network Type BROADCAST, Cost: 10
R1#show ip ospf interface ethernet 0/2 | include Cost
  Process ID 1, Router ID 192.168.114.1, Network Type BROADCAST, Cost: 10
R1
R1#show ip route | begin 5.5.5.5
O       5.5.5.5 [110/21] via 192.168.14.4, 00:00:54, Ethernet0/2
                [110/21] via 192.168.14.3, 00:00:54, Ethernet0/1
                [110/21] via 192.168.14.2, 00:00:54, Ethernet0/0
O    192.168.35.0/24 [110/20] via 192.168.13.3, 00:00:54, Ethernet0/1
```

OSPF预设4条Load Balance是可以修改的，只要在OSPF设定下用指令`maximum-paths <number of path>`便可以设定1至16条Load Balance Route。例如，把R1改成maximum-paths 2 ，Route就会变成2条Load Balance Route

```
R1(config)#router ospf 1
R1(config-router)#maximum-paths ?
  <1-16>  Number of paths

R1(config-router)#maximum-paths 2
R1(config-router)#end
R1#
R1#show ip route | begin 5.5.5.5
O       5.5.5.5 [110/21] via 192.168.13.3, 00:00:13, Ethernet0/1
                [110/21] via 192.168.12.2, 00:00:13, Ethernet0/0
O    192.168.35.0/24 [110/20] via 192.168.13.3, 00:00:13, Ethernet0/1
```

## Virtual Link

Virtual Link的概念非常简单，就是把两个实体上分离的Area相连起来。一般有两个用途，他可以帮没有与Backbon Area0相邻的Area建立一条Virtual Link连接起来。另外，可以将断开的Area0接驳起来

### 接驳至Area0

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211203/Snipaste_2021-08-11_20-17-25.36xv87jzgi00.png)

以上图为例，==Area20没有连接至Area0，就算把所有OSPF配置好，R1与R2都不会收到Area20的Route==

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
O IA 192.168.23.0/24 [110/20] via 192.168.12.2, 00:15:19, Ethernet0/0
```

```
R2#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/1
     2.0.0.0/32 is subnetted, 1 subnets
C       2.2.2.2 is directly connected, Loopback0
C    192.168.23.0/24 is directly connected, Ethernet0/0
```

然后，我们在R2和R3用`area <transit area> virtual-link <router-id>`来设定virtual link

Transit area是virtual link所经过的Area，而router-id这是对方的Router ID

```
R2(config)#router ospf 1
R2(config-router)#area 10 virtual-link 3.3.3.3
```

```
R3(config)#router ospf 1
R3(config-router)#area 10 virtual-link 2.2.2.2
```

于是R1和R2都收到了Area 20的Route

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
O IA 192.168.23.0/24 [110/20] via 192.168.12.2, 00:01:06, Ethernet0/0
O IA 192.168.34.0/24 [110/30] via 192.168.12.2, 00:01:06, Ethernet0/0
```

```
R2#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/1
     2.0.0.0/32 is subnetted, 1 subnets
C       2.2.2.2 is directly connected, Loopback0
C    192.168.23.0/24 is directly connected, Ethernet0/0
O IA 192.168.34.0/24 [110/20] via 192.168.23.3, 00:01:40, Ethernet0/0
```

==注意route 都是inter area route(即需要跨Area)==

### 接驳两个Area0

在整个OSPF network中只可以存在一个Backbon Area0，如果有两个Area0的话，必须用Virtual Link将其相连

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211203/Snipaste_2021-08-11_20-17-25.6k8nb8xvf8g0.png)

按照上图如果没有设定Virtual LIink的情况下，两个Area0没法相连

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
O IA 192.168.23.0/24 [110/20] via 192.168.12.2, 00:03:36, Ethernet0/0
```

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
O IA 192.168.23.0/24 [110/20] via 192.168.12.2, 00:03:36, Ethernet0/0
```

所以要在R2和R3 之间建立Virtual Link

```
R2(config)#router ospf 1
R2(config-router)#area 10 virtual-link 3.3.3.3
```

```
R3(config)#router ospf 1
R3(config-router)#area 10 virtual-link 2.2.2.2
```

两个Area0 互相看到对方了，留意一点，虽然经过了Area10，==但这条Area0去Area0的Route依然被视为intra-area Route==

```
R1#show ip route | begin Gateway
Gateway of last resort is not set

C    192.168.12.0/24 is directly connected, Ethernet0/0
     1.0.0.0/32 is subnetted, 1 subnets
C       1.1.1.1 is directly connected, Loopback0
O IA 192.168.23.0/24 [110/20] via 192.168.12.2, 00:00:54, Ethernet0/0
O    192.168.34.0/24 [110/30] via 192.168.12.2, 00:00:54, Ethernet0/0
```

```
R4#show ip route | begin Gateway
Gateway of last resort is not set
O    192.168.12.0/24 [110/30] via 192.168.34.3, 00:01:09, Ethernet0/1
     4.0.0.0/32 is subnetted, 1 subnets
C       4.4.4.4 is directly connected, Loopback0
O IA 192.168.23.0/24 [110/20] via 192.168.34.3, 00:01:09, Ethernet0/1
C    192.168.34.0/24 is directly connected, Ethernet0/1
```

## Summarization

==Summarization可以把相连的Network组合起来，有效减少Route的数量==。OSPF的Summarization有两种，一个用在ABR的Area Range，而另一个用在ASBR的Summary Address。用一下网络做例子，==ABR R2会把172.16.0.1/32 及 172.16.1.1/32两个Network Summary，而ASBR R3则会把172.16.2.1/32 及 172.16.3.1/32 Redistribution 到 OSPF 之中，并把他们Summary==

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211203/Snipaste_2021-08-11_20-17-25.3npncqo1opi0.png)

### Area Range

在Summary之前，先在R3看看Route Table，正分别收到172.16.0.1/32 及 172.16.1.1/32 两条Route

```
R3#show ip route ospf
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       + - replicated route, % - next hop override

Gateway of last resort is not set

      172.16.0.0/32 is subnetted, 4 subnets
O IA     172.16.0.1 [110/21] via 192.168.23.2, 00:09:51, Ethernet1/0
O IA     172.16.1.1 [110/21] via 192.168.23.2, 00:09:51, Ethernet1/0
O IA  192.168.12.0/24 [110/20] via 192.168.23.2, 00:09:51, Ethernet1/0
```

在R2下Area Range指令：

```
R2(config)#router ospf 1
R2(config-router)#area 0 range 172.16.0.0 255.255.254.0。
```

这样R3只会收到Summary Route 172.16.0.0/23

```
R3#show ip route ospf
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       + - replicated route, % - next hop override

Gateway of last resort is not set

      172.16.0.0/16 is variably subnetted, 3 subnets, 2 masks
O IA     172.16.0.0/23 [110/21] via 192.168.23.2, 00:02:02, Ethernet1/0
O IA  192.168.12.0/24 [110/20] via 192.168.23.2, 00:17:06, Ethernet1/0
```

### Summary Address

R3把172.16.2.1/32 及 172.16.3.1/32 Redistribute 到 OSPF使其成为ASBR，在R1看Route Table会看到这两条External Route

```
R1#show ip route ospf
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       + - replicated route, % - next hop override

Gateway of last resort is not set

      172.16.0.0/32 is subnetted, 4 subnets
O E2     172.16.2.1 [110/20] via 192.168.12.2, 00:00:09, Ethernet1/0
O E2     172.16.3.1 [110/20] via 192.168.12.2, 00:00:03, Ethernet1/0
O IA  192.168.23.0/24 [110/20] via 192.168.12.2, 00:21:24, Ethernet1/0
```

现在于R3下 Summary Address指令把两条External Route 组起来

```
R3(config)#router ospf 1
R3(config-router)#summary-address 172.16.2.0 255.255.254.0
```

于是R1收到了一条已被Summary的External Route

```
R1#show ip route ospf
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       + - replicated route, % - next hop override

Gateway of last resort is not set

      172.16.0.0/16 is variably subnetted, 3 subnets, 2 masks
O E2     172.16.2.0/23 [110/20] via 192.168.12.2, 00:00:31, Ethernet1/0
O IA  192.168.23.0/24 [110/20] via 192.168.12.2, 00:23:40, Ethernet1/0
```



## Example

![2021-11-09_01-10](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211108/2021-11-09_01-10.1u7c56nzfq74.png)

1. Gonfigure Mode使用`router ospf <process id>`启动OSPF，留意process ID只是本机执行OSPF process的一个ID，和设定EIGRP时的AS number不同，两只要成为Neighbor的router==不需要==拥有相同ID
2. 使用`network <network no> <wildcard> area <area id>`来宣告哪一个interface会参与OSPF，参与OSPF的Interface会发布hello packet尝试与对方成为neighbor，然后再成为Adjacency

```
hostname R1

conf t
router ospf 1
network 192.168.80.0 0.0.0.255 area0
network 192.168.79.0 0.0.0.255 area1

R1#sh running-config interface fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.80.1 255.255.255.0
 duplex auto
 speed auto
end

R1#sh running-config interface fa0/1
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/1
 ip address 192.168.79.1 255.255.255.0
 duplex auto
 speed auto
end
```

```
hostname R2

conf t
router ospf 1
network 192.168.80.0 0.0.0.255 area0
network 192.168.78.0 0.0.0.255 area2

R2#show running-config interface fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.80.2 255.255.255.0
 duplex auto
 speed auto
end

R2#show running-config interface fa0/1
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/1
 ip address 192.168.78.2 255.255.255.0
 duplex auto
 speed auto
end
```

```
hostname R3

conf t
router ospf 1
network 192.168.78.0 0.0.0.255 area2
network 192.168.79.0 0.0.0.255 area1

R3#sh running-config int fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.79.3 255.255.255.0
 duplex auto
 speed auto
end

R3#sh running-config int fa0/1
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/1
 ip address 192.168.78.3 255.255.255.0
 duplex auto
 speed auto
end
```

如果配置成功后会显示，LOADING to FULL

```
*Mar  1 01:11:44.203: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.80.2 on FastEthernet0/1 from LOADING to FULL, Loading Done
```

## Command

### network IP wildcard  area number

IP 为具体Interface IP

IP 为Interface network-id

### show running-config | se r o

查看ospf配置

```
R1#show running-config | se r o
router ospf 1
 log-adjacency-changes
 network 192.168.79.0 0.0.0.255 area 1
 network 192.168.80.0 0.0.0.255 area 0
```

### show ip ospf neighbor

> OSPF如果对端down了，show ip ospf neighbor 就不会显示相关信息，如果端口重新up，就又会显示

```
R1#show ip ospf neighbor 

Neighbor ID     Pri   State           Dead Time   Address         Interface
192.168.80.2      1   FULL/BDR        00:00:36    192.168.80.2    FastEthernet0/0
192.168.79.3      1   FULL/BDR        00:00:32    192.168.79.3    FastEthernet0/1
```

- Neighbor ID：Router ID(这里即对端的)，每一个参与OSPF的Router都会有一个名字，称为Router ID。会以如下顺序命名

  1. 根据Router-id 指令来设定

     ```
     R1(config)#router ospf 1
     R1(config-router)#router-id 1.1.1.1
     Reload or use "clear ip ospf process" command, for this to take effect
     R1(config-router)#end
     R1#ena
     *Mar  1 00:33:30.711: %SYS-5-CONFIG_I: Configured from console by console
     R1#enable
     R1#clear ip ospf pro
     R1#clear ip ospf process 
     Reset ALL OSPF processes? [no]: yes
     R1#
     *Mar  1 00:33:45.143: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.80.2 on FastEthernet0/0 from FULL to DOWN, Neighbor Down: Interface down or detached
     *Mar  1 00:33:45.183: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.79.3 on FastEthernet0/1 from FULL to DOWN, Neighbor Down: Interface down or detached
     *Mar  1 00:33:45.303: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.80.2 on FastEthernet0/0 from LOADING to FULL, Loading Done
     *Mar  1 00:33:45.307: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.79.3 on FastEthernet0/1 from LOADING to FULL, Loading Done
     ```

     在R2中查看

     ```
     Neighbor ID     Pri   State           Dead Time   Address         Interface
     1.1.1.1           1   FULL/DROTHER    00:00:35    192.168.80.1    FastEthernet0/0
     192.168.80.3      1   FULL/DR         00:00:01    192.168.80.1    FastEthernet0/0
     192.168.79.3      1   FULL/BDR        00:00:35    192.168.78.3    FastEthernet0/1
     R2#
     *Mar  1 00:31:45.083: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.80.3 on FastEthernet0/0 from FULL to DOWN, Neighbor Down: Dead timer expired
     R2#show ip ospf neighbor 
     
     Neighbor ID     Pri   State           Dead Time   Address         Interface
     1.1.1.1           1   FULL/BDR        00:00:38    192.168.80.1    FastEthernet0/0
     192.168.79.3      1   FULL/BDR        00:00:35    192.168.78.3    FastEthernet0/1
     ```

     

  2. 如果没有设置Router-id，则使用Loopback Interface里最大的IP Address来做Router ID

  3. 如果没有设置Loopback interface，使用参与OSPF的的其他Interface里最大的IP Address来做Router ID

- Pri

  优先级，优先机比较大的router会称为DR，第二大的会称为BDR，其他的会成为DROTHER

- State

  如果两组neighbor能保持full就是正常的。这里也可以看是DR还是BDR

- Address

  ==对端互联的IP==

- Dead Time

  Dead Time 从40 sec开始计算，通常到30 sec ，因为收到回应就会重新设为 40 sec，但是如果倒数到0 sec也收不到回应，就表示neighbor 下线了

- interface

  ==本端互联的端口==