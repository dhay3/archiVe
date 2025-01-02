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
- parity - 校验

## 0x02 Striping

> data striping is the technique of segmenting logically sequential data, such as a file, so that consecutive segments are stored on different physical storage devices.[^1]

striping 是指将完整的数据块，以分段的形式存储在不同的磁盘上

![300x400](https://upload.wikimedia.org/wikipedia/commons/c/cf/Data_striping_example.svg)

例如一个文件 A 被拆分成 8 段 segments，A1/A3/A5/A7 存储在 Disk0，A2/A4/A6/A8 存储在 Disk1

这么做有好处

1. 系统可以并发读取或者写入分段的数据，为磁盘提供 I/O 负载均衡

但是也有缺点

1. 如果任意一块磁盘出现问题，就会导致原先完整的数据 corruption

> [!NOTE]
> stripping 提高 I/O 性能

## 0x03 Mirroring

> disk mirroring is the replication of logical disk volumes onto separate physical hard disks in real time to ensure continuous availability.[^2]

mirroring 是指将数据以镜像(复制)的形式，存储在不同的磁盘上

![](https://upload.wikimedia.org/wikipedia/commons/0/05/Raid1_Version_for_Wiki.jpg)

例如一个文件 A 被拆分成 4 段 segments， A1/A2/A3/A4 存储在 Disk0，同时也会将 A1/A2/A3/A4 镜像到 Disk1

这么做有好处

1. 如果任意一块磁盘出现问题，仍可以获取数据

但是也有缺点

1. 多块磁盘用作一块，磁盘容量使用率低

> [!NOTE] 
> mirroring 提供数据容灾

## 0x04 Parity

> A parity bit, or check bit, is a bit added to a string of binary code. Parity bits are a simple form of error detecting code.[^3]

parity 是指数据错误校验，根据 transfer data 1 的个数奇偶性会在数据结尾增加一位 bit 用于校验数据在传输过程中是否出现错包 

如果 transfer data 1 的个数为奇数(odd) 则校验位为 1

如果 transfer data 1 的个数为偶数(even) 则校验位为 0

例如

`data = 1101` 计算 1 的个数为 3 奇数，所以校验位为 1，即 `transfer data = 11011`

如果 `recive data = 11010`，计算 1 的个数为 3 奇数，但是校验位为 0，说明传输过程中出现错包，反之说明没有错包

| 7 bits of data | (count of 1-bits) | 8 bits including parity |              |
| -------------- | ----------------- | ----------------------- | ------------ |
| even           | odd               |                         |              |
| 0000000        | 0                 | 0000000**0**            | 0000000**1** |
| 1010001        | 3                 | 1010001**1**            | 1010001**0** |
| 1101001        | 4                 | 1101001**0**            | 1101001**1** |
| 1111111        | 7                 | 1111111**1**            | 1111111**0** |

这么做有好处

1. 保证了数据出现错误时能快速重传

但是也有缺点

1. I/O 性能相对降低

> [!NOTE] 
> parity 提供数据错误校验

## 0x05 RAID LEVEL

RAID 根据是否提供 striping/mirroring/parity 将其划分为 7 个 level

### 0x05a RAID 0[^1]

提供 striping，不提供 mirroring/parity

![300x400](https://upload.wikimedia.org/wikipedia/commons/9/9b/RAID_0.svg)

**Performance**

- n 块 read/write I/O 性能相同的磁盘组成 RAID 0，提供 n 倍的 read/write I/O 性能
- n 块 read/write I/O 性能不同的磁盘组成 RAID 0，提供 $n \times min(write I/O)$ write I/O 性能，$n \times min(read I/O)$ read I/O 性能

**Capacity**

- n 块容量相同的磁盘组成 RAID 0，提供 n 倍的容量
- n 块容量不同的磁盘组成 RAID 0，提供 n 倍的容量，但是只有 $n \times min(capacity)$ 的容量提供 stripping，剩余的容量只用于存储

假设 

SSD1 120GB 500MB/s read/write

SSD2 340GB 300MB/s read/write

那么组成 RAID 0 后，read I/O 为 600MB/s，write I/O 为 600MB/s，240GB 提供 stripping，剩余的 120GB 只提供存储

> [!NOTE] Choosing Criteria
> 如果需要磁盘性能优先可以选择 RAID 0

### RAID 1

提供 mirroring，不提供 stripping/parity

![300x400](https://upload.wikimedia.org/wikipedia/commons/b/b7/RAID_1.svg)

**Performance**

- n 块磁盘组成的 RAID 1 write I/O 为随机使用的磁盘 write I/O

**Capacity**

- n 块磁盘组成的 RAID 1 提供 $sum(n disk read I/O)$

假设


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
- [Understanding RAID 0: Benefits, Risks, and Applications \| DiskInternals](https://www.diskinternals.com/raid-recovery/understanding-raid-0/)

***References***

[^1]:[Data striping - Wikipedia](https://en.wikipedia.org/wiki/Data_striping)
[^2]:[Disk mirroring - Wikipedia](https://en.wikipedia.org/wiki/Disk_mirroring)
[^3]:[Parity bit - Wikipedia](https://en.wikipedia.org/wiki/Parity_bit)
