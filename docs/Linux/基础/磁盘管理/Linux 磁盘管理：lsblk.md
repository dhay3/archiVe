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
2. MAJ表示磁盘的identifier，MIN表示分区的identifier
3. 是否可以移动，1表示设备可以移动
4. 容量大小
5. 设备是否只是只读，1表示只读
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

  