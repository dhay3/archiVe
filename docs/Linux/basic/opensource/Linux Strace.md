# Linux Strace 

## 概述

strace intercepts and records the system calls which are called by a process and the signals which are received by a process.

EBNF：`strace ::= "debuging tool","diagnostic";`

syntax：`strace [options] command`

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

### output

- `-o filename`

   吧strace的输出重定向到文件中

- `-i | --instruction-pointer`

  在

### filtering

- `-z | --successful-only`

  只打印syscall成功的

- `-Z | --failed-only`

  只打印syscal失败的



## 输出分析

```
open("/foo/bar", O_RDONLY) = -1
```

这里表示调用open函数（a type fo systemcall），返回-1（system call return value），因为文件不存在，通常返回-1

```
open("xyzzy", O_WRONLY|O_APPEND|O_CREAT, 0666) = 3
```

这里表示生成一个只写的文件，mask（==这里不是umask==） 0666即文件的permission是rwrwrw

