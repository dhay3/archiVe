# Linux exec 

> bash -c 'help exec'
>
> exec只能在行首才能被识别，如果被插入使用，不会被识别
>
> ```
> [root@8d3d229c-4aab-4812-96b9-37c8bc47a1d8 opt]# printf 'top' | xargs exec 
> xargs: exec: No such file or directory
> 
> ```

exec 内建命令不启动新的shell，而是使用命令产生的新命令替换，执行exec命令的进程。

执行`exec top`，当退出top进程时，执行该命令的shell 退出。

一般将`exec`命令放到一个shell脚本中，当执行到exec 命令就会退出创建shell脚本的进程

```
root in /opt λcat test.sh
       File: test.sh
   1   # !/bin/bash
   2   
   3   echo 123
   4   exec ls
   5   top                                                                /0.1s
root in /opt λsh test.sh 
123
 bak.xml	  ibak.xml			    lsd_0.18.0_amd64.deb
'burpsuite pro'   jdk-14.0.2			    test
 containerd	  jdk-14.0.2_linux-x64_bin.tar.gz   test.sh               /0.0s
root in /opt λ 


```

==还有另外一种如果没有指定命令，就是重定向==

```
root in /opt λ exec < /dev/random
```

默认重定向到stdout(忽略了1)，shell会当成命令来执行，我们可以重定向到0，1，2以外的流，就会输出内容

```
root in /opt λ exec 3 < /dev/random
```

