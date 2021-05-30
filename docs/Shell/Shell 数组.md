# Shell 数组

> Shell中数组的下标与其他编程语言一样都从0开始

## 创建数组

1. 逐个赋值

   ```
   $ array[0]=val
   $ array[1]=val
   $ array[2]=val
   ```

2. 一次性赋值

   ```
   ARRAY=(value1 value2 ... valueN)
   ```

   等价于

   ```
   ARRAY=(
     value1
     value2
     value3
   )
   ```

   采用上面方式创建数组时，可以按照默认顺序赋值，也可以在每个值前面指定位置。

   ```
   $ array=(a b c)
   $ array=([2]=c [0]=a [1]=b)
   
   $ days=(Sun Mon Tue Wed Thu Fri Sat)
   $ days=([0]=Sun [1]=Mon [2]=Tue [3]=Wed [4]=Thu [5]=Fri [6]=Sat)
   ```

3. 通配符

   ```
   [root@cyberpelican ~]# touch -f {a..c}.mp3
   [root@cyberpelican ~]# mp3s=( *.mp3 )
   [root@cyberpelican ~]# echo ${mp3s[@]}
   a.mp3 b.mp3 c.mp3
   ```

   上面例子中，将当前目录的所有 MP3 文件，放进一个数组。

4. read创建数组

   ```
   read -a dice
   ```

   `read -a`命令则是将用户的命令行输入，读入一个数组。

5. map

   可以指定数组的key，来自定义map

   ```
   [root@cyberpelican ~]# colors["red"]=red
   [root@cyberpelican ~]# echo ${colors["red"]}
   red
   ```

   

## 读取数组

1. 读取单个元素，必须使用`${}`来获取变量

   ```
   [root@cyberpelican ~]# array=(11 12 13)
   [root@cyberpelican ~]# echo ${array[2]}
   13
   ```

2. `[@]`读取所有成员

   ```
   [root@cyberpelican ~]# foo=(a b c d e f)
   [root@cyberpelican ~]# echo ${foo[@]}
   a b c d e f
   ```

   ```
   [root@cyberpelican ~]# for i in ${foo[@]};do echo $i;done
   a
   b
   c
   d
   e
   f
   ```

   > 注意list是否带有引号有区别，一般需要带有引号

   ```
   [root@cyberpelican ~]# cat test.sh 
   activities=( swimming "water skiing" canoeing "white-water rafting" surfing )
   
   for act in ${activities[@]}
   do
   echo "Activity: $act"
   done
   [root@cyberpelican ~]# ./test.sh 
   Activity: swimming
   Activity: water
   Activity: skiing
   Activity: canoeing
   Activity: white-water
   Activity: rafting
   Activity: surfing
   ```

   上面的例子，数组`activities`实际包含5个元素，但是`for...in`循环直接遍历`${activities[@]}`，会导致返回7个结果。为了避免这种情况，一般把`${activities[@]}`放在双引号之中。

   ```
   [root@cyberpelican ~]# cat test.sh 
   activities=( swimming "water skiing" canoeing "white-water rafting" surfing )
   
   for act in "${activities[@]}"
   do
   echo "Activity: $act"
   done
   [root@cyberpelican ~]# ./test.sh 
   Activity: swimming
   Activity: water skiing
   Activity: canoeing
   Activity: white-water rafting
   Activity: surfing
   
   ```

## ==拷贝数组==

```
[root@cyberpelican ~]# hobbies=( "${activities[@]}" )
[root@cyberpelican ~]# echo ${hobbies[@]}
swimming water skiing canoeing white-water rafting surfing
```

## 数组的长度

数组没有固定的长度，所以只是==读取数组中有效的元素个数==

1. `${#array[*]}`
2. `${#array[@]}`

```
[root@cyberpelican ~]# a[100]=1
[root@cyberpelican ~]# echo ${a[0]}

[root@cyberpelican ~]# echo ${#a[@]}
1

```

因为只给`a[100]`赋值了，所以数组的中有效的元素个数为1

注意如果读取某一个具体的元素，就是读取元素的长度

```
[root@cyberpelican ~]# a[100]=foo
[root@cyberpelican ~]# echo ${#a[100]}
3
```

## 提取数组序号

1. `${!array[@]}`
2. `${!array[*]}`

```
[root@cyberpelican ~]# arr=([5]=a [9]=b [23]=c)
[root@cyberpelican ~]# echo ${!arr[@]}
5 9 23
```

通过`for`循环遍历数组。

```
arr=(a b c d)

for i in ${!arr[@]};do
  echo ${arr[i]}
done
```

## 提取数组成员

`${array[@]:position:length}`的语法可以提取数组成员。

```
$ food=( apples bananas cucumbers dates eggs fajitas grapes )
$ echo ${food[@]:1:1}
bananas
$ echo ${food[@]:1:3}
bananas cucumbers dates
```

上面例子中，`${food[@]:1:1}`返回从数组1号位置开始的1个成员，`${food[@]:1:3}`返回从1号位置开始的3个成员。

如果省略长度参数`length`，则返回从指定位置开始的所有成员。

```
$ echo ${food[@]:4}
eggs fajitas grapes
```

上面例子返回从4号位置开始到结束的所有成员。

## 追加数组成员

数组末尾追加成员，可以使用`+=`赋值运算符。它能够自动地把值追加到数组末尾。否则，就需要知道数组的最大序号，比较麻烦。

```
$ foo=(a b c)
$ echo ${foo[@]}
a b c

$ foo+=(d e f)
$ echo ${foo[@]}
a b c d e f
```

## 删除数组

1. 使用`unset`命令。

   ```
   $ foo=(a b c d e f)
   $ echo ${foo[@]}
   a b c d e f
   
   $ unset foo[2]
   $ echo ${foo[@]}
   a b d e f
   ```

2. 将某个成员设为空值，可以从返回值中“隐藏”这个成员。

   ```
   $ foo=(a b c d e f)
   $ foo[1]=''
   $ echo ${foo[@]}
   a c d e f
   ```

   > 注意，这里是“隐藏”，而不是删除，因为这个成员仍然存在，只是值变成了空值。

    ```
 [root@cyberpelican ~]# foo=(a b c d e f)
    [root@cyberpelican ~]# foo[1]=''
    [root@cyberpelican ~]# echo ${#foo[@]}
    6
    [root@cyberpelican ~]# echo ${!foo[@]}
    0 1 2 3 4 5
    ```
   
    上面代码中，第二个成员设为空值后，数组仍然包含6个成员。

3. 清空整个数组

   ```
   $ unset ARRAY
   
   $ echo ${ARRAY[*]}
   <--no output-->
   ```

   
