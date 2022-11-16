# IPtables-extensions SNAT Target

ref

https://www.zsythink.net/archives/1764

## Digest

SNAT target 只能在 nat table 中的 POSTROUTING 和 INPUT 中使用，或者是 user-defined chain 被上述条件调用

## Optional args

> 需要注意的是，需要先指定 SNAT Target 才能使用 `--to-source` 参数

- `--to-source [ipaddr[-ipaddr]][:port[-port]]`

  which  can  specify a single new source IP address, an inclusive range of IP addresses. Optionally a port range,  if  the rule also specifies one of the following protocols: tcp, udp, dccp or sctp.  If no port range  is  specified,  then  source ports  below  512  will  be  mapped to other ports below 512: those between 512 and 1023 inclusive will be mapped to  ports below  1024, and other ports will be mapped to 1024 or above.Where possible, no port alteration will occur.

## Example

### Lab1

以从 192.168.1.1 访问 192.168.4.1 SNAT 成 192.168.4.3 为例子

![2022-11-02_13-03](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221102/2022-11-02_13-03.2xvk36zki4hs.webp)

设置 SNAT 规则

```
[root@labos-1 /]# iptables -t nat -A POSTROUTING -s 192.168.1.1 -j SNAT --to-source 192.168.4.3
```

从 192.168.1.1 ping 192.168.4.1

这里可以看到 100% 丢包，因为目前 192.168.4.3 和 192.168.4.1 在一个 LAN 中且并没有配置，所以在 ARP 的时候就会失败

```
[root@labos-1 /]# ping -s 192.168.1.1 192.168.4.1PING 192.168.4.1 (192.168.4.1) 192(220) bytes of data.
^C  
--- 192.168.4.1 ping statistics ---
173 packets transmitted, 0 received, 100% packet loss, time 176106ms
```

labos-3 抓包 

这里可以看到过来的报文源 IP 已经是 192.168.4.3 但是只有 request 没有 reply, 因为在 labos-3 arp 表中没有 192.168.4.3 的 mac

```
05:12:35.646453 ca:02:09:8d:00:38 > de:cd:45:9e:59:00, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 62, id 9060, offset 0, flags [DF], proto ICMP (1), length 220)
    192.168.4.3 > 192.168.4.1: ICMP echo request, id 9, seq 1, length 200
05:12:35.646473 de:cd:45:9e:59:00 > Broadcast, ethertype ARP (0x0806), length 42: Ethernet (len 6), IPv4 (len 4), Request who-has 192.168.4.3 tell 192.168.4.1, length 28
05:12:36.651080 de:cd:45:9e:59:00 > Broadcast, ethertype ARP (0x0806), length 42: Ethernet (len 6), IPv4 (len 4), Request who-has 192.168.4.3 tell 192.168.4.1, length 28

[root@labos-3 /]# ip n 
192.168.4.3 dev eth0  INCOMPLETE
192.168.4.2 dev eth0 lladdr ca:02:09:8d:00:38 STALE
```

### Lab2

以 192.168.1.1 访问 192.168.5.2 SNAT 成 192.168.4.1 为例子

![2022-11-10_22-31](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221110/2022-11-10_22-31.2te6bh9llqbk.webp)

labos-3 route 和 NIC 配置

```
[root@labos-3 /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
9: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether fe:83:0f:60:29:25 brd ff:ff:ff:ff:ff:ff
    inet 192.168.4.1/24 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::fc83:fff:fe60:2925/64 scope link 
       valid_lft forever preferred_lft forever
11: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether 3a:b1:f4:73:c0:87 brd ff:ff:ff:ff:ff:ff
    inet 192.168.5.1/24 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::38b1:f4ff:fe73:c087/64 scope link 
       valid_lft forever preferred_lft forever
[root@labos-3 /]# ip r
default via 192.168.4.2 dev eth0 
192.168.4.0/24 dev eth0 proto kernel scope link src 192.168.4.1 
192.168.5.0/24 dev eth1 proto kernel scope link src 192.168.5.1
```

设置 SNAT 规则

```
[root@labos-3 /]# iptables -t nat -A POSTROUTING -s 192.168.1.1 -d 192.168.5.2 -j SNAT --to-source 192.168.4.1        
```

从 192.168.1.1 ping 192.168.5.2

```
[root@labos-1 /]# ping -s 192.168.1.1 192.168.5.2
PING 192.168.5.2 (192.168.5.2) 192(220) bytes of data.
200 bytes from 192.168.5.2: icmp_seq=1 ttl=61 time=40.9 ms
200 bytes from 192.168.5.2: icmp_seq=2 ttl=61 time=31.5 ms
```

labos-4 抓包

这里可以看到来的报文的源 IP 已经是 192.168.4.1

```
    192.168.4.1 > 192.168.5.2: ICMP echo request, id 9, seq 1, length 200
14:56:08.333034 b6:8a:65:67:3b:72 > 3a:b1:f4:73:c0:87, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 64, id 53058, offset 0, flags [none], proto ICMP (1), length 220)
    192.168.5.2 > 192.168.4.1: ICMP echo reply, id 9, seq 1, length 200
14:56:09.342114 3a:b1:f4:73:c0:87 > b6:8a:65:67:3b:72, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 61, id 61165, offset 0, flags [DF], proto ICMP (1), length 220)
```

