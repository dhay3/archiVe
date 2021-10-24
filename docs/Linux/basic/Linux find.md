# Linux find

## 概述

pattern：`find [starting-point] [expression]`

find命令直接搜索硬盘速度较慢，如果没有指定starting-point默认使用`.`，==默认使用递归查询==

expression通常由global options，tests，actions，position options，oprators一个或多个组成

## 全局条件

全局条件必须在starting-point之后的第一个参数

- -maxdepth

  find默认会遍历当前目录及其子目录，可以通过该参数指定允许遍历的层级，1表示当前目录

  ```
  find . -maxdepth  1 -name "*."
  ```

## 基础条件

- `-name pattern`

  匹配文件名，可以使用shell的模式扩展。==如果没有使用模式扩展表示精确匹配==。只匹配文件名不匹配路径

  ```
  root in /usr/local/\/test1 λ find . -name "*2020-*"
  ./error.xin.2020-12-01.log.gz  
  ```

- `-perm pattern`

  匹配有指定权限的文件

  ```
  [root@VM-0-4-centos opt]# find . -perm -751
  .
  ./rh
  ./mellanox
  ./mellanox/iproute2
  ./mellanox/iproute2/etc
  ./mellanox/iproute2/share
  ./mellanox/iproute2/share/man
  ./mellanox/iproute2/share/man/man7
  ./mellanox/iproute2/share/man/man3
  ```

- `-user username`

  按照用户的名字来查找文件

  ```
  find ~ -user sam
  ```

- `-regex pattern`

  使用正则表达式来匹配文件，如果使用了regex，默认使用EMac regex，如果想要使用posix regex(basic regex)需要使用`-regextyep`指定

  ```
  root in /opt λ find . -regextype posix-egrep  -regex "\./[[:digit:]]{4}-[[:digit:]]{2}.log"
  ./2020-04.log
  ./2021-01.log
  ./2020-02.log
  ./2021-08.log
  ./2021-04.log
  ./2020-05.log
  ./2020-03.log
  ./2021-07.log
  ./2021-10.log
  ./2020-09.log
  ./2021-06.log
  ```

  支持的regextype可以使用`find -regextype help`来查看

  ```
  cpl in ~ λ find -regextype help
  find: Unknown regular expression type ‘help’; valid types are ‘findutils-default’, ‘ed’, ‘emacs’, ‘gnu-awk’, ‘grep’, ‘posix-awk’, ‘awk’, ‘posix-basic’, ‘posix-egrep’, ‘egrep’, ‘posix-extended’, ‘posix-minimal-basic’, ‘sed’.
  ```

  this is a match on the whole path

  如果想要子目录遍历，需要在pattern前加`.*/`表示目录，否则会找不到。如果是想在当前目录下遍历子目录需要在pattern前加`\./`

  ```
  
  ```

  

- `-type c`

  可以一次指定多种类型，`-type f,l,s`
  
  1. b：块文件
  2. c：字符设备文件
  3. d：目录
  4. p：具名管道符
  5. f：普通文件
  6. l：链接文件
  7. s：socket文件

## 时间条件

> `+`表示之前，`-`表示之内

- `-amin n `

  匹配被n分钟之前查看过的文件

- `-atime n`

  匹配被n*24小时之前查看过的文件

- `-cmin n`

  匹配n分钟之前文件状态被修改的文件

- `-ctime n`

  匹配n*24小时之前文件状态被修改的文件

- `-mmin`

  匹配n分钟之前被修改的文件

- `-mtime`

  匹配n*24小时之前被修改的文件

## 动作

- `-delete`

  找到文件后删除

- `-exec`

  https://stackoverflow.com/questions/2961673/find-missing-argument-to-exec

  找到文件后执行，直到遇到第一个`;`终止(也就是说需要在exec末尾添加`\;`)。`{}`内的值会被替换成find找到的值

  ```
  find /tmp/foo -exec echo {} \;
  ```

  如果想要执行shell，可以使用如下方法

  ```
  find /tmp/foo -exec sh -c 'ffmpeg ... && rm'
  ```

- `-ok command`

- `-print | -print0`  

  将匹配后的文件输出并换行，默认动作。print0不换行输出并且会原样输出文件名不转义

- `-printf`

  格式化输出，`%f`表示只输出文件名

  ```
  find . -maxdepth 1 -type d -printf '%f\n'
  ```

## 案例

找到文件后删除

```
find /tmp -name core -type f -print | xargs /bin/rm -f
```

查找隐藏文件，==posix regex中dot并不表示any character==

```
find /tmp -type f -name ".*"
```

找到当前目录下的所以文件并移动

```
find /tmp -maxdepth 1 -type f -exec mv {} dev/ \;
```

