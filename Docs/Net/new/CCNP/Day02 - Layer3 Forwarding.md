# Day02 - Layer3 Forwarding

## Bitwise XOR and AND

*How does a network host(ie. your PC) determine if a packet should be forwarded to the default gateway, or if it can be forwarded directly to the destination host*

- Source and Destination in same subnet = forward directly

  > 意味着在同 LAN 中

- Source and Destination in differen subnets = forward to default gateway

首先会使用 XOR(异或，相同为 0 不同为 1)来计算 Source 和 Destination

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_20-29.6q4n2vc620g0.webp)

例如 192.168.1.10/24 访问 192.168.1.99

192.168.1.10 XOR 192.168.1.99 结果为 0.0.0.105



然后使用 AND(与，相同为 1 不同为 0)来计算 XOR 的结果和 Subnetmask

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_20-35.3mtufm37ah00.webp)

例如 192.168.1.10/24 访问 192.168.1.99

0.0.0.105 AND 255.255.255.0，结果为 0.0.0.0

当 4 组 octect 都为 0 时，表示 Source 和 Destination 在相同的 subnet 中，如果其中一组不为 0 就表示在不同的 subnet 中，需要将报文发送到 default gateway

用代码的逻辑表示

```
s xor d and sub == 0.0.0.0 ？ same subnet : not same subnet
```

例子

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_20-45.3rnjai7scda0.webp)

更推荐直接将 Source 和 Destination 都和 subnet mask 做与运算即可，如果结果相同表示在同 subnet，反之不在同 subnet

代码的逻辑为

```
s and sub == d and sub ？ same subnet : not same subnet
```

## Forwarding IP Packets within a LAN

当设备判断 Source 和 Destination 在同 Subnet 中

*The Packet will be encapsulated in an Ethernet frame, and the destination MAC address will be the destination host’s MAC address*

> 目的 MAC 是通过 ARP 学习而来的

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_20-49.4z7dm6dnqww0.webp)

例如现在 PC1 10.0.0.11 想要访问 PC3 10.0.0.13

PC1 首先会查找自己的 ARP table，如果有对应 10.0.0.13 的条目，Dst MAC 就会使用对应条目中的 Physical Address(即 PC3 MAC)

如果没有对应的条目，就会通过 ARP request/reply 学习对应的条目

## Routing IP Packet between Networks

当设置判断 Source 和 Destination 不在同 Subnet 中

- The source host will send the packets to its default gateway

  > 3 层地址不变，2 层 Dst MAC 变为 default gateway MAC

- The router receives the frame/packet, it will change the destination MAC address to that of the next-hop router, and the source MAC address to that of its sending interface

  > 3 层地址始终是不变的，变的只有 2 层地址

- The final router in the path receives the frame/packet, it will change the destination MAC address to that of the destination host and source MAC address to that of its sending interface

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_20-59.14ae3bhdhwdc.webp)

例如 PC1 访问 PC2，subnet mask 255.255.255.0 所以不在一个 subnet 中

1. 因为不在一个 subnet，所以 PC1 发送报文对应的 2 层 Dst MAC 为 default gateway 的(通过 ARP 学习而来)
2. 当 R1 收到报文后，计算发现还是不在同 subnet 中，所以按照路由表通过指定端口转发报文到 next-hop，2 层 Src MAC 为路由表中对应的转发口，2 层 Dst MAC 为 next hop MAC
3. 当 R2 收到报文后，计算发现在同 subnet 中，所以按照路由表通过指定端口转发报文到 Dst, 2 层 Src MAC 为路由表中对应的转发口，2 层 Dst MAC 为对应的目标设备的 MAC

## Layer 2 vs Layer 3 Forwarding

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_21-18.1sexxe1zfuzk.webp)

### Layer 3 Forwarding Decisions

这里扩展 layer 3 forwarding 中 most specific 含义

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_21-24.u9ir1x1cdb4.webp)

*When it comes to making a Layer 3 forwarding decision, the matching route with the longest prefix length wins*

即只按照最长 prefix 来匹配

> 那么之前学的 AD 和 metric 有啥用呢？
>
> AD 和 metric 只是用于决定那条 route 应该加入到 routing table，并不会影响 route decision

## IPv4 Address

### Configuring IPv4 Address

Cisco 的设备还可以配置 secondary IP

> 使用场景可能比较少，通常用于网段迁移，会在后面详细讲解
>
> 逻辑上理解成类似 router-on-a-stick 中的子接口，但是这个 IP 是直接配置在物理接口上

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_21-32.72lcaoc1kq40.webp)

可以通过 `R1(config)#ip address <ip> <subnet mask> secondary` 来配置 secondary IP

*An interface can have one primary IPv4 address and an unlimited number of secondary IPv4 addresses*

> secondary 翻译次接口比较合理，而不是第二个接口，因为接口上可以配置多个 secondary IP，并不会直接 override

### Verifying IPv4 address

![](https://github.com/dhay3/image-repo/raw/master/20230823/2023-08-23_21-44.33il9zwe45k0.webp)

- `IP-Address`

  只会显示 primary IP address， 不会显示 secondary IP address

- `Method`

  - `manual`

    表示当前 IP address 有会修改过

  - `unset`

    表示 IP address 没有配置过，还未使用 `startup-config`(即一次都没有使用 `write` 之类的命令保存 running config)

  - `NVRAM`

    表示当前的 IP address 是通过读取 `startup-config` 来的

    > 从 `manual` 到 `NVRAM` 需要设备保存 running config 并重启

虽然 `show ip int br` 不会显示 secondary address，但是可以通过 `show ip interface <interface-name>` 来查看

![](https://github.com/dhay3/image-repo/raw/master/20230824/2023-08-24_20-09.3sx0y0ydnvs0.webp)

## Directed Broadcast

在使用 `show ip interface <interface-name>` 后，会显示一行

```
Directed broadcast forwarding is disabled
```

什么是 Directed Broadcast 呢？

*A directed broadcast message is a message sent to the broadcast address of another subnet(not the subnet of the sending host)*

- Routers in the path will forward these directed broadcat packets as normal packets

  > 因为报文中的 Dst Address 只包含 IP 地址，不包含 subnet mask，所以 Router 并不知道这个地址是一个广播地址，所以会 Unicast 这个广播报文

- When the router connected to the destination subnet receives the message, it will know that the destination is a broadcast address(because it knowns the subnet mask of the destination network)

当和目的地址互联的 router 收到这个广播的报文，router 并不会将其广播到 network 中其他的所有 hosts(即理解成只有当前这个互联的 router 会回送报文，其他 hosts 并不会回送报文)。如果需要 router 正常处理该报文需要使用 `R1（config-if）#ip directed-broadcast`

例如

![](https://github.com/dhay3/image-repo/raw/master/20230824/2023-08-24_20-24.qkzkvbhj4qo.webp)

1. 首先通过 `R1#debug ip icmp` 进入 ICMP debugging 模式

2. R1 ping 192.168.34.4(R4 G0/0)，可以看到 192.168.34.4 是正常回包的

3. 现在 R1 ping 192.168.34.255 是 192.168.34.0/24 的广播地址，这里可以看到可以正常 ping 通，但是 source address 是 192.168.23.3(R3 G0/0) 而不是 192.168.34.255

   > 这里 source address 为什么是 192.168.23.3 而不是 192.168.34.3?
   >
   > 可以理解成实际是 192.168.34.3 回的包，和 traceroute 类似，报文在经过 R3 G0/0 后，source address 会变为 192.168.23.3

现在使用 `R1(config)#ip direceted-broadcast` 开启 directed-broadcast

> 需要在发送广播帧的出接口上使用该命令

![](https://github.com/dhay3/image-repo/raw/master/20230824/2023-08-24_20-47.3ob0knvrai00.webp)

> 报文的 Dst MAC 如上图

R1 ping 192.168.34.255，这里可以看到 192.168.23.3 和 192.168.34.4 都回包了

- 192.168.23.3 回包和上面的一样，实际是 192.168.34.3 回的包，经过 R3 G0/0 source address 变为 192.168.23.3

- 192.168.34.4 回包因为在 192.168.34.0/24 范围内，所以收到广播帧后回送报文

  > 如果 192.168.34.0/24 还有其他的 host，同样也会回送，因为收到的报文目的地址是一个 broadcast address

**references**

1. ^https://www.youtube.com/watch?v=3jCd2T6R6h8&list=PLxbwE86jKRgOb2uny1CYEzyRy_mc-lE39&index=7