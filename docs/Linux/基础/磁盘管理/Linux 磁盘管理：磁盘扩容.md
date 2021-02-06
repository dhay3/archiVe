# 磁盘扩容

> 整个过程无需修改`/etc/fstab`，因为重新挂载后同样也需要编辑持久挂载。
>
> 扩容时千万不能格式化filesys，否则原由的数据会消失

1. 使用sdb1模拟现有挂载的100G

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
   Disk identifier: 0xcc083120
   
      Device Boot      Start         End      Blocks   Id  System
   /dev/sdb1            2048      206847      102400   83  Linux
   
   [root@chz opt]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sdb               8:16   0    1G  0 disk 
   └─sdb1            8:17   0  100M  0 part /opt
   sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
   
   ```

2. 卸载

   ```
   [root@chz /]# umount /dev/sdb1
   [root@chz /]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sdb               8:16   0    1G  0 disk 
   └─sdb1            8:17   0  100M  0 part 
   sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
   ```

3. 删除分区

   ```
   [root@chz /]# fdisk /dev/sdb
   Welcome to fdisk (util-linux 2.23.2).
   
   Changes will remain in memory only, until you decide to write them.
   Be careful before using the write command.
   
   
   Command (m for help): d
   Selected partition 1
   Partition 1 is deleted
   
   Command (m for help): w
   The partition table has been altered!
   
   Calling ioctl() to re-read partition table.
   Syncing disks.
   [root@chz /]# fdisk -l /dev/sdb
   
   Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0xcc083120
   
      Device Boot      Start         End      Blocks   Id  System
   ```

4. 新建分区并挂载

   ```
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
   Using default response p 
   Partition number (1-4, default 1): 
   First sector (2048-2097151, default 2048): 
   Using default value 2048
   Last sector, +sectors or +size{K,M,G} (2048-2097151, default 2097151): 
   Partition 1 of type Linux and of size 2097151 is set
   
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
      p   primary (1 primary, 0 extended, 3 free) 
      e   extended
   Select (default p):  
   [root@chz ~]# mount /dev/sdb1 /opt
   ```

5. 扩容文件系统

   ```
   [root@chz ~]# resize2fs /dev/sdb1
   resize2fs 1.42.9 (28-Dec-2013)
   Filesystem at /dev/sdb1 is mounted on /opt; on-line resizing required
   old_desc_blocks = 1, new_desc_blocks = 8
   The filesystem on /dev/sdb1 is now 1047552 blocks long.
   
   [root@chz ~]# df -hT
   Filesystem              Type      Size  Used Avail Use% Mounted on
   devtmpfs                devtmpfs  470M     0  470M   0% /dev
   tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
   tmpfs                   tmpfs     487M  8.5M  478M   2% /run
   tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
   /dev/mapper/centos-root xfs        17G  6.1G   11G  36% /
   /dev/sda1               xfs      1014M  172M  843M  17% /boot
   /dev/sdb1               ext4      988M  2.8M  942M   1% /opt
   tmpfs                   tmpfs      98M   24K   98M   1% /run/user/0
   /dev/sr0                iso9660   4.4G  4.4G     0 100% /run/media/root/CentOS 7 x86_64
   ```



