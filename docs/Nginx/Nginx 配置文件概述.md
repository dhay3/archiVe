# Nginx 配置文件概述

参考：

https://docs.nginx.com/nginx/admin-guide/basic-functionality/managing-configuration-files/

https://www.yiibai.com/nginx/nginx-web-server.html#article-start

> ==通过`https://nginx.org/en/docs/`的Modules reference模块中的`Alphabetical index of directives`来查看具体的指令集==

## Directives

```
Syntax:	index file ...;
Default:	
index index.html;
Context:	http, server, location
```

- Syntax：使用语法
- Default：默认指令
- Context：指令允许出现的Context

==配置文件由许多的Directives组成==

- 简单指令

  ```
  user nginx;
  worker_processes auto;
  error_log /var/log/nginx/error.log;
  pid /run/nginx.pid;
  ```

- 指令集合，以curly brace为一组

  ```
  location / {
            proxy_pass http://baidu.com;
          }
  ```

## include

为了方便管理配置文件，可以将一些特殊的配置文件存储在`/etc/nginx/conf.d`下，然后使用`include`引入到`/etc/nginx/nginx.conf`主配置文件中

```
    # Load modular configuration files from the /etc/nginx/conf.d directory.
    # See http://nginx.org/en/docs/ngx_core_module.html#include
    # for more information.
    include /etc/nginx/conf.d/*.conf;
```

例如默认配置文件中，使用模式扩展将`/etc/nginx/conf.d`所有的配置文件引入

## Context

高等级Directives，称为Context

### events

连接处理

```
events {
    worker_connections 1024;
}
```

### http

管理http协议流量

### mail

管理mail协议流量

### stream

管理tcp和udp协议流量

## Virtual Servers

在Context中你可以有一个或多个`server blocks`来控制请求

- http

  指定traffic流向某一个IP或domain

  ```
   server {
          listen       80 default_server;
          listen       [::]:80 default_server;
          server_name  _;
          root         /usr/share/nginx/html;
  
          # Load configuration files for the default server block.
          include /etc/nginx/default.d/*.conf;
  
          location / {
            proxy_pass http://baidu.com;
          }
  
          error_page 404 /404.html;
          location = /404.html {
          }
  
          error_page 500 502 503 504 /50x.html;
          location = /50x.html {
          }
  ```

- mail || stream

  指定traffic流向某一个端口或套接字

## 代理

```
server {
    location / {
        proxy_pass http://localhost:8080;
    }
}
```

当访问根路径时转发到本地8080端口

## 变量

https://nginx.org/en/docs/http/ngx_http_core_module.html?&_ga=1.51463017.1509956953.1490042234#variables

## rewrite

https://segmentfault.com/a/1190000008102599

## 实例

https://www.cnblogs.com/54chensongxia/p/12938929.html

https://note.youdao.com/ynoteshare1/index.html?id=0f7ae7918a68fc7650f362ef4dc325bf&type=note

- 全局块：==相当于全局变量，如果有的指令块中没有指定，就使用全局块中的值==
- events块：Nginx服务器与用户的网络连接
- http块：配置代理，缓存和日志
- mail块：mail协议流量设置
- stream：tcp协议流量设置

```
#nginx worker进程
user www-data;

#工作进程数
worker_processes 2;

#工作进程与cpu映射
worker_cpu_affinity 0001 0010;

#全局错误日志<path level>
error_log logs/error.log info

#进程ID存储位置
pid /run/nginx.pid;

#包括的其他配置文件
include /etc/nginx/modules-enabled/*.conf;

#work processes能打开文件的最大数该值应该与 ulimit -n 显示的值相同
worker_rlimit_nofile 1024;

#是否以守候线程的方式运行nginx，只在生产环境下生效
daemon on;

events {
	#单个work process 连接的最大并发数，不能超过worker_rlimit_nofile的值
	worker_connections 768;
	
	#单个work process是否并行
	# multi_accept on;
	
	#指定使用哪种网络IO模型，可选择的内容有：select、poll、kqueue、epoll、rtsig、/dev/poll以及eventport
	use epoll;
}

http {
	
	#是否开启sendfile 高效传输方式
	sendfile on;
	#防止网络阻塞
	tcp_nopush on;
	tcp_nodelay on;
	
	#指定客户端请求单个文件的最大字节数
	client_max_body_size 8m;
	#指定来自客户端请求头的headerbuffer大小
	client_header_buffer_size 32k;
	#指定客户端请求中较大的消息头的缓存最大数量和大小
	large_client_header_buffers 4 64k;
	#hashtables的最大值
	types_hash_max_size 2048;
	#是否在错误页面显示nginx的版本
	server_tokens on;
	#存储服务器名字的hashtable最大容量
	# server_names_hash_bucket_size 64;
	# server_name_in_redirect off;
	
	#包含的其他配置文件
	include /etc/nginx/mime.types;
	
	#默认文件类型
	default_type application/octet-stream;
	#客户端连接超时时间，单位秒
	keepalive_timeout 60;
	#客户端请求主体读取超时时间
	client_body_timeout 10;
	#响应客户端超时时间
	send_timeout 10;
	
	#fastcgi相关参数时为了改善网站的性能；减少资源占用，提高访问速度。
	fastcgi_connect_timeout 300;
    fastcgi_send_timeout 300;
    fastcgi_read_timeout 300;
    fastcgi_buffer_size 64k;
    fastcgi_buffers 4 64k;
    fastcgi_busy_buffers_size 128k;
    fastcgi_temp_file_write_size 128k;


	##
	# SSL Settings
	##

	ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3; # Dropping SSLv3, ref: POODLE
	ssl_prefer_server_ciphers on;

	##
	# Logging Settings
	##
    #连接日志的存储位置
	access_log /var/log/nginx/access.log;
	#错误日志存储位置
	error_log /var/log/nginx/error.log;
	#默认编码
	charset utf-8;
	  log_format  main  '$remote_addr - $remote_user [$time_local] "$request" ''$status $body_bytes_sent "$http_referer" ''"$http_user_agent" "$http_x_forwarded_for"';
    
    #定义日志的格式。后面定义要输出的内容。
    #1.$remote_addr 与$http_x_forwarded_for 用以记录客户端的ip地址；
    #2.$remote_user ：用来记录客户端用户名称；
    #3.$time_local ：用来记录访问时间与时区；
    #4.$request  ：用来记录请求的url与http协议；
    #5.$status ：用来记录请求状态； 
    #6.$body_bytes_sent ：记录发送给客户端文件主体内容大小；
    #7.$http_referer ：用来记录从那个页面链接访问过来的；
    #8.$http_user_agent ：记录客户端浏览器的相关信息

	##
	# Gzip Settings
	##
	#开启gzip压缩输出
	gzip on;
	# gzip_vary on;
	# gzip_proxied any;
	# gzip_comp_level 6;
	#压缩缓冲区
	# gzip_buffers 16 8k;
	# gzip_http_version 1.1;
	# gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;

	##
	# Virtual Host Configs
	##

	include /etc/nginx/conf.d/*.conf;
	include /etc/nginx/sites-enabled/*;
}

   server
    {
        #监听端口
        listen 80;

        #域名可以有多个，用空格隔开，如果请求的域名不匹配，就会终止当前指令块
        server_name www.w3cschool.cn w3cschool.cn;
        #以下指令做为默认值
        index index.html index.htm index.php;
        root /data/www/w3cschool;

        #对******进行负载均衡
        location ~ .*.(php|php5)?$
        {
            fastcgi_pass 127.0.0.1:9000;
            fastcgi_index index.php;
            include fastcgi.conf;
        }
         
        #图片缓存时间设置
        location ~ .*.(gif|jpg|jpeg|png|bmp|swf)$
        {
            expires 10d;
        }
         
        #JS和CSS缓存时间设置
        location ~ .*.(js|css)?$
        {
            expires 1h;
        }
         
        #局部变量
        #日志格式设定
        #$remote_addr与$http_x_forwarded_for用以记录客户端的ip地址；
        #$remote_user：用来记录客户端用户名称；
        #$time_local： 用来记录访问时间与时区；
        #$request： 用来记录请求的url与http协议；
        #$status： 用来记录请求状态；成功是200，
        #$body_bytes_sent ：记录发送给客户端文件主体内容大小；
        #$http_referer：用来记录从那个页面链接访问过来的；
        #$http_user_agent：记录客户浏览器的相关信息；
        #通常web服务器放在反向代理的后面，这样就不能获取到客户的IP地址了，通过$remote_add拿到的IP地址是反向代理服务器的iP地址。反向代理服务器在转发请求的http头信息中，可以增加x_forwarded_for信息，用以记录原有客户端的IP地址和原来客户端的请求的服务器地址。
        log_format access '$remote_addr - $remote_user [$time_local] "$request" '
        '$status $body_bytes_sent "$http_referer" '
        '"$http_user_agent" $http_x_forwarded_for';
         
        #定义本虚拟主机的访问日志
        access_log  /usr/local/nginx/logs/host.access.log  main;
        access_log  /usr/local/nginx/logs/host.access.404.log  log404;
         
        #对 "/" 启用反向代理
        location / {
            proxy_pass http://127.0.0.1:88;
            proxy_redirect off;
            proxy_set_header X-Real-IP $remote_addr;
             
            #后端的Web服务器可以通过X-Forwarded-For获取用户真实IP
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
             
            #以下是一些反向代理的配置，可选。
            proxy_set_header Host $host;

            #允许客户端请求的最大单文件字节数
            client_max_body_size 10m;

            #缓冲区代理缓冲用户端请求的最大字节数，
            #如果把它设置为比较大的数值，例如256k，那么，无论使用firefox还是IE浏览器，来提交任意小于256k的图片，都很正常。如果注释该指令，使用默认的client_body_buffer_size设置，也就是操作系统页面大小的两倍，8k或者16k，问题就出现了。
            #无论使用firefox4.0还是IE8.0，提交一个比较大，200k左右的图片，都返回500 Internal Server Error错误
            client_body_buffer_size 128k;

            #表示使nginx阻止HTTP应答代码为400或者更高的应答。
            proxy_intercept_errors on;

            #后端服务器连接的超时时间_发起握手等候响应超时时间
            #nginx跟后端服务器连接超时时间(代理连接超时)
            proxy_connect_timeout 90;

            #后端服务器数据回传时间(代理发送超时)
            #后端服务器数据回传时间_就是在规定时间之内后端服务器必须传完所有的数据
            proxy_send_timeout 90;

            #连接成功后，后端服务器响应时间(代理接收超时)
            #连接成功后_等候后端服务器响应时间_其实已经进入后端的排队之中等候处理（也可以说是后端服务器处理请求的时间）
            proxy_read_timeout 90;

            #设置代理服务器（nginx）保存用户头信息的缓冲区大小
            #设置从被代理服务器读取的第一部分应答的缓冲区大小，通常情况下这部分应答中包含一个小的应答头，默认情况下这个值的大小为指令proxy_buffers中指定的一个缓冲区的大小，不过可以将其设置为更小
            proxy_buffer_size 4k;

            #proxy_buffers缓冲区，网页平均在32k以下的设置
            #设置用于读取应答（来自被代理服务器）的缓冲区数目和大小，默认情况也为分页大小，根据操作系统的不同可能是4k或者8k
            proxy_buffers 4 32k;

            #高负荷下缓冲大小（proxy_buffers*2）
            proxy_busy_buffers_size 64k;

            #设置在写入proxy_temp_path时数据的大小，预防一个工作进程在传递文件时阻塞太长
            #设定缓存文件夹大小，大于这个值，将从upstream服务器传
            proxy_temp_file_write_size 64k;
        }
         
         
        #设定查看Nginx状态的地址
        location /NginxStatus {
            stub_status on;
            access_log on;
            auth_basic "NginxStatus";
            auth_basic_user_file confpasswd;
            #htpasswd文件的内容可以用apache提供的htpasswd工具来产生。
        }
         
        #本地动静分离反向代理配置
        #所有jsp的页面均交由tomcat或resin处理
        location ~ .(jsp|jspx|do)?$ {
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_pass http://127.0.0.1:8080;
        }
         
        #所有静态文件由nginx直接读取不经过tomcat或resin
        location ~ .*.(htm|html|gif|jpg|jpeg|png|bmp|swf|ioc|rar|zip|txt|flv|mid|doc|ppt|
        pdf|xls|mp3|wma)$
        {
            expires 15d; 
        }
         
        location ~ .*.(js|css)?$
        {
            expires 1h;
        }
    }
```

