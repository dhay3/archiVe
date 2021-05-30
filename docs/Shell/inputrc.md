# Inputrc

参考：

https://linux.die.net/man/3/readline

`inputrc`文件用于处理键盘的映射，`~/.inputrc`作为用户配置文件，`/etc/inputrc`作为全局配置。

## variables

通过`set variables`可以为readline设置一些特殊的配置，值为on 和 1 表示开启其他值关闭(通常使用off)

- bell-style

  是否开启terminal bell，设置为none关闭

- colored-completion-prefix

  被补全的部分高亮显示

- colored-stats

  补全命令时使用不同颜色区分类别

- completion-ignore-case

  补全命令时忽略大小写

- expand-tilde

  对`~`进行模式扩展

- history-size

  是否对`~/.history`文件中记录的条数进行约束，默认没有上限。

- mark-directories

  对补全的目录会在其后面添加slash

- mark-symlinked-directories

  对连接到目录的链接补全时，在末尾添加slash

- match-hidden-files

  对文件名补全时，也会显示隐藏文件

## keybinding

一组keybinding以`keyseq:function-name`来定义

keyseq表示键盘的输入，function-name表示输入后调用的函数或功能

### keyseq

keyseq通常由如下几个键组成

- `\C-`表示ctrl
- `\M-`表示meta键，==通常是right alt==
- `-e`表示转义符
- `\\`表示backslash
- `\"`和`\'`分别表示双单引号

### function-name

具体查看readline(3)的commands部分。function-name后面表示默认调用的输入

1. beginning-of-line(C-a)
2. end-of-line(C-e)
3. backward-char(C-b)
4. forward-char(C-f)
5. clear-screen(C-l)
6. reverse-search-history(C-r)搜索history中指定的关键词
7. kill-line(C-k)
8. unix-line-discard(C-u)
9. yank(C-y)
10. unix-word-rubox(C-w)前向删除单个词语