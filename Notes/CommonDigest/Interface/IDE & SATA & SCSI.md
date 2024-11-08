# Hard Disk Interfaces

ref:

https://www.cnblogs.com/awpatp/archive/2013/01/29/2881431.html

https://en.wikipedia.org/wiki/Hard_disk_drive#Competition_from_SSDs

## IDE

IDE(Integrated Drive Electronics), 本意是指把控制器与盘体集成在一起的硬盘驱动器，是一种硬盘的传输接口, 有另一个名称叫做ATA（Advanced Technology Attachment），指的是相同的东西。

<img src="https://images0.cnblogs.com/blog/34420/201301/29144617-8c1ec011bbc84e0d8b9c138ea67c9003.png"/>



- 特点:

  一般使用16-bit数据总线, 每次总线处理时传送2个字节。PATA接口一般是100Mbytes/sec带宽，数据总线必须锁定在50MHz，为了减小滤波设计的复杂性，PATA使用Ultra总线，通过“双倍数据比率”或者2个边缘(上升沿和下降沿)时钟机制用来进行DMA传输。这样在数据滤波的上升沿和下降沿都采集数据，就降低一半所需要的滤波频率。这样带宽就是:25MHz 时钟频率x 2 双倍时钟频率x 16 位/每一个边缘/ 8 位/每个字节= 100 Mbytes/sec。

## SATA

SATA（Serial ATA）口的硬盘又叫串口硬盘. SATA以它串行的数据发送方式得名。在数据传输的过程中，数据线和信号线独立使用，并且传输的时钟频率保持独立，因此同以往的PATA相比，SATA的传输速率可以达到并行的**30**倍。可以说:SATA技术并不是简单意义上的PATA技术的改进，而是一种全新的总线架构。

<img src="https://images0.cnblogs.com/blog/34420/201301/29144621-4227c55b80804e9f9d0d73170f719fb0.png"/>

<img src="https://images0.cnblogs.com/blog/34420/201301/29144626-a5ce91d1d06f4adab76b80a246f8892e.png"/>

## SCSI

SCSI英文全称：Small Computer System Interface，它出现的原因主要是因为原来的IDE接口的硬盘转速太慢，传输速率太低，因此高速的SCSI硬盘出现。其实SCSI并不是专为硬盘设计的，实际上它是一种总线型接口。独立于系统总线工作.

<img src="https://images0.cnblogs.com/blog/34420/201301/29144631-d4cf3905a28f480a876dd7e5ac503ca4.png"/>

 

<img src="https://images0.cnblogs.com/blog/34420/201301/29144633-e02259701add495aaf74ff2b2192eee6.png"/>

## SAS

SAS(Serial Attached SCSI)即串行连接SCSI，是新一代的SCSI技术。和现在流行的Serial ATA(SATA)硬盘相同，都是采用串行技术以获得更高的传输速度，并通过缩短连结线改善内部空间等。SAS是并行SCSI接口之后开发出的全新接口。此接口的设计是为了改善存储系统的效能、可用性和扩充性，并且提供与SATA硬盘的兼容性。

SAS的接口技术可以向下兼容SATA。具体来说，二者的兼容性主要体现在物理层和协议层的兼容。

<img src="https://images0.cnblogs.com/blog/34420/201301/29144640-67734dd4bd4343a09528d84fc755a5a0.png"/>

## FC

光纤通道的英文拼写是Fibre Channel，和SCIS接口一样光纤通道最初也不是为硬盘设计开发的接口技术，是专门为网络系统设计的，但随着存储系统对速度的需求，才逐渐应用到硬盘系统中。光纤通道硬盘是为提高多硬盘存储系统的速度和灵活性才开发的，它的出现大大提高了多硬盘系统的通信速度。它以点对点(或是交换)的配置方式在系统之间采用了光缆连接。

即, 硬盘本身是不具备FC接口的, 插硬盘的机柜上带有FC接口, 通过光纤与光纤交换机互联.

## 对比

不同磁盘种类的Metrics:

| 磁盘种类 | 最大IOPS | 最大响应时间 |
| -------- | -------- | ------------ |
| ATA/IDE  | 70       | 15ms         |
| FC/SAS   | 140~160  | 10ms         |
| SSD/EFD  | 2500     | 1ms          |

