# Linux ifconfig

## ifconfig

> 用于配置网络，已过时(不能正确显示MAC地址)使用`ip addr`和`ip link`替换。

1. 查看网卡，接口，网络配置

   ```
   [root@chz ~]# ifconfig
   ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
           inet 192.168.80.131  netmask 255.255.255.0  broadcast 192.168.80.255
           inet6 fe80::a164:2ef4:8841:5fc7  prefixlen 64  scopeid 0x20<link>
           ether 00:0c:29:d7:d1:68  txqueuelen 1000  (Ethernet)
           RX packets 1045  bytes 120646 (117.8 KiB)
           RX errors 0  dropped 0  overruns 0  frame 0
           TX packets 1220  bytes 100883 (98.5 KiB)
           TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
   
   lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
           inet 127.0.0.1  netmask 255.0.0.0
           inet6 ::1  prefixlen 128  scopeid 0x10<host>
           loop  txqueuelen 1000  (Local Loopback)
           RX packets 283  bytes 40151 (39.2 KiB)
           RX errors 0  dropped 0  overruns 0  frame 0
           TX packets 283  bytes 40151 (39.2 KiB)
           TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
   ```

   单网卡ens33，本地环回地址lo

2. 开启关闭指定网卡

   > 可以使用ifup，ifdown替换

   ```
   ifconfig ens33 down
   ifconfig ens33 up
   ```

3. 配置IP地址

   ==重启网卡就会失效==，如果需要配置为静态IP，需要写入配置文件。如果只是A，B，C类网络也可以不指定netmask或是broadcast。==同时也起到变更IP的作用。==

   ```
   ifconfig ens33 192.168.80.200 netmask 255.255.255.0 broadcast 192.168.80.255
   ```









