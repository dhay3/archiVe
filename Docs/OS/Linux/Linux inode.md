# Linux inode

ref:

https://en.wikipedia.org/wiki/Inode

https://man7.org/linux/man-pages/man7/inode.7.html

https://www.ruanyifeng.com/blog/2011/12/inode.html

## Digest

A file system relies on data structures about the files( which called metadata that describes data ), as opposed to the contents of that file. Each file is associated with an inode, which is identified by an integer, often referred to as an i-number or inode number

Inodes store information about files adn directories(floders), such as file ownership, accessmode( read, write, execute permissions ), time(mtime,ctime,atime), etc

简而言之 inode 就是一个存储文件源信息的一个数据结构

## Why inode

大家都知道文件是存储在硬盘上的，硬盘能操作最小单位为 Sector, 中文也叫扇区。而 file system 能操作最小单位为 Block，中文也叫块。Block 是由 一个或者多个 Sector 组成的

文件都被存储在 Block 中，很显然我们需要通过一张表来找到 Block。这张表就叫做 inode table

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/f/f8/File_table_and_inode_table.svg/220px-File_table_and_inode_table.svg.png)

这里提一嘴为了适应普遍 OS file system 的 4KB block size，现在大多数硬盘厂商的 Sector 出厂会为设置 4KB，所以才有了 4K 对这一说法。具体可以参考[4k对齐.md]()

## What are in inode 

具体查看 `man inode`

文件 inode 包含的信息可以使用 `stat` 来查看

```
➜  /etc stat whois.conf 
  File: whois.conf
  Size: 380             Blocks: 8          IO Block: 4096   regular file
Device: 259,7   Inode: 12849129    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2022-06-13 12:03:09.000000000 +0800
Modify: 2022-04-08 21:52:24.000000000 +0800
Change: 2022-06-13 12:03:09.561954812 +0800
 Birth: 2022-06-13 12:03:09.558621456 +0800
```

包含如下几个字段

- File type

  file type 由 `stat.st_mode & S_IFMT` 确认。

  其中 `S_IFMT` 是源码中值为`0170000`的一个 const，作为 `stat.st_mode` 的 mask(与运算掩码)。逻辑代码如下

  ```
  stat(pathname, &sb);
  if ((sb.st_mode & S_IFMT) == S_IFREG) {
  /* Handle regular file */
  }
  ```

  与运算的结果就是 file type 可以是如下几个值

  1. S_IFSOCK(0140000): socket
  2. S_IFLINK(0120000): symbolic link
  3. S_IFREG(0100000): regular file
  4. S_IFBLK(0060000): block device
  5. S_IFDIR(0040000): direcotory
  6. S_IFCHR(0020000): character device
  7. S_IFIFO(0010000): FIFO(具名管道符)

  例如例子就表示当前文件是一个 REG file

- Access mode

  也被称为 File mode，对应源码中的`st_mode`成员变量

  > 3 bit 为一组，所以转成 decimal 最大值为 7

  the 12 bits corresponding to the mask 07777 as the file mode bits

  the least significant 9 bits (0777) as the file permission bits

  第一组 3 bit 和 UGID 有关，可以是如下的值：

  1. S_ISUID(04000): set user id bit
  2. S_ISGID(02000): set group id bit
  3. S_ISVTX(01000): sticky bit, on a directory means that a file in that directory can be renamed or deleted only by the owner of the file, by the onwer of the directory, and by a priviledged process

  第一组没有 07000 对应的值

  第二组 3 bit 和 owner permission 有关，可以是如下的值：

  1. S_IRWXU(00700): owner has read, write, and execute permission
  2. S_IRUSR(00400): owner has read permission
  3. S_IWUSR(00200): onwer has write permission
  4. S_IXUSR(00100): onwer has excute permission

  第 3 组 bit 和 group permission 有关，可以是如下的值：

  1. S_IRWXG(00070): group has read, write, and execute permission 
  2. S_IRGRP(00040): group has read permission
  3. S_IWGRP(00020): group has write permission
  4. S_IXGRP(00010): group has execute permission

  第 4 组 bit 和 other permission 有关，而可以是如下的值

  1. S_IRWXO(00007): others(not in the group) have read, write, and execute permission
  2. S_IROTH(00004): others have read permission
  3. S_IWOTH(00002): others have write permission
  4. S_IXOTH(00001): others have execute permission

  例如

  `Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)`

  转成 decimal 值就为 06644

- Size

  if it is a regular file the field shows the size of the file in bytes

  ==if it is a symbolic link the field shows the length of the pathname it contains==

  例子中为 380 bytes

- Blocks

  number of blocks allocated to the file

  例子中为 8 , 该文件由 8 blocks 组成，即 4 * 8 = 32 bytes

- Device

  each inode(as well as the associated file) resides in a filesystem that is hosted on a device. That device is identified by the combination of its major ID (which identifies the general class of device) and minor ID (which identifies a specific instance in the gernal class ==这里的 instance 表示的是 device 上划分的 partition==)

  例如例子中的 259,7 表示在 259 号设备上的第 7 个分区。可以通过 `lsblk` 来查看具体在那个设备那个分区上

- Inode number

  用于标识文件，==Unix-like OS 是通过 inode 来区分文件的并不是文件名(学过Database的都应该知道为什么)==

  这里需要注意 link 文件的 inode。具体参考下面的 Inode in links

  例子中是 12849129

- Links

  the number of ==hard links== to the file

  需要注意的一点是，原始文件也是一个 hard link  所以会算 1

- Uid

  the user ID of the owner of the file

- Gid

  the group ID of the owner of the file

- Last access timestamp ( atime )

  the last access timestamp, it is changed by file accesses

  当系统调用 execve, mknod, pipe, utime, read 等 readIO 的系统函数时会修改

- Last modification timestamp( mtime )

  the last modification timestamp, it is changed by file modifications

  当系统调用 mknod, truncate, utime, wrete 等 writeIO 的系统函数是会修改

- Last status change timestamp（ctime）

  the file’s last status change timestamp, it is changed by writing or by setting inode information(eg. owner, group, link count, mode, etc..)

  和 mtime 很像，但是 inode 信息修改也会影响该值

- File creation(birth) timestamp (btime)

  the file’s creation timestamp

  顾名思义

## How inode works

The inode number indexes a table of inodes in a known location on the device. From the inode number, the kernel’s file system driver can access the inode contents, including the location of the file, thereby allowing access to the file

内核通过 inode number 查 inode table 找到对应的 inode，通过 inode 找到对应的文件

## Limitation

## Inode in links

## Cautions

for pseudofiles that are autogenerated by the kernel, the file size reported by the kernel is not accurate