# ARP

[TOC]

已经知道了一个机器（主机或路由器）的IP地址，如何找出其相应的硬件地址？

地址解析协议 ARP 就是用来解决这样的问题的

==ARP的作用==：从网络层使用的 IP 地址，解析出在数据链路层使用的硬件地址。

通信时使用了两个地址：

- IP 地址（网络层地址）
- MAC 地址（数据链路层地址）

<img src="..\..\..\imgs\_Net\计算机网络\Snipaste_2020-08-23_15-15-25.png" style="zoom:80%;" />

### APR要点

- 不管网络层使用什么协议, 在实际网络的链路上传送数据时, 最终还是==必须使用硬件地址==

- 当主机 A 欲向本局域网上的某个主机 B 发送 IP 数据报时，就先在其 ARP 高速缓存中查看有无主机 B 的 IP 地址。
  如有，就可查出其对应的硬件地址，再将此硬件地址写入 MAC 帧，然后通过局域网将该 MAC 帧发往此硬件地址。
  如没有， ARP 进程在本局域网上广播发送一个 ARP 请求分组。收到 ARP 响应分组后，将得到的 IP 地址和硬件地址的映射写入 ARP 高速缓存。

  <img src="..\..\..\imgs\_Net\计算机网络\Snipaste_2020-08-23_15-16-20.png" style="zoom:80%;" />

### 使用ARP的四种典型情况

- 发送方是主机，要把 IP 数据报发送到本网络上的另一个主机。这时用 ARP 找到目的主机的硬件地址。 
- 发送方是主机，要把 IP 数据报发送到另一个网络上的一个主机。==这时用 ARP 找到本网络上的一个路由器的硬件地址(由路由器执行ARP)。==剩下的工作由这个路由器来完成。 
- 发送方是路由器，要把 IP 数据报转发到本网络上的一个主机。这时用 ARP 找到目的主机的硬件地址。 
- 发送方是路由器，要把 IP 数据报转发到另一个网络上的一个主机。这时用 ARP 找到本网络上另一个路由器的硬件地址。剩下的工作由这个路由器来完成。 

### 应当注意的问题

- ARP 是解决==同一个局域网==上的主机或路由器的 IP 地址和硬件地址的映射问题。
- 如果所要找的主机和源主机不在同一个局域网上，那么就要通过 ARP 找到一个位于本局域网上的某个路由器的硬件地址，然后把分组发送给这个路由器（网关），让这个路由器把分组转发给下一个网络。剩下的工作就由下一个网络来做。

<img src="..\..\..\imgs\_Net\计算机网络\Snipaste_2020-08-23_17-27-20.png" style="zoom:80%;" />

## 流程

参考：

https://www.quora.com/How-does-a-host-know-if-it-has-to-send-the-packet-to-the-switch-or-the-router

Address Resolution Protocol简称arp协议(基于layer 2)，网络设备要互相通信既要知道IP也要知道MAC，将IP与MAC之间相互映射就是arp协议的主要功能。主要在LAN中使用

> tip 可以使用`sudo ip n flush dev <NIC>`手动清空arp cache

arp分以下几种情况：

==但是不管怎样本机的subnet mask都会被使用，用于校验dst是否在同一LAN，即判断`目标IP & 本机subnet mask ==本机IP & 本机subnet mask  `==

### dst in the LAN

```mermaid
graph LR
a(a 192.168.100.103) -->|ping| b(b 192.168.0.100)
```

1. 先校验是否在同一LAN

2. a首先会在自己的neighbor table中查找是否有b的映射，如果有就不发送arp请求。直接发送数据帧

3. 否则a会向广播地址(IP/MAC 255.255.255.255/00:00:00:00:00:00)发送arp包询问b的Mac地址

   ![2021-06-21_22-04](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-06-21_22-04.28k56b6upfq.png)

   ==LAN中所有的host都可以收到广播地址发送的arp请求==如果和自己的IP不匹配就会丢弃arp请求。这router同样可以收到arp请求，但是router工作在layer 3不会发送layer 2的请求

4. 然后b返回一个arp响应给a，并在b的neighbor table中记录a IP和MAC地址的映射

   ![2021-06-21_22-10](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-06-21_22-10.6byw4e6y4bs0.png)

5. a在收到b的响应后在自己的neighbor table中同样记录b相对的映射

   ```
   cpl in ~ λ ip n
   192.168.0.1 dev wlp1s0 lladdr c0:a5:dd:d4:80:82 REACHABLE
   192.168.0.100 dev wlp1s0 lladdr 12:83:10:10:9c:65 STALE
   ```

==广播地址也会记录a和b的MAC地址和IP映射==

### dst do not in the LAN

如果目的地不在LAN中

```mermaid
graph LR
a(a 192.168.0.103)-->b(router 192.168.0.1) -->c(b 153.81.21.19)
```

1. 先校验是否在同一LAN
2. a先查询route table发现next hop，如果next hop MAC地址在neighbor table中就不执行arp请求，反之会向广播地址询问Router的MAC地址
3. ==IP数据包头src为192.168.0.103，dst为153.81.21.19。MAC数据包头src为192.168.0.103MAC，dst为192.168.0.1MAC==
4. 交给路由器封装IP头部src为192.168.0.103 dest为153.81.21.19

