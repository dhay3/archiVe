# Shell 变量

> 1. 获取变量最好添加双引号
>
> 2. 如果赋值的命令，不会在stoud中输出
>
>    ```
>    root in /usr/local/\/shell_test λ a=$(date)                               /0.0s
>    root in /usr/local/\/shell_test λ echo $a
>    Fri 05 Feb 2021 07:16:01 PM CST   
>    ```
>
> 2. 由于Shell中所有的变量都是字符串，所以变量之直接可以做字符串拼接
>
> ```
> [root@cyberpelican bin]# ok=111
> [root@cyberpelican bin]# echo $ok
> 111
> [root@cyberpelican bin]# ok=$ok:222
> [root@cyberpelican bin]# echo $ok
> 111:222
> ```

## 简介

Bash 变量分成环境变量和自定义变量两类。

### 环境变量

==环境变量是 Bash 环境自带的变量，进入 Shell 时已经定义好了，可以直接使用。它们通常是系统定义好的，也可以由用户从父 Shell 传入子 Shell。==

`env`命令或`printenv`命令，可以显示所有环境变量。

```
$ env
# 或者
$ printenv
```

下面是一些常见的环境变量。

- `BASHPID`：Bash 进程的进程 ID。
- `BASHOPTS`：当前 Shell 的参数，可以用`shopt`命令修改。
- `DISPLAY`：图形环境的显示器名字，通常是`:0`，表示 X Server 的第一个显示器。
- `EDITOR`：默认的文本编辑器。
- `HOME`：用户的主目录。
- `HOST`：当前主机的名称。
- `IFS`：词与词之间的分隔符，默认为空格。
- `LANG`：字符集以及语言编码，比如`zh_CN.UTF-8`。
- `PATH`：由冒号分开的目录列表，当输入可执行程序名后，会搜索这个目录列表。
- `PS1`：Shell 提示符。
- `PS2`： 输入多行命令时，次要的 Shell 提示符。
- `PWD`：当前工作目录。
- `RANDOM`：返回一个0到32767之间的随机数。
- `SHELL`：Shell 的名字。
- `SHELLOPTS`：启动当前 Shell 的`set`命令的参数，参见《set 命令》一章。
- `TERM`：终端类型名，即终端仿真器所用的协议。
- `UID`：当前用户的 ID 编号。
- `USER`：当前用户的用户名。

很多环境变量很少发生变化，而且是只读的，可以视为常量。由于它们的变量名全部都是大写，所以传统上，如果用户要自己定义一个常量，也会使用全部大写的变量名。

注意，Bash 变量名区分大小写，`HOME`和`home`是两个不同的变量。

查看单个环境变量的值，可以使用`printenv`命令或`echo`命令。

```
$ printenv PATH
# 或者
$ echo $PATH
```

注意，`printenv`命令后面的变量名，不用加前缀`$`。

### 自定义变量

自定义变量是用户在当前 Shell 里面自己定义的变量，必须先定义后使用，而且仅在当前 Shell 可用。一旦退出当前 Shell，该变量就不存在了。

`set`命令可以显示所有变量（包括环境变量和自定义变量），以及所有的 Bash 函数。

```
$ set
```

## 创建变量

> `$a=$b`不表示赋值，因为会被渲染为命令，所以不存在

用户创建变量的时候，变量名必须遵守下面的规则。

- 字母、数字和下划线字符组成。
- 第一个字符必须是一个字母或一个下划线，不能是数字。
- ==不允许出现空格和标点符号。==

变量声明的语法如下。

```
variable=value
```

上面命令中，等号左边是变量名，右边是变量。注意，等号两边不能有空格。

如果变量的值包含空格，则必须将值放在引号中。

```
myvar="hello world"
```

==Bash 没有数据类型的概念，所有的变量值都是字符串。==

下面是一些自定义变量的例子。

```
a=z                     # 变量 a 赋值为字符串 z
b="a string"            # 变量值包含空格，就必须放在引号里面
c="a string and $b"     # 变量值可以引用其他变量的值
d="\t\ta string\n"      # 变量值可以使用转义字符
e=$(ls -l foo.txt)      # 变量值可以是命令的执行结果
f=$((5 * 7))            # 变量值可以是数学运算的结果
```

变量可以重复赋值，后面的赋值会覆盖前面的赋值。

```
$ foo=1
$ foo=2
$ echo $foo
2
```

上面例子中，变量`foo`的第二次赋值会覆盖第一次赋值。

如果同一行定义多个变量，必须使用分号（`;`）分隔。

```
$ foo=1;bar=2
```

上面例子中，同一行定义了`foo`和`bar`两个变量。

## 读取变量

读取变量的时候，直接在变量名前加上`$`就可以了。

```
$ foo=bar
$ echo $foo
bar
```

每当 Shell 看到以`$`开头的单词时，就会尝试读取这个变量名对应的值。

如果变量不存在，Bash 不会报错，而会输出空字符。

由于`$`在 Bash 中有特殊含义，把它当作美元符号使用时，一定要非常小心，

```
$ echo The total is $100.00
The total is 00.00
```

上面命令的原意是输入`$100`，但是 Bash 将`$1`解释成了变量，该变量为空，因此输入就变成了`00.00`。所以，如果要使用`$`的原义，需要在`$`前面放上反斜杠，进行转义。

```
$ echo The total is \$100.00
The total is $100.00
```

读取变量的时候，变量名也可以使用花括号`{}`包围，比如`$a`也可以写成`${a}`。这种写法可以用于变量名与其他字符连用的情况。

```
$ a=foo
$ echo $a_file

$ echo ${a}_file
foo_file
```

上面代码中，变量名`a_file`不会有任何输出，因为 Bash 将其整个解释为变量，而这个变量是不存在的。只有用花括号区分`$a`，Bash 才能正确解读。

事实上，读取变量的语法`$foo`，可以看作是`${foo}`的简写形式。

如果变量的值本身也是变量，可以使用`${!varname}`的语法，读取最终的值。

```
$ myvar=USER
$ echo ${!myvar}
ruanyf
```

上面的例子中，变量`myvar`的值是`USER`，`${!myvar}`的写法将其展开成最终的值。

## 删除变量

`unset`命令用来删除一个变量。

```
unset NAME
```

这个命令不是很有用。因为不存在的 Bash 变量一律等于空字符串，所以即使`unset`命令删除了变量，还是可以读取这个变量，值为空字符串。

所以，删除一个变量，也可以将这个变量设成空字符串。

```
$ foo=''
$ foo=
```

上面两种写法，都是删除了变量`foo`。由于不存在的值默认为空字符串，所以后一种写法可以在等号右边不写任何值。

## 特殊变量

Bash 提供一些特殊变量。这些变量的值由 Shell 提供，用户不能进

1. `$?`

   > 可以使用短路与和短路或替代

   `$?`为上一个命令的退出码，用来判断上一个命令是否执行成功。==返回值是`0`，表示上一个命令执行成功；如果是非零，上一个命令执行失败。==

   ```
   $ ls doesnotexist
   ls: doesnotexist: No such file or directory
   
   $ echo $?
   1
   ```

   上面例子中，`ls`命令查看一个不存在的文件，导致报错。`$?`为1，表示上一个命令执行失败。

2. `$$`

   `$$`为当前 Shell 的进程 ID。

   ```
   $ echo $$
   10662
   ```

   这个特殊变量可以用来命名临时文件。

   ```
   LOGFILE=/tmp/output_log.$$
   ```


3. `$_`

   `$_`为上一个命令的最后一个参数。

   ```
   $ grep dictionary /usr/share/dict/words
   dictionary
   
   $ echo $_
   /usr/share/dict/words
   ```


4. `$!`

   `$!`为最近一个后台执行的异步命令的进程 ID。

   ```
   $ firefox &
   [1] 11064
   
   $ echo $!
   11064
   ```

   上面例子中，`firefox`是后台运行的命令，`$!`返回该命令的进程 ID。

5. `$0`

   `$0`为当前 Shell 的名称（在命令行直接执行时）或者脚本名（在脚本中执行时）。

   ```shell
   $ echo $0
   bash
   ```

   上面例子中，`$0`返回当前运行的是 Bash。

6. `$0~$9`

   对应脚本的第一个参数到第九个，如果参数多于九个，使用`${10}`的形式引用，以此类推。

   ```
   command -o foo bar
   ```

   那么`-o`是`$1`，`foo`是`$2`，`bar`是`$3`。

7. `$#`

   获取输入参数的所有个数

   ```
   [root@cyberpelican opt]# /bin/bash test.sh a b c
   0
   1
   2
   [root@cyberpelican opt]# cat test.sh 
          File: test.sh
      1   # /usr/bin/env bash
      2   
      3   for ((i=0;i<$#;i++))
      4   do
      5       echo $i
      6   done
   ```

8. `$-`

   `$-`为当前 Shell 的启动参数。

   ```
   $ echo $-
   himBHs
   ```


9. `$@`

   全部的参数，参数之间使用空格分隔

   ```
   [root@cyberpelican opt]# cat test.sh 
   #!/bin/bash
   for i in $@
   do
   echo "$i"
   done
   [root@cyberpelican opt]# ./test.sh 1 2 3 4
   1
   2
   3
   4
   ```


10. `$*`

    全部的参数，参数之间使用变量`$IFS`值的第一个字符分隔，默认为空格，但是可以自定义。

    ```
    [root@cyberpelican opt]# echo $IFS
    
    [root@cyberpelican opt]# cat test.sh 1 2 4
    #!/bin/bash
    for i in $*
    do
    echo "$i"
    done
    
    ```


## 变量的默认值/elvis

Bash 提供四个特殊语法，跟变量的默认值有关，目的是保证变量不为空。

> 特殊写法`${debian_chroot:-}`表示如果没有设置，就置为null

```
${varname:-word}
```

上面语法的含义是，如果变量`varname`存在且不为空，则返回它的值，否则返回`word`。它的目的是返回一个默认值，比如`${count:-0}`表示变量`count`不存在时返回`0`。

```
${varname:=word}
```

上面语法的含义是，如果变量`varname`存在且不为空，则返回它的值，否则将它设为`word`，并且返回`word`。它的目的是设置变量的默认值，比如`${count:=0}`表示变量`count`不存在时返回`0`，且将`count`设为`0`。

```
${varname:+word}
```

上面语法的含义是，如果变量名存在且不为空，则返回`word`，否则返回空值。它的目的是测试变量是否存在，比如`${count:+1}`表示变量`count`存在时返回`1`（表示`true`），否则返回空值。

```
${varname:?message}
```

上面语法的含义是，如果变量`varname`存在且不为空，则返回它的值，否则打印出`varname: message`，并中断脚本的执行。如果省略了`message`，则输出默认的信息“parameter null or not set.”。它的目的是防止变量未定义，比如`${count:?"undefined!"}`表示变量`count`未定义时就中断执行，抛出错误，返回给定的报错信息`undefined!`。

上面四种语法如果用在脚本中，变量名的部分可以用数字`1`到`9`，表示脚本的参数。

```
filename=${1:?"filename missing."}
```

上面代码出现在脚本中，`1`表示脚本的第一个参数。如果该参数不存在，就退出脚本并报错。

