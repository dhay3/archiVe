# Linux sar

ref:
[https://shockerli.net/post/linux-tool-sar/](https://shockerli.net/post/linux-tool-sar/)

## Digest
System Activity Reporter(sar)，是 linux 中的一款性能分析工具。
下载：`apt install sysstat`
启用服务：`systemctl start sysstat`
包含如下软件

```
ubuntu@VM-12-16-ubuntu:/etc/sysstat$ apt info sysstat
...
Description: system performance tools for Linux
 The sysstat package contains the following system performance tools:
  - sar: collects and reports system activity information;
  - iostat: reports CPU utilization and disk I/O statistics;
  - tapestat: reports statistics for tapes connected to the system;
  - mpstat: reports global and per-processor statistics;
  - pidstat: reports statistics for Linux tasks (processes);
  - sadf: displays data collected by sar in various formats;
  - cifsiostat: reports I/O statistics for CIFS filesystems.
...
```

## EBNF
syntax：`sar [options] INTERVAL COUNT`
```
sar [ -A ] [ -B ] [ -b ] [ -C ] [ -d ] [ -F ] [ -H ] [ -h ] [ -p ] [ -q ] [ -R ] [ -r ]
[ -S ] [ -t ] [ -u [ ALL ] ] [ -V ] [ -v ] [ -W ] [ -w ] [ -y ] [ -I { int [,...] | SUM
|  ALL  |  XALL  } ] [ -P { cpu [,...] | ALL } ] [ -m { keyword [,...] | ALL } ] [ -n {
keyword [,...] | ALL } ] [ -j { ID | LABEL | PATH | UUID | ... } ] [ -f [ filename ]  |
-o  [  filename ] | -[0-9]+ ] [ -i interval ] [ -s [ hh:mm:ss ] ] [ -e [ hh:mm:ss ] ] [
interval [ count ] ]
```
如果没有指定任何系统指标参数，默认查看CPU，即`-u`
## File

- /var/log/sa/sadd

  默认会从 `/var/log/sa/sadd`(大多distro都是这个文件)中读取可以从 kernel 获取的信息，其中的dd表示当前的日期

- /proc

  探测的系统指标源文件

## Postional args

- INTERVAL

  sar探测的时间间隔，如果该值为0，显示从系统开机时到现在的平均静态统计

  ```
  λ ~/ sar -n DEV 0
  Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
  
  11:06:32 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
  11:06:32 PM        lo      0.46      0.46      0.02      0.02      0.00      0.00      0.00
  11:06:32 PM      eth0      4.27      8.20      0.59      2.53      0.00      0.00      0.00
  ```

  如果该值非0，且没有指定 COUNT，会以 INTERVAL 间隔一直探测

- COUNT

  sar探测的次数，如果没有指定改值，表示以 INTERVAL 的间隔一直探测

## Optional args
### Gernal args

- `-A`

  等价于-bBdFHqrRSuvwWy -I SUM -I XALL -m ALL -n ALL -u ALL -P ALL

- `-f [filename]`

  从指定的文件的中读取sar所需的信息

  ```
  $sar -f sa24
  Linux 5.10.60-002.ali5000.alios7.x86_64 (first-line-dev011122132160.ea134)      11/24/2021      _x86_64_        (8 CPU)
  
  05:28:05 PM       LINUX RESTART
  ```

- `-o [filename]`

  将输出内容以binary的公式存储在指定文件中，如果没有指定filename，默认`/var/log/sa/sadd`，其中的dd表示当前的日期

  ```
  $sudo sar -o  -n DEV 1 1
  Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/24/2022      _x86_64_        (8 CPU)
  
  02:44:32 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
  02:44:33 PM        lo      9.00      9.00      4.31      4.31      0.00      0.00      0.00
  02:44:33 PM      eth0      8.00      5.00      0.53      0.35      0.00      0.00      0.00
  
  Average:        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
  Average:           lo      9.00      9.00      4.31      4.31      0.00      0.00      0.00
  Average:         eth0      8.00      5.00      0.53      0.35      0.00      0.00      0.00
  
  $ll
  total 64
  -rw-r--r-- 1 root root   656 Apr  1 15:25 sa01
  -rw-r--r-- 1 root root   704 May 20 22:03 sa20
  -rw-r--r-- 1 root root 53152 May 24 14:44 sa24
  -rw-r--r-- 1 root root   704 Jan 27 15:12 sa27
  ```

- `-p`

  pretty-print device names, sar 默认以`dev m-n`来输出设备名，m表示major number, n 表示 minor number。如果使用了改参数会从`/etc/sysconfig/sysstat.ioconf`中找到对应的映射然后显示

  ```
  λ ~/ sar -d 0 
  Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/24/2022      _x86_64_        (8 CPU)
  
  01:48:07 PM       DEV       tps  rd_sec/s  wr_sec/s  avgrq-sz  avgqu-sz     await     svctm     %util
  01:48:07 PM  dev253-0     13.49      5.39    169.23     12.94      0.03      2.24      0.08      0.11
  λ ~/ sar -dp 0 
  Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/24/2022      _x86_64_        (8 CPU)
  
  01:48:11 PM       DEV       tps  rd_sec/s  wr_sec/s  avgrq-sz  avgqu-sz     await     svctm     %util
  01:48:11 PM       vda     13.49      5.39    169.23     12.94      0.03      2.24      0.08      0.11
  ```

- `-s [hh:mm:ss] | -e [hh:mm:ss]`

  sar 测试开始和结束的时间，必须和`-f`一起使用

### Disk args

- `-B`

  显示paging statistics

  ```
  λ ~/ sar -B 0
  
  10:18:36 PM  pgpgin/s pgpgout/s   fault/s  majflt/s  pgfree/s pgscank/s pgscand/s pgsteal/s    %vmeff
  10:18:36 PM    712.90     88.47   4452.86      2.85   7443.58      0.00      0.00      0.00      0.00
  ```

  包含如下内容

     - `pgpgin/s`

       total number of kilobytes the system paged in from disk per second

     - `pgpout/s`

       total number of kilobytes the system paged out to disk per second

     - `faults/s`

       number of page faults (majar + minor) made by the system per second. This is not a count of page faults that generate I/O, because some page faults can be resolved without I/O

     - `majflt/s`

       Number of major faults the system has made per second, those  which  have required loading a memory page from disk.

     - `pgfree/s`

       Number of pages placed on the free list by the system per second.

     - `pgscank/s`

       Number of pages scanned by the kswapd daemon per second.

     - `pgscand/s`

       Number of pages scanned directly per second.

     - `pgsteal/s`

       Number  of pages the system has reclaimed from cache (pagecache and swap‐cache) per second to satisfy its memory demands.

     - `%vmeff`

       Calculated as pgsteal / pgscan, this is a metric  of  the  efficiency  of page  reclaim.  If  it is near 100% then almost every page coming off the tail of the inactive list is being reaped. If it gets too low (e.g.  less than  30%) then the virtual memory is having some difficulty.  This field is displayed as zero if no pages have been scanned during the interval of time.

- `-b`

  report I/O and transfer rate statistics
  磁盘IO统计(总的)，如果需要查看各个磁盘的统计，使用`-d`

  ```
  root in /var/log/sysstat λ sar -b 1 1
  Linux 4.15.0-124-generic (ubuntu18.04)  01/20/2021      _x86_64_        (2 CPU)
  
  05:51:55 PM       tps      rtps      wtps   bread/s   bwrtn/s
  05:51:56 PM      0.00      0.00      0.00      0.00      0.00
  Average:         0.00      0.00      0.00      0.00      0.00                 /1.0s
  ```

     - tps

       total number of transfer per second that were issued to physical devices
       每秒从物理磁盘 I/O 的次数。注意，多个逻辑请求会被合并为一个 I/O 磁盘请求，一次传输的大小是不确定的；

     - rtps

       total number of read requests per second issued to physical devices
       每秒读IO的请求次数；

     - wtps

       total number of write requests per second issued to physical devices

       每秒写IO的请求次数

     - bread/s

       total amount of data read from the devices in blocks per second
       每秒读扇区的次数

     - bwrtn/s

       total amount of data written to devices in blocks per second
       每秒写扇区的次数

- `-d`

  查看每个块设备的IO情况，如果需要打印块设备对应的逻辑名需要使用`-p`参数

  ```
  10:39:17 PM  dev253-0     16.39    610.59    153.34     46.62      0.03      1.68      0.19      0.31
  λ ~/ sar -pd 0
  Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
  
  10:40:53 PM       DEV       tps  rd_sec/s  wr_sec/s  avgrq-sz  avgqu-sz     await     svctm     %util
  10:40:53 PM       vda     16.10    584.65    152.42     45.78      0.03      1.70      0.19      0.30
  ```

     - tps

       每秒从物理磁盘 I/O 的次数。注意，多个逻辑请求会被合并为一个 I/O 磁盘请求，一次传输的大小是不确定的；

     - rd_sec/s

       每秒读扇区的次数；

     - wr_sec/s

       每秒写扇区的次数；

     - avgrq-sz

       平均每次设备 I/O 操作的数据大小（扇区）；

     - avgqu-sz

       磁盘请求队列的平均长度；

     - await

       从请求磁盘操作到系统完成处理，每次请求的平均消耗时间，包括请求队列等待时间，单位是毫秒（1 秒=1000 毫秒）；

     - svctm

       系统处理每次请求的平均时间，不包括在请求队列中消耗的时间；

     - %util

       I/O 请求占 CPU 的百分比，比率越大，说明越饱和。

- `-H`

  显示 hugepage 使用的情况

### Filesystem args

- `-F`

  显示当前系统挂载的filesystem统计，对比 df

  ```
  Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
  
  10:46:23 PM  MBfsfree  MBfsused   %fsused  %ufsused     Ifree     Iused    %Iused FILESYSTEM
  10:46:23 PM     43683     15655     26.38     31.50   3680314    186310      4.82 /dev/vda2
  10:46:23 PM       832       144     14.78     21.67     65188       348      0.53 /dev/vda1
  ```

     - MBfsfree

       total amount of free space in megabytes 

     - MBfsued

       total amount of space used in megabytes

     - %fsused

       percentage of filesystem space used, as seen by a privileged user

     - %ufsused 

       percentage of filesystem space used, as seen by an unpriviledged user

     - Ifree

       totol number of free file nodes in filesystem

     - Iused

       total number of file nodes used in filesystem

     - %Iused

       percentage of file nodes used in filesystem

- `-v`

  显示inode使用的情况

### Network args

- `-n {keyword[,...]|ALL}`

  keyword 可以是 DEV, EDEV, NFS, NFSD, SOCK, IP, EIP, ICMP, EICMP, TCP, ETCP, UDP, SOCKS6, IP6, EIP6, ICMP6, EICMP6, UDP6。这里只记录几个 keyword，具体的可以查看man page。==E表示error==

     - DEV

       如果指定的是DEV，显示所有网络设备(iface)的信息，包括以下几个属性

       ```
       λ ~/ sar -n DEV 0
       Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
       
       11:06:32 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
       11:06:32 PM        lo      0.46      0.46      0.02      0.02      0.00      0.00      0.00
       11:06:32 PM      eth0      4.27      8.20      0.59      2.53      0.00      0.00      0.00
       ```

       1. IFACE

          name of the network interface for which statistics are reported

       2. rxpck/s

          total number of packets received per second

       3. txpck/s

          total number of packets transmitted per second

       4. rxKB/s

          total number of kilobytes received per second

       5. txKB/s

          total number of kilobytes transmitted per second

       6. rxcmp/s

          number of compressed packets received per second 

       7. txcmp/s

          number of compressed packets transmitted per second

       8. rxmcst/s

          number of multicast packets received per second

     - EDEV

       如果指定的是EDEV，显示所网络设备的错包和丢包信息

       ```
       λ ~/ sar -n EDEV 1
       Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
       
       11:16:47 PM     IFACE   rxerr/s   txerr/s    coll/s  rxdrop/s  txdrop/s  txcarr/s  rxfram/s  rxfifo/s  txfifo/s
       11:16:48 PM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
       11:16:48 PM      eth0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
       ```

       1. IFACE

          Name of the network interface for which statistics are reported.

       2. rxerr/s

          Total number of bad packets received per second.

       3. txerr/s

          Total  number of errors that happened per second while transmitting packets.

       4. coll/s

          Number of collisions that happened per second while transmitting packets.

       5. rxdrop/s

          Number of received packets dropped per second because of a lack of  space in linux buffers.

       6. txdrop/s

          Number  of  transmitted  packets  dropped per second because of a lack of space in linux buffers.

       7. txcarr/s

          Number of carrier-errors that  happened  per  second  while  transmitting packets.

       8. rxfram/s

          Number  of  frame  alignment  errors that happened per second on received packets.

       9. rxfifo/s

          Number of FIFO overrun errors that happened per second on received  packets.

       10. txfifo/s

           Number  of  FIFO  overrun  errors that happened per second on transmitted packets.

     - SOCK

       如果指定的是SOCK，会显示IPv4在使用的SOCK(两元组)

       ```
       λ ~/ sar -n SOCK 0   
       Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
       
       11:21:49 PM    totsck    tcpsck    udpsck    rawsck   ip-frag    tcp-tw
       11:21:49 PM        99        19         4         0         0         8
       ```

       1. totsck

          Total number of sockets used by the system.对比`ss -npa`

       2. tcpsck

          Number of TCP sockets currently in use. 只要TCP状态机制不是CLOSED就是在使用，可以对比`ss -lnpt`

       3. udpsck

          Number of UDP sockets currently in use.

       4. rawsck

          Number of RAW sockets currently in use.

       5. ip-frag

          Number of IP fragments currently in queue.队列中的IP分片

       6. tcp-tw

          Number of TCP sockets in TIME_WAIT state. 处于 TIME_WIAT状态的SOCK

     - TCP

       如果指定的是TCP，会显示和TCP相关的信息

       ```
       λ ~/  sar -n TCP 0
       Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
       
       11:29:16 PM  active/s passive/s    iseg/s    oseg/s
       11:29:16 PM      0.43      0.00      3.58      4.59
       ```

       1. active/s

          The  number of times TCP connections have made a direct transition to the SYN-SENT state from the CLOSED state per second [tcpActiveOpens].每秒从 CLOSED 到 SYN-SNT 的 TCP连接数，即主动发起连接(client 行为)

       2. passive/s

          The number of times TCP connections have made a direct transition to  the SYN-RCVD state from the LISTEN state per second [tcpPassiveOpens].每秒从 LISTEN 到 SYN-RCVD 的 TCP连接数，即被动连接(serve 行为)

       3. iseg/s

          The  total  number  of  segments  received  per  second,  ==including those received in error== [tcpInSegs].  This count includes segments received  on currently established connections.

       4. oseg/s

          The  total number of segments sent per second, including those on currentconnections but ==excluding  those  containing  only  retransmitted  octets== [tcpOutSegs].

     - ETCP

       如果指定的是TCP，会显示和TCP错误的信息(重传)

       ```
       λ ~/  sar -n "ETCP" 0    
       Linux 5.10.60-002.ali5000.alios7.x86_64 (ali-test)      05/20/2022      _x86_64_        (8 CPU)
       
       11:36:03 PM  atmptf/s  estres/s retrans/s isegerr/s   orsts/s
       11:36:03 PM      0.21      0.00      0.13      0.00      0.87
       ```

       1. atmptf/s

          The number of times per second TCP connections have made a direct transition to the CLOSED state from either the SYN-SENT state or  the  SYN-RCVD state,  plus  the  number of times per second TCP connections have made a direct transition to the LISTEN state from  the  SYN-RCVD  state  [tcpAttemptFails].每秒从SYN-SENT或SYN-RCVD或LISTEN到CLOSED状态的连接，即每建连成功的

       2. estres/s

          The number of times per second TCP connections have made a direct transition to the CLOSED state from either the ESTABLISHED state or the  CLOSE-WAIT state [tcpEstabResets].每秒从ESTABLISHED 或 CLOSE-WAIT 到 CLOSED 的 TCP 连接

       3. retrans/s

          The total number of segments retransmitted per second - that is, the number of TCP segments transmitted containing one or more previously  transmitted octets [tcpRetransSegs].每秒重传的 TCP segments 个数

       4. isegerr/s

          The  total number of segments received in error (e.g., bad TCP checksums) per second [tcpInErrs].每秒收错包格式

       5. orsts/s

          The number of TCP segments  sent  per  second  containing  the  RST  flag [tcpOutRsts].每秒发RST包的次数

### Memormy args

- `-R`

  显示当前内存的计数统计

  ```
  λ ~/ sar -R 0
  Linux 5.10.60-002.ali5000.alios7.x86_64     05/24/2022      _x86_64_        (8 CPU)
  
  01:52:34 PM   frmpg/s   bufpg/s   campg/s
  01:52:34 PM      9.07      0.14      2.07
  ```

     - frmpg/s

       number of memory pages freed by the system per second. 每秒释放的内存页数，内存页大小基于机器的架构，通常为 4 KB 或 8 KB

     - bufpg/s

       number of additional memory pages used as buffers by the system per second. 每秒缓存的内存页数

     - campg/s

       number of addition memory pages cahec by the systemd per second.每秒缓存的内存页数

- `-r`

  显示内存使用的情况

  ```
  ubuntu@VM-12-16-ubuntu:~$ sar -r 0
  Linux 5.4.0-109-generic (VM-12-16-ubuntu)       05/24/2022      _x86_64_        (2 CPU)
  
  02:02:32 PM kbmemfree   kbavail kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit  kbactive   kbinact   kbdirty
  02:02:32 PM   1319036   1606652    208920     10.29     48444    359208   1091768     53.78    422828    141220        60
  ```

     - kbmemfree

       Amount of free memory available in kilobytes.

     - kbmemused

       Amount of used memory in kilobytes. This does not take into account memory used by the kernel itself.
       不包含内核使用的内存大小

     - %memused

       Percentage of used memory.

     - kbbuffers

       Amount of memory used as buffers by the kernel in kilobytes.

     - kbcached

       Amount of memory used to cache data by the kernel in kilobytes.

     - kbcommit

       Amount of memory in kilobytes needed for current workload. This is an estimate of how much RAM/swap is needed to guarantee that there never is out of memory.
       估算需要的内存大小，才不在引发OOM

     - %commit

       Percentage  of memory needed for current workload in relation to the total amount of memory (RAM+swap).  This number may be greater than 100% because the kernel usually overcommits memory.

     - kbactive

       Amount of active memory in kilobytes (memory that has been used more recently and usually not reclaimed unless absolutely necessary).

     - kbinact

       Amount of inactive memory in kilobytes (memory which has been less recently used. It is more eligible to be reclaimed for other purposes).

     - kbdirty

       Amount of memory in kilobytes waiting to get written back to the disk.

- `-S`

  显示Swap使用的情况

  ```
  ubuntu@VM-12-16-ubuntu:/etc/sysstat$ sar -u 0
  Linux 5.4.0-109-generic (VM-12-16-ubuntu)       05/24/2022      _x86_64_        (2 CPU)
  
  02:18:28 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
  02:18:28 PM     all      0.29      0.03      0.33      0.07      0.00     99.27
  ```

### CPU args

- `-u [ ALL ]`

  显示CPU使用的情况

  ```
  $sar -u 0
  Linux 5.10.60-002.ali5000.alios7.x86_64     05/24/2022      _x86_64_        (8 CPU)
  
  03:00:35 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
  03:00:35 PM     all      0.16      0.00      0.18      0.00      0.00     99.66
  ```

     - %user

       Percentage of CPU utilization that occurred while executing at the user level (application). Note that this field includes time spent running virtual processors.

     - %usr

       Percentage of CPU utilization that occurred while executing at the user level (application). Note that this field does NOT include time spent running virtual processors.

     - %nice

       Percentage of CPU utilization that occurred while executing at the user level with nice priority.

     - %system

       Percentage of CPU utilization that occurred while executing at the system level (kernel). Note that this field includes time spent servicing hardware and software interrupts.

     - %sys

       Percentage  of CPU utilization that occurred while executing at the system level (kernel). Note that this field does NOT include time spent servicing hardware or software interrupts.

     - %iowait

       Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.

     - %steal

       Percentage of time spent in involuntary wait by the virtual CPU or CPUs while the hypervisor was servicing another virtual processor.

     - %irq

       Percentage of time spent by the CPU or CPUs to service hardware interrupts.

     - %soft

       Percentage of time spent by the CPU or CPUs to service software interrupts.

     - %guest

       Percentage of time spent by the CPU or CPUs to run a virtual processor.

     - %gnice

       Percentage of time spent by the CPU or CPUs to run a niced guest.

     - %idle

       Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.

## Conf
sar 默认会读取`/etc/sysconfig/sysstat`配置文件(可能会因为distro不同，配置的文件的位置也不同)，主要涵盖以下几个指令块

- COMPESSAFTER

  在日志多少天后会被压缩，压缩的工具使用ZIP指令块

- HISTORY

  日志保留的时间，通常储存在`/var/log/sadd`，其中dd表示日期，如果超过改时间日志会被删除。通过mtime来界定

- SADC_OPTIONS

  指定`sadc`命令使用的参数，即 daily file 采集的数据。具体参考`sadc`

- ZIP

  压缩使用的文件

  ```
  # sysstat-10.1.5 configuration file.
  
  # How long to keep log files (in days).
  # If value is greater than 28, then log files are kept in
  # multiple directories, one for each month.
  HISTORY=28
  
  # Compress (using gzip or bzip2) sa and sar files older than (in days):
  COMPRESSAFTER=31
  
  # Parameters for the system activity data collector (see sadc manual page)
  # which are used for the generation of log files.
  SADC_OPTIONS="-S DISK"
  
  # Compression program to use.
  ZIP="bzip2"
  ```

## Example
可以组合使用同时显示，会按照优先级重新排序参数用于显示
```
ubuntu@VM-12-16-ubuntu:/etc/sysstat$ sar -r -n DEV 0
Linux 5.4.0-109-generic (VM-12-16-ubuntu)       05/24/2022      _x86_64_        (2 CPU)

02:21:32 PM kbmemfree   kbavail kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit  kbactive   kbinact   kbdirty
02:21:32 PM   1308680   1605292    209752     10.33     49884    366516   1093032     53.84    430144    144236        64

02:21:32 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
02:21:32 PM        lo      0.36      0.36      0.03      0.03      0.00      0.00      0.00      0.00
02:21:32 PM      eth0      4.93      4.75      0.92      0.74      0.00      0.00      0.00      0.00

ubuntu@VM-12-16-ubuntu:/etc/sysstat$ sar  -n DEV -n TCP  0
Linux 5.4.0-109-generic (VM-12-16-ubuntu)       05/24/2022      _x86_64_        (2 CPU)

02:22:07 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
02:22:07 PM        lo      0.36      0.36      0.03      0.03      0.00      0.00      0.00      0.00
02:22:07 PM      eth0      4.94      4.75      0.91      0.74      0.00      0.00      0.00      0.00

02:22:07 PM  active/s passive/s    iseg/s    oseg/s
02:22:07 PM      0.49      0.05      3.78      3.82
```
