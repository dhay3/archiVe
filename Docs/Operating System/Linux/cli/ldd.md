---
title: ldd
author: "0x00"
createTime: 2024-06-03
lastModifiedTime: 2024-06-03-16:16
draft: true
---

# ldd

## 0x01 Libraries

> [!NOTE]
> 这里主要以介绍 Linux Libraries 为主

在介绍 `ldd` 这个命令前，需要先了解一下 Libraries

*In general, libraries are collections of data and functions which is precompiled and written to be reused by other programs.*[^1]

你可以将 libraries 想象成编程语言中的库，是一些可以重复调用的“轮子”，实际上是一些用 C 写的 precompiled binraries
例如你想用 C 写一个进程，要输出 "hello world"。你并不需要自己实现输出的函数，而是可以直接通过 `#include <stdio.h>` 引入 C standard library(libc)[^7]，来调用对应已经预先写好的函数
```shell
   $ cat test.c
   #include <stdio.h>
   
   int main() {
       printf("hello world\n");
       return 0;
   }
   ```
 程序在编译过程中，会将对应需要调用的 libraries 关联到程序
   ```shell
   $ gcc test.c -o test.o
   $ ldd test.o
           linux-vdso.so.1 =>  (0x00007ffe0e17a000)
           libc.so.6 => /lib64/libc.so.6 (0x00007efcd1bd0000)
           /lib64/ld-linux-x86-64.so.2 (0x00007efcd1f9e000)
   ```
虽然 libraries 是用 C 写并编译的，但是不只有 C/C++ 可以调用，其他编程语言都可以调用这些 libraries。例如 Python 中可以通过 `cdll` 来调用 library
```python
from ctypes import cdll
libc=cdll.LoadLibrary("libc.so.6")
libc.printf(b"hello world\n")
```

### Linking

在系统层面 libraries 会在 [link time](https://en.wikipedia.org/wiki/Link_time) 被调用，根据调用的时机不同可以分为 2 类[^2]
1. static linking - linking be done when an executable file is created(compiled)
	在应用编译时执行 linking 调用 libraries，称为 static linking，这些 libraries 被称为 static(archive) libraries
2. dynamic linking - linking be done when an executable file is running
	在应用运行时执行 linking 调用 libraries，称为 dynamic linking， 这些 libraries 被称为 dynamic(shared) libraries 或者 shared object

### Static libraries

*In computer science, a static library or statically linked library is a set of routines, external functions and variables which are resolved in a caller at compile-time and copied into a target application by a compiler, linker, or binder, producing an object file and a stand-alone executable.*[^3]

在 executable file 编译过程中，被调用的库称为 static libraries 也被称为 archive libraries

有如下特点[^5]
1. 在编译的过程中 libraries 会被直接==拷贝==进 executable file，如果调用的 static libraries 很多，生成的 executable file 就会很大
2. 如果一个 static library 被多个 executable files 调用，static library 出现问题，不会导致所有的 executable files 出现问题，因为调用的是拷贝进 executable file 的 library
3. 只使用了 static libraries 的应用比只使用了 dynamic libraries 的应用运行更快，因为 libraries 直接被编译进 executable file
4. 如果 executable file 依赖的 static libraries 改变(例如 函数被重构，Bug 修复)，executable file 想要调用这些改变的部分，就需要重新编译 executable file
5. 在 Linux 中以 `.a` 作为文件后缀，Windows 中以 `.lib` 作为文件后缀

### Shared libraries

*A shared library or shared object is a computer file that contains excutable code designed to be used by multiple computer programs or other libraries at runtime*[^4]

dynamic libraries 也被称为 shared libraries[^3] 或者 shared object，会在应用运行的时候会被调用

有如下特点[^5]
1. 在 executabl file 运行时，会将对应的 dynamic libraries 载入到 memory 中，executable file 通过地址来调用 dynamic libraries
2. 如果一个 dynamic library 被多个 executable files 调用，dynamic library 出现问题，会导致所有的 executable files 出现问题，因为调用的是 地址，指向同一个 dynamic library
3. 只使用了 dynamic libraries 的应用比只使用了 static libraries 的应用运行要慢，因为需要通过 地址 调用 libraries
4. 如果 executable file 依赖的 dynamic libraries 改变(例如 函数被重构，Bug 修复)，executable file 想要调用这些改变的部分，无需重新编译。但是需要通过 `ldconfig` 更新系统的 cache
5. 在 Linux 中以 `.so` (shared object)作为文件后缀，Windows 中以 `.ddl` 作为文件后缀

## 0x02 ldd

`ldd` 用于输出 program 使用的 dynamic libraries，例如
```
ldd /usr/bin/ls
        linux-vdso.so.1 =>  (0x00007ffc97397000)
        libselinux.so.1 => /lib64/libselinux.so.1 (0x00007fb3f875f000)
        libcap.so.2 => /lib64/libcap.so.2 (0x00007fb3f855a000)
        libacl.so.1 => /lib64/libacl.so.1 (0x00007fb3f8351000)
        libc.so.6 => /lib64/libc.so.6 (0x00007fb3f7f83000)
        libpcre.so.1 => /lib64/libpcre.so.1 (0x00007fb3f7d21000)
        libdl.so.2 => /lib64/libdl.so.2 (0x00007fb3f7b1d000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fb3f8986000)
        libattr.so.1 => /lib64/libattr.so.1 (0x00007fb3f7918000)
        libpthread.so.0 => /lib64/libpthread.so.0 (0x00007fb3f76fc000)
```
如果没找到依赖的 dynamic libraries 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[What are .a and .so Files? | Baeldung on Linux](https://www.baeldung.com/linux/a-so-extension-files)
[^2]:[Library (computing) - Wikipedia](https://en.wikipedia.org/wiki/Library_(computing))
[^3]:[Static library - Wikipedia](https://en.wikipedia.org/wiki/Static_library)
[^4]:[Shared library - Wikipedia](https://en.wikipedia.org/wiki/Shared_library)
[^5]:[Linux Basics: Static Libraries vs. Dynamic Libraries | by Erika Caoili | The Startup | Medium](https://medium.com/swlh/linux-basics-static-libraries-vs-dynamic-libraries-a7bcf8157779)
[^6]:[What is difference between Dynamic and Static library(Static and Dynamic linking) - YouTube](https://www.youtube.com/watch?v=eW5he5uFBNM)
[^7]:[C standard library - Wikipedia](https://en.wikipedia.org/wiki/C_standard_library)
