# Shell set

https://www.gnu.org/software/bash/manual/bash.html#The-Set-Builtin

set用于设定shell该怎么执。

> 所有参数都可以使用`+`来关闭功能选项
>
> 通常使用`set -euo pipefail`，也可以使用`bash -euo pipefail script.sh`
>
> 如果想要查看具体的内建命令如何使用
>
> 1. 非bash使用`bash -c "help command"`
> 2. bash使用`help command`

## set -a

子shell默认获取不到父shell的变量，执行`set -a`之后默认将之后创建的变量导出给子shell

```
[root@cyberpelican \]# sh sc.sh 

[root@cyberpelican \]# set -a
[root@cyberpelican \]# a=20
[root@cyberpelican \]# sh sc.sh 
20
[root@cyberpelican \]# cat sc.sh 
  File: sc.sh
  echo $a
```

## set -u

执行脚本时，如果遇到不存在的变量（除`$@`和`$*`），Bash 默认忽略它。

`set -u`就用来改变这种行为。脚本在头部加上它，遇到不存在的变量就会报错，并停止执行。等价于`set -o nounset`

```
[root@cyberpelican ~]# echo $a
bash: a: unbound variable
```

## set -x

在输出结果之前，先输出执行的那一行命令。

```
[root@cyberpelican opt]# cat test.sh 
set -x
echo foo
echo bar
[root@cyberpelican opt]# ./test.sh 
++ echo foo
foo
++ echo bar
bar
```

## set -e

脚本只要错误就不会继续执行

```
[root@cyberpelican opt]# cat test.sh 
set -e
foo
echo bar

[root@cyberpelican opt]# ./test.sh 
./test.sh: line 2: foo: command not found
```

## set -f

关闭文件的模式扩展

## set -n

读取命令，但是不执行他们。对交互式shell不生效，一般在脚本中，用于验证语法

```
[root@cyberpelican \]# a=10
[root@cyberpelican \]# cat sc.sh 
  File: sc.sh
  set -n
  echo $a
[root@cyberpelican \]# sh sc.sh 
[root@cyberpelican \]# 
```

## set -t

执行完一个命令就退出，如果包括执行完`set -t`

## set -v

运行命令的同时会将输入的命令打印

```
[root@cyberpelican ~]# set -v
__vte_prompt_command
[root@cyberpelican ~]# echo 123
echo 123
123
__vte_prompt_command
```

## set -C

禁止使用重定向将文件覆盖

```
[root@cyberpelican ~]# set -C
[root@cyberpelican ~]# cd /opt/
[root@cyberpelican opt]# echo 123 > a
bash: a: cannot overwrite existing file
```

## set -H

可以使用`!`来执行命令

```
[root@cyberpelican 1557]# set +H
[root@cyberpelican 1557]# !519
bash: !519: command not found
```

## set -o 

通过`set -o option`可以设置一些选项

### pipefail

`set -e`有一个例外情况，就是不适用于管道命令。

只要最后一个子命令不失败，管道命令总是会执行成功，因此它后面命令依然会执行，`set -e`就失效了。

```
[root@cyberpelican opt]# cat test.sh 
set -e
foo|echo a
echo bar

[root@cyberpelican opt]# ./test.sh 
./test.sh: line 2: foo: command not found
a
bar

---

[root@cyberpelican opt]# cat test.sh 
set -eo pipefail
foo|echo a
echo bar

[root@cyberpelican opt]# ./test.sh 
./test.sh: line 2: foo: command not found
a
```

### errorexit

等价`set -e`，命令出错后退出

### noclobber

等价`set -C`，禁止文件重定向后覆盖

### verbose

等价`set -v`