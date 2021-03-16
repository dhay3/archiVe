# awk 概述

`awk`是面向行(line-oriented)的语言，和`grep`，`sed`被称为linux中的三剑客

我们可以使用如下命令查看，当前使用的具体awk。`gawk`是有GNU实现的开源awk

```
[root@chz sda]# readlink -f $(which awk)
/usr/bin/gawk
```

syntax：`awk [options] [--] [-f program-file] <file...>`

program-file表示awk脚本，这里的`--`表示options结束，输入固定参数。

## 注释换行

awk和shell采用相同的注释和换行符

```
[root@chz opt]# df -hT | awk '{ \
> #this is a comment
> print $1
> '}
Filesystem
devtmpfs
tmpfs
tmpfs
tmpfs
/dev/mapper/centos-root
/dev/sda1
tmpfs
/dev/sr0
```

## 入门案例

awk默认以换行符为标记，标识每一个行，每次遇到CR，LF就认为当前任务结束，新的一行开始。awk会按照用户指定的分隔符去分隔当前行，如果没有指定分隔符默认使用空格。

```
[root@chz opt]# df -hT  | awk '{print $6}'
Use%
0%
0%
2%
0%
28%
17%
1%
100%
```

awk中最常用的两个函数是`printf` 和 `print`

```
#这里需要使用逗号分隔，如果没有逗号会拼接在一起。和js中console.log类似于go中的fmt.print不同(没有空格)
[root@chz opt]# df -hT  | awk '{print $1,$6}'
Filesystem Use%
devtmpfs 0%
tmpfs 0%
tmpfs 2%
tmpfs 0%
/dev/mapper/centos-root 28%
/dev/sda1 17%
tmpfs 1%
/dev/sr0 100%

#printf的使用方法与C的printf相同
ps auxf | grep -v ps |awk '{printf "%s %s \n",$2,$8}' | grep Z | awk '{print $1}' | xargs kill -SIGCHLD
```

除了输出文本以外，可以拼接自己的字段，==在awk中字符常量需要在双引号中，取值不能带有双引号==

```
[root@chz opt]# df -hT  | awk '{print $1,"cpl"}'
Filesystem cpl
devtmpfs cpl
tmpfs cpl
tmpfs cpl
tmpfs cpl
/dev/mapper/centos-root cpl
/dev/sda1 cpl
tmpfs cpl
/dev/sr0 cpl
```