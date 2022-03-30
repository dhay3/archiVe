# Linux IP link

ref：

https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking#

## 0x1 Digest

查看或配置设备上L2信息

## 0x2 Cmmand

### 0x2.1 ip link add

syntax：`ip link add [link DEVICE][name]NAME [txqueuelen PACKETS] [address LLADDR] [broadcast LLADDR] [mtu MTU] [index IDX] [numtxqueues QUEUE_COUNT] [numrxqueues QUEUE_COUNT] type TYPE [ARGS]`

添加virtual link

- `DEVICE`

  指定实际的物理网卡，部分link type 无须指定实际的物理网卡

- `NAME`

  虚拟设备的名字

- `TYPE`

  指定虚拟设备的link type，可以是如下几种

- `numtxqueues <queue_count>`

  指定发送数据的队列数

- `numrxqueues <queue_count>`

  指定接受数据的队列数

- `gso_max_size <bytes>`

  指定接受packet的大小

- `gso_max_segs <segments>`

  指定接受segment(不包含包头)的大小

例如添加一个

### 0x2.2 ip link delete

syntax：`ip link delete dev DEVICE [group GROUP] type [TYPE] `

删除virtual link

- `dev DEVICE`

  指定需要删除的 virtual link device，只能是virtual link device

- `group GROUP`

  指定需要删除的group，group 0 不允许删除

  

例如删除 dummy virutal device

```
ip link delete dev dummy0 type dummy
```



## 0x3 Link type

以下均为虚拟设备类型

### 0x3.1 bridge

syntax：`ip link add DEVICE type bridge [ ageing_time AGEING_TIME ] [ group_fwd_mask MASK ] [ group_address ADDRESS ] [ forward_delay FORWARD_DELAY ] [ hello_time HELLO_TIME ] [ max_age MAX_AGE ] [
stp_state STP_STATE ] [ priority PRIORITY ] [ vlan_filtering
VLAN_FILTERING ] [ vlan_protocol VLAN_PROTOCOL ] [ vlan_default_pvid VLAN_DEFAULT_PVID ] [ vlan_stats_enabled VLAN_STATS_ENABLED ] [ mcast_snooping MULTICAST_SNOOPING ] [ mcast_router MULTICAST_ROUTER ] [ mcast_query_use_ifaddr MCAST_QUERY_USE_IFADDR ] [ mcast_querier MULTICAST_QUERIER ] [
mcast_hash_elasticity HASH_ELASTICITY ] [ mcast_hash_max HASH_MAX ] [ mcast_last_member_count LAST_MEMBER_COUNT ] [
mcast_startup_query_count STARTUP_QUERY_COUNT ] [
mcast_last_member_interval LAST_MEMBER_INTERVAL ] [ mcast_mem‐
bership_interval MEMBERSHIP_INTERVAL ] [ mcast_querier_interval
QUERIER_INTERVAL ] [ mcast_query_interval QUERY_INTERVAL ] [ mcast_query_response_interval QUERY_RESPONSE_INTERVAL ] [ mcast_startup_query_interval STARTUP_QUERY_INTERVAL ] [ mcast_stats_enabled MCAST_STATS_ENABLED ] [ mcast_igmp_version IGMP_VERSION ] [ mcast_mld_version MLD_VERSION ] [ nf_call_iptables NF_CALL_IPTABLES ] [ nf_call_ip6tables NF_CALL_IP6TABLES
] [ nf_call_arptables NF_CALL_ARPTABLES ]`

ethernet bridge device

Linux中的Bridge和物理设备的网桥一样

*A Linux bridge behaves like a network switch. It forwards packets between interfaces that are connected to it. It's usually used for forwarding packets on routers, on gateways, or between VMs and network namespaces on a host. It also supports STP, VLAN filter, and multicast snooping.*

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/bridge.png)

==Use a bridge when you want to establish communication channels between VMs, containers, and your hosts.==

Here's how to create a bridge:

```
# ip link add br0 type bridge
# ip link set eth0 master br0
# ip link set tap1 master br0
# ip link set tap2 master br0
# ip link set veth1 master br0
```

This creates a bridge device named `br0` and sets two TAP devices (`tap1`, `tap2`), a VETH device (`veth1`), and a physical device (`eth0`) as its slaves, as shown in the diagram above.

如果link type 选择是 bridge，还支持如下几个常用的参数

参数太多这里不展开，具体可以使用 docker 虚拟出来的bridge interface 参考 man page

### 0x3.2 bond 

 bonding device

将不同的iface聚合成一个逻辑iface，用于LB

*The Linux bonding driver provides a method for aggregating multiple network interfaces into a single logical "bonded" interface. The behavior of the bonded interface depends on the mode; generally speaking, modes provide either hot standby or load balancing services.*

![Bonded interface](https://developers.redhat.com/blog/wp-content/uploads/2018/10/bond.png)

==Use a bonded interface when you want to increase your link speed or do a failover on your server.==

Here's how to create a bonded interface:

```
ip link add bond1 type bond miimon 100 mode active-backup
ip link set eth0 master bond1
ip link set eth1 master bond1
```

This creates a bonded interface named `bond1` with mode active-backup. For other modes, please see the [kernel documentation](https://www.kernel.org/doc/Documentation/networking/bonding.txt).

### 0x3.3 dummy

dummy network interface

并不会发送到网络中，只会在本地网络链处理

A dummy interface is entirely virtual like, for example, the loopback interface. The purpose of a dummy interface is to provide a device to route packets through without actually transmitting them.

*Use a dummy interface to make an inactive SLIP (Serial Line Internet Protocol) address look like a real address for local programs. Nowadays, a dummy interface is mostly used for testing and debugging.*

Here's how to create a dummy interface:

```
# ip link add dummy1 type dummy
# ip addr add 1.1.1.1/24 dev dummy1
# ip link set dummy1 up
```

### 0x3.4 hsr

high-availability seamless redundancy device

### 0x3.5 ifb

intermediate functional block device

The IFB (Intermediate Functional Block) driver supplies a device that allows the concentration of traffic from several sources and the shaping incoming traffic instead of dropping it.

Use an IFB interface when you want to queue and shape incoming traffic.

Here's how to create an IFB interface:

```
# ip link add ifb0 type ifb
# ip link set ifb0 up
# tc qdisc add dev ifb0 root sfq
# tc qdisc add dev eth0 handle ffff: ingress
# tc filter add dev eth0 parent ffff: u32 match u32 0 0 action mirred egress redirect dev ifb0
```

This creates an IFB device named `ifb0` and replaces the root qdisc scheduler with SFQ (Stochastic Fairness Queueing), which is a classless queueing scheduler. Then it adds an ingress qdisc scheduler on `eth0` and redirects all ingress traffic to `ifb0`.

### 0x3.6 ipoid

IP over Infiniband device

### 0x3.7 macvlan

syntax：` ip link add link DEVICE name NAME type { macvlan | macvtap } mode { private | vepa | bridge | passthru  [ nopromisc ] | source }`

virtual interface base on link layer address（MAC）

顾名思义VLAN是virutal LAN，macVLAN是MAC virutal LAN

*With VLAN, you can create multiple interfaces on top of a single one and filter packages based on a VLAN tag. With MACVLAN, you can create multiple interfaces with different Layer 2 (that is, Ethernet MAC) addresses on top of a single one.*

Before MACVLAN, if you wanted to connect to physical network from a VM or namespace, you would have needed to create TAP/VETH devices and attach one side to a bridge and attach a physical interface to the bridge on the host at the same time, as shown below.

[![Configuration before MACVLAN](https://developers.redhat.com/blog/wp-content/uploads/2018/10/br_ns.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/br_ns.png)

Now, with MACVLAN, you can bind a physical interface that is associated with a MACVLAN directly to namespaces, without the need for a bridge.

[![Configuration with MACVLAN](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan.png)

There are five MACVLAN types:

1. Private: doesn't allow communication between MACVLAN instances on the same physical interface, even if the external switch supports hairpin mode.

   [![Private MACVLAN configuration](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_01.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_01.png)

2. VEPA: data from one MACVLAN instance to the other on the same physical interface is transmitted over the physical interface. Either the attached switch needs to support hairpin mode or there must be a TCP/IP router forwarding the packets in order to allow communication.

   [![VEPA MACVLAN configuration](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_02.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_02.png)

3. Bridge: all endpoints are directly connected to each other with a simple bridge via the physical interface.

   [![Bridge MACVLAN configuration](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_03.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_03.png)

4. Passthru: allows a single VM to be connected directly to the physical interface.

   [![Passthru MACVLAN configuration](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_04.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_04.png)

5. Source: the source mode is used to filter traffic based on a list of allowed source MAC addresses to create MAC-based VLAN associations. Please see the [commit message](https://git.kernel.org/pub/scm/linux/kernel/git/davem/net.git/commit/?id=79cf79abce71).

The type is chosen according to different needs. Bridge mode is the most commonly used.

Use a MACVLAN when you want to connect directly to a physical network from containers.

Here's how to set up a MACVLAN:

```
# ip link add macvlan1 link eth0 type macvlan mode bridge
# ip link add macvlan2 link eth0 type macvlan mode bridge
# ip netns add net1
# ip netns add net2
# ip link set macvlan1 netns net1
# ip link set macvlan2 netns net2
```

This creates two new MACVLAN devices in bridge mode and assigns these two devices to two different namespaces.

如果link type 选择是 macvlan，还支持如下几个常用的参数

- `mode {private | vepa | bridge | passthru [nopromisc] | source}`

  具体模型参考上述图片

### 0x3.8 macvtap

syntax：` ip link add link DEVICE name NAME type { macvlan | macvtap } mode { private | vepa | bridge | passthru  [ nopromisc ] | source }`

virutal interface based on link layer address（MAC）and TAP

MACVTAP/IPVTAP is a new device driver meant to simplify virtualized bridged networking. When a MACVTAP/IPVTAP instance is created on top of a physical interface, the kernel also creates a character device/dev/tapX to be used just like a [TUN/TAP](https://en.wikipedia.org/wiki/TUN/TAP) device, which can be directly used by KVM/QEMU.

With MACVTAP/IPVTAP, you can replace the combination of TUN/TAP and bridge drivers with a single module:

[![MACVTAP/IPVTAP instance](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvtap.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvtap.png)

Typically, MACVLAN/IPVLAN is used to make both the guest and the host show up directly on the switch to which the host is connected. The difference between MACVTAP and IPVTAP is same as with MACVLAN/IPVLAN.

Here's how to create a MACVTAP instance:

```
# ip link add link eth0 name macvtap0 type macvtap
```

如果link type 选择是 macvlan，还支持如下几个常用的参数

- `mode {private | vepa | bridge | passthru [nopromisc] | source}`

  具体模型参考MACVLAN图片

### 0x3.9 vcan

virtual controller area network interface

Similar to the network loopback devices, the VCAN (virtual CAN) driver offers a virtual local CAN (Controller Area Network) interface, so users can send/receive CAN messages via a VCAN interface. CAN is mostly used in the automotive field nowadays.

For more CAN protocol information, please refer to the [kernel CAN documentation](https://www.kernel.org/doc/Documentation/networking/can.txt).

Use a VCAN when you want to test a CAN protocol implementation on the local host.

Here's how to create a VCAN:

```
# ip link add dev vcan1 type vcan
```

### 0x3.a vxcan

syntax：`ip link add DEVICE type {veth | vxcan [peer name NAME]}`

virtual controller area network tunnel interface

imilar to the VETH driver, a VXCAN (Virtual CAN tunnel) implements a local CAN traffic tunnel between two VCAN network devices. When you create a VXCAN instance, two VXCAN devices are created as a pair. When one end receives the packet, the packet appears on the device's pair and vice versa. VXCAN can be used for cross-namespace communication.

Use a VXCAN configuration when you want to send CAN message across namespaces.

Here's how to set up a VXCAN instance:

```
# ip netns add net1
# ip netns add net2
# ip link add vxcan1 netns net1 type vxcan peer name vxcan2 netns net2
```

***Note\***: VXCAN is not yet supported in Red Hat Enterprise Linux.

如果link type 选择是 veth，还支持如下几个常用的参数

- `peer name NAME`

  specifies the virtual pair device name of the VETH/VXCAN tunnel

### 0x3.b veth

syntax：`ip link add DEVICE type {veth | vxcan [peer name NAME]}`

virtual ethernet interface

The VETH (virtual Ethernet) device is a local Ethernet tunnel. Devices are created in pairs, as shown in the diagram below.

Packets transmitted on one device in the pair are immediately received on the other device. When either device is down, the link state of the pair is down.

[![Pair of VETH devices](https://developers.redhat.com/blog/wp-content/uploads/2018/10/veth.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/veth.png)

Use a VETH configuration when namespaces need to communicate to the main host namespace or between each other.

Here's how to set up a VETH configuration:

```
# ip netns add net1
# ip netns add net2
# ip link add veth1 netns net1 type veth peer name veth2 netns net2
```

This creates two namespaces, `net1` and `net2`, and a pair of VETH devices, and it assigns `veth1` to namespace `net1` and `veth2` to namespace `net2`. These two namespaces are connected with this VETH pair. Assign a pair of IP addresses, and you can ping and communicate between the two namespaces.

如果link type 选择是 veth，还支持如下几个常用的参数

- `peer name NAME`

  specifies the virtual pair device name of the VETH/VXCAN tunnel

### 0x3.c vlan

syntax：`ip link add link DEVICE name NAME type vlan [ protocol VLAN_PROTO ] id VLANID [ reorder_hdr { on | off } ] [ gvrp { on| off } ] [ mvrp { on | off } ] [ loose_binding { on | off } ][ ingress-qos-map QOS-MAP ] [ egress-qos-map QOS-MAP ]`

802.1 q tagged virtual LAN interface

A VLAN, aka virtual LAN, separates broadcast domains by adding tags to network packets. VLANs allow network administrators to group hosts under the same switch or between different switches.

The VLAN header looks like:

[![VLAN header](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan_01.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan_01.png)

Use a VLAN when you want to separate subnet in VMs, namespaces, or hosts.

Here's how to create a VLAN:

```
# ip link add link eth0 name eth0.2 type vlan id 2
# ip link add link eth0 name eth0.3 type vlan id 3
```

This adds VLAN 2 with name `eth0.2` and VLAN 3 with name `eth0.3`. The topology looks like this:

[![VLAN topology](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan.png)

***Note\***: When configuring a VLAN, you need to make sure the switch connected to the host is able to handle VLAN tags, for example, by setting the switch port to trunk mode.

如果link type 选择是 VLAN，还支持如下几个常用的参数

- `id VLANID`

  指定VLAN id，不能使用0 (保留数，表示所有的VLAN id)

- `reorder {on | off}`

  是否将 VLAN header 立即插入报文头，如果on表示不会立即插入(只有在网卡不支持 vlan offloading 才可以，可以使用`ethtool -k <phyiscal_device> | grep tx-valn-offload` 来校验)，之后再发往实际物理网卡时才会插入，reorder会accelerate tagging egress and hide VLAN header on ingress ，所以报文就会像普通的以太网报文。但是这也会造成抓包时混乱

- `loose_binding {on | off}`

  指定VLAN device 的 state 是否和物理网卡的state 一样

### 0x3.e vxlan

syntax：`ip link add DEVICE type vxlan id VNI [ dev PHYS_DEV  ] [ {group | remote } IPADDR ] [ local { IPADDR | any } ] [ ttl TTL] [ tos TOS ] [ flowlabel FLOWLABEL ] [ dstport PORT ] [ srcport MIN MAX ] [ [no]learning ] [ [no]proxy ] [ [no]rsc ] [[no]l2miss ] [ [no]l3miss ] [ [no]udpcsum ] [ [no]udp6zerocsumtx ] [ [no]udp6zerocsumrx ] [ ageing SECONDS ] [ maxaddress NUMBER ] [ [no]external ] [ gbp ] [ gpe ]`

virtual extended LAN

VXLAN (Virtual eXtensible Local Area Network) is a tunneling protocol designed to solve the problem of limited VLAN IDs (4,096) in IEEE 802.1q. It is described by [IETF RFC 7348](https://tools.ietf.org/html/rfc7348).

With a 24-bit segment ID, aka VXLAN Network Identifier (VNI), VXLAN allows up to 2^24 (16,777,216) virtual LANs, which is 4,096 times the VLAN capacity.

VXLAN encapsulates Layer 2 frames with a VXLAN header into a UDP-IP packet, which looks like this:

[![VXLAN encapsulates Layer 2 frames with a VXLAN header into a UDP-IP packet](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vxlan_01.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vxlan_01.png)

VXLAN is typically deployed in data centers on virtualized hosts, which may be spread across multiple racks.

[![Typical VXLAN deployment](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vxlan.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vxlan.png)

Here's how to use VXLAN:

```
# ip link add vx0 type vxlan id 100 local 1.1.1.1 remote 2.2.2.2 dev eth0 dstport 4789
```

For reference, you can read the [VXLAN kernel documentation](https://www.kernel.org/doc/Documentation/networking/vxlan.txt) or [this VXLAN introduction](https://vincent.bernat.ch/en/blog/2017-vxlan-linux).

如果link type 选择是 VXLAN，还支持如下几个常用的参数

- `id VNI`

  指定VXLAN id

- `dev PHYS_DEV`

  指定用于endpoint communication 的本端物理网卡

- `group IPADDR`

  specifies the multicaaset（多播） ip address to join，不能和 `remote`一起使用

- `remote IPADDR`

  speicifies the unicast（单播）destination IP address to use in outgoing packets when the destination link address is not know in the VXLAN device forwarding database，不能和`group`一起使用

- `local IPADDR`

  指定本端使用的IP

- `ttl`

  specifies the TTL value to use in outgoing packets

- `tos TOS`

  specifes the TOS value to use in outgoing packets

- `dstport PORT`

  指定使用VXLAN UDP 目的端口

- `srcport MIN MAX`

  指定使用VXLAN UDP 源端口的范围

- `[no]learning`

  是否会将未知的源端 MAC 和 IP 地址记录到VXLAN device forwarding database

### 0x3.f ipip

vitrual tunnel interface ipv6 over ipv4

### 0x3.g ipvlan

interface for L3(IPv4/IPv6) base VLANs

*IPVLAN is similar to MACVLAN with the difference being that the endpoints have the same MAC address.*

[![IPVLAN configuration](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan.png)

IPVLAN supports L2 and L3 mode. IPVLAN L2 mode acts like a MACVLAN in bridge mode. The parent interface looks like a bridge or switch.

[![IPVLAN L2 mode](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan_01.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan_01.png)

In IPVLAN L3 mode, the parent interface acts like a router and packets are routed between endpoints, which gives better scalability.

[![IPVLAN L3 mode](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan_02.png)](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan_02.png)

Regarding when to use an IPVLAN, the [IPVLAN kernel documentation](https://www.kernel.org/doc/Documentation/networking/ipvlan.txt) says that MACVLAN and IPVLAN "are very similar in many regards and the specific use case could very well define which device to choose. if one of the following situations defines your use case then you can choose to use ipvlan -
(a) The Linux host that is connected to the external switch / router has policy configured that allows only one mac per port.
(b) No of virtual devices created on a master exceed the mac capacity and puts the NIC in promiscuous mode and degraded performance is a concern.
(c) If the slave device is to be put into the hostile / untrusted network namespace where L2 on the slave could be changed / misused."

Here's how to set up an IPVLAN instance:

```
# ip netns add ns0
# ip link add name ipvl0 link eth0 type ipvlan mode l2
# ip link set dev ipvl0 netns ns0
```

This creates an IPVLAN device named `ipvl0` with mode L2, assigned to namespace `ns0`.

