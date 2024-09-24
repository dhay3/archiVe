# Day03 - Address Resolution Protocol

## Layer2 and Layer3 Address

![](https://github.com/dhay3/image-repo/raw/master/20230825/2023-08-25_09-40.mmwkkutjk1c.webp)

Layer2 和 Layer3 地址提供的功能把不一样

- Layer2 provides hop-to-hop addressing ==within each network segment==
  - Deals with directly connected devices
- Layer3 provides end-to-end addressing: from the source host to destination host
  - Deals with indirectly(and directly) connected devices

> 简而言之就是
>
> The Layer 3 address is destined for the end host, and Layer 2 address is used to pass the packet to the next hop in the path to the end host

报文在传输的过程中(这里不考虑 NAT)，Layer3 address 不会改变，只会改变 Layer2 address

## ARP Overview

![](https://github.com/dhay3/image-repo/raw/master/20230825/2023-08-25_09-55.2lcapvdvnri0.webp)

ARP 是介于 L2 和 L3 之间的(所以 ARP 也比较难划分进 OSI 模型中特定的层)，用于将 Layer3 地址映射成 Layer2 地址

- The sender will use ARP to learn the MAC address of the next hop(not necessarily Dst IP of the packet)

  直接将理解成 ARP 是用于学习 next-hop(gateway) MAC 的协议(Dst IP 逻辑上也是 next hop)

- ARP information is stored in a cache so ARP doesn’t have to be done for every single packet

  逻辑上也可想而知，如果没有 cache，每发一个包都要做 ARP request，会对网络造成负担

例如 PC1 10.0.1.10 访问 PC2 10.0.3.10

> ARP request 会被广播

1. 因为 Dst IP 10.0.3.10 和 Src IP 10.0.1.10 不在同一个 subnet 中，所以 PC1 需要将报文发送到 gateway(next hop)，这里对应的是 R1 G0/1 10.0.1.1
2. 因为 PC1 当前 ARP cache 中没有对应 R1 G0/1 的条目，所以 PC1 会做 ARP request 询问 R1 G0/1 的 MAC，以将报文发送到 gateway
3. 当报文到达 R1 G0/1 时因为 IP 匹配，所以会回送 ARP reply 告诉 PC1，10.0.1.1 的 MAC
4. 当 PC1 收到 R1 回送的 ARP reply，会将对应 10.0.1.1 的 MAC 填入到原始报文的 Dst MAC，并发送给 gateway(R1)
5. R1 收到 PC1 发过来的报文，解封装报文，发现目的 IP 10.0.3.10 可以通过路由表找到，但是和任意一个接口的 IP 都不在同一个 subnet，所以 R1 需要将报文发送到 next-hop，这里对应 R2 G0/0 10.0.12.2
6. 因为 R1 当前 ARP cache 中没有对应 R2 G0/0 的条目，所以 R1 会做 ARP request 询问 R2 G0/0 的 MAC，以将报文发送到 next-hop
7. 当报文达到 R2 G0/0 时，因为 IP 匹配，所以会回送 ARP reply 告诉 R1, 10.0.12.2 的地址
8. 当 R1 收到 R2 回送的 ARP reply，会将对应 10.0.12.2 的 MAC 填入到原始报文的 Dst MAC，并发送给 next hop(R2)
9. R2 收到 R1 发送过来的报文，解封装报文，发现目的 IP 10.0.3.10 可通过路由表找到，但是和任意一个接口的 IP 都不在同一个 subnet，所以 R2 需要将报文发送到 next-hop, 这里对应 R3 G0/0 10.0.23.2
10. 因为 R2 当前 ARP cache 中没有对应 R3 G0/0 的条目，所以 R2 会做 ARP request 询问 R3 G0/0 的 MAC，以将报文发送到 next-hop
11. 当报文达到 R3 G0/0 时，因为 IP 匹配，所以会回送 ARP reply 告诉 R2, 10.0.23.2 的地址
12. 当 R2 收到 R3 回送的 ARP reply，会将对应 10.0.23.2 的 MAC 填入到原始报文的 Dst MAC，并发送给 next hop(R3)
13. R3 收到 R2 发送过来的报文，解封装报文，发现目的 IP 10.0.3.10 可通过路由表找到，和 10.0.3.1/24 在一个 subnet，所以 R3 会直接将报文发送到 10.0.3.10
14. 因为 R2 当前 ARP cache 中没有对应 10.0.3.10 的条目，所以 R2 会做 ARP request 询问 10.0.3.10 的 MAC，以将报文发送到 Dst PC3
15. 当报文达到 PC3 时，因为 IP 匹配，所以会回送 ARP reply 告诉 R3, 10.0.3.10 的地址
16. 当 R3 收到 PC3 回送的 ARP reply，会将对应 10.0.33.10 的 MAC 填入到原始报文的 Dst MAC，并发送给 next hop(PC3)
17. 当报文达到 PC3, PC3 就会回送对应 Unicast 响应，因为链路中的设备已经知道对应的 Source/Dst IP 和 MAC 了

## ARP message Format

ARP 报文是直接被封装在 Ethernet header 和 trailer 内的，不包含 IP header(也是这种原因 ARP 被大多数人视为 L2 协议)。如果 Ehthernet 帧中 Type 值为 0x0806 就表示 payload 部分是 ARP message

![](https://github.com/dhay3/image-repo/raw/master/20230825/2023-08-25_10-39.73mra8ao1f40.webp)

ARP message 格式都一样，主要包含如下几个字段

![](https://github.com/dhay3/image-repo/raw/master/20230825/2023-08-25_10-45.4zt89x0i5n0.webp)

> 针对 GARP Operation 字段的值均为 2，因为 GARP 逻辑上就是 ARP reply

### ARP packet

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_16-49.4cbx978b3gu0.webp)

ARP request 中的 Target MAC address 为 0000.0000.0000 表示当前 Target IP 对应的 MAC 未知

## ARP Process

ARP 处理的过程按照 发送 和 回送 分为

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_16-21.5ile0v9hwa00.png)

- ARP Source Host
- ARP Destination Host

### ARP Source Host

1. 发送报文主机，首先会检查自己的 ARP cache，里面是否有对应 Destination Host’s IP(这里指的是 next-hop 的 IP)

   - If there is an entry, no need to proceed with ARP

     也就是不会做 ARP request

   - If there is no entry, create an incomplete entry

     incomplete 意味着 ARP entry 被创建，但是没有对应 IP 的 MAC

2. 生成 Broadcast ARP request

5. 处理 ARP reply，并更新 ARP cache
   - 意味着之前创建的 incomplete ARP entry 会变为 complete

### ARP Destination Host

3. 处理 ARP request, 更新 ARP cache
3. 生成并发送 ARP reply

### Example

例如 R1 192.168.1.1 ping R2 192.168.1.2，如果使用 `debug arp` 就可以看到输出如下内容

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_16-28.27bdh50ei134.webp)

1. R1 会先生成一条 192.168.1.2 incomplete ARP entry
2. 然后 R1 发送 ARP request，日志中的 0000.0000.0000 是 ARP 头中的，表示 192.168.1.2 目前 MAC 未知，和二层头中的 ffff.ffff.ffff 做区别
3. R2 收到 R1 发送过来的 ARP request
4. R2 生成一条对应 192.168.1.1 的 complete ARP entry
5. R2 回送 ARP reply
6. R1 收到 R2 发送过来的 ARP reply，将对应 192.168.1.2 对应的 incomplete ARP entry 改为 complete

## ARP table

在设备上可以使用 `show arp` 来查看 arp table

> 如果是主机，ARP table 只会显示同端内的 IP 对应条目，因为 ARP 只用于同段内

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_17-11.4b71y6gskes.webp)

主要有几个字段

- `Protocol`

  表示 3 层使用的协议

- `Address`

  3 层的 Source 地址

- `Age(min)`

  ARP entry elapsed 的时间，和 OSPF dead timer 一样，收到对应的 ARP message(只有 ARP message 才可以) 后会刷新时间，从 0 开始计时

- `Hardware Addr`

  2 层的 Source 地址

- `Type`

  2 层使用的封包协议，ARPA 即为 Ethernet II

- `Interface`

  对应的 ARP entry 是从那个端口学过来的

  如果对应的 IP 是配置在接口上的，那么 Age 就会显示 `-` 表示永不过期

### Incomplete ARP table entry

例如在 R1 上 ping 192.168.1.3-5，因为当前 192.168.1.0/24 中只存在 192.168.1.1-2，且只有这两个 IP 能收到 ARP request，因为没有对应 IP 的主机能回送 ARP reply 所以对应的 ARP entry 就会显示 incomplete

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_17-20.23bwubeyk3xc.webp)

在 Cisco 中，在没有收到对应的 ARP message 时，默认的 incomplete ARP entry 会在 1 分钟后从 ARP table 中移除，而对应的 complete ARP entry 会在 4 小时后从 ARP table 中移除

> `show arp` 会显示所有的 ARP entries，而 `show ip arp` 只会显示和 IP 协议相关的 ARP entries
>
> 其实在现在的网络环境中，两者并没有多大的区别，因为网络主要还是基于 IP 通信的

## Proxy ARP

> 和正常的 ARP 报文一样，针对 Proxy ARP request 报文使用 operation 1，Proxy ARP reply 报文使用 operation 2

*It allows a device(usually a router) to respond to ARP Requests for IP addresses that are not its own.*

*Modern use cases involve end hosts with incorrect subnet masks, ‘directly connected’ static routes, and some NAT scenarios*

Proxy ARP 是一种特殊的 ARP，可以让非对应 Target IP 的设备回送 ARP reply, 思科的设备默认会开启 Proxy ARP

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_17-31.7l63xe1ztio0.webp)

如果想要手动配置 Proxy ARP 需要使用如下命令

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_17-34.4vqq0r51kgi0.webp)

- `R1(config)#ip arp proxy disable`

  全局关闭 ARP proxy，如果添加 no 表示开启 ARP proxy

- `R1(config-if)#ip proxy-arp`

  当的端口开启 ARP proxy，如果添加 no 表示关闭 ARP proxy

如果想要校验 ARP proxy 是否启用，可以使用 `show ip interface <interface-name>`, 会显示 `Proxy ARP is enabled` 表示已经启用 ARP proxy



==Proxy ARP 默认只会处理从不同 subnets 来的 ARP message，不会处理相同 subnets 中来的 ARP message==

而 local Proxy ARP 可以处理相同 subnets 中来的 ARP message，所以在收到相同 subnets 中来的 ARP message 会优先转发到 router

思科的设备默认关闭 local Proxy ARP，如果想要开启 local Proxy ARP 可以使用 `R1(config-if)# ip local-proxy-arp`

### incorrect subnet mask

当想要通信段的主机配置错误的 subnet，但是现在想要两两通信就可以通过 Proxy ARP 来实现

例如

PC1 192.168.0.11/16 想要访问 PC3 192.168.1.13/24

> 这里的拓扑忽略了 PC1 PC2 之间的交换机，同理 PC3 PC4

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-04_18-22.1fffo7n7pj34.webp)

当 PC1 想要访问 PC3 时，因为 IP 层报文中并不包含 subnet mask，所以会认为 192.168.1.13 和自己在同一个段中，所以会直接发送 ARP request(无需按照正常的逻辑判断在不同 subnet 中，先请求 router 的 ARP 信息)，同理 PC1 访问 PC2/PC4

但是对应的 ARP request 到达 R1 时，R1 并不会将这个报文广播到 192.168.1.13/24 这个段(因为 ARP 广播只用于同段)，所以 PC3 也就收不到对应的报文

但是如果 R1 开启了 Proxy ARP

- R1 在收到 PC1 发送过来的 ARP request 后，因为 192.168.1.0/24 并不是和 G0/0 直联的，而是需要通过 G0/1 转发
- 虽然 192.168.1.13 不是 R1 上配置的任何一个接口的 IP, 但是 R1 有到 192.168.1.13 的路由
- 所以 R1 会针对 PC1 的 ARP reuquest 回送 ARP reply，使用 R1 G0/0 的 MAC(回程使用的接口 MAC)
- 当 PC1 收到 R1 回送的 ARP reply，就会认为 192.168.1.13/24 的 MAC 是 R1 G0/0 的 MAC，将对应的 MAC 记录到 ARP table 中，后续报文 2 层头对应 192.168.1.13 的都会使用 R1 G0/0 MAC

> PC1 自己逻辑上认为是和 192.168.1.13/24 直联的，但是实际需要通过 Router 转发

### directly connected static route

Proxy ARP 还有一个使用的场景，就是当目的包含在静态直联路由中且不是实际上直联时，允许机器正常进行 ARP

例如

在 R1 分别 ping 192.168.12.2/192.168.34.3/192.168.34.4

> `R1(config)#ip route 192.168.34.0 255.255.255.0 g0/0`
>
> 是静态直联路由(directly connected static route)，逻辑上 192.168.34.0/24 是和 R1 直联，但是实际并不是，当访问 192.168.34.4 时会直接尝试往目的发包，而不是 next-hop
>
> 对比 connected route 虽然显示 `is directly connected` 但是会以 `S` 标识，表示是手动配置的静态路由

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_10-26.730p9qxsi7o0.webp)

这里可以观察到 192.168.12.2/192.168.34.3/192.168.34.4 对应的 MAC 条目均为 R2 G0/0 MAC(回程使用的接口 MAC)

> 虽然 192.168.34.0/24 和 192.168.12.0/24 不在一个段内，但是配置了对应的静态直联路由，仍会在 ARP table 中显示对应的条目



如果关闭 R2 G0/0 Proxy ARP，任然保持添加那条静态直联路由，如果这是 R1 ping 192.168.34.3 就会超时，对应 ARP table 中的条目也会显示 incomplete

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_10-33.5gq4zoly9tc0.webp)

## Gratuitous ARP

Gratuitous ARP 免费的 ARP，顾名思义无需收到 ARP request 也能直接发送 ARP reply(Gratuitous ARP 报文 operation 值为 2)

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_10-48.2o0ozfqh3dm0.webp)

有如下几种场景会发送 Gratuitous ARP

- Announcing when an interface is enabled
- Announcing a change in MAC address(ie. if an admin manually configure the MAC)
- Failover between redundant devices(ie. when using an FHRP)

在思科的设备中当 ARP table 有对应 IP 的条目时，如果收到 Gratuitous ARP 设备会更新对应的条目(MAC 和 timer);如果没有对应 IP 的条目时，设备并不会添加 Gratuitous ARP 对应的条目

例如

R1 g0/0 192.168.1.1 up/up 就会发送 Gratuitous ARP

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_10-52.5ugiy86f6ds0.webp)

这里可以观察到 R1 发送的 ARP 报文如下 

Sender IP 192.168.1.1 

Sender MAC R1 g0/0

Target IP 192.168.1.1

Sender MAC ffff.ffff.ffff

SW 会更新 MAC table，R2/3 刷新 ARP table entry(如果已经存在 192.168.1.1 的条目才会)

### Gratuitous ARP packet

Gratuitous ARP 报文主要如下，如果在 wireshark 中还会显示 `Is gratuitous` 这并不是 Gratuitous ARP 报文中字段，而是 wireshark 按照 ARP 报文中 sender/target 地址推断出来的

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_10-54.3brya8vakns0.webp)  

## Manual ARP Entry Configuration

在大多数的场景中并不会手动配置 ARP 条目，但是思科设备是支持的

可以通过 `R1(config)#arp <ip address> <mac address> arpa `(arpa 其实也是一个多选参数，但是在现在的网络中一般只会使用该参数表示 ethernet II)

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_11-00.k9muw7sscgg.webp)

配置完成后使用 `show arp` 查看 ARP table 时，可以发现对应的条目，但是 Interface 处的值是空的

而使用 `show arp <ip> detail` 查看详细信息时，可以发现对应的条目是关联 g0/0 的，因为 router 知道对应 IP 的直联路由，所以知道对应 IP 的 ARP 应该从 g0/0 学来 

## Clearing Dynamic ARP Entry

假设现在使用了 `R1#clear arp` 来删除 R1 上对应的所有 ARP 条目，可以发现，前后竟然没什么区别(timer 会被 refresh)

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_11-06.6mxqerukwtc0.webp)

因为在使用 `clear arp` 命令之后，设备会先发送 unicast ARP request 刷新 ARP table，如果在 3 次内没有收到对应的 ARP reply 才会执行清空对应对的条目

> 注意区别与 `clear mac-address-table dynamic`

## Dynamic ARP aging

和 dynamic MAC address table 一样，ARP table 也有针对每条记录的 timer，来规定条目的过期时间(ARP cache 默认在 4 hours 后过期),但是细微逻辑有别

假设现在使用 `clear arp` 来 refresh ARP table

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_11-24.5qnnt8bzo280.webp)

当 refresh timer 显示为 0， 设备会重新尝试 refresh，如果 refresh 次数超过 3 次，就会清空对应的条目

大概逻辑如下

```
if refresh timer reaches 0
	reponse = unicast arp request
	if no reponse
		reponse = unicast arp request
		chance = 1
		if no reponse
			response = unicast arp request
			chance = 0
			clear entry
		else
    	reset refresh timer to 4 hours and 2 chances
  else
  	reset refresh timer to 4 hours and 2 chances
else
	reset refresh timer to 4 hours and 2 chances
```

可以发现前后两次 refresh 显示 refresh 的时间是不同的。主要原因如下

*This is to prevent all ARP entries from expiring at the same time, triggering an ‘ARP’ storm*

通过如下方式来防止 ARP storm

Cisco IOS devices add a random jitter between 0 seconds and 30 minutes to the timeout counter of each ARP entry

> 这个逻辑和 redis 的缓存雪崩类似
>
> jitter 按照 aging timer 大小来选择，如果值大 jitter 就大，反之则小

当然 ARP aging 也是可以手动配置的，但是不推荐，可以通过 `R1(config-if)#arp timeout <seconds>` 来实现

![](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_11-34.179m9wx2kjxc.webp)

> 这里为什么 configured timeout + jitter 小于 3 mintus 没明白

## Command Review

![	](https://github.com/dhay3/image-repo/raw/master/20230904/2023-09-05_11-40.47dwfhxb8iw0.webp)



**references**

1. ^https://www.youtube.com/watch?v=k3oda32jmWY&list=PLxbwE86jKRgOb2uny1CYEzyRy_mc-lE39&index=9