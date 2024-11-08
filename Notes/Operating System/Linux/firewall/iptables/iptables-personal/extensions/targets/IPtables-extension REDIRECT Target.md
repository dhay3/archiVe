# IPtables-extension REDIRECT Target

## Digest

REDIRECT target 只能用在 nat table 中的 PREOUTING 和 OUTPUT 中使用，或者是 user-defined chain 被上述条件调用

和 DNAT 很像，但是 REDIRECT 是将匹配的目的报文转发到本机。也可以理解层就是 DNAT 的子集

It redirects the packet to the machine itself by changing the destination IP to the primary address of the incoming interface (locally-generated packets are mapped to the localhost address, 127.0.0.1 for IPv4 and ::1 for IPv6, and  packets arriving on interfaces that don't have an IP address configured are dropped)

## Optional args

- `--to-ports port[-port]` 

  This  specifies  a destination port or range of ports to use: without this, the destination port is never altered.

## Example

以 192.168.1.1 访问 192.168.5.2 为例子

![2022-11-11_01-08](https://github.com/dhay3/image-repo/raw/master/20221110/2022-11-11_01-08.3insanea9j9c.webp)

labos-3 配置规则

```
[root@labos-3 /]# iptables -t nat -A PREROUTING -d 192.168.5.2 -j REDIRECT
```

从 192.168.1.1 ping 192.168.5.2

```
[root@labos-1 /]# ping 192.168.5.2
PING 192.168.5.2 (192.168.5.2) 56(84) bytes of data.
64 bytes from 192.168.5.2: icmp_seq=1 ttl=62 time=34.5 ms
```

在 labos-4 抓包

没有抓到包

```
[root@labos-4 /]# tcpdump -nvei eth0
dropped privs to tcpdump
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes

^C
0 packets captured
0 packets received by filter
0 packets dropped by kernel
```

在 labos-3 抓包

可以抓到包, 这里可以看到 labos-3 回包的 3 层 IP 是不变的，所以源端并不感知，和 DNAT 一个道理

```
17:03:06.619019  In ca:02:09:8d:00:38 ethertype IPv4 (0x0800), length 236: (tos 0x0, ttl 62, id 57440, offset 0, flags [DF], proto ICMP (1), length 220)
    192.168.1.1 > 192.168.5.2: ICMP echo request, id 30, seq 2, length 200
17:03:06.619058 Out b2:35:90:9b:a7:a7 ethertype IPv4 (0x0800), length 236: (tos 0x0, ttl 64, id 5794, offset 0, flags [none], proto ICMP (1), length 220)
    192.168.5.2 > 192.168.1.1: ICMP echo reply, id 30, seq 2, length 200
```

