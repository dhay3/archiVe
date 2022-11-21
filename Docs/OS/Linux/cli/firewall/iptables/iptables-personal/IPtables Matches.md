# IPtables Matches

## Digest

Matches 是报文匹配 rule 必须满足的条件，如果满足条件就会执行 target，如果不满足就会匹配后面的 rules

伪代码如下

```
rules = [...]

match_rule(pkt,0)

func match_rule(pkt,rule_idx)
if pkt match rules[rule_idx].match && rule[rule_idx].target is terminated target
	return
else
	return match_rule(pkt,rule_idx++) 
```

Matches 可以按照类型一般可以是

1. protocol type
2. destination/source address
3. destination/source port
4. destination/source network
5. input/output interface
6. headers
7. connection state
8. etc.

## built-in matches

iptables 提供了一些参数用于匹配默认 built-in matches （即无需指定 extensions matches）

- `-s `

  匹配报文源 IP 或者是 IP 段

- `-d`

  匹配报文目的 IP 或者是 IP 段

- `-p`

  匹配协议，如果指定了协议，会默认加载协议对应的 extension matches 参数

- `-i | -o`

  匹配 in 或者 out interface

## extension matches

具体参考 iptables-extensions

## Cautions

### multiple matches

在 IPtables CLI 大部分 Matches args 都是可变参数，例如 `-s` 和 `-d` ，可以同时指定多个源目(中间以逗号分割)来设置多条 rules

```
[root@localhost vagrant]# iptables -t filter -I INPUT -s 192.168.1.146,192.168.1.150 -j DROP
[root@localhost vagrant]# iptables -nvL INPUTChain INPUT (policy ACCEPT 45 packets, 3208 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      *       192.168.1.150        0.0.0.0/0           
    0     0 DROP       all  --  *      *       192.168.1.146        0.0.0.0/0
```

### negate

iptables 还支持取反(negate), 需要注意的是`-s` 和 `!` 中间必须要有一个空格且只能在`-s` 之前

```
[root@localhost /]# iptables -t filter -A INPUT ! -s 192.168.1.150 -j ACCEPT
[root@localhost /]# iptables -nvL INPUT
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
  136  9784 ACCEPT     all  --  *      *      !192.168.1.150        0.0.0.0/0 
```