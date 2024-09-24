# Linux ip-address

## 0x1 Digest

ip address 用于管理 IPv4 和 IPv6，each device 必须要有一个 IP address 才可以使用不同的网络协议，针对同一个 device 可以配置多个 IP

## 0x2 EBNF

```
       IFADDR := PREFIX | ADDR peer PREFIX [ broadcast ADDR ] [ anycast ADDR
               ] [ label LABEL ] [ scope SCOPE-ID ]

       SCOPE-ID := [ host | link | global | NUMBER ]

       FLAG-LIST := [ FLAG-LIST ] FLAG

       FLAG := [ permanent | dynamic | secondary | primary | [-]tentative |
               [-]deprecated | [-]dadfailed | temporary | CONFFLAG-LIST ]

       CONFFLAG-LIST := [ CONFFLAG-LIST ] CONFFLAG

       CONFFLAG := [ home | mngtmpaddr | nodad | noprefixroute | autojoin ]

       LIFETIME := [ valid_lft LFT ] [ preferred_lft LFT ]

       LFT := [ forever | SECONDS ]

       TYPE := [ bridge | bridge_slave | bond | bond_slave | can | dummy |
               hsr | ifb | ipoib | macvlan | macvtap | vcan | veth | vlan |
               vxlan | ip6tnl | ipip | sit | gre | gretap | erspan | ip6gre |
               ip6gretap | ip6erspan | vti | vrf | nlmon | ipvlan | lowpan |
               geneve | macsec ]
```

## 0x3 Commands

syntax：`       ip address { add | change | replace } IFADDR dev IFNAME [ LIFETIME ] [ CONFFLAG-LIST ]`

### ip address add

为指定device 配置ip

**IFADDR**

- `dev IFNAME`

  the name of the device to add the address to

- `local ADDRESS`

  device 的 IP，可使用无类IP，同时表示IP和掩码

- `peer ADDRESS`

  点对点对端IP，可以使用无类IP。如果指定了peer，local address 不能指定prefix

  ```
  [vagrant@10 ~]$ sudo ip a a local 192.168.80.1 peer 192.168.80.2 dev d0
  ```

- `broadcast ADDRESS`

  the broadcast address on the device

  可使用`+`或`-`表示会根据interface 的 prefix 自动设置广播地址

  ```
  [root@10 vagrant]# ip a add local 192.168.80.1 peer 192.168.80.2 broadcast + dev d0
  ```

- `label LABEL`

  每个地址可以关联一个label，为了 linux-2.0 net aliases 统一。每个label需要以`devicename:x`来表示，x 的值最大 15 chars

  ```
  [root@10 vagrant]# ip a a 192.168.80.1 label d0:haha dev d0
  ```

- `scope SCOPE_VALUE`

  the scope of the area where this address is valid

  所有可用的scopes记录在`/etc/iproute2/rt_scope`，预设的scope有

  1. global，ip  全局有效
  2. site，ip 只在当前站点有效
  3. link，只在当前device上有效
  4. host，ip  只在本机有效

**LIFE_TIME**

- `valid_lft LFT`

  如果超时ip会被kernel移除，默认 forever。必须要比`preferred_lft`值大

  ```
  [root@10 vagrant]# ip a a local 192.168.80.1 valid_lft forever preferred 100 dev d0
  ```

- `preferred_lft LFT`

  如果超时ip不在被new outgoing connections 使用，默认 forever

**CONFLAG**

- `home`

  指定 home address，只针对 IPv6

- `mngtmpaddr`

- `nodad`

- `noprefixroute`

  不会自动生成对应当前ip prefix的路由

- `autojoin`

  自动加入组播组

### ip address delete

syntax：`ip address del IFADDR dev IFNAME [ mngtmpaddr ]`

删除device ip，具体用法和`ip addr add`一样

### ip address show

syntax：`ip address [ show [ dev IFNAME ] [ scope SCOPE-ID ] [ to PREFIX ] [ FLAG-LIST ] [ label PATTERN ] [ master DEVICE ] [ type TYPE ] [ vrf NAME ] [ up ] ]`

查看 L4 信息

**IFADDR**

- `dev IFNAME`(default)

  查看具体device L4 信息

- `scope SCOPE_VAL`

  查看指定scope的device L4 信息

- `label PATTERN`

  查看label匹配PATTERN的device L4信息

- `type TYPE`

  查看指定link type 的device L4 信息

- `up`

  only list running interfaces

#### keyword

- mtu

  maximal transform unit，单包最大的传输值

  以太网卡默认1500

  byte

- qdisc

  queueing discipline

  接受和发送数据的缓冲队列使用的算法，具体可以使用的值参考[链接](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/configuring_and_managing_networking/linux-traffic-control_configuring-and-managing-networking)，常见的有noqueue表示不使用队列缓存信息，noop使用blackhole，所有发送到该iface的数据包都会被丢弃

  可以用`tc qdisc show dev <iface>`查看

- state

  表明iface的状态是否启用

- master

  该iface是另一个iface的子接口

  ```
  4: virbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
      link/ether 52:54:00:12:da:f9 brd ff:ff:ff:ff:ff:ff
      inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
         valid_lft forever preferred_lft forever
  5: virbr0-nic: <BROADCAST,MULTICAST> mtu 1500 qdisc fq_codel master virbr0 state DOWN group default qlen 1000
      link/ether 52:54:00:12:da:f9 brd ff:ff:ff:ff:ff:ff
  ```

- group

  ip 组

- qlen

  缓存队列最大的值

### ip addrss flush

> flush命令是不可恢复的

删除

```
➜  cpl ip a add 81.81.81.81 dev docker0
➜  cpl ip a flush dev docker0 
➜  cpl ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: wlp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 64:bc:58:bd:a6:19 brd ff:ff:ff:ff:ff:ff
    inet 192.168.0.103/24 brd 192.168.0.255 scope global dynamic noprefixroute wlp1s0
       valid_lft 6907sec preferred_lft 6907sec
    inet6 240e:390:e52:7cf1::1001/128 scope global dynamic noprefixroute 
       valid_lft 53710sec preferred_lft 53710sec
    inet6 240e:390:e52:7cf1:a44d:89dc:dc5:6395/64 scope global dynamic noprefixroute 
       valid_lft 86042sec preferred_lft 14042sec
    inet6 fe80::5e7a:b687:b448:fd5f/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:4c:00:fb:f7 brd ff:ff:ff:ff:ff:ff
```

和show有一样的参数，当和`-s`一起使用时可以输出详细的信息，包括删除的address

```
➜  cpl ip -s a flush dev docker0 

*** Round 1, deleting 1 addresses ***
*** Flush is complete after 1 round ***
```

如果使用两次`-s`可以dump所有删除的address

```
➜  cpl ip -s -s a flush dev docker0
3: docker0    inet 81.81.81.81/32 scope global docker0
       valid_lft forever preferred_lft forever

*** Round 1, deleting 1 addresses ***
*** Flush is complete after 1 round ***
```

