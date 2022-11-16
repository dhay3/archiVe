# IPtables CLI

## EBNF

> 部分参数有默认值，所以在组合参数时不支持前后倒置，例如 
>
> `iptables -Lnv` 是错误的
>
> `iptables -nvL` 是正确的

```
iptables [-t table] {-A|-C|-D|-V} chain rule-specification

ip6tables [-t table] {-A|-C|-D|-V} chain rule-specification

iptables [-t table] -I chain [rulenum] rule-specification

iptables [-t table] -R chain rulenum rule-specification

iptables [-t table] -D chain rulenum

iptables [-t table] -S [chain [rulenum]]

iptables [-t table] {-F|-L|-Z} [chain [rulenum]] [options...]

iptables [-t table] -N chain

iptables [-t table] -X [chain]

iptables [-t table] -P chain target

iptables [-t table] -E old-chain-name new-chain-name

rule-specification = [matches...] [target]

match = -m matchname [per-match-options]

target = -j targetname [per-target-options]

```

## Optional args

### Tables args

- `-t | --table [table]`

  matching table which the command should operate on

  如果没有指定`-t`参数默认读取 filter table，值可以是

  1. raw
  2. filter
  3. nat
  4. mangle
  5. security

### Chinas args

下面这些参数不能组合使用

#### Read args

- `-L | --list [chian]`

  list all rules in the selected chain, if no chain is selected, all chains are listed

- `-S | --list-rules [chain]`

  和 `-L` 一样，但是以 iptables 命令行的格式输出 rules

  ```
  cpl in ~ λ sudo iptables -S   
  
  -A FORWARD -o br-73775b359618 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
  ```

- `-C | --check chain rule-specification`

  check whether a rule matching the specification does exist in the selected chain

#### Create/Update/Delete args

- `-A | --append chain rule-specification`

  append one or more rules to the end of the selected chain

- `-I | --insert chain [rulenum] rule-specification`

  insert one or more rules in the selected chain as the given rule number

  在指定 rulenum 前插入 rule-sepcification, 如果没有指定 rulenum 默认为 1

- `-D | --delete chain rule-specification | rulenum`

  delete one or more rules from the selected chain

  支持两种格式 rulenum 和 rule-specification

- `-F | --flush [china]`

  flush the selected chain, all the chians in the table if none is given

  等价于清空 table

  ==这里需要注意的一定是，手动设置的 Policy 并不会被清除==

- `-R | --replace chain rule-specification | rulenum`

  replace a rule in the selected chain

  支持两种格式 rulenum 和 rule-specification

- `-N | --new-chain chain`

  create a new user-defined chain by the given name

- `-E | --rename-chain old-chain new-chain`

  rename the user specified chain to the user supplied name

- `-X | --delete-chain [chain]`

  delete the empty chain which specified

  一般用于删除没有 rules 的 user-defined chains，且不能删除 built-in chains (INPUT, OUTOUT, FOWARD...)

  empty built-in chains 只能通过 `iptables-nft` 删除 

- `-P | --policy chain target`

  set the policy for the built-in chain to the given target

  修改 chain 缺省 target（只能是 ACCEPT 或 DROP），==user-defined chain 不支持==

### Matches args

- `-p | --protocol protocol`

  the protocol of the rule or of the packet to check

  匹配的协议可以是数字也可以是字符串，具体支持的 proto 参考`/etc/protocols`

  1. all 代表所有 proto
  2. 支持`!`( 取反 )

- `-s | --source address[/mask][...]`

  source specification. Address can be either a network name, a hostname, a network IP addrss with `/mask`, or a plain IP address

  可以被使用多次，在对应的 CURD args 中会被扩展成多条 rules

  1. 支持`!`( 取反 )

- `-d | --destination address[/mask][...]`

  和 `-s` 一样，但是匹配 destination

- `-i | --in-interfaace name`

  name of an interface via which a packet was received

  1. 支持`!`( 取反 )
  2. 可以使用`+`来做简单的 wildcard 匹配

- `-o | --out-interface name`

  和`-i` 一样，但是匹配 out interface

### Targets args

- `-j | --jump`

  set the target of the rule

  1. 可以是 userd-defined chain
  2. 如果没有指定该参数同时也没有指定`-g`，==报文同样会被匹配，但是不会做 return 的逻辑而是会匹配下面的 rules(这种情况一般只用在 recent moudle --set 用做计数)==

### Other args

- `-m | --match match`

  指定使用的模块，使用了特定模块就可以使用一些额外的参数，具体参考 `iptables-extensions`。可以使用多个 `-m` 来匹配多个模块

- `-v | --verbose`

  可以被重复使用表示详细程度

- `-n | numeric`

  numberic output

- `--line-numbers`

  显示的时候在规则前面加一列序号, 也可以简写成`--line`

## Cautions

1. Matches args + Targets args 组合是 rule-specification 

2. 在 IPtables CLI 大部分 Matches args 都是可变参数，例如 `-s` 和 `-d` ，可以同时指定多个源目(中间以逗号分割)来设置多条 rules

   ```
   [root@localhost vagrant]# iptables -t filter -I INPUT -s 192.168.1.146,192.168.1.150 -j DROP
   [root@localhost vagrant]# iptables -nvL INPUTChain INPUT (policy ACCEPT 45 packets, 3208 bytes)
    pkts bytes target     prot opt in     out     source               destination         
       0     0 DROP       all  --  *      *       192.168.1.150        0.0.0.0/0           
       0     0 DROP       all  --  *      *       192.168.1.146        0.0.0.0/0
   ```

3. iptables 还支持取反(negate), 需要注意的是`-s` 和 `!` 中间必须要有一个空格且只能在`-s` 之前

   ```
   [root@localhost /]# iptables -t filter -A INPUT ! -s 192.168.1.150 -j ACCEPT
   [root@localhost /]# iptables -nvL INPUT
   Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
    pkts bytes target     prot opt in     out     source               destination         
     136  9784 ACCEPT     all  --  *      *      !192.168.1.150        0.0.0.0/0 
   ```

   