# mac address table vs arp table vs route table

ref：
[https://blog.51cto.com/dengqi/1223132](https://blog.51cto.com/dengqi/1223132)

## macaddress table
在说 MAC 地址表之前需要了解一下 交换机的工作原理
交换机在接受到数据帧后，会记录数据帧中的源MAC和收到包的接口到MAC地址表中。接着检查自己MAC表中的是否有数据帧中的目标MAC地址，如果有则会根据MAC表中记录的对应接口将数据帧发送出去(单播)，如果没有则从其他非接受口发送出去(广播)。而 MAC 地址表就是用来记录，交换机从那个端口学到的MAC
同时根据MAC学习来的方式分成2种
[static mac vs dynamic mac](https://community.cisco.com/t5/switching/mac-address-dynamic-and-static/td-p/1660466#:~:text=A%20Static%20MAC%20address%20is,dynamically%20sourced%20the%20MAC%20address.)

- Static: Static entries are manually added to the table by a switch administrator. Static entries have higher priority than dynamic entries. Static entries remain active until they are removed by the switch administrator.

配置当前交换机上，在交换机重启后仍保留静态MAC

- Dynamic: Dynamic entries are automatically added to the table through a process called MAC learning, in which the switch retrieves the source MAC address (and VLAN ID, if present) of each Ethernet frame received on a port. If the retrieved address does not exist in the table, it is added. Dynamic entries remain in the table for a predetermined amount of time (defined with the command mac-address-table age-time), after which they are automatically deleted.

mac address-table 只会显示从实际的那个物理口学的MAC
```
DSW# show mac address-table
Legend:
        * - primary entry, G - Gateway MAC, (R) - Routed MAC, O - Overlay MAC
        age - seconds since last seen,+ - primary entry using vPC Peer-Link
   VLAN     MAC Address      Type      age     Secure NTFY   Ports/SWID.SSID.LID
---------+-----------------+--------+---------+------+----+------------------
* 3000     8c60.4f46.74fc    static    0          F    F  Po4096
* 10       444c.a827.59df    dynamic   60         F    F  Eth108/1/43
* 10       444c.a8a3.031a    dynamic   60         F    F  Eth108/1/43


```
例如上述表示
 8c60.4f46.74fc 是静态配置的 MAC，该MAC一定从Po4096来，改物理口划分给vlan3000
444c.a827.59df、444c.a8a3.031a 都是从 Eth108/1/43 动态学来的
### 局域网一台交换机
[https://s4.51cto.com/attachment/201306/172813479.jpg](https://s4.51cto.com/attachment/201306/172813479.jpg)

1. 主机A将一个源MAC地址为自己，目标MAC地址为主机B的数据帧发送给交换机
1. 交换机在收到此数据帧后，首先将数据帧中的源MAC地址和对应的接口(f0/1)记录到MAC地址表中
1. 然后交换机检查自己的MAC地址表中是否有数据帧中的目标MAC地址的信息，如果有，则从MAC地址中的记录的接口发送出去，如果没有，则会将次数据从非接收接口的其他所有接口广播出去(除了f0/1)
1. 这时，局域网中所有主机都会收到此帧，但是只有主机B收到此数据帧时会响应并回应一个数据帧，此数据帧中包括主机B的MAC地址
1. 当交换机收到主机B回应的数据帧后，也会记录数据帧中的源MAC地址(也就是主机B的MAC地址)，这时，再当主机A和主机B通信时，交换机根据MAC地址表中的记录，实现单播
### 局域网多台交换机
[https://s4.51cto.com/attachment/201306/174620975.jpg](https://s4.51cto.com/attachment/201306/174620975.jpg)

1. 主机A将一个源MAC地址为自己，目标MAC地址为主机C的数据帧发送给交换机
1. 交换机1收到此数据帧后，会学习源MAC地址，并检查MAC地址表，发现没有目标MAC地址记录，则会将数据帧广播出去，主机B和交换机2都会收到此数据帧
1. 交换机2收到此数据帧后也会将数据帧中的源MAC地址和对应的接口记录记录到MAC地址表中，并检查自己的MAC地址表，发现没有目标MAC地址记录，则会广播次数据帧
1. 主机C收到数据帧后，会响应此数据帧，并回复一个源MAC地址为自己的数据帧，这是交换机1和交换机2都将主机C的MAC地址记录到自己的MAC地址表中，并以单播的形式将次数据帧发送给主机A
1. 这是主机A和主机C通信就是以单播的形式传输数据帧了，主机B和主机C通信如上述过程一样，一次交换机2的MAC地址表中记录这主机A和主机B的MAC地址都对应接口f0/1
## arp table
交换机是通过MAC地址通信的，但是如果获取目标地址的MAC呢？这是就需要ARP协议(将IP解析成MAC)，在每台主机(这里的主机包括3层交换机但是不包括2层交换机)中有一张ARP表，它记录IP地址和MAC地址对应关系
arp table 会显示从物理口或者逻辑口来的arp记录
```
DSW# show ip arp

Flags: * - Adjacencies learnt on non-active FHRP router
       + - Adjacencies synced via CFSoE
       # - Adjacencies Throttled for Glean
       D - Static Adjacencies attached to down interface

IP ARP Table for context default
Total number of entries: 518
Address         Age       MAC Address     Interface
10.133.145.2    00:14:16  70c7.f27d.71e7  Ethernet1/1/1
10.85.17.1      00:02:01  2880.23a2.1c08  Vlan10
```
上面表示 10.133.145.2 对应的MAC是 70c7.f27d.71e7，从 Ethernet1/1/1 学来
[https://s4.51cto.com/attachment/201306/181713341.jpg](https://s4.51cto.com/attachment/201306/181713341.jpg)

1. 如果主机A想发送数据给主机B，主机A首先会检查自己的ARP缓存表，查看是否有主机B的IP地址和MAC地址的对应关闭，如果有，则会将主机B的MAC地址作为源MAC地址封装到数据帧中。如果没有，主机A则会发送一个ARP请求，请求的目标IP地址是主机B的IP地址，目标MAC地址是广播地址(ff:ff:ff:ff:ff:ff)，源IP地址和MAC地址是主机A的IP地址和MAC地址
1. 当交换机接受到此数据帧后，发现此数据帧是广播帧，因此，会将此数据帧从其他接口发送出去
1. 当主机B接受到此数据帧后，会校对IP地址是否是自己的并将主机A的IP地址和MAC地址的对应关闭记录到自己的ARP缓存表中，同时会发送一个ARP应答，其中包括自己的MAC地址
1. 主机A在收到这个回应的数据帧之后，在自己的ARP缓存表中记录主机B的IP地址和MAC地址对应关系。而交换机已经学习到了主机A和主机B的MAC地址
## route table
路由表记录这到不同网段的信息。路由表中的信息2种

1. 直连路由：直接连接在路由器接口的网段，有路由器自动生成
1. 非直连路由：不是直接连接在路由器接口上的网段，需要手动添加或者使用动态路由

路由器工作在网络层，可以识别逻辑地址。当路由器的某个接口收到一个包时，路由器会读取包中相应的目标的逻辑地址的网络部分，然后在路由表中查看。如果路由表中找到目标地址的路由条目，这把包转发到路由对应的接口。如果在路由表中没有找到目标地址的路由条目，但是如果路由配置默认路由，就使用默认路由转发对应的路由接口，如果没有配置默认路由，就会丢弃改包，并返回不可达的信息
[https://s7.51cto.com/attachment/201306/191339247.jpg](https://s7.51cto.com/attachment/201306/191339247.jpg)

1. hosta 在网络层将上层来的数据报文封装成IP包，其中源IP地址为自己，目标IP地址是hostb，hosta用本机配置的24位子网掩码与目标地址进行“与”运算，得出目标地址与本机不是同一地址段，因此发送hostb的数据包需要经过网关路由a的转发
1. hosta 通过arp请求获取网关路由a的 e0 的MAC地址，并在链路层将路由器e0接口的MAC地址封装成目标MAC地址，源MAC地址是自己
1. 路由器a从e0接收到数据帧，把数据链路层的封装去掉，并检查路由表中是否有目标IP地址段(192.168.2.2)匹配的路由条目，根据路由表中记录到192.168.2.0网段的数据发送到下一跳地址10.1.1.2，因此数据在路由器a的e1口重新封装，此时，源MAC地址是路由器a的e1接口的MAC地址，封装的目标MAC地址这是路由2的e1接口的MAC地址、
1. 路由器b从e1口接收到数据帧，同样会把数据链路层的封装去掉，对目标IP进行检测，并与路由表进行匹配，此时发现目标地址段正好是自己e0口直连的网段，路由器b通过arp广播，获知hostb的mac地址，此时数据包在路由器b的e0接口再次封装，源MAC地址是路由器b的e0接口MAC地址，目标MAC地址是hostb的MAC地址。封装完成后直接从路由的e0接口发送给hostb
1. 此时hostb才会收到来自hosta发送的数据
