# Unit 

## 概述

Systemd 可以管理所有系统资源。不同的资源统称为 Unit（单位）。通过`systemd.unit`来查看，具体使用查看systmctl

```
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

## Unit 管理

> 1. 对unit的操作等价于操作`/etc/systemd/system`目录下的配置文件，
>
> 2. Sytemd默认的后缀名为`.service`，所以`httpd`会被理解成`httpd.service`。有socket和service两种
>
> 3. Unit 之间存在依赖关系：A 依赖于 B，就意味着 Systemd 在启动 A 的时候，同时会去启动 B。`systemctl list-dependencies`命令列出一个 Unit 的所有依赖。

- `systemctl`

  输出所有active的unit，等价于`systemctl list-units`

  ```
  [root@cyberpelican init.d]# systemctl
    UNIT                                                         LOAD   ACTIVE SUB       DESCRIPTION
    proc-sys-fs-binfmt_misc.automount                            loaded active waiting   Arbitrary Executable File Formats File System Automount Point
    sys-devices-pci0000:00-0000:00:07.1-ata2-host2-target2:0:0-2:0:0:0-block-sr0.device loaded active plugged   VMware_Virtual_IDE_CDROM_Drive CentOS_
    sys-devices-pci0000:00-0000:00:10.0-host0-target0:0:0-0:0:0:0-block-sda-sda1.device loaded active plugged   VMware_Virtual_S 1
    sys-devices-pci0000:00-0000:00:10.0-host0-target0:0:0-0:0:0:0-block-sda-sda2.device loaded active plugged   LVM PV eA52jE-SFuU-BG5t-Isyw-wWdY-lj4K
    sys-devices-pci0000:00-0000:00:10.0-host0-target0:0:0-0:0:0:0-block-sda.device loaded active plugged   VMware_Virtual_S
  ```

- `systemctl list-units`

  显示所有运行的unit，可以使用`--type`和`--state`参数

  ```
  [root@cyberpelican ~]# systemctl list-units|grep httpd
    httpd.service                                                                                                    loaded active running   The Apache HTTP Server
    
  [root@cyberpelican ~]# systemctl list-units --type=service --state=active
  ```

- `systemctl list-unit-files`

  列出所有配置文件，使用`-t`参数指定配置文件类型

  ```shell
  [root@cyberpelican init.d]# systemctl list-unit-files 
  UNIT FILE                                     STATE   
  proc-sys-fs-binfmt_misc.automount             static  
  dev-hugepages.mount                           static  
  dev-mqueue.mount                              static  
  proc-fs-nfsd.mount                            static  
  proc-sys-fs-binfmt_misc.mount                 static  
  sys-fs-fuse-connections.mount                 static  
  sys-kernel-config.mount                       static  
  sys-kernel-debug.mount                        static  
  
  ---
  
  [root@cyberpelican init.d]# systemctl list-unit-files -t service
  UNIT FILE                                     STATE   
  abrt-ccpp.service                             enabled 
  abrt-oops.service                             enabled 
  abrt-pstoreoops.service                       disabled
  abrt-vmcore.service                           enabled 
  abrt-xorg.service                             enabled 
  abrtd.service                                 enabled 
  accounts-daemon.service                       enabled 
  alsa-restore.service                          static  
  alsa-state.service                            static  
  anaconda-direct.service                       static  
  ```

  四种状态

  |         |                                                              |
  | ------- | ------------------------------------------------------------ |
  | enable  | 已建立启动链接，开机自动启动                                 |
  | disable | 没有建立链接                                                 |
  | static  | 该配置文件没有[Install]部分（无法执行），只能作为其他配置文件的依赖 |
  | masked  | 该配置文件被禁止建立启动链接，无法被启动                     |

  > 注意，从配置文件的状态无法看出，该 Unit 是否正在运行。这必须执行前面提到的`systemctl status`命令

- `systemctl is-active|is-failed|is-enabled`

  ```
  # 显示某个 Unit 是否正在运行
  [root@cyberpelican ~]# systemctl is-active httpd
  active
  
  # 显示某个 Unit 是否处于启动失败状态
  [root@cyberpelican ~]# systemctl is-failed httpd
  active
  
  # 显示某个 Unit 服务是否建立了启动链接
  [root@cyberpelican ~]# systemctl is-enabled httpd
  enabled
  ```

- `systemctl mask <unit>`

  禁用服务，取消链接

  ```
  [root@cyberpelican init.d]# systemctl mask httpd
  Created symlink from /etc/systemd/system/httpd.service to /dev/null.
  ```

- `systemctl unmask <unit>`

  取消禁用服务

  ```
  [root@cyberpelican init.d]# systemctl unmask httpd
  Removed symlink /etc/systemd/system/httpd.service.
  ```

- `systemctl status [pid|unit name]`

  查看当前运行所有unit的状态，无参表示查看所有

  ```shell
  [root@cyberpelican init.d]# systemctl status
  ● cyberpelican
      State: degraded
       Jobs: 0 queued
     Failed: 1 units
      Since: Thu 2020-10-29 17:15:22 CST; 1h 0min ago
     CGroup: /
     
     ...
  ---
  
  [root@cyberpelican init.d]# systemctl status httpd
  ● httpd.service - The Apache HTTP Server
     Loaded: loaded (/usr/lib/systemd/system/httpd.service; enabled; vendor preset: disabled)
     Active: active (running) since Thu 2020-10-29 17:25:15 CST; 52min ago
       Docs: man:httpd(8)
             man:apachectl(8)
   Main PID: 1610 (httpd)
     Status: "Total requests: 0; Current requests/sec: 0; Current traffic:   0 B/sec"
      Tasks: 7
     CGroup: /system.slice/httpd.service
             ├─1610 /usr/sbin/httpd -DFOREGROUND
             ├─2208 /usr/sbin/httpd -DFOREGROUND
             ├─2209 /usr/sbin/httpd -DFOREGROUND
             ├─2210 /usr/sbin/httpd -DFOREGROUND
             ├─2211 /usr/sbin/httpd -DFOREGROUND
             ├─2212 /usr/sbin/httpd -DFOREGROUND
             └─2213 /usr/sbin/httpd -DFOREGROUND
  
  Oct 29 17:24:49 cyberpelican systemd[1]: Starting The Apache HTTP Server...
  Oct 29 17:25:05 cyberpelican httpd[1610]: AH00558: httpd: Could not reliably determine the server's fully qualified domain name, using fe... message
  Oct 29 17:25:15 cyberpelican systemd[1]: Started The Apache HTTP Server.
  Hint: Some lines were ellipsized, use -l to show in full.
  
  
  
  
  ```

  |           |                                                     |
  | --------- | --------------------------------------------------- |
  | Loaded    | 配置文件的位置；是否开机启动                        |
  | Active    | 当前状态 active\|inactive\|activing；状态切换的时间 |
  | Docs      | 帮助文档                                            |
  | Main PID  | 主进程ID                                            |
  | status    | 应用软件提供的状态                                  |
  | CGroup    | 所有子进程                                          |
  | log block | 应用在systemd中的日志                               |

- `systemctl start <unit>`

  启动服务

  ```
  [root@cyberpelican local]# systemctl start httpd
  可以一次启动多个服务
  [root@cyberpelican local]# systemctl restart zabbix-server zabbix-agent httpd
  ```

- `systemctl enable <unit>`

  开机启动服务

  对于那些支持 Systemd 的软件，安装的时候，会自动在`/usr/lib/systemd/system`目录添加一个配置文件。

  ```
  systemctl enable httpd
  等价于
   ln -s '/usr/lib/systemd/system/httpd.service' '/etc/systemd/system/multi-user.target.wants/httpd.service'
  ```

  ==上面的命令相当于在`/etc/systemd/system`目录添加一个符号链接，指向`/usr/lib/systemd/system`里面的`httpd.service`文件。===

  > 这是因为开机时，`Systemd`只执行`/etc/systemd/system`目录里面的配置文件。这也意味着，如果把修改后的配置文件放在该目录，就可以达到覆盖原始配置的效果

- `systemctl disable <unit>`

  取消开机启动服务，撤销两个目录之间的链接

  ```
  systemctl disable httpd
  ```

- `systemctl start <unit>`

  启动服务

  ```
  [root@cyberpelican init.d]# systemctl start httpd
  ```

- `systemctl stop <unit>`

  停止服务，如果无法终止进程使用`systemctl kill`

  ```
  [root@cyberpelican init.d]# systemctl stop httpd
  ```

- `systemctl retart <unit>`

  ```
  [root@cyberpelican init.d]# systemctl restart httpd
  ```

- `systemctl kill httpd`

  ```
  [root@cyberpelican init.d]# systemctl kill httpd
  [root@cyberpelican init.d]# ps -ef|grep httpd
  root       4681   3674  0 18:55 pts/1    00:00:00 grep --color=auto httpd
  ```

## 配置文件

配置文件主要放在`/usr/lib/systemd/system`目录，也可能在`/etc/systemd/system`目录。找到配置文件以后，使用文本编辑器打开即可。==使用`systemctl cat <unit>`，查看配置文件的内容==

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

==注意，键值对的等号两侧不能有空格。==

### 配置文件的区块

#### Unit

`[Unit]`区块通常是配置文件的第一个区块，用来定义 Unit 的元数据，以及配置与其他 Unit 的关系。它的主要字段如下。

```
Description：简短描述
Documentation：文档地址
Requires：当前 Unit 依赖的其他 Unit，如果它们没有运行，当前 Unit 会启动失败
Wants：与当前 Unit 配合的其他 Unit，如果它们没有运行，当前 Unit 不会启动失败
BindsTo：与Requires类似，它指定的 Unit 如果退出，会导致当前 Unit 停止运行
Before：如果该字段指定的 Unit 也要启动，那么必须在当前 Unit 之后启动
After：如果该字段指定的 Unit 也要启动，那么必须在当前 Unit 之前启动
Conflicts：这里指定的 Unit 不能与当前 Unit 同时运行
Condition...：当前 Unit 运行必须满足的条件，否则不会运行
Assert...：当前 Unit 运行必须满足的条件，否则会报启动失败
```

#### Service

`[Service]`区块用来 Service 的配置，只有 Service 类型的 Unit 才有这个区块。它的主要字段如下。

```
Type：定义启动时的进程行为。它有以下几种值。
Type=simple：默认值，执行ExecStart指定的命令，启动主进程
Type=forking：以 fork 方式从父进程创建子进程，创建后父进程会立即退出
Type=oneshot：一次性进程，Systemd 会等当前服务退出，再继续往下执行
Type=dbus：当前服务通过D-Bus启动
Type=notify：当前服务启动完毕，会通知Systemd，再继续往下执行
Type=idle：若有其他任务执行完毕，当前服务才会运行
ExecStart：启动当前服务的命令
ExecStartPre：启动当前服务之前执行的命令
ExecStartPost：启动当前服务之后执行的命令
ExecReload：重启当前服务时执行的命令
ExecStop：停止当前服务时执行的命令
ExecStopPost：停止当其服务之后执行的命令
RestartSec：自动重启当前服务间隔的秒数
Restart：定义何种情况 Systemd 会自动重启当前服务，可能的值包括always（总是重启）、on-success、on-failure、on-abnormal、on-abort、on-watchdog
TimeoutSec：定义 Systemd 停止当前服务之前等待的秒数
Environment：指定环境变量
```

#### Install

`[Install]`通常是配置文件的最后一个区块，用来定义如何启动，以及是否开机启动。它的主要字段如下。

```
WantedBy：它的值是一个或多个 Target，当前 Unit 激活时（enable）符号链接会放入/etc/systemd/system目录下面以 Target 名 + .wants后缀构成的子目录中
RequiredBy：它的值是一个或多个 Target，当前 Unit 激活时，符号链接会放入/etc/systemd/system目录下面以 Target 名 + .required后缀构成的子目录中
Alias：当前 Unit 可用于启动的别名
Also：当前 Unit 激活（enable）时，会被同时激活的其他 Unit
```

Unit 配置文件的完整字段清单，请参考[官方文档](https://www.freedesktop.org/software/systemd/man/systemd.unit.html)。

### 例子

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

**Unit**

- After

  表示sshd需要在`network.target `和 `sshd-keygen.service`启动之后启动

- Wants

  表示`sshd.service`与`sshd-keygen.service`之间存在"弱依赖"关系，即如果"sshd-keygen.service"启动失败或停止运行，不影响`sshd.service`继续执行。

- Requires

  表示"强依赖"关系，即如果该服务启动失败或异常退出，那么`sshd.service`也必须退出。

- EnvironmentFile

  指定当前服务的环境参数文件。该文件内部的`key=value`键值对，可以用`$key`的形式，在当前配置文件中获取。

  > 许多软件都有自己的环境参数文件，该文件可以用`EnvironmentFile`字段读取。

**Service**

- ExecStart

  定义启动进程时执行的命令。上面的例子中，启动`sshd`，执行的命令是`/usr/sbin/sshd -D $OPTIONS`，其中的变量`$OPTIONS`就来自`EnvironmentFile`字段指定的环境参数文件。与之作用相似的，还有如下这些字段。

  ```
  ExecReload字段：重启服务时执行的命令
  ExecStop字段：停止服务时执行的命令
  ExecStartPre字段：启动服务之前执行的命令
  ExecStartPost字段：启动服务之后执行的命令
  ExecStopPost字段：停止服务之后执行的命令
  ```

  请看下面的例子。

  ```shell
  [Service]
  ExecStart=/bin/echo execstart1
  ExecStart=
  ExecStart=/bin/echo execstart2
  ExecStartPost=/bin/echo post1
  ExecStartPost=/bin/echo post2
  ```

  上面这个配置文件，第二行`ExecStart`设为空值，等于取消了第一行的设置，运行结果如下。

  ```
  execstart2
  post1
  post2
  ```

  所有的启动设置之前，都可以加上一个连词号（`-`），表示"抑制错误"，即发生错误的时候，不影响其他命令的执行。比如，`EnvironmentFile=-/etc/sysconfig/sshd`（注意等号后面的那个连词号），就表示即使`/etc/sysconfig/sshd`文件不存在，也不会抛出错误。

- Type

  定义启动类型。它可以设置的值如下。

  ```
  simple（默认值）：ExecStart字段启动的进程为主进程
  forking：ExecStart字段将以fork()方式启动，此时父进程将会退出，子进程将成为主进程
  oneshot：类似于simple，但只执行一次，Systemd 会等它执行完，才启动其他服务
  dbus：类似于simple，但会等待 D-Bus 信号后启动
  notify：类似于simple，启动结束后会发出通知信号，然后 Systemd 再启动其他服务
  idle：类似于simple，但是要等到其他任务都执行完，才会启动该服务。一种使用场合是为让该服务的输出，不与其他服务的输出相混合
  ```

  下面是一个`oneshot`的例子，笔记本电脑启动时，要把触摸板关掉，配置文件可以这样写。

  ```shell
  [Unit]
  Description=Switch-off Touchpad
  
  [Service]
  Type=oneshot
  ExecStart=/usr/bin/touchpad-off
  
  [Install]
  WantedBy=multi-user.target
  ```

  上面的配置文件，启动类型设为`oneshot`，就表明这个服务只要运行一次就够了，不需要长期运行。

  如果关闭以后，将来某个时候还想打开，配置文件修改如下。

  ```shell
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

  上面配置文件中，`RemainAfterExit`字段设为`yes`，表示进程退出以后，服务仍然保持执行。这样的话，一旦使用`systemctl stop`命令停止服务，`ExecStop`指定的命令就会执行，从而重新开启触摸板。

- KillMode

  定义 Systemd 如何停止 sshd 服务。

  上面这个例子中，将`KillMode`设为`process`，表示只停止主进程，不停止任何sshd 子进程，即子进程打开的 SSH session 仍然保持连接。这个设置不太常见，但对 sshd 很重要，否则你停止服务的时候，会连自己打开的 SSH session 一起杀掉。

  `KillMode`字段可以设置的值如下。

  ```
  control-group（默认值）：当前控制组里面的所有子进程，都会被杀掉
  process：只杀主进程
  mixed：主进程将收到 SIGTERM 信号，子进程收到 SIGKILL 信号
  none：没有进程会被杀掉，只是执行服务的 stop 命令。
  ```

- Restart

  定义了 sshd 退出后，Systemd 的重启方式。

  上面的例子中，`Restart`设为`on-failure`，表示任何意外的失败，就将重启sshd。如果 sshd 正常停止（比如执行`systemctl stop`命令），它就不会重启。

  `Restart`字段可以设置的值如下。

  ```
  no（默认值）：退出后不会重启
  on-success：只有正常退出时（退出状态码为0），才会重启
  on-failure：非正常退出时（退出状态码非0），包括被信号终止和超时，才会重启
  on-abnormal：只有被信号终止和超时，才会重启
  on-abort：只有在收到没有捕捉到的信号终止时，才会重启
  on-watchdog：超时退出，才会重启
  always：不管是什么退出原因，总是重启
  ```

  对于守护进程，推荐设为`on-failure`。对于那些允许发生错误退出的服务，可以设为`on-abnormal`。

- RestartSec

  表示 Systemd 重启服务之前，需要等待的秒数。上面的例子设为等待42秒。

**install**

- WantedBy

  表示该服务所在的 Target。`Target`的含义是服务组，表示一组服务。`WantedBy=multi-user.target`指的是，==sshd 所在的 Target 是`multi-user.target`。==

  这个设置非常重要，因为执行`systemctl enable sshd.service`命令时，`sshd.service`的一个符号链接，就会放在`/etc/systemd/system`目录下面的`multi-user.target.wants`子目录之中。(开机启动的服务)
