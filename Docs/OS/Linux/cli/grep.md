# Linux grep

ref:

http://www.zsythink.net/archives/1733

## Digest

syntax:

```
grep [OPTION...] PATTERNS [FILE...]
#等价于 egrep
grep [OPTION...] -e PATTERNS ... [FILE...]
#等价与 fgrep
grep [OPTION...] -f PATTERN_FILE ... [FILE...]
```

grep 是一个从 FILE 中找到 PATTERNs 的工具，通常 PATTERENS 需要被 qutoed

## Positional args

如果 FILE 的位置是`-`表示使用 stdin

如果 FILE 的位置为空，recursive searches 会查看当前目录。如果 nonrecursive searches read standard input

## Optional args

### Pattern args

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

### Ouput args

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

  当匹配到 NUM 行后，停止匹配 

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

## Pattern/Regexp

> 这里的 regexp 也是 shell 中的，需要对比区别与其他 高级编程语言中的 regexp
>
> 查看 [Regexp]()

A regular expression is a pattern that describes a set of strings

grep 可以识别 3 种 regexp —— BRE, ERE, PCRE，另外在 GNU grep 中 BRE 和 ERE 是没有区别的

### Symbolic

下面介绍一下符号含义

#### characters/letters/digits

represent themselves

字面意思，没有特殊含义

#### backslash

转义符

#### period

`.` 

匹配任意一个字符

#### asterisk

`*`

匹配一个或多个任意字符

#### bracket

`[]` bracket 根据使用的场景有多种含义

1. range expression

   例如`[abc]`, `[a-z]`

   匹配 bracket 中 range 指定的任意一个字符。range 可以是指明范围内的字符，也可以是使用 hypen(短横) 连接的

2. not expression

   例如`[^abc]`

   取反，匹配 bracket 中 非 range 指定的任意一个字符。range 可以是指明范围内的字符，也可以是使用 hypen(短横) 连接的

3. predefined expression

   `[:alnum:]` 等价与 `[0-9A-Za-z]`

   `[:alpha:]` 等价与 `[A-Za-z]`

   `[:digit:]` 等价与 `[0-9]`

   `[:xdigit:]` 匹配 hex 中出现的数字和字母

   `[:blank:]` 匹配 space 和 tab

   `[:space:]` 匹配 space, tab, newline, carriage return, line feed

    `[:cntrl:]`,  `[:graph:]`, `[:lowewr:]`, `[:print:]`, `[:punct:]`, `[:upper:]`

   顾名思义

   需要注意的是这些只是 symbolic names, 如果需要使用还需要在外面加一层`[]`, 例如`[[:alpha:]]`

   同样的`^`也表示取非, 例如`[^[:alpha:]]`，表示取字母的反，即数字

#### anchoring

1. caret `^word`

   匹配首字符是 word 的

2. dollor sign `word$`

   匹配尾字符是 word 的

#### extensions

> 需要注意的是在 ERE 或是 BRE 中没有 `\d`

1. `\<`

   匹配字符开头

   ```
   $ echo "abc %-= def." | sed 's/\</X/g'
   Xabc %-= Xdef.
   ```

2. `\>`

   匹配字符结尾

   ```
   $ echo "abc %-= def." | sed 's/\>/X/g'
   abcX %-= defX.
   ```

3. `\bword`

   匹配 word edge 是空字符的

   ```
   $ echo "abc %-= def." | sed 's/\b/X/g'
   XabcX %-= XdefX.
   ```

4. `\Bword`

   匹配 word edge 不是空字符的

5. `\s`

   匹配空格

6. `\S`

   匹配非空格

#### repetition

1. `?`

   the preceding item is optional and matched at most once

2. `*`

   the preceding item will be matched zero or more times

3. `+`

   the preceding item will be matched one or more times

4. `{n}`

   the preceding item is matched exactly n times

5. `{n,}`

   the preceding item is matched n or more times

6. `{,m}`

   the preceding item is matched at most m times

   这只有在 GNU 中有

7. `{n,m}`

   the preceding item is matched at least n times, but not more than m times

#### concatenation

两个 regexp 之间默认使用 and 逻辑拼接

#### alternation

`|`，两个 regexp 使用 or 逻辑拼接

#### subexpression

`()` parenthesized 中的表示子表达式，`\n` 如果 n 是一个数字，表示匹配之前的第 n 个子表达式

## BRE vs ERE

The only difference between basic and extended regular expressions is in the behavior of a few characters: ‘?’, ‘+’, parentheses, braces (‘{}’), and ‘|’. 

*While basic regular expressions require these to be escaped if you want them to behave as special characters, when using extended regular expressions you must escape them if you want them to match a literal character.* 

BRE 和 ERE 的区别主要在于，在 BRE 中 `?`, `+`, `{`, `|`, `(`, `)` 没有任何特殊含义。需要使用 backslash 转义才有特殊含义，而 ERE 不需要

## Exit status

如果有匹配到的内容以 0 为 exit code

如果没有匹配到的内容以 1 为 exit code

如果报错以 2 为 exit code







