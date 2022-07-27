# Linux File Descriptor(Handle)

ref:

https://en.wikipedia.org/wiki/File_descriptor

## Digest

> File descriptor 直接可以理解成 File handle

In Unix like OS, a file descriptor(FD) is a unique identifier (handle) for a file or other input/output resource, such as a pipe or network socket

FD 中文通常也叫文件描述符，是文件的唯一标识符。用于跟踪打开的文件。==需要注意的是任意文件都有 FD==

File descriptors typically have non-negative integer values, with negative values being reserved to indicate “no value” or error conditions

FD 实际取自 C lib，且通常关联一个 unsigned integer，negative integer 一般表示空值。

下面是 Unix like OS 中常见的 FD

| Integer value | Name                                                    | <[unistd.h> symbolic constant | <stdio.h> file stream |
| ------------- | ------------------------------------------------------- | ----------------------------- | --------------------- |
| 0             | [Standard input](https://en.wikipedia.org/wiki/Stdin)   | STDIN_FILENO                  | stdin                 |
| 1             | [Standard output](https://en.wikipedia.org/wiki/Stdout) | STDOUT_FILENO                 | stdout                |
| 2             | [Standard error](https://en.wikipedia.org/wiki/Stderr)  | STDERR_FILENO                 | stderr                |

> FD 生命周期具体调用的函数参考 wiki

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

![480px-Stdstreams-notitle](https://git.poker/dhay3/image-repo/blob/master/20220727/480px-Stdstreams-notitle.6bfsgmevmmbk.webp?raw=true)

```
➜  /tmp ll /dev/std*
lrwxrwxrwx root root 15 B Wed Jul 27 20:09:14 2022  /dev/stderr ⇒ /proc/self/fd/2
lrwxrwxrwx root root 15 B Wed Jul 27 20:09:14 2022  /dev/stdin ⇒ /proc/self/fd/0
lrwxrwxrwx root root 15 B Wed Jul 27 20:09:14 2022  /dev/stdout ⇒ /proc/self/fd/1
```

## FD vs file table vs inode

![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/f8/File_table_and_inode_table.svg/450px-File_table_and_inode_table.svg.png)

File descriptors for a single process, file table and inode table.==Note that multiple file descriptors can refer to the same file table entry and that multiple file table entries can in turn refer to the same inode.== File descriptor 3 does not refer to anything in the file table, signifying that it has been closed

FD/filetable/inode 各自都是多对一关系

