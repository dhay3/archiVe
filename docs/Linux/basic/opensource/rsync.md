# rsync

[https://linux.die.net/man/1/rsync](https://linux.die.net/man/1/rsync)

[https://einverne.github.io/post/2017/07/rsync-introduction.html](https://einverne.github.io/post/2017/07/rsync-introduction.html)

[https://www.cnblogs.com/f-ck-need-u/p/7220009.html](https://www.cnblogs.com/f-ck-need-u/p/7220009.html)

[https://www.huaweicloud.com/articles/51b251bb13e4ad517a86c4533d277636.html](https://www.huaweicloud.com/articles/51b251bb13e4ad517a86c4533d277636.html)

https://www.ruanyifeng.com/blog/2020/08/rsync.html

> 常用参数`-ahvvz`

## Digest
用于本机计算机与远程计算机之间(也可以和rsync daemon 进行PIC proccess Inter communication)，或者两个本地目录之间同步文件。与其他文件传输工具（FTP或scp）不同，rsync 的最大特点是会检查发送方和接收方已有的文件，仅传输有变动的部分（增量传输delta-transfer algorithm，默认规则是文件大小或修改时间有变动）。
rsync 默认使用ssh通道(除rsync daemon外)，所以使用rsync的前提是ssh必须是通的

```protobuf
[root@centos7 a]# rsync -av -e 'ssh -p 65522' ubuntu@82.157.1.139:/etc
The authenticity of host '[82.157.1.139]:65522 ([82.157.1.137]:65522)' can't be established.
ECDSA key fingerprint is SHA256:R2YLjnj2RLeXTkeRpF+DTeR1n/eR3k2m1L8hYKFQX9M.
ECDSA key fingerprint is MD5:5d:a0:01:73:e8:ac:ed:3b:9d:e5:2f:b9:e3:22:7b:63.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '[82.157.1.139]:65522' (ECDSA) to the list of known hosts.
ubuntu@82.157.1.139's password: 

```

rsync有如下几种特性

1. 支持拷贝links，devices，owners，groups，and permissions
1. 和 GNU tar 一样的 exclude 和 exclude-from 选项功能
1. CVS(Concurrent Version System) exclude mode 会忽略未做变更的文件
1. 可以使用任何remote shell(ssh，rsh)
1. 不需要使用super-user权限
1. 流水式的传输文件以减小latency
1. 针对 rsync daemons 支持 anonymous 或 authenticated
## Syntax
如果只指定源，rsync只会对源目录做遍历，即等价`ls`
```protobuf
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
### local
本地rsync无需指定 host 或 rsync daemon
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
### remote shell
```protobuf
Access via remote shell:
	Pull: rsync [OPTION...] [USER@]HOST:SRC... [DEST]
	Push: rsync [OPTION...] SRC... [USER@]HOST:DEST
```
```protobuf
rsync -av host:file1 :file2 host:file{3,4} /dest/
```
### rsync daemon
remote rsync daemon 默认使用 TCP 873(如果协议不同端口是可以被复用的)，为了连上rsync daemon，目的端必须运行rsync daemon 进程
如果连接rsync daemon 绝对不能使用`--rsh`参数
```protobuf
Access via rsync daemon:
	Pull: rsync [OPTION...] [USER@]HOST::SRC... [DEST]
      	rsync [OPTION...] rsync://[USER@]HOST[:PORT]/SRC... [DEST]
	Push: rsync [OPTION...] SRC... [USER@]HOST::DEST
        rsync [OPTION...] SRC... rsync://[USER@]HOST[:PORT]/DEST
```
除了使用double colon 外，还可以使用`rsync://`替代
```protobuf
rsync -av host::modname/file{1,2} host::modname/file3 /dest/
rsync -av host::modname/file1 ::modname/file{3,4}
```
当连上rsync deamons 是打印出时间戳
如果rsync daemon 需要密码，可以使用`--passwrod-file`参数或`RSYNC_PASSWORD`来指定password，通常用于scripting rsync
## cautions

- rsync可以和shell的globbing一起使用

- 在rsync中，目录与文件有严格的区别，如果想要表示目录下的所有文件，需要在最后添加slash，否则会将文件夹发送。

_A  trailing slash on the source changes this behavior to avoid creating an additional 		directory level at the destination. You can think of a trailing / on a source as meaning 	"copy the contents of this directory" as opposed to "copy the directory  by  name"_
```protobuf
rsync -av /src/foo /dest
rsync -av /src/foo/ /dest/foo
```
如上得到的效果一样
```protobuf
[root@centos7 b]# ls && echo '---' && ls ../a
---
1  2

[root@centos7 tmp]# rsync -av a/ b
sending incremental file list
./
1
2

sent 173 bytes  received 57 bytes  460.00 bytes/sec
total size is 4  speedup is 0.02
[root@centos7 tmp]# ls b
1  2

[root@centos7 tmp]# rsync -av a b
sending incremental file list
a/
a/1
a/2

sent 184 bytes  received 58 bytes  484.00 bytes/sec
total size is 4  speedup is 0.02
[root@centos7 tmp]# ls -R b
b:
a

b/a:
1  2
```

如果文件名包含空格，必须转译或使用`--protect-args`参数

```protobuf
rsync -av host:'file\ name\ with\ spaces' /dest
```
## Positonal Argus

- SRC

  源文件

## Optional Argus
### special

- `-n,--dry-run`
  模拟执行后的结果，不会生效  

  ```
  root in /opt λ rsync -anv hydra.restore cpl@8.135.0.171:/opt 
  sending incremental file list
  hydra.restore
  
  sent 54 bytes  received 19 bytes  48.67 bytes/sec
  total size is 5,457  speedup is 74.75 (DRY RUN)
  ```

- `-r | --recursive`

  rsync 对文件递归拷贝，需要源目rsync的版本都是3.0.0以上的

- `-d | --dirs`

  对文件递归，和`-r`不一样的是只会对 the directory name specified is `.` or ends with a trailing slash的文件进行拷贝。如果和`-r`一起使用`-r`优先生效

- `-R,--relative`
  使用相对路径生成文件 
  会在remote的`/tmp`下生成一个baz.c文件 
  会在remote的`/tmp`下一个`/tmp/foo/bar/baz.c`的文件包括目录 

  ```
  rsync -av /foo/bar/baz.c remote:/tmp/
  rsync -avR /foo/bar/baz.c remote:/tmp/
  ```

- `--backup,-b`
  备份，被传输覆盖(preexisting destination files)或是使用`--delete`被删除的文件会以`~`为后缀备份，可以通过`--backup-dir`来指定备份存储的位置，`--suffix`指定后缀  

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

- `--super`

  receiving rsync 会尝试使用super-user

- `-W,--whole-file`
  告诉rsync 不使用delta-transfer algorithm增量传输，而是使用全量传输。如果使用local syntax 默认使用该选项。 

- `--force`

  和`--delete`一起使用，强制删除非空文件

- `-e | --rsh=COMMAND`

  告诉rsync使用的rmote shell，默认使用ssh

  ```
  rsync -avz -e "ssh -p $port" /local/path/ user@remoteip:/path/to/files/
  
  -e 'ssh -p 2234'
  -e 'ssh -o "ProxyCommand nohup ssh firewall nc -w1 %h %p"'
  ```

- `--rsync-path=PROGRAM`

  如果rsync在reciver side 不是使用默认path（`/usr/local/bin/rsync`），需要使用该参数指定。可使用如下方式修改在reciver side 调用rsync的目录

  ```
  rsync -avR --rsync-path="cd /a/b && rsync" host:c/d /e/
  ```

  也意味着PROGRAM的值可以是任意的

- `-z,--compress`
  传输数据时使用压缩数据，不是在目的地址生成压缩文件

- `-P | --progress`
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

  

- `--chown=USER:GROUP`

  指定transfer 的文件所有者

- `--address=ADDRESS`

  和rsync daemon 交互的sender side 绑定IP，默认使用wildcard即本机所有路由可达IP

- `--port=PORT`

  指定rsync daemon 使用的port，默认使用873

- `--password-file=FILE`

  指定rsync daemon 使用的密码，但是同样需提供ssh通道的密码

- `--list-only`

  对sender side 文件做遍历

### transfer

- `-I | --ignore-times`

  此选项会修改rsync的quick check algorithm。如果文件大小相同且mtime相同(不比较permissions、owner)，rsync会跳过此文件。

- `--size-only`

  此选项会修改rsync的quick chekc algorithm。只针对文件大小变化的文件传输，在不想保留原文件的时间戳时使用

- `-c | --checksum`

  修改rsync任务文件修改的逻辑，没有此参数默认使用quick check。使用MD5校验src和dst是否相同了，会增大disk I/O的压力

- `-a | --archive`
  将源地址中的内容同步到目的地址，同时会将隐藏文件同步。等价于`-rlptgoD`参数，同步文件的同时会同步文件的metadata（修改时间，权限等）。  

  ```
  root in /opt λ rsync -av hydra.restore cpl@8.135.0.171:/home/cpl
  sending incremental file list
  hydra.restore
  
  sent 5,555 bytes  received 35 bytes  3,726.67 bytes/sec
  total size is 5,457  speedup is 0.98
  ```

  a. `-r, --recursive`
  告诉rsync递归拷贝 

  b. `-l,--link`
  告诉rsync源地址如果遇到链接文件将链接文件拷贝到目的地址（自动解析路径关系） 

  c. `-p,--perms`
  告诉rsync将源地址的文件权限同样复制到目的地址 

  d. `-g,--group`
  告诉rsync将源地址组信息复制到目的地址 

  e. `-o,--omit-dir-time`
  告诉rsync复制时忽略修改文件的时间 

  f. `-D,--device`
  告诉rsync可以将字节和块文件复制到目的地址 

- `--update,-u`
  如果目标地址上的文件时间戳晚于源地址上的，rsync不对它做同步。只对目标地址上时间戳早于源地址上的文件做同步  。不会对 dirs，symlinks 或者 特殊文件(dev、socket file)生效

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

- `--inplace`

  rsync默认当文件需要更新时，是将源文件拷贝到目的，使用该参数rsync instead writes the updated data directly to the desitination file
  有几点需要注意的是

  a. hard links are note broken. 意味着可以在reciever hard link 看到新增修改的内容

  b. In-use binaries 不能被更新

  c. 传输过程中内容保持不一致

  d. 如果运行rsync的用户写权限不够，rsync不会对文件更新

  e. 如果reiever文件被覆写会降低rsync的效率

- `--append`

  将新的内容增加到EOF(end of the file)，源文件和目的文件会被认为完全相同的。如果receiver的文件大小和sender的一样或更大，文件就会被跳过。可以使用`--appned-verify`来校验整个文件的内容，如果文件校验失败会重传

- `-l | --links`

  rsync对symlink默认处理，使用该参数会在reciever处生成symlink

  ```
  [root@centos7 tmp]# ll
  total 0
  lrwxrwxrwx. 1 root root  4 Feb 27 15:01 a -> /etc
  [root@centos7 tmp]# rsync a b
  skipping non-regular file "a"
  [root@centos7 tmp]# ll
  total 0
  lrwxrwxrwx. 1 root root  4 Feb 27 15:01 a -> /etc
  [root@centos7 tmp]# rsync -l a b
  [root@centos7 tmp]# ll
  total 0
  lrwxrwxrwx. 1 root root  4 Feb 27 15:01 a -> /etc
  lrwxrwxrwx. 1 root root  4 Feb 27 15:03 b -> /etc
  [root@centos7 tmp]# 
  ```

- `-L | --copy-links`

  当遇见symlink时，拷贝symlink映射的文件而不是symlink。如果symlink映射的是一个目录，modern rsync 自动跳过

  ```
  [root@centos7 tmp]# rsync -L a b
  skipping directory a
  [root@centos7 tmp]# ll
  total 0
  lrwxrwxrwx. 1 root root  4 Feb 27 15:01 a -> /etc
  
  [root@centos7 tmp]# rsync -L a.link b
  [root@centos7 tmp]# ll
  total 8
  -rw-r--r--. 1 root root  2 Feb 27 15:08 a
  lrwxrwxrwx. 1 root root  1 Feb 27 15:08 a.link -> a
  -rw-r--r--. 1 root root  2 Feb 27 15:09 b
  ```

- `-p | --perms`

  reciever side 的文件的 permissions 和 sender side 的一样。如果没有使用该参数，permission 会使用和`cp`及`tar`一样的规则。和`-A | --acls`等价

- `--chmod`

  指定sender side 文件的permissions(实际permissions不会改变)，也意味着如果没有指定`--perms`reciever side 已经存在的文件的permissions不会改变

- `-o | --owner`

  rsync 默认reciver side 文件的owner 是 the invoking user on the receiving，如果使用该参数onwer会与sender side 文件的一样

- `-g | --group`

  和`-o`一样，表示设置文件的group

- `--devices`

  rsync可以传输block device files，但是 receiver side running user 必须是 root

- `--specails`

  rsync 可以传输sockets fifos files

- `-D`

  等价于`--devices`+ `--specails`

- `-O | --omit-dir-times`

  如果mtime没有修改，不会传输directories

- `-J | --omit-link-times`

  如果mtime没有修改，不会传输symlinks

- `--existing`
  只同步目的地址中有的文件  

  ```
  root in /opt/test/test3 λ rsync -av --existing /opt/test/test1/ /opt/test/test3/
  sending incremental file list
  
  sent 41 bytes  received 12 bytes  106.00 bytes/sec
  total size is 0  speedup is 0.00
  ```

- `--remove-source-file`
  移动文件(意味着不包含目录)而不是拷贝  

  ```
  [root@centos7 tmp]# rsync --remove-source-files a a.d  b
  skipping directory a.d
  [root@centos7 tmp]# ll
  total 4
  drwxr-xr-x. 2 root root  6 Feb 27 16:08 a.d
  -rw-r--r--. 1 root root  2 Feb 27 16:08 b
  ```

- `--delete`
  同步src_path的内容到dest_path，如果dest_path中内容比src_path多就会删除dest_path中的这些文件。这个参数非常危险，在使用前最好使用dry-run 。还可以结合`--delete-before`、`--delete-during`、`delete-delay`、`--delete-after`、`--delete-excluded`，具体查看man page

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

- `--max-size=SIZE`

  告诉rsync只传输小于SIZE的文件或目录，可以使用K、M、G作为单位

  ```
  --max-size=1.5mb-1 is 1499999 bytes, and --max-size=2g+1 is 2147483649 bytes.
  ```

- `--min-size=SIZE`

  逻辑含义同上

- `-C | --cvs-exclude`

  去掉cvs(control version system)文件，默认包含如下格式的文件，以及`$HOME/.cvsignore`文件中记录的 和 CVSIGNORE 环境变量，可以和`--filter`兼容

  ```
  RCS SCCS CVS CVS.adm RCSLOG cvslog.* tags TAGS .make.state .nse_depinfo *~ #* .#* ,*  _$*  *$  *.old  *.bak
  *.BAK *.orig *.rej .del-* *.a *.olb *.o *.obj *.so *.exe *.Z *.elc *.ln core .svn/ .git/ .hg/ .bzr/
  ```

- `-f | --filter=RULE`

  exclude certain files from the list of files to be transferred，规则具体查看filter rules

- `--exclude=PATTERN`
  排除文件。rsync默认会将隐藏文件同步，可以使用该参数排除，规则具体查看filter rules
  可以使用curly bracket 的模式扩展  

  ```
  root in /opt/test1 λ rsync --exclude=".*" -av ~/ /opt/test1
  root in /opt/test1 λ ls -a                                 
   \u{1b}\u{1b}\u{1b}q   Desktop     Music       Public    Templates
   .                     Documents   nohup.out   revokey   Videos
   ..                    Downloads   Pictures    src
  ```

  ```
  rsync -av --exclude={'file1.txt','dir1/*'} source/ destination
  ```

- `--exclude-from=FILE`

  需要和`--exclude`一起使用，表示从具体某个目录下剔除。FILE可是使用`-`表示从stdin中读取

- `--include=PATTERN`
  只同步指定规则的文件，具体规则查看filter rules  

  ```
  rsync -av --include="*.txt" --exclude='*' source/ destination
  ```

- `--include-frome=FILE`

  字面意思，同`--exclude-from`

- `--bwllimit`
  最大的IOPS,可以使用单位，如果没有指定单位默认1024bytes  

  ```
   rsync -avz   -e 'ssh -p 41456' --exclude='html/bridd.pub/application/config' --bwlimit=1.5m
  ```

- `--compare-dest=DIR`

  this option instructs rsync to use DIR on the destination machine as an additional hierarchy to compare destination files aganinst doing transfers, if the files are missing in the destination directory. If a file is found in DIR that is identical to the sender's file, the file wil NOT be transferred to the destination directory
  以reciver side DIR为基准，如果该文件内包含sender side 所有文件就不会通过，如果少只同步少的部分。可以将DIR作为back

- `--copy-dest=DIR`

  behaves like `--compare-dest`，但是如果文件相同还是会做强制更新

- `--link-dest=DIR`

  behaves like `--copy-dest`，但是如果文件相同会生成hard link

  ```
  ➜  test ll a.d
  .rw-r--r-- root root 2 B Wed Mar  9 22:41:58 2022  a
  .rw-r--r-- root root 2 B Wed Mar  9 23:03:30 2022  b
  ➜  test ll b.d
  .rw-r--r-- root root 2 B Wed Mar  9 22:41:58 2022  a
  .rw-r--r-- root root 0 B Wed Mar  9 23:43:24 2022  c
  .rw-r--r-- root root 0 B Wed Mar  9 23:43:24 2022  d
  ➜  test ll c.d
  ```

- `--timeout=SECONDS`

  如果没有数据传输了，等待最大的I/O timeout。默认0表示无限制

- `--contimeout=SECONDS`

  等待连接rsync daemon的最大超时时间

- `--partial`

  rsync 默认会删除传输未完整的文件，使用该参数可以保留这些不完整的文件

- `--partial-dir=DIR`

  指定partial 存储的文件，以便调用rsync是加速

- `--prune-empty-dirs | -m`

  传输是删除 receiving side 的空文件，包括级联文件

### output

- `-v | --verbose`

  rsync 默认sliently，可以使用多个v详细输出。如果连接的是rsync daemon，且设置了 max verbosity ，就只能使用指定的 verbosity 等级

- `-q | --quiet`

  不输出从remote server的回显，一般用于crontab

- `--human-readable | -h`

  字面意思

- `--msgs2stderr`

  将stdout输出的内容输出到stderr

  ```
  [root@centos7 tmp]# rsync -av --msgs2stderr a/ b > /dev/null
  sending incremental file list
  
  sent 76 bytes  received 12 bytes  176.00 bytes/sec
  total size is 4  speedup is 0.05
  ```

- `--itemize-change | -i`

  以简短的方式显示rsync做了什么，可以使用`-ii`未改变的文件也会显示出来，一般以`YXcstpoguax`格式展示

  X 和 Y 可以代表的值具体查看man page

- `--stats`

  输出rsync statistics 信息

  ```
  ➜  test rsync  --stats a e 
  
  Number of files: 1 (reg: 1)
  Number of created files: 0
  Number of deleted files: 0
  Number of regular files transferred: 1
  Total file size: 2 bytes
  Total transferred file size: 2 bytes
  Literal data: 2 bytes
  Matched data: 0 bytes
  File list size: 0
  File list generation time: 0.001 seconds
  File list transfer time: 0.000 seconds
  Total bytes sent: 82
  Total bytes received: 35
  
  sent 82 bytes  received 35 bytes  234.00 bytes/se
  ```

  具体字段含义查看man page

- `--log-file=FILE`

  指定rsync daemon 日志存储的位置

  ```
  rsync -av --remote-option=--log-file=/tmp/rlog src/ dest/
  
  ```

  

### Daemon Argus

- `--daemon`

  告诉rsync 以 daemon 的方式运行，clinet 必须使用`host::module`或`rsync://host/module`来连接，daemon 默认会读取**rsyncd.conf**，具体配置查看man page

- `bwlimit=RATE`

  限制daemon的接受速率

- `--address=ADDRESS`

  daemon 绑定的IP，默认wildcard 绑定  

- `--config=FILE`

  指定deamon读取的配置文件

- `--dparam=OVERRIDE | -M`

  在启动daemon时指定的全局配置变量，优于配置文件

  ```
  rsync --daemon -M pidfile=/path/rsync.pid
  ```

- `--no-detach`

  以前端的方式运行rsync daemon

- `--port`

  daemon 使用的端口默认873

## filter rules

rsync 可以通过`--filter,-f`参数来过滤同步的文件，后面跟pattern

- `-`：exclude
- `+`：include
- `.`：merge
- `:`：dir-merge
- `H`：hide，指定传输隐藏文件的格式
- `P`：protect，指定不能被删除的文件的格式
- `R`：risk，指定可以被删除的文件的格式

```
➜  test rsync -a  --filter "- 5" a/ b/
➜  test ls a
 1   2   5   6   7   8
➜  test ls b
 1   2   3   6   7   8
```
## Examples

- a large of MS Word files and mail folers
```protobuf
rsync -Cavz . arvidsjaur:backup
```
