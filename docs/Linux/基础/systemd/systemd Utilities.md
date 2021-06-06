# systemd Utilities

### systemctl 

systemctl 用于管理系统电源，以及Unit（服务）

#### 系统电源

> reboot，poweroff ，halt，suspend 无需添加systemctl，具体查看manual

```
systemctl reboot #重启
systemctl poweroff #关闭系统，切断电源 等价于 shutdown -h 0
systemctl halt #cpu停止工作
systemctl suspend #暂停系统
systemctl hibernate #进入冬眠状态
systemctl hybrid-sleep # 等价于 suspend 和 hibernate
```

### analyze

- `systemd-analyze`

  查看启动(boot-up)耗时

  ```
  [root@cyberpelican init.d]# systemd-analyze 
  Startup finished in 677ms (kernel) + 2.579s (initrd) + 47.339s (userspace) = 50.597s
  ```

- `systemd-analyze critical-chain httpd`

  查看某个unit启动的过程和时间

  ```
  [root@cyberpelican ~]# systemd-analyze critical-chain httpd.service
  The time after the unit is active or started is printed after the "@" character.
  The time the unit takes to start is printed after the "+" character.
  
  httpd.service +25.646s
  └─remote-fs.target @8.921s
    └─remote-fs-pre.target @8.919s
      └─iscsi-shutdown.service @8.855s +57ms
        └─network.target @8.835s
          └─wpa_supplicant.service @13.631s +67ms
            └─basic.target @4.799s
              └─sockets.target @4.798s
                └─dbus.socket @4.798s
                  └─sysinit.target @4.784s
                    └─sys-fs-fuse-connections.mount @30.645s +24ms
                      └─system.slice
                        └─-.slice
  ```

### loginctl

- `loginctl`

  显示当前登入的用户，等价于`loginctl list-sessions`，==一个账号能有多个session(tty)==

  ```
  [root@cyberpelican systemd]# loginctl 
     SESSION        UID USER             SEAT            
          c1         42 gdm              seat0           
           1          0 root             seat0           
          22          0 root                             
  
  3 sessions listed.
  ```

- `loginctl lock-session <session id>`

  lock指定session ID，==对非图形化界面不生效==

  ```
  [root@cyberpelican systemd]# loginctl lock-session 1
  ```

- `loginctl kill-session <session id>`

  直接关闭session，等价于`terminate-sessiono`，回到登入界面

  ```
  [root@cyberpelican systemd]# loginctl kill-session 22
  ```

<img src="D:\asset\note\imgs\_Linux\Snipaste_2020-10-29_22-10-23.png"/>

- `loginctl list-users`

  列出当前登入的账号

  ```
  [root@cyberpelican systemd]# loginctl list-users
         UID USER            
          42 gdm             
           0 root            
  
  2 users listed.
  ```

- `loginctl show-user <logined user>`

  显示具体用户的信息，对比`id`

  ```
  [root@cyberpelican systemd]# loginctl show-user root
  UID=0
  GID=0
  Name=root
  Timestamp=Thu 2020-10-29 19:16:46 CST
  TimestampMonotonic=30183376
  RuntimePath=/run/user/0
  Slice=user-0.slice
  Display=1
  State=active
  Sessions=1
  IdleHint=no
  IdleSinceHint=0
  IdleSinceHintMonotonic=0
  Linger=no
  [root@cyberpelican systemd]# id root
  uid=0(root) gid=0(root) groups=0(root)
  ```

- `loginctl kill-user <user>`

  关闭指定user的所有session，等价于`terminate-user`，回到登入界面

  ```
  [root@cyberpelican ~]# loginctl kill-user root
  ```

### journalctl

参考:

https://www.cnblogs.com/sparkdev/p/8795141.html

> systemd-journald.service默认将日志保存在`/var/log/journal`下，用户没有权限修改日志，重启系统后日志丢失。

- `journalctl`

  不带任何参数显示启动系统后的所有日志

  ```
  [root@cyberpelican ~]# journalctl 
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Thu 2020-10-29 19:24:08 CST. --
  Oct 29 19:07:10 chz systemd-journal[91]: Runtime journal is using 6.0M (max allowed 48.6M, trying to lea
  Oct 29 19:07:10 chz kernel: Initializing cgroup subsys cpuset
  Oct 29 19:07:10 chz kernel: Initializing cgroup subsys cpu
  Oct 29 19:07:10 chz kernel: Initializing cgroup subsys cpuacct
  Oct 29 19:07:10 chz kernel: Linux version 3.10.0-1062.el7.x86_64 (mockbuild@kbuilder.bsys.centos.org) (g
  Oct 29 19:07:10 chz kernel: Command line: BOOT_IMAGE=/vmlinuz-3.10.0-1062.el7.x86_64 root=/dev/mapper/ce
  Oct 29 19:07:10 chz kernel: Disabled fast string operations
  ```

- `journalctl -f`

  实时滚动显示日志

- `journalctl -e`

  查看日志并使用G

  ```
  Oct 29 20:09:02 cyberpelican systemd[1]: Started Network Manager Script Dispatcher Service.
  Oct 29 20:09:02 cyberpelican nm-dispatcher[3671]: req:1 'dhcp4-change' [ens33]: new request (4 scripts)
  Oct 29 20:09:02 cyberpelican nm-dispatcher[3671]: req:1 'dhcp4-change' [ens33]: start running ordered sc
  Oct 29 20:10:01 cyberpelican systemd[1]: Started Session 8 of user root.
  Oct 29 20:10:01 cyberpelican CROND[3701]: (root) CMD (/usr/lib64/sa/sa1 1 1)
  lines 1001-1038/1038 (END)
  ```

- `journalctl -x`

  当存在错误信息时，使用`-x`参数可以给出简单的帮助信息，可以配合`-e`参数一起使用

  ```
  Oct 29 20:27:01 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_start.timer:12] Failed to pars
  Oct 29 20:27:01 cyberpelican systemd[1]: mdcheck_start.timer lacks value setting. Refusing.
  Oct 29 20:27:01 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_start.service:14] Invalid envi
  Oct 29 20:27:01 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_continue.service:14] Invalid e
  Oct 29 20:27:03 cyberpelican polkitd[1118]: Registered Authentication Agent for unix-process:4091:424719
  Oct 29 20:27:03 cyberpelican polkitd[1118]: Unregistered Authentication Agent for unix-process:4091:4247
  lines 2774-2811/2811 (END)
  ```

- `journalctl _PID=<pid>`

  查看指定pid进程日志，同样的也可以使用`_UID`来查看某一个用户的日志

  ```shell
  [root@cyberpelican system]# journalctl _PID=1593
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Fri 2020-10-30 09:04:15 CST. --
  Oct 29 19:16:42 cyberpelican httpd[1593]: AH00558: httpd: Could not reliably determin
  
  ---
  
  [root@cyberpelican system]# journalctl _UID=0
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Fri 2020-10-30 09:04:15 CST. --
  Oct 29 19:07:10 chz systemd-journal[91]: Runtime journal is using 6.0M (max allowed 4
  Oct 29 19:07:10 chz systemd-journal[91]: Journal started
  Oct 29 19:07:10 chz systemd[1]: Started dracut cmdline hook.
  Oct 29 19:07:10 chz systemd[1]: Starting dracut pre-udev hook...
  Oct 29 19:07:10 chz systemd[1]: Started dracut pre-udev hook.
  Oct 29 19:07:10 chz systemd[1]: Starting udev Kernel Device Manager...
  Oct 29 19:07:10 chz systemd-udevd[242]: starting version 219
  
  ```

- ==`journalctl -u <unit>`==

  指定显示某一个unit，使用`--since`和`--until`参数查看指定时间段内的日志

  ```
  [root@cyberpelican ~]# journalctl -u httpd.service  --since "2020-10-29"
  -- Logs begin at Thu 2020-10-29 19:07:10 CST, end at Thu 2020-10-29 20:27:03 CST. --
  Oct 29 19:16:27 cyberpelican systemd[1]: Starting The Apache HTTP Server...
  Oct 29 19:16:42 cyberpelican httpd[1593]: AH00558: httpd: Could not reliably determine the server's full
  Oct 29 19:16:53 cyberpelican systemd[1]: Started The Apache HTTP Server.
  ```

- `journactl --since | --until`

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

- `journalctl -o`

  指定output输出的内容格式

  ```
  [root@cyberpelican systemd]# journalctl -o json-pretty
  {
          "__CURSOR" : "s=80e1f974ad0643998becb58e51847d4a;i=1;b=d466e1cc725d4316875231666fd7de59;m=ad297;t=5b2cd480f
          "__REALTIME_TIMESTAMP" : "1603969630547679",
          "__MONOTONIC_TIMESTAMP" : "709271",
          "_BOOT_ID" : "d466e1cc725d4316875231666fd7de59",
          "PRIORITY" : "6",
          "_TRANSPORT" : "driver",
          "MESSAGE" : "Runtime journal is using 6.0M (max allowed 48.6M, trying to leave 72.9M free of 480.1M availab
          "MESSAGE_ID" : "ec387f577b844b8fa948f33cad9a75e6",
          "_PID" : "91",
          "_UID" : "0",
          "_GID" : "0",
          "_COMM" : "systemd-journal",
          "_EXE" : "/usr/lib/systemd/systemd-journald",
          "_CMDLINE" : "/usr/lib/systemd/systemd-journald",
          "_CAP_EFFECTIVE" : "5402800cf",
          "_SYSTEMD_CGROUP" : "/system.slice/systemd-journald.service",
          "_SYSTEMD_UNIT" : "systemd-journald.service",
          "_SYSTEMD_SLICE" : "system.slice",
          "_MACHINE_ID" : "fbb74b6620184684961580de92e236c2",
          "_HOSTNAME" : "chz"
  }
  
  ```

- `journalctl -p`

  指定输出日志的等级，显示当前等级及以下的内容

  ```
  [root@cyberpelican ~]# journalctl -e -p 3
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_continue.service:14] Invalid e
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_continue.service:14] Invalid e
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_start.timer:12] Failed to pars
  Oct 29 20:26:58 cyberpelican systemd[1]: mdcheck_start.timer lacks value setting. Refusing.
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_start.service:14] Invalid envi
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_continue.service:14] Invalid e
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_continue.service:14] Invalid e
  Oct 29 20:26:58 cyberpelican systemd[1]: [/usr/lib/systemd/system/mdcheck_start.service:14] Invalid envi
  ```

- `journalctl --disk-usage`

  查看当前主机存储日志的大小

  ```
  [root@cyberpelican ~]# journalctl --disk-usage 
  Archived and active journals take up 6.0M on disk.
  ```

- `journalctl --vacuum-size | vacuum-time`

  删除指定大小或是指定时间内的日志

  ```
  [root@cyberpelican systemd]# journalctl --vacuum-size=1G
  Vacuuming done, freed 0B of archived journals on disk.
  
  
  #在 1 hour之前的内容删除
  [root@cyberpelican systemd]# journalctl --vacuum-time=1h
  Vacuuming done, freed 0B of archived journals on disk.=
  ```

#### 配置文件

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
#SplitMode=uid
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

### hostnamectl

查看当前主机的信息

- `hostnamectl`

  等价于`hostnamectl status`，对比`uname -r`和`cat /proc/version`

  ```
  [root@cyberpelican systemd]# hostnamectl 
     Static hostname: cyberpelican
           Icon name: computer-vm
             Chassis: vm
          Machine ID: fbb74b6620184684961580de92e236c2
             Boot ID: d466e1cc725d4316875231666fd7de59
      Virtualization: vmware
    Operating System: CentOS Linux 7 (Core)
         CPE OS Name: cpe:/o:centos:centos:7
              Kernel: Linux 3.10.0-1062.el7.x86_64
        Architecture: x86-64
        
  [root@cyberpelican systemd]# cat /proc/version 
  Linux version 3.10.0-1062.el7.x86_64 (mockbuild@kbuilder.bsys.centos.org) (gcc version 4.8.5 20150623 (Red Hat 4.8.5-36) (GCC) ) #1 SMP Wed Aug 7 18:08:02 UTC 2019
  [root@cyberpelican systemd]# uname -r
  3.10.0-1062.el7.x86_64
  
  
  ```

- `hostnamectl set-hostname `

  设置主机名

  ```
  [root@cyberpelican systemd]# hostnamectl set-hostname  cyberpelican
  ```

### localectl

查看和设置本地化(语言)

- `localectl`

  等价于`localectl status`，查看当前系统的本地化设置

  ```
  [root@chz systemd]# localectl 
     System Locale: LANG=en_US.UTF-8
         VC Keymap: us
        X11 Layout: us
  ```

- `localectl set-locale | set-keymap`

  设置显示语言或输入语言，结合`--list-locales`和`--list-keymap`一起使用

  ```
  [root@cyberpelican systemd]# localectl set-locale LANG=zh_CN.UTF-8
  [root@cyberpelican systemd]# localectl
     System Locale: LANG=zh_CN.UTF-8
         VC Keymap: us
        X11 Layout: us
        
  [root@cyberpelican systemd]# localectl set-keymap zh_CN
  [root@cyberpelican systemd]# localectl
     System Locale: LANG=zh_CN.UTF-8
         VC Keymap: zh_CN
        X11 Layout: us
  ```

### timedatectl

与`locatectl`相似，具体查看manual

- `timedatectl`

  等价于`timedatectl status`，区别于`date`

  ```
  [root@cyberpelican systemd]# timedatectl 
        Local time: Thu 2020-10-29 21:34:37 CST
    Universal time: Thu 2020-10-29 13:34:37 UTC
          RTC time: Thu 2020-10-29 13:25:31
         Time zone: Asia/Shanghai (CST, +0800)
       NTP enabled: no
  NTP synchronized: no
   RTC in local TZ: no
        DST active: n/a
  [root@cyberpelican systemd]# date
  Thu Oct 29 21:34:41 CST 2020
  ```

