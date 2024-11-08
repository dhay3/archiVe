# IPtables-extensions recent Module

ref

https://www.cnblogs.com/kevingrace/p/10008487.html

## Digest

IPtables 中提供了一个 recent match module，用于动态的匹配报文(一段时间或者指定个数)。通过 recent 就可以实现一些特殊的防火墙规则（DDos）

## List

recent module 核心是一张动态的表，用于记录通过`--set`添加的 IP -- 被称为 list，对应`/proc/net/xt_recnet/<listname>`文件

可以通过

`echo +addr >/proc/net/xt_recent/<listname>` 手动的添加 IP 到 list 中

`echo -addr >/proc/net/xt_recent/<listname>` 手动从 list 中删除 IP

`echo / >/proc/net/xt_recent/<listname>` 手动的清空 list

例如下面就是一个 list 文件

```
[root@labos-1 /]# cat /proc/net/xt_recent/badguy 
src=192.168.1.1 ttl: 62 last_seen: 4298558272 oldest_pkt: 2 4298558022, 4298558272
src=192.168.3.1 ttl: 62 last_seen: 4298543229 oldest_pkt: 6 4298541976, 4298542226, 4298542478, 4298542728, 4298542980, 4298543229
```

其中 

- src 表示 IP 地址(并不特指源地址)
- last_seen 报文出现的最后时间戳，oldest_pkt 后面的数组对应所有记录报文的时间戳
- oldest_pkt 表示报文记录的次数

此外 list 有一些默认限制

- `ip_list_tot=100`

  list 最大能存储 100 条不同 IP 的记录

- `ip_pkt_list_tot`

  list 中同 IP 最大能存储 20 条记录

## Optional args

> 注意`--set | --rcheck | --update | --remove` 是互斥的，同时只能使用一个
>

- `--name`

  the name of list. If no name is given then DEFAULT will be used

- `[!] --set`

  add the address of the packet to the list. If the  address is already in the list, this will update the existing entry（==include last_seen,oldest_pkt==）. This will always return success

- `--rsource`

  指定 `--set` 保存源地址到 list，缺省值

- `--rdest`

  指定 `--set` 保存目的地址到 list

- `--mask netmask`

  netmask that will be applied to this  recent list

  默认 32 位

- `[!] --rcheck`

  check if the address of the packet is currently in the list

- `[!] --update`

  like `--rcheck`, except it will update the “last_seen” timestamp if it matches

  即使会自动更新 last_seen, 也需要和`--set`规则一起使用

- `[!] --remove`

  check if the source address of the packet is currently in the list and if so that address will be removed from the list and the rule will return true

- `--seconds seconds`

  this option must be used in conjunction with one of `--rcheck` or `--update`. When used, this will narrrow the match to only happen when the address is in the list and was seen within the last given number of seconds

  匹配从 last_seen + seconds 之内，如果当前计算机时间戳超过 last_seen + seconds 之后重新计时

- `--hitcount hits`

  this option must be used in conjunction with one of `--rcheck` or `--update`. When used, this will narrow the match to only happen when the address is in the list and packets had been received greater than or equal to the given value
  
  oldest_pkt 值达到 hits 之后匹配

- `--reap`

  this option can only be used in conjunction with `--seconds`. when used, this will cause entries older than the last given number of seconds to be purged

## Examples

> 因为`--set` 会更新 list ，所以需要注意规则的先后顺序

- 报文数达到 n 后，丢弃。设置下面两条规则后，会出现前 2 个报文正常，第 3 个报文及之后的报文被丢弃的现象

  ```
  #将源地址加入到 badguy 中
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --hitcount 3 -j DROP
  ```

- 第一个报文正常，从第一个报文开始后 5 秒内的报文会被丢弃；然后后面一个的报文正常，从后面一个的报文后 n 秒内的报文都会丢出；以此递归。（实际是从 从 last_seen 开始计算后面 n 秒内的报文都会被丢弃）

  ```
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --seconds 5 -j DROP
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  ```

  交换一下两条规则

  ```
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --seconds 5 -j DROP
  ```

  就出现直接丢包，因为匹配第一条规则的时候 last_seen 一直在更新，匹配到第二条规则时所以一直都在 5 秒内，所以直接丢包

- 从 last_seen 之后 5 秒内 2 个报文正常，第 3 个报文(包括第 3 个报文)之后直接丢包。如果一直发包会出现前 2 个报文正常，之后丢包。因为先匹配第一条规则 last_seen 会一直更新。如果中间停止发包，会在 last_seen + 5 后重新递归

  ```
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --seconds 5 --hitcount 3 -j DROP
  ```

  交换一下两条规则

  ```
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --seconds 5 --hitcount 3 -j DROP
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  ```

  如果以 1 秒间隔发包会出现前 3 个报文正常后面 2 个报文丢弃，因为前 3 个报文没有匹配到 hitcount，后面 2 两个报文匹配到第 1 条规则时满足在 last_seen + 5 内且 hitcount 到达 3 次，因为`--rcheck` 不会更新 last_seen，所以当前时间戳大于 last_seen + 5 之后以此递归

- 第 2 条规则使用`--update` 匹配时会自动更新 last_seen， 因为 `--set` 规则在之前，所以效果上和`--rcheck` 一样

  ```
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  iptables -t filter -A INPUT -m recent --name badguy --update --seconds 5 --hitcount 3 -j DROP
  ```

  交换一下两条规则，会出现前 3 个报文正常，之后直接丢包因为 last_seen 一直更新在 5 sec 内且 oldest_pkt 大于 3 。之后以此递归
  
  ```
  iptables -t filter -A INPUT -m recent --name badguy --update --seconds 5 --hitcount 3 -j DROP
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  ```
  
  