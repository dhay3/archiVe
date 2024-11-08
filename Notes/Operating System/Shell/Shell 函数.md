# Bash Function

> <font color="red">Shell中函数体内直接声明的变量，属于全局变量，整个脚本都可以读取，与javascript相似</font>

## 概述

与javascript类似shell也会进行变量提升，==同样也有闭包，可以将函数定义在函数中==

```
func1()
{
    func2
}

func2()
{
    echo "In func2, var = $var"
}

var=global
func1

----

echo In func2, var = global
```

pattern：

1. ```
   fn() {
     # codes
   }
   ```

2. ```
   function fn() {
     # codes
   }
   ```

> 使用`declare -f`来查看所有已经定义的函数
>
> `declare -f functionname`来查看某一特定的函数

下面是一个简单函数的例子。

> 函数内部的`$1`和`$2`分别代表形参，调用函数时直接在函数后面带上形参

```
[root@cyberpelican ~]# cat test.sh 
add(){
 echo $(($1+$2))
}
add 2 3
[root@cyberpelican ~]# ./test.sh 
5
```

==函数中同样可以使用特殊变量==

```
[root@cyberpelican ~]# cat test.sh 
function alice {
  echo "alice: $@"
  echo "$0: $1 $2 $3 $4"
  echo "$# arguments"

}
alice in wonderland
[root@cyberpelican ~]# ./test.sh 
alice: in wonderland
./test.sh: in wonderland  
2 arguments
```

## return

`return`命令用于从函数返回一个==数值==。函数执行到这条命令，就不再往下执行了，直接返回了。==不能返回数组或字符串，否则就会numeric argument required，可以设置一个全局变量来接受==

```
[root@cyberpelican ~]# cat test.sh 
return_func (){
return 10
}
return_func
echo "return_func $?" 
[root@cyberpelican ~]# ./test.sh 
return_func 10
```

## 全局变量/局部变量

==Bash 函数体内直接声明的变量，属于全局变量，整个脚本都可以读取。这一点需要特别小心。==

```
[root@cyberpelican ~]# cat test.sh 
fn(){
foo=1
echo "fn: foo=$foo"
}
fn
echo "global: foo=$foo"
[root@cyberpelican ~]# ./test.sh 
fn: foo=1
global: foo=1
```

==函数体内不仅可以声明全局变量，还可以修改全局变量。==

```
[root@cyberpelican ~]# cat test.sh 
foo=1
fn(){
foo=2
}
fn
echo $foo
[root@cyberpelican ~]# ./test.sh 
2
```

==函数里面可以用`local`命令声明局部变量。==

```
[root@cyberpelican ~]# cat test.sh 
foo=1
fn(){
local foo=2
echo "fn: foo = $foo"
}
fn
echo "global: foo = $foo"
[root@cyberpelican ~]# ./test.sh 
fn: foo = 2
global: foo = 1
```

