# IPtables-extensions connlimit Module

## Digest

> 可以结合 recent 一起使用，限制连接状态并作出处罚

IPtables 中提供了一个模块用于限制从 client 来的连接，==这里连接不仅仅只面向连接的协议(TCP)，也适用于不面向连接的协议(ICMP, UDP)==

## Optional args

- `--connlimit-upto n`

  match if the number of existing connections is below or equal n

- `--connlimit-above n`

  match if the number of exsiting connections is above n

## Examples

- allow 2 telnet connections per client host

  ```
  iptables -A INPUT -p tcp --syn --dport 23 -m connlimit --connlimit-upto 2 -j ACCEP
  #等价
  iptables -A INPUT -p tcp --syn --dport 23 -m connlimit --connlimit-above 2 -j REJECT
  ```

- limit the number of parallel HTTP requests to 16 per class C sized source network (24 bit netmask)

  ```
  iptables -p tcp --syn --dport 80 -m connlimit --connlimit-above 16 --connlimit-mask 24 -j REJECT
  ```

