# UEFI & BIOS

参考:

https://www.linuxdashen.com/linux%E7%94%A8%E6%88%B7%E7%9A%84uefi%E5%9B%BA%E4%BB%B6%E6%8C%87%E5%8D%97

https://blog.csdn.net/F8qG7f9YD02Pe/article/details/79551663

https://www.jianshu.com/p/32adf21317cd

https://www.freecodecamp.org/news/uefi-vs-bios/

## 概述

bios和efi是固件接口标准，操作系统通过它来完成引导。

## BIOS

Basic Input/Output System(BIOS，通常也叫做Legacy BIOS)

**启动流程**

1. 初始化CPU和RAM
2. 开机自检，power on self test (POST)
3. 初始化LAN，PCLe(固态)接口
4. 启动MBR分区的boot loader(例如GRUB)，或是USB和网络上的boot loader
5. ==boot loader加载kernel到内存中==，windows中加载wininit.exe

当完成以上任务后，bios将控制权交给操作系统。

bios由于是汇编语言编写的，所以被刷在了固件上，也就不能更改。运行后存储在EPROM (Erasable Programmable Read-Only Memory)

## EFI/UEFI

Unified Extensible Firmware Interface(UEFI，统一可扩展固件接口)，前生是Extensible Firmware Interface (EFI)

**启动流程**

1. 初始化CPU和RAM
2. 初始化LAN，PCLe(固态)接口
3. 启动MBR分区的boot loader，或是USB和网络上的boot loader
4. ==boot loader加载kernel到内存中==，windows中加载wininit.exe

当完成以上任务后，uefi将控制权交给操作系统。

UEFI的所有信息存储在以`.efi`文件中(binary)，而不是存储在固件上，可以通过UEFI Shell来自定义引导过程。

- linux

![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-10_11-16-44.png)



- windows

  ![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-10_11-17-47.png)

`.efi`文件存储在一个特殊的分区EFI System Partition(ESP)

## BIOS vs UEFI

1. BIOS是用汇编语言写的存储在固件上，UEFI是用C语言写的存储在`.efi`文件中
2. BIOS使用的分区表为MBR，==UEFI使用的分区表是GPT(说明支持2T硬盘)==
3. 使用Driver/protocal替换中断硬件端口的操作方法
4. UEFI启动需要一个独立的分区，它可以将系统启动文件和操作系统隔离
5. UEFI减少了自检，加快的启动时间。同时提供安全引导
6. UEFI运行32bits或64bits，bios运行16bits。大大加快了运行速度，同时支持图形化界面

## 启动方式

> 可以在启动过程中使用F2，F10，F12在Boot tab中选择

![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-10_11-35-45.png)

- Legacy mode

  通过MBR/BIOS引导

- UEFI mode

  通过GPT/UEFI引导

- CSM mode

  兼容模式

## 查看引导方式

- windows

  使用命令`msinfo32`来查看，如果在BIOS模式一栏是UEFI就是uefi方式启动，如果是legacy就是普通bios启动。

- linux

  查找`/sys/firmware/efi`文件。如果使用BIOS就没有该文件

  ![img](https://ss.csdn.net/p?https://mmbiz.qpic.cn/mmbiz_png/W9DqKgFsc6ibHJT2OmUdcfSvXr2icU8tDrx7jHhAkM18ib0RAkicxpTIiaURU4X5hpMs330vbbYlgsNhcRRrSvSK46Q/640?wx_fmt=png)
  
  `fw_platform_size`是efi使用的系统位数

## 引导分区

https://docs.voidlinux.org/installation/live-images/partitions.html

以下使用GRUB作为boot loader

- BIOS

  | 分区表 | 分区位置 | 分区大小 | 分区类型  | 挂载点 |
  | ------ | -------- | -------- | --------- | ------ |
  | MBR    | 第一     | 1MB      | BIOS Boot | 无     |

- UEFI

  | 分区表 | 分区位置 | 分区大小  | 分区类型   | 挂载点      |
  | ------ | -------- | --------- | ---------- | ----------- |
  | GPT    | 任意     | 200MB-1GB | EFI System | `/boot/efi` |

  









