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

例如一个文件 A 被拆分成 4 段 segments，A1/A3/A5/A7 存储在 Disk0，A2/A4/A6/A8 存储在 Disk1

这么做有好处

1. 系统可以并发读取或者写入分段的数据，为磁盘提供 I/O 负载均衡

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

> [!note] 
> parity 提供轻量数据容灾

## 0x05 RAID LEVEL

> [!important]
> 例子中均为理论值，实际要比这个小，因为 I/O 快的磁盘需要等 I/O 慢的磁盘完成 read 或者是 write

RAID 根据是否提供 striping（包括 striping 的最小单元）/mirroring/parity(包括 parity 数据存储的方式)，将其划分为 7 个 level

### 0x05a RAID 0

提供 striping，不提供 mirroring/parity

![300x400](https://upload.wikimedia.org/wikipedia/commons/9/9b/RAID_0.svg)

使用 block-striping

**Performance**

- n 块 read/write I/O 性能相同(或者不相同)的磁盘组成 RAID 0，提供 $n \times min(read\ I/O)$ read I/O 性能，$n \times min(write\ I/O)$ write I/O 性能

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 0，提供 $sum(capacity)$ 容量，只有 $n \times min(capacity)$ 的容量提供 striping，剩余的容量只用于存储

假设 

SSD1 120GB 500MB/s read/write I/O

SSD2 340GB 300MB/s read/write I/O

那么组成 RAID 0 后，read I/O 为 600MB/s，write I/O 为 600MB/s，SSD1 SSD2 各提供 120GB 用于 striping，剩余的 120GB 只提供存储

> [!note] Choosing Criteria
> 如果需要最好的 read/write I/O 性能优先选择 RAID 0

### 0x05b RAID 1

提供 mirroring，不提供 striping/parity

![300x400](https://upload.wikimedia.org/wikipedia/commons/b/b7/RAID_1.svg)

**Performance**

- n 块 read I/O 性能相同(或者不相同)的磁盘组成的 RAID 1，提供大于 $min(read\ I/O)$ 但是小于 $max(read\ I/O)$ read I/O 性能（read request 会被 RAID controller 发送到不同的磁盘上，因为磁盘上的数据都相同）
- n 块 write I/O 性能相同(或者不相同)的磁盘组成的 RAID 1，wirte I/O 为 $min(write\ I/O)$

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 1，提供 $sum(capacity)$ 容量，只有 $min(capacity)$ 的容量提供 mirroring，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O
   
SSD2 300GB 300MB/s read/write I/O

SSD3 400GB 400MB/s read/write I/O

那么组成 RAID 1 后，read I/O 在 100MB/s 至 400MB/s 之间，write I/O 为 100MB/s，100GB 提供 mirroring，剩余的 700GB 只提供存储

> [!note] Choosing Criteria
> 如果需要最好的数据高可用优先选择 RAID 1

### 0x05c RAID 2

提供 striping/parity，不提供 mirroring

![800x400](https://upload.wikimedia.org/wikipedia/commons/b/b5/RAID2_arch.svg)

使用 bit-level striping，因为以 1 bit 为一个单元，大量写入或者修改数据，需要非常频繁写入或者更新 parity 数据，所以综合 write I/O 性能非常差，有可能不如 RAID 1；同时使用复杂的 hamming code 用于错误校验，write I/O 性能进一步降低。几乎没有设备会使用 RAID 2，所以这里就不做过多的介绍

### 0x05d RAID 3

提供 striping/parity，不提供 mirroring

![600x400](https://upload.wikimedia.org/wikipedia/commons/f/f9/RAID_3.svg)

使用 byte-level striping，因为以 1 byte 为一个单元，大量写入或者是修改数据，需要频繁写入或者更新 parity 数据，所以综合 write I/O 性能较差，相对 RAID 2 高；parity 的数据单独存储在一块磁盘上，允许有一块坏盘(non-parity disk 或者是 parity disk 都可以)；至少 3 块磁盘(至少 2 块盘才能 striping，额外任意 1 块盘存储 parity data)

**Performance**

- n 块 read I/O 性能相同(或者不相同)的磁盘组成的 RAID 3，提供 $(n - 1) \times min(\mbox{non-parity\ disk\ read\ I/O})$ read I/O 性能
- n 块 write I/O 性能相同(或者不相同)的磁盘组成 RAID 3，提供 $(n - 1) \times min(\mbox{non-parity\ disk\ write\ I/O}) - (parity\ write\ I/O\ overhead)$ write I/O 性能，由于 byte-level striping，实际远比这个小的多

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 3，提供 $sum(capacity)$ 容量，只有 $(n - 1) \times min(\mbox{non-parity\ disk\ capacity})$ 的容量提供 striping/parity， $min(\mbox{non-parity\ disk\ capacity})$ 的容量存储 parity 数据，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 300GB 300MB/s read/write I/O

SSD3 400GB 400MB/s read/write I/O 用于 parity

那么组成 RAID 3 后，read I/O 为 200MB/s，write I/O 为 $\lt 200MB/s$，SSD1 SSD2 各提供 100GB striping/parity，SSD3 提供 100GB 存储 parity，剩余的 500GB 只提供存储

> [!note] Choosing Criteria
> 通常不会使用 RAID 3，使用 RAID 5 替代

### 0x05e RAID 4

提供 striping/parity，不提供 mirroring

![600x400](https://upload.wikimedia.org/wikipedia/commons/a/ad/RAID_4.svg)

使用 block-level striping，以 1 block(通常为 64KB/128KB) 为一个单元，大量写入或者修改数据，write I/O 相对 RAID 3 高；parity 的数据单独存储在一块磁盘上，允许有一块坏盘(non-parity disk 或者是 parity disk 都可以)；至少 3 块磁盘(至少 2 块盘才能 striping，额外任意 1 块盘存储 parity data)

**Performance**

- n 块 read I/O 性能相同(或者不相同)的磁盘组成的 RAID 4，提供 $(n - 1) \times min(\mbox{non-parity\ disk\ read\ I/O})$ read I/O 性能
- n 块 write I/O 性能相同(或者不相同)的磁盘组成 RAID 4，提供 $(n - 1) \times min(\mbox{non-parity\ disk\ write\ I/O}) - (parity\ write\ I/O\ overhead)$ write I/O 性能

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 4，提供 $sum(capacity)$ 容量，只有 $(n - 1) \times min(\mbox{non-parity\ disk\ capacity})$ 的容量提供 striping/parity， $min(\mbox{non-parity\ disk\ capacity})$ 的容量存储 parity 数据，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 300GB 300MB/s read/write I/O

SSD3 400GB 400MB/s read/write I/O 用于 parity

那么组成 RAID 4 后，read I/O 为 200MB/s，write I/O 为 $\lt 200MB/s$，SSD1 SSD2 各提供 100GB striping/parity，SSD3 提供 100GB 存储 parity，剩余的 500GB 只提供存储

> [!note] Choosing Criteria
> 通常不会使用 RAID 4，使用 RAID 5 替代

### 0x05f RAID 5

提供 striping/parity，不提供 mirroring

![600x400](https://upload.wikimedia.org/wikipedia/commons/6/64/RAID_5.svg)

使用 block-level striping；parity 的数据分布存储在不同磁盘上，允许有一块坏盘(non-parity disk 或者是 parity disk 都可以)；至少 3 块磁盘(至少 2 块盘才能 striping，额外任意 1 块盘存储 parity data)

数据高可用性相对 RAID 1 较低，read I/O 相较 RAID 3/4 高，write I/O 相较 RAID 3 高

**Performance**

- n 块 read I/O 性能相同(或者不相同)的磁盘组成的 RAID 5，提供 $(n - 1) \times min(\mbox{non-parity\ disk\ read\ I/O})$ read I/O 性能
- n 块 write I/O 性能相同(或者不相同)的磁盘组成 RAID 5，提供 $(n - 1) \times min(\mbox{non-parity\ disk\ write\ I/O}) - (parity\ write\ I/O\ overhead)$ write I/O 性能

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 5，提供 $sum(capacity)$ 容量，只有 $(n - 1) \times min(capacity)$ 的容量提供 striping/parity， $min(capacity)$ 的容量存储 parity 数据，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 300GB 300MB/s read/write I/O

SSD3 400GB 400MB/s read/write I/O

那么组成 RAID 5 后，read I/O 大于 200MB/s（当 SSD3 存储 parity） 小于 600MB/s（当 SSD1 存储 parity），write I/O 为 $\lt 300MB/s$，SSD1 SSD2 SSD3 加起来可以提供 200GB 用于 striping/parity，SSD1 SSD2 SSD3 加起来可以提供100GB 用于存储 parity，剩余 500GB 只提供存储

> [!note] Choosing Criteria
> 需要较好的 read/write I/O 性能，不要求数据完全高可用优先选择 RAID 5(可以看作是 RAID 0 和 RIAD 1 的折中)

### 0x05g RAID 6

提供 striping/parity，不提供 mirroring

![800x400](https://upload.wikimedia.org/wikipedia/commons/7/70/RAID_6.svg)

使用 block-level striping；同时 parity 的数据分布存储在不同磁盘上；在 RIAD 5 的基础上额外添加了 second parity，所以如果 first parity data 丢失，还可以通过 second parity data 还原数据，允许有两块坏盘(non-parity disk 或者是 parity disk 都可以)，但是相对 RAID 5 需要额外的 write I/O 开销；至少 4 块磁盘(至少 2 块盘才能 striping，额外任意 2 块盘存储 parity data)

**Performance**

- n 块 read I/O 性能相同(或者不相同)的磁盘组成的 RAID 6，提供 $(n - 2) \times min(\mbox{non-parity\ disk\ read\ I/O})$ read I/O 性能
- n 块 write I/O 性能相同(或者不相同)的磁盘组成 RAID 6，提供 $(n - 2) \times min(\mbox{non-parity\ disk\ write\ I/O}) - (parity\ write\ I/O\ overhead)$ write I/O 性能

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 6，提供 $sum(capacity)$ 容量，只有 $(n - 2) \times min(\mbox{non-parity\ disk\ capacity})$ 的容量提供 striping/parity， $2 \times min(\mbox{non-parity\ disk\ capacity})$ 的容量存储 parity 数据，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 300GB 300MB/s read/write I/O

SSD3 400GB 400MB/s read/write I/O

SSD4 500GB 500MB/s read/write I/O

那么组成 RAID 6 后，read I/O 大于 200MB/s（当 SSD3 SSD4 存储 parity） 小于 800MB/s（当 SSD1 SSD2 存储 parity） ，write I/O 为 $\lt 400MB/s$，SSD1 SSD2 SSD3 SSD4 加起来可以提供 200GB 用于 striping/parity，SSD1 SSD2 SSD3 SSD4 加起来可以提供 200GB 用于存储 parity，剩余 900GB 只提供存储

> [!note] Choosing Criteria
> 需要较好的 read/write I/O 性能，数据需要较高的可用优先选择 RAID 6

## 0x06 Comparision

使用性能容量相同的 n 块磁盘组成 RAID

| Level  | striping/mirroring/parity   | Minimum number of drivers | Space efficiency | Fault tolerence | Read performance | Write performance |
| ------ | --------------------------- | ------------------------- | ---------------- | --------------- | ---------------- | ----------------- |
| RAID 0 | block-striping              | 1                         | 1                | None            | n                | n                 |
| RAID 1 | mirroring                   | 2                         | 1/n              | n - 1  drivers  | n                | 1                 |
| RAID 3 | byte-striping parity        | 3                         | (n - 1)/n        | 1 driver        | n - 1            | n - 1             |
| RAID 4 | byte-striping parity        | 3                         | (n - 1)/n        | 1 driver        | n - 1            | n - 1             |
| RAID 5 | byte-striping parity        | 3                         | (n - 1)/n        | 1 driver        | n - 1            | n - 1             |
| RAID 6 | byte-striping double parity | 4                         | (n - 2)/n        | 2 driver        | n - 2            | n - 2             |

## 0x07 Nested RAID

除了磁盘之间可以互相组 RIAD，还可以 RAID 之间互相组 RAID，这种形式的 RAID 被称为 Nested(Hybrid) RAID。通常用 2 个数字来标示 RAID level

数字的先后顺序不同 Nested RAID 也不同

> The first number in the numeric designation denotes the lowest RAID level in the "stack", while the rightmost one denotes the highest layered RAID level

例如 RAID 50 表示先将磁盘组成 RAID 5，然后将 RAID 5 组成 RAID 0；而 RAID 05 表示先将磁盘组成 RAID 0，然后将 RIAD 0 组成 RAID 5

常见的 Nested RAID 如下

### 0x07a RAID 01

先将磁盘组成 RAID 0，提供 striping；然后将 RAID 0 组成 RAID 1，提供 mirroring

![500](https://upload.wikimedia.org/wikipedia/commons/a/ad/RAID_01.svg)

至少 4 块磁盘

**Performance**

- n 块 read I/O 性能相同(或者不相同)的磁盘组成的 RAID 01，提供大于 $min(RAID\ 0\ read\ I/O)$ 但是小于 $max(RAID\ 0\ read\ I/O)$ read I/O 性能（read request 会被 RAID controller 发送到不同的 RAID 0）
- n 块 write I/O 性能相同(或者不相同)的磁盘组成的 RAID 01，wirte I/O 为 $min(RAID\ 0\ write\ I/O)$

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 01，提供 $sum(capacity)$ 容量，只有 $min(RAID\ 0\ capacity)$ 的容量提供 striping/mirroring，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 200GB 200MB/s read/write I/O

SSD3 300GB 300MB/s read/write I/O

SSD4 400GB 400MB/s read/write I/O

SSD1/2 和 SSD3/4 两两组成 RAID0

那么 

SSD1/2 RAID 0 read I/0 为 200MB/s，write I/O 为 200 MB/s，SSD1 SSD2 各提供 100GB 用于 striping

SSD3/4 RAID 0 read I/0 为 600MB/s，write I/0 为 600 MB/s，SSD3 SSD4 各提供 300GB 用于 striping

所以组成 RAID 01 后，read I/O 在 200MB/s 至 600MB/s 之间 ，write I/O 为 200MB/s，只有 200GB 提供 striping/miroring

### 0x07b RAID 10

先将磁盘组成 RAID 1，提供 mirroring；然后将 RAID 1 组成 RAID 0，提供 striping

![500](https://upload.wikimedia.org/wikipedia/commons/e/e6/RAID_10_01.svg)

至少 4 块磁盘

**Performance**

- n 块 read/write I/O 性能相同(或者不相同)的磁盘组成 RAID 10，提供$n \times min(RAID\ 1\ read\ I/O)$ read I/O 性能，$n \times min(RAID\ 1\ write\ I/O)$ write I/O 性能

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 10，提供 $sum(capacity)$ 容量，只有 $n \times min(RAID\ 1\ capacity)$ 的容量提供 striping/mirroring，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 200GB 200MB/s read/write I/O

SSD3 300GB 300MB/s read/write I/O

SSD4 400GB 400MB/s read/write I/O

SSD1/2 和 SSD3/4 两两组成 RAID1

那么 

SSD1/2 RAID 1 read I/0 在 100MB/s 至 200MB/s 之间，write I/O 为 100 MB/s，提供 100GB 用于 mirroring

SSD3/4 RAID 1 read I/0 为 300MB/s 至 400MB/s 之间，write I/0 为 300 MB/s，提供 300GB 用于 mirroring

所以组成 RAID 10 后，read I/O 在 200MB/s 至 400MB/s 之间 ，write I/O 为 200MB/s，只有 200GB 提供 striping/miroring

### 0x07c RAID 50

先将磁盘组成 RAID 5，提供 striping/parity；然后将 RAID 5 组成 RAID 0，提供 striping

![800](https://upload.wikimedia.org/wikipedia/commons/9/9d/RAID_50.png)

允许有 $n \div 3$ 块坏盘(non-parity disk 或者是 parity disk 都可以)；至少 6 块磁盘

**Performance**

- n 块 read/write I/O 性能相同(或者不相同)的磁盘组成 RAID 50，提供$n \times min(RAID\ 5\ read\ I/O)$ read I/O 性能，$n \times min(RAID\ 5\ write\ I/O)$ write I/O 性能

**Capacity**

- n 块容量相同(或者不相同)的磁盘组成 RAID 50，提供 $sum(capacity)$ 容量，只有 $n \times min(RAID\ 5\ capacity)$ 的容量提供 striping/parity，剩余的容量只用于存储

假设

SSD1 100GB 100MB/s read/write I/O

SSD2 200GB 200MB/s read/write I/O

SSD3 300GB 300MB/s read/write I/O

SSD4 400GB 400MB/s read/write I/O

SSD5 500GB 500MB/s read/write I/O

SSD6 600GB 600MB/s read/write I/O 

SSD1/2/3 和 SSD4/5/6 两两组成 RAID5

那么 

SSD1/2/3 RAID 5 read I/0 大于 200MB/s（当 SSD3 存储 parity） 小于 400MB/s（当 SSD1 存储 parity），write I/O 为 $\lt 300MB/s$，SSD1 SSD2 SSD3 加起来可以提供 200GB 用于 striping/parity，SSD1 SSD2 SSD3 加起来可以提供 100GB 用于存储 parity

SSD4/5/6 RAID 5 read I/0 大于 800MB/s（当 SSD6 存储 parity） 小于 1000MB/s（当 SSD4 存储 parity），write I/O 为 $\lt 1200MB/s$，SSD4 SSD5 SSD6 加起来可以提供 800GB 用于 striping/parity，SSD4 SSD5 SSD6 加起来可以提供 400GB 用于存储 parity

所以组成 RAID 50 后，read I/O 在 400MB/s 至 800MB/s 之间 ，write I/O 为 600MB/s，只有 200GB 提供 striping/parity

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [RAID - Wikipedia](https://en.wikipedia.org/wiki/RAID)
- [Standard RAID levels - Wikipedia](https://en.wikipedia.org/wiki/Standard_RAID_levels)
- [Understanding RAID 0: Benefits, Risks, and Applications \| DiskInternals](https://www.diskinternals.com/raid-recovery/understanding-raid-0/)
- [Nested RAID levels - Wikipedia](https://en.wikipedia.org/wiki/Nested_RAID_levels)

***References***

[^1]:[Data striping - Wikipedia](https://en.wikipedia.org/wiki/Data_striping)
[^2]:[Disk mirroring - Wikipedia](https://en.wikipedia.org/wiki/Disk_mirroring)
[^3]:[Parity bit - Wikipedia](https://en.wikipedia.org/wiki/Parity_bit)
