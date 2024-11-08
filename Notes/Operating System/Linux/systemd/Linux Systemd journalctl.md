# Linux Systemd journalctl

ref:

https://www.cnblogs.com/sparkdev/p/8795141.html

## digest

用于查询 systemd-journald.service 记录的日志，如果没有指定任何参数会显示最早的记录。默认输出less模式，内容较长的会被truncated

systemd-journald.service默认将日志保存在`/var/log/journal`下，用户没有权限修改日志，重启系统后日志丢失。

## optional args

### output

- `-f | --follow`

  以流方式显示最新的日志，会持续更新

- `-e | --pager-end`

  ==直接到最新的日志==

- `-r | --reverse`

  等价与`-e`

- `-n | --lines=`

  指定显示行数，一般与`-f`一起使用

- `-o | --output=`

  以指定格式输出，默认以`short`格式输出，具体查看man page，常用short-iso、verbose、json、json-pretty

- `--no-hostname`

  不输出 hostname field

  ```
  May 29 19:18:45 kernel: BIOS-e820: [mem 0x00000003ae340000-0x000000042fffffff] reserved
  
  ```

- `-x | --catalog`

  ==为错误日志提供帮助==

- `-b [_BOOT_ID] | --boot=[_BOOT_ID]`

  查看指定boot阶段的日志，1 表示boot 1阶段，2表示boot 2阶段以此类推，0表示最后。具体的ID 范围可以使用`--list-boots`查看，详细查看 journal-fields

- `-k | --dmesg`

  等价使用`dmesg`查看kernal 日志

- `-u=UNIT|PATTERN`

  指定显示 service unit，或者匹配Pattern的。如果是一个 unit systemd.slice 会显示所有的子unit。通常和`--since`和`--until`参数查看指定时间段内的日志

  ```
  [root@cyberpelican ~]# journalctl -u httpd.service  --since "2020-10-29"
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Thu 2020-10-29 20:27:03 CST. --
  Oct 29 19:16:27 cyberpelican systemd[1]: Starting The Apache HTTP Server...
  Oct 29 19:16:42 cyberpelican httpd[1593]: AH00558: httpd: Could not reliably determine the server's full
  Oct 29 19:16:53 cyberpelican systemd[1]: Started The Apache HTTP Server.
  ```

- `-p | --priority=`

  按照指定优先级显示日志，可以使用 text(emerg..debug) 或者 number(1..7)，具体等级查看 `syslog` man page

  ```
  cpl in ~/note/docs/Linux/basic/systemd on master λ journalctl -rkp 4
  -- Journal begins at Sat 2021-05-29 19:18:45 HKT, ends at Mon 2022-03-14 21:58:21 HKT. --
  Mar 14 20:52:11 cyberpelican kernel: kauditd_printk_skb: 16 callbacks suppressed
  Mar 14 20:51:52 cyberpelican kernel: kauditd_printk_skb: 204 callbacks suppressed
  Mar 14 20:51:51 cyberpelican kernel: amdgpu: SRAT table not found
  Mar 14 20:51:50 cyberpelican kernel: ucsi_acpi: probe of USBC000:00 failed with error -5
  Mar 14 20:51:50 cyberpelican kernel: amdgpu 0000:03:00.0: amdgpu: PSP runtime database doesn't exist
  Mar 14 20:51:49 cyberpelican kernel: clocksource: Checking clocksource tsc synchronization from CPU 0 to CPUs 1-2,6,8,12-13.
  Mar 14 20:51:49 cyberpelican kernel: TSC found unstable after boot, most likely due to broken BIOS. Use 'tsc=unstable'.
  Mar 14 20:51:49 cyberpelican kernel: clocksource:                       'tsc' is current clocksource.
  Mar 14 20:51:49 cyberpelican kernel: clocksource:                       'tsc' cs_nsec: 506652289 cs_now: 8a6f3da80 cs_last: >
  Mar 14 20:51:49 cyberpelican kernel: clocksource:                       'hpet' wd_nsec: 510550572 wd_now: 1baa648 wd_last: 1>
  Mar 14 20:51:49 cyberpelican kernel: clocksource: timekeeping watchdog on CPU3: Marking clocksource 'tsc' as unstable becaus>
  Mar 14 20:51:49 cyberpelican kernel: thermal thermal_zone0: failed to read out thermal zone (-61)
  ```

- `-g | --grep=`

  使用pcre2pattern过滤message部分内容

- `--no-pager`

  直接输出所有内容到stdout

### time

- ` -S | --since , -U | --until`

  指定日志的开始时间或结束时间

  ```
  root in ~ λ journalctl --since="2021-01-06"
  -- Logs begin at Fri 2020-11-20 18:08:39 HKT, end at Thu 2021-01-07 09:17:01 HKT. --
  Jan 06 00:05:01 ubuntu18.04 CRON[6290]: pam_unix(cron:session): session opened for user root by (uid=0)
  Jan 06 00:05:01 ubuntu18.04 CRON[6291]: (root) CMD (command -v debian-sa1 > /dev/null && debian-sa1 1 1)
  Jan 06 00:05:01 ubuntu18.04 CRON[6290]: pam_unix(cron:session): session closed for user root
  Jan 06 00:15:01 ubuntu18.04 CRON[6304]: pam_unix(cron:session): session opened for user root by (uid=0)
  Jan 06 00:15:01 ubuntu18.04 CRON[6305]: (root) CMD (command -v debian-sa1 > /dev/null && debian-sa1 1 1)
  Jan 06 00:15:01 ubuntu18.04 CRON[6304]: pam_unix(cron:session): session closed for user root
  ```

### journal managment

- `-N | --fields`

  查看当前使用的fields

- `-F | --field=`

  只打印指定field内容

- `--disk-usage`

  查看journal日志占用系统的大小

- `--vacuum-size, --vacuum-time, --vacuum-files`

  删除指定大小或是指定时间内的日志

  ```
  [root@cyberpelican systemd]# journalctl --vacuum-size=1G
  Vacuuming done, freed 0B of archived journals on disk.
  
  
  #在 1 hour之前的内容删除
  [root@cyberpelican systemd]# journalctl --vacuum-time=1h
  Vacuuming done, freed 0B of archived journals on disk.=
  ```

- `--sync`

   Asks the journal daemon to write all yet unwritten journal data to the backing file system and synchronize all journals

- `--flush`

  将journal中的当前存储的日志(/run/log/journal)同步到 /var/log/journal

## journal-fields

以`FIELD=VALUE`的格式，用于过滤journal

```
#过滤指定的unit service
journalctl _SYSTEMD_UNIT=avahi-daemon.service

#过滤unit service的同时按照PID过滤
journalctl _SYSTEMD_UNIT=avahi-daemon.service _PID=28097
```

如果使用了`+`表示逻辑或

```
journalctl _SYSTEMD_UNIT=avahi-daemon.service _PID=28097 + _SYSTEMD_UNIT=dbus.service
```

- `-PID=,_UID=,_GID=`

  顾名思义

  ```
  [root@cyberpelican system]# journalctl _PID=1593
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Fri 2020-10-30 09:04:15 CST. --
  Oct 29 19:16:42 cyberpelican httpd[1593]: AH00558: httpd: Could not reliably determin
  
  ---
  
  [root@cyberpelican system]# journalctl _UID=0
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Fri 2020-10-30 09:04:15 CST. --
  ```

- `_BOOT_ID=`

  the kernel boot id （have no idea about this）

- `_HOSTNAME=`

  the name of the originatting host

- `_TRANSPORT=`

  journal收到日志的来源

  1. audit
  2. driver
  3. syslog
  4. journal
  5. stdout
  6. kernel

## conf

具体查看man page

```
[root@cyberpelican systemd]# cat journald.conf 
#  This file is part of systemd.
#
#  systemd is free software; you can redistribute it and/or modify it
#  under the terms of the GNU Lesser General Public License as published by
#  the Free Software Foundation; either version 2.1 of the License, or
#  (at your option) any later version.
#
# Entries in this file show the compile time defaults.
# You can change settings by editing this file.
# Defaults can be restored by simply deleting this file.
#
# See journald.conf(5) for details.

[Journal]
#Storage=auto  #默认使用auto，volatile存储在内存中，persistence存储在物理硬盘中
#Compress=yes #是否压缩日志
#Seal=yes
#SplitMode=uid #按照uid分割日志
#SyncIntervalSec=5m
#RateLimitInterval=30s
#RateLimitBurst=1000
#SystemMaxUse= #指定journal能使用的最高持久存储容量
#SystemKeepFree= #指定journal在添加新条目时需要保留的剩余空间
#SystemMaxFileSize=	#控制单一journal文件大小，符合条件方可被转为持久存储
#RuntimeMaxUse= #指定易失性存储纵的最大可用磁盘容量
#RuntimeKeepFree= #指定向易失性存储内写入数据时为其它应用保留的空间量
#RuntimeMaxFileSize= #指定单一journal文件可占用的最大易失性存储容量
#MaxRetentionSec=
#MaxFileSec=1month
#ForwardToSyslog=yes
#ForwardToKMsg=no
#ForwardToConsole=no
#ForwardToWall=yes
#TTYPath=/dev/console
#MaxLevelStore=debug
#MaxLevelSyslog=debug
#MaxLevelKMsg=notice
#MaxLevelConsole=info
#MaxLevelWall=emerg
#LineMax=48K
```
