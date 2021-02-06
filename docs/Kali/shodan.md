# shodan

[TOC]

## 参考:

https://blog.csdn.net/weinierbian/article/details/88823586

搜索过滤

- hostname: 搜索指定的host, 例如 `hostname:"google"`
- port: 搜索指定的端口或服务, 例如 `port:"21"`
- country: 搜索指定的城市, 例如 `city:"Hefei"`
- org: 搜索指定的组织或公司, 例如 `org:"google"`
- isp: 搜索指定的ISP供应商, 例如 `isp:"China Telecom"`
- product: 搜索指定的操作系统软件/平台, 例如 `product:"Apache httpd"`
- version: 搜索指定的软件版本, 例如 `product:"Apache httpd"`
- `before/after`：搜索指定收录时间前后的数据，格式为dd-mm-yy，例如 `before:"11-11-15"`
- ==`net`==：搜索指定的IP地址或子网，例如 `net:"210.45.240.0/24"`

### 搜索实例

查找位于合肥的 Apache 服务器：

```
apache city:"Hefei"
```

查找位于国内的 Nginx 服务器：

```
nginx country:"CN"
```

查找 GWS(Google Web Server) 服务器：

```
"Server: gws" hostname:"google"
```

查找指定网段的华为设备：

```
huawei net:"61.191.146.0/24"
```
