# Linux ip-route

## 0x1 Digest

ip route 用于管控 kernel routing table

## 0x2 Terms

### Route tables

Linux 2.x 将路由分成多种 routing tables 用 1 - $2^{32}$ 或 字符来标识，默认记录在`/etc/iproute2/rt_tables`。其中`local` route table 是不可见的（包含local路由和broadcast路由）

### defualt

如果参数有default标识，表示使用 command 时的缺省值，这时可以省略参数，例如

`ip route get 1.1.1.1`等价与`ip route get to 1.1.1.1`

## 0x3 ENBF

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

## 0x4 Route types

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

## 0x5 Commands

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

- to TYPE PREFIX（default）

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

### ip route delete

syntax：`ip route { add | del | change | append | replace } ROUTE`

delete route

参数和 ip route add 相同

### ip route show

syntax：`ip route { show | flush } SELECTOR`

list routes

- to SELECTOR（default）

  使用指定的SELECTOR过滤目的IP（如果没有指定SELECTOR，默认使用root 0/0 即列出 entire table），可以是如下值

  1. root

     root PREFIX selects routes with prefixes not shorter than PREFIX

     例如 root 0/0 会选择 entire routing table

     ```
     [vagrant@localhost ~]$ ip route show root 10.0.2.0/23
     10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 100
     ```

  2. match 

     match PREFIX selects routes with prefixes not longer than PREFIX

     例如 10.0/16 匹配 10.0/16，10/8，0/0，但是不匹配 10.1/16 和 10.0.0/24

  3. exact

     exact PREFIX or just PREFIX selects routes with this exact PREFIX

- table TABLEID

  show routes from table

  默认只显示 mian table，TABLEID的值可以是

  1. all - list all of the tables
  2. cache - dump the routing cache

- vrf NAME

  show the routes for the table associated with the vrf name

- from SELECTOR

  使用SELECTOR过滤源IP，SELECTOR可以使用的值和 to 参数相同

- protocol PTROTO

  查看指定proto类型的路由

  ```
  cpl in ~/VagrantMachines/centos7 λ ip route show proto kernel
  30.131.78.0/24 dev wlp1s0 scope link src 30.131.78.32 metric 600 
  172.17.0.0/16 dev docker0 scope link src 172.17.0.1 linkdown
  ```

- type TYPE

  只显示指定 TYPE 路由

- dev NAME

  only list routes of this type

- via [FAMILY] PREFIX

  only list routes going via the nexthop routers selected by PREFIX

  ```
  [vagrant@localhost ~]$ ip route show via 10.0.2.2/24
  default via 10.0.2.2 dev eth0 proto dhcp metric 100
  ```

- src PREFIX

  only list routes with preferred source addresses selected by PREFIX

### ip route flush

syntax：`ip route { show | flush } SELECTOR`

清空指定路由条目

### ip route get

获得指定route条目，该命令不等价与`ip route show`

syntax：`ip route get ADDRESS [ from ADDRESS iif STRING  ] [ oif STRING ] [ tos TOS ] [ vrf NAME ]`

- to ADDRESS(default)

  查看到指定ADDRESS的路由

  ```
  [vagrant@localhost ~]$ ip route get to 1.1.1.1
  1.1.1.1 via 10.0.2.2 dev eth0 src 10.0.2.15 
      cache 
  ```

- from ADDRESS

  查看源地址

### ip route save | restore

save将路由保存成binary，restore 读取 binary并载入路由

```
ip route save > out
ip route restart < out
```

## 0x6 Examples

```
[vagrant@localhost ~]$ sudo ip route add to default via 10.0.2.1 dev eth0
```







