# Day22 - Rapid Spanning Tree Protocol

## STP version

在介绍 RSPT 前，对比一下 IEEE标准 和 Cisco标准 的 STP

> 最大的区别是，IEEE 不支持每一个 VLAN 单独配置 STP(除 MSTP 外)，而 Cisco 支持(STP Load Balancing)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_14-22.w9edsihiajk.webp)

在大型的网络拓扑中，一般会优先使用 Multiple Spanning Tree Protocol(MSTP)。因为支持为一组 VLAN 配置 STP Load Balancing

## RSTP

*“RSTP” is not a timer-based spanning tree algorithm like 802.1D. Therefore, RSTP offers an improvement over the 30 seconds or more that 802.1D takes to move a link to forwarding .The heart of the protocol is a new bridge-bridge handshake mechanism, which allows ports to move directly to forwarding*

普通的 SPT 如果初始端口到 forwarding 状态需要消耗 30 秒（15 delay timer * 2）

从 blocking 的状态转为 forwarding 需要消耗 50 秒（20 max age + 15 delay timer * 2）

RSTP 相比普通的 STP，会使用 handshake 的机制来判断端口应该处于的状态

### similarities between STP and RSTP

- RSTP serves the same purpose as STP, blocking specific ports to prevent Layer2 loops
- RSTP elects a root bridge with the same rules as STP
- RSTP elects root ports with the same rule as STP
- RSTP elects designated ports with the same rule as STP

### differences between STP and RSTP

#### Port cost

端口花费不一样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_14-47.533otewecxs0.webp)

#### Port state

端口状态不一样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_14-49.2z51qy9ijds0.webp)

对比普通的 STP

- Blocking/Disabled 合并成 discarding
- Listening 状态被去除

#### Port roles

端口角色不同

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_14-56.6vi0vjobteyo.webp)

> root port 和 designated port 角色和逻辑都相同，但是 non-designated port(blocking port) 被拆分成两个角色
>
> 1. alternate port
> 2. backup port
>
> 状态仍为 blocking

##### Alternate Port

> 如果端口能从其他交换机收到 BPDUs，说明这个端口就是 Alternate Port(需要先判断 root port 和 designated port)

- The RSTP alternate port role is a discarding port that receives a superior BPDU **from another switch**
- This is the same as what you’ve learned about blocking ports in classic STP
- Functions as a backup to the root port

> 前 3 点和普通的 STP non-designated port 一样

- If the root port fails the switch can immediately move its best alternate port to forwarding

  如果 root port 有问题了，会立马使用一个替代 root port 的端口作为 root port

##### Backup Port

> 如果能从不同端口收到从同一个交换来的 BPDUs，说明这个端口就是 Backup port(需要先判断 root port 和 designated port)

- The RSTP backup port role is a discarding port that receives a superior BPDU **from another interface on the same switch**
- This only happens when two interfaces are connected to the same collision domain(via a hub)
- Hubs are not used in modern networks, so you will probably not encounter an RSTP backup port
- Functions as a backup for a designated port 

### UplinkFast

例如有如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_15-06.7choktplax6o.webp)

假设现在 SW3 和 SW1 互联的链路有问题，端口收不到 SW1 BPDUs 了，那么 SW3 和 SW2 互联的端口就会从 non-designate port 变成 root port，从 blocking 转为 fowarding

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_15-07.7j6ixal2vo8w.webp)

*This immediate move to forwarding state functions like a classic STP optional feature called UplinkFast. Because it built into RSTP, you do not need to activate UplinkFast when using RSTP/Rapid PVST+*

> 在普通的 STP 中，端口状态不能立马变为 forwarding，需要先经过 listening 和 learning。而在 RSTP 中可以直接将 non-designated port 变为 designate port
>
> UplinkFast 是 RSTP 中默认启用的功能

### BackboneFast

有如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230606/2023-06-06_15-24.1s2qdqtab868.webp)

假设现在 SW1 和 SW2 之间的链路出问题了，SW2 收不到 BPDUs 了，就会认为自己就是 root bridge，然后往 SW3 发送 BPDUs

但是呢 SW2 发送的 BPDUs 中 bridge ID, 是比 SW1 往 SW3 发送的大的，所以 SW3 并不会认为 SW2 是 root bridge

所以 SW3 会忽略 SW2 发送过来的 BPDUs，直到端口的状态从 non-designated 变为 designated(forwarding)，并转发从 SW1 来的 BPDUs 到 SW2

SW2 重新认定 SW1 为 root bridge，然后 SW2 变更端口角色

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230607/2023-06-07_17-56.6eazbkfx0etc.webp)

*BackboneFast allows SW3(和 SW2 互联的端口) to expire the max age timers(20 秒) on its itnerface and rapidly forward the superior BPDUs to SW2*

> 使用 BackboneFast(在 RSTP 中) ，在端口收不到 BPDUs 的情况下，可以无需等待 max age(20)
>
> 上述的例子 SW3 和 SW2 互联的端口，在收不到从 SW2 转发过来的 BPDUs 时。不会等待 20 秒，判断已经完全收不到 BPDUs
>
> BackboneFast 是 RSTP 中默认启用的功能

## RSTP Configuration

可以使用 `spanning-tree mode rapid-pvst` 来使用 RSTP(在 Cisco 中默认使用 RSTP，所以一般无须使用该命令)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230607/2023-06-08_13-57.66fsr2vd3wxs.webp)

使用上述命令后，可以看到 `spanning tree enabled protocol rstp` 就表明开启并使用了 RSTP

同时 Interface Role，就会出现 Back 和 Altn，分别表示 Backup port 和 Alternate port

> 如果互联的交换机端口使用了普通的 STP 没有使用 RSTP，那么 2 台交换机互联的端口都会使用普通 STP(即 RSTP 兼容 STP)。但是非互联的端口仍然会使用 RSTP

## RSTP message

对比一下和普通 STP 报文，左边为普通 STP 报文，右边为 RSTP 报文

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230607/2023-06-08_14-09.67jyzozv12bk.webp)

1. Protocol Version Identifier

   普通 STP 使用 0， RSTP 使用 2

2. BPDU Type

   普通 STP 使用 0x00, RSTP 使用 0x02

3. BPDU flags

   These flags are used in the negotiation process that allows rapid STP to converge much faster than classic STP

## RSTP Link Types

RSTP 将互联的链路分为 3 种

> RSTP 会**自动**识别 link 属于那种类型，然后将端口分别**自动**配置成对应 link 的类型 

1. Edge

   A port that is connected to an end host. Moves directly to forwarding, without negotiation

   和 portfast 功能类似，可以使用 `spanning-tree portfast` 将端口手动变为 edge port

2. Point-to-point

   A direct connection between two switches

   可以使用 `spanning-tree link-type point-to-point` 手动将端口变为 point-to-point port

3. Shared

   > connection 是指 Switch 和 Switch 通过 Hub 互联

   A connection to a hub. Must operate in half-duplex mode

   可以使用 `spanning-tree link-type shared` 手动将端口变为 shared

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_14-24.4ah9hko3dcw0.webp)

## RSTP VS STP

- All switches running Rapid STP send their own BPDUs every hello time(2 sec)

  **所有使用 RSTP 的 switches 都会自动的发送自己的BPDUs**

- Switches ‘age’ the BPDU information much more quickly. In classic STP, a switch waits 10 hello intervals(20 senconds). In rapid STP, a switch considers a neighbor lost if it misses 3 BPDUs(6 seconds). It will then ‘flush’(删除) all MAC addresses learning on that interface

## QUIZ

### 0x001

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230607/2023-06-08_13-38.cvs5jx7llmo.webp)

1. 先判断 root bridge。因为在 priority ID 相同的情况下 SW1 MAC 最小，所以 SW1 是 root bridge，SW1 G0/0 G0/1 均为 designated port
2. 再判断 root port。SW3 G0/2 cost 最小所以是 root port；SW2 G0/0 cost 最小所以是 root port；SW4 在 cost 相同的情况下, neighbor SW2 MAC 小，所以 SW4 G0/1 是 root port
3. 再判断剩下的 designated port。SW2 G0/1 designated port;SW3 和 SW4 之间是一个 collision domain(一个 collision domain 中必须要有一个 designated port), 因为 SW3 cost 小，同时 SW3 G0/0 port ID 比 SW3 G0/1 port ID 小，所以 SW3 G0/0 是 designated port
4. 然后判断 alternate port。SW4 G0/0 收到的 BPUDs 是从 SW3 转发过来的，所以 SW4 G0/0 是 alternate port
5. 最后判断 backup port。SW3 G0/1 BPDUs 可以是从 SW3 G0/0 来的，所以 SW3 G0/1 是 backup port

```
if a is root_port then:
	...
else if a is designated_port then:
	...
else if a is alternate_port then:
	...
else if a is backup_port then:
	...
```

## Lab

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230608/2023-06-08_14-45.2cze3m3xgzi8.webp)

### 0x01

SW1 是 root bridge

使用 `show spanning-tree` 查看

### 0x02

#### SW1

F0/1 F0/2 F0/3 F0/24 designated port

```
Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Fa0/24           Desg FWD 19        128.24   Shr
Fa0/1            Desg FWD 19        128.1    P2p
Fa0/2            Desg FWD 19        128.2    Shr
Fa0/3            Back BLK 19        128.3    Shr
```

这里 F0/3 为 Back up port，因为 SW1 和 SW3 通过 hub 互联是一个 collision domain, 一个 collision domain 中只能有一个 designated port，因为 F0/2 port id 比较小，所以 F0/2 是 designate port，那么另外一个端口 F0/3 一定是 backup port

#### SW2

F0/1 root port

F0/23 F0/24 designated port

G0/1 alternate port

F0/2 designated port

```
Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Fa0/1            Root FWD 19        128.1    P2p
Fa0/2            Desg FWD 19        128.2    P2p
Fa0/23           Desg FWD 19        128.23   P2p
Fa0/24           Desg FWD 19        128.24   P2p
Gi0/1            Altn BLK 4         128.25   P2p
```

#### SW3

F0/2 root port

F0/24 designated port

G0/1 designated port

F0/1 designated port

```
Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Fa0/1            Desg BLK 19        128.1    P2p
Fa0/2            Root FWD 19        128.2    Shr
Fa0/24           Desg BLK 19        128.24   P2p
Gi0/1            Desg BLK 4         128.25   P2p
```

#### SW4

F0/1 root port

F0/24 designated port

F0/2 alternate port

```
Interface        Role Sts Cost      Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Fa0/1            Root FWD 19        128.1    P2p
Fa0/2            Altn BLK 19        128.2    P2p
Fa0/24           Desg FWD 19        128.24   P2p
```

### 0x03

> edge port 对应的 Type 也会显示成 P2p

#### SW1

```
SW1(config)#int fa0/24
SW1(config-if)#spanning-tree portfast
SW1(config)#int fa0/1
SW1(config-if)#spanning-tree link-type point-to-point 
SW1(config)#int fa0/2,f0/2
SW1(config-if)#spanning-tree link-type shared
```

SW1 的 f0/24 应该是 edge link type，因为只有 Switch 之间通过 hub 互联才会是 shared link type

#### SW2

```
SW1(config)#int fa0/23,fa0/24
SW1(config-if)#spanning-tree portfast
SW1(config)#int fa0/1,g0/1
SW1(config-if)#spanning-tree link-type point-to-point 
```

#### SW3

```
SW1(config)#int fa0/24
SW1(config-if)#spanning-tree portfast
SW1(config)#int fa0/1,g0/1
SW1(config-if)#spanning-tree link-type point-to-point 
SW1(config)#int fa0/2
SW1(config-if)#spanning-tree link-type shared
```

#### SW4

```
SW1(config)#int fa0/24
SW1(config-if)#spanning-tree portfast
SW1(config)#int fa0/1,f0/2
SW1(config-if)#spanning-tree link-type point-to-point 
```

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=EpazNsLlPps&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=41