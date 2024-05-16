# server_name

## 0x00 Overview

```
Syntax: 	server_name name ...;
Default: 	server_name "";
Context: 	server
```

指定 Virtual Server 的名字，Nginx 会根据 HTTP 请求中的 HOST 字段对比 `server_name` 的值判断使用具体的那个  VS(server directive block)

### server_name ...

name 是一个变参，可以指定多个值

```
server {
    server_name example.com www.example.com;
}
```

### server_name *.com

server_name 中可以包含 `*`(wildcard)，例如

```
server {
    server_name example.com *.example.com www.example.*;
}
```

也可以合并成一个

```
server {
    server_name .example.com;
}
```

也可以和 PCRE 一起使用，例如

```
server {
    server_name ~^(www\.)?(.+)$;

    location / {
        root /sites/$2;
    }
}

server {
    server_name _;

    location / {
        root /sites/default;
    }
}
```

### server_name  ""

也可以将 servername 设置为 `""` ，表示允许处理没有 HOST 字段的 HTTP 请求，实际测试发现如果 HTTP 请求不带 Host 字段，会报错返回 400

```
worker_processes auto;
events {
}
http {
  server {
    listen 80;
    server_name "";
    location / {
      return 200;
    }
  }
}

$ curl -v -H'Host:' localhost
* About to connect() to localhost port 80 (#0)
*   Trying ::1...
* Connection refused
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.29.0
> Accept: */*
>
< HTTP/1.1 400 Bad Request
< Server: nginx/1.24.0
< Date: Tue, 14 May 2024 05:47:51 GMT
< Content-Type: text/html
< Content-Length: 157
< Connection: close
```

### Server_name _

但是你也会看到许多的配置文件使用 `_` 来代替 name，例如

```
server {
    listen       80  default_server;
    server_name  _;
    return       444;
}
```

这表示 catch-all server(匹配所有的 Host)，实际上只要是无效的域名都可以，例如  `!@#` 

```
worker_processes auto;
events {
}
http {
  server {
    listen 80;
    server_name @;
    location / {
      return 200;
    }
  }
}
[root@centos nginx]# curl -v  localhost
* About to connect() to localhost port 80 (#0)
*   Trying ::1...
* Connection refused
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.29.0
> Host: localhost
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: nginx/1.24.0
< Date: Tue, 14 May 2024 05:53:28 GMT
< Content-Type: text/plain
< Content-Length: 0
< Connection: keep-alive
<
* Connection #0 to host localhost left intact
```

## 0x01 server_name Searching

在 Nginx 收到 HTTP 请求时，是如何判断使用具体那个 server 

例如 有如下配置

```
server {
    listen      80;
    server_name example.org www.example.org;
    ...
}

server {
    listen      80;
    server_name example.net www.example.net;
    ...
}

server {
    listen      80;
    server_name example.com www.example.com;
    ...
}
```

Nginx 首先会从 HTTP 请求中拆解 HOST 字段，用于决定那个 server directive block 来处理请求。如果 HOST 为 `example.org` 就会使用第一个 server directive block，即 HOST 精确匹配 `server_name`。如果 HOST 不匹配任何一个 server directive block 中的 `server_name`  或者请求不包含 HOST 字段，就会默认使用按照配置中出现顺序的第一个 server directive block 来处理请求

这种默认行为也可以通过 `listen` directive 的 `default_server` 来显式声明如果没有匹配或者不包含 HOST 字段的请求，使用该 server directive block

```
server {
    listen      80;
    server_name example.org www.example.org;
    ...
}

server {
    listen      80 default_server;
    server_name example.net www.example.net;
    ...
}

server {
    listen      80;
    server_name example.com www.example.com;
    ...
}
```

如果请求为 `examples.org` 就会匹配第二个 server directive block

## 0x02 Matches More than one server_name

HOST 也可能存在匹配多个 `server_name` 的情况，例如

```
worker_processes auto;
events {
}
http {
  server {
    listen 80;
    server_name ~example[s]?\.com;
    location / {
      return 200;
    }
  }
  server {
    listen 80;
    server_name example.com;
    location / {
      return 300;
    }
  }
  server {
    listen 80;
    server_name *.com;
    location / {
      return 400;
    }
  } 
}
```

当访问 `example.com`　时会返回 300

当访问 `examples.com`　时会返回 400

## 0x03 Conclusion

总结可以得出 Nginx 会按照如下顺序规则使用 HOST 以及`  server_name` 去匹配 server directive block

1. the exact name matches the `server_name`
2. the longest wildcard name starting with an asterisk, e.g. “`*.example.com`”
3. the longest wildcard name ending with an asterisk, e.g. “`mail.*`”
4. the first matching regular expression (in order of appearance in the configuration file)
5. the server directive block wich the server’s `listen` directive has `default_server` 
6. the firtst server directive block in order of appearance in the configuration file

**references**

[^1]:http://nginx.org/en/docs/http/server_names.html
[^2]: http://nginx.org/en/docs/http/request_processing.html

