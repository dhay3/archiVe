---
createTime: 2025-03-05 13:14
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID 05 - Common Terms

## 0x01 Preface

RAID 管理中还有一些常用的术语

## 0x02 PD

PD(Physical Drives)

物理磁盘，用于构建 VD 或者直接使用

## 0x03 VD

VD(Virtual Disks)

虚拟磁盘，由多块或者是单块 PD 组成（可以将其看作 RAID/JBOD 的产物）

## 0x04 JBOD

JBOD(Just a Bunch Of Drives)

磁盘的状态，表示当前磁盘只是独立的磁盘，或者是非 RAID VD 中的一块

> JBOD (just a bunch of disks or just a bunch of drives) is an architecture using multiple hard drives exposed as individual devices. Hard drives may be treated independently or may be combined into one or more logical volumes[^1]

JBOD 不具备 striping/mirroring/parity 的功能，只是将单块或者是多块磁盘组成一块 VD（容量增加，顺序写，性能按照磁盘），可以将其理解成 LVM

## 0x UG

UG(Unconfigured Good)

磁盘的状态，表示当前磁盘没有任何配置，可以配置 RAID（只有是 UG 的磁盘，才可以组 RAID）

## 0x Optl

Optl(Optimal)

## 0x Onln



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Non-RAID drive architectures - Wikipedia](https://en.wikipedia.org/wiki/Non-RAID_drive_architectures)


***References***

[^1]:[Non-RAID drive architectures - Wikipedia](https://en.wikipedia.org/wiki/Non-RAID_drive_architectures)



