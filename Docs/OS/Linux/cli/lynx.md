# lynx

ref

https://lynx.invisible-island.net/

https://lynx.invisible-island.net/lynx_help/

## Digest

lynx 是一个 TUI 浏览器

## Help

lynx 有这完整的帮助信息，可以在运行 lynx 时使用 `?` 或者 `H` 来查看

## How to use

这里只显示常用的命令

- right-arrow 或者 enter

  “点击” link

- q

  退出 lynx 需 prompt

- ctrl + D 或者 Q

  退出 lynx 且不 prompt

- ==backspace==

  历史记录

- ctrl + R

  刷新

- 7/:arrow_up:/9 

  top of text/prev lnk/pageup

- 1/:arrow_down:/3

  end of text/next link/page down

- 4/:arrow_left:

  back prev doc

- 6/:arrow_right:

  click link

- / or s

  search

- d

  下载

- ==ctrl + n==

  鼠标滚轮向下

- ==ctrl + p==

  鼠标滚轮向上

- g

  跳转到指定 url

- !

  spawn a shell

  when you quit or exit the shell you will return to lynx

## lynx.cfg

lynx 启动时会读取 lynx.cfg 作为配置文件，默认`/etc/lynx.cfg`。如果想要指定配置文件，需要通过`-cfg` 参数或者 `LYNX_CFG` 来手动自动配置文件

启动 lynx 时可以通过 `g lynxcfg:` 来查看当前配置文件

## Options

https://lynx.invisible-island.net/release/lynx_help/cattoc.html

lynx 内置了很多的 options ，我们可以在 lynx.cfg 中调用这些 options 以达到客制化

例如 `/etc/lynx.cfg` 中添加如下配置

```
ACCEPT_ALL_COOKIES:TRUE
VI_KEYS_ALWAYS_ON:TRUE
XLOADIMAGE_COMMAND:gwenview %s &
VIEWER:application/postscript:gwenview %s&:XWINDOWS
VIEWER:image/gif:gwenview %s&:XWINDOWS
VIEWER:image/x-xbm:gwenview %s&:XWINDOWS
VIEWER:image/png:gwenview %s&:XWINDOWS
VIEWER:image/tiff:gwenview %s&:XWINDOWS
VIEWER:image/jpeg:gwenview %s&:XWINDOWS
VIEWER:video/mpeg:gwenview %s &:XWINDOWS
```

## Images

https://lynx.invisible-island.net/release/lynx_help/body.html#VIEWER

https://unix.stackexchange.com/questions/180760/navigation-and-images-in-lynx-text-browser

https://www.iana.org/assignments/media-types/media-types.xhtml#image

lynx 作为一个 TUI 浏览器，默认不显示 images，但是可以通过 `lynx.cfg` 设置 viewer 来显示 images	

```
VIEWER:<mime type>:<viewer command>[:<enviroment>]
```

配置文件如下

```
VIEWER:image/gif:gwenview %s&:XWINDOWS
VIEWER:image/jpeg:gwenview %s&:XWINDOWS
VIEWER:image/png:gwenview %s&:XWINDOWS
VIEWER:image/bmp:gwenview %s&:XWINDOWS
```

但是实际测试没有生效的需要 xli

## Cautions

### Installation

https://www.linode.com/community/questions/22943/how-do-i-fix-error-while-loading-shared-libraries-libsslso11

https://blog.csdn.net/sunhson/article/details/106476185

https://bbs.archlinux.org/viewtopic.php?id=225478 

安装 lynx 后，可能会遇到一些错误

```
lynx: error while loading shared libraries: libssl.so.3: cannot open shared object file: No such file or directory
```

这是因为 lynx 使用 openssl 3 编译，但是当前系统的 openssl 不适配, 所以报错

可以到 [openssl 官网](https://www.openssl.org/source/) 下载 openssl 3，然后执行

```
./config
make
```

会生成两个 lib( libssl.so.x 和 libcrypto.so.x )，将这两个 lib 复制到 `/usr/lib` 即可

但是这样可能会导致系统不能正确识别 libssl 和 libcrypto 导致 pacman 不能正常更新和安装。这时候可以参考

https://bbs.archlinux.org/viewtopic.php?id=225528

https://unix.stackexchange.com/questions/240252/pacman-exists-on-filesystem-error

```
#回退到正确的 openssl 版本，这里一定要用 overwrite，否则会报 libssl.so.1.1 和 libcrypto.so.1.1 already exist
sudo pacman -U  /var/cache/pacman/pkg/openssl-1.1.1.q-1-x86_64.pkg.tar.zst --overwrite *
```

### Charset

https://stackoverflow.com/questions/51548306/lynx-utf-8-support

```
lynx -display_charset=utf-8 http://www.aliyun.com/
```

可以通过 alias 设置快捷命令

