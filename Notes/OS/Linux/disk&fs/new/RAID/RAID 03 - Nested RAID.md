---
createTime: 2025-02-06 14:43
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID 03 - Nested RAID

## 0x01 Preface

除了磁盘之间可以互相组 RIAD，还可以 RAID 之间互相组 RAID，这种形式的 RAID 被称为 Nested(Hybrid) RAID。通常用 2 个数字来标识 RAID level

数字的先后顺序不同 Nested RAID 也不同

> The first number in the numeric designation denotes the lowest RAID level in the "stack", while the rightmost one denotes the highest layered RAID level

例如 RAID 50 表示先将磁盘组成 RAID 5，然后将 RAID 5 组成 RAID 0；而 RAID 05 表示先将磁盘组成 RAID 0，然后将 RIAD 0 组成 RAID 5

常见的 Nested RAID 有 RAID 01/RAID 10/RAID 50

## 0x02 Nested RAID Level

### 0x02a RAID 01

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

### 0x02b RAID 10

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

### 0x02c RAID 50

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

- [Nested RAID levels - Wikipedia](https://en.wikipedia.org/wiki/Nested_RAID_levels)

***References***


