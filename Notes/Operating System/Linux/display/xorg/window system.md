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

