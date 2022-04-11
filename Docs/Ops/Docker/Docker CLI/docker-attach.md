# docker-attach

> 如果容器以`-it`参数运行，可以通过`ctrl+p+q`脱离容器

## 概述

pattern：`docker attach [options] <container>`

将本地的stdout，stderr，stdin附加到==运行中==的容器。如果使用`ctrl+c`默认向容器发送SIGINT。

**注意！！！**

> 在linux中PID为1的进程，会忽略发送SIGINT(ctrl + c)。所以由docker创建的shell可能无法正常关闭

如果想要将detach的快捷键替换使用`--detach-keys=""`

```
root in ~ λ docker attach --detach-keys="@" t3
[root@06dcdb3c3c9c etc]# read escape sequence        
```

