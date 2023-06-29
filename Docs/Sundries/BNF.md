# BNF & EBNF

## What is BNF?

Backus-Naur form(BNF) 是一种用于描述 syntax of languages 的范式, 也可以直接理解成 the language of languages

被大量的用在计算机领域的文档中, 用于表示语法的使用方式

例如

Java 中类修饰符的 BNF 可以使用如下方式表示

```
<class modifiers> ::= <class modifier> | <class modifiers> <class modifier>
<class modifier> ::= public | abstract | final 
```

*A BNF sepcification is a set of derivation rules*

> 可以把 BNF 想象成 推导 或者是 等价 的关系

```
 <name> ::= expansion | <nick-name>
 <nick-name> ::= expansion
```

BNF 范式有几条规则

1. `::=` 表示为 may expand into 或者是 may be replaced with
2. `<name>` 部分也被称为 non-terminal symbols
3. non-terminal 必须要要在 `< >`(angle brackets) 内，不管在 `::=` 左侧还是右侧
4. expansion 由 terminal symbol 和 non-terminal symbols组成
5. `|`(vertical bar) 表示或者

## Terminal symbols/Non-terminal symbols

在 BNF 中最小的单元是 terminal symbols 也被称为 tokens，non-terminal symbols 由 terminal symbols 或者 non-terminal symbols 组成

可以表示为如下 BNF 规则

```
<terminal symbol> ::= "identifier" | "literal" | "keyword" | "white space" | "comment" | "delimiter"
<non-terminal symbol> ::= <terminal symbol> | <terminal symbol> "+" <non-terminal symbol>
```

> 这里 non-terminal symbol 表示的方式不准确，因为 BNF 不能表示组合以及变参

### Terminal symbols

terminal symbols 可以使用 a literal 例如 “a literal word” 就表示字面意思，也可以是语言中的 keywords 例如 JAVA 中的 “switch” 或者 Python 中的 “def”，也可以是分隔符号 例如 “，” 等等

> keywords 并没有规则规定必须在 double quoted 内，所以为了方便记忆均无需在 double quoted 内

具体可以使用的 terminal symbols 形式可以参考下表(Definition 部分使用 regexp 表示)

| **Category**          | **Terminal**                        | **Definition**                                               | **Note**                                                     |
| --------------------- | ----------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| *Identifier*          | Java identifier                     | `/[a-zA-Z$_][a-zA-Z0-9$_]*/`                                 | This is a simplification because it is technically possible to use also UTF escape sequences in Java identifiers |
|                       | Haskell type identifier             | `/[A-Z][a-zA-Z0-9$_’]*/`                                     |                                                              |
|                       | Haskell value identifier            | `/[a-z][a-zA-Z0-9_’]*/`                                      |                                                              |
| *Keyword*             | Some Java keywords                  | `“abstract”, “assert”, “boolean”, “break”, “byte”, “case”, “catch”, “char”, “class”, “const”…` | “const” is a keyword in Java even if it is not used by any construct. It is a  reserved keyword for future usages (same is true for “goto”) |
|                       | Some Python 3 keywords/td>          | `“def”, “return”, “raise”, “from”, “import”, “as”, “global”, “nonlocal”, “assert”…` |                                                              |
| *Literal*             | Java string literal                 | `/'”‘ (~[“] | ” [\b\t\n\f\r”‘])* ‘”‘/`                       | This is a simplification. We are ignoring the octal and unicode escape sequences |
|                       | Java character literal              | `/”’ (~[‘] |” [\b\t\n\f\r”‘]) ”’`                            |                                                              |
|                       | Java integer literal                | `/[“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?/`                    | A Java integer literal can actually be expressed in decimal, hexadecimal, octal or binary format. We are just considering the decimal format  here. This is true also for Java long literals |
|                       | Java long literal                   | `/[“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?(‘l’|’L’)`            |                                                              |
|                       | Java float literal                  | `/[“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?’.'([“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?)?(‘f’|’F’)` | A Java float literal can actually be expressed in decimal or hexadecimal  format. We are just considering the decimal format here. We are also  ignoring the possibility of specifying the exponent |
|                       | Java boolean literal                | `/”true”|”false”/`                                           |                                                              |
| *Separator/Delimiter* | Some Java separators and delimiters | `“(“, “)”, “{“, “}”, “,”, “;”…`                              |                                                              |
|                       | Some Ruby separators and delimiters | `“,”, “;”…`                                                  |                                                              |
| *Whitespace*          | Java whitespace                     | `/[ \t\r\n\u000C]+/`                                         |                                                              |
|                       | Ruby whitespace                     | `/(‘ ‘|’\t’)+/`                                              |                                                              |
| *Comment*             | Java line comment                   | `/’//’ ~[\r\n]*/`                                            |                                                              |
|                       | Java block comment                  | `/’\/*’ .*? ‘*\/’/`                                          |                                                              |
|                       | Python line comment                 | `/’#’ ~[\r\n\f]*/`                                           |                                                              |

### Non-terminal symbols

Non-terminal symbols 由一个或者多个 terminal symbols 组成

例如

```
<expr> ::= <term> "+" <expr>  |  <term>
<term> ::= <factor> "*" <term>  |  <factor>
<factor> ::= "(" <expr> ")"  |  <const>
<const> ::= integer
```

-  `<const>::=integer` 中 `<const>` 是一个 non-terminal symbol，`integer` 是 terminal symbols
- `<factor>::="(" <expr> ")" | <const>` 中 `<factor>` 是 non-terminal symbol，由 `“(”` terminal symbol，`<expr>` non-terminal symbol， `<const>` non-terminal symbol 组成

可以将 non-terminal symbol 和 terminal symbol 的关系想象成从下往上的树结构

![2021-08-12_12-22](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-08-12_12-22.4oq5txwzat20.png)

## EBNF

Extended-BNF 是 BNF 的扩展

修改了一些规则

- 将 `::=` 改成 `=`
- non-terminal 无需在 `< >`(angled brackets) 内
- `;`(semicolon) 后面部分是 comment

增加了一些规则

- `{ }`(curly brackets) 表示 repeated zero or more times.
- `[ ]`(square brackets) 表示 optional
- `( )`(parenthness) 表示 grouping
- `,` 表示 concatention(并列)

> 逻辑上大体和 regex 一样

例如

sudoer files(`/etc/sudoers`) 就是使用 EBNF 的

```
Cmnd_Alias  PROCESSES = /usr/bin/nice, /bin/kill, /usr/bin/renice, /usr/bin/pkill, /usr/bin/top
Cmnd_Alias  REBOOT = /sbin/halt, /sbin/reboot, /sbin/poweroff
```

## ABNF

Augmented-BNF 同样也是 BNF 的扩展，更多的被用在 IEEE RFC 中

ABNF 和 EBNF 类似

相同的是

- 将 `::=` 改成 `=`
- non-terminal 无需在 `< >`(angled brackets) 内
- `[ ]`(square brackets) 表示 optional
- `( )`(parenthness) 表示 grouping
- `;`(semicolon) 后面部分是 comment

不同的是

- BNF 中表示或者的符号 `|` ，使用 `/` 替代

- EBNF 中表示 repeation 的符号`{ }`, 使用 `*` 替代，如果想要表示 `*` 原来的意思，使用 literal 即 `“*”` 或者 `‘*’`

  如果需要表示重复 n 次，使用 `n*`

  如果需要表示重复 n - m 次，使用 `n*m`

  例如 `3*Rule`, `1*3Rule`

例如

```
date-time       =   [ day-of-week "," ] date time [CFWS]
day-of-week     =   ([FWS] day-name) / obs-day-of-week
day-name        =   "Mon" / "Tue" / "Wed" / "Thu" /
"Fri" / "Sat" / "Sun"
date            =   day month year
day             =   ([FWS] 1*2DIGIT FWS) / obs-day
month           =   "Jan" / "Feb" / "Mar" / "Apr" /
"May" / "Jun" / "Jul" / "Aug" /
"Sep" / "Oct" / "Nov" / "Dec"
year            =   (FWS 4*DIGIT FWS) / obs-year
time            =   time-of-day zone
time-of-day     =   hour ":" minute [ ":" second ]
hour            =   2DIGIT / obs-hour
minute          =   2DIGIT / obs-minute
second          =   2DIGIT / obs-second
zone            =   (FWS ( "+" / "-" ) 4DIGIT) / obs-zone
```

**references**

1. [^wikipedia]:https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form

2. [^geeksforgeeks]:https://www.geeksforgeeks.org/bnf-notation-in-compiler-design/

3. [^EBNF]:https://matt.might.net/articles/grammars-bnf-ebnf/

4. [^EBNF]:https://tomassetti.me/ebnf/

5. [^instaparse]:https://github.com/Engelberg/instaparse/blob/master/docs/ABNF.md