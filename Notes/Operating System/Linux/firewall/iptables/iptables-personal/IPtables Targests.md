# IPtables Targests

## Digest

当报文匹配一条 rule 时，可能会执行一系类的 action，而 action 在 iptables 中也被称为 targets 

targets 的值可以是 

1. user-defined chains
2. built-in targets( ACCEPT, DROP, RETURN, etc) 

### built-in targets

1. ACCEPT

2. DROP

3. RETURN

   直接理解成编程语言中的 return 关键字

### non-terminated targets VS terminated targets

同时 targets 也被分为 non-terminated targets 和 terminiated targets

**non-terminated targets**

在匹配 rules 是如果 target 不会导致 rules 匹配停止的，这些 target 被称为 non-terminated targets

常见的 non-terminated targets 有：

LOG, MARK, etc

**terminated targets**

在匹配 rules 是如果 target 会导致 rules 匹配停止的，这些 target 被称为 terminated targets

大多数的 targets 都是 terminated targets, 常见的有：

ACCEPT, DROP, REJECT, RETURN, SNAT, DNAT, etc

### extension targets

具体查看 `iptables-extensions`

## None Target

在 iptables rules 中 target 不是必须的，可以是空值。如果为空值，iptables 只会记录 counter 值不会执行动作(可以把 None Target 直接理解成 non-terminated target)

例如

192.168.1.1 访问 192.168.3.1

192.168.3.1 配置规则

```
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.1.1    
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.1.1 -j ACCEPT
#这里为了排除 policy 的影响将规则设置成丢包
[root@labos-2 /]# iptables -P INPUT DROP
```

然后从 192.168.1.1 ping 192.168.3.1

这里可以发现是正常回包的

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
64 bytes from 192.168.3.1: icmp_seq=1 ttl=62 time=37.8 ms
64 bytes from 192.168.3.1: icmp_seq=2 ttl=62 time=37.8 ms
64 bytes from 192.168.3.1: icmp_seq=3 ttl=62 time=38.3 ms
^C
--- 192.168.3.1 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2003ms
rtt min/avg/max/mdev = 37.775/37.952/38.303/0.248 ms
```

查看 192.168.3.1 计数

```
[root@labos-2 /]# iptables -nvL
Chain INPUT (policy DROP 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    3   252            all  --  *      *       192.168.1.1          0.0.0.0/0           
    3   252 ACCEPT     all  --  *      *       192.168.1.1          0.0.0.0/0
```

