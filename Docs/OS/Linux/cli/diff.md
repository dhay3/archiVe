# Linux diff

ref:

https://unix.stackexchange.com/questions/81998/understanding-of-diff-output

https://askubuntu.com/questions/229447/how-do-i-diff-the-output-of-two-commands

## Digest

syntax：`diff [options] ...files`

Compare files ==line by line==

一般 file1 是对比的原始文件，file2 是来对比的文件

if a file is ‘-’, read stdout

exit status is 0 if inputs are the same, 1 if different, 2 if trouble

## Output

以下面的为例子说明(这里使用defualt format)

```
cpl in /tmp λ cat file1
2
1
4
8
5
cpl in /tmp λ cat file2
3
4
5
6
7
8
cpl in /tmp λ diff file1 file2
1,2c1
< 2
< 1
---
> 3
4d2
< 8
5a4,6
> 6
> 7
> 8
```

- `>`

  表示内容属于右边的文件即 file2，但是 file1 缺失该部分

- `<`

  表示内容属于左边的文件即 file1，但是 file2 缺少该部分

- `a/d/c`

  adding/deletion/change

  字母左边的数字通常表示左边的文件的行数

  字母右面的数字通常表示右边的文件的行数

  例如：

  1. `1,2c1`

     表示 file1 对比 file2，如果 file1 需要变成 file2 需要将 1,2 行闭区间替换成 file2 的第 1 行

     文件 a 则为

     ```
     3
     4
     8
     5
     ```

  2. `4d2`

     表示 file1 对比 file2, 如果 file1 需要变成 file2 需要在 file1 的第 4 行删除 file1 的第 4 行。==d 右边的数字比较特殊，表示 file1 删除后当前行的计数器(行)，会计算之前操作过的(1,2c1)==

     文件 a 则为

     ```
     3
     4
     5
     ```

  3. `4a4,6`

     表示 file1 对比 file2, 如果 file1 需要变成 file2 需要在第 4 行增加 file2 4,6 行闭区间的内容

     文件 a 则为

     ```
     3
     4
     5
     6
     7
     8
     ```

## Optional args

- `--normal`

  比较不同使用 normanl format 输出内容，缺省值

  ```
  cpl in /tmp λ seq 5 > a
  cpl in /tmp λ seq 10 > b
  cpl in /tmp λ diff a b
  5a6,10
  > 6
  > 7
  > 8
  > 9
  > 10
  ```

- `-u | --unified`

  使用 unified format 输出内容，==提一嘴 github 就是使用这种格式 diff history code==

  ```
  cpl in /tmp λ diff -u file1 file2
  --- file1       2022-07-27 23:21:55.338059251 +0800
  +++ file2       2022-07-27 21:58:39.446240093 +0800
  @@ -1,5 +1,6 @@
  -2
  -1
  +3
   4
  -8
   5
  +6
  +7
  +8
  ```

  可以对比参考 #Output 部分

  `--- file1` 表示为需要对比的源文件

  `+++ file2 ` 表示为对比文件

  `@@ -1,5 +1,6 @@` 表示对比的内容行数；file1 为 1,5 闭区间，file2 为 1,6 闭区间

  `-2 -1 -8` 表示在 file1 中存在，但是不在 file2 中不存在，file1 需要删除后才能变成 file2

  `+3 +6 +7 +8` 表示在 file1 中不存在，但是在 file2 中存在，file1 需要添加后才能变成 file2

- `-r | --recursive`

  recursively campare any subdirectories found

- `--no-dereference`

  don’t follow symbolic links, 意味这默认会 follow symbolic link

- `--ignore-case`

  ignore case differences in file contents

  ```
  cpl in /tmp λ diff --ignore-case a b
  cpl in /tmp λ diff  a b                       
  1c1
  < a b c d e f g h i j k l m n o p q r s t u v w x y z
  ---
  > A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
  ```

  默认和 grep 一样大小写敏感

- `--ignore-space-change`

  忽略空格导致的不一致

- `--strip-trailing-cr`

  strip trailing carriage return on input

  非常有用，可以排除 OS 导致不一致

- `-q | --brief`

  report only when files differ

- `-s | --report-identical-files`

  report when two files are the same

- `-y | --side-by-side`

  输出原始内容比对比不同

  ```
  cpl in /tmp λ diff -y a b 
  1                                                               1
  2                                                               2
  3                                                               3
  4                                                               4
  5                                                               5
                                                                > 6
                                                                > 7
                                                                > 8
                                                                > 9
                                                                > 10
  ```

- `--color[=when]`

  colored output, default auto

## Tricks

### stdout cmp to file

```
cpl in /tmp λ seq 10 | diff - file1
1d0
< 1
3c2
< 3
---
> 1
5,7d3
< 5
< 6
< 7
9,10c5
< 9
< 10
---
> 5
```

### stdout cmp to stdout

这里巧妙的使用了 substitution shell，将 stdin 重定向到 sub shell 进程打开的 fd

```
cpl in /tmp λ diff <(seq 5) <(seq 3 8)
1,2d0
< 1
< 2
5a4,6
> 6
> 7
> 8
```

```
cpl in / λ echo <(echo) <(echo)  
/proc/self/fd/11 /proc/self/fd/12
```

