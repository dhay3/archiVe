# Day17 - VLANs part2

以如下拓扑为例

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-03.6vql9jubyqrk.webp)

Engineering 部门划分在 VLAN10，HR 部门划分在 VLAN20,Sales 部门划分在 VLAN30，SW1 和 SW2 之间通过 2 个 access port 建立连接，橙色的是 VLAN10,紫色的是 VLAN30

> There is no link in VLAN20 between SW1 and SW2. This is because there are no PCs in VLAN20 connected to SW1，PCs in VLAN 20 can still reach PCs connected to SW1, R1 will perform **inter-VLAN routing**

如果 VLAN 间想相互通信，就必须要通过 router

这里 VLAN 30 访问 VLAN 10 同样需要经过 router，只有 VLAN 10 访问另外一个 VLAN 10 不需要经过 router



## What is a trunk port

在小型的网络拓扑中只需要用几个 VLAN 对应几个 access port 来连接 switches 或者 routers。但是在大型的网络拓扑中如果一个接口分配给一个 VLAN，明显太浪费接口了，而且通常 router 的接口都比较少，还不够 VLAN 用

所以就需要使用 trunk ports

> you can use trunk ports to carry traffic from multiple VLANs over a single interface

VLAN 10 中的一台机器，想要访问另外一个 VLAN 10 中的一台机器

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-12.46574qsnoy80.webp)

现在有一个问题

*how does SW1 know which VLAN the traffic belongs to? Both VLAN30 and VLAN10 are allowed on the interfaces the traffic was received on SW1*

答案是 VLAN tagging

*Switch will tag all frames that they send **over a trunk link(这里注意只有出 trunk port 或者是路由器的子接口会打上 VLAN tag，access port 或者是 入 trunk port 都不会打上 VLAN tagging)**. This allows the receiving switch to known which VLAN the frame belongs to*

因为这个原因通常有如下几个别名

- Trunk port = tagged port
- acess port = untagged port

## VLAN tagging

有两种主要的 trunking protocols

- ISL(Inter-Switch Link) 是思科独有的协议，现在已过时
- IEEE 802.1q(dot1q) 是标准公认的协议

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-22.2b9unrcos20w.webp)

dot1q 会在原始的 2 层帧头 Source 和 Type/Length 之间添加一个 802.1Q 的字段

- The dot1q tag is inserted between the source and type/length fields of the Ethernet frame
- The tag is 4 bytes(32bits) in length
- the tag consists of two main fields
  - Tag protocol indentifier(TPID)
  - Tag control information(TCI)，consists of three sub-fields

### dot1q tag

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-27.56wy4u6e9a0w.webp)

#### TPID

- 16 bits(2bytes) in length
- Always set to a value of 0x8100. This indicates that the frame is 802.1q-tagged

主要就是表示当前帧是一个 dot1q 帧

#### TCI

##### PCP

Priority Code Point

- 3 bits in length
- Used for Class of Service(Cos), which priorities important traffic in congested networks

##### DEI

Drop Eligible Indicator

- 1 bits in length
- Used to indicate frames that can be dropped if the network is congested

##### VID

VLAN ID

- 12 bit in length，所以最多只有 4096 个 VLAN (0 - 4095，2 x 1024)，但是 0 和 4095 是保留的 VLAN，不能被直接使用，所以实际可以使用的是 1 - 4094
- Identifies the VLAN the frame belongs to

### VLAN range

The range of VLANs(1-4094) is divided into two sections

1. Normal VLANs: 1 - 1005
2. Extended VLANs: 1006 - 4094

> 大多数交换机都支持这两个部分，但是一些比较老的设备不支持 Extended VLANs

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-41.1j61a3220tvk.webp)

当 SW2 接受到 VLAN10 内的机器发送的报文后；会根据是否是广播报文，选择转发或者是广播到 VLAN10 接口，出向 trunk 接口会加上 VLAN10 tag

## Native VLAN

有一个比较特殊的 VLAN，被称为 Native VLAN

*Switch does not add an dot1q tag to frames in the native VLAN*

所有的 trunk port 默认都以 1 做为 Native VLAN，但是也可以手动配置（出于安全考虑，一般会将 Native VLAN 设置成除 1 外的 vlan-id）

*When a switch receives an untagged frame on a trunk port, it assumes the frame belongs to the native VLAN. **It is vary important that the native VLAN matches***

假设现在 SW1 和 SW2 互联的 trunk port 配置的 Native VLAN 都是 10

现在 VLAN10 的一台机器想和 VLAN10 的另外一台机器通信，当报文到达 SW2 时因为是 VLAN10 和 Native VLAN 相同，所以 SW2 不会对出向的报文加上 VLAN10 tag；当报文到达 SW1 时，因为收到的报文没有 VLAN tag，所以会认为是 Native VLAN，所以会直接转包到 VLAN10 内的另外一台机器或者是广播

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-53.2hjmnzxitxog.webp)

现在将 SW1 的 Native VLAN 配置成 30，当报文到达 SW1 时因为没有 VLAN tag，所以会认为是 Native VLAN，所以会转发到 VLAN30，但是 VLAN30 中并不包含实际想要访问 VLAN 10 中的机器，所以显然就不会转发直接丢包

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_16-59.2e1at7art94w.webp)

## Trunk Configuration

- `switchport mode trunk`

  将当前端口转换为 trunk port，但是如果交换机支持 ISL 和 dot1q 两种 VLAN 协议，直接使用可能会有问题

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_17-06.183crtufqmbk.webp)

  需要使用 `switchport trunk encapsulation dot1q` 来指定使用 dot1q 协议，才可以使用 `switchport mode trunk`

- `show interface trunk`

  来查看所有的 trunk port

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_17-10.87wflnuubg0.webp)

  - mode on 表示端口是手动配置的 trunk port
  - Encapsulation 802.1q 表示端口使用 dot1q 协议
  - Native Vlan 1 表示端口使用的 Native VLAN
  - Vlans allowed on trunk 1 - 4094 表示 trunk port 支持的 VLAN，1 - 4094 表示所有的 VLAN
  - Vlans allowed and active in management domain 1,10,30 当前允许且生效的 VLAN，在这个例子中因为允许所有的 VLAN，所以只要增加一个 VLAN 就会出现在该列中

- `switchport trunk allowed vlan <vlan-id>`

  配置当前 trunk 接口允许通过的 VLAN

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_17-17.13kw0ehligu8.webp)

- `switchport trunk allowed vlan add|remove <vlan-id>`

  对当前 trunk 接口允许的 VLAN 做增删

- `switchport trunk native vlan <vlan-id>`

  修改当前 trunk 接口的 Native VLAN

## Router on a stick

你可能已经注意到了上面的例子中 R1 和 SW2 是通过 1 个接口互联的。如果需要使用这种拓扑，就需要使用 Router on a stick(ROAS)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_17-33.238grvqumigw.webp)

将和 SW2 互联的 R1 G0/0 分成 3 个 subinterfaces(子接口)，subinterface number 可以不用和 VLAN id 相同，但是为了方便阅读一般设置成和 VLAN id 相同

- G0/0.10 对应 VLAN10
- G0/0.20 对应 VLAN20
- G0/0.30 对应 VLAN30

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_17-40.6ztw54yvolfk.webp)

当然也需要为每一个 子接口 单独配置一个 3 层的 IP 地址

如果收到的 2 层帧包含 VLAN 10 就会从 interface g0/0.10 入，如果 2 层帧需要发送到 VLAN 10 就会从 interface g0/0.20 出，同理 g0/0.20, g0/0.30

假设 VLAN 10 中的一台机器需要访问 VLAN 30 中的一台机器

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_17-54.2scnfm6b654.webp)

1. 因为不在一个段,首先会做 ARP request 找 GW MAC
2. ARP request 到达 SW2 G0/2 因为入口是 VLAN10，所以 SW2 只会广播到 VLAN10 中所有的设备，同时记录入向 MAC 到自己的 MAC address table,因为 SW G0/1 是一个 trunk port,通过该端口发送的会带上 VLAN tag
3. 因为收到的 ARP request 有 VLAN10 tag，所以会从逻辑子接口 G0/0.10 入，回 ARP reply，到 SW2，并记录对应的入向 MAC 到自己的 MAC address table，以及 ARP table
4. ARP reply 到 SW2，unicast 到 VLAN10 中发报文的机器
5. VLAN10 发报文的机器通过 SW2 转发到 R1，查路由(这里省略 R1 ARP request,实际需要做)，转发到 G0/0.30 并打上 VLAN30 tag 到 SW2
6. SW2 de-encapsulation，查 MAC 并转发到 SW1，同时打上 VLAN30 tag
7. SW1 因为收到的报文有 VLAN30，所以识别成 VLAN30，并转发对应的 VLAN30 access port

### Summary

- ROAS is used to route between multiple VLANs using a single interface on the router and the switch
- The switch interface is configured as a regular trunk
- The router interface is configured using subinterfaces. You configure the VLAN tag and IP address on each subinterface
- The router will behave as if frames arriving with a certain VLAN tag have arrived on the subinterface configured with that VLAN tag
- The router will tag frames sent out of each subinterface with the VLAN tag configured on the interface

> router 的逻辑子接口和 trunk port 一样，只有出向的才会带上 VLAN tag

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-29_20-05.v16t6oil140.webp)

这里需要注意 SW1 和 SW2 之间互联的 trunk port，并不需要 VLAN20

假设 VLAN30 的机器访问 VLAN20 的机器(不考虑广播的情况)

1. VLAN30 的报文到 SW1, 因为入口是 VLAN30 access port，de-encapsulation 2 层,查 MAC address table 转发到 G0/1,并加上 VLAN30 tag
2. SW2 收到报文后，因为有 VLAN30 tag，de-encapsulation 2 层, 查 MAC address table 转发到 G0/2，并加上 VLAN30 tag
3. R1 收到报文后，因为有 VLAN30 tag, de-encapsulation 3 层,查 routing table 转发到 G0/0.20(VLAN20)，2 层加上 VLAN20 tag
4. SW2 收到报文后，因为有 VLAN20 tag，de-encapsulation 2 层，查 MAC address table 转发到 F0/1 

但是配置了 VLAN20 同样也可以，只是 SW2 VLAN20 广播的时候，会多做一个 G0/1 口

### SW1

```
SW1>en
SW1#conf t
Enter configuration commands, one per line.  End with CNTL/Z. 
SW1(config-if-range)#int ran f0/1,f0/2,f0/3,f0/4
SW1(config-if-range)#switchport mode access 
SW1(config-if)#int range f0/1,f0/2
SW1(config-if-range)#switchport access vlan 10
SW1(config-if)#int range f0/3,f0/4
SW1(config-if-range)#switchport access vlan 30
SW1(config-if-range)#int g0/1
SW1(config-if)#switchport mode trunk
SW1(config-if)#switchport trunk native vlan 101
SW1(config-if)#switchport trunk allowed vlan 10,30
```

### SW2

> 注意这里需要给 SW2 创建一个 VLAN 30，否则 VLAN30 的报文不能正常到 SW2

```
SW2>en
SW2#conf t
Enter configuration commands, one per line.  End with CNTL/Z. 
SW2(config-if-range)#int ran f0/1,f0/2,f0/3,f0/4
SW2(config-if-range)#switchport mode access 
SW2(config-if)#int f0/1
SW2(config-if)#switchport access vlan 20
SW2(config-if)#int range f0/2,f0/3
SW2(config-if-range)#switchport access vlan 10
SW2(config-if-range)#vlan 30
SW2(config-if)#int g0/1
SW2(config-if)#switchport mode trunk
SW2(config-if)#switchport trunk native vlan 101
SW2(config-if)#switchport trunk allowed vlan 10,30
SW2(config-if-range)#int g0/2
SW2(config-if)#switchport mode trunk
SW2(config-if)#switchport trunk native vlan 101
SW2(config-if)#switchport trunk allowed vlan 10,20,30
```

### R1

```
R1>en
R1#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
R1(config)#int g0/0
R1(config-if)#no shutdown
%LINK-5-CHANGED: Interface GigabitEthernet0/0, changed state to up
%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet0/0, changed state to up
R1(config-if)#interface g0/0.10
%LINK-5-CHANGED: Interface GigabitEthernet0/0.10, changed state to up
%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet0/0.10, changed state to up
R1(config-subif)#encapsulation dot1Q 10
R1(config-subif)#ip add 10.0.0.62 255.255.255.192
R1(config-subif)#encapsulation dot1Q 20
R1(config-subif)#ip add 10.0.0.126 255.255.255.192
R1(config-subif)#encapsulation dot1Q 30
R1(config-subif)#ip add 10.0.0.192 255.255.255.192
```

**referenes**

[^jeremy’s IT Lab]: https://www.youtube.com/watch?v=aHwAm8GYbn8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=19
