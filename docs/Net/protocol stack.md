# protocol stack

参考：

https://www.pcmag.com/encyclopedia/term/protocol-stack

协议栈，也称为协议堆叠。一个协议依赖于另一个协议，例如TCP/IP模型

![Snipaste_2021-08-23_22-38-06](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210823/Snipaste_2021-08-23_22-38-06.m0qzo3cex00.png)

3台主机，A和B通过无线电设备互联(IEEE 802.11)，B和C通过物理电缆互联(以太网)。使用这两项协议A不能直接和C通信。因为这些电脑在概念上是连接在不同网络上的。因此需要一个跨网络协议来连接它们。

你可以结合两个网络来建立一个更加强大的第三个网络协议，能够控制无线和有线传输。但是一个更简单的办法是在这两个协议上建立一个协议。这就是一个协议栈