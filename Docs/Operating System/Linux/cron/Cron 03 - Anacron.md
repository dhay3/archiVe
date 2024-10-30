---
createTime: 2024-10-28 15:00
tags:
  - "#hash1"
  - "#hash2"
---

# Cron 03 - Anacron

## 0x01 Preface

> [!important]
> 本文会以 Modern Version cron 作为基础来介绍 cron

cron 是 Linux 上的一个 job scheduler，可以通过 cron 实现在指定时间运行指定指令

## 0x02 History[^1]

> cron originates from Chronos

### 0x02a Early Versions

在 Versino 7 Unix 的时代，只能以 root 的身份运行 cron，逻辑非常直白

1. Read `/usr/lib/crontab`
2. Determine if any commands must run at the current date and time, and if so, run them as the superuser, root
3. Sleep for one minute
4. Repeat from step 1

### 0x02b Multi-user Capability

随着 SysV 的到来，cron 从只针对 root，扩展为其他的用户也可以使用 cron

1. On start-up, look for a file name `.crontab` in the home directory of all account holders
2. For each crontab file found, determine the next time in the future that each command must run
3. Place those commands on the Franta–Maly event list(相对早期的版本，效率更高) with their corresponding time and their "five field" time specifier.
4. Enter main loop:
	- Examine the task entry at the head of the queue, compute how far in the future it must run.
	- Sleep for that period of time.
	- On awakening and after verifying the correct time, execute the task at the head of the queue (in background) with the privileges of the user who created it.
	- Determine the next time in the future to run this command and place it back on the event list at that time value.

同时 cron 默认会将 stdout 和 stderr 输出的内容发送到 `/var/spool/mail/${USER}`(可以通过 `MAILTO` 环境变量设置)

再后来，Keith Williamson 将 `at` 和 cron 做了合并，并将 crontab file 从用户目录移动至 spool directory(`/var/spool/cron`，文件名为用户名)。同时增加了 `crontab` 命令，可以让用户直接在 spool directory 下生成 crontab file

这也是我们最熟悉的形式

### 0x02c Modern Versions

在 GNU/Linux 的运动下，由 Pual Vixie 开发的 cron 占据了主流，大多数的 distros 都使用了 Vixie cron

在 2007 Red hat 克隆了 Vixie cron，添加了 PAM 以及 SELinux 的支持，同时加入了 anacron，并将这个项目称为 *cronie*（这也是目前大多数 distro 使用的 cron）


## 0x08 Anacron[^4]

anacron 是 cronie 中的一部分，但是和 cron 不一样。cron 会认为主机是一直运行的，而 anacron 不认为主机是一直运行的

假设 cron 和 anacron 的使用方式都一致，有如下 scheduled task

```
0 4 * * * /usr/bin/local/trojan-shell -h 223.5.5.5
```

如果配置了该 scheduled task 的主机，每天的 23:00 - 06:00 都执行关机，在第二天的 10:00 开机

- cron 就不会在开机后执行该 scheduled task，因为 cron 认为主机是一直运行的，过了就不会执行
- anacron 会根据 timestamp files(存储在 `/var/spool/anacron`，后面会提到) 来判断开机后是否需要执行该 scheduled task

但是实际上 cron 和 anacron 的字段不一样，cron 有 5 个字段来表示 cron part，而 anacron 只有 2 个字段来表示 cron part

## 0x09 Anacrontab Files

anacron 默认会从 `/etc/anacrontab` 读取 anacron 任务

anacrontab files 的格式和 crontab files 不同

### 0x09a Comments and Nonsense Characters

anacrontab files 注释，空行 和 crontab files 一致([0x04a Comments and Nonsense Characters](#0x04a%20Comments%20and%20Nonsense%20Characters))

### 0x09b Anacron Entries

Anacron Entries 每一行代表一个任务

```
#period in days   delay in minutes   job-identifier   command
1       5       cron.daily              nice run-parts /etc/cron.daily
```

细分为 4 部分

#### period in days

以天为单位，执行的频率。可以是一个 integer 也可以是一个 marco 例如

- `@daily`
	等价于 1
- `@weekly`
	等价于 7
- `@month`
	根据月份来判断

#### delay in minutes

anacron 会在 delay in minutes 后执行 scheduled job

#### job identifier

对应 schedule job 的唯一标识符

#### command

实际执行的命令或者是脚本

### 0x09c Environments

anacrontab files 中的环境变量大体上和 crontab files 一致([0x04b Environments](#0x04c%20Environments))，但额外增加一个变量

- `START_HOURS_RANGE`
	指定了 anacron scheduled job 允许在什么时间段内运行，如果 scheduled job 不在这个时间段内就不会运行
- `RANDOM_DELAY`
	为每个 schedule job 额外随机增加的最大的 delay(in mintues)，例如 `RANDOM_DELAY=12` 表示为每个 schedule job 额外随机增加 0 - 12 minutes delay。如果值为 0 表示不额外增加

## 0x09 Anacron Timestamp Files

Anacron timestamp files 是 anacron 判断 scheduled tasks 是否执行的因子，

## 0x10 Anacron Cmd[^5]

Anacron 没有类似 `crontab` 的命令，用于管理 anacrontab files。但是提供了一个告诉 anacron 该如何执行 scheduled task 的指令 `anacron`

```
 anacron [-s] [-f] [-n] [-d] [-q] [-t anacrontab] [-S spooldir] [job]
 anacron [-S spooldir] -u [-t anacrontab] [job]
 anacron [-V|-h]
 anacron -T [-t anacrontab]
```

### 0x10a Optional args

- `-f`
	不考虑 timestamp files，立刻执行 `/etc/anacrontab` 中的内容
- `-u`
	将 timestamp files




## 0x11 Source Code Analyzing

### 0x11a Cron


### 0x11b Anacron


## 0x10 Cron VS Anacron Summrize

***References***

- [cron - Wikipedia](https://en.wikipedia.org/wiki/Cron)
- [GitHub - cronie-crond/cronie: Cronie cron daemon project](https://github.com/cronie-crond/cronie?tab=readme-ov-file)
- `man cron.8`
- `man crontab.1`
- `man crontab.5`
- `man run-parts`
- `man anacron.8`
- `man anacrontab.5`
- [anacron - Wikipedia](https://en.wikipedia.org/wiki/Anacron)
- [Confused about relationship between cron and anacron - Ask Ubuntu](https://askubuntu.com/questions/848610/confused-about-relationship-between-cron-and-anacron)

***FootNotes***

[^1]:[cron - Wikipedia](https://en.wikipedia.org/wiki/Cron#History)
[^2]:[crontab(5) - Linux manual page](https://www.man7.org/linux/man-pages/man5/crontab.5.html)
[^3]:[crontab(1) - Linux manual page](https://www.man7.org/linux/man-pages/man1/crontab.1.html)
[^4]:[anacron(8) - Linux manual page](https://www.man7.org/linux/man-pages/man8/anacron.8.html)
[^5]:[anacrontab(5) - Linux manual page](https://www.man7.org/linux/man-pages/man5/anacrontab.5.html)


