# swap分区

参考：

https://www.cnblogs.com/kerrycode/p/5246383.html

## 概述

类似于windows中的虚拟内存，当内存不足时，将一部分硬盘空间虚拟成内存使用。虽然这个SWAP分区能够作为"虚拟"的内存,但它的速度比物理内存可是慢多了,因此如果需要更快的速度的话,并不能寄厚望于SWAP,最好的办法仍然是加大物理内存。

![Snipaste_2020-10-21_18-39-37](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-10-21_18-39-37.1ef3nfzlrkgw.png)

## 查看swap分区大小

- free

  ```
  [root@chz dev]# free -h
                total        used        free      shared  buff/cache   available
  Mem:           972M        585M         67M         19M        319M        215M
  Swap:          2.0G        371M        1.6G
  ```

- swapon

  查看交换分区，等价于`cat /proc/swaps`，

  ```
  [root@chz dev]# swapon
  NAME      TYPE      SIZE USED PRIO
  /dev/dm-1 partition   2G 347M   -2
  ```

  启用交换分区

  ```
  [root@chz dev]# swapon /dev/dm-1
  ```

- swapoff

  关闭交换功能。如果需要永久替换修改`/etc/fstab`

  ```
  [root@chz Desktop]# swapoff -a
  [root@chz Desktop]# swapoff /dev/dm-1 
  ```

  > 参考：https://www.lijiaocn.com/%E9%97%AE%E9%A2%98/2019/02/27/linux-swap-off-fail.html
  >
  > 出现swapoff failed: Cannot allocate memory。是因为主存空间不够将交换分区中的内容写入。

  解决方法：

  ```
  sync ; echo 3 > /proc/sys/vm/drop_caches
  ```

  `drop_caches`接受的参数是`1`、`2`、`3`，分别清空pagecache、slab对象、pagecahce和slab对象。

  `dirty`状态的内存缓存不会被释放，如果要释放尽可能多的内存缓存，先执行命令`sync`，减少dirty状态的内存缓存。如果要disable，输入参数`4`，注意`0`不被接受：

  ==上述方法可能不会成功，当你的swap分区使用太多的时候。==

  ## swap分区大小设置

  | **RAM**                         | **Swap Space**             |
  | ------------------------------- | -------------------------- |
  | **Up to 512 MB**                | 2 times the size of RAM    |
  | **Between 1024 MB and 2048 MB** | 1.5 times the size of RAM  |
  | **Between 2049 MB and 8192 MB** | Equal to the size of RAM   |
  | **More than 8192 MB**           | 0.75 times the size of RAM |

## 积极使用swap

参考：https://blog.csdn.net/tianlesoftware/article/details/8741873

0表示最大限度使用物理内存，然后才是swap空间，100的时候表示积极使用swap分区，并且把内存上的数据及时搬运到swap空间中。==我们可以将该参数设置的低一些，让操作系统尽可能使用物理内存，降低系统对swap的使用，从而提高系统的性能==

```
[root@chz ~]# cat /proc/sys/vm/swappiness 
30
```

- 临时修改

  重启后失效

  ```
  [root@chz ~]# sysctl vm.swappiness=10
  vm.swappiness = 10
  [root@chz ~]# cat /proc/sys/vm/swappiness 
  10
  ```

- 永久修改

  ```
  [root@chz ~]# echo 'vm.swappiness=10' >> /etc/sysctl.conf 
  [root@chz ~]# cat /etc/sysctl.conf 
  vm.overcommit_memory=0
  net.ipv4.ip_forward=1
  vm.swappiness=10
  ```

## 添加swap分区

> 推荐使用磁盘来添加交换分区，或是在创建虚拟机时添加

1. 添加磁盘

2. 分区

3. 修改分区类型

   ```
   [root@chz ~]# fdisk /dev/sdb
   Welcome to fdisk (util-linux 2.23.2).
   
   Changes will remain in memory only, until you decide to write them.
   Be careful before using the write command.
   
   
   Command (m for help): t
   Partition number (1-3, default 3): 1 
   Hex code (type L to list all codes): L
   
    0  Empty           24  NEC DOS         81  Minix / old Lin bf  Solaris        
    1  FAT12           27  Hidden NTFS Win 82  Linux swap / So c1  DRDOS/sec (FAT-
    2  XENIX root      39  Plan 9          83  Linux           c4  DRDOS/sec (FAT-
    3  XENIX usr       3c  PartitionMagic  84  OS/2 hidden C:  c6  DRDOS/sec (FAT-
    4  FAT16 <32M      40  Venix 80286     85  Linux extended  c7  Syrinx         
    5  Extended        41  PPC PReP Boot   86  NTFS volume set da  Non-FS data    
    6  FAT16           42  SFS             87  NTFS volume set db  CP/M / CTOS / .
    7  HPFS/NTFS/exFAT 4d  QNX4.x          88  Linux plaintext de  Dell Utility   
    8  AIX             4e  QNX4.x 2nd part 8e  Linux LVM       df  BootIt         
    9  AIX bootable    4f  QNX4.x 3rd part 93  Amoeba          e1  DOS access     
    a  OS/2 Boot Manag 50  OnTrack DM      94  Amoeba BBT      e3  DOS R/O        
    b  W95 FAT32       51  OnTrack DM6 Aux 9f  BSD/OS          e4  SpeedStor      
    c  W95 FAT32 (LBA) 52  CP/M            a0  IBM Thinkpad hi eb  BeOS fs        
    e  W95 FAT16 (LBA) 53  OnTrack DM6 Aux a5  FreeBSD         ee  GPT            
    f  W95 Ext'd (LBA) 54  OnTrackDM6      a6  OpenBSD         ef  EFI (FAT-12/16/
   10  OPUS            55  EZ-Drive        a7  NeXTSTEP        f0  Linux/PA-RISC b
   11  Hidden FAT12    56  Golden Bow      a8  Darwin UFS      f1  SpeedStor      
   12  Compaq diagnost 5c  Priam Edisk     a9  NetBSD          f4  SpeedStor      
   14  Hidden FAT16 <3 61  SpeedStor       ab  Darwin boot     f2  DOS secondary  
   16  Hidden FAT16    63  GNU HURD or Sys af  HFS / HFS+      fb  VMware VMFS    
   17  Hidden HPFS/NTF 64  Novell Netware  b7  BSDI fs         fc  VMware VMKCORE 
   18  AST SmartSleep  65  Novell Netware  b8  BSDI swap       fd  Linux raid auto
   1b  Hidden W95 FAT3 70  DiskSecure Mult bb  Boot Wizard hid fe  LANstep        
   1c  Hidden W95 FAT3 75  PC/IX           be  Solaris boot    ff  BBT            
   1e  Hidden W95 FAT1 80  Old Minix      
   Hex code (type L to list all codes): 82
   Changed type of partition 'Linux LVM' to 'Linux swap / Solaris'
   
   Command (m for help): w
   The partition table has been altered!
   
   Calling ioctl() to re-read partition table.
   Syncing disks.
   
   ```

4. 格式化

   ```
   [root@chz ~]# mkswap /dev/sdb1 
   Setting up swapspace version 1, size = 511996 KiB
   no label, UUID=88e7c6ab-b5f5-4d61-8225-7746cb404e24
   
   [root@chz ~]# lsblk -f
   NAME            FSTYPE      LABEL UUID                                   MOUNTPOINT
   sda                                                                      
   ├─sda1          xfs               52ff1027-e9d7-427d-9f43-3a98ba708796   /boot
   └─sda2          LVM2_member       eA52jE-SFuU-BG5t-Isyw-wWdY-lj4K-xso1bu 
     ├─centos-root xfs               c18c2095-72b4-49f4-95be-a58d0a6cc2ad   /
     └─centos-swap swap              0b4a0dde-4db6-4494-a067-80f3318849f7   [SWAP]
   sdb                                                                      
   ├─sdb1          swap              88e7c6ab-b5f5-4d61-8225-7746cb404e24   
   ├─sdb2          LVM2_member       nrkgKn-zDje-od4w-k0Is-QQ8m-m3CZ-eVFsM3 
   └─sdb3          LVM2_member       qTscE4-fYC3-ctCg-3mAe-12NA-d61g-LslY83 
   
   ```

5. 启动新增分区

   ```
   [root@chz ~]# swapon /dev/sdb1 
   [root@chz ~]# swapon
   NAME      TYPE      SIZE  USED PRIO
   /dev/dm-1 partition   2G 74.3M   -2
   /dev/sdb1 partition 500M    0B   -3
   ```

6. 添加永久设置

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
   /dev/sdb1 swap swap default 0 0
   ```

## 使用文件做为swap分区

1. 写入文件

   > 写入的文件最大和if有关

   ```
   [root@chz ~]# dd if=/dev/sdb1 of=/opt/swap-file bs=1M count=1024
   500+0 records in
   500+0 records out
   524288000 bytes (524 MB) copied, 1.34818 s, 389 MB/s
   [root@chz ~]# ll -h /opt
   total 500M
   drwxr-xr-x. 2 root root    6 Oct 31  2018 rh
   -rw-r--r--. 1 root root 500M Oct 21 19:14 swap-file
   ```

2. 格式化

   ```
   [root@chz opt]# mkswap swap-file 
   mkswap: swap-file: warning: wiping old swap signature.
   Setting up swapspace version 1, size = 511996 KiB
   no label, UUID=37d97f81-88c2-4dad-967c-fa3f486f2f21
   ```

3. 启用分区

   ```
   [root@chz opt]# swapon 
   NAME      TYPE      SIZE  USED PRIO
   /dev/dm-1 partition   2G 50.3M   -2
   [root@chz opt]# swapon /opt/swap-file
   swapon: /opt/swap-file: insecure permissions 0644, 0600 suggested.
   [root@chz opt]# swapon 
   NAME           TYPE      SIZE  USED PRIO
   /opt/swap-file file      500M    0B   -3
   /dev/dm-1      partition   2G 50.3M   -2
   ```

## 删除swap分区

1. swapoff

   ```
   [root@chz ~]# swapoff /dev/sdb1 
   [root@chz ~]# swapon
   NAME      TYPE      SIZE USED PRIO
   /dev/dm-1 partition   2G 3.3M   -2
   ```

2. 删除配置文件

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
   ```

   
