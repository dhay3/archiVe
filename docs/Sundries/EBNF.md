# EBNF

参考：

https://matt.might.net/articles/grammars-bnf-ebnf/

https://tomassetti.me/ebnf/

EBNF(Extended Backups-Naur Form)是一个用于定义语言的语言(sudoers配置文件就是使用EBNF来定义的)，在BNF的基础上扩展了。EBNF由terminal symbols和non-terminal production rules组成



## terminal symbols

terminals也被称为tokens，是EBNF中最小的单元，==可以理解为二叉树的叶子节点==

terminals通常由如下表格中的类型组成，也可以是a quoted literal(表示字面上的意思)，a regular expression，keyword(描述代码时语言的关键字)。表格中的Definition是正则

| **Category**          | **Terminal**                        | **Definition**                                               | **Note**                                                     |
| --------------------- | ----------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| *Identifier*          | Java identifier                     | `/[a-zA-Z$_][a-zA-Z0-9$_]*/`                                 | This is a simplification because it is technically possible to use also UTF escape sequences in Java identifiers |
|                       | Haskell type identifier             | `/[A-Z][a-zA-Z0-9$_’]*/`                                     |                                                              |
|                       | Haskell value identifier            | `/[a-z][a-zA-Z0-9_’]*/`                                      |                                                              |
| *Keyword*             | Some Java keywords                  | “abstract”, “assert”, “boolean”, “break”, “byte”, “case”, “catch”, “char”, “class”, “const”… | “const” is a keyword in Java even if it is not used by any construct. It is a reserved keyword for future usages (same is true for “goto”) |
|                       | Some Python 3 keywords/td>          | “def”, “return”, “raise”, “from”, “import”, “as”, “global”, “nonlocal”, “assert”… |                                                              |
| *Literal*             | Java string literal                 | `/'”‘ (~[“] | ” [btnfr”‘])* ‘”‘/`                            | This is a simplification. We are ignoring the octal and unicode escape sequences |
|                       | Java character literal              | `/”’ (~[‘] |” [btnfr”‘]) ”’`                                 |                                                              |
|                       | Java integer literal                | `/[“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?/`                    | A Java integer literal can actually be expressed in decimal, hexadecimal, octal or binary format. We are just considering the decimal format here. This is true also for Java long literals |
|                       | Java long literal                   | `/[“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?(‘l’|’L’)`            |                                                              |
|                       | Java float literal                  | `/[“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?’.'([“0”-“9”](([“0”-“9″,”_”])*[“0”-“9”])?)?(‘f’|’F’)` | A Java float literal can actually be expressed in decimal or hexadecimal format. We are just considering the decimal format here. We are also ignoring the possibility of specifying the exponent |
|                       | Java boolean literal                | /”true”\|”false”/                                            |                                                              |
| *Separator/Delimiter* | Some Java separators and delimiters | “(“, “)”, “{“, “}”, “,”, “;”…                                |                                                              |
|                       | Some Ruby separators and delimiters | “,”, “;”…                                                    |                                                              |
| *Whitespace*          | Java whitespace                     | /[ trnu000C]+/                                               |                                                              |
|                       | Ruby whitespace                     | /(‘ ‘\|’t’)+/                                                |                                                              |
| *Comment*             | Java line comment                   | /’//’ ~[rn]*/                                                |                                                              |
|                       | Java block comment                  | /’/*’ .*? ‘*/’/                                              |                                                              |
|                       | Python line comment                 | /’#’ ~[rnf]*/                                                |                                                              |

## non-terminal 

由terminal symbols和其他的non-terminal组成，==可以理解为二叉树的父节点==

![2021-08-12_12-22](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-08-12_12-22.4oq5txwzat20.png)

production rules用于定义non-terminal

## 常见符号

definition也可以使用`::=`

|        Usage         | Notation  |
| :------------------: | :-------: |
|      definition      |     =     |
|  concatenation 组合  |     ,     |
|     termination      |     ;     |
|     alternation      |    \|     |
|       optional       |  [ ... ]  |
| repetition 0次或多次 |  { ... }  |
|       grouping       |  ( ... )  |
|   terminal string    |  " ... "  |
|   terminal string    |  ' ... '  |
|       comment        | (* ... *) |
|   special sequence   |  ? ... ?  |
|      exception       |     -     |

## 例子

### 0x001

```enbf
digit with zero ::= "0" | "1";
digit :: = "2" | digit with zero;
numbers ::= {digit};
```

### 0x002

```
aa = "A";
bb = 3 * aa, "B";
cc = 3 * [aa], "C";
dd = {aa}, "D";
ee = aa, {aa}, "E";
ff = 3 * aa, 3 * [aa], "F";
gg = {3 * aa}, "G";
```

可以推出

```
aa: A
bb: AAAB
cc: C AC AAC AAAC
dd: D AD AAD AAAD AAAAD etc.
ee: AE AAE AAAE AAAAE AAAAAE etc.
ff: AAAF AAAAF AAAAAF AAAAAAF
gg: G AAAG AAAAAAG etc.
```

### 0x003

用于描述脚本

```
who ::= "alice";
say ::= func "say(who string)";
do ::= who,say;
```

