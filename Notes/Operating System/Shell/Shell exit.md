# exit

> The return status (see [Exit Status](https://www.gnu.org/software/bash/manual/bash.html#Exit-Status)) of a simple command is its exit status as provided by the POSIX 1003.1 `waitpid` function, or 128+n if the command was terminated by signal n.

退出时，脚本会返回一个退出值。脚本的退出值，==`0`表示正常，`1`表示发生错误，`2`表示用法不对==，`126`表示不是可执行脚本，`127`表示命令没有发现。如果脚本被信号`N`终止，则退出值为`128 + N`。简单来说，只要退出值非0，就认为执行出错。

```
if [ $(id -u) != "0" ]; then
  echo "根用户才能执行当前脚本"
  exit 1
fi
```

调用`exit(int status)`函数，使进程以指定状态退出，如果只调用`exit`相当于`exit $?`

退出码规定：

- 0表示成功（Zero - Success）

- 非0表示失败（Non-Zero - Failure）

- 2表示用法不当（Incorrect Usage）

- 127表示命令没有找到（Command Not Found）

- 126表示不是可执行的（Not an executable）

- \>=128  减去127表示导致进程退出的信号的值

