# IPtables Matches

ref

https://www.zsythink.net/archives/1544

## Multiple rules

在 IPtables CLI 中 `-s` 和 `-d` 的是可变参数，可以同时指定多个源目(中间以逗号分割)来设置多条 rules

```
[root@localhost vagrant]# iptables -t filter -I INPUT -s 192.168.1.146,192.168.1.150 -j DROP
[root@localhost vagrant]# iptables -nvL INPUTChain INPUT (policy ACCEPT 45 packets, 3208 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       192.168.1.150        0.0.0.0/0           
    0     0 DROP       all  --  *      *       192.168.1.146        0.0.0.0/0
```

## Negate

iptables 还支持取反, 需要注意的是`-s` 和 `!` 中间必须要有一个空格且只能在`-s` 之前

```
[root@localhost /]# iptables -t filter -A INPUT ! -s 192.168.1.150 -j ACCEPT
[root@localhost /]# iptables -nvL INPUT
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
  136  9784 ACCEPT     all  --  *      *      !192.168.1.150        0.0.0.0/0 
```

