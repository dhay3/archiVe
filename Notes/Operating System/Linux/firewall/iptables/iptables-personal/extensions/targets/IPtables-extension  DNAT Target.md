# IPtables-extension  DNAT Target

ref

https://www.zsythink.net/archives/1764

## Digest

DNAT target 只能在 nat table 中的 PREROUTING 和 OUTPUT 中使用，或者是 user-defined chain 被上述条件调用

## Optional args

- `--to-destination [[ipaddr[-ipaddr]][:port[-port[/baseport]]]`

  which can specify a single new destination IP address, an inclusive range of IP addresses. Optionally a port range, if the rule also specifies  one of the following protocols: tcp, udp, dccp or sctp.  If no port range is specified, then the destination port will never be modified. If no IP address is specified then only the destination port will be modified.  If baseport is given, the difference of the  original destination port and its value is used as offset into the mapping port range. This allows to create shifted portmap ranges and is available since kernel version 4.18.  For a single port or baseport, a service name as listed in /etc/services may be used.

  根据指定的参数决定是否使用 NAPT

## Example

### Lab1

以访问 39.156.66.10 DNAT 成 180.97.251.233 为例子

设置 DNAT 规则

```
sudo iptables -t nat -A OUTPUT -d 39.156.66.10 -j DNAT --to-destination 180.97.251.233

sudo iptables -t nat -nvL OUTPUT
Chain OUTPUT (policy ACCEPT 350 packets, 37995 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0   180 DNAT       all  --  *      *       0.0.0.0/0            39.156.66.10         to:180.97.251.233
```

这是使用 curl 访问 39.156.66.10

```
curl -svL 39.156.66.10
*   Trying 39.156.66.10:80...
* Connected to 39.156.66.10 (39.156.66.10) port 80 (#0)
> GET / HTTP/1.1
> Host: 39.156.66.10
> User-Agent: curl/7.84.0
> Accept: */*
> 
```

这里显示的是 `connected to 39.156.66.10` 实际应该是 180.97.251.233。因为应用输出到 stdout 的内容是最优先的(在 OSI7 之前)，但是可以通过抓包看到报文是已经 DNAT 的了

![2022-10-31_23-16](https://github.com/dhay3/image-repo/raw/master/20221031/2022-10-31_23-16.1noi1a38ay3k.webp)

因为 nfilter 在 tcpdump 或者 wireshark 之前也在 NIC 处理报文之前

另外这里还需要注意的一点是一次请求的所有报文 IPtables pkts 只会记录一次，并没按照报文个数计算

```
cpl in ~ λ sudo iptables -t nat -nvL OUTPUT
Chain OUTPUT (policy ACCEPT 13 packets, 2311 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    1    60 DNAT       all  --  *      *       0.0.0.0/0            39.156.66.10         to:180.97.251.233
```

### Lab2

以 192.168.1.1 访问 192.168.4.1 DNAT 成 192.168.5.2 为例子

![2022-11-02_13-34](https://github.com/dhay3/image-repo/raw/master/20221102/2022-11-02_13-34.4jnwlj1u7zwg.webp)

先看一下 labos-3 的 NIC 信息和路由

```
[root@labos-3 /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
16: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether ce:ed:90:c0:14:31 brd ff:ff:ff:ff:ff:ff
    inet 192.168.4.1/24 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::cced:90ff:fec0:1431/64 scope link 
       valid_lft forever preferred_lft forever
17: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether 1e:7c:22:23:e4:c8 brd ff:ff:ff:ff:ff:ff
    inet 192.168.5.1/24 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::1c7c:22ff:fe23:e4c8/64 scope link 
       valid_lft forever preferred_lft forever

[root@labos-3 /]# ip r
default via 192.168.4.2 dev eth0 
192.168.4.0/24 dev eth0 proto kernel scope link src 192.168.4.1 
192.168.5.0/24 dev eth1 proto kernel scope link src 192.168.5.1 
```

配置 DNAT 规则

```
[root@labos-3 /]# iptables -t nat -A PREROUTING -d 192.168.4.1 -j DNAT --to-destination 192.168.5.2 
```

从 192.168.1.1 ping 192.168.4.1

```
[root@labos-1 /]# ping -s 192.168.1.1 192.168.4.1
PING 192.168.4.1 (192.168.4.1) 192(220) bytes of data.
200 bytes from 192.168.4.1: icmp_seq=1 ttl=61 time=43.2 ms
200 bytes from 192.168.4.1: icmp_seq=2 ttl=61 time=34.10 ms
```

在 labos-4 上抓包

这里可以看到 labos-4 上以及收到了 labos-1 来的报文了

```
15:24:52.023679 3a:b1:f4:73:c0:87 > b6:8a:65:67:3b:72, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 61, id 36458, offset 0, flags [DF], proto ICMP (1), length 220)
    192.168.1.1 > 192.168.5.2: ICMP echo request, id 12, seq 1, length 200
15:24:52.023704 b6:8a:65:67:3b:72 > 3a:b1:f4:73:c0:87, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 64, id 59092, offset 0, flags [none], proto ICMP (1), length 220)
    192.168.5.2 > 192.168.1.1: ICMP echo reply, id 12, seq 1, length 200
15:24:53.016740 3a:b1:f4:73:c0:87 > b6:8a:65:67:3b:72, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 61, id 36675, offset 0, flags [DF], proto ICMP (1), length 220)
```

