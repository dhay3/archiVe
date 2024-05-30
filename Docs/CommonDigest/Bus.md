# Bus

参考：

https://zhuanlan.zhihu.com/p/136854650

Bus我们也称为总线，在计算机中担任不同组件之间数据传输的物理或逻辑通道，或是计算机与计算机之间。按照连接的方向可以分为Internal buses(也被称为local buses)和External buses。

- Internal buses：用于连接计算机内部不同组件，例如CPU，memory连接到motherboard
- Extenal buses：用于连接外部设备到计算机，例如打印机

按照种类来分可以分为

- 数据总线，传输数据信息
- 地址总线，传输地址信息
- 控制总线，传输控制信息
- 电源总线，为系统提供电源信号

按照传输方式分类可以分为串行总线和并行总线

常见的串行总线有USB(universal serial bus)，SPI，I2C等

常见的并行总线有PCI，PCIe等

各种总线之间的关系：

![2021-06-23_22-40](https://github.com/dhay3/image-repo/raw/master/20210601/2021-06-23_22-40.q4xro82zxtc.png)

使用总线的优点：

1. 简化软硬件设计
2. 简化系统结构
3. 便于系统扩充

![2021-06-23_22-30](https://github.com/dhay3/image-repo/raw/master/20210601/2021-06-23_22-30.7n80v096p2s.png)

上图是一块有四个PCIe bus卡槽的PBC板