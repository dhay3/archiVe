# LVS Digest

ref

https://en.wikipedia.org/wiki/Linux_Virtual_Server

http://www.linuxvirtualserver.org/whatis.html

http://www.austintek.com/LVS/LVS-HOWTO/mini-HOWTO/LVS-mini-HOWTO.html

Linux Virtual Server ( LVS ) is load balancing software for Linux kernel-based operating systems.

Standard client-server semantics are preserved. Each client thinks that it has connected directly with the realserver. Each realserver thinks that it has connected directly to the client. Neither the client nor the realservers have any way of telling that a director has intervened in the connection. 

说白就是一个 VS，也可以理解层 layer 4 switch

LVS 主要依靠 Linux Netfilter framework（source code 在 [net/netfilter/ipvs](https://github.com/torvalds/linux/tree/master/net/netfilter/ipvs) ) 实现 4 层流量转发 ( UDP 和 TCP)。用户可以通过 userland utility - `ipvs` 来配置和管理 LVS，构建 scalable, available network services

目前大多数互联网公司都采用了 LVS 作为负载工具

## What is virtual server

# 

Virtual server is a highly scalable and highly available server built on a cluster of real servers. The architecture of server cluster is *fully transparent* to end users, and the users interact with the cluster system as if it were only a single high-performance virtual server. Please consider the following figure.

![img](http://www.linuxvirtualserver.org/VirtualServer.png)

The real servers and the load balancers may be interconnected by either high-speed LAN or by geographically dispersed WAN. The load balancers can dispatch requests to the different servers and make parallel services of the cluster to appear as a virtual service on a single IP address, and request dispatching can use IP [load balancing](http://kb.linuxvirtualserver.org/wiki/Load_balancing) technolgies or application-level [load balancing](http://kb.linuxvirtualserver.org/wiki/Load_balancing) technologies. Scalability of the system is achieved by transparently adding or removing nodes in the cluster. High availability is provided by detecting node or daemon failures and reconfiguring the system appropriately.