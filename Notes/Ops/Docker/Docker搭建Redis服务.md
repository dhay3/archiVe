# Docker搭建Redis服务

> 开机自动启动
>
> docker update redis  --restart=always

参考：

http://www.redis.cn/download.html

[TOC]

## 拉取镜像

```
docker pull redis
```

## 检查镜像

```
docker images
```

## 开启Redis并建立目录映射

==注意==Redis镜像不一定带有`redis.conf`文件, 所以系统会认为主机的`redis.conf`是文件，所以容器也会认为`redis.conf`为文件，所以在配置映射时，一定要==先创建配置文件==

```
mkdir -p /mydata/redis/conf #-p表示级联创建目录
touch /mydata/redis/conf/redis.conf
```

开启Redis

```
docker run -p 6379:6379 --name redis \
-v /mydata/redis/data:/data \
-v/mydata/redis/conf/redis.conf:/etc/redis/redis.conf \
-d redis redis-server /etc/redis/redis.conf
```

## 连接Redis

这里无需使用redis-server开启redis服务器

```
docker exec -it redis /bin/bash
redis-cli //开启redis客户端
```

测试

```
127.0.0.1:6379> set a 10
OK
127.0.0.1:6379> get a
"10"

```

==注意==，这里redis没有配置文件

## 修改配置文件

https://raw.githubusercontent.com/redis/redis/6.0/redis.conf

==由于容器每次重启都会是一个新的==，所以需要开启redis的持久化， 这里使用aof

```
appendonly yes
```

## RedisDesktopManager创建连接

<img src="..\..\imgs\_docker\Snipaste_2020-08-20_14-59-01.png" style="zoom:80%;" />

