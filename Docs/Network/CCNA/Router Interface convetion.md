# Router Interface convetion

参考：

https://www.juniper.net/documentation/us/en/software/junos/interfaces-fundamentals/topics/topic-map/router-interfaces-overview.html

## 概述

每个router上的interface都按`physical[:channel].logical`格式命名channcle可以为空，在Cisco的Router上可以使用`show interfaces`命令可以查看当前Router上的所有interfaces

## pysical part

pysical part用于标识物理设备，该部分有由`type[fpc_slot/]pic_slot/port`可以将fpc_slot和pic_slot理解为NIC的坐标

- type指的是media type，常见的有fast ethernet interfaces(fe)，Gigabit Ethernet interface(ge)，Serial Interface(se)

> slot通常用于代表FPC或PIC(如果没有FPC)所在的位置

- FPC Flexible PIC Concentrator, An interface concentrator on which physical interface cards (PICs) are mounted.
- PIC Physical Interface Card, Network interface–specific card that can be installed on an FPC in the router. Also known as card(NIC), blade, module.
- port用于标识是PIC上的第几个port

![](https://www.juniper.net/documentation/us/en/software/junos/interfaces-fundamentals/images/h1415.gif)

例如上图`media_type 5/1/0`到`media_type 5/1/3`

![](https://www.computernetworkingnotes.org/images/cisco/ccna-study-guide/cisco-2960-switch.jpg)

这是一个有4个slot的router

![](https://blog.router-switch.com/wp-content/uploads/2012/06/router_int_naming_2.gif)

## logical part

逻辑上区分router interface