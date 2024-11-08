# 广播域 VS 冲突域

refernce:

https://www.guru99.com/collision-broadcast-domain.html

https://study-ccna.com/collision-broadcast-domain/

## 冲突域

collision domain，Layer 1

由HUB或repeater(中继器)，连接起来的网络(同一物理网段上所有节点的集合)。由于半双工，如果同时传输就会导致数据冲突，这样一个域被称为冲突域，这是就采用了CSMA/CD来解决这个问题。由Layer 2即以上的设备划分

![2021-09-06_13-00](https://github.com/dhay3/image-repo/raw/master/20210601/2021-09-06_13-00.2tje6atlnym0.png)

==each port on a hub is in the same collision domain.Each port on a bridge, a switch or router is in a sperate collision==

上图有6个冲突域

## 广播域

broadcast domain，Layer 2	

所有的节点都可以通过数据链路层的广播得到数据（无须经过路由器），这样的一个域被称为广播域。Layer 1 和 Layer 2连接的设备同属于一个广播域。通过Layer 3 设备（路由器，不会转发广播）划分广播域

![2021-09-06_13-05](https://github.com/dhay3/image-repo/raw/master/20210601/2021-09-06_13-05.6mshhfm4m20.png)

==all ports on a hub or a switch are in the same broadcast domain, and all ports on a router are in a different broadcast domain==

上图有三个广播域

