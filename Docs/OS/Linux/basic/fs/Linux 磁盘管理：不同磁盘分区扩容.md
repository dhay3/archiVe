# 不同磁盘分区扩容

参考：

http://blog.51yip.com/linux/1778.html

有些时候磁盘不够用了，我们需要添加新的磁盘，对现有的分区进行扩容。

有两种方法，普通方式通过对指定目录挂载，LVM方式

## 普通方法

这种方式不是正真意义上的扩容，只是挂载，且如果挂载的磁盘原本有数据，会导致数据因为更换磁盘而消失。所以==一般需要创建一个空文件做为挂载点。==

```
[root@chz /]# lsblk
NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda               8:0    0   20G  0 disk 
├─sda1            8:1    0    1G  0 part /boot
└─sda2            8:2    0   19G  0 part 
  ├─centos-root 253:0    0   17G  0 lvm  /
  └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
sdb               8:16   0    1G  0 disk 
└─sdb1            8:17   0 1023M  0 part 
sr0              11:0    1 1024M  0 rom  

[root@chz /]# stat /opt/t
  File: ‘/opt/t’
  Size: 6         	Blocks: 0          IO Block: 4096   directory
Device: fd00h/64768d	Inode: 16785990    Links: 2
Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:usr_t:s0
Access: 2021-03-09 10:43:20.828013078 +0800
Modify: 2021-03-09 10:43:20.828013078 +0800
Change: 2021-03-09 10:43:20.828013078 +0800
 Birth: -
```

---



对`/opt/t`目录进行扩容

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
   └─sdb1            8:17   0 1023M  0 part /opt
   sdc               8:32   0    1G  0 disk 
   sr0              11:0    1 1024M  0 rom  
   
   [root@chz /]# stat /etc/
     File: ‘/etc/’
     Size: 12288     	Blocks: 32         IO Block: 4096   directory
   Device: fd00h/64768d	Inode: 16777281    Links: 150
   Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
   Context: system_u:object_r:etc_t:s0
   Access: 2021-03-09 10:00:44.087014311 +0800
   Modify: 2021-03-09 10:01:21.895014976 +0800
   Change: 2021-03-09 10:01:21.895014976 +0800
    Birth: -
   
   
   ```

2. 挂载分区，这里的sdb1已格式化。注意查看Device行的编码替换了

   ```
   [root@chz /]# mount /dev/sdb1 /opt/t
   [root@chz /]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sdb               8:16   0    1G  0 disk 
   └─sdb1            8:17   0 1023M  0 part /opt/t
   sr0              11:0    1 1024M  0 rom  
   
   [root@chz /]# stat /opt/t
     File: ‘/opt/t’
     Size: 4096      	Blocks: 8          IO Block: 4096   directory
   Device: 811h/2065d	Inode: 2           Links: 3
   Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
   Context: system_u:object_r:unlabeled_t:s0
   Access: 2021-03-09 10:43:53.000000000 +0800
   Modify: 2021-03-09 10:43:53.000000000 +0800
   Change: 2021-03-09 10:43:53.000000000 +0800
    Birth: -
   
   ```

## LVM

使用LVM是最通用且方便的，在==LVM中只有逻辑卷组才能被使用==。

1. 划分区，并将分区类型修改为LVM

   ```
   [root@chz /]# fdisk -l /dev/sdb
   
   Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
   Units = sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes
   Disk label type: dos
   Disk identifier: 0xf78e552c
   
      Device Boot      Start         End      Blocks   Id  System
   /dev/sdb1            2048      206847      102400   8e  Linux LVM
   /dev/sdb2          206848      411647      102400   8e  Linux LVM
   /dev/sdb3          411648      616447      102400   8e  Linux LVM
   
   ```

2. 如果原本分区有filesys，使用`wipefs`清除filesys

   ```
   [root@chz /]# wipefs -a /dev/sdb1
   /dev/sdb1: 2 bytes were erased at offset 0x00000438 (ext4): 53 ef
   ```

3. 创建物理卷

   ```
   [root@chz /]# pvcreate /dev/sdb1 /dev/sdb2 /dev/sdb3
   WARNING: ext4 signature detected on /dev/sdb3 at offset 1080. Wipe it? [y/n]: y
     Wiping ext4 signature on /dev/sdb3.
     Physical volume "/dev/sdb1" successfully created.
     Physical volume "/dev/sdb2" successfully created.
     Physical volume "/dev/sdb3" successfully created.
   [root@chz /]# pvs
     PV         VG     Fmt  Attr PSize   PFree  
     /dev/sda2  centos lvm2 a--  <19.00g      0 
     /dev/sdb1         lvm2 ---  100.00m 100.00m
     /dev/sdb2         lvm2 ---  100.00m 100.00m
     /dev/sdb3         lvm2 ---  100.00m 100.00m
   ```

4. 创建卷组

   ```
   [root@chz /]# vgcreate vg1 /dev/sdb1 /dev/sdb2
     Volume group "vg1" successfully created
   [root@chz /]# vgcreate vg2 /dev/sdb3
     Volume group "vg2" successfully created
   [root@chz /]# vgs
     VG     #PV #LV #SN Attr   VSize   VFree  
     centos   1   2   0 wz--n- <19.00g      0 
     vg1      2   0   0 wz--n- 192.00m 192.00m
     vg2      1   0   0 wz--n-  96.00m  96.00m
   ```

4. 创建逻辑卷

   ```
   [root@chz /]# lvcreate -L 100 vg1 -n lv01
     Logical volume "lv01" created.
   [root@chz /]# lvcreate -l 100%free vg1 -n lv02
     Logical volume "lv02" created.
   [root@chz /]# vgdisplay vg2 | grep "VG Size" | awk '{print $3}' | xargs -i  lvcreate -L {} vg2  -n lv03
     Logical volume "lv03" created.
   [root@chz /]# lvs
     LV   VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
     root centos -wi-ao---- <17.00g                                                    
     swap centos -wi-ao----   2.00g                                                    
     lv01 vg1    -wi-a----- 100.00m                                                    
     lv02 vg1    -wi-a-----  92.00m                                                    
     lv03 vg2    -wi-a-----  96.00m   
   ```

5. 格式化逻辑卷

   ```
   [root@chz /]# mkfs.ext4 /dev/vg1/lv01
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

5. 挂载

   ```
   [root@chz /]# mount /dev/vg1/lv01 /opt
   ```

6. 永久挂载

   





