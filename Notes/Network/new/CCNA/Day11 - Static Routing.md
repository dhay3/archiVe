# Day11 - Static Routing

如果需要往不是直接互联的设备发包，只使用 connected route 和 local route 显然是不够的。这里主要说实现的方式中的一种 —— static route

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_17-31.17ai6a89sq3k.webp)

以 R2 为例

在配置如上图的 IP address 后，会自动添加 2 对 connected route 和 local route

显然 R2 可以直接访问 192.168.12.0/24, 192.168.24.0/24

但是不能直接访问 192.168.1.0/24, 192.168.13.0/24, 192.168.34.0/24, 192.168.4.0/24

同理 R3

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_17-35.4awf0a5v2a2.webp)

R3 可以直接访问 192.168.13.0/24, 192.168.34.0/24

但是不能直接访问 192.168.1.0/24, 192.168.12.0/24, 192.168.24.0/24, 192.168.4.0/24

同理 R4

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_17-37.4qorek92s54w.webp)

R4 可以直接访问 192.168.24.0/24, 192.168.34.0/24, 192.168.4.0/24

但是不能直接访问 192.168.1.0/24, 192.168.13.0/24, 192.168.12.0/24

## Static Route Configuration

现在 PC1 需要访问 PC4，不考虑 ARP(即 GW MAC 地址已知)

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_17-54.gl3olgooz28.webp)

那么 PC1 发出的报文会包含如下信息

- Src IP 192.168.1.10
- Dst IP 192.168.4.10
- Src MAC PC1 eth0 MAC
- Dst MAC R1 G0/2 MAC

当报文到达 R1 时，R1 会 de-encapsulate it (remove L2 header/tailer)， 并检查 routing table 中是否含有匹配 Dst IP 的路由

因为 R1 没有匹配的路由，所以 R1 会将报文丢弃

如果 R1 需要转发报文，就需要有到 192.168.4.0/24 的路由

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_18-00.66849j896v0g.webp)

PC1 和 PC4 需要通信，光有 Dst IP 的路由是不够的，还需要有 Src IP 的路由

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_18-04.6emsoo9y1yf4.webp)

路由必须要是 **two-way reachability**

如果需要配置 static route 需要使用 `ip route <ip-address> <netmask> <next-hop>` 命令

R1 配置如下

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_18-07.mj9exiwt7kg.webp)

因为 R3 和 192.168.1.0/24 以及 192.168.4.0/24 不是直联的所以需要添加 2 条路由

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_18-10.capl4wmv6ns.webp)

R4 配置如下

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_18-10_1.3fg39vn67wzk.webp)

配置完上述内容，PC1 报文到 PC4 的路径如下图 

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_18-12.4ude4ctqy0e8.webp)

3 层 IP address 是不是会改变的，2 层 MAC address 会根据 recive interface 和 next-hop interface 而改变 

当然除了使用 `ip route <ip-address> <netmask> <next-hop>` 来添加静态路由，还可以使用 `ip route <ip-address> <netmask> <exit-interface>` 来指定使用对应的出端口。还可以同时指定 next-hop 和 exit-interface

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_14-20.1tlveytj1leo.webp)

如果指定了 exit-interface 和 `ip route <ip-address> <netmask> <next-hop>` 有区别。路由会被标识成 directly connected，**但是实际上并不是直接互联的**

## Default Gateway/Route

End hosts PC1 和 PC4 可以直接发送 packets 到直联的 network。对于 PC1 来说是 192.168.1.0/24，对于 PC4 来说是 192.168.4.0/24。如果需要发送到不在直联的 network 就需要通过 default gateway,对应的路由也被称为 default route

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_17-44.1hjxwka0k1q8.webp)

*Default route is a route to 0.0.0.0/0 = all netmask bits set 0. Includes all address from 0.0.0.0 to 255.255.255.255*

因为 0.0.0.0/0 含有所有的 IPv4 address，所以 specific 最低

可以使用 `ip route 0.0.0.0 0.0.0.0 <next-hop>` 来配置 default route

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-25_14-25.5qmvg8r9c45c.webp)

candidate default 表示是一条候选的 default route，在 Cisco 中可以有多条 default route

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=aHwAm8GYbn8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=19

