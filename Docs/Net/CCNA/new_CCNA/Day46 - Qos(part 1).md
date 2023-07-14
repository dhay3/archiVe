# Day46 - Qos(part 1)

## IP Phones

传统的手机使用 Public Switched telephone network(PSTN)，通常也被称为 POSTS(Plain Old Telephone Service)

而 IP phone 使用 VoIP(Voice over IP), 可以通过 network 来实现电话互通

思科的 IP phone 长这样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_14-09.6f0h374dgdc.webp)

IP phone 和 end host 一样，都是通过 Switch 或者是 WIFI 接入网络的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_14-10.50ucw4pnrc74.webp)

IP phone 内置了一个 3 口的交换机

1. 1 port is the ‘uplink’ to the external switch
2. 1 port is the ‘downlink’ to the PC
3. 1 port connects internally to the phone itself

逻辑上可以想象成下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_14-12.2g0l6f4hyi2o.webp)

因为这种模型，通常需要将 voice traffic(from the IP phone) and data traffic(from the PC) 通过 VLAN 隔离

通常使用 voice VLAN 来隔离，当流量是从 PC 来的不会被打上 VLAN tag，而当流量是从 phone 来的会被打上 VLAN tag

因为 IP phone 逻辑上内部包含一个 3 口交换机，所以上面的拓扑就可以抽象成下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_16-47.4fb5fq0trb40.webp)

### IP phones/Voice VLAN

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_14-20.66jmonhbth0.webp)

所有的配置均在 SW1 上完成

PC1 正常的流量走 VLAN10，如果是从 PH1 来的走 VLAN11 会打上 VLAN11 tag(通过 CDP 来告诉 PH1 打 VLAN11 tag，并不需要 PH1 和 SW1 互联的链路是 trunk link)

还可以通过 `show interfaces trunk` 来确定 SW1 g0/0 不是 trunk 口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_14-54.7j0wvlmn7q4g.webp)

## Power over Ethernet(PoE)

IP phone 也是一台设备，所以需要电源来供电。但是有一种技术 PoE(Power over Ethernet) 可以让电流通过 Ethernet cable 来传导

*PoE allows Power Sourcing Equipment(PSE) to provide power to Powered Devices(PD) over an Ethernet cable*

通常 PSE 会是 Switch

PD 可以是 IP phones, IP cameras, wireless access points 等

PoE 主要有如下几种协议

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_17-24.br46xgy2lu8.webp)

最早的 PoE 协议是 Cisco inline Power，PSE 每个端口可以提供 7 watts 的功率，网线中的 2 对绞线提供电流

使用 PoE 后拓扑如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_17-30.10ot9d9qo500.webp)

### PSE/PD

*The PSE receives AC power from the outlet, converts it to DC power, and supplies that DC power to the PDs*

> PSE 提供直流转交流的功能为 PD 供电

例如下图中

交换机电源接入到插座，就可以将 AC 转发为 DC 为 PD 供电

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_16-54.6j0fraom64u8.webp)

如果 PSE 提供太高的电流给 PC 就可能造成损毁 PD

PoE 有自己独立的一套逻辑来判断连接的设备是否需要电流，以及电流的大小

- When a device is connected to a PoE-enabled port, the PSE(switch) sends low power signals, monitors the response, and determines how much power the PD needs
- If the device needs power, the PSE supplies the power to allow the PD to boot
- The PSE continues to monitor the PD and supply the required amount of  power(but not too much)

*Power policing can be configured to prevent a PD from taking too much power*

可以通过 `power inline police` 来配置 power policing

*disable the port and send a Syslog message if a PD draws too much power by default*

`power inline police` 等价于 `power inline police action err-disable`

和 STP 中的 BPDU Guard 类似如果需要让端口正常需要先使用 `shutdown` 命令，然后使用 `no shutdown` 来启用端口

*`poewr inline police action log` does not shutdown the interface if the PD draws(吸收) too much poewr. It will restart the interface and send a Syslog message*

也可以使用 `power inline police action log` 来不让端口自动关闭当前收到过多的电流，他会重启对应的端口然后发送 Syslog message

> 因为重启端口，所以 PD 会和 PSE 重新协商需要的电流

### PoE Configuration

例如让 SW1 G0/0 为 PH1 供电

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_17-18.5atcvthl6ao0.webp)

只需要配置 `SW1(config-if)#power inline police` 指定 PSE 使用的供电策略，这里使用默认的

配置完成后可以使用 `SW1#show power inline police g0/0` 来查看供电的情况和使用的策略，主要看 Admin Police 列

如果使用了 `power inline police action log` 就会修改 Admin Police 到 log

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_17-21.7aejyxj6rksg.webp)

## QoS

在介绍 QoS(Quality of Service) 之前为什么需要了解 IP phones？ 是因为通常我们通过 QoS 来保证 IP phone 发出的报文的处理优先级，以确保 IP phone 通话的质量

Voice traffic 和 Data traffic 使用两种完全隔离的网络模式

- Voice traffic(from IP Phone) used the PSTN(Public Switched telephone network)
- Data traffic(from PC) used the IP network(enterprise WAN, Internet,etc)

在这种场景中 QoS 并不是必须的，两种不同的流量之间并不需要竞争 bandwidth

> 即使 Data traffic 导致网络拥塞，Voice traffic 也不会受影响

但是在现在的网络架构中，通常将 IP phones,服务器，个人主机都接入一个 network 中(也就不存在 PSTN 了)

这样做有一个好处就是 IP phone 可以和其他的软件(服务器)交互，但是这样也就导致了不同的流量之间会竞争 bandwidth。如果 bandwidth 够大那可以，但是如果 bandwidth 并不是特别大，如果流量之间没有优先处理的顺序，就会导致一些对时延非常敏感的服务(例如 VoIP)出问题

*QoS is a set of tools used by network devices to apply different treatment to different packets*

QoS 主要被用于控制 4 个方面

1. Bandwidth 带宽

   - The overall capacity of the link, measured in bits per second(Kbps, Mbps, Gbps, etc)

   - QoS tools allow you to reserve a certain amount of a link’s bandwidth for specific kinds of traffic

     For example: 20% voice traffic, 30% for specific kinds of data traffic, leaving 50% for all other traffic

2. Delay 时延

   - The amount of the time it takes traffic to go from source to destination = one-way delay
   - The amount of the time it takes traffic to go from source to destination and return = two-way delay

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_17-57.5lxm3toqornk.webp)

3. Jitter 抖动

   - The variation in one-way delay between packets sent by the same application

     例如发送一个请求时延是 10ms，在发送一个相同的请求时延是 100ms，中间的差值就是 Jitter

   - IP phones have a ‘jitter buffer’ to provide a fixed delay audio packet

4. Loss 丢包率

   - The % of packets sent that do not reach their destination

   - Can be caused by faulty cables(错报导致的丢包)
   - Can also be caused when a device’s packet queues get full and the device starts discarding packets(接受队列或者是窗口满了丢包)

通常针对像 IP phone call 这种流量有一种标准

- One-way delay: 150 ms or less
- Jitter: 30 ms or less
- Loss: 1% or less

如果没有达到这个标准，一般用户体验会比较差

### Qos - Queuing

*If a network device receives  messages faster than it can forward them out of the appropriate interface, the message are placed in a queue*

例如

当 Router 同时收到 G0/0-1 接口过来的流量时，只有一个 G0/2 做转发，转发效率比不上接受的效率，在 Router 内有一个抽象的 Queue，里面填充了接受但是还没有转发的报文

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_18-13.u24wb02g0bk.webp)

*By default, queued messages will be forwarded in a First In First Out(FIFO) manner*

*Messages will be sent in the order they are received*

针对队列中的报文会以先进先出的顺序对报文转发，这也是 Queue 的特性

假设现在队列满了，这时入向的报文就会被丢弃

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_18-17.2byyk6xh7sqo.webp)

这种现象也被称为 **Tail drop**，可能会导致 **TCP global synchronization**

TCP sliding window 会根据发送的流量来动态的调整

- When a packet is dropped it will be re-tranmitted
- When a drop occurs, the sender will reduce the rate it sends traffic
- It will then gradually increase the rate agian
- Then repeat all the steps

*When the queue fills up and tail drop occurs, all TCP hosts sending traffic will slow down the rate at which they send traffic*

*They will all then increase the rate at which they send traffic, which rapidly leads to more congestion*

> 正是因为 TCP 的 二进制退避 规则会导致更多的丢包，网络更加拥塞

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-12_18-25.772tlyniy0e8.webp)

解决 tail drop 和 TCP global synchronization 有一种方法就是 **Random Early Detection(RED)**

*When the amount of traffic in the queue reaches a certain threshold, the device will start **randomly dropping** packets from select TCP flows*

丢掉一整个 TCP stream(flow)

*Those TCP flows that dropped packets will reduce the rate at which traffic is sent, but you will avoid global TCP synchronization, in which ALL TCP flows reduce and the increase the rate of transmission at the same time in waves*

在标准的 RED 中，所有类型的流量都一样，都可能会被丢弃。但是在后来的版本中的 **Weighted Random Early Detection(WRED)** 可以按照流量类型或者是流量的优先级来决定是否丢弃报文

例如当 Queue 队列满了，你可以选择丢弃 HTTP 流量或者是 FTP 流量

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_16-28.2oqvvgqca5og.webp)

### 0x01

Configure SW1’s interfaces in the appropriate VLANs

```
SW1(config)#int range g1/0/2-3
SW1(config-if-range)#switchport mode access
SW1(config-if-range)#switchport access vlan 10
% Access VLAN does not exist. Creating vlan 10
SW1(config-if-range)#switchport voice vlan 20
% Voice VLAN does not exist. Creating vlan 20
```

使用 `show vlan br` 来查看

```
SW1(config-if)#do show vlan br

VLAN Name                             Status    Ports
---- -------------------------------- --------- -------------------------------
1    default                          active    Gig1/0/1, Gig1/0/4, Gig1/0/5, Gig1/0/6
                                                Gig1/0/7, Gig1/0/8, Gig1/0/9, Gig1/0/10
                                                Gig1/0/11, Gig1/0/12, Gig1/0/13, Gig1/0/14
                                                Gig1/0/15, Gig1/0/16, Gig1/0/17, Gig1/0/18
                                                Gig1/0/19, Gig1/0/20, Gig1/0/21, Gig1/0/22
                                                Gig1/0/23, Gig1/0/24, Gig1/1/1, Gig1/1/2
                                                Gig1/1/3, Gig1/1/4
10   VLAN0010                         active    Gig1/0/2, Gig1/0/3
20   VLAN0020                         active    Gig1/0/2, Gig1/0/3
```

### 0x02

Configure ROAS for the connection between SW1 and R1

这里因为 SW1 是一个三层设备可以有两种方式来配置

第一种方法路由用 R1

```
SW1(config-if-range)#int g1/0/1
SW1(config-if)#switchport mode trunk 
SW1(config-if)#switchport trunk allowed vlan 10,20

R1(config)#int f0/0
R1(config-if)#no shutdown
R1(config)#int f0/0.10
R1(config-subif)#encapsulation dot1Q 10
R1(config-subif)#ip add 192.168.10.1 255.255.255.0
R1(config)#int f0/0.20
R1(config-subif)#encapsulation dot1Q 20
R1(config-subif)#ip add 192.168.20.1 255.255.255.0
```

> 这里 packettracer 会有一个 bug, 如果使用 `show interfaces trunk` 来查看 trunk 会显示为空

方法二路由用 SW1

```
SW1(config-if-range)#int g1/0/1
SW1(config-if)#switchport mode trunk 
SW1(config-if)#switchport trunk allowed vlan 10,20

SW1(config)#ip routing 
SW1(config)#int vlan10
SW1(config-if)#ip addr 192.168.10.1 255.255.255.0
SW1(config-if)#no shutdown
SW1(config)#int vlan20
SW1(config-if)#ip addr 192.168.20.1 255.255.255.0
SW1(config-if)#no shutdown

R1(config)#int f0/0
R1(config-if)#no shutdown
...
```

### 0x03

In simulation mode, ping PC2 from PC1. Is the traffic tagged with a VLAN ID

不会带 VLAN ID，因为发送 data traffic 过的是 access mode 口

### 0x04

In simulation mode, call PH1 from PH2. Is the traffic tagged with a VLAN ID

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_17-32.4qbdbqv0t668.webp)

框中的就是 PH1 的号码，如果从 PH2 拨打 2020 就可以发现发送的报文是带 VLAN ID 的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230713/2023-07-13_17-35.7e6fnu1xaglc.webp)

这里可以发现 PH2 发出的报文 TCI 是 0x0014 即 VLAN20

**references**

1. [^https://www.youtube.com/watch?v=H6FKJMiiL6E&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=89]