# Shell 运算

## $((...))

> 由于Shell中所有的默认变量都是字符串，所以才需要模式扩展
>
> tip：`$[]`是以前的语法，也可以做整数，不建议使用
>
> ==注意只有在需要获取变量字符值时才需要使用`$`，否则会报错==
>
> ```
> [root@cyberpelican opt]# ((a++))
> [root@cyberpelican opt]# echo $a
> 1
> [root@cyberpelican opt]# $((a++))
> bash: 1: command not found...
> ```
>
> 

`((...))`语法可以进行==整数==的算术运算。

这个语法不返回值，命令执行的结果根据算术运算的结果而定。==只要算术结果不是`0`，命令就算执行成功。==

`((...))`会自动忽略内部的空格，所以中间有多少个空格都一样

```
[root@cyberpelican ~]# ((3+3))
[root@cyberpelican ~]# echo $?
0
[root@cyberpelican ~]# ((3-3))
[root@cyberpelican ~]# echo $?
1
```

> tip：`$?`如果前一条命令执行错误返回1，否则返回0

`((...))`语法支持的算术运算符如下。

- `+`：加法
- `-`：减法
- `*`：乘法
- `/`：除法（整除）
- `%`：余数
- `**`：指数
- `++`：自增运算（前缀或后缀）
- `--`：自减运算（前缀或后缀）

> tip：除法运算符的返回结果总是整数，比如`5`除以`2`，得到的结果是`2`，而不是`2.5`。

`++`和`--`这两个运算符有前缀和后缀的区别。==作为前缀是先后返回值，作为后缀是先返回值后运算。==

```
[root@cyberpelican ~]# i=0
[root@cyberpelican ~]# echo $((i++))
0
[root@cyberpelican ~]# echo $i
1
[root@cyberpelican ~]# echo $((++i))
2
[root@cyberpelican ~]# echo $i
2
```

`$((...))`内部可以用圆括号改变运算顺序

```
[root@cyberpelican ~]# echo $(((2+3)*4))
20
```

## 进制

Bash 的数值默认都是十进制，但是在算术表达式中，也可以使用其他进制。

- 二进制

  ```
  [root@cyberpelican ~]# echo $((2#111))
  7
  ```

- 十六进制

  ```
  [root@cyberpelican ~]# echo $((0xff))
  255
  ```

## 位运算

`$((...))`支持以下的二进制位运算符。

- `<<`：位左移运算，把一个数字的所有位向左移动指定的位。乘
- `>>`：位右移运算，把一个数字的所有位向右移动指定的位。除
- `&`：位的“与”运算，对两个数字的所有位执行一个`AND`操作。
- `|`：位的“或”运算，对两个数字的所有位执行一个`OR`操作。
- `~`：位的“否”运算，对一个数字的所有位取反。
- `^`：位的异或运算（exclusive or），对两个数字的所有位执行一个异或操作。

## 逻辑运算

`$((...))`支持以下的逻辑运算符。

- `<`：小于

- `>`：大于

- `<=`：小于或相等

- `>=`：大于或相等

- `==`：相等

- `!=`：不相等

- `&&`：逻辑与

- `||`：逻辑或

- `!`：逻辑否

- `expr1?expr2:expr3`：三元条件运算符。==若表达式`expr1`的计算结果为非零值（算术真）==，则执行表达式`expr2`，否则执行表达式`expr3`。

  ```
  [root@cyberpelican ~]# echo $((a>2?10:20))
  20
  [root@cyberpelican ~]# a=0
  [root@cyberpelican ~]# echo $((a==0?10:20))
  10
  ```

## 赋值运算

算术表达式`$((...))`可以执行赋值运算。

```
$ echo $((a=1))
1
$ echo $a
1
```

上面例子中，`a=1`对变量`a`进行赋值。这个式子本身也是一个表达式，返回值就是等号右边的值。

`$((...))`支持的赋值运算符，有以下这些。

- `parameter = value`：简单赋值。
- `parameter += value`：等价于`parameter = parameter + value`。
- `parameter -= value`：等价于`parameter = parameter – value`。
- `parameter *= value`：等价于`parameter = parameter * value`。
- `parameter /= value`：等价于`parameter = parameter / value`。
- `parameter %= value`：等价于`parameter = parameter % value`。
- `parameter <<= value`：等价于`parameter = parameter << value`。
- `parameter >>= value`：等价于`parameter = parameter >> value`。
- `parameter &= value`：等价于`parameter = parameter & value`。
- `parameter |= value`：等价于`parameter = parameter | value`。
- `parameter ^= value`：等价于`parameter = parameter ^ value`。

下面是一个例子。

```
$ foo=5
$ echo $((foo*=2))
10
```

如果在表达式内部赋值，可以放在圆括号中，否则会报错。

```
$ echo $(( a<1 ? (a+=1) : (a-=1) ))
```

## expr

`expr`命令支持==整数==算术运算，可以不使用`((...))`语法。

> tip：支持变量替换，expr需要添加空格，但是如果使用复合运算需要反引号

```
[root@cyberpelican ~]# echo $ok
1
[root@cyberpelican ~]# expr ok+2
ok+2
[root@cyberpelican ~]# expr $ok + 1
2
```

## let

> `let`如果不使用双引号不能在运算符两侧添加空格

`let`命令声明变量时，可以直接执行算术表达式。

```
$ let foo=1+2
$ echo $foo
3
```

上面例子中，`let`命令可以直接计算`1 + 2`。

`let`命令的参数表达式如果包含空格，就需要使用引号。

```
$ let "foo = 1 + 2"
```

`let`可以同时对多个变量赋值，赋值表达式之间使用空格分隔。

```
$ let "v1 = 1" "v2 = v1++"
$ echo $v1,$v2
2,1
```

上面例子中，`let`声明了两个变量`v1`和`v2`，其中`v2`等于`v1++`，表示先返回`v1`的值，然后`v1`自增。

这种语法支持的运算符，参考《Bash 的算术运算》一章。
