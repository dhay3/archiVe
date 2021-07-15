# SysRq

参考：

https://en.wikipedia.org/wiki/Magic_SysRq_key

https://www.kernel.org/doc/html/latest/admin-guide/sysrq.html?highlight=sysrq

https://ubuntuqa.com/article/9526.html

## 概述

magic SysRq key is a key combination，allow the user perform low-level commands regardless of the system's state  expcept fro kernel panic or hardware failure

## 启用SysRq

> man proc /sysrq

想要使用SysRq，内核配置必须已使用`CONFIG_MAGIC_SYSRQ = y`，怎么查看内核配置呢？使用如下方法

```
#查看linux 的内核版本
ubuntu@win2k:/boot$ uname -r
4.15.0-118-generic

#查看内核配置文件
ubuntu@win2k:/boot$ cat config-4.15.0-118-generic | grep -i config_magic_sysrq
CONFIG_MAGIC_SYSRQ=y
#默认启用的功能，必须是十六进制。十进制413 = 256 + 128 + 16 + 8 + 4 + 1
CONFIG_MAGIC_SYSRQ_DEFAULT_ENABLE=0x01b6
CONFIG_MAGIC_SYSRQ_SERIAL=y
```

## 修改SysRq

> 有些distro为了安全不会分配1的权限

`/proc/sys/kernel/sysrq`控制SysRq允许调用的功能，可以出现如下值的组合

    0    Disable sysrq completely
    
    1    Enable all functions of sysrq
    
    > 1  Bit mask of allowed sysrq functions, as follows:
         2  Enable control of console logging level
         4  Enable control of keyboard (SAK, unraw)
         8  Enable debugging dumps of processes etc.
         16  Enable sync command
         32  Enable remount read-only
         64  Enable signaling of processes (term, kill, oom-kill)
         128  Allow reboot/poweroff
         256  Allow nicing of all real-time tasks
通过如下方式修改，number可以是hexdeicimal或decimal

```
echo "number" >/proc/sys/kernel/sysrq
```

也可以通过修改`sysctl.conf`持久配置

```
kernel.sysrq = "number"
```

## 调用SysRq

### x86

alt + SysRq + command key

### SPARC

alt + stop + command key

### PowerPC

alt + printScreen

### ==通用==

```
echo "command key" > /proc/sysrq-trigger
```

## command key

| Command | Function                                                     |
| ------- | ------------------------------------------------------------ |
| `b`     | Will immediately reboot the system without syncing or unmounting your disks. |
| `c`     | Will perform a system crash by a NULL pointer dereference. A crashdump will be taken if configured. |
| `d`     | Shows all locks that are held.                               |
| `e`     | Send a SIGTERM to all processes, except for init.            |
| `f`     | Will call the oom killer to kill a memory hog process, but do not panic if nothing can be killed. |
| `g`     | Used by kgdb (kernel debugger)                               |
| `h`     | Will display help (actually any other key than those listed here will display help. but `h` is easy to remember :-) |
| `i`     | Send a SIGKILL to all processes, except for init.            |
| `j`     | Forcibly “Just thaw it” - filesystems frozen by the FIFREEZE ioctl. |
| `k`     | Secure Access Key (SAK) Kills all programs on the current virtual console. NOTE: See important comments below in SAK section. |
| `l`     | Shows a stack backtrace for all active CPUs.                 |
| `m`     | Will dump current memory info to your console.               |
| `n`     | Used to make RT tasks nice-able                              |
| `o`     | Will shut your system off (if configured and supported).     |
| `p`     | Will dump the current registers and flags to your console.   |
| `q`     | Will dump per CPU lists of all armed hrtimers (but NOT regular timer_list timers) and detailed information about all clockevent devices. |
| `r`     | Turns off keyboard raw mode and sets it to XLATE.            |
| `s`     | Will attempt to sync all mounted filesystems.                |
| `t`     | Will dump a list of current tasks and their information to your console. |
| `u`     | Will attempt to remount all mounted filesystems read-only.   |
| `v`     | Forcefully restores framebuffer console                      |
| `v`     | Causes ETM buffer dump [ARM-specific]                        |
| `w`     | Dumps tasks that are in uninterruptable (blocked) state.     |
| `x`     | Used by xmon interface on ppc/powerpc platforms. Show global PMU Registers on sparc64. Dump all TLB entries on MIPS. |
| `y`     | Show global CPU Registers [SPARC-64 specific]                |
| `z`     | Dump the ftrace buffer                                       |
| `0`-`9` | Sets the console log level, controlling which kernel messages will be printed to your console. (`0`, for example would make it so that only emergency messages like PANICs or OOPSes would make it to your console.) |

## tricks

1. 使用alt + sysrq + k (secure access key)会杀掉当前console的所有进程(如果不是login shell)，保证是正真的login shell。防止是trojan program调用的假login shell。
2. 无法通过命令重启，alt + sysrq + b
3. 保证数据的写入，alt + sysrq + s
4. 修改kernel 日志等级，低到高越来越详细













