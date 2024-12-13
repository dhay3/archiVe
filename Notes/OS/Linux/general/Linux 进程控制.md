# 进程控制

参考：

https://program-think.blogspot.com/2019/11/POSIX-TUI-from-TTY-to-Shell-Programming.html?q=bash&scope=all

[TOC]

　一般来说，每个“进程”都是由另一个进程启动滴。如果“进程A”创建了“进程B”，则 A 是【父进程】，B 是【子进程】（这个“父子关系”很好理解——因为完全符合直觉）
　　一般来说，第一个进程由【操作系统内核】（kernel）亲自操刀运行起来；而 kernel 又是由“引导扇区”中的“boot loader”加载。==通过pstree来查看进程数，通过 -p  参数来查看pid==

```
[root@chz Desktop]# pstree
systemd─┬─ModemManager───2*[{ModemManager}]
        ├─NetworkManager───2*[{NetworkManager}]
        ├─VGAuthService
        ├─2*[abrt-watch-log]
        ├─abrtd
        ├─accounts-daemon───2*[{accounts-daemon}]
        ├─alsactl
        ├─at-spi-bus-laun─┬─dbus-daemon───{dbus-daemon}
        │                 └─3*[{at-spi-bus-laun}]
        ├─at-spi2-registr───2*[{at-spi2-registr}]
        ├─atd
        ├─auditd─┬─audispd─┬─sedispatch
        │        │         └─{audispd}
        │        └─{auditd}
```

## 进程终止的信号

> 我们可以通过kill -l 来查看所有的信号，==数字对应编号，使用名称时可以省略SIG==
>
> 我们可以通过`kill`来发送信号
>
> http://c.biancheng.net/view/1069.html

- SIGINT

  在大部分 POSIX 系统的各种终端上，==`Ctrl + C` 组合键触发的就是这个信号。==通常情况下，进程收到这个信号后，做完相关的善后工作，就自行了断（自杀）。

- SIGTERM

  这个信号基本类似于 SIGINT。它是 `kill` ＆ `killall` 这两个命令【默认】使用的信号。也就是说，当你用这俩命令杀进程，并且【没有】指定信号类型，==那么 `kill` 或 `killall` 用的就是这个 SIGTERM 信号。==

- SIGQUIT

  这个信号类似于前两个（SIGINT ＆ SIGINT），差别在于——进程在退出前会执行“[core dump](https://en.wikipedia.org/wiki/Core_dump)”操作。
  一般而言，只有程序员才会去关心“core dump”这个玩意儿，所以这里就不细聊了。

- SIGKILL

  在杀进程的几个信号中，这个信号是是最牛逼的（也是最粗暴的）。前面三个信号都是【可屏蔽】滴，而这个信号是【不可屏蔽】滴。==当某个进程收到了【SIGKILL】信号，该进程自己【完全没有】处理信号的机会，而是由操作系统内核直接把这个进程干掉。==此种行为可以形象地称之为“它杀”。当你用下列这些命令杀进程，本质上就是在发送这个信号进行【它杀】。【SIGKILL】这个信号的编号是 `9`，下列这些命令中的 `-9` 参数就是这么来滴。

  ```
  kill -9 进程号
  kill -KILL 进程号
  
  killall -9 进程名称
  killall -KILL 进程名称 
  killall -SIGKILL 进程名称
  ```

　为了方便对照上述这4种，俺放一个表格如下

| 信号名称 | 编号 | 能否屏蔽 | 默认动作                  | 俗称 |
| -------- | ---- | -------- | ------------------------- | ---- |
| SIGINT   | 2    | YES      | 进程自己退出              | 自杀 |
| SIGTERM  | 15   | YES      | 进程自己退出              | 自杀 |
| SIGQUIT  | 3    | YES      | 执行core dump进程自己退出 | 自杀 |
| SIGKILL  | 9    | NO       | 进程被kernel干掉          | 他杀 |

> 请注意==他杀是一种比较危险的做法，可能导致一些副作用==，只有在使用其他方法无法干掉某个进程，才考虑用这招

例如：

某个进程正在保存文件。这时候遭遇“它杀”可能会导致文件损坏。
　　（注：虽然某些操作系统能做到“写操作的原子性”，但数据存储可能会涉及多个写操作。当进程在作【多个】关键性写操作时，遭遇它杀。可能导致数据文件【逻辑上】的损坏）

## kill vs killall

两者差别在于--前者使用进程号，后者使用进程名

```
[root@chz Desktop]# firefox &
[1] 22518
[root@chz Desktop]# kill 22518
[1]+  Terminated  

[root@chz Desktop]# killall firefox
[1]+  Terminated              firefox
```

## 进程暂停

> 暂停的进程将失去响应

- 温柔式暂停（SIGTSTP）

  ```
  kill -TSTP <pid>
  ```

  【SIGTSTP】默认绑定到组合键【`Ctrl + Z`】，是【可】屏蔽的信号。也就是说，如果某个进程屏蔽了【SIGTSTP】信号，你就【无法】用该方式暂停它。这时候你就得改用【粗暴】的方式

- 粗暴式暂停（SIGSTOP）

  ```
  kill -STOP <pid>
  ```

  这个【SIGSTOP】信号与前面提及的【SIGKILL】有某种相同之处——这两个信号都属于【不可屏蔽】的信号。也就是说，收到【SIGSTOP】信号的进程【无法】抗拒被暂停（suspend）的命运。

## 恢复进程

如果想要resume被暂停的进程使用【SIGCONT】

```
kill -CONT <pid>
```

## ps VS top

- ps 显示的是当前正在运行的进程的快照，top动态显示
- top显示的字段与ps -aux差不多

