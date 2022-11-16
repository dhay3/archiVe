# IPtables-extensions connbytes Module

## Digest

如果想要显示报文的个数或者是大小，IPtables 中同样提供了一个 connbytes 模块。不仅限于面向连接的协议，对非面向连接的协议也适用

## Optional args

- `[!] --connbytes from[:to]`

  match  packets  from a connection whose packets/bytes/average packet size is more than FROM and less than TO bytes/packets. if  TO  is  omitted  only  FROM check is done. "!" is used to match packets not falling in the range.

- `--connbytes-mode {packets|bytes|avgpkt}`

  决定`--connbytes`的数值对应的类型

- `--connbytes-dir {original|reply|both}`

  决定 `--connbytes`的计数方式

## Examples

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

因为 iptables 的规则指定从 192.168.3.1 来的包或者是回的包都会被计数（both），当在`3:5`之间时就会执行 DROP target，所以 icmp seq 2 192.168.1.1 不回包