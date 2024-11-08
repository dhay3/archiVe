# Shell eval

参考：

https://blog.csdn.net/baidu_37964071/article/details/80930704

eval 会对后面的命令进行两遍扫描

如果第一遍扫描后，命令是个普通命令，则执行命令；==如果命令中含有变量的间接引用==，则保证间接引用的语义。也就是说，eval命令将会首先扫描命令行进行所有的置换，然后再执行该命令。

```
root in /etc λ a="cat networks"                                                  root in /etc λ $($a)
zsh: command not found: cat networks                                           
root in /etc λ eval $a
  File: networks
  default     0.0.0.0
  loopback    127.0.0.0
  link-local  169.254.0.0
  
chz@cyberpelican:/etc$ echo $($a)
default 0.0.0.0 loopback 127.0.0.0 link-local 169.254.0.0

```

这么看和`$()`没什么区别，那我们看几个特殊的案例

```
root in /usr/local/\/shell_test λ cat eval.sh 
  File: eval.sh
  #!/usr/bin/env bash
  
#这里必须加转义符，否则就会被解析为$$即打印pid
  echo "\$$#"
  
  eval echo "\$$#"
root in /usr/local/\/shell_test λ zsh eval.sh da lei wa
$3
wa

```

这里eval第一遍先解析`$#`获取到值为3，然后再解析`$3`

## eval结合$()

我们有看到过`eval $(ssh-agent)`设置环境变量，这里先获取ssh-agent返回的结果，第二遍扫描时导出变量到shell

```
#这里ssh-agent是二进制文件，不是通过shell运行，只是输出字符串
root in /usr/local/\/shell_test λ ssh-agent
SSH_AUTH_SOCK=/tmp/ssh-9p5dKsS8gair/agent.7221; export SSH_AUTH_SOCK;
SSH_AGENT_PID=7222; export SSH_AGENT_PID;
echo Agent pid 7222;
```

如果直接`$(ssh-agent)`就会将返回值当作命令，而不会赋值。

## 特殊用法

使用eval可以作为map处理。eval第一遍先处理key，然后处理value

```
[root@localhost ~]# cat file
NAME ZHANG
AGE 20
SEX NV
[root@localhost ~]# cat test.sh
#!/bin/bash

while read KEY VALUE
do
#这里如果使用普通的赋值，变量名含有特殊字符就会报错
    eval "${KEY}=${VALUE}"
done < file
echo "$NAME $AGE $SEX"
[root@localhost ~]# ./test.sh
ZHANG 20 NV
```

