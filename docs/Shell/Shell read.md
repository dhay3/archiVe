# Shell read

参考：

https://www.runoob.com/linux/linux-comm-read.html

https://wangdoc.com/bash/read.html

## 概述

pattern：`read [options] [variable...]`

read命令将用户的输入存入变量，用户按下回车键，就表示输入结束。

variable 用来保存输入数值的一个或多个变量名。==如果没有提供变量名，环境变量REPLY会包含用户输入的一整行数据。==

> 如果用户输入项小于read命令给出的变量数，那么额外的变量值为空。
>
> 如果用户输入项多于定义的变量，那么多余的输入项都会包含到最后一个变量中

### 0x001

默认变量REPLY

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
echo -n "enter your firstname and lastname > "
read 
echo "Hi! $REPLY"
[root@cyberpelican opt]# ./test.sh 
enter your firstname and lastname > henery terry
Hi! henery terry
```

### 0x002

单一变量，会将所有的输入最为变量

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
echo -n "put some words > "
read text
echo "out put > $text"
[root@cyberpelican opt]# ./test.sh 
put some words > a b c
out put > a b c
```

### 0x003

输入项 大于 变量数，可以做为类似map结合eval

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
echo -n "enter your firstname and lastname > "
read FN LN
echo "firstname = $FN"
echo "lastname = $LN"
[root@cyberpelican opt]# ./test.sh 
enter your firstname and lastname > ted teddy woofy mickey
firstname = ted
lastname = teddy woofy mickey
```

### 0x004

输入项 小于 变量数

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
echo -n "enter your firstname and lastname > "
read FN LN
echo "firstname = $FN"
echo "lastname = $LN"
[root@cyberpelican opt]# ./test.sh 
enter your firstname and lastname > ted
firstname = ted
lastname = 
```

## 读取文件

> read读取文件每次读取一行

read还可以用于读取文件

```
[root@cyberpelican opt]# cat test.sh 
filename='/etc/hosts'
cat $filename | while read line
do
	echo $line
done
[root@cyberpelican opt]# ./test.sh 
127.0.0.1 localhost localhost.localdomain localhost4 localhost4.localdomain4
::1 localhost localhost.localdomain localhost6 localhost6.localdomain6
```

同样可以使用重定向符

```
[root@cyberpelican opt]# cat test.sh 
filename='/etc/hosts'
while read line
do
	echo $line
done < $filename
```

使用eval可以作为map处理

```
[root@localhost ~]# cat file
NAME ZHANG
AGE 20
SEX NV
[root@localhost ~]# cat test.sh
#!/bin/bash

while read KEY VALUE
do
    eval "${KEY}=${VALUE}"
done < file
echo "$NAME $AGE $SEX"
[root@localhost ~]# ./test.sh
ZHANG 20 NV
```



## 参数

- `-t`

  设置超时秒数，如果超过了指定时间，用户仍然没有输入，脚本将放弃等待，继续向下执行。

  ```
  [root@cyberpelican opt]# cat test.sh 
  echo -n "输入一些文本 > "
  if read -t 3 response; then
    echo "用户已经输入了"
  else
    echo "用户没有输入"
  fi
  [root@cyberpelican opt]# ./test.sh 
  输入一些文本 > 用户没有输入
  ```

  环境变量`TMOUT`也起同样的作用

  ```
  [root@cyberpelican opt]# cat test.sh 
  TMOUT=3
  echo -n "输入一些文本 > "
  if read response; then
    echo "用户已经输入了"
  else
    echo "用户没有输入"
  fi
  ```

- `-P`

  > 注意单引号在双引号内同样失去特殊含义

  指定用户输入的提示信息

  ```
  [root@cyberpelican opt]# cat test.sh 
  read -p "Enter one or more values > "
  echo "REPLY = '$REPLY'"
  [root@cyberpelican opt]# ./test.sh 
  Enter one or more values > ok
  REPLY = 'ok'
  ```

- `-a`

  > 注意如果读取数值中的值需要使用`${}`，不能使用缩写

  将用户的输入的值赋值给一个数组，下标从0开始

  ```
  [root@cyberpelican opt]# read -a array
  jinger mitty mewo
  [root@cyberpelican opt]# echo $array[2]
  jinger[2]
  [root@cyberpelican opt]# echo ${array[2]}
  mewo
  [root@cyberpelican opt]# 
  ```

- `-n`

  读取指定字符数做为变量，输出不会换行

  ```
  hel[root@cyberpelican opt]# read -n 5 letter
  [root@cyberpelican opt]# 
  ```

- `-s`

  静默模式，不显示输入的内容在控制台中，常用于密码

  ```
  [root@cyberpelican opt]# read -s passwd
  [root@cyberpelican opt]# echo $passwd
  password
  ```

  
