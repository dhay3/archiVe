# Linux man 

ref:

https://superuser.com/questions/346703/linux-apropos-command-always-returns-nothing-appropriate

https://www.howtogeek.com/682871/how-to-create-a-man-page-on-linux/

> 无法查看内建命令，如果想要查看内建命令
>
> 1. 非bash使用`bash -c "help command"`
>2. bash使用`help command`
> 
> man page 全文支持使用 ERE 过滤

## Digest

syntax:

```
man [man options] [[section] page ...] ...
man -k [apropos options] regexp ...
man -K [man options] [section] term ...
man -f [whatis options] page ...
man -l [man options] file ...
man -w|-W [man options] page ...
```

man is the system’s manual pager

系统通过 man 来管理 manual page

默认 man 会按照 predefined 顺序查找所有匹配的 sections。==但是只会显示匹配的 first page，即使 page 在其他许多的 sections 中存在==

sections 如下：

1.  Executable programs or shell commands
2.  System calls (functions provided by the kernel)
3.  Library calls (functions within program libraries)
4.  Special files (usually found in /dev)
5.  File formats and conventions, e.g. /etc/passwd
6.  Games
7.  Miscellaneous (including  macro  packages  and  conventions),  e.g.
        man(7), groff(7)
8.  System administration commands (usually only for root)
9.  Kernel routines [Non standard]

一般在`man command`后显示在左上角，例如`sudo(8)`表示第 8 章

## nroff

> 如果需要编写 man page 可以使用 pandoc
>
> 具体查看
>
> [Pandoc]()

man pages 一般以 nroff（在 GNU 体系中一般是 groff） 的格式存储在`/usr/share/man`下

具体查看

[norff]()

## Output

通常一个完整的 man page 都会包含

1. NAME
2. SYNOPSIS
3. CONFIGURATION
4. DESCRIPTION
5. OPTIONS
6. EXIT STATUS
7. RETURN VALUE
8. ERRORS
9. ENVIRONMENT
10. FILES
11. VERSIONS
12. CONFORMING TO
13. NOTES
14. BUGS
15. EXAMPLE
16. AUTHORS

在 man page 中可能会出现一些有特殊含义的标识

1. bold text

   type exactly as shown

2. italic text

   replace with appropriate argument

3. [-abc]

   any or all arguments with [ ] are optional

4. -a | -b

   options delimited by | cannot be used together

5. argument...

   argument is repeatable

6. [expression] ...

   entire expression within [ ] is repeatable

## mandb

mandb is used to initialise or manually update index data caches of the manual page system and the information

mandb 用于更新和 `whatis`, `appropos` 相关命令的 index cache

## Optional args

> 如果`-f`或`-k`显示 nothing appropriate，但是实际有对应的 cli
>
> 需要执行 `sudo mandb`

- `-f | --whatis`

  等价与 `whatis`

  display one-line manual page descriptions

- `-k | --apropos`

  等价与 `apropos`

- `-i | --ignore-case`

  搜索 man page 时，大小写不敏感

- `-I | --match-case`

  搜索 man page 时，大小写敏感

## Examples

1. `man stat.2` 和 `man 2 stat` 等价，都表示查看 sections 2 中的 stat manual page

2. `man -k stat`

3. `man ascii`

   ascii 表快查

