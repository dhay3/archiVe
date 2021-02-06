# ifcfg 配置文件

参考：

https://blog.csdn.net/u011857683/article/details/80950466

## 概述

在Linux中往往是通过命令修改文件的方式配置网络，因此不仅需要知道配置哪个文件，还要知道文件中每个配置参数的功能。在Redhat/Fedora等Linux中，网络配置文件一般是`/etc/sysconfig/network-scripts/ifcfg-eth0`；而在SLES 10中却是`/etc/sysconfig/network/ifcfg-eth-id-xx:xx:xx:xx:xx:xx`（后面是该网络接口的MAC地址）；在SLES 11中是`/etc/sysconfig/network/ifcfg-eth0`。

 

在一个计算机系统中，可以有多个网络接口，分别对应多个网络接口配置文件，在`/etc/sysconfig/network-scripts/`目录下，依次编号的文件是`ifcfg-eth0`，`ifcfg-eth1`，...，`ifcfg-eth<X>`。常用的是`ifcfg-eth0`，表示第一个网络接口配置文件。

##  配置参数

> 这些参数值不区分大小写，不区分单引号和双引号，甚至可以不用引号。

| 参数                | 值                                                           | 说明                                                         |
| ------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| TYPE                | (a) Ethernet(b) IPsec                                        | 配置文件接口类型。有Ethernet 、IPsec等类型，以太网网络接口类型为Ethernet，默认为Ethernet。 |
| DEVICE              | (a) eth0(b) eth1(c) ......                                   | 这是物理设备的名字（动态分配的PPP设备应当除外，它的名字是“逻辑名”）。 |
| BOOTPROTO           | (a) none：不使用启动地址协议。设置为none禁止DHCP。(b) bootp：BOOTP协议。(c) dhcp：DHCP动态地址协议。(d) static：静态地址协议。 | 系统启动地址协议。                                           |
| ONBOOT              | (a) yes：系统启动时激活该网络接口。(b) no：系统启动时不激活该网络接口。 | 系统启动时是否激活。                                         |
| IPADDR              | 192.168.1.123                                                | IP地址。                                                     |
| NETMASK             | 255.255.255.0                                                | 子网掩码。                                                   |
| NETWORK             | 192.168.1.0                                                  | 网络地址。                                                   |
| GATEWAY             | 192.168.1.254                                                | 网关地址。                                                   |
| BROADCAST           | 192.168.1.255                                                | 广播地址。                                                   |
| HWADDR/MACADDR      | AA:BB:CC:DD:EE:FF                                            | MAC地址。只需设置其中一个，同时设置时不能相互冲突。          |
| PREFIX              | 24                                                           | 子网掩码前缀。                                               |
| UUID                | 27137241-842e-4e50-88dd-8d8da1305dc0                         | 设备唯一标识。                                               |
| NAME                | System eth0                                                  | 网络连接的名字。                                             |
| MASTER              | bond0                                                        | 指定主的名字，绑定网口会用到。                               |
| SLAVE               | (a) yes：是主的组件。(b) no：非主的组件。                    | 指定该网口是主的组件，绑定网口会用到。。                     |
| ARPCHECK            | (a) yes：网卡启动需要检测。(b) no：网卡启动不需要检测。      | 如果服务启动时显示ip is already in use for device eth0, 这个不是ip地址冲突,因为有的linux默认开启ARPCHECK此时在配置网卡的文件中添加 ARPCHECK=no 。 |
| PEERDNS             | (a) yes：如果DNS设置，修改/etc/resolv.conf中的DNS. (b) 不修改/etc/resolv.conf中的DNS. | 是否允许DHCP获得的DNS覆盖本地的DNS。如果使用DHCP协议，默认为yes。 |
| PEERROUTES          | (a) yes：从DHCP服务器获取路由。(b) no：不从DHCP服务器获取路由。 | 是否从DHCP服务器获取用于定义接口的默认网关的信息的路由表条目。 |
| DNS{1, 2}           | (a) 8.8.8.8(b) 9.9.9.9(c) ......                             | DNS地址。当PEERDNS为yes时会被写入/etc/resolv.conf中。        |
| NM_CONTROLLED       | (a) yes: 由Network Manager控制。(b) no: 不由Network Manager控制。 | 是否由Network Manager控制该网络接口。修改保存后立即生效，无需重启。被其坑过几次，建议一般设为no。 |
| USERCTL             | (a) yes：非root用户允许控制该网络接口。(b) no：非root用户不允许控制该网络接口。 | 用户权限控制。                                               |
| IPV6INIT            | (a) yes：支持IPv6。(b) no：不支持IPv6。                      | 是否执行IPv6。                                               |
| IPV6ADDR            | FD55:faaf:e1ab:1B0D:10:14:24:106/64                          | IPv6地址/前缀长度。                                          |
| IPV4_FAILURE_FATAL  | (a) yes：如果ipv4配置失败，即使ipv6连接成功也报告失败。(b) no：不进行整体检查。 | 是否一定要进行ipv4检查。理论上，如果开启了ipv4和ipv6，那么其中一个获取ip成功，就是成功了。如果开启该选项，如果ipv4配置失败，即使ipv6连接成功也报告失败。 |
| IPV6_FAILURE_FATA   | (a) yes：如果ipv6配置失败，即使ipv4连接成功也报告失败。(b) no：不进行整体检查。 | 是否一定要进行ipv6检查。理论上，如果开启了ipv4和ipv6，那么其中一个获取ip成功，就是成功了。如果开启该选项，如果ipv6配置失败，即使ipv4连接成功也报告失败。 |
| DEFROUTE            | (a) yes：该网卡启动时，会自动生成一条默认路由。如果多个网卡设置了yes，由于只保存一条默认路由，那么每个网卡启动都会先去掉原先的默认路由，然后添加一条该网口的默认路由。(b) no：该网卡启动时，不会自动生成默认路由。 | centos只允许自动生成一条默认路由(当然你可以通过命令route add 手动添加多条默认路由，但当网口重启时，又会去掉其它路由，只保留一条该网卡的默认路由)。 如果有多个网卡，比如eth0、eth1，就有可能会出现这样的情况：系统默认路由选择eth0网卡，而实际使用的是eth1网卡，这个时候就需要配置DEFROUTE参数，在eth1的配置文件内加上DEFROUTE=yes |
| PERSISTENT_DHCLIENT | (a) yes：dhcp续约失败进程dhclient不退出，休息一段时间继续请求dhcp服务器。(b) no：dhcp续约失败进程dhclient退出。 | 控制dhclient进程续约失败是否退出。可用于解决dhclient 进程检测ip冲突后，发送 Decline报文退出的问题。 |

## 配置生效

修改网络配置，最简单地是使用ifconfig命令，无需重启，立即生效。ifconfig配置的网络只是当前临时有效，当计算机重启之后就失效了。为了使网络配置永久有效，就需要在`/etc/sysconfig/network-scripts/`下修改网络接口配置文件。但是，这些文件修改后，并不能立即生效。有三种方式使修改文件的网络配置生效：

 

(1) 重启所有网口，执行service network restart命令，重启网络服务

(2) 重启某个网口，执行ifdown eth0 ，然后 ifup eth0

(3) 重启计算机，执行reboot

## 配置文件示例

### 静态地址

```shell

[root@omp120 ~]# cat /etc/sysconfig/network-scripts/ifcfg-eth0 
DEVICE="eth0"
BOOTPROTO="static"
NM_CONTROLLED="no"
ONBOOT="yes"   // 开机启动网卡
TYPE="Ethernet"
UUID=27137241-842e-4e50-88dd-8d8da1305dc0
DEFROUTE=yes   // 启动该网卡时，自动生成默认路由
IPADDR=192.168.254.109
NETMASK=255.255.255.0
GATEWAY=
DNS1=
DNS2=
HWADDR=00:90:27:50:5B:30
ARPCHECK=no
```

### 动态地址

```shell
[root@omp120 ~]# cat /etc/sysconfig/network-scripts/ifcfg-eth1
DEVICE=eth1
TYPE="Ethernet"
ONBOOT=yes
NM_CONTROLLED=no
BOOTPROTO=dhcp
HWADDR=00:90:27:50:5B:31
UUID=27137241-842e-4e50-88dd-8d8da1305dc1
DEFROUTE=yes
```

### pppoe拨号

```shell
[root@wensen ~]# cat /etc/sysconfig/network-scripts/ifcfg-ppp2 
USERCTL=no
BOOTPROTO=dialup
NAME=DSLppp2
DEVICE=ppp2
TYPE=xDSL
ONBOOT=yes
PIDFILE=/var/run/ppp2_adsl.pid
FIREWALL=NONE
PING=.
PPPOE_TIMEOUT=80
LCP_FAILURE=3
LCP_INTERVAL=20
CLAMPMSS=1412
CONNECT_POLL=6
CONNECT_TIMEOUT=60
DEFAULTROUTE=no       // 默认路由
SYNCHRONOUS=no
ETH=eth2   // 物理口eth2，使用pppoe拨号形式
PROVIDER=DSLppp2
USER=12345678
PEERDNS=yes
DEMAND=no
```

## 绑定口

绑定eth1和eth2

```shell
[root@wensen ~]# cat /etc/sysconfig/network-scripts/ifcfg-eth1
DEVICE=eth1
NM_CONTROLLED=no
BOOTPROTO=none
MASTER=bond0   // 绑定口名称
SLAVE=yes  // 该网卡为绑定口的一部分
USERCTL=no
HWADDR=00:90:27:50:5B:31
UUID=27137241-842e-4e50-88dd-8d8da1305dc1
```

```shell
[root@wensen ~]# cat /etc/sysconfig/network-scripts/ifcfg-eth2
DEVICE=eth2
NM_CONTROLLED=no
BOOTPROTO=none
MASTER=bond0   // 绑定口名称
SLAVE=yes    // 该网卡为绑定口的一部分
USERCTL=no
HWADDR=00:90:27:50:5B:32
UUID=27137241-842e-4e50-88dd-8d8da1305dc2
```

```shell
[root@wensen ~]# cat /etc/sysconfig/network-scripts/ifcfg-bond0 
DEVICE=bond0
ONBOOT=yes
NM_CONTROLLED=no
BOOTPROTO=static
IPADDR=192.168.88.23
NETMASK=255.255.255.0
GATEWAY=
DNS1=
DNS2=
```

