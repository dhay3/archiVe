# OSPF

reference：

https://www.jannet.hk/open-shortest-path-first-ospf-zh-hant/

https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/iproute_ospf/configuration/xe-16/iro-xe-16-book/iro-cfg.html

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

OSPF是一个link-state routing protocal（链路状态协议），link简单的说就是router的interface。OSPF的运作原理就是把自己每个interface连接的network告诉其他的router，于是router便可以计算出自己的routing table

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

要看external  LSA，可以在Router输入指令`show ip ospf database external`，以下用R1做例子。R1 只收到一条来自远方R6(6.6.6.6)所发布的external LSA，告诉R1从它那里可以到达192.168.67.0这个network，cost 是20，而metric type是2，又称E2。

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

收到了router LSA和external LSA

## Network Type

根据 自动/手动Neighbor、是否DR选举、10,40/30,120 Timer设定这三大元素，Cisco router把不同的设定组合归纳为5中模式，我们把这玩意称为Network Type

![2021-11-25_22-20](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211123/2021-11-25_22-20.v60hfqvsywg.png)

## connect State

Router们称为neighbor其实中间经过很多状态（称为Adjacency），包含如下几种

1. Down：没有发放Hello Message
2. Init：刚刚向对方发放Hello Message
3. 2-way：他们开始沟通，选出DR和BDR。
4. ExStart：预备交换Link信息
5. Exchange：正在交换DBD（Database Descriptors），可以把DBD看成Link信息的一些目录，先给对方目录，让对方回复那些LSA（link state advertisements）是他需要的
6. Loading：正在交换LSA
7. Full：两只router称为adjacency，在这是，同一个area的router里面的topology table应该完全相同

## neighbor

### 成为neighbor的必要条件

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

## command

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
