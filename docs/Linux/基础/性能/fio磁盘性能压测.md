# fio磁盘性能压测

参考：

https://help.aliyun.com/document_detail/147897.html?spm=a2c4g.11186623.6.892.6de418cdpOtRdM

> 使用磁盘性能压测，需要先查看是否[4k对齐](..\..\..\Hardware\4k对齐.md)
>
> ==如果文件已存在，会对文件覆写==

syntax：`fio [options] [jobfile]`

fio有两种方式，通过cmdline和文件读取的方式来对磁盘性能进行压测

## 通用参数

- `--eta=<when>`

  显示eta进度，when取值always，never或则auto

- `--max-jobs=<nr>`

  设置的最大线程

- `--client=<hostname>`

  测试hostname的IO，而不是本地

- `--showcmd=<jobfile>`

  ==将jobfile转为cmdline options的形式==

- `--readonly`

  只执行读操作，不进行写操作。

  ```
  [root@chz opt]# fio --readonly -direct=1 -iodepth=128 -rw=randwrite -ioengine=libaio -bs=4k -size=1G -numjobs=1 -runtime=1000 -group_reporting -filename=iotest -name=Rand_Write_Testing
  fio: job <(null)> has write bit set, but fio is in read-only mode
  Correct format for offending option
                    rw: IO direction
                 alias: readwrite
                  type: string (opt=bla)
               default: read
          valid values: read       Sequential read
                      : write      Sequential write
                      : trim       Sequential trim
                      : randread   Random read
                      : randwrite  Random write
                      : randtrim   Random trim
                      : rw         Sequential read and write mix
                      : readwrite  Sequential read and write mix
                      : randrw     Random read and write mix
                      : trimwrite  Trim and write mix, trims preceding writes
  
  fio: failed parsing rw=randwrite
  ```

- `--warnings-fatal`

  所有warngins被fio认为是fatal，会直接导致fio退出

## jobfile参考

https://github.com/axboe/fio/tree/master/examples

job name 定义在`[]`中，注释以`;`或`#`开头

`[global]`表示全局变量，可以在当都的job中继承或重写全局变量

```
[root@chz opt]# cat jobfile 
[global]
size=10m
rw=read
directory=/tmp
[file1]
[file2]
[file3]

[root@chz opt]# fio jobfile
```

执行后会在`/tmp`目录下创建三个文件

```
[root@chz opt]# ls /tmp | grep  -v files | grep file
file1.0.0
file2.0.0
file3.0.0
```

## jobfile参数operation

- addition +
- substraction - 
- multiplication *
- division /
- modulus %
- exponentiation ^

## jobfile参数类型

当有多个值时，可以使用`:`分隔

- str
- time：带有单位的数字类型，默认s
- int：可选前缀(进制)和后缀(单位，校验数据使用bytes，校验时间使用s)
- bool：使用1和0代表true or false
- irange：integer range
- float_list

## target file/device

- directory=str

  指定生成文件的路径，而不是默认使用`./`，可以通过`:`分隔指定多个文件

- filename_format=str

  ==如果多个jobs共享多个files，需要格式化filename。默认fio使用jobname(线程或进程的名字).jobnumber(进程号或线程号).filenumber(进程或线程执行的任务数)。==

- filename=str

  fio默认以job name，thread number，file number做为filename，如果想要多个线程运行同一个fliename需要指定一个固定的filename。如果使用了该参数忽略nrfiles。如果使用了backslash或是colon需要使用` '转义。

  例如：

  ```
  filename=`/dev/sdb:/dev/sda'
  ```

  `-'表示filename从stdin和stdout获取

- lockflie=str

  io操作时使用的锁

  1. none无锁默认
  2. exclusive只有一个线程能操作
  3. readwrite读写锁

- nrfiles=int

  job使用files数量

- openfiles=int

  job一次能打开的file数量，默认与nrfiles相同

## job descritption

- name=str

  name of job(随意设置)，On the command line this parameter has the special purpose of also signaling the start of a new job.

- description=str

  job的描述，不会被解析

- loops=int

  重复运行这个job的次数，默认1

- numjobs=int

  jobs调用的线程数

## I/O depth

- iodepth=int

  如果时异步IO，表示同时发出的I/O units 

## I/O type

- buffered=bool

  默认true，使用buffered I/O模式测试

- direct=bool

  默认false，是否使用direct I/O。约束条件具体查看manual page

- atomic=bool

  atomic direct I/O，只有linux支持

- rw=str

  I/O模式，支持如下几个值

  1. read：顺序读，缺省值
  2. write：顺序写
  3. trim：顺序删除，只有linux block支持
  4. randread：随机读
  5. randwrite：随机写
  6. randtrim：随机删除
  7. rw：顺序读写混合
  8. randrw：随机读写混合
  9. trimwrite：先顺序删除然后再顺序写

## I/O engine

- ioengine=str

  定义如何读写

  1. sync：使用read()和write()函数
  2. libaio：linux原生的async I/O，必须和direct=1或是buffered=0一起使用
  3. windows：windows原生的async I/O，windows上默认使用这个值
  4. mmap：使用mmap()函数
  5. null：不进行数据传输，通常用于fio debug
  6. net：通过网络传输

## Block size

- `bs=int[,int][,int]`

  块操作时使用的最小单元，从左至右为reads，writes，trims。缺省值4096bytes

  bs=4k表示reads，writes，trims都为4k

## I/O size

- size=int

  每个线程测试的文件大小，通过filename或nrfiles决定。size可以是百分比的形式。

## Time related parameters

- runtime=time

  告诉fio在指定的time后终止所有进程

- time_based

  If set, fio will run for the duration of the runtime specified even if the file(s) are completely read or written. It will simply loop over the same workload as many times as the runtime allows.

- startdelay=irange

  file启动的延迟时间

## Measurement and reporting

- group_reporting

  ==如果使用numjobs，将不同的jobs以组的形式展示==

  ```
  [root@chz opt]# cat jobfile 
  [global]
  group_reporting
  size=10m
  rw=read
  directory=/tmp
  numjobs=2
  [file]
  -----
  file: (groupid=0, jobs=2): err= 0: pid=3376: Mon Mar 15 18:40:31 2021
  ```

## Output

```
Jobs: 20 (f=20): [R(10),W(10)][4.0%][r=20.5MiB/s,w=23.5MiB/s][r=82,w=94 IOPS][eta 57m:36s]
```

1. 20

   当前运行的线程数

2. (f=20)

   表示当前打开的文件数

3. [_(1),M(1)]

   RW表示线程的任务，10表示线程数

4. [4.0%]

   当前完成的进度

5. [r=20.5MiB/s,w=23.5MiB/s]

   当前读的速度和当前写的速度

6. [r=82,w=94 IOPS]

   按照IOPS的规定速率

7. [eta 57m:36s]

   任务跑完结束的大概时间

```
               Client1: (groupid=0, jobs=1): err= 0: pid=16109: Sat Jun 24 12:07:54 2017
                   write: IOPS=88, BW=623KiB/s (638kB/s)(30.4MiB/50032msec)
                     slat (nsec): min=500, max=145500, avg=8318.00, stdev=4781.50
                     clat (usec): min=170, max=78367, avg=4019.02, stdev=8293.31
                      lat (usec): min=174, max=78375, avg=4027.34, stdev=8291.79
                     clat percentiles (usec):
                      |  1.00th=[  302],  5.00th=[  326], 10.00th=[  343], 20.00th=[  363],
                      | 30.00th=[  392], 40.00th=[  404], 50.00th=[  416], 60.00th=[  445],
                      | 70.00th=[  816], 80.00th=[ 6718], 90.00th=[12911], 95.00th=[21627],
                      | 99.00th=[43779], 99.50th=[51643], 99.90th=[68682], 99.95th=[72877],
                      | 99.99th=[78119]
                    bw (  KiB/s): min=  532, max=  686, per=0.10%, avg=622.87, stdev=24.82, samples=  100
                    iops        : min=   76, max=   98, avg=88.98, stdev= 3.54, samples=  100
                   lat (usec)   : 250=0.04%, 500=64.11%, 750=4.81%, 1000=2.79%
                   lat (msec)   : 2=4.16%, 4=1.84%, 10=4.90%, 20=11.33%, 50=5.37%
                   lat (msec)   : 100=0.65%
                   cpu          : usr=0.27%, sys=0.18%, ctx=12072, majf=0, minf=21
                   IO depths    : 1=85.0%, 2=13.1%, 4=1.8%, 8=0.1%, 16=0.0%, 32=0.0%, >=64=0.0%
                      submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
                      complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
                      issued rwt: total=0,4450,0, short=0,0,0, dropped=0,0,0
                      latency   : target=0, window=0, percentile=100.00%, depth=8

```

1. Client1

   job的名字，当使用了group_reporting显示

2. (groupid=0, jobs=1)

   组号，一共的jobs

3. err=0

   error的数量

4.  write: IOPS=88, BW=623KiB/s (638kB/s)(30.4MiB/50032msec)

   ==平均值，IOPS表示每秒IO次数，BW表示带宽623(1KB=1024B)，638(1kB=1000B)==，30.4MiB表示IO以1024为进制，50032msec当前线程占用的微妙数

5. slat (nsec): min=500, max=145500, avg=8318.00, stdev=4781.50‘

   ==submission latency 提交IO的延迟==，fio会自动选择单位，nsec表示nano sec

6. clat (usec): min=170, max=78367, avg=4019.02, stdev=8293.31

   ==completion latency 完成IO的延迟==

7. lat (usec): min=174, max=78375, avg=4027.34, stdev=8291.79

   total latency = submission latency + completion latency / 2

8. bw (  KiB/s): min=  532, max=  686, per=0.10%, avg=622.87, stdev=24.82, samples=  100

   带宽的具体值

9.  iops        : min=   76, max=   98, avg=88.98, stdev= 3.54, samples=  100

   iops的具体值

10. cpu          : usr=0.27%, sys=0.18%, ctx=12072, majf=0, minf=21

    cpu使用的情况

11. IO depths    : 1=85.0%, 2=13.1%, 4=1.8%, 8=0.1%, 16=0.0%, 32=0.0%, >=64=0.0%

    job运行过程中发出I/O depths

13. submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%

14. complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%

15. issued rwts: total=0,262144,0,0 short=0,0,0,0 dropped=0,0,0,0

    read/write/trim请求的次数

## 例子

在使用前可以使用`-ioengine=null`来测试fio是否有问题， 不会做实际的IO操作

```
[root@cpl ~]# fio -direct=1 -iodepth=128 -rw=randwrite -ioengine=null -bs=4k -size=1G -numjobs=1 -runtime=1000 -group_reporting -filename=iotest -name=Rand_Write_Testing
```





















































