# 查看LINUX发行版的名称及其版本号

转载:

https://www.qiancheng.me/post/coding/show-linux-issue-version

2017-01-16

最近跟合作方支付公司（一个北京的互联网支付公司，就不具体提名字啦）沟通的时候，需要对方生成非对称加密密钥公钥提供给我方，对方技术是个妹子。不懂怎么在预发／生产机器上面生成加密算法的公私钥，也不知道怎么查看系统版本。属于一问三不知类型，怎么办～ 我也只能打电话过去一步步手把手教如何查看发行版，如何安装命令，如何生成对应的公私钥。下面讲讲如何查看系统发行版和版本号。

查看LINUX发行版的名称及其版本号的命令,这些信息对于添加正确的软件更新源很有用，而当你只能在命令行下工作的时候，下面的方法可以帮忙。

## 一、查看Linux内核版本命令（两种方法）：

1、`cat /proc/version`

```
[root@localhost ~]# cat /proc/version
Linux version 2.6.18-194.8.1.el5.centos.plus
(mockbuild@builder17.centos.org) (gcc version 4.1.220080704
(Red Hat 4.1.2-48)) #1 SMP Wed Jul 7 11:50:45 EDT 2010
```

2、`uname -a`

```
[root@localhost ~]# uname -a
Linux localhost.localdomain 2.6.18-194.8.1.el5.centos.plus 
#1 SMP Wed Jul 7 11:50:45 EDT 2010 i686 i686 i386 GNU/Linux
```

## 二、查看Linux系统版本的命令（3种方法）

1. `lsb_release -a`，即可列出所有版本信息：

```
[root@localhost ~]# lsb_release -a
LSB Version: :core-3.1-ia32:core-3.1-noarch:graphics-3.1-ia32:graphics-3.1-noarch
Distributor ID: CentOS
Description: CentOS release 6.5 (Final)
Release: 6.5
Codename: Final
```

这个命令适用于所有的Linux发行版，包括Redhat、SuSE、Debian…等发行版。

2. `cat /etc/redhat-release`，这种方法只适合Redhat系的Linux：

```
[root@localhost ~]# cat /etc/redhat-release
CentOS release 6.7 (Final)
```

3. `cat /etc/issue`，此命令也适用于所有的Linux发行版。

```
[root@localhost ~]# cat /etc/issue
CentOS release 6.7 (Final)
Kernel \r on an \m
```

4. `cat /etc/os-release`

```
[root@chz Desktop]# cat /etc/os-release 
PRETTY_NAME="Kali GNU/Linux Rolling"
NAME="Kali GNU/Linux"
ID=kali
VERSION="2020.3"
VERSION_ID="2020.3"
VERSION_CODENAME="kali-rolling"
ID_LIKE=debian
ANSI_COLOR="1;31"
HOME_URL="https://www.kali.org/"
SUPPORT_URL="https://forums.kali.org/"
BUG_REPORT_URL="https://bugs.kali.org/"
```

