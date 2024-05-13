# location

## 0x00 Overview

在使用 location 之前需要知道 Nginx 会对 URI 做那些操作

1. 注意 Nginx 中的 URI 和传统意义上的 URI (传统 URI 的定义为`URL = scheme ":" ["//" authority] path ["?" query] ["#" fragment]` )不同(这点可以从 variable `$uri` 中得出)

   ```
   worker_processes  1;
   events {
       worker_connections  1024;
   }
   http {
       server {
           listen       80;
           server_name  _;
           location / {
             return 200 '$uri';
           }
       }
   }
   ```

   调用接口可以发现在 Nginx URI 的部分为 `URL.split('/')[1].split('?')[0]`

   ```
   $ curl -s 'localhost/path?token=aabb'
   /path
   ```

2. 如果 URL 做了 encoding，例如 URL中有中文, Nginx 会自动 decoding 然后再做 URI 匹配

   例如有如下配置

   ```
   worker_processes auto;
   events {
   
   }
   http {
     server {
     listen 80;
     location /无 {
       return 300;
     }
    }
   }
   ```

   发送请求 `localhost/%E6%97%A0` ，Nginx 就会先将 URL 解码

   ```
   curl -I 'localhost/%E6%97%A0'
   HTTP/1.1 300
   Server: nginx/1.24.0
   Date: Sat, 11 May 2024 09:19:47 GMT
   Content-Type: text/plain
   Content-Length: 0
   Connection: keep-alive
   ```

location direcitve 使用格式如下

```
Syntax: 	location [ = | ~ | ~* | ^~ ] uri { ... }
					location @name { ... }
Default: 	—
Context: 	server, location
```



另外补充说明这里的 Context 指明 location 同样也能出现在 location 指令块中，例如

```
location ^~ /images/ {
    ...
    location ~ /*.jpg/ {
    	...
    }
    location ~ /*.png/ {
    	...
    }
}
```

## 0x01 Modifiers

可以将这些符号统一称为 modifiers，可以分为两类

1. prefix modifiers - `=`, `^~`, none
2. regex modifiers - `~` , `~*`

- `=`

  exact matching, the search terminate immediately

  uri 必须精确匹配才会执行，匹配后不会匹配其他的 modifiers

  属于 prefix modifiers

  ```
  location = /pics/ {
      [ configuration ]
  }
  ```

  如果请求为 `/pics/1.img`  就不匹配

- `^~`

  以 regex 匹配 uri ，如果匹配，就不会再去匹配 regex modifiers 的规则

  属于 prefix modifiers

  ```
  location ^~ /images/ {
      [ configuration ]
  }
  ```

  如果请求 `/image/2.img` 则匹配

- none

  即没有符号

  属于 prefix modifiers

  ```
  location /documents/ {
      [ configuration ]
  }
  ```

- `~*`

  case-insensitive matching

  uri 使用正则匹配，忽略大小写

  属于 regex modifiers

  ```
  location ~* /MAP/ {
      [ configuration ]
  }
  ```

- `~`

  case-sensitive matching

  uri 使用正则匹配，大小写敏感

  属于 regex modifiers

  ```
  location ~ /Map/ {
      [ configuration ]
  }
  ```


### cautions

1. `~` 和 `^~` 

   不能同时使用，例如 `~ /get`  和 ` ^~ /get` 

2. 针对大小写不敏感的系统(MacOS/Windows) `~` 等价与 `~*`

## 0x02 Matching rules

> *To find location matching a given request, nginx first checks locations defined using the prefix strings (prefix locations). Among them, the location with the longest matching prefix is selected and remembered. Then regular expressions are checked, in the order of their appearance in the configuration file. The search of regular expressions terminates on the first match, and the corresponding configuration is used. If no match with a regular expression is found then the configuration of the prefix location remembered earlier is used.*

简单的说就是

1. Nginx 首先会使用 URI 去匹配 prefix modifiers，然后选择 the longest prefix  (这点和路由的逻辑很像)，如果匹配的规则为  `=`  以及  `^~`  就会直接返回，如果不是( 即 none )就会将匹配的规则寄存，
2. 然后使用 URI 按照在配置文件中出现的位置去匹配 regular modifiers，如果匹配到任意一个就会返回，如果没有匹配到任意一个就返回第一步寄存的规则

伪代码逻辑如下

```
def match()
	read rules from top to bottom

def config()
	while rules && match(= uri) then
		apply(= config)
		return
  shift rule
  while rules && match(^~ uri) then
      apply(^= config)
      return
    shift rule
  while rules && (match(~ uri) or match(~* uri)) then
      apply(~ uri) or apply(~*)
      return
    shift rule
  while rules && match(uri)
    if match(uri) then
      apply(uri)
      return
    shift rule
```

1. 先查看是否有 `locatoin = uri` 的，如果有匹配，则使用该配置
2. 如果没有，查看 `location ^~ uri` 的，如果有匹配，则使用该配置
3. 如果没有，查看 `location ~ uri` 和 `location ~* uri` 的，如果有匹配，则使用该配置
4. 如果没有，查看 `location uri` 的,如果有匹配，则使用该配置，如果没有就报错

可以归纳得出 modifiers 优先级如下

1. `=`
2. `^~`
3. `~` `~*`
4. none

例如有如下配置

```
worker_processes auto;
events {

}
http {
  server {
  listen 80;
  location /get {
    return 300;
  }
  location / {
  	return 400;
  }
  location ~ /get {
    return 500;
  }
  location ~* /get {
    return 600;
  }
 }
}
```

当请求 `localhost/get`  时，就会匹配到 `~ /get`

```
$ curl -I localhost/get
HTTP/1.1 500 Internal Server Error
Server: nginx/1.24.0
Date: Sat, 11 May 2024 08:21:35 GMT
Content-Type: text/html
Content-Length: 177
Connection: close
```

## 0x03 @

还有一种特殊的 prefix modifiers，只用于内部转发，外部的请求不会匹配该规则，被称为 named location(`@`)

例如

```
worker_processes auto;
events {
}
http {
  server {
  listen 80;
  location /api {
    rewrite ^/api/(.*)$ @get;
  }

  location @get {
    proxy_pass http://127.0.0.1:5000/get;
  }
 }
}
```

0x04 Prefix

**references**

[^1]:http://nginx.org/en/docs/http/ngx_http_core_module.html#location
[^2]:https://www.digitalocean.com/community/tutorials/understanding-nginx-server-and-location-block-selection-algorithms
[^3]:https://serverfault.com/questions/674425/what-does-location-mean-in-an-nginx-location-block
[^4]:https://serverfault.com/questions/738452/what-does-the-at-sign-mean-in-nginx-location-blocks



