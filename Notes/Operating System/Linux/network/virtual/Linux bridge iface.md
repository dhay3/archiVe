ref
[https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/networking/bridge.rst?h=v5.19-rc1](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/networking/bridge.rst?h=v5.19-rc1)
[https://wiki.linuxfoundation.org/networking/bridge](https://wiki.linuxfoundation.org/networking/bridge)
[https://wiki.debian.org/BridgeNetworkConnections](https://wiki.debian.org/BridgeNetworkConnections)
[https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/deployment_guide/s2-networkscripts-interfaces_network-bridge](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/deployment_guide/s2-networkscripts-interfaces_network-bridge)
## Digest
a bridge is a way to connect two ethernet segments together in a protocol independent way. Packets are forwarded based on ethernet address(MAC address)，rather than IP address(like a router)
bridge用于就两个以太网连接，不使用IP来转发包，而根据2层MAC地址来转发包(寻址)。可以通过bridge-utils(内含`brctl`，当前已无人在维护)或者iproute2来管理bridge设备
## Kernel Configuration
### cautions

1. make sure both network cards are set up and working properly. Don't set the IP address, and don't let the startup scripts run DHCP on the ethernet interfaces. The IP address needs to be set after the bridge has been configured

如果直接NIC桥接到bridge，IP会被清除。需要谨慎操作

2. 如果iface已经是SLAVE了，就不能被桥接。例如 eth0 eth1 是 bond0 的 SLAVE，就不能被用在bridge下
2. bridge配置的 IP 不会显示在 traceroute 中，bridge对3层网络来说是透明的
2. 为了可以桥接到bridge，NIC需要被设置成 promiscuous mode，所有从 bridge 来的流量都会被SLAVE收到，会导致占用大量内存可用带宽减小。这一点和物理设备的网桥是不同的
### brctl
需要通过`yum install bridge-utils`来安装，未添加bridge iface。不是持久配置
```
[vagrant@localhost ~]$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 52:54:00:4d:77:d3 brd ff:ff:ff:ff:ff:ff
4: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:89:8e:10 brd ff:ff:ff:ff:ff:ff
5: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:10:a5:6e brd ff:ff:ff:ff:ff:ff
    inet 10.0.4.15/24 brd 10.0.4.255 scope global noprefixroute dynamic eth2
       valid_lft 84516sec preferred_lft 84516sec
    inet6 fe80::320c:c66a:c15a:7a67/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
9: bond0: <NO-CARRIER,BROADCAST,MULTICAST,MASTER,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 0e:d3:d7:8d:c8:4d brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.1/24 brd 192.168.1.255 scope global noprefixroute bond0
       valid_lft forever preferred_lft forever
```
使用brctl添加br0
```
[vagrant@localhost ~]$ sudo brctl addbr br0
[vagrant@localhost ~]$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 52:54:00:4d:77:d3 brd ff:ff:ff:ff:ff:ff
4: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:89:8e:10 brd ff:ff:ff:ff:ff:ff
5: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:10:a5:6e brd ff:ff:ff:ff:ff:ff
    inet 10.0.4.15/24 brd 10.0.4.255 scope global noprefixroute dynamic eth2
       valid_lft 84437sec preferred_lft 84437sec
    inet6 fe80::320c:c66a:c15a:7a67/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
9: bond0: <NO-CARRIER,BROADCAST,MULTICAST,MASTER,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 0e:d3:d7:8d:c8:4d brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.1/24 brd 192.168.1.255 scope global noprefixroute bond0
       valid_lft forever preferred_lft forever
10: br0: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether ce:0d:0e:da:6d:75 brd ff:ff:ff:ff:ff:ff

[vagrant@localhost ~]$ sudo brctl addif br0 eth1
[vagrant@localhost ~]$ sudo brctl addif br0 eth1
[vagrant@localhost ~]$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master br0 state UP group default qlen 1000
    link/ether 52:54:00:4d:77:d3 brd ff:ff:ff:ff:ff:ff
4: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master br0 state UP group default qlen 1000
    link/ether 08:00:27:89:8e:10 brd ff:ff:ff:ff:ff:ff
5: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:10:a5:6e brd ff:ff:ff:ff:ff:ff
    inet 10.0.4.15/24 brd 10.0.4.255 scope global noprefixroute dynamic eth2
       valid_lft 84302sec preferred_lft 84302sec
    inet6 fe80::320c:c66a:c15a:7a67/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
9: bond0: <NO-CARRIER,BROADCAST,MULTICAST,MASTER,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 0e:d3:d7:8d:c8:4d brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.1/24 brd 192.168.1.255 scope global noprefixroute bond0
       valid_lft forever preferred_lft forever
10: br0: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 08:00:27:89:8e:10 brd ff:ff:ff:ff:ff:ff
```
eth0和eth1都被桥接到了br0，因为没有配置IP，br0现在还不具备TCP/IP的功能
```
[vagrant@localhost ~]$ brctl show
bridge name     bridge id               STP enabled     interfaces
br0             8000.080027898e10       no              eth0
                                                        eth1
```
使用 iproute2 设置 IP 并 启用 iface
```
[vagrant@localhost ~]$ sudo ip address change dev br0 192.168.1.1/24
[vagrant@localhost ~]$ sudo ip link set br0 up

10: br0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 08:00:27:89:8e:10 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.1/24 scope global br0
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe89:8e10/64 scope link 
       valid_lft forever preferred_lft forever
```
通过 bridge IP ping 测宿主机
```
[vagrant@localhost ~]$ brctl showmacs br0
port no mac addr                is local?       ageing timer
  2     08:00:27:89:8e:10       yes                0.00
  2     08:00:27:89:8e:10       yes                0.00
  # 这里从port2学到宿主机MAC
  2     0a:00:27:00:00:1a       no                 0.12
  1     52:54:00:4d:77:d3       yes                0.00
  1     52:54:00:4d:77:d3       yes                0.00
  
  
  Ethernet adapter VirtualBox Host-Only Network #3:

Connection-specific DNS Suffix  . :
Description . . . . . . . . . . . : VirtualBox Host-Only Ethernet Adapter #3
Physical Address. . . . . . . . . : 0A-00-27-00-00-1A
DHCP Enabled. . . . . . . . . . . : No
Autoconfiguration Enabled . . . . : Yes
Link-local IPv6 Address . . . . . : fe80::9188:a4a3:6787:254d%26(Preferred)
IPv4 Address. . . . . . . . . . . : 192.168.1.2(Preferred)
Subnet Mask . . . . . . . . . . . : 255.255.255.0
Default Gateway . . . . . . . . . :
DHCPv6 IAID . . . . . . . . . . . : 436863015
DHCPv6 Client DUID. . . . . . . . : 00-01-00-01-27-FF-28-FC-E8-6A-64-87-7E-63
DNS Servers . . . . . . . . . . . : fec0:0:0:ffff::1%1
fec0:0:0:ffff::2%1
fec0:0:0:ffff::3%1
NetBIOS over Tcpip. . . . . . . . : Enabled
```
### iproute2
非持久配置
```
[vagrant@localhost ~]$ sudo ip link add br0 type bridge
[vagrant@localhost ~]$ sudo ip link set eth0 master br0
[vagrant@localhost ~]$ sudo ip link set eth1 master br0
[vagrant@localhost ~]$ sudo ip addr change dev br0 192.168.1.1/24
[vagrant@localhost ~]$ sudo ip link set br0 up
```
### initscripts
```
[root@localhost network-scripts]# cat ifcfg-bridge0
ipaddr=192.168.1.1
netmask=255.255.255.0
onboot=yes
bootproto=none
type=bridge

[root@localhost network-scripts]# cat ifcfg-eth0
DEVICE=eth0
BOOTPROTO=none
ONBOOT=yes
SLAVE=yes
bridge=bridge0

[root@localhost network-scripts]# cat ifcfg-eth1
DEVICE=eth1
SLAVE=yes
BRIDGE=bridge0
ONBOOT=yes
BOOTPROTO=none
```
