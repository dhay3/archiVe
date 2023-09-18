# Day2 - Interfaces and Cables



## Interfaces

### RJ45

![switch interfaces](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-16_17-59.5gm776povf28.webp)

RJ45 是最常见的接口类型，全称为 Registrered Jack 45 connector，使用 8PIN。通常被用在双绞网线中

### SFP

SFP (Small Form-Factor Pluggable)，通常和 fiber-optic cables(光纤) 一起使用

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_13-50.1qfk0i4rrkjk.webp)

黄色的部分是 RJ45 口，红色的部分是 SFP 口。但是现在还不能使用，还需要插入 SFP transceiver 才能性(即支持 hot-pluggable)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_13-57.4ie7v09a1am8.webp)



## Cables

按照线缆是否使用屏蔽材质分为两种

- Unshiedled Twsited Pair (UTP) 非屏蔽双绞线
- Shiedled Twsited Pair (STP) 屏蔽双绞线

按照线缆传输介质使用的材质分为两种

- Copper cable 铜质线缆
- Fiber-optic cable 光纤

### UTP Cables

UTP Cables 通常使用 RJ45 作为接口，使用 copper (铜) 做传输介质,

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-16_18-15.566r8qj8kigw.webp)

在 Ethernet 中特性如下表

| Speed   | Common Name      | IEEE Standard | Informal Name | Maximum Length |
| ------- | ---------------- | ------------- | ------------- | -------------- |
| 10Mbps  | Ethernet         | 802.3i        | 10BASE-T      | 100m           |
| 100Mbps | Fast Ethernet    | 802.3u        | 100BASE-T     | 100m           |
| 1Gpbs   | Gigabit Ethernet | 802.3ab       | 1000BASE-T    | 100m           |
| 10Gbps  | 10 Gig Ethernet  | 802.3an       | 10GBASE-T     | 100m           |

- BASE = refers to baseband signaling

- T = twisted pair (双绞)

  使用 twisted pair 主要用于防止 Electromagnetic interference (EMI 电磁辐射干扰)

#### 10BASE-T/100BASE-T

10BASE-T/100BASE-T 通常使用 2 pairs，即 4 wires 用于收发数据

假设现在有一台 PC 和 Switch 需要互联，就会使用如下的方式通过 cable 对接 Connector PIN 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_11-04.3prlfy2j70xs.webp)

- 在 PC 侧 RJ45 母口 PIN12 用于传输数据，母口 PIN36 用于接收数据
- 在 Switch 侧 RJ45 母口 PIN12 用于接收数据，母口 PIN36 用于发送数据
- PIN12 和 PIN12 使用 cable 中同一条线互联，同理 PIN36

因为数据的传输和发送都在独立的通道，所以不会有冲突。也被称为 ==Full-Duplex== 中文也叫做全双工(interface 和 cable 都要支持)

假设现在有一台 Router 和 Switch 需要互联，相同的也会采用上述的方式

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_11-13.38ohoc973g00.webp)

总结一下，在 UTP cables 10BASE-T/100BASE-T 的情况下

- PC 和 Router RJ45 母口 PIN12 用于传输数据，PIN36 用于接收数据
- Switch RJ45 母口 PIN12 用于接收数据，PIN36 用于发送数据

像这种 PIN1 和 PIN1 通过 cable 中同一条线对接，PIN3 和 PIN3 通过 cable 中同一条线对接的 cable 也被称为 ==Straight-through cable==

假设现在 2 台 Router 需要互联

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_11-42.259mhm4zvq00.webp)

就不能使用 Straight-through cable，因为 Router PIN12 都用于发送数据，如果这种情况下还使用 Straight-through cable 逻辑上明显不合理(同理两台 Switch 或者 PC，或者 PC 和 Router)

为了解决这种情况引入了 ==Crossover cable==

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_11-46.65n1p4igvq0w.webp)

- Switch1 RJ45 母口 PIN1 和 Switch2 RJ45 母口 PIN3，通过 cable 中同一条线对接，同理 Switch1 PIN2 和 Switch2 PIN6 
- Switch1 RJ45 母口 PIN3 和 Switch2 RJ45 母口 PIN1，通过 cable 中同一条先对接，同理 Switch1 PIN6 和 Switch2 PIN2

为了方便记忆不同设备不同 PIN 的功能参考下图 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_11-52.54nf4a7efslc.webp)

##### Auto MDI-X

还要按照连接的不同设备来使用 Straight-through cable 还是 Crossover cable 是不是太麻烦了？我连接设备的时候也没区别是 Straight-through cable 还是 Crossover cable 啊，网线插上去就能用了

那是因为现在支持入网的或者是网络设备都支持一个叫做 Auto MDI-X 的技术，会根据对端的设备自动选择 PIN 是用于发送数据还是接收数据

Switch1 和 Switch 互联，在两个接口都不支持 Auto MDI-X 的情况下，数据明显是不能正常传输

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_12-10.5vpmf6q7zf9c.webp)

现在使用 Auto MDI-X 的接口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_12-12.l6pi4bfwra.webp)

那么 Switch2 接口的 PIN12 就会从接收数据变成发送数据，PIN36 从发送数据变成接收数据，这样数据就可以在两台交换机间正常传输了

#### 1000BASE-T/10GBASE-T

1000BASE-T/10GBASE-T，使用 4 pairs，即 8 wires 用于收发数据

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_12-20.5df3fuvqsg00.webp)

且每对 pair 都是 bidirectional，即可以用于传输也可以用于接收数据

### Fiber-Optic Cables

Fiber-Optic cables 中文也叫做光纤，使用光信号替代电信号，通常和 SFP 一起使用

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_14-03.79yi6jn4e2v4.webp)

光纤长成这样，是因为一侧的两个头分别用于传输和接收数据

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_14-05.412xvtpfwkhs.webp)

一般由 4 层构成

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_14-11.3fno7sh4rrb4.webp)

根据一些特性分为两类

- Multimode Fiber
- Sigblemode Fiber

#### Multimode Fiber

Multimode Fiber 横切面张这样

米黄色部分代表 fiberglass core

蓝色部分代表 cladding

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_14-17.581bgnp1h474.webp)

- Multimode Fiber fiberglass core 直径要比 Singlemode Fiber 要大
- 允许传输多角度的光波
- 最大长度比 UTP cables 长，但是比 Singlemode Fiber 短
- 比 Singlemode Fiber 造价便宜，因为使用 LED SFP transciever

#### Singlemode Fiber

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_14-19.4er5uyo4ud4w.webp)

- Singhlemode fiber fiberglass core 直接要比 Multimode Fiber 要小
- 光波以直线输入
- 比 UTP cables 和 Singlemode Fiber 都要长
- 比 Singlemode Fiber 造价贵，因为使用 laser-based SFP transciever

#### Fiber-Optic cable Standards

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230516/2023-05-17_14-26.3js96arx1uww.webp)

## UTP VS Fiber-Optic Cables

| UTP                                                          | Fiber-optic Cables                                           |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Lower cost than fiber-optic                                  | Higher cost than UTP                                         |
| Shorter maximum distance than fiber-optic (~100m)            | Longer maximum distance than UTP                             |
| Camm be vulnerable to EMI (Eletromagnetic Interference)      | No vulnerable to EMI                                         |
| RJ45 Ports used with UTP are cheaper than SFP ports          | SFP ports are more expensive than than RJ45(Signlemode is more expensive than multimode) |
| Emit (leak) a faint signal outside of the cable, which can be copied (=security risk) | Does not emit any signal outside of the cable(=no security risk) |

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=ieTH5lVhNaY&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=5
[^Straight-throught cable vs crossover cable]:https://www.cables-solutions.com/difference-between-straight-through-and-crossover-cable.html
[^Auto MDI-x]:https://en.wikipedia.org/wiki/Medium-dependent_interface