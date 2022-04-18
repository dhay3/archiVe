# Switch

> 可以理解为是一个多口的网桥

外观上和HUB没有很大的区别，工作在Layer 2（现在也有三层交换机甚至是更高层的，==以下都以二层交换机为例==）。

![2021-09-06_12-01](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-09-06_12-01.4fdihlqo4nq0.png)

交换机内部的CPU会在每个端口成功连接是，通过将MAC地址和端口对应，形成一张MAC表。发往该MAC地址的数据包将进送完其对应的端口。==所以可以划分冲突域==。交换机在转发数据包时，不知道也无须知道信源和信宿机的IP地址，只需要知道物理地址。

## 工作方式

- 收到某网段（设为A）MAC地址为X的计算机发给MAC地址为Y的计算机的数据包。交换机从而记下了MAC地址X在网段A。这称为学习（learning）。
- 交换机还不知道MAC地址Y在哪个网段上，于是向除了A以外的所有网段转发该数据包。这称为泛洪（flooding）。
- MAC地址Y的计算机收到该数据包，向MAC地址X发出确认包。交换机收到该包后，从而记录下MAC地址Y所在的网段。
- 交换机向MAC地址X转发确认包。这称为转发（forwarding）。
- 交换机收到一个数据包，查表后发现该数据包的来源地址与目的地址属于同一网段。交换机将不处理该数据包。这称为过滤（filtering）。==直接建立peer to peer==
- 交换机内部的MAC地址-网段查询表的每条记录采用时间戳记录最后一次访问的时间。早于某个阈值（用户可配置）的记录被清除。这称为老化（aging）。