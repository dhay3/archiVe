# netplan

参考：

https://netplan.io/reference/

## 概述

在Ubuntu 17.10后引入netplan来管理网络，使用yaml来配置。

==再通过`ifupdonw`工具包或修改`/etc/network/interfaces`中的配置是无效的，需要通过netplan==

所有`/{lib,etc,run}/netplan/*.yaml`中的配置文件都会被考虑

<img src="https://assets.ubuntu.com/v1/a1a80854-netplan_design_overview.svg" alt="img" style="zoom:80%;" />

netplan中定义了两种设备

- physical devices

  有三种值ethernet，modem，wifi。==支持热插拔（修改配置后无需重启网络）==。

  通过`math`属性匹配iface。如果没有`math`属性，使用`id`匹配iface 

- virtual devices

  veth，bridge，bond

## subcommands

- netplan generate

  将配置文件转成systemd-networkd和NetworkManager可以读懂的

  不需要手动执行，会在boot阶段，netplan apply和netplan try时被调用

- netplan apply

  让yaml配置文件生效

- netplan try

  ==测试配置是否正确==

## 物理设备通用属性

- match(mapping)

  通过如下hardware properties匹配iface
  - name(scalar)

    当前iface的名字，支持通配符用于匹配名字

  - macaddress(scalar)

    设备的macaddress，不支持通配符

  - driver(scalar)

    kernel driver name

  eg.

  ```
  #匹配所有名字含有enp的iface
  match:
     name: enp2*
  ----
  #匹配mac地址为如下地址的iface
   match:
     macaddress: 11:22:33:AA:BB:FF
     
  ----
  
   match:
     driver: ixgbe
     name: en*s0
  ```

- set-name

  表示只精确匹配

  会匹配eth0，eth01，eht0*

  ```
   match:
    name: eth0
  ```

  但是使用set-name只匹配eth0

## 通用配置属性

- renderer(scalar)：

  实际管理网络的后端程序，可选`networkd`和`NetworkManager`，默认值为`networkd`。可以出现在network，或者device type(eg. ethernets)中。

- dhcp4(bool)

  是否对IPv4启用dhcp，默认false

- dhcp6(bool)

  是否对IPv6启用dhcp，默认false

- critical(bool)

  当rederer的守护线程重启后不会重启分配IP

- addresses(sequence scalar and mappings)

  指定静态地址使用（==可以配置多个ip==），以下的属性只有networkd才支持

  - lifetime

  - label

  ```
    ethernets:
      eth0:
        addresses:
          - 10.0.0.15/24:
              lifetime: 0
              label: "maas"
          - "2001:1::1/64"
  ```

- gateway4，gateway6(scalar)

  手动指定的IPv4/6的网关

- nameservers(mapping)

  设置DNS nameservers和search domain，和`resolv.conf`中的含义相同

  ```
    ethernets:
      id0:
        [...]
        nameservers:
          search: [lab, home]
          addresses: [8.8.8.8, "FEDC::1"]
  ```

- macaddress(scalar)

  设置NIC的mac地址

- mtu(scalar)

  设置传输的最大unit，默认1500byte

## 路由

- routes(mapping)

  - from(scalar)

    指定src 

  - to(scalar)

    指定dest

  - via(scalar)

    指定路由的网关

  - metric(scalar)

    指定路由的优先级

  - scope(scalar)

    指定路由的类型，可以使用global，link，host

## 例子

最上层节点使用两个属性，network和version(表示使用yaml几)，然后是device(ethernets,modems,wifis,bridges)，然后是id。

```
root in /etc/netplan λ cat 50-cloud-init.yaml
# This file is generated from information provided by
# the datasource.  Changes to it will not persist across an instance.
# To disable cloud-init's network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
network:
    version: 2
    ethernets:
        eth0:
            dhcp4: true
            match:
                macaddress: 00:16:3e:0a:be:8b
            set-name: eth0
```

配置多个IPv4地址多条指定路由，路由的优先级相同

```
network:
  version: 2
  renderer: networkd
  ethernets:
    eno1:
      addresses:
      - 10.0.0.10/24
      - 11.0.0.11/24
      nameservers:
        addresses:
          - 8.8.8.8
          - 8.8.4.4
      routes:
      - to: 0.0.0.0/0
        via: 10.0.0.1
        metric: 100
      - to: 0.0.0.0/0
        via: 11.0.0.1
        metric: 100
```

桥接配置

```
# Let NetworkManager manage all devices on this system
network:
  version: 2
  renderer: NetworkManager
  bridges:
    br0:
      addresses: [192.168.80.100/24]
      gateway4: 192.168.80.1
      mtu: 1500
      nameservers:
        addresses: [8.8.8.8]
      dhcp4: no
      dhcp6: no
```

