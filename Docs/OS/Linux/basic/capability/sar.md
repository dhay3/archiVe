# Linux sar

参考：

https://shockerli.net/post/linux-tool-sar/

## 概述

System Activity Reporter (sar)系统活动情况报告，是目前 Linux 上最为全面的系统性能分析工具之一。

可以有效的找出系统的瓶颈，使用`systemctl status sysstat.service`来查看sar的状态

pattern：`sar [option] [interval] [times]`

每隔interval执行一次，一共执行times次

```
root in /var/log/sysstat λ systemctl status sysstat.service
● sysstat.service - Resets System Activity Data Collector
   Loaded: loaded (/lib/systemd/system/sysstat.service; enabled; vendor preset: enable
   Active: active (exited) since Wed 2021-01-20 17:28:35 HKT; 8min ago
     Docs: man:sa1(8)
           man:sadc(8)
           man:sar(1)
  Process: 5532 ExecStart=/usr/lib/sysstat/debian-sa1 --boot (code=exited, status=0/SU
 Main PID: 5532 (code=exited, status=0/SUCCESS)

Jan 20 17:28:35 ubuntu18.04 systemd[1]: Stopped Resets System Activity Data Collector.
Jan 20 17:28:35 ubuntu18.04 systemd[1]: Stopping Resets System Activity Data Collector
Jan 20 17:28:35 ubuntu18.04 systemd[1]: Starting Resets System Activity Data Collector
Jan 20 17:28:35 ubuntu18.04 systemd[1]: Started Resets System Activity Data Collector.
```

## 使用

- 持久化

  将收集的内容输入到`/var/log/sysstat`中，默认只标准输出

  ```
  root in /var/log/sysstat λ sar -o 1 5
  ```

- 内存统计，ram

  ```
  root in /var/log/sysstat λ sar -r 1 5
  Linux 4.15.0-124-generic (ubuntu18.04)  01/20/2021      _x86_64_        (2 CPU)
  
  05:46:08 PM kbmemfree   kbavail kbmemused  %memused kbbuffers  kbcached  kbcommit   %commit  kbactive   kbinact   kbdirty
  05:46:09 PM   3596180   3691248    442972     10.97     44972    235288   1071768     26.53    221288    128268         8
  ```

- cpu使用率

  ==不带任何参数默认统计cpu==

  ```
  root in /var/log/sysstat λ sar -u 1 1
  Linux 4.15.0-124-generic (ubuntu18.04)  01/20/2021      _x86_64_        (2 CPU)
  
  05:48:15 PM     CPU     %user     %nice   %system   %iowait    %steal     %idle
  05:48:16 PM     all      1.00      0.00      0.00      0.00      0.00     99.00
  Average:        all      1.00      0.00      0.00      0.00      0.00     99.00                                 
  ```

  - %user：用于表示用户模式下消耗的 CPU 时间的比例；
  - %nice：通过 nice 改变了进程调度优先级的进程，在用户模式下消耗的 CPU 时间的比例；
  - %system：系统模式下消耗的 CPU 时间的比例；
  - %iowait：CPU 等待磁盘 I/O 导致空闲状态消耗的时间比例；
  - %steal：利用 Xen 等操作系统虚拟化技术，等待其它虚拟 CPU 计算占用的时间比例；
  - %idle：CPU 空闲时间比例。

- 磁盘IO

  ```
  root in /var/log/sysstat λ sar -b 1 1
  Linux 4.15.0-124-generic (ubuntu18.04)  01/20/2021      _x86_64_        (2 CPU)
  
  05:51:55 PM       tps      rtps      wtps   bread/s   bwrtn/s
  05:51:56 PM      0.00      0.00      0.00      0.00      0.00
  Average:         0.00      0.00      0.00      0.00      0.00                 /1.0s
  ```

  - tps：每秒从物理磁盘 I/O 的次数。注意，多个逻辑请求会被合并为一个 I/O 磁盘请求，一次传输的大小是不确定的；
  - rd_sec/s：每秒读扇区的次数；
  - wr_sec/s：每秒写扇区的次数；
  - avgrq-sz：平均每次设备 I/O 操作的数据大小（扇区）；
  - avgqu-sz：磁盘请求队列的平均长度；
  - await：从请求磁盘操作到系统完成处理，每次请求的平均消耗时间，包括请求队列等待时间，单位是毫秒（1 秒=1000 毫秒）；
  - svctm：系统处理每次请求的平均时间，不包括在请求队列中消耗的时间；
  - %util：I/O 请求占 CPU 的百分比，比率越大，说明越饱和。
  
- 网络流量统计

  需要修改配置文件`/etc/default/sysstat`然后重启服务`systemctl restart sysstat`

  ```
  root in /opt λ sar -n DEV 1 3
  Linux 5.7.0-kali1-amd64 (cyberpelican) 	02/05/2021 	_x86_64_	(4 CPU)
  
  01:58:51 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
  01:58:52 PM        lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
  01:58:52 PM      eth0      1.00      1.00      0.10      0.10      0.00      0.00      0.00      0.00
  01:58:52 PM      eth1      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
  01:58:52 PM   docker0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
  ```

  1. IFACE: Name of the network interface for which statistics are reported.
  2. rxpck/s: packet receiving rate (unit: packets/second)
  3. txpck/s: packet transmitting rate (unit: packets/second)
  4. rxkB/s: data receiving rate (unit: Kbytes/second)
  5. txkB/s: data transmitting rate (unit: Kbytes/second)
  6. rxcmp/s: compressed packets receiving rate (unit: Kbytes/second)
  7. txcmp/s: compressed packets transmitting rate (unit: Kbytes/second)
  8. rxmcst/s: multicast packets receiving rate (unit: Kbytes/second)

