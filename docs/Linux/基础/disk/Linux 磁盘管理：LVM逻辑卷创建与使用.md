# LVM逻辑卷创建与使用

参考：

https://blog.51cto.com/3069201/2056653

https://blog.51cto.com/dreamfire/1084729

## 概述

LVM(Logical volume Manager)是逻辑卷管理的简称。它是Linux环境下对磁盘分区进行管理的一种机制。现在不仅仅是Linux系统上可以使用LVM这种磁盘管理机制，对于其它的类UNIX操作系统，以及windows操作系统都有类似与LVM这种磁盘管理软件。

LVM的工作原理其实很简单，它就是通过将底层的物理硬盘抽象的封装起来，然后以逻辑卷的方式呈现给上层应用。在传统的磁盘管理机制中，我们的上层应用是直接访问文件系统，从而对底层的物理硬盘进行读取，而在LVM中，其通过对底层的硬盘进行封装，当我们对底层的物理硬盘进行操作时，其不再是针对于分区进行操作，而是通过一个叫做逻辑卷的东西来对其进行底层的磁盘管理操作。比如说我增加一个物理硬盘，这个时候上层的服务是感觉不到的，因为呈现给上次服务的是以逻辑卷的方式。
**LVM最大的特点就是可以对磁盘进行动态管理。因为逻辑卷的大小是可以动态调整的，而且不会丢失现有的数据。我们如果新增加了硬盘，其也不会改变现有上层的逻辑卷。作为一个动态磁盘管理机制，逻辑卷技术大大提高了磁盘管理的灵活性！！！**
**原理：创建物理分区-->创建物理卷-->创建卷组-->创建逻辑卷**

<img src="https://s4.51cto.com/images/blog/201801/02/029a87f59131bdf50b2c27ee148b8908.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk="/>

## 创建逻辑卷

### 准备

1. 添加物理磁盘

2. 划分区，将分区类型改为Linux LVM(通过l参数来查看)

   ```
   [root@chz ~]# fdisk /dev/sdb
   Welcome to fdisk (util-linux 2.23.2).
   
   Changes will remain in memory only, until you decide to write them.
   Be careful before using the write command.
   
   
   Command (m for help): t
   Partition number (1-3, default 3): 3
   Hex code (type L to list all codes): 8e
   Changed type of partition 'Linux' to 'Linux LVM'
   
   Command (m for help): p
   
   Disk /dev/sdb: 2147 MB, 2147483648 bytes, 4194304 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0x97ca0327
   
      Device Boot      Start         End      Blocks   Id  System
   /dev/sdb1            2048     1026047      512000   8e  Linux LVM
   /dev/sdb2         1026048     2050047      512000   8e  Linux LVM
   /dev/sdb3         2050048     3074047      512000   8e  Linux LVM
   
   ```

> 需要使用w参数保存

### 物理卷

physical Volume

1. 将分区转为PV

   ```
   [root@chz ~]# pvcreate /dev/sdb1 /dev/sdb2
     Physical volume "/dev/sdb1" successfully created.
     Physical volume "/dev/sdb2" successfully created.
   ```

2. 查看创建的PV

   > PV由多个PE(最小的单位)组成

   ```
   [root@chz ~]# pvs
     PV         VG     Fmt  Attr PSize   PFree  
     /dev/sda2  centos lvm2 a--  <19.00g      0 
     /dev/sdb1         lvm2 ---  500.00m 500.00m
     /dev/sdb2         lvm2 ---  500.00m 500.00m
   [root@chz ~]# pvdisplay
     --- Physical volume ---
     PV Name               /dev/sda2
     VG Name               centos
     PV Size               <19.00 GiB / not usable 3.00 MiB
     Allocatable           yes (but full)
     PE Size               4.00 MiB
     Total PE              4863
     Free PE               0
     Allocated PE          4863
     PV UUID               eA52jE-SFuU-BG5t-Isyw-wWdY-lj4K-xso1bu
   
   ```

### 卷组

Volume Group

1. 将两个物理卷划分到一个VG

   卷组的名字为vg1

   ```
   [root@chz ~]# vgcreate vg1 /dev/sdb1  /dev/sdb2 
     Volume group "vg1" successfully created
   ```

2. 查看卷组

   ```
   [root@chz ~]# vgs
     VG     #PV #LV #SN Attr   VSize   VFree  
     centos   1   2   0 wz--n- <19.00g      0 
     vg1      2   0   0 wz--n- 992.00m 992.00m
   [root@chz ~]# vgdisplay
     --- Volume group ---
     VG Name               vg1
     System ID             
     Format                lvm2
     Metadata Areas        2
     Metadata Sequence No  1
     VG Access             read/write
     VG Status             resizable
     MAX LV                0
     Cur LV                0
     Open LV               0
     Max PV                0
     Cur PV                2
     Act PV                2
     VG Size               992.00 MiB
     PE Size               4.00 MiB
     Total PE              248
     Alloc PE / Size       0 / 0   
     Free  PE / Size       248 / 992.00 MiB
     VG UUID               CVbfWw-PGjT-B5dk-eA0J-FgES-heCz-yuhjCB
   ```

### 逻辑卷

Logincal Volume

1. 划分容量给逻辑卷

   `-L`指定容量大小，`-n`指定LV名字

   ```
   [root@chz ~]# lvcreate -L 200M -n lv01 vg1
     Logical volume "lv01" created.
   ```

   使用`-l 100%FREE`表示分配剩余空间

   ```
   [root@chz ~]# lvcreate -n lv02 -l 100%FREE vg1
     Logical volume "lv02" created.
   [root@chz ~]# lvs
     LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
     root centos -wi-ao---- <17.00g                                                    
     swap centos -wi-ao----   2.00g                                                    
     lv01 vg1    -wi-ao---- 300.00m                                                    
     lv02 vg1    -wi-a----- 692.00m 
   ```

2. 查看LV

   物理位置为`dev/VG/LV`

   ```
   [root@chz ~]# lvs
     LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
     root centos -wi-ao---- <17.00g                                                    
     swap centos -wi-ao----   2.00g                                                    
     lv01 vg1    -wi-a----- 200.00m                                                    
   [root@chz ~]# lvdisplay 
     --- Logical volume ---
     LV Path                /dev/vg1/lv01
     LV Name                lv01
     VG Name                vg1
     LV UUID                R35G7T-ZGgn-P0i2-aSCo-ZJDB-U0gS-3Whjsf
     LV Write Access        read/write
     LV Creation host, time chz, 2020-10-21 11:31:24 +0800
     LV Status              available
     # open                 0
     LV Size                200.00 MiB
     Current LE             50
     Segments               1
     Allocation             inherit
     Read ahead sectors     auto
     - currently set to     8192
     Block device           253:2
   ```

### 格式化文件系统

1. 使用ext4文件系统

   ```
   [root@chz ~]# mkfs.ext4 /dev/vg1/lv01 
   mke2fs 1.42.9 (28-Dec-2013)
   Filesystem label=
   OS type: Linux
   Block size=1024 (log=0)
   Fragment size=1024 (log=0)
   Stride=0 blocks, Stripe width=0 blocks
   51200 inodes, 204800 blocks
   10240 blocks (5.00%) reserved for the super user
   First data block=1
   Maximum filesystem blocks=33816576
   25 block groups
   8192 blocks per group, 8192 fragments per group
   2048 inodes per group
   Superblock backups stored on blocks: 
   	8193, 24577, 40961, 57345, 73729
   
   Allocating group tables: done                            
   Writing inode tables: done                            
   Creating journal (4096 blocks): done
   Writing superblocks and filesystem accounting information: done 
   ```

### 挂载

1. 挂载==（不会永久挂载）==

   需要修改`/etc/fstab`

   ```
   [root@chz Desktop]# mount /dev/vg1/lv01  /root/Desktop/test
   ```

2. 查看

   ```
   [root@chz Desktop]# df -Th
   Filesystem              Type      Size  Used Avail Use% Mounted on
   devtmpfs                devtmpfs  470M     0  470M   0% /dev
   tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
   tmpfs                   tmpfs     487M  8.6M  478M   2% /run
   tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
   /dev/mapper/centos-root xfs        17G  5.6G   12G  33% /
   /dev/sda1               xfs      1014M  172M  843M  17% /boot
   tmpfs                   tmpfs      98M  4.0K   98M   1% /run/user/42
   tmpfs                   tmpfs      98M   20K   98M   1% /run/user/0
   /dev/sr0                iso9660   4.4G  4.4G     0 100% /run/media/root/CentOS 7 x86_64
   /dev/mapper/vg1-lv01    ext4      190M  1.6M  175M   1% /root/Desktop/test
   ```

   > 使用lsblk查看

   ```
   [root@chz ~]# lsblk
   NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda            8:0    0   20G  0 disk 
   ├─sda1         8:1    0    1G  0 part /boot
   └─sda2         8:2    0   19G  0 part 
     ├─centos-root
                253:0    0   17G  0 lvm  /
     └─centos-swap
                253:1    0    2G  0 lvm  [SWAP]
   sdb            8:16   0    2G  0 disk 
   ├─sdb1         8:17   0  500M  0 part 
   │ └─vg1-lv01 253:2    0  300M  0 lvm  /root/Desktop/test
   ├─sdb2         8:18   0  500M  0 part 
   │ └─vg1-lv02 253:3    0  100M  0 lvm  
   └─sdb3         8:19   0  500M  0 part 
   sr0           11:0    1  4.4G  0 rom  /run/media/root/CentO
   ```

## 物理卷转普通卷

```
[root@chz ~]# pvremove /dev/sdb1
  Labels on physical volume "/dev/sdb1" successfully wiped.
[root@chz ~]# pvs
  PV         VG     Fmt  Attr PSize   PFree  
  /dev/sda2  centos lvm2 a--  <19.00g      0 
  /dev/sdb2         lvm2 ---  500.00m 500.00m
  /dev/sdb3         lvm2 ---  500.00m 500.00m
[root@chz ~]# 
```

> 之后还需要将System置为Linux

```
Disk /dev/sdb: 2147 MB, 2147483648 bytes, 4194304 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk label type: dos
Disk identifier: 0x97ca0327

   Device Boot      Start         End      Blocks   Id  System
/dev/sdb1            2048     1026047      512000   83  Linux
/dev/sdb2         1026048     2050047      512000   8e  Linux LVM
/dev/sdb3         2050048     3074047      512000   8e  Linux LVM

```

## 扩容卷组

```
[root@chz ~]# vgextend vg1 /dev/sdb3
  Physical volume "/dev/sdb3" successfully created.
  Volume group "vg1" successfully extended
[root@chz ~]# vgs
  VG     #PV #LV #SN Attr   VSize   VFree 
  centos   1   2   0 wz--n- <19.00g     0 
  vg1      3   2   0 wz--n-   1.45g <1.26g
[root@chz ~]# 

```

## 删除卷组

删除卷组时，也会删除卷组上的逻辑卷

```
[root@chz ~]# vgremove vg1 
Do you really want to remove volume group "vg1" containing 1 logical volumes? [y/n]: y
Do you really want to remove active logical volume vg1/lv02? [y/n]: y
  Logical volume "lv02" successfully removed
  Volume group "vg1" successfully removed
[root@chz ~]# vgs
  VG     #PV #LV #SN Attr   VSize   VFree
  centos   1   2   0 wz--n- <19.00g    0 
```

## 扩容逻辑卷

> 卷组需要由足够的空间

- lvresize

  ```
  [root@chz Desktop]# lvresize -L 300M /dev/vg1/lv01 
    Size of logical volume vg1/lv01 changed from 200.00 MiB (50 extents) to 300.00 MiB (75 extents).
    Logical volume vg1/lv01 successfully resized.
    
  [root@chz Desktop]# lvs
    LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
    root centos -wi-ao---- <17.00g                                                    
    swap centos -wi-ao----   2.00g                                                    
    lv01 vg1    -wi-ao---- 300.00m  
  ```

- lvextend

  ==可与指定一个数字，或是在原有的基础上加一个数字==

  ```
  [root@chz ~]# lvs
    LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
    root centos -wi-ao---- <17.00g                                                    
    swap centos -wi-ao----   2.00g                                                    
    lv01 vg1    -wi-ao---- 300.00m                                                    
    lv02 vg1    -wi-a----- 100.00m                                                    
  [root@chz ~]# lvextend -L +100M /dev/vg1/lv01
    Size of logical volume vg1/lv01 changed from 300.00 MiB (75 extents) to 400.00 MiB (100 extents).
    Logical volume vg1/lv01 successfully resized.
  [root@chz ~]# lvs
    LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
    root centos -wi-ao---- <17.00g                                                    
    swap centos -wi-ao----   2.00g                                                    
    lv01 vg1    -wi-ao---- 400.00m                                                    
    lv02 vg1    -wi-a----- 100.00m            
  ```

## 缩减逻辑卷

- lvresize

  可能坏造成数据丢失

  ```
  [root@chz ~]# lvresize -L 100 /dev/vg1/lv01
    WARNING: Reducing active and open logical volume to 100.00 MiB.
    THIS MAY DESTROY YOUR DATA (filesystem etc.)
  Do you really want to reduce vg1/lv01? [y/n]: y
    Size of logical volume vg1/lv01 changed from 400.00 MiB (100 extents) to 100.00 MiB (25 extents).
    Logical volume vg1/lv01 successfully resized.
  [root@chz ~]# lvs
    LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
    root centos -wi-ao---- <17.00g                                                    
    swap centos -wi-ao----   2.00g                                                    
    lv01 vg1    -wi-ao---- 100.00m                                                    
    lv02 vg1    -wi-a----- 100.00m          
  ```

- lvreduce

  ```
  [root@chz ~]# lvreduce -L 100M /dev/vg1/lv01
    WARNING: Reducing active and open logical volume to 100.00 MiB.
    THIS MAY DESTROY YOUR DATA (filesystem etc.)
  Do you really want to reduce vg1/lv01? [y/n]: y
    Size of logical volume vg1/lv01 changed from 200.00 MiB (50 extents) to 100.00 MiB (25 extents).
    Logical volume vg1/lv01 successfully resized.
  ```

> 正确步骤，==在检查分区(fsck)时需要先卸载(umount)==

1. 先卸载逻辑卷

   ```
   [root@chz ~]# lsblk
   NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda            8:0    0   20G  0 disk 
   ├─sda1         8:1    0    1G  0 part /boot
   └─sda2         8:2    0   19G  0 part 
     ├─centos-root
                253:0    0   17G  0 lvm  /
     └─centos-swap
                253:1    0    2G  0 lvm  [SWAP]
   sdb            8:16   0    2G  0 disk 
   ├─sdb1         8:17   0  500M  0 part 
   │ └─vg1-lv01 253:2    0  300M  0 lvm  /root/Desktop/test
   ├─sdb2         8:18   0  500M  0 part 
   │ └─vg1-lv02 253:3    0  100M  0 lvm  
   └─sdb3         8:19   0  500M  0 part 
   sr0           11:0    1  4.4G  0 rom  /run/media/root/CentO
   [root@chz ~]# umount /dev/vg1/lv01 /root/Desktop/test
   umount: /root/Desktop/test: not mounted
   [root@chz ~]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sdb               8:16   0    2G  0 disk 
   ├─sdb1            8:17   0  500M  0 part 
   │ └─vg1-lv01    253:2    0  300M  0 lvm  
   ├─sdb2            8:18   0  500M  0 part 
   │ └─vg1-lv02    253:3    0  100M  0 lvm  
   └─sdb3            8:19   0  500M  0 part 
   sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
   
   ```

2. 使用`fsck`查看文件系统损坏，`e2fsck`用于检查ext2/ext3

   ```
   [root@chz ~]# fsck /dev/vg1/lv01
   e2fsck 1.42.9 (28-Dec-2013)
   /dev/vg1/lv01: clean, 11/51200 files, 12115/204800 blocks
   ```

3. 使用`resize2fs`增大或缩减文件系统大小

   > ==需要保持文件系统和逻辑卷大小一致，否则回到空间浪费==

   ```
   [root@chz ~]# resize2fs /dev/vg1/lv01 300M
   resize2fs 1.42.9 (28-Dec-2013)
   Resizing the filesystem on /dev/vg1/lv01 to 307200 (1k) blocks.
   The filesystem on /dev/vg1/lv01 is now 307200 blocks long.
   ```

4. 使用`lvreduce`修改逻辑卷大小

   ```
   [root@chz ~]# lvreduce -L 300M /dev/vg1/lv01
     WARNING: Reducing active logical volume to 300.00 MiB.
     THIS MAY DESTROY YOUR DATA (filesystem etc.)
   Do you really want to reduce vg1/lv01? [y/n]: y
     Size of logical volume vg1/lv01 changed from 500.00 MiB (125 extents) to 300.00 MiB (75 extents).
     Logical volume vg1/lv01 successfully resized.
   [root@chz ~]# lvs
     LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
     root centos -wi-ao---- <17.00g                                                    
     swap centos -wi-ao----   2.00g                                                    
     lv01 vg1    -wi-a----- 300.00m                                                    
     lv02 vg1    -wi-a----- 100.00m                           
   ```

5. 重新挂载

## 删除逻辑卷

- lvremove

  ```
  [root@chz ~]# lvremove /dev/vg1/lv01
  Do you really want to remove active logical volume vg1/lv01? [y/n]: y
    Logical volume "lv01" successfully removed
  [root@chz ~]# lvs
    LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
    root centos -wi-ao---- <17.00g                                                    
    swap centos -wi-ao----   2.00g                                                    
    lv02 vg1    -wi-a----- 100.00m          
  ```











