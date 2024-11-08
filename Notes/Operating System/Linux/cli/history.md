# Linux history

ref:

https://null-byte.wonderhowto.com/how-to/clear-logs-bash-history-hacked-linux-systems-cover-your-tracks-remain-undetected-0244768/

## Digest

> 这里使用的是 GNU history，区别其他版本的 history。具体使用 `man -k history` 查看

syntax：`history option arg`

`history`命令用于显示或管理 shell 的操作历史

例如：当前的 shell 是 GNU Bash，那么就会显示`~/.bash_history`中的内容。如果是 Zsh 那么就会显示`~/.zsh_history`中的内容。一般对应有一个环境变量表示这个文件

```
cpl in ~ λ echo $HISTFILE
/home/cpl/.zsh_history

cpl in ~ λ history | less
1  cat ~/.zsh_history
2  clear
3  ls
4  cd package

cpl in ~ λ head -4 ~/.zsh_history 
: 1626398747:0;cat ~/.zsh_history
: 1626398796:0;clear
: 1626398890:0;ls
: 1626398892:0;cd package
```

在 `history` 中 shell 的操作历史也叫做 event，event 可以是如下的 2 种数据类型

1. a number

   如果是一个 positive number，表示从 1 开始计算的第 number 条命令

   如果是一个 negative number，表示从 currnet event 开始计算的前 number 条命令。例如 `-3` 表示前 3 条命令

2. a string

   selects the most recent event that matches the string. 选择按照 string 匹配的最近的一条

## Optional args

> help history

- `-c`

  clear the history list

  等价于清空对应的文件内容

- `-d offset`

  delete the history entry at postition OFFSET

## Shortcut

- `!!`：执行前一条命令
- `!n`：执行历史文件里面行号为`n`的命令。
- `!-n`：执行当前命令之前`n`条的命令。
- `!string`：执行最近一个以指定字符串`string`开头的命令。
- `!?string`：执行最近一条包含字符串`string`的命令。
- `^string1^string2`：执行最近一条包含`string1`的命令，将其替换成`string2`

## ENV

> 只对 bash 生效

- HISTTIMEFORMAT

  可以让 history 命令显示操作的时间

  ```
  $ export HISTTIMEFORMAT='%F %T  '
  $ history
  1  2013-06-09 10:40:12   cat /etc/issue
  2  2013-06-09 10:40:12   clear
  ```

  format 使用`strftime`

- HISTSIZE

  如果`HISTSIZE=0`写入用户主目录的`~/.bashrc`文件，那么就不会保留该用户的操作历史。如果写入`/etc/profile`，整个系统都不会保留操作历史。

  ```
  export HISTSIZE=0
  ```

## Special

如果是 Zsh，可以在命令前加空格，这样就不会被记录到`~/.zsh_history`中

```
cpl in /tmp λ : > ~/.zsh_history 
cpl in /tmp λ cat ~/.zsh_history 
: 1626398747:0;cat ~/.zsh_history
cpl in /tmp λ  echo hidden msg
hidden msg
cpl in /tmp λ cat ~/.zsh_history
: 1626398747:0;cat ~/.zsh_history
```





