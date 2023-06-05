# Day21 - Spanning Tree Protocol(Part2)

## STP Port States

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_19-47.4o1qrjxkbdds.webp)

### Blocking

- Non-designated ports are in a Blocking state

- Interfaces in a Blocking state are effectively disabled to prevent loops

- Interfaces in a Blocking state are not send/receive regular network traffic

- Interfaces in a Blocking state receive STP BPDUs

- Interfaces in a Blocking state do NOT forward STP BPDUs

- Interfaces in a Blocking state do NOT learn MAC addresses

  如果来的报文中的 Dst MAC 不在 MAC address table 中，就会被直接丢弃

### Listening

- After the Blocking state, interfaces with the Designated or Root role enter the Listening state

  由 Blocking 状态转为 Designated port 或者 Root port 后会先经过 Listening 状态

- Only Designated or Root ports enter the Listenging state(Non-designated ports are always Blocking)
- The Listenging state is 15 seconds long by defualt. This is determined by the Forward delay timer
- An interface in the Listening state ONLY forwards/receives STP BPDUs
- An interface in the Listening state does NOT send/receive regular traffic

### Learning

- After the Listening state, a Designated or Root port will enter the Learning state
- The Learning state is 15 seconds long by default. This is determined by the Forward delay timer(the same timer is used for both the Listening and Learning states)

- An interface in the Learning state ONLY sends/receives STP BPDUs
- An interface in the Learning state does NOT send/receive regular traffic
- An interface in the Learning state learns MAC addresses from regular traffic that arrives on the interfaces

### Forwarding

- Root and Designated ports are in a Forwarding state
- A port in the Forwarding state operate as normal
- A port in the Forwarding state sends/receives BPDUs
- A port in the Forwarding state sends/receives normal traffic
- A port in the Forwarding state learns MAC addresses

上述内容可以精简成下表，Stable/Transitional 表示状态是否处于中间态

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_20-12.1js6u72kcy8w.webp)

## STP Timers

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_20-14.2x8imudgkdkw.webp)

### Hello Timer

在初始的状态下，所有的 Switch 都会认为自己是 root bridge，然后往自己所有的端口发送 STP Hello BPDU

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_20-17.34kw5y3a8h6o.webp)

但是一旦确认了每个 Switch 在 STP 中的角色，只有 root bridge 会发送 Hello BPDUs

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_20-19.422dz7n1dvls.webp)

然后除 root bridge 外的其他 Swtiches 会 forwarding Hello BPDUs

> 只有 designated port 会转发 Hello BPDUs，non-designated port 或者是 root port 都不会转发 Hello BPDUs

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_20-21.rmwadq64ryo.webp)

并更新 root cost, neighbor bridge ID, neighbor port ID 等

然后在 2 秒后，root bridge 会重新发送 Hello BPDUs，其他 Switches 同样会 forwarding Hello BPDUs

### Foward Delay Timer

在 Listening 或者是 Learning 状态下的时间，所以从 blocking 到 fowarding 一共需要 30 秒 (15 + 15)

### Max Age Timer

当网络 STP 拓扑改变时，Switch 判断端口不能收到 BDPUs 的时间，如果收到一个 BDPU 就会重新计时

以下面的拓扑为例，主要关注 SW2 G0/1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_20-52.1a34v1w9nkcg.webp)

当 SW1 和 SW2 之间的链路正常的情况下

1. SW3 发送 BPDUs 到 SW1 通过 G0/0 转发到 SW2 G0/1
2. SW2 G0/1 启动一个 Max Age 计时器，从默认的 20 开始
3. 如果了 2 秒后 SW3 再次发送 BPDUs， SW2 G0/1 Max Age 计时器为 18，因为收到了 BPDUs 所以重置计时器到 20
4. 突然 SW1 和 SW2 之间的链路出现了问题
5. SW2 G0/1 不能收到 BPDUs 了，Max Age 到 0，就会重新评估整个 STP 拓扑中的 root bridge, 以及各种端口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-01.6kam81h7gwao.webp)

> 这里的 50 秒，是 20（Max Age，重新评估） + 15 + 15  
>
> A forwarding interface can move directly to a blocking state(there is no worry about creating a loop by blockng an interface)
>
> A blocking interface cannot move directly to forwarding state. It must go through the listening and learning states

## STP BPDU

看一下 BPDU 报文

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-04.26p4dbqmppq8.webp)

- Dst: PVST+ (01:00:0c:cc:cc:cd)

  Cisco 默认使用该 MAC 作为 PVST+ MAC(交换机端口并没有 MAC)

  PVST = Only ISL trunk encapsulation(思科独有的协议)

  PVST+ = Supports dot1q(公用的标准协议)

- Root Identifier

  Root bridge 相关的信息

  - Root Bridge Priotiry：32768

  - Root Bridge System ID Extension: 10

    VLAN10

  - Root Bridge System ID: MAC

    root bridge 的 MAC

- Root Path Cost

  到 root bridge 的花费，这里为 0，表示这就是 root bridge

- Bridge Indentifier

  发送或转发 BPUD 设备 bridge ID 相关的信息

- Port identifier

  发送 BPDUs 的端口标识符，前一半是 Port priority(Port ID)，例子中为 80 hexadecimal 即为 128

## STP Toolkits

STP 还有一些 toolkits 用于增强某些特质 

### Portfast

Portfast 解决了和 PC 或者 Router 互联的端口，状态必须要从 Listening 到 Learning 然后才能到 Forwarding 的问题

一共需要花费 30 秒，但是和 PC 或者 Router 互联，即使有环对整个网络拓扑来说也是没有问题的，因为只是增加 PC 或者是 Router 的负载

如果使用了 Portfast 就可以使端口直接从 Listening 或者是 Learning 状态直接转到 Forwarding，而不需要等待 30 秒

*Portfast allows a port to move immediately to the Forwarding state, bypassing Listening and Learning*

*If used, it must be enabled only on ports connected to end hosts*

*If enable on a port connected to another switch it could cause a Layer 2 loop*

可以使用 `spanning-tree portfast` 来开启端口 portfast 的功能

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-31.4hzqsgchwww0.webp)

也可以使用 `spanning-tree portfast default` 来为所有的 access mode ports(只会对 access mode 有效， 对 trunk mode 没有任何效果) 开启 portfast 的功能

### BPDU Guard

使用 portfast 会有一个问题，那就是当网络拓扑改变了，交换机连的不再是 PC 或者是 Router，而是另外一台 Switch

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-43.5vh94ew064u8.webp)

因为 portfast ，红框中的端口直接是 forwarding 的不，这样就会有环

> portfast 并不是安全的，为了避免这种情况，可以使用 BPDU Guard

If an interface with BPDU Guard enabled receives a BPDU from another switch, the interface will be shut down to prevent a loop from forming

可以使用 `spanning-tree bpduguard enable` 来开启端口 BPDU Guard 功能

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-40.173eyfa1ta9s.webp)

也可以使用 `spanning-tree portfast bpduguard enable default` 来为所有的 access mode ports 开启 bpduguard 功能 

> 需要先开启 portfast 后，才可以使用 bpduguard

### Root Guard

If you enable root guard on an interface, even if it receives a superior BPDU(lower bridge ID) on that interface, the switch will not accept the new switch as the root bridge. The interface will be disable

顾名思义就是方式 root bridge 被篡改

### Loop Guard

If you enable loop guard on an interface, even if the interface stops receiving BPDUs, it will not start forwarding. The interface will be disable

## STP Configuration

Cisco 的所有设备默认都会开启 rapid-pvst，但是如果想要修改 STP 模式，可以使用下面的命令

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-55.10xdncai3w6o.webp)

在下面的拓扑中默认 SW1 应该会是 root bridge，但是我们也可以通过手动的方式来指定 SW3 为 root bridge，SW2 为 secondary bridge

> 当 root bridge 失效时，secondary bridge 会上升为 root bridge

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_21-57.740jjg7ynvgg.webp)

可以使用 `spanning-tree vlan <vlan-id> root primary` 来指定 root bridge

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_22-01.o6c8jyf1f8w.webp)

可以使用 `spanning-tree vlan <vlan-id> secondary ` 来指定 secondary bridge

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_22-04.3kdbqpyxm2o0.webp)

使用上述的命令后拓扑如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_22-08.d7090l5kgkw.webp)

只会对 VLAN1 生效，如果现在有一个 VLAN2，那么对应的 priority ID 以及 root bridge 选择都会是默认的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_22-10.6lg2jks2yo74.webp)

因为在不同的 VLAN 中 STP 拓扑的结构不同，这也被称为 **STP loading-balancing**

端口对应的 cost 也可以通过 `spanning-tree vlan <vlan-id> cost <number>` 来修改

同样的 port priority(port id) 可以通过 `spanning-tree vlan <vlan-id> port-priority <number>` 来修改

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_22-17.45wj74nxpwe8.webp)



**references**

[^jeremy’s IT Lab]: https://www.youtube.com/watch?v=nWpldCc8msY&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=39