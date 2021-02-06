# LVS

参考：

http://www.zsythink.net/archives/2134

## 概述

LVS，全称Linux Virtual Server，是一个虚拟的服务器集群系统。通过LVS实现负载均衡集群的方案属于“软件方案”。实现负载均衡的软件不知LVS一种，同样的可以通过nginx，haproxy等软件实现

**三层结构**

- 负载调度器（balancer）：集群对外的入口主机，负责接受请求，负载均衡
- 服务器池（server pool）：正真执行客户请求的服务器
- 共享存储（shared storage）：为服务器池提供一个共享的存储区

**角色**

- Director：调度器，用于接受用户请求，又称Balancer，Scheduler
- Real Server：简称RS，用于正真处理用户请求的服务器
- Client：客户端
- Virtual IP；简称VIP，LVS服务器的公网IP，与客户端通信
- Director IP：简称DIP，与RS通信的IP
- Real Server IP：简称RIP
- Client IP：简称CIP

<img src="..\..\imgs\_LVS\Snipaste_2020-11-22_13-40-23.png"/>

## 原理

借助iptables的INPUT链，如果符合规则，则将报文转发到POSTROUTING链，最终到达RS

<img src="..\..\imgs\_LVS\Snipaste_2020-11-22_17-47-05.png"/>

## LVS 调度模式

### LVS-NAT

<img src="..\..\imgs\_LVS\Snipaste_2020-11-22_13-50-25.png"/>

- 请求：源地址CIP，目标地址VIP，LVS将目标地址转换为RIP，具体请求那一台主机更具配置的算法
- 响应：源地址RIP，目标地址CIP，LVS将源地址转为换为VIP

