# IPtables-extensions statistic Match

## Digest

iptables 还可以根据 counter 来匹配，这样我们就可以通过 iptables 实现负债均衡

## Optional args

- `--mode mode`

  设置 statistic 匹配的模式，可以是 random 和 nth

- `--probability p`

- `--every n`

  匹配每 nth 报文，只在 mode 为 nth 是生效

- `--packet p`

  设置 nth 初始的计数值 (只能在 $[0,n-1]$ 范围内)

## Examples

### Lab1

以 192.168.1.1 访问 192.168.3.1 为例子

![2022-11-15_15-14](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221115/2022-11-15_15-14.56e538jrk8w0.webp)

192.168.3.1 设置规则

```
iptables -t filter -A INPUT -s 192.168.1.1 -m statistic --mode nth --packet 0 --every 3 -j DROP
```

192.168.1.1 ping 192.168.3.1

每隔 3 个报文丢弃，从 0 开始计数( icmp_seq - 1)

```
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
64 bytes from 192.168.3.1: icmp_seq=2 ttl=62 time=31.3 ms
64 bytes from 192.168.3.1: icmp_seq=3 ttl=62 time=43.1 ms
64 bytes from 192.168.3.1: icmp_seq=5 ttl=62 time=38.7 ms
64 bytes from 192.168.3.1: icmp_seq=6 ttl=62 time=38.4 ms
64 bytes from 192.168.3.1: icmp_seq=8 ttl=62 time=34.8 ms
64 bytes from 192.168.3.1: icmp_seq=9 ttl=62 time=34.3 ms
64 bytes from 192.168.3.1: icmp_seq=11 ttl=62 time=41.2 ms
```

### Lab2

> 只能做简单的 round-robin 规则

例如 192.168.1.1 访问 192.168.4.2 负载均衡到 192.168.5.2 和 192.168.6.2 

![2022-11-21_22-34](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221117/2022-11-21_22-34.59hp8mxzyznk.webp)

192.168.4.1 配置规则

```
[root@labos-3 /]# iptables -t nat -A PREROUTING -p tcp -d 192.168.4.1 --dport 22 -m statistic --mode nth --every 2 --packet 0 -j DNAT --to-destination 192.168.5.2
[root@labos-3 /]# iptables -t nat -A PREROUTING -p tcp -d 192.168.4.1 --dport 22 -m statistic --mode nth --every 2 --packet 0 -j DNAT --to-destination 192.168.6.2
```

192.168.4.1 监听端口

```
#labos-3/labos-4/labos-5
nc -lvk 22
```

192.168.1.1 访问 192.168.4.1

可以看到在 192.168.5.2 上收到了 192.168.1.1 的请求了

```
[root@labos-4 /]# Ncat: Connection from 192.168.1.1.
Ncat: Connection from 192.168.1.1:51130.
GET / HTTP/1.1
Host: 192.168.4.1:22
User-Agent: curl/7.61.1
Accept: */*
```

看一下 192.168.4.1 上的计数

```
[root@labos-3 /]# iptables -t nat -nvL PREROUTING
Chain PREROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    1    60 DNAT       tcp  --  *      *       0.0.0.0/0            192.168.4.1          tcp dpt:22 statistic mode nth every 2 to:192.168.5.2
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            192.168.4.1          tcp dpt:22 statistic mode nth every 2 to:192.168.6.2
```

192.168.1.1 再访问 192.168.4.1

可以看到在 192.168.6.2 上收到了 192.168.1.1 的请求了

```
[root@labos-5 /]# Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Listening on :::22
Ncat: Listening on 0.0.0.0:22
Ncat: Connection from 192.168.1.1.
Ncat: Connection from 192.168.1.1:51132.
GET / HTTP/1.1
Host: 192.168.4.1:22
User-Agent: curl/7.61.1
Accept: */*
```

192.168.3.1 计数

```
[root@labos-3 /]# iptables -t nat -nvL PREROUTING
Chain PREROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    1    60 DNAT       tcp  --  *      *       0.0.0.0/0            192.168.4.1          tcp dpt:22 statistic mode nth every 2 to:192.168.5.2
    1    60 DNAT       tcp  --  *      *       0.0.0.0/0            192.168.4.1          tcp dpt:22 statistic mode nth every 2 to:192.168.6.2
```

