# blkid

展示块设备的属性，可以使用`lsblk --fs`替代

```
[root@chz opt]# blkid
/dev/sda1: UUID="52ff1027-e9d7-427d-9f43-3a98ba708796" TYPE="xfs" 
/dev/sda2: UUID="eA52jE-SFuU-BG5t-Isyw-wWdY-lj4K-xso1bu" TYPE="LVM2_member" 
/dev/sdb1: UUID="eupqtA-26Oe-RVXH-evXk-lRwN-flPd-wF8avK" TYPE="LVM2_member" 
/dev/sdb2: UUID="NhBE0y-Qrt6-l9zy-4tnL-hKQf-lvyw-NoyiAW" TYPE="LVM2_member" 
/dev/sdb3: UUID="5X5jb8-QswZ-8S9F-lttD-nYgI-adXX-0I1Arc" TYPE="LVM2_member" 
/dev/mapper/centos-root: UUID="c18c2095-72b4-49f4-95be-a58d0a6cc2ad" TYPE="xfs" 
/dev/mapper/centos-swap: UUID="0b4a0dde-4db6-4494-a067-80f3318849f7" TYPE="swap" 
/dev/mapper/vg1-lv01: UUID="95a68178-1d5b-46aa-93fe-230c5664e1c3" TYPE="ext4"
```

常用

```
#磁盘分区表
[root@chz opt]# blkid /dev/sda
/dev/sda: PTTYPE="dos" 

#分区的属性
[root@chz opt]# blkid /dev/sda1
/dev/sda1: UUID="52ff1027-e9d7-427d-9f43-3a98ba708796" TYPE="xfs"

#根据UUID查询分区
[root@chz opt]# blkid -U 52ff1027-e9d7-427d-9f43-3a98ba708796
/dev/sda1
```

使用awk对`/etc/fstab`写入

```
[root@chz opt]# blkid /dev/vg1/lv01 | awk -F '[= ]' '{printf "UUID=%s \t /opt%s \t defaults \t 0 \t 0\n",$3,$5 }' >> /etc/fstab
[root@chz opt]# cat /etc/fstab

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
UUID="95a68178-1d5b-46aa-93fe-230c5664e1c3" 	 "ext4" 	 defualts 	 0 	 0
```

