# Linux 磁盘管理：parted

> ==在parted中的命令是立即生效的==

虽然我们可以使用 fdisk命令对硬盘进行快速的分区，但对高于 2TB 的硬盘分区(GPT)，此命令却无能为力，此时就需要使用==parted 命令或是 gdisk（会将MBR转为GPT）==。

syntax：`parted [option] [device]`

==如果没有指定块设备默认使用第一块==

## 参数

- `-l | --list`

  显示所有块设备的分区

  ```
  root in ~ λ parted -l
  Model: VMware, VMware Virtual S (scsi)
  Disk /dev/sda: 21.5GB
  Sector size (logical/physical): 512B/512B
  Partition Table: msdos
  Disk Flags: 
  
  Number  Start   End     Size    Type      File system     Flags
   1      1049kB  20.5GB  20.4GB  primary   ext4            boot
   2      20.5GB  21.5GB  1022MB  extended
   5      20.5GB  21.5GB  1022MB  logical   linux-swap(v1)
  ```

- `-s`

  不进入交互模式

  ```
  root in ~ λ parted -s
  root in ~ λ parted  
  GNU Parted 3.3
  Using /dev/sda
  Welcome to GNU Parted! Type 'help' to view a list of commands.
                                                                         (parted) ^C
  ```

## command

- print

  等价于`-l`参数

  ```
  (parted) print                                                            
  Model: VMware, VMware Virtual S (scsi)
  Disk /dev/sda: 21.5GB
  Sector size (logical/physical): 512B/512B
  Partition Table: msdos
  Disk Flags: 
  
  Number  Start   End     Size    Type     File system  Flags
   1      1049kB  1075MB  1074MB  primary  xfs          boot
   2      1075MB  21.5GB  20.4GB  primary               lvm
  ```

- mklabel

  修改磁盘的分区表，msdos表示mrb，gpt表示gpt。由于修改了分区表导致原来的数据丢失，分区也消失了。==所以需要修改`/etc/fstab`文件==

  ```
  (parted) print
  Model: VMware, VMware Virtual S (scsi)
  Disk /dev/sda: 21.5GB
  Sector size (logical/physical): 512B/512B
  Partition Table: msdos
  Disk Flags: 
  
  Number  Start   End     Size    Type     File system  Flags
   1      1049kB  1075MB  1074MB  primary  xfs          boot
   2      1075MB  21.5GB  20.4GB  primary               lvm
  
  (parted) mklabel gpt
  Warning: Partition(s) on /dev/sda are being used.
  Ignore/Cancel? ignore                                                     
  Warning: The existing disk label on /dev/sda will be destroyed and all data on
  this disk will be lost. Do you want to continue?
  Yes/No? yes                                                               
  Error: Partition(s) 1, 2 on /dev/sda have been written, but we have been unable
  to inform the kernel of the change, probably because it/they are in use.  As a
  result, the old partition(s) will remain in use.  You should reboot now before
  making further changes.
  Ignore/Cancel? ignore                                                     
  (parted) print                                                            
  Model: VMware, VMware Virtual S (scsi)
  Disk /dev/sda: 21.5GB
  Sector size (logical/physical): 512B/512B
  Partition Table: gpt
  Disk Flags: 
  
  Number  Start  End  Size  File system  Name  Flags
  
  ```

- mkpart

  生成分区，依次选择分区的名字，文件系统，起始点，终点。==注意这里少了一个Type字段，因为在GPT中没有主分区和扩展分区之分==

  ```
  (parted) mkpart                                                           
  Partition name?  []? disk1                                                
  File system type?  [ext2]? ext4                                           
  Start?                                                                    
  Start? 1MB                                                                
  End? 5GB
  Error: Error informing the kernel about modifications to partition /dev/sda1 --
  Device or resource busy.  This means Linux won't know about any changes you made
  to /dev/sda1 until you reboot -- so you shouldn't mount it or use it in any way
  before rebooting.
  Ignore/Cancel? ignore                                                     
  Error: Partition(s) 2 on /dev/sda have been written, but we have been unable to
  inform the kernel of the change, probably because it/they are in use.  As a
  result, the old partition(s) will remain in use.  You should reboot now before
  making further changes.
  Ignore/Cancel? ignore                                                     
  (parted) print                                                            
  Model: VMware, VMware Virtual S (scsi)
  Disk /dev/sda: 21.5GB
  Sector size (logical/physical): 512B/512B
  Partition Table: gpt
  Disk Flags: 
  
  Number  Start   End     Size    File system  Name   Flags
   1      1049kB  5000MB  4999MB  xfs          disk1
  
  (parted)                                                         
  ```

- resizepart

  改变分区大小

  ```
  (parted) resize 
  Error: The resize command has been removed in parted 3.0
  (parted) resizepart 
  Partition number? 1                                                       
  End?  [100MB]? 50MB                                                       
  Warning: Shrinking a partition can cause data loss, are you sure you want to continue?
  Yes/No? yes                                                               
  (parted) p                                                                
  Model: VMware, VMware Virtual S (scsi)
  Disk /dev/sdb: 1074MB
  Sector size (logical/physical): 512B/512B
  Partition Table: gpt
  Disk Flags: 
  
  Number  Start   End     Size    File system  Name    Flags
   1      17.4kB  50.0MB  50.0MB               disk01
  ```

- rm

  删除分区



