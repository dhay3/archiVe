# Linux ip-route

## 0x1 Digest

ip route 用于管控 kernel routing table

## 0x Terms

### Route tables

Linux 2.x 将路由分成多种 routing tables 用 1 - $2^{32}$ 或 字符来标识，默认记录在`/etc/iproute2/rt_tables`。其中`local` route table 是不可见的（包含local路由和broadcast路由）

## 0x2 ENBF

- SELECTOR := [ root PREFIX ] [ match PREFIX ] [ exact PREFIX ] [ table
  TABLE_ID ] [ vrf NAME ] [ proto RTPROTO ] [ type TYPE ] [
  scope SCOPE ]
- ROUTE := NODE_SPEC [ INFO_SPEC ]
- NODE_SPEC := [ TYPE ] PREFIX [ tos TOS ] [ table TABLE_ID ] [ proto
  RTPROTO ] [ scope SCOPE ] [ metric METRIC ] [ ttl-propagate { enabled | disabled } ]
- INFO_SPEC := NH OPTIONS FLAGS [ nexthop NH ] ...
- NH := [ encap ENCAP ] [ via [ FAMILY ] ADDRESS ] [ dev STRING ] [ weight NUMBER ] NHFLAGS

- FAMILY := [ inet | inet6 | ipx | dnet | mpls | bridge | link ]

- OPTIONS := FLAGS [ mtu NUMBER ] [ advmss NUMBER ] [ as [ to ] ADDRESS ] rtt TIME ] [ rttvar TIME ] [ reordering NUMBER ] [ window NUMBER ] [ cwnd NUMBER ] [ ssthresh REALM ] [ realms REALM ] [ rto_min TIME ] [ initcwnd NUMBER ] [ initrwnd NUMBER ] [ features FEATURES ] [ quickack BOOL ] [ congctl NAME ] [ pref PREF ] [ expires TIME ] [ fastopen_no_cookie BOOL ]

- TYPE := [ unicast | local | broadcast | multicast | throw | unreachable | prohibit | blackhole | nat ]
  TABLE_ID := [ local| main | default | all | NUMBER ]

- SCOPE := [ host | link | global | NUMBER ]

- NHFLAGS := [ onlink | pervasive ]

- RTPROTO := [ kernel | boot | static | NUMBER ]

- FEATURES := [ ecn | ]

- PREF := [ low | medium | high ]

- ENCAP := [ MPLS | IP | BPF | SEG6 | SEG6LOCAL ]

- ENCAP_MPLS := mpls [ LABEL ] [ ttl TTL ]

- ENCAP_IP := ip id TUNNEL_ID dst REMOTE_IP [ tos TOS ] [ ttl TTL ]

- ENCAP_BPF := bpf [ in PROG ] [ out PROG ] [ xmit PROG ] [ headroom SIZE ]
- ENCAP_SEG6 := seg6 mode [ encap | inline | l2encap ] segs SEGMENTS [ hmac KEYID ]

- ENCAP_SEG6LOCAL := seg6local action SEG6_ACTION [ SEG6_ACTION_PARAM ]

- ROUTE_GET_FLAGS :=  [ fibmatch  ]

## 0x3 Route types

- unicast

  单播路由

- unreachable

  these destinations are unreachable

  包会被丢弃同时会回送 ICMP host unreachable，源端收到EHOSTUNEACH erro

- blackhole

  these destinations are unreachable

  包会被discarded silently，源端收到 EINVAL erro

- prohibit

  these destinations are unreachable

  包会被丢弃同时回送ICMP communication adminstratively prohibited，源端收到 EACCES error

- local

  these destinations are assigned to this host

  包被looped back，只能在本地

- broadcast

  these destinations are broadcast addresses

  包通过广播的方式发送出去

- throw

  会假设没有路由条目，包会被丢弃回送ICMP net unreachable，源端收到 ENETUNERACH

- nat

  DNAT，由`via`字段选择规则，Linux 2.6 后不支持

- anycast

  the destinations are anycast addresses

  多播

- multicast

  the destinations are unicast addresses

  多播

## 0x04 Commands

具体字段可以使用的值查看EBNF

### ip route add

syntax：`ip route { add | del | change | append | replace } ROUTE`

add new route

### ip route change

syntax：`ip route { add | del | change | append | replace } ROUTE`

change route

### ip route replace

syntax：`ip route { add | del | change | append | replace } ROUTE`

change or add new add

- to TYPE PREFIX

  the destinatino prefix of the route

  如果没有指定TYPE(route types)，默认使用 unicast. PREFIX 可是点对点IP或是网段，如果没有指定 prefix，掩码为32位。也可以使用 defult 表示`0/0`或`::/0`

- tos TOS | dsfield TOS

  the type of service（TOS）

- metric NUMBER | preference NUMBER

  the preference value of the route

  路由的优先值，值越小优先值越大

- table TABLEID

  tehi table to add this route to

  TABLEID 使用`/etc/iproute2/rt_tables`中记录的值，如果没有指定TABLEID默认使用main，如果添加的route type 是local，broadcast，nat会自动添加到 local table

- vrf NAME

  the vrff name to add this route to

- dev NAME

  the out device name

  流量出设备名

- via [FAMLIY] ADDRESS

  the address of the nexthop router，in the address family  FAMILY

  下一跳地址

- src ADDRESS

  the source address to prefer when sending to the destinations convered by the route prefix

- realm REALMID

  the realm to which this route is assigned. REALMID 使用`/etc/iproute2/rt_realm`中记录的值

- mtu MTU | mtu lock MTU

  the mtu along the path to the destination

- window NUMBER

  TCP 窗口的最大值

- rtt TIME | rttvar TIME

  the initial RTT

- rto_min TIME

  the minimum TCP retransmission timeout

  TCP重传超时时间

- ssthresh NUMBER

  the initial slow start threshold

- cwnd NUMBER

  the clamp for congestion window

- initcwnd NUMBER

  the initial congestion window

- initrwnd NUMBER

  initial receive window size

- features  FEATURES

  enable or disable pre-route features

- quickack BOOL

  enable or disable quick ack for connections to this destination

- congctl NAMC

- advmss

  the maximal segment size to advertise to these destinations

- nexthop

  the nexthop of a multipath route

  nexthop 是一个复杂参数格式如下

  1. via [FAMILY] ADDRESS

     the nexthop router

  2. dev NAME

     the output device

  3. weight NUMBER

     路由权重

- scop SCOPE_VAL

  路由在那个scop使用，如果没有指定默认使用global，表示全局生效

- protocol PTROTO

  标识当前路由的来源，PTROTO值使用`/etc/iproute2/rt_protos`中的值，如果没有指定默认使用 boot，可以是如下的值

  1. redirect

     the route was installed due to an ICMP redirect

  2. kernel

     the route was installed by the kernel during autoconfiguration

  3. boot 

     the route was installed during the bootup sequence. It assumes the route was added by someone who does’t understand what the are doing

  4. static

     the route was installed by the administrator to override dynamic routing

  5. ra

     the route was installed by router discovery

  配置文件中的值没有保留，可以由用户自行设置

- onlink

  nexthop 会被任务直联，即使实际不是

- pref PREF

  the ipv6 route preference

- encap ENCAPTYPE ENCAPHDR

- expires TIME

  route will be deleted after the expires time

  目前只支持IPv6





