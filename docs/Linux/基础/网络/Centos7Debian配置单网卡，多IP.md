# Centos7/Debain配置单网卡，多IP

> 可以配置不同net-id的IP，用于局域网通信

## 暂时配置

```
ifconfig ens33:0 192.168.80.101 netmask 255.255.255.0 up
```

## 永久配置

复制一份原有的网卡配置文件，以 `identifier:number`来表示网卡

<img src="..\..\..\imgs\_Linux\Snipaste_2020-09-18_14-01-45.png"/>

修改内容如下，注意NAME和DEVICE

```
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=static
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=ens33:0
DEVICE=ens33:0
ONBOOT=yes
IPADDR=192.168.80.101
NETMASK=255.255.255.0
GATEWAY=192.168.80.2
```

重启网络即可，`systemclt restart network`

> 如果手动ifconfig ens33:0 down, 重新使用ifconfig ens33:0 up就会报错。由于MAC地址重复了
>
> SIOCSIFFLAGS: Cannot assign requested address

解决方法

手动分配一个MAC地址

`ifconfig ens33:0 hw class address`

## Debian配置

```
# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).

source /etc/network/interfaces.d/*

# The loopback network interface
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet static
address 192.168.80.200
netmask 255.255.255.0
broadcast 192.168.80.255
network 192.168.80.0
gateway 192.168.80.2


auto eth0:1
iface eth0:1 inet static
address 192.168.80.230
netmask 255.255.255
broadcast 192.168.80.255
```

