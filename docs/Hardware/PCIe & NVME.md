# PCIe & NVME

参考：

https://zhuanlan.zhihu.com/p/62426408

## PCIe

> M.2 connector就是使用PCIe 4.0作为逻辑通道

Peripheral Component Interconnect Express，也被缩写为PCIe或PCI-e是一种==高速的串行总线标准（既是通道也是物理接口）==，一般用在motherboard上连接graphics card(GPU)，hard disk drive，SSD，Wi-Fi。PCIe在高速的同时还提供了I/O virutalization，取代了传统的PCI/PCI-x bus。经过多年的迭代PCIe已经到了PCIe 6.0。

[PCIe各版本速率](https://zh.wikipedia.org/wiki/PCI_Express)

以PCIe 2.0為例，每秒5GT（Gigatransfer）原始數據傳輸率，編碼方式為8b/10b（每10個位元只有8個有效數據），即有效頻寬為4Gb/s = 500MByte/s。

### 背景

早年电脑还没有标准化，各种配件的接口和协议都不同。声卡用声卡的接口，网卡用网卡的接口，大家都不统一，不同品牌不同厂商的接口都不一样，不统一带来的问题就是不方便扩展。为了解决这种不统一，业内当时统一了一个规格**ISA**。随着计算机的发展，带宽受到了接口的限制，所以也就使用**PCI**替代了ISA，PCI相对于ISA不仅提高了带宽，还做到了即插即用，PCI同样还是受到了带宽的限制，所以也就出现了PCIe

## NVME

NVM(Non-Volatile Memory) Express也被称为NVMe，是一种专门为SSD设计的新型传输协议通过PCIe物理通道来传输数据。相较SATA，NVME大大提高了数据的吞吐量（也和NVME协议版本有关）。









