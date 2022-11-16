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

在 iptables rules 中 target 不是必须的，可以是空值。如果为空值，默认会使用 Policy 作为匹配后的动作即 traget

例如

192.168.3.1 添加一条规则匹配从 192.168.1.1 来的报文

```
[root@labos-2 /]# iptables -A INPUT -s 192.168.1.1 
```

然后从 192.168.1.1 ping 192.168.3.1

这里可以发现是正常回包的，因为默认 built-in chain  policy 均为 ACCEPT

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
64 bytes from 192.168.3.1: icmp_seq=1 ttl=62 time=40.10 ms
```

修改 192.168.3.1  INPUT policy

```
[root@labos-2 /]# iptables -P INPUT DROP
```

然后从 192.168.1.1 ping 192.168.3.1

这是可以发现丢包了，因为 192.168.3.1 INPUT 设置了 DROP policy

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
^C
--- 192.168.3.1 ping statistics ---
2 packets transmitted, 0 received, 100% packet loss, time 1026ms
```

