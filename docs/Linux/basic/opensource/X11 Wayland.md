# window system

参考：

https://en.wikipedia.org/wiki/Windowing_system

window system 主要由两个组件

- Display client：Any application that runs and presents its GUI in a window, is a client of the display server. 
- Display server：the display server being the mediator between the clients and the user（用于处理input）

Display server 和 Display client通过display server protocols(例如：X11或Wayland)

![](/home/cpl/note/imgs/_Linux/2021-07-05_23-12.png)

如果使用X11，display server 也会依赖与window manager(例如：kwin，或Mutter)

==但是Wayland直接通过wayland compositor==

## X11

参考：

https://en.wikipedia.org/wiki/X_Window_System

https://medium.com/mindorks/x-server-client-what-the-hell-305bd0dc857f

### 概述

==X window system==也被称为X(X11就是第11个版本)，为GUI提供基础的框架以及键盘和鼠标的交互，通常被用在UNIX操作系统上。

X有如下几个特点：

1. X使用CS模型，client 和 server可以在同一台主机上也可以不在，甚至可以是不同的arch和OS
2. ==X和常见的CS模型相反==。X server为application提供display和I/O services，application使用这些服务所以是X client
3. X可以通过加密隧道进行CS通信，例如SSH，==需要使用X forwarding参数==
4. X client可以模拟X server为其他的X client提供display服务，这被称为X nesting

![Screenshot from 2021-05-24 01-31-44](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Screenshot from 2021-05-24 01-31-44.zsikn5q9e1s.png)

**the X server receives input from a local keyboard and mouse and displays to a screen. A web browser and a terminal emulator run on the user's workstation and a terminal emulator runs on a remote computer but is controlled and monitored from the user's machine**

### X server

X server(常见的有x terminals简称xterm)可以和多种不同的client程序交互。X server为X client提供graphics resources并将input events 发送给X client

*meaning that the X server is usually running on the computer in front of a human user*

### X client

X client 是要求渲染图形化内容和接受输入的application，例如firefox,gnome

### remote SSH X11

为了远程使用X client application，流程如下

1. 连接sshd时使用ForwardX11

   ```
   ssh -Y host firefox
   ```

2. 请求本地的 display/input 服务
3. remote X client application连接到用户本地(请求建立连接的主机)的local X server，并为用户提供display/input服务

> 这里的X client application 是firefox，X server 是用户连接用的terminal

## xhost

安装`x11-xserver-utils`来添加xhost

xhost 用于将host names 或 user names 添加到允许连接到X server。

For workstations, this is the same machine as the server.  For X terminals, it is the login host.

使用`xhost +<Name>`表示添加，`xhost -<Name>`表示去除。==local machine可以被移除，只用通过reset才可以重新允许连接==

Name使用如下格式`famliy:name`，family可以是如下的值：

- inet：IPv4
- inet6：IPv6
- krb：Kerberos
- localf

name由`host@name`或`host`组成，例如：

```
xhost +inet:gns3vm:root
```

初始的access control list默认读取`/etc/X11.hosts`

## xinput



## Wayland

wayland compositor作为Wayland display server

![2021-07-05_23-39](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-07-05_23-39.4xl34l0bojk0.png)

## X11 vs Wayland

参考：https://www.secjuice.com/wayland-vs-xorg/

X11存在安全隐患，但是为了兼容大多数distro还是使用X11。Wayland虽然比较安全，但是存在兼容问题

## change X11 to Wayland

参考：https://itsfoss.com/switch-xorg-wayland/

可以通过`XDG_SESSION_TYPE`变量来查看使用的Display protocol

```
➜  system-connections echo $XDG_SESSION_TYPE
x11
```

