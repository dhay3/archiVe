# Day13 - Subnetting

## CIDR

Classless Inter-Domain Routing (CIDR) 可以让我们更加灵活地使用 IPv4 address，不限于 ABC 类地址

假设有如下一个 Point-to-Point 网络

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-26.6877jhheykxs.webp)

显然的在 203.0.133.0/24 中只需要 4 个 IP address，所以有 252 IP addresses 被浪费了

假设 Company X 需要 5000 个 end hosts，如果不使用 CIDR，采用 class C network 地址是不够的，如果采用 class B network 可用地址为 60000 又太大了

*CIDR 就是为了解决这个问题*

With CIDR, the requirements of

Class A = /8

Class B = /16

Class C = /24

were removed

This allowd larger networks to be split into smaller networks, allowing greater efficiency

These smaller network are called ‘subnetworks’ or ‘subnets’

## Point-to-Point

假设现在一个点对点网络配置了 203.0.113.0/24 network

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-35.6q2rq6yy2qo0.webp)

就会有 254 个可用的地址，公式如下
$$
2^{n} - 2 = usable\ address\\
n = number\ of\ host\ bits
$$
我们使用的 number of host bits 越小，浪费的地址也就越小

如果使用 30 bits mask，那么一共可用的地址就是 2 个，做到地址零浪费

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-40.2iwdamdvy5c0.webp)

> 但是在 Point-to-Point 网络中，实际是不需要 3 层广播地址的或者是 network 地址的，所以 203.0.113.1 和 203.0.113.0 也是可以被使用的
>
> *在 Point-to-Point network	/31 可以用做 mask*

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-51.4s7epr3kdr0g.webp)

如果使用 32 bits mask 呢？

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-53.2dohmwtldtk.webp)

显然是不行的，32 bits mask 是不能被用在 interface 上表示一个网段的，但是可以用在路由中表示最高精度的匹配

## CIDR Notation

在 CIDR 中，mask 部分被称为 CIDR Notation

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-56.5o438f8omvwg.webp)

## Subnetting

假设有 192.168.1.0/24 network 需要划分成 4 个 subnets，每个 subnets 中可以包含 45 hosts

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_20-57.4mhbndu56xs0.webp)

$47 \times 4=188$ 显然是小于 256 的，所以这个需求是可以被满足的，大于 47 最小的 2 进制数是 64，即 $2^6$

所以 host bits 至少是 6 位，即 CIDR notation 为 26

所以可以划分成

- 192.168.1.0/26

  192.168.1.1 - 192.168.1.63

- 192.168.1.64/26

  192.168.1.65 - 192.168.1.127

- 192.168.1.128/26

  192.168.1.129 - 192.168.1.191

- 192.168.1.192/26

  192.168.1.193 - 192.168.1.254

是怎么计算的呢？

1. Find the broadcast address of first subnet
2. The next address of first subnet is the network address of second subnet

> 即为前一个 subnet 广播地址加 1

1. reapeat the process 

### Another example for subnetting

现在需要将 192.168.255.0/24 划分成 5 个 subnets

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_21-21.6ap3utkzaruo.webp)

现在并不知道 host number 所以也就不能使用上面的方法

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_21-27.54kuyg6rhqbk.webp)

我们可以使用“借位法”，1 bit 有 2 种可能 0 和 1，所以借 1 bit 就可以划分出 2 subnets，如果需要划分出 5 subnets，至少需要 3 位，即 CIDR notation /27

- 192.168.255.0/27
- 192.168.255.32/27
- 192.168.255.64/27
- 192.168.255.96/27
- 192.168.255.128/27

### Subnetting Trick

subnetting 有一个小技巧

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_21-13.6j370b0cwx34.webp)

因为至少需要 6 bits CIDR natation，所以只关注 the last of octets

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_21-16.1jzdlo08ra4g.webp)

因为 the first subnet 是很容易就计算出来的，所以在 the first subnet 的基础上 **add the last bit of the network portion** 就是 next subnet

## Identify the subnet

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_21-32.4n3311tlxo1s.webp)

只需要去掉，host portion 部分的 bit 即可，所以为 192.168.5.32

## VLSM

Variable-Length Subnet Masks (VLSM) is the process of creating subnets of different sizes, to make your use of network address more efficient

VLSM 就是一种划分子网的方法，主要有如下几步

1. Assign the largest subnet at the start of the address sapce
2. Assign the second-largest subnet after it
3. Repeat the process until all subnets have been assigned

> 总的来说就是按照，实际需要的 end-host 大小顺序来划分子网

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_22-26.6uw9hv8753pc.webp)

> 还需要判断一些 192.168.1.0/24 IP 数是否够划分
>
> 256 - 110 - 45 - 29 - 8 >= 0

### Tokyo LAN A

因为需要 110 hosts，所以 host portion 至少需要 7 bis(128)，所以可以推出

| Network address | Broadcast address | First usable address | Last usable address | Total number of usable host address |
| --------------- | ----------------- | -------------------- | ------------------- | ----------------------------------- |
| 192.168.1.0/25  | 192.168.1.127/25  | 192.168.1.1/25       | 192.168.1.126/25    | 126                                 |

## Toronto LAN B

因为 192.168.1.0/25 可用的范围为 192.168.1.1/25 - 192.168.1.127/25 被 Tokyo LAN A 占用，所以需要从下个 subnet 开始计算 192.168.1.128

因为 45 hosts，所以 host portion 至少需要 6 bits(64)，所以可以推出

| Network address  | Broadcast address | First usable address | Last usable address | Total number of usable host address |
| ---------------- | ----------------- | -------------------- | ------------------- | ----------------------------------- |
| 192.168.1.128/26 | 192.168.1.191/26  | 192.168.1.129/26     | 192.168.1.190/26    | 62                                  |

## Toronto LAN A

因为 192.168.1.129/26 - 192.168.1.191/26 被 Toronto LAN B 占用，所以需要从下个 subnet 开始计算 192.168.1.192

因为 29 hosts，所以 host portion 至少 5 bits(32)，所以推出

| Network address  | Broadcast address | First usable address | Last usable address | Total number of usable host address |
| ---------------- | ----------------- | -------------------- | ------------------- | ----------------------------------- |
| 192.168.1.192/27 | 192.168.1.223/27  | 192.168.1.193/27     | 192.168.1.222/27    | 30                                  |

## Tokyo LAN B

因为 192.168.1.193/27 - 192.168.1.206/27 被 Toronto LAN A 占用，所以需要从下个 subnet 开始计算 192.168.1.224

因为 8 hosts，所以 host portion 至少 4 bits(16)，所以推出

| Network address  | Broadcast address | First usable address | Last usable address | Total number of usable host address |
| ---------------- | ----------------- | -------------------- | ------------------- | ----------------------------------- |
| 192.168.1.224/28 | 192.168.1.239/28  | 192.168.1.225/28     | 192.168.1.238/28    | 14                                  |

## Point-to-Point Connection

因为 192.168.1.225/28 - 192.168.1.239/28 被 Tokyo LAN B 占用，所以需要从下个 subnet 开始计算 192.168.1.240

Point-to-Point connection 只需要分配 2 个可用的 IP address，所以可以使用 /30 和 /31，但是优先使用 /30

| Network address  | Broadcast address | First usable address | Last usable address | Total number of usable host address |
| ---------------- | ----------------- | -------------------- | ------------------- | ----------------------------------- |
| 192.168.1.240/30 | 192.168.1.243/30  | 192.168.1.241/30     | 192.168.1.242/30    | 2                                   |

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-26_14-07.2j5q1ie12juo.webp)

LAN2
192.168.5.0/25
192.168.5.1/25
R1 G0/1 192.168.5.126/25

LAN1
192.168.5.128/26
192.168.5.129/26
R1 G0/0 192.168.5.190/26

LAN3
192.168.5.192/28
192.168.5.193/28
R2 G0/0 192.168.5.206/28

LAN4
192.168.5.208/28
192.168.5.209/28
R2 G0/1 192.168.5.222/28

P2P
192.168.5.224/30
R1 G0/0/0 192.168.5.225/30
R2 G0/0/0 192.168.5.226/30

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=aHwAm8GYbn8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=19