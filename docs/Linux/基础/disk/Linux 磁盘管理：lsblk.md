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

- `-f`

  列出磁盘的使用情况，同时展示文件系统

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

  
