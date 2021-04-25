# Linux top

top 流式展示当前linux的进程信息，是一个交互式的命令

## 常用参数

- `-d`

  指定top一次更新的时间间隔

  ```
  #10s更新一次
  root in /tmp λ top -d 10
  ```

- `-b`

  以batch-mode运行top，通常与`-n`(指定top需要等待更新多少次后终止)一起使用。将top中的内容输出到文件

  ```
  root in /tmp λ top -b -n 1 > top.dump
  root in /tmp λ head -10 top.dump
  top - 14:00:44 up 2 days, 20:04,  2 users,  load average: 0.03, 0.05, 0.02
  Tasks: 115 total,   1 running,  77 sleeping,   0 stopped,   0 zombie
  %Cpu(s):  0.5 us,  0.4 sy,  0.0 ni, 99.0 id,  0.1 wa,  0.0 hi,  0.0 si,  0.0 st
  KiB Mem :  1877052 total,   147324 free,   871488 used,   858240 buff/cache
  KiB Swap:        0 total,        0 free,        0 used.   850268 avail Mem
  
    PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
      1 root      20   0  160232   6980   4280 S  0.0  0.4   0:04.22 systemd
      2 root      20   0       0      0      0 S  0.0  0.0   0:00.13 kthreadd
      4 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 kworker/0:0H
  ```

- `-o <field-names>`

  按照指定列排序，在列名前添加`+`表示从高到低，`-`表示从低到高

  ```
  #按照内存使用率从高到低排序
  root in /tmp λ top -o +%MEM
  ```

- `-p <pid>`

  监控指定pid的进程

  ```
  root in /tmp λ top -p 1 -p 2
  ```

- `-u <uid>`

  展示指定uid调用的进程，可以是名字也可以是数字

  ```
  root in /tmp λ top -u postfix
  ```

## 展示界面

### header

```
top - 14:15:57 up 2 days, 20:20,  2 users,  load average: 0.00, 0.00, 0.00
Tasks: 115 total,   1 running,  77 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.7 sy,  0.0 ni, 99.3 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :  1877052 total,   143992 free,   873124 used,   859936 buff/cache
KiB Swap:        0 total,        0 free,        0 used.   848624 avail Mem
```

1. 主机从上一次关机到现在的时间，当前登入系统的用户数

2. 当前运行的进程或线程数，状态有四种：runngin，sleeping，stopped，zombie
3. CPU使用的情况，具体查看`man top /^\s+2b\.`(使用`t`可以以不同的形式显示CPU信息)
4. 内存使用的情况，使用`m`可以以不同的形式显示内存信息，==使用`E`以不同单位显示内存使用情况==

### body

> 可以通过`man /fields\/columns`查看

top默认显示如下几个字段，==但是可以通过`f`来自定义显示的列==

```
 PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
```

- pid：进程的pid
- user：调用进程的user
- pr：priority 进程调度未来的优先级，值越小优先级越高，rt表示进程的优先级动态变化
- ni：nice value进程调度的优先级，值越小优先级越高
- virt：virtual memory size，进程使用的虚拟内存
- res：resident memory size，non-swapped phsical memory
- shr：shared memory size，和其他进程共享的内存
- S：进程的状态，D=uniterruptibele sleep，R=running，S=sleeping，T=stopped by job control signal，t=stopped by debugger during trace，Z=zombie
- time：从进程启动后占用cpu的时间

## Interactive commands

### global command

==`h`显示帮助信息==

- d

  修改delay time(刷新时间)

- E | e

  E以不同单位显示header中内存内容，e以不同单位显示body中内存内容

- H

  不以Task为单位显示，而是以Threads

- k

  终止一个进程

- r

  renice-a-task，为task重新分配优先值

- W

  ==持久化当前在top中的修改的设置，默认存储在`~/.toprc`中，下次使用top是自动调用该配置文件==

- Z

  ==修改top展示信息的颜色==

- X

  设置列宽，`-1`表示自动scale

- Y

  在`.toprc`中输入如下内容，==可以查看指定进程打开的文件，日志，NUMA==，同样还是从syslog中读取日志

  ```
   /bin/echo -e "pipe\tOpen Files\tlsof -P -p %d 2>&1" >> ~/.toprc
   /bin/echo -e "file\tNUMA Info\t/proc/%d/numa_maps" >> ~/.toprc
   /bin/echo -e "pipe\tLog\ttail -n200 /var/log/syslog | sort -Mr" >> ~/.toprc
  ```

### header command

- t

  以不同形式显示header 部分cpu信息

- m

  以不同形式显示header 部分mem信息

- 1

  %CPU(S)表示显示核CPU使用的情况，使用`1`显示每一个核CPU使用的情况，结合`lscpu`命令来查看

### body commands

- j

  信息以中间对齐的方式显示

- c

  切换command列显示的内容

- u

  显示指定user打开的进程

- i

  只显示当前正在运行的task

- f

  显示指定列

- M

  ==按照%MEM排序，从大到小==

- N

  ==按照%PID排序，从大到小==

- P

  ==按照%CPU排序，从大到小==

- L

  ==locate-a-string，搜索指定关键字，大小写敏感==，`&`表示下一个

