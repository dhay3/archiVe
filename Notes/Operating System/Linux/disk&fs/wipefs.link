# Linux 磁盘管理：wipefs

wipefs用于抹除分区上的fs

```
[root@chz /]# lsblk 
NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda               8:0    0   20G  0 disk 
├─sda1            8:1    0    1G  0 part /boot
└─sda2            8:2    0   19G  0 part 
  ├─centos-root 253:0    0   17G  0 lvm  /
  └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
sdb               8:16   0  102M  0 disk 
└─sdb1            8:17   0  101M  0 part /opt/t
sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64

```

如果wipe没有带任何参数，不会擦除fs，只打印分区信息

```
[root@chz t]# wipefs /dev/sdb1
offset               type
----------------------------------------------------------------
0x438                ext4   [filesystem]
                     UUID:  665e4ddc-41ad-4a7c-8baf-242616261c5c
```

使用-a参数擦除fs，如果分区正在使用需要使用--force

```
[root@chz t]# wipefs -a --force /dev/sdb1
/dev/sdb1: 2 bytes were erased at offset 0x00000438 (ext4): 53 ef
```

