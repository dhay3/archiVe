# MySQL Installation

## Binary Installation

### [Download](https://www.mysql.com/downloads/)


![](https://github.com/dhay3/image-repo/raw/master/20230802/2023-08-02_20-17.73ke7kx6670g.webp)

选择 [MySQL Community Server](https://dev.mysql.com/downloads/mysql/) 按照系统的架构选择版本即可

### [Installation](https://dev.mysql.com/doc/refman/5.7/en/binary-installation.html)

1. 将 tar 包解压到自定义目录

   ```
   mv mysql-5.7.43-linux-glibc2.12-x86_64.tar.gz /sharing/env
   tar -xzvf mysql-5.7.43-linux-glibc2.12-x86_64.tar.gz
   mkdir mysql 
   mv mysql-5.7.43-linux-glibc2.12-x86_64 mysql
   ```

2. 创建配置文件

   ```
   cd mysql/mysql-5.7.43-linux-glibc2.12-x86_64
   mkdir etc
   cd etc
   touch my.cnf
   ```

   内容如下

   ```
   [client]
   port=3306
   socket=/tmp/mysql.sock
   [mysqld]
   port = 3306
   log-error=/sharing/env/mysql/mysql-5.7.43-linux-glibc2.12-x86_64/error.log
   socket=/tmp/mysql.sock
   basedir=/sharing/env/mysql/mysql-5.7.43-linux-glibc2.12-x86_64 
   datadir=/sharing/env/mysql/mysql-5.7.43-linux-glibc2.12-x86_64/data
   ```

3. 初始化 data 目录

   ```
   bin/mysqld  --defaults-file=etc/my.cnf --initialize
   ```

   > `--initialize` 必须在最后，注意这里需要记录一下临时密码，使用 `mysql` 时需要使用
   >
   > 也可以使用 `--initialize-insecure`, 默认就不会生成密码，使用 `mysql` 时无需指定密码，但是需要和 `--skip-password` 一起使用

4. 建立软连接

   ```
   cd /usr/local/bin 
   sudo ln -s /sharing/env/mysql/mysql-5.7.43-linux-glibc2.12-x86_64/bin/mysql mysql
   ```

5. 启动 mysqld

   ```
   cd -
   bin/mysqld  --defaults-file=etc/my.cnf
   ```

6. 修改初始密码

   ```
   bin/mysql -u root -p
   alter user 'root'@'localhost' identified by 'password';
   exit
   ```

   > 如果出现 `./mysql: error while loading shared libraries: libncurses.so.5: cannot open shared object file: No such file or directory` 安装 `ncurses5-compat-libs` 即可
   >
   > 例如 Arch
   >
   > `yay -S aur/ncurses5-compat-libs`

6. 校验密码是否成功

### [Systemd](https://dev.mysql.com/doc/mysql-secure-deployment-guide/8.0/en/secure-deployment-post-install.html)

在 `/etc/systemd/system/` 下创建一个 `mysqld.service`

```
[Unit]
Description=MySQL Server
Documentation=man:mysqld(8)
Documentation=http://dev.mysql.com/doc/refman/en/using-systemd.html
After=network.target
After=syslog.target

[Install]
WantedBy=multi-user.target

[Service]
User=cpl
Group=cpl
Type=notify
TimeoutSec=0
ExecStart=/sharing/env/mysql/mysql-5.7.43-linux-glibc2.12-x86_64/bin/mysqld --defaults-file=/sharing/env/mysql/mysql-5.7.43-linux-glibc2.12-x86_64/etc/my.cnf --daemonize $MYSQLD_OPTS 

EnvironmentFile=-/etc/sysconfig/mysql

LimitNOFILE = 10000

Restart=on-failure

RestartPreventExitStatus=1

# Set environment variable MYSQLD_PARENT_PID. This is required for restart.
Environment=MYSQLD_PARENT_PID=1

PrivateTmp=false
```

**references**

1. [^https://dev.mysql.com/doc/refman/5.7/en/binary-installation.html]
2. [^https://dev.mysql.com/doc/refman/5.7/en/starting-server.html]


