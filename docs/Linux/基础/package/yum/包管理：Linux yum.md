# Linux yum

## 概述

yum（ Yellow dog Updater, Modified）是一个在 Fedora 和 RedHat 以及 SUSE 中的 Shell 前端软件包管理器。会下载相关依赖，==支持wildcard，先缓存安装包，然后安装==

pattern `yum [options] [command] [package...]`

## 命令

- install

  安装

  ```
  [root@chz yum.repos.d]# yum install vim*
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  base                                                                                                                       | 3.6 kB  00:00:00     
  extras                                                                                                                     | 2.9 kB  00:00:00     
  updates                                                                                                                    | 2.9 kB  00:00:00     
  Package 2:vim-minimal-7.4.629-6.el7.x86_64 already installed and latest version
  Package 2:vim-filesystem-7.4.629-6.el7.x86_64 already installed and latest version
  Package 2:vim-common-7.4.629-6.el7.x86_64 already installed and latest version
  Package 2:vim-enhanced-7.4.629-6.el7.x86_64 already installed and latest version
  ```

  ```
  yum -y install vim*
  yum -y install wget
  yum -y install net-tools
  ```

- update

  询问，并更新。不带包表示全部软件包

  ```
  [root@chz yum.repos.d]# yum update 
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  Resolving Dependencies
  --> Running transaction check
  ---> Package LibRaw.x86_64 0:0.19.2-1.el7 will be updated
  ---> Package LibRaw.x86_64 0:0.19.4-1.el7 will be an update
  ---> Package NetworkManager.x86_64 1:1.18.0-5.el7 will be updated
  ---> Package NetworkManager.x86_64 1:1.18.4-3.el7 will be an update
  ---> Package NetworkManager-adsl.x86_64 1:1.18.0-5.el7 will be updated
  ---> Package NetworkManager-adsl.x86_64 1:1.18.4-3.el7 will be an update
  ---> Package NetworkManager-glib.x86_64 1:1.18.0-5.el7 will be updated
  ---> Package NetworkManager-glib.x86_64 1:1.18.4-3.el7 will be an update
  ```

- check-update

  检查是否有更新

  ```
  [root@chz yum.repos.d]# yum check-update 
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  
  LibRaw.x86_64                                                                0.19.4-1.el7                                                base     
  NetworkManager.x86_64                                                        1:1.18.4-3.el7                                              base     
  NetworkManager-adsl.x86_64                                                   1:1.18.4-3.el7                                              base    
  ```

- yum upgrade

  检查更新并询问是否需要更新

  ```
  [root@chz yum.repos.d]# yum upgrade NetworkManager.x86_64
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  Resolving Dependencies
  --> Running transaction check
  ---> Package NetworkManager.x86_64 1:1.18.0-5.el7 will be updated
  --> Processing Dependency: NetworkManager = 1:1.18.0-5.el7 for package: 1:NetworkManager-ppp-1.18.0-5.el7.x86_64
  --> Processing Dependency: NetworkManager = 1:1.18.0-5.el7 for package: 1:NetworkManager-tui-1.18.0-5.el7.x86_64
  ```

- yum remve|erase

  询问并卸载包

  ```
  [root@chz yum.repos.d]# yum remove NetworkManager.x86_64
  Loaded plugins: fastestmirror, langpacks
  Resolving Dependencies
  --> Running transaction check
  ---> Package NetworkManager.x86_64 1:1.18.0-5.el7 will be erased
  --> Processing Dependency: NetworkManager = 1:1.18.0-5.el7 for package: 1:NetworkManager-ppp-1.18.0-5.el7.x86_64
  ```

> list 与 search 的区别就是，==search能搜索的范围包括包的概述，但是list只对包名进行检索==
>
> 如果想要使用模式扩展需要使用双引号，如`apt list "apache2*"`

- yum list

  列出软件包，可以使用`--installed`，`--updateable`来过滤

  ```
  [root@chz yum.repos.d]# yum list NetworkManager.x86_64
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  Installed Packages
  NetworkManager.x86_64                                                   1:1.18.0-5.el7                                                   @anaconda
  Available Packages
  NetworkManager.x86_64                                                   1:1.18.4-3.el7                                                   base  
  ```

- yum search 

  搜索yum源中可用的包

  ```
  [root@cyberpelican local]# yum search mysql|more
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * centos-sclo-rh: mirrors.cn99.com
   * centos-sclo-sclo: mirrors.cn99.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  ============================== N/S matched: mysql ==============================
  MySQL-python.x86_64 : An interface to MySQL
  akonadi-mysql.x86_64 : Akonadi MySQL backend support
  apr-util-mysql.x86_64 : APR utility library MySQL DBD driver
  dovecot-mysql.x86_64 : MySQL back end for dovecot
  freeradius-mysql.x86_64 : MySQL support for freeradius
  libdbi-dbd-mysql.x86_64 : MySQL plugin for libdbi
  
  ```

- yum repolist

  列出当前的源

  ```
  [root@chz yum.repos.d]# yum repolist 
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  repo id                                                   repo name                                                                         status
  base/7/x86_64                                             CentOS-7 - Base - mirrors.aliyun.com                                              10,070
  extras/7/x86_64                                           CentOS-7 - Extras - mirrors.aliyun.com                                               413
  updates/7/x86_64                                          CentOS-7 - Updates - mirrors.aliyun.com                                            1,134
  repolist: 11,617
  ```

- yum clean all

  [ packages | metadata | expire-cache | rpmdb | plugins | all ]

  清除所有下载的缓存（安装包和检索信息）

  ```
  [root@chz yum.repos.d]# yum clean all
  Loaded plugins: fastestmirror, langpacks
  Cleaning repos: base extras updates
  Cleaning up list of fastest mirrors
  ```

- yum makecache

  将服务器上的软件包信息缓存到本地，以提高 搜索 安装软件的速度

  ```
  [root@chz yum.repos.d]# yum makecache 
  Loaded plugins: fastestmirror, langpacks
  Determining fastest mirrors
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  base                                                                                                                       | 3.6 kB  00:00:00     
  extras                                                                                                                     | 2.9 kB  00:00:00     
  updates                                                                                                                    | 2.9 kB  00:00:00     
  (1/10): base/7/x86_64/group_gz                                                                                             | 153 kB  00:00:00     
  (2/10): base/7/x86_64/filelists_db                                                                                         | 7.1 MB  00:00:01     
  (3/10): base/7/x86_64/primary_db                                                                                           | 6.1 MB  00:00:01     
  (4/10): base/7/x86_64/other_db                                                                                             | 2.6 MB  00:00:00     
  (5/10): extras/7/x86_64/filelists_db                                                                                       | 217 kB  00:00:00     
  (6/10): extras/7/x86_64/other_db                                                                                           | 124 kB  00:00:00     
  (7/10): extras/7/x86_64/primary_db                                                                                         | 206 kB  00:00:00     
  (8/10): updates/7/x86_64/filelists_db                                                                                      | 2.4 MB  00:00:00     
  (9/10): updates/7/x86_64/other_db                                                                                          | 318 kB  00:00:00     
  (10/10): updates/7/x86_64/primary_db                                                                                       | 4.5 MB  00:00:00     
  Metadata Cache Created
  ```

- yum info

  列出包的信息包括已安装和未安装的

  ```
  [root@chz yum.repos.d]# yum info gdb|more
  Loaded plugins: fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  Installed Packages
  Name        : gdb
  Arch        : x86_64
  Version     : 7.6.1
  Release     : 115.el7
  Size        : 7.0 M
  Repo        : installed
  From repo   : anaconda
  Summary     : A GNU source-level debugger for C, C++, Fortran, Go and other
              : languages
  URL         : http://gnu.org/software/gdb/
  License     : GPLv3+ and GPLv3+ with exceptions and GPLv2+ and GPLv2+ with
              : exceptions and GPL+ and LGPLv2+ and BSD and Public Domain
  Description : GDB, the GNU debugger, allows you to debug programs written in C,
              : C++, Java, and other languages, by executing them in a controlled
              : fashion and printing their data.
  
  Available Packages
  Name        : gdb
  Arch        : x86_64
  Version     : 7.6.1
  Release     : 119.el7
  ```

- yum grouplist

  列出可用的组

  ```
  [root@chz yum.repos.d]# yum grouplist 
  Loaded plugins: fastestmirror, langpacks
  There is no installed groups file.
  Maybe run: yum groups mark convert (see man yum)
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  Available Environment Groups:
     Minimal Install
     Compute Node
     Infrastructure Server
     File and Print Server
     Basic Web Server
     Virtualization Host
     Server with GUI
     GNOME Desktop
     KDE Plasma Workspaces
     Development and Creative Workstation
  Available Groups:
     Compatibility Libraries
     Console Internet Tools
     Development Tools
     Graphical Administration Tools
     Legacy UNIX Compatibility
     Scientific Support
     Security Tools
     Smart Card Support
     System Administration Tools
     System Management
  Done
  ```

- yum groupinstall

  组安装

  ```
  [root@chz yum.repos.d]# yum groupinstall 'File and Print Server'
  Loaded plugins: fastestmirror, langpacks
  There is no installed groups file.
  Maybe run: yum groups mark convert (see man yum)
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  base                                                                                                                       | 3.6 kB  00:00:00     
  extras                                                                                                                     | 2.9 kB  00:00:00     
  updates                                                                                                                    | 2.9 kB  00:00:00     
  (1/4): base/7/x86_64/group_gz                                                                                              | 153 kB  00:00:00     
  (2/4): extras/7/x86_64/primary_db                                                                                          | 206 kB  00:00:00     
  (3/4): updates/7/x86_64/primary_db                                                                                         | 4.5 MB  00:00:00     
  (4/4): base/7/x86_64/primary_db                                                                                            | 6.1 MB  00:00:01     
  Warning: Group core does not have any packages to install.
  Package rubygem-abrt-0.3.0-1.el7.noarch already installed and latest version
  Warning: Group base does not have any packages to install.
  ```

> yum install yum-fastestmirror
>
> 安装此插件可以位yum安装加速

## 参数

- -y

  默认选中yes

- --color=always

  以彩色显示

## 配置文件

配置文件`/etc/yum.conf`，具体可查看`man yum.conf`。这里不再过多赘述。

- reposdir

  yum 源存储的位置

## 配置阿里云yum源

具体参考：

https://developer.aliyun.com/mirror/centos?spm=a2c6h.13651102.0.0.3e221b11hDwVpB

https://developer.aliyun.com/mirror/epel?spm=a2c6h.13651102.0.0.3e221b11hDwVpB

## 删除yum源

删除`etc/yum.repo.d/`指定的文件，`yum clean all`

## 常见错误

- Another app is currently holding the yum lock; waiting for it to exit...

  ```
  rm -f /var/run/yum.pid
  ```

  

