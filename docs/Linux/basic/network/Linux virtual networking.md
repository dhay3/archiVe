# Linux virtual networking

ref：

https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking#

## Bridge

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

## Bonded interface

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

## Team device

和 bonded interface 类似，但是有额外特性

Similar a bonded interface, the purpose of a team device is to provide a mechanism to group multiple NICs (ports) into one logical one (teamdev) at the L2 layer.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/team.png)

The main thing to realize is that a team device is not trying to replicate or mimic a bonded interface. What it does is to solve the same problem using a different approach, using, for example, a lockless (RCU) TX/RX path and modular design.

But there are also some functional differences between a bonded interface and a team. For example, a team supports LACP load-balancing, NS/NA (IPV6) link monitoring, D-Bus interface, etc., which are absent in bonding. For further details about the differences between bonding and team, see [Bonding vs. Team features](https://github.com/jpirko/libteam/wiki/Bonding-vs.-Team-features).

==Use a team when you want to use some features that bonding doesn't provide.==

Here's how to create a team:

```
# teamd -o -n -U -d -t team0 -c '{"runner": {"name": "activebackup"},"link_watch": {"name": "ethtool"}}'
# ip link set eth0 down
# ip link set eth1 down
# teamdctl team0 port add eth0
# teamdctl team0 port add eth1
```

This creates a team interface named `team0` with mode `active-backup`, and it adds `eth0` and `eth1` as `team0`'s sub-interfaces.

A new driver called [net_failover](https://www.kernel.org/doc/html/latest/networking/net_failover.html) has been added to Linux recently. It's another failover master net device for virtualization and manages a primary ([passthru/VF [Virtual Function\]](https://wiki.libvirt.org/page/Networking#PCI_Passthrough_of_host_network_devices) device) slave net device and a standby (the original paravirtual interface) slave net device.

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/net_failover.png)

## VLAN

A VLAN, aka virtual LAN, ==separates broadcast domains== by adding tags to network packets. VLANs allow network administrators to group hosts under the same switch or between different switches.

The VLAN header looks like:

![VLAN header](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan_01.png)

Use a VLAN when you want to separate subnet in VMs, namespaces, or hosts.

Here's how to create a VLAN:

```
# ip link add link eth0 name eth0.2 type vlan id 2
# ip link add link eth0 name eth0.3 type vlan id 3
```

This adds VLAN 2 with name `eth0.2` and VLAN 3 with name `eth0.3`. The topology looks like this:

![](https://developers.redhat.com/blog/wp-content/uploads/2018/10/vlan.png)

***Note\***: When configuring a VLAN, you need to make sure the switch connected to the host is able to handle VLAN tags, for example, by setting the switch port to trunk mode.

## VXLAN

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

## VETH

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

## Dummy interface



A dummy interface is entirely virtual like, for example, the loopback interface. The purpose of a dummy interface is to provide a device to route packets through without actually transmitting them.

*Use a dummy interface to make an inactive SLIP (Serial Line Internet Protocol) address look like a real address for local programs. Nowadays, a dummy interface is mostly used for testing and debugging.*

Here's how to create a dummy interface:

```
# ip link add dummy1 type dummy
# ip addr add 1.1.1.1/24 dev dummy1
# ip link set dummy1 up
```