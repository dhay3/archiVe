# GRUB 配置文件

参考：

https://www.gnu.org/software/grub/manual/grub/grub.html#Configuration

GRUB使用`grub.cfg`配置文件由`grub-mkcofnig`，通常位于`/boot/grub`下。`/etc/default/grub`控制了`grub-mkconfig`改如何生成配置文件。它会发现可以被使用的kernel然后将它展示在引导阶段的menu entries

![Snipaste_2021-04-15_15-17-23](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-04-15_15-17-23.6tpmm7zqnh40.png)

## /etc/default/grub

> 多个值通过空格隔开
>
> 修改配置后`grub-update`或`grub-mkconfig -o <grub-mkconfig-path>`

- GRUB_DEFAULT

  表示在menu entries中低级显示，从0开始

- GRUB_TIMEOUT

  引导默认的entry在指定的sec之后，默认5，0表示立即引导，-1表示等待手动选择

- GRUB_TIMEOUT_STYLE

  1. menu表示等待timeout
  2. GRUB_TIMEOUT表示等待timeout但是可以被任意一个按键终止
  3. countdown | hidden 表示会等待timeout，如果键入ECS直接展示menu。如果没有直接引导默认的entry

- GRUB_DEFAULT_BUTTON

  GRUB_TIMEOUT_BUTTON

  GRUB_TIMEOUT_STYLE_BUTTON

  GRUB_BUTTON_CMOS_ADDRESS

- GRUB_DISTRIBUTOR

  menu 展示额外的信息

- GRUB_TERMINAL_INPUT

  指定终端的输入设备，默认使用平台本地支持的terminal input

- GRUB_TERMINAL_OUTPUT

  指定终端输出的设备，默认使用平台本地的输出设备。其中有一个`spkmodem`非常有用，当设备的串行接口不可用时。==尚未实验成功==

- GRUB_TERMINAL

  等价INPUT+OUTPUT

- ==GRUB_CMDLINE_LINUX==

  指定的GRUB cli 对所有的menu entries生效

- ==GRUB_CMDLINE_LINUX_DEFAULT==

  在GRUB_CMDLINE_LINUX对defulat entry(非recovery entry)额外增加GRUB cli

- GRUB_DISABLE_LINUX_UUID

  是否为root的filesystem设置UUID以和其他用户的filesystem区别

- GRUB_DISABLE_LINUX_PARTUUID

  如果GRUB_DISABLE_LINUX_UUID为true，则可以通过改选项通过分区的UUID来表示root的filesystem

- GRUB_DISABLE_LINUX_RECOVERY

  只有设置为true，否则会为每个kernel设置两条entry

  ![Snipaste_2021-04-15_16-21-59](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-04-15_16-21-59.wod5lpqgto0.png)

- GRUB_DISABLE_OS_PROBER

  默认grub-mkconfig会使用os-prober去探测安装在统一系统的其他操作系统并将其展示在menu

- GRUB_OS_PROBER_SKIP_LIST

  跳过探测指定UUID的filesystem

- GRUB_DISABLE_SUBMENU

  grub-mkcofnig通常会根据kernel的版本从高到低排序展示在menu，如果设置为false会按照子目录的形式展示

  ![Snipaste_2021-04-15_16-58-56](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-04-15_16-58-56.4u02fbwlbrk.png)

  进入Advanced options 可以选择内核。需要修改GRUB_DEFAULT，和fallback，default环境变量。

- GRUB_ENABLE_CRYPTODISK

  GRUB检查是否由加密的磁盘，引导是需要输入密码才能加载内核，`y`表示开启

- GRUB_PRELOAD_MODULES

  读取`grub.cfg`之前调用的模块

## gfxterm

gfterm是GRUB引导的图形化终端

- GRUB_GFXMODE

  gfxterm使用的模式，默认为auto

- GRUB_BACKGROUND

  gfxterm使用的背景，必须以`.png, .tga, .jpg, or .jpeg.`j结尾，会自动伸缩

- GRUB_THEME

  gfxterm的主题

- GRUB_GFXPAYLOAD_LINUX

  gfxterm展示的大小

  











































