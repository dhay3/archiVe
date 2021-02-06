# Linux vmstat

参考：

https://www.cnblogs.com/ggjucheng/archive/2012/01/05/2312625.html

## 概述

vmstat和sar类似，用于在指定的时间展示服务器的状态值，包括cpu，ram，vram，io等。以byte展示

pattern：`vmstat [options] [interval [count]]`

```
 in / λ vmstat 
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 1  0      0 1972364 111424 1064988    0    0     2     3   28   57  0  0 100  0  0
```

1. r表示运行队列，1表示有多少进程分配到了cpu
2. b表示阻塞队列

3. swpd交换分区使用的大小，如果大于0就表示内存不够使用了
4. free表示空闲的物理内存
5. buff缓冲区使用大小
6. cache文件io使用的缓冲区
7. si从交换分区中每秒读入的大小
8. so从交换分区中每秒输出的大小、
9. bi块设备每秒接受的块数量
10. bo块设备每秒发送的块数量
11. in每秒cpu切换任务的次数
12. cs每秒上下文切换的次数
13. us用户调用cpu的时间
14. sy系统调用cpu的时间
15. id cpu空闲时间
16. wa等待io，cpu时间

## 参数

- `-a`

  将正在使用的内存和不正在使用的内存显示

  ```
  in / λ vmstat -a 1 3
  procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
   r  b   swpd   free  inact active   si   so    bi    bo   in   cs us sy id wa st
   2  0      0 1971052 608792 1051448    0    0     2     3   28   57  0  0 100  0  0
   0  0      0 1971052 608792 1051532    0    0     0     0  196  550  0  1 100  0  0
   0  0      0 1971052 608792 1051532    0    0     0    32  150  305  0  0 100  0  0      
  ```

- `-d`

  展示物理磁盘的使用情况(不显示分区)

  ```
  root in / λ vmstat -d
  disk- ------------reads------------ ------------writes----------- -----IO------
         total merged sectors      ms  total merged sectors      ms    cur    sec
  sr0       10      0       4       2      0      0       0       0      0      0
  sda    17802  11109 1665694  485144 188716  90271 2531176  200855      0    331 
  ```

- `-p <partition>`

  展示指定磁盘分区的使用情况

  ```
  root in / λ vmstat -p sda1 
  sda1            reads      read sectors      writes  requested writes
                  17623           1657178      188887           2533616                                       /0.0s
  ```

- `-w`

  格式化输出

  ```
  root in / λ vmstat -w  
  procs -----------------------memory---------------------- ---swap-- -----io---- -system-- --------cpu--------
   r  b         swpd         free         buff        cache   si   so    bi    bo   in   cs  us  sy  id  wa  st
   0  0            0      1943444       112780      1067692    0    0     2     3   28   57   0   0 100   0   0
  ```

## 案例



