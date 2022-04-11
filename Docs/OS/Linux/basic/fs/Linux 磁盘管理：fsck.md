# fsck

> file system check对于需要修复的磁盘必须`umount`后再执行

用于检查和修复linux filesys

syntax：`fsck [-t fstype] [filesystem]`

filesystem可以是磁盘，也可以是挂载点，或则是分区也可以是UUID表示的分区

如果没有指定filesystem，默认顺序检查`/etc/fstab`中的



## 参数

- `-t <fstype>`

  指定检查分区的filesys类型，如果没有指定默认使用ext2

  ```
  [root@chz etc]# fsck -t ext4 /dev/sdb1
  fsck from util-linux 2.23.2
  e2fsck 1.42.9 (28-Dec-2013)
  /dev/sdb1: clean, 11/65536 files, 8859/261883 blocks
  ```

- `-N`

  不修复filesys，但是打印出错误

  ```
  
  ```

- `-a`

  自动修复filesys

- `-r`

  交互式的修复filesys



























