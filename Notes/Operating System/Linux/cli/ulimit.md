# Linux ulimit

https://man.linuxde.net/ulimit

ulimit是Shell内建命令。==为Shell进程及其子进程的资源使用设置限制（如果打开两个shell，设置了一个shell，但是另一个shell并不会受影响），会话终止时便结束限制==。而对于长期的固定限制，ulimit 命令语句又可以被添加到由登录 shell 读取的文件中，作用于特定的 shell 用户。

我们使用`man bash /ulimit`来查看，使用`ulimit -a`来查看所有的参数

```
root in /opt/test λ ulimit -a
-t: cpu time (seconds)              unlimited
-f: file size (blocks)              unlimited #进程允许创建的最大文件，单位block
-d: data seg size (kbytes)          unlimited #数据段长度
-s: stack size (kbytes)             8192 #栈大小
-c: core file size (blocks)         0 #对core文件的限制，表示不生成core文件
-m: resident set size (kbytes)      unlimited
-u: processes                       15409  #最大进程数
-n: file descriptors                1024  #单个进程允许打开最大文件数
-l: locked-in-memory size (kbytes)  500741 
-v: address space (kbytes)          unlimited
-x: file locks                      unlimited
-i: pending signals                 15409
-q: bytes in POSIX msg queues       819200
-e: max nice                        0
-r: max rt priority                 0
-N 15:                              unlimited                              /0.0s
```

使用`ulimit [option] `来查看具体的某一项参数的值，使用`ulimit [options] [values]`

来设置改项参数的值，==立即生效==

> 如果用户使用bash，可以在`.bashrc`文件中加入`ulimit -u 64`，来限制用户可以打开的进程数

- ` ulimit -u 10240`

  修改Linux允许最大进程数

- `ulimit -n 4069`

  修改每个进程能打开的最大文件数，缺省值为1024

==可以在`/etc/security/limit.conf`中进行持久配置==

