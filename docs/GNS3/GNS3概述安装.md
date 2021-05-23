# GNS3概述安装

参考：

https://docs.gns3.com/docs/

GNS3是一个网络模拟工具，类似的有Cisco tracer。

GNS3由两部分组成：

- The GNS3-all-in-one software (GUI)
- The GNS3 virtual machine (VM)，GNS3的一些组件需要在VM上运行(VM上有写好的API)

GUI是我们画网络拓扑图结构的客户端工具，需要由一个服务端进程(通常是VM)托管运行。有如下几种选择服务端的方式：

1. local GNS3 server
2. local GNS3 VM (recommended)
3. remote GNS3 VM

## 核心组件

1. Dynamips：运行Cisco IOS(Internetwork Operating System)
2. IOU：IOS on Unix
3. Winpcap：协议抓包的底层支持
4. Wireshark：主流抓包和网络协议分析软件
5. Qemu：Cisco ASA防火墙，入侵检测IDS，Juniper Jun OS等设备的仿真
6. Pemu：Cisco PIX 防火墙的仿真
7. VPCS：PC进行的模拟

## GNS3 VM安装

> 出现[No module named 'apt_pkg' error](https://askubuntu.com/questions/1069087/modulenotfounderror-no-module-named-apt-pkg-error) 

这里我采用[remote GNS3 VM](https://docs.gns3.com/docs/getting-started/installation/remote-server/)方式，==这种方式project会存储在remote server opt目录以uuid为文件名==

执行如下方法，会在remote server上安装并运行openVPN，作为服务端

```
cd /tmp
curl https://raw.githubusercontent.com/GNS3/gns3-server/master/scripts/remote-install.sh > gns3-remote-install.sh
bash gns3-remote-install.sh --with-openvpn --with-iou --with-i386-repository
```

同时在$HOME目录下生成客户端的openVPN配置文件，名为`client.opvn`。校验服务端服务是否开启

```
#openvpn默认使用udp1194作为CS之间的通信通道
root in ~ λ netstat -lnpu | grep 1194
udp        0      0 0.0.0.0:1194            0.0.0.0:*                           17386/openvpn  

#创建一个tun
16: tun1194: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 100
    link/none 
    inet 172.16.253.1 peer 172.16.253.2/32 scope global tun1194
```

在客户端运行`client.opvn`，并校验连接

```
#可以使用--daemon以守护进程的方式运行，日志将会被转移到syslog中
cpl in ~/Desktop λ sudo openvpn client.opvn
...
#GNS3默认创建172.16.253.1 3080
nc -nvz 172.16.253.1 3080
```

如果连接不上确保防火墙正常关闭，使用HTTP连接

## IOS安装

参考：

https://www.youtube.com/watch?v=ykF15xI44mE

https://www.youtube.com/watch?v=YQcWuWGjppY

https://docs.gns3.com/docs/emulators/cisco-ios-images-for-dynamips/

 c3640, c3660, c3725, c3745 and c7200官方推荐IOS

单个IOS可能会耗尽CPU，所以需要设置一个==Idle-PC value==，尽可能保证IOS使用的CPU在正常范围内。每个IOS使用的Idle-PC value值都不一样。每个IOS使用的RAM也不同，具体配置参考上面的链接。

由于产品的生命周期，除了c7200外都不能在Cisco官网下载。(==c7200系列GNS3只支持7206==)

[官网IOS下载地址](https://software.cisco.com/download/home/282188585/type?catid=268437899)

```mermaid
graph LR
a(IOS Software)-->b(all release)-->c(15.2M)-->d(15.2.4M10)
```

[非官方下载地址](https://mega.nz/folder/nJR3BTjJ#N5wZsncqDkdKyFQLELU1wQ)

导入IOS参考：https://www.kjnotes.com/devtools/53

1. Network Adapter：不同设备可以选择的network adapters(NIC)数量和类型不同
2. WIC modules：WAN Interface Card(WIC)广域网接口卡

