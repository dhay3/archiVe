# Day47 - QoS(Part 2)

## Classification

*Classification organize network traffic(packets) into traffic classes(categories)*

简单的说 Classification 就是将不同的流量按照规则分类，主要规则有如下几种

1. An ACL. Traffic which is permitted by the ACL will be given certain treament, other traffic will not

   > 这点和 Dynamic NAT 中使用 ACL 的方式类似，并不是直接 deny 或者是 permit 流量，这里可以理解成 map reduce 中的 map 操作

2. NBAR(Network Base Application Recognition) performs a deep packet inspection, looking beyond the Layer3 and Layer 4 information up to Layer 7 to identify the specific kind of traffic

   通过报文的 2 - 7 层的信息来划分

   例如在 2 - 3 层的报文头中有一些特殊的字段

   - *The PCP(Priority Code Point) field of the 802.1Q tag(in the Ethernet header) can be used to identify high/low priority traffic*

     只存在于 802.1Q 报文头中，即 VLAN 报文头

   - *The DSCP(Differentiated Services Code Point) field of the IP header can also be used to identify high/low priority traffic*

     大多数都会使用 DSCP

### PCP

PCP 也被称为 CoS(Class of Service)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_18-04.4jphbk32ng5c.webp)

一共 3 bit 最大的值为 8，不同的值对应的流量类型不同

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_18-07.11k4h5568q9s.webp)

因为 PCP 只在 dot1q 报文头中，所以支持出现在如下的 connections 中

1. trunk links
2. access links with a voice VLAN

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_18-11.3c5qjhuxg18g.webp)

即上图中红色的链路，PH1/2 和 SW1/2 匹配第二种 connection，SW1 - SW2 和 SW1 - R1 匹配第一种 connection

除此外其他链路都不会有 dot1q tag，所以也就没有 PCP，其他的设备就不能根据 PCP 来对流量进行优先级排序

### DSCP

在 IPv4 报文头中有一部分被称为 ToS byte(8 bits)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_18-16.1ap6vl6zqdc0.webp)

即红框中部分，由 DSCP 和 ECN 组成

早的时候 ToS byte 并不是由 DSCP 和 ECN 组成的，而是由如下部分组成

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_18-18.5tflontvfvr4.webp)

实际只有 3 bits IPP 来其划分流量的功能，最大只能有 8 种

而现在的 ToS byte

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_18-19.6nbtaund4lts.webp)

使用 DSCP 来划分流量，一共占了 6 bits，最大 64 种，所以可以提供对流量划分细粒度更小的判断

#### IPP

标准的 IPP 类似与 PCP，如果值为

- 6 和 7 表示 network control traffic(例如 OSFP messages)
- 5 = voice
- 4 = video
- 3 = voice signaling
- 0 = best effort

一共只保留 6 位，可以划分 6 种不同的网络流量，但是实际生产中可能需要更加精细的划分流量，所以 DSCP 取代了 IPP

#### DSCP value set

DSCP 中只要注意几个值的集合

- Default Forwarding(DF) = best effort traffic
- Expedited Forwarding(EF) = low loss/latency/jitter traffic(usually voice)
- Assured Forwarding(AF) = A set of 12 standard values
- Class Selector(CS) = A set of 8 standard values, provides backward compatibility with IPP

> QoS 配置并不是 CCNA 考试中的内容，这边只做简单的介绍

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_09-59.7joqp2cozxxc.webp)

- `R1(config)#class-map <map name>`

  定义一个映射集

- `R1(config)#match dscp <value>`

  映射匹配指定的 DSCP 值

##### DF/EF

###### ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_10-03.4vdq49fnxu68.webp)

DF 对应 DSCP 0

EF 对应 DSCP 46

##### AF

DF 和 EF 比较好理解，AF 有点特殊

在 AF 中有 2 bits 被用于表示 Drop Precedence

*Higher drop precedence = more likely to drop the packet during congestion*

3 bits 表示 class

*Higher class = more priority*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_10-07.5bxj4foi6rnk.webp)

X 部分值为 Class 部分的 decimal, Y 部分值为 Drop Precedence 部分

例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_10-09.5u0kpgcigum8.webp)

就表示 AF11,对应 DSCP 10

有一个公式将 AF 转为对应的 DSCP 值

*Formula to convert from AF value to decimal DSCP value: **8X + 2Y***

所以按照 AF class 以及 drop precedence 的规则，可以得出 AF41 是最优的，对应 DSCP 值 34，AF13 是最差的，对应 DSCP 14

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_10-16.3afb763dvhog.webp)

##### CS

CS 是用于兼容 IPP 的，所以只能占用 3 bits

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_10-29.1p43guzwtlfk.webp)

如果需要将 CSX 转为 DSCP，直接将 X 乘以 8 即可

#### RFC 4954

RFC 4954 是用于规范以及建议 DSCP 值使用的，其中规定

- Voice traffic:EF
- Interactive video:AF4X
- Streaming video:AF3X
- High priority data:AF2x
- Best effort:DF

## Trust Boundaries

*The trust boundary of a network defines where devices trust/don’t trust the QoS marking of received messages*

- If the markings are trusted, the device will forward the message without changing the markings
- If the markings aren’t trusted, the device will change the marking according to the configured policy

> 以出向为例 boundary 左侧过来的流量为不可信的，boundary 右侧及其本身的流量为可信的

例如 假设 SW1 和 SW2 是 trust boundaries

> PCP 和 DSCP 可以同时使用

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_10-55.509vldz7pf28.webp)

PH1 发送 voice traffic EF/CoS5 到 SW1，因为是 trust boundaries,SW1 会发送报文改成 DF/CoS0(这里 SW 也会对 3 层报文中的 DSCP 做改动吗？)，到 R1 发送只会有 DF(因为 CoS 只在 VLAN 中)

这个配置并不是合理的，因为正常我们就是需要保证 voice traffic 的优先级至少不是 DF 的。所以通常有如下规则

*If an IP phone is connected to the switch port, it is recommanded to move the trust boundary to the IP phones*

*This is done vai configuration on the switch port connected to the IP phone*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_11-06.1q1rg3mse1d.webp)

PH1 发送 traffic voice EF/CoS5 到 SW1 发送仍然会是 EF/CoS5，而 PC2 发送过来的 data traffic 虽然 DSCP 值是 EF，但是过了 trusted boundary 就会变成 DF

## Queuing/Congestion Management

*When a network device receives traffic at a faster rate than it can forward the traffic out of the appropriate interface, packets are placed in that interface’s queue as the they wait to be forwarded*

当网络设备的接受的流量大于自己能转发的流量，报文就会在对应转发端口堆积成队列

When the queue becomes full, packets that don’t fit in the queue are dropped(tail drop)

*RED and WRED drop packets early to avoid tail drop*

在 QoS 中 classification 会根据流量的某些值将报文放到合适的 queue 中，一个接口同时只能发送一个报文，所以 scheduler 会被应用，决定那个 queue 会被采用

Prioritization allows the scheduler to give certain queues more priority than others

例如

ingress traffic 到 router，会先做路由判断，然后做 classification，由 classification 决定将报文放到接口的那个 queue 中，然后有 scheduler 决定使用那个 queue

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_11-35.39ldlxdj5ry8.webp)

常用的 scheduling methods 有

1. round-robin = pakcets are taken from each queue in order, cyclically
2. weighted = more data is taken from high priority queues each time the scheduler reaches the queue

而在 QoS 中使用的最多的一种 scheduling method 是 **CBWFQ(Class-Based Weighted Fair Queuing)**，是 round-robin 和 weighted 的结合

先取 weighted 最大的 queue，然后取第二的，以此循环

例如 scheduler 会先取 60% 权重的 queue，然后取 15% 权重，然后 10% 权重的，以此循环

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_12-00.4ssrx1b1po1s.webp)

但是对于 voice/video traffic 来言是不适合使用(含有) round-robin scheduling 的，因为即使 queue wighted 最高还是需要做轮询的操作，这无疑会增加 delay 和 jitter

为了解决这个问题可以使用 **LLQ(Low Latency Queuing)**

*LLQ designates one (or more) queues as strict priority queues*

*This means that if there is traffic in the queue, the scheduler will always take the next packet from that queue until it is empty*

> 可以让 queue 固定出列顺序

例如 使用了 LLQ，红色的 queue 被指定为 strict priority queue，一旦这个 queue 中有报文，scheduler 就会调度这个 queue，直到这个 queue 中没有报文为止

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_11-56.5zcz24feflz4.webp)

如果 strict priority queue 中一直有报文，就会导致一个问题，会让其他 queue 中的报文得不到处理。我们可以通过 *Policing* 来解决这个问题

## Shaping and Policing

shaping 和 policing 都被用于控制流量的速率

- shapping buffers traffic in a queue if the traffic rate goes over the configured rate

  只针对与 queue

- policing drops traffic if the traffic rate goes over the configured rate

  针对全体

  ‘Burst’ traffic over the configured rate is allowed for a short period of time

  The amount of burst traffic allowed is configurable

这里有一个问题，为什么需要限制流量的速率

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_12-15.us4p6btrl6o.webp)

假设 Customer router 和 ISP router 直连，端口为 gigabyteEthernet，逻辑上最大能到 1Gbps，但是 Customer 只向 ISP 付了 300 mbps 带宽的钱，所以 ISP 在 g0/0 inbound 配置了 policing 来限制 Customer 的流量带宽到 300 mbps

因为 Customer 带宽一共有 300 mbps，超过 300 mbps 的流量就会被 ISP 丢弃，这样也会导致有一个问题就是 二进制退避，所以 Customer 在互联的 g0/0 outbound 配置了 shaping 以限制接口出流量为 300 mbps

## LAB

> 注意 QoS 的配置并不在 CCNA 范围内

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_12-27.xe2chtqoiv4.webp)

Configure the following QoS settings on R1 and apply them outbound on interface G0/0/0

> 每台设备的 QoS 配置只表明当前设备对应流量的优先级，并不会影响在整个链路中的优先级
>
> 例如 R1 上 FTP 流量配置成最高的优先级，但是 R2 并不认为 FTP 流量是最高的优先级，会按照默认的方式处理 FTP 流量
>
> 这个现象也被称为 **PHB(Per Hop Behavior)**, QoS 配置在每台设备上独立，流量怎么处理完全看设备自己的 QoS 配置

在测试前先看一下 PC1 ping SRV1(jeremysitlab.com) 的报文中 DSCP 的值是多少

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-14_12-51.40adx6ecl4e8.webp)

在 R1 上可以看到报文的 DSCP 值为 0，即 DF

### 0x01

Mark HTTPS traffic as AF31. Provide minimum 10% bandwidth as a priority queue

> 通常并不会给 HTTPS 流量加入到 priority queue 中，通常都是 vocie traffic/vedio traffic

需要先声明一个 class map 匹配 HTTPS

```
R1(config)#class-map HTTPS
R1(config-cmap)#match protocol https
```

使用 `show run | se class-map` 来查看

```
R1(config-cmap)#do show run | se class-map
class-map match-all HTTPS
 match protocol https
```

这里会多加一个条件 `match-all` 表示只有匹配后面所有的条件才会匹配 class-map，即逻辑与

```
R1(config)#class-map ?
  WORD       class-map name
  match-all  Logical-AND all matching statements under this classmap
  match-any  Logical-OR all matching statements under this classmap
  type       type of the class-map
```

如果想要只配置其中的一个条件，可以使用 `match-any`

现在我们还要指定匹配 class-map 的流量需要被怎么处理，通过 policy-map 来完成

> ==一个接口上只能使用一个 policy-map==

```
#定义一个 policy-map
R1(config)#policy-map G0/0/0_OUT
#匹配条件为满足 class-map HTTPS
R1(config-pmap)#class HTTPS
#匹配条件的报文将 DSCP 值设置为 AF31
R1(config-pmap-c)#set ip dscp af31
#配置最小带宽，并将匹配的报文放入 priority queue 中
R1(config-pmap-c)#priority percent 10
```

现在还需要将 policy-map 应用在端口上

```
R1(config-pmap-c)#int g0/0/0
R1(config-if)#service-policy output G0/0/0_OUT
```

### 0x02

Mark HTTP traffic as AF32. Provide minimum 10% bandwidth

```
R1(config)#class-map HTTP
R1(config-cmap)#match protocol http
R1(config)#policy-map G0/0/0_OUT
R1(config-pmap)#class HTTP
R1(config-pmap-c)#set ip dscp af32
#配置最小带宽
R1(config-pmap-c)#bandwidth percent 10
```

这里因为已经在 0x01 中将 G0/0/0_OUT policy-map 应用在 R1 G0/0/0 outbound 了，所以不需要再配置一遍

### 0x03

Mark ICMP traffic as CS2. Provide minimum 5% bandwidth

```
R1(config)#class-map ICMP
R1(config-cmap)#match protocol icmp
R1(config)#policy-map G0/0/0_OUT
R1(config-pmap)#class ICMP
R1(config-pmap-c)#set ip dscp cs2
#配置最小带宽
R1(config-pmap-c)#bandwidth percent 5
```

这里因为已经在 0x01 中将 G0/0/0_OUT policy-map 应用在 R1 G0/0/0 outbound 了，所以不需要再配置一遍

使用 `show run | se map` 检查一下配置

```
R1(config-pmap-c)#do show run | se map
class-map match-all HTTP
 match protocol http
class-map match-all HTTPS
 match protocol https
class-map match-all ICMP
 match protocol icmp
policy-map G0/0/0_OUT
 class HTTPS
  priority percent 10
  set ip dscp af31
 class HTTP
  bandwidth percent 10
  set ip dscp af32
 class ICMP
  bandwidth percent 5
  set ip dscp cs2
```

### 0x04

Use simulation mode to view the DSCP markings of packets

- When pinging jeremysitlab.com from PC1

  R1 Inbound DSCP 的值为 0x00

  R1 Outbound DSCP 的值为 0x10 = 16 = CS2

  R2 Inbound DSCP 的值为 0x10

  R2 Outbound DSCP 的值为 0x10

  SRV1 回的包 DSCP 值均为 0x00 因为 service-policy 绑定在 R1 G0/0/0 outbound

- When accessing jeremysitlab.com from PC1 via HTTP

  R1 Inboud DSCP 的值为 0x00

  R2 Outbound DSCP 的值为 0x1C = 28 = AF32

- When accessing jeremysitlab.com from PC1 via HTTPS

  R1 Inbound DSCP 的值为 0x00

  R2 Outbound DSCP 的值为 0x1A = 26 = AF31

**references**

1. [^https://www.youtube.com/watch?v=4vurfhVjcMM&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=91]