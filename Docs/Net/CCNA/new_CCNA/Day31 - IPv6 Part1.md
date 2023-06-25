# Day31 - IPv6 Part1

## What about IPv5

IPv4 为什么直接到了 IPv6, 中间的 IPv5 去那了？在谈 IPv6 之前先谈一下 IPv5

*‘Internet Stream Protocol’ was developed in the late 1970s, but never actually introduced for public use, the Version filed of the protocol’s IP header is 5, but never called ‘IPv5’*

*SO, when the successor to IPv4 was being developed, it was name IPv6,  uses a value of 6  in the Version field of the header*

## Hexadecimal

IPv6 和 MAC address 一样都使用 Hexadecimal

先回顾一下 Binary/Decimal/Hexadecimal

- Binary 也被称为 Base 2 (在二进制中每 bit 只能是 0,1 两个数字，所以也被称为 base 2)使用 0b 来标识
- Decimal 也被称为 Base 10 (在十进制中每 bit 只能是 0-9 9 个数字，所以也被称为 base 10)使用 0d 来标识
- Hexadecimal 也被称为 Base 16 (在十六进制中每 bit 只能是 0-F 16个数字，所以也被称为 base 16)使用 0x 来标识

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-56.44w3noporkqo.webp)

> 如果一个数字 10 没有标识是 0b/0d/0x 是分辨不出具体的值的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-57.4shjtjl7emtc.webp)

> 注意在 Binary 列中，部分数字有 leading 0，实际在表示二进制数时可以不需要 leading 0
>
> 例如 0011 完全可以使用 11 来表示
>
> 这里只是为了方便记忆，因为 Hexadeicmal 一位需要由 4 位 Binary 数字来组成(因为 hexdecimal 最大 15，对应 Binary 1111，至少 4 位)

0b11011011 = 0b1101 0b1011 = 0d13 0d11 = 0xD 0xB = 0xDB

> binary 转 hexdecimal 先拆分成 4  bit 一组，然后转成 10 进制后，再转成 16 进制

0b00101111 = 0b0010 0b1111 = 0d2 0d15 = 0x2 0xF = 0x2F

0b10000001 = 0b1000 0b0001 = 0d8 0d1 = 0x8 0x1 = 0x81



0xEC = 0xE 0xC = 0d14 0d12 = 0b1110 0b1100 = 0b11101100

> hexdecimal 转 binary 就按照 binary 转 hexdecimal 逆序

0x2B = 0x2 0xB = 0d2 0d11 = 0b0010 0b1011 = 0b00101011

0xD7 = 0xD 0x7 = 0d13 0d7 = 0b1101 0b0111 = 0b11010111

## Why IPv6

为什么使用 IPv6, 最主要的原因是 

there simply aren’t enough IPv4 address available

因为 IPv4 address 一共 32 bit，所以最多就 $2^{32} = 4294967296$ 个地址。IPv4 在 30 年前被设计，网络发展之迅速，是让所有人都想不到的，所以地址自然也就不够用

为了应对 IPv4 地址不够的原因，有 3 种方案

1. VLSM(variable length subnet mask) 让 IPv4 地址按需划分，更能被有效的使用
2. NAT(network address translation) 无需为每台主机配置单独的公网 IP 也能访问公网
3. Private IPv4 address

这些都是 short-term solutions，long-term solution 应该使用 IPv6

## IPv6 address

- An IPv6 address is 128 bits

  你可能会认为 128 是 32 的 4 倍，所以 IPv6 的可用的地址是 IPv4 的 4 倍。但是这个是错误的，实际 IPv6 可用的地址应该为 $2^{128}$ 

  这几乎是一个天文数字，完全可以满足现在的需求

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_16-27.5m0v2wpvf3wg.webp)

- An IPv6 address is written in hexadecimal

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_16-28.6zdcmsht44g0.webp)

  > 显然如果用 binary 或者是 decimal 来表示 IPv6 地址太复杂了

  IPv6 会以 16 bits 为一组

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_16-45.5tme9ojdq2rk.webp)

  且 IPv6 直接会使用 slash notation 来表示 network portion，在 Cisco 设备中也可以直接使用 slash 的方式来表示 subnet mask，无需使用 dotted decimal subnet mask

## Shortening(abbreviating) IPv6

如果要配置 IPv6 地址要写这么一大串，当然也可以 shortening(abbreviating) IPv6 地址

- Leading 0s can be removed

  例如红框中的前导 0 可以被直接删除

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_16-52.xmmqhd2uokw.webp)

- Consecutive(连续的) quartets(这里指的就是以 : 分隔的部分) of all 0s can be replaced with a double colon(::)

  例如红框中有 4 个连续的 0000，是可以直接使用 :: 替换

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_16-54.5yjj6wg20dc0.webp)

  为什么可以这样来划分呢？因为我们知道 IPv6 是按照 16 bits 划分为 8 个 quartets 的，所以我们可以按照剩余非 :: 部分的 quartets 部分，就可以推算出全零 quartets 的个数。例如上图中的 2001:0DB8::0080:34BD 就可以推出全零 quartets 部分为 4 个

  > *Consecutive quartets of 0s can only be abbreviated once in an IPv6 address*
  >
  > 需要注意的一点是，这条规则在 IPv6 address 中只能被使用一次

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_17-09.6epnhdi9dji8.webp)

  因为你不能判断每一个 :: 到底有多少个 quartets of 0

  我们应该使用 2001::20A1:0:0:34BD 这个地址

  > 0000 去掉 leading 0 就是 0

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_17-12.5y9vh3zy3b4.webp)

理所当然，我们可以组合这两个简写的方式

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_17-07.2ekzr3hlb2ps.webp)

下面几个例子加强一下理解

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-21_18-21.5vmgzq5nsyyo.webp)

==在 RFC 5952 中有几条规则==

1. Leading 0s must be removed
2. :: must be used to shorten the longest string of all-0 quartets
3. if there are two equal-length choices for the ::, use :: to shorten the one on the left

4. Hexadecimal characters abcdef must be written using lower-case, not upper-case ABCDEF

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_19-59.5ykwtfrf5cao.webp)

> 即使在 Cisco 的设备上 IPv6 地址显示的也是大写的，但是为了按照标准尽量使用小写的 IPv6 地址

## Expanding shortened IPv6 address

IPv4 可以 shorten address，当然也可以 expand address

例如 FE80::2:0:0:FBE8

1. 先将 leading 0 添加回来 FE80::0002:0000:0000:FBE8
2. 将 :: 扩展， 因为除 :: 外还有 5 个 quartets，所以全 0 quartets 部分为 3 个， FE80:0000:0000:0000:0002:0000:0000:FBE8

下面几个例子加强一下理解

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-21_18-33.43f128xxmadc.webp)

## IPv6 prefix

通常如果一个公司向 ISP 申请 IPv6 的地址，ISP 会为其分配一段 /48 的地址，而公司 subnets 一般会使用 /64 划分，意味着

- 64 - 48 = 16 bits 用做子网划分
- 128 - 64 = 64 bits 用做 host portion

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-21_18-47.tlw0qmt5xuo.webp)

例如

2001:0DB8:8B00:0001:0000:0000:0000:0001/64 

64/16 = 4 ... 0

所以 network portion 就是 2001:DB8:8B00:1::/64

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-21_18-49.qctd81fwv4g.webp)

300D:00F2:0B34:2100:0000:0000:1200:0001/56

56/16 = 3 ... 8

8/4 = 2 ... 0

所以 network portion 就是 300D:F2:B34:21::/56 对吗？

实际上应该为 300D:F2:B34:2100::/56，如果是上面的在扩展 IPv6 时地址就会变成 300D:F2:B34:0021::/56 显然和原始的地址不一样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-21_18-54.4fmkmdu3ps74.webp)

2001:0D88:8B00:0001:FB89:017B:0020:0011/93

93/16 = 5 ... 13

13/4 = 3 ... 1

13 并不是 4 的倍数，取模后会余 1

关注 92 - 96 对应的 B，转成 10 进制就是 0d11，然后转成 2 进制 0b1011

因为只余 1 bit，所以取消 0b1011 后 3 位，即 0b1000，转成 10 进制就是 0d8，转成 16 进制就是 0x8

所以 network portion 就是 2001:D88:8B00:1:FB89:178::/56

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-21_19-00.2iljdmfifdds.webp)

下面几个例子加强一下理解

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_13-42.3z7ao8tnejr4.webp)

## Configuring IPv6 addresses

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_13-50.t5twj7j59f.webp)

> 这里使用 2001:db8 是保留地址，类似与 IPv4 中的私网地址

1. `ipv6 unicast-routing`

   *allows the router to perform IPv6 routing*

   必须使用该命令来声明并启用 IPv6 路由的功能

2. `ipv6 address 2001:db8:0:0::1/64`

   和配置 IPv4 不同，IPv6 可以直接使用 prefix 来表示 subnet mask

配置完接口之后和 IPv4 类似，可以通过 `show  ipv6 interface brief` 来查看接口的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_13-56.3hviz0l8fb5s.webp)

这里可以看到除了红框外手动配置的 IPv6 地址外，还有类似 IPv6 的地址

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_13-58.1usltgixf7ls.webp)

这些被称为 **link-local addresses**，在端口配置了 IPv6 addresses 后会自动配置在对应的端口上

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_14-45.35wnci47vveo.webp)

### 0x01

Enable IPv6 routing on R1

```
R1(config)#ipv6 unicast-routing 
```

### 0x02

Configure the appropriate IPv6 addresses on R1

```
R1(config)#int g0/0
R1(config-if)#ipv6 add 2001:DB8:0:1::1/64
R1(config)#int g0/1
R1(config-if)#ipv6 add 2001:DB8:0:2::1/64
R1(config)#int g0/2
R1(config-if)#ipv6 add 2001:DB8:0:3::1/64
```

### 0x03

Confirm your configurations. What IPv6 addresses are present on each interface?

```
R1(config-if)#do sh ipv6 int br
GigabitEthernet0/0         [up/up]
    FE80::201:97FF:FE9A:AC01
    2001:DB8:0:1::1
GigabitEthernet0/1         [up/up]
    FE80::201:97FF:FE9A:AC02
    2001:DB8:0:2::1
GigabitEthernet0/2         [up/up]
    FE80::201:97FF:FE9A:AC03
    2001:DB8:0:3::1
```

### 0x04 

Configure the appropriate IPv6 addresses on each PC. Configure the correct default gateway

PC1

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::200:CFF:FE44:7047
   IPv6 Address....................: 2001:DB8:0:1::2
   IPv4 Address....................: 192.168.1.2
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 2001:DB8:0:1::1
                                     192.168.1.1
```

PC2

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::2D0:97FF:FE12:77E2
   IPv6 Address....................: 2001:DB8:0:2::2
   IPv4 Address....................: 192.168.2.2
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 2001:DB8:0:2::1
                                     192.168.2.1

```

PC3

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::230:A3FF:FEB0:AC87
   IPv6 Address....................: 2001:DB8:0:3::2
   IPv4 Address....................: 192.168.3.2
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 2001:DB8:0:3::1
                                     192.168.3.1
```

### 0x05

Attempt to ping between the PCs(IPv4 and IPv6)

```
C:\>ping 2001:DB8:0:1::1

Pinging 2001:DB8:0:1::1 with 32 bytes of data:

Reply from 2001:DB8:0:1::1: bytes=32 time<1ms TTL=255
Reply from 2001:DB8:0:1::1: bytes=32 time<1ms TTL=255

Ping statistics for 2001:DB8:0:1::1:
    Packets: Sent = 2, Received = 2, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 0ms, Maximum = 0ms, Average = 0ms
```

windows 中的 ping 命令可以直接使用 IPv6 地址，而 Linux 中需要使用特定的参数

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=ZNuXyOXae5U&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=59