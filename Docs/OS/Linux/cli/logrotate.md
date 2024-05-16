# Logrotate

## 0x01 Overview

logrotate 是一个日志切割工具, 通常会和 cron 或者是 rsyslog 一起使用，实现自动清理以及切割日志文件

在大多数的 distros 上都默认安装，系统会定时运行 logrotate

以 rhel7 为例，在 `/etc/cron.daily` 下有一个 `logrotate` 文件 （有些会有 `logrotate.service` 由 systemd 来控制）

```
[root@centos /]# cat /etc/cron.daily/logrotate
#!/bin/sh

/usr/sbin/logrotate -s /var/lib/logrotate/logrotate.status /etc/logrotate.conf
EXITVALUE=$?
if [ $EXITVALUE != 0 ]; then
    /usr/bin/logger -t logrotate "ALERT exited abnormally with [$EXITVALUE]"
fi
exit 0
```

即每天都会按照 `/etc/logrotate.conf` 中的内容执行 logrotate

```
[root@centos /]# cat /etc/logrotate.conf
# see "man logrotate" for details
# rotate log files weekly
weekly

# keep 4 weeks worth of backlogs
rotate 4

# create new (empty) log files after rotating old ones
create

# use date as a suffix of the rotated file
dateext

# uncomment this if you want your log files compressed
#compress

# RPM packages drop log rotation information into this directory
include /etc/logrotate.d

# no packages own wtmp and btmp -- we'll rotate them here
/var/log/wtmp {
    monthly
    create 0664 root utmp
        minsize 1M
    rotate 1
}

/var/log/btmp {
    missingok
    monthly
    create 0600 root utmp
    rotate 1
}

# system-specific logs may be also be configured here.
```

## 0x02 Syntax

```
logrotate [options] configfile [configfile...]
```

### Postional Args

- configfile

  logrotate 配置的文件的路径

### Optional Args

- `-f | --force`

  让 logrotate 进行一次切割(==不考虑设置的 frequecy 或者 size 以及其他条件==)

- `-d | --debug`

  ==dry run 只显示切割的结果，但是不实际对文件应用（用于测试配置文件）==

- `-v | --verbose`

  显示详细信息

- `-s | --state <statefile>`

  指定 logrotate 判断文件是否需要 rotate 的状态文件，默认 `/var/lib/logrotate.status`

## 0x03 Configuration

logrotate 配置文件通常存储在以下两个位置，会被系统自动调用（`cron` 或者是 `systemd` 中需要声明对应的配置）

```
(base) 0x00 in /etc/logrotate.d λ pacman -Ql logrotate 
...
logrotate /etc/logrotate.conf
logrotate /etc/logrotate.d/
...
```

### Configuration Sample

```
# sample logrotate configuration file
       compress

       /var/log/messages {
           rotate 5
           weekly
           postrotate
               /usr/bin/killall -HUP syslogd
           endscript
       }

       "/var/log/httpd/access.log" /var/log/httpd/error.log {
           rotate 5
           mail recipient@example.org
           size 100k
           sharedscripts
           postrotate
               /usr/bin/killall -HUP httpd
           endscript
       }

       /var/log/news/* {
           monthly
           rotate 2
           olddir /var/log/news/old
           missingok
           sharedscripts
           postrotate
               kill -HUP $(cat /var/run/inn.pid)
           endscript
           nocompress
       }

       ~/log/*.log {}

```

1. 第 2 行的 `compress` 是一个 global directive, 表示对所有在当前配置文件中的指令集都生效
2. `filename [...] {...}` 表示对应的指令集，对那些日志文件生效。filename 是一个变参，可以使用 wildcard

### Directives

#### Rotation

- `rotate <count>`

  logrotate 为日志维护的队列长度,当队列满的时候会使用 FIFO 清空队列头

  例如 rotate 5，当一个文件历史上已经切割成 5 个后，最旧的文件会从队列移除(文件删除)，新的文件加入队列。即 FIFO

  默认为 0，即只允许有 1 份日志，刚切出来的文件会立即被删除，==必须指定该指令==

- `olddir <directory>; nooldir`

  将切割后的文件（即 old logs）移动到指定的目录。反之保留在和被切割文件同目录

  directory 可以是相对（被切割的日志）路径，也可以是绝对路径

- `su <user> <group>`

  使用指定用户来切割日志，对应的 logrotate 配置文件也必须是匹配的

#### Frequency

> logrotate 是如何判断文件是否需要切割，具体逻辑看 `## logrotate.status` 部分

- `hourly`

  按照小时切割

- `daily`

  按照每天切割

- `weekly [weekday]`

  按照周几切割

- `monthly`

  按照月切割

- `yearly`

  按照年切割

- `size <size>`

  只有当被切割文件大于 size 时才会切割，unit 默认 byte，可以使用 suffix，例如 10K,10M

#### File selection 

- `missingok; nomissingok`

  如果被切割的文件的缺失，logrotate 不会报错。反之 `nomissingok`

- `ignoreduplicates`

  如果被切割的文件有相同文件名的，logrotate 不会报错

- `ifempty; notifempty`

  即使被切割的文件是空的 logrotate 也会切割。反之 `notifempty`

#### Files and Folders

> 为了能周期性切割日志，通常需要指定 `create`,``

- `create [mode] <owner> <group>；nocreate`

  切割生成文件后，logrotate 会创建和被切割文件同名文件(==意味着新创建的被切割文件 inode 和原始文件的不同。有些进程对原始文件的 fd 没有关闭，会向等于原始文件 inode 的文件中继续写入，即切割生成的文件==)，以 `[mode] <owner> <group>` 的格式被创建

  假设对 `test.log` 做了两次切割，生成 `test.log` 和 `test.log.0` 以及 `test.log.1`

  ```
  before rotate(test.log.inode) == first rotate(test.log.0.inode) == second rotate(test.log.1.inode)
  before rotate(test.log.inode) != (first rotate(test.log.inode) || second rotate(test.log.inode))
  ```

- `shred；noshred`

  使用 `shred` 删除老的切割文件(不能恢复)。反之 `noshred`

- `copy; nocopy`

  对被切割文件做复制生成切割文件(==生成的切割文件和被切割的文件 inode 不同==)，原文件中的内容不会改变(==被切割文件 inode 保持不变==)，而不是将被切割文件重命名以生成切割文件

  假设对 `test.log` 做了两次切割，生成 `test.log` 和 `test.log.0` 以及 `test.log.1`

  ```
  before rotate(test.log.inode) == first rotate(test.log.inode) == second rotate(test.log.inode)
  before rotate(test.log.inode) != (first rotate(test.log.0.inode) || second rotate(test.log.0.inode) || second rotate(test.log.1.inode))
  ```

  ==因为切割文件时被切割文件中的内容不会改变，所以切割生成的文件等于本次被切割文件的全量内容,即 `test.log.0` 中保留 `test.log.1` 中的内容，`test.log.0` 中的内容和 `test.log` 相同，依次类推==

  如果有其他的程序需要对应的日志文件做操作，就可以使用该指令，来保证原始的日志文件 inode 不会改变,这样即使其他程序打开了对应 inode 的文件，内容也不会改变

  如果使用了该指令，`create` 指令就不会生效

- `copytruncate; nocopytruncate`

  和 `copy` 类似，但是会对被切割文件做清空的操作，即切割生成的文件对比前一次切割生成的文件为增量（增量部分为 `本次被切割文件的内容 - 前一次切割生成的文件的内容`）

  为增量记录日志
  
  对被切割文件做复制(==生成的切割文件和被切割的文件 inode 不同==)，复制完成后，将原被切割文件清空 to zero size（==被切割文件 inode 保持不变==）,而不是将被切割文件重命名以生成切割文件
  
  假设对 `test.log` 做了两次切割，生成 `test.log` 和 `test.log.0` 以及 `test.log.1`
  
  ```
  before rotate(test.log.inode) == first rotate(test.log.inode) == second rotate(test.log.inode)
  before rotate(test.log.inode) != (first rotate(test.log.0.inode) || second rotate(test.log.0.inode) || second rotate(test.log.1.inode))
  ```
  
  因为会清空文件，所以和重命名差不多，但是不同的是 inode 保持不变，这样可以保证打开对应 inode 文件的程序会读取相同的文件
  
  如果是重命名，以 `test.log` 做一次切割生成 `test.log` 和 `test.log.0` 为例，有的程序因为没有关闭 fd 就可能会读取 `test.log.0` 中的内容，而不是 `test.log` 中的内容，因为 `test.log` 是新创建的文件（虽然文件名相同，但是 inode 改变了），原 `test.log` 被重命名为 `test.log.0` (inode 不变)
  
  如果使用了该指令，`create` 指令就不会生效

#### Compression

- `compress; nocompress`

  切割生成的旧日志会以 `gz` 的格式存储。反之 `nocompress`

- `compresscmd；uncompresscmd`

  指定解压缩使用的命令

- `delaycompress | nodelaycompress`

  文件会在下次执行 logrotate 时压缩（即前一次执行的 logrotate 生成的文件，会在下一次执行 logrotate 时压缩前一次生成的文件）。反之 `nodelaycompress`

#### Filename

> 如果文件名相同，logrotate 不会切割日志

- `start <count>`

  切割生成的文件以 `filename.count` 的格式命名

- `dateext; nodateext`

  切割生成的文件以由 `dateformat` 指定的 `strftime` 的格式命名。反之 `nodateext`

- `dateformat <format_string>`

  指定 `dateext` 使用的后缀格式

- `dateyesterday`

  `dateext` 使用 yesterday 作为后缀

#### Mail

- `mail <address>; nomail`

  如果文件开始切割通知指定的邮箱。反之 `nomail`

#### Additional config fiel

- `include <file_or_directory>`

  和 nginx 中的 `include` 相同表示导入一组文件

  ```
  include /etc/logrotate.d
  ```

#### Scripts

- 

## 0x04 Default vs Create vs Copy vs Copytruncate

是否使用 `create`,`copy`,`copytruncate` 决定了 logrotate 该如何切割文件

现有如下配置文件

```
~/test/test.log {
    daily
    #create 0644 root root
    #copy
    #copytruncate
    rotate 2
    missingok
    ifempty
    noshred
    start 0
    nomail
}
```

### default

表示默认，即配置文件中不含有 `create`，`copy` 或者是 `copytruncate`

切割前先看一下 inode

```
[root@centos test]# stat test.log
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-03 06:55:12.121538789 +0000
Modify: 2024-04-03 06:55:12.121538789 +0000
Change: 2024-04-03 06:55:12.121538789 +0000
 Birth: -
```

使用 logrotate 对 `test.log` 做切割结果如下

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 2
dateext suffix '-20240403'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.1 to /root/test/test.log.2 (rotatecount 2, logstart 0, i 1),
old log /root/test/test.log.1 does not exist
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 2, logstart 0, i 0),
old log /root/test/test.log.0 does not exist
log /root/test/test.log.2 doesn't exist -- won't try to dispose of it
fscreate context set to unconfined_u:object_r:admin_home_t:s0
renaming /root/test/test.log to /root/test/test.log.0
set default create context
```

观察一下生成的文件和 inode

```
[root@centos test]# ll
total 4
-rw-r--r--. 1 root root 141 Apr  8 03:19 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  3 06:55 test.log.0
[root@centos test]# stat test.log.0
  File: ‘test.log.0’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-03 06:55:12.121538789 +0000
Modify: 2024-04-03 06:55:12.121538789 +0000
Change: 2024-04-08 03:19:11.195233296 +0000
 Birth: -
```

这里可以看到 logrotate 并不会生成一个新的 `test.log`，即只是做切割。同时原始的 `test.log` 和 `test.log.0` 的 inode 相同

再次运行 logrotate 时会报错，因为没有对应的 `test.log`

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log ~/test/test.log
set default create context
[root@centos test]# ll
total 8
-rw-r--r--. 1 root root 141 Apr  3 06:44 logrotate.conf
-rw-r--r--. 1 root root   2 Apr  3 06:32 test.log.0
```

### create

使用 `create 0644 root root` ，并复原 `test.log`(还原测试环境)

切割前先看一下 inode

```
[root@centos test]# stat test.log
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-03 06:55:12.121538789 +0000
Modify: 2024-04-03 06:55:12.121538789 +0000
Change: 2024-04-08 03:21:48.535257958 +0000
 Birth: -
```

使用 logrotate 对 `test.log` 做切割结果如下

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 2
dateext suffix '-20240403'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.1 to /root/test/test.log.2 (rotatecount 2, logstart 0, i 1),
old log /root/test/test.log.1 does not exist
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 2, logstart 0, i 0),
old log /root/test/test.log.0 does not exist
log /root/test/test.log.2 doesn't exist -- won't try to dispose of it
fscreate context set to unconfined_u:object_r:admin_home_t:s0
renaming /root/test/test.log to /root/test/test.log.0
creating new /root/test/test.log mode = 0644 uid = 0 gid = 0
set default create context
```

1. `renaming /root/test/test.log to /root/test/test.log.0`

   将原被切割文件重命名，为 `test.log.0`

   此时被切割文件 `test.log` 在切割后和 `test.log.0` inode 相同，即做了重命名的动作

2. `creating new /root/test/test.log mode = 0644 uid = 0 gid = 0`

   生成一个新的 `test.log` 
   
   意味着切割后的创建的 `test.log` 和原始切割文件 `test.log` inode 不同

观察一下生成的文件和 inode

```
[root@centos test]# ll
total 4
-rw-r--r--. 1 root root 140 Apr  8 03:22 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  8 03:23 test.log
-rw-r--r--. 1 root root   0 Apr  3 06:55 test.log.0
[root@centos test]# stat test.log
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110589    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-08 03:23:03.465899163 +0000
Modify: 2024-04-08 03:23:03.465899163 +0000
Change: 2024-04-08 03:23:03.465899163 +0000
 Birth: -
[root@centos test]# stat test.log.0
  File: ‘test.log.0’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-03 06:55:12.121538789 +0000
Modify: 2024-04-03 06:55:12.121538789 +0000
Change: 2024-04-08 03:23:03.465899163 +0000
 Birth: -
```

==即如果有进程的 fd 针对原始被切割文件 `test.log` 没有关闭，可能会使用切割后的 `test.log.0` 因为 inode 和原始的被切割文件 `test.log` 相同==

### copy

使用 `copy`，并复原 `test.log` (还原测试环境)

切割前先看一下 inode 和 文件内容

```
[root@centos test]# stat test.log
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-03 06:55:12.121538789 +0000
Modify: 2024-04-03 06:55:12.121538789 +0000
Change: 2024-04-08 03:26:55.438014257 +0000
 Birth: -
 
[root@centos test]# cat test.log
this is a test
```

使用 logrotate 对 `test.log` 做切割结果如下

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 2
dateext suffix '-20240408'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.1 to /root/test/test.log.2 (rotatecount 2, logstart 0, i 1),
old log /root/test/test.log.1 does not exist
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 2, logstart 0, i 0),
old log /root/test/test.log.0 does not exist
log /root/test/test.log.2 doesn't exist -- won't try to dispose of it
copying /root/test/test.log to /root/test/test.log.0
set default create context to unconfined_u:object_r:admin_home_t:s0
Not truncating /root/test/test.log
set default create context
```

1. `copying /root/test/test.log to /root/test/test.log.0`

   使用原文件复制生成 `test.log.0` （内容相同）

   此时被切割文件 `test.log` 在切割后和 `test.log.0` inode 不同，`test.log` inode 保持不变

观察一下生成的文件和 inode 以及内容

```
[root@centos test]# ll
total 4
-rw-r--r--. 1 root root 140 Apr  8 03:26 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  3 06:55 test.log
-rw-r--r--. 1 root root   0 Apr  8 03:28 test.log.0
[root@centos test]# stat test.log
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-08 03:28:10.500727914 +0000
Modify: 2024-04-03 06:55:12.121538789 +0000
Change: 2024-04-08 03:26:55.438014257 +0000
 Birth: -
[root@centos test]# stat test.log.0
  File: ‘test.log.0’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110589    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-08 03:28:10.500727914 +0000
Modify: 2024-04-08 03:28:10.500727914 +0000
Change: 2024-04-08 03:28:10.500727914 +0000
 Birth: -
[root@centos test]# cat test.log
this is a test
[root@centos test]# cat test.log.0
this is a test
```

往 `test.log` 增加内容

```
[root@centos test]# echo "this is another test" >> test.log
```

再次运行 logrotate

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 2
dateext suffix '-20240408'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.1 to /root/test/test.log.2 (rotatecount 2, logstart 0, i 1),
old log /root/test/test.log.1 does not exist
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 2, logstart 0, i 0),
log /root/test/test.log.2 doesn't exist -- won't try to dispose of it
copying /root/test/test.log to /root/test/test.log.0
set default create context to unconfined_u:object_r:admin_home_t:s0
Not truncating /root/test/test.log
set default create context
```

观察文件内容

```
[root@centos test]# cat test.log
this is a test
this is another test
[root@centos test]# cat test.log.0
this is a test
this is another test
[root@centos test]# cat test.log.1
this is a test
```

这里可以观察到 `test.log.0` 对 `test.log` 做了全量记录(也可以将这种全量记录，理解成快照)

==即如果有进程的 fd 针对原始被切割文件 `test.log` 没有关闭，进程还是会读取原始的被切割文件，因为 `test.log` inode 未发生变化。同时切割生成的文件，为本次被切割文件的全量==

### copytruncate

使用 `copytruncate`，并复原 `test.log` (还原测试环境)

这里只对比文件内容，inode 的逻辑和 `copy` 相同

切割前先看一下文件内容

```
[root@centos test]# cat test.log
this is a test
```

使用 logrotate 对 `test.log` 做切割结果如下

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 2
dateext suffix '-20240408'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.1 to /root/test/test.log.2 (rotatecount 2, logstart 0, i 1),
old log /root/test/test.log.1 does not exist
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 2, logstart 0, i 0),
old log /root/test/test.log.0 does not exist
log /root/test/test.log.2 doesn't exist -- won't try to dispose of it
copying /root/test/test.log to /root/test/test.log.0
set default create context to unconfined_u:object_r:admin_home_t:s0
truncating /root/test/test.log
set default create context
```

1. `copying /root/test/test.log to /root/test/test.log.0`

   使用原文件复制生成 `test.log.0` （内容相同）

   此时被切割文件 `test.log` 在切割后和 `test.log.0` inode 不同，`test.log` inode 保持不变

2. `truncating /root/test/test.log`

   清空 `test.log` 中的内容

观察一下生成的文件以及内容

```
[root@centos test]# ll
total 8
-rw-r--r--. 1 root root 158 Apr  8 07:45 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  8 07:46 test.log
-rw-r--r--. 1 root root  15 Apr  8 07:46 test.log.0
[root@centos test]# cat test.log
[root@centos test]# cat test.log.0
this is a test
```

这里可以发现被切割文件 `test.log` 中的内容为空，因为被 truncate 掉了

往 `test.log` 增加内容

```
[root@centos test]# echo "this is another test" >> test.log
```

再次运行 logrotate

```
[root@centos test]# logrotate -vf logrotate.conf
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (2 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 2
dateext suffix '-20240408'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.1 to /root/test/test.log.2 (rotatecount 2, logstart 0, i 1),
old log /root/test/test.log.1 does not exist
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 2, logstart 0, i 0),
log /root/test/test.log.2 doesn't exist -- won't try to dispose of it
copying /root/test/test.log to /root/test/test.log.0
set default create context to unconfined_u:object_r:admin_home_t:s0
truncating /root/test/test.log
set default create context
```

观察文件内容

```
[root@centos test]# ll
total 12
-rw-r--r--. 1 root root 158 Apr  8 07:45 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  8 07:52 test.log
-rw-r--r--. 1 root root  21 Apr  8 07:52 test.log.0
-rw-r--r--. 1 root root  15 Apr  8 07:46 test.log.1
[root@centos test]# cat test.log
[root@centos test]# cat test.log.0
this is another test
[root@centos test]# cat test.log.1
this is a test
```

这里可以发现每次切割，都会将 `test.log` 清空

==即如果有进程的 fd 针对原始被切割文件 `test.log` 没有关闭，进程还是会读取原始的被切割文件，因为 `test.log` inode 未发生变化。同时切割生成的文件内容，为当前被切割文件对比前一次切割生成的文件中内容的增量==

## 0x05 What is rotate?

例如 rotate 5

你可以将其抽象成一个 length 5 的 quque, 如果指定切割被的文件已经被切割过 5 次（以 `glob pattern` 来确定当前队列中有多少成员，切割次数），就会按照 FIFO 的原则将队列中最旧的日志替换出去，新切割的文件加入队列

## 0x06 Lab

假设现在有一个创建 `test.log` 需要对该文件进行切割

```
[root@centos test]# echo 1 > test.log
```

定义一个配置如下，其中

`rotate 1` (定义一个长度为 1 的一个队列)

`start 0` (防止以日期命名切割文件，导致 logrotate 不成生文件)

```
[root@centos test]# cat logrotate.conf
~/test/test.log {
    #rotation
    daily
    #frequency
    rotate 1
    #file selection
    missingok
    ifempty
    #files and folders
    create 0644 root root
    noshred
    #filename
    start 0
    #mail
    nomail
}
```

> 需要注意几点：
>
> 1. 以 dry-run 的方式运行 logrotate 输出的结果可能不同，但是逻辑相同。这里为了方便对应实际生产直接对文件切割
>
> 2. 因为 `var/lib/logrotate.status` 是 640 的文件(logratate 会读取并写入该文件)为了让 logrotate 记录文件的状态必须使用 root 来执行

```
[root@centos test]# ll && stat test.log && logrotate -vf logrotate.conf && ll && stat test.log
total 8
-rw-r--r--. 1 root root 226 Apr  2 09:25 logrotate.conf
-rw-r--r--. 1 root root   2 Apr  2 09:27 test.log
  File: ‘test.log’
  Size: 2               Blocks: 8          IO Block: 4096   regular file
Device: 801h/2049d      Inode: 68110587    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-02 09:26:44.322474657 +0000
Modify: 2024-04-02 09:27:23.811571197 +0000
Change: 2024-04-02 09:27:23.811571197 +0000
 Birth: -
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (1 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 1
dateext suffix '-20240402'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 1, logstart 0, i 0),
old log /root/test/test.log.0 does not exist
log /root/test/test.log.1 doesn't exist -- won't try to dispose of it
fscreate context set to unconfined_u:object_r:admin_home_t:s0
renaming /root/test/test.log to /root/test/test.log.0
creating new /root/test/test.log mode = 0644 uid = 0 gid = 0
set default create context
total 8
-rw-r--r--. 1 root root 226 Apr  2 09:25 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  2 09:28 test.log
-rw-r--r--. 1 root root   2 Apr  2 09:27 test.log.0
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110589    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-02 09:28:19.992708535 +0000
Modify: 2024-04-02 09:28:19.992708535 +0000
Change: 2024-04-02 09:28:19.992708535 +0000
 Birth: -
[root@centos test]#
```

1. `rotating pattern: ~/test.log  forced from command line (1 rotations)`

   对应配置文件的中 `~/test.log { ... }` 表示指令集对 `test.log` 生效

2. `empty log files are rotated, old logs are removed`

   对应 `ifempty` 如果原始日志文件为空，也会被切割

3. `rotating log /root/test.log, log->rotateCount is 1`

   对应 `rotate 1` 

4. `renaming /root/test.log.0 to /root/test.log.1 (rotatecount 1, logstart 0, i 0),`

   对应 `rotate 1` ， `start 0`

   logrotate 会先将队列中的 `test.log.0` 重命名为 `test.log.1`，腾出队列

5. `old log /root/test.log.0 does not exist`

   `log /root/test.log.1 doesn't exist -- won't try to dispose of it`

   logrotate 会尝试将队列尾 `rotate 1` 的文件重命名，这里会逻辑上会将 `test.log.0` 重命名为 `test.log.1`

   然后 lorotate 会尝试删除 `rotate 1` 队列外的文件，这里对应 `test.log.1` (即切割前的)

6. `creating new /root/test.log mode = 0644 uid = 0 gid = 0`

   创建一个新的空文件 `test.log` (这一点可以由 `stats test.log` 的 inode 和 ctime 确定 )，表示原始日志(用于后续切割)，以 `create` 中定义的模式

再运行一次 logrotate

```
[root@centos test]# ll && stat test.log && logrotate -vf logrotate.conf && ll && stat test.log
total 8
-rw-r--r--. 1 root root 226 Apr  2 09:25 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  2 09:28 test.log
-rw-r--r--. 1 root root   2 Apr  2 09:27 test.log.0
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110589    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-02 09:28:19.992708535 +0000
Modify: 2024-04-02 09:28:19.992708535 +0000
Change: 2024-04-02 09:28:19.992708535 +0000
 Birth: -
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (1 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 1
dateext suffix '-20240402'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 1, logstart 0, i 0),
fscreate context set to unconfined_u:object_r:admin_home_t:s0
renaming /root/test/test.log to /root/test/test.log.0
creating new /root/test/test.log mode = 0644 uid = 0 gid = 0
removing old log /root/test/test.log.1
set default create context
total 4
-rw-r--r--. 1 root root 226 Apr  2 09:25 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  2 09:31 test.log
-rw-r--r--. 1 root root   0 Apr  2 09:28 test.log.0
  File: ‘test.log’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 68110590    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:admin_home_t:s0
Access: 2024-04-02 09:31:22.503154714 +0000
Modify: 2024-04-02 09:31:22.503154714 +0000
Change: 2024-04-02 09:31:22.503154714 +0000
 Birth: -
```

1. `renaming /root/test/test.log.0 to /root/test/test.log.1 (rotatecount 1, logstart 0, i 0),`

   这里可以看到 logrotate 将 `test.log.0` 重命名为 `test.log.1`

2. `removing old log /root/test/test.log.1`

   因为 `rotate 1`, 新生成的 `test.test.log` 加入队列，队列中的 `test.log.1` 从队列中删除

在原有的配置中添加 `compress`

```
[root@centos test]# cat logrotate.conf
~/test/test.log {
  ...
	compress
  ...
}
```

那么切割生成的文件会以 `gzip` 格式存储

```
[root@centos test]# ll &&  logrotate -vf logrotate.conf && ll
total 4
-rw-r--r--. 1 root root 239 Apr  2 09:38 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  2 09:31 test.log
-rw-r--r--. 1 root root   0 Apr  2 09:28 test.log.0
reading config file logrotate.conf
Allocating hash table for state file, size 15360 B

Handling 1 logs

rotating pattern: ~/test/test.log  forced from command line (1 rotations)
empty log files are rotated, old logs are removed
considering log /root/test/test.log
  log needs rotating
rotating log /root/test/test.log, log->rotateCount is 1
dateext suffix '-20240402'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /root/test/test.log.0.gz to /root/test/test.log.1.gz (rotatecount 1, logstart 0, i 0),
old log /root/test/test.log.0.gz does not exist
log /root/test/test.log.1.gz doesn't exist -- won't try to dispose of it
fscreate context set to unconfined_u:object_r:admin_home_t:s0
renaming /root/test/test.log to /root/test/test.log.0
creating new /root/test/test.log mode = 0644 uid = 0 gid = 0
compressing log with: /bin/gzip
set default create context to unconfined_u:object_r:admin_home_t:s0
set default create context
total 8
-rw-r--r--. 1 root root 239 Apr  2 09:38 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  2 09:40 test.log
-rw-r--r--. 1 root root  20 Apr  2 09:31 test.log.0.gz

```

现在将 `rotate` 置为 3，更加贴合实际，同时使用 `delaycompress`

```
[root@centos test]# cat logrotate.conf
~/test/test.log {
  ...
	rotate 3
	...
	compress
	delaycompress
  ...
}
```

那么调用 logrotate ==n== 次，就只会保留 3 个文件（2 个 `gzip` 压缩文件，因为使用了 `delaycompress`），队列头的 `gzip` 文件会被移除

生成的文件如下

```
-rw-r--r--. 1 root root 257 Apr  2 09:42 logrotate.conf
-rw-r--r--. 1 root root   0 Apr  2 09:45 test.log
-rw-r--r--. 1 root root   0 Apr  2 09:45 test.log.0
-rw-r--r--. 1 root root  20 Apr  2 09:45 test.log.1.gz
-rw-r--r--. 1 root root  20 Apr  2 09:45 test.log.2.gz
```

## 0x07 logrotate.status

logrotate 自己维护着一个简单的数据库(默认存储在 `/var/lib/logrotate.status` 也可以使用 `--state <statefile>` 来指定)用于计算 frequency 是否达到匹配规则。

```
(base) 0x00 in ~/test λ sudo cat /var/lib/logrotate.status
logrotate state -- version 2
"/home/0x00/test/test.log" 2023-11-10-17:29:44
```

例如配置文件中使用 `rotate 1 hourly` 那么只有在当前系统的时间大于 statefile 中指定文件的 timestamp  

1 小时 logrotate 才会切割日志（如果使用 logrotate 时如果没有带上 `-f` 参数）

```
(base) 0x00 in ~/test λ echo 5 > test.log && (sudo logrotate -v logrotate.conf && ll)
reading config file logrotate.conf
acquired lock on state file /var/lib/logrotate.statusReading state from file: /var/lib/logrotate.status
Allocating hash table for state file, size 64 entries
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state
Creating new state

Handling 1 logs

rotating pattern: /home/0x00/test/test.log  hourly (3 rotations)
empty log files are not rotated, old logs are removed
considering log /home/0x00/test/test.log
  Now: 2023-11-10 17:54
  Last rotated at 2023-11-10 17:29
  log does not need rotating (log has already been rotated)
.rw-r--r-- root root 243 B Fri Nov 10 17:27:36 2023  logrotate.conf
.rw-r--r-- 0x00 0x00   2 B Fri Nov 10 17:54:04 2023  test.log
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:44 2023  test.log.0.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:37 2023  test.log.1.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:33 2023  test.log.2.gz
```

对应的源码逻辑如下

主要看 [logrotate.c](https://github.com/logrotate/logrotate/blob/main/logrotate.c) 的 readState 函数(configure.ac 是 autobuild 的配置用于 shell 注入)

```
#configure.ac
STATE_FILE_PATH="/var/lib/logrotate.status"
AC_DEFINE_UNQUOTED([STATEFILE], ["$STATE_FILE_PATH"], [State file path.])

#logrotate.c
int main(int argc, const char **argv)
	...
	const char *stateFile = STATEFILE;
	...
  if (readState(stateFile))
  	rc = 1;

static int readState(const char *stateFilename) 	
	...
	fd = open(stateFilename, O_RDONLY);
  ...
  while (fgets(buf, sizeof(buf) - 1, f)) {
    const size_t i = strlen(buf);
    char *filename;
    int argc;
    const char **argv = NULL;
    int year, month, day, hour, minute, second;
    struct logState *st;
    time_t lr_time;
  	...
  	filename = strdup(argv[0]);
  	...
  	#关键在这看配置文件
    if ((st = findState(filename)) == NULL) {
    	free(argv);
    	free(filename);
    	fclose(f);
    	return 1;
    }
    memset(&st->lastRotated, 0, sizeof(st->lastRotated));
    st->lastRotated.tm_year = year;
    st->lastRotated.tm_mon = month;
    st->lastRotated.tm_mday = day;
    st->lastRotated.tm_hour = hour;
    st->lastRotated.tm_min = minute;
    st->lastRotated.tm_sec = second;
    st->lastRotated.tm_isdst = -1;
  	lr_time = mktime(&st->lastRotated);
    localtime_r(&lr_time, &st->lastRotated);
    ...
```

需要提一嘴的是 logrotate.status 记录的时间并不切分文件的 mtime，而是 logrotate 调用 localtime_r 的时间，主要看 [wirteState](https://sourcegraph.com/github.com/logrotate/logrotate/-/blob/logrotate.c?L2616) 这个函数

```
static int writeState(const char *stateFilename)
	...
	localtime_r(&nowSecs, &now);
	...
  now_time = mktime(&now);
  last_time = mktime(&p->lastRotated);
  if (!p->isUsed && difftime(now_time, last_time) > SECONDS_IN_YEAR) {
  	message(MESS_DEBUG, "Removing %s from state file, "
  		"because it does not exist and has not been rotated for one year\n",
  		p->fn);
  	continue;
  }
```

## 0x08 Examples

- `nginx`

  通常使用包管理器下载 nginx 会自动在 `/etc/logrotate.d` 下添加一个 nginx 配置，而如果使用编译安装需要手动添加

  ```
  (base) cc in /etc/logrotate.d λ cat nginx
  /var/log/nginx/*log {
          missingok
          notifempty
          create 640 http root
          sharedscripts
          compress
          postrotate
                  test ! -r /run/nginx.pid || kill -USR1 `cat /run/nginx.pid`
          endscript
  }
  ```

- `zabbix-agent`

  ```
  (base) cc in /etc/logrotate.d λ cat zabbix-agent
  /var/log/zabbix/zabbix_agentd.log {
  	weekly
  	rotate 12
  	compress
  	delaycompress
  	missingok
  	notifempty
  	create 0664 zabbix zabbix
  }
  ```

- `libvirtd`

  ```
  (base) cc in /etc/logrotate.d λ cat libvirtd
  /var/log/libvirt/libvirtd.log {
          weekly
          missingok
          rotate 4
          compress
          delaycompress
          copytruncate
          minsize 100k
  }
  
  ```

## 9x09 Boilerplate Configuration

按照天切割文件，切割的后文件保留 14 天

```
#/etc/logrotate.d/zzcyzs
/data/jar/bigscreen.log {
    #rotation
    daily
    #frequency
    rotate 14
    #file selection
    missingok
    ifempty
    #files and folders
    copytruncate
    noshred
    #compression
    compress
    delaycompress
    #filename
    dateext
    dateformat -%Y%m%d
    #mail
    nomail
}
```

为了让 logrotate 能定时启动还需要和 crontab 一起配置使用

```
sudo crontab -e 
0 3 * * * /usr/bin/logrotate -f /etc/logrotate.d/zzcyzs >& /dev/null &
```

**references**

[^1]:https://wsgzao.github.io/post/logrotate/
[^2]:https://github.com/logrotate/logrotate/tree/main

**changeLog**

| Date     | Comment             |
| -------- | ------------------- |
| 20240403 | Create logrotate.md |

