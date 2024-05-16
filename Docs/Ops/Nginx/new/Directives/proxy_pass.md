# proxy_pass

## 0x00 Overview

```
Syntax: 	proxy_pass URL;
Default: 	—
Context: 	location, if in location, limit_except
```

设置代理，protocol 可以是 http 或者 https，地址可以是 IP address, Unix socket, domain name

例如

```
proxy_pass http://localhost:8000/uri/;
proxy_pass http://10.0.3.75;
proxy_pass http://unix:/tmp/backend.socket:/uri/;
```

如果对应的地址是 domain name 且有多条 A record，会使用 round-robin 代理到实际的地址。也可通过 `upstream` directive 指定一组 `server`

```
upstream servergroup1 {
	server 10.0.3.1;
	server 10.0.3.2;
	server 10.0.3.3;
}
server {
	...
	location / {
		proxy_pass http://servergroup1;
	}
}
```

## 0x01 How proxy_pass process URI

### proxy_pass end with slash



### proxy_pass end without slash

```

```



## 0x03 Embeded Variables





**references**

[^1]:http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass