# Linux pv

参考:

https://www.tecmint.com/monitor-copy-backup-tar-progress-in-linux-using-pv-command/

syntax：`pv [option] [file]`

`pv`可以打印出进程的progress bar。`pv`会拷贝file到stdout，如果没有file就将stdin拷贝到stdout(和`cat`一样)，==然后将progress bar 输出到stderr==

## 参数

> pv默认使用`-pterb`

- -p | --progress

  显示progress bar，如果stdin不是文件或没有大小，progress bar不会显示进度，只会左右移动表示有数据在移动。

- -l | --line-mode

  只左右移动progress bar

- -t | --timer

  显示计时器

- -e | --eta

  eta计时器

- -r | --rate

  显示rate

- -b | --bytes

  显示已经传输的byte count

- -L rate

  传输的最大速率，默认bps

  ```
  pv -L 2 file1 > file2
  ```

- -d pid

  显示进程打开的file descriptor

  ```
  cpl in ~ λ pv -d 91158
     4:/proc/stat: 3.31KiB 0:00:44 [0.00 B/s] [<=>                           ] 
     5:/proc/uptime: 19.0 B 0:00:44 [0.00 B/s] [<=>                          ] 
     6:/proc/meminfo: 1.44KiB 0:00:44 [0.00 B/s] [<=>                        ] 
     7:/proc/loadavg: 28.0 B 0:00:44 [0.00 B/s] [<=>                         ] 
  ```

- -s size 

  假定传输文件的大小，默认bytes，可以使用suffix

## 例子

复制文件

```
➜  test pv file1 > file2
10.0MiB 0:00:00 [ 578MiB/s] [=============================================================================>] 100%
```

压缩文件的时候显示progress bar

```
#zip
cpl in /tmp/test λ pv file1 | zip file1.zip - 
  adding: -10.0MiB 0:00:00 [44.8MiB/s] [=============================================================================>] 100%            
 (deflated 0%)
 
 #tar, 压缩只后的内容重定向到file1.tar
 cpl in /tmp/test λ tar -cf - file1  | pv > file1.tar
10.0MiB 0:00:00 [1.36GiB/s] [ <=> 
```

使用netcat传输文件时，可以使用`pv`显示传输速率和进度

```
#sender,stdout的内容被发送到了reciever，然后输出stderr内容
cpl in ~ λ pv file | nc 82.157.1.137 10086
 100 B 0:00:00 [1.67MiB/s] [==============================================================================>] 100%
 
#reciever
root in /home/ubuntu λ nc -w 10 -lp 10086 > file
```

显示dd进度，dd没有指定`of`默认输出到stdout

```
cpl in /tmp/test λ dd if=/dev/zero bs=10M count=3 | pv | dd of=file3  
dd: warning: partial read (65536 bytes); suggest iflag=fullblock
0+3 records in
0+3 records out
142336 bytes (142 kB, 139 KiB) copied, 0.000314633 s, 452 MB/s
 143KiB 0:00:00 [95.3MiB/s] [ <=>  
```

