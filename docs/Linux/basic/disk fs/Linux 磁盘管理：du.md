# Linux 磁盘管理：du

`du`命令用于查看文件的大小(默认以block size为最小显示单位)

默认会遍历当前目录下的所有==目录(文件夹非文件)==，如果需要查看文件的大小需要使用`*`通配符

syntax：`du [options] [file]`

- -a

  ==du默认只输目录的大小以KB为单位==，可以使用该参数输出文件或是通配符`*`

  ```bash
  cpl in ~/.ssh λ du 
  20	.
  cpl in ~/.ssh λ du -a
  4	./known_hosts
  4	./id_rsa.pub
  4	./id_rsa
  4	./config
  20	.
  
  cpl in ~ λ du * | more
  4       Applications/.trash
  8       Applications
  1345556 Desktop
  ```

- -s

  统计当前目录下所有文件大小的总计

  ```
  cpl in ~/note/docs/Linux/基础/disk on master ● λ du -sh
  132K
  cpl in ~/note/docs/Linux/基础/disk on master ● λ du -sh *
  4.0K    Linux 磁盘管理：blkid.md
  4.0K    Linux 磁盘管理： dd.md
  4.0K    Linux 磁盘管理：du.md
  16K     Linux 磁盘管理：fdisk.md
  4.0K    Linux 磁盘管理：fsck.md
  4.0K    Linux 磁盘管理：gdisk.md
  4.0K    Linux 磁盘管理：growpart.md
  ....
  ```

- -h | --human-readable

  以K，M，G为单位显示

  ```
  cpl in ~/note on master ● ● λ du -hd 1
  352M	./.git
  49M	./imgs
  407M	./docs
  807M	.
  ```

- -B

  指定block size，==但是文件大小不足时以最小单位显示==

  ```bash
  cpl in ~/note on master ● ● λ du -BM -d 1
  352M	./.git
  49M	./imgs
  407M	./docs
  807M	.
  cpl in ~/note on master ● ● λ du -BG -d 1
  1G	./.git
  1G	./imgs
  1G	./docs
  1G	.
  ```

- -d | --max--depth=N

  输出指定目录深度的内容，当前目录从0开始

  ```bash
  cpl in ~/note on master ● ● λ du -d 0
  825360	.
  cpl in ~/note on master ● ● λ du -d 1
  359776	./.git
  49300	./imgs
  416264	./docs
  825360	.
  ```

- --exclude=pattern

  使用posix regex，过滤指定的文件

  ```
  du --exclude='*.o'
  ```

- `--files0-form=E`

  从文件中读file names，`-`表示从stdin中读，可以结合find一起使用

  ```
  find . -name "*txt" -exec echo -n -e {}"\0" \; | du -hc --files0-from=-
  ```

  

## du vs stat

> 如果用`ls -l`看文件夹，是不会显示真实的大小，只会显示block size

1. stat默认以实际存储的数据显示(读取inode信息)，du以block size为单位显示。如果实际大小小于block size时以block size取整显示(因为filesys只能操作block)

   ```
   cpl in ~ λ du -h /etc/resolv.conf
   4.0K    /etc/resolv.conf
   cpl in ~ λ stat /etc/resolv.conf
     File: /etc/resolv.conf
     Size: 258             Blocks: 8          IO Block: 4096   regular file
   Device: 10307h/66311d   Inode: 12845122    Links: 1
   Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
   Access: 2021-07-13 09:12:32.464541631 +0800
   Modify: 2021-07-13 09:12:32.464541631 +0800
   Change: 2021-07-13 09:12:32.464541631 +0800
    Birth: 2021-07-13 09:12:32.464541631 +0800
   ```

2. stat会被dd创建的假文件欺骗，但是du不会

   ```
   cpl in /tmp λ dd if=/dev/zero of=test bs=1G seek=100 count=0
   0+0 records in
   0+0 records out
   0 bytes copied, 6.9299e-05 s, 0.0 kB/s
   cpl in /tmp λ ll test
   .rw-r--r-- cpl cpl 100 GB Tue Nov 30 21:23:18 2021  test
   cpl in /tmp λ stat test
     File: test
     Size: 107374182400    Blocks: 0          IO Block: 4096   regular file
   Device: 0,36    Inode: 269         Links: 1
   Access: (0644/-rw-r--r--)  Uid: ( 1000/     cpl)   Gid: ( 1000/     cpl)
   Access: 2021-11-30 21:23:18.281364762 +0800
   Modify: 2021-11-30 21:23:18.281364762 +0800
   Change: 2021-11-30 21:23:18.281364762 +0800
    Birth: -
   cpl in /tmp λ du -h test
   0       test
   ```

   

