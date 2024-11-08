# Nginx Common Varibles

## 0x01 How to show Nginx vairables’ values?

Chang the nginx.conf as follow

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
            return 200 $args;
        }
    }
}
```

Or use `add_header` directive

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
          add_header body_bytes_sent $body_bytes_sent;
          return 302 'https://baidu.com';
        }
    }
}

```

Use Curl to verify

```
(base) 0x00 in /etc/nginx λ curl 'localhost?a=A&b=B'
a=A&b=B
```

## 0x02 Variables

> The `ngx_http_core_module` module supports embedded variables with names matching the Apache Server variables. First of all, these are variables representing client request header fields, such as `$http_user_agent`, `$http_cookie`, and so on. Also there are other variables:

意味着只要是 request header 中的值都可以使用 `$http_{header}` 的方式来获取。例如

```
(base) 0x00 in /run λ curl -H 'x-debug: test' localhost
test
(base) 0x00 in /run λ cat /etc/nginx/nginx.conf
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    server {
        listen       80;
        server_name  _;

        location / {
            return 200 '$http_x_debug';
        }
    }
}
```

除此外还有如下的 built-in variables

### 0x021 request

- `$args`

  arguments in the request line

  GET request 中的参数

  ```
  (base) 0x00 in /etc/nginx λ curl 'localhost?a=A&b=B'
  a=A&b=B
  ```

- `$arg_{name}`

  argument *name* in the request line

  GET request 中的参数名对应的值

  ```
  (base) 0x00 in /etc/nginx λ curl 'localhost?foo=bar'
  bar                                                                                                                                                                             
  (base) 0x00 in /etc/nginx λ cat /etc/nginx/nginx.conf
  worker_processes  1;
  events {
      worker_connections  1024;
  }
  http {
      server {
          listen       80;
          server_name  _;
          location / {
            return 200 '$arg_foo';
          }
      }
  }
  ```

- `$schema`

  request scheme, “`http`” or “`https`”

  请求的协议

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost'
  http%
  ```

- `$request`

  full original request line

  ```
  (base) 0x00 in /etc/nginx λ curl -s  'localhost?token=aabb'
  GET /?token=aabb HTTP/1.1
  ```

- `$uri`

  current URI in request

  这里特指 Host 后的路径（不包含参数），即 URL path

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost/path?token=aabb'
  /path
  ```

- `$request_uri`

  full original request URI (with arguments)

  Host 后的路径包含参数

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost/path?token'
  /path?token
  ```

- `$request_method`

  request method, usually “`GET`” or “`POST`”

  请求的方法

  ```
  (base) 0x00 in /etc/nginx λ curl -s  -X PUT 'localhost'
  PUT
  ```

- `$request_time`

  request processing time in seconds with a milliseconds resolution (1.3.9, 1.2.6); time elapsed since the first bytes were read from the client

  请求处理完成的时间

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost'
  0.000
  ```

- `$content_length`

  “Content-Length” request header field

  request body 的长度

  ```
  (base) 0x00 in /etc/nginx λ curl -s -d 'data' 'localhost'
  4
  ```

- `$content_type`

  “Content-Type” request header field

  请求中的 content-type

  ```
  (base) 0x00 in /etc/nginx λ curl -s -H "content-type:application/json" localhost
  application/json
  ```

- `$host`

  in this order of precedence: host name from the request line, or host name from the “Host” request header field, or the server name matching a request

  ```
  (base) 0x00 in /etc/nginx λ curl -s -H "host: www.example.org" localhost
  www.example.org
  ```

- `$hostname`

  the host name from the “Host” request header

  可以等价理解成 `$host`

- `$https`

  “`on`” if connection operates in SSL mode, or an empty string otherwise

  是否启用 TLS

### 0x022 response

- `$status`

  response status (1.3.2, 1.2.2)

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost'
  200
  ```

- `$body_bytes_sent`

  number of bytes sent to a client, not counting the response header; this variable is compatible with the “`%B`” parameter of the `mod_log_config` Apache module

  response body 中字节大小

  ```
  (base) 0x00 in /etc/nginx λ curl -d '{k=v}' localhost
  0
  ```

- `$bytes_sent`

  number of bytes sent to a client (1.3.8, 1.2.5)

  reponse 字节大小

  ```
  
  ```

- `$server_protocol`

  request protocol, usually “`HTTP/1.0`”, “`HTTP/1.1`”

  服务器协议

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost/path'
  HTTP/1.1
  ```

### 0x023 miscellaneous

- `$binary_remote_addr`

  client address in a binary form, value’s length is always 4 bytes for IPv4 addresses or 16 bytes for IPv6 addresses

  以二进制的方式显示 client IP

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost' | xxd -b
  00000000: 01111111 00000000 00000000 00000001                    ....
  ```

- `$remote_addr`

  `$remote_port`

  client address port

  ```
  (base) 0x00 in /etc/nginx λ curl -s  localhost
  127.0.0.1 30374
  ```

- `$server_addr`

  `$server_port`

  an address of the server which accepted a request

  port of the server which accepted a request

  ```
  (base) 0x00 in /etc/nginx λ curl -s 'localhost'
  127.0.0.1 80	
  ```

- `$tcpinfo_rtt`, `$tcpinfo_rttvar`, `$tcpinfo_snd_cwnd`, `$tcpinfo_rcv_space`

  information about the client TCP connection; available on systems that support the `TCP_INFO` socket option

  TCP 连接相关的参数。(这里可以发现 Nginx 使用 tcp recive window 实际有别于系统的 tcp_rmem)

  ```
  (base) 0x00 in /etc/nginx λ curl -d '{k=v}' localhost
  tcpinfo_rtt=14, tcpinfo_rttvar=7,tcpinfo_snd_cwnd=10, tcpinfo_rcv_space=26000
  
  (base) 0x00 in /etc/nginx λ sudo sysctl -a | grep tcp_rmem
  net.ipv4.tcp_rmem = 4096        102400  67108864
  ```

  

**referenes**

[^1]:http://nginx.org/en/docs/varindex.html
[^2]:https://serverfault.com/questions/404626/how-to-output-variable-in-nginx-log-for-debugging
[^3]:https://serverfault.com/questions/1033992/whats-the-nginx-equivalent-of-console-log-or-print

[^4]:https://statuslist.app/nginx/variables/
