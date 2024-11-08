# Host-Only（仅主机模式）

> 外部网络对虚拟系统的网卡是【**不可见**】的；虚拟系统的网卡对外部网络也是**不可见**的。
> 这种模式相当于【**双向**】防火墙的效果。相对而言，这种模式用得较少。当你想搭建一个跟外界隔离的虚拟内部网络，可以使用这种模式。

参考:

https://www.linuxidc.com/Linux/2016-09/135521p3.htm

Host-Only模式其实就是NAT模式去除了虚拟NAT设备，然后使用VMware Network Adapter VMnet1虚拟网卡连接VMnet1虚拟交换机来与虚拟机通信的，Host-Only模式将虚拟机与外网隔开，使得虚拟机成为一个独立的系统，只与主机相互通讯。其网络结构如下图所示：

<img src="https://www.linuxidc.com/upload/2016_09/160926204874121.png"/>

通过上图，我们可以发现，如果要使得虚拟机能联网，我们可以将主机网卡共享给VMware Network Adapter VMnet1网卡，从而达到虚拟机联网的目的。接下来，我们就来测试一下。

首先设置“虚拟网络编辑器”，可以设置DHCP的起始范围。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874122.png"/>

设置虚拟机为Host-Only模式。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874123.png"/>

开机启动系统，然后设置网卡文件。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874124.png"/>

保存退出，然后重启网卡，利用远程工具测试能否与主机通信。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874125.png"/>

主机与虚拟机之间可以通信，现在设置虚拟机联通外网。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874126.png"/>

我们可以看到上图有一个提示，强制将VMware Network Adapter VMnet1的ip设置成192.168.137.1，那么接下来，我们就要将虚拟机的DHCP的子网和起始地址进行修改，点击“虚拟网络编辑器”

<img src="https://www.linuxidc.com/upload/2016_09/160926204874127.png"/>

重新配置网卡，将VMware Network Adapter VMnet1虚拟网卡作为虚拟机的路由。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874128.png"/>

重启网卡，然后通过 远程工具测试能否联通外网以及与主机通信。

<img src="https://www.linuxidc.com/upload/2016_09/160926204874129.png"/>

测试结果证明可以使得虚拟机连接外网。
