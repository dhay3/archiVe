# Docker 搭建MySQL服务

> 开机自动启动
>
> docker update mysql  --restart=always

参考:

https://www.cnblogs.com/sablier/p/11605606.html

[TOC]

## 拉取镜像

```
docker pull mysql:5.7   # 拉取 mysql 5.7
docker pull mysql       # 拉取最新版mysql镜像
```

## 检查镜像

```
docker images
```

## 开启MySQL并建立目录映射

```
docker run -p 3306:3306 --name mysql \
-v /mydata/mysql/log:/var/log/mysql \
-v /mydata/mysql/data:/var/lib/mysql \
-v /mydata/mysql/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql:5.7
```

- -p：主机（这里指的是虚拟机）与容器端口的映射关系，":"前为主机目录，之后为容器目录
- -v：主机和容器的目录映射关系，":"前为主机目录，之后为容器目录
- -e ：启动参数
- --name： 指定容器名
- -d ：静默运行容器

修改主机的指定位置就会影响容器对应的位置，==需要重启容器==

> 注意mysql 8 以后没有var/log/mysql文件夹和var/lib/mysql，所以不会正常启动

```
docker run -p 3306:3306 --name mysql \
-v /mydata/mysql/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql:latest
```

如果还是无法解决请参考



https://blog.csdn.net/qq_41999034/article/details/106162366



https://blog.csdn.net/wzyaiwl/article/details/90293453?utm_medium=distribute.pc_relevant.none-task-blog-title-6&spm=1001.2101.3001.4242

## 连接MySQL

==注意==, 这里建议采用`exec`, 因为通过这种方式与`tty`交互, 容器不会退出, 如果采用`attach`, 退出`tty`容器就会退出

```
docker exec -it mysql /bin/bash //表示进入容器后执行 /bin/bash
mysql -uroot -proot //拉取的镜像本身是一个精简的Linux系统
```

<img src="..\..\imgs\_docker\Snipaste_2020-08-20_00-47-07.png" style="zoom:80%;" />

这里只做校验, 具体操作通过Navicat

<img src="..\..\imgs\_docker\Snipaste_2020-08-20_00-49-05.png"/>

==这里说明MySQL容器就一个精简的Linux容器==

## 配置MySQL文件

在主机上操作

```
yum -y install vim* //安装vim
vi my.conf
```

配置文件

```
[client]
default-character-set=utf8

[mysql]
default-character-set=utf8

[mysqld]
init_connect='SET collation_connection = utf8_unicode_ci'
init_connect='SET NAMES utf8'
character-set-server=utf8
collation-server=utf8_unicode_ci
skip-character-set-client-handshake
skip-name-resolve
```

重启MySQL容器

```
docker restart mysql //重启
docker start mysql //开启
```

重新进入MySQL容器

```
docker exec -it mysql /bin/bash
```

查看容器内配置文件

<img src="..\..\imgs\_docker\Snipaste_2020-08-20_01-28-44.png" style="zoom:80%;" />

## 对外开放3306端口

```
systemctl status firewalld //查看防火墙状态
systemctl start firewalld  //开启防火墙
firewall-cmd --add-port=3306/tcp --permanent //永久开放端口
firewall-cmd --add-service=http --permanent //开放http权
firewall-cmd --reload //重启防护墙,但是不会中断用户连接,不丢失状态
firewall-cmd --list-all //查看开放的端口
```

如果这里没有`-cmd`选项需要`yum install net-tools`

## Navicat创建连接

<img src="..\..\imgs\_docker\Snipaste_2020-08-20_00-36-48.png" style="zoom:80%;" />
