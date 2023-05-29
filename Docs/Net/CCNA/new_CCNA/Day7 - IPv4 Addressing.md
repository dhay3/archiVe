# Day7 - IPv4 Addressing



## Routing

Router 是一个 3 层设备，用于划分广播域(LAN)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230519/2023-05-22_14-44.2d0o6ilyhcsg.webp)

假设现在 PC1 访问 PC3

## IPv4

### IPv4 Header

IPv4 Header 长这样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_14-46.4kh1u49auozk.webp)

先关注两个字段

- Source IP address
- Destination IP address

两个字段都为 32bits(4bytes)

### IPv4 address

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_14-50.bijsuyjuee0.webp)

一个 IPv4 address 由 4 部分 8 bits decimal number 组成，通过 dot 分隔，所以也被称为 dotted decimal (中文叫点分十进制)。主要为了方便记忆

> 每 8 bits 一组的部分也被称为 octets

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_15-00.3y75edo8q1w.webp)

IPv4 address 还可以通过 `/` 的方式来划分 network portion 和 host portion

即例子中 `192.168.1` 是 network portion，`254` 是 host portion

#### IPv4 Address Classes

根据 first octet 可以将 IPv4 address 分为 4 类

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_14-58.gzyxlswhwn4.webp)

主要关注 ABC 类地址，D 类地址为保留的 Multicast address, E 类地址为保留的实验地址

- A 类地址，第一组 octet，以 0 开始，network portion 8 位

  在 A 类地址实际范围只有 `0-126`，firtst otect 127 为保留的 Loopback Adress

  - Address range `127.0.0.0` - `127.255.255.255`
  - Used to test the ‘network stack’(think OSI,TCP/IP model) on the local device

- B 类地址，第一组 octet，以 01 开始，network portion 16 位

- C 类地址，第一组 octet，以 110 开始，network portion 24 位

例如

`12.128.251.23/8` 是 class A address

`154.78.111.32/16` 是 class B address

`192.168.1.254/24` 是 class C address



每类地址都有固定可用数量的 network 和 addresses(host) per network 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_15-14.5ckh1ct40adc.webp)

> number of networks 或者 addresses per network 计算的都是可能性(和 binary 转 decimal 没有什么关系)。每 bit 上只有 0 和 1 两种可能，想象成一棵二叉树
>
> 以 class A 为例子
>
> 实际可以分配数字的 network portion 有 7 bit，所以 number of networks 部分就有 $2^7$ 个
>
> 实际可以分为数字的 host portion 有 24 bit，所以 number of networks 部分就有 $2^{24}$ 个 (实际还需要 minus 2，host portion 全 0 表示 network address 和 host portion 全 1 表示 boardcast address)

##### Network Adddress

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_15-34.543fjla8lngg.webp)

- Host portion 全 0 的地址
- 不能被分配给主机

例如上图中的 `192.168.1.0/24` 和 `192.168.2.0/24` 都是 network addresss

##### Boardcast Address

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_15-37.4fshawu1xu9s.webp)

- Host portion 全 1 的地址
- 不能被分配给主机

例如上图中的 `192.168.1.0/24` 对应的 boardcast address 为 `192.168.1.255`

### Netmask

Netmask(Subnet mask) 中文也叫做掩码(子网掩码)，和 IPv4 address 一起使用，标识 IPv4 address 所处的 network address

例如

- Class A /8 对应 netmask `255.0.0.0`
- Class B /16 对应 netmask `255.255.0.0`
- Class C /24 对应 netmask `255.255.255.0`



## Configure IP address

按照如下拓扑配置 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230522/2023-05-22_16-09.7jrw4i8aeef4.webp)

进入 global configuration mode

```
R1>en
R1#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
```

配置 fa0/0

```
R1(config)#int fa0/0           
R1(config-if)#ip address ?
  A.B.C.D  IP address
  dhcp     IP Address negotiated via DHCP
  pool     IP Address autoconfigured from a local DHCP pool

R1(config-if)#ip address 10.0.0.254 ?
  A.B.C.D  IP subnet mask

R1(config-if)#ip address 10.0.0.254 255.0.0.0
```

查看端口状态

```
R1(config-if)#do show ip int br
Interface                  IP-Address      OK? Method Status                Protocol
FastEthernet0/0            10.0.0.254      YES manual administratively down down      
FastEthernet0/1            unassigned      YES unset  administratively down down    
FastEthernet0/2            unassigned      YES unset  administratively down down
```

- Interface

  对应的端口名

- IP-Address

  端口对应的 IP address

- OK?

  端口对应 IP address 是否是 valid

- Method

  IP address 是通过什么方式分配的

- Status

  Layer 1 status of the interface

  - administratively down 表示使用了 `shutdown` 命令，由管理员手动关闭端口的。Cisco Router 端口默认该状态，Cisco Switch 端口默认不是 administratively down
  - up 表示端口，使用了 `no shutdown` 命令，且和对端的设备正确连接
  - down 表示端口，使用了 `no shutdownn` 命令，但是和对端的设备没有正确连接

- Protocol

  Layer 2 status of the interface

  如果 Status 是 down，那么 Protocol 字段也一定是 down

  相反如果 Protocol 字段为 down，但是 Status 字段不一定是 down

使用 `no shutdown` 打开端口

```
R1(config-if)#no shutdown 
```

在使用 `no shutdown` 后，可以看到如下输出

第一行表示 Layer1 状态变成 UP，对应 Status

第二行表示 Layer2 状态变成 UP，对应 Protocol

```
R1(config-if)#
*Mar  1 00:20:53.635: %LINK-3-UPDOWN: Interface FastEthernet0/0, changed state to up
*Mar  1 00:20:54.635: %LINEPROTO-5-UPDOWN: Line protocol on Interface FastEthernet0/0, changed state to up
```

配置 fa0/1

```
R1(config)#int fa0/1
R1(config-if)#ip add 172.16.0.254 255.255.0.0
R1(config-if)#no shutdown 
```

配置 fa0/2

```
R1(config)#int fa0/2
R1(config-if)#ip add 192.168.0.254 255.255.255.0
R1(config-if)#no shutdown 
```

查看端口信息

```
R1(config-if)#do show ip int br
Interface                  IP-Address      OK? Method Status                Protocol
FastEthernet0/0            10.0.0.254      YES manual up                    up      
FastEthernet0/1            172.16.0.254    YES manual up                    up      
FastEthernet0/2            192.168.0.254   YES manual up                    up
```

除 `show ip interface brief` 外还可以使用 `show interface <interface-name>` 来查看端口的详细信息

```
R1#show int fa0/0
FastEthernet0/0 is up, line protocol is up 
  Hardware is AmdFE, address is cc01.057e.0000 (bia cc01.057e.0000)
  Internet address is 10.0.0.254/8
  MTU 1500 bytes, BW 100000 Kbit/sec, DLY 100 usec, 
     reliability 255/255, txload 1/255, rxload 1/255
  Encapsulation ARPA, loopback not set
  Keepalive set (10 sec)
  Full-duplex, 100Mb/s, 100BaseTX/FX
  ARP type: ARPA, ARP Timeout 04:00:00
  Last input 00:25:14, output 00:00:02, output hang never
  Last clearing of "show interface" counters never
  Input queue: 0/75/0/0 (size/max/drops/flushes); Total output drops: 0
  Queueing strategy: fifo
  Output queue: 0/40 (size/max)
  5 minute input rate 0 bits/sec, 0 packets/sec
  5 minute output rate 0 bits/sec, 0 packets/sec
     1 packets input, 70 bytes
     Received 1 broadcasts, 0 runts, 0 giants, 0 throttles
     0 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
     0 watchdog
     0 input packets with dribble condition detected
     264 packets output, 26917 bytes, 0 underruns
     0 output errors, 0 collisions, 1 interface resets
     0 unknown protocol drops
     0 babbles, 0 late collision, 0 deferred
     0 lost carrier, 0 no carrier
     0 output buffer failures, 0 output buffers swapped out
```

- FastEthernet0/0 is up

  表示 Layer1 是正常的

- line protocol is up

  表示 Layer2 是正常的

  - address is cc01.057e.0000 (bia cc01.057e.0000)

    表示 MAC 地址，bia 也就是 bured in address 和 MAC 地址含义一样

- Internet address is 10.0.0.254/8

  表示 IP 地址

除此之外还有一个 `show interfaces description` 命令来查看用户自己为端口设定的 description

```
R1#show int description 
Interface                      Status         Protocol Description
Fa0/0                          up             up       
Fa0/1                          up             up       
Fa0/2                          up             up
```

可以通过 `description <description information>` 来设定

```
R1(config-if)#description //this is a description of fa0/0
R1(config-if)#do show int desc
Interface                      Status         Protocol Description
Fa0/0                          up             up       //this is a description of fa0/0
Fa0/1                          up             up       
Fa0/2                          up             up
```

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=3ROdsfEUuhs&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=13