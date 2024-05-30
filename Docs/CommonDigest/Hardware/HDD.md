# HDD



ref：

https://zhuanlan.zhihu.com/p/46829113

https://en.wikipedia.org/wiki/Hard_disk_drive

https://www.crucial.com/articles/pc-builders/what-is-a-hard-drive



hard disk drive 简称 HDD，中文叫硬盘(也有叫机械硬盘)

![2022-07-22_21-17](https://git.poker/dhay3/image-repo/blob/master/20220722/2022-07-22_21-17.63zkoqyszmkg.webp?raw=true)

我们都知道系统如果需要读取数据，就需要将数据拉取到内存中。但是内存在电脑关机时就会清空。那我们写的程序和文档就无法保存，所以我就需要一个介质来存储这些数据，这就是硬盘。硬盘的容量一般很大，通常是 TB 级别的。

https://en.wikipedia.org/wiki/Hard_disk_drive#Components#Components

- platter

  盘片，盘片表面有很多缝隙(compartment)，这些缝隙就是记录你的数据的

- actuator arm/head

  当需要读取或者写入数据是 HDD 会转动，head 会指向对应的 compartment

> 因为这东西会转，所以会有炒豆子一样的声音

## Drawbacks

HDD can be slow, especially to open large applications or files. ==Because they do not write data sequentially , the data can become fragmented==, with empty space within each compartment

因为HDD不连续写入的特性，寻址需要时间所以速度比较慢。

为了提高读写的速度，所以磁盘需要 defragment ，具体参考

https://www.crucial.com/articles/pc-users/how-to-defrag-hard-drive

或者直接使用 SSD

## Types

1. 按照硬盘尺寸大小有 3.5 inch 和 2.5 inch 的硬盘
2. 按照接口分类，HDD 一般有 [IDE, SATA, SCSI](../Interface/IDE & SATA & SCSI.md)等接口。一般没有[NVME, PCIE, M2](../Interface/PCIe & NVME & M.2.md)等接口（SSD会使用这些接口）

## Speed

硬盘的速度一般和RPM挂钩(==个人理解，SSD不算在HDD中，因为SSD的工作逻辑和HDD不同==)，RPM越快速度越快。不考虑接口等其他因素 5400 RPMps 的硬盘理论速度大概有 100MBps，7200 RPMps 大概有 120 MBps

可以使用 crystalDisk 或者 disk genius 做 benchmark

