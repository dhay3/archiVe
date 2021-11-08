# Linux IP概述通用参数 

reference：

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/configuring_and_managing_networking/linux-traffic-control_configuring-and-managing-networking

http://linux-ip.net/gl/ip-cref/ip-cref-node17.html

## 概述

syntax：`ip [options] object {command}`

用来配置和查询linux的网络

object := {link | address | addrlabel | route | rule | neigh}

具体查看man page

## general options

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

  输出详细的信息

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