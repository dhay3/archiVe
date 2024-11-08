# Shell 脚本错误处理

> https://wangdoc.com/bash/debug.html#bash%E7%9A%84-x%E5%8F%82%E6%95%B0
>
> 可以使用`bash -x`打印出命令(如果以`/urs/bin/env bash -x`会报错)，用作debug，或是使用`set -e`

## LINENO

变量`LINENO`返回它在脚本里面的行号。

```
[root@cyberpelican opt]# cat test.sh 
# !/bin/bash 

echo "the line of $LINENO"
[root@cyberpelican opt]# ./test.sh 
the line of 3
```

## FUNCNAME

变量`FUNCNAME`返回一个数组，内容是当前的函数调用堆栈。该数组的0号成员是当前调用的函数，1号成员是调用当前函数的函数，以此类推。

```shell
[root@cyberpelican opt]# cat test.sh 
# !/bin/bash 
func1 (){
echo "func1 : FUNCNAME0 is ${FUNCNAME[0]}"
echo "func1 : FUNCNAME1 is ${FUNCNAME[1]}"
echo "func1 : FUNCNAME2 is ${FUNCNAME[2]}"
func2
}

func2 (){
echo "func2 : FUNCNAME0 is ${FUNCNAME[0]}"
echo "func2 : FUNCNAME1 is ${FUNCNAME[1]}"
echo "func2 : FUNCNAME2 is ${FUNCNAME[2]}"
}

func1

[root@cyberpelican opt]# ./test.sh 
func1 : FUNCNAME0 is func1
func1 : FUNCNAME1 is main
func1 : FUNCNAME2 is 
func2 : FUNCNAME0 is func2
func2 : FUNCNAME1 is func1
func2 : FUNCNAME2 is main
```

func1：FUNCNAME[0]是func1，1号调用函数是main，，没有2号调用函数

func2：FUNCNAME[0]是func2，1号调用函数是func1，2号调用函数是main

## BASH_SOURCE

> lib1.sh或lib2.sh都不需要执行权限

变量`BASH_SOURCE`返回一个数组，内容是当前的脚本调用堆栈。该数组的0号成员是当前执行的脚本，1号成员是调用当前脚本的脚本，以此类推

```shell
[root@cyberpelican opt]# cat lib1.sh 
# lib1.sh
function func1()
{
  echo "func1: BASH_SOURCE0 is ${BASH_SOURCE[0]}"
  echo "func1: BASH_SOURCE1 is ${BASH_SOURCE[1]}"
  echo "func1: BASH_SOURCE2 is ${BASH_SOURCE[2]}"
  func2
}
[root@cyberpelican opt]# cat lib2.sh 
# lib2.sh
function func2()
{
  echo "func2: BASH_SOURCE0 is ${BASH_SOURCE[0]}"
  echo "func2: BASH_SOURCE1 is ${BASH_SOURCE[1]}"
  echo "func2: BASH_SOURCE2 is ${BASH_SOURCE[2]}"
}

[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
# main.sh

source lib1.sh
source lib2.sh

func1
[root@cyberpelican opt]# ./test.sh 
func1: BASH_SOURCE0 is lib1.sh
func1: BASH_SOURCE1 is ./test.sh
func1: BASH_SOURCE2 is 
func2: BASH_SOURCE0 is lib2.sh
func2: BASH_SOURCE1 is lib1.sh
func2: BASH_SOURCE2 is ./test.sh

```

## BASH_LINENO

变量`BASH_LINENO`返回一个数组，内容是调用该脚本的脚本中对应的行号

- lib1.sh

  ```
  # lib1.sh
  function func1()
  {
    echo "func1: BASH_LINENO is ${BASH_LINENO[0]}"
    echo "func1: FUNCNAME is ${FUNCNAME[0]}"
    echo "func1: BASH_SOURCE is ${BASH_SOURCE[1]}"
  
    func2
  }
  ```

- lib2.sh

  ```
  # lib2.sh
  function func2()
  {
    echo "func2: BASH_LINENO is ${BASH_LINENO[0]}"
    echo "func2: FUNCNAME is ${FUNCNAME[0]}"
    echo "func2: BASH_SOURCE is ${BASH_SOURCE[1]}"
  }
  ```

- main.sh

  ```
  #!/bin/bash
  # main.sh
  
  source lib1.sh
  source lib2.sh
  
  func1
  ```

- 结果

  ```
  $ ./main.sh
  func1: BASH_LINENO is 7
  func1: FUNCNAME is func1
  func1: BASH_SOURCE is main.sh
  func2: BASH_LINENO is 8
  func2: FUNCNAME is func2
  func2: BASH_SOURCE is lib1.sh
  ```

  上面例子中，函数`func1`是在`main.sh`的第7行调用，函数`func2`是在`lib1.sh`的第8行调用的。

## 0x001

如果脚本里面有运行失败的命令（返回值非`0`），==Bash 默认会继续执行后面的命令。==

```
[root@cyberpelican opt]# cat test.sh 
foo
echo bar

[root@cyberpelican opt]# ./test.sh 
./test.sh: line 1: foo: command not found
bar
```

上面脚本中，`foo`是一个不存在的命令，执行时会报错。==但是，Bash 会忽略这个错误，继续往下执行。==

这种行为很不利于脚本安全和除错。实际开发中，如果某个命令失败，往往需要脚本停止执行，防止错误累积。这时，一般采用下面的写法。

1. `command || { echo "command failed"; exit 1; }`

   这里的`{}`与模式扩展没有关系

2. `if ! command; then echo "command failed"; exit 1; fi`

3. `command
   if [ "$?" -ne 0 ]; then echo "command failed"; exit 1; fi`

## 0x002

```
#! /bin/bash

dir_name=/path/not/exist

cd $dir_name
rm *
```

1. 如果`dir_name`不存在，`cd $dir_name`执行失败，整条语句执行失败
2. 如果`dir_name`为空，`cd`就会进入用户主目录，从而删光用户主目录的文件。

所以要在执行之前判断目录是否存在，然后执行其他操作

```
[[ -d $dir_name ]] && cd $dir_name && rm *
```

如果放心删除什么文件，可以先打印出来看一下

```
[[ -d $dir_name ]] && cd $dir_name && echo rm *
```

