# Linux stat

> ll使用就是stat中的信息

用于查看文件的具体信息，如果单独查看文件修改日期可以使用`ll -t`

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

- Size：文件字节数(bytes)，比`du`精确(du默认以kb为单位)，但是没有精度转换
- IO Block：数据块大小
- Links：链接数，即有多少文件名指向这个inode
- Access：文件的读写权限，创建文件的用户ID和组ID
- Access-time：文件上一次打开的时间

> modify-time和change-time区别
>
> https://askubuntu.com/questions/600837/what-is-the-difference-between-modify-and-change-in-stat-output

- Modify-time：==文件内容上一次变动的时间==
- Change-time：==metadata上一次变动的时间==

可以通过`-c | --format`指定输出内容

```
#以字节输出大小
cpl in ~/note/docs/Linux/基础/disk on master ● λ stat -c %s Linux\ 磁盘管理：blkid.md 
1754
```

使用`-f`参数输出fs相关的信息，可以获得block size(在windows上被称为簇cluster)

```
cpl in /sys/block/nvme0n1/queue λ stat -f /etc/resolv.conf
  File: "/etc/resolv.conf"
    ID: bd5efaae75a7f210 Namelen: 255     Type: ext2/ext3
Block size: 4096       Fundamental block size: 4096
Blocks: Total: 64230001   Free: 52398172   Available: 49117276
Inodes: Total: 16384000   Free: 15915349
```

