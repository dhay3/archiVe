# IPtables User-defined Chains

https://www.zsythink.net/archives/1625

https://sleeplessbeastie.eu/2018/06/21/how-to-create-iptables-firewall-using-custom-chains/

## Digest

如果我们需要定义一组功能类似的 rules，例如过滤 http 报文。如果这些 rules 都加入到 built-in chains 中，这样就会比较混乱且不利于管理。这是就可以通过创建 user-defined chains 来管理这些 rules 以表示一个集合

同时 user-defined chain 需要遵循如下规则

1. user-defined chain 只能被用做 traget 挂载 built-in chains 中
2. 如果需要删除 user-defined chain，必须先要清空 user-defined chain 然后 delink built-in chain
3. user-defined chain 不支持设置 policy
4. 如果 user-defined chain 中定义了 return target rule, 在匹配的时候会停止匹配 user-defined chain 后面的 rules，但是会继续匹配 applied built-in chain 中的 rules

##  Add user-defined chain

我们可以通过 `-N` 参数来实现添加 user-defined chain

```
[root@labos-1 /]# iptables -N CUZ
```

这样我们就可以看到已经添加了一条 CUS user-defined chain

```
[root@labos-1 /]# iptables -nvL
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain CUZ (0 references)
 pkts bytes target     prot opt in     out     source               destination         
```

## Rename user-defined chain

> 这里需要注意的是，rename 不需要清空 rules 或者 delink

如果 user-defined  chain 名字错了，也没有关系，可以通过 `-E` 来修改

```
[root@labos-1 /]# iptables -E CUZ CUS
```

## Apply user-defined chain

为了使用 user-defined chain 先添加一条 rule

```
[root@labos-1 /]# iptables -t filter -A CUS -d 192.168.3.1 -j REJECT
```

将 user-defined chain 挂到 built-in chain

```
[root@labos-1 /]# iptables -t filter -A OUTPUT -j CUS
```

查看 rules

这时可以看到 CUS chain 已经挂到了 OUTPUT chain 上了

```
[root@labos-1 /]# iptables -t filter -nvL            
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 CUS        all  --  *      *       0.0.0.0/0            0.0.0.0/0           

Chain CUS (1 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 REJECT     all  --  *      *       0.0.0.0/0            192.168.3.1          reject-with icmp-port-unreachable
```

## Test user-defined chain

我们来测试一下，以 192.168.1.1（即 labos-1）ping 192.168.3.1 为例

这时可以看到 192.168.1.1 回送了 destination port unreachable 符合预期

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
From 192.168.1.1 icmp_seq=1 Destination Port Unreachable
```

## Delete user-defined chain

> 可以通过直接清空 tables 来清空 user-defined chain, 但是风险需要自己评估

如果 user-defined chain 中包含 rules ，那么就不能被直接删除

```
[root@labos-1 /]# iptables -X CUS
iptables v1.8.4 (nf_tables):  CHAIN_USER_DEL failed (Device or resource busy): chain CUS
```

必须先要清空 rules，但是同样会报错

```
[root@labos-1 /]# iptables -F CUS 
[root@labos-1 /]# iptables -X CUS
iptables v1.8.4 (nf_tables):  CHAIN_USER_DEL failed (Device or resource busy): chain CUS
```

因为 CUS chain 被引入到了 built-in chain,  iptables 不支持级联删除，必须先要删除 built-in chain 中对应的 rule 后才能被删除

```
[root@labos-1 /]# iptables -D OUTPUT -j CUS   
[root@labos-1 /]# iptables -X CUS
```

