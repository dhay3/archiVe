# Systemd Digest

参考：

[https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html](https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)
[https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html](https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html)
[https://zh.wikipedia.org/wiki/Systemd](https://zh.wikipedia.org/wiki/Systemd)
[https://wiki.archlinux.org/index.php/systemd_(简体中文)](https://wiki.archlinux.org/index.php/systemd_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87))
[http://0pointer.de/blog/projects/systemd.html](http://0pointer.de/blog/projects/systemd.html)

## Digset

systemd is a system and service manager for linux OS，以 PID 1 运行和 init system 的作用一样用于启动和维护 userspace services
Systemd 的优点是功能强大，使用方便，缺点是体系庞大，非常复杂。事实上，现在还有很多人反对使用 Systemd，理由就是它过于复杂，与操作系统的其他部分强耦合，违反"keep simple, keep stupid"

![Snipaste_2020-10-29_12-45-27](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220316/Snipaste_2020-10-29_12-45-27.5l9vtfsxre00.webp)
systemd 在启动时会执行多种任务，会配置hostname、loopback network device，同时也会设置多种系统参数(/sys、/proc目录下的)
在 boot 过程中 systemd 会启用 default.target，用于激活 ob-boot services and other on-boot units. 通常是一个 symlink，指向其他实际的 target unit( graphical.target、multi-use.target、etc ). 具体查看`systemd.special`
## SysV

历史上，[Linux 的启动](http://www.ruanyifeng.com/blog/2013/08/linux_boot_process.html)一直采用`[init](https://en.wikipedia.org/wiki/Init)`进程。

```shell
[root@cyberpelican ~]# /etc/init.d/apache2 start
or
[root@cyberpelican ~]# service apache2 start
```

这种方法有两个缺点。

1.  启动时间长。`init`进程是串行启动，只有前一个进程启动完，才会启动下一个进程。 
1.  启动脚本复杂。`init`进程只是执行启动脚本，不管其他事情。脚本需要自己处理各种情况，这往往使得脚本变得很长。  

Systemd 就是为了解决这些问题而诞生的。它的设计目标是，为系统的启动和管理提供一套完整的解决方案。
systemd is compatibel with the SysV init system to a large degree

1. SysV init scripts are supported and simply read as an alternative（though limited） configuration file format
1. The SysV /dev/initctl( init 命令 ) is provided, and compatibility imlementations of the various SysV client tools are available.
1. various established Unix functionality such as /etc/fstab or the utmp database are supported

```
#网络任然可以通过该方式启动
[root@cyberpelican ~]# cd /etc/init.d/
[root@cyberpelican init.d]# ls
functions  netconsole  network  README
```

根据Unix 惯例，字母是守护进程（daemon）的缩写`d`。 Systemd 这个名字的含义，就是它要守护整个系统。

使用了 Systemd，就不需要再用`init`了。Systemd 取代了，成为系统的第一个进程（PID 等于 1），其他进程都是它的子进程。`initd`

## Concepts
### status
units 有几种状态，根据不同的unit type 状态也会改变

1. active, meaning started
1. inactive, meaning stopped
1. actived, once was active
1. inactived, equal to inactive
1. fail, srvice failed in some way（process returned error code on exit, or crashed, an operation timed out, or fater too many restarts）
### Unit Files
用于描述unit，语法和 windows 的 `.ini`文件类似，具体查看 LInux Systemd Unit
#### unit type
systemd 通过 12 种不同的 unit 来管理, unit type 可以是如下几种

1. Service units, which start and control daemons and the processes they consist of. For details see `systemd.service`
1. Socket units, which encapsulate local IPC(interal process communication) or network sockets in the system, useful for socket-based activation. For details about socket units see `systemd.socket`, for details on socket-based activation and others forms of activation, see `daemon`
1. Target units are useful to group units, or provide well-know synchronization points during boot-up, see `systemd.target`
1. Device units expose kernel devices in systemd and may be used to implement device-based activation. For details see `systemd.device`
1. Mount units control mount points in the file system, for details see `systemd.mount`
1. Automunt units provide automount capbilities, for on-demand mounting of file systems as well as parallelized boot-up. see`systemd.automount`
1. Snapshot units can be used to emporarily save the state of the set of systemd units, which later may be restared by activating the saved snapshot unit. For more information see `systemd.snapshot`
1. Timer units are useful for triggering activation fo other units based on timers. You may find details in `sysmted.timer`
1. Swap units are very similart to mount units and encapsulate memory swap partitions or files of the operation system. They are described in `systemd.swap`
1. Path units may be used to activate other services when file systemd objects change or are modified. See `systemd.path`
1. Slice units may be used to group units which manage system processes (such as service and scop units) in a hierarchical tree for resource management purpose. See `systemd.slice`
1. Scope units are similar to service units, but manage foreign processes instead of starting them as well. See `systemd.scopej`
## Direcotories
### System unit directories
systemd 会从不同的目录读取unit配置，默认使用`pkg-config systemd --variable=systemdsystemconfdir`返回的目录
### SysV init scripts directory
根据不同的 distro来确认目录，如果systemd 不能找到对应的service 的 unit file，就会找对应的 SysV init script (same name without `.service` suffix)
### SysV runlevel link farm directory
根据不同的 distro来确认目录，用来决定service是否enable (自启)

