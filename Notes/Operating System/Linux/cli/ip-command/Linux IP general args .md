# Linux IP general args 

reference：

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/configuring_and_managing_networking/linux-traffic-control_configuring-and-managing-networking

http://linux-ip.net/gl/ip-cref/ip-cref-node17.html

## Digest

syntax：`ip [options] object {command}`

用来配置和查询linux的网络

object := {link | address | addrlabel | route | rule | neigh}

具体查看man page

## Optional args

### batch mode

- `-b | -batch FILENAME`

  从文件中批量执行ip命令

  ```
  [vagrant@10 tmp]$ cat file
  addr show
  route show
  [vagrant@10 tmp]$ ip -b file
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host 
         valid_lft forever preferred_lft forever
  2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
      link/ether 52:54:00:4d:77:d3 brd ff:ff:ff:ff:ff:ff
      inet 10.0.2.15/24 brd 10.0.2.255 scope global noprefixroute dynamic eth0
         valid_lft 85768sec preferred_lft 85768sec
      inet6 fe80::5054:ff:fe4d:77d3/64 scope link 
         valid_lft forever preferred_lft forever
  default via 10.0.2.2 dev eth0 proto dhcp metric 100 
  10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 100
  ```

  无须在每一行命令开头指定`ip`，如果其中一个命令错误就会终止ip命令

- `-force`

  don’t terminate ip on errors in batch mode

### general options

- `-s | -stats | -statistics`

  输出详细的统计信息(收包，丢包)，s越多信息越详细

  ```
  cpl in ~ λ ip -s -s a
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host 
         valid_lft forever preferred_lft forever
      RX: bytes  packets  errors  dropped missed  mcast   
      3719595667 5815589  0       0       0       0       
      RX errors: length   crc     frame   fifo    overrun
                 0        0       0       0       0       
      TX: bytes  packets  errors  dropped carrier collsns 
      3719595667 5815589  0       0       0       0       
      TX errors: aborted  fifo   window heartbeat transns
                 0        0       0       0       0     
  ```

- `-d | --details`

  输出详细的信息，这里可以看interface的类型是eth还是vlan或者是bridge etc

  ```
  cpl in ~ λ ip -d a
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00 promiscuity 0 minmtu 0 maxmtu 0 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host 
         valid_lft forever preferred_lft forever
  ```

- `-4`

  输出ipv4的相关信息

- `-6`

  输出ipv6的相关信息

- `-B`

  输出bridge的相关信息

- `-0`

  输出链路层信息

- `-o | -oneline`

  对每个iface以一行的格式输出，方便wc或grep统计

  ```
  cpl in ~ λ ip -o a | wc -l
  7
  cpl in ~ λ ip a | wc -l
  20
  ```

- `-c=always`

  高亮显示IP，link信息

- `-br | -brief`

  简要输出网络信息

  ```
  cpl in ~ λ ip -br a
  lo               UNKNOWN        127.0.0.1/8 ::1/128 
  wlp1s0           UP             192.168.0.103/24 240e:390:e52:7cf1::1001/128 240e:390:e52:7cf1:a44d:89dc:dc5:6395/64 fe80::5e7a:b687:b448:fd5f/64 
  docker0          DOWN           172.17.0.1/16 
  ```

- `-j | -json`

  以json的格式输出，可以结合json_pp来格式化

  ```
  ip -j a | json_pp
  ```

- `-t | --tmiestamp`

  在monitor模式下显示时间戳

## Object

具体查看具体文档

### address

IPV4/IPV6 相关的信息

### addrlabel

label configuration for protocol address selection

### l2tp

tunnel ethernet over ip(L2TPv3)

### maddress

组播(multicase)地址管理

### monitor

监控设备状态，地址路由变化

### mroute

组播路由cache

### mrule

组播 routing policy

### neigbhour

ARP  和 NDISC cache

### netns

network namespace 管理

### netable

ARP 和 NDISC cache 管理

### route

路由条目

### rule

routing policy

### tcp_metrics/tcpmetrics

管理tcp metrics

### token

manage tokenized interface identifiers

### tunnel

tunnel over IP

### tuntap

mange TUN/TAP devices

### xfrm

mange IPsec polices

## Command

表示 object 的动作，例如show、add、delete等，如果没有指定command，一些 default command 会被使用，通常是 list