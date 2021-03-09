# UEFI & BIOS

https://blog.csdn.net/F8qG7f9YD02Pe/article/details/79551663

如果想要双系统，需要知道主机的启动方式(划分区)。==如果需要双引导，建议始终在相同的引导模式下安装两个操作系统==

UEFI启动对比BIOS启动的优势有三点：

1. 安全性更强
   UEFI启动需要一个独立的分区，它将系统启动文件和操作系统本身隔离，可以更好的保护系统的启动。即使系统启动出错需要重新配置，我们只要简单对启动分区重新进行配置即可。而且，对于win8系统，它利用UEFI安全启动以及固件中存储的证书与平台固件之间创建一个信任源，可以确保在加载操作系统之前，近能够执行已签名并获得认证的“已知安全”代码和启动加载程序，可以防止用户在根路径中执行恶意代码。
2. 启动配置更灵活
   EFI启动和GRUB启动类似，在启动的时候可以调用EFIShell，在此可以加载指定硬件驱动，选择启动文件。比如默认启动失败，在EFIShell加载U盘上的启动文件继续启动系统。
3. 支持容量更大
   传统的BIOS启动由于MBR的限制，默认是无法引导超过2.1TB以上的硬盘的。随着硬盘价格的不断走低，2.1TB以上的硬盘会逐渐普及，因此UEFI启动也是今后主流的启动方式。

### windows查看引导方式

使用命令`msinfo32`来查看，如果在BIOS模式一栏是UEFI就是uefi方式启动，如果是legacy就是普通bios启动。

### linux中查看引导方式

查找`/sys/firmware/efi`文件。如果使用BIOS就没有改文件

![img](https://ss.csdn.net/p?https://mmbiz.qpic.cn/mmbiz_png/W9DqKgFsc6ibHJT2OmUdcfSvXr2icU8tDrx7jHhAkM18ib0RAkicxpTIiaURU4X5hpMs330vbbYlgsNhcRRrSvSK46Q/640?wx_fmt=png)













