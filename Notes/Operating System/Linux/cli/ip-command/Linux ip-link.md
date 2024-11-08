# Linux IP link

ref：

[https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking#](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking#)

https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/networking?h=v5.17

> 通过这种方式增加的device不会持久存在，如需持久存在需写入配置文件

## 0x1 Digest

查看或配置设备上L2信息

这里的 device 表示的是设备，包含虚拟设备、物理网卡、虚拟网卡等

## 0x2 EBNF

```
                 TYPE := [ bridge | bond | can | dummy | hsr | ifb | ipoib | macvlan | macvtap | vcan | vxcan | veth |
                         vlan | vxlan | ip6tnl | ipip | sit | gre | gretap | erspan | ip6gre | ip6gretap | ip6erspan |
                         vti | nlmon | ipvlan | ipvtap | lowpan | geneve | bareudp | vrf | macsec | netdevsim | rmnet |
                         xfrm ]

                 ETYPE := [ TYPE | bridge_slave | bond_slave ]

                 VFVLAN-LIST := [ VFVLAN-LIST ] VFVLAN

                 VFVLAN := [ vlan VLANID [ qos VLAN-QOS ] [ proto VLAN-PROTO ] ]

```

## 0x3 Commands

### ip link show
syntax：`ip link show [ DEVICE | group GROUP ] [ up ] [ master DEVICE ] [ type ETYPE ] [ vrf NAME ]`
查询device link layer 相关的信息

- `dev NAME `

查看指定device 的信息

- `group GROUP`

查看指定group 的信息

- `up`

只查看UP的 link device

- `master DEVICE`

查看master device的slaves
```bash
[root@localhost vagrant]# ip link set dummy0 master br0
[root@localhost vagrant]# ip link show master br0
3: dummy0: <BROADCAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc noqueue master br0 state UNKNOWN mode DEFAULT group default qlen 1000
    link/ether 6a:e9:47:b7:06:20 brd ff:ff:ff:ff:ff:ff
```

- `type TYPE`

查看指定link type device
### ip link add

syntax：`ip link add [link DEVICE][name]NAME [txqueuelen PACKETS] [address LLADDR] [broadcast LLADDR] [mtu MTU] [index IDX] [numtxqueues QUEUE_COUNT] [numrxqueues QUEUE_COUNT] type TYPE [ARGS]`

添加virtual link

-  `DEVICE`
指定实际的物理网卡，部分link type 无须指定实际的物理网卡 
-  `NAME`
虚拟设备的名字 
-  `TYPE`
指定虚拟设备的link type，可以是如下几种 
-  `numtxqueues <queue_count>`
指定发送数据的队列数 
-  `numrxqueues <queue_count>`
指定接受数据的队列数 
-  `gso_max_size <bytes>`
指定接受packet的大小 
-  `gso_max_segs <segments>`
指定接受segment(不包含包头)的大小 

例如添加一个

### ip link delete

syntax：`ip link delete dev DEVICE [group GROUP] type [TYPE]`

删除virtual link

-  `dev DEVICE`
指定需要删除的 virtual link device，只能是virtual link device 
-  `group GROUP`
指定需要删除的group，group 0 不允许删除 

例如删除 dummy virutal device

```
ip link delete dev dummy0 type dummy
```
### ip link set
syntax：`ip link set {DEVICE | group GROUP}[options]`

- `dev DEVICE`

指定需要操作的网络设备

- `group GROUP`

以组名指定一组设备
#### options
##### gernal options

- `up | down`

  设置device 的 state 为 UP 或 DOWN，如果状态显示UNKOWN

- `arp on | arp off`

  change the NOARP flag on the device

- `multicast on | multicast off`

  change the MULTICAST flag on the device

- `protodown on | protodown off`

  change the PROTODOWN state on 
  如果在改device上检测到protocol 错误，交换机会对互联的端口进行物理关闭(shutdown)

- `dynamic on | dynamic off`

  change the DYNAMIC flag on the device
  当 interface goes down 网络地址会变化

- `name NAME`

  change the name of the device
  如果设备正在运行或已经有配置网络地址的话，不推荐使用该选项

- `txqueuelen NUMBER | txqlem NUMBER`

  change the transimit queue length of the device

- `mtu NUMBER`

  change the MTU of the device

- `address LLADDRESS(link layer address)`

  change the station address of the interface
  修改MAC地址

  ```
  [root@localhost vagrant]# ip link set dummy0 address 6a:e9:47:b7:06:20
  [root@localhost vagrant]# ip a show dev dummy0
  3: dummy0: <BROADCAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN group default qlen 1000
      link/ether 6a:e9:47:b7:06:20 brd ff:ff:ff:ff:ff:ff
      inet6 fe80::68e9:47ff:feb7:614/64 scope link 
         valid_lft forever preferred_lft forever
  ```

- `broadcast LLADDRESS | brd LLADDRESS | peer LLADDRESS`

  change the link layer broadcast address or the peer address when the interface is point to point
  修改广播地址

- `netns NETNSNAME | PID`

  move the device to the network namespace associated with name NETNSNAME or process PID
  将 device 关联到指定 network namespace 或 pid
  一些 device 不允许修改name space(name space local device)，例如：loopback，bridge，wireless(可以使用 `ethtool -k DEVICE | grep netns-local`来查看是否是network namespace local device)

- `alias NAME`

  为 device 设置别名

- `group GROUP`

  设置 device 的group，可以使用的group值查看`/etc/iproute2/group`

- `master DEVICE`

  set master device of the device，如果设置了master会显示

  ```
  3: dummy0: <BROADCAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc noqueue master br0 state UNKNOWN group default qlen 1000
      link/ether 6a:e9:47:b7:06:20 brd ff:ff:ff:ff:ff:ff
      inet6 fe80::68e9:47ff:feb7:620/64 scope link 
         valid_lft forever preferred_lft forever
  ```

- `nomaster`

  unset master device of the device

- `link-netnsid`

  set peer netnsid for a cross-netns interface

- `type ETYPE TYPE_ARGS`

  change type-specific settings

##### virutal device option

- `vf`

  用于设置虚拟网络设备的参数，具体值查看man page

- `Bridge Slave Support`

  如果 link type 是 bridge 且是 master 的 slave，还支持如下参数

   - fdb_flush 

     flush bridge slave's fdb dynamic entries

   - state STATE

     set port state，STATE 的值可以是0(disable), 1(listening), 2(learning), 3(forwarding), 4(blocking)

   - prioritiy PRIO

     set port priority，值可以是 0 - 64

   - cost COST

     set port cost，值可以是 1 - 65535

   - guard {on | off}

     blocking incoming BPDU packets on this port

   - learning { on | off}

     allow MAC address learning on this port

- `Bonding Slave Support`

  如果 link type 是 bond 且是master 的 slave，还支持如下参数

   - queue_id ID

     设置 bond slave 的 queue ID

## 0x3 Link type

以下均为虚拟设备（不是虚拟网卡）类型

### bridge

syntax：`ip link add DEVICE type bridge [ ageing_time AGEING_TIME ] [ group_fwd_mask MASK ] [ group_address ADDRESS ] [ forward_delay FORWARD_DELAY ] [ hello_time HELLO_TIME ] [ max_age MAX_AGE ] [ stp_state STP_STATE ] [ priority PRIORITY ] [ vlan_filtering VLAN_FILTERING ] [ vlan_protocol VLAN_PROTOCOL ] [ vlan_default_pvid VLAN_DEFAULT_PVID ] [ vlan_stats_enabled VLAN_STATS_ENABLED ] [ mcast_snooping MULTICAST_SNOOPING ] [ mcast_router MULTICAST_ROUTER ] [ mcast_query_use_ifaddr MCAST_QUERY_USE_IFADDR ] [ mcast_querier MULTICAST_QUERIER ] [ mcast_hash_elasticity HASH_ELASTICITY ] [ mcast_hash_max HASH_MAX ] [ mcast_last_member_count LAST_MEMBER_COUNT ] [ mcast_startup_query_count STARTUP_QUERY_COUNT ] [ mcast_last_member_interval LAST_MEMBER_INTERVAL ] [ mcast_mem‐ bership_interval MEMBERSHIP_INTERVAL ] [ mcast_querier_interval QUERIER_INTERVAL ] [ mcast_query_interval QUERY_INTERVAL ] [ mcast_query_response_interval QUERY_RESPONSE_INTERVAL ] [ mcast_startup_query_interval STARTUP_QUERY_INTERVAL ] [ mcast_stats_enabled MCAST_STATS_ENABLED ] [ mcast_igmp_version IGMP_VERSION ] [ mcast_mld_version MLD_VERSION ] [ nf_call_iptables NF_CALL_IPTABLES ] [ nf_call_ip6tables NF_CALL_IP6TABLES ] [ nf_call_arptables NF_CALL_ARPTABLES ]`

ethernet bridge device

Linux中的Bridge和物理设备的网桥一样

_A Linux bridge behaves like a network switch. It forwards packets between interfaces that are connected to it. It's usually used for forwarding packets on routers, on gateways, or between VMs and network namespaces on a host. It also supports STP, VLAN filter, and multicast snooping._

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/bridge.png#crop=0&crop=0&crop=1&crop=1&id=KMNSR&originHeight=439&originWidth=654&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Use a bridge when you want to establish communication channels between VMs, containers, and your hosts.

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

### bond

bonding device

将不同的iface聚合成一个逻辑iface，用于LB

_The Linux bonding driver provides a method for aggregating multiple network interfaces into a single logical "bonded" interface. The behavior of the bonded interface depends on the mode; generally speaking, modes provide either hot standby or load balancing services._

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/bond.png#crop=0&crop=0&crop=1&crop=1&id=qBzeD&originHeight=270&originWidth=319&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Use a bonded interface when you want to increase your link speed or do a failover on your server.

Here's how to create a bonded interface:

```
ip link add bond1 type bond miimon 100 mode active-backup
ip link set eth0 master bond1
ip link set eth1 master bond1
```

This creates a bonded interface named `bond1` with mode active-backup. For other modes, please see the [kernel documentation](https://www.kernel.org/doc/Documentation/networking/bonding.txt).

### dummy

dummy network interface

并不会发送到网络中，只会在本地网络链处理

A dummy interface is entirely virtual like, for example, the loopback interface. The purpose of a dummy interface is to provide a device to route packets through without actually transmitting them.

_Use a dummy interface to make an inactive SLIP (Serial Line Internet Protocol) address look like a real address for local programs. Nowadays, a dummy interface is mostly used for testing and debugging._

Here's how to create a dummy interface:

```
# ip link add dummy1 type dummy
# ip addr add 1.1.1.1/24 dev dummy1
# ip link set dummy1 up
```

### hsr

high-availability seamless redundancy device

### ifb

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

### ipoid

IP over Infiniband device

### macvlan

syntax：`ip link add link DEVICE name NAME type { macvlan | macvtap } mode { private | vepa | bridge | passthru [ nopromisc ] | source }`

virtual interface base on link layer address（MAC）

顾名思义VLAN是virutal LAN，macVLAN是MAC virutal LAN

_With VLAN, you can create multiple interfaces on top of a single one and filter packages based on a VLAN tag. With MACVLAN, you can create multiple interfaces with different Layer 2 (that is, Ethernet MAC) addresses on top of a single one._

Before MACVLAN, if you wanted to connect to physical network from a VM or namespace, you would have needed to create TAP/VETH devices and attach one side to a bridge and attach a physical interface to the bridge on the host at the same time, as shown below.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/br_ns.png#crop=0&crop=0&crop=1&crop=1&id=uOtCH&originHeight=376&originWidth=436&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Now, with MACVLAN, you can bind a physical interface that is associated with a MACVLAN directly to namespaces, without the need for a bridge.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan.png#crop=0&crop=0&crop=1&crop=1&id=dY0cf&originHeight=372&originWidth=439&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

There are five MACVLAN types:

1.  Private: doesn't allow communication between MACVLAN instances on the same physical interface, even if the external switch supports hairpin mode.
![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_01.png#crop=0&crop=0&crop=1&crop=1&id=O9342&originHeight=367&originWidth=428&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=) 
1.  VEPA: data from one MACVLAN instance to the other on the same physical interface is transmitted over the physical interface. Either the attached switch needs to support hairpin mode or there must be a TCP/IP router forwarding the packets in order to allow communication.
![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_02.png#crop=0&crop=0&crop=1&crop=1&id=asKBT&originHeight=372&originWidth=433&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=) 
1.  Bridge: all endpoints are directly connected to each other with a simple bridge via the physical interface.
![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_03.png#crop=0&crop=0&crop=1&crop=1&id=Ox8XF&originHeight=369&originWidth=435&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=) 
1.  Passthru: allows a single VM to be connected directly to the physical interface.
![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvlan_04.png#crop=0&crop=0&crop=1&crop=1&id=F0Qpw&originHeight=318&originWidth=291&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=) 
1.  Source: the source mode is used to filter traffic based on a list of allowed source MAC addresses to create MAC-based VLAN associations. Please see the [commit message](https://git.kernel.org/pub/scm/linux/kernel/git/davem/net.git/commit/?id=79cf79abce71). 

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

-  `mode {private | vepa | bridge | passthru [nopromisc] | source}`
具体模型参考上述图片 

### macvtap

syntax：`ip link add link DEVICE name NAME type { macvlan | macvtap } mode { private | vepa | bridge | passthru [ nopromisc ] | source }`

virutal interface based on link layer address（MAC）and TAP

MACVTAP/IPVTAP is a new device driver meant to simplify virtualized bridged networking. When a MACVTAP/IPVTAP instance is created on top of a physical interface, the kernel also creates a character device/dev/tapX to be used just like a [TUN/TAP](https://en.wikipedia.org/wiki/TUN/TAP) device, which can be directly used by KVM/QEMU.

With MACVTAP/IPVTAP, you can replace the combination of TUN/TAP and bridge drivers with a single module:

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/macvtap.png#crop=0&crop=0&crop=1&crop=1&id=OhX0I&originHeight=375&originWidth=436&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Typically, MACVLAN/IPVLAN is used to make both the guest and the host show up directly on the switch to which the host is connected. The difference between MACVTAP and IPVTAP is same as with MACVLAN/IPVLAN.

Here's how to create a MACVTAP instance:

```
# ip link add link eth0 name macvtap0 type macvtap
```

如果link type 选择是 macvlan，还支持如下几个常用的参数

-  `mode {private | vepa | bridge | passthru [nopromisc] | source}`
具体模型参考MACVLAN图片 

### vcan

virtual controller area network interface

Similar to the network loopback devices, the VCAN (virtual CAN) driver offers a virtual local CAN (Controller Area Network) interface, so users can send/receive CAN messages via a VCAN interface. CAN is mostly used in the automotive field nowadays.

For more CAN protocol information, please refer to the [kernel CAN documentation](https://www.kernel.org/doc/Documentation/networking/can.txt).

Use a VCAN when you want to test a CAN protocol implementation on the local host.

Here's how to create a VCAN:

```
# ip link add dev vcan1 type vcan
```

### vxcan

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

**_Note*_: VXCAN is not yet supported in Red Hat Enterprise Linux.

如果link type 选择是 veth，还支持如下几个常用的参数

-  `peer name NAME`
specifies the virtual pair device name of the VETH/VXCAN tunnel 

### veth

syntax：`ip link add DEVICE type {veth | vxcan [peer name NAME]}`

virtual ethernet interface

The VETH (virtual Ethernet) device is a local Ethernet tunnel. Devices are created in pairs, as shown in the diagram below.

Packets transmitted on one device in the pair are immediately received on the other device. When either device is down, the link state of the pair is down.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/veth.png#crop=0&crop=0&crop=1&crop=1&id=UuLJi&originHeight=367&originWidth=437&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Use a VETH configuration when namespaces need to communicate to the main host namespace or between each other.

Here's how to set up a VETH configuration:

```
# ip netns add net1
# ip netns add net2
# ip link add veth1 netns net1 type veth peer name veth2 netns net2
```

This creates two namespaces, `net1` and `net2`, and a pair of VETH devices, and it assigns `veth1` to namespace `net1` and `veth2` to namespace `net2`. These two namespaces are connected with this VETH pair. Assign a pair of IP addresses, and you can ping and communicate between the two namespaces.

如果link type 选择是 veth，还支持如下几个常用的参数

-  `peer name NAME`
specifies the virtual pair device name of the VETH/VXCAN tunnel 

### vlan

syntax：`ip link add link DEVICE name NAME type vlan [ protocol VLAN_PROTO ] id VLANID [ reorder_hdr { on | off } ] [ gvrp { on| off } ] [ mvrp { on | off } ] [ loose_binding { on | off } ][ ingress-qos-map QOS-MAP ] [ egress-qos-map QOS-MAP ]`

802.1 q tagged virtual LAN interface

A VLAN, aka virtual LAN, separates broadcast domains by adding tags to network packets. VLANs allow network administrators to group hosts under the same switch or between different switches.

The VLAN header looks like:

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan_01.png#crop=0&crop=0&crop=1&crop=1&id=aQVNK&originHeight=351&originWidth=891&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Use a VLAN when you want to separate subnet in VMs, namespaces, or hosts.

Here's how to create a VLAN:

```
# ip link add link eth0 name eth0.2 type vlan id 2
# ip link add link eth0 name eth0.3 type vlan id 3
```

This adds VLAN 2 with name `eth0.2` and VLAN 3 with name `eth0.3`. The topology looks like this:

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan.png#crop=0&crop=0&crop=1&crop=1&id=vjTGl&originHeight=282&originWidth=283&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

**_Note*_: When configuring a VLAN, you need to make sure the switch connected to the host is able to handle VLAN tags, for example, by setting the switch port to trunk mode.

如果link type 选择是 VLAN，还支持如下几个常用的参数

-  `id VLANID`
指定VLAN id，不能使用0 (保留数，表示所有的VLAN id) 
-  `reorder {on | off}`
是否将 VLAN header 立即插入报文头，如果on表示不会立即插入(只有在网卡不支持 vlan offloading 才可以，可以使用`ethtool -k <phyiscal_device> | grep tx-valn-offload` 来校验)，之后再发往实际物理网卡时才会插入，reorder会accelerate tagging egress and hide VLAN header on ingress ，所以报文就会像普通的以太网报文。但是这也会造成抓包时混乱 
-  `loose_binding {on | off}`
指定VLAN device 的 state 是否和物理网卡的state 一样 

### vxlan

syntax：`ip link add DEVICE type vxlan id VNI [ dev PHYS_DEV ] [ {group | remote } IPADDR ] [ local { IPADDR | any } ] [ ttl TTL] [ tos TOS ] [ flowlabel FLOWLABEL ] [ dstport PORT ] [ srcport MIN MAX ] [ [no]learning ] [ [no]proxy ] [ [no]rsc ] [[no]l2miss ] [ [no]l3miss ] [ [no]udpcsum ] [ [no]udp6zerocsumtx ] [ [no]udp6zerocsumrx ] [ ageing SECONDS ] [ maxaddress NUMBER ] [ [no]external ] [ gbp ] [ gpe ]`

virtual extended LAN

VXLAN (Virtual eXtensible Local Area Network) is a tunneling protocol designed to solve the problem of limited VLAN IDs (4,096) in IEEE 802.1q. It is described by [IETF RFC 7348](https://tools.ietf.org/html/rfc7348).

With a 24-bit segment ID, aka VXLAN Network Identifier (VNI), VXLAN allows up to 2^24 (16,777,216) virtual LANs, which is 4,096 times the VLAN capacity.

VXLAN encapsulates Layer 2 frames with a VXLAN header into a UDP-IP packet, which looks like this:

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vxlan_01.png#crop=0&crop=0&crop=1&crop=1&id=CEMV8&originHeight=276&originWidth=983&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

VXLAN is typically deployed in data centers on virtualized hosts, which may be spread across multiple racks.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vxlan.png#crop=0&crop=0&crop=1&crop=1&id=DFJvV&originHeight=301&originWidth=813&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

Here's how to use VXLAN:

```
# ip link add vx0 type vxlan id 100 local 1.1.1.1 remote 2.2.2.2 dev eth0 dstport 4789
```

For reference, you can read the [VXLAN kernel documentation](https://www.kernel.org/doc/Documentation/networking/vxlan.txt) or [this VXLAN introduction](https://vincent.bernat.ch/en/blog/2017-vxlan-linux).

如果link type 选择是 VXLAN，还支持如下几个常用的参数

-  `id VNI`
指定VXLAN id 
-  `dev PHYS_DEV`
指定用于endpoint communication 的本端物理网卡 
-  `group IPADDR`
specifies the multicaaset（多播） ip address to join，不能和 `remote`一起使用 
-  `remote IPADDR`
speicifies the unicast（单播）destination IP address to use in outgoing packets when the destination link address is not know in the VXLAN device forwarding database，不能和`group`一起使用 
-  `local IPADDR`
指定本端使用的IP 
-  `ttl`
specifies the TTL value to use in outgoing packets 
-  `tos TOS`
specifes the TOS value to use in outgoing packets 
-  `dstport PORT`
指定使用VXLAN UDP 目的端口 
-  `srcport MIN MAX`
指定使用VXLAN UDP 源端口的范围 
-  `[no]learning`
是否会将未知的源端 MAC 和 IP 地址记录到VXLAN device forwarding database 

### ipip

vitrual tunnel interface ipv6 over ipv4

### ipvlan

interface for L3(IPv4/IPv6) base VLANs

_IPVLAN is similar to MACVLAN with the difference being that the endpoints have the same MAC address._

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan.png#crop=0&crop=0&crop=1&crop=1&id=cc1Gu&originHeight=373&originWidth=861&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

IPVLAN supports L2 and L3 mode. IPVLAN L2 mode acts like a MACVLAN in bridge mode. The parent interface looks like a bridge or switch.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan_01.png#crop=0&crop=0&crop=1&crop=1&id=ZWVfJ&originHeight=367&originWidth=438&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

In IPVLAN L3 mode, the parent interface acts like a router and packets are routed between endpoints, which gives better scalability.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/ipvlan_02.png#crop=0&crop=0&crop=1&crop=1&id=xhcJF&originHeight=373&originWidth=437&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=)

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
## 0x5 Master/Slave
[https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/networking_guide/sec-team-understanding_the_default_behavior_of_master_and_slave_interfaces](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/networking_guide/sec-team-understanding_the_default_behavior_of_master_and_slave_interfaces)

