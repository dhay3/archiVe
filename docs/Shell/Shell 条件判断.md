# Shell 条件判断

## if

> `[]`两旁必须要有空格，可以使用逻辑运算符，如果是命令无需添加`[]`
>
> ==`if`后面可以跟任意数量的命令。这时，所有命令都会执行，但是判断真伪只看最后一个命令，即使前面所有命令都失败，只要最后一个命令返回`0`，就会执行`then`的部分。==
>
> ```
> [root@cyberpelican opt]# if false; true; then echo 'hello world'; fi
> hello world
> ```

pattern：

```shell
if [ condition ]; then
  commands
elif [ codition ]; then
  commands...
else
  commands
fi
```

==condition两旁的中括号可以省略，如果if与then不在同一行，分号可以省略，fi也可以添加分号在同一行`if [codition];then commands;fi`==

```shell
if [condition] 
then
  commands
elif [codition]
then
  commands...
else
  commands
fi
```

## case

`case`结构用于多值判断，可以为每个值指定对应的命令，跟包含多个`elif`的`if`结构等价，但是语义更好。

```shell
case expression in
  pattern )
    commands ;;
  pattern )
    commands ;;
  ...
esac
```

上面代码中，`expression`是一个表达式，`pattern`是表达式的值或者一个模式，可以有多条，用来匹配多个值，每条以两个分号（`;`）结尾(break)。

> 可以使用wildcard expasion 和 globbing

### 匹配模式

- `a)`：匹配`a`。
- `a|b)`：匹配`a`或`b`。
- `[[:alpha:]])`：匹配单个字母。
- `???)`：匹配3个字符的单词。
- `*.txt)`：匹配`.txt`结尾。
- `*)`：匹配任意输入，通过作为`case`结构的最后一个模式。

### example1

```shell
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
read -p "Enter a number between 1 to 3 > " num
case $num in
1) echo 1
;;
2) echo 2
;;
3) echo 3
;;
*) echo "not a number"
;;
esac 


[root@cyberpelican opt]# ./test.sh 
Enter a number between 1 to 3 > 2
2
```

### example2

```shell
#!/bin/bash

OS=$(uname -s)

case "$OS" in
  FreeBSD) echo "This is FreeBSD" ;;
  Darwin) echo "This is Mac OSX" ;;
  AIX) echo "This is AIX" ;;
  Minix) echo "This is Minix" ;;
  Linux) echo "This is Linux" ;;
  *) echo "Failed to identify this OS" ;;
esac
```

### example3

Bash 4.0之前，`case`结构只能匹配一个条件，然后就会退出`case`结构。Bash 4.0之后，允许匹配多个条件，这时可以用`;;&`终止每个条件块。

```shell
#!/bin/bash
# test.sh

read -n 1 -p "Type a character > "
echo
case $REPLY in
  [[:upper:]])    echo "'$REPLY' is upper case." ;;&
  [[:lower:]])    echo "'$REPLY' is lower case." ;;&
  [[:alpha:]])    echo "'$REPLY' is alphabetic." ;;&
  [[:digit:]])    echo "'$REPLY' is a digit." ;;&
  [[:graph:]])    echo "'$REPLY' is a visible character." ;;&
  [[:punct:]])    echo "'$REPLY' is a punctuation symbol." ;;&
  [[:space:]])    echo "'$REPLY' is a whitespace character." ;;&
  [[:xdigit:]])   echo "'$REPLY' is a hexadecimal digit." ;;&
esac

$ test.sh
Type a character > a
'a' is lower case.
'a' is alphabetic.
'a' is a visible character.
'a' is a hexadecimal digit.
```

### example4

```sh
    while [ $# -gt 0 ]
    do
        key="$1"

        case $key in
            --fixup)
                COMMIT="$2"
                shift # past argument
                shift # past value
                ;;
            --amend)
                AMEND=1
                shift # past argument
                ;;
            -f|--force)
                FORCE=1
                shift # past argument
                ;;
            -h|--help)
                echo "$USAGE"
                return 0
                ;;
            *)
                MESSAGE="$1"
                shift # past argument
                ;;
        esac
    done
```

