# Logrotate

## Overview

logrotate 是一个日志拆分工具, 通常会和 cron 一起使用

## Syntax

```
logrotate [options] configfile [configfile...]
```

## Postional Args

- configfile

  logrotate 配置的文件的路径

## Optional Args

- `-f | --force`

  让 logrotate 进行一次切割(不考虑设置的 frequecy 或者 size 以及其他条件)

- `-d | --debug`

  dry run 只显示切割的结果，但是不实际对文件应用

- `-v | --verbose`

## Configuration

可以将 logrotate 配置文件通常存储在以下两个位置 

```
(base) 0x00 in /etc/logrotate.d λ pacman -Ql logrotate 
...
logrotate /etc/logrotate.conf
logrotate /etc/logrotate.d/
...
```

### Directives

#### Rotation

- `rotate <count>`

  logrotate 为日志维护的队列长度,当队列满的时候会使用 FIFO 清空队列头

  例如 rotate 5，当一个文件历史上已经切割成 5 个后，最旧的文件会从队列移除(文件删除)，新的文件加入队列。即 FIFO

- `olddir <directory>; nooldir`

  将切割后的文件移动到指定的目录。反之保留在和被切割文件同目录

- `su <user> <group>`

  使用指定用户来切割日志，对应的 logrotate 配置文件也必须是匹配的

#### Frequency

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

  只有当文件大于 size 时才会切割，unit 默认 byte，可以使用 suffix

#### File selection

- `missingok; nomissingok`

  如果切割的文件的缺失，logrotate 不会报错。反之 `nomissingok`

- `ignoreduplicates`

  如果切割的文件有相同文件名的，logrotate 不会报错

- `ifempty; notifempty`

  即使切割的文件是空的 logrotate 也会切割。反之 `notifempty`

#### Files and Folders

- `shred；noshred`

  使用 `shred` 删除老的切割文件(不能恢复)。反之 `noshred`

#### compression

- `compress; nocompress`

  切割的旧日志会以 `gz` 的格式存储。反之 `nocompress`

- `compresscmd；uncompresscmd`

  指定解压缩使用的命令

- `delaycompress | nodelaycompress`

  文件会在下次执行 logrotate 时压缩。反之 `nodelaycompress`

#### filename

- `start <count>`

  切割的文件以 `filename.count` 的格式命名

- `dateext; nodateext`

  压缩后的文件以由 `dateformat` 指定的 `strftime` 的格式命名。反之 `nodateext`

- `dateformat <format_string>`

  指定 `dateext` 使用的后缀格式

- `dateyesterday`

  `dateext` 使用 yesterday 作为后缀

#### mail

- `mail <address>; nomail`

  如果文件开始切割通知指定的邮箱。反之 `nomail`

## What is rotate?

例如 rotate 5

你可以将其抽象成一个 length 5 的 quque, 如果指定切割被的文件已经被切割过 5 次，就会按照 FIFO 的原则将队列中最旧的日志替换出去，新切割的文件加入队列

## Example

假设现在有一个创建 `test.log`

```
(base) 0x00 in ~/test λ echo 1 > test.log
```

定义一个配置如下

> 为了方便测试定义以下几个参数
>
> `horly`
>
> `rotate 1`
>
> `start 0`

```
(base) 0x00 in ~/test λ cat logrotate.conf           
/home/0x00/test/test.log {
    #rotation
    hourly
    #frequency
    rotate 1
    #file selection
    missingok
    notifempty
    #files and folders
    noshred
    #compression
    compress
    #filename
    start 0
    #mail
    nomail
}

```

我们可以使用 `sudo logrotate -vdf /etc/logrotate.d/lastlog.conf` 以 dry run 的方式来校验配置是否生效

> 因为 `var/lib/logrotate.status` 是 640 的文件为了让 logrotate 记录文件的状态必须使用 root 来执行

```
(base) 0x00 in ~/test λ sudo logrotate -vdf logrotate.conf       
warning: logrotate in debug mode does nothing except printing debug messages!  Consider using verbose mode (-v) instead if this is not what you want.

reading config file logrotate.conf
Reading state from file: /var/lib/logrotate.status
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

rotating pattern: /home/0x00/test/test.log  forced from command line (1 rotations)
empty log files are not rotated, old logs are removed
considering log /home/0x00/test/test.log
  Now: 2023-11-10 17:09
  Last rotated at 2023-11-10 17:06
  log needs rotating
rotating log /home/0x00/test/test.log, log->rotateCount is 1
dateext suffix '-2023111017'
glob pattern '-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]'
renaming /home/0x00/test/test.log.0.gz to /home/0x00/test/test.log.1.gz (rotatecount 1, logstart 0, i 0), 
log /home/0x00/test/test.log.1.gz doesn't exist -- won't try to dispose of it
renaming /home/0x00/test/test.log to /home/0x00/test/test.log.0
compressing log with: /usr/bin/gzip
```

我们直接运行 `sudo logrotate -f logrotate.conf` 来生成一个切割的日志,这里会直接生成一个压缩文件(如果使用来 `compressdelay` 会在下次运行 `logrotate` 时对前一次切分的文件压缩)

```
(base) 0x00 in ~/test λ sudo logrotate -f logrotate.conf && ll
.rw-r--r-- root root 243 B Fri Nov 10 17:08:49 2023  logrotate.conf
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:18:09 2023  test.log.0.gz
```

然后我们再创建一个 `test.log` 并使用 `logrotate` 切分

```
(base) 0x00 in ~/test λ echo 2 >test.log && (sudo logrotate -f logrotate.conf && ll)
.rw-r--r-- root root 243 B Fri Nov 10 17:08:49 2023  logrotate.conf
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:21:17 2023  test.log.0.gz
```

这里可以看到原来的 `test.log.0.gz` 被替换掉了。因为 rotate 为 1 即队列长度只有 1，前一次切分的时候队列满了，按照 FIFO 的逻辑，旧的 `test.log.0.gz` 从队列中移除(即从磁盘中删除)，生成一个新的压缩文件对应内容为 2

```
(base) 0x00 in ~/test λ gzip -d test.log.0.gz && cat test.log.0 
2
```

我们现在将 `logrotate.conf` 中的 `rotate` 的值改为 3, 并将测试环境还原，然后使用 `logrotate` 来切分文件

```
(base) 0x00 in ~/test λ echo 1 > test.log                                           
(base) 0x00 in ~/test λ sudo logrotate -f logrotate.conf && ll
.rw-r--r-- root root 243 B Fri Nov 10 17:27:36 2023  logrotate.conf
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:16 2023  test.log.0.gz
(base) 0x00 in ~/test λ echo 2 > test.log && (sudo logrotate -f logrotate.conf && ll)
.rw-r--r-- root root 243 B Fri Nov 10 17:27:36 2023  logrotate.conf
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:33 2023  test.log.0.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:16 2023  test.log.1.gz
(base) 0x00 in ~/test λ echo 3 > test.log && (sudo logrotate -f logrotate.conf && ll)
.rw-r--r-- root root 243 B Fri Nov 10 17:27:36 2023  logrotate.conf
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:37 2023  test.log.0.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:33 2023  test.log.1.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:16 2023  test.log.2.gz
(base) 0x00 in ~/test λ echo 4 > test.log && (sudo logrotate -f logrotate.conf && ll)
.rw-r--r-- root root 243 B Fri Nov 10 17:27:36 2023  logrotate.conf
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:44 2023  test.log.0.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:37 2023  test.log.1.gz
.rw-r--r-- 0x00 0x00  22 B Fri Nov 10 17:29:33 2023  test.log.2.gz

```

可以发现压缩的文件只会有 3 个因为 rotate 值为 3

## How does logrotate work?

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

## Boilerplate Configuration

按照天切割文件，切后的后文件保留 14 天

```
#/etc/logrotate.d/zzcyzs
/data/jar/bigscreen.log {
    #rotation
    daily
    su root root
    #frequency
    rotate 14
    #file selection
    missingok
    notifempty
    #files and folders
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