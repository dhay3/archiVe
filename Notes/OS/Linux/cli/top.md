---
createTime: 2024-12-13 17:41
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# top

## 0x01 Preface

`top` 是一个查看系统资源(主要是 CPU/MEM)使用率以及 processes 或者是 threads 占用系统资源的工具，默认会以 interactive mode 运行，由 4 部分组成

- Summary Area
- Fields/Columns Header
- Task Area
- Input/Message Line

例如

```
# Summary Area
top - 17:55:06 up 10 days,  3:28,  1 user,  load average: 0.00, 0.01, 0.05
Tasks: 221 total,   1 running, 220 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni, 99.9 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem : 32779900 total, 31618788 free,   733032 used,   428080 buff/cache
KiB Swap:  2097148 total,  2097148 free,        0 used. 31646984 avail Mem

# Input/Message Line
Locate string

# Fields/Columns Header
  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
  
# Task Area
 1803 mysql     20   0 2250968 399904  16284 S   0.7  1.2  83:03.81 mysqld
 1195 root      20   0  210648   8932   5508 S   0.3  0.0  11:16.35 vmtoolsd
25502 root      20   0  162064   2348   1584 R   0.3  0.0   0:00.06 top
    1 root      20   0  191404   4452   2600 S   0.0  0.0   0:06.46 systemd
    2 root      20   0       0      0      0 S   0.0  0.0   0:00.09 kthreadd
```

## 0x02 Linux Memory Type

在详细介绍 `top` 前，还需要了解 Linux memory type，主要有 3 种

- Physical memory
- Virtual memory
- Swap memory(file)

同时进程是不能直接调用内存的，需要通过系统提供的 API 寻址(mapping)来调用内存。API 能够寻址的内存类型严格按照如下 4 象限

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

### 0x02a Physical Memory

*Physical memory, a limited resource where code and data must reside when executed or referenced.*

物理内存，也被称为 RAM，其容量由内存颗粒决定。可以通过 `dmidecode -t memory` 来查其看容量

```
$ sudo dmidecode -t memory
...
Memory Device
...
        Size: 32 GB
Memory Device
...
        Size: 32 GB
```

也是速度最快的

### 0x02c Swap Memory(File)

*Swap file(memory), where modified(dirty) memory can be saved and later retrieved if too many demands are made on.*

以文件形式存在的内存，在 Windows 上称为 page memory(file)

内存会以文件的形式存储，如果后续需要这部分的内存，系统会将文件载入到物理内存中。其容量由 swap partition 决定。可以通过 `lsblk` 来查看

```
$ lsblk
NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
nvme0n1     259:0    0   1.8T  0 disk
├─nvme0n1p1 259:2    0   300M  0 part /boot/efi
├─nvme0n1p2 259:3    0   1.8T  0 part /
└─nvme0n1p3 259:4    0    10G  0 part [SWAP]
```

速度较 physical memory 慢

### 0x02b Virtual Memory

*Virtual memory, a nearly unlimited resource serving the following*

虚拟内存，可以看作 physical memory 和 swap memory 的结合

![](https://www.baeldung.com/wp-content/uploads/sites/4/2021/06/virtual-mem.png)

其容量由 CPU 架构决定。例如 32bit CPU 可以支持 $2^{32}bit = 4GB$ 的虚拟内存，而 64bit CPU 可以支持 $2^{64}bit = 2^{31}GB$ 的虚拟内存

可以通过 `dmidecode -t processor` 来查看 CPU 架构

```
$ sudo dmidecode -t processor
...
       Characteristics:
                64-bit capable
                Multi-Core
                Hardware Thread
                Execute Protection
                Enhanced Virtualization
                Power/Performance Control
```

## 0x03 Display Area

### 0x03a Summary Area

展示一些资源的总使用率或者是总量，可以通过 [Summary Area Command](#Summary%20Area%20Command) 修改显示的方式

例如

```
top - 17:55:06 up 10 days,  3:28,  1 user,  load average: 0.00, 0.01, 0.05
Tasks: 221 total,   1 running, 220 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni, 99.9 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem : 32779900 total, 31618788 free,   733032 used,   428080 buff/cache
KiB Swap:  2097148 total,  2097148 free,        0 used. 31646984 avail Mem
```

- `top - 17:55:06 up 10 days,  3:28,  1 user,  load average: 0.00, 0.01, 0.05`

	输出 uptime 以及 load average 等价于 `uptime`

- `Tasks: 221 total,   1 running, 220 sleeping,   0 stopped,   0 zombie`

	输出正在运行的进程数，不同状态的进程数(即 `task = processes | threads`)。可以通过 `H` 将进程数转为线程数

	有 4 种状态

	- running
	
		当前占用 CPU 资源的进程数

		task area `S` 字段会以 `R` 标示 running 进程
	
	- sleeping

		当前没有使用 CPU 资源或者等待 I/O 的进程数

		task area `S` 字段会以 `S` 标示 sleeping 进程

	- stopped

		暂时停止的进程，通常是 `ctrl + z` 或者是 `kill -STOP` 触发的进程

		task area `S` 字段会以 `T` 标示 stopped 进程

	- zombie

		通常是子进程已经退出了，但是父进程并没有通过 `wait()` 来读取子进程的状态导致的

		task area `S` 字段会以 `Z` 标示 zombie 进程

- `%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni, 99.9 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st`

	所有 tasks 从 last delay 开始占用 CPU 不同类型资源时间的平均值。根据 Thread mode on/off Irix mode on/off

	计算公式为 

	%% TODO %%

	主要分为如下几种 CPU 资源

	- us(user space)

		time running un-niced user processes

		未使用 `nice` 调整优先级的 tasks 占用 CPU 的时间，通常就是 tasks 默认的状态

	- sy(system space)

		time running kernel processes

		系统 tasks 占用 CPU 的时间 

	- ni(nice)

		time running niced user processes

		通过 `nice` 调整优先级的 tasks 占用 CPU 的时间

	- id(idle)

		time spent in the kernel idle handler

		CPU 没有处理任何 tasks 的时间

	- wa(I/O wait)

		time waiting for I/O completion

		CPU 处于 I/O wati 的时间

	- hi(hardware interrupts)

		time spent servicing hardware interrupts

		CPU 处理硬件中断的时间。比如 keyboards, mice 发送的信号

	- si(software interrupts)

		time spent servicing software interrupts

		CPU 处理软件中断的时间。比如 drivers 发送的信号

	- st(steal time)

		time stolen from this vm by the hypervisor

		CPU 用于 virtualization 的时间

- `KiB Mem : 32779900 total, 31618788 free,   733032 used,   428080 buff/cache`

	%% TODO %%

- `KiB Swap:  2097148 total,  2097148 free,        0 used. 31646984 avail Mem`

	%% TODO %%

### 0x03b Input/Message Line

```
Locate string
```

### 0x03c Fields/Columns Header

task area 显示的字段，可以通过 `f` interacrive command 来设置显示的字段

例如

```
PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
```

常见的字段有

#### %CPU -- CPU Usage

> [!important]
> 千万不要将 `%CPU` 和 `%CPU(s)` 搞混，两者计算方式不同

*The task's share of the elapsed CPU  time  since the  last  screen  update,  expressed  as  a percentage of total CPU time.*

单 task 从 last delay 开始占用 CPU 时间的均值，没有做任何配置的情况下 task 指 process(可能由多个 threads 组成)

计算公式为 $elasped\ CPU\ time \div total\ CPU\ time$

elasped CPU time 会根据 thread mode on/off 来取值

total CPU time 会根据 Irix mode on/off 来取值

> [!note]
> 通过 interactive command `H` 设置 thread mode
> 
> 通过 interactive command `I` 设置 Irix mode

定义如下 EBNF

```
task = process with one thread | process with multiple threads | thread

elasped CPU time = sum(one task's us,sy,ni,wa,hi,si,st) since the last delay

total CPU time = time of delay [multily the number of cores]
```

如果 thread mode on

```
elasped CPU time = sum(one thread's us,sy,ni,wa,hi,si,st) since the last delay
```

如果 thread mode off(缺省值)

```
elasped CPU time = sum(one task's thread's us,sy,ni,wa,hi,si,st[,...]) since the last delay
```

如果 Irix mode on(缺省值)

```
total CPU time = time of delay
```

如果 Irix mode off

```
total CPU time = time of delay multily the number of cores
```

假设有一个 4 core server，delay time 为 3s，一个进程有 2 个线程，从 last delay 开始计算分别占用了 0.5s/1s

- 如果 thread mode on，Irix mode on，那么可以得出这 2 个 task $CPU\% = 0.5 \div 3$，$CPU\% = 1 \div 3$
- 如果 thread mode on，Irix mode off，那么可以得出这 2 个 task $CPU\% = 0.5 \div (3 \times 4)$，$CPU\% = 1 \div (3 \times 4)$
- 如果 thread mode off，Irix mode on，那么可以得出 $CPU\% = (0.5 + 1)\div 3$
- 如果 thread mode off，Irix mode off，那么可以得出 $CPU\% = (0.5 + 1) \div (3 \times 4)$

**需要注意的是在 thread mode off，Irix mode on，multi-core 的场景下(即缺省) CPU% 是有可能大于 100%**

假设一个 4 core server，delay time 为 3s，一个进程有 3 个线程，thread mode off，从 last delay 开始计算 3 个线程分别占用了不同 cores 1s 1.5s 1s，那么可以得出 $CPU\% = (1 + 1.5 + 1)\div 3 \gt 1$

可以得出如下结论

- thread mode on，Irix mode on

	CPU% = 进程每个线程每秒占用 CPU 的时间比率

- thread mode on，Irix mode off

	CPU% = 进程每个线程每秒占用每 Core 的时间比率

- thread mode off，Irix mode on

	CPU% = 进程每秒占用 CPU 的时间比率

- thread mode off，Irix mode off

	CPU% = 进程每秒占用每 Core 的时间比率

#### %MEM -- Memory Usage

*A task's currently resident share of available physical memory.*

#### CGROUPS -- Control Groups



#### CODE -- Code Size(KiB)


#### COMMAND -- Command Name or Command Line

#### DATA -- Data + Stack Size(KiB)

#### ENVIRON -- Environment Variables

#### Flags -- Task Flags

#### GID -- Group Id

#### GROUP -- Group Name

#### NI -- Nice Value

#### PGRP -- Process Group Id

#### PID -- Process Id

#### PPID -- Process Id

#### PR -- Priority

#### RES -- Resident Memory Size(KiB)



#### RSS -- Residient Memory, smaps(KiB)



### 0x03d Task Area

```
 1803 mysql     20   0 2250968 399904  16284 S   0.7  1.2  83:03.81 mysqld
 1195 root      20   0  210648   8932   5508 S   0.3  0.0  11:16.35 vmtoolsd
```


## 0x04 Interactive Mode

在没有使用任何参数的情况下 `top` 会进入 interactive mode，用户可以通过 keystrokes 来修改/增加/删除显示的一些选项

keystrokes 按照组成的部分分类

### Global Command

- `Enter | Space`

	refresh-display

	刷新 Summary Area/Task Area

- `? | h`

	help

	显示帮助信息

- `=`

	exit-display-limits

	主要用于清空 `L` 过滤的条件

- `0`

	zero-suppress

	0 值是否显示，UID/GID/NI/PR/P 不受影响

- `A`

	alternate-display-mode

- `B`

	bold-disable/enable

	summary area 和 task area 部分字体是否加粗

- `d | s`

	change-delay-time-interval

	修改 refresh delay，默认 3s，可以通过 `? | h` 来查看当前的 delay 值

- `H`

	threads-mode

	以 thread 展示 tasks，而不是默认的 processes

- `I`

	Irix-mods

	规定了 CPU% 的计算方式

	

- `k`

	kill-a-task

	kill 掉指定 PID 的进程

- `q`

	quit

	退出 `top` interactive mode

- 



### Summary Area Command

- `E`

	enforce-summary-memory-scale

	调整 summary area memory 显示的单位，KiB - EiB

- `1 | 4`

  显示
  显示每核cpu使用率，使用4多行显示

- `l`

  toggle uptime, load average

- `t`

  切换summary area cpu 显示的模式

```
	
```

- `m`

  切换summary area memory 显示的模式

- `E`

  调整summary area memmory 显示的单位

### Task Aread Command

- `e`

	enforce-task-memory-scale

	调整 task area memory 显示的单位，KiB - PiB

- `g`

	choose-another-window/field-group

	选择面板，不做任何配置不同面板 task area 字段不同，1 - 4

## Non-interactive Mode

### Optional Args


## The CPU% shows greater than 100%

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man top.1`


***References***

- [operating system - What's the difference between "virtual memory" and "swap space"? - Stack Overflow](https://stackoverflow.com/questions/4970421/whats-the-difference-between-virtual-memory-and-swap-space)












