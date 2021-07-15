# Linux 磁盘管理：du

`du`命令用于查看文件的大小(默认以block size为最小显示单位)

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

## du vs stat

stat默认以实际存储的数据显示，du以block size为单位显示。如果实际大小小于block size时以block size取整显示(因为filesys只能操作block)

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



