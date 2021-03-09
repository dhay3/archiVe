# 同磁盘分区扩容

> 整个过程无需修改`/etc/fstab`或格式化fs
>
> 一旦对新的分区格式化后会导致原分区中的内容丢失

有些时候我们对一块磁盘分区，但是还留下了一部分。当我们想要扩容改磁盘上的某一个分区时可以使用如下方法。

==对sdb1分区进行扩容==

1. 原有状态

   ```
   [root@chz opt]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sdb               8:16   0    1G  0 disk 
   └─sdb1            8:17   0  100M  0 part /opt
   sr0              11:0    1 1024M  0 rom  
   [root@chz opt]# ls /opt
   1  10  2  3  4  5  6  7  8  9  lost+found
   
   [root@chz opt]# df -hT
   Filesystem              Type      Size  Used Avail Use% Mounted on
   devtmpfs                devtmpfs  470M     0  470M   0% /dev
   tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
   tmpfs                   tmpfs     487M  8.6M  478M   2% /run
   tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
   /dev/mapper/centos-root xfs        17G  4.6G   13G  27% /
   /dev/sda1               xfs      1014M  172M  843M  17% /boot
   tmpfs                   tmpfs      98M  4.0K   98M   1% /run/user/42
   tmpfs                   tmpfs      98M   16K   98M   1% /run/user/0
   /dev/sdb1               ext4       93M  1.6M   85M   2% /opt
   ```

2. 删除原有的分区

   ```
   [root@chz opt]# fdisk /dev/sdb
   Welcome to fdisk (util-linux 2.23.2).
   
   Changes will remain in memory only, until you decide to write them.
   Be careful before using the write command.
   
   
   Command (m for help): d
   Selected partition 1
   Partition 1 is deleted
   
   Command (m for help): w
   The partition table has been altered!
   
   Calling ioctl() to re-read partition table.
   
   WARNING: Re-reading the partition table failed with error 16: Device or resource busy.
   The kernel still uses the old table. The new table will be used at
   the next reboot or after you run partprobe(8) or kpartx(8)
   Syncing disks.
   ```

3. 添加分区

   ```
   [root@chz opt]# fdisk /dev/sdb
   Welcome to fdisk (util-linux 2.23.2).
   
   Changes will remain in memory only, until you decide to write them.
   Be careful before using the write command.
   
   
   Command (m for help): p
   
   Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0xf78e552c
   
      Device Boot      Start         End      Blocks   Id  System
   
   Command (m for help): n
   Partition type:
      p   primary (0 primary, 0 extended, 4 free)
      e   extended
   Select (default p): 
   Using default response p
   Partition number (1-4, default 1): 
   First sector (2048-2097151, default 2048): 
   Using default value 2048
   Last sector, +sectors or +size{K,M,G} (2048-2097151, default 2097151): 
   Using default value 2097151
   Partition 1 of type Linux and of size 1023 MiB is set
   
   Command (m for help): w
   The partition table has been altered!
   
   Calling ioctl() to re-read partition table.
   
   WARNING: Re-reading the partition table failed with error 16: Device or resource busy.
   The kernel still uses the old table. The new table will be used at
   the next reboot or after you run partprobe(8) or kpartx(8)
   Syncing disks.
   ```

4. partprobe告诉OS分区改变，也可以等待下次重启后生效

   ```
   [root@chz opt]# partprobe 
   ```

5. 扩容文件系统

   ```
   [root@chz opt]# df -hT
   Filesystem              Type      Size  Used Avail Use% Mounted on
   devtmpfs                devtmpfs  470M     0  470M   0% /dev
   tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
   tmpfs                   tmpfs     487M  8.6M  478M   2% /run
   tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
   /dev/mapper/centos-root xfs        17G  4.6G   13G  27% /
   /dev/sda1               xfs      1014M  172M  843M  17% /boot
   tmpfs                   tmpfs      98M  4.0K   98M   1% /run/user/42
   tmpfs                   tmpfs      98M   16K   98M   1% /run/user/0
   /dev/sdb1               ext4       93M  1.6M   85M   2% /opt
   [root@chz opt]# resize2fs /dev/sdb1
   resize2fs 1.42.9 (28-Dec-2013)
   Filesystem at /dev/sdb1 is mounted on /opt; on-line resizing required
   old_desc_blocks = 1, new_desc_blocks = 8
   The filesystem on /dev/sdb1 is now 1047552 blocks long.
   
   [root@chz opt]# df -hT
   Filesystem              Type      Size  Used Avail Use% Mounted on
   devtmpfs                devtmpfs  470M     0  470M   0% /dev
   tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
   tmpfs                   tmpfs     487M  8.6M  478M   2% /run
   tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
   /dev/mapper/centos-root xfs        17G  4.6G   13G  27% /
   /dev/sda1               xfs      1014M  172M  843M  17% /boot
   tmpfs                   tmpfs      98M  4.0K   98M   1% /run/user/42
   tmpfs                   tmpfs      98M   16K   98M   1% /run/user/0
   /dev/sdb1               ext4      988M  2.7M  942M   1% /opt
   ```

   