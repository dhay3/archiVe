# IPtables-extensions conntrack/state Module

## Conntrack

### Digest

> 和防火墙上的 session 类似

IPtables 中提供了conntrack 模块用于跟踪连的状态，==这里连接不仅仅只面向连接的协议(TCP)，也适用于不面向连接的协议(ICMP, UDP)==

### Optional args

- `[!] --ctstate statelist`

  statelist 可以一个值也可以是一个数组。具体可以是 

  - INVALID

    The packet is associated with no known connection

    在 TCP 中通常会匹配未发送 SYN 报文，直接发送 SYN-ACK 的畸形连接 


  - NEW

    The packet has started a new connection or otherwise  associated with a connection which has not seen packets in both directions

    在 TCP 中不仅限于匹配 SYN 报文。如果未发送 SYN 报文，直接发送 ACK 报文也会匹配 


  - ESTABLISHED

    The packet is associated with a  connection  which  has  seen packets in both directions.

    在 TCP 中的表现就是连接建立之后的报文(不仅仅是包含 ACK 标志位的报文)


  - RELATED

    The  packet  is  starting a new connection, but is associated with an existing connection, such as an FTP data transfer  or an ICMP error.

    因为 ICMP error 通常都是目的发送给源的 


  - UNTRACKED

    The  packet  is  not tracked at all, which happens if you explicitly untrack it by using -j CT --notrack in the  raw  table


  - SNAT

    A virtual state, matching if the original source address differs from the reply destination

    从目的到本机的报文的即说明源IP做了 SNAT


  - DNAT

    A virtual state, matching if the original destination differs  from the reply source

    即说明目IP做了 DNAT

- `[!] --ctproto l4proto`

  匹配 layer-4 协议。例如 tcp, udp 等

## state

### Digest

state 是 conntrack 的一个子集, 不涵盖 NAT 的场景

### Optional args

- `[!] --state state`

  where state is a comma separated list of the connection states to match. Only a subset of the states unterstood by “conntrack” are recognized: INVALID, ESTABLISHED, NEW, RELATED or UNTRACKED

## Exmaple

- 丢弃从 192.168.3.1 来的所有新建报文。如果构造 ACK 连接同样会被丢弃

  ```
  iptables -t filter -A INPUT -s 192.168.1.1 -m conntrack --ctstate NEW -j DROP
  #等价
  iptables -t filter -A INPUT -s 192.168.1.1 -m state --ctstate NEW -j DROP
  ```

- 丢弃 SNAT/DNAT 的报文

  ```
  
  ```

  