# Go4期 day03

# 每周分享

the beat way to learn code is to code.

走上坡路总是会累的。

# 内容回顾

## 数组

### 数组的定义

数组是带长度和元素类型！！！

```go

var a [3]int
var b [3]string
var c [5]int
```

### 数组的初始化

```go
var a = [3]int{1,2,3}
var b = [...]bool{true, false}
var c = [5]string{1:"娜扎", 3:"小王子"}
```

### 多维数组

```go
var a = [3][2]int{
    {1, 2},
    {3, 4},
    {5, 6},
}
```

### 数组的遍历

1. 根据索引index遍历
2. for range遍历

## 切片

### 切片的定义

切片与数组的区别是没有固定的长度。

切片本质上就是一个数组。是语言层面对数组的封装。

切片是引用类型。

```go
var a []int
var b []string
var c []bool
```

<img src="..\..\..\..\..\..\..\..\Go\src\code.oldboy.com\studygolang\day03\assets\1555208079167.png"/>



<img src="..\..\..\..\..\..\..\..\Go\src\code.oldboy.com\studygolang\day03\assets\1555208177025.png"/>

### Go语言中类型的初始值

#### 值类型

```go

var a int//0
var b string//
var c float32//0
var d [3]int//[0 0 0]
var e bool//false
```

#### 引用类型 make()是给引用类型申请内存空间

```go
var x []int//nil

```



### 切片的初始化

1. 声明变量时初始化
2. 数组切片得到
3. 切片再切片得到
4. make()函数创建

### 切片追加元素 append()

必须用变量接收`append`函数的返回值

### 切片的复制

copy(目标切片, 原切片)

### 切片删除元素

```go
a = append(a[:1], a[2:]...)
```



## map

### map的定义

map是引用类型， 默认值是`nil`

map专门用来存储键值对类型的数据

### 判断某个键是否存在

```go
v, ok := m[key]
```



## 函数

### 函数的定义

### 函数的参数

不带参数

多个相同类型的参数可以省略前面参数的类型

可变参数(`...`)

```go
func f1(){
	fmt.Println("hello 沙河！")
}
//参数类型简写
func f2(name1, name2 string){
	fmt.Println(name1)
	fmt.Println(name2)
}
//可变参数 0个或多个
func f3(names ...string){
	fmt.Println(names)//[]string
}
```



### 函数的返回值

不带返回值

单个返回值

多个返回值

命名返回值



```go
func f4(){
	fmt.Println("MJJ真烦")
}
func f5(a, b int) int{
	return a+b
}
//多返回值,必须用括号括起来，用英文逗号分隔
func f6(a, b int) (int, int){
	return a+b, a-b
}

//命名的返回值
func f6(a, b int) (sum int, sub int){
	sum = a + b
	sub = a - b
	return
}
```



## 匿名函数和闭包

#### 变量的作用域

 	1. 全局作用域
 	2. 函数作用域
 	3. 语句块作用域



### defer

defer延迟执行函数

defer先注册后执行（在函数即将要结束但还没有结束的时候执行）

多用来做一些资源回收类的工作。



# 今日内容

## 指针

### 指针和地址有什么区别？

地址：就是内存地址（用字节来描述的内存地址）

指针：指针是带类型的。

### `&`和`*`

`&`：表示取地址

`*`:根据地址取值

## new和make

new:是用来初始化值类型指针的

make:是用来初始化`slice` 、`map`、 `chan`

```go
	// 以下是错误的写法
	// var a *int //a是一个int类型的指针
	// var b *string
	// var c *[3]int
	// 以上是错误的写法
	var a = new(int) //得到一个int类型的指针
	fmt.Println(a)

	*a = 10
	fmt.Println(a)
	fmt.Println(*a)

	var c = new([3]int)
	fmt.Println(c)
	c[0] = 1
	fmt.Println(*c)
```

## panic和recover

panic：运行时异常

recover：用来将函数在panic时恢复回来，用于做一些资源回收的操作

**注意：**程序会出现panic一定是不正常的。

## 结构体（struct）和方法

`type`关键字用来在Go语言中定义新的类型。

### 自定义类型和类型别名

#### 自定义类型

创造一个新类型

```go

type NewInt int
```

#### 类型别名（软链）

```go
var MyInt = int
```

`byte`: uint8 和 `rune`:int32是Go语言内置的别名。

类型别名只在代码编写过程中生效，编译完不存在。

## 结构体的定义

```go
//结构体
// 创在新的类型要使用type关键字
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var haojie = student{
		name:   "豪杰",
		age:    19,
		gender: "男",
		hobby:  []string{"篮球", "足球", "双色球"},
	}
	//结构体支持.访问属性
	fmt.Println(haojie)
	fmt.Println(haojie.name)
	fmt.Println(haojie.age)
	fmt.Println(haojie.gender)
	fmt.Println(haojie.hobby)
}

```

## 结构体的实例化

1. 基本实例化

   ```go
   var haojie = student{
   		name:   "豪杰",
   		age:    19,
   		gender: "男",
   		hobby:  []string{"篮球", "足球", "双色球"},
   	}
   ```

## 结构体的初始化

## 结构体的内存布局

了解为主

结构体的字段在内存上是连续的。

```go
func main() {
	type test struct {
		a int16
		b int16
		c int16
	}

	var t = test{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(&(t.a))
	fmt.Println(&(t.b))
	fmt.Println(&(t.c))
}
```

## 构造函数

```go
func newStudent(n string, age int, g string, h []string) *student {
	return &student{
		name:   n,
		age:    age,
		gender: g,
		hobby:  h,
	}
```

## 基于函数版学员信息管理系统

### 获取终端输入

[fmt.Scan系列](<https://www.liwenzhou.com/posts/Go/fmt_Scan/>)

### 需求分析

```go
// 使用函数实现一个简单的图书管理系统。
// 每本书有书名、作者、价格、上架信息，
// 用户可以在控制台添加书籍、修改书籍信息、打印所有的书籍列表。

// 0. 定义结构体
// 1. 打印菜单
// 2. 等待用户输入菜单选项
// 3. 添加书籍的函数
// 4. 修改书籍的函数
// 5. 展示书籍的函数
// 6. 退出 os.Exit(0)
```



### 代码实现

详见课上代码 11book_mgr



## 结构体的嵌套

## 结构体的匿名字段

# 本周作业

1. 把上课的内容整理一下笔记
2. 把上课的图书管理系统自己写一遍函数版和方法版
3. 写一个学生管理系统（要交的作业）
   1. 学生有姓名 年龄 id 班级 
   2. 增加学生/修改学生/删除学生/展示学生
   3. 用结构体+方法的形式（面向对象的思维方式）



作业一定要交。
