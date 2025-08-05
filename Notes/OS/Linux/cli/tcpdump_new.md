---
createTime: 2025-04-15 16:33
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---
# Linux tcpdump_new

## 0x01 Preface

`tcpdump`[^1] 是一个基于 `libpcap`[^2] 开发的网络抓包 CLI 工具，类似的还有 wireshark 的 `tshark`（可以直接使用 wireshark filter expression 过滤）

## 0x02 EBNF

```
tcpdump [ -AbdDefhHIJKlLnNOpqStuUvxX# ] [ -B buffer_size ]
			 [ -c count ] [ --count ] [ -C file_size ]
			 [ -E spi@ipaddr algo:secret,...  ]
			 [ -F file ] [ -G rotate_seconds ] [ -i interface ]
			 [ --immediate-mode ] [ -j tstamp_type ] [ -m module ]
			 [ -M secret ] [ --number ] [ --print ] [ -Q in|out|inout ]
			 [ -r file ] [ -s snaplen ] [ -T type ] [ --version ]
			 [ -V file ] [ -w file ] [ -W filecount ] [ -y datalinktype ]
			 [ -z postrotate-command ] [ -Z user ]
			 [ --time-stamp-precision=tstamp_precision ]
			 [ --micro ] [ --nano ]
			 [ expression ]
```

如果没有使用 `-i` 指定监听的 NIC，默认会使用 `-D` 中显示的第一个 NIC(不包括 loopback)

```
$ tcpdump -D
1.wlp5s0 [Up, Running, Wireless, Associated]
2.meta0 [Up, Running, Connected]
3.any (Pseudo-device that captures on all interfaces) [Up, Running]
4.lo [Up, Running, Loopback]
5.enp4s0 [Up, Disconnected]
6.br-66d90a5c39b1 [Up, Disconnected]
7.docker0 [Up, Disconnected]
8.br-ebcbf65fa79f [Up, Disconnected]
9.bluetooth0 (Bluetooth adapter number 0) [Wireless, Association status unknown]
...
```

例如这里会监听 `wlp5s0`

```
$ sudo tcpdump
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on wlp5s0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
```

其中 expression 是过滤表达式，如果没有指定会抓取所有的包，反之只会抓取 expression boolean "true" 的包

> [!NOTE]
> 具体的 expression 可以参考 `man pcap-filter.7`

## 0x03 Terms[^3]

### 0x03a Buffer Size

> Packets that arrive for a capture are stored in a buffer, so that they do not have to be read by the application as soon as they arrive. On some platforms, the buffer’s size can be set; a size that’s too small could mean that, if too many packets are being captured and the snapshot length dosen’t limit the amount of data that’s buffered, packets could be dropped if the buffer fills up before the application can read packets from it, while a size that’s too large could use more non-pageable operating system memory than is necessary to prevent packets from being dropping

在系统中包并不会立即发送或者是被应用立即处理，而是会先存储在一个空间，这个空间被称为 buffer，在 Linux 中可以通过 `sudo sysctl -a | grep -P 'rmem|wmem'` 来查看(单位 byte)

```
$ sudo sysctl -a | grep -P 'rmem|wmem'
net.core.rmem_default = 262144
net.core.rmem_max = 67108864
net.core.wmem_default = 262144
net.core.wmem_max = 67108864
net.ipv4.tcp_rmem = 4096        102400  67108864
net.ipv4.tcp_wmem = 4096        102400  67108864
net.ipv4.udp_rmem_min = 4096
net.ipv4.udp_wmem_min = 4096
```

如果 buffer 设置的太小会导致包被系统直接丢弃，可以通过 `ethtool -S <NIC>` 来查看是否有被丢弃的包

```
$ ethtool -S ensp4s0
NIC statistics:
	 rx_packets: 3162478
	 rx_bytes: 2863210135
	 rx_duplicates: 0
	 rx_fragments: 75249
	 rx_dropped: 0
	 tx_packets: 839069
	 tx_bytes: 93148153
	 tx_filtered: 0
	 tx_retry_failed: 0
	 tx_retries: 149
	 sta_state: 4
	 txrate: 0
	 rxrate: 286700000
	 signal: 0
	 channel: 2437
	 noise: 160
	 ch_time: 56
	 ch_time_busy: 24
	 ch_time_ext_busy: 18446744073709551615
	 ch_time_rx: 18446744073709551615
	 ch_time_tx: 18446744073709551615
```

在 `tcpdump` 中可以直接通过 `--buffer-size=<buffer_size>` 来设置 buffer size，无需修改 kernel tunable paramters

### 0x03b Snapshot Length

> If, when capturing, you capture the entire contents of the packet, that requires more CPU time to copy the packet to your application, more disk and possibly network bandwidth to write the packet data to a file, and more disk space to save the packet. If you don't need the entire contents of the packet - for example, if you are only interested in the TCP headers of packets - you can set the "snapshot length" for the capture to an appropriate value. If the snapshot length is set to snaplen, and snaplen is less than the size of a packet that is captured, only the first snaplen bytes of that packet will be captured and provided as packet data

`tcpdump` 默认会截取每个包 262144 bytes 范围内的数据，可以直接理解为抓完整的包

```
listening on any, link-type LINUX_SLL2 (Linux cooked v2), snapshot length 262144 bytes
```

通常意味需要消耗更多的 CPU/buffer/带宽/存储 资源。而我们分析包的时候，通常只分析包头(少数情况下需要分析 payload)，就可以通过 `-s <snapshot>`(单位 byte) 来抓取指定长度的包，这个包也被称为 snapshot

例如只想分析 TCP header 就可以使用 `tcpdump -i any -s 54 tcp` 来抓取每个包前 54 bytes 的数据($14\ bytes\ Ethernet\ header\ +\ 20\ bytes\ IP\ header\ +\ 20\ bytes\ TCP\ header$)，如果不带额外的 options 的话

### 0x03c Mode

#### Promiscuous Mode

> On broadcast LANs such as Ethernet, if the network isn’t switched, or if the adapter is connected to a “mirror port” on a switch to which all packets passing through the switch are sent, a network adapter receives all packets on the LAN, including unicast or multicast packets not sent to a network address that the netwrok adapter isn’t configured to recognize
> 
> Normally, the adapter will discard those packets; however, many network adpaters support “promiscuous mode”, ==which is a mode in which all packets, even if they are not sent to an address that the adpater recognizes, are provided to the host.==This is useful for passively capturing traffic between two or more other hosts for analysis
> 
> Note that even if an application dose not set promiscuous mode, the adpter could well be in promiscuous mode for some other reason
> 
> ==For now, this doesn’t work on the “any” device; if an argument of “any” or NULL is supplied, the setting of promiscuous mode is ignored==

switch 的出现可以将网络划分为 2 类

- isn't switched network(hubed network) 由 hub 构建的网络
- switched network 由 switch 构建的网络

hubed network 和 switched network 最大的区别就在于，hub 不具备 MAC 学习的能力，而 switch 具备。所以 hubed network 会把收到的 unicast/multicast 包转发到除收包端口外的所有端口，而 switched network 会根据 unicast/multicast 目的 MAC 把包转发到特定的端口。当主机收到 unicast/multicast 包时，如果目的 MAC 和本机的不一致，就会被 NIC 丢弃，而一些 NIC 可以设置为 promiscuous mode(混杂模式)，从而将这些本应该丢弃的包发往本机

> [!NOTE]
> 在 switched network 中混杂模式通常没啥用，因为 switch 会按照目的 MAC 过滤，所以本机也就几乎不会存在 unicast/multicast 目的 MAC 和本机不一致的情况。现今只有在 mirror port(span 镜像端口) 的场景中会用到混杂模式

`tcpdump` 通过 `-i <interface>` 调用 `libpcap` 将指定 NIC 置为混杂模式（并不一定会在 `iproute2` 或者 `ifconfig` 中显示[^4]），使用 `-p` 来强制关闭混杂模式，但是有几种情况不支持混杂模式

- `-i` 使用 any pseudo-interface(`tcpdump -i any`)
- `-i` 使用 WIFI NIC(`tcpdump -i wlpxx`)

> [!IMPORTANT]
> Promiscuous mode is often used to diagnose network connectivity issues. There are programs that make use of this feature ot show the user all the data being transferrred over the network. Some protocols like FTP and Telnet transfer data and passwords in clear text, without encryption, and network scanners can see this data. Therefore, computer users are encourage to stay away from insecure protocols like telnet and use more ones such as SSH
> 
> 但是需要注意一点是混杂模式并不安全。例如宿主机的 NIC 设置成了混杂模式，那么宿主机上的所有使用了 bridge 模式的虚拟机都能获取到其他虚拟机上的数据包

#### Monitor Mode

> On IEEE 802.11 wireless LANs, even if an adapter is in promiscuous mode, it will supply to the host only frames for the network with which it’s associated. It might also supply only data frames
> 
> In “monitor mode”, sometimes also called “rfmon mode” ( for “Radio Frequency MONitor” ), the adapter will supply all frames that it receives, with 802.11 headers, and might supply a pseudo-header with radia information about the frames as well.
> 
> ==Note that in monitor mode the adapter might disassociate from the network with which it’s associated, so that you will not be able to use any wireless networks with that adapter==. This could prevent accessing files on a network server, or resolving host names or network addresses, if you are capturing in monitor mode and are not connected to anohter network with which it’s associated, so that you will not be able to use any wireless networks with that adapter. 

在 802.11 wireless LANs(WIFI) 中，即使 WIFI NIC 被置为混杂模式，也只会接收关联本机的包。想要达到混杂模式的效果，需要将 WIFI NIC 置为 monitor mode(监听模式，也被称为 rfmon mode)，`aircrack-ng` 就是使用 monitor mode 来监听数据包的

==需要注意的一点是监听模式会将网络断开==

#### Immediate Mode

> In immediate mode, packets are always delivered as soon as they arrive, with no buffering. 

包直接跳过系统的 buffer 打到 `tcpdump`，可以使用 `--immediate-mode` 来开启

## 0x04 Optional Args

> 具体请查看 `man tcpdump`

### 0x04a Common Args

- `-D | --list-interfaces`

	打印出可以使用的 NIC

- `-w <file>`

	将 stdout 内容保存到文件

- `-r <file>`

	读取 `-w` 生成的文件

### 0x04b Mode Related Args

- `-i | --interface=<interface>`

	监听指定 NIC，并将其置为 promiscuous mode，除 pseudo-interface any 外

- `-I | --monitor-mode`

	将 Wireless NIC 置为 monitor mode

- `--immediate-mode`

	将 NIC 置为 immedidate mode

- `-p | --no-promiscucous-mode`

	NIC 不会置为 promiscuous mode

### 0x04c Input Related Args

- `-B | --buffer-size=<buffer_size>`

	临时设置系统的 buffer size，单位 1024bytes

- `-F <file>`

	使用文件作为 filter expression，命令行的 filter expression 会被忽略

- `-Q | --direction=<direction>`

	==只抓特定方向的包，值可以是 `in`，`out` 或者是 `in-out`==

- `-c <count>`

	抓指定个数的包后退出 `tcpdump`

- `-s | --snapshot-length=<snaplen>`

	​每个包只截取 snaplen byte 大小，而不是默认抓完整的包

- `-e`

	==打印 link-level header，在 802.q 的场景下很有用，可以看到 vlan id==

- `-C <file_size>`

	需要和 `-w <file>` 一起使用，当抓的包累计大于 `file_size * 1,000,000 bytes` 时会写入到一个新文件，文件名会以 `file<count>` 的格式命名

	```
	$ sudo tcpdump  -i any -C 1 -w a.pcap
	$ ls | grep a.pcap
	  a.pcap a.pcap1 a.pcap2 
	```

- `-G <rotate>`

	每隔 rotate seconds 将内容存到 `-w` 指定的文件中

	如下表示每隔 3 秒，将抓包内容写入文件 `/tmp/a`

  ```
	tcpdump -i any -G 3 -w /tmp/a
  ```

	通常 `-w` 使用的文件名需要以 `strftime` 的格式命名，如果使用固定文件名，文件会被复写

  ```
	tcpdump -i any -G 3 -w %F.pcap
  ```

	如果和 `-C <file_size>` 一起使用，文件名会以 `file<count>` 的格式命名。如下表示每隔 30 秒抓包，如果文件大小超过 100 m就写入到新文件 

  ```
	tcpdump -i any -G 30 -C 100 -w tmp.pcap
  ```

	还可以和 `-W <file_count>` 一起使用，表示每隔多少秒，抓包并写到文件，一共抓多少个文件。如下表示每隔 30 秒抓包并写入到 `%F.pcap`，一共抓 2 个文件

  ```
	tcpdump -i any -G 30 -W 2 -w %F.pcap
  ```

	==需要注意的一点时，如果在指定使用内没有抓到包，就不会写入到文件。==如果没有指定 `-W`，需要使用 SIGINT 或类似的 posix signal 终止

- `-W <filecounot>`

	和 `-C` 一起使用时，表示最多生成文件的个数

	和 `-G` 一起使用时，表示最多生成文件的个数

	和 `-G` 以及 `-C` 一起使用时，忽略该选项

- `-U | --packet-buffered`

	通常和 `-w` 一起使用，`tcpdump` 会将抓取的包立即写入 file 中，而不是默认的以 chunk 的形式写入(可能会导致数据丢失)

### 0x04d Output Related Args

- `-#|--number`

	多一列表示包的序号

	```
	1  15:38:38.031051 wlp5s0 In  IP 192.168.137.1.3389 > 192.168.137.141.55306: Flags [P.], seq 2505589589:2505589640, ack 1827918166, win 63671, length 51
	2  15:38:38.031074 wlp5s0 Out IP 192.168.137.141.55306 > 192.168.137.1.3389: Flags [.], ack 51, win 3612, length 0
	```

- `-S | --absolute-tcp-sequence-numbers`

	以绝对值输出 Tcp Sequence Number

- `-A`

	以 ASCII 的形式，明文输出包的内容

- `-q`

	输出的内容更精简

	```
	15:41:26.925695 wlp5s0 In  IP 192.168.137.1.3389 > 192.168.137.141.55306: tcp 51
	15:41:26.925709 wlp5s0 Out IP 192.168.137.141.55306 > 192.168.137.1.3389: tcp 0
	```

- `-n`

	解析 host address 和 port number，某些版本需要 repeat 该参数才可以实现 port number 解析

- `-t|-tt|-ttt|-tttt|-ttttt`

	具体看man page，==其中包含 delta time 设置==

- `-v|-vv|-vvv`

	详细输出包内内容，具体查看 man page

## 0x05 Output Format

`tcpdump` 根据协议不同输出的内容的也不同，但是通常有 3 列是固定的

```
16:29:45.650367 wlp5s0 In PROTO
```

1. timestamp，通常以 `hh:mm:ss.frac` 的格式输出，由 kernel 计算而来。可以使用 `-t` 来改变 timestamp 格式
2. interface，监听的 NIC 会在 timestamp 后输出
3. direction，包的方向
4. PROTO，包解析的协议

一些版本可能只有 timestamp 是固定的

### 0x05a Link Level Headers

`tcpdump` 默认不会输出 2 层的报文头，可以使用 `-e` 来输出，PROTO 部分会显示为 ifindex

```
16:49:05.987885 wlp5s0 In  ifindex 3 2e:33:58:0a:b4:82 192.168.137.1.3389 > 192.168.137.141.55306: tcp 66
16:49:05.987998 wlp5s0 Out ifindex 3 a8:3b:76:1e:25:b1 192.168.137.141.55306 > 192.168.137.1.3389: tcp 74
```

### 0x05b ARP Packets

PROTO 部分会显示为 ARP

```
17:15:41.714116 wlp5s0 Out ARP, Request who-has 192.168.137.2 tell 192.168.137.141, length 28
```

### 0x05c IPv4 packets

PROTO 部分会显示为 IP，需要结合 `-v` 才会显示 IP 报文头的其他细节

```
18:03:40.576127 meta0 Out IP (tos 0x0, ttl 64, id 18944, offset 0, flags [DF], proto ICMP (1), length 84)
           100.64.0.1 > 110.242.68.66: ICMP echo request, id 7, seq 1, length 64
```

通常包含 `tos`, `ttl`, `id`, `offset`, `flags [flags]`, `proto`, `length`, `options` 几个字段

`flags` 通常是 `MF` 或者 `DF`，如果显示的值是 `+` 就表示 `MF`(the more fragments)，如果显示的值是 `DF` 就表示 `DF`(don’t fragments)，如果显示的 `.` 就表示没有设置分片位

### 0x05d TCP Packets

PROTO 部分显示为 IP(TCP Over IP)，需要结合 `-v` 才会显示 TCP 报文头的其他细节

```
src > dst: Flags [tcpflags], seq data-seqno, ack ackno, win window, urg urgent, options [opts], length len
```

例如

```
17:56:13.875338 wlp5s0 In  IP (tos 0x74, ttl 49, id 45573, offset 0, flags [DF], proto TCP (6), length 60)
       110.242.68.66.80 > 192.168.137.141.39716: Flags [S.], cksum 0x500b (correct), seq 591794355, ack 3988398177, win 8192, options [mss 1452,sackOK,nop,nop,nop,nop,nop,nop,nop,nop,nop,nop,nop,wscale 5], length 0
```

- `tcpflags` 标志位，通常是如下的值 `S(SYN)`, `.(ACK)`, `F(FIN)`, `P(PUSH)`, `R(RST)`, `U(URG)`, `W(ECN CWR)`, `E(ECN-Echo)`, `none(if no flags are set)`
- `data-seqno` SEQ 序列号，例如 `seq 2:21` 表示当前 SEQ 序列号为 2，Next SEQ 序列号为 21
- `ackno` ACK 序列号
- `window` 系统 receive buffer 剩余可用的大小
- `urgent` 是否携带了 URG 标志位
- `opts` TCP options，例如 `mss`，`wscale`

### 0x05e DNS Packets

PROTO 部分显示为 IP

```
src > dst:id op rcode flags qtype qclass name(len)
```

```
17:40:38.316295 IP 192.168.137.141.44355 > 192.168.137.1.53: 39676+ A? baidu.com. (27)
17:40:38.322220 IP 192.168.137.1.53 > 192.168.137.141.44355: 39676 2/0/0 A 39.156.66.10, A 110.242.68.66 (59)   
```

## 0x06 Statistics

tcpdump 抓包结束后会显示抓包的数量

```
82 packets captured
157 packets received by filter
0 packets dropped by kernel
```

- captured，this is the number of packets that tcpdump has received and processed

	`tcpdump` 按照 expression 抓到的包数量

- received by filter，the meaning of this depneds on the OS on which you are running tcpdump. If a filter was specified on the command line, on some OSes it counts packets regardless of whether they were matched by the filter expression and, even if they were matched by the filter expression, regardless of whether tcpdump has read and process them yet, on other OSes it counts only packets that were matched by the filter expression regardless of whether tcpdump has read and processed them yet, and on ohter OSes it counts only packets that were matched by the filter expression and were processed by tcpdump

	系统判断 `tcpdump` 抓到的包数量，系统不同判断的方式也不同

- dropped by kernel, this is the number of packets that were dropped, due to a lack of buffer space, by the packet capture machanism in the OS

	系统丢包数量，不包括 `iptables`/`ntftables` 的丢包，只计算 buffer 不足情况下的丢包数量

## 0x07 Filter Expressions

> man pcap-filter
>
> ==需要注意的是 filter expressions 应用于每一个包，和 wireshark filter 区别比较大==

过滤表达式，由多个 premitives 组成，premitives 由 id 组成，id 由多个 qulifier 组成。qulifier 可以是如下几种

1. type

   what kind of thing the id name or number refers to  . Possible types are `host`, `net`,  `port` and `portrange`. ==If there is no type qualifier, `host` is assumed.==

2. dir

   dir qualifiers specify a particular transfer direction to and/or from id. Possible protos are: `ether`, `fddi`, `tr`, `wlan`, `ip`, `ip6`, `arp`, `rarp`, `decnet`, `tcp` and `udp`

3. proto

   proto qualifiers restrict the match to a particular protocol. Possibel protos are：`ether`, `fddi`, `tr`，`wlan`，`ip`，`ip6`，`arp`，`rarp`，`decent`，`tcp` and `udp`. ==If there is no proto qualifier, all protocols consistent with the type are assumed==

4. logical expression

   逻辑表达式, `and(&&)`, `or(||)`, `not(!)`

### 0x07a Primitives

顾名思义

- `src|dst host <host>`
	`host <host>`

	按照 IPv4/v6 address 过滤

- `ether src|dst <ehost>`
	`ether host <ehost>`

	按照 MAC address 过滤，格式可以是 `xx:xx:xx:xx:xx:xx`， `xx-xx-xx-xx-xx-xx`，`xxxx.xxxx.xxxx`，`xxxxxxxxxxxx`

- `gateway <host>`

	按照 gateway IPv4/v6 address 过滤

- `src|dst net <net>`
- `net <net>`

	按照 prefix(网络位) 过滤，格式可以是 `192.168.1.0`，`192.168.1`，`172.16`，`10`

  an IPv4 network number can be written as a dotted quad(e.g., 192.168.1.0)，dotted triple (e.g., 192.168.1), dotted pair (e.g., 172.16), or single number (e.g., 10); the netmask is 255.255.255.255 for a dotted quad (which means that it;s really a host match), 255.255.255.0 for a dotted triple, 255.255.0.0 for a dotted pair, or 255.0.0.0 for a single number

- `net <net> mask <netmask>`
	`net net/len`

	按照 CIDR prefix(网络位) 过滤

- `src|dst port <port>`
	`port <port>`

	按照端口过滤

- `src|dst portrange <port1-port2>`
	`portrange <port1-port2>`

	按照端口范围过滤

- `less|greater <length>`

- `proto \<protocol>`

  Proto 过滤器用来过滤某个协议的数据，关键字为 `proto`，可省略。proto 后面可以跟上协议号或协议名称，支持 `icmp`, `igmp`, `igrp`, `pim`, `ah`, `esp`, `carp`, `vrrp`, `udp`和 `tcp`。因为通常的协议名称是保留字段，所以在于 proto 指令一起使用时，必须根据 shell 类型使用一个或两个反斜杠（\）来转义

- `tcp,udp,icmp`

  abbreviation for `proto \protocol`

- `ether proto <protocol>`

- `ip,arp`

  abbreviation for `ether proto \protocol`

- `inbound | outbound`

  packet was recieved | sent by the host performing the capturee rather than being sent | received by that host

- `vlan [vlan_id]`

  true if the packet is an IEEE 802.1Q VLAN packet

### 0x07b TCP Flag Filter

![Snipaste_2020-08-25_00-39-07](https://github.com/dhay3/image-repo/raw/master/20220719/Snipaste_2020-08-25_00-39-07.4kvfcqtsrsow.webp)

```
0                            15                              31
-----------------------------------------------------------------
|          source port          |       destination port        |
-----------------------------------------------------------------
|                        sequence number                        |
-----------------------------------------------------------------
|                     acknowledgment number                     |
-----------------------------------------------------------------
|  HL   | rsvd  |C|E|U|A|P|R|S|F|        window size            |
-----------------------------------------------------------------
|         TCP checksum          |       urgent pointer          |
-----------------------------------------------------------------
```

TCP header 通常 20 字节(octets)，除非指定了 TCP options。从 0 开始算，标志位出现在第 13 个字节

```
0             7|             15|             23|             31
----------------|---------------|---------------|----------------
|  HL   | rsvd  |C|E|U|A|P|R|S|F|        window size            |
----------------|---------------|---------------|----------------
|               |  13th octet   |               |               |
```

假设 13th octect number 是一个 8 bit unsinged integer，需要由 binary number 转成 decimal number

```
|C|E|U|A|P|R|S|F|
|---------------|
|0 0 0 0 0 0 1 0|
|---------------|
|7 6 5 4 3 2 1 0|
```

所以如果需要表示 SYN，那么可以使用`tcp[13] == 2` 表示 13th octect 的值是 `00000010 == 2`

如果需要表示 SYN + ACK，那么可以使用`tcp[13] == 18` 表是 13th octect 的值是 `00010010 == 18`

如果需要匹配含有某个标志位的包需要怎么办？这时就需要`&`操作(与计算)

`tcp[13] & 2 == 2`，表示 13th octect 的值 & 2 的值 一定是 2，即表示 一定包含 SYN。==需要注意的一点&在shell中有特殊的含有( 表示 async )，所有在 tcpdump 中需要将 filter expression 加上 single qutoed==。==同时抓 syn-ack 非常有助于对 TCP 异常的问题排查，例如 TCP 参数错误（通常都是由某几项 TCP options 导致，而 TCP options 都是在握手时协商预设的）导致连接异常，都是在 3 way-handshakes 中体现的==。同理的如果需要抓 RST 包，可以使用`tcp[13] & 4 == 4`

如果需要只匹配含有某个标志位的包需要怎么办？这时就需要组合表达式。例如

`tcp[13] & 2 == 2 and tcp[13] & 16 == 16` 就只会匹配含有 SYN 包的，不会匹配含有 ACK 包的

## 0x08 Domain Name Packets

`tcpdump` 和 `wireshark` 不一样， `tcpdump` 可以直接抓指定域名的包

```
$ sudo tcpdump -nni wlp1s0 host baidu.com 
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on wlp1s0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
22:17:13.680118 IP 192.168.2.194.52000 > 39.156.66.10.80: Flags [S], seq 1951551414, win 64240, options [mss 1460,sackOK,TS val 23985928 ecr 0,nop,wscale 7], length 0
22:17:13.733831 IP 39.156.66.10.80 > 192.168.2.194.52000: Flags [S.], seq 1780741195, ack 1951551415, win 8192, options [mss 1400,sackOK,nop,nop,nop,nop,nop,nop,nop,nop,nop,nop,nop,wscale 5], length 0
22:17:13.733907 IP 192.168.2.194.52000 > 39.156.66.10.80: Flags [.], ack 1, win 502, length 0
...
```

## 0x09 iptables

As a matter of fact, *tcpdump* is the first software found after the wire (and the NIC, if you will) on the way *IN*, and the last one on the way *OUT*.  

```
#IN Direction
Wire -> NIC -> tcpdump -> netfilter/iptables

#OUT Direction
iptables -> tcpdump -> NIC -> Wire
```

做一个小实验分别在 filter 表 INPUT 和 OUTPUT chian 添加规则并抓包

### OUTPUT Chain

```
cpl in ~ λ sudo iptables -t filter -A OUTPUT -d 39.156.66.10 -j DROP
cpl in ~ λ sudo iptables -nvL OUTPUT                             
cpl in ~ λ sudo iptables -t filter -nvL OUTPUT                        
Chain OUTPUT (policy ACCEPT 13 packets, 1003 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       0.0.0.0/0            39.156.66.10
```

ICMP 包

```
cpl in ~ λ ping 39.156.66.10
PING 39.156.66.10 (39.156.66.10) 56(84) bytes of data.
^C
--- 39.156.66.10 ping statistics ---
2 packets transmitted, 0 received, 100% packet loss, time 1008ms
```

抓 39.156.66.10 的包

```
cpl in ~ λ sudo tcpdump -nni any host 39.156.66.10  
tcpdump: data link type LINUX_SLL2
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on any, link-type LINUX_SLL2 (Linux cooked v2), snapshot length 262144 bytes
^C
0 packets captured
0 packets received by filter
0 packets dropped by kernel
```

未抓到，==同时也没有显示丢包, ip statistic 同样未显示==

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    RX:  bytes packets errors dropped  missed   mcast           
       5209039   16216      0       0       0       0 
    TX:  bytes packets errors dropped carrier collsns           
       5209039   16216      0       0       0       0 
2: wlp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DORMANT group default qlen 1000
    link/ether 64:bc:58:bd:a6:19 brd ff:ff:ff:ff:ff:ff
    RX:  bytes packets errors dropped  missed   mcast           
     174107124  188640      0       0       0       0 
    TX:  bytes packets errors dropped carrier collsns           
      16661891  111714      0       0       0       0
```

符合规则，在 iptables 层面丢包

### INPUT Chain

```
cpl in ~ λ sudo iptables -t filter -A INPUT -s 39.156.66.10 -j DROP 

cpl in ~ λ sudo iptables -nvL INPUT                                 
Chain INPUT (policy ACCEPT 5 packets, 518 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    1    84 DROP       all  --  *      *       39.156.66.10         0.0.0.0/0
```

抓 39.156.66.10 的包

```
$ sudo tcpdump -nni any host 39.156.66.10
tcpdump: data link type LINUX_SLL2
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on any, link-type LINUX_SLL2 (Linux cooked v2), snapshot length 262144 bytes
01:37:04.090673 wlp1s0 Out IP 192.168.2.194 > 39.156.66.10: ICMP echo request, id 8, seq 1, length 64
01:37:04.131768 wlp1s0 In  IP 39.156.66.10 > 192.168.2.194: ICMP echo reply, id 8, seq 1, length 64
^C
2 packets captured
2 packets received by filter
0 packets dropped by kernel
```

可以抓到发包和会包，==同时也没有显示丢包, ip statistic 同样未显示==

```
cpl in ~ λ ip -s link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    RX:  bytes packets errors dropped  missed   mcast           
       5209039   16216      0       0       0       0 
    TX:  bytes packets errors dropped carrier collsns           
       5209039   16216      0       0       0       0 
2: wlp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DORMANT group default qlen 1000
    link/ether 64:bc:58:bd:a6:19 brd ff:ff:ff:ff:ff:ff
    RX:  bytes packets errors dropped  missed   mcast           
     174648521  190837      0       0       0       0 
    TX:  bytes packets errors dropped carrier collsns           
      17067025  113850      0       0       0       0
```

符合规则，在 iptables 层面丢包

## 0x10 Tcpdump VS Wireshark

1. tcpdump 不能智能的分析重传的包，但是可以从 seq number 来分析。如果需要 TUI 类型的工具来分析，可以使用 tshark
2. tcpdump 使用的 filter expressions 和 wireshark 的大相径庭


## 0x11 Examples

同时匹配一个包源目IP是192.168.80.200 和 192.168.80.100

```
tcpdump -i ens33 host 192.168.80.200 && 192.168.80.100
```

匹配包中 TCP flags 只含有 SYN 的

```
tcpdump -i any 'tcp[13] & 2 == 2 and tcp[13] & 16 == 16'
```

抓30秒的包

```
tcpdump -i any -G 30 -W 1 -w /tmp/a.pcap
```

超详细抓包

```
tcpdump -nveSA#i ens0 
```

将包写入文件的同时，输出到 stdout

```
tcpdump -i any host 10.0.1.75 -w- -U | tee /tmp/a.pcap | tcpdump -nnr- 
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [tcpdump-little-book](https://nanxiao.github.io/tcpdump-little-book/)
- `man tcpdump`
- `man pcap.3pcap`
- `man pcap-filter.7`

***References***

[^1]:[GitHub - the-tcpdump-group/tcpdump: the TCPdump network dissector](https://github.com/the-tcpdump-group/tcpdump)
[^2]:[GitHub - the-tcpdump-group/libpcap: the LIBpcap interface to various kernel packet capture mechanism](https://github.com/the-tcpdump-group/libpcap)
[^3]:`man pcap.3pcap`
[^4]:[FAQ \| TCPDUMP & LIBPCAP](https://www.tcpdump.org/faq.html#q7)


[^1]:https://www.tcpdump.org/
[^2]:https://en.wikipedia.org/wiki/Promiscuous_mode
[^3]:https://superuser.com/questions/925286/does-tcpdump-bypass-iptables
[^4]:man pcap.3pcap
[^5]:https://stackoverflow.com/questions/25603831/how-can-i-have-tcpdump-write-to-file-and-standard-output-the-appropriate-data#25604237