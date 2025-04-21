---
createTime: 2025-02-06 14:48
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID 04 - Hardware And Software RAID

## 0x01 Preface

根据 RAID 管理的方式可以将其分为 Hardware RAID 和 Software RAID

- Hardware RAID - 通过硬件来管理 RAID
- Software RAID - 通过软件来管理 RAID

## 0x02 Hardware RAID

Hardware RAID 是指所有 RAID 的操作都通过硬件完成，这个硬件被统称为 RAID Card，通常独立于服务器或者是内嵌在 motherboard 上。可以将 RAID Card 看作一台 micro-server，有自己的 processor, memory，甚至有些还有自己的 battery

RAID Card 大概长这样

![](https://www.fujitsu.com/global/Images/W-RAID-Ctrl-SAS-6Gb-1GB-D3116_tcm100-165887.png)

其中核心是 RAID Controller，是实际完成 RAID 操作的芯片

> [!important]
> 由于 RAID 这项技术早于 NVME，所以一些早期或者是低端的 RAID Card 只支持 SATA/SCSI/SAS 接口的磁盘，并不支持 M2/NVME 接口的磁盘(这时通常需要通过 Software RAID 来配置 M2/NVME RAID)

### 0x02a MegaRAID

MegaRAID 是 Boardcom 出品的一系列 RAID Card，其特殊在可以通过 `storcli` （Storage Command Line Tool）命令来直接和 RAID Controller 进行通信，所以也归属为 Hardware RAID 但是可以通过 Software 来配置 RAID

### 0x02b Common RAID Card

#### LSI SAS3108[^1]

- 支持 RAID 0,1,5,6,10,50,60
- 支持 SAS/SATA/PCIe3.0 接口

#### LSI SAS3508[^2]

- 支持 RAID 0,1,5,6,10,50,60
- 

## 0x03 Software RAID

所有 RAID 的操作通过软件完成，基于操作系统，Linux 上通常使用 `mdadm`

## 0x04 Hardware RAID VS Software RAID

### 0x04a Function

- Hardware RAID 通常支持 Hotspare failover(热备故障转移)，无需关机就可以自动替换 RAID 中有故障的磁盘。而 Software RAID 必须要关机才可以替换 RAID 中有故障的磁盘

### 0x04b Reliability

- Hardware RAID 有自己的 battery，所以当服务器异常关机时，RAID card memory 中正在写入的数据不会丢失，会继续写入
- Hardware RAID 有自己的 processor 和 memory，所以当系统 panic 时，RAID card memory 中正在写入的数据不会丢失，会继续写入

### 0x04c Performace

- Hardware RAID （read/write I/O）性能比 Software RAID 高，因为 Hardware RAID 有自己的 processor 和 memory 不占用服务器的资源（Software RAID 间接增加系统资源的开销）

### 0x04d Compatibility

- 当服务器更换操作系统时，Hardware RAID 几乎不用动，而 Software RAID 需要重新配置
- 当在不同硬件平台上迁移 RAID 配置时，如果 RAID Card 不同 Hardware RAID 几乎无法迁移，而 Software RAID 较为容易
- Software RAID 没有磁盘接口的限制，而 Hardware RAID 根据 RAID Card 来判断有无限制

### 0x04e Fee

- Software RAID 没有硬件成本，而 Hardware RAID 有硬件以及维护成本

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Understanding Hardware RAID: Comprehensive Guide to RAID Levels and Benefits \| DiskInternals](https://www.diskinternals.com/raid-recovery/understanding-hardware-raid/)
- [Hardware RAID - GeeksforGeeks](https://www.geeksforgeeks.org/hardware-raid/)
- [What is hardware RAID (hardware redundant array of independent disk)? \| Definition from TechTarget](https://www.techtarget.com/searchstorage/definition/hardware-RAID-hardware-redundant-array-of-independent-disk)


***References***

[^1]:[LSISAS3108](https://docs.broadcom.com/doc/LSISAS3108)
