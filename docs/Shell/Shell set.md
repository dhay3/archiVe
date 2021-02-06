# Shell set

> 所有参数都可以使用`+`来重新开启
>
> 通常使用`set -euo pipefail`，也可以使用`bash -euo pipefail script.sh`

## set -u

执行脚本时，如果遇到不存在的变量，Bash 默认忽略它。

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

## set -o pipefail

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

