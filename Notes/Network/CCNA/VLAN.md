# VLAN

ref

https://zhuanlan.zhihu.com/p/343467875

https://www.jannet.hk/virtual-lan-vlan-zh-hans/

https://mp.weixin.qq.com/s?src=11&timestamp=1655404792&ver=3864&signature=zITBQ3DyP*bRi6HAWkG8WD5AL3EtEs3-eHOFfF0Iw2ZtHHlWaq2WmrKiO*Ch6ioR0N5btbeLLX1MrOKAtI8jg8sjp*LsftW9mqlDSfUo4U*TdzZ5r4QTKgLCaYEqWwey&new=1

## 什么是VLAN

Virtual LAN 即 VLAN 是一种可以将物理LAN逻辑上划分层多个广播域的通信技术。将2层广播报文限制在一个VLAN中（划分广播域），这样就很好解决了VLAN的问题

VLAN使用两种协议，IEEE 802.1q（也被称为dot1q，公开RFC中定义的）和 ISL（Inter-Switch Link，Cisco专有）

## 为什么需要VLAN

在没有VLAN前，如果需要通过ARP学到一个地址的MAC，就需要广播，而这种广播包也会泛洪到LAN中的所有设备(主机，交换机，互联的路由器单口)。这对整个网络的性能都是一种很大的损耗

![2022-06-17_02-43](https://github.com/dhay3/image-repo/raw/master/20220617/2022-06-17_02-43.20j6lgnaix28.webp)

图中，是一个由5台二层交换机连接所构成的网络。假设这是计算机A需要与计算机B通信。在基于以太网的通信中，必须在数据帧中指定目的MAC地址才能正常通信，因此计算机A必须先广播ARP请求，来获取计算机B的MAC地址

交换机1收到广播帧(ARP请求)，会将它转发给除接受端口外的其他端口，也就是flooding。接着，交换机2收到广播帧后也会flooding。交换机3，4，5也会flooding。最终ARP请求会被转发到同一网络中的所有客户机上

![2022-06-17_02-49](https://github.com/dhay3/image-repo/raw/master/20220617/2022-06-17_02-49.58oy6j75oxs0.webp)

需要注意以下，ARP请求原本就是为了获取计算机B的MAC地址而发出的。也就是说：只要计算机B能收到就行了。可是事实上，数据帧却传遍了整个网络，导致所有的计算机都收到了它。如此一来，一方面关播信息消耗了网络的整体带宽，另一方面，收到的关播信息的计算机还要消耗一部分CPU时间来对它进行处理。造成了网络带宽和CPU运算能力的大量无谓消耗

其次就是安全风险，因为在一个LAN中，广播包会被所有主机都收到，这也导致了这些主机利用这种特性可以获取源主机的IP和MAC，有一定的风险

## 广播信息经常发出吗

除了ARP之外，还有其他的协议会发出广播帧，主要包括

1. RIP
2. DHCP
3. NetBEUI
4. IPX
5. Apple Talk

## 广播与的分割与VLAN的必要性

分割广播域时，一般都必须使用到路由器。使用路由器后，可以以路由器上的网络接口(LAN Interface)为单位分割广播域。

但是，通常情况下路由器上不会有太多的网络接口，其数目多在1～4个左右。随着宽带连接的普及，宽带路由器(或者叫IP共享器)变得较为常见，但是需要注意的是，它们上面虽然带着多个(一般为4个左右)连接LAN一侧的网络接口，但那实际上是路由器内置的交换机，并不能分割广播域。

况且使用路由器分割广播域的话，所能分割的个数完全取决于路由器的网络接口个数，使得用户无法自由地根据实际需要分割广播域。

与路由器相比，二层交换机一般带有多个网络接口。因此如果能使用它分割广播域，那么无疑运用上的灵活性会大大提高。

用于在二层交换机上分割广播域的技术，就是VLAN。通过利用VLAN，我们可以自由设计广播域的构成，提高网络设计的自由度。

## VLAN实现的机制

在理解了“为什么需要VLAN”之后，接下来让我们来了解一下交换机是如何使用VLAN分割广播域的。

首先，在一台未设置任何VLAN的二层交换机上，任何广播帧都会被转发给除接收端口外的所有其他端口(Flooding)。例如，计算机A发送广播信息后，会被转发给端口2、3、4。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnu0PB2ibnPzcHNichE3ticGiahtEgYz067gDAvgTqIxft0wGvibe6ibCgghx0Q/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

这时，如果在交换机上生成红、蓝两个VLAN;同时设置端口1、2属于红色VLAN、端口3、4属于蓝色VLAN。再从A发出广播帧的话，交换机就只会把它转发给同属于一个VLAN的其他端口——也就是同属于红色VLAN的端口2，不会再转发给属于蓝色VLAN的端口。

同样，C发送广播信息时，只会被转发给其他属于蓝色VLAN的端口，不会被转发给属于红色VLAN的端口。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnu8vO0ibpXHL8iaF1pRMJQPznYCStWM2oBspDUbBdcu5tl4b3R5dyyibI6w/640?wx_fmt=jpeg&wxfrom=5&wx_lazy=1&wx_co=1)

就这样，VLAN通过限制广播帧转发的范围分割了广播域。上图中为了便于说明，以红、蓝两色识别不同的VLAN，在实际使用中则是用“VLAN ID”来区分的。

如果要更为直观地描述VLAN的话，我们可以把它理解为将一台交换机在逻辑上分割成了数台交换机。在一台交换机上生成红、蓝两个VLAN，也可以看作是将一台交换机换做一红一蓝两台虚拟的交换机。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnuCQc2dpaYsRp95yRfZTMKgrMia2ncOY6gaaiabABsyT5sOBDCPWgTEkeQ/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

在红、蓝两个VLAN之外生成新的VLAN时，可以想象成又添加了新的交换机。

但是，VLAN生成的逻辑上的交换机是互不相通的。因此，在交换机上设置VLAN后，如果未做其他处理，VLAN间是无法通信的。

明明接在同一台交换机上，但却偏偏无法通信——这个事实也许让人难以接受。但它既是VLAN方便易用的特征，又是使VLAN令人难以理解的原因。

## VLAN之间怎么通信

请大家再次回忆一下：VLAN是广播域。而通常两个广播域之间由路由器连接，广播域之间来往的数据包都是由路由器中继的。因此，VLAN间的通信也需要路由器提供中继服务，这被称作“VLAN间路由”。

VLAN间路由，可以使用普通的路由器，也可以使用三层交换机。

## VLAN的访问方式

按照交换机的端口，可以分为以下两种访问方式：

- 访问链接(Access Link)
- 汇聚链接(Trunk Link)

### Access Link

访问链接，指的是“只属于一个VLAN，且仅向该VLAN转发数据帧”的端口。在大多数情况下，访问链接所连的是客户机。

通常设置VLAN的顺序是：

- 生成VLAN
- 设定访问链接(决定各端口属于哪一个VLAN)

设定访问链接的手法，可以是事先固定的、也可以是根据所连的计算机而动态改变设定。前者被称为“静态VLAN”、后者自然就是“动态VLAN”了。

#### 静态VLAN

静态VLAN又被称为基于端口的VLAN(Port Based VLAN)。顾名思义，就是明确指定各端口属于哪个VLAN的设定方法。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnut9PxYibfVGsLntibKBGicLAVQvBL3yWz5dXicKIukibGJwpJQBrPWD2BvLg/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

由于需要一个个端口地指定，因此当网络中的计算机数目超过一定数字(比如数百台)后，设定操作就会变得烦杂无比。并且，客户机每次变更所连端口，都必须同时更改该端口所属VLAN的设定——这显然静态VLAN不适合那些需要频繁改变拓补结构的网络。

#### 动态VLAN

另一方面，动态VLAN则是根据每个端口所连的计算机，随时改变端口所属的VLAN。这就可以避免上述的更改设定之类的操作。动态VLAN可以大致分为3类：

- 基于MAC地址的VLAN(MAC Based VLAN)
- 基于子网的VLAN(Subnet Based VLAN)
- 基于用户的VLAN(User Based VLAN)

其间的差异，主要在于根据OSI参照模型哪一层的信息决定端口所属的VLAN。

①、基于MAC地址的VLAN，就是通过查询并记录端口所连计算机上网卡的MAC地址来决定端口的所属。假定有一个MAC地址“A”被交换机设定为属于VLAN“10”，那么不论MAC地址为“A”的这台计算机连在交换机哪个端口，该端口都会被划分到VLAN10中去。计算机连在端口1时，端口1属于VLAN10;而计算机连在端口2时，则是端口2属于VLAN10。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnuJ0LGJzxRmJLXibSAJaAjOSVAYxPqjH6nWSiaCLvoE7tj3uA0icAwgtYMg/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

②、基于子网的VLAN，则是通过所连计算机的IP地址，来决定端口所属VLAN的。不像基于MAC地址的VLAN，即使计算机因为交换了网卡或是其他原因导致MAC地址改变，只要它的IP地址不变，就仍可以加入原先设定的VLAN。

因此，与基于MAC地址的VLAN相比，能够更为简便地改变网络结构。IP地址是OSI参照模型中第三层的信息，所以我们可以理解为基于子网的VLAN是一种在OSI的第三层设定访问链接的方法。

③、基于用户的VLAN，则是根据交换机各端口所连的计算机上当前登录的用户，来决定该端口属于哪个VLAN。这里的用户识别信息，一般是计算机操作系统登录的用户，比如可以是Windows域中使用的用户名。这些用户名信息，属于OSI第四层以上的信息。

总的来说，决定端口所属VLAN时利用的信息在OSI中的层面越高，就越适于构建灵活多变的网络。

### Trunk Link

到此为止，我们学习的都是使用单台交换机设置VLAN时的情况。那么，如果需要设置跨越多台交换机的VLAN时又如何呢?

在规划企业级网络时，很有可能会遇到隶属于同一部门的用户分散在同一座建筑物中的不同楼层的情况，这时可能就需要考虑到如何跨越多台交换机设置VLAN的问题了。假设有如下图所示的网络，且需要将不同楼层的A、C和B、D设置为同一个VLAN。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnuAsTt5kbn2oElPYGjn7wLpHSBRpTRfc3LOMmTfefRMw0X2LYj4AgE2Q/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

这时最关键的就是“交换机1和交换机2该如何连接才好呢?”

最简单的方法，自然是在交换机1和交换机2上各设一个红、蓝VLAN专用的接口并互联了。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnubyIaLSjTkTYUcicqEFewyEjeJDiaxCKua6PTP8l5NsK8rPWs64XZNUicA/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

但是，这个办法从扩展性和管理效率来看都不好。例如，在现有网络基础上再新建VLAN时，为了让这个VLAN能够互通，就需要在交换机间连接新的网线。建筑物楼层间的纵向布线是比较麻烦的，一般不能由基层管理人员随意进行。并且，VLAN越多，楼层间(严格地说是交换机间)互联所需的端口也越来越多，交换机端口的利用效率低是对资源的一种浪费、也限制了网络的扩展。

为了避免这种低效率的连接方式，人们想办法让交换机间互联的网线集中到一根上，这时使用的就是汇聚链接(Trunk Link)。

#### 实现方式

汇聚链接(Trunk Link)指的是能够转发多个不同VLAN的通信的端口。

汇聚链路上流通的数据帧，都被附加了用于识别分属于哪个VLAN的特殊信息。

现在再让我们回过头来考虑一下刚才那个网络如果采用汇聚链路又会如何呢?用户只需要简单地将交换机间互联的端口设定为汇聚链接就可以了。这时使用的网线还是普通的UTP线，而不是什么其他的特殊布线。图例中是交换机间互联，因此需要用交叉线来连接。

接下来，让我们具体看看汇聚链接是如何实现跨越交换机间的VLAN的。

①、A发送的数据帧从交换机1经过汇聚链路到达交换机2时，在数据帧上附加了表示属于红色VLAN的标记。

②、交换机2收到数据帧后，经过检查VLAN标识发现这个数据帧是属于红色VLAN的。

③、因此去除标记后根据需要将复原的数据帧只转发给其他属于红色VLAN的端口。

这时的转送，是指经过确认目标MAC地址并与MAC地址列表比对后只转发给目标MAC地址所连的端口。

只有当数据帧是一个广播帧、多播帧或是目标不明的帧时，它才会被转发到所有属于红色VLAN的端口。

同理，蓝色VLAN发送数据帧时的情形也与此相同。

![图片](https://mmbiz.qpic.cn/mmbiz_png/MQVibZibkKWNPJicScCSbvoyAaTdNrWLbnuRtbzgE7QVvyibAoJKHTIMcoicxBQqJOhrsLY2eyTnfic7Q0AxqWbhEkBA/640?wx_fmt=jpeg&wxfrom=5&wx_lazy=1&wx_co=1)



通过汇聚链路时附加的VLAN识别信息，有可能支持标准的“IEEE 802.1Q”协议，也可能是Cisco产品独有的“ISL(Inter Switch Link)”。如果交换机支持这些规格，那么用户就能够高效率地构筑横跨多台交换机的VLAN。

另外，汇聚链路上流通着多个VLAN的数据，自然负载较重。因此，在设定汇聚链接时，有一个前提就是必须支持100Mbps以上的传输速度。

另外，默认条件下，汇聚链接会转发交换机上存在的所有VLAN的数据。换一个角度看，可以认为汇聚链接(端口)同时属于交换机上所有的VLAN。由于实际应用中很可能并不需要转发所有VLAN的数据，因此为了减轻交换机的负载、也为了减少对带宽的浪费，我们可以通过用户设定限制能够经由汇聚链路互联的VLAN。

## 实战

### Cisco

