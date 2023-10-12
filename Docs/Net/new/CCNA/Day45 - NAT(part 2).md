# Day45 - NAT(part 2)

## Dynamic NAT

*In dynamic NAT, the router dynamically maps inside local addresses to inside global addresses as needed*

Dynamic NAT 通过 ACL 来判断那些流量需要被 NAT

- If the source IP is permitted by the ACL, the source IP will be translated
- If the source IP is denied by the ACL, the source IP will NOT be translated.***the traffic will NOT be dropped***

> 如果匹配 permit ACL 就会被 NAT，如果匹配 deny ACL 或者是 implicit deny 仅仅只是不会被 NAT，并不会被丢包

Dynamic NAT 通过 NAT pool 来定义可以使用的 inside global addresses

例如

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_11-21.fgfpzl8ou0w.webp)

虽然被称为 Dynamic NAT，但是实际上还是一对一

*one inside local address per inside global IP address*

## NAT pool exhaustion

如果没有足够的 inside global 地址(所有的 inside global 地址都被分配了)，被称为 **NAT pool exhaustion**

例如上图中 ACL 允许 192.168.0.0/24 最大一共 256 台机器，但是 NAT pool 一个只有 10 个地址可以用，如果出现 10 个地址都被分配的现象，其他机器要想出公网就会出现 NAT pool exhaustion 的现象

- If a packet from another inside host arrives and needs NAT but there are no available addresses, the router will drop the packet

- The host will be unable to access outside networks until one of the inside global IP addresses becomes available

  *Dynamic NAT entries will time out automatically if not used, or you can clear them manually(`clear ip nat translation *`)*

例如

100.0.0.1 - 10 都被分配了，这时 192.168.0.98 需要出公网，就不能分配 NAT 地址，Router 就会==丢弃==这个报文

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_11-49.36ds4jr2m6kg.webp)

当 192.168.0.167 数据传输完，地址次的 100.0.0.1 变为可用的状态，这时如果 192.168.0.98 想要出公网就会 SNAT 成 100.0.0.1

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_11-52.2lg8i73srn40.webp)

## Dynamic NAT Configuration

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_12-07.3v3p5uk8t8n4.webp)

- `R1(config)#int <interface-id>`

  `R1(config-if)#ip nat inside`

  和 Static NAT 一样，指定 inside local 互联的接口

- `R1(config)#int <interface-id>`

  `R1(config-if)#ip nat outside`

  和 Static NAT 一样，指定 inside global 出接口

- `R1(config)#access-list <num> permit <source> <subnet mask>`

  指定 Dynamic NAT 需要使用的 ACL

- `R1(config)#ip nat pool <pool name> <start address> <end address> prefix-length <prefix-length>`

  规定 NAT pool 使用的范围

- `R1(config)#ip nat inside source list <num> pool <pool name>`

  指定 Dynamic NAT 使用的 ACL 和 NAT pool

这时使用 `show ip nat translations` 来查看 NAT table

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_12-23.5x4mkw3lkw00.webp)

其中第 3 第 6 行表示只要从 inside local 来就会 SNAT 成 inside global，如果期间没有流量默认 24 消失会清空，其余部分是 dynamic entries，一分钟没有流量就会自动从 NAT table 中清空

同样的和 Static NAT 一样使用 `clear ip nat translations *` 会将 dynamic entries 从 NAT table 中手动清空

看一下 `R1#show ip nat statistics`

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_12-34.39dso06j7934.webp) 

可以看到 6 dynamic 表示 NAT table 中一共有 6 条 entries, 4 extended 表示 dynamic entries

还可以看到使用的 ACL 和 NAT pool

## PAT

如果想要不同 inside local 地址使用相同的 inside global 地址就需要使用 PAT 也被称为 NAT Overload

*By using a unique port number for each communication flow, a single public IP address can be used by many different internal hosts.(port number are 16 bits = over 65535 avaliable port numbers)*

例如

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_12-47.5evh83d61wxs.webp)

1. PC1 192.168.0.167 使用 port 54321 访问 8.8.8.8:53，同时 PC2 192.168.0.168 使用 port 54321 访问 8.8.8.8:53
2. R1 会将 PC1 发送过来的报文源地址 NAT 成 100.0.0.1:54321，将 PC2 发送过来的报文源地址 NAT 成 100.0.0.1:543322(这个端口一般随机)
3. SRV 收到 R1 发送过来的报文后，返回响应报文
4. R1 在收到报文后，将目的地址是 100.0.0.1:54321 的按照 Session 转成 192.168.0.167:54321; 将目的地址是 100.0.0.1:54322 的按照 Session 转发 192.168.0.168:54321

> 如果 PC2 使用的端口不是 54321，R1 不会使用随机的端口，而是会按照 PC2 使用的端口来转换。==通常只有两台或者是多台机器使用了相同的源端口，且 SNAT IP 地址相同，才会出现 SNAT 后源端口改变的情况==

PAT 转换的逻辑如下

a b 分别为两台不同 IP 的主机, r 为 a b NAT 后的 3 层地址

```
if a.port == b.port then
	nat(a) = r + new(a.port)
	nat(b) = r + new(b.port)
else
	nat(a) = r + a.port
	nat(b) = r + b.port
```

*Because many inside hosts can share a single public IP, PAT is very useful for preserving public IP addresses, and it is used in networks all over the world*

因为 PAT 的这种特性，也是生产中使用最多的方案

### PAT Configuration

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_12-58.5xn8j244djb4.webp)

和 Dynamic NAT 的配置类似，只有一点不同

`R1(config)#ip nat inside source list <num> pool <pool name> overload`

在原先的配置命令后加 overload

看一下 `show ip nat translations` 和 `show ip nat statistics` 的信息

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_13-02.4zkys7y415hc.webp)

这里并没有类似 `---...` 的 entries，因为 PAT 是一对多的关系，不存在 Staic NAT 和 Dynamic NAT 中一对一的关系，所以只会显示 dynamic entries

### PAT Configuration(interface)

除了上面的方式配置 PAT，还有另外一种方式，NAT 地址使用自己接口的地址而不使用 NAT pool

使用 `R1(config)#ip nat inside source list <num> interface <interface-id> overload` 命令，指定 inside global 地址为 interface g0/0 的 IP，同时启用 PAT

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_13-07.mq938yizozk.webp)

例如

PC1/2 想要出公网就会使用 R1 G0/0 接口的 IP 203.0.113.1 作为 SNAT 地址

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_13-09.3fem3sj4xgqo.webp)

## Command Summary

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_13-11.53jxozv11wg0.webp)

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230706/2023-07-12_13-26.3siy1ub0tsxs.webp)

### 0x01

Configure dynamic NAT on R1

Configure the appropriate inside/outside interfaces

```
R1(config)#int g0/1
R1(config-if)#ip nat inside 

R1(config-if)#int g0/0
R1(config-if)#ip nat outside
```

Translate all traffic from 172.16.0.0/24

```
R1(config)#access-list 1 permit 172.16.0.0 0.0.0.255
```

Create a pool of 100.0.0.1 to 100.0.0.2 from the 100.0.0.0/24 subnet

```
R1(config)#ip nat pool POOL1 100.0.0.1 100.0.0.2 netmask 255.255.255.0
R1(config)#ip nat inside source list 1 pool POOL1 
```

### 0x02

Ping google.com from PC1 and PC2. Then, ping it from PC3. What happens to PC3’s ping?

PC3 会失败，因为 100.0.0.1 和 100.0.0.2 都被分配了

```
R1(config)#do show ip nat trans
Pro  Inside global     Inside local       Outside local      Outside global
icmp 100.0.0.1:1       172.16.0.1:1       172.217.175.238:1  172.217.175.238:1
icmp 100.0.0.1:2       172.16.0.1:2       172.217.175.238:2  172.217.175.238:2
icmp 100.0.0.1:3       172.16.0.1:3       172.217.175.238:3  172.217.175.238:3
icmp 100.0.0.1:4       172.16.0.1:4       172.217.175.238:4  172.217.175.238:4
icmp 100.0.0.2:1       172.16.0.2:1       172.217.175.238:1  172.217.175.238:1
icmp 100.0.0.2:2       172.16.0.2:2       172.217.175.238:2  172.217.175.238:2
icmp 100.0.0.2:3       172.16.0.2:3       172.217.175.238:3  172.217.175.238:3
icmp 100.0.0.2:4       172.16.0.2:4       172.217.175.238:4  172.217.175.238:4
udp 100.0.0.1:1025     172.16.0.1:1025    8.8.8.8:53         8.8.8.8:53
udp 100.0.0.2:1025     172.16.0.2:1025    8.8.8.8:53         8.8.8.8:53
```

> 注意这里的 ICMP 报文 `:` 后面的部分并不是端口，而是 3 层 ICMP 报文中的 sequence number

### 0x03

Clear the NAT translations and remove the current NAT configuration. 

```
R1(config)#do clear ip nat tra *
R1(config)#do show ip nat tra

R1(config)#do show run | se nat
 ip nat outside
 ip nat inside
ip nat pool POOL1 100.0.0.1 100.0.0.2 netmask 255.255.255.0
ip nat inside source list 1 pool POOL1
R1(config)#no ip nat pool POOL1 100.0.0.1 100.0.0.2 netmask 255.255.255.0
%Pool POOL1 in use, cannot destroy
R1(config)#no ip nat inside source list 1 pool POOL1
R1(config)#no ip nat pool POOL1 100.0.0.1 100.0.0.2 netmask 255.255.255.0
```

这里需要先取消 `ip nat inside source...` 才可以取消 NAT pool 配置

Switch the Configuration to PAT using R1’s public IP address

```
R1(config)#ip nat inside source list 1 interface g0/0 overload 
```

这里也可以直接使用这条命令，无需使用 `no`，会指定替换对应的命令，但是不能取消 NAT pool

### 0x04

Ping google.com from each PC. Do the pings work? Examine the NAT translations on R1

正常，可以在 R1 上使用 `show ip nat trans` 来查看

```
R1(config)#do show ip nat tran
Pro  Inside global     Inside local       Outside local      Outside global
icmp 203.0.113.1:1024  172.16.0.2:5       172.217.175.238:5  172.217.175.238:1024
icmp 203.0.113.1:1025  172.16.0.2:6       172.217.175.238:6  172.217.175.238:1025
icmp 203.0.113.1:1026  172.16.0.2:7       172.217.175.238:7  172.217.175.238:1026
icmp 203.0.113.1:1027  172.16.0.2:8       172.217.175.238:8  172.217.175.238:1027
icmp 203.0.113.1:1     172.16.0.3:1       172.217.175.238:1  172.217.175.238:1
icmp 203.0.113.1:2     172.16.0.3:2       172.217.175.238:2  172.217.175.238:2
icmp 203.0.113.1:3     172.16.0.3:3       172.217.175.238:3  172.217.175.238:3
icmp 203.0.113.1:4     172.16.0.3:4       172.217.175.238:4  172.217.175.238:4
icmp 203.0.113.1:5     172.16.0.1:5       172.217.175.238:5  172.217.175.238:5
icmp 203.0.113.1:6     172.16.0.1:6       172.217.175.238:6  172.217.175.238:6
icmp 203.0.113.1:7     172.16.0.1:7       172.217.175.238:7  172.217.175.238:7
icmp 203.0.113.1:8     172.16.0.1:8       172.217.175.238:8  172.217.175.238:8
udp 203.0.113.1:1024   172.16.0.2:1026    8.8.8.8:53         8.8.8.8:53
udp 203.0.113.1:1025   172.16.0.3:1026    8.8.8.8:53         8.8.8.8:53
udp 203.0.113.1:1026   172.16.0.1:1026    8.8.8.8:53         8.8.8.8:53
```

**references**

1. [^https://community.cisco.com/t5/other-network-architecture-subjects/difference-between-enable-secret-command-and-service-password/td-p/173461]
2. [^https://www.cisco.com/c/en/us/support/docs/ip/network-address-translation-nat/13771-10.html]