# Linux ip address

## 概述

> 开启和关闭iface同样需要使用`ifup`和`ifdown`或`nmcli`

查看或配置设备NIC L4信息

`ip address`可以缩写成`ip a`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: wlp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 64:bc:58:bd:a6:19 brd ff:ff:ff:ff:ff:ff
    inet 192.168.0.103/24 brd 192.168.0.255 scope global dynamic noprefixroute wlp1s0
       valid_lft 6485sec preferred_lft 6485sec
    inet6 240e:390:e52:7cf1::1001/128 scope global dynamic noprefixroute 
       valid_lft 53289sec preferred_lft 53289sec
    inet6 240e:390:e52:7cf1:a44d:89dc:dc5:6395/64 scope global dynamic noprefixroute 
       valid_lft 86285sec preferred_lft 14285sec
    inet6 fe80::5e7a:b687:b448:fd5f/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
```

## keyword

- mtu

  maximal transform unit，单包最大的传输值

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

## ip address add

alias：ip a add

为指定iface添加protocol address

syntax：`ip address add <address> dev <device_name>`

```
➜  cpl ip a add 81.81.81.81/24 dev docker0 
➜  cpl ip a show dev docker0
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:4c:00:fb:f7 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet 81.81.81.81/24 scope global docker0
       valid_lft forever preferred_lft forever

```

如果没有指定net mask位数默认使用32位

- `peer <address>`

  指定remote endpoint用于p2p，但是可以使用CIDR

  ```
  ➜  cpl ip a add 82.82.82.82/32 peer 81.81.81.81/32  dev docker0
  ➜  cpl ip a show dev docker0
  3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
      link/ether 02:42:4c:00:fb:f7 brd ff:ff:ff:ff:ff:ff
      inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
         valid_lft forever preferred_lft forever
      inet 82.82.82.82 peer 81.81.81.81/32 scope global docker0
         valid_lft forever preferred_lft forever
  
  ```

- `label <label>`

  将address与指定的label name绑定，必须以`device’s name:label`的方式命名

  ```
  ➜  cpl ip a add 81.81.81.81 label docker0:01  dev docker0
  ➜  cpl ip a show dev docker0
  3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
      link/ether 02:42:4c:00:fb:f7 brd ff:ff:ff:ff:ff:ff
      inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
         valid_lft forever preferred_lft forever
      inet 81.81.81.81/32 scope global docker0:01
         valid_lft forever preferred_lft forever
  ```

- `scop <scop_value>`

  address在指定范围内有效，所有有效的scope在`/etc/iproute2/rt_scops`

  ```
  ➜  cpl cat /etc/iproute2/rt_scopes 
  #全局生效
  0       global
  #不生效
  255     nowhere
  #只在本机生效
  254     host
  #只在链路层有效
  253     link
  200     site
  ```

- `mertric <number>`

  address的优先级，值越高优先级越低

  ```
   ➜  cpl ip a add 82.82.82.82/32 metric 1000 dev docker0
  ➜  cpl ip a show dev docker0 
  3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
      link/ether 02:42:4c:00:fb:f7 brd ff:ff:ff:ff:ff:ff
      inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
         valid_lft forever preferred_lft forever
      inet 81.81.81.81/32 scope global docker0:01
         valid_lft forever preferred_lft forever
      inet 82.82.82.82/32 metric 1000 scope global docker0
         valid_lft forever preferred_lft forever
  ```

- `valid_lft <lft>`

  address有效的时间，如果过期就被kernel删除。默认使用forever。具体查看RFC 5.5.4 

- `preferred_lft <lft>`

  address有效的时间，如果过期就不再被使用，但是不删除。默认使用forever。具体查看RFC5.5.4

- `noprefixroute`

  带有该标记的address不会自动生成路由信息

## ip address delete

alias：ip a del

和add方式一样，device name是必须参数，==如果没有指定address，默认删除第一个address，最好指定net mask位数，否则会使用统配删除，应该指定所有的具体信息（scope，metrics）==

syntax：`ip address delete <address> dev <device_name>`

删除前

```
➜  cpl ip a show label docker0
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet 81.81.81.81/32 scope global docker0
       valid_lft forever preferred_lft forever
    inet 81.81.81.81/24 scope global docker0
       valid_lft forever preferred_lft forever
```

删除

```
➜  cpl ip a delete 81.81.81.81/32 dev docker0
```

删除后

```
➜  cpl ip a show label docker0
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet 81.81.81.81/24 scope global docker0
       valid_lft forever preferred_lft forever
```

删除peer mode IP

```
➜  cpl ip a delete 82.82.82.82 peer 81.81.81.81/32 dev docker0
```

## ip address show

> 与`ifconfig`不同的是，`ip a`会将同一张NIC上的iface划到一起，而前者不会。==相同的是重启后都会还原==

alias：ip a 

- `dev <ifname>`

  查看指定device相关的L4信息

  ```
  ➜  cpl ip addr show lo   
  ```

- `scope <scope_value>`

  查看指定scope相关L4的信息

  ```
  ➜  cpl ip addr show scope host  
  ```

- `to <prefix>`

  只查看相关prefix的L4信息

  ```
  ➜  cpl ip addr show to 127.0.0.1
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
  ```

- `label <pattern>`

  只查看和label相关的L4信息

  ```
  ➜  cpl ip addr show label lo
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host 
         valid_lft forever preferred_lft forever
  ```

- `master <device>`

  只显示该设备上的子接口

  ```
  ip a show master docker0
  ```

- `up`

  只显示up的interface

  ```
  ip a show up
  ```

## ip address flush

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

