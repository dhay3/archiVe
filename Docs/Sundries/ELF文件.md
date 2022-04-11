# ELF文件

参考：

https://blog.csdn.net/xuehuafeiwu123/article/details/72963229

Executable and Linkable Format 简称ELF，是一种对象文件格式，==做为Unix应用程序二进制接口而开发的==。用于定义不同类型的对象文件(object files)中放了什么东西。由4部分组成，分别是ELF头（ELF header）、程序头表（Program header table）、节（Section）和节头表（Section header table）。实际上，一个文件中不一定包含全部内容，而且它们的位置也未必如同所示这样安排，只有ELF头的位置是固定的，其余各部分的位置、大小等信息由ELF头中的各项值来决定。

## 格式分析

可以使用linux中`readelf`命令来查看

### 源文件

```
root in /usr/local/script λ gcc readelf_t.c -o readelf_t
root in /usr/local/script λ cat readelf_t.c
#include <stdio.h>
int main(void)
{
printf("hello world");
return 0;
}
```

### ELF header

```
root in /usr/local/script λ readelf -h readelf_t
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              DYN (Shared object file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x540
  Start of program headers:          64 (bytes into file)
  Start of section headers:          6448 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         9
  Size of section headers:           64 (bytes)
  Number of section headers:         29
  Section header string table index: 28
```

