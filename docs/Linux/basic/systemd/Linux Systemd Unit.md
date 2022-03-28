ref:
[http://standards.freedesktop.org/desktop-entry-spec/latest/](http://standards.freedesktop.org/desktop-entry-spec/latest/)
systemd.unit(5)
[官方文档](https://www.freedesktop.org/software/systemd/man/systemd.unit.html)

## Digest

用于管理系统资源，可以如下几种Unit，具体查看对应的man page

```bash
Service unit：系统服务
Target unit：多个 Unit 构成的一个组
Device Unit：硬件设备
Mount Unit：文件系统的挂载点
Automount Unit：自动挂载点
Path Unit：文件或路径
Scope Unit：不是由 Systemd 启动的外部进程
Slice Unit：进程组
Snapshot Unit：Systemd 快照，可以切回某个快照
Socket Unit：进程间通信的 socket
Swap Unit：swap 文件
Timer Unit：定时器
```

有如下几条规则需要注意

1. 如果systemd在读取unit配置是出错了，会报错同时continue loading the unit
1. boolean 可以有多种形式。1、yes、true 等价，0、no、false 等价
1. 支持 time span，默认以 sec 为单位
1. empty lines and lines starting with # or ; are ignored. Lines ending in a backslash are concatednated with the following line while reading and teh backslash is repalced by a space character
1. 如果一个文件叫 foo.service，他可以使用 foo.service.wants( 是一个文件 ) 来指代 `Wants=` directive
1. systemd 在找 unit file，默认会先按照 literal unit name 找，如果没有找到且 unit name 中包含 @ ，就会按照 the part between the @ character and the suffix 去找。例如：if a service getty[@tty3.service ](/tty3.service ) is requested and no file by that name is found, systemd will look for getty[@.service ](/.service ) and instantiate a service from that configuration file if it is found. 
1. 如果 unit file 大小为 0 或者 symlinked to  /dev/null，改配置不会被加载也不能被手动启动

## Unit load path

unit files 默认从如下目录中载入，也可以使用`systemctl`来指定其他目录。按照表格中出现的path表示优先级

```bash
┌────────────────────────┬─────────────────────────────┐
│Path                    │ Description                 │
├────────────────────────┼─────────────────────────────┤
│/etc/systemd/system     │ Local configuration         │
├────────────────────────┼─────────────────────────────┤
│/run/systemd/system     │ Runtime units               │
├────────────────────────┼─────────────────────────────┤
│/usr/lib/systemd/system │ Units of installed packages │
└────────────────────────┴─────────────────────────────┘
```

## Unit file
> 查看某个 directives 能在那种 unit files 中使用，可以查看`systemd.direcitves(7)`


配置文件主要放在`/usr/lib/systemd/system`目录，也可能在`/etc/systemd/system`目录。找到配置文件以后，使用文本编辑器打开即可。使用，查看配置文件的内容`systemctl cat <unit>`

```shell
[root@cyberpelican ~]# systemctl cat httpd
# /usr/lib/systemd/system/httpd.service
[Unit]
Description=The Apache HTTP Server
After=network.target remote-fs.target nss-lookup.target
Documentation=man:httpd(8)
Documentation=man:apachectl(8)

[Service]
Type=notify
EnvironmentFile=/etc/sysconfig/httpd
ExecStart=/usr/sbin/httpd $OPTIONS -DFOREGROUND
ExecReload=/usr/sbin/httpd $OPTIONS -k graceful
ExecStop=/bin/kill -WINCH ${MAINPID}
# We want systemd to give httpd some time to finish gracefully, but still want
# it to kill httpd after TimeoutStopSec if something went wrong during the
# graceful stop. Normally, Systemd sends SIGTERM signal right after the
# ExecStop, which would kill httpd. We are sending useless SIGCONT here to give
# httpd time to finish.
KillSignal=SIGCONT
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

从上面的输出可以看到，配置文件分成几个区块。每个区块的第一行，是用方括号表示的区别名，比如`[Unit]`。注意，配置文件的区块名和字段名，都是大小写敏感的。

每个区块内部是一些等号连接的键值对。

```
[Section]
Directive1=value
Directive2=value
. . .
```

注意，键值对的等号两侧不能有空格。

### Unit directives

`[Unit]`指令块通常是配置文件的第一个区块，用来定义 Unit 的元数据，以及配置与其他 Unit 的关系。它的主要字段如下：

- Description：简短描述
- Documentation：文档地址
- Requires：当前 Unit 依赖的其他 Unit，如果它们没有运行，当前 Unit 会启动失败
- Wants：与当前 Unit 配合的其他 Unit，如果它们启动失败，当前 Unit 不会启动失败
- BindsTo：与Requires类似，它指定的 Unit 如果退出，会导致当前 Unit 停止运行
- Before：如果该字段指定的 Unit 也要启动，那么必须在当前 Unit 之后启动
- After：如果该字段指定的 Unit 也要启动，那么必须在当前 Unit 之前启动
- Conflicts：这里指定的 Unit 不能与当前 Unit 同时运行
- Condition...：当前 Unit 运行必须满足的条件，否则不会运行，显示condition failed

```bash
[root@centos7 default.target.wants]# systemctl status systemd-readahead-replay.service 
● systemd-readahead-replay.service - Replay Read-Ahead Data
   Loaded: loaded (/usr/lib/systemd/system/systemd-readahead-replay.service; enabled; vendor preset: enabled)
   Active: inactive (dead)
Condition: start condition failed at Tue 2022-02-22 07:27:21 UTC; 1 weeks 0 days ago
```

- Assert...：当前 Unit 运行必须满足的条件，否则会报启动失败
- OnFailure：单 Unit 启动失败是调用的 Units
- StopWhenUnneeded: 当前 Unit 不在被使用是时会自动停止，默认 false
- RefusemanualStart：Unit不能被用户手动启动，默认 false
- RefusemanualStop：Unit不能被用户手动关闭，默认 false
- AllowIsolate：Unit可以被 `systemctl isolate`隔离，默认false

### Type-specific directives

不同的 unit 有不通的指令块，这里只介绍几个，具体查看 man page

#### Service directives

A unit configuration file whose name ends in `.service` encodes informatin about a process controlled and supervised by systemd
`[Service]`区块用来 Service 的配置，只有 Service 类型的 Unit 才有这个区块。它的主要字段如下。

-  Type：定义启动时的进程行为。它有以下几种值。 
   1. simple：默认值(如果没有Type, BusName 指令块，但是有 ExecStart 指令块)，执行ExecStart指定的命令作为主进程
   1. forking：以 fork 方式从父进程创建子进程，创建后父进程会立即退出，最好指定PIDFILE指令块
   1. oneshot：一次性进程，Systemd 会等当前服务退出，再继续往下执行。不能和Type，ExecStart 指令块一起使用
   1. dbus：当前服务通过D-Bus启动
   1. notify：当前服务启动完毕，会通知Systemd，再继续往下执行
   1. idle：若有其他任务执行完毕，当前服务才会运行
-  RemainAfterExit：在该 service 所有的进程停止后，服务状态是否为active，默认 no 
-  GuessMainPID：计算 service 的PID，如果 service 由多个进程组成，结果可能不准确 
-  PIDFILE：指定 service 使用的PID file 
-  ExecStart：启动当前服务的命令，command 必须是绝对路径，如果指定了多条 command，按照出现的先后顺序执行，如果其中的一条命令执行失败了， service 也会失败。如果 Type 是 forking 表示该参数对应的是主进程 
-  ExecStartPre：启动当前服务之前执行的命令 
-  ExecStartPost：启动当前服务之后执行的命令 
-  ExecReload：重启当前服务时执行的命令 
-  ExecStop：停止当前服务时执行的命令 
-  ExecStopPost：停止当其服务之后执行的命令 
-  RestartSec：自动重启当前服务间隔的秒数 
-  Restart：定义何种情况 Systemd 会自动重启当前服务，可能的值包括always（总是重启）、on-success、on-failure、on-abnormal、on-abort、on-watchdog 

```bash
┌───────────────┬────┬────────┬────────────┬────────────┬─────────────┬──────────┬─────────────┐
│Restart        │ no │ always │ on-success │ on-failure │ on-abnormal │ on-abort │ on-watchdog │
│settings/Exit  │    │        │            │            │             │          │             │
│causes         │    │        │            │            │             │          │             │
├───────────────┼────┼────────┼────────────┼────────────┼─────────────┼──────────┼─────────────┤
│Clean exit     │    │ X      │ X          │            │             │          │             │
│code or signal │    │        │            │            │             │          │             │
├───────────────┼────┼────────┼────────────┼────────────┼─────────────┼──────────┼─────────────┤
│Unclean exit   │    │ X      │            │ X          │             │          │             │
│code           │    │        │            │            │             │          │             │
├───────────────┼────┼────────┼────────────┼────────────┼─────────────┼──────────┼─────────────┤
│Unclean signal │    │ X      │            │ X          │ X           │ X        │             │
├───────────────┼────┼────────┼────────────┼────────────┼─────────────┼──────────┼─────────────┤
│Timeout        │    │ X      │            │ X          │ X           │          │             │
├───────────────┼────┼────────┼────────────┼────────────┼─────────────┼──────────┼─────────────┤
│Watchdog       │    │ X      │            │ X          │ X           │          │ X           │
└───────────────┴────┴────────┴────────────┴────────────┴─────────────┴──────────┴─────────────┘
```

- TimeoutSec：定义 Systemd 停止当前服务之前等待的秒数
- Environment：指定环境变量

### Install directives

`[Install]`通常是配置文件的最后一个区块，用来定义如何启动，以及是否开机启动。它的主要字段如下：

-  WantedBy：如果指定的 Unit started，当前 Unit 会自动启动
[https://unix.stackexchange.com/questions/506347/why-do-most-systemd-examples-contain-wantedby-multi-user-target](https://unix.stackexchange.com/questions/506347/why-do-most-systemd-examples-contain-wantedby-multi-user-target) 
-  RequiredBy：如果指定的 Unit started，当前 Unit 会自动启动 
-  Alias：当前 Unit 可用于启动的别名，必须和 unit file 有一样的 suffix 
-  Also：当前 Unit enable 时，会同时 install 的其他 Unit。当前 Unit disable 时，会被 uninstall 的其他 Units 

## Instances

### 0x001 foo.service

```bash
[Unit]
Description=Foo

[Service]
ExecStart=/usr/sbin/foo-daemon

[Install]
WantedBy=multi-user.target
```

当运行 `systemctl enable foo.service`时，会生成`/etc/systemd/system/multi-user.target.wants/foo.service`symlink 指向实际的 unit file。即当 multi-user.target 启动时，自动启动 foo.service。对应的`systemcl disable foo.service`会删除symlink

### 0x002 sshd.service

```shell
[root@cyberpelican system]# systemctl cat sshd
# /usr/lib/systemd/system/sshd.service
[Unit]
Description=OpenSSH server daemon
Documentation=man:sshd(8) man:sshd_config(5)
After=network.target sshd-keygen.service
Wants=sshd-keygen.service

[Service]
Type=notify
EnvironmentFile=/etc/sysconfig/sshd
ExecStart=/usr/sbin/sshd -D $OPTIONS
ExecReload=/bin/kill -HUP $MAINPID
KillMode=process
Restart=on-failure
RestartSec=42s

[Install]
WantedBy=multi-user.target
```

#### Unit block

- After
表示sshd需要在`network.target`和 `sshd-keygen.service`启动之后启动
- Wants
表示`sshd.service`与`sshd-keygen.service`之间存在"弱依赖"关系，即如果"sshd-keygen.service"启动失败或停止运行，不影响`sshd.service`继续执行。
- Requires
表示"强依赖"关系，即如果该服务启动失败或异常退出，那么`sshd.service`也必须退出。
- EnvironmentFile
指定当前服务的环境参数文件。该文件内部的`key=value`键值对，可以用`$key`的形式，在当前配置文件中获取。

> 许多软件都有自己的环境参数文件，该文件可以用`EnvironmentFile`字段读取。


#### Service block

-  ExecStart
定义启动进程时执行的命令。上面的例子中，启动`sshd`，执行的命令是`/usr/sbin/sshd -D $OPTIONS`，其中的变量`$OPTIONS`就来自`EnvironmentFile`字段指定的环境参数文件。与之作用相似的，还有如下这些字段。
请看下面的例子。
上面这个配置文件，第二行`ExecStart`设为空值，等于取消了第一行的设置，运行结果如下。
所有的启动设置之前，都可以加上一个连词号（`-`），表示"抑制错误"，即发生错误的时候，不影响其他命令的执行。比如，`EnvironmentFile=-/etc/sysconfig/sshd`（注意等号后面的那个连词号），就表示即使`/etc/sysconfig/sshd`文件不存在，也不会抛出错误。    
```
ExecReload字段：重启服务时执行的命令
ExecStop字段：停止服务时执行的命令
ExecStartPre字段：启动服务之前执行的命令
ExecStartPost字段：启动服务之后执行的命令
ExecStopPost字段：停止服务之后执行的命令
```
```
ExecReload字段：重启服务时执行的命令
ExecStop字段：停止服务时执行的命令
ExecStartPre字段：启动服务之前执行的命令
ExecStartPost字段：启动服务之后执行的命令
ExecStopPost字段：停止服务之后执行的命令
```
```
execstart2
post1
post2
```

-  Type
定义启动类型。它可以设置的值如下。
下面是一个`oneshot`的例子，笔记本电脑启动时，要把触摸板关掉，配置文件可以这样写。
上面的配置文件，启动类型设为`oneshot`，就表明这个服务只要运行一次就够了，不需要长期运行。
如果关闭以后，将来某个时候还想打开，配置文件修改如下。
上面配置文件中，`RemainAfterExit`字段设为`yes`，表示进程退出以后，服务仍然保持执行。这样的话，一旦使用`systemctl stop`命令停止服务，`ExecStop`指定的命令就会执行，从而重新开启触摸板。    
```
simple（默认值）：ExecStart字段启动的进程为主进程
forking：ExecStart字段将以fork()方式启动，此时父进程将会退出，子进程将成为主进程
oneshot：类似于simple，但只执行一次，Systemd 会等它执行完，才启动其他服务
dbus：类似于simple，但会等待 D-Bus 信号后启动
notify：类似于simple，启动结束后会发出通知信号，然后 Systemd 再启动其他服务
idle：类似于simple，但是要等到其他任务都执行完，才会启动该服务。一种使用场合是为让该服务的输出，不与其他服务的输出相混合
```
```
[Unit]
Description=Switch-off Touchpad

[Service]
Type=oneshot
ExecStart=/usr/bin/touchpad-off

[Install]
WantedBy=multi-user.target
```
```
[Unit]
Description=Switch-off Touchpad

[Service]
Type=oneshot
ExecStart=/usr/bin/touchpad-off start
ExecStop=/usr/bin/touchpad-off stop
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
```

-  KillMode
定义 Systemd 如何停止 sshd 服务。
上面这个例子中，将`KillMode`设为`process`，表示只停止主进程，不停止任何sshd 子进程，即子进程打开的 SSH session 仍然保持连接。这个设置不太常见，但对 sshd 很重要，否则你停止服务的时候，会连自己打开的 SSH session 一起杀掉。
`KillMode`字段可以设置的值如下。  
```
control-group（默认值）：当前控制组里面的所有子进程，都会被杀掉
process：只杀主进程
mixed：主进程将收到 SIGTERM 信号，子进程收到 SIGKILL 信号
none：没有进程会被杀掉，只是执行服务的 stop 命令。
```

-  Restart
定义了 sshd 退出后，Systemd 的重启方式。
上面的例子中，`Restart`设为`on-failure`，表示任何意外的失败，就将重启sshd。如果 sshd 正常停止（比如执行`systemctl stop`命令），它就不会重启。
`Restart`字段可以设置的值如下。
对于守护进程，推荐设为`on-failure`。对于那些允许发生错误退出的服务，可以设为`on-abnormal`。  
```
no（默认值）：退出后不会重启
on-success：只有正常退出时（退出状态码为0），才会重启
on-failure：非正常退出时（退出状态码非0），包括被信号终止和超时，才会重启
on-abnormal：只有被信号终止和超时，才会重启
on-abort：只有在收到没有捕捉到的信号终止时，才会重启
on-watchdog：超时退出，才会重启
always：不管是什么退出原因，总是重启
```

-  RestartSec
表示 Systemd 重启服务之前，需要等待的秒数。上面的例子设为等待42秒。 

#### Install block

- WantedBy
表示该服务所在的 Target。`Target`的含义是服务组，表示一组服务。`WantedBy=multi-user.target`指的是，sshd 所在的 Target 是。`multi-user.target`
这个设置非常重要，因为执行`systemctl enable sshd.service`命令时，`sshd.service`的一个符号链接，就会放在`/etc/systemd/system`目录下面的`multi-user.target.wants`子目录之中。(开机启动的服务)
