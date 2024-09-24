# IPtables-extensions MASQUERADE Target

ref

https://www.zsythink.net/archives/1764

## Digest

只在 nat table 中的 POSTROUTING chain 中生效

用于 dynmically assigned IP connections 动态伪装源 IP 。即从配置策略的机器的上分配可用的 IP 自动做 SNAT, 如果 IP 不可用就不分配

和 SNAT taregt 的区别就是 SNAT target 需要指定具体的 IP 或者 IP 段，但是 MASQUERADE 不需要只需要指定 NIC 即可

## Optional args

- `-to-ports port[-port]`

  指定伪装的源端口，只能在 tcp, udp, dccp, sctp 模块中被使用

- `--random`

  随机分配源端口

- `--random-fully`

  随机分配源端口

## Examples

以 192.168.1.1 访问 192.168.5.2 自动 SNAT 成 labos-3   eth0 配置的 IP 为例

![2022-11-11_00-26](https://github.com/dhay3/image-repo/raw/master/20221110/2022-11-11_00-26.77dy5jt8no1s.webp)

### Lab1

labos-3 网络配置

```
[root@labos-3 /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
14: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether b2:35:90:9b:a7:a7 brd ff:ff:ff:ff:ff:ff
    inet 192.168.4.1/24 scope global eth0
       valid_lft forever preferred_lft forever
    inet 192.168.4.3/24 scope global secondary eth0:0
       valid_lft forever preferred_lft forever
    inet6 fe80::b035:90ff:fe9b:a7a7/64 scope link 
       valid_lft forever preferred_lft forever
15: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether 9e:08:88:3a:ef:f8 brd ff:ff:ff:ff:ff:ff
    inet 192.168.5.1/24 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::9c08:88ff:fe3a:eff8/64 scope link 
       valid_lft forever preferred_lft forever
```

配置 MASQUERADE 规则

```
[root@labos-3 /]# iptables -t nat -A POSTROUTING -d 192.168.5.2 -s 192.168.1.1 -j MASQUERADE 
```

从 192.168.1.1 ping 192.168.5.2

```
[root@labos-1 /]# ping -s 192.168.1.1 192.168.5.2
PING 192.168.5.2 (192.168.5.2) 192(220) bytes of data.
200 bytes from 192.168.5.2: icmp_seq=1 ttl=61 time=41.1 ms
200 bytes from 192.168.5.2: icmp_seq=2 ttl=61 time=32.4 ms
```

在 labos-4 上抓包

这里已经看到，labos-4 收到的报文已经自动 SNAT 192.168.5.1 了

```
    192.168.5.1 > 192.168.5.2: ICMP echo request, id 26, seq 1, length 200
16:23:47.114934 b6:8a:65:67:3b:72 > 9e:08:88:3a:ef:f8, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 64, id 33662, offset 0, flags [none], proto ICMP (1), length 220)
    192.168.5.2 > 192.168.5.1: ICMP echo reply, id 26, seq 1, length 200
16:23:48.106863 9e:08:88:3a:ef:f8 > b6:8a:65:67:3b:72, ethertype IPv4 (0x0800), length 234: (tos 0x0, ttl 61, id 32721, offset 0, flags [DF], proto ICMP (1), length 220)
```