# Linux rpm 

参考：http://c.biancheng.net/view/2868.html

## 命名规则

rpm二进制包命名格式一般如下：

```
包名-版本号-发布次数-发行商-Linux平台-适合的硬件平台-包扩展名
```

例如：`httpd-2.2.15-15.el6.centos.1.i686.rpm`

- 包名：httpd
- 版本号：2.2.15
- 发布次数：15
- 发行商：el6 表示此包是由 Red Hat 公司发布，适合在 RHEL 6.x (Red Hat Enterprise Unux) 和 CentOS 6.x 上使用

- linux平台：表示此包适用于 CentOS 系统

- 适合的硬件平台：表示此包使用的硬件平台

  | 平台名称 | 适用平台信息                                                 |
  | -------- | ------------------------------------------------------------ |
  | i386     | 386 以上的计算机都可以安装                                   |
  | i586     | 686 以上的计算机都可以安装                                   |
  | i686     | 奔腾 II 以上的计算机都可以安装，目前所有的 CPU 是奔腾 II 以上的，所以这个软件版本居多 |
  | x86_64   | 64 位 CPU 可以安装                                           |
  | noarch   | 没有硬件限制                                                 |

- 包扩展名：RPM 包的扩展名，表明这是编译好的二进制包，可以使用 rpm 命令直接安装

## 默认安装路径

| 安装路径        | 含 义                      |
| --------------- | -------------------------- |
| /etc/           | 配置文件安装目录           |
| /usr/bin/       | 可执行的命令安装目录       |
| /usr/lib/       | 程序所使用的函数库保存位置 |
| /usr/share/doc/ | 基本的软件使用手册保存位置 |
| /usr/share/man/ | 帮助文件保存位置           |

## rpm管理

### 安装

> 此命令还可以一次性安装多个软件包，仅需将包全名用空格分开即可`[root@localhost ~]# rpm -ivh a.rpm b.rpm c.rpm`

```
rpm -ivh 包名
```

- i：安装install
- -v：显示详细信息verbose
- -h：显示安装进程ETA

例如：

```
[root@localhost ~]# rpm -ivh \
/mnt/cdrom/Packages/httpd-2.2.15-15.el6.centos.1.i686.rpm
Preparing...
####################
[100%]
1:httpd
####################
[100%]
```

注意，直到出现两个 100% 才是真正的安装成功，第一个 100% 仅表示完成了安装准备工作。

### 更新

```
[root@localhost ~]# rpm -Uvh 包全名
```

-U（大写）选项的含义是：如果该软件没安装过则直接安装；若没安装则升级至最新版本。

```
[root@localhost ~]# rpm -Fvh 包全名
```

-F（大写）选项的含义是：如果该软件没有安装，则不会安装，必须安装有较低版本才能升级。

### 卸载

> rpm卸载包之间需要考虑依赖性，按照层级卸载
>
> 强制卸载，可以使用通配符`rpm -e --nodeps mysql*`

```
[root@localhost ~]# rpm -e httpd
error: Failed dependencies:
httpd-mmn = 20051115 is needed by (installed) mod_wsgi-3.2-1.el6.i686
httpd-mmn = 20051115 is needed by (installed) php-5.3.3-3.el6_2.8.i686
httpd-mmn = 20051115 is needed by (installed) mod_ssl-1:2.2.15-15.el6.
centos.1.i686
httpd-mmn = 20051115 is needed by (installed) mod_perl-2.0.4-10.el6.i686
httpd = 2.2.15-15.el6.centos.1 is needed by (installed) httpd-manual-2.2.
15-15.el6.centos.1 .noarch
httpd is needed by (installed) webalizer-2.21_02-3.3.el6.i686
httpd is needed by (installed) mod_ssl-1:2.2.15-15.el6.centos.1.i686
httpd=0:2.2.15-15.el6.centos.1 is needed by(installed)mod_ssl-1:2.2.15-15.el6.centos.1.i686
```

### 查询

- `-q`

  查询软件包是否安装

  ```
  [root@localhost ~]# rpm -q httpd
  httpd-2.2.15-15.el6.centos.1.i686
  ```

- `-qa`

  查询所有安装的软件

  ```
  [root@cyberpelican local]# rpm -qa | grep  mysql
  zabbix-server-mysql-5.0.5-1.el7.x86_64
  ```

- `-qi`

  输出详细信息

  ```
  [root@localhost ~]# rpm -qi httpd
  Name : httpd Relocations:(not relocatable)
  #包名
  Version : 2.2.15 Vendor:CentOS
  #版本和厂商
  Release : 15.el6.centos.1 Build Date: 2012年02月14日星期二 06时27分1秒
  #发行版本和建立时间
  Install Date: 2013年01月07日星期一19时22分43秒
  Build Host:
  c6b18n2.bsys.dev.centos.org
  #安装时间
  Group : System Environment/Daemons Source RPM:
  httpd-2.2.15-15.el6.centos.1.src.rpm
  #组和源RPM包文件名
  Size : 2896132 License: ASL 2.0
  #软件包大小和许可协议
  Signature :RSA/SHA1,2012年02月14日星期二 19时11分00秒，Key ID
  0946fca2c105b9de
  #数字签名
  Packager：CentOS BuildSystem <http://bugs.centos.org>
  URL : http://httpd.apache.org/
  #厂商网址
  Summary : Apache HTTP Server
  #软件包说明
  Description:
  The Apache HTTP Server is a powerful, efficient, and extensible web server.
  #描述
  ```

- `-ql`

  查看软件包的安装路径

  ```
  [root@localhost ~]# rpm -ql httpd
  /etc/httpd
  /etc/httpd/conf
  /etc/httpd/conf.d
  /etc/httpd/conf.d/README
  /etc/httpd/conf.d/welcome.conf
  /etc/httpd/conf/httpd.conf
  /etc/httpd/conf/magic
  …省略部分输出…
  ```

- `-qR`

  查询依赖关系

  ```
  [root@cyberpelican local]# rpm -qR httpd|more
  /etc/mime.types
  system-logos >= 7.92.1-1
  httpd-tools = 2.4.6-90.el7.centos
  /usr/sbin/useradd
  /usr/sbin/groupadd
  systemd-units
  systemd-units
  systemd-units
  /bin/sh
  /bin/sh
  ```

  





