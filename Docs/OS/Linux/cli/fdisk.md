# Linux 磁盘管理：fdisk

参考:

https://www.cnblogs.com/itech/archive/2010/12/24/1916255.html

https://www.jianshu.com/p/bf939474d69b

https://www.thegeekdiary.com/understanding-linux-fdisk-utility/

gdisk使用参考

https://www.cnblogs.com/Sunzz/p/6908329.html

```mermaid
graph LR
添加物理硬盘-->分区-->格式化文件系统-->挂载
```



> 对于fdisk不熟的朋友还可以使用cfdisk来分区
>
> ==同一个目录下只能挂载一个分区==
>
> 最好以sector为单位，方便4k对齐

## 概念

优于一般linux分配硬盘空间不会超过2T，所以也就无需使用GPT分区表，所以使用MBR分区工具`fdisk`即可（**GPT分区表使用gpart 或是 gdisk**）

1. 在linux下SCSI接口设备以sd命名，第一个是sda，第二个sdb，依次类推。IDE接口设备用hd(hard disk)命名，第一个是hda，第二个是hdb，依次类推。==sr0表示是是SCSI的磁盘驱动==。软盘以fd(floppy disk)开头
2. 分区是用设备名称加数字命名。例如sda1代表sda这个硬盘设备上的第一个分区。
3. MBR分区、表最多有四个主分区（windows对应C,D,E,F 盘），一个扩展分区，扩张分区可以在分为多个逻辑分区。

> 我们可以通过 lsblk 命令来查看硬盘，分区以及挂载点
>
> linux 中1-4都是主分区，从5开始为逻辑分区

```bash
[root@chz ~]# lsblk
NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda               8:0    0   20G  0 disk 
├─sda1            8:1    0    1G  0 part /boot
└─sda2            8:2    0   19G  0 part 
  ├─centos-root 253:0    0   17G  0 lvm  /
  └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
sdb               8:16   0    1G  0 disk 
sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
[root@chz ~]# 

```

## fdisk

> 如果自建分区系统，指定主引导分区需要使用`-t`参数来改变分区的类型

- `fdisk -l`

  查看所有磁盘和分区

  1. Disk：磁盘容量，和逻辑扇区总数
  2. Units：簇，一个簇包含一个逻辑扇区大小为512bytes
  3. Sector size：逻辑扇区和物理扇区大小
  4. I/O size：磁盘每次IO的最小和最佳大小
  5. Disklable：磁盘的分区表类型。dos代表mbr，gpt代表gpt
  6. Disk identifier：磁盘的标识符

  ```bash
  $ sudo fdisk -l
  Disk /dev/sda: 30 GiB, 32212254720 bytes, 62914560 sectors
  Units: sectors of 1 * 512 = 512 bytes
  Sector size (logical/physical): 512 bytes / 512 bytes
  I/O size (minimum/optimal): 512 bytes / 512 bytes
  Disklabel type: dos
  Disk identifier: 0xeab59449
  
  Device     Boot    Start      End  Sectors Size Id Type
  /dev/sda1  *    20973568 62914559 41940992  20G 83 Linux
  
  Disk /dev/sdb: 10 GiB, 10737418240 bytes, 20971520 sectors
  Units: sectors of 1 * 512 = 512 bytes
  Sector size (logical/physical): 512 bytes / 512 bytes
  I/O size (minimum/optimal): 512 bytes / 512 bytes
  
  Disk /dev/sdc: 10 GiB, 10737418240 bytes, 20971520 sectors
  Units: sectors of 1 * 512 = 512 bytes
  Sector size (logical/physical): 512 bytes / 512 bytes
  I/O size (minimum/optimal): 512 bytes / 512 bytes
  
  Disk /dev/sdd: 10 GiB, 10737418240 bytes, 20971520 sectors
  Units: sectors of 1 * 512 = 512 bytes
  Sector size (logical/physical): 512 bytes / 512 bytes
  I/O size (minimum/optimal): 512 bytes / 512 bytes
  
  Disk /dev/sde: 10 GiB, 10737418240 bytes, 20971520 sectors
  Units: sectors of 1 * 512 = 512 bytes
  Sector size (logical/physical): 512 bytes / 512 bytes
  I/O size (minimum/optimal): 512 bytes / 512 bytes
  ```

  这里可以发现，5块硬盘，sda有两个主分区；sda1做为主引导分区。Unites和Sector分别表示柱面单元大小和扇区大小

  Sectors = End - Start + 1

  Size = Sectors * logical sector size / 1024 ^ 3

  |                                        | 时否是主引导分区 （*表示是） | 起始柱面(开始扇区) | 结束柱面(结束扇区) | 分区总扇区 | 分区类型 | 对分区类型的解析 |       |
  | -------------------------------------- | ---------------------------- | ------------------ | ------------------ | ---------- | -------- | ---------------- | ----- |
  | Device Boot Start End Blocks Id System | Boot                         | Start              | End                | Sectors    | Id       | Type             | Size  |
  | /dev/vda1                              | *                            | 2048               | 1026047            | 41940992   | 83       | Linux            | 500M  |
  | /dev/vda2                              |                              | 1026048            | 209715199          | 104344576  | 8e       | Linux LVM        | 99.5G |

- `fdisk /dev/sda`

  操作具体某块硬盘

## 挂载磁盘

> ==注意如果挂载的目录之前有东西，最好先备份到一个新文件夹，否则后面挂载的新磁盘会覆盖之前的内容(只要umount就可以显示会原来的内容)。==

1. 在虚拟机中添加一块硬盘

   > 添加硬盘后重启，让硬盘生效
   >
   > 通过`lsblk`或是`fdisk -l`来查看新添加的硬盘

2. 分区

   > 可以通过 + size {K,M,G} 来指定结束柱面
   >
   > 这里创建了一个100M的主分区sdb1

   创建主分区

   ```shell
   [root@chz ~]# fdisk /dev/sdb
   Welcome to fdisk (util-linux 2.23.2).
   
   Changes will remain in memory only, until you decide to write them.
   Be careful before using the write command.
   
   Device does not contain a recognized partition table
   Building a new DOS disklabel with disk identifier 0xac359141.
   
   Command (m for help): n
   Partition type:
      p   primary (0 primary, 0 extended, 4 free)
      e   extended
   Select (default p):     
   Using default response p #选择主分区
   Partition number (1-4, default 1): 
   First sector (2048-2097151, default 2048): 
   Using default value 2048
   Last sector, +sectors or +size{K,M,G} (2048-2097151, default 2097151): +100M #结束柱面
   Partition 1 of type Linux and of size 100 MiB is set
   
   Command (m for help): p
   
   Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0xac359141
   
      Device Boot      Start         End      Blocks   Id  System
   /dev/sdb1            2048      206847      102400   83  Linux
   
   Command (m for help): n
   Partition type:
      p   primary (1 primary, 0 extended, 3 free) #有1个主分区，0个扩展分区，还可以创建3个主分区
      e   extended
   Select (default p):  
   ```

   创建扩展分区（扩展分区需要创建逻辑分区才能使用）

   [传送门](..\..\Hardware\主分区、扩展分区、逻辑分区、动态分区.md)

   ```bash
   Command (m for help): n
   Partition type:
      p   primary (1 primary, 0 extended, 3 free)
      e   extended
   Select (default p): e #选择扩展分区
   Partition number (2-4, default 2): 2
   First sector (206848-2097151, default 206848): 
   Using default value 206848
   Last sector, +sectors or +size{K,M,G} (206848-2097151, default 2097151): +100M
   Partition 2 of type Extended and of size 100 MiB is set
   
   Command (m for help): p
   
   Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0xac359141
   
      Device Boot      Start         End      Blocks   Id  System
   /dev/sdb1            2048      206847      102400   83  Linux #主分区
   /dev/sdb2          206848      411647      102400    5  Extended #扩展分区
   
   Command (m for help): 
   ```

   在扩展分区上创建逻辑分区

   ```bash
   Command (m for help): n
   Partition type:
      p   primary (1 primary, 1 extended, 2 free)
      l   logical (numbered from 5) #逻辑分区从5开始计数
   Select (default p): l
   Adding logical partition 5 
   First sector (208896-411647, default 208896): 
   Using default value 208896
   Last sector, +sectors or +size{K,M,G} (208896-411647, default 411647): +20M
   Partition 5 of type Linux and of size 20 MiB is set
   
   Command (m for help): p
   
   Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0xf9fe9356
   
      Device Boot      Start         End      Blocks   Id  System
   /dev/sdb1            2048      206847      102400   83  Linux
   /dev/sdb2          206848      411647      102400    5  Extended
   /dev/sdb5          208896      249855       20480   83  Linux
   
   Command (m for help): 
   
   ```

   保存分区

   ```bash
   Command (m for help): w
   The partition table has been altered!
   
   Calling ioctl() to re-read partition table.
   ```

   `lsblk`查看分区

   ``` bash
   [root@chz ~]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sdb               8:16   0    1G  0 disk 
   ├─sdb1            8:17   0  100M  0 part 
   ├─sdb2            8:18   0    1K  0 part 
   └─sdb5            8:21   0   20M  0 part 
   sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
   ```

3. 格式化

   这里将 /dev/sdb1分区格式化为ext4文件系统，默认使用ext2

   > 这里mkfs.ext4是mke2fs的衍生命令， 也可以使用mkfs -t ext4  /dev/sdb1
   >
   > 将磁盘挂载后通过 df -T 来查看磁盘的类型
   >
   > 生成ext4文件系统后系统会创建一个`lost + found`文件

   ```bash
   [root@chz ~]# mkfs.ext4 /dev/sdb1
   mke2fs 1.42.9 (28-Dec-2013)
   Filesystem label=
   OS type: Linux
   Block size=1024 (log=0)
   Fragment size=1024 (log=0)
   Stride=0 blocks, Stripe width=0 blocks
   25688 inodes, 102400 blocks
   5120 blocks (5.00%) reserved for the super user
   First data block=1
   Maximum filesystem blocks=33685504
   13 block groups
   8192 blocks per group, 8192 fragments per group
   1976 inodes per group
   Superblock backups stored on blocks: 
   	8193, 24577, 40961, 57345, 73729
   
   Allocating group tables: done                            
   Writing inode tables: done                            
   Creating journal (4096 blocks): done
   Writing superblocks and filesystem accounting information: done 
   
   ```

4. 挂载

   > 这种方式挂载并不会永久挂载，当重启后失效。在`fstab`中添加后使用该命令可以不用重启

   ```bash
   [root@chz ~]# mount /dev/sdb1 /root/test
   [root@chz ~]# df -T
   Filesystem              Type     1K-blocks    Used Available Use% Mounted on
   devtmpfs                devtmpfs    480800       0    480800   0% /dev
   tmpfs                   tmpfs       497840       0    497840   0% /dev/shm
   tmpfs                   tmpfs       497840    8740    489100   2% /run
   tmpfs                   tmpfs       497840       0    497840   0% /sys/fs/cgroup
   /dev/mapper/centos-root xfs       17811456 7798620  10012836  44% /
   /dev/sda1               xfs        1038336  217148    821188  21% /boot
   tmpfs                   tmpfs        99572      20     99552   1% /run/user/0
   /dev/sr0                iso9660    4554702 4554702         0 100% /run/media/root/CentOS 7 x86_64
   /dev/sdb1               ext4         95054    1550     86336   2% /root/test
   
   ```

5. 永久挂载

   ==注意备份==

   从左至右字段分别为，device or filesystem，mount point，type of filesystem，mount options，which filesystems need to be dumped(0表示不存档)，which filesystems  need to be checked(0表示在引导时不需要检查)。
   
   ```
   [root@chz ~]# cat /etc/fstab
   
   #
   # /etc/fstab
   # Created by anaconda on Mon Aug 24 07:49:09 2020
   #
   # Accessible filesystems, by reference, are maintained under '/dev/disk'
   # See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info
   #
   /dev/mapper/centos-root /                       xfs     defaults        0 0
   UUID=52ff1027-e9d7-427d-9f43-3a98ba708796 /boot                   xfs     defaults        0 0
   /dev/mapper/centos-swap swap                    swap    defaults        0 0
   /dev/sdb1 /root/test ext4 defaults 0 0
   
   ```

   我们也可以通过UUID来挂载，通过`blkid`来获取UUID
   
   ```bash
   [root@chz ~]# blkid
   /dev/sda1: UUID="52ff1027-e9d7-427d-9f43-3a98ba708796" TYPE="xfs" 
   /dev/sda2: UUID="eA52jE-SFuU-BG5t-Isyw-wWdY-lj4K-xso1bu" TYPE="LVM2_member" 
   /dev/sdb1: UUID="8492fdac-fc1c-40e7-9548-ba9d3944e348" TYPE="ext4" 
   /dev/sr0: UUID="2019-09-11-18-50-31-00" LABEL="CentOS 7 x86_64" TYPE="iso9660" PTTYPE="dos" 
   /dev/mapper/centos-root: UUID="c18c2095-72b4-49f4-95be-a58d0a6cc2ad" TYPE="xfs" 
   /dev/mapper/centos-swap: UUID="0b4a0dde-4db6-4494-a067-80f3318849f7" TYPE="swap" 
   ```

## 删除分区

> 注意如果直接删除分区，但是没有修改 /etc/fstab中挂载的就会出现give root password for maintenance
>
> ==删除分区后，如果没有重新在原来磁盘块位置写入内容或是格式化文件系统，磁盘上的内容都不会消失==















