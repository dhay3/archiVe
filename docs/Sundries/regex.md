# posix regex

https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap09.html

https://remram44.github.io/regex-cheatsheet/regex.html

https://www.gnu.org/savannah-checkouts/gnu/sed/manual/sed.html#BRE-vs-ERE

regular expression（regex）正则表达式，遵从POSIX.1-2017准则的工具（shell，awk，sed，etc）大都支持两种regex，默认一般使用BRE

Basic Regular Expression，Extended Regular Expression

## RE Bracket Expression

### RE extensions

> 注意区别于高级编程语言，无`\d`

- `\w`

  匹配任意一个单词（前后导空格）

- `\W`

  匹配非单词字符

- `\b`

  匹配字符边界

  ```
  $ echo "abc %-= def." | sed 's/\b/X/g'
  XabcX %-= XdefX.
  ```

- `\s`

  匹配空格

- `\S`

  匹配非空格

- `\<`

  匹配字符开头

  ```
  $ echo "abc %-= def." | sed 's/\</X/g'
  Xabc %-= Xdef.
  ```

- `\>`

  匹配字符结尾

  ```
  $ echo "abc %-= def." | sed 's/\>/X/g'
  abcX %-= defX.
  ```

### predefined bracket expression

具体看sed man page，每个都需要在外层添加`[]`

- `[:alnum:]`

  等价`[0-9A-Za-z]`

- `[:alpha:]`,`[:lower:]`,`[:upper:]`

  匹配alphanumeric characters，字面意思

- `[:blank:]`

  匹配空格和tab

- `[:space:]`

  匹配tab，newline，carriage return，line feed，space

- `[:digit:]`

  等价`[0-9]`

- `[:xdigit:]`

  匹配16进制正出现的数字和字母

## BRE

Basic Regular Expression（BRE）

### ordinanry chars

- `(`,`)`,`{`,`}`
- 1 - 9
- a character inside a bracket expression。（ig：ipv6 address could in [ ]）

### special chars

- `\`

  转义符

- `.` period

  如果在bracket外表示任意一个字符，包括换行符

- `*` asterisk

  匹配一个或多个字符

- `[list]` square-bracket

  匹配list中的任意一个字符，可以使用char1-char2来表示中间的字符（按照ascii）

- `\+`

  和`*`一样，但至少需要匹配一次

- `\?`

  和`.`一样，表示任意一个字符

- `\{i\}`，`\{i,j}\`，`{\i,\}`

  和其他的regex一样表示匹配的次数

- `\(regex\)`

  subexpression

- `regexp1\|rgexp2`

  逻辑或

- `\digit`

  匹配第几个subexpression，复用subexpression

  ```
  $echo aa | grep '\(a\)\1'
  aa
  ```

- `^`，`$`

  same as other regex，present anchor。如果在`[]`中表示取反
  
  | 字符 | 描述                                                         |
  | ---- | ------------------------------------------------------------ |
  | ^    | 匹配输入字符串的开始位置，除非在方括号表达式中使用，当该符号在方括号表达式中使用时，表示不接受该方括号表达式中的字符集合。要匹配 ^ 字符本身，请使用 \^。 |
  | $    | 匹配输入字符串的结尾位置。如果设置了` RegExp` 对象的 `Multiline` 属性，则 $ 也匹配 '\n' 或 '\r'。要匹配 $ 字符本身，请使用 `\$`。 |
  
  看一个`$`的例子
  
  如果没有设置`multiline`只能匹配后面没有换行符的
  
  <img src="/home/cpl/note/docs/Sundries/..\..\imgs\_Net\Snipaste_2020-08-22_21-41-34.png" style="zoom:80%;" />
  
  <img src="/home/cpl/note/docs/Sundries/..\..\imgs\_Net\Snipaste_2020-08-22_21-44-11.png" style="zoom:80%;" />

## ERE

The only difference between basic and extended regular expressions is in the behavior of a few characters: ‘?’, ‘+’, parentheses, braces (‘{}’), and ‘|’. 

*While basic regular expressions require these to be escaped if you want them to behave as special characters, when using extended regular expressions you must escape them if you want them to match a literal character.* 

==即一些 speical chars 无需转译==

## Python/Go/Java regex

https://www.runoob.com/regexp/regexp-syntax.html

https://studygolang.com/pkgdoc `golang regex`

https://regexr.com/ 正则在线检测

> 如果没有开启multi-line，`$`和`^`不匹配结尾和开头

### 非打印字符

| 字符    | 描述                                         |
| ------- | -------------------------------------------- |
| \n      | 换行                                         |
| \r      | 回车                                         |
| \t      | 制表符                                       |
| \s 小写 | ==匹配任何空白字符==, 包括空格, 制表符, 换行 |
| \S 大写 | ==匹配任何非空字符==                         |

### 特殊字符

| 字符 | 描述                                                         |
| ---- | ------------------------------------------------------------ |
| .    | ==匹配除换行符 \n 之外的任何单字符==。要匹配 . ，请使用 `\. `。 |
| \    | 将下一个字符标记为或特殊字符、或原义字符、或向后引用、或八进制转义符。例如， 'n' 匹配字符 'n'。'\n' 匹配换行符。序列 `'\\' `匹配 `"\"`，而` '\(' `则匹配 `"("`。 |
| \|   | 指明两项之间的一个==选择==(或)。要匹配 \， 请使用`\|`        |
| ( )  | 标记一个子表达式的开始和结束位置。子表达式可以获取供以后使用。要匹配这些字符，请使用 `\(` 和 `\)`。 |
| \w   | [0-9A-Za-z_]， 表示从 0 - 9 ，A - Z，a - z 中选出任意一个字符 |
| \d   | [0- 9]                                                       |

### 限定符

| 字符  | 描述                                                         |
| ----- | ------------------------------------------------------------ |
| *     | 匹配前面的子表达式零次或多次。要匹配 * 字符，请使用 `\*`。   |
| +     | 匹配前面的子表达式一次或多次。要匹配 + 字符，请使用 `\+`。   |
| ?     | 匹配前面的子表达式零次或一次，或指明一个==非贪婪限定符==。要匹配 ? 字符，请使用` \?`。 |
| [ ]   | 标记一个中括号表达式的开始和结束位置。==表示从该范围中取任意一个值==，要匹配这些字符，请使用 `\[` 和 `\]`。 |
| {n}   | n 是一个非负整数。匹配确定的 n 次。例如，'o{2}' 不能匹配 "Bob" 中的 'o'，但是能匹配 "food" 中的两个 o。 |
| {n,}  | n 是一个非负整数。至少匹配n 次。例如，'o{2,}' 不能匹配 "Bob" 中的 'o'，但能匹配 "foooood" 中的所有 o。'o{1,}' 等价于 'o+'。'o{0,}' 则等价于 'o*'。 |
| {n,m} | m 和 n 均为非负整数，其中n <= m。最少匹配 n 次且最多匹配 m 次。例如，"o{1,3}" 将匹配 "fooooood" 中的前三个 o。'o{0,1}' 等价于 'o?'。请注意在逗号和两个数之间不能有空格。 |

## Cheat Sheet

| What                                                | [Perl](http://perldoc.perl.org/perlre.html)/PCRE | [Python's `re`](https://docs.python.org/library/re.html) | POSIX (BRE)                           | POSIX extended (ERE)                  | Vim                                      |
| :-------------------------------------------------- | :----------------------------------------------- | :------------------------------------------------------- | :------------------------------------ | :------------------------------------ | :--------------------------------------- |
| Basics                                              |                                                  |                                                          |                                       |                                       |                                          |
| Custom character class                              | `[...]`                                          | `[...]`                                                  | `[...]`                               | `[...]`                               | `[...]`                                  |
| Negated custom character class                      | `[^...]`                                         | `[^...]`                                                 | `[^...]`                              | `[^...]`                              | `[^...]`                                 |
| \ special in class?                                 | yes                                              | yes                                                      | no, `]` escaped if comes first        | no, `]` escaped if comes first        | yes                                      |
| Ranges                                              | `[a-z]`, `-` escaped if first or last            | `[a-z]`, `-` escaped if first or last                    | `[a-z]`, `-` escaped if first or last | `[a-z]`, `-` escaped if first or last | `[a-z]`, `-` escaped if first or last    |
| Alternation                                         | `|`                                              | `|`                                                      | `\|`                                  | `|`                                   | `\|` `\&` (low precedence)               |
| Escaped character                                   | `\033` `\x1B` `\x{1234}` `\N{name}` `\N{U+263D}` | `\x12`                                                   |                                       |                                       | `\%d123` `\%x2A` `\%u1234` `\%U1234ABCD` |
| Character classes                                   |                                                  |                                                          |                                       |                                       |                                          |
| Any character (except newline)                      | `.`                                              | `.`                                                      | `.`                                   | `.`                                   | `.`                                      |
| Any character (including newline)                   |                                                  |                                                          |                                       |                                       | `\_.`                                    |
| Match a "word" character (alphanumeric plus `_`)    | `\w` `[[:word:]]`                                | `\w`                                                     | `\w`                                  | `\w`                                  | `\w`                                     |
| Case                                                | `[[:upper:]]` / `[[:lower:]]`                    |                                                          | `[[:upper:]]` / `[[:lower:]]`         | `[[:upper:]]` / `[[:lower:]]`         | `\u` `[[:upper:]]` / `\l` `[[:lower:]]`  |
| Match a non-"word" character                        | `\W`                                             | `\W`                                                     |                                       |                                       | `\W`                                     |
| Match a whitespace character (except newline)       |                                                  |                                                          | `\s` `[[:space:]]`                    | `\s` `[[:space:]]`                    | `\s` `[[:space:]]`                       |
| Whitespace including newline                        | `\s` `[[:space:]]`                               | `\s`                                                     |                                       |                                       | `\_s`                                    |
| Match a non-whitespace character                    | `\S`                                             | `\S`                                                     | `[^[:space:]]`                        | `[^[:space:]]`                        | `\S` `[^[:space:]]`                      |
| Match a digit character                             | `\d` `[[:digit:]]`                               | `\d`                                                     | `[[:digit:]]`                         | `[[:digit:]]`                         | `\d` `[[:digit:]]`                       |
| Match a non-digit character                         | `\D`                                             | `\D`                                                     | `[^[:digit:]]`                        | `[^[:digit:]]`                        | `\D` `[^[:digit:]]`                      |
| Any hexadecimal digit                               | `[[:xdigit:]]`                                   |                                                          | `[[:xdigit:]]`                        | `[[:xdigit:]]`                        | `\x` `[[:xdigit:]]`                      |
| Any octal digit                                     |                                                  |                                                          |                                       |                                       | `\o`                                     |
| Any graphical character excluding "word" characters | `[[:punct:]]`                                    |                                                          | `[[:punct:]]`                         | `[[:punct:]]`                         | `[[:punct:]]`                            |
| Any alphabetical character                          | `[[:alpha:]]`                                    |                                                          | `[[:alpha:]]`                         | `[[:alpha:]]`                         | `\a` `[[:alpha:]]`                       |
| Non-alphabetical character                          |                                                  |                                                          | `[^[:alpha:]]`                        | `[^[:alpha:]]`                        | `\A` `[^[:alpha:]]`                      |
| Any alphanumerical character                        | `[[:alnum:]]`                                    |                                                          | `[[:alnum:]]`                         | `[[:alnum:]]`                         | `[[:alnum:]]`                            |
| ASCII                                               | `[[:ascii:]]`                                    |                                                          |                                       |                                       |                                          |
| Character equivalents (e = é = è) (as per locale)   |                                                  |                                                          | `[[=e=]]`                             | `[[=e=]]`                             | `[[=e=]]`                                |
| Zero-width assertions                               |                                                  |                                                          |                                       |                                       |                                          |
| Word boundary                                       | `\b`                                             | `\b`                                                     | `\b` / `\<` (start) / `\>` (end)      | `\b` / `\<` (start) / `\>` (end)      | `\<` (start) / `\>` (end)                |
| Anywhere but word boundary                          | `\B`                                             | `\B`                                                     | `\B`                                  | `\B`                                  |                                          |
| Beginning of line/string                            | `^` / `\A`                                       | `^` / `\A`                                               | `^`                                   | `^`                                   | `^` (beginning of pattern ) `\_^`        |
| End of line/string                                  | `$` / `\Z`                                       | `$` / `\Z`                                               | `$`                                   | `$`                                   | `$` (end of pattern) `\_$`               |
| Captures and groups                                 |                                                  |                                                          |                                       |                                       |                                          |
| Capturing group                                     | `(...)` `(?<name>...)`                           | `(...)` `(?P<name>...)`                                  | `\(...\)`                             | `(...)`                               | `\(...\)`                                |
| Non-capturing group                                 | `(?:...)`                                        | `(?:...)`                                                |                                       |                                       | `\%(...\)`                               |
| Backreference to a specific group.                  | `\1` `\g1` `\g{-1}`                              | `\1`                                                     | `\1`                                  | `\1` non-official                     | `\1`                                     |
| Named backreference                                 | `\g{name}` `\k<name>`                            | `(?P=name)`                                              |                                       |                                       |                                          |
| Look-around                                         |                                                  |                                                          |                                       |                                       |                                          |
| Positive look-ahead                                 | `(?=...)`                                        | `(?=...)`                                                |                                       |                                       | `\(...\)\@=`                             |
| Negative look-ahead                                 | `(?!...)`                                        | `(?!...)`                                                |                                       |                                       | `\(...\)\@!`                             |
| Positive look-behind                                | `(?<=...)`                                       | `(?<=...)`                                               |                                       |                                       | `\(...\)\@<=`                            |
| Negative look-behind                                | `(?<!...)`                                       | `(?<!...)`                                               |                                       |                                       | `\(...\)\@<!`                            |
| Multiplicity                                        |                                                  |                                                          |                                       |                                       |                                          |
| 0 or 1                                              | `?`                                              | `?`                                                      | `\?`                                  | `?`                                   | `\?`                                     |
| 0 or more                                           | `*`                                              | `*`                                                      | `*`                                   | `*`                                   | `*`                                      |
| 1 or more                                           | `+`                                              | `+`                                                      | `\+`                                  | `+`                                   | `\+`                                     |
| Specific number                                     | `{n}` `{n,m}` `{n,}`                             | `{n}` `{n,m}` `{n,}`                                     | `\{n\}` `\{n,m\}` `\{n,\}`            | `{n}` `{n,m}` `{n,}`                  | `\{n}` `\{n,m}` `\{n,}`                  |
| 0 or 1, non-greedy                                  | `??`                                             | `??`                                                     |                                       |                                       |                                          |
| 0 or more, non-greedy                               | `*?`                                             | `*?`                                                     |                                       |                                       | `\{-}`                                   |
| 1 or more, non-greedy                               | `+?`                                             | `+?`                                                     |                                       |                                       |                                          |
| Specific number, non-greedy                         | `{n,m}?` `{n,}?`                                 | `{n,m}?` `{n,}?`                                         |                                       |                                       | `\{-n,m}` `\{-n,}`                       |
| 0 or 1, don't give back on backtrack                | `?+`                                             |                                                          |                                       |                                       |                                          |
| 0 or more, don't give back on backtrack             | `*+`                                             |                                                          |                                       |                                       |                                          |
| 1 or more, don't give back on backtrack             | `++`                                             |                                                          |                                       |                                       |                                          |
| Specific number, don't give back on backtrack       | `{n,m}+` `{n,}+`                                 |                                                          |                                       |                                       |                                          |
| Other                                               |                                                  |                                                          |                                       |                                       |                                          |
| Independent non-backtracking pattern                | `(?>...)`                                        |                                                          |                                       |                                       | `\(...\)\@>`                             |
| Make case-sensitive/insensitive                     | `(?i)` / `(?-i)`                                 | `(?i)` / `(?-i)`                                         |                                       |                                       | `\c` / `\C`                              |

## example

### 电话正则：

- 通用规则

  `1[34578]\d{9}`

### 身份证匹配：

- 通用规则

  `[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[\dXx]`

  ```go
  区号[1-9]\d{5}
  年(18|19|20)\d{2}
  月((0[1-9])|(1[0-2]))
  日(([0-2][1-9])|10|20|30|31) 
  校验码\d{3}[\dXx]  //[\dXx]表示从该范围中取任意一字符
  ```

### 超链接：

- 通用规则

  ``http(s)?://\w+\.\w+(\.\w+)?``

### 邮箱正则：

- qq邮箱

  `\d+@qq\.com`

- 163邮箱

  `\w+@163\.com`

- 通用规则

  `\w+@\w+\.\w+`