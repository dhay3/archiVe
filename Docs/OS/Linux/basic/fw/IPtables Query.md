# IPtables Read Rules  

ref

https://www.zsythink.net/archives/1493

如果需要查询 IPtables rules 主要使用下面几个参数

- `-t | --table [table]`

  查看指定 table，如果没有指定 table 默认查看 filter

- `-L | --list [chain]`

  查看指定 chain rules，如果没有指定 chain 默认查看所有 chain

- `-n`

  numberic ouput

  默认以字符串的形式显示

- `-v`

  verbose ouput

  会显示额外显示 pkts, bytes, in, out 字段

- `--line-numbers`

  在 rules 前显示 num 序列号

以添加了如下规则为例子

```
cpl in ~/ λ sudo iptables -t filter -S OUTPUT
-P OUTPUT ACCEPT
-A OUTPUT -d 39.156.66.10/32 -j REJECT --reject-with icmp-port-unreachable
```

## Header

```
cpl in ~ λ sudo iptables -t filter --line -nvL OUTPUT
Chain OUTPUT (policy ACCEPT 118K packets, 11M bytes)
num   pkts bytes target     prot opt in     out     source               destination         
1        0     0 REJECT     all  --  *      *       0.0.0.0/0            39.156.66.10         reject-with icmp-port-unreachable
```

- policy

  当前 chain 的默认 Target, 当不匹配 chain 中的所有 rules 时，使用 policy 

  如果值为 ACCEPT 即为 黑名单 机制

  如果值为 DROP 即为 白名单 机制

- packets

  当前 chain 匹配的所有包数量，需要和 `-v` 一起使用时才会出现

- bytes

  当前 chain匹配所有包的大小总和，需要和 `-v` 一起使用时才会出现

## Columns

### default

```
cpl in ~ λ sudo iptables -t filter -L OUTPUT
Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         
REJECT     all  --  anywhere             39.156.66.10         reject-with icmp-port-unreachable
```

- target

- proto

  需要匹配的协议

- opt

  target 额外的参数选项

- source

  需要匹配的源地址

- destination

  需要匹配的目的地址

### verbose

使用`-v` 额外显示字段

```
cpl in ~ λ sudo iptables -t filter -vL OUTPUT
Chain OUTPUT (policy ACCEPT 84473 packets, 7590K bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 REJECT     all  --  any    any     anywhere             39.156.66.10         reject-with icmp-port-unreachable
```

- pkts

  rule 匹配的报文数

- bytes

  rule 匹配报文的大小总和

- in

  从那个网卡来

- out

  从那个网卡出

其中 pkts 和 bytes 也被称为 counters

