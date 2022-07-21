# CPU

ref:

https://en.wikipedia.org/wiki/Central_processing_unit

https://stackoverflow.com/questions/40163095/what-is-socket-core-threads-cpu

https://cloud.tencent.com/developer/article/1736628

https://blogs.vmware.com/customer-experience-and-success/2021/06/sockets-cpus-cores-and-threads-the-history-of-x86-processor-terms.html

https://virtual-dba.com/blog/sockets-cores-and-threads/

https://www.digitaltrends.com/computing/how-to-overclock-your-cpu/

## Digest

Central processsing unit (CPU)，中文通常翻译成中央处理器。主要负责有应用调用的的 basic arithemtic, logic, controlling, and input/ouput(I/O) 。需要和GPU做区别，GPU 通常用只做 I/O 和 main memory。

关于CPU有如下几个特殊的Terms

## Terms

### Socket

> numa node is not related to the CPU. A numa node is the relationship between the CPU socket adn the closet memory banks

At the most basic level, there is a motherboard that can do nothing without a CPU chip with pins that are inserted into the socket. The more correct term is CPU socket

区别 OSI socket 套接字，CPU socket 是一个物理的固件。指的是 CPU 安装的卡槽，常见的个人电脑通常一块主板只有一个Socket，而商业服务器可能有几十个 Socket。大概长这样

![img](https://ask.qcloudimg.com/http-save/yehe-5449215/4nnikthwwa.jpeg?imageView2/2/w/1620)

Socket 和 CPU 是 1 对 1 的关系，1 个 Socket 只能安装 1 块 CPU

### Core

历史上传统的CPU通常都是一个完整的物理处理单元，但是由于多核技术的发展，CPU渐渐的转向了容器的概念（think of this a two CPUs sharing the same socket）。而 Core 这变成真正的物理处理单元。一个CPU中可以有多个 Core，各个 Core 之前相互独立，可以并行处理逻辑。每个 Core 都有自己的独立的寄存器，L1/L2 cache 等物理硬件

![img](https://ask.qcloudimg.com/http-save/yehe-5449215/wa0tx39ep9.jpeg?imageView2/2/w/1620)

### Thread

CPU通常处理速度很快，但是磁盘可能I/O处理很慢。为了充分利用CPU的资源，在Core的基础上提出了Hyper-Threading(HT) 的概念。即一个 core 里可以模拟多个逻辑核，而这个逻辑核就表示 Thread

![img](https://ask.qcloudimg.com/http-save/yehe-5449215/o3drphu5c1.jpeg?imageView2/2/w/1620)

假设一个 10 Core 的 CPU 开启了 HT 的功能(通常是 2 线程)。那么就等于 10 Cores 20 Threads

这么看HT是一个好东西啊，所有的 CPU 应该能都应该开启啊。事实上相反，有些管理者会关闭 HT。因为 HT 可能会导致默写古早的应用(当年的CPU还不支持HT)异常

### Cache

A CPU cache is a hardware cache used by the CPU of a computer to redue the average cost (time or energy) to access data from the main memory.

简单来说，CPU缓存就是加快内存读取速度的。当然越快越好。通常 CPU Cache 有加快指令集和数据 的 Cache。且按照L1,L2,L3 ... 来命名 Cache。

后来由于技术的发展 Cache L1 裂变成 L1i(for instruction) 和 L1d (for data)，Cache L2 保持不变，作为 Cache L1 的候补

### Clock rate

这东西中文一般叫时脉频率，最小值(base clock)叫基频，最大值(turbo clock)叫___

you CPU processes many instructions from different programs every second. The clock speed measure the number of cycles your CPU executes per second, measured in GHz（集显CPU一般以MHz为单位）. A CPU with a clock speed of 3.2 GHz executes 3.2 billion cycles per second

the faster the clock, the more instructions and, consequently, the faster the clock, the more instructions the CPU will execute each second

overclocking（超频），用来提升 CPU 的 clock rate 加快 CPU 处理指令的速度，通常可以通过BIOS来或者CPU厂商推出的管理工具来设置超频。同时也有副作用，会导致 CPU 使用寿命减短（CPU 基础功率上升导致温度身高）。另外一般的集显笔记本的U都不支持超频。

## lscpu

在 linux 中可以通过 `lscpu`或者`lshw -class cpu` 来查看 CPU 相关的性能

```
cpl in ~ λ lscpu
Architecture:            x86_64
  CPU op-mode(s):        32-bit, 64-bit
  Address sizes:         48 bits physical, 48 bits virtual
  Byte Order:            Little Endian
CPU(s):                  16
  On-line CPU(s) list:   0-15
Vendor ID:               AuthenticAMD
  Model name:            AMD Ryzen 7 5800H with Radeon Graphics
    CPU family:          25
    Model:               80
    Thread(s) per core:  2
    Core(s) per socket:  8
    Socket(s):           1
...
    CPU max MHz:         4462.5000
    CPU min MHz:         1200.0000
...
Virtualization features: 
  Virtualization:        AMD-V
...
Caches (sum of all):     
  L1d:                   256 KiB (8 instances)
  L1i:                   256 KiB (8 instances)
  L2:                    4 MiB (8 instances)
  L3:                    16 MiB (1 instance)
```

从上面可以看出 1 socket，8 cores，2 threads per core，一共有 3 级缓存，支持 AMD VT，基频 1200 MHz。这样就可以计算出R7 5800 有 16 个逻辑 CPU。计算公式如下

$socket \times cores \times thread = Logical CPU$

当然逻辑CPU数字越大，性能越强。另外提一嘴，在 Linux 中还可以通过`cat /proc/cpuinfno`来查看 CPU 信息，其中的 max(processor) 的值就是 LogicalCPU 的值