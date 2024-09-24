# Day32 - IPv6 Part2

## Configuring IPv6 addresses(EUI-64)

除了手动配置 IPv6 地址外，还可以通过一种规则来配置 IPv6

EUI(Extended Unique Identifier) 也被称为 Modified EUI-64

*(Modified) EUI-64 is a method of converting a MAC address(48 bits) into a 64-bit interface identifier*

> EUI-64 可以将 MAC 地址转为 64 bit 的 interface identifier

*This interface identifer can then become the ‘host portion’ of a /64 IPv6 address*

> 意味着 64 bits host portion，64 bits network portion

EUI-64 是怎么将 MAC address 转成 host portion 的呢？

1.  Divide the MAC address in half

   例如 1234 5678 90AB 会被分成两部分

   1234 56 和 78 90AB

2. Insert FFFE in the middle

   上面就会变成 1234 56FF FE78 90AB

3. Invert(取反) the 7th bit

   上面就会变成 1034 56FF FE78 90AB

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-00.6j9zlduqxmdc.webp)

下面几个例子加深一些印象

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-04.2p6fxpzsz2ww.webp)

如果需要使用 EUI-64 来配置 IPv6 地址，可以参考下面的例子

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-07.4gbdk263qvq.webp)

需要使用 `ipv6 address <ipv6 network portion> eui-64` 来使用 EUI-64 配置端口 IPv6 地址

> 这里必须带上 IPv6 network portion,因为 EUI-64 转化而来的部分仅仅只作为 IPv6 host portion

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-10.5fmgh1bfyy9s.webp)

上图为接口对应的 MAC address

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-13.32v23apcz8qo.webp)

例如 

g0/0 network portion 2001:db8:0:0::/64, host portion 可以计算出为 0ef8:22ff:fe36:8500

组合为 2001:db8:0:0:0ef8:22ff:fe36:8500 缩写 2001:db8::ef8：22ff:fe36:8500

### Why invert the 7th bit?

在 EUI-64 中 7th bit 为什么需要取反呢？

MAC addresses 可以分为两种

1. UAA(Universally Administered Address)

   Uniquely assigned to the device by the manufacturer

2. LAA(Locally Administered Address)

   Manually assigned by an admin(with the `mac-address` command on the interface) or protocol. Doesn’t have to be globally unique

我们可以通过 7th bit of the MAC address(也被称为 U/L bit Universal/Local bit) 来识别 MAC address 是 UAA 还是 LAA

1. if U/L bit set to 0 => UAA
2. if U/L bit set to 1 => LAA

在 IPv6 中正好相反

1. if U/L bit set to 0 = The MAC address the EUI-64 interface ID was made from was an LAA
2. if U/L bit set to 1 = The MAC address the EUI-64 interface ID was made from was an UAA

> 这个并不表示 IPv6 和 MAC address 类似也可以分为类似 UAA 或者是 LAA 

## Global unicast addresses

*Global unicast IPv6 addresses are public addresses which can used over the Internet*

> 可以将 Global unicast IPv6 addresses 直接理解成 IPv6 公网，全球唯一的地址

最初 2000::/3(2000:: to 3FFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF) 被定义为 Global unicast addresses 范围，后来规定只要不是保留的地址都是 Global unicast addresses

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-45.z0ng1cby4uo.webp)

## Unique Local addresses

*Unique local IPv6 addresses are private addresses which cannot be used over the Internet*

> 可以将 Unique Local addresses 直接理解成 IPv6 私网，不需要全球唯一

最初 FC00::/7 (FC00:: to FDFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF) 被定义为 Unique Local addresses 范围，后来规则要求 8th bit 必须为 1，所以必须以 FD 开头，即 FD00::/7

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-46.47j7ubndwzcw.webp)

## Link local addresses

Link local addresses 会在使用 IPv6 地址的接口上自动生成，范围在 FE80::/10(FE80:: to FEBF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF)，但是后来的的规则要求只能以 FE80/10 开头，10 bits - 64 bits 部分都需要是 0， 64 bits 部分拼接 EUI-64 转换的 host portion 部分

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_15-56.6wewt38z37nk.webp)

*Link-local means that these addresses are used for communication within a single link(subnet). Routers will not route packets with a link-local destination IPv6 address*

> link-local address 只能在同 subnet 中使用

link-local address 通常被用在如下几个场景

1. Routing protocol peerings(OSPFv3 uses link-local addresses for neighbor adjacencies)
2. Next-hop addresses for static routes
3. Neighbor Discovery Protocol(NDP,IPv6’s replacement for ARP)uses link-local addresses to function

例如

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_16-06.12is7m3ob40w.webp)

PC1 需要访问 PC2，R1/2/3/4 上均只有 link-local address(可以使用 `ipv6 enable` 来自动生成 link-local address，而无需通过手动配置端口 IPv6 地址来生成 link-local address)

这时 PC1 的报文是可以通过 R1/2/3/4 正常到 PC2 的，但是在 R1 上并不能 ping 通 

R3/R4 任意一个 link-local address，但是可以 ping 通 R2 和 R1 互联的接口地址，因为 link-local address 只能用在同 subnet 中

## Multicast addresses

- Unicast addresses are one-to-one

  One source to one destination

- Broadcast addresses are one-to-all

  One source to all destinations(within the subnet)

  IPv6 不支持 broadcast，所以没有 broadcast address

- Multicast addresses are one-to-many

  One source to multiple destinations(that have joined the specific multicast group)

  IPv6 使用 FF00::/8 for multicast

虽然 IPv6 没有 broadcast address 不支持 broadcast，但是可以通过 mutlicast address 中的 FF02::1 来实现 “broadcast”

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_16-43.2y63dnja9uv4.webp)

> 这里 IPv6 address 都是 Link-local scope multicast address
>
> 不要将 Link-local multicast address 和 Link-local address 混淆
>
> 在 IPv6 中 multicast address 均以 FF 开头

### Mutlicast address scopes

IPv6 定义了多种 multicast scopes

- Interface-local(FF01)

  the packet doesn’t leave the local device. Can be used to send traffic to a service within the local device

- Link-local(FF02)

  the packet remains in the local subnet. Routers will not route the packet between subnets

- Site-local(FF05)

  the packet can be forwarded by routers. Should be limited to a single pyhsical location(not forwarded over a WAN)

- Organization-local(FF08)

  Wider in scope than site-local(an entire company/organization)

- Global(FF0E)

  No boundaries. Possible to be routed over the Internet

例如下图

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_17-17.w5ybhhpusao.webp)

如果使用 `show ipv6 int <interface-name>` 可以看到加入的 multicast group

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_17-19.40fsvo5w70ao.webp)

## Anycast

Anycast 是 IPv6 中的一种新特质

one-to-(one-of-many)

有多个可能的目的地址，但是每次只和一台目的地址交互

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_17-25.7hv9gts9i4n4.webp)

Multiple routers are configured with the same IPv6 address

- They use a routing protocol to advertise the address
- When hosts sends packets to that destination address, routers will forward it to the nearest router configured with that IP address(based on routing metric)

> anycast address 并没有一个特定的地址范围

可以通过 `ipv6 address <ipv6-address> anycast` 来使一个地址变为 anycast address

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_17-30.1fuax58yk4g0.webp)

如果使用了上述命令，可以通过 `show ipv6 interface <interface-name>` 就可以看到 unicast address 多路一个 2001:DB8:1:1::99 的地址对应配置的的 anycast 地址，使用 ANY 来标识

## Other IPv6 Addresses

- ::

  The unspecified IPv6 address

  > 可以理解成 IPv4 中的 0.0.0.0 表示 wildcard address

  Can be used when a device doesn’t yet know its IPv6 address.

  IPv6 default routes are configured to ::/0

- ::1

  The loopback address

  > 等价于 IPv4 127.0.0.0/8
  >
  > 在 IPv6 中只使用 ::1 这一个地址作为 loopback address，而 IPv4 中是 127.0.0.0/8 这一个段

  Used to test the protocol stack on the local device

  Messages sent to this address are processed within the local device, but not sent to other devices

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230616/2023-06-25_18-32.42fkyvxkkou8.webp)

### 0x01

Before Configureing the addresses, calculate the EUI-64 interface ID that will be generated on each interface

```
R1#show int g0/1
GigabitEthernet0/1 is up, line protocol is up (connected)
  Hardware is CN Gigabit Ethernet, address is 0030.f236.4502 (bia 0030.f236.4502)
---
EUI-64
2001:db8::230:f2ff:fe36:4502
```

```
R2#show int g0/1
GigabitEthernet0/1 is up, line protocol is up (connected)
  Hardware is CN Gigabit Ethernet, address is 0001.63b0.b802 (bia 0001.63b0.b802)
---
EUI-64
2001:db8:0:1:201:63ff:feb0:b802
```

Use EUI-64 to configure IPv6 addresses on G0/1 fo R1/R2

```
R1(config-if)#ipv6 add 2001:db8::/64 eui-64
R2(config-if)#ipv6 add 2001:db8:0:1::/64 eui-64
```

可以使用 `show ipv6 int br` 来查看对应的 IPv6 address

```
R1(config-if)#do show ipv6 int br
GigabitEthernet0/0         [up/up]
    unassigned
GigabitEthernet0/1         [up/up]
    FE80::230:F2FF:FE36:4502
    2001:DB8::230:F2FF:FE36:4502

R2(config-if)#do sh ipv6 int br
GigabitEthernet0/0         [up/up]
    unassigned
GigabitEthernet0/1         [up/up]
    FE80::201:63FF:FEB0:B802
    2001:DB8:0:1:201:63FF:FEB0:B802
```

### 0x02

Configure the appropriate IPv6 addresses/default gateways on PC

PC1

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::260:70FF:FE35:C92E
   IPv6 Address....................: 2001:DB8::1
   IPv4 Address....................: 10.0.1.2
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 2001:DB8::230:F2FF:FE36:4502
                                     10.0.1.254
```

PC2

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::201:43FF:FE8B:2D12
   IPv6 Address....................: 2001:DB8:0:1::2
   IPv4 Address....................: 10.0.2.2
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 2001:DB8:0:1:201:63FF:FEB0:B802
                                     10.0.2.254
```

### 0x03

Enable IPv6 on G0/0 of R1/R2 without explicitly configuring an IPv6 address

```
R1(config-if)#int g0/0
R1(config-if)#ipv6 enable

R2(config-if)#int g0/0
R2(config-if)#ipv6 enable
```

同样的也可以使用 `show ipv6 int br` 来查看

```
R1(config-if)#do sh ipv6 int br
GigabitEthernet0/0         [up/up]
    FE80::230:F2FF:FE36:4501
    
R2(config-if)#do sh ipv6 int br
GigabitEthernet0/0         [up/up]
    FE80::201:63FF:FEB0:B801
```

### 0x04

Configure Static routes on R1/R2 to enable PC1 to ping PC2

```
R1(config)#ipv6 unicast-routing 
R2(config)#ipv6 unicast-routing
```

如果这时直接指定 next hop 为对端的 link-local address 就会报错

```
R1(config)#ipv6 route 2001:db8:0:1::/64 FE80::201:63FF:FEB0:B801
% Interface has to be specified for a link-local nexthop
```

需要指定互联的接口

```
R1(config)#ipv6 route 2001:db8:0:1::/64 g0/0 FE80::201:63FF:FEB0:B801
```

这样就可以通过 `show ipv6 route` 看到刚才配置的路由

```
S   2001:DB8:0:1::/64 [1/0]
     via FE80::201:63FF:FEB0:B801, GigabitEthernet0/0
```

同理 R2

```
R2(config)#ipv6 route 2001:db8::/64 g0/0 FE80::230:F2FF:FE36:4501
```

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=BrTMMOXFhDU&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=61