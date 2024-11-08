# ngx_http_upstream_module

ngx_http_upstream_module 用于指定一组上游服务器

### upstream

```
Syntax: 	upstream name { ... }
Default: 	—
Context: 	http
```

upstream context，用于定义一组 servers。servers 可以监听不同的端口，或者使用不同的协议。例如

```
upstream backend {
	server backend1.example.com weight=5;
	server 127.0.0.1:8080       max_fails=3 fail_timeout=30s;
	server unix:/tmp/backend3;
	server backup1.example.com  backup;
}
```

默认使用 weighted round-robin 负载均衡

例如上述例子，每 7 个请求中的 5 个请求会使用第 1 个 `server` directive ，剩下的 2 个请求会使用分别会使用第 2 个和第 3 个 `server` directive。如果 前 3 个 server 均没有响应，就会使用第 4 个 `server` directive

### server

```
Syntax: 	server address [parameters...];
Default: 	—
Context: 	upstream
```

指定 `upstream` directive 中使用的 server

address 可以是 IP 或者是 domain-name(如果域名对应有多条解析记录等价于同时声明使用多个 server)，如果没有指定端口默认使用 80

可用 parameters 如下

- `weight=number`

  weighted round-robin 中的权重，默认 1

- `max_conns=number`

  server 同一时间内 active 连接数，默认 0 表示没有限制，只在企业版中支持

- `max_fails=number`

  如果 server 连接失败尝试的次数(server 是否失败由 `fail_timeout` 判断)，默认为 1，如果置为 0 表示没有限制

- `fail_timeout=time`

  server 判定连接失败的时间，默认为 10 秒

- `backup`

  备用 server，只有 primary servers (不包含 `backup` 的 `server`)都不可用时使用，不能和 `hash` , `ip_hash`, `random`  direcitve 一起使用

-  `down`

  标识服务器永远不可用，只在企业版中支持

- `slow_start=time`

  当服务器状态实际从不可用变为可用时，被 Nginx 标识为可用的时间，默认 0s

### state

```
Syntax: 	state file;
Default: 	—
Context: 	upstream
This directive appeared in version 1.9.7. 
```

### keepalive

```
Syntax: 	keepalive connections;
Default: 	—
Context: 	upstream
This directive appeared in version 1.1.4. 
```

### queue(commercial only)

```
Syntax: 	queue number [timeout=time];
Default: 	—
Context: 	upstream

This directive appeared in version 1.5.12. 
```

指定 `upstream` directive 中的 server 不能处理请求时，请求被放入队列的大小(number)以及队列满了之后超过指定时间(timeout=time 默认 60 秒)返回 502

如果没有使用默认的 round-robin lb，需要在 `server` directive 之后声明

### resolver(commercial only)

```
Syntax: 	resolver address ... [valid=time] [ipv4=on|off] [ipv6=on|off] [status_zone=zone];
Default: 	—
Context: 	upstream

This directive appeared in version 1.17.5. 
```

指定用于解析域名的 nameserver

### resolver_timeout(commercial only)

```
Syntax: 	resolver_timeout time;
Default: 	resolver_timeout 30s;
Context: 	upstream

This directive appeared in version 1.17.5. 
```

设置域名解析的超时时间

## LB algorithm relevant

### ip_hash

```
Syntax: 	ip_hash;
Default: 	—
Context: 	upstream
```

表示`upstream` direcitve 使用 ip_hash 算法

The method ensures that requests from the same client will always be passed to the same server except when this server is unavailable. In the latter case client requests will be passed to another server. Most probably, it will always be the same server as well.

### least_conn

```
Syntax: 	least_conn;
Default: 	—
Context: 	upstream

This directive appeared in versions 1.3.1 and 1.2.2. 
```

表示`upstream` direcitve 使用 least_conn 算法

A request is passed to the server with the least number of active connections, taking into account weights of servers. If there are several such servers, they are tried in turn using a weighted round-robin balancing method.

### least_time(commercial only)

```
Syntax: 	least_time header | last_byte [inflight];
Default: 	—
Context: 	upstream

This directive appeared in version 1.7.10. 
```

表示`upstream` direcitve 使用 least_time 算法

a request is passed to the server with the least average response time and least number of active connections, taking into account weights of servers. If there are several such servers, they are tried in turn using a weighted round-robin balancing method.

如果使用了 header，使用 reponse header 响应的时间做为因子

如果使用了 last_byte，使用 full reponse 响应的时间作为因子。如果还带有 inflight，不完整的 HTTP 响应报文的响应时间，也会做因子

只在企业版中支持

### random

```
Syntax: 	random [two [method]];
Default: 	—
Context: 	upstream

This directive appeared in version 1.15.1. 
```

表示`upstream` direcitve 使用 ramdon 算法

Specifies that a group should use a load balancing method where a request is passed to a randomly selected server, taking into account weights of servers.

two 表示 Nginx 会随即选择 2 个 server directive，然后使用指定的 method 选择一个 server，method 默认为 least_conn，也可以指定 least_time(只在企业版中支持)

**references**

[^1]:http://nginx.org/en/docs/http/ngx_http_upstream_module.html