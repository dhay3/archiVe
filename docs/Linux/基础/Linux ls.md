# linux ls

> 实际上调用`printf  '%s\n' *`

参考：

https://blog.csdn.net/JenMinZhang/article/details/9816853

## 概述

用于查看指定文件夹下的目录与文件

## 参数

- `-l`

  以list形式显示内容，centos中配置了`alias ll=ls -l --color`

  ```
  drwxr-x--- 2 root              adm    4096 2013-08-07 11:03 apache2
  drwxr-xr-x 2 root              root   4096 2013-08-07 09:43 apparmor
  drwxr-xr-x 2 root              root   4096 2013-08-07 09:44 apt
  -rw-r----- 1 syslog            adm   16802 2013-08-07 14:30 auth.log
  -rw-r--r-- 1 root              root    642 2013-08-07 11:03 boot.log
  drwxr-xr-x 2 root              root   4096 2013-08-06 18:34 ConsoleKit
  drwxr-xr-x 2 root              root   4096 2013-08-07 09:44 cups
  -rw-r----- 1 syslog            adm   10824 2013-08-07 11:08 daemon.log
  drwxr-xr-x 2 root              root   4096 2013-08-07 09:45 dbconfig-common
  -rw-r----- 1 syslog            adm   21582 2013-08-07 11:03 debug
  drwxr-xr-x 2 root              root   4096 2013-08-07 09:45 dist-upgrade
  -rw-r--r-- 1 root              adm   59891 2013-08-07 11:03 dmesg
  ```

  第一列：标识文件的类型，组与用户权限

  第二列：文件的链接数

  第三列：文件的所有者

  第四列：文件所在的组（有root组）

  第五列：文件长度（大小）

  第六列：文件最后修改的时间

  第七列：文件的名称

  **文件类型**

  - d：目录

    -：普通文件

    l：链接

    s：socket

    p：name pipe

    b：block device

    c：character device

  **文件权限**

  - r：表示“可读”，用数字4表示
  - w：表示“可写”，用数字2表示
  - x：表示“可执行”，用数字1表示

- -h

  以认类可读的显示

  ```
  [root@chz network-scripts]# ll -h
  total 256K
  -rw-r--r--. 1 root root  348 Oct 24 09:16 ifcfg-ens33
  -rw-r--r--. 1 root root  320 Oct 24 09:17 ifcfg-ens34
  -rw-r--r--. 1 root root  254 Mar 29  2019 ifcfg-lo
  lrwxrwxrwx. 1 root root   24 Aug 24 07:51 ifdown -> ../../../usr/sbin/ifdown
  -rwxr-xr-x. 1 root root  654 Mar 29  2019 ifdown-bnep
  ```

- -Z

  显示文件的security context

  ```
  [root@chz network-scripts]# ls -Z
  -rw-r--r--. root root system_u:object_r:net_conf_t:s0  ifcfg-ens33
  -rw-r--r--. root root system_u:object_r:net_conf_t:s0  ifcfg-ens34
  -rw-r--r--. root root system_u:object_r:net_conf_t:s0  ifcfg-lo
  lrwxrwxrwx. root root system_u:object_r:bin_t:s0       ifdown -> ../../../usr/sbin/ifdown
  ```

- -t

  按照文件修改的时间排序，最新修改的放在前面

- -r

  对排序后的文件逆序

