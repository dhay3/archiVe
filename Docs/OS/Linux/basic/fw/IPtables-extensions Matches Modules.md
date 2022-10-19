# IPtables-extensions Match Modules

ref:

https://www.zsythink.net/archives/1544

只写一些常用的 match modules，具体参考 man page

## Modules

为了方便归类，这边按照协议进行划分 Modules

以 192.168.3.1 访问 192.168.1.1 为例

### IP

#### addrtype

IP地址模块

- `--src-type | --dst-type type`

  匹配报文 IP 地址是 type 的

type 的值可以是

- `UNSPEC` an unspecified address (i.e. 0.0.0.0)
- `UNICAST` an unicast address
- `LOCAL` a local address
- `BROADCAST` a broadcast address
- `ANYCAST` a anycast address
- `MULTICAST` a multicast address
- `BLACKHOLE` a blackhole address
- `UNREACHABLE` an unreachable address
- `PROHIBIT` an prohibited address

#### ttl



### ICMP

#### icmp

当指定 `-p icmp` 时自动加载, 支持如下参数

- `[!] --icmp-type {type[/code]|typename}`

例如

```
[root@netos-1 /]# iptables -t filter -A INPUT -p icmp -s 192.168.3.1 --icmp-type echo-request -j DROP 

[root@netos-1 /]# iptables -nvL INPUT
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       icmp --  *      *       192.168.3.1          0.0.0.0/0            icmptype 8
```

从 192.168.3.1 来的 icmp type 8 (echo-request) 在 192.168.1.1 上通过 iptables 丢掉

### TCP

#### tcp

当指定 `-p tcp` 时自动加载, 支持如下参数

- `[!] --source-port | --sport port[:port]`

  source port, a port number or a range of ports

- `[!] --destination-port | --dport port[:port]`

  destinatin port, a port number or a range of ports

- `[!] --tcp-flags mask comp`

  匹配指定的 TCP flags

  mask is the flags which we should examine, written as a comma-separated list

  comp is a comma-separated list of flags which must be set

  Flags are SYN ACK FIN RST URG PSH ALL NONE

  例如

  `iptables -A FORWARD -p tcp --tcp-flags SYN,ACK,FIN,RST SYN`

  只会匹配含有 SYN 但是不含有 ACK, FIN, RST 的报文

- `[!] --syn`

  只匹配 SYN 报文，不会匹配 SYN-ACK 报文，等价与 `--tcp-flags SYN,RST,ACK,FIN SYN`

- `[!] --tcp-option number`

  匹配指定 tcp option

#### tcpmss

tcp mss 模块，只会匹配 SYN，SYN-ACK 报文，因为 mss 只会在 TCP three-way handshaks 中宣告

- `[!] --mss value[:value]`

  match a give tcp mss value or range

#### multiport

matching a set oof source or destination ports

只能和 4 层协议一起使用，即 `-p` 可以使用的值是 tcp, udp, udplite, dccp 和 sctp

- `[!] --source-ports | --sports port,[port|,port:port]...`
- `[!] --destination-ports,--dports port[,port|,port:port]...`
- `[!] --ports port[,port|,port:port]...`

例如

```
[root@labos-1 /]# iptables -t filter -A INPUT -p tcp -s 192.168.3.1 -m multiport --dports 1,2:65535 -j DROP
[root@labos-1 /]# iptables -nvL INPUT
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       tcp  --  *      *       192.168.3.1          0.0.0.0/0            multiport dports 1,2:65535

[root@labos-2 /]# nc -nv 192.168.1.1 65533
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connection timed out.
```



#### connbytes

连接(这里的连接并不特指 TCP 报文，也包含 ICMP 报文，这里为了方便记忆就放在 tcp 下 )模块

- `[!] --connbytes from[:to]`

  match  packets  from a connection whose packets/bytes/average packet size is more than FROM and less than TO bytes/packets. if  TO  is  omitted  only  FROM check is done. "!" is used to match packets not falling in the range.

- `--connbytes-mode {packets|bytes|avgpkt}`

  决定`--connbytes`的数值对应的类型

- `--connbytes-dir {original|reply|both}`

  决定 `--connbytes`的计数方式

例如

192.168.3.1 访问 192.168.1.1

```
#192.168.1.1 iptables
[root@netos-1 /]# iptables -t filter -A OUTPUT -d 192.168.3.1 -m connbytes --connbytes 3:5 --connbytes-dir both --connbytes-mode packets -j DROP
[root@netos-1 /]# iptables -nvL OUTPUT
Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       0.0.0.0/0            192.168.3.1          connbytes 3:5 connbytes mode packets connbytes direction both

#192.168.3.1 ping
[root@netos-2 /]# ping 192.168.1.1
PING 192.168.1.1 (192.168.1.1) 56(84) bytes of data.
64 bytes from 192.168.1.1: icmp_seq=1 ttl=62 time=47.9 ms
64 bytes from 192.168.1.1: icmp_seq=3 ttl=62 time=37.3 ms

#192.168.1.1 tcpdump
12:48:20.990985 IP 192.168.3.1 > 192.168.1.1: ICMP echo request, id 8, seq 1, length 64
12:48:20.991014 IP 192.168.1.1 > 192.168.3.1: ICMP echo reply, id 8, seq 1, length 64
12:48:21.981651 IP 192.168.3.1 > 192.168.1.1: ICMP echo request, id 8, seq 2, length 64
12:48:23.007687 IP 192.168.3.1 > 192.168.1.1: ICMP echo request, id 8, seq 3, length 64
12:48:23.007710 IP 192.168.1.1 > 192.168.3.1: ICMP echo reply, id 8, seq 3, length 64
```

从 ping 和 tcpdump 的结果看，192.168.1.1 icmp seq 2 没有回包

因为 iptables 的规则指定从 192.168.1.1 来的包或者是回的包都会被计数（both），当在`3:5`之间时就会执行 DROP target，所以 icmp seq 2 192.168.1.1 不回包

#### connlimit

用来限制从 client 来的连接数，不仅对 TCP 生效，ICMP 也生效

- `--connlimit-upto n`

  match if the number of existing connections is below or equal n

- `--connlimit-above n`

  match if the number of exsiting connections is above n

```
# allow 2 telnet connections per client host
iptables -A INPUT  -p  tcp  --syn  --dport  23  -m  connlimit
--connlimit-above 2 -j REJECT

# you can also match the other way around:
iptables  -A  INPUT  -p  tcp  --syn  --dport  23 -m connlimit
--connlimit-upto 2 -j ACCEPT

# limit the number of parallel HTTP requests to 16 per class C sized
source network (24 bit netmask)
iptables    -p    tcp   --syn   --dport   80   -m   connlimit
--connlimit-above 16 --connlimit-mask 24 -j REJECT

```

#### conntrack

用来匹配 connection 状态(不仅限于 TCP 报文，ICMP 报文同样生效)

- `[!] --ctstate statelist`

  statelist 可以一个值也可以是一个数组。具体可以是 

  - INVALID

    The packet is associated with no known connection

  - NEW

    The packet has started a new connection or otherwise  associated with a connection which has not seen packets in both directions

  -  ESTABLISHED

    The packet is associated with a  connection  which  has  seen packets in both directions.

  - RELATED

    The  packet  is  starting a new connection, but is associated with an existing connection, such as an FTP data transfer  or an ICMP error.

  - UNTRACKED

    The  packet  is  not tracked at all, which happens if you explicitly untrack it by using -j CT --notrack in the  raw  table

  - SNAT

    A virtual state, matching if the original source address differs from the reply destination

    即说明源IP做了 SNAT

  - DNAT

    A virtual state, matching if the original destination differs  from the reply source

    即说明目IP做了 DNAT

例如

```
[root@netos-1 /]# iptables -t filter -A INPUT -s 192.168.3.1 -m conntrack --ctstate NEW -j DROP

[root@netos-2 /]# ping 192.168.1.1
PING 192.168.1.1 (192.168.1.1) 56(84) bytes of data.
^C
--- 192.168.1.1 ping statistics ---
2 packets transmitted, 0 received, 100% packet loss, time 1019ms

[root@netos-1 /]# iptables -nvL INPUT
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    2   168 DROP       all  --  *      *       192.168.3.1          0.0.0.0/0            ctstate NEW
```

#### state

是 conntrack 的子集

### Mixin

#### recent

具体查看 IPtables-extensions rcent Module

#### comment

评论模块

- `--comment commetn`

  可以为 rule 添加 comment

  ```
  [root@localhost /]#iptables -A INPUT -i eth1 -m comment --comment "my local LAN"
  [root@localhost /]# iptables -nvL 
  Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
   pkts bytes target     prot opt in     out     source               destination         
    305 23542 ACCEPT     all  --  *      *      !192.168.1.150        0.0.0.0/0           
      0     0            all  --  eth1   *       0.0.0.0/0            0.0.0.0/0            /* my local LAN */
  ```

#### ipvs

匹配 ipvs（LVS）connection

- `[!] --ipvs`

  packet belongs to an IPVS connection

- `[!] --vporto protocol`

  VIP protocol

- `[!] --vaddr address[/mask]`

  VIP address to match

- `[!] --vport port`

  VIP port to match

#### devgroup

match device group of packet’s incoming/outgoing interface

- `[!] --src-group|--dst-group name`

  match device group of incoming/outgoing device

#### string

#### time
