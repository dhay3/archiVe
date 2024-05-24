# CPU

## Overview

> *在介绍 CPU 前，需要声明的一点是 CPU 这个名词的定义是灾难性的。你会在各种文章或者工具的定义中看到 CPU 等价与 Socket，Core 或者是 Thread。本文基于个人的理解，只为自己的逻辑做统一* 

Central Processsing Unit (CPU)，是电脑中最重要的部分之一(需要和 GPU 做区别, 传统意义上的 GPU 通常用只做图形处理。但是由于 CUDA 的出现也让 GPU 越来越趋近于 CPU 的逻辑)，负责执行应用发出的指令，包括有

1. arithemtic - 数字计算指令
2. logic/controlling - 逻辑控制指令
3. input/ouput(I/O) - IO 指令

随着技术的发展，CPU 在实现的方式上和概念出现不同。但是主要的组件保持不变，主要有

1. arithemtic logic unit(ALU) - 负责数字和逻辑运算，中文通常翻译为算术逻辑单元
2. processor registers - 负责存储 ALU 的结果，中文通常翻译为寄存器
3. control unit(CU) - 负责调度 ALU/processor registers, 存取 memory，decode, 以及指令地执行

### The Shifting Of CPU’s Meaning

在没有出现 Core 的概念前，CPU 指的是物理意义上的，你可以说下图的设备是一个 CPU

<img src="https://upload.wikimedia.org/wikipedia/commons/1/12/Intel_i9-14900K.webp" style="width:300px">

但是随着技术的发展，出现了 Core/Thread 的概念，CPU 从物理上的意义变成逻辑上的意义。整个物理设备也从 CPU 这一名词转变为 Processor(也有人将其称为 Chip，但是大多数人仍使用 CPU 这个称呼，==本文会以 Processor 命名这个物理设备==)

![img](https://media.licdn.com/dms/image/C5112AQG7WX-ECFUVAg/article-inline_image-shrink_400_744/0/1520148418058?e=1721865600&v=beta&t=p0JUXg6Y8QF9jZESTLkYH6EzQKsfdgEyOKlaAvN2Ang)

## Terminologies

和 CPU 关联的有如下几个特殊的 Terms

### Socket 

CPU Socket 是一个物理的固件，区别于 OSI Socket 套接字，指的是主板上安装 CPU 的卡槽。个人电脑通常一块主板只有一个 Socket，而商业服务器可能有几十个 Socket

Socket 大概长这样

<img src="https://pic1.zhimg.com/80/v2-d11b887550690c62655490bfa38dabd4_1440w.webp" style="width:400px">

在没有 Core 的概念前，Socket 和 CPU 是 一对一 的关系，一个 Socket 只能有 一个 CPU，CPU 是物理上的

在出现 Core 的概念后，Socket 和 CPU 是 一对多 的关系，一个 Socket 可以有 多个 CPU，CPU 是逻辑上的

### Core

早期在没有 Core 的概念时，CPU 指的就是物理意义上的设备，同一时间只能运行一个应用发出的指令。随着需求的发展，厂商为了提升性能，引入 Core 的概念。这个阶段单个 Core 几乎等价于单个 CPU(可以将 Core 理解成是 CPU 的容器，整个物理设备也从 CPU 这一名词转变为 Processor)，单个 Processor 可以有多个 Core，同一时间可以让多条指令分发到不同的 Core

Processor 根据 Core 的数量可以分为 multi-core processor 和 single-core processor。在 multi-core processor 中，各个 Core 之前相互独立，可以并行处理逻辑，每个 Core 都有自己的独立 ALU,register,CU,cache 等组件。因为 Core 在功能上和原始含义的 CPU 相同，也被称为 logical CPU

Core 物理上是一组晶体管，大概长这样

<img src="https://www.techpowerup.com/img/15-08-18/77a.jpg" width="500px">

### Thread

随着技术的发展出现了 Hyper-Threading(HT), 单个 Core 可以被划分成 2 个逻辑上的 Core(也被称为 Thread)。这些逻辑上的 Core 也有自己的独立 ALU,register,CU,cache 等组件。因为这些逻辑上的 Core 同样也被称为 logical CPU

假设一个有 10 个 Core 的 Processor 开启了 HT 的功能(可以通过 `sudo dmidecode -t processor` 中 flag 是否包含 HTT 来判断，现在的 CPU 默认会开启)，那么就有 20 个逻辑上的 Core，也就是有 20 个 logical CPUs

Thread 是逻辑上的概念，没有实际的物理形式

<img src="https://blogs.vmware.com/customer-experience-and-sucess/files/2021/06/Screen-Shot-2021-06-09-at-1.48.25-AM.png" style="width:300px">

因为 HT 是在 2002 后才推出的技术，一些古早的应用没有经过重构编译可能会出现异常(因为早时的 CPU 并不支持 HT)

### Cache

简单来说，Cache 就是加快 Memory 读取速度的，当然越大越好。对比 Memory， Cache 更靠近 CPU，存取数据快，但是容量也相对的较小

会按照存取的速度从快到慢排序，容量依次递减（只有几 KB 或者几 MB）。例如 L1,L2,L3 ... Ln

当没有命中 Ln 的 Cache，就会使用 Ln+1 的 Cache

后来由于技术的发展 Cache L1 裂变成 L1i(for instruction) 和 L1d (for data)

Cache 和 Core 类似，也是由一组晶体管组成

![img](https://qph.cf2.quoracdn.net/main-qimg-4beade6f4c50c0ad463e9a49d8fa7c54-lq)

### Die

Die 是 Processor 中最核心的部分，由 Core,Cache,Controller 等组件构成

![img](https://qph.cf2.quoracdn.net/main-qimg-86c1d7d49c32c6532eccd15768ee6b84-pjlq)

物理上是一块连续的半导体材料，下图中亮银色的东西(也被称为 Heatsink)下面就是 Die

![img](https://qph.cf2.quoracdn.net/main-qimg-22c01a46fbafac67ce3a5f64a14f9f64-pjlq)

### Package

Package 是 Processor 的保护外壳(早期的 Processor 并没有 Package)

下图中的左侧的盖子就是 Package

![img](https://qph.cf2.quoracdn.net/main-qimg-df8f1d338549ce8ffc9edab8d029a87c-lq)

### Clock Rate

在说明 Clock Rate 之前，需要先介绍下 Clock Cycle，指的是 2 个 Pulses 之前的时间。一个 Clock Cycle 可以执行一个或者多个 Instruction

而 Clock Rate，也被称为 Clock Speed 或者 frequencey(中文叫频率 或者 时脉频率)，指的是每秒包含的 Clock cycle 数量，通常以 Hz 为单位(MHz,GHz)。该值越大，说明每秒能执行的 Instruction 也越多，CPU 性能也越强。

例如一个 CPU 的 Clock Rate 为 4GHz, 那么每秒就有 4000,000,000 Clock Cycle

### Turbo Boost

Turbo Boost 中文也叫做睿频，是 Intel 推出的一种技术，可以让 CPU 根据负载动态的调整 Clock Rate

### Base Clock/Current Clock/Boost Clock

> 都可以通过 `inxi -C` 获取

Base Clock 中文也叫基频，指是 CPU 能到的最小 Clock Rate

Current Clock 中文没有明确的叫法，指的是 CPU 当前的 Clock Rate

Boost Clock 中文没有明确的叫法, 指的是 CPU 能到的最大 Clock Rate

### Overclocking/Underclocking

Overclocking 中文也叫超频，指的是人为提高 Boost Clock 上限这一动作(通常可以在 Bios 中设置)。如果在正常使用中经常顶到 Boost Clock，就可以通过 Overclocking 来解决这个问题，但是也会导致 Proccessor 的温度上升

Underclocking 中文也叫降频，指的是人为降低 Boost Clock 上限这一动作

### NUMA

> 强烈推荐看注脚的文档

在没有 NUMA 前，假设有一台 2 板卡(意味着有多个 Socket, 这里以 2 个 Socket 为例，就能安装 2 个 Processor，以及若干在不同板卡上的内存条)的服务器， A 板卡上的 CPU 想要获得 B 板卡上内存条的数据，这时 CPU 到 内存条 的信道长度(这里主要指物理意义上的)是影响 throughput 的一个因子，信道越长，throughput 就越小。为了解决这个问题就引入了 Non-Uniform Memory Access(NUMA)[^17] 的逻辑，在系统为应用分配 CPU 以及 内存 资源时，会将其分配在同一板卡上，这样就不存在跨板卡获取数据的问题，CPU 到 内存条 的信道相对也是最短的，throught 比没有使用 NUMA 的就高

在 NUMA 中，会将一块板卡(英文为 bank)上的 Processor 以及 内存条，作为一个 NUMA Node (通常个人电脑只有一块主板，也就只有一个 NUMA Node)

![A System with Two NUMA Nodes and Eight Processors](https://dl.acm.org/cms/attachment/fcd8af7a-9903-4a8b-b01c-627fabd852f2/lameter1.png)

在 Linux 中可以通过 `numactl` 来管理 NUMA，例如 `numactl --hardware` 就可以查看服务器上 NUMA node 的信息

```
$ numactl --hardware
available: 2 nodes (0-1)
node 0 cpus: 0 2 4 6 8 10 12 14 16 18 20 22 24 26 28 30
node 0 size: 131026 MB
node 0 free: 588 MB
node 1 cpus: 1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31
node 1 size: 131072 MB
node 1 free: 169 MB
node distances:
node  0  1
  0: 10 20
  1: 20 10  
```

### Tctl/Tdie

Tctl 特指 AMD Processor 上 T Control 温度

Tdie 特制 Die 的温度

## Instruction Set Archtecture

The Instruction Set Archtecture(ISA) acts as an interface between the hardware and the software, which defines the supported data types, the registers, how the hardware manages main memory, key features (such as virtual memory),  which instructions a microprocessor can execute, and the input/output  model of multiple ISA implementations. 

Instruction Set Archtecture(ISA), 中文通常翻译为指令集，是一套为应用提供和硬件交互的底层 API

不同的指令集通常互不兼容，如果在 x64 上编译的应用就不能在 arm64 上运行

### x86/i386

x86 也被称为 i386, 在 intel/AMD 的处理器上使用。有 32bit 指令集, 即应用在每个 clock cycle 最大支持调用 $2^{32}$ memory locatio(大约 4GB 内存)

### x64/x86_64/amd64

也被称为 x86_64 或者 amd64(在 Amd 的帮助下才产生), 在 intel/AMD 的处理器上使用。基于 x86，有 64bit 指令集, 即应用在每个 clock cycle 最大支持调用 $2^{64}$ memory location(大约 16EB 的内存)

通常在 x86 上编译的应用，同样也可以在 x64 的机器上运行。Linux 上可能需要 32bit 关联的包才可以运行，否则会出现 file or directory not found

### aarch64/arm64

aarch64 也被称为 arm64, 在 ARM 的处理器上使用，类似 x64 有 64bit 指令集

ARM 对比 Intel/AMD 处理器在于耗电量小，通常被用在手机平板，也有少数电脑会使用(苹果 Macbook m1/m2/m3)，也有一些云提供商使用 ARM 架构的 CPU 做虚拟化。但是功能性以及兼容性上没有 Intel/AMD 强

### RISC-V

开源的指令集

## lscpu/lshw/dmidecode/inxi

在 Linux 中可以通过多种方式来查看 CPU 相关的信息，以 `lscpu` 为例

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

**refences**

[^1]:https://en.wikipedia.org/wiki/Central_processing_unit
[^2]:https://stackoverflow.com/questions/40163095/what-is-socket-core-threads-cpu
[^3]:https://cloud.tencent.com/developer/article/1736628
[^4]:https://blogs.vmware.com/customer-experience-and-success/2021/06/sockets-cpus-cores-and-threads-the-history-of-x86-processor-terms.html
[^5]:https://virtual-dba.com/blog/sockets-cores-and-threads/
[^6]:https://www.digitaltrends.com/computing/how-to-overclock-your-cpu/
[^7]:https://phoenixnap.com/kb/x64-vs-x86
[^8]:https://www.howtogeek.com/194756/cpu-basics-multiple-cpus-cores-and-hyper-threading-explained/
[^9]:https://en.wikipedia.org/wiki/Instruction_set_architecture
[^10]:https://en.wikipedia.org/wiki/Multi-core_processor
[^11]:https://www.techspot.com/article/2000-anatomy-cpu/
[^12]:https://www.reddit.com/r/explainlikeimfive/comments/1bkacsu/eli5_what_is_core_and_what_is_processor_how_are/
[^13]:https://www.linkedin.com/pulse/understanding-physical-logical-cpus-akshay-deshpande
[^14]:https://www.quora.com/What-do-CPU-cores-physically-look-like?no_redirect=1
[^15]:https://stackoverflow.com/questions/43651954/what-is-a-clock-cycle-and-clock-speed
[^16]:https://www.intel.com/content/www/us/en/gaming/resources/cpu-clock-speed.html
[^17]:https://queue.acm.org/detail.cfm?id=2513149