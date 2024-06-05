---
author: "0x00"
createTime: 2024-06-03
lastModifiedTime: 2024-06-03-11:18
draft: true
---
# MySQL 01 - Binary Installation


## 0x01 Overview

MySQL 安装的方式有多种
1. Binary Installation - 使用已经编译好的二进制包安装，适用于所有的 Linux，也被成为 Generic Installation
2. Source Installation - 使用源码编译安装，按照 Distro 使用特定的包
3. Package Installation - 使用包管理器安装，安装包可以是对应源提供的，也可以是 MySQL 站点下载的

> [!NOTE]
> 这里只记录 Binary Installation

## 0x02 Which Package to Install[^1]

在安装 MySQL 之前需要先知道那个安装包是适合且可以在目标系统上安装的，安装包名由如下几部分组成
```
mysql-<MYSQL_VERSION>-linux-<GLIBC_VERSION>-<ARCH>-[FORMAT].<COMPRESSION>
```
- MYSQL_VERSION 
  	MySQL 的版本，由 3 部分组成
  	1. major version
  	2. minor version
  	3. the version number of LTS releases or innovation releases 在 innovation releases 中第三部分永远为 0
  	例如 8.4.1 LTS, 8.3.0 Innovaton
- GLIBC_VERSION
	系统使用的 glibc 版本，可以通过 `ldconfig --version` 来查看(同理 `ldd --version`)。例如
	```shell
	$ ldconfig --version
	ldconfig (GNU libc) 2.17
	```

	就需要使用 glibc 2.17 的安装包
	如果使用了错误的 glibc 包，就会出现诸如 `version GLIBCXX_3.4.20 not found` 的错误
- ARCH
	CPU 使用的架构，必须匹配，可以通过 `arch` 来查看
- FORMAT
	可选，只有一种 minimal，表示安装包不包括 debug binaries 以及 regular binaries
- COMPRESSION
	安装包使用的压缩格式

## 0x03 Prerequisites

在安装 MySQL 前，还需要满足如下条件

1. 系统上有 libaio(MySQL 依赖 libaio)[^2]
	如果没有安装就会出现如下错误
	```shell
	error while loading shared libraries: libaio.so.1: cannot open shared object file: No such file or directory
	```
	可以通过包管理安装 `yum install libaio` 
	或者编译安装[^3]
	```shell
	$ sudo wget https://pagure.io/libaio/archive/libaio-0.3.113/libaio-0.3.113.tar.gz && tar -xzvf libaio-0.3.113.tar.gz && cd libaio-0.3.113
	$ sudo sed -i '/install.*libaio.a/s/^/#/' src/Makefile
	$ sudo make && sudo make install
	...
	install -D -m 644 libaio.h /usr/include/libaio.h
	install -D -m 755 libaio.so.1.0.2 /usr/lib/libaio.so.1.0.2
	ln -sf libaio.so.1.0.2 /usr/lib/libaio.so.1
	ln -sf libaio.so.1.0.2 /usr/lib/libaio.so
	...
	```
	编译完成后可能还不能运行
	```shell
	$ sudo ldd mysqld | grep libaio
			libaio.so.1 => not found
	```
	执行 `ldconfig` 即可

## 0x04 Installation

选择 [MySQL Community Server](https://dev.mysql.com/downloads/mysql/) 进入下载页面

![](https://github.com/dhay3/image-repo/raw/master/20230802/2023-08-02_20-17.73ke7kx6670g.webp)
这里以 version 8.4.0 Linux-Generic glib 2.17 x86_64 为例
1. 下载 binary package
	```
	$ sudo wget https://dev.mysql.com/get/Downloads/MySQL-8.4/mysql-8.4.0-linux-glibc2.17-x86_64.tar.xz
	```

2. 校验 hash 和 签名，都可以在下载页面获取到
3. 解压
	```
	xz -d mysql-8.4.0-linux-glibc2.17-x86_64.tar.xz && tar -xvf mysql-8.4.0-linux-glibc2.17-x86_64.tar
	```
	解压后会包含如下几个目录
	
	| Directory       | Contents of Directory                                         |
	| --------------- | ------------------------------------------------------------- |
	| `bin`           | mysqldserver, client and utility programs                     |
	| `docs`          | MySQL manual in Info format                                   |
	| `man`           | Unix manual pages                                             |
	| `include`       | Include (header) files                                        |
	| `lib`           | Libraries                                                     |
	| `share`         | Error messages, dictionary, and SQL for database installation |
	| `support-files` | Miscellaneous support files                                   |
	
4. 创建配置文件
   ```
	cd mysql-8.4.0-linux-glibc2.17-x86_64 && mkdir etc && cd etc
	touch my.cnf
	```



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[MySQL :: MySQL 8.4 Reference Manual :: 2.1.2 Which MySQL Version and Distribution to Install](https://dev.mysql.com/doc/refman/8.4/en/which-version.html)
[^2]:[MySQL :: MySQL 8.4 Reference Manual :: 2.2 Installing MySQL on Unix/Linux Using Generic Binaries](https://dev.mysql.com/doc/refman/8.4/en/binary-installation.html)
[^3]:[libaio-0.3.113](https://www.linuxfromscratch.org/blfs/view/svn/general/libaio.html)
