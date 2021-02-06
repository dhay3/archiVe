# trap

trap命令用来捕获指定SIG，并执行指定ACTION，我们可以使用`trap -l`或是`kill -l`

来查看所有的SIG

如果`trap`需要触发多条命令，可以封装一个 Bash 函数。

pattern：`trap [ACTION] [SIG1] [SIG2] ...`

### 常用的SIG

- HUP：编号1，脚本与所在的终端脱离联系。
- INT：编号2，用户按下 Ctrl + C，意图让脚本终止运行。
- QUIT：编号3，用户按下 Ctrl + 斜杠，意图退出脚本。
- KILL：编号9，该信号用于杀死进程。
- TERM：编号15，这是`kill`命令发出的默认信号。
- EXIT：编号0，这不是系统信号，而是 Bash 脚本特有的信号，不管什么情况，只要退出脚本就会产生。

`trap`命令响应`EXIT`信号的写法如下。

```
 $ trap 'rm -f "$TMPFILE"' EXIT
```

上面命令中，脚本遇到`EXIT`信号时，就会执行`rm -f "$TMPFILE"`。

### 0x100

> 注意，`trap`命令必须放在脚本的开头。否则，它上方的任何命令导致脚本退出，都不会被它捕获。

```
 #!/bin/bash
 
 trap 'rm -f "$TMPFILE"' EXIT
 
 TMPFILE=$(mktemp) || exit 1
 ls /etc > $TMPFILE
 if grep -qi "kernel" $TMPFILE; then
   echo 'find'
 fi
```

上面代码中，不管是脚本正常执行结束，还是用户按 Ctrl + C 终止，都会产生`EXIT`信号，从而触发删除临时文件。