# Shell coproc

https://www.gnu.org/software/bash/manual/bash.html#Coprocesses

coproc是shell中的关键字，作用与`&`相同，==异步调用进程(在当前shell的子shell中运行)==。

```
coproc [NAME] command [redirections]
```

如果没有提供NAME，默认使用COPROC(如果是单行命令，不能指定NAME)

可以通过jobs命令来查看后台异步启动的进程

```
root in ~ λ jobs
[1]  + running    firefox
```

