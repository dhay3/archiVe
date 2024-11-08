# Day51 - Dynamic ARP Inspection

## ARP

回忆一下普通的 ARP，例如 192.168.1.10 需要访问 8.8.8.8

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_20-24.6ysrshz8ilmo.png)

因为 PC1 目前还没有 8.8.8.8 的 MAC，且 8.8.8.8 不在一个 LAN 中，需要通过 Router 转发，所以会先问 Router 的 MAC 地址。即 PC1 首先会广播一个 ARP Request, 报文内容大致如下

> ARP 只用于 LAN，不需要被发送到 LAN 外，所以也就不需要被 router 发送到其他的 network，所以也就不需要 IP header 

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_20-30.53ewwyghmlo0.webp)

ARP Request 到 SW1/SW2 分别会记录相应的条目到对应设备的 MAC table，R1 在收到 ARP Request 后因为 Target IP address 对应接口的地址，R1 会将相应的条目记录到 MAC table 和 ARP table 中，同时回送 ARP Reply 到 PC1

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_20-40.79yi6n847680.webp)

ARP Reply 到 SW1/SW2 分别记录相应的条目到对应设备的 MAC table，PC1 在收到 ARP reply 将相应的条目记录到自己的 MAC table 和 ARP table 中，同时发送报文

Src IP 192.168.1.10 Dst IP 8.8.8.8

Src MAC 192.168.1.10’s mac Dst MAC 192.168.1.1’s mac

## Gratuitous ARP

*A Gratuitous ARP message is an **ARP reply** that is sent without receiving an ARP request.*

Gratuitous ARP(GARP) 顾名思义免费的 ARP，有几点特质

1. It is send to the broadcast MAC address

   相比普通的 ARP reply，是 unicast 的

2. It allows other devices to learn the MAC address of the sending device without having to send ARP requests
3. Some devices automatically send GARP messages when an interface is enabled, IP address is changed, MAC address is changed, etc

例如当 PC1 开机时，就会发送 GARP 到整个网络

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_20-51.30xxpl9biw20.webp)

## ARP Poisoning

ARP poisoning 和 DHCP poisoning 类似

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-04.9meva8arfbw.webp)

1. 假设 PC2 发送了 GARP(当然也可以是正常的 ARP reply) 声明 192.168.1.1 的 MAC 是 PC2，而实际 PC2 的 IP 是 192.168.1.11
2. LAN 中所有的设备除 192.168.1.1 外所有的设备会认为 PC2 就是 192.168.1.1
3. 如果这时 PC1 访问 192.168.1.1 就会将报文发送到 PC2

需要防止这种情况的出现，可以使用 DAI

## Dynamic ARP Inspection

*DAI is a security feature of switches that is used to filter ARP messages received on untrusted ports*

Dynamic ARP Inspection(DAI) 是交换机用于过滤端口接受到 ARP messages 的一个安全功能

当设备启用 DAI 时，收到 ARP messages 会

> DAI 和 DHCP snooping 一样只会校验 untrusted ports 收到的 ARP 报文，并不会校验 trusted ports 收到的 ARP 报文

1. DAI inspects the **sender MAC** and **sender IP** fields of ARP messages received on untrusted ports and checks that there is a matching entry in the **DHCP snoop binding table** 
2. If there is a matching entry, the ARP message is forwarded normally
3. If there isn’t a matching entry, the ARP message is discard

DAI 有一些注意点

- DAI only filters ARP messsages. Non-ARP messages aren’t affected

- All ports are untrusted by default

  和 DHCP 一样， 默认所有的端口都是 untrusted 的，但是也有一些逻辑规则来分配端口是否是 trusted

  Typically, all ports connected other network devices(switches, routers) should be configured as trusted, while interfaces connected to end hosts should remain untrusted

  例如下图中蓝色的接口都是 trusted ports，橙色的接口都是 untrusted ports

  ![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-01.77zigay9m400.webp)

- ARP ACLs can be manually configured to map IP addresses/MAC address for DAI to check

  主要针对不使用 DHCP 的机器，如果机器不是使用 DHCP 也就无法通过 DHCP snooping binding table 来判断设备的 ARP request 或者 repy 是否有效 

- Like DHCP snooping, DAI also supports rate-limiting to prevent attackers from overwhelming the switch with messages

  但是需要注意的一点是，虽然可以限制 overwhelming, 但是设备 CPU 仍要处理报文，大量的报文是会导致 CPU 过载的，从而影响转发

## DAI Configuration

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-25.1dxkpvjool8g.webp)

DAI 配置很简单

1. `SW1(config)# ip arp inspection vlan 1`

   表示设备针对具体 VLAN 开启 DAI

2. `SW1(config-if)#ip arp inspection trust`

   表示当前端口为 trusted port

配置完成后可以使用 `SW1#show ip arp inspection interfaces` 来查看那些端口是 trusted 或 untrusted 的

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-31.159xqa9rpakg.webp)

> DAI 比 DHCP snooping 更灵活

### DAI rate limiting

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-34.5y8en3uceq00.webp)

配置 rate limiting 也相对简单，只需要使用 `SW1(config-if)#ip arp inspection limit rate <counts> [burst interval <seconds>]` 即可

例如 `ip arp inspection limit rate 25 burst interval 2` 表示每 2 秒可以接受的 ARP messages 为 25，超过 25 端口变为 errdisable；`ip arp inspection limit rate 10` 表示每秒可以接受的 ARP messages 为 10，超过 10 端口变为 errdisable

这里和 portsecurity/DHCP snooping 导致端口 errdisable 一样，也可以设置自动将状态改为 enable

只需要使用 `SW1(config)#errdisable recovery cause arp-inspection` 即可

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-40.6bgvgqybuvo0.webp)

### DAI Optional Checks

DAI 除了校验 DHCP snooping table 还可以校验 dst-mac, ip, src-mac

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_21-52.2t16fco5zz40.webp)

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_20-30.53ewwyghmlo0.webp)

- dst-mac

  即对比 ARP 报文中 Target MAC 和 Destination MAC, 如果不匹配就丢弃

- ip

  即对比 ARP 报文 Sender IP 是否是 0.0.0.0/255.255.255.255/multicast address，如果是就丢弃

- src-mac

  即对比 ARP 报文中 sender MAC 和 Source MAC, 如果不匹配就丢弃

如果要同时校验 dst-mac, ip, src-mac 需要使用 `SW1(config)#ip arp inspection validate dst-mac ip src-mac`

> dst-mac ip src-mac 没有先后顺序

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_22-00.3doab41glzc0.webp)

### ARP ACLs

如果设备使用静态地址不使用 DHCP，因为对应的信息不在 DHCP snooping table 中就会直接被丢弃

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_22-01.6kiz90bgv94.webp)

为了让这些设备的 ARP request 或者 reply 正常转发可以通过 ARP ACL

![](https://github.com/dhay3/image-repo/raw/master/20230815/2023-08-15_22-27.6x3x5x6bnis0.webp)

- `SW1(config)#arp access-list <name>`

  创建一个 ARP ACL

- `SW1(config-arp-nacl)#permit ip host <source-ip> mac host <source-mac>`

  放通指定 source IP MAC 的 ARP message

- `SW1(config)#ip arp inspection filter <name> vlan <id>`

  将 ARP ACL 应用到指定的 VLAN

配置完成后可以使用 `show ip arp inspection` 来查看 DAI 配置

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_13-41.5fo8a28k8vc0.webp)

## Command Summary

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_13-42.1ror1amu9gbk.webp)

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230816/2023-08-16_13-56.4uwapfz5umc0.webp)

### 0x01

Configure R1 as a DHCP server

Exclude 192.168.1.1 - 192.168.1.9 from the pool

Default gateway：R1

```
R1(config)#ip dhcp excluded-address 192.168.1.1 192.168.1.9
R1(config)#ip dhcp pool 192.168.1.0/24
R1(dhcp-config)#network 192.168.1.0 255.255.255.0
R1(dhcp-config)#default-router 192.168.1.1
```

将 PC1 配置成 DHCP，然后使用 `show ip dhcp binding` 来校验 DHCP 是否成功

```
R1(dhcp-config)#do show ip dhcp binding
IP address       Client-ID/              Lease expiration        Type
                 Hardware address
192.168.1.10     0001.6432.B922           --                     Automatic
```

### 0x02

Configure DHCP snooping on SW1 and SW2

```
SW1(config)#ip dhcp snooping 
SW1(config)#ip dhcp snooping vlan 1
SW1(config)#no ip dhcp snooping information option 
SW1(config-if)#int g0/2
SW1(config-if)#ip dhcp snooping trust
SW2(config)#ip dhcp snooping 
SW2(config)#ip dhcp snooping vlan 1
SW2(config)#no ip dhcp snooping information option 
SW2(config-if)#int g0/1
SW2(config-if)#ip dhcp snooping trust
```

### 0x03

Configure DAI on SW1 and SW2

Enable all additional validation checks

Trust ports connected to a router or switch

```
SW1(config)#ip arp inspection validate dst-mac ip src-mac 
SW1(config)#ip arp inspection vlan 1
SW1(config)#int range g0/1 - 2
SW1(config-if-range)#ip arp inspection trust 
SW2(config)#ip arp inspection validate dst-mac ip src-mac 
SW1(config)#ip arp inspection vlan 1
SW2(config)#int g0/1
SW2(config-if)#ip arp inspection trust 
```

配置完成后可以使用 `show ip arp inspection interfaces` 和 `show ip arp inspection` 来查看配置

**references**

1. ^https://www.youtube.com/watch?v=HwbTKaIvL6s&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=99