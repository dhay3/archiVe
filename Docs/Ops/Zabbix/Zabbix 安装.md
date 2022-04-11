# Zabbix 安装

参考：

https://www.zabbix.com/cn/download?zabbix=5.0&os_distribution=centos&os_version=7&db=mysql&ws=apache

https://developer.aliyun.com/mirror/zabbix?spm=a2c6h.13651102.0.0.65101b11D12qAI

https://www.jianshu.com/p/16682746137b

https://www.zabbix.com/documentation/5.0/manual/installation/install_from_packages/rhel_centos

> 注意如果前端界面有问题，可能需要升级火狐浏览器

## Server

1. 添加官方镜像yum 

   ```
   rpm -Uvh https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm
   ```

2. 修改repo文件

   ```
   sed -i 's/repo.zabbix.com/mirrors.aliyun.com\/zabbix/g' /etc/yum.repos.d/zabbix.repo && yum clean all
   ```

3. 安装

   ```
   yum install zabbix-server-mysql zabbix-agent && yum install centos-release-scl
   ```

4. 安装Zabbix frornt，并修改配置`/etc/yum.repos.d/zabbix.repo`

   ```
   yum install zabbix-web-mysql-scl zabbix-apache-conf-scl
   
   ---
   
   [zabbix-frontend] #注意是zabbix-frontend块下
   ...
   enabled=1
   ...
   ```

5. 安装mysql组件，这里使用mariadb替代mysql

   ```
   yum install -y mariadb mariadb-server
   mysql_secure_installation
   ```

6. 创建数据库，用户，授权

   ```
   mysql -uroot -p
   
   your_password
   mysql> create database zabbix character set utf8 collate utf8_bin;
   mysql> create user zabbix@localhost identified by 'password';
   mysql> grant all privileges on zabbix.* to zabbix@localhost;
   mysql> quit;
   ```

7. 导入数据

   > 参考：https://www.cnblogs.com/opsprobe/p/10812274.html
   >
   > 如果出现`ERROR 1046 (3D000) at line 1: No database selected`
   >
   > vim /usr/share/doc/zabbix-server-mysql-4.0.7/create.sql.gz
   >
   > 添加
   >
   > use zabbix;

   ```
   zcat /usr/share/doc/zabbix-server-mysql*/create.sql.gz | mysql -uzabbix -p  your_password
   ```

8. 为zabbix server指定数据库的密码

   ```
   vim /etc/zabbix/zabbix_server.conf
   DBPassword=your_password
   ```

9. 配置前端PHP时区，需要去掉前面的分号，`/etc/opt/rh/rh-php72/php-fpm.d/zabbix.conf`

   ```
    php_value[date.timezone] = Asia/Shanghai
   ```

10. 开启服务并置为开机自启

    ```
    # systemctl restart zabbix-server zabbix-agent httpd rh-php72-php-fpm
    # systemctl enable zabbix-server zabbix-agent httpd rh-php72-php-fpm
    ```

11. 访问`http://host/zabbix`，默认超级用户账户为`Admin`，密码为`zabbix`

12. 修改超级用户密码

    ```mermaid
    graph LR
    sidebar --> Administration --> Users --> Admin --> a(change password)
    ```

13. 进入zabbix管理界面出现`zabbix server is not running: the information displayed may not be current`

    我们使用`journalctl -xe`来查看日志

    - 方法一：关闭SELinux，不推荐

    - 方法二：

      参考

      https://www.zabbix.com/forum/zabbix-help/367261-selinux-and-zabbix

      https://www.zabbix.com/documentation/5.0/manual/installation/install_from_packages/rhel_centos

      ```
      getsebool | grep 
      setsebool -P httpd_can_connect_zabbix on
      setsebool -P httpd_can_network_connect_db on
      systemctl restart httpd
      ```

## agent

1. 配置yum源

   ```
   rpm -Uvh https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm
   yum clean all
   ```

2. 安装zabbix agent

   ```
   yum install zabbix-agent
   ```

3. 开启服务

   ```
   yum start zabbix-agent && yum enable zabbix-agent
   ```

   
