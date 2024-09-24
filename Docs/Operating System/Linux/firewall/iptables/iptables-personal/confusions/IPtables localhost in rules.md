# IPtables localhost in rules

如果我们设置了一个规则正好匹配到 localhost，这时 IPtables 会怎么处理呢？

## INPUT

设置 INPUT 规则

```
[root@labos-1 /]# iptables -A INPUT -p tcp --dport 80 -j DROP
```

访问 localhost 80

```
[root@labos-1 /]# nc -v 127.0.0.1 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connection timed out
```

查看规则

匹配规则丢包

```
[root@labos-1 /]# iptables -nvL 
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    8   480 DROP       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp dpt:80

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination        
```

## OUTPUT

设置 OUTPUT 规则

```
[root@labos-1 /]# iptables -A OUTPUT -p tcp --dport 80 -j DROP
```

访问 localhost 80

```
[root@labos-1 /]# nc -v 127.0.0.1 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connection timed out
```

查看规则

匹配规则丢包

```
[root@labos-1 /]# iptables -nvL
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    4   240 DROP       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp dpt:80
```

## Summary

Chain 只和 NIC 有关系，即使是虚拟的 loopback NIC，IPtables 同样会生效