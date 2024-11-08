# Regular Expression

## 0x01 Overview

Regular Expression 正则表达式，通常也被简写为 regex 或者 regexp，通过一系列特殊含义的字符来表示匹配的字符

按照 syntaxes 可以大概分为 2 大类

1. POSIX Regular Expression

   可以细分为 3 种

   1. BRE(Basic Regular Expressions)
   2. ERE(Extended Regular Expressions)
   3. SRE(Simple Regular Expressions) 已废弃

   Unix

2. PCRE(Perl Compatiable Regular Expression)

   比 POSIX Regular Expression 更强大且易懂

## 0x02 POSIX Regular Expression

### BRE

BRE (Basic Regular Expression)

| Metacharacter | Description                                                  |
| ------------- | ------------------------------------------------------------ |
| `^`           | Matches the starting position within the string. In line-based tools, it matches the starting position of any line. |
| `.`           | Matches any single character (many applications exclude newlines(/n), and exactly which characters are considered newlines is flavor-,  character-encoding-, and platform-specific, but it is safe to assume  that the line feed character is included). Within POSIX bracket  expressions, the dot character matches a literal dot. For example, `a.c` matches "abc", etc., but `[a.c]` matches only "a", ".", or "c". |
| `[ ]`         | A bracket expression. Matches a single character that is contained within the brackets. For example, `[abc]` matches "a", "b", or "c". `[a-z]` specifies a range which matches any lowercase letter from "a" to "z".(applyd to digits`[1-9]` as well) These forms can be mixed: `[abcx-z]` matches "a", "b", "c", "x", "y", or "z", as does `[a-cx-z]`. The `-` character is treated as a literal character if it is the last or the first (after the `^`, if present) character within the brackets: `[abc-]`, `[-abc]`. Backslash escapes are not allowed. The `]` character can be included in a bracket expression if it is the first (after the `^`) character: `[]abc]`. |
| `[^ ]`        | Matches a single character that is not contained within the brackets. For example, `[^abc]` matches any character other than "a", "b", or "c". `[^a-z]` matches any single character that is not a lowercase letter from "a" to "z". Likewise, literal characters and ranges can be mixed. |
| `$`           | Matches the ending position of the string or the position just  before a string-ending newline. In line-based tools, it matches the  ending position of any line. |
| `( )`         | Defines a marked subexpression. The string matched within the parentheses can be recalled later (see the next entry, `\n`). A marked subexpression is also called a block or capturing group. *BRE mode requires `\( \)`.* |
| `\n`          | Matches what the *n*th marked subexpression matched, where *n* is a digit from 1 to 9. This construct is defined in the POSIX standard. Some tools allow referencing more than nine capturing groups. Also  known as a back-reference, this feature is supported in BRE mode. |
| `*`           | Matches the preceding element zero or more times. For example, `ab*c` matches "ac", "abc", "abbbc", etc. `[xyz]*` matches "", "x", "y", "z", "zx", "zyx", "xyzzy", and so on. `(ab)*` matches "", "ab", "abab", "ababab", and so on. |
| `{m,n}`       | Matches the preceding element at least *m* and not more than *n* times. For example, `a{3,5}` matches only "aaa", "aaaa", and "aaaaa". This is not found in a few older instances of regexes. BRE mode requires `\{m,n\}`. |

### ERE

*ERE adds `?`, `+`, and `|`, and it removes the need to escape the metacharacters `( )` and `{ }`, which are required in BRE.* 

ERE (Extended Regular Expression)是 BRE 的扩展，在 BRE 的基础上添加了 `?`, `+` ,`|` 

且无需转译 `()`, `{}` ，即在 BRE 中的 `\{\}` 等价于 ERE 中的 `{}`，在 ERE 中 `\{\}` 表是字面意义

| Metacharacter | Description                                                  |
| ------------- | ------------------------------------------------------------ |
| `?`           | Matches the preceding element zero or one time. For example, `ab?c` matches only "ac" or "abc". |
| `+`           | Matches the preceding element one or more times. For example, `ab+c` matches "abc", "abbc", "abbbc", and so on, but not "ac". |
| `|`           | The choice (also known as alternation or set union) operator matches either the expression before or the expression after the operator. For  example, `abc|def` matches "abc" or "def". |

### Character Class

BRE 和 ERE 在除上部分谈到的规则外，还支持 character class(a sequence of characters)

只能被用在 bracket expression `[]`中，例如 `[[:digit:]]`

对比 PCRE

| Description                                | POSIX        | Perl/Tcl | Vim     | Java                | ASCII                                      |
| ------------------------------------------ | ------------ | -------- | ------- | ------------------- | ------------------------------------------ |
| ASCII characters                           |              |          |         | `\p{ASCII}`         | `[\x00-\x7F]`                              |
| Alphanumeric characters                    | `[:alnum:]`  |          |         | `\p{Alnum}`         | `[A-Za-z0-9]`                              |
| Alphanumeric characters plus "_"           |              | `\w`     | `\w`    | `\w`                | `[A-Za-z0-9_]`                             |
| Non-word characters                        |              | `\W`     | `\W`    | `\W`                | `[^A-Za-z0-9_]`                            |
| Alphabetic characters                      | `[:alpha:]`  |          | `\a`    | `\p{Alpha}`         | `[A-Za-z]`                                 |
| Space and tab                              | `[:blank:]`  |          | `\s`    | `\p{Blank}`         | `[ \t]`                                    |
| Word boundaries                            |              | `\b`     | `\< \>` | `\b`                | `(?<=\W)(?=\w)|(?<=\w)(?=\W)`              |
| Non-word boundaries                        |              |          |         | `\B`                | `(?<=\W)(?=\W)|(?<=\w)(?=\w)`              |
| Control characters                         | `[:cntrl:]`  |          |         | `\p{Cntrl}`         | `[\x00-\x1F\x7F]`                          |
| Digits                                     | `[:digit:]`  | `\d`     | `\d`    | `\p{Digit}` or `\d` | `[0-9]`                                    |
| Non-digits                                 |              | `\D`     | `\D`    | `\D`                | `[^0-9]`                                   |
| Visible characters                         | `[:graph:]`  |          |         | `\p{Graph}`         | `[\x21-\x7E]`                              |
| Lowercase letters                          | `[:lower:]`  |          | `\l`    | `\p{Lower}`         | `[a-z]`                                    |
| Visible characters and the space character | `[:print:]`  |          | `\p`    | `\p{Print}`         | `[\x20-\x7E]`                              |
| Punctuation characters                     | `[:punct:]`  |          |         | `\p{Punct}`         | <code>[][!"#$%&'()*+,./:;<=>?@\^_`{</code> |
| Whitespace characters                      | `[:space:]`  | `\s`     | `\_s`   | `\p{Space}` or `\s` | `[ \t\r\n\v\f]`                            |
| Non-whitespace characters                  |              | `\S`     | `\S`    | `\S`                | `[^ \t\r\n\v\f]`                           |
| Uppercase letters                          | `[:upper:]`  |          | `\u`    | `\p{Upper}`         | `[A-Z]`                                    |
| Hexadecimal digits                         | `[:xdigit:]` |          | `\x`    | `\p{XDigit}`        | `[A-Fa-f0-9]`                              |

## 0x03 PCRE

PCRE 是另外一种格式，对比 POSIX Regular Expression 更加强大且易懂

很多的高级编程语言或者应用都采用了 PCRE 或者类似的语法，例如 Java,Python,JavaScript,Ruby,Nginx 等等

| Meta­character(s) | Description                                                  | Example                                                      |
| ----------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `.`               | Normally matches any character except a newline.  Within square brackets the dot is literal. | `$string1 = "Hello World\n"; if ($string1 =~ m/...../) {  print "$string1 has length >= 5.\n"; } ` **Output:** `Hello World has length >= 5. ` |
| `( )`             | Groups a series of pattern elements to a single element.  When you match a pattern within parentheses, you can use any of `$1`, `$2`, ... later to refer to the previously matched pattern. Some implementations may use a backslash notation instead, like `\1`, `\2`. | `$string1 = "Hello World\n"; if ($string1 =~ m/(H..).(o..)/) {  print "We matched '$1' and '$2'.\n"; } ` **Output:** `We matched 'Hel' and 'o W'. ` |
| `+`               | Matches the preceding pattern element one or more times.     | `$string1 = "Hello World\n"; if ($string1 =~ m/l+/) {  print "There are one or more consecutive letter \"l\"'s in $string1.\n"; } ` **Output:** `There are one or more consecutive letter "l"'s in Hello World. ` |
| `?`               | Matches the preceding pattern element zero or one time.      | `$string1 = "Hello World\n"; if ($string1 =~ m/H.?e/) {  print "There is an 'H' and a 'e' separated by ";  print "0-1 characters (e.g., He Hue Hee).\n"; } ` **Output:** `There is an 'H' and a 'e' separated by 0-1 characters (e.g., He Hue Hee). ` |
| `?`               | Modifies the `*`, `+`, `?` or `{M,N}`'d regex that comes before to match as few times as possible. | `$string1 = "Hello World\n"; if ($string1 =~ m/(l.+?o)/) {  print "The non-greedy match with 'l' followed by one or ";  print "more characters is 'llo' rather than 'llo Wo'.\n"; } ` **Output:** `The non-greedy match with 'l' followed by one or more characters is 'llo' rather than 'llo Wo'. ` |
| `*`               | Matches the preceding pattern element zero or more times.    | `$string1 = "Hello World\n"; if ($string1 =~ m/el*o/) {  print "There is an 'e' followed by zero to many ";  print "'l' followed by 'o' (e.g., eo, elo, ello, elllo).\n"; } ` **Output:** `There is an 'e' followed by zero to many 'l' followed by 'o' (e.g., eo, elo, ello, elllo). ` |
| `{M,N}`           | Denotes the minimum M and the maximum N match count. N can be omitted and M can be 0: `{M}` matches "exactly" M times; `{M,}` matches "at least" M times; `{0,N}` matches "at most" N times. `x* y+ z?` is thus equivalent to `x{0,} y{1,} z{0,1}`. | `$string1 = "Hello World\n"; if ($string1 =~ m/l{1,2}/) {  print "There exists a substring with at least 1 ";  print "and at most 2 l's in $string1\n"; } ` **Output:** `There exists a substring with at least 1 and at most 2 l's in Hello World ` |
| `[…]`             | Denotes a set of possible character matches.                 | `$string1 = "Hello World\n"; if ($string1 =~ m/[aeiou]+/) {  print "$string1 contains one or more vowels.\n"; } ` **Output:** `Hello World contains one or more vowels. ` |
| `|`               | Separates alternate possibilities.                           | `$string1 = "Hello World\n"; if ($string1 =~ m/(Hello|Hi|Pogo)/) {  print "$string1 contains at least one of Hello, Hi, or Pogo."; } ` **Output:** `Hello World contains at least one of Hello, Hi, or Pogo. ` |
| `\b`              | Matches a zero-width boundary between a word-class character (see  next) and either a non-word class character or an edge; same as `(^\w|\w$|\W\w|\w\W)`. | `$string1 = "Hello World\n"; if ($string1 =~ m/llo\b/) {  print "There is a word that ends with 'llo'.\n"; } ` **Output:** `There is a word that ends with 'llo'. ` |
| `\w`              | Matches an alphanumeric character, including "_";  same as `[A-Za-z0-9_]` in ASCII, and `[\p{Alphabetic}\p{GC=Mark}\p{GC=Decimal_Number}\p{GC=Connector_Punctuation}]` in Unicode, where the `Alphabetic` property contains more than Latin letters, and the `Decimal_Number` property contains more than Arab digits. | `$string1 = "Hello World\n"; if ($string1 =~ m/\w/) {  print "There is at least one alphanumeric ";  print "character in $string1 (A-Z, a-z, 0-9, _).\n"; } ` **Output:** `There is at least one alphanumeric character in Hello World (A-Z, a-z, 0-9, _). ` |
| `\W`              | Matches a *non*-alphanumeric character, excluding "_";  same as `[^A-Za-z0-9_]` in ASCII, and `[^\p{Alphabetic}\p{GC=Mark}\p{GC=Decimal_Number}\p{GC=Connector_Punctuation}]` in Unicode. | `$string1 = "Hello World\n"; if ($string1 =~ m/\W/) {  print "The space between Hello and ";  print "World is not alphanumeric.\n"; } ` **Output:** `The space between Hello and World is not alphanumeric. ` |
| `\s`              | Matches a whitespace character,  which in ASCII are tab, line feed, form feed, carriage return, and space;  in Unicode, also matches no-break spaces, next line, and the variable-width spaces (amongst others). | `$string1 = "Hello World\n"; if ($string1 =~ m/\s.*\s/) {  print "In $string1 there are TWO whitespace characters, which may";  print " be separated by other characters.\n"; } ` **Output:** `In Hello World there are TWO whitespace characters, which may be separated by other characters. ` |
| `\S`              | Matches anything *but* a whitespace.                         | `$string1 = "Hello World\n"; if ($string1 =~ m/\S.*\S/) {  print "In $string1 there are TWO non-whitespace characters, which";  print " may be separated by other characters.\n"; } ` **Output:** `In Hello World there are TWO non-whitespace characters, which may be separated by other characters. ` |
| `\d`              | Matches a digit;  same as `[0-9]` in ASCII;  in Unicode, same as the `\p{Digit}` or `\p{GC=Decimal_Number}` property, which itself the same as the `\p{Numeric_Type=Decimal}` property. | `$string1 = "99 bottles of beer on the wall."; if ($string1 =~ m/(\d+)/) {  print "$1 is the first number in '$string1'\n"; } ` **Output:** `99 is the first number in '99 bottles of beer on the wall.' ` |
| `\D`              | Matches a non-digit;  same as `[^0-9]` in ASCII or `\P{Digit}` in Unicode. | `$string1 = "Hello World\n"; if ($string1 =~ m/\D/) {  print "There is at least one character in $string1";  print " that is not a digit.\n"; } ` **Output:** `There is at least one character in Hello World that is not a digit. ` |
| `^`               | Matches the beginning of a line or string.                   | `$string1 = "Hello World\n"; if ($string1 =~ m/^He/) {  print "$string1 starts with the characters 'He'.\n"; } ` **Output:** `Hello World starts with the characters 'He'. ` |
| `$`               | Matches the end of a line or string.                         | `$string1 = "Hello World\n"; if ($string1 =~ m/rld$/) {  print "$string1 is a line or string ";  print "that ends with 'rld'.\n"; } ` **Output:** `Hello World is a line or string that ends with 'rld'. ` |
| `\A`              | Matches the beginning of a string (but not an internal line). | `$string1 = "Hello\nWorld\n"; if ($string1 =~ m/\AH/) {  print "$string1 is a string ";  print "that starts with 'H'.\n"; } ` **Output:** `Hello World is a string that starts with 'H'. ` |
| `\z`              | Matches the end of a string (but not an internal line).[[72\]](https://en.wikipedia.org/wiki/Regular_expression#cite_note-Perl_Best_Practices-72) | `$string1 = "Hello\nWorld\n"; if ($string1 =~ m/d\n\z/) {  print "$string1 is a string ";  print "that ends with 'd\\n'.\n"; } ` **Output:** `Hello World is a string that ends with 'd\n'. ` |
| `[^…]`            | Matches every character except the ones inside brackets.     | `$string1 = "Hello World\n"; if ($string1 =~ m/[^abc]/) { print "$string1 contains a character other than "; print "a, b, and c.\n"; } ` **Output:** `Hello World contains a character other than a, b, and c. ` |

例如：

```
#身份证 PCRE
[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[\dXx]

区号[1-9]\d{5}
年(18|19|20)\d{2}
月((0[1-9])|(1[0-2]))
日(([0-2][1-9])|10|20|30|31) 
校验码\d{3}[\dXx]  //[\dXx]表示从该范围中取任意一字符

#超链接
http(s)?://\w+\.\w+(\.\w+)?
```

**references获取**

[^1]:https://en.wikipedia.org/wiki/Regular_expression