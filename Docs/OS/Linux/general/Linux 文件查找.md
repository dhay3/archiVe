# Linux 文件查找

## which

==一般用于查找系统内建命令==，只能查找可执行文件，该命令基本只在$PATH路径中搜索，查找范围小，速度快。

## whereis

查找源代码，二进制文件，帮助手册文件。与which不同的是通过文件索引库查询，所以比which查找范围广。

## locate

> 使用`locate`出现
>
> ```
> [root@cyberpelican ~]# locate nginx
> locate: can not stat () `/var/lib/mlocate/mlocate.db': No such file or directory
> ```
>
> 解决方法
>
> ```
> updatedb
> ```

locate命令其实是`find -name`的另一种写法，但是要比后者快得多。这个命令可以找到任意你指定要找到的文件，并且==可以只输入部分文件名==

## find

直接搜索硬盘的方式查找，搜索速度慢，但是更加全面

pattern

`find [path...] [[option] [expression]`

```
[root@chz Desktop]# find . -name 'test*'
./test.bak
./test
./test.ttt
./testsource
```

## 同异点

| 对比选项     | which        | whereis                      | locate             | find              |
| ------------ | ------------ | ---------------------------- | ------------------ | ----------------- |
| **搜索目标** | 可执行文件   | 二进制文件、源文件、帮助文件 | 所有类型           | 所有类型          |
| **查找路径** | PATH所含目录 | 索引数据库所含目录           | 索引数据库所含目录 | 当前目录/指定目录 |
| **搜索原理** | 完全匹配     | 去除.之后的完全匹配          | 部分匹配即可       | 遍历寻找          |
| **查找速度** | 非常快       | 比较快                       | 比较快             | 比较慢            |
