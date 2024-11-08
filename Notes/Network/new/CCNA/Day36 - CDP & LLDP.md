# Day36 - CDP & LLDP



*Layer 2 discovery protocols such as CDP and LLDP share information with and discover information about neighboring(connected) devices*

*the shared information includes host name, IP address, device type, etc*

CDP 和 LLDP 都是 2 层的协议，用于共享和学习互联设备的一些信息，CDP 是思科独有的协议，而 LLDP 是 IEEE 标准的协议

因为共享的这个特性，可能存在安全隐患，所以需要根据实际的需求来判断是否使用 CDP/LLDP

例如 R1 和 SW1 互联，会分享类似下图中的内容

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_16-51.5qdipe7qnzpc.webp)

## CDP

CDP(Cisco Discovery Protocol) 是思科独有而协议，在思科所有的设备(routers, switches, firewalls, IP phones, etc)上都默认开启

> 因为默认开启，所以不用额外的配置，也能查看 CDP 信息

- CDP messages are periodically sent to multicast MAC address 0100.0ccc.cccc

- When a device receives a CDP message, it processes and discard the message. It does not forward it to other devices

  要不处理，要不丢弃，即只在 neighbor 之间有效

- By default, CDP messages are sent once every **60 seconds**. And device will add the entry in its CDP neighbor table
- By defualt, the CDP holdtime is **180 seconds**. If a message isn’t received from a neighbor for 180 seconds, the nighbor is removed from the CDP neighbor table
- CDPv2 messages are sent by default

### Show CDP Information

以下图拓扑为例

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-09.6unaxgf35ukg.webp)

可以使用 `show cdp` 来查看 CDP sending time 和 hold time，使用 `show cdp traffic` 可以看收到和发送的 CDP 报文数

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-01.6rlt7uxwsjr4.webp)

> 如果设备没有启用 CDP `show cdp` 会得到如下内容
>
> ![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-03.5gqq99kwjzi8.webp)

同样的信息也可以使用 `show cdp interface [interface-name]` 来查看，如果没有指定 interface-name 默认查看接受的端口

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-05.343nuc6d32tc.webp)

蓝框内表示当前设备上，一共有几个端口使用了 CDP，几个端口没有使用

如果需要查看 CDP neighbor table，可以使用 `show cdp neighbors`

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-08.34a28y1ynkn4.webp)

- Device ID

  当前设备的 neighbor 的 hostname

- Local Interface

  当前设备和 Device ID 对应的设备，互联的接口

  例子中表示 R1 通过自己的 G0/0 和 SW1 互联，通过自己的 G0/1 和 R2 互联

- Holdtme

  只要收到 CDP message 就就会重置为 180(holdtime 默认值)，如果该值为 0， 对应的条目就会从 CDP neighbor table 中删除

- Capability

  互联的设备是什么，可以参考输出的 code lengend 部分

  因为这里 SW1 是一个 3 层交换机，所以会显示 R S 表名既是 router 有是 switch

- Platform

  互联的设备对应的型号，因为这里是使用 GNS3 模拟的，所以不会显示

  如果在 packettracer 中会显示

  ![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-22.4h0ap45w9v5s.webp)

- Port ID

  当前设备和 Device ID 对应的设备，互联的对端接口，即 neighbor 的接口

  例子中表示 R1 连接到 R2 G0/0，连接到 SW1 G0/0

我们还可以使用 `show cdp neighbors detail` 来查看和 CDP 相关的详细信息

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-28.6kvvvorkc0ao.webp)

> CDP 可以看 VTP 相关的信息，但是 LLDP 不能，因为 VTP 同样也是思科独有的协议

==可以通过 Native VLAN 值，来判断互联的设备 Native VLAN 是否配置相同==

如果互联的设备是 Router，还可以看到对应设备的 Management address

如果需要查看单独的设备 cdp 详细的信息可以使用 `show cdp entry <device-id>`

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-35.zgeljeaal80.webp)

#### Summary

查看 CDP 的信息可以使用如下几个命令

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-36.4chqe2vcbwn4.webp)

### CDP Configuration

- CDP is globally enabled by default, CDP is also enabled on each interface by default
- to enable/disable CDP globally, use `R1(config)#[no] cdp run`
- to enable/disable CDP on specific interfaces. use `R1(config-if)#[no] cdp enable`

- Configure the CDP timer, use `R1(config)#cdp timer <seconds>`

  每次发送 CDP message 的间隔

- Configure the CDP holdtime, use `R1(config)#cdp holdtime <seconds>`

  max-age

- enable/disable CDPv2, use `R1(config)#[no] cdp advertise-v2`

## LLDP

(LLDP)Link Layer Discovery Protocol 是 IEEE 标准的协议，在思科的设备上默认不启用，必须手动配置

- **A devie can run CDP and LLDP at the same, although usually you’ll just use one**

- LLDP messages are periodically sent to multicast MAC address 0180.c200.000e

- When a device receives an LLDP message, it processes and discards the message. It does not forward it to other devices

- By default, LLDP messages are sent onece every **30 seconds**

  是 CDP 的一半

- LLDP default holdtime is **120 seconds**

  CDP 180 seconds

- LLDP has an additional timer called the ‘reinitialization delay’. If LLDP is enabled(globally or on an interface)， this timer will delay the actual initialization of LLDP. **2 seconds** by default

  为了防止状态 flapping， 所以设备不会立即发送 LLDP message 到对端

### LLDP Configuration 

以下图拓扑为例

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_17-09.6unaxgf35ukg.webp)

- LLDP is usually globally disabled by default

  这个需要看设备信号，有些设备默认启用 LLDP

  > 只要是默认的就不会显示的声明在 `show running-config` 中

- LLDP is also disabled on each interface by default

因为 LLDP 默认不会启用，所以需要手动配置

- to enable LLDP globally, use `R1(config)# lldp run`

- to enable LLDP on specific interfaces(tx), use `R1(config-if)# lldp transmit`

  只是允许发送 LLDP message

- to enable LLDP on specific a interface(rx), use `R1(config-if)# lldp receive`

  只是运行接受 LLDP message

> 对比 CDP，CDP 只需要使用一条命令就可以，而 LLDP 需要三条

- Configure the LLDP timer, use `R1(config)#lldp timer <seconds>`
- Configure the LLDP holdtime, use `R1(config)# lldp holdtime <seconds>`
- Configure the LLDP reinitialization delay timer, use `R1(config)# lldp reinit <seconds>`

### Show LLDP

和 CDP 一样，可以通过 `show lldp` 来查看对应 timer 的信息

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-03.2f1mt36380qo.webp)

使用 `show lldp traffic` 来查看发送和接受的 LLDP message 数量

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-06.1yrvd2gdlhvk.webp)

使用 `show lldp interface [interface-name]` 来查看对应接口 trimission 和 recive 的状态

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-06_1.2bo13jhqvq9s.webp)

其中的 `Tx state: IDLE` 表示当前没有发送 LLDP message，`Rx state: WAIT FOR FRAME` 表示等待接受 LLDP message

可以使用 `show lldp neighbors` 来查看 lldp neighbors

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-09.4ashl2bh5fwg.webp)

- Deivce ID

  和 CDP 中的一样，表示互联 neighbor 的 hostname

- Local Intf

  和 CDP 中的 Local Intrfc 一样，表示本端和 neighor 互联的接口

- Hold time

  ==和 CDP 不一样，只显示 hold time 对应配置的时间，如果收到 LLDP message 不会从 hold time 减==

- Capability

  对端设备的类型，具体看输出的 code lengend

  > 这里 SW1 没有值，因为是 CNS3 中虚拟的设备，如果正常会显示 B（Switch 也可以归类到 Bridge，只有 2 口）

- Port ID

  和 CDP 一样，表示当前设备和 Device ID 对应的设备，互联的对端接口，即 neighbor 的接口

除此外同样也可以使用 `show lldp neighbors detail` 来查看 lldp 详细的信息

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-18.40c99qnafx6o.webp)

这里会有两个 Capabilities 相关的字段

- `System Capabilities: B,R`

  对端互联设备的角色，这里是 3 层交换机，所以显示 B R

- `Enabled Capabilities - not advertised`

  对端互联设备开启的功能，例如 3 层交换机使用 `ip routing` 开启路由的功能，这里就会显示 R

当然也可以使用 `show lldp entry <device-id>` 来查看对应 neighbor device ID 的信息

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-23.4juzhzbcsus.webp)

#### Summary

查看 LLDP 的信息可以使用如下几个命令

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-24.2n9f3knm51hc.webp)

## Wireshark Captures

### CDP

以 R1 发送 CDP 到 SW1 为例

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_18-25.78e6m1o74pz4.webp)

- Destination: CDP/VTP/DTP/PAgP/UDLD (01:00:0c:cc:cc:cc)

  这里并没有限制只为 CDP 的多播地址，因为在 VTP/DTP 等中也使用一样的多播地址

- Cisco Discovery Protocol

  CDP 相关的详细信息

### LDP

以 SW1 发送 LLDP 到 R1 为例

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_14-37.31uj08jskwsg.webp)

- Dst: LLDP_Multicast (01:80:c2:00:00:0e)

- Link Layer Discovery Protocol

  LLDP 详细的信息

  - Time To Live = 120 sec

    holdtime

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_14-56.7cjnh1ihzy80.webp)

### 0x01

Use CDP(and other commands) to identify and label the missing IP addresses and interface IDs of the devices in the network

这里需要先查看 PC 的 IP 和 gateway 地址，然后使用 `show cdp neighbors` 来查看设备的 cdp neighbor table 来确定互联端口是那个

例如 SW1

```
SW1#show cdp neighbors 
Capability Codes: R - Router, T - Trans Bridge, B - Source Route Bridge
                  S - Switch, H - Host, I - IGMP, r - Repeater, P - Phone
Device ID    Local Intrfce   Holdtme    Capability   Platform    Port ID
R1           Gig 0/1          166            R       C2900       Gig 0/2
```

可以知道 SW1 是通过 G0/1 和 R1 G0/2 互联的，同理其他的设备

最后的拓扑如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-00.7l3qup2spegw.webp)

### 0x02

Disable CDP on the switch interfaces currently connected to PCs	

```
SW1(config)#int g0/1
SW1(config-if)#no cdp enable
R1(config)#int rang g0/0-2
R1(config-if-range)#no cdp enable
R2(config)#int range g0/0-2
R2(config-if-range)#no cdp enable
R3(config)#int range g0/0-2
R3(config-if-range)#no cdp enable
SW3(config)#int g0/1
SW3(config-if)#no cdp enable
SW2(config)#int g0/1
SW2(config-if)#no cdp enable
```

### 0x03

Disable CDP globally on each network device

```
SW1(config)#no cdp run 
SW2(config)#no cdp run 
SW3(config)#no cdp run 
R1(config)#no cdp run 
R2(config)#no cdp run 
S3(config)#no cdp run 
```

### 0x04

Enable LLDP globally on each network device, and enable Tx/Rx on the interfaces connected to other network devices.

```
SW1(config)#lldp run
SW1(config)#int g0/1
SW1(config-if)#lldp transmit 
SW1(config-if)#lldp receive 
SW2(config)#lldp run 
SW2(config)#int g0/1
SW2(config-if)#lldp transmit 
SW2(config-if)#lldp receive
SW3(config)#lldp run 
SW3(config)#int g0/1
SW3(config-if)#lldp transmit 
SW3(config-if)#lldp receive
SW1(config)#int g0/1
SW1(config-if)#lldp transmit 
SW1(config-if)#lldp receive
R1(config)#lldp run 
R1(config)#int range g0/0-2
R1(config-if-range)#lldp transmit 
R1(config-if-range)#lldp receive 
R2(config)#lldp run 
R2(config)#int range g0/0-2
R2(config-if-range)#lldp transmit 
R2(config-if-range)#lldp receive
R3(config)#lldp run 
R3(config)#int range g0/0-2
R3(config-if-range)#lldp transmit 
R3(config-if-range)#lldp receive
```

**references**

1. [^jeremy’s IT Lab]:https://www.youtube.com/watch?v=_hnMZBzXRRk&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=69