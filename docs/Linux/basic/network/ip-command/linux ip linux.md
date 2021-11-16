# linux ip linux

## 概述

查看或配置设备上L2信息



## keyword

## link type

- bridge - ethernet bridge device
- bond - bonding device
- dummy - dummy network interface
- hsr - high-availability seamless redundancy
- macvlan - virtual interface based on link layer address and  tap
- vcan - virtual controller area network interface
- vxcan - virtual controller area network ethernet interface
- vlan - virtual lan interface
- vxlan - virtual extended LAN
- ipip - virtual tunnel interface ipv4 over ipv4
- vti - virtual tunnel interface
- ipvlan - interface for L3 based vlans
- vrf - interface for L3 vrf domain

## ip link add

- `link <device>`

  

- `numtxqueues <queue_count>`

  指定发送数据的队列数

- `numrxqueues <queue_count>`

  指定接受数据的队列数

- `gso_max_size <bytes>`

  指定接受packet的大小

- `gso_max_segs <segments>`

  指定接受segment(不包含包头)的大小

## ip link delete



