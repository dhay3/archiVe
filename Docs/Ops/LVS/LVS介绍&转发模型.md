# LVS介绍&转发模型

## 1、基本介绍

​    LVS是 Linux Virtual Server 的简称，也就是Linux虚拟服务器，已经是 Linux 内核标准的一部分。SLB的四层负载均衡采用LVS实现。



#### （1）相关术语

1. DS：Director Server。指的是前端负载均衡器节点。

1. RS：Real Server。后端真实的工作服务器。

1. VIP：向外部直接面向用户请求，作为用户请求的目标的IP地址。

1. DIP：Director Server IP，主要用于和内部主机通讯的IP地址。

1. RIP：Real Server IP，后端服务器的IP地址。

1. CIP：Client IP，访问客户端的IP地址。



#### （2）LVS的组成

​    LVS 由2部分程序组成，包括 ipvs 和 ipvsadm。

1. ipvs(ip virtual server)：一段代码工作在内核空间，叫ipvs，是真正生效实现调度的代码。

1. ipvsadm：另外一段是工作在用户空间，叫ipvsadm，负责为ipvs内核框架编写规则，定义谁是集群服务，而谁是后端真实的服务器(Real Server)



#### （3）LVS工作原理

​    LVS是基于linux netfilter框架实现（同iptables）的一个内核模块，名称为ipvs。

![img](https://cdn.nlark.com/lark/0/2018/png/66340/1543814191187-ff20ee9c-906d-4522-a1d1-db25a79f4a36.png)

1. 当用户向DS发起请求，调度器将请求发往至内核空间

1. PREROUTING链首先会接收到用户请求，判断目标IP确定是本机IP，将数据包发往INPUT链

1. IPVS是工作在INPUT链上的，当用户请求到达INPUT时，IPVS会将用户请求和自己已定义好的集群服务进行比对，如果用户请求的就是定义的集群服务，那么此时IPVS会强行修改数据包里的目标IP地址及端口，并将新的数据包发往POSTROUTING链

1. POSTROUTING链接收数据包后发现目标IP地址刚好是自己的后端服务器，那么此时通过选路，将数据包最终发送给后端的服务器



#### （4）LVS调度算法

1. **轮叫调度 rr**
   这种算法是最简单的，就是按依次循环的方式将请求调度到不同的服务器上，该算法最大的特点就是简单。轮询算法假设所有的服务器处理请求的能力都是一样的，调度器会将所有的请求平均分配给每个真实服务器，不管后端 RS 配置和处理能力，非常均衡地分发下去。

1. **加权轮叫 wrr**
   这种算法比 rr 的算法多了一个权重的概念，可以给 RS 设置权重，权重越高，那么分发的请求数越多，权重的取值范围 0 – 100。主要是对rr算法的一种优化和补充， LVS 会考虑每台服务器的性能，并给每台服务器添加要给权值，如果服务器A的权值为1，服务器B的权值为2，则调度到服务器B的请求会是服务器A的2倍。权值越高的服务器，处理的请求越多。

1. 最少链接 lc
   这个算法会根据后端 RS 的连接数来决定把请求分发给谁，比如 RS1 连接数比 RS2 连接数少，那么请求就优先发给 RS1

1. **加权最少链接 wlc**
   这个算法比 lc 多了一个权重的概念。

1. 基于局部性的最少连接调度算法 lblc
   这个算法是请求数据包的目标 IP 地址的一种调度算法，该算法先根据请求的目标 IP 地址寻找最近的该目标 IP 地址所有使用的服务器，如果这台服务器依然可用，并且有能力处理该请求，调度器会尽量选择相同的服务器，否则会继续选择其它可行的服务器

1. 复杂的基于局部性最少的连接算法 lblcr
   记录的不是要给目标 IP 与一台服务器之间的连接记录，它会维护一个目标 IP 到一组服务器之间的映射关系，防止单点服务器负载过高。

1. 目标地址散列调度算法 dh
   该算法是根据目标 IP 地址通过散列函数将目标 IP 与服务器建立映射关系，出现服务器不可用或负载过高的情况下，发往该目标 IP 的请求会固定发给该服务器。

1. 源地址散列调度算法 sh
   与目标地址散列调度算法类似，但它是根据源地址散列算法进行静态分配固定的服务器资源。



  9、一致性哈希 CH，不在LVS调度算法中，是SLB所支持的算法。https://www.cnblogs.com/lpfuture/p/5796398.html





## 2、转发模型

官方LVS支持LVS支持NAT/DR/TUNNEL三种转发模型，在多vlan网络环境下部署时，这三种存在网络拓扑复杂，运维成本高的问题；

SLB基于以上三种模型，新增转发模型FULLNAT，实现LVS-RealServer间跨vlan通讯；

#### （1）NAT模型

![img](https://cdn.nlark.com/lark/0/2018/png/66340/1543814817840-6d21aef8-2c2e-4ba6-a10b-d04d620e4d31.png)

1.  当用户请求到达DS，此时请求的数据报文会先到内核空间的PREROUTING链。 此时报文的源IP为CIP，目标IP为VIP

1. PREROUTING检查发现数据包的目标IP是本机，将数据包送至INPUT链

1. IPVS比对数据包请求的服务是否为集群服务，若是，修改数据包的目标IP地址为后端服务器IP，然后将数据包发至POSTROUTING链。 此时报文的源IP为CIP，目标IP为RIP

1. POSTROUTING链通过选路，将数据包发送给Real Server

1. RS比对发现目标为自己的IP，开始构建响应报文发回给Director Server。 此时报文的源IP为RIP，目标IP为CIP 

1. DS在响应客户端前，此时会将源IP地址修改为自己的VIP地址，然后响应给客户端。 此时报文的源IP为VIP，目标IP为CIP



特点：

- RS应该使用私有地址，RS的网关必须指向DIP

- DIP和RIP必须在同一个网段内

- 请求和响应报文都需要经过DS，高负载场景中，DS易成为性能瓶颈

- 支持端口映射

- RS可以使用任意操作系统

- 缺陷：对DS压力会比较大，请求和响应都需经过DS



#### （2）DR模型

![img](https://cdn.nlark.com/lark/0/2018/png/66340/1543814952802-e071767a-4e68-4232-a6db-de66ad1e550a.png)

1. 当用户请求到达Director Server，此时请求的数据报文会先到内核空间的PREROUTING链。 此时报文的源IP为CIP，目标IP为VIP

1. PREROUTING检查发现数据包的目标IP是本机，将数据包送至INPUT链

1. IPVS比对数据包请求的服务是否为集群服务，若是，将请求报文中的源MAC地址修改为DIP的MAC地址，将目标MAC地址修改RIP的MAC地址，然后将数据包发至POSTROUTING链。 此时的源IP和目的IP均未修改，仅修改了源MAC地址为DIP的MAC地址，目标MAC地址为RIP的MAC地址 

1. 由于DS和RS在同一个网络中，所以是通过二层来传输。POSTROUTING链检查目标MAC地址为RIP的MAC地址，那么此时数据包将会发至Real Server。

1. RS发现请求报文的MAC地址是自己的MAC地址，就接收此报文。处理完成之后，将响应报文通过lo接口传送给eth0网卡然后向外发出。 此时的源IP地址为VIP，目标IP为CIP 

1. 响应报文最终送达至客户端



特性

- 保证前端路由将目标地址为VIP报文统统发给DS，而不是RS

- RS可以使用私有地址；也可以是公网地址，如果使用公网地址，此时可以通过互联网对RIP进行直接访问

- RS跟DS必须在同一个物理网络中

- 所有的请求报文经由Director Server，但响应报文必须不能进过DS

- 不支持地址转换，也不支持端口映射

- RS可以是大多数常见的操作系统

- RS的网关绝不允许指向DIP

- RS上的lo接口配置VIP的IP地址

- RS和DS必须在同一机房中



#### （3）Tunnel模型

在原有的IP报文外再次封装多一层IP首部，内部IP首部(源地址为CIP，目标IIP为VIP)，外层IP首部(源地址为DIP，目标IP为RIP)

![img](https://cdn.nlark.com/lark/0/2018/png/66340/1543817100195-04710e56-12ee-46e1-a351-f46b2c4eca14.png)



1. 用户请求到达Director Server，此时请求的数据报文会先到内核空间的PREROUTING链。 此时报文的源IP为CIP，目标IP为VIP 。

1. IPVS比对数据包请求的服务是否为集群服务，若是，在请求报文的首部再次封\装一层IP报文，封装源IP为为DIP，目标IP为RIP。然后发至POSTROUTING链。 此时源IP为DIP，目标IP为RIP 

1. IPVS比对数据包请求的服务是否为集群服务，若是，在请求报文的首部再次封装一层IP报文，封装源IP为为DIP，目标IP为RIP。然后发至POSTROUTING链。 此时源IP为DIP，目标IP为RIP 

1. POSTROUTING链根据最新封装的IP报文，将数据包发至RS（因为在外层封装多了一层IP首部，所以可以理解为此时通过隧道传输）。 此时源IP为DIP，目标IP为RIP

1. RS接收到报文后发现是自己的IP地址，就将报文接收下来，拆除掉最外层的IP后，会发现里面还有一层IP首部，而且目标是自己的lo接口VIP，那么此时RS开始处理此请求，处理完成之后，通过lo接口送给eth0网卡，然后向外传递。 此时的源IP地址为VIP，目标IP为CIP

1. 响应报文最终送达至客户端



**特性**

- RIP、VIP、DIP全是公网地址

- RS的网关不会也不可能指向DIP

- 所有的请求报文经由Director Server，但响应报文必须不能进过Director Server

- 不支持端口映射

- RS的系统必须支持隧道



#### （4）FullNAT模型





nat/fullnat之间区别：

![img](https://cdn.nlark.com/lark/0/2018/png/66340/1543817806426-7ab5d665-e409-41cd-809b-b8417690f18d.png)



FULLNAT一个最大的问题是：RealServer无法获得用户IP；为了解决这个问题我们提出了TOA的概念，主要原理是：将client address放到了TCP Option里面带给后端RealServer，RealServer上通过toa内核模块hack了getname函数，给用户态返回TCP Option中的client ip。