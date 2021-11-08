## objects

### address

输出L4相关的信息

### addrlabel

### link

### monitor

### netns

### ntable

### route

### rule

### vrf

管理vitrual routing 和 forwarding devices

## ip

> 使用 `ip link`，`ip addr`，`ip route`替代`ifconfig`，`route` 。
>
> 当启用dhcp，我们可以使用`dhclient -r`来释放IP(等价于windows的`ipconfig /release`)，`dhclient`重新获取IP。

pattern：`ip [options] OBJECT {COMMAND|help}`

每一个对象都有help命令，可以查看具体的`COMMAND`。使用`-s`和`-d`参数显示更加详细的信息

## keyword

- mtu

  maximal transform unit，单包最大的传输值

  byte

- qdisc

  queueing discipline

  接受和发送数据的缓冲队列使用的算法，具体可以使用的值参考[链接](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/configuring_and_managing_networking/linux-traffic-control_configuring-and-managing-networking)，常见的有noqueue表示不使用队列缓存信息，noop使用blackhole，所有发送到该iface的数据包都会被丢弃

  可以用`tc qdisc show dev <iface>`查看

- state

  表明iface的状态是否启用

- master

  该iface是另一个iface的子接口

  ```
  4: virbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
      link/ether 52:54:00:12:da:f9 brd ff:ff:ff:ff:ff:ff
      inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
         valid_lft forever preferred_lft forever
  5: virbr0-nic: <BROADCAST,MULTICAST> mtu 1500 qdisc fq_codel master virbr0 state DOWN group default qlen 1000
      link/ether 52:54:00:12:da:f9 brd ff:ff:ff:ff:ff:ff
  ```

- group

  ip 组

- qlen

  缓存队列最大的值

## ip link

==设置和查看NIC的接口(数据链路层)==，可以使用`ip l`来缩写`ip link show`

- ip link show

  显示NIC(network interface controller)的运行状态，包括DOWN和UP。show可以被省略。可以指定具体的NIC，使用`ip link show ens33`

  ```
  C:\root> ip l show
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
  2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
      link/ether 00:50:56:32:1b:8c brd ff:ff:ff:ff:ff:ff
  3: eth1: <BROADCAST,MULTICAST> mtu 1500 qdisc pfifo_fast state DOWN mode DEFAULT group default qlen 1000
      link/ether 00:0c:29:a0:ef:a3 brd ff:ff:ff:ff:ff:ff
  4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN mode DEFAULT group default 
      link/ether 02:42:58:78:34:6a brd ff:ff:ff:ff:ff:ff
  ```

  

- ==ip link set==

  ```
  ip l set ens33 up | down #开启或关闭指定接口
  ip l set ens33 address <MAC ADDRESS> #设置接口的mac地址
  ```

- ip link add

  具体查看使用`man ip link`, 这里使用`type==vlan`。为NIC添加一个子接口

  pattern：`ip link add link <device> name <identifier> type <type> id <支持十六进制>`

  ```
  [root@chz network-scripts]# ip l add link ens33 name ens33.1 type vlan id 1
  [root@chz network-scripts]# ip a
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host 
         valid_lft forever preferred_lft forever
  2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
      link/ether 00:0c:29:d7:d1:68 brd ff:ff:ff:ff:ff:ff
      inet 192.168.80.131/24 brd 192.168.80.255 scope global noprefixroute dynamic ens33
         valid_lft 1703sec preferred_lft 1703sec
      inet6 fe80::a164:2ef4:8841:5fc7/64 scope link noprefixroute 
         valid_lft forever preferred_lft forever
  3: ens33.1@ens33: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
      link/ether 00:0c:29:d7:d1:68 brd ff:ff:ff:ff:ff:ff
  
  ```

  在ens33 NIC上添加一个ens33.1接口

- ip link delete

  可以使用`del`缩写

  ```
  [root@chz network-scripts]# ip link del ens33.1
  ```

## ip route

> 替换原来的route 命令，==重启主机后失效==

查看路由表，使用`ip r`缩写`ip route show`

- ip r

  对比`route`命令

  ```
  [root@chz ~]# ip r
  default via 192.168.80.2 dev ens33 proto dhcp metric 100 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.131 metric 100 
  192.168.122.0/24 dev virbr0 proto kernel scope link src 192.168.122.1 
  [root@chz ~]# route -n
  Kernel IP routing table
  Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
  0.0.0.0         192.168.80.2    0.0.0.0         UG    100    0        0 ens33
  192.168.80.0    0.0.0.0         255.255.255.0   U     100    0        0 ens33
  192.168.122.0   0.0.0.0         255.255.255.0   U     0      0        0 virbr0
  ```

- ip r add

  可以指定具体主机，也可以使用CIDR

  ```
  [root@chz network-scripts]# ip r add 192.168.10.0/24 dev ens34 #仅主机模式，无需指定网关
  [root@chz network-scripts]# ip r
  default via 192.168.80.2 dev ens33 proto static metric 100 
  192.168.10.0/24 dev ens34 scope link 
  192.168.10.0/24 dev ens34 proto kernel scope link src 192.168.10.100 metric 101 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.100 metric 100 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.101 metric 100 
  [root@chz network-scripts]# 
  
  
  
  [root@chz network-scripts]# ip r add 192.168.10.200 dev ens34 #指定LAN中主机不需要掩码
  [root@chz network-scripts]# ip r
  default via 192.168.80.2 dev ens33 proto static metric 100 
  192.168.10.0/24 dev ens34 proto kernel scope link src 192.168.10.100 metric 101 
  192.168.10.200 dev ens34 scope link 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.100 metric 100 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.101 metric 100 
  
  
  
  [root@chz network-scripts]# ip r add 39.156.69.79 via 192.168.80.2 dev ens33 #指定外网IP需要指定网关
  [root@chz network-scripts]# ip r
  default via 192.168.80.2 dev ens33 proto static metric 100 
  39.156.69.79 via 192.168.80.2 dev ens33 
  192.168.10.0/24 dev ens33 scope link 
  192.168.10.0/24 dev ens34 proto kernel scope link src 192.168.10.100 metric 101 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.100 metric 100 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.101 metric 100 
  [root@chz network-scripts]# 
  ```

- ip r del

  删除默认路由，使用`ping 39.156.69.79`来检测路由是否生效

  ```
  [root@chz network-scripts]# ip r del default dev ens33
  [root@chz network-scripts]# ip r
  39.156.69.79 via 192.168.80.2 dev ens33 
  192.168.10.0/24 dev ens33 scope link 
  192.168.10.0/24 dev ens34 proto kernel scope link src 192.168.10.100 metric 101 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.100 metric 100 
  192.168.80.0/24 dev ens33 proto kernel scope link src 192.168.80.101 metric 100 
  ```

- ip r change

  使用删除后添加路由替代修改功能

## ip neighbour

查询neighbour table（==ARP table==），使用`ip n`缩写`ip neighbour`。我们也可以通过`apr`命令来查看

> arp 表只会记录当前局域网内IP对应的MAC地址，访问外网交给gateway(记录网关IP和MAC)

- ip n

  ```
  [root@chz ~]# ip n
  192.168.80.254 dev ens33 lladdr 00:50:56:e0:99:7d STALE
  192.168.80.2 dev ens33 lladdr 00:50:56:e3:f8:b0 REACHABLE
  192.168.80.200 dev ens33 lladdr 00:50:56:32:1b:8c STALE
  ```

- ip n flush 

  删除指定NIC的ARP table entry，与`arp -d`的区别在于前者指定NIC，后者指定IP或CIDR

  ```
  [root@chz network-scripts]# arp 
  Address                  HWtype  HWaddress           Flags Mask            Iface
  gateway                  ether   00:50:56:e3:f8:b0   C                     ens33
  [root@chz network-scripts]# ip n flush dev ens33
  [root@chz network-scripts]# ip n
  [root@chz network-scripts]# arp 
  [root@chz network-scripts]# 
  ```
