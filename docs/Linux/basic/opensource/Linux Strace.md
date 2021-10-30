# Linux Strace 

## 概述

strace intercepts and records the system calls which are called by a process and the signals which are received by a process.

EBNF：`strace ::= "debuging tool","diagnostic";`

syntax：`strace [options] command`

当command退出时，strace会以同样的方式退出(exit code相同)

## system call

> 所有的syscall都可以通过manual page查看，大都数的syscall都在高级编程语言中有库

- execve(const char *pathname,char *const argv[],char *const envp[])

  execute the program referred to  by pathname

  ```
  execve("/usr/bin/ls", ["ls"], 0x7ffeaf291a38 /* 70 vars */)
  ```

- brk(void *addr)

  https://www.geeksforgeeks.org/void-pointer-c-cpp/

  void pointer 可以理解为JAVA中的多态，可以表示任意类型的变量，例如malloc和calloc返回的值是*void，说明可以分配任意类型的memory

  change the location fo the program break which defines the end of the process’s data segment

  ```
  [00007f4c27e5da6b] brk(NULL)
  brk(0x55965e0d2000)
  ```

- mmap(void *addr,size_t len,int prot,int flags,int fildes,off_t off)

  map pages of memory

  ```
  mmap(0x7f2b44c6d000, 311296, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x171000)
  ```

- open()

- openat(int fd,const char *path,int oflag,...)

  打开文件相关的fd，使用path(==相对路径==)下的fd

  ```
  openat(AT_FDCWD, "/usr/lib/locale/locale-archive", O_RDONLY|O_CLOEXEC) = 3
  ```

  

- close(int fd)

  `man -k close | grep -w close;man close.2`

  close a file descriptor

  ```
  close(3)
  ```

- read

- write

- exit_group

## options

### filtering    

- `-e expr`

  指定trace的方式，`expr:=[qualifier=][!]value[,value]...`。

  qualifier有如下几种

  1. trace(t)

     输出和syscall相关的信息

     value有如下几种，syscall，?value，/regex(使用POSIX Extended Regex)，%process(输出进程creation，exec，termination相关的syscall)，%network(==输出与网络有关的syscall==)，%signal(==输出与posix signal相关的syscall==)，%memory(输出与内存相关的syscall)。默认使用all表示输出所有的syscall

  2. signal(s)

     输出和signal相关的信息

     默认all

  3. status

     输出指定返回状态的syscall

     状态可以是如下几种

     successful表示无错误返回码（等价于-z）

     failed错误返回码(等价与-Z)

     unfinished没有终止的

     unavailable不能获取错误信息的

     detached没返回就终止的

     默认all

  4. abbrev(a)

     简要的输出syscall，默认all

  5. verbose(v)

     详细的输出syscall，默认all

  6. raw(x)

     以16进制的方式输出相关信息，默认all

  7. read(r)

     读取read syscall读取的内容以ascii的方式输出，例如`-e read=3,5`表示只读3和5的stream

     ```
     strace -e trace=read -e read pwd
     ```

  8. write(w)

     和read syscall类似

  9. quiet(q)

     不输出指定的内容，可以如下几种值，具体查看man page

     attach

     ```
     ("[Process NNNN attached ]",  "[Process NNNN detached ]")
     ```

     exit

     ```
     +++  exitedwith SSS +++
     ```

  例如`-e trace=open`表示只看和open相关的syscall

### startup

- `-p pid`

  以流的方式跟踪进程

  ```
  sudo strace -p <pid>
  ```

- `-E env`

  运行命令前去掉特定的环境变量

  ```
  cpl in /tmp λ strace -E PATH -o 1 ls
  ```

  这里就不会运行，我们可以使用diff来观察

  ```
  write(1, " 2\n", 3) = 3
  ```

- `-u username`

  以特定用户执行命令，并将信息返回给strace

### tracing

- `-b | --detach-on <syscall>`

  如果接受到指定syscall，就退出strace。目前只支持`execve`

- `-f | --follow-forks`

  strace以fork的方式跟踪当前进程，如果和`-p`一起使用，表示fork指定进程的所有值进程

- `--output-separately`

  和`--output=filename`一起使用时才有效，将pid的信息和输出信息分离。文件以`fliename.pid`命名

- `-ff`

  等价与`--follow-forks`加`--out-separately`

- `-P path`

  只追踪path下的syscall

### output

- `-o filename`

   把strace的输出重定向到文件中，然后可以结合grep来查看具体的信息

- `-k | --stack-traces`

  输出每个syscall的堆栈信息

- `-n | --syscall-number`

  输出syscall number

- `-z | --successful-only`

  只打印syscall成功的

- `-Z | --failed-only`

  只打印syscal失败的

- `-q[q...]`

  q越多，输出的信息越少

- `-t | --absolute-timestamps`

  在每个syscall之前都加时间戳

- `-d`

  输出和strace相关的debug信息

### statistics

- `-c | --summary-only`

  统计时间，syscall数量，错误数，当strace进程退出时

  ```
  cpl in ~/note on master λ strace -c pwd
  /home/cpl/note
  % time     seconds  usecs/call     calls    errors syscall
  ------ ----------- ----------- --------- --------- ----------------
    0.00    0.000000           0         1           read
    0.00    0.000000           0         1           write
    0.00    0.000000           0         5           close
    0.00    0.000000           0         9           mmap
    0.00    0.000000           0         3           mprotect
    0.00    0.000000           0         1           munmap
    0.00    0.000000           0         3           brk
    0.00    0.000000           0         4           pread64
    0.00    0.000000           0         1         1 access
    0.00    0.000000           0         1           execve
    0.00    0.000000           0         1           getcwd
    0.00    0.000000           0         2         1 arch_prctl
    0.00    0.000000           0         3           openat
    0.00    0.000000           0         4           newfstatat
  ------ ----------- ----------- --------- --------- ----------------
  100.00    0.000000           0        39         2 total
  ```

- `-C | --summary`

  和`-c`类似，但是会输出syscall的信息

## 输出分析

```
open("/foo/bar", O_RDONLY) = -1
```

这里表示调用open函数（a type fo systemcall），返回-1（system call return value），因为文件不存在，通常返回-1

```
open("xyzzy", O_WRONLY|O_APPEND|O_CREAT, 0666) = 3
```

这里表示生成一个只写的文件，mask（==这里不是umask==） 0666即文件的permission是rwrwrw

