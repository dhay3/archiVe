ref
[https://docs.kernel.org/networking/bonding.html?highlight=bond](https://docs.kernel.org/networking/bonding.html?highlight=bond)
[https://blog.51cto.com/liuqun/2044047](https://blog.51cto.com/liuqun/2044047)
## Digest
The linux bonding driver provides a method for aggregating multiple network interfaces into a single logical bonded interface. The behavior of the bonded interfaces depends upon the mode; generally speaking, modes provide either hot standy or load balancing services.
bond iface 是逻辑聚合的接口(聚合的接口需要是实际存在的nic)，通常被用来 hot standy 或者 load balancing
```
λ /lib/modules/5.10.archclear.x86_64/ cat modules.networking | grep bon
bonding.ko
```
如果不支持 bonding，可以参考kernel doc 安装模块
### requriments
现在主流的 kernel 版本已经编译了 bonding driver 模块。在使用`/etc/init.d/network restart`重启时，如果有bond设备，会自动加载bonding模块。如果bonding 模块被删除了，bond设备也会被删除
## Bonding Dirver Options
目前通过 iproute2 来管理和配置bonding iface，这里只介绍 mode，具体参数查看 kernel doc
bond 设备支持如下6种调度算法，默认 balance-rr

- balance-rr or 0

  round robin
  this mode provides load balancing and fault tolerance

- active-backup or 1

  only one slave in the bond is active.A different salve becomes active if, and only if the active slave fails
  this mode provides fault tolerance

- balance-xor or 2

  hash 一致
  this mode provides load balancing and fault tolerance

- broadcast or 3

  transimits everything on all slave interfaces
  this mode provides fault tolerance

- 802.3ad or 4

  802.3ad dynamic link aggregation

- balance-tlb or 5
- balance-alb or 6
## Configuring Bonding Devices
bonding 可以由多种方式配置

1. your distro's network initialization scripts，distro generally use one of three packages for the network initialization scripts：initscripts，sysconfig or interfaces
1. manually using either iproute2
1. the sysfs interfaces

如何判断使用的是那个package
```
if exsit /etc/network
  use interface
else rpm -qf /sbin/ifup
  use initscripts or sysconfig
   if grep ifenslave /sbin/ifup returns any matches
       support for bonding 
```
### Configuration with sysconfig
> 需要注意的是使用`/sbin/ifdown`会把bonding moudle 删除掉，需要将改文件给屏蔽掉

suse linux 通常使用sysconfig，网络配置文件通常在`/etc/sysconfig/network`
第一步先修改需要聚合的 slave 配置，将BOOTPROTO置为none，STARTMODE置为off，其他参数不动
```
BOOTPROTO='none'
STARTMODE='off'
USERCTL='no'
UNIQUE='XNzu.WeZGOGF+4wE'
_nm_name='bus-pci-0001:61:01.0'
```
第二步，新增一个 bonding device 配置文件，以 ifcfg-bondX 命名(X从0开始，例如ifcfg-bond0)
```
BOOTPROTO="static"
BROADCAST="10.0.2.255"
IPADDR="10.0.2.10"
NETMASK="255.255.0.0"
NETWORK="10.0.2.0"
REMOTE_IPADDR=""
STARTMODE="onboot"
BONDING_MASTER="yes"
BONDING_MODULE_OPTS="mode=active-backup miimon=100"
BONDING_SLAVE0="eth0"
BONDING_SLAVE1="bus-pci-0000:06:08.1"
```
Replace the sample `BROADCAST`, `IPADDR`, `NETMASK` and `NETWORK` values with the appropriate values for your network.

- STARTMODE的值可以使用如下几种
| onboot | The device is started at boot time. If you’re not sure, this is probably what you want. |
| --- | --- |
| manual | The device is started only when ifup is called manually. Bonding devices may be configured this way if you do not wish them to start automatically at boot for some reason. |
| hotplug | The device is started by a hotplug event. This is not a valid choice for a bonding device. |
| off or | The device configuration is ignored. |
| ignore |  |

- `BONDING_MASTER='yes'`

  表示当前设备是 bonding master device

- ` BONDING_MODULE_OPTS`

  表示当前设备使用的bonding driver options

- `BONDING_SLAVEn="salve device"`

  需要聚合slave，n从0开始

第三步重启网络，让配置生效

```
/etc/init.d/network restart
```
### Configuration with initscripts
red hat linux, Fedora 通常使用 initscripts，网络配置文件通常在`/etc/sysconfig/network-scripts`
第一步先修改需要聚合的 slave 配置，将SLAVE置为yes，MASTER置为需要bond的iface，BOOTPROTO置为none
```
DEVICE=eth0
USERCTL=no
ONBOOT=yes
MASTER=bond0
SLAVE=yes
BOOTPROTO=none
```
第二步，新增一个 bonding device 配置文件，以 ifcfg-bondX 命名(X从0开始，例如ifcfg-bond0)
```
DEVICE=bond0
IPADDR=192.168.1.1
NETMASK=255.255.255.0
NETWORK=192.168.1.0
BROADCAST=192.168.1.255
ONBOOT=yes
BOOTPROTO=none
USERCTL=no
```
Be sure to change the networking specific lines (IPADDR, NETMASK, NETWORK and BROADCAST) to match your network configuration.
如果是较新版本的(red hat 5之后的)，还支持`BONDING_OPTS` ，==注意有些distro需要指定`TYPE=bond0`配置才会生效==
第三步重启网络，让配置生效
```
/etc/init.d/network restart
```
### Configuration with iproute2
使用iproute2手动添加持久配置，在`/etc/init.d/boot.local`或者`/etc/rc.d/rc.local`中添加如下配置
```
modprobe bonding mode=balance-alb miimon=100
modprobe e100
ifconfig bond0 192.168.1.1 netmask 255.255.255.0 up
ip link set eth0 master bond0
ip link set eth1 master bond0
```
如果需要卸载配置注释如下内容
```
# ifconfig bond0 down
# rmmod bonding
# rmmod e100
```
## Querying Bonding configuration
方法一，可以在`/proc/net/bonding`查看bond设备的配置
```
Ethernet Channel Bonding Driver: 2.6.1 (October 29, 2004)
Bonding Mode: load balancing (round-robin)
Currently Active Slave: eth0
MII Status: up
MII Polling Interval (ms): 1000
Up Delay (ms): 0
Down Delay (ms): 0

Slave Interface: eth1
MII Status: up
Link Failure Count: 1

Slave Interface: eth0
MII Status: up
Link Failure Count: 1
```
方法二，使用`ifconfig`。所有的SLAVE和MASTER使用的MAC地址都一样
```
# /sbin/ifconfig
bond0     Link encap:Ethernet  HWaddr 00:C0:F0:1F:37:B4
          inet addr:XXX.XXX.XXX.YYY  Bcast:XXX.XXX.XXX.255  Mask:255.255.252.0
          UP BROADCAST RUNNING MASTER MULTICAST  MTU:1500  Metric:1
          RX packets:7224794 errors:0 dropped:0 overruns:0 frame:0
          TX packets:3286647 errors:1 dropped:0 overruns:1 carrier:0
          collisions:0 txqueuelen:0

eth0      Link encap:Ethernet  HWaddr 00:C0:F0:1F:37:B4
          UP BROADCAST RUNNING SLAVE MULTICAST  MTU:1500  Metric:1
          RX packets:3573025 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1643167 errors:1 dropped:0 overruns:1 carrier:0
          collisions:0 txqueuelen:100
          Interrupt:10 Base address:0x1080

eth1      Link encap:Ethernet  HWaddr 00:C0:F0:1F:37:B4
          UP BROADCAST RUNNING SLAVE MULTICAST  MTU:1500  Metric:1
          RX packets:3651769 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1643480 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:100
          Interrupt:9 Base address:0x1400
```
## Bonding Routes
当配置了bonding iface，slave会没有路由取而代之的是以master显示的路由
```
Kernel IP routing table
Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
10.0.0.0        0.0.0.0         255.255.0.0     U        40 0          0 eth0
10.0.0.0        0.0.0.0         255.255.0.0     U        40 0          0 eth1
10.0.0.0        0.0.0.0         255.255.0.0     U        40 0          0 bond0
127.0.0.0       0.0.0.0         255.0.0.0       U        40 0          0 lo
```
实际10.0.0.0的路由会从bond0发出
## Modules
### Bonding for High Availability
eth0/eth1 聚合bond0，上联两台swith，如果一台switch挂了，流量仍然能打到host1上
```
      |                                     |
      |port3                           port3|
+-----+----+                          +-----+----+
|          |port2       ISL      port2|          |
| switch A +--------------------------+ switch B |
|          |                          |          |
+-----+----+                          +-----++---+
      |port1                           port1|
      |             +-------+               |
      +-------------+ host1 +---------------+
               eth0 +-------+ eth1
```
### Bonding for Maximum Throughput
eth0/eth1聚合bond0，上联一台路由器或者3层交换机的两个口。以增大流量的吞吐
```
+----------+                     +----------+
|          |eth0            port1|          | to other networks
| Host A   +---------------------+ router   +------------------->
|          +---------------------+          | Hosts B and C are out
|          |eth1            port2|          | here somewhere
+----------+                     +----------+
```



