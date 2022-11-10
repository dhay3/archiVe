# IPtables Digest

ref

https://www.netfilter.org/

https://wiki.archlinux.org/title/iptables

https://www.zsythink.net/archives/1199

## Netfilter

在介绍 iptables 之前需要了解一下 netfilter

netfilter 是 FOSS 一个开源的软件，作为 framework 内嵌在 linux kernel 中。主要提供如下功能：

1. packet filtering
2. network address [ and port ] translation (NA[P]T)
3. packet logging
4. userspace packet queueing
5. packet manglilng

而 nftables(iptables 的升级版)/iptables 就是基于 netfilter，所以也就继承了 netfilter 的所有功能

## IPtables

IPtables 是一个通过 netfilter 来管理 kernel 层面的数据报文的 userspace CML，主要有如下功能

1. listing the contents of the packet filter ruleset
2. adding/removing/modifying rules in the packet filter ruleset
3. listing/zeroing per-rule counters of the packet filter ruleset

在 Linux 还提供一些包用于管理 iptables，例如 firewalld 就是，同时 firewalld 还有提供 GUI 的包

## IPtables VS NIC

![Snipaste_2020-10-13_23-09-07](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-10-13_23-09-07.2bumkikl390k.png)

入向报文

NIC -> IPtables

出向报文

IPtables -> NIC
