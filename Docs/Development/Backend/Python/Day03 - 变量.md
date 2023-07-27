# Day03 - 变量 & 数据类型

## 变量

在 Python 中和 Java 类似也可以在一行中定义多个变量，语法和 go 相同

```
x, y, z = "Orange", "Banana", "Cherry"
```

同样的 Python 中也分为 全局变量 和 局部变量

### 全局变量

即在函数外部的变量，可以被函数调用

```
x = "awesome"

def myfunc():
  print("Python is " + x)

myfunc()
```

在 Python 中还可以使用 `global` 将局部变量转为全局变量

```
def myfunc():
  global x
  x = "fantastic"

myfunc()

print("Python is " + x)
```

需要注意的一点是 Python 和 Java 或者是 C++ 不一样，在函数中不能直接对全局变量赋值

> Python 会理解成定义一个新的局部变量，因为在 Python 中声明变量和赋值并没有明显的区分方式 

```
In [1]: a = 30

In [2]: def func():
   ...:     a = 3
   ...: 

In [3]: func()

In [4]: a
Out[4]: 30

```

如果想要在函数内对变量赋值，需要使用 `global` 或者是 `nonlocal`

### 局部变量

和 Java 中的一样，函数内定义的变量，只能在函数中使用

```
In [1]: def func():
   ...:     a = 30
   ...: 

In [2]: a
---------------------------------------------------------------------------
NameError                                 Traceback (most recent call last)
<ipython-input-2-3f786850e387> in <module>
----> 1 a

NameError: name 'a' is not defined

```

## 数据类型

在 Python 中数据类型可以分为 2 大类

数字型(number) 和 非数字型(non-number)

### 数字型

1. 整型(int)，类似 Java 中的 int，但是包含 long 同时 Python 中整型支持直接表示 binary(0b001)、octect(0o100)、hex decimal(0xff)
2. 浮点型(float)，类似 Java 中 float 和 double 的组合，所有的小数都是浮点型没有精度之分
3. 布尔值(bool)，在 Python 1 和 True 是等价的，0 和 False 是等价的
4. 复数型(complex)，就是数学上的复数例如 3 + 5j

### 非数字型

1. 字符型(string)，在 Python 中不区分双引号或者是单引号均可表示字符串, 也不区分 char 或者是 string，统一都是 string
2. 列表(list)，类似与 Java 中的 ArrayList，元素类型随意
3. 元组(tuple)，类似与数组，长度不能改变，但是和数组不同的是 tuple 中的元素不能变更
4. 字典(dict)，在 3.7 之前是无序的类似 HashMap；之后是有序的，按照 key 插入的顺序输出，类似 LinkedHashMap

## type()

`type()` 是 Python 中用于查看变量对应的类型的函数

```
In [9]: type(2**32)
Out[9]: int

In [10]: type(2**64)
Out[10]: int
```

## 类型转换

| 示例                                     | 数据类型       |
| ---------------------------------------- | -------------- |
| x = str(111)                             | 转换成 str     |
| x = int(‘29’)                            | 转换成 int     |
| x = float(‘29.5’)                        | 转换成 float   |
| x = complex(1j)                          | 转换成 complex |
| x = list(("apple", "banana", "cherry"))  | 转换成 list    |
| x = tuple(("apple", "banana", "cherry")) | 转换成 tuple   |
| x = dict(name="Bill", age=36)            | 转换成 dict    |
| x = set(("apple", "banana", "cherry"))   | 转换成 set     |
| x = bool(5)                              | 转换成 bool    |

## 变量之间的计算

在 Python 中因为不需要考虑数据类型，所以数值型的变量之间可以使用运算符

```
n [11]: i = 10

In [12]: f = 10.5

In [13]: b = True

In [15]: f - i
Out[15]: 0.5

In [16]: f + b
Out[16]: 11.5 #因为 True == 1
```

当然和 Java 中一样， Python 也可以通过 `+` 来拼接字符串

```
In [20]: a = 'be'

In [21]: b = 'ship'

In [22]: a + b
Out[22]: 'beship'
```

还可以使用 `*` 来重复字符串

```
In [23]: a = 'f'

In [24]: a*4
Out[24]: 'ffff'
```

虽然 Python 支持的计算的方式多，但是同样的也不支持字符串和数字型变量相加