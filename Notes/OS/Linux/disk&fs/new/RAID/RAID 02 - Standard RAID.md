---
createTime: 2025-02-06 14:38
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID 02 - Standard RAID

## 0x01 Preface

RAID 根据是否提供 striping（包括 striping 的最小单元）/mirroring/parity(包括 parity 数据存储的方式)，将其划分为 7 个 standard RAID level

- RAID 0
- RAID 1
- RAID 2
- RAID 3
- RAID 4
- RAID 5
- RAID 6

其中 RAID 0/RAID 1/RAID 5 最常用

## 0x02 Standard RAID LEVEL

> [!important]
> 下列例子中均为理论值，实际要比这个小，因为 I/O 快的磁盘需要等 I/O 慢的磁盘完成 read 或者是 write

### 0x02a RAID 0

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

### 0x02b RAID 1

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

### 0x02c RAID 2

提供 striping/parity，不提供 mirroring

![800x400](https://upload.wikimedia.org/wikipedia/commons/b/b5/RAID2_arch.svg)

使用 bit-level striping，因为以 1 bit 为一个单元，大量写入或者修改数据，需要非常频繁写入或者更新 parity 数据，所以综合 write I/O 性能非常差，有可能不如 RAID 1；同时使用复杂的 hamming code 用于错误校验，write I/O 性能进一步降低。几乎没有设备会使用 RAID 2，所以这里就不做过多的介绍

### 0x02d RAID 3

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

### 0x02e RAID 4

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

### 0x02f RAID 5

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

### 0x02g RAID 6

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

## 0x03 Comparision

使用性能容量相同的 n 块磁盘组成 RAID

| Level  | striping/mirroring/parity   | Minimum number of drivers | Space efficiency | Fault tolerence | Read performance | Write performance |
| ------ | --------------------------- | ------------------------- | ---------------- | --------------- | ---------------- | ----------------- |
| RAID 0 | block-striping              | 1                         | 1                | None            | n                | n                 |
| RAID 1 | mirroring                   | 2                         | 1/n              | n - 1  drivers  | n                | 1                 |
| RAID 3 | byte-striping parity        | 3                         | (n - 1)/n        | 1 driver        | n - 1            | n - 1             |
| RAID 4 | byte-striping parity        | 3                         | (n - 1)/n        | 1 driver        | n - 1            | n - 1             |
| RAID 5 | byte-striping parity        | 3                         | (n - 1)/n        | 1 driver        | n - 1            | n - 1             |
| RAID 6 | byte-striping double parity | 4                         | (n - 2)/n        | 2 driver        | n - 2            | n - 2             |


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Standard RAID levels - Wikipedia](https://en.wikipedia.org/wiki/Standard_RAID_levels)
- [Understanding RAID 0: Benefits, Risks, and Applications \| DiskInternals](https://www.diskinternals.com/raid-recovery/understanding-raid-0/)

***References***


