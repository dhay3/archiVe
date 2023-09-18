# Day01 - Layer2 Forwarding

## Review

### OSI Model

先回顾一下 OSI 模型

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230821/2023-08-21_22-52.4ipceywzosy0.png)

对比 TCP/IP 模型，以及常用的模型

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230821/2023-08-21_22-54.5faqnl51jqw0.webp)

### Encapsulation/De-Encapsulation

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230821/2023-08-21_22-57.6cr5h5dwtwo0.webp)

报文从上层到下层叫做 encapsulation(封装)

> 4 层报文叫做 segment
>
> 3 层报文叫做 packet
>
> 2 层报文叫做 frame

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230821/2023-08-21_22-59.3r57z5xq1ko0.webp)

报文从下层到上层叫做 de-encapsulation(解封装)

### Collision Domains

- 半双工是一种通信模型，要求同一时间网络上只可以有一个发送者，其他的设备只能是接受者，不能是发送者

- 如果同时存在多个发送者，就会出现冲突(这里的逻辑无需深入，将其看成一个定理即可)，即 Collision Domain

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_11-08.7ep8p9k1t740.webp)

### Hub

- Hub 并不是一个 2 层设备，因为 Hub 并不会解析 2 层 MAC，只是做报文转发
- Hub 没有 buffer，Hub 收到报文后会立即转发到除接收端口外的其他所有端口。因为这种特性就要求 Hub 工作在半双工的环境下，以减少冲突的几率(同样还是有概率冲突的)
- 所有通过 Hub 互联的设备构成的网络是一个冲突域

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_11-39.5652zhj4axs0.webp)

### Switch

- Switch 是一个 2 层设备，会根据报文 MAC 进行转发
- 和 Hub 不同，Switch 每一个端口对应的链路都是一个 Collision Domain

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_11-22.fp3nnoc9vmg.webp)

Switch 有 buffer，所以可以决定报文的发送顺序，减小冲突

1. 当两台 PC 同时向 Switch 发送广播帧时
2. Switch 不会同时将收到的广播帧转发
3. Switch 会将其中一台机器的广播帧缓存，先传输一台机器的广播帧。当 CSMA/CD 监听到当前网络中空闲后，在转发另外一台机器的广播帧

所以同一时间网络中可以同时存在多个发送者，称为全双工。Switch 可以工作在全双工下

### quiz

例如下图中一共有 9 个冲突域

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_13-56.21jilsertg2o.webp)

### Broadcast Domains

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_13-59.2mcfcw3il3o.webp)

> 通常我们并不需要注意 Collision domain，因为 Switch 会为我们划分 Collision domain
>
> 但是我们需要考虑 Broadcast domain 的范围，以减少报文广播的范围(通常使用 VLAN 实现)

### quiz

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_16-16.6n1h3wxjzas0.webp)

## Layer 2 Forwarding

*Layer 2 Forwarding refers to thee process switches use to forward frames within a LAN.*

> 虽然 routers 是 3 层设备，但是 router 任然会注意以及使用 2 层 MAC

Layer 2 Forwarding 可以划分成 4 大类

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_14-10.4hy2j6jr7b40.webp)

- known Unicast

  一对一，forward

- unknown Unicast

  一对所有，flooding

- Broadcast

  一对所有，flooding

- Multicast

  一对多，flooding

### unknown Unicast

例如 R1 访问 PC1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_14-42.2h1pj6vt0a20.webp)

1. R1 发送报文 Src MAC aaaa.aaaa.aaaa Dst MAC 1111.1111.1111

2. 当报文到达 SW1 时 de-encapsulation，因为 SW1 MAC address table 中还没有对应的 aaaa.aaaa.aaaa 的条目，所以会将对应的条目记录到 MAC address table 中

   > 这也被称为 Dynamically learning

3. 但是现在 SW1 中没有对应 Dst MAC 1111.1111.1111 的条目，所以不知道往那发送这个报文，所以被转发到除接受报文端口 G0/0 外的其他所有端口(Flood)。这也是被称为 Unkown Unicast 的原因

4. 当 PC2/PC3 收到报文后，de-encapsulation 报文，因为目的 MAC 不匹配所以丢弃

5. 当 PC1 收到报文后，de-encapsulation 报文，因为目的 MAC 匹配，所以接受

### Known Unicast

例如 PC1 给 R1 回包

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-06.5ql3ndgzvf80.webp)

1. PC1 回送 Src MAC 1111.1111.1111 Dst MAC aaaa.aaaa.aaaa

2. 当报文到 SW1 时 de-encapsulation，因为 SW1 MAC address table 中还没有对应的 1111.1111.1111 的条目，所以会将对应的条目记录到 MAC address table 中

3. 对应的 Dst MAC aaaa.aaaa.aaaa 已经在 SW1 MAC address table 中，所以 SW1 不会广播报文，而是直接转发到 G0/0，即 Knwon Unicast

   > Known Unicast 和 Unicast 本质上没有区分，可以归为一类

### Broadcast

例如 PC2 发送广播帧

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-15.141tpaznxkm8.webp)

1. PC2 发送 Src MAC 2222.2222.2222 Dst MAC ffff.ffff.ffff

   > 目的 MAC 为 ffff.ffff.ffff 的报文也被称为广播帧

2. SW1 收到报文时，会 de-encapsulation，因为 Src MAC 2222.2222.2222 不在 MAC address table 中，所以会将其记录，同时发现 Dst MAC 是 ffff.ffff.ffff 对应二层广播地址，所以会转发到除接受端口外的其他所有端口

### Forwarding

例如 PC3 访问 PC2

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-23.1hlf2sxnppmo.webp)

1. PC3 发送 Src MAC 3333.3333.3333 Dst MAC 2222.2222.2222
2. SW1 收到报文时，会 de-encapsulation，因为 Src MAC 3333.3333.3333 不在 MAC address table 中，所以会将其记录，同时发现 Dst MAC 是 2222.2222.2222 在 MAC address table 中有记录，所以直接转发到 G0/2

### Multicast

> 多播只需记住默认 flood 即可，会在后面的章节详细介绍

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-36.y07wn5dsuds.webp)

## MAC Address Table

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-42.11c8wfu9hsps.webp)

- `Type`

  STATIC 并不是表示对应的 MAC 是手动配置的，而是 Switch 默认自带的。常见的对应多播和广播地址

- `Port`

  CPU 表示当 Switch 收到对应的 MAC 地址时，是直接转发到 CPU 处理的。常见的对应多播和广播地址

### MAC Address aging

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-49.6nr69ydmsbg0.webp)

MAC Address 默认并不会永久的保存在 MAC Address table 中

默认如果没有收到对应地址的报文在 5 mins 后会被清除，如果期间收到报文，就会重新从 5 mins 开始计时

可以使用 `SW1#show mac address-table aging-time` 来查看配置的 aging time，如果想要配置 aging time 可以使用 `SW1(config)#mac address-table aging-time <secs>`

如果 `<secs>` 为 0 表示永不过期(通常没有意义)

### MAC Address learning

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-53.429ru00f86c0.webp)

还可以通过 `SW1(config)#no mac address-table learning vlan <vlan-id>` 来关闭指定 VLAN 的 dynamically learning

> 通常不会使用该命令，即使收到 MAC address flooding

可以通过 `SW1#show mac address-table learning` 来查看那些 VLAN 开启了 dynamically learning 的功能

### MAC Address static configuration

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_15-57.6f4ptocbks00.webp)

当然 MAC Address 也可以和 Route 一样，配置成 Static

可以通过 `SW1(config)#mac address-table static <mac> vlan <vlan-id> interface <interfance-name>` 来配置当指定 VLAN 收到指定 MAC 时，通过指定端口转发

还可以使用 `SW1(config)#mac address-table static <mac> vlan <vlan-id> drop` 来配置当指定 VLAN 收到指定 MAC 时，直接丢弃

配置完成后 `show mac address-table` 可以发现多了 2 条对应的条目

### Mac Address Table count

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_16-10.4cwp6rddtfy0.webp)

可以使用 `show mac address-table count` 来显示 mac address table 中一共的条目数

> 这里只显示了 2 条
>
> 因为该命令不会计算 Switch 中默认的 MAC address(通常是多播和广播地址)

### Clearing dynamic MAC Address

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_16-06.13i17tqmet34.webp)

除了 aging time 后自动删除 MAC address，还可以通过 `SW1#clear mac address-table dynamic` 来清空所有 dynamically learning 学来的 MAC address

> 当然也可以指定 vlan/interface/address

## Command Summary

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230823/2023-08-23_16-14.z2bpgluocsw.webp)



**references**

1. ^https://www.youtube.com/watch?v=VNMyGOA_LoY&list=PLxbwE86jKRgOb2uny1CYEzyRy_mc-lE39&index=5