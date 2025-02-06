---
createTime: 2024-12-31 12:34
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# RAID Overview

## 0x01 Preface

RAID(Redundant Array of Independent Disks) 直译过来就是独立磁盘的冗余队列。是一种将多个物理磁盘虚拟成一个逻辑磁盘的技术，可以为数据提供

- striping - 分段
- mirroring - 镜像
- parity - 校验

## 0x02 Striping

> data striping is the technique of segmenting logically sequential data, such as a file, so that consecutive segments are stored on different physical storage devices.[^1]

striping 是指将完整的数据块，以分段的形式存储在不同的磁盘上

![300x400](https://upload.wikimedia.org/wikipedia/commons/c/cf/Data_striping_example.svg)

例如一个文件 A 被拆分成 4 段 segments，A1/A3/A5/A7 存储在 Disk0，A2/A4/A6/A8 存储在 Disk1

这么做有好处

1. 系统可以并发读取或者写入分段的数据，提高 I/O 性能
2. 为磁盘提供 I/O 负载均衡

但是也有缺点

1. 如果任意一块磁盘出现问题，就会导致原先完整的数据 corruption

根据 striping 的最小单元分为

- bit-striping - 数据以 bit 的形式分段，即最小的单元为 bit
- byte-striping - 数据以 byte 的形式分段，即最小的单元为 byte
- block-striping - 数据以 block 的形式分段，即最小的单元为 block(通常为 64KB/128KB)

> [!note]
> striping 提高 I/O 性能

## 0x03 Mirroring

> disk mirroring is the replication of logical disk volumes onto separate physical hard disks in real time to ensure continuous availability.[^2]

mirroring 是指将数据以镜像(复制)的形式，存储在不同的磁盘上

![](https://upload.wikimedia.org/wikipedia/commons/0/05/Raid1_Version_for_Wiki.jpg)

例如一个文件 A 被拆分成 4 段 segments， A1/A2/A3/A4 存储在 Disk0，同时也会将 A1/A2/A3/A4 镜像到 Disk1

这么做有好处

1. 如果任意一块磁盘出现问题，仍可以获取数据

但是也有缺点

1. 多块磁盘用作一块，磁盘容量使用率低

> [!note] 
> mirroring 提供数据容灾

## 0x04 Parity

> A parity bit, or check bit, is a bit added to a string of binary code. Parity bits are a simple form of error detecting code.[^3]

parity 中文译为奇偶性，根据 transfer data 1 的个数奇偶性会在数据结尾增加一位 bit 用于校验数据在传输过程中是否出现错包(比 CRC 简单)

如果 transfer data 1 的个数为奇数(odd) 则校验位为 1

如果 transfer data 1 的个数为偶数(even) 则校验位为 0

例如

`data = 1101` 计算 1 的个数为 3 奇数，所以校验位为 1，即 `transfer data = 11011`

如果 `recive data = 11010`，计算 1 的个数为 3 奇数，但是校验位为 0，说明传输过程中出现错包，反之说明没有错包

但是 RAID 中 parity 不仅仅提供错误校验也会提供数据容灾。RAID 会对磁盘上不同的数据执行 XOR 计算 parity

例如

SSD1 上有数据 101101

SSD2 上有数据 110001

$parity = XOR(101101,110001) = 011100$

如果任意一块磁盘上数据出现丢失，可以通过 parity 计算出原先的数据

假设 SSD2 上 110001 数据丢失，可以通过 parity 对 SSD2 上的数据 110001 执行 XOR 还原 $XOR(011100,101101)=110001$

这么做有好处

1. 如果一块磁盘部分数据出现问题，可以复原数据

但是也有缺点

1. 写入或者更新数据需要额外的 write I/O 开销(因为需要计算 parity 并写入)
2. parity 需要额外的存储空间

> [!note] 
> parity 提供轻量数据容灾

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [RAID - Wikipedia](https://en.wikipedia.org/wiki/RAID)

***References***

[^1]:[Data striping - Wikipedia](https://en.wikipedia.org/wiki/Data_striping)
[^2]:[Disk mirroring - Wikipedia](https://en.wikipedia.org/wiki/Disk_mirroring)
[^3]:[Parity bit - Wikipedia](https://en.wikipedia.org/wiki/Parity_bit)
