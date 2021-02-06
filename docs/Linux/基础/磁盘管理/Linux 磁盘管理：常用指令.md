# 磁盘管理常用指令

> 使用`mount`命令在重启后挂载会失效，需要在`/etc/fstab`中修改

- lsblk

  查看磁盘，分区，挂载点。使用`-f`查看文件系统类型

  > sdx为SCSI，SATA，USB类型的硬盘
  >
  > hdx为IDE类型硬盘
  >
  > srx由内核分配，表示CD/DVD

  ```
  [root@chz network-scripts]# lsblk
  NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
  sda               8:0    0   20G  0 disk 
  ├─sda1            8:1    0    1G  0 part /boot
  └─sda2            8:2    0   19G  0 part 
    ├─centos-root 253:0    0   17G  0 lvm  /
    └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
  sdb               8:16   0    1G  0 disk 
  ├─sdb2            8:18   0    1K  0 part 
  └─sdb5            8:21   0   20M  0 part 
  sr0              11:0    1 1024M  0 rom  
  
  [root@chz ~]# lsblk -f
  NAME            FSTYPE      LABEL UUID                                   MOUNTPOINT
  sda                                                                      
  ├─sda1          xfs               52ff1027-e9d7-427d-9f43-3a98ba708796   /boot
  └─sda2          LVM2_member       eA52jE-SFuU-BG5t-Isyw-wWdY-lj4K-xso1bu 
    ├─centos-root xfs               c18c2095-72b4-49f4-95be-a58d0a6cc2ad   /
    └─centos-swap swap              0b4a0dde-4db6-4494-a067-80f3318849f7   [SWAP]
  sdb                                                                      
  ├─sdb1                                                                   
  ├─sdb2          LVM2_member       nrkgKn-zDje-od4w-k0Is-QQ8m-m3CZ-eVFsM3 
  └─sdb3          LVM2_member       qTscE4-fYC3-ctCg-3mAe-12NA-d61g-LslY83 
  sr0                                               
  ```

- fdisk

  MBR磁盘管理器

- gdisk

  GPT磁盘管理器

- df -h

  以可读的方式显示磁盘使用情况，和文件系统的路径，==但是不会显示尚未挂载的分区==

  ```
  [root@chz network-scripts]# df -h
  Filesystem               Size  Used Avail Use% Mounted on
  devtmpfs                 470M     0  470M   0% /dev
  tmpfs                    487M     0  487M   0% /dev/shm
  tmpfs                    487M  8.6M  478M   2% /run
  tmpfs                    487M     0  487M   0% /sys/fs/cgroup
  /dev/mapper/centos-root   17G  8.1G  9.0G  48% /
  /dev/sda1               1014M  213M  802M  21% /boot
  tmpfs                     98M   44K   98M   1% /run/user/0
  ```

  > df -T 会显示文件系统，但是不会显示没有扩容的文件系统
  >
  > ```
  > [root@chz ~]# df -hT
  > Filesystem              Type      Size  Used Avail Use% Mounted on
  > devtmpfs                devtmpfs  470M     0  470M   0% /dev
  > tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
  > tmpfs                   tmpfs     487M  8.6M  478M   2% /run
  > tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
  > /dev/mapper/centos-root xfs        17G  6.1G   11G  36% /
  > /dev/sda1               xfs      1014M  172M  843M  17% /boot
  > /dev/sdb1               ext4       93M  1.6M   85M   2% /opt
  > tmpfs                   tmpfs      98M   24K   98M   1% /run/user/0
  > /dev/sr0                iso9660   4.4G  4.4G     0 100% /run/media/root/CentOS 7 x86_64
  > [root@chz ~]# lsblk
  > NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
  > sda               8:0    0   20G  0 disk 
  > ├─sda1            8:1    0    1G  0 part /boot
  > └─sda2            8:2    0   19G  0 part 
  >   ├─centos-root 253:0    0   17G  0 lvm  /
  >   └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
  > sdb               8:16   0    1G  0 disk 
  > └─sdb1            8:17   0 1023M  0 part /opt
  > sr0              11:0    1  4.4G  0 rom  /run/media/root/CentOS 7 x86_64
  > ```

- du

  用于显示目录或文件的大小，默认会递归当前目录，默认以字节为单位

  ```
  root in /opt λ du
  16      ./lsd-0.18.0-x86_64-unknown-linux-gnu/lsd-0.18.0-x86_64-unknown-linux-gnu/autocomplete
  2496    ./lsd-0.18.0-x86_64-unknown-linux-gnu/lsd-0.18.0-x86_64-unknown-linux-gnu
  2500    ./lsd-0.18.0-x86_64-unknown-linux-gnu
  2504    .                                                                                                                                                /0.0s
  root in /opt λ du -s
  2504    .            
  
  root in /var/www λ du -h /var/www/* #使用-h参数显示可读信息
  20K     /var/www/html               
  ```

- mount

  将分区挂载`mount <source> <directory>`

  ```
  mount /dev/hda1 /mnt
  mount -a  #挂载/etc/fstab中的所有filesystem
  ```

- umount

  将分区卸载`umount <dev> `

  ```
  umount /dev/hda1 
  ```

- dd

  将if中的内容写入of，一般用于制作启动盘  

  ```
  dd if=<input file> of=<output file> 
  ```

  > 如果不指定`if`,将采用stdin

- free

  查看系统内存和交换分区的使用情况

  ```
  [root@chz network-scripts]# free -h
                total        used        free      shared  buff/cache   available
  Mem:           972M        570M         64M         19M        337M        231M
  Swap:          2.0G        371M        1.6G
  ```

- resize2fs

  重新生成文件系统，到分区的最大容量

  ```
  [root@chz ~]# resize2fs /dev/sdb1
  resize2fs 1.42.9 (28-Dec-2013)
  Filesystem at /dev/sdb1 is mounted on /opt; on-line resizing required
  old_desc_blocks = 1, new_desc_blocks = 8
  The filesystem on /dev/sdb1 is now 1047552 blocks long.
  
  ```

  























