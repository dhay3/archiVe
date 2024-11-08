# Linux 磁盘管理： dd

> dd和rufus同样都可以制作live USB，但是rufus会自动识别drive。

## Digest

用于生成该指定要求生成文件，可以与netcat实现文件备份与backdoor脚本

## Optional args

- `if=`

  使用文件替代标准输入流

- `of=`

  将标准输出流以文件的形式存储

- `bs=`

  指定每次读写块的大小

- `count=`

  指定拷贝多少bs*count的内容

- `seek`

  跳过多少块后才真正写入内容，bs*seek

```
root in /opt λ dd if=/dev/zero of=t1 bs=1G count=0 seek=1000
0+0 records in
0+0 records out
0 bytes copied, 7.8543e-05 s, 0.0 kB/s                                                                                                       /0.0s
root in /opt λ ll
.rw-r--r-- root root  1000 GB Wed Jan 27 07:08:12 2021  t1 
```

- `status=progress`

  显示dd的进度

## Examples

1. 使用`/dev/zero`创建空文件

```
root in /opt λ dd if=/dev/zero of=test bs=1M count=10
10+0 records in
10+0 records out
10485760 bytes (10 MB, 10 MiB) copied, 0.0435582 s, 241 MB/s                                                         /0.1s
root in /opt λ ll
.rw-r--r-- root root  10 MB Wed Jan 27 13:36:57 2021  test             
```

2. 创建假文件，不占有实际空间

```
root in /opt λ dd if=/dev/zero of=t1 bs=1G count=0 seek=1000
0+0 records in
0+0 records out
0 bytes copied, 7.8543e-05 s, 0.0 kB/s       /0.0s
root in /opt λ ll
.rw-r--r-- root root  1000 GB Wed Jan 27 07:08:12 2021  t1
```

