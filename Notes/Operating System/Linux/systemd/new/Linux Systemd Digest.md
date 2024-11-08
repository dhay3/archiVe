# Linux Systemd Digest

## Overview

systemd is a system and service manager for linux OS，以 PID 1 运行和 init system 的作用一样用于启动和维护 userspace services
Systemd 的优点是功能强大，使用方便，缺点是体系庞大，非常复杂。事实上，现在还有很多人反对使用 Systemd，理由就是它过于复杂，与操作系统的其他部分强耦合，违反"keep simple, keep stupid"

![Snipaste_2020-10-29_12-45-27](https://github.com/dhay3/image-repo/raw/master/20220316/Snipaste_2020-10-29_12-45-27.5l9vtfsxre00.webp)
systemd 在启动时会执行多种任务，会配置hostname、loopback network device，同时也会设置多种系统参数(/sys、/proc目录下的)
在 boot 过程中 systemd 会启用 default.target，用于激活 ob-boot services and other on-boot units. 通常是一个 symlink，指向其他实际的 target unit( graphical.target、multi-use.target、etc ). 具体查看`systemd.special`



**references**

1. [^1]:https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html

2. [^2]:https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html

3. [^3]:https://zh.wikipedia.org/wiki/Systemd

4. [^4]:https://wiki.archlinux.org/index.php/systemd_(简体中文)

5. [^5]:http://0pointer.de/blog/projects/systemd.html