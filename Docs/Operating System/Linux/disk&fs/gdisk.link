# gdisk

`gdisk`是`fdisk`的形式操作GPT分区表的磁盘。

如果磁盘的分区表是MBR会要求转为GPT

```
[root@chz etc]# gdisk /dev/sdb
GPT fdisk (gdisk) version 0.8.10

Partition table scan:
  MBR: MBR only
  BSD: not present
  APM: not present
  GPT: not present


***************************************************************
Found invalid GPT and valid MBR; converting MBR to GPT format
in memory. THIS OPERATION IS POTENTIALLY DESTRUCTIVE! Exit by
typing 'q' if you don't want to convert your MBR partitions
to GPT format!
***************************************************************

```

原始状态

```
[root@chz /]# fdisk -l /dev/sdb

Disk /dev/sdb: 1073 MB, 1073741824 bytes, 2097152 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk label type: dos
Disk identifier: 0x0006b41d

   Device Boot      Start         End      Blocks   Id  System
/dev/sdb1               1      195312       97656   83  Linux

[root@chz /]# lsblk
NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda               8:0    0   20G  0 disk 
├─sda1            8:1    0    1G  0 part /boot
└─sda2            8:2    0   19G  0 part 
  ├─centos-root 253:0    0   17G  0 lvm  /
  └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
sdb               8:16   0    1G  0 disk 
└─sdb1            8:17   0 95.4M  0 part 
sr0              11:0    1 1024M  0 rom  

```

注意这里的分区表是MBR类型（label type dos），==一个磁盘只能存在一种分区表，所以需要先删除之前的分区==

```
Command (? for help): n
Partition number (2-128, default 2): 
First sector (195313-2097118, default = 195313) or {+-}size{KMGTP}: 
Last sector (195313-2097118, default = 2097118) or {+-}size{KMGTP}: +200M
Current type is 'Linux filesystem'
Hex code or GUID (L to show codes, Enter = 8300): 
Changed type of partition to 'Linux filesystem'

Command (? for help): w
Warning! Main partition table overlaps the first partition by 33 blocks!
You will need to delete this partition or resize it in another utility.
Aborting write of new partition table.

Command (? for help): d
Partition number (1-2): 1

Command (? for help): w

Final checks complete. About to write GPT data. THIS WILL OVERWRITE EXISTING
PARTITIONS!!

Do you want to proceed? (Y/N): y
OK; writing new GUID partition table (GPT) to /dev/sdb.
The operation has completed successfully.

```

