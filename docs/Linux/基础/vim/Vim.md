# Vim

参考：

https://coolshell.cn/articles/5426.html

https://github.com/wsdjeg/vim-galore-zh_cn

> 如果操作做错了，可以使用`u`撤回，`ctrl+r`反撤销
>
> 使用`:sp`可以分屏, `:e`可以编辑新文件

## 概述

Vim是一个文本编辑工具，是基于`vi`的克隆，在兼容模式下运行vim意味者使用vi的默认设置，而不是vim的默认设置。除非新建一个用户`vimrc`或者使用`vim -N`启动vim，否则就是在兼容模式下运行。

我们可以使用`vimtutor`命令来了解vim的基础信息，或是`:help usr_01.txt`

==右下角的`39:5`，表示光标的坐标==

## 概念

- **缓冲区**：每一个文件都有自己的缓冲区，文件中的内容做为缓冲区的一部份。
- **窗口**：窗口时缓冲区的view，如果想要多种窗口布局，使用标签页

> tip：使用`:ls`列出所有缓冲区
>
> ```
> :ls
>   1 #    "1.txt"                        line 1
>   2 %a   "bak.xml"                      line 1090
> 
> ```

- **已激活缓冲区**：当前窗口上显示的缓冲区
- **隐藏缓冲区**：file1已激活缓冲区，使用`e file2`加载的文件，这时file1就会变为隐藏缓冲区
- **标签页**：和浏览器的tab类似
- **Normal-mode**：正常模式，包括简单的范围选中，复制，粘贴，移动等操作
- **Insert-mode**：插入模式，`i`在当前位置插入，==`a`在当前位置的下一个位置插入，`o`在当前位置的下一行位置插入一行空格，并格式化前导空格==
- **Command-mode**：命令行模式，`/`或`:`进入的模式，包括复杂的范围选中
- **Repalce-mode**：替换模式，`R`，编辑文本时直接替换内容

简单的说，如果你启动 Vim 的时候没有附带任何参数，你会得到一个包含着一个呈现一个缓冲区的窗口的标签。

## version

在vim中我们可以使用`:version`来查看当前运行的vim

```
VIM - Vi IMproved 8.2 (2019 Dec 12, compiled May 12 2020 02:37:13)
Included patches: 1-716
Modified by team+vim@tracker.debian.org
Compiled by team+vim@tracker.debian.org
Huge version without GUI.  Features included (+) or not (-):
+acl               +cryptv            +fork()            +modify_fname      +persistent_undo   +syntax            +visualextra
+arabic            +cscope            +gettext           +mouse             +popupwin          +tag_binary        +viminfo
+autocmd           +cursorbind        -hangul_input      -mouseshape        +postscript        -tag_old_static    +vreplace
```

1. vim的版本号
2. 补丁版本包

5. 告诉你当前的vim是否支持GUI。从终端运行`gvim`或是在vim中运行`:gui`。另一个重要的信息是 `Tiny` 和 `Huge`。Vim 的特性集区分被叫做 `tiny`，`small`，`normal`，`big` and `huge`，所有的都实现不同的功能子集。

   特性列表告诉你当前vim中安装的特性

## moving cursor

参考：`:help usr_03.txt`

| h 或 向左箭头键(←)               | 光标向左移动一个字符                                         |
| -------------------------------- | ------------------------------------------------------------ |
| j 或 向下箭头键(↓)               | 光标向下移动一个字符                                         |
| k 或 向上箭头键(↑)               | 光标向上移动一个字符                                         |
| l 或 向右箭头键(→)               | 光标向右移动一个字符                                         |
| $(在输入模式下表示最后一行)或END | 光标移动到当前行的最后一个字符                               |
| num$                             | 移动指定行数的最后一个字符                                   |
| w                                | 移动光标到下一个单词                                         |
| b                                | 移动光标到上一个单词                                         |
| 0或HOME                          | 光标移动到当前行的第一个字符                                 |
| gg                               | 移动到第一行                                                 |
| G                                | 移动到最后一行                                               |
| num j                            | 当前行向下多少行。可以将j省略                                |
| `f alpha`                        | 移动到当前行的第一个出现alpha的位置，大写表示反向寻找        |
| %                                | 如果光标在`(`那么使用该符号，会移动到匹配的`)`，同样也可以匹配`[]`,`{}` |
| H，M，L                          | H移动可见的第一行，M移动到可见的中间，L移动到可见的最后一行  |
| zz                               | 在窗口中将光标的位置置中                                     |

## set

1. set nu：设置行号
2. set ignorecase：搜索时忽略大小写，set noignorecase取消
3. set hlsearch：将搜索出来的结果高亮显示

## Search

`:help usr_03.txt`

1. `/`向下搜索，`?`向上搜索

2. 使用`n`向下一个，`N`上一个。
3. 支持通配符，`.`，`*`
4. 如果`.*[]^%/\?~$`想要表达原来的意思需要添加转义符`\`

### 精确匹配

- `/the\>`

  表示匹配只有以the结尾的单词，不会匹配there，但是会匹配athe

- `/\<the`

  表示只匹配以the开头的单词，不会匹配father，但是会匹配they

- `/\<the\>`

  精确匹配，只会匹配the

## Ranges

使用`:h range`来查看，一般与operators一起使用

- 大部分命令默认作用于当前行

- 一般在操作后面跟上范围。但是除了`dd`,`yy`外，需要先声明数字。如`10dd`
- `:write`和`:global`代表所有行，即`1,$`

| 操作  |                  |
| ----- | ---------------- |
| `.`   | 当前行           |
| `.,5` | 当前行到第5行    |
| `1,$` | 第一行到最后一行 |
| `%`   | `1,$`的语法糖    |

## Operatos

使用`:h operator`来查看所有的动作，一般配合ranges一起使用，可以在

| 操作      |                                                              |
| --------- | ------------------------------------------------------------ |
| y（yank） | 复制（将选中的复制到寄存器中），输入`yy`表示选中并确认，替代回车 |
| gu        | 将选中的内容小写                                             |
| gU        | 将选中的内容大写                                             |
| <         | 去掉前导空格                                                 |
| >         | 去掉后缀空格                                                 |
| d         | 删除，输入`dd`表示删除并确认，替代回车                       |
| zf        | 将选中的内容折叠，以`+-- 16 lines: ** TYPING A NUMBER WITH AN OPERATOR REPEATS IT THAT MANY TIMES. **`形式显示，进入输入模式自动展开 |
| g？       | 将选中的内容使用ROT13编码                                    |
| c         | 见选中的内容删除，置为输入模式                               |
| p         | 在光标的下一行粘贴选中的内容                                 |
| J         | 将下面指定行数的内容和当前行合并成一行                       |
|           |                                                              |
|           |                                                              |

## replace

| /word                                          | 搜索文本                                                     |
| ---------------------------------------------- | ------------------------------------------------------------ |
| :n1,n2s/word1/word2/g                          | n1 与 n2 为数字。在第 n1 与 n2 行之间寻找 word1 这个字符串，并将该字符串取代为 word2 ！举例来说，在 100 到 200 行之间搜寻 vbird 并取代为 VBIRD 则：『:100,200s/vbird/VBIRD/g』。(常用) |
| `:1,$s/word1/word2/g` 或 `:%s/word1/word2/g`   | 从第一行到最后一行寻找 word1 字符串，并将该字符串取代为 word2 ！(常用) |
| `:1,$s/word1/word2/gc` 或 `:%s/word1/word2/gc` | 从第一行到最后一行寻找 word1 字符串，并将该字符串取代为 word2 ！且在取代前显示提示字符给用户确认 (confirm) 是否需要取代！(常用) |

## Getting out

1. `wq`

   写入并保存

2. `wq!`

   强制写入并保存

3. `q`

   退出

4. `q!`

   退出不保存

5. `ZZ`

   退出并保存