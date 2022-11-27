# linux 日志系统

参考：

https://www.cnblogs.com/yingsong/p/6022181.html

linux日志系统一般由syslog或systemd-jourald.service来管理。不同的distro日志生成的方式也不同

## wtmp

`/var/log/wtmp`是一个二进制文件记录了登入的详细信息，不能直接查看，所以我们需要使用`last`命令来查看。

## btmp

`/var/log/btmp`是一个二进制文件记录了登入失败的详细信息，不能直接查看，需要通过`lastb`来查看

## utmp

`/var/run/utmp`，二进制文件记录当前登入的用户，可以使用`users`或是`w`查看

```
root in /var/log λ w
 13:35:57 up 16:49,  1 user,  load average: 0.00, 0.01, 0.00
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
root     pts/0    115.233.222.34   13:29    1.00s  0.20s  0.00s w
root in /var/log λ users
root
```

## auth

`/var/log/auth*`，用户认证产生的信息，如login，su

```
Feb 28 12:40:10 cyberpelican CRON[7032]: pam_unix(cron:session): session opened for user root by (uid=0)
Feb 28 12:40:10 cyberpelican CRON[7031]: pam_unix(cron:session): session opened for user root by (uid=0)
Feb 28 12:40:10 cyberpelican CRON[7031]: pam_unix(cron:session): session closed for user root
Feb 28 12:40:12 cyberpelican CRON[7032]: pam_unix(cron:session): session closed for user root
Feb 28 12:41:01 cyberpelican CRON[7046]: pam_unix(cron:session): session opened for user root by (uid=0)
```

## lastlog

`/var/log/lastlog`，记录所有的登入信息，包括系统级别的账号，使用`lastlog`命令来查看

## journal

`/var/log/journal/`，systemd-jourald.service存储日志的地方。

删除`rm -rf /var/log/journal`

## cron

`/var/log/cron*`，系统crontab产生的日志

## secure

`/var/log/secure*`，常见的系统和服务错误信息

## syslog

`/var/log/syslog*`，syslog的日志

## boot

`/var/log/boot*`，系统的引导日志

## 清除日志

使用`cat /dev/null >| log`来清除日志，这里需要使用循环。重定向的文件不支持模式扩展

```
[root@chz log]# for i in $(ls | grep secure);do cat /dev/null >| "$i";done
```

其实也可以只对secure文件进行清除，因为按照linux的自动压缩规则，没有日期的都是最新的日志文件

```
[root@chz log]# cat /dev/null >| secure
```

删除所有敏感日志

```
rm -rf /var/log/journal
cat /dev/null >| /var/log/secure
cat /dev/null >| /var/log/cron
cat /dev/null >| /var/log/wtmp
cat /dev/null >| /var/log/lastlog
cat /dev/null >| /var/log/auth
cat /dev/null >| /var/log/secure
```

