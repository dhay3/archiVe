# FreeRDP

## 0x01 Overview

FreeRDP 是 RDP 的开源实现，支持跨平台

## 0x02 Installation

直接用包管理器安装即可

```
sudo pacman -S freerdp
```

## 0x03 xfreerdp/wlfreerdp

xfreerdp/wlfreerdp(wayland 下 icon 显示会有问题，需要通过 `/wm-class` 修改) 分别是 FreeRDP xorg/wayland 用户态工具(cli)，目前更推荐使用 xfreerdp，wlfreerdp 延迟感非常强，而且图标不正确

支持两种格式的命令行

- windows style syntax(默认)
- posix-style syntax

2 者等价(即 `+` 和 `\` 等价)，但是推荐使用 windows style syntax

例如

```
xfreerdp +u:admin  +p: +v:10.100.4.127 +dynamic-resolution
xfreerdp /u:admin  /p: /v:10.100.4.127 /dynamic-resolution
```

一个 RDP 连接需要提供如下几种信息

1. username

   例如 JohnDoe

2. domain (可以为空)

   例如 CONTOSO

3. password (可以为空)

   例如 Password123!

4. server-addess

   例如 rdp.contoso.com 默认端口 3389

在 xfreerdp 中对应到具体参数如下

```
xfreerdp /u:JohnDoe /d:CONTOSO /p:Password123! /v:rdp.contoso.com
```

### Optional args

> 只列举常用的参数，具体查看官方文档

- `/v:<server>[:port]`

  连接的地址

- `/u:[<domain>\\]<user> or <user>[@<domain>]`

  连接的账户和域

- `/p:<password>`

  账户对应的密码

- `/g:<gateway>[:port]`

  跳板机地址

- `/gu:[<domain>\\]<user> or <user>[@<domain>]`

  跳板机对应的账户

- `/gp:<password>`

  跳板机对应账户的密码

- `/dynamic-resolution`

  可以设置 RDP 框大小

- `/fonts`

  让字体更加流畅，默认不启用

- `/bpp:<depth>`

  多少 bit 位来表示一个像素点，值越高颜色越细腻，通用设置为 64 即可

- `/menu-anims`

  菜单的动画效果，默认不开启

- `/cert:tofu`

  连接时自动接受证书

- `/tls-seclevel:0`

  在碰到 ERRCONNECT_TLS_CONNECT_FAILED 时使用，使用 TLS1

- `/f`

  以全屏模式启动，使用 <kbd>Ctrl</kbd> + <kbd>Alt</kbd> + <kbd>Enter</kbd> 切换

- `/scale:[100|140|180]`

  屏幕缩放比例，只能是这几个值

- `/sacle-desktop:<percentage>`

  和 `/scale` 一起使用才会生效，表示

- `/w:<width>` | `/h:<height>`

  设置长宽分辨率分辨率

- `-decorations`

  不显示 window decoratoins，即无窗框

- `--compression-level:[0|1|2]`

  压缩的等级

- `/vedio`

  优化视频通道

## 0x04 Ungrab Window

freerdp 同时也支持类似 Virtualbox 中的 <kbd>Right-Ctrl</kbd> 临时切换到主机使用键盘[^2]，这样就可以使用 <kbd>Alt</kbd> + <kbd>TAB</kbd> 来切换来

在 freerdp 中使用 <kbd>Alt</kbd> + <kbd>Right-Ctrl</kbd> 中实现

目前 freerdp 中还有一个 bug 就是连续点击 2 次 <kbd>Right-Ctrl</kbd> 会导致 freerdp lose focus[^4]（ungrab），只能点击 freerdp 的 window boarder 才会重新 focus on freerdp。目前已经被修复~~，但是没有被整合到 mainline~~

## 0x05 HDPI Display

通常 HDPI 的屏幕显示会有问题，可以使用 `/w:1920 /h:1080` 来指定分辨率为 1080P，或者使用 `/scale:140 /scale-desktop:125` 放大 DPI(推荐使用该方法)

部分操作系统（例如 Windows7），缩放可能不会生效，可以使用 `/f /dynamic-resolution` 强制全屏然后再调整

## 0x06 Examples

```
/usr/bin/xfreerdp /u:admin /p: /v:10.100.4.127 /bpp:64 /fonts /cert:tofu /scale:140 /scale-desktop:125 /dynamic-resolution
```

**references**

[^1]:https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown
[^2]:https://github.com/FreeRDP/FreeRDP/issues/2585
[^3]:https://github.com/FreeRDP/FreeRDP/issues/5114
[^4]:https://github.com/FreeRDP/FreeRDP/issues/9959