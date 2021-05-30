# Linux touch

参考：

https://blog.csdn.net/qq_29503203/article/details/53862790

touch除了创建文件之外还可以修改文件的atime，ctime，mtime等。

- 修改atime到当前主机的时间

  ```
  root in /usr/local λ stat cus_man.sh
    File: cus_man.sh
    Size: 254             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407745      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-08 11:41:40.078382050 +0800
  Modify: 2021-04-08 11:41:38.718395171 +0800
  Change: 2021-04-08 11:41:38.718395171 +0800
   Birth: -
  root in /usr/local λ date
  Tue Apr 13 10:30:30 CST 2021
  root in /usr/local λ touch -a cus_man.sh
  root in /usr/local λ stat cus_man.sh
    File: cus_man.sh
    Size: 254             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407745      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:30:41.022036228 +0800
  Modify: 2021-04-08 11:41:38.718395171 +0800
  Change: 2021-04-13 10:30:41.022036228 +0800
   Birth: -
  ```

- 修改mtime，因为mtime，atime，ctime都存储在inode中，所以修改mtime的同时也会修改ctime

  ```
  root in /usr/local λ touch -m cus_man.sh;stat cus_man.sh
    File: cus_man.sh
    Size: 254             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407745      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:32:16.285146222 +0800
  Modify: 2021-04-13 10:32:31.669002490 +0800
  Change: 2021-04-13 10:32:31.669002490 +0800
   Birth: -
  ```

- ==修改atime和mtime到指定时间==

  ```
  root in /usr/local λ touch -d "2020-01-15 10:30:45" cus_man.sh
  root in /usr/local λ stat cus_man.sh
    File: cus_man.sh
    Size: 254             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407745      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2020-01-15 10:30:45.000000000 +0800
  Modify: 2020-01-15 10:30:45.000000000 +0800
  Change: 2021-04-13 11:44:01.584611571 +0800
   Birth: -
  ```

- 使用指定文件的时间戳修改，并求改ctime到当前时间

  ```
  root in /usr/local λ stat bin
    File: bin
    Size: 4096            Blocks: 8          IO Block: 4096   directory
  Device: fc01h/64513d    Inode: 396189      Links: 2
  Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 09:57:56.820566390 +0800
  Modify: 2020-03-11 16:52:58.561788635 +0800
  Change: 2020-03-11 16:52:58.561788635 +0800
   Birth: -
   
   root in /usr/local λ stat cus_man.sh
    File: cus_man.sh
    Size: 254             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407745      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:33:27.280482886 +0800
  Modify: 2021-04-13 10:33:26.824487146 +0800
  Change: 2021-04-13 10:33:26.824487146 +0800
   Birth: -
  
  root in /usr/local λ touch -r cus_man.sh bin
  
  root in /usr/local λ stat bin
    File: bin
    Size: 4096            Blocks: 8          IO Block: 4096   directory
  Device: fc01h/64513d    Inode: 396189      Links: 2
  Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:33:27.280482886 +0800
  Modify: 2021-04-13 10:33:26.824487146 +0800
  Change: 2021-04-13 10:36:59.422500498 +0800
   Birth: -
  ```

  ## ctime
  
  参考：
  
  https://unix.stackexchange.com/questions/36021/how-can-i-change-change-date-of-file
  
  linux默认不能手动修改ctime，只有文件的metadata修改后自动更新ctime
  
  ```
  # Update ctime
  debugfs -w -R 'set_inode_field /tmp/foo ctime 201001010101' /dev/sda1
  
  # Drop vm cache so ctime update is reflected
  echo 2 > /proc/sys/vm/drop_caches
  ```
  
  



