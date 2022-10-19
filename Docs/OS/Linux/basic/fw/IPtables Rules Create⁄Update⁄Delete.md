# IPtables Rules Create/Update/Delete

ref

https://www.zsythink.net/archives/1517

> CUD 都包含 rule-specification ( Matches args + Targets arg )

## Create

如果需要 create IPtables rules 主要使用下面几个参数

- `-t | --table [table]`

  使用指定 table，如果没有指定 table 默认使用 filter

- `-A | --append chain rule-specification`

  append one or more rules to the end of the selected chain

- `-I | --insert chain [rulenum] rule-specification`

  insert one or more rules in the selected chain as the given rule number

  在指定 rulenum 前插入 rule-sepcification, 如果没有指定 rulenum 默认为 1

```
#头插
[root@localhost /]# iptables -t filter -I INPUT -s 192.168.1.146 -j DROP
#尾插
[root@localhost /]# iptables -t filter -A INPUT -s 192.168.1.150 -j DROP
[root@localhost /]# iptables -t filter -nvL INPUT
Chain INPUT (policy ACCEPT 103 packets, 7384 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       192.168.1.146        0.0.0.0/0           
    0     0 DROP       all  --  *      *       192.168.1.150        0.0.0.0/0
    
#通过 rulenum 插入，只有头插没有尾插，从 1 开始计算
[root@localhost /]# iptables -t filter -I INPUT 2 -s 192.168.1.150 -j REJECT
[root@localhost /]# iptables -t filter -nvL INPUT
Chain INPUT (policy ACCEPT 240 packets, 17296 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       192.168.1.146        0.0.0.0/0           
    0     0 REJECT     all  --  *      *       192.168.1.150        0.0.0.0/0            reject-with icmp-port-unreachable
    0     0 DROP       all  --  *      *       192.168.1.150        0.0.0.0/0 
```

## Update

> 此处有坑！！！
>
> 这里需要注意的是如果没有指定`-s`或者`-d` 指定源目时，默认会使用 anywhere 即匹配 `0.0.0.0/0`, IPv6 同理 

如果需要 update IPtables rules 主要使用下面几个参数

- `-R | --replace chain rule-specification | rulenum`

  replace a rule in the selected chain

  支持两种格式 rulenum 和 rule-specification

```
[root@localhost /]# iptables -t filter -R INPUT 1 -s 192.168.1.150 -j DROP
[root@localhost /]# iptables -t filter -nvL INPUT
Chain INPUT (policy ACCEPT 86 packets, 6208 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       192.168.1.150        0.0.0.0/0
```

## Delete

如果需要 delte IPtables rules 主要使用下面几个参数

- `-t | --table [table]`

  使用指定 table，如果没有指定 table 默认使用 filter

- `-F | --flush [chain]`

  清空指定 chain’s rules , 如果没有指定默认清空指定 table 所有 chain’s rules 

- `-D | --delete chain rule-specification | rulenum`

  delete one or more rules from the selected chain

  支持两种格式 rulenum 和 rule-specification

```
#flush
cpl in ~ λ sudo iptables -t filter -F OUTPUT 
cpl in ~ λ sudo iptables -t filter --line -nL OUTPUT
Chain OUTPUT (policy ACCEPT)
num  target     prot opt source               destination
#delete
[root@localhost /]# iptables -t filter -D INPUT -s 192.168.1.150 -j DROP
#delete by rulenum
#通过这种方式删除的好处是不需要指定 rule-specification
[root@localhost /]# iptables -t filter -D INPUT 2
```