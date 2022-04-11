# VM openWRT

参考：https://www.jianshu.com/p/832af1b3b4cd

[TOC]

## 准备工作

1. 下载[openWRT ](https://firmware.koolshare.cn/LEDE_X64_fw867/)，这里选用通用lede版本

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_11-23-17.png"/>

   根据自己固件的类型选择是BIOS或是EFI，我们可以通过`msinfo32`来查看

2. 创建虚拟机

   > 为了使openWRT能连接外玩选用NAT或是bridge（校园网桥接模式存在问题）
   >
   > 也可以采用双网卡，host-only+bridge|NAT

   这里注意的步骤

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_11-32-17.png"/>
<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_11-33-54.png"/>

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_11-34-50.png"/>

   到这里就可以正常启动openWRT

## 网络配置

我这里使用NAT模式，在该模式下hostOS无法访问guestOS，所以我们想要访问openWRT需要将openWRT的IP置为与vmnet8 adaptor 同一网段

可以通过`ipconfig`或是vmware中编辑-->虚拟网络编辑器 来查看vmnet8交换机的net-id

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_11-47-35.png"/>

我们这里就需要将openWRT的IP置为`192.168.80.X`或是将vmnet8 adaptor的IP置为`192.168.1.X`。由于小编vmware中有多台虚拟机与vmnet8连接设置了静态IP，所以这里采用修改openWRT IP的方案。

> 注意在修改网络配置文件时，先备份一份
>
> cp  /etc/config/network /etc/config/network.bak

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_12-01-56.png"/>

> 如果出现下载文件校验不一致错误时, 请检查网关和DNS配置

然后我们通过hostOS访问`192.168.80.254`就会到openWRT的控制台，默认密码`koolshare`。

目前还无法访问酷软，那是因为没有配置DNS，选中正在使用网卡

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_12-23-34.png"/>

添加DNS服务器

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_12-43-21.png"/>

