# Linux 磁盘管理：lsblk

lsblk用于列出所有的块设备(磁盘，也会展示分区)，默认不会展示RAM。

```
root in / λ lsblk 
NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda      8:0    0   20G  0 disk 
├─sda1   8:1    0   19G  0 part /
├─sda2   8:2    0    1K  0 part 
└─sda5   8:5    0  975M  0 part [SWAP]
sr0     11:0    1 1024M  0 rom          
```

1. 设备块名字

2. MAJ(major)表示磁盘的identifier，MIN(minor)表示分区的identifier

   我们可以在`/dev/block`下查看具体的信息

   ```
   [root@chz block]# ll 
   total 0
   lrwxrwxrwx. 1 root root 6 Mar 16 09:50 11:0 -> ../sr0
   lrwxrwxrwx. 1 root root 7 Mar 16 09:50 253:0 -> ../dm-0
   lrwxrwxrwx. 1 root root 7 Mar 16 09:50 253:1 -> ../dm-1
   lrwxrwxrwx. 1 root root 6 Mar 16 09:50 8:0 -> ../sda
   lrwxrwxrwx. 1 root root 7 Mar 16 09:50 8:1 -> ../sda1
   lrwxrwxrwx. 1 root root 7 Mar 16 09:50 8:2 -> ../sda2
   [root@chz block]# lsblk
   NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
   sda               8:0    0   20G  0 disk 
   ├─sda1            8:1    0    1G  0 part /boot
   └─sda2            8:2    0   19G  0 part 
     ├─centos-root 253:0    0   17G  0 lvm  /
     └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
   sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
   ```

3. 是否可以移动，1表示设备可以移动

4. 容量大小

5. ==Rotation，1表示可以旋转就是机械硬盘，0表示不可以旋转就是固态硬盘(固态的磁盘不会转)，但是不具备通用性==

6. disk表示磁盘，part表示分区，rom表示只读存储

7. 挂载点

## 常用参数

- `-o`

  输出指定列，可以使用`+UUID`表示在默认输出添加额外的列。使用`--help`输出可用的列

  ```
  cpl in /mnt λ lsblk -o +label
  NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT                                 LABEL
  loop0         7:0    0 162.9M  1 loop /var/lib/snapd/snap/gnome-3-28-1804/145    
  loop1         7:1    0  32.1M  1 loop /var/lib/snapd/snap/snapd/12057            
  loop2         7:2    0  55.4M  1 loop /var/lib/snapd/snap/core18/2066            
  loop3         7:3    0  76.8M  1 loop /var/lib/snapd/snap/termius-app/77         
  loop4         7:4    0  65.1M  1 loop /var/lib/snapd/snap/gtk-common-themes/1515 
  loop5         7:5    0  32.1M  1 loop /var/lib/snapd/snap/snapd/11841            
  loop6         7:6    0  76.8M  1 loop /var/lib/snapd/snap/termius-app/76         
  sda           8:0    0 238.5G  0 disk                                            
  ├─sda1        8:1    0   3.6G  0 part                                            Kali Live
  ├─sda2        8:2    0   736K  0 part                                            
  ├─sda3        8:3    0  56.4G  0 part                                            persistence
  ├─sda4        8:4    0    60G  0 part                                            qemu-disk
  └─sda5        8:5    0    80G  0 part                                            store
  nvme0n1     259:0    0 476.9G  0 disk                                            
  ├─nvme0n1p1 259:1    0   260M  0 part                                            SYSTEM_DRV
  ├─nvme0n1p2 259:2    0    16M  0 part                                            
  ├─nvme0n1p3 259:3    0   100G  0 part                                            WIN
  ├─nvme0n1p4 259:4    0  1000M  0 part                                            WINRE_DRV
  ├─nvme0n1p5 259:5    0     2G  0 part /boot/efi                                  NO_LABEL
  ├─nvme0n1p6 259:6    0     4G  0 part [SWAP]                                     
  └─nvme0n1p7 259:7    0   250G  0 part /                                         
  ```

- `-f`

  列出磁盘的使用情况，同时展示文件系统==替代blkid==

  ```
  root in / λ lsblk -f
  NAME   FSTYPE FSVER LABEL UUID                                 FSAVAIL FSUSE% MOUNTPOINT
  sda                                                                           
  ├─sda1 ext4   1.0         b0c99072-74fd-4a32-8dbf-f12652a25e67    4.3G    72% /
  ├─sda2                                                                        
  └─sda5 swap   1           2ac953bd-7edb-4b1f-98cc-30b1437cf54f                [SWAP]
  ```

- `-J`

  以JSON的格式输出

  ```
  root in / λ lsblk -J
  {
     "blockdevices": [
        {"name":"sda", "maj:min":"8:0", "rm":false, "size":"20G", "ro":false, "type":"disk", "mountpoint":null,
           "children": [
              {"name":"sda1", "maj:min":"8:1", "rm":false, "size":"19G", "ro":false, "type":"part", "mountpoint":"/"},
              {"name":"sda2", "maj:min":"8:2", "rm":false, "size":"1K", "ro":false, "type":"part", "mountpoint":null},
              {"name":"sda5", "maj:min":"8:5", "rm":false, "size":"975M", "ro":false, "type":"part", "mountpoint":"[SWAP]"}
           ]
        },
        {"name":"sr0", "maj:min":"11:0", "rm":true, "size":"1024M", "ro":false, "type":"rom", "mountpoint":null}
     ]
  }                          
  ```

  
