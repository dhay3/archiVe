# Linux gcc

ref：

http://c.biancheng.net/gcc/

## Digest

> gcc会自动根据文件的后缀名来编译文件
>
> 例如：gcc test.c ，会自动将文件编译

GUN project C and C++ compiler

## Examples

```
gcc --version 查看gcc版本
gcc <file> 编译自动生成a.out, 类似于javac
gcc -o <out-file> <src-file> 指定编译生成的文件，文件名可以任意
gcc -S <file> 生成指定文件的汇编文件，类似于javap <filename>
```

