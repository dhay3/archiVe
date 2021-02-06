# Linux chkconfig

## 概述

参考：

https://zh.wikipedia.org/wiki/%E8%BF%90%E8%A1%8C%E7%BA%A7%E5%88%AB

https://developer.aliyun.com/article/537443

> chkconfig用于设置开机启动服务，在新版中被systemd替换

主要用于更新和查询系统服务运行级别。用于操作`/etc/rc[0-6].d`

运行级别：

- 0 ：停机，关机
- 1： 单用户，无网络连接，不运行守护进程，不允许非root用户登入
- 2：多用户，无网络连接，不运行守护进程
- 3：多用户，正常启动系统
- 4：用户自定义
- 5：多用户，带图形界面
- 6：重启

> 在全新的Linux  systemd中已经使用target代替Runlevel，如multi-user.target相当于init 3，graphical.target相当于init 5，但是SystemD仍然兼容运行级别（Runlevel）。目前绝大多数发行版已采用systemd代替UNIX System V。

## 参数

- `chkconfig`

  显示所有服务的运行级别，等价于`--list`

  ```
  [root@chz yum]# chkconfig |more
  
  Note: This output shows SysV services only and does not include native
        systemd services. SysV configuration data might be overridden by native
        systemd configuration.
  
        If you want to list systemd services use 'systemctl list-unit-files'.
        To see services enabled on particular target use
        'systemctl list-dependencies [target]'.
  
  netconsole     	0:off	1:off	2:off	3:off	4:off	5:off	6:off
  network        	0:off	1:off	2:on	3:on	4:on	5:on	6:off
  
  ```

- `chkconfig --level 35 <serverice> on`

  `--level 35`表示在3和5运行级别下==开机自动启动==，如果没有指定`--level`则表示2345，可选on, off, reset, resetpriorites

  ```
  [root@chz yum]# chkconfig httpd on
  Note: Forwarding request to 'systemctl enable httpd.service'.
  Created symlink from /etc/systemd/system/multi-user.target.wants/httpd.service to /usr/lib/systemd/system/httpd.service.
  [root@chz yum]# chkconfig 
  
  Note: This output shows SysV services only and does not include native
        systemd services. SysV configuration data might be overridden by native
        systemd configuration.
  
        If you want to list systemd services use 'systemctl list-unit-files'.
        To see services enabled on particular target use
        'systemctl list-dependencies [target]'.
  
  netconsole     	0:off	1:off	2:off	3:off	4:off	5:off	6:off
  network        	0:off	1:off	2:on	3:on	4:on	5:on	6:off
  [root@chz yum]# systemctl list-unit-files |grep http
  httpd.service                                 enabled 
  ```

- `chkconfig --add `

  添加服务让chkconfig管理

- `chkconfig --del`

  删除服务不让chkconfig管理
