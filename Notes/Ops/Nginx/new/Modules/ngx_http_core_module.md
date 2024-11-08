# ngx_http_core_module

Nginx http 最核心的模块，编译时自动选择

## Request relevant

### client_max_body_size

```
Syntax: 	client_max_body_size size;
Default: 	client_max_body_size 1m;
Context: 	http, server, location
```

Sets the maximum allowed size of the client request body. If the size in a request exceeds the configured value, the 413 (Request Entity Too Large) error is returned to the client. Please be aware that browsers cannot correctly display this error. Setting *size* to 0 disables checking of client request body size.

client request body 能承载的最大值，如果为 0 表示无限制。如果超出配置的值会回送 413

### client_body_timeout

```
Syntax: 	client_body_timeout time;
Default: 	client_body_timeout 60s;
Context: 	http, server, location
```

Defines a timeout for reading client request header. If a client does not transmit the entire header within this time, the request is terminated with the 408 (Request Time-out) error.

### client_header_buffer_size

```
Syntax: 	client_header_buffer_size size;
Default: 	client_header_buffer_size 1k;
Context: 	http, server
```

Sets buffer size for reading client request header.

### large_client_header_buffers

```
Syntax: 	large_client_header_buffers number size;
Default: 	large_client_header_buffers 4 8k;
Context: 	http, server
```



## Response relevant

### alias

```
Syntax: 	alias path;
Default: 	—
Context: 	location
```

例如定义如下配置

```
location /i/ {
    alias /data/w3/images/;
}
```

如果请求 `/i/top.gif` 就会变为 `/data/w3/images/top.gif`

```
location ~ ^/users/(.+\.(?:gif|jpe?g|png))$ {
    alias /data/w3/images/$1;
}
```

如果请求`/users/a.jpeg`就会变为`/data/w3/images/a.jpeg`

最好直接使用 `root` 替代 `alias`

### default-type

```
Syntax: 	default_type mime-type;
Default: 	default_type text/plain;
Context: 	http, server, location
```

Defines the default MIME type of a response. Mapping of file name extensions to MIME types can be set with the [types](http://nginx.org/en/docs/http/ngx_http_core_module.html#types) directive.

### error_page

```
Syntax: 	error_page code ... [=[response]] uri;
Default: 	—
Context: 	http, server, location, if in location
```

### etag

```
Syntax: 	etag on | off;
Default: 	etag on;
Context: 	http, server, location
This directive appeared in version 1.3.3. 
```

Enables or disables automatic generation of the “ETag”[^2] response header field for static resources.

开启该参数后

1. Nginx 会计算 static resources 的 ETag values(直接理解成 hash)
2. 当 client 请求某一 static resource 时，Nginx 在回送响应时会将计算的 ETag 放在响应报文头中
3. 然后浏览器会将对应的 static resource 缓存(没有开启浏览器的 disable cache)
4. 当 client 再次请求 static resouce 时，client 会将 Nginx 回送的 ETag 值放在请求头中的 If-None-Match，如果对应的文件没有改变过，Nginx 计算出来的 ETag 会和请求中的 If-None-Match 值相同
5. 这时 Nginx 就不会将对应 static resource 传给 client

该参数可以有效节省服务器的带宽

### internal

## keepalive relevant

### keepalive_request

```
Syntax: 	keepalive_requests number;
Default: 	keepalive_requests 1000;
Context: 	http, server, location
This directive appeared in version 0.8.0. 
```

keepalive 信道允许的最大请求数，超过该数值 keepalive 信道关闭

### keepalive_time

```
Syntax: 	keepalive_time time;
Default: 	keepalive_time 1h;
Context: 	http, server, location
This directive appeared in version 1.19.10. 
```

keepalive 信道允许打开的最长时间，超过该时间 keepalive 信道关闭

### keepalive_timeout

```
Syntax: 	keepalive_timeout timeout [header_timeout];
Default: 	keepalive_timeout 75s;
Context: 	http, server, location
```

The first parameter sets a timeout during which a keep-alive client connection will stay open on the server side. The zero value disables keep-alive client connections. The optional second parameter sets a value in the “Keep-Alive: timeout=`*time*`” response header field. Two parameters may differ.

The “Keep-Alive: timeout=`*time*`” header field is recognized by Mozilla and Konqueror. MSIE closes keep-alive connections by itself in about 60 seconds.

## server relevant

### listen

```
Syntax: 	
listen address[:port] [default_server] [ssl] [http2 | quic] [proxy_protocol] [setfib=number] [fastopen=number] [backlog=number] [rcvbuf=size] [sndbuf=size] [accept_filter=filter] [deferred] [bind] [ipv6only=on|off] [reuseport] [so_keepalive=on|off|[keepidle]:[keepintvl]:[keepcnt]];

listen port [default_server] [ssl] [http2 | quic] [proxy_protocol] [setfib=number] [fastopen=number] [backlog=number] [rcvbuf=size] [sndbuf=size] [accept_filter=filter] [deferred] [bind] [ipv6only=on|off] [reuseport] [so_keepalive=on|off|[keepidle]:[keepintvl]:[keepcnt]];

listen unix:path [default_server] [ssl] [http2 | quic] [proxy_protocol] [backlog=number] [rcvbuf=size] [sndbuf=size] [accept_filter=filter] [deferred] [bind] [so_keepalive=on|off|[keepidle]:[keepintvl]:[keepcnt]];
Default: 	listen *:80 | *:8000;
Context: 	server
```

指定 Nginx 监听具体那个地址的端口，一共 3 种 格式

例如

```
listen 127.0.0.1:8000;
listen 127.0.0.1;
listen 8000;
listen *:8000;
listen localhost:8000;
```

1. 如果只指定了 address 默认使用 80 端口
2. 如果没有指定 listen directive，如果是以 root 运行 nginx 的，默认 80 端口，如果非 root，默认 8000 端口

### location

```
Syntax: 	location [ = | ~ | ~* | ^~ ] uri { ... }
location @name { ... }
Default: 	—
Context: 	server, location
```

具体查看 Direcitves 中的 location.md

**references**

[^1]:http://nginx.org/en/docs/http/ngx_http_core_module.html#http
[^2]:https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag