#  systemd 概述

参考：

https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html

https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html

https://zh.wikipedia.org/wiki/Systemd

https://wiki.archlinux.org/index.php/systemd_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)

## 概述

历史上，[Linux 的启动](http://www.ruanyifeng.com/blog/2013/08/linux_boot_process.html)一直采用[`init`](https://en.wikipedia.org/wiki/Init)进程。

```shell
[root@cyberpelican ~]# /etc/init.d/apache2 start
or
[root@cyberpelican ~]# service apache2 start
```

这种方法有两个缺点。

1. 启动时间长。`init`进程是串行启动，只有前一个进程启动完，才会启动下一个进程。

2. 启动脚本复杂。`init`进程只是执行启动脚本，不管其他事情。脚本需要自己处理各种情况，这往往使得脚本变得很长。

> 注意网络任然可以通过该方式启动
>
> ```
> [root@cyberpelican ~]# cd /etc/init.d/
> [root@cyberpelican init.d]# ls
> functions  netconsole  network  README
> ```

Systemd 就是为了解决这些问题而诞生的。它的设计目标是，为系统的启动和管理提供一套完整的解决方案。==兼容SysV==

根据Unix 惯例，==字母`d`是守护进程（daemon）的缩写==。 Systemd 这个名字的含义，就是它要守护整个系统。

使用了 Systemd，就不需要再用`init`了。==Systemd 取代了`initd`，成为系统的第一个进程（PID 等于 1），其他进程都是它的子进程。==

通过如下命令可以查看systemd的版本

```shell
[root@cyberpelican init.d]# systemctl --version
systemd 219
+PAM +AUDIT +SELINUX +IMA -APPARMOR +SMACK +SYSVINIT +UTMP +LIBCRYPTSETUP +GCRYPT +GNUTLS +ACL +XZ +LZ4 -SECCOMP +BLKID +ELFUTILS +KMOD +IDN
```

Systemd 的优点是功能强大，使用方便，缺点是体系庞大，非常复杂。事实上，现在还有很多人反对使用 Systemd，理由就是它过于复杂，与操作系统的其他部分强耦合，违反"keep simple, keep stupid"

<img src="..\..\..\..\imgs\_Linux\Snipaste_2020-10-29_12-45-27.png"/>

