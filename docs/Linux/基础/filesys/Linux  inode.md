# Linux  inode

参考：https://www.ruanyifeng.com/blog/2011/12/inode.html

Linux文件数据都存储在"block"中，显然，我们还需要找到一个地方存储文件的元信息，比如文件的创建者、文件的创建日期、文件的大小等等。这种存储文件元信息的区域就叫做inode，每一个文件都有对应的inode

==Unix/Linux系统内部不使用文件名，而使用inode号码来标识文件==

系统内部打开文件实际分成三步

1. 系统首先找到文件名对应的inode号码
2. 其次通过inode号码，获取inode信息
3. 根据inode信息，找到文件数据所在block，读取数据

我们可以通过`stat`命令来查看文件的inode，可以通过`ls -i <filename>`来查看文件inode的值

> 如果我们要查看目录的inode可以通过`ls -di <dire>`

```
root in /opt λ stat bak.xml 
  File: bak.xml
  Size: 1317387   	Blocks: 2576       IO Block: 4096   regular file
Device: 801h/2049d	Inode: 24          Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2020-12-18 20:26:21.793017524 -0500
Modify: 2020-09-09 00:12:37.989966656 -0400
Change: 2020-09-09 00:12:37.989966656 -0400
 Birth: -                    
```

- Size：文件字节数
- IO Block：数据块大小
- Links：链接数，即有多少文件名指向这个inode
- Access：文件的读写权限，创建文件的用户ID和组ID
- Access-time：文件上一次打开的时间
- Modify-time：==文件内容上一次变动的时间==
- Change-time：==inode上一次变动的时间==

我们可以使用`df -i`来查看各磁盘上的inode情况

```
root in /opt λ df -i
Filesystem      Inodes  IUsed  IFree IUse% Mounted on
udev            493107    358 492749    1% /dev
tmpfs           500741    660 500081    1% /run
/dev/sda1      1248480 333548 914932   27% /
tmpfs           500741      1 500740    1% /dev/shm
tmpfs           500741      3 500738    1% /run/lock
tmpfs           500741     17 500724    1% /sys/fs/cgroup
tmpfs           500741     46 500695    1% /run/user/0        
```

使用`dumpe2fs`来查看inode的大小

```
root in /opt λdumpe2fs -h /dev/sda1 | grep "Inode size"
dumpe2fs 1.45.6 (20-Mar-2020)
Inode size:	          256    
```

> 每个文件都必须有一个inode，如果inode已经用完了，就无法在硬盘上创建文件
