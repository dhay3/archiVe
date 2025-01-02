---
createTime: 2024-12-31 12:34
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# RAID

## 0x01 Preface

RAID(Redundant Array of Independent Disks) 直译过来就是独立磁盘的冗余队列。是一种将多个物理磁盘虚拟成一个逻辑磁盘的技术，可以为数据提供

- striping - 分段
- mirroring - 镜像
- parity - 

## 0x02 Striping

> data striping is the technique of segmenting logically sequential data, such as a file, so that consecutive segments are stored on different physical storage devices.

striping 是指将完整的数据块，以分段的形式存储在不同的磁盘上

![300x400](https://upload.wikimedia.org/wikipedia/commons/c/cf/Data_striping_example.svg)

例如一个文件 A 被拆分成 8 段 segments，A1/A3/A5/A7 存储在 Disk0，A2/A4/A6/A8 存储在 Disk1

这么做有好处

1. 系统可以并发读取分段的数据，I/O 高
2. 为磁盘提供 I/O 负载均衡

但是也有缺点

1. 如果任意一块磁盘出现问题，就会导致原先完整的数据 corruption

## 0x03 Mirroring

> disk mirroring is the replication of logical disk volumes onto separate physical hard disks in real time to ensure continuous availability.

mirroring 是指将数据以镜像(复制)的形式，存储在不同的磁盘上

![](https://upload.wikimedia.org/wikipedia/commons/0/05/Raid1_Version_for_Wiki.jpg)

例如一个文件 A 被拆分成 4 段 segments， A1/A2/A3/A4 存储在 Disk0，同时也会将 A1/A2/A3/A4 镜像到 Disk1

这么做有好处

1. 如果任意一块磁盘出现问题，仍可以获取数据。即数据容灾

但是也有缺点

1. 多块磁盘用作一块，磁盘容量使用率低，成本也高

## 0x04 Parity

> A parity bit, or check bit, is a bit added to a string of binary code. Parity bits are a simple form of error detecting code.

parity 是指数据错误校验，根据 transfer data 1 的个数奇偶性会在数据结尾增加一位 bit 用于校验数据在传输过程中是否出现错包 

如果 transfer data 1 的个数为奇数(odd) 则校验位为 1

如果 transfer data 1 的个数为偶数(even) 则校验位为 0

例如

`data = 1101` 计算 1 的个数为 3 奇数，所以校验位为 1，即 `transfer data = 11011`

如果 `recive data = 11010`，计算 1 的个数为 3 奇数，但是校验位为 0，说明传输过程中出现错包，反之说明没有错包

## 0x05 RAID LEVEL

RAID 根据是否提供 striping/mirroring/parity 将其划分为 7 个 level

### 0x05a RAID 0

提供 striping，不提供 mirroring/parity

![300x400](https://upload.wikimedia.org/wikipedia/commons/9/9b/RAID_0.svg)

- n 块磁盘组成的 RAID0 提供 n 倍的 I/O 性能
- 没有数据容灾和错误校验
- RAID0 的磁盘大小可以不一致，但是 RAID0 的容量为 $min(disk) \times n$。例如两块磁盘分别为 120GB 和 320GB，组成 RAID0 其容量为 $120GB \times 2 = 240GB$

### RAID 1

提供 mirroring，不提供 stripping/parity

![300x400](https://upload.wikimedia.org/wikipedia/commons/b/b7/RAID_1.svg)

- 

### RAID 2

### RAID 3

### RAID 4

### RAID 5

### RAID 6

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [RAID - Wikipedia](https://en.wikipedia.org/wiki/RAID)
- [Standard RAID levels - Wikipedia](https://en.wikipedia.org/wiki/Standard_RAID_levels)

***References***


