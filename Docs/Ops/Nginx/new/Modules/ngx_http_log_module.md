# ngx_http_log_module

ngx_http_log_module 用于将请求以指定的格式记录到日志中，编译时自动选择

### log_format

```
Syntax: 	log_format name [escape=default|json|none] string ...;
Default: 	log_format combined "...";
Context: 	http
```

用于指定 `access_log` directive 日志的格式，例如

```
log_format combined '$remote_addr - $remote_user [$time_local] '
                    '"$request" $status $body_bytes_sent '
                    '"$http_referer" "$http_user_agent"';
```

string 部分表示日志的格式，可以包含 variables，以 `$varname` 的形式调用

escape 用于指定日志中某些字符是否需要转译，支持 3 种，默认使用 default escaping。下面例子均使用 combined 格式

1. default

   会将 `"`，`\`  或者 ASCII  值小于 32 大于 126 的 字符转译成，`\xXX` (16进制)

   例如

   ```
   curl -v 127.0.0.1/\\
   127.0.0.1 - - [13/May/2024:07:37:23 +0000] "GET /\x5C HTTP/1.1" 404 153 "-" "curl/7.29.0"
   ```

   如果请求中没有对应 string 中的 variables 值的话，会以 `-` (hypen)表示，例如 `$remote_user`

2. json

   所有在 json 中不合法的字符都会被转译，例如 `"` 转译为 `\"`，`;` 转译为 `\;`

   例如

   ```
   curl -v 127.0.0.1/\\
   127.0.0.1 -  [13/May/2024:07:57:40 +0000] "GET /get\\ HTTP/1.1" 200 0 "" "curl/7.29.0"
   ```

   如果请求中没有对应 string 中的 variables 值的话，直接以空格表示，例如 `$remote_user`

3. none 

   不对任何字符做转译

   例如

   ```
   curl -v 127.0.0.1/\\
   127.0.0.1 -  [13/May/2024:08:07:30 +0000] "GET /get\ HTTP/1.1" 200 0 "" "curl/7.29.0"
   ```

### access_log

```
Syntax: 	access_log path [format [buffer=size] [gzip[=level]] [flush=time] [if=condition]];
					access_log off;
Default: 	access_log logs/access.log combined;
Context: 	http, server, location, if in location, limit_except
```

设置所有请求日志(包含 500/400 httpcode)写入的位置(path 可以和 [syslog](http://nginx.org/en/docs/syslog.html) 一起使用)以及 format，其中的 format 值对应 `log_format` directive 中的 name，如果没有指定默认使用 combined

`access_log off` 表示不记录日志

**references**

[^1]:http://nginx.org/en/docs/http/ngx_http_log_module.html#access_log