# Target

### 概述

> 与传统方式中的RunLevel相似，由于systemd取代了init（SysV），所以chkconfig命令也应该取消使用

启动计算机的时候，需要启动大量的 Unit。如果每一次启动，都要一一写明本次启动需要哪些 Unit，显然非常不方便。Systemd 的解决方案就是 Target。

简单说，Target 就是一个 Unit 组，包含许多相关的 Unit 。==启动某个 Target 的时候，Systemd 就会启动里面所有的 Unit。==从这个意义上说，Target 这个概念类似于"状态点"，启动某个 Target 就好比启动到某种状态。

传统的`init`启动模式里面，有 RunLevel 的概念，跟 Target 的作用很类似。==不同的是，RunLevel 是互斥的，不可能多个 RunLevel 同时启动，但是多个 Target 可以同时启动(意味着可以，即是graphical.target 也是multi-user.target)。==

### "SysV 运行级别" 与 "systemd 目标" 对照表

| SysV 运行级别 |                     Systemd 目标                      |                            注释                             |
| :-----------: | :---------------------------------------------------: | :---------------------------------------------------------: |
|       0       |           runlevel0.target, poweroff.target           |                      中断系统（halt）                       |
| 1, s, single  |            runlevel1.target, rescue.target            |                         单用户模式                          |
|     2, 4      | runlevel2.target, runlevel4.target, multi-user.target |            用户自定义运行级别，通常识别为级别3。            |
|       3       |          runlevel3.target, multi-user.target          |      多用户，无图形界面。用户可以通过终端或网络登录。       |
|       5       |          runlevel5.target, graphical.target           | 多用户，图形界面。==继承级别3的服务，并启动图形界面服务。== |
|       6       |            runlevel6.target, reboot.target            |                            重启                             |
|   emergency   |                   emergency.target                    |                 急救模式（Emergency shell）                 |

- `systemctl list-unit-files -t type`

  列出所有的target

  ```
  [root@cyberpelican multi-user.target.wants]# systemctl list-unit-files -t target
  UNIT FILE                  STATE   
  anaconda.target            static  
  basic.target               static  
  bluetooth.target           static  
  cryptsetup-pre.target      static  
  cryptsetup.target          static  
  ctrl-alt-del.target        disabled
  default.target             enabled 
  emergency.target           static  
  final.target               static  
  getty-pre.target           static  
  getty.target               static  
  graphical.target           enabled 
  ```

- `systemctl get-default`

  ==获取默认启动的target。这个组内的所有服务，都将开机启动，关联配置文件`WantedBy`字段==

  ```
  [root@cyberpelican system]# systemctl get-default 
  graphical.target
  ```

  上面的结果表示，默认的启动 Target 是`multi-user.target`。在这个组里的所有服务，都将开机启动。这就是为什么`systemctl enable`命令能设置开机启动的原因。

- `systemctl list-denpendencies multi-user.target`

  列出指定target包含的所有服务

  ```
  [root@cyberpelican system]# systemctl list-dependencies multi-user.target
  multi-user.target
  ● ├─abrt-ccpp.service
  ● ├─abrt-oops.service
  ● ├─abrt-vmcore.service
  ● ├─abrt-xorg.service
  ● ├─abrtd.service
  ● ├─atd.service
  ```

- `systemctl isolate rescue.target`

  直接切换target，终止其他所有非指定target的服务，`.target`可以被省略

  ```
  [root@cyberpelican system]# systemctl isolate rescue
  ```

- `systemctl set-default rescue.target`

  设置启动默认的target，target可以被省略

  ```
  [root@cyberpelican ~]# systemctl set-default multi-user
  Removed symlink /etc/systemd/system/default.target.
  Created symlink from /etc/systemd/system/default.target to /usr/lib/systemd/system/multi-user.target.
  ```

### Target配置文件

`systemctl cat <target>`查看配置文件

```
[root@cyberpelican ~]# systemctl cat multi-user.target 
# /usr/lib/systemd/system/multi-user.target
#  This file is part of systemd.
#
#  systemd is free software; you can redistribute it and/or modify it
#  under the terms of the GNU Lesser General Public License as published
#  the Free Software Foundation; either version 2.1 of the License, or
#  (at your option) any later version.

[Unit]
Description=Multi-User System
Documentation=man:systemd.special(7)
Requires=basic.target
Conflicts=rescue.service rescue.target
After=basic.target rescue.service rescue.target
AllowIsolate=yes
```

注意，Target 配置文件里面没有启动命令。

- Requires：要求与`basic.target`一起运行
- conflicts：冲突字段。如果`rescue.service`或`rescue.target`正在运行，`multi-user.target`就不能运行，反之亦然。
- After：表示`multi-user.target`在`basic.target` 、 `rescue.service`、 `rescue.target`之后启动，如果它们有启动的话。
- AllowIsolate：允许使用`systemctl isolate`命令切换到`multi-user.target`。

> 如果修改了配置文件，需要使用`systemctl daemon-reload`使配置文件重新加载

### 与SysV的差别 

1. **默认的 RunLevel**（在`/etc/inittab`文件设置）现在被默认的 Target 取代，位置是`/etc/systemd/system/default.target`，通常符号链接到`graphical.target`（图形界面）或者`multi-user.target`（多用户命令行）。

   ```shell
   [root@cyberpelican system]# ll|grep default.target
   lrwxrwxrwx. 1 root root   36 Aug 24 07:59 default.target -> /lib/systemd/system/graphical.target
   
   ```

2. **启动脚本的位置**，以前是`/etc/init.d`目录，符号链接到不同的 RunLevel 目录 （比如`/etc/rc3.d`、`/etc/rc5.d`等），现在则存放在`/lib/systemd/system`和`/etc/systemd/system`目录。

3. **配置文件的位置**，以前`init`进程的配置文件是`/etc/inittab`，各种服务的配置文件存放在`/etc/sysconfig`目录。现在的配置文件主要存放在`/lib/systemd`目录，在`/etc/systemd`目录里面的修改可以覆盖原始设置。

   



