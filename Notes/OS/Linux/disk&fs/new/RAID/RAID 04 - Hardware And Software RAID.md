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

所有的 RAID 操作通过硬件完成，这个硬件被称为 RAID Card（独立于服务器或者是内嵌在 motherboard 上，可以看作一台 micro-server，有自己的 processor, memory，甚至有些还有自己的 battery），核心是 RAID Controller 是实际完成 RAID 操作的芯片

RAID Card 大概长这样

![](https://www.fujitsu.com/global/Images/W-RAID-Ctrl-SAS-6Gb-1GB-D3116_tcm100-165887.png)

## 0x03 Software RAID

所有的 RAID 操作通过软件完成，基于操作系统，Linux 上通常使用 `mdadm`

## 0x04 Hardware RAID VS Software RAID

### 0x04 Function

- Hardware RAID 通常支持 Hotspare failover(热备故障转移)，无需关机就可以自动替换 RAID 中有故障的磁盘。而 Software RAID 必须要关机才可以替换 RAID 中有故障的磁盘

### 0x04 Reliability

- 因为有自己的 battery，所以当服务器异常关机时，RAID card memory 中正在写入的数据不会丢失，会继续写入
- 因为有自己的 processor 和 memory，所以当系统 panic 时，RAID card memory 中正在写入的数据不会丢失，会继续写入

### 0x04 Performace

- Hardware RAID （read/write I/O）性能比 Software RAID 高，因为 Hardware RAID 有自己的 processor 和 memory 不占用服务器的资源（Software RAID 间接增加系统资源的开销）
- 在系统负载高的情况下，Software RAID 操作会受影响，而 Hardware RAID 不会

### 0x04 Compatibility

- 当服务器更换操作系统时，Hardware RAID 几乎不用动，而 Software RAID 需要重新配置
- 当 Hardware RAID 配置需要迁移到不同硬件平台上时，几乎不太可能，而 Software RAID 较为容易

### 0x04 Fee

- Software RAID 没有硬件成本，而 Hardware RAID 有硬件以及维护成本

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Understanding Hardware RAID: Comprehensive Guide to RAID Levels and Benefits \| DiskInternals](https://www.diskinternals.com/raid-recovery/understanding-hardware-raid/)
- [Hardware RAID - GeeksforGeeks](https://www.geeksforgeeks.org/hardware-raid/)
- [What is hardware RAID (hardware redundant array of independent disk)? \| Definition from TechTarget](https://www.techtarget.com/searchstorage/definition/hardware-RAID-hardware-redundant-array-of-independent-disk)

***References***
