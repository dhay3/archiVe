Dict protocol

ref:

https://curl.se/rfc/rfc2229.txt

https://curl.se/mail/archive-2015-12/0011.html

## 0x1 Digest

dict 是一个基于 TCP 的 L7 协议，顾名思义就是查字典（查找不同的数据库）。可以通过 curl 来使用该协议，dict://dict.org 是常用的站点

## 0x2 Terms

### database

词库，freely-distributable database ，其中Free online dictionary of computing (FOLDOC)包含多种语言，==目前没有中的数据库==

### url

dict 通常使用两种格式的url

1. 其中的 d 表示使用define命令

```
dict://<user>;<auth>@<host>:<port>/d:<word>:<database>:<n>
```

2. 其中的 m 表示使用match命令

```
dict://<user>;<auth>@<host>:<port>/m:<word>:<database>:<strat>:<n>
```

`<user>;<auth>@`可以不用指定表示无authentication

`<port>`也可以是被忽略，默认端口 2628

`<database>`如果被忽略需要使用`!`来占位

`<start>`如果被忽略需要使用`.`来占位

### commands

需要在 dict 服务器上执行的命令，使用以下 EBNF 来定义，不区分大小写

```
command     = cmd-word *<WS cmd-param>
cmd-word    = atom
cmd-param   = database / strategy / word
database    = atom
strategy    = atom
```

## 0x3 Commands

### help

查看帮助信息

```
cpl in ~ λ curl dict.org/help      
220 dict.dict.org dictd 1.12.1/rf on Linux 4.19.0-10-amd64 <auth.mime> <110649318.6487.1650208645@dict.dict.org>
250 ok
113 help text follows
DEFINE database word         -- look up word in database
MATCH database strategy word -- match word in database using strategy
SHOW DB                      -- list all accessible databases
SHOW DATABASES               -- list all accessible databases
SHOW STRAT                   -- list available matching strategies
SHOW STRATEGIES              -- list available matching strategies
SHOW INFO database           -- provide information about the database
SHOW SERVER                  -- provide site-specific information
OPTION MIME                  -- use MIME headers
CLIENT info                  -- identify client to server
AUTH user string             -- provide authentication information
STATUS                       -- display timing information
HELP                         -- display this help information
QUIT                         -- terminate connection

The following commands are unofficial server extensions for debugging
only.  You may find them useful if you are using telnet as a client.
If you are writing a client, you MUST NOT use these commands, since
they won't be supported on any other server!

D word                       -- DEFINE * word
D database word              -- DEFINE database word
M word                       -- MATCH * . word
M strategy word              -- MATCH * strategy word
M database strategy word     -- MATCH database strategy word
S                            -- STATUS
H                            -- HELP
Q                            -- QUIT
.
250 ok
221 bye [d/m/c = 0/0/0; 0.000r 0.000u 0.000s]
```

### define

从指定的database中查找指定的词，如果database的位置使用了`!`表示查找所有的database，直到找到一个匹配的后停止，如果没有找到该词会返回 552 code

```
cpl in ~ λ curl -Ss dict.org/d:hello
220 dict.dict.org dictd 1.12.1/rf on Linux 4.19.0-10-amd64 <auth.mime> <110648506.5533.1650208295@dict.dict.org>
250 ok
150 1 definitions retrieved
151 "Hello" gcide "The Collaborative International Dictionary of English v.0.48"
Hello \Hel*lo"\, interj. & n.
   An exclamation used as a greeting, to call attention, as an
   exclamation of surprise, or to encourage one. This variant of
   {Halloo} and {Holloo} has become the dominant form. In the
   United States, it is the most common greeting used in
   answering a telephone.
   [1913 Webster +PJC]
```

如果database使用了`*`表示查找所有database（dict.org使用all替代），并列出匹配的，等价与show db

```
cpl in ~ λ curl dict.org/d:和:all
cpl in ~ λ curl dict.org/d:和:fd-jpn-eng
```

### match

按照词语匹配可能的单词，match 支持两种策略

1. excat 精确匹配
2. prefix 最长前缀匹配

```
cpl in ~ λ curl -Ss dict.org/m:hello
220 dict.dict.org dictd 1.12.1/rf on Linux 4.19.0-10-amd64 <auth.mime> <110648430.5449.1650208264@dict.dict.org>
250 ok
152 8 matches found
gcide "Hell"
gcide "Hello"
gcide "Cello"
gcide "Jell-O"
gcide "Hollo"
gcide "Hullo"
gcide "Helio-"
gcide "Helly"
```