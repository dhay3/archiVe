# Day44 - NAT(Part 1)

因为 IPv4 地址并不能为所有入网的设备提供 IP，所以需要方案来解决这个问题

- The long-term solution is to switch to IPv6

- There are three main short-term solutions

  1. CIDR

     无需按照 8 16 24 subnet mask 来划分网段

  2. Private IPv4 address and NAT

## Private IPv4 address

RFC 1918 规定了 Private IPv4 address 地址范围

10.0.0.0/8(10.0.0.0 to 10.255.255.255)

172.16.0.0/12(172.16.0.0 to 172.16.255.255)

192.168.0.0/16(192.168.0.0 to 192.168.255.255)

*You are free to use these addresses in your networks. They don’t have to be globally unique*

Private IP addresses cannot be used over the Internet

> 如果运营收到私网来的 IP 报文，就会丢弃该报文

既然私网 IP 不能出公网，那么对应的机器要怎么出公网呢？那就需要使用到 NAT

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_10-47.4f6pxats3ocg.webp)

*Although the private IP addresses are unique, the public addresses must be unique.*

> 如果没有 NAT 就会有问题
>
> 1. Duplicate addresses
>
>    如果有个报文目的地址是 192.168.0.167，那台机器会是实际的目的机器呢？
>
> 2. Private IP addresses can’t be used over the Internet, so the PCs can’t access the Internet

## NAT

*Network Address Translation(NAT) is used to modify the source and/or destination IP addresses of packets*

使用 NAT 的最主要的原因是为了让私网内的机器可以访问公网

> 在 CCNA 中主要是 source NAT

例如

192.168.0.167 需要访问 8.8.8.8

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_10-53.6o7ip24wg5j4.webp)

1. R1 在收到 PC1 过来的请求，会将源 IP 替换层 203.0.113.1(这里对应 R1 的接口 IP，但是实际有很多种不同的方式来选择源 IP)
2. 服务器在收到从 R1 来的请求，会回目的到 203.0.113.1
3. 响应到 R1 会替换目的地址为原发送报文的 IP 192.168.0.167，这里并不是 Destination NAT

## Static NAT

Static NAT involves statically configuring **one-to-one** mappings of private IP addresses to public IP addresses

将其想象成一个私网 IP 对应一个公网 IP，两者是一对一的关系

- An *inside local* IP address is mapped to an *inside global* IP address

  - Inside Local = The IP address of the inside host, from the perspective of the local network

    the IP address actually configured on the inside host, usually a private address

    简单理解就是出向源私网地址

  - Inside Global = The IP address of the inside host, from the perspective of outside hosts

    the IP address of the inside host <u>after NAT</u>, usually a public address

    简单理解就是出向源私网 NAT 后的地址

例如下图

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_11-26.6t55cbq6jg1s.webp)

*Static NAT allows devices with private IP addresses to communicate over the Internet*

*However, because it requires a one-to-one IP address mapping, it doesn’t help preserve IP addresses*

> 也是因为一对一的关系，外部的设备也可以通过 inside global 地址直接访问 inside local 设备

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_11-15.hticgsmbekg.webp)

### Static NAT Configuration

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_14-45.4ms20us2mxkw.webp)

- `R1(config)#int <interface-id>`

  `R1(config-if)#ip nat inside`

  声明 inside local 互联的接口

- `R1(config)#int <interface-id>`

  `R1(config-if)#ip nat outside`

  声明 inside golbal 使用的接口

- `R1(config)#ip nat inside source static <source> <source nat>`

  声明 SNAT 规则

  > 例子中使用的 source nat 地址是公网，你并不能随便的使用这些 IP，只有公网的地址是被注册给你的才可以使用

配置完后可以使用 `R1#show ip nat translations` 查看 NAT table

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_14-49.6plcbe1vgvi8.webp)

只要你使用了 Static NAT 就会出现第 2 第 4 行的情形，其余的都是 dynamic translation

- `Pro`

- `Inside global`

  SNAT 后的源地址

- `Inside local`

  SNAT 前的源地址

- `Outside local`

  The IP address of the outside host, from the perspective of the local network

  DNAT 前的地址

- `Outside global`

  The IP address of the outside host,f rom the perspective of the outside network

  DNAT 后的地址

> 除非使用 Destination NAT 否则 Outside local 和 Outside global 地址都会一样

如果需要手动清空 NAT table 在的 dynamic translation 可以使用 `R1#clear ip nat translation *`

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_15-59.2936fxqbl7eo.webp)

> 即使不手动清空 NAT table, 在 192.168.0.167/168 不在需要访问 100.0.0.1/2 一段时间后也会从 NAT table 中自动清空

还可以使用 `R1#show ip nat statistics` 来查看 NAT 统计的信息

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_16-08.5hwa4f938r9c.webp)

- `Total active translations`

  当前 NAT table 中一个有多少条 NAT 条目

- `Peak translations`

  设备最多到达的 NAT 条目(不是 capacity 只是一个状态)

- `Outside interfaces`

  `Inside interfaces`

## Command Summary

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_16-09.5ttvo0ju2glc.webp)

## Quiz

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_16-14.1h7xg9x1xsw0.webp)

这里选 C 是因为，如果使用 Static NAT 私网地址和公网地址是一一对应的，在输入第二条命令后设备会直接报错

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_16-17.5psn883g2o74.webp)

只有先声明映射关系的可以使用对应的公网地址，如果需要 10.0.0.2 也可以出公网，要使用另外一个公网的 IP

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-11_16-10.32cqrz4xu3uo.webp)

### 0x01

Attempt to ping from PC1 to 8.8.8.8. Does the ping work

不通，因为需要过公网

### 0x02

Configure static NAT on R1

- Configure the appropriate inside/outside

  ```
  R1(config)#int g0/1
  R1(config-if)#ip nat inside 
  R1(config-if)#int g0/0
  R1(config-if)#ip nat outside
  ```

  可以使用 `show ip nat statistics` 查看配置的端口是否正确

- Map the IP addresses of PC1, PC2 and PC3 to 100.0.0.x/24

  ```
  R1(config)#ip nat inside source static 172.16.0.1 100.0.0.1
  R1(config)#ip nat inside source static 172.16.0.2 100.0.0.2
  R1(config)#ip nat inside source static 172.16.0.3 100.0.0.3
  ```

### 0x03

Ping 8.8.8.8 from PC1 again. Does the ping work

可以

### 0x04

Ping google.com from each PC, and the check the NAT translations on R1

```
R1(config)#do show ip nat tra
Pro  Inside global     Inside local       Outside local      Outside global
icmp 100.0.0.2:1       172.16.0.2:1       8.8.8.8:1          8.8.8.8:1
icmp 100.0.0.2:2       172.16.0.2:2       8.8.8.8:2          8.8.8.8:2
icmp 100.0.0.2:3       172.16.0.2:3       8.8.8.8:3          8.8.8.8:3
icmp 100.0.0.2:4       172.16.0.2:4       8.8.8.8:4          8.8.8.8:4
icmp 100.0.0.3:1       172.16.0.3:1       8.8.8.8:1          8.8.8.8:1
icmp 100.0.0.3:2       172.16.0.3:2       8.8.8.8:2          8.8.8.8:2
---  100.0.0.1         172.16.0.1         ---                ---
---  100.0.0.2         172.16.0.2         ---                ---
---  100.0.0.3         172.16.0.3         ---                ---
```

### 0x05

Clear the NAT translations on R1. Which entries remain

```
R1(config)#do clear ip nat tran *
R1(config)#do show ip nat tran
Pro  Inside global     Inside local       Outside local      Outside global
---  100.0.0.1         172.16.0.1         ---                ---
---  100.0.0.2         172.16.0.2         ---                ---
---  100.0.0.3         172.16.0.3         ---                ---
```

**references**

1. [^https://community.cisco.com/t5/other-network-architecture-subjects/difference-between-enable-secret-command-and-service-password/td-p/173461]