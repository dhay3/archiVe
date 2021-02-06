# Shell 循环

## while

> `while`的条件部分可以执行任意数量的命令，但是执行结果的真伪只看最后一个命令的执行结果。
>
> 如果无需使用`test`命令，无需使用`[]`

pattern：

```shell
while [ condition ]; do
  commands
done
```

如果`do`与`while`不在同一行就不需要分号；也可以在同一行，使用分号分隔；如果`condition`是命令无需添加`[]`

### example1

```shell
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash

number=0
while [ "$number" -lt 10 ]; do
  echo "Number = $number"
  ((number++))
done
[root@cyberpelican opt]# ./test.sh 
Number = 0
Number = 1
Number = 2
Number = 3
Number = 4
Number = 5
Number = 6
Number = 7
Number = 8
Number = 9
```

### example2

```shell
while echo 'ECHO'; do echo 'Hi, while looping ...'; done
```

## until

==可以统一都使用`while`。==

`until`循环与`while`循环恰好相反，只要不符合判断条件（判断条件失败），就不断循环执行指定的语句。一旦符合判断条件，就退出循环。

## for

> 注意是双括号

pattern：

```
for (( expression1; expression2; expression3 )); do
  commands
done
```

`expression1`用来初始化循环条件，`expression2`用来决定循环结束的条件，`expression3`在每次循环迭代的末尾执行，用于更新值。

注意，循环条件放在双重圆括号之中。另外，圆括号之中使用变量，不必加上美元符号`$`。

### exmaple1

支持i++

```
for (( i=0; i<5; i=i+1 )); do
  echo $i
done
```

## for ... in 

> 可以使用wildcard expasion

pattern：

```shell
for variable in list;do
  commands
done
```

`do`与`for`可以不在同一行。`for`循环会依次从`list`列表中取出一项，作为变量`variable`，然后在循环体中进行处理。

### exmaple1

```
[root@cyberpelican opt]# for i in word1 word2 word3;do echo $i;done
word1
word2
word3
```

### example2

```
[root@cyberpelican opt]# for i in *.sh;do ll $i;done
-rwxr-xr-x. 1 root root 100 Nov 20 11:29 test.sh
```

### exmaple3

使用子命令

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
for i in $(cat /opt/data)
do
echo "$i"
done
[root@cyberpelican opt]# ./test.sh 
word1
word2
word3
```

### example4

省略`list`，默认使用`$@`

```
for filename; do
  echo "$filename"
done

# 等同于

for filename in "$@" ; do
  echo "$filename"
done
```

### example5/遍历文件

```
path=$1
files=$(ls $path)
for filename in $files
do
   echo $filename >> filename.txt
done
```

## break/continue

Shell中同样支持，break/continue

```
#!/bin/bash

for number in 1 2 3 4 5 6
do
  echo "number is $number"
  if [ "$number" = "3" ]; then
    break
  fi
done
```

