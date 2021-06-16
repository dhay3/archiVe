# USB & Thunderbolt

## USB

USB(Universal Serial Bus)是一种串口总线标准，至今已有四代，USB1.x，USB2.0，USB3.x和USB4。USB在很大的程度上取代了serial ports和parallel ports

类别和对应接口参考[en_wiki](https://en.wikipedia.org/wiki/USB)/[zh_wiki](https://zh.wikipedia.org/wiki/USB)

![2021-06-16_01-51](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-06-16_01-51.4ebzycmsxo60.png)

> 现有的USB都兼容USB2.0，使用了SuperSpeed的数据线一般会标有ss的标志，也有厂商不标明

**USB3.0**添加了SuperSpeed 最大的传输速率达到了5Gbps(5/8GBps)，比USB2.0快了10倍。USB2.0和USB3.0可以通过Standard-A receptacles的颜色来区分两者(==USB3.0 TypeA 接口为蓝色，但是USB2.0 TypeA 接口为白色==)。USB3.1保存了USB3.0的SuperSpeed也被称为**USB 3.1 Gen1**。在基础之上研发了SuperSpeed+使用了这种传输模式的也被称为**USB3.1 Gen2**速率到达了10Gbps(10/8GBps)，==USB3.1需要电源两端提供100W的电压==。USB3.2保存了USB3.1SuperSpeed+的同时将Type A接口改用Type C，使用双通道传输(==USB Type-A和Type-B不支援雙通道==)速率在10-20Gbps(1.25-2.5GBps)也被**USB3.2 Gen2**，如果是在USB3.0基础之上使用Type C双通道就是**USB3.1 Gen1**。USB4在2019年发布，**USB4**基于Thunderbolt 3协议，传输速率大概40Gbps，兼容Thunderbolt 3、USB3.2和USB2.0

## Thunderbolt

Thunberbolt(雷电)是==硬件接口(母头)==，Thunderbolt 1和2 都使用Mini DisplayPort(MDP)接口 ，而Thunderbolt 3 和 4 采用了USB-C接口。Thunderbolt可以将其理解为USB4（40Gbps）

