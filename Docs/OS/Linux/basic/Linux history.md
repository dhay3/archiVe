# Linux history 

`history`命令能显示Shell操作历史，即`.bash_history`或是其他shell的history文件的内容。

```
$ history
...
498 echo Goodbye
499 ls ~
500 cd
```

使用该命令，而不是直接读取`.bash_history`文件的好处是，它会在所有的操作前加上行号，最近的操作在最后面，行号最大。

通过定制环境变量`HISTTIMEFORMAT`，可以显示每个操作的时间。

```
$ export HISTTIMEFORMAT='%F %T  '
$ history
1  2013-06-09 10:40:12   cat /etc/issue
2  2013-06-09 10:40:12   clear
```

上面代码中，`%F`相当于`%Y - %m - %d`，`%T`相当于`%H : %M : %S`。

只要设置`HISTTIMEFORMAT`这个环境变量，就会在`.bash_history`文件保存命令的执行时间戳。如果不设置，就不会保存时间戳。

如果不希望保存本次操作的历史，可以设置环境变量`HISTSIZE`等于0。

```
export HISTSIZE=0
```

如果`HISTSIZE=0`写入用户主目录的`~/.bashrc`文件，那么就不会保留该用户的操作历史。如果写入`/etc/profile`，整个系统都不会保留操作历史。

如果想搜索某个以前执行的命令，可以配合`grep`命令搜索操作历史。

```
$ history | grep /usr/bin
```

上面命令返回`.bash_history`文件里面，那些包含`/usr/bin`的命令。

操作历史的每一条记录都有编号。知道了命令的编号以后，可以用`感叹号 + 编号`执行该命令。如果想要执行`.bash_history`里面的第8条命令，可以像下面这样操作。

```
$ !8
```

`history`命令的`-c`参数可以清除操作历史。

```
$ history -c
```

## 快捷键

- `!!`：执行上一个命令。
- `!n`：执行历史文件里面行号为`n`的命令。
- `!-n`：执行当前命令之前`n`条的命令。
- `!string`：执行最近一个以指定字符串`string`开头的命令。
- `!?string`：执行最近一条包含字符串`string`的命令。
- `^string1^string2`：执行最近一条包含`string1`的命令，将其替换成`string2`

## 特殊

==如果在history命令之前加一个空格，那么命令不记录在history文件中==

```
cpl in /tmp λ : > ~/.zsh_history 
cpl in /tmp λ cat ~/.zsh_history 
: 1626398747:0;cat ~/.zsh_history
cpl in /tmp λ  echo hidden msg
hidden msg
cpl in /tmp λ cat ~/.zsh_history
: 1626398747:0;cat ~/.zsh_history
```

