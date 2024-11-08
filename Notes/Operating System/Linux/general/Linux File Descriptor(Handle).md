# Linux File Descriptor(Handle)

ref:

https://en.wikipedia.org/wiki/File_descriptor

https://www.bottomupcs.com/file_descriptors.xhtml

https://stackoverflow.com/questions/5256599/what-are-file-descriptors-explained-in-simple-terms

https://man7.org/training/download/lusp_fileio_slides.pdf

https://igm.univ-mlv.fr/~yahya/progsys/linux.pdf

https://sciencesoftcode.files.wordpress.com/2018/12/the-linux-programming-interface-michael-kerrisk-1.pdf

## Digest

> File descriptor 直接可以理解成 File handle

In Unix like OS, a file descriptor(FD) is a unique identifier (handle) for a file or other input/output resource, such as a pipe or network socket

FD 中文通常也叫文件描述符，是==open file==的唯一标识符（区别与 inode）。用于跟踪==打开的文件==。POSIX 中也有一种叫法 open file description(OFD)

==需要注意的是任意只要是打开的文件都有 FD==

File descriptors typically have non-negative integer values, with negative values being reserved to indicate “no value” or error conditions

这个 unsigned integer 通常也叫做 FD flags

FD 实际取自 C lib，且通常关联一个 unsigned integer，negative integer 一般表示空值。

下面是 Unix-like OS 中常见的 FD

| Integer value | Name                                                    | <[unistd.h> symbolic constant | <stdio.h> file stream |
| ------------- | ------------------------------------------------------- | ----------------------------- | --------------------- |
| 0             | [Standard input](https://en.wikipedia.org/wiki/Stdin)   | STDIN_FILENO                  | stdin                 |
| 1             | [Standard output](https://en.wikipedia.org/wiki/Stdout) | STDOUT_FILENO                 | stdout                |
| 2             | [Standard error](https://en.wikipedia.org/wiki/Stderr)  | STDERR_FILENO                 | stderr                |

## FD in Linux

On Linux, the set of file descriptors open in process can be accessed under the path `/proc/PID/fd`，where PID is the process identifier. Any running process can also access its own file descriptors through the folders `/proc/self/fd`（self 下的文件就是当前运行的 shell 命令的各种信息，等价于当前运行 shell 命令的`proc/pid/`下的内容） and `/dev/fd`

```
➜  self ll /proc/self/fd
lr-x------ root root 64 B Thu Jul 28 01:26:43 2022  3 ⇒ /proc/18094/fd
lrwx------ root root 64 B Thu Jul 28 01:26:43 2022  0 ⇒ /dev/pts/1
lrwx------ root root 64 B Thu Jul 28 01:26:43 2022  1 ⇒ /dev/pts/1
lrwx------ root root 64 B Thu Jul 28 01:26:43 2022  2 ⇒ /dev/pts/1
```

因为`proc/self/fd`下有 stdin, stdout, stderr。==所以所有的 Shell 命令都会默认有 3 个文件FD==

![480px-Stdstreams-notitle](https://github.com/dhay3/image-repo/raw/master/20220727/480px-Stdstreams-notitle.6bfsgmevmmbk.webp)

```
➜  /tmp ll /dev/std*
lrwxrwxrwx root root 15 B Wed Jul 27 20:09:14 2022  /dev/stderr ⇒ /proc/self/fd/2
lrwxrwxrwx root root 15 B Wed Jul 27 20:09:14 2022  /dev/stdin ⇒ /proc/self/fd/0
lrwxrwxrwx root root 15 B Wed Jul 27 20:09:14 2022  /dev/stdout ⇒ /proc/self/fd/1
```

## How FD Comes

FD 是有系统函数调用 return 回的一个变量

下面是 Unix-like OS 中比较常见的可以生成 FD 的系统函数

> FD 生命周期具体调用的函数参考 wiki

1. open()
2. create()
3. socket()
4. accept()
5. pipe()

## How FD works

To perform input or output, the process passes the file descriptor to the kernel through a system call, and the kernel will access the file on behalf of the process. ==The process does not have direct access to the file or inode tables==

在 bottomupcs 看到一个比较形象的比喻

“the file descriptor is the gateway into the kernel’s abstractions fo underlying harddware”

## Limitation on FD

在 Unix-like OS 中 FD 是有数量限制的，通常为 1024 。 我们可以通过 `ulimit -n` 来查看

```
➜  /etc ulimit -n 
1024
```

这也就表示同时能打开的文件只能有 1024 个

## FD table vs file table vs inode table

In the traditional implementation of Unix, file descriptors index into a per-process file descriptor table maintained by the kernel, that in turn indexes into a system-wide table of files opened by processes, called the file table. The file table records the mode with which the file( or other resource ) has been opened: for reading, wrting, appending, and possibly other modes. It also indexes into a third table called inode table that desribes hte actual underlying files

从上面的简述中可以获知以下几点

1. FD table 是针对==单进程的==, 每个进程都有一张 FD table
2. file table 是==system-wide==, 全局只有一张 file table。记录了 file’s mode (acess mode)，具体可以参考[Linux inode]()
3. inode table 才是记录文件地址的实际 table

这是从 wiki 上摘抄下来的一段内容

![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/f8/File_table_and_inode_table.svg/450px-File_table_and_inode_table.svg.png)

File descriptors for a single process, file table and inode table.==Note that multiple file descriptors can refer to the same file table entry and that multiple file table entries can in turn refer to the same inode.== File descriptor 3 does not refer to anything in the file table, signifying that it has been closed

FD/filetable/inode 各自都是多对一关系



学过数据结构的你可能想为什么只存值，FD 怎么找到 File table 和 inode table 呢？这不符合逻辑啊

其实 wiki 上的这张图有比较大的误导性，用下面这张图来解释

![2022-08-02_21-19](https://git.poker/dhay3/image-repo/blob/master/20220802/2022-08-02_21-19.67uwfznwk1hc.webp?raw=true)

实际 FD table 每条 entry 有如下几个字段

具体参考源码中的`include/linux/fdtable.h`

1. FD flag which is an unique unsigned integer
2. file point which reference the specify entry of open file table

对应的 open file table 每条 entry 有如下几个字段

具体参考源码中的`include/linux/fs.h`

1. file offset
2. file status flags which is one of  the return value from open()
3. inode point which reference the specify entry of inode table

对应的 inode table 每条 entry 有如下几个字段，具体查看[Linux inode]()

具体参考源码中的`include/linux/fs.h`

1. file type
2. file permissions
3. other file properties(size,timestamps...)

## Duplicated file descriptors



## Source code





