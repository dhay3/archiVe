# rsync

https://linux.die.net/man/1/rsync

https://einverne.github.io/post/2017/07/rsync-introduction.html

https://www.cnblogs.com/f-ck-need-u/p/7220009.html

https://www.huaweicloud.com/articles/51b251bb13e4ad517a86c4533d277636.html

> 常用参数`-avvz`

## 概述

用于本机计算机与远程计算机之间，或者两个本地目录之间同步文件。与其他文件传输工具（FTP或scp）不同，==rsync 的最大特点是会检查发送方和接收方已有的文件，仅传输有变动的部分==（增量传输delta-transfer algorithm，默认规则是文件大小或修改时间有变动）。

增量传输，因为使用`-a`参数还会将时间戳传输

```
root in /opt/test1 λ rsync  -av /opt/test1/ /opt/test2/
sending incremental file list
./
1
10
2
3
4
5
6
7
8
9

sent 472 bytes  received 209 bytes  1,362.00 bytes/sec
total size is 0  speedup is 0.00
root in /opt/test1 λ rsync  -av /opt/test1/ /opt/test2/
sending incremental file list

sent 155 bytes  received 12 bytes  334.00 bytes/sec
total size is 0  speedup is 0.00
```

syntax：

> 在rsync中，目录与文件有严格的区别，如果想要表示目录下的所有文件，需要在最后添加slash，否则会将文件夹发送。

- local 

  `rsync [options] <src_path ...> [dest_path]`

  ```
  root in /opt/test1 λ ls && cd ../test2 && echo -e "\n" && ls
   1   10   2   3   4   5   6   7   8   9
  
  
   10   11   12   13   14   15   16   17   18   19   20
  
  root in /opt/test2 λ rsync -av /opt/test1/ /opt/test2/
  sending incremental file list
  ./
  1
  10
  2
  3
  4
  5
  6
  7
  8
  9
  
  root in /opt/test2 λls && cd ../test1 && echo -e "\n" && ls
   1    11   13   15   17   19   20   4   6   8
   10   12   14   16   18   2    3    5   7   9
  
  
   1   10   2   3   4   5   6   7   8   9
  ```

- remote shell

  pull，远程同步到本地

  `rsync [options] [user@]<host:src_path ...>[dest_path]`

  push，本地同步到远程

  `rsync [options] <src> [user@]<host:dest>`

- rsync daemon

  ==目的地址需要安装rsync daemon才能使用==

  这里`::`和`rsync`表示使用rsync daemon传输

  pull

  `rsync [options] <src...[user@]host::dest>`

  push
  
  `rsync [options] <src... rsync://[user@]host:[port]/dest>`

==如果没有源地址，等价于`ls`==

```
root in /opt/test/t2 λ rsync cpl@8.135.0.171:/opt/
drwxr-xr-x          4,096 2021/02/25 10:30:50 .
-rwxr-xr-x              0 2021/02/22 16:51:23 .dockerenv
-rw-r--r--            126 2021/01/22 14:00:51 Dockerfile
-rw-r--r--            615 2021/02/22 16:51:23 resolv.conf
-rwxr-xr-x            107 2021/02/05 16:05:06 tput_t.sh
drwxr-xr-x          4,096 2021/01/27 17:39:00 alibabacloud
drwx--x--x          4,096 2021/01/21 10:58:32 containerd
drwxr-xr-x          4,096 2020/12/31 14:17:18 lsd-0.18.0-x86_64-unknown-linux-gnu
drwxr-xr-x          4,096 2021/02/25 10:32:16 t
```

## 参数

- --progress

  显示进度条

- `-e`

  选择一个连接远程终端的进程，当ssh不是默认端口时非常有用

  ```
  rsync -avz -e "ssh -p $port" /local/path/ user@remoteip:/path/to/files/
  ```

- `-z,--compress`

  传输数据时使用压缩数据，不是在目的地址生成压缩文件

- `-v | -vv`

  输出详细信息

- `-q, --quiet`

  不会讲传输的信息输出，如果是以cron启动这个参数非常有用

- `--bwllimit`

  最大的IOPS,可以使用单位，如果没有指定单位默认1024bytes

  ```
   rsync -avz   -e 'ssh -p 41456' --exclude='html/bridd.pub/application/config' --bwlimit=1.5m
  ```

- `-a, --archive`

  将源地址中的内容同步到目的地址，==同时会将隐藏文件同步==。等价于`-rlptgoD`参数，同步文件的同时会同步文件的metadata（修改时间，权限等）。

  ```
  root in /opt λ rsync -av hydra.restore cpl@8.135.0.171:/home/cpl
  sending incremental file list
  hydra.restore
  
  sent 5,555 bytes  received 35 bytes  3,726.67 bytes/sec
  total size is 5,457  speedup is 0.98
  ```

  1. `-r, --recursive`

     告诉rsync递归拷贝

  2. `-l,--link`

     告诉rsync源地址如果遇到链接文件将链接文件拷贝到目的地址

  3. `-p,--perms`

     告诉rsync将源地址的文件权限同样复制到目的地址

  4. `-g,--group`

     告诉rsync将源地址组信息复制到目的地址

  5. `-o,--omit-dir-time`

     告诉rsync复制时忽略修改文件的时间

  6. `-D,--device`

     告诉rsync可以将字节和块文件复制到目的地址

- `-n,--dry-run`

  模拟执行后的结果，不会生效

  ```
  root in /opt λ rsync -anv hydra.restore cpl@8.135.0.171:/opt 
  sending incremental file list
  hydra.restore
  
  sent 54 bytes  received 19 bytes  48.67 bytes/sec
  total size is 5,457  speedup is 74.75 (DRY RUN)
  ```

- `-R,--relative`

  使用相对路径生成文件

  ```
  rsync -av /foo/bar/baz.c remote:/tmp/
  ```

  会在remote的`/tmp`下生成一个baz.c文件

  ```
  rsync -avR /foo/bar/baz.c remote:/tmp/
  ```

  会在remote的`/tmp`下一个`/tmp/foo/bar/baz.c`的文件包括目录

- `-P`

  打印文件传输进度，与`-v`参数冲突
  
  ```
  root in /opt/test λrsync -aP /opt/test/test2 /opt/test/test3
  sending incremental file list
  test2/
  test2/1
                0 100%    0.00kB/s    0:00:00 (xfr#1, to-chk=19/21)
  test2/10
                0 100%    0.00kB/s    0:00:00 (xfr#2, to-chk=18/21)
  test2/11
                0 100%    0.00kB/s    0:00:00 (xfr#3, to-chk=17/21)
  test2/12
                0 100%    0.00kB/s    0:00:00 (xfr#4, to-chk=16/21)
  test2/13
                0 100%    0.00kB/s    0:00:00 (xfr#5, to-chk=15/21)
  test2/14
                0 100%    0.00kB/s    0:00:00 (xfr#6, to-chk=14/21)
  ```

- `--existing`

  只同步目的地址中有的文件

  ```
  root in /opt/test/test3 λ rsync -av --existing /opt/test/test1/ /opt/test/test3/
  sending incremental file list
  
  sent 41 bytes  received 12 bytes  106.00 bytes/sec
  total size is 0  speedup is 0.00
  ```

- `--update,-u`

  ==如果目标地址上的文件时间戳晚于源地址上的，rsync不对它做同步==。只对目标地址上时间戳早于源地址上的文件做同步
  
  ```
  root in /opt/test/t2 λ stat 3
    File: 3
    Size: 4         	Blocks: 8          IO Block: 4096   regular file
  Device: 801h/2049d	Inode: 10653       Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-02-25 14:29:14.573356713 +0800
  Modify: 2021-02-25 14:29:12.357356808 +0800
  Change: 2021-02-25 14:29:12.357356808 +0800
   Birth: -
  root in /opt/test/t2 λ cd ../t1 && stat 3
    File: 3
    Size: 0         	Blocks: 0          IO Block: 4096   regular empty file
  Device: 801h/2049d	Inode: 10647       Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-02-25 13:59:35.217433003 +0800
  Modify: 2021-02-25 13:59:35.217433003 +0800
  Change: 2021-02-25 13:59:35.217433003 +0800
   Birth: -
  
  root in /opt/test/t2 λ rsync -avvu /opt/test/t1/ /opt/test/t2/
  sending incremental file list
  delta-transmission disabled for local transfer or --whole-file
  1 is newer
  2 is newer
  3 is newer
  total: matches=0  hash_hits=0  false_alarms=0 data=0
  
  sent 84 bytes  received 131 bytes  430.00 bytes/sec
  total size is 0  speedup is 0.00
  ```
  
- `--delete`

  同步src_path的内容到dest_path，如果dest_path中内容比src_path多就会删除这些文件。这个参数非常危险，==在使用前最好使用dry-run==
  
  ```
  root in /opt/test λ rsync --delete -av /opt/test/test2/ /opt/test/test3/
  sending incremental file list
  deleting test2/9
  deleting test2/8
  deleting test2/7
  deleting test2/6
  deleting test2/5
  deleting test2/4
  deleting test2/3
  deleting test2/20
  deleting test2/2
  deleting test2/19
  deleting test2/18
  deleting test2/17
  deleting test2/16
  deleting test2/15
  deleting test2/14
  deleting test2/13
  deleting test2/12
  deleting test2/11
  deleting test2/10
  deleting test2/1
  deleting test2/
  ./
  1
  10
  11
  12
  ....
  ```

- `--exclude`

  排除文件。rsync默认会将隐藏文件同步，可以使用该参数排除

  ```
  root in /opt/test1 λ rsync --exclude=".*" -av ~/ /opt/test1
  root in /opt/test1 λ ls -a                                 
   \u{1b}\u{1b}\u{1b}q   Desktop     Music       Public    Templates
   .                     Documents   nohup.out   revokey   Videos
   ..                    Downloads   Pictures    src      
  ```

  可以使用curly bracket 的模式扩展

  ```
  rsync -av --exclude={'file1.txt','dir1/*'} source/ destination
  ```

- `--include`

  只同步指定文件

  ```
  rsync -av --include="*.txt" --exclude='*' source/ destination
  ```

- `-W,--whole-file`

  告诉rsync 不使用delta-transfer algorithm增量传输，而是使用全量传输。==如果使用local syntax 默认使用该选项==。

- `--link-dest`

  选定一个目录做为同步的基准。test1做为源地址，test2做为同步的基准，test3做为目的地址

  ```
  root in /opt/test λ ls
   test1   test2   test3
  root in /opt/test λ ls test2
   1    11   13   15   17   19   20   4   6   8
   10   12   14   16   18   2    3    5   7   9
  ```

- `--backup,-b`

  备份，被传输覆盖或是使用`--delete`被删除的文件会以`~`为后缀备份，可以通过`--backup-dir`来指定备份存储的位置，`--suffix`指定后缀

  ```
  #被删除的文件
  root in /opt/test/t2 λ rsync  -b  -avvz --delete /opt/test/t1/ /opt/test/t2/
  sending incremental file list
  delta-transmission disabled for local transfer or --whole-file
  backed up 4 to 4~
  deleting 4
  1 is uptodate
  2 is uptodate
  3 is uptodate
  total: matches=0  hash_hits=0  false_alarms=0 data=0
  
  sent 89 bytes  received 172 bytes  522.00 bytes/sec
  total size is 0  speedup is 0.00
  root in /opt/test/t2 λ ls
   1   2   3   4~
  
  ---
  #被覆盖的文件
  root in /opt/test/t2 λ rsync  -b  -avvz /opt/test/t1/ /opt/test/t2/
  sending incremental file list
  delta-transmission disabled for local transfer or --whole-file
  2 is uptodate
  3 is uptodate
  1
  backed up 1 to 1~
  total: matches=0  hash_hits=0  false_alarms=0 data=0
  
  sent 118 bytes  received 169 bytes  574.00 bytes/sec
  total size is 0  speedup is 0.00
  root in /opt/test/t2 λ ls
   1   1~   2   3   4~
  root in /opt/test/t2 λ cat 1
    File: 1   <EMPTY>
  root in /opt/test/t2 λ cat 1~
    File: 1~
    1
  ```

- `--prune-empty-dirs,-m`

  该选项告诉receiver端的rsync从文件列表中删除空目录，包括那些没有文件的空的嵌套目录

## filter rules

rsync 可以通过`--filter,-f`参数来过滤同步的文件，后面跟pattern

- `-`：exclude
- `+`：include
- `.`：merge
- `:`：dir-merge

- `H`：hide，指定传输隐藏文件的格式

- `P`：protect，指定不能被删除的文件的格式

- `R`：risk，指定可以被删除的文件的格式

  
