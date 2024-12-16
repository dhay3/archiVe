# top

## Digest

以 flow 的方式展示 system summary information (包含 process 和 thread 的使用率)，默认以interactive mode 运行

top 由 3 部分组成：Summary Area，Fields/Columns Header，Task Area

### Linux memory

Linux 中有 3 种 memory，一种是 optional，

- physical memory，a limited resource where code and data must reside when executed or referenced

- optional swap file，where modified (dirty) memory can be saved and later retrieved if too many demands are made on

- virtual memory，a nearly unlimited resource serving the following goals：

  ```
  1. abstraction, free from physical memory addresses/limits
  2. isolation, every process in a separate address space
  3. sharing, a single mapping can serve multiple needs
  4. flexibility, assign a virtual address to a file
  ```

每种内存都严格在如下 4 像限使用

> Both physical memory and virtual memory can include any of the four, while the swap file only includes #1 through #3.  The memory in quadrant #4, when modified, acts as its own dedicated swap file.

```
                                     Private | Shared
                                 1           |          2
            Anonymous  . stack               |
                       . malloc()            |
                       . brk()/sbrk()        | . POSIX shm*
                       . mmap(PRIVATE, ANON) | . mmap(SHARED, ANON)
                      -----------------------+----------------------
                       . mmap(PRIVATE, fd)   | . mmap(SHARED, fd)
          File-backed  . pgms/shared libs    |
                                 3           |          4
```

- RES  - anything occupying physical memory which, beginning with Linux-4.5, is the sum of the following three fields:
  RSan - quadrant 1 pages, which include any former quadrant 3 pages if modified
  RSfd - quadrant 3 and quadrant 4 pages
  RSsh - quadrant 2 pages

  使用的实际物理内存

- SHR  - subset of RES (excludes 1, includes all 2 & 4, some 3)

  共享内存

- VIRT - everything in-use and/or reserved (all quadrants)

  虚拟内存

- SWAP - potentially any quadrant except 4

### Summary Area

```
top - 18:13:37 up 1 day, 12:35,  1 user,  load average: 0.00, 0.01, 0.05
Tasks:  81 total,   1 running,  79 sleeping,   1 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni,100.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :   498684 total,    73004 free,    90840 used,   334840 buff/cache
KiB Swap:  2097148 total,  2097148 free,        0 used.   390192 avail Mem 
```

#### uptime and load averages

`top - 18:15:11 up 1 day, 12:36,  1 user,  load average: 0.00, 0.01, 0.05`

load average的时间间隔分别是从1,5,15 mins 计算

#### task and cpu states

`Tasks:  81 total,   1 running,  79 sleeping,   1 stopped,   0 zombie`

任务状态计数，包含4中：running；sleeping；stopped；zombie

`%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni,100.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st`

cpu状态，已百分比显示

```
us, user    : time running un-niced user processes
sy, system  : time running kernel processes
ni, nice    : time running niced user processes
id, idle    : time spent in the kernel idle handler
wa, IO-wait : time waiting for I/O completion
hi : time spent servicing hardware interrupts
si : time spent servicing software interrupts
st : time stolen from this vm by the hypervisor
```

如果是多选的cpu显示模式（interactive mode 使用 t 切换）

```
                      a    b     c    d
           %Cpu(s):  75.0/25.0  100[ ...
```

a) is the usesr (us + ni) percentage

b) is the system(sy + hi + si) percentage

c) is the total

d) 图表

#### memory usage

```
GiB Mem :    1.790 total,    0.107 free,    0.381 used,    1.302 buff/cache
GiB Swap:    0.001 total,    0.000 free,    0.001 used.    1.240 avail Mem
```

第一行显示physical memory，第二行显示virtual memory

vitrual momory avail is an estimation of physical memory（包括可回收的内存页），而 free 不包含。

单位默认以kibibytes显示，在interactive mode 中可以使用 `E` 来调整单位

如果以多选的memory模式显示，interactive mode使用`m`来替换

```
                      a    b          c
           GiB Mem : 18.7/15.738   [ ...
           GiB Swap:  0.0/7.999    [ ...
```

a) is the percentage used

b) is the total available，not percentage

c) 图表

### Fields/Columns Header

部分不常用的参数不记录在内，其他具体查看man page

==和内存相关(VIRT，RES，SHR，CODE..)的默认以`kiB`显示（无单位），使用`e`来改变，从KiB 到 PiB==

interactive mode 使用`f`(fileds management)来设置显示的fields

进入field management后使用

`d`或空格添加删除显示的feild

`a`切换预支方案

`s`指定排序的主键

#### RES

resident memory size

a subset of virtual address space(VIRT)，representing the non-swapped physical memory a task is currently using. It is also the sum of the RSan, RSfd and RSsh fields.

VIRT 占用的虚拟内存，RES占用的物理内存，SHR占用的共享内存

#### SHR

A subset of resident memory(RES) that may be used by other processes

It will include shared anonymous  pages  and  shared  file-backed  pages.  It  also  includes private pages mapped to files representing program images and shared libraries

#### SWAP

The formerly resident portion of a task's address space  written to the swap file when physical memory becomes over committed.

#### VIRT

The total amount  of  virtual  memory  used  by  the  task.  It  includes  all  code,  data  and shared libraries plus pages that  have been swapped out and pages that have been  mapped  but  not  used.

#### %CPU

The task's share of the elapsed CPU time since the last screen update

==如果一个进程是多线程的，如果top没有以thread mode 运行可能会显示超过100%==，interactive mode 使用`H`可以切换

#### %MEM

进程占用的physical memory比率(RES/total physical memory)

#### CODE

可执行代码占用的物理内存

#### COMMAND

显示启动进程的命令，interactive mode可以使用`c`来切换，如果命令太长就会被truncate

```
1 root      20   0  128136   6288   3760 S  0.0  1.3   0:01.87 /usr/lib/systemd/systemd --switched-root --system --deseriali+
```

#### DATA

进程占用的private memory，也被称为DRS，不会映射到physical memory（RES），但是会计算到vitrual memory（VIRT）

#### ENVIRON

显示进程使用的环境变量，会被truncate

#### GID/GROUP

进程对应的GID/group name

#### NI

nice value，==A negative nice value means higher priority，whereas a positive nice value means lower priority==  Zero in this field simply means priority will not be adjusted in determining a task's dispatch-ability.

优先级越高，被cpu处理越早

#### OOMa/OOMs

当内存耗尽时，进程kill掉的优先级，1000 always kill，0 never kill

#### P

a number representing the last used cpu

#### PID

进程号

This value may also be used as: a process group ID (see PGRP); a session ID for the session leader (see SID); a thread group ID for the thread group leader (see TGID); and a TTY process group ID for the process group leader (see TPGID).

#### PPID

父进程号，使用forest mode 比较好

#### PR

cpu 调度进程的优先级，如果值为rt 表示 rael time scheduling priority 即当前调用被cpu调用的

#### S

进程状态

1. D = uninterruptibel slepp
2. I = idle
3. R = running(ready to run)
4. S = sleeping
5. T = stopped by job control signal
6. t = stopped by dubugger during trace
7. Z = zombie

#### TIME/TIME+

从进程启动后占用CPU的时间，TIME+毫秒级单位(别和 `ps` 的 STIME 混淆)

#### TTY

进程关联的TTY

#### UID

启动进程的uid

#### USER

启动进程的用户

#### USED

This field represents the non-swapped physical memory a task  is  using  (RES)  plus  the swapped out portion of its address space  (SWAP).

#### nDRT

https://www.thegeekdiary.com/what-are-dirty-pages-in-linux/

进程使用的dirty pages

if old page has been modified already then it must be preserved somewhere so application/database can re-used later on – this is called dirty page.

dirty page 一般存储在swap中

#### nMaj/nMin/VMj/VMn

https://www.quora.com/What-is-the-difference-between-minor-and-major-page-fault-in-Linux

marjor page faults delta

minor page faults delta

#### nTH

进程关联的线程数

## mode optionals

- `-H`

  以threads-mode 运行top，如果没有改参数，一个进程显示所有线程占用的系统资源

- `-b`

  以 batch-mode 运行 top，可以将top的输出输入到文件，通常和`-n`一起使用表示更新多少次后终止

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

## none interactive optionals

- `-O`

  输出top可以使用的field-names

### input optionals

- `-d`

  指定top一次更新的时间间隔

  ```
  #10s更新一次
  root in /tmp λ top -d 10
  ```

- `-p <pid>`

  只查看pid关联的进程，可以使用comma分隔查看多个pid，pid 0 表示 top 进程

- `-u <uid | username>`

  查看指定用的启用的进程

### output optionals

- `-o <field-names>`

  按照指定列排序，在列名前添加`+`表示从高到低，`-`表示从低到高，具体可以使用的列名使用`-O`参数查看

  ```
  #按照内存使用率从高到低排序
  root in /tmp λ top -o +%MEM
  ```

- `-c`

  输出的command以全路径参数显示

  ```
   PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND                                                         
      1 root      20   0  128136   6288   3760 S  0.0  1.3   0:01.63 /usr/lib/systemd/systemd --switched-root --system --deserialize+
  ```

## interactive optionals

### common

- `? | help`

  帮助信息

- `enter | space`

  主动刷新top显示 

- `=`

  清空当前的过滤条件

- `0`

  0值是否显示

- `B`

  字体加粗功能

- `d`

  修改delay 时间

- `H`

  以thread mod显示

- `k`

  杀死指定进程

- `r`

  renice a task

- `W`

  ==持久化当前在top中的修改的设置，默认存储在`~/.toprc`中，下次使用top是自动调用该配置文件==

- `Z`

  ==修改top展示信息的颜色==，使用`w`来当前选中的更改面板，默认USR，使用RGB 256

  https://www.ditig.com/256-colors-cheat-sheet

- `X`

  设置列宽，`-1`表示自动scale

- `Y`

  在`.toprc`中输入如下内容，==可以查看指定进程打开的文件，日志，NUMA==，同样还是从syslog中读取日志

  ```
   /bin/echo -e "pipe\tOpen Files\tlsof -P -p %d 2>&1" >> ~/.toprc
   /bin/echo -e "file\tNUMA Info\t/proc/%d/numa_maps" >> ~/.toprc
   /bin/echo -e "pipe\tLog\ttail -n200 /var/log/syslog | sort -Mr" >> ~/.toprc
  ```

### sorting

- `M`

  按照进程使用的内存大小排序

- `N`

  按照进程的PID排序

- `P`

  按照进程的CPU使用率排序

- `T`

  按照进程启用的时间排序

### summary area

- `1 | 4`

  显示每核cpu使用率，使用4多行显示

- `l`

  toggle uptime, load average

- `t`

  切换summary area cpu 显示的模式

- `m`

  切换summary area memory 显示的模式

- `E`

  调整summary area memmory 显示的单位

###  searching 

- `L`

  从所有field中选取含有指定字符的进程，使用`&`匹配下一个

### Fields/Columns Header

- `f`

  调整显示的fields

- `c`

  显示command column

### Task Area

- `J | j`

  显示内容左右对齐  

- `x | y` 

  主键列高亮，running tasks 高亮 

- `S`

  累计占用cpu的时间

- `U | u`

  查看指定用户起的进程

- `z`

  更改主题颜色到黑白

- `A`

  进入多选模式，使用`g`选择使用的面板,再次键入`A`使用选择的面板

- `i`

  idle-process toggle，只显示active tasks

- `V`

  以forest mode 显示，和pstree类似显示子进程

  ```
    910 root      20   0   89704   2204   1168 S  0.0  0.4   0:00.34  `- master                                                    
    912 postfix   20   0   89876   4084   3076 S  0.0  0.8   0:00.05      `- qmgr                                                  
  22167 postfix   20   0   89808   4056   3060 S  0.0  0.8   0:00.00      `- pickup 
  ```

- `n`

  只显示指定数目的进程

## config

> 注意该配置文件不能被直接使用，需要通过设置然后使用 `W` 保存生成

![2022-02-24_23-46](https://github.com/dhay3/image-repo/raw/master/20220224/2022-02-24_23-46.2jvf71r4shk0.png)

```
cpl in ~/.config/procps λ cat toprc 
top's Config File (Linux processes with windows)
Id:j, Mode_altscr=1, Mode_irixps=1, Delay_time=1.0, Curwin=0
Def     fieldscur=����'34�E79:@)=;*+,-./01268<>?ABCFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz
        winflags=195892, sortindx=0, maxtasks=0, graph_cpus=0, graph_mems=0, double_up=0, combine_cpus=0
        summclr=3, msgsclr=3, headclr=3, taskclr=3
Job     fieldscur=����:(���;=@<�E)*+,-./012568>?ABCFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz
        winflags=195892, sortindx=18, maxtasks=0, graph_cpus=0, graph_mems=0, double_up=0, combine_cpus=0
        summclr=208, msgsclr=208, headclr=208, taskclr=208
Mem     fieldscur=���<�����MBN�D347E&'()*+,-./0125689FGHIJKLOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz
        winflags=195892, sortindx=21, maxtasks=0, graph_cpus=0, graph_mems=0, double_up=0, combine_cpus=0
        summclr=93, msgsclr=93, headclr=93, taskclr=93
Usr     fieldscur=�&D5'(*798:=M�)+,-./123406;<>?@ABCFGHIJKLNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz
        winflags=64950, sortindx=20, maxtasks=0, graph_cpus=0, graph_mems=0, double_up=0, combine_cpus=0
        summclr=2, msgsclr=2, headclr=2, taskclr=2
Fixed_widest=0, Summ_mscale=2, Task_mscale=0, Zero_suppress=0

pipe    Open Files      lsof -P -p %d 2>&1
file    NUMA Info       /proc/%d/numa_maps
pipe    Log     tail -n200 /var/log/syslog | sort -Mr
```































