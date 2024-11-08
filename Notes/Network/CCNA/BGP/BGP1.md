# BGP1

ref:

https://en.wikipedia.org/wiki/Border_Gateway_Protocol

https://zhuanlan.zhihu.com/p/25433049

https://zhuanlan.zhihu.com/p/126754314

https://www.jannet.hk/border-gateway-protocol-bgp-zh-hans/

https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus3000/sw/unicast/503_u1_2/nexus3000_unicast_config_gd_503_u1_2/l3_overview.html#wp1114179

https://www.omnisecu.com/cisco-certified-network-associate-ccna/what-is-autonomous-system.php

## Digest

Border Gateway Protocol ( BGP )，中文通常翻译做边界网关协议。简单的可以理解为是 Internet 上一个核心的去中心化自治路由协议

BGP 通常用在大型的网络结构中，用作交换不同 AS 之间的路由信息，例如 ISP 与 ISP 之间的路由交换。

BGP 只是负责宣告 AS 与 AS 之间的路由

![2022-09-03_13-00](https://git.poker/dhay3/image-repo/blob/master/20220903/2022-09-03_13-00.2hpfsnjn57eo.webp?raw=true)

## Terms

### AS

Autonomous system( AS )，通常也被叫做自治系统。指由一个机构完全管理的 network（这里的 network 指的是通过 3 层路由设备构建的网络），例如 ISP 可能就会被划分到一个或者多个 AS。通过 AS，将 Internet 分隔层成独立的 routing domains

### ASN

Autonomous system number 是一个 16 bit 的数，即 10 进制最大值 65535 (BGP4 中该值扩展成 4 byte 32 bit)

0 和 65535 都是保留的 ASN

| 2-Byte Numbers | 4-Byte Numbers in AS.dot Notation | 4-Byte Numbers in plaintext Notation | Purpose                                                      |
| -------------- | --------------------------------- | ------------------------------------ | ------------------------------------------------------------ |
| 1 to 64511     | 0.1 to 0.64511                    | 1 to 64511                           | Public AS (assigned by RIR)[1](https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus3000/sw/unicast/503_u1_2/nexus3000_unicast_config_gd_503_u1_2/l3_overview.html#wp1114223) |
| 64512 to 65534 | 0.64512 to 0.65534                | 64512 to 65534                       | Private AS (assigned by local administrator)                 |
| 65535          | 0.65535                           | 65535                                | Reserved                                                     |
| N/A            | 1.0 to 65535.65535                | 65536 to 4294967295                  | Public AS (assigned by RIR)                                  |

### Router ID

用于标识 BGP 设备的 32 bit 的值（通常用 IPv4 的形式），==在 BGP 网络中必须是唯一的==

### Speaker

发送 BGP 报文(报文包含路由信息)的设备，它接受或产生新的报文信息，并发布（Advertise）给其他 BGP Speaker

### Peer/Neighbors

Peer，在 Cisco 中叫 Neighbours，中文也叫对等体。相互交换报文的 Speaker 就被称为 Peer

多个 Peer 可以组成 Peer Group，应用在 Peer Group 上的配置也会被应用在 Peer

### iBGP

Interior Gateway Protocol (IGP) 内部网关协议，在一个 AS 内部使用的路由协议。一个 AS 的路由器如果需要动态的知道路由，就需要通过 IGP 来实现。例如：OSPF, RIP, ISIS

Interior Border Gateway Protocol (iBGP) 通常也被叫做内部边界网关协议。如果两个 router 在相同的 AS 之间组成 Peers，就会称为 IGBP peers

但是 IGP 协议都好好的为什么还有一个 iBGP 呢？因为 IGP 例如 OSPF 存在性能瓶颈

BGP 使用 TCP，且 TCP 窗口能达到最大 65535，其他 IGP 例如 OSPF TCP 窗口只有默认的一个数据包大小。如果当网络规模庞大，传输的数据就会变大，如果每次只能发一个包，数据效率大大折扣

简单的可以概括 小规模网络 AS 内可以使用 IGP，大规模网络 AS 内使用 BGP

通过 IBGP peer 学到的路由，==会 re-advertised 到所有的 eBGP peers (没有 iBGP peers)==



### eBGP

External Gateway Protocol (EGP) 外部网关协议，多个 AS 之间交换路由的协议，已被 eBGP 替代

External Border Gateway Protocol (eBGP) 通常也被叫做外部边界网关协议。如果两个 router 在不同的AS之间组成Peers，就会称为 EBGP peers

如果运行 BGP 的设备在收到一条同设备的 ASN 就会丢弃该条路由，防止环路

通过 eBGP peer 学到的路由会 re-advertised 到所有的 iBGP and eBGP peers

## Why BGP comes

> zhihu 上的例子真的很形象。所以直接抄袭。。

BGP号称是使互联网工作的协议，看起来似乎很重要，为什么平常的生活中很少接触呢？似乎云里面也不怎么提BGP，我们来看看原因吧。

假设小明正在搭建一个云环境，提供虚拟机服务

![img](https://pic1.zhimg.com/80/v2-73b552f494d3fba7c3d29e9690052dd8_720w.png)

云里的虚机需要有互联网访问能力，于是小明向ISP（Internet service  provider）申请了一个公网IP，这里的ISP可以是联通，移动，电信等等。虚机们可以通过路由器的NAT/PAT（Network / Port address translation）将自己的私网IP转换成这个公网IP，然后小明在云中路由器上将ISP  router的地址设为默认路由。这样地址转换之后的IP包都发送到了ISP，进而发送到了互联网（这也是我们家用路由器能让家里的设备上网的原理)。这样小明的1.0版本云简单上线。这里小明不需要BGP。

![img](https://pic2.zhimg.com/80/v2-691fbc26690af6a9c1d999695b6b9239_720w.png)

版本上线之后怎么办？当然是开发下一版本！下一版本的需求是可以通过互联网访问虚机（也就是从互联网访问我们家里的电脑）。这个也不难，可以通过端口转发（Port  Forward），将虚机的一个端口与公网IP的端口进行映射。例如将虚机的22端口映射到公网IP的1122端口，那么可以通过互联网ssh到公网IP:1122，登陆虚机。这部分工作仍然是在小明的云中路由器完成。这样，小明的2.0版本云上线了，这里小明还是不需要BGP。

2.0版本虽然支持了从互联网访问虚机，但是还有问题：

1. 每个虚机每开放一个端口都需要映射一次
2. 公网IP的端口是有限的

为了解决这些问题，小明向联通申请了一些公网IP地址，对于需要从外网访问的虚机，直接给它们分配公网IP。这样小明的3.0版本云上线了，这里小明还是不需要BGP。因为：

1. 联通是小明云唯一连接的ISP，小明只能通过联通访问互联网，所以小明的云中路由器的默认路由只能设置成ISP 路由器的地址。
2. 小明云里面的公网IP都是联通分配的，联通当然知道该从哪个IP地址作为下一跳去访问那些IP地址。

![img](https://pic1.zhimg.com/80/v2-befa9834049a20930b56f67deb2802c0_720w.png)

联通的IP毕竟是有限的，而且联通还老是断线。这都发布3个版本了，小明决定干一票大的。



首先，小明向IANA（Internet Assigned Numbers  Authority）申请了自己的公网IP池。因为有了自己的公网IP，也必须要考虑申请AS号。AS号是一个16bit的数字，全球共用这60000多个编号。1 – 64511 是全球唯一的，而 64512 – 65535  是可以自用的，类似于私网网段。每个自治网络都需要申请自己的AS编号，联通的AS号是9800。

然后，小明分别向联通和电信买了线路，这样就算联通断线还能用电信。

那现在问题来了：

1. 联通或者电信怎么知道小明申请的公网IP是什么。换言之，我现在拨号拨进了联通宽带，我怎么才能访问到小明云的公网IP？
2. 小明的云中路由器的默认路由该设置到联通的ISP路由器，还是电信的？



终于，在小明的4.0版本云上，小明需要用BGP了。通过BGP，小明可以将自己云中的路由信息发送到联通，电信，这样ISP就知道了改如何访问小明的公网虚拟机，也就是说我们普通的使用者通过ISP，能访问到小明的网络。另一方面，通过在云中运行BGP服务，小明可以管理云中路由器的默认路由。

总的来说，要是你之前没有听过或者用过BGP，只能说你的网络还没有到那个规模  ：）

## Inside BGP

BGP 是 7 层协议，默认使用 179 端口作为 Socket Pair

### BGP state

一般来说，如果设定没有问题的话，BGP peers 就会变成 Established 的状态。但是实际上，peers 在进入 Established 之前还需要经过几个状态

![img](https://pic2.zhimg.com/80/v2-c4ee273cc1c32325c14b0f6a87da93d1_720w.jpg)

- IDLE

  Router正在搜寻Routing Table，找一条能够连接Neighbor路径（==不会使用Default Route==）

- CONNECT

  Router 已经找到连接Neighbor的路径，并且完成了TCP 3-way handshake

- OPEN SENT

  已经发送了BGP的OPEN封包，告诉对方希望建立Peers

- OPEN CONFIRM

  收到了Neighbor 回传封包，对方赞成建立Peers

- ESTABLISHED

  两个Neighbor已经成功建立了Peers

- ACTIVE

  Router仍然处于主动传送封包的状态，收不到对方回传，如果持续见到此状态的话，==代表Peers并未成功建立==

### BGP 引入 IGP 路由

BGP 协议本身不发现路由，因此需要将其他路由引入到 BGP 路由表，实现 AS 间的路由互通。当一个 AS 需要将路由发布到其他 AS 时，AS 边缘路由器会在 BGP 路由表中引入 IGP 的路由。为了更好的回话网络， BGP 在引入 IGP 的路由时，可以使用路由策略进行路由过滤和路由属性设置

BGP  引入路由时的两种方式

1. Import 

   按照路由协议类型引入，例如 OSPF、ISIS 等路由协议引入到 BGP 路由表。为了保证引入的 IGP 路由有效性，Import 方式还可以引入静态路由和直连路由

2. Network

   逐条将 IP 路由表中已经存在的路由引入到 BGP 路由表中

### BGP 选路规则

在 BGP 路由表中，到达同一目的可能存在多条路由。此时 BGP 会选择其中一条路由作为最佳路由，并只把此路由发送给 BGP Peers

BGP 为了选出最佳路由，会根据 BGP 的属性优先级来选择。主要包含如下几种属性

1. Origin

   标记一条路由是怎么成为 BGP 路由的。可以是如下几个值

   - IGP

     具有最高的优先级。通过 network 命令注入到 BGP 路由标的路由，其 Original 属性为 IGP

   - EGP

     优先级次之。通过 EGP 得到的信息，其 Origin 属性为 EGP

   - Incomplete

     优先级最低。通过其他方式学习到的路由。比如 BGP 通过 import-route 命令引入的路由，其 Origin 属性为 Incomplete

2. AS_path

   

