# GNS3概述安装

参考：

https://docs.gns3.com/docs/

GNS3是一个网络模拟工具，类似的有Cisco tracer。

GNS3由两部分组成：

- The GNS3-all-in-one software (GUI)

  GNS3的客户端

- The GNS3 virtual machine (VM)

  GNS3的一些组件需要在VM上运行(VM上有写好的API)

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

