#  BIOS & UEFI

参考:

https://www.linuxdashen.com/linux%E7%94%A8%E6%88%B7%E7%9A%84uefi%E5%9B%BA%E4%BB%B6%E6%8C%87%E5%8D%97

https://en.wikipedia.org/wiki/BIOS

https://www.jianshu.com/p/32adf21317cd

https://www.freecodecamp.org/news/uefi-vs-bios/

https://www.computernetworkingnotes.com/linux-tutorials/linux-disk-management-tutorial.html

## Digest

当计算机插电启动时， BIOS 或者是 EFI 作为第一个启动的进程，主要负责检查校验系统硬件的完整性，例如 CPU, Memory, Keyboad, Hard disk 以及其关联的周边设备（peripherals）。可以简称 Power On Self Test 即 POST。当设备校验通过会启动 Boot-loader 载入系统

BIOS 和 EFI 也被称为固件，因为其程序直接被刷在 motherboad 上 ( 实际为 ROM chip )。系统或者主板不同 BIOS 或者 UEFI 的固件也不一样，互不兼容

## Boot Procedure

正如上节说的 BIOS/UEFI 会负责校验系统的硬件已经引导 Boot-loader，主要流程如下

1. 初始化 CPU 和 RAM
2. 开机自检，Power On Self Test (POST)
3. 初始化 LAN，PCLe 接口
4. 启动 MBR 分区的 Boot-loader (例如GRUB)，或是 USB 和网络上的 Boot-loader
5. Boot-loader 加载 kernel 到内存中，windows 中加载 wininit.exe

当完成以上任务后，BIOS/UEFI 将控制权交给操作系统。

## BIOS

Basic Input/Output System 简称 BIOS 也被称为 Legacy BIOS

![Snipaste_2021-03-10_11-35-45](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230214/Snipaste_2021-03-10_11-35-45.36v0g0rzwvpc.webp)

由 16 汇编语言开发

BIOS 由于是汇编语言编写的，所以被刷在了固件上，也就不能更改。运行后存储在EPROM (Erasable Programmable Read-Only Memory)



## EFI/UEFI

Unified Extensible Firmware Interface 简称 UEFI, 前身是 Extensible Firmware Interface 简称 EFI

Legacy BIOS 最初是用 16 bit 汇编语言编写的，随着 CPU 架构的升级到 64 bit，使用 BIOS 就不能充分的发挥 64 bit 的处理器。所以 Intel 公司就开发了一种新的固件 EFI  用来替代 BIOS ( BIOS )。后来在演变中变成通用的规则，也就是 UEFI

UEFI的所有信息存储在以`.efi`文件中(binary)，而不是存储在固件上，可以通过UEFI Shell来自定义引导过程。

- linux

![Snipaste_2021-03-10_11-16-44](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230214/Snipaste_2021-03-10_11-16-44.47njszmw0dc0.webp)

- windows

  ![Snipaste_2021-03-10_11-17-47](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230214/Snipaste_2021-03-10_11-17-47.3twdhkkpcncw.webp)

`.efi`文件存储在一个特殊的分区EFI System Partition(ESP)

## BIOS vs UEFI

1. BIOS 是用汇编语言写的存储在固件上，UEFI 是用 C 语言写的存储在`.efi`文件中
2. BIOS 只支持 MBR 分区表即最大支持 2TB 分区，UEFI 支持 MBR/GPT 分区表最大支持 9ZB  
3. 使用 Driver/protocal 替换中断硬件端口的操作方法
4. UEFI启动需要一个独立的分区，它可以将系统启动文件和操作系统隔离
5. UEFI减少了自检，加快的启动时间。同时提供安全引导
6. UEFI 运行32bits 或 64bits，运行 BIOS 16bits。大大加快了运行速度，同时支持图形化界面

## Boot mode

> 可以在启动过程中使用F2，F10，F12在Boot tab中选择



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

  | 分区表 | 分区位置 | 分区大小  | 分区类型   | 挂载点               | 文件系统 |
  | ------ | -------- | --------- | ---------- | -------------------- | -------- |
  | GPT    | 任意     | 200MB-1GB | EFI System | `/boot/efi`或`/boot` | fat32    |

  









