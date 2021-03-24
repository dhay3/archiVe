# 标准流

参考：

https://program-think.blogspot.com/2019/11/POSIX-TUI-from-TTY-to-Shell-Programming.html?q=bash&scope=all

https://xz.aliyun.com/t/2548



> pipeline(管道符)，==将前一个命令的stdout做为后一个命令的stdin==，采用FIFO

```
cat filename | command1 && >> filname.bak
```

会先执行command1 然后将filename中的内容写入到filename.bak中

## 文件描述符

POSIX 操作系统会分配一个数字用于跟踪打开的文件，这个数字就叫做文件描述符。Linux 启动时会默认打开三个文件描述符：

1. 标准输入stdin 0 `/dev/stdin`
2. 标准输出 stdout 1 `/dev/stdout`
3. 错误输出 stderr 2 `/dev/stderr`



<img src="..\..\..\imgs\_Linux\480px-Stdstreams-notitle.svg.png"/>

一条shell命令，都会继承其父进程的文件描述符，因此所有的shell命令，都会默认有三个文件描述符。

## 重定向

命令执行的结果默认会在stdout上显示，如果想要输出到其他文件中就需要使用重定向。

- `>`

  等价于 `1 > `

- `<`

  等价于 `< 0`

- ``

## 标准流的重定向

> 如果流A被重定向到B流，但是B流被重定向到文件C，那么，A和B都会被保存到文件C中。

- 输入流重定向

  把某个文件重定向为stdin；此时进程通过stdin读取的是该文件内容。使用`<`来表示

  ```
  cat < 文件名
  ```

  > 注意echo 使用的并不是输入流，而是命令行参数，cat命令就是输入流，类似的还有wc

<img src="..\..\..\imgs\_Linux\Snipaste_2020-10-11_13-49-37.png"/>

- 输出流重定向

  把stdout重定向到==某个文件==，`>`覆盖文件内容，`>>`追加文件内容

  ==`>`等价于`1 >`，将标准输出重定向==
  
  ```
  echo hello > redirect
echo hello >> redirect
  ```
  
  > ==如果是整合流不能有空格==
  >
  > 2>&1 表示将stderr(2)==合并==到stdout(1)，错误信息将会显示在屏幕
  >
  > ```
  > echo hello > redirect 2>&1
  > echo hello >> redirect 2>&1
  > ```
  >
  > 可以将&理解为整合通道，如果没有&，将被bash理解为文件
  >

`cat`命令还可以起到类似“文件复制”效果

```
cat < src > dest
```

> 某些同学可能会问了：既然能这么玩，为啥还需要用 `cp` 命令进行文件复制捏？
> 原因在于：`cat` 的玩法，只保证内容一样，其它的不管；而 `cp` 除了复制文件内容，还会确保“目标文件”与“源文件”具有相同的属性（比如 mode）。

### `>|`

https://unix.stackexchange.com/questions/45201/bash-what-does-do

在一些shell中有一个noclobber选项，保护文件因为重定向而被覆盖或销毁。

如果noclobber设置为true，并且`/tmp/output.txt`文件存在，那么下面的命令将会失败

```
some-command > /tmp/output.txt
```

但是你能使用`>|`符号，表示强制重定向

## <>

双向重定向，会导致阻塞，可以通过`timeout`命令来终止

```
[root@8d3d229c-4aab-4812-96b9-37c8bc47a1d8 opt]# timeout 2s  cat <> /dev/zero
```

## example

参考：

https://linux.cn/article-3464-1.html

https://blog.csdn.net/huangjuegeek/article/details/21713809

1. `command 2>file1`

   将stderr保存到file1中，标准输出还是显示在屏幕

2. `command 1>file1 2>file2`

   将stdout和stderr分别保存到file1和file2

3. `>&2`

   等价于`1>&2`，将stdout重定向到stderr

4. `>&file1`

将stoud和stderr都重定向到file1

5. `|&`

   等价于`2>&1 |`





