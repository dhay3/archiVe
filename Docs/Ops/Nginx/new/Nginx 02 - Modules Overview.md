# Nginx 02 - Modules Overview

## 0x01 Overview

Nginx 中所有的 Directives 均和 Modules 关联，只有指定的 Modules 被编译进 Nginx 才可以使用对应的 Directives

## 0x02 Modules

Nginx 默认编译的模块可以使用 `./configure --help | grep without` 查看，默认不会包含 ngx_http_ssl_module (如需使用 HTTPS 必须指定该模块)

Nginx 官网已经将所有 Modules 可以使用的 Directives 按照 Module 为维度划分，具体可以查看 Modules reference[^1]

Nginx 会标明每个 Directives 使用的方式 ，以  ` ngx_http_core_module` 的 `absolute_redirect` directive 做为例子

```
Syntax: 	absolute_redirect on | off;
Default: 	absolute_redirect on;
Context: 	http, server, location

This directive appeared in version 1.11.8. 
```

- Syntax 表示使用的语法
- Default 表示如果没有指定该 directive 默认的值
- Context 表示该 directive 可以出现的位置，上例就表示 `absolute_redirect` 可以出现在 `http`, `server`, `location` 内

**referneces**

[^1]:http://nginx.org/en/docs/#Modules%20reference
