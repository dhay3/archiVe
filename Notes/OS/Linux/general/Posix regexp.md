# Posix regexp

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

## 