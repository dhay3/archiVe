# growpart

参考：

https://help.aliyun.com/document_detail/111738.htm?spm=a2c4g.11186623.2.30.37e84eb7XegIi3#section-vxq-3tw-dhb

是一个分区扩展工具，近最大容量进行扩容

```
[root@chz ~]# yum search growpart
Loaded plugins: fastestmirror, langpacks
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.aliyun.com
============================================================= N/S matched: growpart ==============================================================
cloud-utils-growpart.noarch : Script for growing a partition
```

syntax：`growpart <disk> <partition-number>`

## 参数

- -N | --dry-run

  不真正的执行，用于检测

## 例子

现在要对sdb1分区进行扩容

```
[root@chz /]# lsblk
NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda               8:0    0   20G  0 disk 
├─sda1            8:1    0    1G  0 part /boot
└─sda2            8:2    0   19G  0 part 
  ├─centos-root 253:0    0   17G  0 lvm  /
  └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
sdb               8:16   0    1G  0 disk 
└─sdb1            8:17   0  100M  0 part 
sr0              11:0    1 1024M  0 rom  

[root@chz /]# growpart /dev/sdb 1
CHANGED: partition=1 start=2048 old: size=204800 end=206848 new: size=2095071 end=2097119
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
```

最后对fs进行扩容

```
resize2fs <PartitionName>
```

