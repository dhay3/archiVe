# Nginx Master and Work Processes

Nginx 启动时会创建 2 种进程

1. Master process

   The main purpose of the master process is to read and evaluate configuration files, as well as maintain the worker processes.

2. Work processes

   The worker processes do the actual processing of requests.

   由 `work_processes` directive 指定

master processes 的所有者和 `user`  directives 无关，`user` directives 只影响 work processes

```
$ ps aux  | grep nginx | grep -v grep
root     24769  0.0  0.0  46152  1932 ?        Ss   03:49   0:00 nginx: master process ./sbin/nginx
nobody   24853  0.0  0.0  46348  1684 ?        S<   05:29   0:00 nginx: worker process
nobody   24854  0.0  0.0  46348  1684 ?        S<   05:29   0:00 nginx: worker process
nobody   24855  0.0  0.0  46348  1684 ?        S<   05:29   0:00 nginx: worker process
nobody   24856  0.0  0.0  46348  1684 ?        S<   05:29   0:00 nginx: worker process
```

**references**

[^1]:https://docs.nginx.com/nginx/admin-guide/basic-functionality/runtime-control/