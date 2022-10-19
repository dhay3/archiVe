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
> 

- `--name`

  the name of list. If no name is given then DEFAULT will be used

- `[!] --set`

  add the address of the packet to the list. If the  address is already in the list, this will update the existing entry（include last_seen,oldest_pkt）. This will always return success

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

  like `--rcheck`, except it will update the “last seen” timestamp if it matches

- `[!] --remove`

  check if the source address of the packet is currently in the list and if so that address will be removed from the list and the rule will return true

- `--seconds seconds`

  this option must be used in conjunction with one of `--rcheck` or `--update`. When used, this will narrrow the match to only happen when the address is in the list and was seen within the last given number of seconds

  只会在第 second 内生效

- `--hitcount hits`

  this option must be used in conjunction with one of `--rcheck` or `--update`. When used, this will narrow the match to only happen when the address is in the list and packets had been received greater than or equal to the given value

## Examples

> 因为`--set` 会更新 list ，所以需要注意规则的先后顺序

- 报文数达到 n 后，丢弃。设置下面两条规则后，会出现前 3 个报文正常，但是后面报文被丢弃的现象

  ```
  #将源地址加入到 badguy 中
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --hitcount 3 -j DROP
  ```

- 每隔第 n 秒内的报文会被丢弃。设置下面两条规则后，会出现每隔第 3 秒的报文都会被丢弃

  ```
  iptables -t filter -A INPUT -m recent --name badguy --rcheck --seconds 3 -j DROP
  
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  ```

- 11

  ```
  iptables -t filter -A INPUT -m recent --name badguy --update --seconds 10 -j DROP
  iptables -t filter -A INPUT -m recent --name badguy --set --rsource
  ```

  