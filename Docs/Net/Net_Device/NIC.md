# NIC

参考：

https://uk.rs-online.com/web/generalDisplay.html?id=ideas-and-advice/network-interface-cards-guide

network interface card也是我们常说的网卡，NIC通过RJ45 cable或fibre我们可以连接到Wi-Fi 或 fast Ethernet.



> 现在的NIC一般使用PCIe接口

NIC有如下几种

- wireless NIC

  通过Wi-Fi微波传播信号，一般被焊死在motherboad上。

  ![Snipaste_2021-08-23_23-28-42](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210823/Snipaste_2021-08-23_23-28-42.1e523e2mjj4w.png)

- wired NIC

  通过电信号传播型号，一般使用RJ-45 socket或是相近的接口，比Wi-Fi稳定能提示更稳定的传输速率

  ![Snipaste_2021-08-23_23-18-28](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210823/Snipaste_2021-08-23_23-18-28.6sy96wx1y6c0.png)

  <img src="https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-08-26_01-11.4rr29wb6oz60.png" alt="2021-08-26_01-11" style="zoom:50%;" />

  

- USB NIC

  通过USB连接主机，扩展坞

  ![Snipaste_2021-08-23_23-27-11](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210823/Snipaste_2021-08-23_23-27-11.36z8ummyq1s0.png)
  
  做数据转发

## 注意点

- 过去一张物理网卡只能有一个IP，但是现在可以通过virtual networking在一张物理网卡上虚拟网卡(IP和MAC都是不同的)

- 路由器只有WAN口配置网卡

- 一张多口的物理网卡，一个端口一个mac分配一个IP(一片上==多个芯片==)。例如企业级3层交换机

  ​	https://v2ex.com/t/688891

