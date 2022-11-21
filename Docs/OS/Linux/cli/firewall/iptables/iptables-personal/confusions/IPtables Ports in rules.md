# IPtables Ports in rules

学过 TCP 大家都知道由四元组定义为唯一一条连接，同时 A 发送报文给 B，B 会回送报文给 A。这个在我们看来是两个方向的报文，那么现在有一个问题

以 A:sport 访问 B:dport 为例

如果我们在 A 上就设置了 ACCEPT OUTPUT dport 其他规则均为 DROP

 B 上没有设置规则

会出现什么现象呢？是 A 上将 B 的报文丢弃(因为 INPUT DROP)，还是正常就收报文呢？

我们以一个试验切入

## Lab

A 上设置规则

```
[root@labos-A /]# iptables -P INPUT DROP
[root@labos-A /]# iptables -P OUTPUT DROP
[root@labos-A /]# iptables -A OUTPUT -p tcp --dport 80 -j ACCEPT
```

A 访问 B:80

这里可以看到连接超时了

```
[root@labos-A /]# nc -v 192.168.3.1 80
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connection timed out.
```

A 上抓包

这里可以看到 B 是有回包的，但是 A 没有回送 ACK  报文

```
18:14:13.458029 IP 192.168.1.1.51348 > 192.168.3.1.http: Flags [S], seq 3671141429, win 64240, options [mss 1460,sackOK,TS val 4232990929 ecr 0,nop,wscale 7], length 0
18:14:13.495474 IP 192.168.3.1.http > 192.168.1.1.51348: Flags [S.], seq 3276581094, ack 3671141430, win 65160, options [mss 1460,sackOK,TS val 2267044554 ecr 4232990929,nop,wscale 7], length 0
18:14:14.477015 IP 192.168.1.1.51348 > 192.168.3.1.http: Flags [S], seq 3671141429, win 64240, options [mss 1460,sackOK,TS val 4232991948 ecr 0,nop,wscale 7], length 0
18:14:14.495058 IP 192.168.3.1.http > 192.168.1.1.51348: Flags [S.], seq 3276581094, ack 3671141430, win 65160, options [mss 1460,sackOK,TS val 2267045556 ecr 4232990929,nop,wscale 7], length 0
...
```

之后重传到达设定阈值后断开连接

## Summary

上面的实验说明 IPtables 中

1. 如果是 CS 模型，Client 只开放一个目的端口的出向规则是不能正常建立 TCP 连接的。还需要开放入向规则任意源端口或者是指定目的端口的规则才能正常建连
2. 如果是 CS 模型，Server 只开放一个服务端口的入向规则是不能正常建立 TCP 连接的。还需要开放出向规则任意目的端口或者是指定源端口为服务端口的规则才能正常建连
3. 从上面的结论看，IPtables 和 firewall 一样。但是和市面上云产品的安全组不一样(安全组只要能出就能进)