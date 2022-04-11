# Linux file descriptor

参考：

https://en.wikipedia.org/wiki/File_descriptor

在Unix中file descriptor是文件(linux 上一切皆文件)的唯一标识符，通常用一个非负正整数表示。==每一个Unix process都有下表所示的3个fd==，其中

| Integer value |      Name       | <unistd.h> symbolic constant | <stdio.h> file stream |
| :-----------: | :-------------: | :--------------------------: | :-------------------: |
|       0       | Standard input  |         STDIN_FILENO         |         stdin         |
|       1       | Standard output |        STDOUT_FILENO         |        stdout         |
|       2       | Standard error  |        STDERR_FILENO         |        stderr         |

fd通过system call来进行读写操作，进程通过file table找到fd对应的数值，然后在`/proc/PID/fd`目录下找到对应的fd，然后才能进行读写。

linux中通过open()，create()，socket()等函数来创建fd，通过read()，write()

通常在unistd.h类库中









