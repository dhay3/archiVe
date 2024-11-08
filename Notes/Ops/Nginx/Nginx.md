[TOC]

## 安装Nginx



##### 一. 

从nginx官网下载tar包

##### 二.

 1). 然后下载pcre文件依赖

```
wget http://downloads.sourceforge.net/project/pcre/pcre/8.37/pcre-8.37.tar.gz
```

2). 解压pcere

3).在pcere目录下 

```
./configure
```

4). 编译并安装pcere

```
make && make install
```

5) .安装 openssl 、 zlib 、 gcc 依赖

```
yum -y install make zlib zlib-devel gcc-c++ libtool openssl openssl-devel
```

##### 三. 

1). 进入nginx目录

```
./configure
```

2). 安装nginx

```
make && make install
```

3). ==进入 目录 /usr/local/nginx/sbin/nginx 启动服务==

#### 四

开放防火墙

```
firewall-cmd --add-service=http –permanent
firewall-cmd --add-port=80/tcp --permanent
firewall-cmd –reload
```

## 常用命令

### linux

以下命令都在`/usr/local/nginx/sbin`目录下执行

- 启动nginx

```
./nginx
```

- 关闭

```
./nginx -s stop
```

- 重启(linux下会热部署)

```
./nginx -s reload
```

- windows

一下命令都在解压后的窗口运行

- 启动nginx

```
 nginx.exe
```

- 关闭

```
nginx.exe -s stop 
```

- 重启

```
nginx.exe -s reload
```

- 查看版本

```
./nginx -v
```

## Nginx配置反向代理

访问linux主机地址自动会监听80端口,所以不用手动输入端口

##### 首先来到nginx安装目录

<img src="..\..\imgs\_Nginx\1.PNG"/>

==先复制一份原文件,然后再修改配置文件==

##### 修改配置文件

```
cd conf
vim nginx.conf
```

<img src="..\..\imgs\_Nginx\2.PNG"/>

配置server_name 和 proxy_pass 

表示访问server_name的域名,跳转到proxy_pass 指定的地址

==需要开放80和8080端口==

这里的localhost 指的是linux上的本机ip 127.0.0.1

##### 修改windows HOST 文件

来到如下位置

<img src="..\..\imgs\_Nginx\3.png"/>

再HOSTS中添加如下配置

```
192.168.1.130 www.chz.com
```

表示访问` www.chz.com` 域名匹配 192.168.1.130

这里的ip是linux 主机的ip

nginx接收到访问`www.chz.com`的请求, 就会转发到proxy_pass

此时,用户并不知道访问的具体服务器这就是反向代理

## 搞清楚proxy_pass 后的 /的区别

Nginx的[官网](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass)将proxy_pass分为两种类型：一种是只包含IP和端口号的（连端口之后的`/`也没有，这里要特别注意），比如`proxy_pass http://localhost:8080`，这种方式称为不带URI方式；另一种是在端口号之后有其他路径的，包含了只有单个`/`的如`proxy_pass http://localhost:8080/`，以及其他路径，比如`proxy_pass http://localhost:8080/abc`。

也即：`proxy_pass http://localhost:8080`和`proxy_pass http://localhost:8080/`(多了末尾的`/`)是不同的的处理方式，

而`proxy_pass http://localhost:8080/`和`proxy_pass http://localhost:8080/abc`是相同的处理方式。

==对于不带URI方式，nginx将会保留location中路径部分及以前路径，比如：==

```
 location /api1/ {
           proxy_pass http://localhost:8080;
        }
```

在访问`http://localhost/api1/xxx`时，会代理`http://localhost:8080/api1/xxx`

```
		location ~/ucenter/ {
			proxy_pass http://localhost:8006;
        }
```

在访问`http://localhost/api/ucenter/hello`, 会代理到`http:localhost:8006/api/ucenter/hello`

==对于带URI方式，nginx将使用诸如alias的替换方式对URL进行替换，并且这种替换只是字面上的替换，比如：==

```

location /api2/ {
           proxy_pass http://localhost:8080/;
        }
```

当访问`http://localhost/api2/xxx`时，`http://localhost/api2/`（注意最后的`/`）被替换成了`http://localhost:8080/`，然后再加上剩下的`xxx`，于是变成了`http://localhost:8080/xxx`。

```
 location /api5/ {
           proxy_pass http://localhost:8080/haha;
        }
```

当访问`http://localhost/api5/xxx`时，`http://localhost/api5/`被替换成了`http://localhost:8080/haha`，请注意这里`haha`后面没有`/`，然后再加上剩下的`xxx`，即`http://localhost:8080/haha`+`xxx`=`http://localhost:8080/hahaxxx`

## location中的正则匹配

参考: https://www.cnblogs.com/jpfss/p/10418150.html

- `=` 开头表示精确匹配
- `^~` 开头表示uri以某个常规字符串开头，理解为匹配 url路径即可。nginx不对url做编码，因此请求为/static/20%/aa，可以被规则^~ /static/ /aa匹配到（注意是空格）。以xx开头
- `~` 开头表示区分大小写的正则匹配           以xx结尾
- `~*` 开头表示不区分大小写的正则匹配        以xx结尾
- `!~`和`!~*`分别为区分大小写不匹配及不区分大小写不匹配 的正则
- `/` 通用匹配，任何请求都会匹配到。

## Nginx负载均衡

<img src="..\..\imgs\_Nginx\10.PNG"/>

访问server_name `www.domain.com` 时 通过nginx转发到 http//myproject  ,即 upstream

分发到三个不同端口的服务器来处理请求

```
http {
   upstream myproject {
 	server 127.0.0.1:8000 weight=3;
 	server 127.0.0.1:8001;
 	server 127.0.0.1:8002;
   }

   server {
 	listen 80;
 	server_name www.domain.com;
 	location / {
 		proxy_pass http//myproject;
	}
   }
}
```

###### ==配置负载均衡后要重启nginx的配置文件==

```
sbin/ngix -s reload
```

### ==负载均衡, 分配服务器策略==

##### 1). 轮询 (默认)

所有请求都按照时间顺序分配到不同的服务上，如果服务Down掉，可以自动剔除宕机的服务器

```
upstream  dalaoyang-server {
       server    localhost:8080;
       server    localhost:8081;
}
```

##### 2). 权重

指定每个服务的权重比例，weight和访问比率成正比，通常用于后端服务机器性能不统一，将性能好的分配权重高来发挥服务器最大性能

```
upstream  dalaoyang-server {
       server    localhost:8080 weight=1;
       server    localhost:8081 weight=2;
}
```

##### 3) ip_hash

每个请求都根据访问ip的hash结果分配，经过这样的处理，每个访客固定访问一个后端服务，如下配置（ip_hash可以和weight配合使用）。

ip_hash机制能够让某一客户机在相当长的一段时间内只访问固定的后端的某台真实的web服务器,这样会话就会得以保持,在网站页面进行login的时候就不会在后面的web服务器之间跳来跳去了,也不会出现登录一次的网站又提醒重新登录的情况.

```undefined
upstream  dalaoyang-server {
       ip_hash; 
       server    localhost:8080 weight=1;
       server    localhost:8081 weight=2;
}
```

##### 4) 最少连接

将请求分配到连接数最少的服务上。

```
upstream  dalaoyang-server {
       least_conn;
       server    localhost:8080 weight=1;
       server    localhost:8081 weight=2;
}
```

##### 5) fair

按后端服务器的响应时间来分配请求，响应时间短的优先分配。

```
upstream  dalaoyang-server {
       server    localhost:8080;
       server    localhost:8081;
       fair;  
}
```

## 搞清楚root后的/ 区别

```
location /img/ {
    alias /var/www/image/;
}
#若按照上述配置的话，则访问/img/目录里面的文件时，ningx会自动去/var/www/image/目录找文件
location /img/ {
    root /var/www/image;
}
#若按照这种配置的话，则访问/img/目录下的文件时，nginx会去/var/www/image/img/目录下找文件。]
```

==还有一个重要的区别是alias后面必须要用“/”结束，否则会找不到文件的。。。而root则可有可无~~==

## Nginx动静分离

<img src="..\..\imgs\_Nginx\9.png"/>

目录结构如下

<img src="..\..\imgs\_Nginx\7.PNG"/>

访问 192.168.1.130/www/a.html  --> /data/www/a.html ,  保留location (是linxu上的路径)

<img src="..\..\imgs\_Nginx\5.png"/>

==重启nginx==

如果开启autoindex on;

访问192.168.1.130/image/    能目录层级显示

不开启就不能显示层级目录

<img src="..\..\imgs\_Nginx\6.PNG"/>

## Nginx配置高可用的集群

每个nginx服务器下连接多个work,通过work访问tomcat和静态资源

<img src="..\..\imgs\_Nginx\8.png"/>

==主机宕机,同样可以访问,类似于redis的哨兵模式==

#### 安装

1)在两台虚拟机上安装nginx和keepalived

```
yum install keepalived –y
查看keeppalived是否安装完毕
rpm -qa keepalived
```

==安装目录 /etc/keepalived==

#### 配置

1)进入keepalived 目录,修改配置文件

配置如下 [keepalived.conf](image\keepalived.conf) 

2)将bash文件放在/usr/local/src下

 [nginx_check.sh](image\nginx_check.sh) 

3)启动两台主机的keepalived 和 nginx

```
/usr/local/nginx/sbin/nginx
systemctl start keepalived.service 
```

#### keepalived.conf配置文件

```
! Configuration File for keepalived

global_defs {
   notification_email {
     acassen@firewall.loc
     failover@firewall.loc
     sysadmin@firewall.loc
   }
   notification_email_from Alexandre.Cassen@firewall.loc
   smtp_server 192.168.1.100 #MASTER IP
   smtp_connect_timeout 30
   router_id LVS_DEVEL
}
vrrp_script chk_http_port { #用于检测nginx是否还存活
    script "/usr/local/src/nginx_check.sh"
	interval 2 				#(检测脚本执行的间隔)
	weight 2
}

vrrp_instance VI_1 {
    state MASTER 		  #如果是主服务器MASTER,如果是备份服务器改为BACKUP
    interface ens32 	  #网卡名字
    virtual_router_id 51  #主,备机的vitrual_router_id 必须相同
    priority 100  		  #主,备机取不同的优先级,主机值较大,备份机值较小
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1111
    }
    virtual_ipaddress {
        192.168.1.50   #虚拟的ip地址
        192.168.200.17 #与主机的网段不同不能联机
        192.168.200.18
    }
}
```

