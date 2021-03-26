# Linux 磁盘管理： dd

## 概述

用于生成该指定要求生成文件，可以与netcat实现文件备份与backdoor脚本

## 参数

- `if=`

  使用文件替代标准输入流

- `of=`

  将标准输出流以文件的形式存储

- `bs=`

  指定每次读写块的大小

- `count=`

  指定拷贝多少bs*count的内容

- `seek`

  跳过多少块后才真正写入内容

```
root in /opt λ dd if=/dev/zero of=t1 bs=1G count=0 seek=1000
0+0 records in
0+0 records out
0 bytes copied, 7.8543e-05 s, 0.0 kB/s                                                                                                       /0.0s
root in /opt λ ll
.rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  bak.xml
drwxrwxrwx root root     4 KB Sun Sep 13 23:45:24 2020  burpsuite pro
drwx--x--x root root     4 KB Sat Sep 19 22:45:25 2020  containerd
.rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  ibak.xml
drwxr-xr-x root root     4 KB Thu Sep 10 07:28:47 2020  jdk-14.0.2
.rwxr--r-- root root 182.1 MB Thu Sep 10 07:28:10 2020  jdk-14.0.2_linux-x64_bin.tar.gz
.rw-r--r-- root root 661.5 KB Sat Dec 12 09:07:29 2020  lsd_0.18.0_amd64.deb
.rw-r--r-- root root  1000 GB Wed Jan 27 07:08:12 2021  t1
drwxr-xr-x root root     4 KB Thu Jan 21 06:42:41 2021  test
.rw-r--r-- root root    35 B  Thu Jan 21 06:43:05 2021  test.sh   
```

## 案例

1. 使用`/dev/zero`创建空文件

```
root in /opt λ dd if=/dev/zero of=test bs=1M count=10
10+0 records in
10+0 records out
10485760 bytes (10 MB, 10 MiB) copied, 0.0435582 s, 241 MB/s                                                         /0.1s
root in /opt λ ll
drwx--x--x root root   4 KB Thu Jan 21 10:58:32 2021  containerd
.rw-r--r-- root root 126 B  Fri Jan 22 14:00:51 2021  Dockerfile
.rwxr-xr-x root root   0 B  Mon Jan 25 19:49:41 2021  hostname
.rwxr-xr-x root root   0 B  Mon Jan 25 19:49:41 2021  hosts
drwxr-xr-x root root   4 KB Thu Dec 31 14:17:18 2020  lsd-0.18.0-x86_64-unknown-linux-gnu
.rwxr-xr-x root root   0 B  Mon Jan 25 19:49:41 2021  resolv.conf
.rw-r--r-- root root  10 MB Wed Jan 27 13:36:57 2021  test             
```

2. 创建假文件，不占有实际空间

```
root in /opt λ dd if=/dev/zero of=t1 bs=1G count=0 seek=1000
0+0 records in
0+0 records out
0 bytes copied, 7.8543e-05 s, 0.0 kB/s                                                                                                       /0.0s
root in /opt λ ll
.rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  bak.xml
drwxrwxrwx root root     4 KB Sun Sep 13 23:45:24 2020  burpsuite pro
drwx--x--x root root     4 KB Sat Sep 19 22:45:25 2020  containerd
.rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  ibak.xml
drwxr-xr-x root root     4 KB Thu Sep 10 07:28:47 2020  jdk-14.0.2
.rwxr--r-- root root 182.1 MB Thu Sep 10 07:28:10 2020  jdk-14.0.2_linux-x64_bin.tar.gz
.rw-r--r-- root root 661.5 KB Sat Dec 12 09:07:29 2020  lsd_0.18.0_amd64.deb
.rw-r--r-- root root  1000 GB Wed Jan 27 07:08:12 2021  t1
drwxr-xr-x root root     4 KB Thu Jan 21 06:42:41 2021  test
.rw-r--r-- root root    35 B  Thu Jan 21 06:43:05 2021  test.sh   
```

