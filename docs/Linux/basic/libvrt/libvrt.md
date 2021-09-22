# brctl

> bridge对应端设备是透明的，所以不会在traceroute中出现

## 概述

bridge controller 用于管理linux kernel的bridge。

## command

brctl还可以配置STP(spanning tree protocal)，具体查看manual page

- `brctl show `

  查看当前所有的bridge

  ```
  root in /home/ubuntu λ brctl show
  bridge name     bridge id               STP enabled     interfaces
  br-30bc447f31f1         8000.02427be9755a       no
  docker0         8000.0242185f73f2       no
  virbr0          8000.52540012daf9       yes             virbr0-nic
  ```

- `brctl addbr <name>`

  添加一个bridge实例

  ```
  root in /home/ubuntu λ ip a show br0
  52: br0: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
      link/ether ce:c1:8b:fd:06:4e brd ff:ff:ff:ff:ff:ff
  ```

  默认创建或为down的状态

- `brctl delbr <name>`

  删除一个bridge实例

- `brctl addif <brname> <ifname>`

- `brctl delif <brname> <ifname>`

- `brctl show <brname>`

- `brctl showmacs <brname>`

  

  