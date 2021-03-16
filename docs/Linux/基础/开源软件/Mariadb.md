# Linux Mariadb

参考：

https://www.jianshu.com/p/16682746137b

MariaDB是MySQL的一个分支，为了防止MySQL被Oracle收购后闭源

## 安装

1. 检查是否已有与MySQL相关的东西

   ```
   rpm -qa | grep -i mysql
   ```

2. 卸载与mysql相关的软件

   ```
   rpm -e --nodeps mysql*
   ```

3. 安装mariadb

   ```
   yum install mariadb mariadb-server
   ```

4. 启动服务

   ```
   systemctl start mariadb && systemctl enable mariadb
   ```

5. 设置密码

   ```
   [root@cyberpelican local]# mysql_secure_installation 
   
   NOTE: RUNNING ALL PARTS OF THIS SCRIPT IS RECOMMENDED FOR ALL MariaDB
         SERVERS IN PRODUCTION USE!  PLEASE READ EACH STEP CAREFULLY!
   
   In order to log into MariaDB to secure it, we'll need the current
   password for the root user.  If you've just installed MariaDB, and
   you haven't set the root password yet, the password will be blank,
   so you should just press enter here.
   Setting the root password ensures that nobody can log into the MariaDB
   root user without the proper authorisation.
   
   Set root password? [Y/n] Y
   New password: 
   Re-enter new password: 
   Password updated successfully!
   
   ```

6. 其他设置

   ```
   Remove anonymous users? [Y/n] <– 是否删除匿名用户，回车
   Disallow root login remotely? [Y/n] <–是否禁止root远程登录,回车,
   Remove test database and access to it? [Y/n] <– 是否删除test数据库，回车
   Reload privilege tables now? [Y/n] <– 是否重新加载权限表，回车
   ```

















