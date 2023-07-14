# Day16 - VLANs part 1

## What is a LAN

在了解 VLAN 前，需要先回顾一下 LAN

*A LAN is a group of devices(PCs, servers, routers, switches, etc.) in a single location(home, office, etc.)*

- A LAN is a single broadcast domain, including all devices in that broadcast domain
- A broadcast domain is the group of devices which will receive a broadcast frame (destination MAC FFFF.FFFF.FFFF) sent by any one of the members

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-16.499n6m8xki80.webp)

例如上图就有 4 个 boardcast domain(也对应这 4 个 LAN)

可以明显看出 boardcast domain 是通过 Router 来划分的。因为 Router 在收到 Dst MAC FFFF.FFFF.FFFF 的报文时，并不会和 Switch 一样转发到除接受端口外的所有端口

## Why needs VLAN

假设有 3 个部门在 192.1681.0/24 内，Engineering 内的一台 PC 需要发包到 engineering 内的另外一台机器，在不知道对方的 MAC address 的情况下，就需要做 boardcast

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-22.r1tpyw42clc.webp)

因为 3 个部分都在一个 boardcast domain(LAN) 中，除 engineering 部门外，sales 和 HR 部门都能收到发送的广播包。有会有两个问题

1. 效率和性能

   Lots of unnecessary broadcast traffic can reduce network performance

2. 安全

   Because this is one LAN, PCs can reach each other directly, **without traffic passing through the router**. So, even if you configure security policies, the won’t have any effect(交换机没有安全策略)

那么现在把 192.168.1.0/24 划分成 3 个子网

192.168.1.0/26, 192.168.1.64/26, 192.168.1.128/26

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-35.4my9c6yhytxc.webp)

这样也就需要 3 条 cable 和 Switch 互联(实际可以只使用一条，单臂路由)

假设现在PC1 在不需要做 ARP request 的情况下(即不需要 boardcast/unkown unicast，只有 unicast )，PC1 192.168.1.1 访问 PC2 192.168.1.128

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-39.6m8x53se3j0g.webp)

1. 报文在 PC1

   Src IP: 192.168.1.1

   Dst IP: 192.168.1.129

   Dst MAC: R1 MAC

   Src MAC: PC1 MAc

2. 报文在 SW，因为不需要做 ARP request，SW 也就不用广播，报文同 PC1 报文

3. 报文在 R1

   Src IP: 192.168.1.1

   Dst IP: 192.168.1.129

   Dst MAC: PC2 MAC

   Src MAC: R1 MAC

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-44.1cm0mgyalgjk.webp)

假设现在是正常对的，需要做 ARP request，即使划分了subnets，同样还是有问题的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-46.46r5lmj5dhds.webp)

在报文到达 SW 时，因为 Dst MAC FFFF.FFFF.FFFF 是 2 层广播地址，所以 SW 会做 boardcast 转发到除接收端口外的所有端口

> Although we separated the three departments into three subnets(Layer3), they are still in the same broadcast domain(Layer2)

当然可以为每一个 subnet 配一个 SW 来解决，但是会增加设备成本

但是我们可以使用 VLAN 来解决上面的种种问题

## What is a VLAN

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_14-55.5cqrwcwos4n4.webp)

假设现在已经将网络改造成上图中的 3 个 VLAN

> You configure the switch interface to be in a specific VLAN, and then the end host connected to that interface is part of that VLAN

通过配置 Switch interface 来指定 VLAN，例如上图中将 Switch 连接黄色的接口配置 VLAN 10, 蓝色的配置成 VLAN 20, 紫色的配置成 VLAN 30

*A switch will not forward traffic between VLANs including broadcast/unknown unicast traffic(即一个 VLAN 就是一个 boardcast domain)*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_15-01.7drnydnnhk3k.webp)

同样的 PC1 访问 PC2，因为现在不知道 R1 的 MAC 地址，所以要做 ARP request，在 ARP request 到达 SW 时，因为 Dst MAC FFFF.FFFF.FFFF 就只会广播到 VLAN10

PC1 访问 PC2 后面部分就和正常的一样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_15-10.5ln2tklplwg0.webp)

R1 回送 ARP request 到 PC1，PC1 发送

Src IP: 192.168.1.1

Dst IP: 192.178.1.129

Dst MAC: R1 MAC

Src MAC: PC1 MAC

在 R1 收到 PC1 发来的报文后，de-encapsulation 2 层报文头，对比 3 层报文头 Dst IP，查路由，转发

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_15-13.3qckgqxc2u68.webp)

报文为

Src IP: 192.168.1.1

Dst IP: 192.168.1.129

Dst MAC: PC2 MAC

Src MAC: R1 MAC

> **Router is used to route between VLANs**
>
> **The switch does not perform inter-VLAN routing(VLAN 之间的路由). It must send the traffic through the router**
>
> **even if PC1 and PC2 were in the same subnet, the switch wouldn’t forwrd the traffic from PC1 to PC2, because they are in separate VLANs**

### Summary 

- VLANs are configured on **switches** on a per-interface basis (router 不需要做额外的配置，router 并不会广播报文)

- logically separate end hosts at Layer2

- Per VLAN, Per boardcast domain

- Switch do not forwrd traffic directly between hosts in different VLANs(the switch msut forward the traffic to a router)

  不同 VLANs 间通信需要经过 router，相同 VLAN 间通信不需要

## Access Configuration

- `show vlan br`

  显示 Switch 上所有的 VLAN (不会显示 trunk mode 的接口，只会显示 access mode 的接口，如果需要查看 trunk mode 的接口需要使用 `show interfaces trunk`)

  ```
  Switch>show vlan br
  
  VLAN Name                             Status    Ports
  ---- -------------------------------- --------- -------------------------------
  1    default                          active    Fa0/1, Fa0/2, Fa0/3, Fa0/4
                                                  Fa0/5, Fa0/6, Fa0/7, Fa0/8
                                                  Fa0/9, Fa0/10, Fa0/11, Fa0/12
                                                  Fa0/13, Fa0/14, Fa0/15, Fa0/16
                                                  Fa0/17, Fa0/18, Fa0/19, Fa0/20
                                                  Fa0/21, Fa0/22, Fa0/23, Fa0/24
                                                  Gig0/1, Gig0/2
  1002 fddi-default                     active    
  1003 token-ring-default               active    
  1004 fddinet-default                  active    
  1005 trnet-default                    active  
  ```

  > 如果没有对端口分配 VLAN，默认都在 VLAN1
  >
  > VLAN1,1002-1005 是默认的，不能被删除

- `switchport mode access`

  将当前交换机端口转为 access port

  > An access port is a switchport which belongs to a single VLAN, and usually connectes to end hosts like PCs
  >
  > Switchports which carry multiple VLANs are called ‘trunk ports’.

- `switchport access vlan <vlan-id>`

  将当前交换机端口分配到指定的 VLAN 

- `VLAN <vlan-id>`

  配置指定的 VLAN

- `name <vlan-name>`

  配置 VLAN 的名字，而不是默认的 vlan-id

## LAB

这里已经手动接线了，不需要

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-26_16-42.1du872hag2ww.webp)

### SW1

```
SW1(config-if)#int range f3/1,f4/1,g0/1
SW1(config-if)#switchport mode access
SW1(config-if)#switchport access vlan 10
SW1(config-if)#int range f5/1,f6/1,g1/1
SW1(config-if)#switchport mode access
SW1(config-if)#switchport access vlan 20
SW1(config-if)#int range f7/1,f8/1,g2/1
SW1(config-if)#switchport mode access
SW1(config-if)#switchport access vlan 30
SW1(config-if)#vlan 10
SW1(config-if)#name engineering
SW1(config-if)#vlan 20
SW1(config-if)#name hr
SW1(config-if)#vlan 10
SW1(config-if)#name sales
 
SW1#sh vlan br

VLAN Name                             Status    Ports
---- -------------------------------- --------- -------------------------------
1    default                          active    Fa9/1
10   engineering	                    active    Gig0/1, Fa3/1, Fa4/1
20   hr		                            active    Gig1/1, Fa5/1, Fa6/1
30   sales                            active    Gig2/1, Fa7/1, Fa8/1
1002 fddi-default                     active    
1003 token-ring-default               active    
1004 fddinet-default                  active    
1005 trnet-default                    active
```

### R1

```
R1(config)#int g0/0
R1(config-if)#ip add 10.0.0.62 255.255.255.192
R1(config-if)#no shutdonw
R1(config-if)#int g0/1
R1(config-if)#ip add 10.0.0.126 255.255.255.192
R1(config-if)#no shutdonw
R1(config-if)#int g0/2
R1(config-if)#ip add 10.0.0.190 255.255.255.192
R1(config-if)#no shutdonw
```

**referenes**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=aHwAm8GYbn8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=19