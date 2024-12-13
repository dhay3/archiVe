# IPtables Terms

ref

https://wiki.archlinux.org/title/iptables

https://www.zsythink.net/archives/1199

https://arthurchiao.art/blog/deep-dive-into-iptables-and-netfilter-arch-zh/

理解 IPtables 的核心在于下面这张图表，为了方便记忆 security table 被移除（实际情况下很少会使用 security table）

![Snipaste_2020-10-13_23-24-43](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-10-13_23-24-43.1q1a0cer9hvk.png)

 任何数据报文都会按照上述图表的顺序

> tables 和 chains 大小写敏感

## Tables

iptables 有下面 5 张表，每张表都是逻辑上功能类似的 rules 集合

1. raw

   只用于关闭 connection tracking for packets

2. filter

   缺省 table, 负责过滤包

3. nat

   用于 NAT 或者是 NPAT

4. mangle

   用于修改包 IP 层报文头

5. security

   管理 Mandotory Access Control, 在 linux 的具体实现是 SELinux

## Chains

每张表格都有不同的 chains，每一条 chain 由多条 rules 组成

为了方便下面用一张表格展示，其中 Y 标识这个 tables 里含有这个 chian。

另外为了展示每条 chain 在 table 中执行的优先级，这边将 nat 细分成了 SNAT 和 DNAT (DNAT 执行的规则总是优先与 SNAT)

这里可以思考一下为什么 DNAT 总是优先于 SNAT 的。==因为当一台主机收到报文时假设目的不是自己就会转发报文，如果是就会进入系统栈中处理，这时和源并没有关系。==

| Tables/Chains  | PREROUTING | INPUT | FORWARD | OUTPUT | POSTROUTING |
| :------------- | :--------- | :---- | :------ | :----- | :---------- |
| (路由判断)     |            |       |         |        |             |
| **raw**        | Y          |       |         | Y      |             |
| (连接跟踪）    |            |       |         |        |             |
| **mangle**     | Y          | Y     | Y       | Y      | Y           |
| **nat (DNAT)** | Y          |       |         | Y      |             |
| (路由判断)     |            |       |         |        |             |
| **filter**     |            | Y     | Y       | Y      |             |
| **security**   |            | Y     | Y       | Y      |             |
| **nat (SNAT)** |            | Y     |         | Y      | Y           |

==包的执行顺序从左至右，从上到下。但是根据场景不同，包经过的 chains 也不同===

1. 收到的包目的是本机

   PREROUTING -> INPUT -> OUTPUT -> POSTROUTING

2. 收到的包目的不是本机

   PREROUTING -> INPUT -> FORWARD -> POSTROUTING

3. 本地发出的包

   OUTPUT -> POSTROUTING

## Rules

每张 table 的每条 chain 都有一组 rules, 当 chain 被调用时会依次读取 chain 里面的 rules 按照下面的伪代码执行

```
rules = [...]

match_rule(pkt,0)

func match_rule(pkt,rule_idx)
if pkt match rules[rule_idx].match && rule[rule_idx].target is terminated target
	return
else
	return match_rule(pkt,rule_idx++)   
```

每条 rule 按照在 table 中的先后顺序匹配

### Matches

> 具体可以参考 iptables-extensions

必须满足的条件, 条件可以是

1. protocol type
2. destination/source address
3. destination/source port
4. destination/source network
5. input/output interface
6. headers
7. connection state
8. etc.

### Targets

> 具体可以参考 iptables-extensions

满足条件后执行的动作 ( action ) 叫做目标 ( target )

## User-defined chains

iptables 不仅有 built-in chains, 同时还支持使用 user-defined chains

和 built-in chains 不同的是 user-defined chains 只能被用做为 targets

## Policy

iptables 中的每一张 table 每一条 chain 都有一条默认的规则用于匹配，没有匹配到任何具名(手动设置) rules 的报文，即 fallback 规则

```
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination           
```

上述中的 policy ACCEPT 表示如果当前报文不匹配 INPUT chian 中具名的 rules 时，就指定当前报文为 ACCEPT target 

chain’s Policy 可以通过 `-P` 来修改，例如

```
iptables -t filter -P INPUT DROP
```

## Modules

> 具体可以参考 iptables-extensions

iptables 按照场景还提供了许多 modules，例如 connlimit, conntrack, limit, recent 等等

