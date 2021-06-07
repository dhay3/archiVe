# atime，mtime，ctime

参考：

https://www.howtogeek.com/517098/linux-file-timestamps-explained-atime-mtime-and-ctime/

> 可以通过`touch`命令来修改atime，mtime，ctime

```
root in /usr/local λ stat cus_alias.sh
  File: cus_alias.sh
  Size: 231             Blocks: 8          IO Block: 4096   regular file
Device: fc01h/64513d    Inode: 407741      Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2021-04-08 11:41:09.078681113 +0800
Modify: 2021-04-08 11:41:07.906692420 +0800
Change: 2021-04-08 11:41:07.906692420 +0800
 Birth: -
```

- atime：access-time ，上次读取(cat等命令)文件的时间

  ```
  root in /usr/local λ cat cus_alias.sh >| /dev/null;stat cus_alias.sh
    File: cus_alias.sh
    Size: 231             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407741      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:20:24.583847003 +0800
  Modify: 2021-04-08 11:41:07.906692420 +0800
  Change: 2021-04-08 11:41:07.906692420 +0800
   Birth: -
  ```

- mtime：modify-time，上次文件内容改变的时间

  ```
  root in /usr/local λ echo "#test" >> cus_alias.sh;stat cus_alias.sh
    File: cus_alias.sh
    Size: 237             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407741      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:20:24.583847003 +0800
  Modify: 2021-04-13 10:23:18.910202946 +0800
  Change: 2021-04-13 10:23:18.910202946 +0800
   Birth: -
  ```

- ctime：change-time，上次metadata(权限)改变的时间

  ```
  root in /usr/local λ chmod +x cus_alias.sh
  root in /usr/local λ stat cus_alias.sh
    File: cus_alias.sh
    Size: 237             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407741      Links: 1
  Access: (0755/-rwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:23:19.282199440 +0800
  Modify: 2021-04-13 10:23:18.910202946 +0800
  Change: 2021-04-13 10:27:59.135560640 +0800
  Birth: -
  ```

## 特殊

- 参考：

  https://unix.stackexchange.com/questions/36021/how-can-i-change-change-date-of-file

  linux默认不能手动修改ctime，只有文件的metadata修改后自动更新ctime

  ```
  # Update ctime
  debugfs -w -R 'set_inode_field /tmp/foo ctime 201001010101' /dev/sda1
  
  # Drop vm cache so ctime update is reflected
  echo 2 > /proc/sys/vm/drop_caches
  ```

