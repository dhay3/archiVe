# Linux grep

## 0x01 Overview

syntax

```
grep [OPTION...] PATTERNS [FILE...]
#等价于 egrep
grep [OPTION...] -e PATTERNS ... [FILE...]
#等价与 fgrep
grep [OPTION...] -f PATTERN_FILE ... [FILE...]
```

`grep` 是一个从 FILE 中找到 PATTERNs 的工具，通常 PATTERENS 需要被 qutoed(如果不，有些符号会被 Shell 解析成模式扩展)

## 0x02 Positional args

如果 FILE 的位置是 `-` 表示使用 stdin

如果 FILE 的位置为空

如果是 recursive searches 会查看当前目录

如果是 nonrecursive searches 会从 stdin 中读取

## 0x03 Optional args

### Pattern args

> 默认是 BRE，推荐使用 PCRE

- `-E | --extended-regexp`

  Interpret PATTERNS as extended regular expressions(EREs)

- `-F | --fixed-strings`

  Interpret PATTERNS as ==fixed strings, not regular expressions==

- `-G | --basic-regexp`

  Interpret PATTERNS as basic regular expressions(BREs)

  ==this is the default==

- `-P | --perl-regexp`

  Interpret PATTERNS as perl-compatibel regular expressions(PCREs)

### Matching args

- `-i | --ignore-case`

  Ignore case distinctions in patterns and input data

  忽略大小写敏感，==grep 默认大小写敏感==

- `--no-ignore-case`

  大小写敏感，缺省值

- `-v | --invert-match`

  invert the sense of matching

  等于与 regexp 中`[^word]`表示取反，即匹配内容中无 PATTERN 的

  ```
  [root@chz etc]# grep -v root passwd
  ...输出的内容不包含root
  ```

- `-w | --word-regexp`

  select only those lines containing matches that form whole words

  ==匹配子串==

- `-x | --line-regexp`

  select only those matches that exactly math the whole line

  ==只有整行的内容匹配到 PATTERN 才匹配==

- `-q | --quiet | --silent`

  do not write anything to standard output. Exit immediately with zero status if any match is found, even ==if an error was detected==

- `-c | --count`

  suppress normal ouput; instead print a count of mathing lines for each input file

  只输出匹配内容的一共有几行，所以尽量避免和使用`wc`，显得没看过 man page

- `--color[=WHNE]`

  colorful ouput，WHEN 的值可以是 never, always, auto

- `-L | --files-without-math`

  ==打印没有匹配 PATTERN 内容的文件名==

  ```
  cpl in ~/gitcwdon master ● λ grep -LR traceroute
  ls.md
  mv.md
  grep.md
  ```

- `-l | --files-with-match`

  ==打印匹配 PATTEREN 内容的文件名==

  ```
  cpl in ~/gitcwd on master ● λ grep -l traceroute ./*  
  grep: ./built-in: Is a directory
  ./grep.md
  ./mtr.md
  grep: ./network: Is a directory
  ./ping.md
  ./traceroute.md
  ```

- `-m NUM | --max-count=NUM`

  当匹配有 NUM 行后，停止匹配 

- `-A NUM | --after-context=NUM`

  在输出匹配 PATTERN 的内容后同时打印出后 NUM 行

- `-B NUM | --before-context=NUM`

  在输出匹配 PATTERN 的内容后同时打印出前 NUM 行

- `-H | --with-filename`

  print the file name for each match. this is the default when there is more than one file to search

- `-h | --no-filename`

  suppress the prefixing of file name on output. This is the default when there is only one file to search

- `-n | --line-number`

  prefix each line of output

  输出的内容会带上一个序号字段，对应在输入文件中的行数或者 stdin 中的行数

### File and directory selection

- `-a | --text`

  process a binary file as if it were text

  grep 默认不会读取 binary file

- `--exclude=GLOB`

  exclude 目录下匹配 BREs 的文件 ，这里叫 GLOB 也没错。bash 中一般称为 globbing

- `--include=GLOB`

  include 同上

- `-r | --recursive`

  read all files under each directory without followling symbolic link

- `-R | --dereference-recursive`

  read all files under each directory , following symbolic link

## 0x04 Exit status

如果有匹配到的内容以 0 为 exit code

如果没有匹配到的内容以 1 为 exit code

如果报错以 2 为 exit code

**references**

[^1]:http://www.zsythink.net/archives/1733