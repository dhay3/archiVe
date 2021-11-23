# OSPF

reference：

https://www.jannet.hk/open-shortest-path-first-ospf-zh-hant/

## 概述

Open Shortest Path First是一个基于链路状态(Link-State Routing Protocol)的IGP协议，由于RIP的种种限制，RIP组件被OSPF取代。每个在OSPF里的Router都会向Neighbor交换自己的Link-State，当Router收到这些Link-State之后，就会运用Dijkstra Alogrith来计算出最短的路径

## Link-state Advertisement (LSA)

OSPF是一个link-state routing protocal（链路状态协议），link简单的说就是router的interface。OSPF的运作原理就是把自己每个interface连接的network告诉其他的router，于是router便可以计算出自己的routing table

LSA中包含几项信息：LSA由谁传送出来、连着的是什么network、以及去这个network的cost

LSA有不同的类型，用下面的一个例子说明

![2021-11-24_00-30](https://raw.githubusercontent.com/dhay3/image-repo/master/20211123/2021-11-24_00-30.62tmmwf8jrg0.png)

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

### router LSA

### network LSA



## Area

OSPF是设计给一个大型网络使用的，为了解决管理上的问题，OSPF采用了一个hierachical system(分层系统)，把大型的OSPF分割成多个Area。

Area有两种表达方式，可以是一个16bits数字（0-65535），或者使用类似IP的方式（OID），例如：192.168.1.1，前者比较常见。Area0（或0.0.0.0）是一个特别的Area，我们称为Backbone Area（骨干），==所有其他Area必须与Backbone Area连接，这是规定==

![2021-11-08_23-24](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211108/2021-11-08_23-24.18o156168r1c.png)

### Internal Router

Router上的所有interface都属于同一个area，R3和R5就是Internal Router

### Backbone Router

最少一个interface连接Backbone Area（Area 0），所以R2、R3和R4都是Backbone Router

### Area Border Router(ABR)

连接连个Area或以上的Router称为ABR，R2和R4都是ABR

### Autonomous System Border Router(ASBR)

有interface连接其他AS的Router就是ASBR，在这个网络中，除了执行OSPF之外，最右边蓝色的区域正在执行另一个Routing Protocol RIP。R6为OSPF 与 RIP的连接点，即连个AS的连接点，称为ASBR

## Area Type

OSPF 的 Area分为几种Backbone Area(Area 0)、Standard Area、Stub Area、Totally Stubby Area、Not-so-stubby Area、Totally Not-so-stubby Area

## Network Type

pending to write



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

- Neighbor ID：Router ID，每一个参与OSPF的Router都会有一个名字，称为Router ID。会以如下顺序命名

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

  对端的IP

- Dead Time

  Dead Time 从40 sec开始计算，通常到30 sec ，因为收到回应就会重新设为 40 sec，但是如果倒数到0 sec也收不到回应，就表示neighbor 下线了
