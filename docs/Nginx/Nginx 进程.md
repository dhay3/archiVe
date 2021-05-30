# Nginx 进程

启动Nginx后，进程不只有一个，默认情况下至少有两个Nginx进程

```
root in /etc/ssh λ ps -ef | grep nginx
root       14865       1  0 02:37 ?        00:00:00 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
www-data   14866   14865  0 02:37 ?        00:00:00 nginx: worker process
www-data   14867   14865  0 02:37 ?        00:00:00 nginx: worker process
www-data   14868   14865  0 02:37 ?        00:00:00 nginx: worker process
www-data   14869   14865  0 02:37 ?        00:00:00 nginx: worker process
root       14966   13840  0 02:41 pts/2    00:00:00 grep --color nginx    
```

- master进程负责管理worker进程，同时负责读取配置文件。
- worker进程负责处理请求

生成的worker 进程个数由配置文件中的worker_process决定

```
Syntax:	worker_processes number | auto;
Default:	
worker_processes 1;
Context:	main
```

如果值为auto，nginx自动获取cpu核数(可以使用`lscpu`来查看)。

==为了避免cpu在切换进程时产生性能损耗==，我们也可以将worker进程与cpu核心进行绑定。我们需要使用work_cpu_affinity

```
Syntax:	worker_cpu_affinity cpumask ...;
worker_cpu_affinity auto [cpumask];
Default:	—
Context:	main
```

我们通过`lscpu | grep cpus`来获取核数，决定了cpumask的位数，例如

```
worker_processes    4;
worker_cpu_affinity 0001 0010 0100 1000;
```

表示将工作线程分别绑定1号核，2号核，3号核，4号核

