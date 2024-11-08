# Shell 快捷键

参考：

https://wangdoc.com/bash/readline.html

## 概述

Bash 内置了 Readline 库，具有这个库提供的很多“行操作”功能(快捷键)，比如命令的自动补全(TAB)，可以大大加快操作速度。==默认配置参考inputrc==

这个库默认采用Emacs快捷键，也可以改为Vi快捷键

```bash
$ set -o vi
```

下面的命令可以改回 Emacs 快捷键。

```
$ set -o emacs
```

如果想永久性更改编辑模式（Emacs / Vi），可以将命令写在`~/.inputrc`文件，这个文件是 Readline 的配置文件。

```
set editing-mode vi
```

本章介绍的快捷键都属于 Emacs 模式。Vi 模式的快捷键，读者可以参考 Vi 编辑器的教程。

Bash 默认开启这个库，但是允许关闭。

```
$ bash --noediting
```

上面命令中，`--noediting`参数关闭了 Readline 库，启动的 Bash 就不带有行操作功能。

## 自动补全

命令输入到一半的时候，可以按一下 Tab 键，Readline 会自动补全命令或路径。比如，输入`cle`，再按下 Tab 键，Bash 会自动将这个命令补全为`clear`。

如果符合条件的命令或路径有多个，就需要连续按两次 Tab 键，Bash 会提示所有符合条件的命令或路径。

除了命令或路径，==Tab 还可以补全其他值==。如果一个值以`$`开头，则按下 Tab 键会补全变量；如果以`~`开头，则补全用户名；如果以`@`开头，则补全主机名（hostname），主机名以列在`/etc/hosts`文件里面的主机为准。

## 光标移动

- `Ctrl + a`：移到行首。
- `Ctrl + e`：==移到行尾。==
- `Alt + f`：移动到当前单词的词尾。
- `Alt + b`：移动到当前单词的词首。

上面快捷键的 Alt 键，也可以用 ESC 键代替。

## 清除屏幕

`Ctrl + l`快捷键可以清除屏幕，即将当前行移到屏幕的第一行，与`clear`命令作用相同。

## 编辑操作

- `Ctrl + d`：删除光标位置的字符（delete）。

  > 使用`Ctrl + d`的时候，如果当前行没有任何字符，会导致退出当前 Shell，所以要小心。

- `Ctrl + w`：删除光标前面的单词。
- `Alt + l`：将光标位置至词尾转为小写（lowercase）。
- `Alt + u`：将光标位置至词尾转为大写（uppercase）。
- `Ctrl + k`：剪切光标位置到行尾的文本。
- `Ctrl + u`：剪切光标位置到行首的文本。
- `Ctrl + y`：在光标位置粘贴文本。

同样地，Alt 键可以用 Esc 键代替。

## 操作历史

Bash 会保留用户的操作历史，即用户输入的每一条命令都会记录。退出当前 Shell 的时候，Bash 会将用户在当前 Shell 的操作历史写入`~/.bash_history`文件，该文件默认储存500个操作。

环境变量`HISTFILE`总是指向这个文件。

```
$ echo $HISTFILE
/home/me/.bash_history
```

有了操作历史以后，就可以使用方向键的`↑`和`↓`，快速浏览上一条和下一条命令。

下面的方法可以快速执行以前执行过的命令。

```
$ echo Hello World
Hello World

$ echo Goodbye
Goodbye

$ !e
echo Goodbye
Goodbye
```

上面例子中，`!e`表示找出操作历史之中，==最近的那一条==以`e`开头的命令并执行。Bash 会先输出那一条命令`echo Goodbye`，然后直接执行

同理，`!echo`也会执行最近一条以`echo`开头的命令。

```
$ !echo
echo Goodbye
Goodbye

$ !echo H
echo Goodbye H
Goodbye H

$ !echo H G
echo Goodbye H G
Goodbye H G
```

注意，`!string`语法只会匹配命令，不会匹配参数。所以`!echo H`不会执行`echo Hello World`，而是会执行`echo Goodbye`，并把参数`H`附加在这条命令之后。同理，`!echo H G`也是等同于`echo Goodbye`命令之后附加`H G`。

最后，按下`Ctrl + r`会显示操作历史，可以用方向键上下移动，选择其中要执行的命令。也可以键入命令的首字母，Shell 就会自动在历史文件中，查询并显示匹配的结果。





