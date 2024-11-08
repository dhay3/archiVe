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

  > 注意iproute2和arp中显示的状态不同，如
  >
  > iproute2 中显示failed 的条目，在arp 中显示incomplete

  ```
  cpl in ~ λ ip neigh         
  192.168.2.1 dev wlp1s0 lladdr 86:2a:e4:7d:fc:ea DELAY 
  192.168.2.241 dev wlp1s0 lladdr 98:f1:81:ac:3a:07 STALE 
  192.168.2.10 dev wlp1s0 FAILED 
  192.168.2.2 dev wlp1s0 FAILED 
  fe80::842a:e4ff:fe7d:fcea dev wlp1s0 lladdr 86:2a:e4:7d:fc:ea router STALE 
  cpl in ~ λ arp -n
  Address                  HWtype  HWaddress           Flags Mask            Iface
  192.168.2.1              ether   86:2a:e4:7d:fc:ea   C                     wlp1s0
  192.168.2.241            ether   98:f1:81:ac:3a:07   C                     wlp1s0
  192.168.2.10                     (incomplete)                              wlp1s0
  192.168.2.2                      (incomplete)                              wlp1s0
  ```

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

     造成这种的原因可能是

     https://networkengineering.stackexchange.com/questions/50843/what-are-the-reasons-for-seeing-an-incomplete-arp#:~:text=Incomplete%20ARP%20requests%20have%20two,destination%20node%20may%20be%20down.

     1. 对端端口或节点物理关闭
     2. a node within the 192.168.0.0/24 subnet is incorrectly set up with 192.168.0.10/*16* it considers a destination like 192.168.16.1 local. It will not try to use a gateway but attempt a direct ARP which stays incomplete.
     3. 30位掩码的网段点对点，例如192.168.80.1需要知道192.168.80.2的MAC，当前IP的端口或主机存在会返回给192.168.80.1MAC。但是现在线接错了，当前LAN中没有192.168.80.2，就会返回incomplete。即当前LAN中根本就没有个IP的主机在线

  7. delay

     neighbor entry validation is currently delayed

  8. probe

     neighbor is being probed

  9. failed

     max bumber of probes exceed without success neighbor validation has ultimately failed

### ip neighbour delete

syntax：`ip neigh { add | del | change | replace } { ADDR [ lladdr LLADDR ] [ nud STATE ] | proxy ADDR } [ dev DEV ]`

delete a neighbour entry

和`ip neigh add`参数相同，手动删除 noarp  的条目会导致

### ip neighbour show

syntax：`ip neigh { show | flush } [ proxy ] [ to PREFIX ] [ dev DEV ] [ nud STATE ] [ vrf NAME `

list neighbour entries

- `to ADDRESS`(default)

  查看到指定IP是否有arp记录

- `dev NAME`

  只查看和device关联的arp记录

- `unused`

  查看当前没有被使用的arp记录，一般是条目无效的（状态 incomplete，failed，state）

  ```
  pl in ~ λ ip neigh show unused
  192.168.2.241 dev wlp1s0 lladdr 98:f1:81:ac:3a:07 STALE 
  192.168.2.10 dev wlp1s0 FAILED 
  192.168.2.2 dev wlp1s0 FAILED 
  ```

- `nud STATE`

  只查看指定num STATE条目的



### ip neighbour flush

syntax：`ip neigh { show | flush } [ proxy ] [ to PREFIX ] [ dev DEV ] [ nud STATE ] [ vrf NAME`

flush neighbour entries

清空 neighbour entries，参数和`ip neigh show`相同。必须要执行参数才能运行



