# IPtables-extensions REJECT Target 

## Digest

REJECT target 用于回送指定的报文，只在 INPUT, FORWARD, OUPUT chains 以及 user-defined chains 中生效

## Optional args

- `--reject-with type`

  type 可以是以下的值

  1. icmp-net-unreachable
  2. icmp-host-unreachable
  3. icmp-port-unreachable
  4. icmp-proto-unreachable
  5. icmp-net-prohibited
  6. icmp-host-prohibited
  7. icmp-admin-prohibited

  如果和`-p tcp` 一起使用 type 还可以是 tcp-reset

## Exmaples

### Lab01

以 192.168.1.1 ping 192.168.3.1 回送 icmp-port-unreachable 为例子

![2022-11-14_11-50](https://github.com/dhay3/image-repo/raw/master/20221114/2022-11-14_11-50.559yumdnaxog.webp)

192.168.3.1 设置规则

```
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.1.1 -j REJECT --reject-with icmp-port-unreachable
```

192.168.1.1 ping 192.168.3.1

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
From 192.168.3.1 icmp_seq=1 Destination Port Unreachable
From 192.168.3.1 icmp_seq=2 Destination Port Unreachable
```

这里可以看到 192.168.1.1 上已经收到了从 192.168.3.1 上回送的 icmp-port-unreachable

### Lab02

以 192.168.1.1 访问 192.168.3.1 80 端口回送 icmp-port-reachable

![2022-11-14_11-50](https://github.com/dhay3/image-repo/raw/master/20221114/2022-11-14_11-50.559yumdnaxog.webp)

192.168.3.1 设置规则

```
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.1.1 -j REJECT --reject-with icmp-port-unreachable
```

192.168.3.1 监听 80 端口

```
[root@labos-2 /]# nc -lvk 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Listening on :::80
Ncat: Listening on 0.0.0.0:80
```

192.168.1.1 curl 192.168.3.1:80

```
[root@labos-1 /]# curl -v 192.168.3.1
* Rebuilt URL to: 192.168.3.1/
*   Trying 192.168.3.1...
* TCP_NODELAY set
* connect to 192.168.3.1 port 80 failed: Connection refused
* Failed to connect to 192.168.3.1 port 80: Connection refused
* Closing connection 0
curl: (7) Failed to connect to 192.168.3.1 port 80: Connection refused
```

192.168.1.1 抓包

这里可以看到收到 192.168.3.1 报文是 icmp unreachable 并不是 TCP Reset 报文，但是 curl 显示的是 connection refuesd（==这是 curl 进程显示的内容和 TCP 并没有关联==）

```
03:59:20.222218 IP (tos 0x0, ttl 64, id 28048, offset 0, flags [DF], proto ICMP (1), length 84)
    192.168.1.1 > 192.168.3.1: ICMP echo request, id 4, seq 7, length 64
03:59:20.262353 IP (tos 0xc0, ttl 62, id 53529, offset 0, flags [none], proto ICMP (1), length 112)
    192.168.3.1 > 192.168.1.1: ICMP 192.168.3.1 protocol 1 port 48797 unreachable, length 92
        IP (tos 0x0, ttl 62, id 28048, offset 0, flags [DF], proto ICMP (1), length 84)
```

### Lab03

以 192.168.1.1 访问 192.168.3.1 为例

![2022-11-14_11-50](https://github.com/dhay3/image-repo/raw/master/20221114/2022-11-14_11-50.559yumdnaxog.webp)

192.168.3.1 设置规则

```
[root@labos-2 /]# iptables -t filter -A INPUT -p tcp -s 192.168.1.1 -j REJECT --reject-with  tcp-reset
```

192.168.3.1 监听 80 端口

```
[root@labos-2 /]# nc -lvk 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Listening on :::80
Ncat: Listening on 0.0.0.0:80
```

192.168.1.1 curl 192.168.3.1:80

```
[root@labos-1 /]# curl -v 192.168.3.1
* Rebuilt URL to: 192.168.3.1/
*   Trying 192.168.3.1...
* TCP_NODELAY set
* connect to 192.168.3.1 port 80 failed: Connection refused
* Failed to connect to 192.168.3.1 port 80: Connection refused
* Closing connection 0
curl: (7) Failed to connect to 192.168.3.1 port 80: Connection refused
```

192.168.1.1 抓包

这里可以看到 192.168.3.1 回送了 Reset 报文

```
04:12:59.187379 IP (tos 0x0, ttl 64, id 64713, offset 0, flags [DF], proto TCP (6), length 60)
    192.168.1.1.59756 > 192.168.3.1.http: Flags [S], cksum 0x849b (correct), seq 1702402993, win 64240, options [mss 1460,sackOK,TS val 3679878371 ecr 0,nop,wscale 7], length 0
04:12:59.228341 IP (tos 0x0, ttl 62, id 0, offset 0, flags [DF], proto TCP (6), length 40)
    192.168.3.1.http > 192.168.1.1.59756: Flags [R.], cksum 0x3f96 (correct), seq 0, ack 1702402994, win 0, length 0
```

