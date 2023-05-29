# Day9 - Switch Interfaces

## Full/Half Duplex

### Half duplex (半双工)

The device cannot send and receive data at the same time. If it is receiving a frame, it must wait before sending a frame

在现代的网络架构中一般不存在 Half duplex，过去常见的设备 Hub 就是一个半双工的设备

假设 PC1 需要发包到 PC2, PC3 需要发包到 PC1。因为 Hub 会 Flood 所有的报文到所有的端口，所以报文到连接 PC2 的端口时，就会有冲突，PC2 不会收到任意一个报文

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230523/2023-05-23_21-04.5zrjo8zqmjk0.webp)

这种情况所在的组网也被称为 Collision domain (冲突域)，一般通过 CSMA/CD 来解决

#### CSMA/CD

Carrier Sense Multiple Access with Collision Detection (CSMA/CD)

- Before sending frames, deviccs ‘listen’ to the collision domain until they detect that other devices are not sending
- If a collision does occur, the device sends a jamming(拥塞) signal to inform the other devices that a collsion happened
- Each device will wait a random period of time before sending frames again
- The process repeats 

### Full duplex (全双工)

The device can send and receive data at the same time. It does not have to wait

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

  Frames that have an incorrect format (due to an error)

- 0 input errors

  Frames the switch recived, but failed due to an error

  包含 runts/giants/CRC/frame 的计数

- 0 output errors

  Frames the switch tried to send, but failed due to an error

**references**

[^jeremy’s It Lab]:https://www.youtube.com/watch?v=cCqluocfQe0&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=16
