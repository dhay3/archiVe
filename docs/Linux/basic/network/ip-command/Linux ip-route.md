# Linux ip-route

## 0x1 Digest

ip route 用于管控 kernel routing table

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

  

