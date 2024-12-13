# Linux ip-rule

## 0x1 Digest

管理 policy routing（策略路由类似于ACL） 

传统的路由选路规则是按照目的IP来选择路由，但是现在需要根据对IP数据包中的其他字段(比如端口、源IP、Tos)来实现自定义选路规则。实现该功能的同时结合传统路由规则产生了 routing policy database（RPDB）

每条 policy routing 由 selector 和 action predicate 组成，当匹配到 selector 时会执行 action predicate

## 0x2 Terms

### Preference

优先级，数值越大优先级越低

### RPDB

在 kernel 启动过程中会初始会 RPDB，其中包含 3 条规则，优先级从上到下

1. priority：0

   selector：match anything

   action predicate：lookup up routing table local

   匹配所有会先去查询local table（`ip route show table local`）

2. priority 32766

   selector：match anything

   action：lookup routing table main

   查询 main table，通常是管理员手动添加的路由

3. priority 32767

   selector：match anything

   action：lookup routing table default

   查询 default table （空表）

### type

RPDB中可以设定的rule type

- `unicast`

  直接从路由表中找

- `blackhole`

  执行把包丢弃

- `unreachable`

  generate a ‘network is unreachable’ error

- `prohibit`

  generate a ‘communication is adminstratively prohibit’ error

- `nat`

  translate the source address of the IP packet into some other value

## 0x3 EBNF

```
ip [ OPTIONS ] rule { COMMAND | help }

ip rule [ list [ SELECTOR ]]

ip rule { add | del } SELECTOR ACTION

ip rule { flush | save | restore }

SELECTOR := [ not ] [ from PREFIX ] [ to PREFIX ] [ tos TOS ] [ fw‐
mark FWMARK[/MASK] ] [ iif STRING ] [ oif STRING ] [ pref
NUMBER ] [ l3mdev ] [ uidrange NUMBER-NUMBER ] [ ipproto
PROTOCOL ] [ sport [ NUMBER | NUMBER-NUMBER ] ] [ dport [
NUMBER | NUMBER-NUMBER ] ] [ tun_id TUN_ID ]

ACTION := [ table TABLE_ID ] [ protocol PROTO ] [ nat ADDRESS ] [
realms [SRCREALM/]DSTREALM ] [ goto NUMBER ] SUPPRESSOR

SUPPRESSOR := [ suppress_prefixlength NUMBER ] [ suppress_ifgroup
GROUP ]

TABLE_ID := [ local | main | default | NUMBER ]
```

## 0x4 Commands

### ip rule list

syntax：`ip rule [ list [SELECTOR]]`

show ruting policy

### ip rule add

syntax: `ip rule { add | del } SELECTOR ACTION`

insert a new rule

### ip rule delete

syntax：`ip rule { add | del } SELECTOR ACTION`

delete a rule

### ip rule flush

syntax：`ip rule { flush | save | restore }`

清空ip rule，谨慎使用

### ip rule save

syntax：`ip rule { flush | save | restore }`

behaves like `ip rule restore`

### ip rule restore

syntax：`ip rule { flush | save | restore }`

## 0x5 Args

### selector

- from PREFIX

  select the source prefix to match

- to

  select the destination prefix to match

- tos

  select the tos value to match

- fwmark

- iif

  select the incoming device to match

- oif

  select the outgoing device to match

- pref

  the priority of this rule

- sport

- dport

  







