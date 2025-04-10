---
createTime: 2025-04-10 22:57
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Freerdp 02 - xfreerdp

## 0x01 Preface

xfreerdp 是 freerdp 基于 X11 协议的实现

## 0x02 Options

> 只列举常用的参数，具体查看 manual page

### 0x02a Authentication Related Options

- `/v:<server>[:port]`

  RDP server 的地址，端口默认 3389

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

- `/cert:[ignore|tofu]`

  - ignore: 连接时忽略证书校验
  - tofu: trust on first use 连接时只有第一次需要手动信任证书

- `/tls:seclevel:0`

  兼容一些老的 RDP server 例如 Win7，在碰到 ERRCONNECT_TLS_CONNECT_FAILED 时使用[^1]

### 0x02b Vedio Related Options

- `-decorations`

  不显示 window decoratoins，即无窗框显示 freerdp

- `/w:<width>` | `/h:<height>`
	`/size:<widthxheight>`

  设置长宽分辨率

- `+dynamic-resolution`

  动态设置分辨率

- `/f`

  以全屏模式启动，使用 <kbd>Ctrl</kbd> + <kbd>Alt</kbd> + <kbd>Enter</kbd> 切换

- `/sacle-desktop:<100-500>`

  桌面应用按照百分比方案

- `/scale:[100|140|180]`

  屏幕缩放比例，只能是这几个值，更推荐使用 `/scale-desktop`

- `/fonts`

  让字体更加流畅，默认不启用，开启会占用带宽[^2]

- `/menu-anims`

	开启菜单的动画效果，默认不开启[^3]

- `-wallpaper`

	不显示 windows 桌面，可以减少渲染桌面的时间[^5]

- `-themes`

	不显示 windows 的主题，可以减少 windows 应用渲染的时间[^6]

- `-aero`

	不显示 windows aero 特效[^4]

- `/bpp:<8|16|24|32>`

  多少 bit 位来表示一个像素点，值越高颜色越细腻，占用的带宽也越大

### 0x02c Audio Related Options

TODO

### 0x02d Performance Related Options

- `/cache:[bitmap[:on|off],codec[:rfx|nsc],glyph[:on|off],offscreen[:on|off]]`

	缓存数据，在网络环境差的情况下可以减少传输的流量，提高稳定性

	- bitmap
	
		允许缓存 bitmap(静态图片)

	- codec
	
		允许缓存 codec(压缩的图片)，只推荐使用 rfx

	- glyph
	
		允许缓存 glyph(字形)

	- offscreen
	
		允许缓存 offscreen(UI)

- `/gfx:avc444`

	使用 avc444 codec 对视频优化程度最大，**如果图片加载卡顿建议启用**

- `/compression`

	对数据压缩，默认开启

- `/compression-level:[0|1|2]`

  数据压缩的等级

## 0x03 Keyboard Shortcut

- <kbd>Right-Ctrl</kbd>

	类似于 Virtualbox，releases keyboard and mouse grab

- <kbd>Ctrl</kbd> + <kbd>Alt</kbd> + <kbd>Enter</kbd>

	切换全屏


## 0x04 HDPI Display

因为屏幕的分辨率不同，freerdp 有些时候需要缩放才可以看的清楚，可以通过 `/scale:[100|140|180]` 或者 `/sacle-desktop:<100-500>` 放大 DPI，更推荐使用 scale-desktop

```
/usr/bin/xfreerdp /u:admin /p: /v:10.100.4.127 /scale-desktop:125 /dynamic-resolution
```

## 0x05 Examples

```
/usr/bin/xfreerdp /u:admin /p: /v:10.100.4.127 /bpp:16 /fonts /cert:ignore /scale:140 /scale-desktop:125 /dynamic-resolution

# 适用于 poor network connection
/usr/bin/xfreerdp /u:"admin" /p:"toor" /v:192.168.137.1 /tls:seclevel:0 /cert:ignore /bpp:16 /fonts /gfx:avc444 /cache:bitmap:on,codec:rfx,glyph:on,offscreen:on -aero -wallpaper -themes -decorations /scale-desktop:125 /f
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown)
- `man xfreerdp3`

***References***

[^1]:[FAQ · FreeRDP/FreeRDP Wiki · GitHub](https://github.com/FreeRDP/FreeRDP/wiki/FAQ#windows-7-errconnect_tls_connect_failed)
[^2]:[FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown#smooth-fonts)
[^3]:[FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown#menu-animations)
[^4]:[FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown#desktop-composition)
[^5]:[FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown#wallpaper)
[^6]:[FreeRDP-Manuals/User/FreeRDP-User-Manual.markdown at master · awakecoding/FreeRDP-Manuals · GitHub](https://github.com/awakecoding/FreeRDP-Manuals/blob/master/User/FreeRDP-User-Manual.markdown#themes)