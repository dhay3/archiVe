# IPtables Return Target

ref

https://www.linuxtopia.org/Linux_Firewall_iptables/x4604.html



RETURN target 是 IPtables built-in Targets 之一，且是一个 terminated Target

同时 RETURN target 遵循下面几条规则

1. 当报文匹配 RETURN rules 时会停止匹配下面的规则, 然后使用 chian 默认 policy
2. 当匹配 father chain 中某一条 subchain RETURN rule 时，会停止匹配 subchain 后面的 rules 跳转到 father chain 然后继续匹配 father chain 后面的 rules

## Examples

### Lab01

例如设置 192.168.1.1 访问 192.168.3.1 不通，但是 192.168.4.1 可以访问 192.168.3.1

![2022-11-15_15-14](https://github.com/dhay3/image-repo/raw/master/20221115/2022-11-15_15-14.56e538jrk8w0.webp)

设置 192.168.3.1 规则

```
#设置 192.168.3.1 INPUT defualt DROP policy
[root@labos-2 /]# iptables -t filter -P INPUT DROP 
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.1.1 -j RETURN
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.4.1 -j ACCEPT
```

192.168.1.1 ping 192.168.3.1

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
^C
--- 192.168.3.1 ping statistics ---
1 packets transmitted, 0 received, 100% packet loss, time 0ms
```

192.168.4.1 ping 192.168.3.1

```
[root@labos-3 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
64 bytes from 192.168.3.1: icmp_seq=1 ttl=63 time=31.4 ms
^C
--- 192.168.3.1 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 31.440/31.440/31.440/0.000 ms
```

192.168.3.1 iptables counters

```
[root@labos-2 /]# iptables -t filter -nvL INPUT
Chain INPUT (policy DROP 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    1    84 RETURN     all  --  *      *       192.168.1.1          0.0.0.0/0           
    1    84 ACCEPT     all  --  *      *       192.168.4.1          0.0.0.0/0
```

### Lab02

以 192.168.1.1 访问 192.168.3.1 回送 icmp port unreachable 为例子

设置 192.168.3.1 规则

```
[root@labos-2 /]# iptables -N CUS
[root@labos-2 /]# iptables -A CUS -s 192.168.1.1 -j RETURN
[root@labos-2 /]# iptables -A CUS -s 192.168.1.1 -j DROP
[root@labos-2 /]# iptables -A INPUT -j CUS
[root@labos-2 /]# iptables -A INPUT -s 192.168.1.1 -j REJECT
```

让我们看一下规则

```
[root@labos-2 /]# iptables -nvL
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 CUS        all  --  *      *       0.0.0.0/0            0.0.0.0/0           
    0     0 REJECT     all  --  *      *       192.168.1.1          0.0.0.0/0            reject-with icmp-port-unreachable

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain CUS (1 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 RETURN     all  --  *      *       192.168.1.1          0.0.0.0/0           
    0     0 DROP       all  --  *      *       192.168.1.1          0.0.0.0/0 
```

从 192.168.1.1 ping 192.168.3.1

这是可以看到回送的报文的 icmp port unreachable

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
From 192.168.3.1 icmp_seq=1 Destination Port Unreachable
```

看一下眼 counters

```
[root@labos-2 /]# iptables -nvL
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    2   168 CUS        all  --  *      *       0.0.0.0/0            0.0.0.0/0           
    2   168 REJECT     all  --  *      *       192.168.1.1          0.0.0.0/0            reject-with icmp-port-unreachable

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain CUS (1 references)
 pkts bytes target     prot opt in     out     source               destination         
    2   168 RETURN     all  --  *      *       192.168.1.1          0.0.0.0/0           
    0     0 DROP       all  --  *      *       192.168.1.1          0.0.0.0/0
```

同时匹配了 INPUT 中的两条规则
