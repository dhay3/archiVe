# schtasks.exe

`schtasks`是windows上的计划任务管理器

## /create

创建定计划任务

- /tn

  指定计划任务的名字

- /tr

  指定计划任务调用的程序路径

- /st

  指定计划任务开始的时间，如果没有指定默认从现在开始。如果`/sc once`必须要有该选项

- /et

  指定计划任务的结束时间

- /sc 

  指定计频率，有效值为MINUTE，HOURLY，DAILY，WEEKLY，MONTHLY，ONCE，ONSTART，ONLOGON，ONIDLE，ONEVENT

创建mysql定时备份计划，如果不是系统自带的程序调用时需要带上双引号

```
@echo off
set uname="root"
set passwd="12345"
set databases="db2019"
set format="%date:~,4%%date:~5,2%%date:~8,2%"
set bakPath="f:\mysqldump\"
"D:\mysql\mysql-8.0.17-winx64\bin\mysqldump.exe"  -P 3306 --user=%uname% --password=%passwd%   %databases% >  "%bakPath%%format%.sql"

PS C:\Users\82341> schtasks.exe /create /tn mysql-bak /tr "C:\Users\82341\mysqldump.bat" /sc daily /st 00:14
```

