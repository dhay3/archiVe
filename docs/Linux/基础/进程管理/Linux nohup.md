# Linux nohup

参考：https://www.cnblogs.com/jinxiao-pu/p/9131057.html

使用`nohup`会忽略挂起和非强制的POSIX信号(==不会因为父进程的退出而退出，由于Linux的进程机制会被父进程的父进程过继，但是由于父进程打开的sockets会关闭所以被过继的进程的内容产生异常==)，保证进程一直运行。会将输出的内容生成一个`nohup.out`文件。

```
nohup: ignoring input and appending output to 'nohup.out'                                          /0.3s
root in ~ λ cat nohup.out 
       File: nohup.out
   1   [Tue Jan 26 06:17:48.776382 2021] [core:warn] [pid 3719] AH00111: Config variable ${APACHE_RUN_DI
       R} is not defined
   2   apache2: Syntax error on line 80 of /etc/apache2/apache2.conf: DefaultRuntimeDir must be a valid 
       directory, absolute or relative to ServerRoot
   3   /etc/nginx

```

如果启动的进程需要过长时间，我们可以使用异步方式（`&`）启动进程。

```
root in ~ λ nohup firefox &
[2] 4287
nohup: ignoring input and appending output to 'nohup.out'     
```

可以通过jobs来查看通过nohup方式启动的进程

```
root in ~/Desktop λ jobs
[1]  + running    nohup firefox
```