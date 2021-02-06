# Linux Make

参考：

https://www.ruanyifeng.com/blog/2015/02/make.html

## 概述

> 代码变成可执行文件，叫做编译（compile）；先编译这个，还是先编译那个（即编译的安排），叫做构建（build）。
>
> ==make支持Bash模式扩展==

Make是最常用的构建工具，用于更新；可以在名为`makefile`或`Makefile`的文件中定义规则，然后使用make来构建（推荐使用`Makefile`）

## 文件格式

Makefile文件由一系列规则（rules）构成。每条规则的形式如下。

```
<target> : <prerequisites> 
[tab]  <commands>
make <target> 运行makefile
```

1. target：目标，不可省略
2. prerequisites：前置条件，可缺，与command至少存在一个
3. ==换行的tab键是必须的==
4. command：命令，可缺，与prerequisites至少存在一个

### target

> 如果make命令运行时没有指定target，默认执行Makefile的第一个target

**0x001**

target通常是文件名，指明Make命令要构建的对象。target可以是一个文件也可以是多个文件，之间用空格隔开

```
[root@cyberpelican opt]# cat Makefile 
a.txt: /proc/version
	cat /proc/version >> a.txt
[root@cyberpelican opt]# make a.txt
cat /proc/version >> a.txt
```

如上命令将`/proc/version`中的内容写入到`a.txt`文件

**0x002**

目标还可以是操作名字(自定义的名字)，称为phony target

```
[root@cyberpelican opt]# cat Makefile 
clean:
	rm test*
[root@cyberpelican opt]# make clean
rm test*

```

但是，如果当前目录中，正好有一个文件叫做clean，那么这个命令不会被执行。

为了避免这种情况，可以明确声明clean是phony target

```
[root@cyberpelican opt]# cat Makefile 
.PHONY: clean
clean:
	rm test*
[root@cyberpelican opt]# make clean
rm test*
```

声明clean是"伪目标"之后，make就不会去检查是否存在一个叫做clean的文件，而是每次运行都执行对应的命令。像.PHONY这样的内置目标名还有不少，可以查看[手册](http://www.gnu.org/software/make/manual/html_node/Special-Targets.html#Special-Targets)。

### prerequisites

前置条件通常为一组文件名，之间用空格分隔。它指定了target是否需要重新构建的判断标准：只要有一个前置文件不存在，"target"就不能构建。

```
[root@cyberpelican opt]# cat Makefile 
source:
	echo "this is source" >> source.txt
	
[root@cyberpelican opt]# make source
echo "this is source" >> source.txt
[root@cyberpelican opt]# cat source.txt 
this is source
[root@cyberpelican opt]# make source
echo "this is source" >> source.txt
[root@cyberpelican opt]# cat source.txt 
this is source
this is source
```

### commands

commands表示如何更新target，由一行或多行Shell命令构成。

每行命令之前必须有一个`tab`键。如果想使用其他键，可以用内置变量`.PECIPEPREFIX`声明

```
.RECIPEPREFIX= >
golang:
>echo "hello world"
```

上述使用`>`替换了`tab`

==每行命令在一个单独的shell中执行，这些shell之间没有继承关系，所以不能在不同行取到变量==

```
[root@cyberpelican opt]# cat Makefile 
var-lost:
	foo=bar
	echo "$$foo"

[root@cyberpelican opt]# make var-lost
foo=bar
echo ""

```

> 注意make会将`$`转译所以需要再添加一个`$`

通过如下几种方式解决

1. 将命令写在同一行

   ```
   [root@cyberpelican opt]# cat  Makefile 
   var-lost:
   	foo=bar;echo "foo=$${foo}"
   
   [root@cyberpelican opt]# make var-lost
   foo=bar;echo "foo=${foo}"
   foo=bar
   ```

2. 使用反斜杠转译

   注意这里虽然使用了backslash，但是同样需要添加分号来表示同一行

   ```
   [root@cyberpelican opt]# cat Makefile 
   var-lost:
   	foo=bar;\
   	echo "foo=$${foo}"
   
   [root@cyberpelican opt]# make var-lost
   foo=bar;\
   echo "foo=${foo}"
   foo=bar
   ```

3. 使用`.ONESHELL`变量

   ```
   [root@cyberpelican opt]# cat Makefile 
   .ONESHELL:
   var-lost:
   	foo=bar;
   	echo "foo=$${foo}"
   
   [root@cyberpelican opt]# make var-lost
   foo=bar;
   echo "foo=${foo}"
   foo=bar
   ```

### 注释

`#`再Makefile中表示注释

```
[root@cyberpelican opt]# cat Makefile 
comment:
	#this is a comment	
	echo "comment"

[root@cyberpelican opt]# make comment
#this is a comment	
echo "comment"
comment
```

### echoing/回声

和windows中的bat编程一样，make会回显每一条命令，然后执行。==注释同样会回显==

```
[root@cyberpelican opt]# cat Makefile 
comment:
	@#this is a comment	
	@echo "comment"

[root@cyberpelican opt]# make comment
comment
```

## 变量

Makefile允许自定义变量，调用变量必须使用`$(variable)`，与Bash调用子命令不可混乱

```
[root@cyberpelican opt]# cat Makefile 
var = hello world
var:
	@#this is a comment	
	echo "$(var)"

[root@cyberpelican opt]# make var
echo "hello world"
hello world
```

调用Shell定义的变量，需要再加一个`$`

```
[root@cyberpelican opt]# cat Makefile 
var-lost:
	foo=bar
	echo "$$foo"
```

## 赋值

`v1=$(v2)`产生一个问题，v1的值是再定义时扩展（静态扩展），还是运行时扩展（动态扩展）。如果v2的值时动态的，结果将会大相径庭。

为了解决类似问题，Makefile一共提供了四个赋值运算符 （=、:=、？=、+=），它们的区别请看[StackOverflow](http://stackoverflow.com/questions/448910/makefile-variable-assignment)。

- `VARIABLE = value`

  在执行时扩展，允许递归扩展

- `VARIABLE := value`

  在定义时扩展。

- `VARIABLE ?= value`

  只有在该变量为空时才设置值。

- `VARIABLE += value`

  将值追加到变量的尾端。

## 隐藏变量

Make命令提供一系列内置变量，比如，$(CC) 指向当前使用的编译器，$(MAKE) 指向当前使用的Make工具。这主要是为了跨平台的兼容性，详细的内置变量清单见[手册](https://www.gnu.org/software/make/manual/html_node/Implicit-Variables.html)。

```
output:
    $(CC) -o output input.c
```

## 自动变量

所有的自动变量清单，请看[手册](https://www.gnu.org/software/make/manual/html_node/Automatic-Variables.html)。

- `$@`

  Makefile中的targets

  ```
  [root@cyberpelican opt]# cat Makefile 
  a.txt b.txt:
  	touch $@
  [root@cyberpelican opt]# make b.txt
  touch b.txt
  
  ---
  等价于
  a.txt:
      touch a.txt
  b.txt:
      touch b.txt
  ```

- `$<`

  代表第一个prerequisites。比如，规则为 t: p1 p2，那么$< 就指代p1。

- `$?`

  `$?` 指代比目标更新的所有前置条件，之间以空格分隔。比如，规则为 t: p1 p2，其中 p2 的时间戳比 t 新，`$?`就指代p2。

- `$^`

  `$^` 指代所有前置条件，之间以空格分隔。比如，规则为 t: p1 p2，那么 `$^` 就指代 p1 p2 。

- `$(@D)` 和` $(@F)`

  `$(@D)` 和 `$(@F) `分别指向 `$@` 的目录名和文件名。比如，`$@`是 `src/input.c`，那么`$(@D) `的值为 src ，`$(@F) `的值为 `input.c`。

- `$(<D)` 和 `$(<F)`

  `$(<D)` 和 `$(<F) `分别指向 `$< `的目录名和文件名。

```
dest/%.txt: src/%.txt
    @[ -d dest ] || mkdir dest
    cp $< $@
```

上面代码将 src 目录下的 txt 文件，拷贝到 dest 目录下。首先判断 dest 目录是否存在，如果不存在就新建，然后，`$<` 指代前置文件（src/%.txt）， `$@` 指代目标文件（dest/%.txt）。@关闭回显
