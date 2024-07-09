# OpenWrt  opkgs

ref

https://openwrt.org/docs/guide-user/installation/after.installation

https://openwrt.org/docs/guide-user/additional-software/opkg

## Digest

syntax

```
opkg must have one sub-command argument:
usage: opkg [options...] sub-command [arguments...]
where sub-command is one of:
Package Manipulation:
        update                  Update list of available packages
        upgrade <pkgs>          Upgrade packages
        install <pkgs>          Install package(s)
        configure <pkgs>        Configure unpacked package(s)
        remove <pkgs|regexp>    Remove package(s)
        flag <flag> <pkgs>      Flag package(s)
         <flag>=hold|noprune|user|ok|installed|unpacked (one per invocation)
```

和大多数的 Linux distro 一样，OpenWrt 也有自己的包管理工具。`opkg` 就是 OpenWrt 的包管理工具

## Package manipulation sub coommands

- `update`

  用于更新 available packages list

- `upgrade <pkgs>`

  用于更新 package 例如更新 package1 和 package2 可以使用下面的命令

  ```
  opkg upgrade package1 package2
  ```

  可以使用 `opkg list-upgradeble` 来查看可以更新的 package

- `install <pkgs|url>`

  用于安装 packages, 例如

  ```
  opkg install hiawatha
  opkg install http://downloads.openwrt.org/snapshots/trunk/ar71xx/packages/hiawatha_7.7-2_ar71xx.ipk
  opkg install /tmp/hiawatha_7.7-2_ar71xx.ipk
  ```

- `configure <pkgs>`

  用于配置 uppacked packages

- `remove <pkgs|globp>`

  删除 packages

- `flag <flag> <pkgs>`

  给 packages 打标，flag 的值可以是  hold, noprune, user, ok, installed, upacked

## Informational subcommands

- `list [pkg|globp]`

  显示可以安装 packages，或者显示指定的

- `list-installed | list-upgradeble`

  顾名思义

- `files <pkg>`

  查看 package 包含的文件

- `search <file|globp>`

  查看 file 属于那个 package

  ```
  root@OpenWrt:~# opkg files busybox | grep zip
  /bin/gunzip
  /bin/gzip
  root@OpenWrt:~# opkg search /bin/gzip
  busybox - 1.35.0-4
  ```

- `info [pkg|globp]`

  查看 package 的具体信息

- `download <pkg>`

  将 package 下载到当前目录

- `whatdepends [-A] [pkgname | pat]+`

  查看 installed package 的依赖

## Optional args

- `-A`

  query all package not just those installed

  只能和 `whatdepends` 一起使用

- `-V [<level>]`

  verbose output

- `-f <conf_file>`

  制定使用 opkg 的配置文件，默认使用 `/etc/opkg.conf`

- `--force-mainainer`

  覆写现有的配置文件

- `--force-reinstall`

  强制重新安装

- `--force-overwrite`

  从其他 package 覆写文件

- `--force-downgrade`

  允许 downgrade package

- `--noaction`

  只用测试不安装 dryrun

- `--autoremove`

  安装完成后，自动删除依赖

## Feeds

Feed 是 opkg 的镜像源配置文件，在 `/etc/opkg/distfeeds.conf`

可以修改 `customfeed.conf` 来添加指定的镜像源

国内如果不能科学上网，可以添加国内镜像源

清华

https://mirrors.tuna.tsinghua.edu.cn/help/openwrt/

腾讯

https://mirrors.tencent.com/

```
src/gz tuna_core https://mirrors.tuna.tsinghua.edu.cn/openwrt/releases/22.03.2/targets/x86/64/packages
src/gz tuna_base https://mirrors.tuna.tsinghua.edu.cn/openwrt/releases/22.03.2/packages/x86_64/base
src/gz tuna_luci https://mirrors.tuna.tsinghua.edu.cn/openwrt/releases/22.03.2/packages/x86_64/luci
src/gz tuna_routing https://mirrors.tuna.tsinghua.edu.cn/openwrt/releases/22.03.2/packages/x86_64/routing
```

