---
createTime: 2024-10-30 16:09
tags:
  - "#hash1"
  - "#hash2"
---

# Cron 01 - Cron Overview

## 0x01 Preface

> [!important]
> 本系列会以 Modern Version cron 作为基础来介绍 cron

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

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [cron - Wikipedia](https://en.wikipedia.org/wiki/Cron)


***FootNotes***

[^1]:[cron - Wikipedia](https://en.wikipedia.org/wiki/Cron#History)
