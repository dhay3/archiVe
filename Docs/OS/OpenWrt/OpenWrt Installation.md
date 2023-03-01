# OpenWrt Installation

ref

https://openwrt.org/docs/guide-quick-start/start

https://openwrt.org/docs/guide-quick-start/factory_installation

https://openwrt.org/docs/guide-user/installation/before.installation

https://post.smzdm.com/p/akx3okxk/

https://blog.csdn.net/zhangzejin3883/article/details/108904399

## Firmware

根据需求选择对应的 firmware

#### Official

原生 OpenWrt 固件，无后门安全。大部分 翻墙 插件不能直接通过 opkg 下载，或者没有 ipk。下载原生 OpenWrt 参考下面内容操作

先按照 [对照表](https://openwrt.org/toh/views/toh_fwdownload) 查看时候有对应的固件。如果没找到说明对应设备不支持 OpenWrt。如果是自己组装的小主机，大部分都支持 stable release。这里选择 All firmware images

![Snipaste_2022-12-19_00-24-29](https://cdn.staticaly.com/gh/dhay3/image-repo@master/Snipaste_2022-12-19_00-24-29.4gwnjctlnea0.webp)

按照 CPU 的指令集选择，会出现如下几个版本的 img 文件下载链接，按需选择。Supplementary Files 为校验文件

![Snipaste_2022-12-19_20-15-26](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221219/Snipaste_2022-12-19_20-15-26.63ptrl73ork0.webp)

- **ext4-combined-efi.img.gz** This disk image uses a single read-write ext4 partition without a read-only squashfs root filesystem. As a result, the root partition can be expanded to fill a large drive (e.g. SSD/SATA/mSATA/SATA DOM/NVMe/etc). Features like Failsafe Mode or Factory Reset will not be available as they need a read-only squashfs partition in order to function. It has both the boot and root partitions and Master Boot Record (MBR) area with updated GRUB2.
- **ext4-combined.img.gz** This disk image is the same as above but it is intended to be booted with PC BIOS instead of EFI.
- **ext4-rootfs.img.gz** This is a partition image of only the root partition. It can be used to install OpenWRT without overwriting the boot partition and Master Boot Record (MBR).
- **kernel.bin**
- **squashfs-combined-efi.img.gz** This disk image uses the traditional OpenWrt layout, a squashfs read-only root filesystem and a read-write partition where settings and packages you install are stored. Due to how this image is assembled, you will have less than 100MB of space to store additional packages and configuration, and extroot does not work. It supports booting from EFI.
- **squashfs-combined.img.gz** This disk image is the same as above but it is intended to be booted with PC BIOS instead of EFI. (==多数软路由 CPU 不支持 UEFI，所以选这个就好了==)
- **squashfs-rootfs.img.gz**
- **rootfs.tar.gz** This contains all the files from the root partition. It can be extracted onto a root filesystem without the need of overwriting the partition. To avoid conflicts, it is highly recommended you backup any older files and extract this file onto an empty filesystem.

这里千万不要选择 EFI 的，如果 CPU 不支持 EFI 启动的话，会导致刷机后无法正常进入系统

#### Lean's LEDE

由国人开发的原生 OpenWrt 分支。通过编译开源的代码，可以制作高度客制化的固件，支持 翻墙 插件

具体参考 [OpenWrt LEDE firmware]()

#### right

koolshare 已死，[恩山](https://www.right.com.cn/forum/forum.php) 论坛找编译好的固件。新手推荐使用 高大全

https://www.right.com.cn/forum/thread-7048868-1-1.html

## Flash firmware

如果刷机后，连接显示屏不能正常显示，需要考虑一下显示屏是否兼容。尽量使用支持 VGA 口的显示屏

### Linux live stick plus dd

> 推荐使用 ventoy + ventoy injection plugin, 这样就不需要配置网络去下载固件了

官方推荐的方案是使用 Linux live stick ( iso 可以选择 finnix ) 写  drive，按照官方文档操作即可但是不能将磁盘写满，需要手动扩展分区。参考视频指路

https://www.youtube.com/watch?v=cOLn2H1FZEI

### WinPE plus physdiskwrite

由于测试时的小主机 不支持 2K HDMI 显示，手头没有 VGA 兼容 VGA 的显示屏，所以无法使用 Linux live stick 的方式。

这里使用 WinPE 工具 + [physdiskwrite](https://m0n0.ch/wall/physdiskwrite.php)( imagedisk,rufus 都可以 ) 的方式写入，然后使用 DiskGenius 扩展分区。具体可以参考

https://blog.csdn.net/zhangzejin3883/article/details/108904399

## Extend partition

如果使用的是 WinPE plus physdiskwrite 的方式刷入固件，可以跳过此步骤

## Connection

OpenWrt 默认 eth0 配置 LAN，刷完机后将第一个网口连接主机 RJ45 口。打开网络适配器查看详细信息中的网关 IP，即刷好 OpenWrt 的机器的 eth0

![Snipaste_2022-12-19_23-57-53](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221219/Snipaste_2022-12-19_23-57-53.24hgp6dsjv34.webp)

这样我们就可以通过 192.168.1.1 来连接我们的软路由了

## Configuration

出于习惯，把将 eth0 刷成 WAN 口，eth1 刷成 LAN，同时为 eth1 设置固定 IP 以方便管理

```
#debian
vim /etc/config/network

config interface 'loopback'
        option device 'lo'
        option proto 'static'
        option ipaddr '127.0.0.1'
        option netmask '255.0.0.0'

config globals 'globals'
        option ula_prefix 'fdc5:7e22:24a6::/48'

config device
        option name 'br-lan'
        option type 'bridge'
        list ports 'eth1'


config interface 'lan'
        option device 'br-lan'
        option proto 'static'
        option ipaddr '192.168.2.100'
        option netmask '255.255.255.0'
        option ip6assign '60'

config interface 'wan'
        option device 'eth0'
        option proto 'dhcp'

config interface 'wan6'
        option device 'eth0'
        option proto 'dhcpv6'
```

这样我们就可以把软路由挂在交换机下了，让主机连接 AP，这样我们就可以访问软路由了。==这里注意和交换机互联的应该是 eth1==

