## Administrative Distance

一次路由中不同协议都为路由器提供了路由信息。这时候就需要Administrative distance来决定采用哪一个协议。==值越小协议优先级越高==。例如：

router 同时接受到Interior Gateway Routing(IGRP - default administrative distance - 100)和Open Shortest Path First(OSPF - default administrative distance - 100)。这时候就会使用IGRP的路由信息到routing table

## Defaul Administrative Distance

| Route Source                                                 | Default Distance Values |
| ------------------------------------------------------------ | ----------------------- |
| Connected interface                                          | 0                       |
| Static route                                                 | 1                       |
| Enhanced Interior Gateway Routing Protocol (EIGRP) summary route | 5                       |
| External Border Gateway Protocol (BGP)                       | 20                      |
| Internal EIGRP                                               | 90                      |
| IGRP                                                         | 100                     |
| OSPF                                                         | 110                     |
| Intermediate System-to-Intermediate System (IS-IS)           | 115                     |
| Routing Information Protocol (RIP)                           | 120                     |
| Exterior Gateway Protocol (EGP)                              | 140                     |
| On Demand Routing (ODR)                                      | 160                     |
| External EIGRP                                               | 170                     |
| Internal BGP                                                 | 200                     |
| Unknown*                                                     | 255                     |