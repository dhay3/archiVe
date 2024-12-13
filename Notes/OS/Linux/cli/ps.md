# Linux ps

参考：

https://www.cnblogs.com/peida/archive/2012/12/19/2824418.html

http://c.biancheng.net/view/1062.html

> `ps -lef | more`查看进程
>
> `ps auxf`信息更加详细
>
> `ps -ef | wc -l`查看一共打开了多少进程

## 概述

显示当前进程的快照（process status），CMD等同于COMMAND

## 参数

- -a

  查看所有进程，但是不显示与当前终端无关的进程

- -l

  显示更加详细的信息

- -A

  查看所有进程，等价于`-e`

  ```
  [root@chz html]# ps -e
     PID TTY          TIME CMD
       1 ?        00:00:04 systemd
       2 ?        00:00:00 kthreadd
       4 ?        00:00:00 kworker/0:0H
       6 ?        00:00:00 ksoftirqd/0
       7 ?        00:00:00 migration/0
       8 ?        00:00:00 rcu_bh
  ```

- -U

  根据用户ID查询，等价于`-u`

  ```
  [root@chz html]# ps -u
  USER        PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
  root       1730  0.0  5.9 353024 59008 tty1     Ssl+ 13:13   0:07 /usr/bin/X :0 -background none -noreset -au
  root       2827  0.0  0.3 116996  3132 pts/0    Ss   13:14   0:00 bash
  root       7278  0.0  0.1 155372  1872 pts/0    R+   18:52   0:00 ps -u
  
  ```

- -f 

  全格式输出

  ```
  [root@chz html]# ps -f
  UID         PID   PPID  C STIME TTY          TIME CMD
  root       2827   2820  0 13:14 pts/0    00:00:00 bash
  root       7340   2827  0 18:56 pts/0    00:00:00 ps -f
  ```

## 例子

- `ps -ef`

  查询所有终端的进程，包括内核的

  ```
  [root@chz html]# ps -ef
  UID         PID   PPID  C STIME TTY          TIME CMD
  root          1      0  0 13:13 ?        00:00:04 /usr/lib/systemd/systemd --switched-root --system --deseria
  root          2      0  0 13:13 ?        00:00:00 [kthreadd]
  root          4      2  0 13:13 ?        00:00:00 [kworker/0:0H]
  root          6      2  0 13:13 ?        00:00:00 [ksoftirqd/0]
  root          7      2  0 13:13 ?        00:00:00 [migration/0]
  ```

- `ps -ef1`

  查看指定pid 进程

  ```
  cpl in /usr/local/bin λ ps -ef1
      PID TTY      STAT   TIME COMMAND
        1 ?        Ss     0:03 /sbin/init
  ```

- `ps -af`

  查询当前终端的所有进程

  ```
  [root@chz html]# ps -af
  UID         PID   PPID  C STIME TTY          TIME CMD
  root       7545   2827  0 19:12 pts/0    00:00:00 ps -af
  ```

- `ps aux`

  以BSD格式显示输出内容

  ```
  [root@chz html]# ps aux|more
  USER        PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
  root          1  0.0  0.4 194292  4548 ?        Ss   13:13   0:04 /usr/lib/systemd/systemd --switched-root --
  system --deserialize 22
  root          2  0.0  0.0      0     0 ?        S    13:13   0:00 [kthreadd]
  root          4  0.0  0.0      0     0 ?        S<   13:13   0:00 [kworker/0:0H]
  ```

  | 表头    | 含义                                                         |
  | ------- | ------------------------------------------------------------ |
  | USER    | 该进程是由哪个用户产生的。                                   |
  | PID     | 进程的 ID。                                                  |
  | %CPU    | 该进程占用 CPU 资源的百分比，占用的百分比越高，进程越耗费资源。 |
  | %MEM    | 该进程占用物理内存的百分比，占用的百分比越高，进程越耗费资源。 |
  | VSZ     | 该进程占用虚拟内存的大小，单位为 KB。                        |
  | RSS     | 该进程占用实际物理内存的大小，单位为 KB。                    |
  | TTY     | 该进程是在哪个终端运行的。其中，tty1 ~ tty7 代表本地控制台终端（可以通过 Alt+F1 ~ F7 快捷键切换不同的终端），tty1~tty6 是本地的字符界面终端，tty7 是图形终端。pts/0 ~ 255 代表虚拟终端，一般是远程连接的终端，第一个远程连接占用 pts/0，第二个远程连接占用 pts/1，依次増长。 |
  | STAT    | 进程状态。常见的状态有以下几种：-D：不可被唤醒的睡眠状态，通常用于 I/O 情况。-R：该进程正在运行。-S：该进程处于睡眠状态，可被唤醒。-T：停止状态，可能是在后台暂停或进程处于除错状态。-W：内存交互状态（从 2.6 内核开始无效）。-X：死掉的进程（应该不会出现）。-Z：僵尸进程。进程已经中止，但是部分程序还在内存当中。-<：高优先级（以下状态在 BSD 格式中出现）。-N：低优先级。-L：被锁入内存。-s：包含子进程。-l：多线程（小写 L）。-+：位于后台。 |
  | START   | 该进程的启动时间。                                           |
  | TIME    | 该进程占用 CPU 的运算时间，注意不是系统时间。                |
  | COMMAND | 产生此进程的命令名。                                         |

- `ps -le`

  ```
  root in ~ λ ps -le | more
  F S   UID     PID    PPID  C PRI  NI ADDR SZ WCHAN  TTY          TIME CMD
  4 S     0       1       0  0  80   0 - 41800 -      ?        00:00:03 systemd
  1 S     0       2       0  0  80   0 -     0 -      ?        00:00:00 kthreadd
  1 I     0       3       2  0  60 -20 -     0 -      ?        00:00:00 rcu_gp
  1 I     0       4       2  0  60 -20 -     0 -      ?        00:00:00 rcu_par_gp
  ```

  | 表头  | 含义                                                         |
  | ----- | ------------------------------------------------------------ |
  | F     | 进程标志，说明进程的权限，常见的标志有两个: 1：进程可以被复制，但是不能被执行；4：进程使用超级用户权限； |
  | S     | 进程状态。具体的状态和"psaux"命令中的 STAT 状态一致；        |
  | UID   | 运行此进程的用户的 ID；                                      |
  | PID   | 进程的 ID；                                                  |
  | PPID  | 父进程的 ID；                                                |
  | C     | 该进程的 CPU 使用率，单位是百分比；                          |
  | PRI   | 进程的优先级，数值越小，该进程的优先级越高，越早被 CPU 执行； |
  | NI    | 进程的优先级，数值越小，该进程越早被执行；                   |
  | ADDR  | 该进程在内存的哪个位置；                                     |
  | SZ    | 该进程占用多大内存；                                         |
  | WCHAN | 该进程是否运行。"-"代表正在运行；                            |
  | TTY   | 该进程由哪个终端产生；                                       |
  | TIME  | 该进程占用 CPU 的运算时间，注意不是系统时间；                |
  | CMD   | 产生此进程的命令名；                                         |
