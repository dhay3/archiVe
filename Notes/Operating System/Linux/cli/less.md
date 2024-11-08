# less

## 0x01 Overview

syntax

```
       less -?
       less --help
       less -V
       less --version
       less [-[+]aABcCdeEfFgGiIJKLmMnNqQrRsSuUVwWX~]
            [-b space] [-h lines] [-j line] [-k keyfile]
            [-{oO} logfile] [-p pattern] [-P prompt] [-t tag]
            [-T tagsfile] [-x tab,...] [-y lines] [-[z] lines]
            [-# shift] [+[+]cmd] [--] [filename]...
```

`less` 类似 `more`, 但是允许 backward movement 和 forward movement(`more` 同样也支持，但是每次以一屏为单位滚动)。同时 `less` 不会读取文件所有的内容，所以对大文件的读取比 `cat`, `vim` 都快

## 0x02 Commands

`less` 和 `vim` 类似，支持多种指令 (大多数和 `vim` 类似，具体查看 man page) 这里只列举常用的以及方便记忆的 commands 

- `h`

  查看帮助信息

- `j` | `downarrow` | `enter`

  向下一行

- `space`

  向下一屏

- `d`

  向下半屏

- `y` | `k`

  向上一行

- `u`

  向上半屏

- `/pattern`

  向下查询

- `?pattern`

  向上查询

- `n`

  向下匹配查询

- `N`

  向上匹配查询

- `v`

  进入编辑模式，使用 `$VISUAL` 变量指定的进程

- `-I`

  查询 pattern 时忽略大小写（在 Linux 中默认大小写敏感）



