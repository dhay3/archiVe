# IPtables-extensions Match Modules

ref:

https://www.zsythink.net/archives/1544

## Digest

IPtables 除了提供一些简单的 matches，还按照场景提供了大量的 matches modules。只写一些常用的 match modules，具体参考 man page

## Modules

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

具体查看 [IPtables-extensions connbytes Match]()

#### connlimit

具体查看 [IPtables-extensions connlimit Match]()

#### conntrack

具体查看 [IPtables-extensions conntrack/state Match]()

#### state

具体查看 [IPtables-extensions conntrack/state Match]()

#### recent

具体查看 [IPtables-extensions rcent Match]()

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

匹配报文的 ascii 内容，一般用于过滤 URL

- `--algo {bm|kmp}`

  选择字符串的匹配模式，指定

- `[!] --string pattern`

  matches the given pattern

- `--icase`

  Ignore case when searching

```
iptables  -A  INPUT  -p  tcp  --dport  80 -m string --algo bm --string 'example' -j DROP
```

如果使用了 TLS 就不能按照字符串匹配了，因为加密了，但是可以指定过滤 DNS 报文

```
iptables  -A  INPUT  -p  udp  --dport  53 -m string --algo bm --string 'example' -j DROP
```

#### time

具体参考 [IPtables-extension time Match]()

#### hashlimit





