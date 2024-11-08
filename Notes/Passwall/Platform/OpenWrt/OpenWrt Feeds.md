# OpenWrt Feeds

ref

https://openwrt.org/docs/guide-developer/feeds

## Digest

OpenWrt Feeds 简单的可以理解成是 OpenWrt packages 的镜像源 ，但是只在编译固件时用到。

## Feeds configuration

Openwrt 默认把 Feeds 存储在 `/etc/opkg/feeds.conf` 

一条 feeds 条目由 3 个组件组成

1. feed method
2. feed name
3. feed source

### feed method

| Method       | Function                                                     |
| :----------- | :----------------------------------------------------------- |
| src-bzr      | Data is downloaded from the source path/URL using `bzr`      |
| src-cpy      | Data is copied from the source path. The path can be specified as either relative to OpenWrt repository root or absolute. |
| src-darcs    | Data is downloaded from the source path/URL using `darcs`    |
| src-git      | Data is downloaded from the source path/URL using `git` as a shallow (depth of 1) clone |
| src-git-full | Data is downloaded from the source path/URL using `git` as a full clone |
| src-gitsvn   | Bidirectional operation between a Subversion repository and git |
| src-hg       | Data is downloaded from the source path/URL using `hg`       |
| src-link     | A symlink to the source path is created. The path must be absolute. |
| src-svn      | Data is downloaded from the source path/URL using `svn`      |

### feed name

顾名思义是 feed 的名字，用于标识

### feed source

feed 存储的位置

### Exmaple

```
#lede 原生自带的包
src-git packages https://github.com/coolsnowwolf/packages
src-git luci https://github.com/coolsnowwolf/luci
src-git routing https://github.com/coolsnowwolf/routing
#openwrt 官网的包
src-git telephony https://git.openwrt.org/feed/telephony.git
#自定义的包
src-git helloworld https://github.com/fw876/helloworld.git
src-git https://github.com/kenzok8/openwrt-packages.git
#包含 passwall 依赖，如果不指定无法正常编译 passwall
src-git https://github.com/kenzok8/small.git
```

kenzok8 中已包含多数常用的 feeds 源, 为了防止不可靠因素最好 fork 一份。下面是单独的 git 地址

```
https://github.com/xiaorouji/openwrt-passwall.git
https://github.com/xiaorouji/openwrt-passwall2.git
https://github.com/frainzy1477/luci-app-clash.git
adguardhome https://github.com/rufengsuixing/luci-app-adguardhome
https://github.com/destan19/OpenAppFilter.git
https://github.com/esirplayground/luci-app-poweroff.git
https://github.com/kenzok8/openwrt-packages.git
```

## Run feeds

在 2018 年后，feeds 都由 `~/scripts/feeds` 管理

```
$ ./scripts/feeds 
Usage: ./scripts/feeds <command> [options]

Commands:
	list [options]: List feeds, their content and revisions (if installed)
	Options:
	    -n :            List of feed names.
	    -s :            List of feed names and their URL.
	    -r <feedname>:  List packages of specified feed.
	    -d <delimiter>: Use specified delimiter to distinguish rows (default: spaces)
	    -f :            List feeds in feeds.conf compatible format (when using -s).

	install [options] <package>: Install a package
	Options:
	    -a :           Install all packages from all feeds or from the specified feed using the -p option.
	    -p <feedname>: Prefer this feed when installing packages.
	    -d <y|m|n>:    Set default for newly installed packages.
	    -f :           Install will be forced even if the package exists in core OpenWrt (override)

	search [options] <substring>: Search for a package
	Options:
	    -r <feedname>: Only search in this feed

	uninstall -a|<package>: Uninstall a package
	Options:
	    -a :           Uninstalls all packages.

	update -a|<feedname(s)>: Update packages and lists of feeds in feeds.conf .
	Options:
	    -a :           Update all feeds listed within feeds.conf. Otherwise the specified feeds will be updated.
	    -i :           Recreate the index only. No feed update from repository is performed.
	    -f :           Force updating feeds even if there are changed, uncommitted files.

	clean:             Remove downloaded/generated files.
```

其实只要注意 2 点，修改完 feeds 后需要执行一下

```
#将 feeds 内的插件拷贝到本地，具体参考 openwrt 官网
./scripts/feeds update -a
#将 feeds 内的插件安装，具体参考 openwrt 官网
./scripts/feeds install -a
```

## Using feeds

上面的操作都完成后，就可以通过 `make menuconfig` 生成固件的配置文件，记录固件需要安装的插件











