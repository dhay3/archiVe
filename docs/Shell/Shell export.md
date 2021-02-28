# Shell export

==用户创建的变量仅可用于当前 Shell，子 Shell 默认读取不到父 Shell 定义的变量。为了把变量传递给子 Shell，需要使用`export`命令。==这样输出的变量，对于子 Shell 来说就是环境变量。

`export`命令用来向子 Shell 输出变量。

```
NAME=foo
export NAME
```

上面命令输出了变量`NAME`。变量的赋值和输出也可以在一个步骤中完成。

```
export NAME=value
```

> 上面命令执行后，==当前 Shell 及随后新建的子 Shell，都可以读取变量`$NAME`。(父shell不能读取变量)==
>
> 也就是说在Shell脚本中的变量不能被父Shell使用，但是可以在子Shell中使用。==两个不同的Shell脚本变量不会被覆盖==

子 Shell 如果修改继承的变量，不会影响父 Shell。

```
# 输出变量 $foo
$ export foo=bar

# 新建子 Shell
$ bash

# 读取 $foo
$ echo $foo
bar

# 修改继承的变量
$ foo=baz

# 退出子 Shell
$ exit

# 读取 $foo
$ echo $foo
bar
```

上面例子中，子 Shell 修改了继承的变量`$foo`，==对父 Shell 没有影响。==

## 分号

https://stackoverflow.com/questions/27962161/what-does-the-operator-colon-in-the-satement-export-variable-lib-dev-input-eve

导出变量时检查的路径

```
root in /opt/go λ echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

export PATH+$JAVA_HOME/bin:$PATH
```

这里获取默认shell中的`$PATH`值，并组合`$JAVA_HOME`做为新的`$PATH`

PATH是执行二进制文件是的路径





