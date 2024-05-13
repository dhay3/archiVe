# Nginx CLI

## 0x01 Overview

syntax

```
nginx [-?hqTtVv] [-c file] [-e file] [-g directives] [-p prefix] [-s signal]
```

Nginx 通过 `nginx` CLI 来控制和管理 Nginx

直接使用 `nginx` 可以启动 Nginx Server

## 0x02 Options

- `-c file`

  指定 Nginx 使用的配置文件，默认使用 `/etc/nginx/nginx.conf`

- `-t`

  测试 Nginx 的配置文件语法是否正确。也可以查看当前 nginx 使用的配置文件路径

- `-s [stop | quit | reopen | reload ]`

  用于向 Nginx 发送指定的指令

  1. stop

     shutdown Nginx quickly

  2. quit

     shutdown Nginx gracefully

  3. reload

     reload Nginx configuration

  4. reopen

     reopen Nginx log files

- `-V`

  输出 Nginx 详细的版本信息，以及编译的模块

**references**

[^1]:http://nginx.org/en/docs/switches.html