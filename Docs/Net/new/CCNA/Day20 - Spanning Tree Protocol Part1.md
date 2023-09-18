# Day20 - Spanning Tree Protocol Part1

## Network Redundancy

在网络中 Redundancy(冗余) 是完全有必要的

*Modern networks are expected to run 24/7/365. Even a short downtime can be disastrous for a business*

所以当网络中的一个 component fails，你必须要保证网络中的其他 components 能替代 failed component 的功能。需要尽可能地为网络中每一个节点，都配置 redundancy

例如下面的这个拓扑结构中就不具备 redundancy

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-42.4aeoyws0zdc.webp)

如果 Router 和 Internet 中间的链路有问题，那么整个网络中的所有主机就会有问题

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-44.6xstz050aa2o.webp)

如果交换之间的链路有问题，那个连在交换机下的主机就会有问题

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-45.3p429btshsao.webp)

对比一下一个设计地比较完整的拓扑，如果两个 components(除主机外) 中的链路出现问题，主机还是可以正常访问 Internet 的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-46.4fjiqgzzv79c.webp)

例如

PC 访问 Internet

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-48.75h7yg4aixz4.webp)

或者同 LAN PC 访问 PC

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-49.6jc5kri5m6f4.webp)

但是如果和主机互联的 SW 出现问题了，一般就会有问题

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-52.2sm431du8800.webp)

但是呢，这个拓扑任然是有问题的，会导致 Broadcast Storms

### Broadcast Storms

Broadcast Storms 中文也叫做广播风暴

*The Ethernet header doen’t have a TTL field. These broadcast frames will loop around the network indenfinitely(无限地). If enough of these looped broadcasts accumulate in the network, the network will be too congested for legitimate(合法的) traffic to use the network. This is called a broadcast strom*

以如下拓扑为例

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_14-00.3uyz0jzd10w0.webp)

假设 PC1 需要访问 PC2，首先需要知道 PC2 的 MAC address，所以需要做 ARP request。因为 ARP request 目的 MAC 是 FFFF.FFFF.FFFF，所以会被 SW 广播

1. 当 SW1 收到 ARP request 时，会广播到 和 SW2 以及 SW3 互联的端口
2. 当 SW2 收到 ARP request 时，会广播到 和 SW3 以及 PC2 互联的端口
3. 当 SW3 收到 ARP request 时，会广播到 和 SW2 以及 PC3 互联的端口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_14-23.4al3qbx1pp4w.webp)

这样就有一个问题，SW2 会和 SW3 互相收到从对方发来的广播帧，因为目的 MAC 任然是 FFFF.FFFF.FFFF，所以他们又会广播到 SW1，SW1 又会分别广播到 SW2 或者 SW3。这样就形成了一个环



因为 SW 只是一个二层的设备，不会考虑 3 层的报文头，所以 3 层报文头里 TTL 对他们来说是没有意义的。所以报文会在链路中反复广播，直到打垮网络

### MAC Address Flapping

上面的拓扑还有一个问题就是 MAC address Flapping

SW 在 de-encapsulation 报文的时候并不会修改报文中的 MAC address。但是会将收到的报文中的 MAC source Address 是从那个端口来的记录到 MAC address table 中

1. 在 SW1 将广播帧发送到 SW3 时，PC1 Source MAC address 是从 SW3 互联 SW1 的端口来的，并记录到自己 MAC address table 中

2. 在 SW2 将广播帧发送到 SW3 时，PC1 Source MAC address 是从 SW3 互联 SW2 的端口来的，会更新 MAC address table 对应 PC1 Source MAC address 的条目，将端口值变更为 SW3 和 SW2 互联的端口

因为 boardcast storm 的原因，上述 2 步会重复，所以 MAC address table 也就会不断的更新 PC1 Source MAC address 对应条目的端口值，这个也被称为 MAC address flapping

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_15-57.6psfj7wefi0w.webp)



## STP

STP 就是为了解决 Boardcast Storm 和 MAC Address Flapping

Spanning Tree Protocol 是一种为了保证 2 层 network Redundancy 正常运行的协议，所以只会在 LAN 中生效

默认所有 vendors 的网络设备，都会开启并使用 STP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_13-58.1lummuyftqww.webp)

使用了 STP 端口有两种状态

- Forwarding

  Interfaces in a forwarding state behave normally. They send and receive all normal traffic

  **如果互联的是 Router 或者是 PC，一定是 forwarding state**

- Blocking

  Interfaces in a blocking state only send or receive STP messages(called BPDUs = Bridge Protocol Data Units)

> These interfaces act as backups that can enter a forwarding state if an active (=currently forwarding) interface fails

这里需要提一嘴 STP Bridge Protocol Data Units 中的 Bridge

都知道 Hub 是不会管报文是几层的，只要是报文都会 flood 其他的端口

在 Hub 往 Switch 发展的过程中，还有一个 Bridge，是一个只有 2 口的 Switch，一样会做 Mac address 寻址

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-09.23bwu7n56t28.webp)

这里直接将 Bridge 理解成 Switch 即可

### Example

假设

绿色的端口都是 forwarding 状态的，会正常转发报文

橙色的端口都是 blocking 状态的，只会发送或者接受 STP 协议的报文，不会转发其他的报文(比如 ARP)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-15.4n55gurcnsow.webp)

那么可以逻辑上的将 SW3 和 SW2 间的链路想象成不存在，以虚线表示

如果 PC1 发送 ARP request，那么流量就会如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-18.2t2nleqf5a9s.webp)

如果 SW2 和 SW1 之间的链路出问题了，链路就会自动地调整成如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-21.jukik49f14g.webp)

*By selecting which ports are forwarding and which ports are blocking, STP creates a single path to/form each point in the network. This prevents Layer2 loops*

### BPDUs

STP-enabled switches send/receive Hello BPDUs out of all interfaces, the default timer is 2 seconds(the switch will send a Hello BPDU out of every interface, once every 2 seconds)

If a switch receives a Hello BPDU on an interface, it knows that interface is connected to another switch(routers, PCs,etc. do not use STP, so they do not send Hello BPDUs)

即如果收到 BDPUs Hello 报文，表示对端互联的设备是一台 Switch;如果没有收到 BDPUs Hello 就表示对端互联的设备是 Router 或者是 PC

设备收发 BDPUs Hello 图如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-29.59jv895r8fb4.webp)

#### Root bridge

那么 BPDUs Hello 有什么用呢，主要用于选 root bridge

1. Switches use one field in the STP BPDU, the Bridge ID field, to elect a **root bridge** for the network
2. The switch with the **lowest Bridge ID** becomes the root bridge
3. All ports on the root bridge are put in a **forwarding state**, and other switches in the topology must have a path to reach the root bridge

> All interfaces on the root bridge are designated(代理功能上就是转发) ports Designated ports are in a forwarding state

#### Bridge ID

Birdge ID 长得类似下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-34.4y439ydru64g.webp)

> 这里的 MAC address 对应 Switch MAC[^STP Switch MAC address] 是 SWitch 的一个唯一标识符
>
> 逻辑上 Switch 设备本身并不需要 MAC address，但是 STP 需要 MAC address 来决定 root bridge，所以一般的 Switch 都会有一个 MAC 来代指设备本身

*The default bridge priority is 32768 on all switches, so by default the MAC address is used as the tie-breaker(lowest MAC address becames the root bridge)*

这里的 tie 有平局的意义，tie-breaker 就是平局后的决定因素，==即 bridge priority 值相同，就比较 MAC address==

*The Bridge Priority is compared first. If they tie, the MAC address is then compared*

假设拓扑如下

因为 SW 和 PC 互联的端口，一定是不会出现环的，所以一定是 forwarding state(designated port)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-41.19g6y2i9qw2o.webp)

现在因为都是默认的状态，所以 Bridge priority 默认都是 32768，因为 3 台 SW 都一样，所以会比较 3 台 SW 的 MAC address

因为 SW1 在 3 台 SW 中的 MAC address 最小，所以 SW1 就会选举成 root bridge，那么他所有的端口就都会是 forwarding state

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-46.1nawvfa9pk2o.webp)

#### Bridge Priority

在 Cisco 设备上实际 Bridge Priority 还可以分成两部分

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_16-48.45bd344rkwow.webp)

> 这里为什么包含 VLAN ID 呢，因为 Cisco 实际使用的是 PVTS(Per-VLAN Spanning Tree)。每一个 VLAN 都会运行一个单独的 STP instance，在不同的 VLAN 中，不同的接口状态可以是不同的
>
> 例如 SW1 G0/1 在 VLAN10 是 forwarding，但是在 VLAN20 中可以是 blocking

为什么默认的 bridge priority 是 32768

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_17-08.6huamg77xri8.webp)

因为 bridge prority 一共占 16 bits，默认的 prority bit 是最高为，所以 32768 = $2^{15}$

但是因为 Cisco 的设备使用的是 PVTS 还会加上 Extend System ID(VLAN ID) 到 bridge priority，上图中因为是 1(VLAN 1)，所以最后实际的 bridge priority 为 32769 

*In the default VLAN of 1, the default bridge priority is actually 32769(32768 + 1)*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_17-20.5iq7pmdxq4u8.webp)

> 如果是 VLAN2，那么对应的 bridge priority 就会是 32770
>
> 如果是 VLAN10, 那么对应的 bridge prioriry 就会是 32779

现在假设你想减小或者增加 bridge priority 那么最小的差值单位(the minimum unit fo increase/decrease)为多少？

因为实际的 bridge priority 是由 bridge priority 中的 bridge priority 和 Extended Sytem ID(VLAN ID) 组成的，但是 Extended System ID 是不能修改的(因为和 VLAN 有关，如果 VLAN 变了，Extended System ID 就会改变)。所以 bridge priority 可以使用的 bit 只有 12 - 15，最小的 bit 对应的值为 4096 = $2^{12}$，即最小的差值单位为 4096

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_17-18.75lngv9rcvi8.webp)

因为 bridge priority 是可以增加或者减小的，所以 root bridge 可以有用户自己设置，例如

在 VLAN1 中 root bridge 可以是 SW1

在 VLAN2 中 root bridge 可以是 SW2

在 VLAN3 中 root bridge 可以是 SW3

### STP Port Role

在 STP 中 Port 是有角色的，角色不同端口的状态也不同

> designated port 和 root port 端口的状态都是 forwarding 的

#### Designated Port

- The Switch with the lowsest bridge ID is elected as the root bridge. All ports on the root bridge are **designated ports**(forwarding state).

- Every collision domain has a single STP designated port

  会比较 switch root cost,如果一样比较 bridge ID(自己的 bridge ID)

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_18-32.22vkf2c005ds.webp)

  SW2 最小的 root cost 为 SW2 G0/1 - SW1 G0/0 4

  SW3 最小的 root cost 为 SW3 G0/0 - SW1 G0/1 4

  比较 bridge ID，因为 SW2 MAC 比 SW3 小，所以 G0/0 为 designated port

- Ports connect to Routers or PCs are designated ports

#### Root Port

> root port 互联的端口一定是 designated port

Each remaning switch will select **ONE(root port 只能有一个)** of its interfaces to be its root port. The interface with the lowset root cost will be the **root port**. Root ports are also in the forwarding state

root port 会先由 root cost 决定

*The root cost is the total cost of the outgoing interfaces along the path to the root bridge*

即 SW 到 root bridge 经过的**出向**的端口花费总值 

端口花费值参考下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_18-06.59dyjoyu8xa8.webp)

例如，下面拓扑，所有的端口都是 Gigabits 的所以端口的花费均为 4

> 课件里的图和逻辑 cost 不好理解
>
> 按照文字里的逻辑理解

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_18-41.4vuyfun6uvls.webp)

因为 SW1 是 root bridge 所以到 root bridge 不需要经过任何的出端口(自己本身就是 root bridge)，所以 root cost 是 0，每个端口都是 root port(这种情况也将 port记为 designated port)

因为 SW2 不是 root bridge，到 root bridge 的路有有两条，分别对应是

1. SW2 G0/1 -> SW1 G0/0

   路径中只有 SW2 G0/1 是出向的端口，所以 root cost 4

2. SW2 G0/0 -> SW3 G0/1 -> SW3 G0/0 -> SW1 G0/1

   路径中 SW2 G0/0 和 SW3 G0/0 都是出向的端口，所以 root cost 4 + 4 = 8

因为路径 1 root cost 最小，所以 SW2 G0/1 是 root port

因为 SW3 不是 root bridge，到 root bridge 的路有有两条，分别对应是

1. SW3 G0/0 -> SW1 G0/1

   路径中只有 SW3 G0/0 是出向的端口，所以 root cost 4

2. SW3 G0/1 -> SW2 G0/0 -> SW2 G0/1 -> SW1 G0/0

   路径中 SW3 G0/1 和 SW2 G0/1 都是出向的端口，所以 root cost 4 + 4 = 8

因为路径 1 root cost 最小，所以 SW3 G0/0 是 root port

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-01_18-48.6gu5pth82yo0.webp)

*The ports connected to another switch’s root port MUST be designated. Because the root port is the switch’s path to the root bridge, another switch must not block it*

> 但是也有可能会出现 root cost 相同的情况，这时就会比较互联设备发送过来的 BPDUs 中的 bridge ID
>
> 收到 bridge ID 值小的端口作为 root port，如果 bridge ID 相同比较 MAC address

例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-02_12-29.48ba2ynmxgjk.webp)

4 台 SW 中，SW2 会被选为 root bridge，因为 Bridge Id 最小，所以 SW2 所有的端口就是 designated ports

SW1 G0/0 和 SW4 G0/1 分别是 root ports，因为另外一条到 root bridge 的路径 cost 为 12

SW3 两条路径到 root bridge cost 均为 8，因为相同，所以会比较互联设备 bridge ID。因为 SW1 和 SW4 bridge ID 都为 32769，所以就比较 MAC address，因为 SW1 MAC 比 SW4 小，所以 SW3 和 SW1 互联的端口 G0/0 就是 root port，那么 SW1 G0/1 就是 designated port

##### Port ID

> 但是也有可能会出现 root cost 相同，并且互联设备发过来的 BPDUs bridge ID 和 MAC address 相同的情况
>
> 这时就会比较互联设备端口的 port ID，收到 port ID 值小的端口就会是 root port

下图红框在的部分关联 port ID

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-02_13-30.7dcp5rkljwn4.webp)

*STP port ID = port priority(default 128) + port number*

例如

128.1 对应的 port ID 为 129，128.7 对应的 port ID 为 135

以下图拓扑为例

SW1 和 SW2 通过两条 link 互联

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-02_13-36.17qcvedbdnfk.webp)

SW3 有 3 条路径到 root bridge，cost 均为 8

比较 bridge ID，SW4 和 SW1 bridge ID 相同，比较 MAC，因为 SW1 MAC 小，所以看 SW1

SW1 有两路径和 SW3 互联，端口分别为 G0/1 和 G0/2。G0/1 的序号比 G0/2 的小，所以 port ID 比 G0/2 小，所以和 SW1 G0/1 互联的端口 G0/2 就是 SW3 的 root port(因为 SW3 G0/2 是 root port 所以和他互联的 SW1 G0/1 是 designated port)

*The NEIGHBOR switch’s port ID is used to break the tie, not the local switch’s port ID*

#### Blocking Port

> 即不是 root port 也不是 designated port 就是 blocking port

上面说了 designated port 和 root port,但是还没有说处于 blocking 状态的接口叫什么

和 HUB 不一样，Switch 之间互联的 link 都是一个 collision domain ,每一个 collision domain 至少有一个 designated port，剩下的就是 blocking port

### The process of STP

1. When a switch is powerd on, it assumes it is the root bridge(send BPDUs)

2. It will only give up its position if it receives a ‘superior’ BPDU(lower bridge ID).

   > 这里不管什么情况，只要 Switch 没有收到较优的 BPDU 就会认为自己就是 root bridge

   One switch is elected as the root bridge. All ports on the root bridge are designated ports(forwarding state). 

   > root bridge 判断逻辑如下
   >
   > 1. Lowest bridge ID
   > 2. Lowest MAC address
   >
   > a b 分别为 2 台 SW

   ```
   if a.bridge_id < b.bridge_id then:
   	a = root_bridge
   else if a.bridge_id > b.bridge_id then:
   	b = root_brige
   else if a.bridge_id == b.bridge_id then:
   	if a.MAC < b.MAC then:
   		a = root_bridge
   	else if a.MAC > b.MAC
   		b = root_bridge
   ```

   Each remaining switch wil select ONE of its interfaces to be its root port(forwarding state). Ports across from the root port are always designated ports

   > root port 判断逻辑如下
   >
   > 1. Lowest root cost
   > 2. Lowest neighbor bridge ID
   > 3. Lowest neigbor MAC address
   > 4. Lowest neighbor port ID
   >
   > a b 分别为 1 台 SW 的两个端口

   ```
   if cost(a to root_bridge) < cost(a to root_bridge) then:
   	a = root_port
   else if cost(a to root_bridge) > cost(b to root_bridge) then:
   	b = root_port
   else if cost(a to root_bridge) == cost(b to root_bridge) then:
   	if a.neighbor.bridge_id < a.neighbor.bridge_id then:
   		a = root_port
   	else if a.neighbor.bridge_id > b.neighbor.bridge_id then:
   		b = root_port
     else if a.neighbor.bridge_id == b.neighbor.bridge_id then:
     	if a.neighbour.MAC < b.neighbour.MAC then:
     		a = root_port
       else if a.neighbor.MAC > b.neighbor.MAC then:
       	b = root_port
       else if a.neighbor.MAC == b.neighbor.MAC then:
       	if a.neighbor.port_id < b.neighbor.port_id then:
       		a = root_port
         else if a.neighbor.port_id > b.neighbor.port_id then
         	b = root_port
   ```

3. Each remaining collision domain will select ONE interface to be a designated port(forwarding state). The other port in the collision domain will be non-designated(blocking)

   > designated port(blocking port) in remainging collision domain
   >
   > a b 分别为 2 台 SW 的两个端口(也可以是一台 SW 两个端口，例如中间通过 hub 互联，hub 互联的整个网络就是一个 collision domain)

   ```
   if cost(a to root_bridge) < cost(a to root_bridge) then:
   	a = designated_port
   	b = blocking_port
   else if cost(a to root_bridge) > cost(b to root_bridge) then:
   	a = blocking_port
   	b = designated_port
   else if cost(a to root_bridge) == cost(b to root_bridge) then:
   	if a.bridge_id < b.bridge_id then:
   			a = designated_port
   			b = blocking_port
     else if a.bridge_id > b.bridge_id then:
     		a = blocking_port
   			b = designated_port
     else if a.bridge_id == b.bridge_id then:
     		if a.port_id < b.port_id then:
     			a = designated_port
     			b = blocking_port
         else a.port_id > b.prot_id then:
         	a = blocking_port
         	b = designated_port
   ```

4. Once the topology has converged(网络拓扑发生改变) and all switches agree on the root bridge, only the root bridge sends BPDUs

5. Other switches in the network will forward these BPDUs, but will not generate their own original BPDUs

## Quiz

### 0x01

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230601/2023-06-05_18-56.43k1ub2nkbls.webp)

1. 先判断 root bridge，因为 priority ID 都一样，所处比较 MAC，SW3 MAC 最小，所以 SW3 是 root bridge。所以 SW3 G0/0 和 SW3 G0/1 都是 designated port
2. SW1 G0/1 到 root bridge cost 最小，所以 SW1 G0/1 是 root port；SW4 G0/0 到 root bridge cost 最小，所以 SW4 G0/0 是 root port
3. SW2 G0/1 G0/0 G0/2 到 root bridge cost 都一样均为 8，所以比较 neighbor bridge ID，因为 SW1 最小，所以在 SW2 G0/1 和 SW2 G0/2 之间选 root port; 因为 neighbor SW1 G0/0 port ID 比较小，所以互联的 SW2 G0/2 是 root port
4. 因为每个 collision domain 中必须要有一个 designated port，所以 SW1 G0/0 为 designated port
5. SW1 G0/2 和 SW2 G0/2 比较 root cost，因为 SW1 root cost 比 SW2 小，所以 SW1 G0/2 为 designated port，所以 SW2 G0/1 为 blocking port
6. SW2 G0/0 和 SW4 G0/1 比较 root cost，因为 SW4 root cost 比 SW2 小，所以 SW4 G0/1 位 designated port，所以 SW2 G0/0 为 blocking port

**referneces**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=j-bK-EFt9cY&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=38
[^STP Switch MAC address]:https://networkengineering.stackexchange.com/questions/54568/does-switch-need-its-own-mac-address