# nginx_core_module



## worker process relavant

### main

#### worker_processes

```
Syntax: 	worker_processes number | auto;
Default: 	worker_processes 1;
Context: 	main
```

worker processes 的数量，通常使用 `auto` 即可，让系统自己选择

#### worker_priority

```
Syntax: 	worker_priority number;
Default: 	worker_priority 0;
Context: 	main
```

woker process 调度 CPU 的优先级，范围为 `-20 - 20`， 值越小优先级越高

可以通过 `top`  的 NI 列来校验是否生效

#### worker_cpu_affinity

```
Syntax: 	worker_cpu_affinity cpumask ...;
worker_cpu_affinity auto [cpumask];
Default: 	—
Context: 	main
```

将 worker process 绑定到指定的 CPU

例如

```
worker_processes    4;
worker_cpu_affinity 0001 0010 0100 1000
```

4 个 worker processes 会被分别绑定到 CPU0,CPU1,CPU3,CPU4

```
worker_processes    2;
worker_cpu_affinity 0101 1010;
```

2 个 worker processes 会被分别绑定到 CPU0/CPU2, CPU1/CPU4

对应 bit 位的值为 1 就表示，绑定对应序号的 CPU

或者使用 `auto` 让系统自己决定该怎么绑定

#### worker_rlimit_core

```
Syntax: 	worker_rlimit_core size;
Default: 	—
Context: 	main
```

#### worker_rlimit_nofile

```
Syntax: 	worker_rlimit_nofile number;
Default: 	—
Context: 	main
```

#### working_direcotry

```
Syntax: 	working_directory directory;
Default: 	—
Context: 	main
```

指定 work processes 的工作目录，主要被用在写 core-file 的场景中

### event

#### woker_connections

```
Syntax: 	worker_connections number;
Default: 	worker_connections 512;
Context: 	events
```

单个 worker process 在同一时间可以打开的连接 (包括 proxied servers)

其值不能超过 maximum number of open files (可以通过 worker_rlimit_nofile 修改)

Nginx 可以自己单独设置，而不使用系统的 `limits.conf`

## Others

#### user

```
Syntax: 	user user [group];
Default: 	user nobody nobody;
Context: 	main
```

指定启动 work processes 的用户以及组（master process 不受影响，如果启动 nginx 的用户为 root，那么 master process 的所有者就是 root）

#### pid

```
Syntax: 	pid file;
Default: 	pid logs/nginx.pid;
Context: 	main
```

指定存储 Nginx pid 的文件路径，如果编译的时候使用了 `--pid-path=<path>`，默认就会使用 path 对应的值 

#### lock_file

```
Syntax: 	lock_file file;
Default: 	lock_file logs/nginx.lock;
Context: 	main
```

指定存储 Nginx lock 的文件路径，如果编译的时候使用 `--lock-path=<path>`, 默认就会使用 path 对应的值

#### pcre_jit

```
Syntax: 	pcre_jit on | off;
Default: 	pcre_jit off;
Context: 	main
This directive appeared in version 1.1.12. 
```

启用或者关闭 pcre_jit (just in time compilation)

PCRE_JIT 可以加快 regex 的处理

#### thread_pool

```
Syntax: 	thread_pool name threads=number [max_queue=number];
Default: 	thread_pool default threads=32 max_queue=65536;
Context: 	main
This directive appeared in version 1.7.11. 
```

#### error_log

```
Syntax: 	error_log file [level];
Default: 	error_log logs/error.log error;
Context: 	main, http, mail, stream, server, location
```

#### events

```
Syntax: 	events { ... }
Default: 	—
Context: 	main
```

提供 event context

#### accept_mutex

```
Syntax: 	accept_mutex on | off;
Default: 	accept_mutex off;
Context: 	events
```

accept_mutex 控制 worker processes 该如何处理 incoming connections

- 如果为 on, 同一时间只有 one worker process 可以收到 incoming connections，然后轮询。在连接并发量少的情况下，可以减少 CPU 的资源，在连接并发量多的情况下，会增加时延
- 如果为 off, 同一时间所有的 woker processes 都可以收到 incoming connections

#### env

```
Syntax: 	env variable[=value];
Default: 	env TZ;
Context: 	main
```

#### include

```
Syntax: 	include file | mask;
Default: 	—
Context: 	any
```

将指定的配置或者是目录中的配置引入

例如

```
include mime.types;
include vhosts/*.conf;
```

#### load_module

```
Syntax: 	load_module file;
Default: 	—
Context: 	main

This directive appeared in version 1.9.11. 
```

引入 dynamic module

例如

```
load_module modules/ngx_mail_module.so;
```

#### daemon

```
Syntax: 	daemon on | off;
Default: 	daemon on;
Context: 	main
```

Nginx 是否以 daemon 的方式运行

**references**

[^1]:http://nginx.org/en/docs/ngx_core_module.html