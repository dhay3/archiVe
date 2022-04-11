# Nginx 日志服务

> 我们使用`locate nginx | grep log`来查看nginx日志具体存储位置

`access.log`用于记录执行的请求

```
[root@cyberpelican nginx]# cat /var/log/nginx/access.log
192.168.80.1 - - [25/Nov/2020:11:16:53 +0800] "GET / HTTP/1.1" 302 161 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36" "-"
127.0.0.1 - - [25/Nov/2020:12:27:56 +0800] "GET /favicon.ico HTTP/1.1" 404 3650 "-" "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"
127.0.0.1 - - [25/Nov/2020:12:29:41 +0800] "GET / HTTP/1.1" 200 83 "-" "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"

```

默认由IP，访问时间，user-agent组成。

## access_log

https://nginx.org/en/docs/http/ngx_http_log_module.html#access_log

 可以记录特定等级，形式，条件的日志。使用`access_log off`关闭access日志的记录。

```
Syntax:	access_log path [format [buffer=size] [gzip[=level]] [flush=time] [if=condition]];
access_log off;
Default:	
access_log logs/access.log combined;
Context:	http, server, location, if in location, limit_except
```

普通安装默认`access_log logs/access.log combined;`

yum或epel安装默认`access_log  /var/log/nginx/access.log  main;`

## access_format

> 注意一点按照编译顺序，`access_format`需要在`access_log`之前

https://nginx.org/en/docs/http/ngx_http_log_module.html#log_format

指定日志的格式

```
Syntax:	log_format name [escape=default|json|none] string ...;
Default:	
log_format combined "...";
Context:	http
```

可以使用[内建变量](https://nginx.org/en/docs/http/ngx_http_core_module.html#variables)

## Exmaple

```
http {
    keepalive_timeout   65;

    log_format pattern '$remote_addr-"$http_user_agent"';
    access_log /var/log/nginx/access.log pattern;
	...
	}
```

