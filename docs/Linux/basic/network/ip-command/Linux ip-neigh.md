# Linux ip-neighbour

## 0x1 Digest

管理 neighbour/arp 表

## 0x2 EBNF

```
ip [ OPTIONS ] neigh  { COMMAND | help }

ip neigh { add | del | change | replace } { ADDR [ lladdr LLADDR ] [
nud STATE ] | proxy ADDR } [ dev DEV ]

ip neigh { show | flush } [ proxy ] [ to PREFIX ] [ dev DEV ] [ nud
STATE ] [ vrf NAME ]

STATE := { permanent | noarp | stale | reachable | none | incomplete |
delay | probe | failed }
```

## 0x2 Commands

### ip neighbour add

add a new neighbour entry

### ip neighbour change

change an existing entry

### ip  neighbour replace

syntax：`       STATE := { permanent | noarp | stale | reachable | none | incomplete | delay | probe | failed }`

add a new entry or change an existing one

- `to ADDRESS`（default）

  指定neigh IP

- `dev NAME`

  the interface to wich this neighbour is attached

  本端的接口

- `lladdr`

  指定neigh link layer address（MAC地址）

- `nud STATE`

  the state of the neighbour entry

  nud 是 Neighbour unreachability detection 的缩写，可以如下的值

  1. permant

     the neigh entry is valid forever and cna be only be removed administratively

  2. noarp

     the neigh entry is valid. No attemps to validate this entry will be made bbut it can be  removed when its lifetime expires

  3. reachable

     the neighbour entry is valid until the reachability timeout expires

  4. stale

     the neigh entry is valid but suspicious

  5. none

     this is a pseudo state

  6. incomplete

     the neighbour entry has not yet been validated/resolved

     通常可能是

