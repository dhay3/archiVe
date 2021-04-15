# GRUB Theme file 

参考：

https://www.gnu.org/software/grub/manual/grub/grub.html#Theme-file-format

GRUB支持修改 boot menu 和 GRUB graphical menu的样式

==可以使用[这里](https://github.com/vinceliuice/grub2-themes)提供的主题==

修改完后记住让GRUB配置文件生效

## Colors

- 支持HTML格式，例如：“#RRGGBB”
- 支持RGB格式，例如：“128, 128, 255”

- 支持SVG 1.0 color格式，例如：“cornflowerblue”，必须小写

## Fonts

使用"PFF2 font format"

 Fonts are specified with full font names. Currently there is no provision for a preference list of fonts, or deriving one font from another. Fonts are loaded with the “loadfont” command in GRUB ([loadfont](https://www.gnu.org/software/grub/manual/grub/grub.html#loadfont)). To see the list of loaded fonts, execute the “lsfonts” command ([lsfonts](https://www.gnu.org/software/grub/manual/grub/grub.html#lsfonts)). If there are too many fonts to fit on screen, do “set pager=1” before executing “lsfonts”.

## Progress Bar





