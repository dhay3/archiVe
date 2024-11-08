# IPtables-extensions time Module

## Digest

IPtables 中提供了一个 time 模块，用于按照报文的时间戳过滤。

## Optional args

- `--datestart YYYY[-MM[-DD[Thh[:mm[:ss]]]]]`

  `--datestop YYYY[-MM[-DD[Thh[:mm[:ss]]]]]`

  匹配 ISO 8601 的格式指定时间

- `--timestart hh:mm[:ss]`

  `--timestop hh:mm[:ss]`

  匹配 daytime, 范围在 00:00:00 - 23:59:59

- `[!]--monthdays day[,day...]`

  匹配 monthdays 可以是 1 - 31 如果一个月没有 31 就不会匹配(非匹配最后一天)

- `[!]--weekdays day[,day...]`

  匹配 weekdays, 可以是以 Mon - Sun 的格式也可以 1 - 7 的格式

## Examples

- 每周一从 9 点开始丢弃所有 INPUT 的 icmp 报文

  ```
  [root@labos-2 /]# iptables -t filter -A INPUT -p icmp -m time --weekdays 1 --timestart 9:00 -j DROP
  ```