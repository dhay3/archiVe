# OpenWrt LEDE firmware

ref

https://github.com/coolsnowwolf/lede

https://sspai.com/post/61463

https://www.solarck.com/post/install-ssr-plus/

https://openwrt.org/docs/guide-developer/feeds

https://lighti.me/4174.html

## Digest

OpenWrt LEDE 需要编译出符合自己的固件比较麻烦，且相关知识比较凌乱。所以开一篇文章记录一下

## Compile

在编译固件前需要安装好编译需要的依赖，为了避免因为 OS 而导致编译失败，尽量采用 Debian 系的 OS。full-upgrade 以及依赖安装大概会占用 1.5 G 左右的空间，所以虚拟机的磁盘分区不能太小，以防止固件编译的时候磁盘空间不足( ==实际至少需要 40 G 的磁盘空间== )。是否进行 full-upgrade 可以由自己决定

```
sudo apt update -y
sudo apt full-upgrade -y
sudo apt install -y ack antlr3 aria2 asciidoc autoconf automake autopoint binutils bison build-essential \
bzip2 ccache cmake cpio curl device-tree-compiler fastjar flex gawk gettext gcc-multilib g++-multilib \
git gperf haveged help2man intltool libc6-dev-i386 libelf-dev libglib2.0-dev libgmp3-dev libltdl-dev \
libmpc-dev libmpfr-dev libncurses5-dev libncursesw5-dev libreadline-dev libssl-dev libtool lrzsz \
mkisofs msmtp nano ninja-build p7zip p7zip-full patch pkgconf python2.7 python3 python3-pip libpython3-dev qemu-utils \
rsync scons squashfs-tools subversion swig texinfo uglifyjs upx-ucl unzip vim wget xmlto xxd zlib1g-dev
```

git clone lean's lede 代码

```
git clone https://github.com/coolsnowwolf/lede
cd lede
```

修改 `feeds.conf.default` 配置, 增加常用的 feeds 源。如果不修改 feeds 默认不会安装 翻墙 插件

```
src-git packages https://github.com/coolsnowwolf/packages
src-git luci https://github.com/coolsnowwolf/luci
src-git routing https://github.com/coolsnowwolf/routing
src-git telephony https://git.openwrt.org/feed/telephony.git
src-git helloworld https://github.com/fw876/helloworld.git
src-git kenzok8 https://github.com/kenzok8/openwrt-packages.git
src-git kenzok8small https://github.com/kenzok8/small.git
```

更新 feeds 插件列表并将 feeds 中的插件安装到本地用于编译。如果 feeds 更新了同样需要执行一遍这 2 个命令，否则对应的插件不会出现在 menuconfig 中

```
./scripts/feeds update -a
./scripts/feeds install -a
```

选择固件需要安装的插件以及固件的一些配置。默认会编译 x86_64 架构的固件，如果需要安装固件的主机不是 对应架构的需要按需

选择 Target System, Subtraget System

```
make menuconfig
```

Luci 部分插件功能可以参考

https://www.right.com.cn/forum/thread-344825-1-2.html

编译固件，时间根据选择的插件数而异，通常 1 - 2 小时甚至更长。==血的教训尽量别全选，如果需要 lede 库外的插件，需要先编译一遍 lede 库的，然后再编译 lede 库外的==

```
make download -j$(nproc)
make V=s -j$(nproc)
```

固件编译完后会在 `bin/trages`内

## About SSR

https://github.com/coolsnowwolf/lede/commit/af803f2569bbbaac88d001890a8a7646abf0ce0d#comments

SSR 已经被 Helloworld 替代了