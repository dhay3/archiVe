# Linux tcpdump

参考：

https://juejin.im/post/6844904084168769549

```
 tcpdump -i eth0 -nn -vv port 80 
```

## 概述

命令行抓包工具，不仅限于tcp协议。

## 参数

- -c

  抓取指定个数个包后tcpdump自动退出

  

- -D

  显示所有网络接口。tcpdump默认监听第一个网络接口

  ```
  [root@chz Desktop]# tcpdump --list-interfaces
  1.bluetooth0 (Bluetooth adapter number 0)
  2.nflog (Linux netfilter log (NFLOG) interface)
  3.nfqueue (Linux netfilter queue (NFQUEUE) interface)
  4.usbmon1 (USB bus number 1)
  5.usbmon2 (USB bus number 2)
  6.ens33
  7.any (Pseudo-device that captures on all interfaces)
  8.lo [Loopback]
  [root@chz Desktop]# tcpdump 
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on bluetooth0, link-type BLUETOOTH_HCI_H4_WITH_PHDR (Bluetooth HCI UART transport layer plus pseudo-header), capture size 262144 bytes
  ^C
  0 packets captured
  48 packets received by filter
  0 packets dropped by kernel
  ```

- -i

  指定监听的接口

- -nn

  单个n取消DNS解析，两个n取消DNS解析和端口解析

- -t

  不打印时间错

- -r

  读取用-w生成的文件

  ```
  [root@chz Desktop]# tcpdump -r test 
  reading from file test, link-type EN10MB (Ethernet)
  10:55:13.395594 IP 192.168.80.200 > chz: ICMP echo request, id 3248, seq 675, length 64
  10:55:13.395650 IP chz > 192.168.80.200: ICMP echo reply, id 3248, seq 675, length 64
  10:55:14.420008 IP 192.168.80.200 > chz: ICMP echo request, id 3248, seq 676, length 64
  10:55:14.420065 IP chz > 192.168.80.200: ICMP echo reply, id 3248, seq 676, length 64
  10:55:15.443622 IP 192.168.80.200 > chz: ICMP echo request, id 3248, seq 677, length 64
  ```

- -w

  将数据包以==raw==的格式(binary file)写入到文件，所有内容将不会以标准输出流输出

  ```
  [root@chz Desktop]# tcpdump -w test -i ens33
  ```

- ==-W==

  限制生成的文件个数，和-C，-G一起使用。最好将文件按照==strftime==格式命名

  ```
  cpl in /tmp λ sudo tcpdump -i wlp1s0 -n -G 30 -w %M%S.pcap -W 3
  cpl in /tmp λ ls
   0019.pcap
   0049.pcap
   0119.pcap
  ```
  
- -G 

  在指定间隔(sec)后向文件中写入也可以发送SIGINT信号中断进程提前写入，一般与-w或-W一起使用

- -C 

  指定生成文件的大小

  ```
  cpl in /tmp λ sudo tcpdump -i wlp1s0 -n -G 30 -w %M%S.pcap -C 10
  ```
  
- -e

  显示link-layer MAC地址

  ```
  [root@chz Desktop]# tcpdump -e -i ens33 
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
  11:04:30.371765 00:50:56:32:1b:8c (oui Unknown) > 00:0c:29:d7:d1:68 (oui Unknown), ethertype IPv4 (0x0800), length 98: 192.168.80.200 > chz: ICMP echo request, id 3352, seq 1, length 64
  11:04:30.371881 00:0c:29:d7:d1:68 (oui Unknown) > 00:50:56:32:1b:8c (oui Unknown), ethertype IPv4 (0x0800), length 98: chz > 192.168.80.200: ICMP echo reply, id 3352, seq 1, length 64
  11:04:30.372686 00:0c:29:d7:d1:68 (oui Unknown) > 00:50:56:e3:f8:b0 (oui Unknown), ethertype IPv4 (0x0800), length 87: chz.40942 > gateway.domain: 12072+ PTR? 139.80.168.192.in-addr.arpa. (45)
  ```

- -c

  抓取指定数量的数据包，可以使用SIGINT提前结束

  ```
  [root@chz Desktop]# tcpdump -i ens33 -c 3
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
  19:51:46.073738 ARP, Request who-has 192.168.80.200 tell chz, length 28
  19:51:46.074334 IP chz.53701 > gateway.domain: 62718+ PTR? 200.80.168.192.in-addr.arpa. (45)
  19:51:46.086321 IP gateway.domain > chz.53701: 62718 NXDomain 0/1/0 (80)
  3 packets captured
  7 packets received by filter
  0 packets dropped by kernel
  ```



## 过滤器

> 如果没有指定过滤器，所有的数据包会被捕捉。具体查看 pcap-filter(7)。与wireshark的过滤条件相同

### proto

Proto 过滤器用来过滤某个协议的数据，关键字为 `proto`，可省略。proto 后面可以跟上协议号或协议名称，支持 `icmp`, `igmp`, `igrp`, `pim`, `ah`, `esp`, `carp`, `vrrp`, `udp`和 `tcp`。因为通常的协议名称是保留字段，所以在于 proto 指令一起使用时，必须根据 shell 类型使用一个或两个反斜杠（/）来转义。Linux 中的 shell 需要使用两个反斜杠来转义，MacOS 只需要一个

```
[root@chz Desktop]# tcpdump -i ens33 proto \\tcp
或
[root@chz Desktop]# tcpdump -i ens33 tcp
```

### host

Host 过滤器用来过滤某个主机的数据报文。例如：

```
[root@chz Desktop]# tcpdump -i ens33 host 192.168.80.200
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:15:53.352025 IP 192.168.80.200 > chz: ICMP echo request, id 3486, seq 1, length 64
11:15:53.352103 IP chz > 192.168.80.200: ICMP echo reply, id 3486, seq 1, length 64
11:15:54.357307 IP 192.168.80.200 > chz: ICMP echo request, id 3486, seq 2, length 64
11:15:54.357366 IP chz > 192.168.80.200: ICMP echo reply, id 3486, seq 2, length 64
11:15:58.359014 ARP, Request who-has 192.168.80.200 tell chz, length 28
11:15:58.359522 ARP, Reply 192.168.80.200 is-at 00:50:56:32:1b:8c (oui Unknown), length 46
11:15:58.518641 ARP, Request who-has chz tell 192.168.80.200, length 46
11:15:58.518657 ARP, Reply chz is-at 00:0c:29:d7:d1:68 (oui Unknown), length 28
```

该命令会抓取所有发往主机 `192.168.80.200` 或者从主机 `192.168.80.200` 发出的流量。如果想==只抓取从该主机发出的流量(src host)==，可以使用下面的命令：

```
[root@chz Desktop]# tcpdump -i ens33 src host 192.168.80.200
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:17:12.199686 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 1, length 64
11:17:13.210913 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 2, length 64
11:17:14.211471 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 3, length 64
11:17:15.211456 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 4, length 64
11:17:16.212138 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 5, length 64
```

==只抓取目标是该主机的流量(dst host)==

```
[root@chz Desktop]# tcpdump -i ens33 dst host 192.168.80.200
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:17:51.266870 IP chz > 192.168.80.200: ICMP echo request, id 8174, seq 40, length 64
11:17:52.269125 IP chz > 192.168.80.200: ICMP echo request, id 8174, seq 41, length 64
11:17:53.271534 IP chz > 192.168.80.200: ICMP echo request, id 8174, seq 42, length 64
```

### net

过滤指定网段的数据包，使用CIDR格式

```
[root@chz Desktop]# tcpdump -i ens33 net 192.168.80.0/24
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:28:05.550637 IP chz > 192.168.80.200: ICMP echo request, id 8174, seq 653, length 64
11:28:05.550987 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 653, length 64
11:28:05.552051 IP chz.39578 > gateway.domain: 6019+ PTR? 200.80.168.192.in-addr.arpa. (45)
```

==如果只想抓取该网段发出的数据包(src net)==

```
[root@chz Desktop]# tcpdump -i ens33 src net 192.168.80.0/24
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:29:26.704333 IP chz > 192.168.80.200: ICMP echo request, id 8174, seq 734, length 64
11:29:26.704664 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 734, length 64
11:29:26.705091 IP chz.54563 > gateway.domain: 52031+ PTR? 200.80.168.192.in-addr.arpa. (45)
```

==如果只想抓取目标是该网段的数据包(dst net)==

```
[root@chz Desktop]# tcpdump -i ens33 dst net 192.168.80.0/24
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:30:31.828019 IP chz > 192.168.80.200: ICMP echo request, id 8174, seq 799, length 64
11:30:31.828837 IP 192.168.80.200 > chz: ICMP echo reply, id 8174, seq 799, length 64
11:30:31.829975 IP chz.54657 > gateway.domain: 30303+ PTR? 200.80.168.192.in-addr.arpa. (45)
```

### port

Port 过滤器用来过滤通过某个端口的数据报文，关键字为 `port`。==同样也有src port 和 dst port==例如：

```
[root@chz Desktop]# tcpdump -i ens33 port 80
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
11:13:57.551557 IP 192.168.80.200.43430 > chz.http: Flags [S], seq 2715582199, win 64240, options [mss 1460,sackOK,TS val 1613393823 ecr 0,nop,wscale 7], length 0
11:13:57.551664 IP chz.http > 192.168.80.200.43430: Flags [R.], seq 0, ack 2715582200, win 0, length 0
```

### portrange

过滤指定范围内的端口号的数据包。==同样也有src portrange 和 dst portrange==

```
[root@chz Desktop]# tcpdump -i ens33 portrange 80-3306
```

## 例子

> 可以使用逻辑运算符('!' or 'not'；‘&&’ or  ‘and’；'||' or 'or')，也可以使用主机名，也可以使用子表达式

1. ```
   [root@chz Desktop]# tcpdump -i ens33 host 192.168.80.200 && 192.168.80.100
   tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
   listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
   19:59:51.892328 IP 192.168.80.200 > chz: ICMP echo request, id 2177, seq 1, length 64
   19:59:51.892465 IP chz > 192.168.80.200: ICMP echo reply, id 2177, seq 1, length 64
   ```

   同时捕捉192.168.80.200和192.168.80.100数据包

2. ```
   [root@chz opt]# tcpdump -i ens33 not arp
   tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
   listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
   ```

   捕捉非arp协议的

