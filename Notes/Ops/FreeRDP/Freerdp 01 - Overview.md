---
createTime: 2025-04-10 18:30
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Freerdp 01 - Overview

## 0x01 Preface

FreeRDP 是 (Remote Desktop Protocol)RDP 的开源实现，支持跨平台

## 0x02 Installation

直接用包管理器安装即可

这里需要注意的是 `extra/freerdp` 对应的 binary 是 `{x|wl}freerdp3` 而不是 `{x|wl}freerdp`

```
sudo pacman -S freerdp

pacman -Ql freerdp | grep bin
freerdp /usr/bin/
freerdp /usr/bin/freerdp-proxy3
freerdp /usr/bin/freerdp-shadow-cli3
freerdp /usr/bin/sdl-freerdp3
freerdp /usr/bin/sfreerdp-server3
freerdp /usr/bin/sfreerdp3
freerdp /usr/bin/winpr-hash3
freerdp /usr/bin/winpr-makecert3
freerdp /usr/bin/wlfreerdp3
freerdp /usr/bin/xfreerdp3
```

`{x|wl}freerdp` 对应 `extra/freerdp2`，一些软件会依赖 `extra/freerdp2` 两者的一些参数不通用，freerdp2 已经停止更新不建议使用

```
sudo pacman -R freerdp2
checking dependencies...
error: failed to prepare transaction (could not satisfy dependencies)
:: removing freerdp2 breaks dependency 'freerdp2' required by hydra
:: removing freerdp2 breaks dependency 'libfreerdp2.so=2-64' required by hydra
:: removing freerdp2 breaks dependency 'libwinpr2.so=2-64' required by hydra
```

为了方便调用建议创建 alias

```
alias xfreerdp='xfreerdp3'
alias wlfreerdp='wlfreerdp3'
```

## 0x03 Syntax[^1]

freerdp 支持 2 种格式的命令行

- windows-style syntax
- posix-style syntax

syntax 不能混用

更推荐 windows-style syntax 因为简洁，且在 freerdp 中会经常使用 toggle option  

### 0x03a Windows-style Syntax

```
/flag (enables flag)
/option:<value> (specifies option with value)
+toggle -toggle (enables or disables toggle, where '/' is a synonym of '+')
```

使用 `/option` 和 `+option` 来开启一个配置项，例如 `/wallpaper` 等价于 `+wallpaper`

```
xfreerdp /f /bpp:32 /v:rdp.contoso.com +wallpaper -themes
```

### 0x03b Posix-style Syntax

```
--flag (enables flag)
--option:<value> (specifies option with value)
--enable-toggle --disable-toggle (enables or disables toggle)
```

使用 `--option` 或者 `--enable-option` 来开启一个配置项，例如 `--wallpaper` 等价于 `--enable-wallaper`

```
xfreerdp -f --bpp 32 -v rdp.contoso.com --enable-wallpaper --disable-themes
```

## 0x04 Authentication

RDP 连接通常需要提供如下几种信息

```
- /u:user 用户名
- /d:domain 域控(可选)
- /p:password 密码(可选)
- /v:server 服务器地址
```

一个简单的例子如下

```
xfreerdp /u:JohnDoe /d:CONTOSO /p:Password123! /v:rdp.contoso.com
```

`/u` 和 `/d` 可以单独赋值，也可以组合赋值

```
xfreerdp /u:CONTOSO\JohnDoe /p:Password123! /v:rdp.contoso.com
```

虽然这种方式更加简洁，但是 `\` 在 Shell 通常是转义符，所以需要 转义 或者 quoted

```
xfreerdp /u:CONTOSO\\JohnDoe /p:Password123! /v:rdp.contoso.com
xfreerdp "/u:CONTOSO\JohnDoe" /p:Password123! /v:rdp.contoso.com
```

也可以使用 the User Principal Name (UPN) notation

```
xfreerdp /u:JohnDoe@CONTOSO /p:Password123! /v:rdp.contoso.com
```

但是如果用户名包含 `@` 就会将后面的部分识别为 domain，这时需要使用如下格式

```
xfreerdp /u:\john.doe@live.com /p:Password123! /v:rdp.contoso.com
```

除了账号外，freerdp 会自动将密码隐藏

```
awake@workstation:~$ ps aux | grep freerdp
awake           22506   0.0  0.1  2502620  10236 s002  S+   11:10pm   0:01.00 xfreerdp /u:JohnDoe /p:************ /d:CONTOSO /v:rdp.contoso.com
```

## 0x05 freerdp2 VS freerdp3

freerdp3 对 freerdp2 的一些参数做了修改[^2]，例如

- `/tls-seclevel:1` 替换成了 `/tls:seclevel:1`
- `/bitmap-cache` 替换成了 `/cache:bitmap:on`

具体可以查看 manual page

## 0x05 wlfreerdp

freerdp 虽然支持 wayland，但是 wlfreerdp 目前还有很多问题，所以更加推荐使用 xfreerdp

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [FAQ · FreeRDP/FreeRDP Wiki · GitHub](https://github.com/FreeRDP/FreeRDP/wiki/FAQ)
- [FreeRDP-Manuals/Configuration/FreeRDP-Configuration-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/Configuration/FreeRDP-Configuration-Manual.markdown)
- [FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown)

***References***

[^1]:[FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown#syntax)
[^2]:[FAQ · FreeRDP/FreeRDP Wiki · GitHub](https://github.com/FreeRDP/FreeRDP/wiki/FAQ#windows-7-errconnect_tls_connect_failed)
