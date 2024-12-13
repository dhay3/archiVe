ref：

[https://www.gnu.org/software/grub/manual/grub/grub.html#Overview](https://www.gnu.org/software/grub/manual/grub/grub.html#Overview)
## Overview
GRUB (GRand Unified Bootloader)最早是由Erich Boleyn于1995年写的也被叫做GRUB Legacy，1999年正式被加入到GNU package。在2002年，Yoshinori K. Okuji 在GRUB的基础上创建了PUPA(Preliminary Universal Programming Architecture for GNU GRUB)，后来被命名为GRUB2，也就是现在被大家称为的GRUB。现在大多数UNIX系统都已经采用GRUB做为boot loader。

大家都知道在电脑启动时启动的第一个软件是【boot loader】—— 负责加载kernel将操作系统的控制权转交给kernel，然后由kernel初始化操作系统。GRUB 正是一个 可以引导不同OS kernl 的 boot loader，当使用 GRUB 时可以通过 CLI 或者 menu interface 来引导 kernel

## GRUB2 VS GRUB
现在的大家收熟知的GRUB一般为GRUB2，和 legacy GRUB有以下几种不同

1. 配置文件以`grub.cfg`命名而不是`menu.list`或`grub.conf`，使用 new syntax ，增加了许多 new commands
1. `grub.cfg`由`grub-mkconfig`命令自动生成
1. 在GRUB2中device的分区，以 1 开始，而不是 0
1. GRUB2可以直接从 LVM(logical volume)和 RAID 中读取
## Features
> 只针对boot这一阶段


1. 能识别多种可执行文件，ELF, a.out, sysmbol tables
1. 支持【multiple boot】
1. 使用配置文件预设boot
1. 提供menu interface 展示preset boot commands
1. 灵活的命令行接口，可以修改或添加preset boot commnads
1. 支持多filesystem types。例如FAT32, NTFS
1. 支持automatic decompression，自动解压kernel中的`gzip`或`xz`文件
1. 支持从软盘或硬盘中读取能被BIOS识别的数据
1. 支持探测RAM
1. 如果使用Legacy BIOS，BIOS分区必须在1024 cylinders内。如果磁盘Logical Block Address(逻辑地址寻址)，【GRUB就可以将boot loader需要加载的内容放在磁盘上的任意一个位置】。
1. 支持从网络加载bootloader，使用TFTP协议
1. 支持从远程终端连接GRUB

## 

