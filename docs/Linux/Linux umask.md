# Linux umask

参考：

https://blog.csdn.net/guaiguaihenguai/article/details/79934142

https://blog.csdn.net/stpeace/article/details/45509425

该函数为进程(对子进程同样生效)设置文件模式屏蔽字，并返回之前的值

一共有九种权限，用16位表示

![img](https://img-blog.csdn.net/20180413202810248)

https://unix.stackexchange.com/questions/337182/different-umask-for-directories-and-files

可以通过`umask [new_mask]`指定mask，四位对于目录默认从0777中减去，对于文件默认从0666中减去权限。

```
root in /usr/local/\ λumask 
022
root in /usr/local/\ λtouch a.txt
root in /usr/local/\ λ mkdir a
root in /usr/local/\ λ ll
drwxr-xr-x root root   4 KB Sun Feb  7 18:25:23 2021  a
.rw-r--r-- root root   0 B  Sun Feb  7 18:25:19 2021  a.txt

root in /usr/local/\ λ umask 077
root in /usr/local/\ λ mkdir b
root in /usr/local/\ λ touch b.txt
root in /usr/local/\ λ ll
drwxr-xr-x root root   4 KB Sun Feb  7 18:25:23 2021  a
.rw-r--r-- root root   0 B  Sun Feb  7 18:25:19 2021  a.txt
drwx------ root root   4 KB Sun Feb  7 18:26:21 2021  b
.rw------- root root   0 B  Sun Feb  7 18:26:25 2021  b.txt
```

