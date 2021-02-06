# Shell test

> 如果test 命令执行成功，返回值0，表达式为真
>
> 如果test命令执行失败，返回值1，表达式为假
>
> ==`[`和`]`与内部的表达式之间必须有空格，`test`简写，所以必须要有空格==
>
> 支持逻辑运算符，`&&`，`||`，`!` 
>
> `if [] && [];then`

`if`结构的判断条件，一般使用`test`命令，有三种形式。

1. `test expression`
2. `[ expression ]`
3. `[[ expression ]]`

获取变量(字符串)长度

```
[root@cyberpelican opt]# echo ${#a}
0
```

## 字符串判断

> ==如果使用数学符号必须在符号两旁添加空格，因为使用例如`-eq`参数时需要空格==

- `[ string ]`：如果`string`不为空（长度大于0），则判断为真，返回0。==替代`-n`==

  ```
  [root@cyberpelican opt]# if [ $a ];then echo "ok";fi
  [root@cyberpelican opt]# 
  [root@cyberpelican opt]# [ $a ]
  [root@cyberpelican opt]# echo $?
  1
  ```

  ==等价于`if [ $a ];then;fi`==

- `[ -n string ]`：如果字符串`string`的长度大于零，则判断为真。

  ```
  [root@cyberpelican opt]# if [ -n "" ];then echo "ok";else echo "zero";fi
  zero
  ```
  
  ==不能以变量的形式==
  
  ```
  [root@cyberpelican opt]# a=""
  [root@cyberpelican opt]# if [ -n $a ];then echo "ok";fi
  ok
  ```
  
- `[ -z string]`：如果字符串`string`的长度为零，则判断为真。

  ```
  [root@cyberpelican opt]# if [ -z $a ];then echo "ok";else echo "not";fi
  ok
  ```

- `[ string1 = string2 ]`：如果`string1`和`string2`相同，则判断为真。==等号两旁必须要有空格==

  等价于`[ string1 == string2 ]`

  ```
  [root@cyberpelican opt]# b=123
  [root@cyberpelican opt]# if [ $b = 123 ];then echo "ok";fi
  ok
  ```

- `[ string1 != string2 ]`：如果`string1`和`string2`不相同，则判断为真。

  ```
  [root@cyberpelican opt]# cat test.sh 
  #!/bin/bash
  c="123"
  d="123"
  if [ $c != $d ]
  then
  	echo "equal"
  else
  	echo "not equal"
  fi
  [root@cyberpelican opt]# ./test.sh 
  not equal
  ```

> 注意，`test`命令内部的`>`和`<`，必须用引号引起来（或者是用反斜杠转义）。否则，它们会被 shell 解释为重定向操作符。

- `[ string1 '>' string2 ]`：如果按照字典顺序`string1`排列在`string2`之后，则判断为真。

  ```
  [root@cyberpelican opt]# cat test.sh 
  #!/bin/bash
  c="231"
  d="123"
  if [ $c '>' $d ]
  then
  	echo "equal"
  else
  	echo "not equal"
  fi
  [root@cyberpelican opt]# ./test.sh 
  equal
  ```

- `[ string1 '<' string2 ]`：如果按照字典顺序`string1`排列在`string2`之前，则判断为真。

  ```
  [root@cyberpelican opt]# cat test.sh 
  #!/bin/bash
  c="231"
  d="123"
  if [ $c '<' $d ]
  then
  	echo "equal"
  else
  	echo "not equal"
  fi
  [root@cyberpelican opt]# ./test.sh 
  not equal
  ```

### example

> exit 1 表示正常退出

```shell
#!/bin/bash

ANSWER=maybe

if [ -z "$ANSWER" ]; then
  echo "There is no answer." >&2
  exit 1
fi
if [ "$ANSWER" = "yes" ]; then
  echo "The answer is YES."
elif [ "$ANSWER" = "no" ]; then
  echo "The answer is NO."
elif [ "$ANSWER" = "maybe" ]; then
  echo "The answer is MAYBE."
else
  echo "The answer is UNKNOWN."
fi
```

## 整型判断

==Shell中没有实际的整型，而是整型字符串，所以也可以使用上述的数学符号==

- `[ integer1 -eq integer2 ]`：如果`integer1`等于`integer2`，则为`true`。
- `[ integer1 -ne integer2 ]`：如果`integer1`不等于`integer2`，则为`true`。
- `[ integer1 -le integer2 ]`：如果`integer1`小于或等于`integer2`，则为`true`。
- `[ integer1 -lt integer2 ]`：如果`integer1`小于`integer2`，则为`true`。
- `[ integer1 -ge integer2 ]`：如果`integer1`大于或等于`integer2`，则为`true`。
- `[ integer1 -gt integer2 ]`：如果`integer1`大于`integer2`，则为`true`。

### example

```shell
#!/bin/bash

INT=-5

if [ -z "$INT" ]; then
  echo "INT is empty." >&2
  exit 1
fi
if [ $INT -eq 0 ]; then
  echo "INT is zero."
else
  if [ $INT -lt 0 ]; then
    echo "INT is negative."
  else
    echo "INT is positive."
  fi
  if [ $((INT % 2)) -eq 0 ]; then
    echo "INT is even."
  else
    echo "INT is odd."
  fi
fi
```

## 文件判断

```
[root@cyberpelican opt]# ll
total 16
-rw-r--r--. 1 root root 14 Nov 20 09:57 file1
-rw-r--r--. 1 root root 76 Nov 18 10:32 read
-rw-r--r--. 1 root root  6 Nov 18 10:33 redirect
drwxr-xr-x. 2 root root  6 Oct 31  2018 rh
-rwxr-xr-x. 1 root root 67 Nov 20 09:58 test.sh

```

- `[ -f file ]`：如果 file 存在并且是一个普通文件，则为`true`。

  ```
  [root@cyberpelican opt]# if [ -f "/opt/test.sh" ];then echo "test.sh";fi
  test.sh
  ```

- `[ -e file ]`：如果 file 存在，则为`true`。

  ```
  [root@cyberpelican opt]# if [ -e "/opt/test.sh" ];then echo "test.sh";fi
  test.sh
  ```

- `[ -d file ]`：如果 file 存在并且是一个目录，则为`true`。

  ```
  [root@cyberpelican opt]# if [ -d "/opt" ];then echo "/opt";fi
  /opt
  ```

- `[ -r file ]`：如果 file 存在并且可读（当前用户有可读权限），则为`true`。

  ```
  [root@cyberpelican opt]# if [ -r "/opt/test.sh" ];then echo "test.sh";fi
  test.sh
  ```

- `[ -w file ]`：如果 file 存在并且可写（当前用户拥有可写权限），则为`true`。

- `[ -x file ]`：如果 file 存在并且可执行（有效用户有执行／搜索权限），则为`true`。

## 算术判断

> 使用`((...))`，需要注意的是`((!0))`为true，`((0))`为false，与`$?`相反

```shell
#!/bin/bash

INT=-5

if [[ "$INT" =~ ^-?[0-9]+$ ]]; then
  if ((INT == 0)); then
    echo "INT is zero."
  else
    if ((INT < 0)); then
      echo "INT is negative."
    else
      echo "INT is positive."
    fi
    if (( ((INT % 2)) == 0)); then
      echo "INT is even."
    else
      echo "INT is odd."
    fi
  fi
else
  echo "INT is not an integer." >&2
  exit 1
fi
```

## 普通命令与逻辑命令

> 注意如果是普通命令无需`[]`

```shell
[root@cyberpelican opt]# cat test.sh 
filename=$1
word=$2
if grep --color $word $filename
then
	echo "$word is in the $filename"
else
	echo "$word is not in the $filename"
fi
[root@cyberpelican opt]# ./test.sh /opt/test.sh filename
filename=$1
if grep --color $word $filename
	echo "$word is in the $filename"
	echo "$word is not in the $filename"
filename is in the /opt/test.sh

```

### example

```shell
[[ -d "$dir_name" ]] && cd "$dir_name" && rm *

# 等同于

if [[ ! -d "$dir_name" ]]; then
  echo "No such directory: '$dir_name'" >&2
  exit 1
fi
if ! cd "$dir_name"; then
  echo "Cannot cd to '$dir_name'" >&2
  exit 1
fi
if ! rm *; then
  echo "File deletion failed. Check results" >&2
  exit 1
fi
```

