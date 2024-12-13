## Digest
syntax：`systemctl [options] command [name]`
用于管理 systemd 的控制程序
### notes

1. 如果 unit name 没有指定 suffix，默认使用`.service`，例如：

   `systemctl start sshd`和`systemctl start sshd.service`等价，`systemctl isolate default`和 `systemctl isolate default.target`等价

2. 可以使用 shell glob，具体规则查看 glob man page

## Optional args

- `--version`

  查看当前systemd的版本

- `-t | --type=`

  按照 unit type 查看  unit，可以使用`systemctl -t help`查看所有可以使用的type

- `--state`

  按照 unit status 查看unit，可以使用`systemctl --state=failed`查看所有failed units

- `-a | --all`

  显示所有状态的units

- `--after`

  和`list-dependencies`一起使用，即显示`After=`directive的内容

- `--before`

  和`list-dependencies`一起使用，即显示`Before=`directive的内容

- `--show-types`

  显示可用的socket types

- `-n`

  和`status`一起使用，controls the number of journal lines to show，defaults to 10

- `-o | --output=`

  和`status`一起使用，controls the fomatting of the journal entries that are shown，具体值查`journalctl`

## Positional args
### Target management

- `add-wants TARGET NAME..., add-requires TARGET NAME`

  为指定target 在`Wants=`和`Requires=`指令块中添加 unit

- `get-default`

  return the default target to boot into. This returns the target unit name default.target is symlinked to

  ```
  [root@centos7 system]# systemctl get-default 
  multi-user.target
  ```

- `set-default NAME`

  Set the default target to boot into. This sets (symlinks) the default.target alias to the given target unit.

  ```
  [root@centos7 system]# systemctl set-default graphical.target 
  Removed symlink /etc/systemd/system/default.target.
  Created symlink from /etc/systemd/system/default.target to /usr/lib/systemd/system/graphical.target.
  [root@centos7 system]# systemctl get-default 
  graphical.target
  ```

- `isolate NAME`

  start the unit specified on the command line and its dependencies and stop all others. If a unit name with no extension is given, an extension of ".target" will be assumed. 切换系统当前运行的 target 等级

### Unit managment

- `status [PATTERN...|PID...]`

  查看unit的状态，默认只显示10行日志(顺序和`journalctl`默认显示的一样)，可以使用`--lines`和`--full`指定。详细内容可以查看`journalctl -u=NAME`来查看。不会显示 unit 是否 enabled，可以使用`systemctl is-enabled`来校验

  ```
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

|  |  |
| --- | --- |
| Loaded | 配置文件的位置；是否开机启动 |
| Active | 当前状态 active&#124;inactive&#124;activing；状态切换的时间 |
| Docs | 帮助文档 |
| Main PID | 主进程ID |
| status | 应用软件提供的状态 |
| CGroup | 所有子进程 |
| log block | 应用在systemd中的日志 |

​	如果参数是PID，查看进程关联的unit状态

```
[root@centos7 system]# ps -ef | grep -v grep | grep sshd
root       661     1  0 Feb22 ?        00:00:00 /usr/sbin/sshd -D -u0
root     22750   661  0 Feb25 ?        00:00:00 sshd: vagrant [priv]
vagrant  22753 22750  0 Feb25 ?        00:00:01 sshd: vagrant@pts/1
root     22907   661  0 Feb25 ?        00:00:00 sshd: vagrant [priv]
vagrant  22910 22907  0 Feb25 ?        00:00:01 sshd: vagrant@pts/2
[root@centos7 system]# systemctl status 661
● sshd.service - OpenSSH server daemon
   Loaded: loaded (/usr/lib/systemd/system/sshd.service; enabled; vendor preset: enabled)
   Active: active (running) since Tue 2022-02-22 07:27:26 UTC; 1 weeks 0 days ago
```

- `start PATTERN...`

  start one or more units specified on the command line

  ```
  [root@cyberpelican local]# systemctl start httpd
  可以一次启动多个服务
  [root@cyberpelican local]# systemctl restart zabbix-server zabbix-agent httpd
  ```

- `enable NAME...`

  会为 unit files 中的 Install directive 生成 symlink，symlink 生成后 systemd configuration 会重载(等价于`daemon-reload`)。symlink 会出现在 target 开机自启的目录下，即表示当前unit会在指定target下自启

  ```
  systemctl enable httpd
  等价于
  ln -s '/usr/lib/systemd/system/httpd.service' '/etc/systemd/system/multi-user.target.wants/httpd.service'
  ```

  `Systemd`只执行`/etc/systemd/system`目录里面的配置文件。这也意味着，如果把修改后的配置文件放在该目录，就可以达到覆盖原始配置的效果（hacking on it）。
  可以结合`--now`表示`enable`的同时执行`start`

  ```
  [root@centos7 system]# systemctl status sshd
  ● sshd.service - OpenSSH server daemon
     Loaded: loaded (/usr/lib/systemd/system/sshd.service; disabled; vendor preset: enabled)
     Active: inactive (dead)
  
  [root@centos7 system]# systemctl enable sshd --now
  Created symlink from /etc/systemd/system/multi-user.target.wants/sshd.service to /usr/lib/systemd/system/sshd.service.
  
  [root@centos7 system]# systemctl status sshd
  ● sshd.service - OpenSSH server daemon
     Loaded: loaded (/usr/lib/systemd/system/sshd.service; enabled; vendor preset: enabled)
     Active: active (running) since Tue 2022-03-01 20:59:49 UTC; 4s ago
  ```

  如果 unit 被 masked，在使用 enable 时就会报错、

  ```
  [root@centos7 system]# systemctl mask sshd
  Created symlink from /etc/systemd/system/sshd.service to /dev/null.
  [root@centos7 system]# systemctl enable sshd
  Failed to execute operation: Cannot send after transport endpoint shutdown
  ```

- `disable NAME...`

  和`enable`相反表示开启不自启，如果使用`--now`表示`disable`时也执行`stop`

- `reenable NAME...`

  `disable`和 `enable`结合

- `is-enable NAME...`

  校验 unit 是否 enabled

  ```
  [root@centos7 system]# systemctl is-enabled sshd
  enabled
  ```

  有如下几种状态

  ```
  ┌──────────────────┬──────────────────────────────────┬──────────────┐
  │Printed string    │ Meaning                          │ Return value │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"enabled"         │ Enabled through a symlink in     │              │
  ├──────────────────┤ .wants directory (permanently or │ 0            │
  │"enabled-runtime" │ just in /run).                   │              │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"linked"          │ Made available through a symlink │              │
  ├──────────────────┤ to the unit file (permanently or │ 1            │
  │"linked-runtime"  │ just in /run).                   │              │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"masked"          │ Disabled entirely (permanently   │              │
  ├──────────────────┤ or just in /run).                │ 1            │
  │"masked-runtime"  │                                  │              │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"static"          │ Unit file is not enabled, and    │ 0            │
  │                  │ has no provisions for enabling   │              │
  │                  │ in the "[Install]" section.      │              │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"indirect"        │ Unit file itself is not enabled, │ 0            │
  │                  │ but it has a non-empty Also=     │              │
  │                  │ setting in the "[Install]"       │              │
  │                  │ section, listing other unit      │              │
  │                  │ files that might be enabled.     │              │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"disabled"        │ Unit file is not enabled.        │ 1            │
  ├──────────────────┼──────────────────────────────────┼──────────────┤
  │"bad"             │ Unit file is invalid or another  │ > 0          │
  │                  │ error occured. Note that         │              │
  │                  │ is-enabled wil not actually      │              │
  │                  │ return this state, but print an  │              │
  │                  │ error message instead. However   │              │
  │                  │ the unit file listing printed by │              │
  │                  │ list-unit-files might show it.   │              │
  └──────────────────┴──────────────────────────────────┴──────────────┘
  ```

- `stop PATTERN...`

  stop one or more untis specified on the command line

- `reload PATTERN...`

  reload the service-specific configuration, not the unit configuration file of systemd. For the example case of Apache

- `restart PATTERN...`

  restart one or more units specified on the command line

- `mask NAME...`

  将 unit file 链接指向 /dev/null，即不能启动该 unit (手动和自启都一样)，可以使用`--runtime`表示只在本次生效，重启后取消mask，也可以使用`--now`mask 的同时也 stop

  ```
  [root@cyberpelican init.d]# systemctl mask httpd
  Created symlink from /etc/systemd/system/httpd.service to /dev/null.
  ```

- `unmask NAME...`
  this will undo the effect of `mask`

- `kill PATTERN...`

  当stop无法停止unit，可以使用该参数。`--kill-who=`、`--signal=`指定需要kill掉的进程和发送的SIG

- `is-active PATTERN...`

  check whether any of the specified units are active

- `is-failed PATTERN...`

  check whether any of the specified units are in a failed state

- `show [PATTERN...|JOB...]`

  查看unit的属性。读取的配置、执行的命令、PID等

- `cat PATTERN...`

  查看 unit 对应的 unit file

- `list-dependencies [NAME]`

  TODO
  shows units required and wanted by the specified unit，可以和`--after`、`--before`一起使用过滤

- `list-unit-files [PATTERN...]`

  查看unit 对应的 unit file 以及 unit 状态，可以使用`-t`参数指定配置文件类型 ，使用`systemctl`时的默认command

  ```
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

  注意，从配置文件的状态无法看出，该 Unit 是否正在运行。这必须执行前面提到的`systemctl status`命令

| enable | 已建立启动链接，开机自动启动 |
| --- | --- |
| disable | 没有建立链接 |
| static | 该配置文件没有[Install]部分（无法执行），只能作为其他配置文件的依赖 |
| masked | 该配置文件被禁止建立启动链接，无法被启动 |

- `edit NAME...`

  修改指定的unit file

  1. `--full`表示拷贝原始的unit file，在这基础上修改

  2. `--runtime`表示只在本次生效，重启后失效

### Query

- `list-uints [PATTERN]`

  根据 pattern 显示所有的 unit 内容，通常结合`--state`和`-t`一起使用

  ```
  [root@cyberpelican ~]# systemctl list-units|grep httpd
    httpd.service                                                                                                    loaded active running   The Apache HTTP Server
    
  [root@cyberpelican ~]# systemctl list-units --type=service --state=active
  ```

- `list-socket [PATTERN]`

  根据 pattern 显示所有的 socket units 内容

- `help PATTERN...|PID...`

  查看unit对应的man page

### Machines Mangement

- `list-machines [PATTERN...]`

  显示本机运行的container (实测不会关联docker)

### Jobs Management

- `list-jobs [PATTERN...]`

  显示当前运行的jobs，等价于 built-in `jobs`

- `cancel JOB...`

  取消 jobs ，需要指定 job id，如果没有指定 job id，取消所有

### System Commands

- `is-system-running`

  校验当前OS是否是 operational，OS可以是如下几种状态

  ```
  ┌─────────────┬───────────────────────────────────────────┐
  │Name         │ Description                               │
  ├─────────────┼───────────────────────────────────────────┤
  │initializing │ Early bootup, before basic.target is      │
  │             │ reached or the maintenance state entered. │
  ├─────────────┼───────────────────────────────────────────┤
  │starting     │ Late bootup, before the job queue becomes │
  │             │ idle for the first time, or one of the    │
  │             │ rescue targets are reached.               │
  ├─────────────┼───────────────────────────────────────────┤
  │running      │ The system is fully operational.          │
  ├─────────────┼───────────────────────────────────────────┤
  │degraded     │ The system is operational but one or more │
  │             │ units failed.                             │
  ├─────────────┼───────────────────────────────────────────┤
  │maintenance  │ The rescue or emergency target is active. │
  ├─────────────┼───────────────────────────────────────────┤
  │stopping     │ The manager is shutting down.             │
  └─────────────┴───────────────────────────────────────────┘
  ```

- `default`

  等价于`isolate default.target`

- `rescue`

  等价于`isolate rescue.target`

- `emergency`

  等价于`isolate emergency.target`

- `halt`

  等价于执行 halt

- `suspend`

  等价执行suspend 锁屏

- `poweroff`

  等价于执行poweroff

- `rboot`

  等价于reboot
