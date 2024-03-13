# FreeRDP

## 0x01 Overview

FreeRDP 是 RDP 的开源实现

## 0x02 Installation

直接用包管理器安装即可

```
sudo pacman -S freerdp
```

## 0x03 xfreerdp

xfreerdp 是 FreeRDP 用户态工具(cli)，支持两种格式的命令行

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

**references**

[^1]:https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown