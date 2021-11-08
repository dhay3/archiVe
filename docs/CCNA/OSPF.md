# OSPF

reference：

https://www.jannet.hk/open-shortest-path-first-ospf-zh-hant/

## 概述

Open Shortest Path First是一个基于链路状态(Link-State Routing Protocol)的IGP协议，由于RIP的种种限制，RIP组件被OSPF取代。每个在OSPF里的Router都会向Neighbor交换自己的Link-State，当Router收到这些Link-State之后，就会运用Dijkstra Alogrith来计算出最短的路径

## 区域

OSPF是设计给一个大型网络使用的，为了解决管理上的问题，OSPF采用了一个hierachical system(分层系统)，把大型的OSPF分割成多个Area。

Area有两种表达方式，可以是一个16bits数字（0-65535），或者使用类似IP的方式（OID），例如：192.168.1.1，前者比较常见。Area0（或0.0.0.0）是一个特别的Area，我们称为Backbone Area（骨干），==所有其他Area必须与Backbone Area连接，这是规定==

OSPF协议通过将自治系统划分为不同的区域解决LSDB(LInk-State Database)频繁更新的问题。逻辑上将路由器划分为不同的组，每个组区域号(Area ID)来标识。==区域的边界是路由器，而不是链路。一个网段只能属于一个区域。在OSPF里的==

## IGP协议区别

| 对比项   | RIP                                                          | OSPF                                                         | IS-IS                                                        |
| -------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 协议类型 | IP层协议                                                     | IP层协议                                                     | 链路层协议                                                   |
| 适用范围 | 应用于规模较小的网络中。例如：校园网等结构简单的地区性网络   | 应用与规模适中的网络中，最多支持几百台路由器。例如：中小型企业网络 | 应用于规模较大的网络中。例如：大型ISP(Internet Service Provider) |
| 路由算法 | 采用距离矢量(Distance-Vector)算法，通过UDP报文进行路由信息交换 | 采用最短路径SPF(Shortest Path First)算法。通过链路状态通告LSA(Link state Advertisement)描述网络拓扑，依据网络拓扑生成一棵最短路径树SPT，计算书到网络中所有 目的地的最短路径，进行路由信息交换 | 采用最短路径SPF算法。依据网络拓扑生成一棵最短路径树SPT，计算出到网络中所有目的地的最短路径。在IS-IS中，SPF算法分别独立在Level-1和Level-2数据库中运行 |
| 收敛速度 | 慢                                                           | 快，小于1s                                                   | 快，小于1s                                                   |
| 扩展性   | 不能扩展                                                     | 通过划分区域扩展网络支撑能力                                 | 通过Level路由器扩展网络支撑能力                              |

## Router ID

一台路由器要运行OSPF协议，必须存在Router ID是一个32bit无符号整数，是一台路由器在AS中唯一标识符。

Router ID的选取有两种方式：

- 通过命令手动配置

- 路由器自动设定

  如果没有手动配置Router ID，路由器会从当前接口的IP地址中自动选取一个作为Router ID。其选择顺序是：

  1. 优先从Loopback地址中选择最大的IP地址作为Router ID
  2. 如果没有配置Loopback 接口，这在接口地址中选取最大的IP地址作为Router ID

