# Apache2概述

参考：

https://yungyuc.github.io/oldtech/debian/ApacheConf.html

https://www.yiibai.com/apache_http/apache-configuration-files.html

> 在apt源中全部用apache2替代apache，apache2与apache在架构和模块上存在较大的差异
>
> 本教程采用Apache 2.4

## 文件目录

==可以具体参考如下链接==

https://cwiki.apache.org/confluence/display/httpd/DistrosDefaultLayout

使用`whereis apache2`来查看apache2的具体位置

```
#	/etc/apache2/
#	|-- apache2.conf
#	|	`--  ports.conf
#	|-- mods-enabled
#	|	|-- *.load
#	|	`-- *.conf
#	|-- conf-enabled
#	|	`-- *.conf
# 	`-- sites-enabled
#	 	`-- *.conf
```

- envvars

  apache环境变量

- ports.conf

  指定apache监听的端口。可以监听多个端口

  ```
  root in /etc/apache2 λ netstat -lnpt  | grep apache2
  tcp6       0      0 :::8868                 :::*                    LISTEN      1748/apache2        
  tcp6       0      0 :::8848                 :::*                    LISTEN      1748/apache2      
  ```

  ```
  root in /etc/apache2 λ cat ports.conf 
  # If you just change the port or add more ports here, you will likely also
  # have to change the VirtualHost statement in
  # /etc/apache2/sites-enabled/000-default.conf
  
  Listen 80
  <IfModule ssl_module>
  	Listen 443
  </IfModule>
  
  <IfModule mod_gnutls.c>
  	Listen 443
  </IfModule>
  
  # vim: syntax=apache ts=4 sw=4 sts=4 sr noet 
  ```

  ==如果没有指定Listen 端口，就不能通过访问本机端口来访问apache==

- sites-available

  可用虚拟机

- sites-enable

  已启用虚拟机

> 模块的启用建立软连接，模块关闭删除软连接
>
> ```
> ln -s /etc/apache2/mods-available/php5.load /etc/apache2/mods-enabled/php5.load
> rm /etc/apache2/mods-enabled/php5.load
> ```
>
> 在Ubuntu中可以使用`a2dismod`删除模块，`a2enmod`开启模块。模块名为配置文件去掉`.conf`后缀

- mods-avaliable

  可用模块

- mods-enable

  已启用模块

## Directives & Sections

==`_default_`等价于`*`==

Directives是最基础的，如果想要Directives局部生效，可以将其放在Sections下

### Directives

参考：https://httpd.apache.org/docs/2.4/mod/quickreference.html

### Sections

如果只想要指令在局部生效，可以放在`<Directory>`，`<DirectoryMatch>`，`<Files>`，`<FilesMatch>`，`<Location>`和`<LocationMatch>`指令块中（只在指定条件下生效）。

https://httpd.apache.org/docs/2.4/sections.html

==指令块后添加Match可以使用regex==

- IfDefine：只有在命令行在中使用特定参数，section生效

- IfModule：只有在启用指定模块，section生效

- IfVersion：只有在指定apache version，section生效

- Directory：只有在本机特定的目录，section生效。可以内嵌Files

- Files：只有在本机特定的文件，section生效

- Location：只有在请求中含有特定URI，section生效

- VirtualHost：https://httpd.apache.org/docs/2.4/mod/core.html#virtualhost

  可以在一个Apache2软件中配置多个VirtualHost(一般为内网的主机)

  ```
  NameVirtualHost *:80
  <VirtualHost 192.168.0.108:80>
      ServerAdmin webmaster@yiibai.com
      DocumentRoot /var/www/html/example1_com_dir 
      ServerName www.example1.com
  </VirtualHost>
  <VirtualHost 192.168.0.108:80>
      ServerAdmin admin@yiibai.com
      DocumentRoot /var/www/html/example2_com_dir
      ServerName www.example2.com
  </VirtualHost>
  ```

  上述配置表示，如果apache2收到来自192.168.0.108:80的请求，就应用下面的section，访问本机`/var/www/html/example_com/dir`



