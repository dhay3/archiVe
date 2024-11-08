# Shell coproc

https://www.gnu.org/software/bash/manual/bash.html#Coprocesses

coproc是shell中的关键字，作用与`&`相同，==异步调用进程(不与当前的shell在同一个进程)==。

```
coproc [NAME] command [redirections]
```

如果没有提供NAME，默认使用COPROC(如果是单行命令，不能指定NAME)