# OSPF

reference：

https://www.jannet.hk/open-shortest-path-first-ospf-zh-hant/

## 概述

Open Shortest Path First是一个基于链路状态(Link-State Routing Protocol)的IGP协议，由于RIP的种种限制，RIP组件被OSPF取代。每个在OSPF里的Router都会向Neighbor交换自己的Link-State，当Router收到这些Link-State之后，就会运用Dijkstra Alogrith来计算出最短的路径

## 区域

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

有interface连接其他AS的Router就是ASBR，在这个网络中，除了执行OSPF之外，最右边脸色的区域正在执行另一个Routing Protocol RIP。R6为OSPF 与 RIP的连接点，即连个AS的连接点，称为ASBR



## Example

![2021-11-09_01-10](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211108/2021-11-09_01-10.1u7c56nzfq74.png)

### R1

```
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

OSPF配置

```
conf t
router ospf 1
network 192.168.80.0 0.0.0.255 area0
network 192.168.79.0 0.0.0.255 area1
```

### R2

```
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

OSPF配置

```
conf t
router ospf 1
network 192.168.80.0 0.0.0.255 area0
network 192.168.78.0 0.0.0.255 area2
```

### R3

```
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

OSPF配置

```
conf t
router ospf 1
network 192.168.78.0 0.0.0.255 area2
network 192.168.79.0 0.0.0.255 area1
```

如果配置成功后会显示，LOADING to FULL

```
*Mar  1 01:11:44.203: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.80.2 on FastEthernet0/1 from LOADING to FULL, Loading Done
```

Router们称为neighbor其实中间经过很多状态（称为Adjacency），包含如下几种

1. Down：没有发放Hello Message
2. Init：刚刚向对方发放Hello Message
3. 2-way：他们开始沟通，选出DR和BDR。
4. ExStart：预备交换Link信息
5. Exchange：正在交换DBD（Database Descriptors），可以把DBD看成Link信息的一些目录，先给对方目录，让对方回复那些LSA（link state advertisements）是他需要的
6. Loading：正在交换LSA
7. Full：两只router称为adjacency，在这是，同一个area的router里面的topology table应该完全相同

## show running-config | se r o

查看ospf配置

```
R1#show running-config | se r o
router ospf 1
 log-adjacency-changes
 network 192.168.79.0 0.0.0.255 area 1
 network 192.168.80.0 0.0.0.255 area 0
```

## show ip ospf neighbor

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

- 

## IGP协议区别

| 对比项   | RIP                                                          | OSPF                                                         | IS-IS                                                        |
| -------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 协议类型 | IP层协议                                                     | IP层协议                                                     | 链路层协议                                                   |
| 适用范围 | 应用于规模较小的网络中。例如：校园网等结构简单的地区性网络   | 应用与规模适中的网络中，最多支持几百台路由器。例如：中小型企业网络 | 应用于规模较大的网络中。例如：大型ISP(Internet Service Provider) |
| 路由算法 | 采用距离矢量(Distance-Vector)算法，通过UDP报文进行路由信息交换 | 采用最短路径SPF(Shortest Path First)算法。通过链路状态通告LSA(Link state Advertisement)描述网络拓扑，依据网络拓扑生成一棵最短路径树SPT，计算书到网络中所有 目的地的最短路径，进行路由信息交换 | 采用最短路径SPF算法。依据网络拓扑生成一棵最短路径树SPT，计算出到网络中所有目的地的最短路径。在IS-IS中，SPF算法分别独立在Level-1和Level-2数据库中运行 |
| 收敛速度 | 慢                                                           | 快，小于1s                                                   | 快，小于1s                                                   |
| 扩展性   | 不能扩展                                                     | 通过划分区域扩展网络支撑能力                                 | 通过Level路由器扩展网络支撑能力                              |

