# Linux tcpdump11

ref:

https://www.tcpdump.org/

https://en.wikipedia.org/wiki/Promiscuous_mode

pcap(3PCAP)

## Digest

syntax

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

tcpdump 是一个基于 C 开发的 CLI 抓包工具，同样的还有 wireshark 出品的 tshark

## Terms

- buffer size

  packets that arrive for a capture are stored in a buffer, so that they do not have to be read by the application as soon as they arrive. On some platforms, the buffer’s size can be set; a size that’s too small could mean that, if too many packets are being captured and the snapshot length dosen’t limit the amount of data that’s buffered, packets could be dropped if the buffer fills up before the application can read packets from it, while a size that’s too large could use more non-pageable operating system memory than is necessary to prevent packets from being droppingq

- promiscuous mode

  On broadcast LANs such as Ethernet, if the network isn’t switched, or if the adapter is connected to a “mirror port” on a switch to which all packets passing through the switch are sent, a network adapter receives all packets on the LAN, including unicast or multicast packets not sent to a network address that the netwrok adapter isn’t configured to recognize

  Normally, the adapter will discard those packets; however, many network adpaters support “promiscuous mode”, which is a mode in which all packets, even if they are not sent to an mode in which all packets, even if they are not sent to an address that the adapter recognized, are provided to the host. This is useful for passively capturing traffic between two or more other hosts for analysis

  Note that even if an application dose not set promiscuous mode, the adpter could well be in promiscuous mode for some other reason

  ==For now, this doesn’t work on the “any” device; if an argument of “any” or NULL is supplied, the setting of promiscuous mode is ignored==

- monitor mode

  On IEEE 802.11 wireless LANs（可以直接理解成 WIFI）, even if an adapter is in promiscuous mode, it will supply to the host only frames for the network with which it’s associated. It might also supply only data frames

  In “monitor mode”, sometimes also called “rfmon mode” ( for “Radio Frequency MONitor” ), the adapter will supply all frames that it receives, with 802.11 headers, and might supply a pseudo-header with radia information about the frames as well.

  Note that in monitor mode the adapter might disassociate from the network with which it’s associated, so that you will not be able to use any wireless networks with that adapter. This could prevent accessing files on a network server, or resolving host names or network addresses, if you are capturing in monitor mode and are not connected to anohter network with which it’s associated, so that you will not be able to use any wireless networks with that adapter. 

  kali wifi attacker 就是使用 monitor mode 来监听数据包，然后 hack WIFI

## Statistics

tcpdump 后会抓包结束后显示抓包的数量

```
82 packets captured
157 packets received by filter
0 packets dropped by kernel
```

1. captured，this is the number of packets that tcpdump has received and processed

2. received by filter，the meaning of this depneds on the OS on which you are running tcpdump. If a filter was specified on the command line, on some OSes it counts packets regardless of whether they were matched by the filter expression and, even if they were matched by the filter expression, regardless of whether tcpdump has read and process them yet, on other OSes it counts only packets that were matched by the filter expression regardless of whether tcpdump has read and processed them yet, and on ohter OSes it counts only packets that were matched by the filter expression and were processed by tcpdump

   简单来说该值根据系统不同捕捉的数字也不同

3. dropped by kernel, this is the number of packets that were dropped, due to a lack of buffer space, by the packet capture machanism in the OS

## Optional args

### Common args

- `-D | --list-interfaces`

  print the list of the network interfaces available on the system and on which tcpdumo can capture packets

- `-L | --list-data-link-types`

  list the known data link types for the interface

- `-I | --monitor-mode`

  put the interface in “monitor mode”; thisi is supported only on IEEE 802.11 WIFI interface（只支持无线网卡）

- `-w file`

  将 stdout 内容保存到文件

- `-r file`

  读取`-w`生成的文件

### Input args

- `-Q | --direction=direction`

  ==只抓特定方向的包，值可以是`in`、`out`和`in-out`==

- `-i | --interface=interface`

  抓指定的NIC，如果没有指定且没有使用`-d`，默认抓 the system interface list for the lowest humbered, configured up interface( excluding loopback )

  ==after 2.2. kernel, an interface argument of any can be used to capture packets from all interfaces==

  if the `-D` flag is supported, an interface number as printed by that flag can be used as the interfaces argument

  ```
  cpl in ~ λ tcpdump -D
  1.wlp1s0 [Up, Running, Wireless, Associated]
  cpl in ~ λ sudo tcpdump -i 1 -c 10
  ```

  例如上述就表示抓 wlp1s0 无线网卡

- `-I | --monitor-mode`

  put the interface in “monitor mode”

- `-B | --buffer-size=buffer_size`

  set the OS capture buffer size, in units of KiB

- `-c count`

  exit after receiving count packets，抓指定个数的包

- `-e`

  ==print the link-level header on echo dump line==

- `-F file`

  use file as input for the filter expression. An additional expression given on the command line is ignored

- `-G rotate_seconds`

  every rotate_seconds seconds save the dump file named by `-w`

  例如如下表示每隔3秒，将抓包内容写入/tmp/a

  ```
  tcpdump -i any -G 3 -w /tmp/a
  ```

  通常`-w` 使用的文件名需要以`strftime`(c lib)的格式命名，如果没有使用该格式，文件会被复写。

  ```
  tcpdump -i any -G 3 -w %D.pcap
  ```

  如果和`-C`一起使用，文件名会以`file<count>`的格式命名（使用这种方式，文件名就可以不需要使用 `strftime`）

  例如如下表示每隔30秒抓包，如果文件大小超过 100 kib 就写入到新文件 

  ```
  tcpdump -i any -G 30 -C 100 -w tmp.pcap
  ```

  可以和`-W`一起使用，表示每隔多少秒，抓包并写到文件，一共抓多少个文件

  如下表示抓包30秒并写入到tmp.pcap

  ```
  tcpdump -i any -G 30 -W 1 -w tmp.pcap
  ```

  ==需要注意的一点时，如果在指定使用内没有抓到包，就不会写入到文件。==如果没有指定`-W`，需要使用 SIGINT 或类似的 posix signal 终止

- `-C file_size`

  before writing a raw packet to a savefile, check whether the file is currently larger than file_size and, if so , close the current savefile and open a new one. Savefiles after the first savefile will have the name specified with the `-w` flag, with a number after it, starting at 1 and countinuing upward

- `-W filecounot`

  和`-C`一起使用时，表示最多生成文件的个数

  和`-G`一起使用时，表示最多生成文件的个数

  和`-G`以及`-C`一起使用时，忽略该选项

- `-l`

  make stdout line buffered

  几乎等价于`-w`

  ```
  tcpdump -l > dat
  tcpdump -l > dat & tail -f data
  ```

### Output args

- `-#|--number`

  print an optional packet number at the beginning of the line
  
  多一列表示packet的序号

- `-S | --absolute-tcp-sequence-numbers`

  以绝对值输出 tcp seq number

- `-A`

  print echo packet in ASCII，明文输出包 segment

- `-q`

  输出的内容更精简

- `-n`

  don’t convert address，将 host address 和 port number 解析

  某些版本需要repeat改参数才可以实现port number 节解析

- `-j | --time-stamp-type=tstamp_type`

  设置时间戳，具体可用值参考`pcap-tstamp`

- `-t|-tt|-ttt|-tttt|-ttttt`

  具体看man page，==其中包含 delta time 设置==

- `-v|-vv|-vvv`

  verbose out，具体查看 man page

## Filter Expressions

> man pcap-filter

过滤表达式，由多个 premitives 组成，premitives 由 id 组成，id 由多个 qulifier 组成。qulifier 可以是如下几种

1. type

   what kind of thing the id name or number refers to  . Possible types are `host`, `net`,  `port` and `portrange`. ==If there is no type qualifier, `host` is assumed.==

2. dir

   dir qualifiers specify a particular transfer direction to and/or from id. Possible protos are: `ether`, `fddi`, `tr`, `wlan`, `ip`, `ip6`, `arp`, `rarp`, `decnet`, `tcp` and `udp`

3. 









