# Linux stat

> ll使用就是stat中的信息

## Digest

`stat` 用于查看文件的 inode 信息

## Output

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

## Exmaples

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

