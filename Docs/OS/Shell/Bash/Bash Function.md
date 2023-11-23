# 3.3 - Bash Function

Bash 中的 functions 和其他编程语言中的 functions 或者是 methods 逻辑一样，都是用于将某一功能提取出来以便重复利用

> *Shell functions are executed in the current shell context; no new process is created to interpret them*

## Overview

functions 可以通过如下两种方式来声明

```
fname () compound-command [ redirections ]
function fname [()] compound-command [ redirections ]
```

声明 function 时可以不使用 `function` keyword，如果声明了 `function` keyword 可以不使用 `()`,例如下面两个 function 等价

```
#fname 与 compound-command 中间必须要有一个空格
function test1 {
  echo 1
}

test2 (){
  echo 1
}

test1
test2 
```

compound-command 表示 the body of the function, 通常以 `{ commands }` 这种形式展示，但是也可以使用任何一种 compound commands

例如

```
function test1  while true; do echo 1; done
```

为了统一标准应该使用如下这种格式来定义函数

```
function fname(){
	
}
```

**references**

1. [^1]: https://www.gnu.org/software/bash/manual/bash.html#Shell-Functions

2. [^2]:https://www.gnu.org/software/bash/manual/bash.html#Compound-Commands