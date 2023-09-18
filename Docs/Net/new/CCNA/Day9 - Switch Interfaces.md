# Day9 - Switch Interfaces



## Duplex

Duplex 双工分两种

- Half duplex
- Full duplex

### Half duplex (半双工)

通信的两个或者多个参与者之间，数据传输只能在一个方向上进行，而不能同时进行双向传输。当其中的一个参与者发送数据时，其他的参与者必须等待接收完成后才能发送自己的数据

> 例如 walkie-talkies(对讲机)，一个人在说话时，其他人只能接收

在现代的网络拓扑中一般不会使用 Half duplex 设备，但是早期使用的 **Hub(会将收到的报文，flood 到除接收端口外的所有端口，不管报文是几层的)** 就是一个工作在半双工下的网络设备

假设现在 PC1/PC2/PC3 通过 Hub 互联

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-02_14-49.1ekt9lmqjudc.webp)

当 PC1 需要发送报文到 PC2

1. 因为 PC1 发送报文，所以 PC2 和 PC3 只能处于接受状态，不能发送报文
2. 报文到了 Hub 会 flood 到除接收端口外的所有端口，即和 PC2/PC3 互联的端口 
3. 因为 PC3 MAC 不匹配，所以会丢弃报文，PC2 MAC 匹配所以回报文
4. 因为 PC2 回报文，所以 PC1 和 PC3 只能处于接受状态，不能发送报文
5. 回报文到了 Hub 会 flood 到除接收端口外的所有端口，即和 PC1/PC3 互联的端口
6. 因为 PC3 MAC 不匹配，所以会丢弃报文，PC1 MAC 匹配，所以报文到了 PC1

> Half duplex 并不是指对应的设备(例如 Hub)在同一时间内只能接受或者是发送报文，Half duplex 指的是一种通信模型
>
> ==即同一时间 Half duplex 网络上只可以有一个发送者，发送一个方向的传输数据==

### Full duplex (全双工)

通信的两个或者多个参与者之间，数据传输可以在两个方向上进行，发送方和接收方可以同时发送和接收数据(无需监听网络是否空闲)，而不需要等待对方的响应

以上图拓扑为例(将 Hub 换成 Switch)，PC1 需要发送报文到 PC2

PC1 发送报文时，PC2 和 PC3 也可以发送报文，无须处于接收状态等待 PC1 的报文

## Collision Domain

> 在 Half-Duplex 和 Full-Duplex 中都会出现 Collision Domain

*The term **collision domain** is used to describe a part of a network where packet collisions can  occur. Packet collisions occur when two devices on a shared network segment **send** packets simultaneously. The colliding packets must be  discarded and sent again, which reduces network efficency.*

在 Half Duplex 的情况下，网络上只可以有一个发送者发送一个方向的传输数据。那么就需要互联的设备监听网络是否“空闲”，如果网络“空闲”就可以发送数据包。但是这样就会有一个问题，那就是多台设备同时判断网络为“空闲”，这样就不能保证只有一个方向或者是只有一个发送者。他们的信号会相互干扰，导致信息的损坏或者是丢失

例如在同一时间 PC1 需要发包到 PC2, PC3 需要发包到 PC1，两个报文同时到了 Hub

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230523/2023-05-23_21-04.5zrjo8zqmjk0.webp)

因为 Hub 会 Flood 所有的报文到除接收端口外的所有的端口，所以就会将从 PC1 和 PC3 收到的报文同时 flood 到和 PC1 互联的端口

Hub 不会将报文排序发送，而会同时发送到 PC1，这样就会有冲突，PC2 并不会完好无损的收到 PC1 和 PC3 的报文

所以 PC1 和 PC3 会重新发送报文

这种情况所在的组网也被称为 Collision domain[^Collision domain] (冲突域)，一般通过 CSMA/CD 来解决

> 需要注意的一点是
>
> Hub 所在的整个组网都是一个 Collision Domain
>
> 而 Switch 每一个端口所组成的 link 就是一个 Collision Domain

#### CSMA/CD

Carrier Sense Multiple Access with Collision Detection (CSMA/CD)

- Before sending frames, deviccs ‘listen’ to the collision domain until they detect that other devices are not sending
- If a collision does occur, the device sends a jamming(拥塞) signal to inform the other devices that a collsion happened
- Each device will wait a random period of time before sending frames again
- The process repeats 

## Speed/Duplex autonegotiation

> Interfaces that can run at different speeds (10/100 or 10/100/1000) have default settings of speed auto and duplex auto
>
> Interfaces ‘advertise’ their capabilities to the neighboring device, and they negotiate the best speed and duplex settings they are both capcable of
>
> 即如果使用 auto negotiation，speed 默认是两端互联的端口最小的；duplex 默认使用端口支持的，如果两端互联的端口都支持 full duplex，会使用 full duplex，如果其中有一端只支持 half duplex，会使用 half duplex
>
> - Speed
>
>   the switch will try to sense the speed that the other device is operating at
>
>   If it fails to sense the speed, it will use the slowest supported speed (ie 10Mbps on 10/100/1000 interface)
>
> - Duplex
>
>   If the speed is 10 or 100 Mbps, the switch will use the half duplex
>
>   If the speed is 1000 Mbps or greater, use full duplex

如果有以下拓扑，并使用了 auto negotiation

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230523/2023-05-23_21-24.2k2bc7agpkao.webp)

那么最后每个端口的 speed/duplex 如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230523/2023-05-23_21-26.6qd3x68etj0g.webp)

如果 auto negotiation 被 disable 掉了，例如下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230523/2023-05-23_22-11.5v3v1izyl8qo.webp)

那么最后每个端口的 speed/duplex 如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230523/2023-05-23_22-13.5y2w4w27km80.webp)

## Commands

### Query duplex and speed

`show ip int br` 和 Router 一样，都用于查看的 interface 的信息。但是和 Router 不一样，Switch 端口默认打开，所以默认状态为 down 而不是 administratively down

```
SW(config)#do show ip int br
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet0/0     unassigned      YES unset  down                  down
```

> Status/Protocol，例如 up/up 或者 down/down

- Router interfaces have the `shutdown` command applied by default.

  Will be in the administratively down/down state by default

- Switch interfaces do NOT have the `shudown command` applied by default.

  Will be in the up/up state if connected to another device OR in the down/down state if not connected to another device

`show interfaces status` 也用于查看端口的状态，例如 duplex(双工)，speed(速率)，Type(类型)

```
SW#show interfaces status

Port      Name               Status       Vlan       Duplex  Speed Type 
Fa0/1                        connected    1          a-full  a-100 10/100BaseTX
Port      Name               Status       Vlan       Duplex  Speed Type 
Fa0/2                        connected    truck      a-full  a-100 10/100BaseTX
Port      Name               Status       Vlan       Duplex  Speed Type 
Fa0/3                        notconnected 1          auto  	 auto  10/1000BaseTX
Port      Name               Status       Vlan       Duplex  Speed Type 
Fa0/4                        notconnected 1          auto  	 auto  10/1000BaseTX
```

Name 表示 description，使用 `description <description>` 来设定

Duplex column 中的 `auto` (只有在 status 为 notconnected 时)表示端口和对端端口自动协商，如果可以使用 full-duplex 会优先采用。`a-full` (只有在 status 为 notconnected 时)表示端口和对端端口自动协商，并采用 full-duplex

Speed column 中的 `auto` (只有在 status 为 notconnected 时)表示端口和对端端口自动协商，会使用最大的速率。`a-1000` (只有在 status 为 notconnected 时)表示端口和对端端口自动协商，速率为 100 Mbit/s

Type column 表示端口的类型，可以看出来是 UTP 100Base 的，最大速率就是 100 Mbits/s

### Configuring interfaces speed and duplex

> 如果 interface 使用了 `negotiation auto`，那么就需要先使用 `no negotiation auto ` 取消自动协商，然后才可以修改 speed 和 duplex

如果需要修改 speed 和 duplex 可以通过如下方式

```
SW>en
SW#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
SW(config)#int g0/0
SW(config-if)#no negotiation auto
SW(config-if)#speed 10
SW(config-if)#duplex half
SW()
SW#write
Building configuration...
Compressed configuration from 4730 bytes to 1780 bytes[OK]
*May 23 10:35:22.984: %GRUB-5-CONFIG_WRITING: GRUB configuration is being updated on disk. Please wait...
*May 23 10:35:23.613: %GRUB-5-CONFIG_WRITTEN: GRUB configuration was written to disk successfully
```

### Interfaces Errors

当使用 `show interfaces <interfaces-name>` 时会输出端口相关的 error counter

```
SW#show interfaces gi0/0 
GigabitEthernet0/0 is down, line protocol is down (notconnect) 
  Hardware is iGbE, address is 0c8b.7be3.0000 (bia 0c8b.7be3.0000)
  Description: gi0/0 port
....
     269 packets input, 71059 bytes, 0 no buffer
     Received 0 broadcasts (0 multicasts)
     0 runts, 0 giants, 0 throttles 
     0 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
     0 watchdog, 0 multicast, 0 pause input
     7290 packets output, 429075 bytes, 0 underruns
     0 output errors, 0 collisions, 2 interface resets
     0 unknown protocol drops
     0 babbles, 0 late collision, 0 deferred
     1 lost carrier, 0 no carrier, 0 pause output
     0 output buffer failures, 0 output buffers swapped out
```

- 269 packets input, 71059 bytes

  端口一共发送的 packets 个数，以 bytes 计算大小

- 0 runts

  Frames that are smaller that the minimum frame size (64bytes)

- 0 giants

  Frames that are larger than the maximum frame size (1518bytes)

  1518 = 1500 mtu + 14 L2 header + 4 L2 trailer(CRC)

- 0 CRC

  Frames that failed the CRC check(in the Ethernet FCS trailer)

- 0 frame

  Frames that have an incorrect format (due to an error，不符合 IEEE 标准的畸形报文)

- 0 input errors

  Frames the switch recived, but failed due to an error

  包含 runts/giants/CRC/frame 的计数

- 0 output errors

  Frames the switch tried to send, but failed due to an error

**references**

[^jeremy’s It Lab]:https://www.youtube.com/watch?v=cCqluocfQe0&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=16
[^Collision domain]:https://www.uokirkuk.edu.iq/science/images/2019/Lectures_download/Dr.Ahmed_chalack/lecture_8.pdf
