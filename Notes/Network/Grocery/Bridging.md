# Bridging

ref

https://www.zhihu.com/question/263496943

https://baike.baidu.com/item/%E6%A1%A5%E6%8E%A5/3574841

## Digest

Bridging 中文也被称为 桥接，通俗的理解就是将两个 ethernet 合并成一个 ethernet

工作在 L2 对数据链路层的包进行转发，Bridge 和 Switch 都有 briding 的作用

## Exmaple

家庭无线路由器一般有四个LAN口，这四个LAN口任意选取其中的两个，每个LAN口连接一台电脑A、B，两个LAN口就是**桥接关系**，从A发出的以太帧，原封不动地到达B，反之亦然。桥设备几乎什么都没有做，除了学习A、B的MAC地址并动态绑定在LAN口，并在转发帧时查询此种绑定关系。

任何一个无线终端绑定无线路由器，会动态生成一个**逻辑（软件）端口**，这个逻辑端口与任何一个LAN口任意组合也是一个**桥接关系**，但与上文不同的是，桥接设备需要做有线“**802.3**”与无线“**802.11**”帧格式转换（二层），其它与上文类似。

上述两种场景，通信的双方感觉不到桥接的存在，桥接对用户透明，仿佛A、B之间只有一根网线相连。

WAN口与任意一个LAN口组合，则是一个**路由关系**，如果两端连接电脑，电脑处于两个网段，需要三层路由处理才能通信，处理的结果是，变换二层以太帧头。

如果一台**防火墙工作在桥接模式**，意味着可以将此设备看成一根网线，只需要将桥接口的两端，一端连接用户，另外一端连接网关，当流量流经防火墙时，防火墙可以做二、三、四层或**放行**、或**丢弃**、或**更改处理。**