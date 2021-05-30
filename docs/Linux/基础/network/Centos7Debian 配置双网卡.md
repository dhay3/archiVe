# Centos7/Debian 配置双网卡

1. 添加一块虚拟网卡

<img src="..\..\..\..\imgs\_Linux\Snipaste_2020-09-18_16-41-43.png"/>

2. `ifconfig`查看网卡device name，新建一个对应的配置文件`vim /etc/sysconfig/network-script/ifcfg-name`

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
   NAME=ens34
   DEVICE=ens34
   ONBOOT=yes
   IPADDR=192.168.10.100
   NETMASK=255.255.255.0
   ```

   这里使用vmnet01 host-only，所以不需要设置网关

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
#与centos不同
dns-nameserver 8.8.8.8 

auto eth1
iface eth1 inet static
address 192.168.10.200
netmask 255.255.255
broadcast 192.168.10.255
```



