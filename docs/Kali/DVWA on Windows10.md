# DVWA on Windows

## XAMPP

[下载地址](https://sourceforge.net/projects/xampp/)

==如果本地还没有Apache web server, MySQL, PHP, Perl 或是 a FTP服务器可以使用XAMP, 具体可以参考:==

如果本地安装了这些服务器, 可能会存在冲突. 如果存在冲突请使用虚拟机安装

https://www.youtube.com/watch?v=cak2lQvBRAo

如果出现如下图所示, 避免将文件安装到下图路径中

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_14-35-34.png" style="zoom:80%;" />

选择需要的组件

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_14-41-02.png" style="zoom:80%;" />

安装完成后, 进入界面, 开启Apache, 访问`localhost`, 出现下图界面即表示安装成功。将`E:\XAMPP\htdocs`

中的所有内容删除。然后重新访问`localhost`, 出现下图即可

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_14-55-04.png" style="zoom:80%;" />



## DVWA

`git clone https://github.com/ethicalhack3r/DVWA.git`

将文件解压， 然后将里面的内容copy新建的文件夹dvwa到`E:\XAMPP\htdocs`中，重新访问`localhost`, 出现下图即可

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_15-00-23.png" alt="Snipaste_2020-09-04_15-00-23" style="zoom:80%;" />

然后访问文件提示`Copy config/config.inc.php.dist to config/config.inc.php and configure to your environment.`将改文件中的指定的文件复制一份命名位`config.inc.php`, 然后重新检查。进入页面会发现无法连接MySQL。这里需要修改`config.inc.php`中的配置

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_15-07-20.png" style="zoom:80%;" />



然后修改Apache->config->PHP, 将url_include置为On

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_16-15-26.png" style="zoom:80%;" />



配置完成后就会进入登入页面, 默认账户与密码分别为`admin`和`password`

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-04_17-11-53.png" style="zoom:80%;" />

## 错误

Could not connect to the database service.
Please check the config file.
Database Error #1045: Access denied for user 'dvwa'@'localhost' (using password: NO).

出现上述错误, 查看phpmyadmin

<img src="..\..\imgs\_Kali\dvwa\Snipaste_2020-09-06_13-31-31.png"/>

查看账户的登入用户名, 将`$_DVWA[ 'db_user' ] `改为相同的用户名即可

