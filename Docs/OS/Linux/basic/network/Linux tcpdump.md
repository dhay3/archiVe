# Linux tcpdump

参考：

https://juejin.im/post/6844904084168769549

>  tcpdump -Snnvvi eth0 tcp and port 80 

## 概述

syntax：`tcpdump [options] [filters]`

命令行抓包工具，不仅限于tcp协议。发送SINGTERM终止

```
root in /home/ubuntu λ tcpdump -i any -c 3
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
08:15:34.080299 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 2112669508:2112669600, ack 824759699, win 501, length 92
08:15:34.080575 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 92:200, ack 1, win 501, length 108
08:15:34.080596 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 200:236, ack 1, win 501, length 36
3 packets captured
26 packets received by filter
17 packets dropped by kernel
```

received by filter 会跟进OS不同执行不同的操作，查看manual page 获取详细信息

## 参数

- -D | --list-interfaces

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

  指定监听的接口，如果不指定默认使用第一个iface，==可以使用any抓所有的iface==

  ```
  root in /home/ubuntu λ tcpdump -i any
  ```

- -c

  抓取指定个数个包后tcpdump自动退出，例如抓100个包

  ```
  root in /home/ubuntu λ tcpdump -i any -c 10
  ```
  
- -S

  以绝对的方式打印sequence number，默认第一个包先显示seq然后以相对的方式显示seq

  ```
  cpl in ~ λ sudo tcpdump -Snnvvi wlp1s0 tcp and host 1.1.1.1
  tcpdump: listening on wlp1s0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
  13:16:04.968479 IP (tos 0x0, ttl 64, id 58267, offset 0, flags [DF], proto TCP (6), length 60)
      30.226.76.42.52248 > 1.1.1.1.443: Flags [S], cksum 0x6d3c (incorrect -> 0xc0a5), seq 354593100, win 64240, options [mss 1460,sackOK,TS val 1233537427 ecr 0,nop,wscale 7], length 0
  13:16:05.204493 IP (tos 0x14, ttl 48, id 0, offset 0, flags [DF], proto TCP (6), length 52)
      1.1.1.1.443 > 30.226.76.42.52248: Flags [S.], cksum 0x108b (correct), seq 1646134020, ack 354593101, win 65535, options [mss 1460,nop,nop,sackOK,nop,wscale 10], length 0
  13:16:05.204582 IP (tos 0x0, ttl 64, id 58268, offset 0, flags [DF], proto TCP (6), length 40)
      30.226.76.42.52248 > 1.1.1.1.443: Flags [.], cksum 0x6d28 (incorrect -> 0x4f6a), seq 354593101, ack 1646134021, win 502, length 0
  ```

- -nn

  单个n取消DNS解析，两个n取消DNS解析和端口解析

- -v

  打印出ttl,tos,id,length,flags,proto等信息

- -t

  不打印时间戳

- -ttt

  以delta的格式输出时间戳

  ```
  root in /home/ubuntu λ tcpdump -i any -c 10 -ttt
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
   00:00:00.000000 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 2112824520:2112824628, ack 824803895, win 501, length 108
   00:00:00.000025 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 108:144, ack 1, win 501, length 36
   00:00:00.000022 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 144:260, ack 1, win 501, length 116
   00:00:00.000016 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 260:296, ack 1, win 501, length 36
   00:00:00.000229 IP localhost.localdomain.58762 > localhost.domain: 16067+ [1au] PTR? 99.72.120.42.in-addr.arpa. (54)
   00:00:00.134158 IP localhost.localdomain.49429 > localhost.domain: 50225+ [1au] PTR? 53.0.0.127.in-addr.arpa. (52)
   00:00:00.000055 IP gns3vm.65522 > 42.120.72.99.2539: Flags [P.], seq 296:780, ack 1, win 501, length 484
   00:00:00.000092 IP localhost.domain > localhost.localdomain.49429: 50225 1/0/1 PTR localhost. (75)
  ```

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

- -l

  buffer stdout中的信息，可以结合pipeline将信息以==明文==的方式输出到文件中

  ```
  cpl in ~ λ sudo tcpdump -li wlp1s0 tcp and host 1.1.1.1 | tee data
  ```

- -w

  将数据包以==raw==的格式(binary file)写入到文件，所有内容将不会以标准输出流输出。如果和`-G`一起使用需要以`strftime`的格式命名

  ```
  [root@chz Desktop]# tcpdump -w test -i ens33
  ```
  
- -C  file_size

  当抓到的单个文件大于file_size(1 == 1 millions bytes)时，会新生成一个文件，后生成的文件以`test[1++]`的格式命名

  ```
  cpl in /tmp λ sudo tcpdump -i wlp1s0 -n -G 30 -w test.pcap -C 10
  ```
  
- -G 

  在指定间隔(sec)后向文件中写入也可以发送SIGINT信号中断进程提前写入，一般与-w或-W一起使用。`-w`必须以==`strftime`==格式命名

  ```
  sudo tcpdump -i wlp1s0 -n -G 30 -w %M%S.pcap
  ```
  
- ==-W==

  限制生成的文件个数，和-C，-G一起使用。

  ```
  cpl in /tmp λ sudo tcpdump -i wlp1s0 -n -G 30 -w %M%S.pcap -W 3
  cpl in /tmp λ ls
   0019.pcap
   0049.pcap
   0119.pcap
  ```

- -z command

  在每次rotation后，都执行命令。需要和`-G`或`-C`一起使用否则无效

  ```
  
  ```
  
- -e

  ==显示link-layer MAC地址==

  ```
  [root@chz Desktop]# tcpdump -e -i ens33 
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on ens33, link-type EN10MB (Ethernet), capture size 262144 bytes
  11:04:30.371765 00:50:56:32:1b:8c (oui Unknown) > 00:0c:29:d7:d1:68 (oui Unknown), ethertype IPv4 (0x0800), length 98: 192.168.80.200 > chz: ICMP echo request, id 3352, seq 1, length 64
  11:04:30.371881 00:0c:29:d7:d1:68 (oui Unknown) > 00:50:56:32:1b:8c (oui Unknown), ethertype IPv4 (0x0800), length 98: chz > 192.168.80.200: ICMP echo reply, id 3352, seq 1, length 64
  11:04:30.372686 00:0c:29:d7:d1:68 (oui Unknown) > 00:50:56:e3:f8:b0 (oui Unknown), ethertype IPv4 (0x0800), length 87: chz.40942 > gateway.domain: 12072+ PTR? 139.80.168.192.in-addr.arpa. (45)
  ```
  
- -F file

  使用文件作为filter

  ```
  root in /home/ubuntu λ cat > filter <<EOF
  heredoc> tcp[13] == 2
  heredoc> EOF
  root in /home/ubuntu λ cat filter
  tcp[13] == 2
  root in /home/ubuntu λtcpdump -i any -c 3 -F filter
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
  08:38:10.664230 IP gns3vm.48340 > 169.254.0.4.http: Flags [S], seq 2275304406, win 64240, options [mss 1460,sackOK,TS val 2008682951 ecr 0,nop,wscale 7], length 0
  08:38:11.660677 IP gns3vm.48342 > 169.254.0.4.http: Flags [S], seq 2028759254, win 64240, options [mss 1460,sackOK,TS val 2008683948 ecr 0,nop,wscale 7], length 0
  08:38:12.661094 IP gns3vm.48344 > 169.254.0.4.http: Flags [S], seq 459183679, win 64240, options [mss 1460,sackOK,TS val 2008684948 ecr 0,nop,wscale 7], length 0
  3 packets captured
  3 packets received by filter
  0 packets dropped by kernel
  ```

- -K | --dont-verify-checksums

  不计算crc冗余码

- -l

  将输出到stdout的内容输出到line buffered中

  ```
  tcpdump -l | tee dat
  ```

- `-s <snaplen>`

  抓固定长度的包，默认抓262144bytes

  ```
  cpl in ~ λ sudo tcpdump -nni any -s 60 -c 10
  tcpdump: data link type LINUX_SLL2
  tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
  listening on any, link-type LINUX_SLL2 (Linux cooked v2), snapshot length 60 bytes
  16:20:45.271363 lo    In  IP 127.0.0.1.45414 > 127.0.0.1.1089: Flags [.], ack 4088476552, win 512, options [ [|tcp]
  16:20:45.271397 lo    In  IP 127.0.0.1.1089 > 127.0.0.1.45414: Flags [.], ack 1, win 512, options [ [|tcp]
  16:20:45.670890 wlp1s0 B   ARP, Request who-has 192.168.10.1 (ff:ff:ff:ff:ff:ff) tell 192.168.10.1, length 46
  16:20:45.687370 lo    In  IP 127.0.0.1.37506 > 127.0.0.1.15490: Flags [P.], seq 3269200592:3269200707, ack 896792731, win 621, options [ [|tcp]
  ```

## TCP flags

```
IP rtsg.1023 > csam.login: Flags [S], seq 768512:768512, win 4096, opts [mss 1024]
              IP csam.login > rtsg.1023: Flags [S.], seq, 947648:947648, ack 768513, win 4096, opts [mss 1024]
              IP rtsg.1023 > csam.login: Flags [.], ack 1, win 4096
              IP rtsg.1023 > csam.login: Flags [P.], seq 1:2, ack 1, win 4096, length 1
              IP csam.login > rtsg.1023: Flags [.], ack 2, win 4096
              IP rtsg.1023 > csam.login: Flags [P.], seq 2:21, ack 1, win 4096, length 19
              IP csam.login > rtsg.1023: Flags [P.], seq 1:2, ack 21, win 4077, length 1
              IP csam.login > rtsg.1023: Flags [P.], seq 2:3, ack 21, win 4077, urg 1, length 1
              IP csam.login > rtsg.1023: Flags [P.], seq 3:4, ack 21, win 4077, urg 1, length 1
```

- S：代表SYN
- . : 代表ACK
- F：代表FIN
- P：代表PUSH
- R：代表RST
- U：代表URG
- W：代表ECN CWR
- E：代表ECN-Echo

例如第一行表示从rtsg 1023端口发往csam login端口，tcp接受窗口为4096byte，seq为768512，没有发送数据

注意，如下并不是 tcp flags，DF 表示 3 层 IP 的 don’t fragement 标志 

```
21:54:19.343958 wlp1s0 Out IP (tos 0xc0, ttl 1, id 0, offset 0, flags [DF], proto IGMP (2), length 32, options (RA))
    192.168.2.194 > 224.0.0.251: igmp v2 report 224.0.0.251
```

## Filter



> 如果没有指定过滤器，所有的数据包(==所有的协议包==)会被捕捉。具体查看 pcap-filter(7)。与wireshark的过滤条件相同

### proto

Proto 过滤器用来过滤某个协议的数据，关键字为 `proto`，可省略。proto 后面可以跟上协议号或协议名称，支持 `icmp`, `igmp`, `igrp`, `pim`, `ah`, `esp`, `carp`, `vrrp`, `udp`和 `tcp`。因为通常的协议名称是保留字段，所以在于 proto 指令一起使用时，必须根据 shell 类型使用一个或两个反斜杠（\）来转义。Linux 中的 shell 需要使用两个反斜杠来转义，MacOS 只需要一个

```
[root@chz Desktop]# tcpdump -i ens33 proto \\tcp
或
[root@chz Desktop]# tcpdump -i ens33 tcp
#如果后面还有参数需要使用and连接
cpl in ~ λ sudo tcpdump -nni wlp1s0 tcp and host 1.1.1.1
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

### flags

![Snipaste_2020-08-25_00-39-07](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220719/Snipaste_2020-08-25_00-39-07.4kvfcqtsrsow.webp)

首部20bytes，标志位从第13个octet（8bits一组）算起

    |        		| 
    |---------------| 
    |C|E|U|A|P|R|S|F| 
    |---------------| 
    |7 6 5 4 3 2 1 0|

ACK $2^4$

SYN $2^1$ 

FIN $2^0$

如果只想要表示SYN包可以使用`tcp[13] == 2`，如果想表示包含SYN包的可以使用`tcp[13] == 2 & 2 == 2`，也可以使用name的形式`tcp-fin`, `tcp-syn`, `tcp-rst`, `tcp-push`, `tcp-ack`, `tcp-urg`.例如

```
tcpdump -i xl0 'tcp[tcpflags] & tcp-push != 0'
```

 只抓PUSH的包

## 例子

> 可以使用逻辑运算符('!' or 'not'；‘&&’ or  ‘and’；'||' or 'or')，也可以使用主机名，也可以使用子表达式

1. 
   
   ```
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

## TCP抓包

> 可以使用nc来测试

- 建立连接 `nc 1.1.1.1 80`

  ==tcpdump只会在第一个包以绝对的方式显示seq，之后都会以相对的方式显示seq，可以使用`-S`参数==

  ```
  cpl in ~ λ sudo tcpdump -nni wlp1s0 tcp and host 1.1.1.1
  tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
  listening on wlp1s0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
  12:59:39.143646 IP 30.226.76.42.52018 > 1.1.1.1.443: Flags [S], seq 2620238540, win 64240, options [mss 1460,sackOK,TS val 1232551602 ecr 0,nop,wscale 7], length 0
  12:59:39.390211 IP 1.1.1.1.443 > 30.226.76.42.52018: Flags [S.], seq 3357268989, ack 2620238541, win 65535, options [mss 1460,nop,nop,sackOK,nop,wscale 10], length 0
  12:59:39.390295 IP 30.226.76.42.52018 > 1.1.1.1.443: Flags [.], ack 1, win 502, length 0
  ```

- 关闭连接

  ```
  
  ```

  
