# Csico static route

参考：https://www.cisco.com/c/en/us/support/docs/dial-access/floating-static-route/118263-technote-nexthop-00.html

## directly connect

iface配置IP后就会出现directly connect，会在`show ip route`中显示

```
C       100.0.0.0 is directly connected, FastEthernet0/1
C    192.168.80.0/24 is directly connected, FastEthernet0/0
```

但是不会出现在`show running-config`的route中，因为这不算一条路由

## static route to iface without next hop

参考：https://serverfault.com/questions/226864/specify-the-next-hop-or-the-interface-for-static-routes

如果使用静态路由但是没有指定next hop只有在iface启用的情况下才会生效。例如：

```
#ip route dest-prefix dest-prefix-mask iface
ip route 0.0.0.0 0.0.0.0 f0/0
```

router会将dest address认为和iface直连(所以会执行ARP)。如果大量使用这种静态路由，就会造成CPU打高和ARP cache堆积(影响RAM的分配)。官方不推荐这种配置方法。

## static route with next hop

static route 最常用的方式是指定一个next hop来进行路由，例如：

```
#ip route dest-prefix dest-prefix-mask  next-hop
ip route 0.0.0.0 0.0.0.0 192.168.80.1
```

## floating static route 

由于 Enhanced Interior Gateway Routing Protocol(EIGRP)，默认所有的internal route entry的administrative distance值为90，external routes的administrative distance值为170，==值越大优先级越小==

如果route entry的administrative distance的值大于默认的，就被称为floating static route。只有当dynamically learned route从route table中删除floating static route才会生效(shutdown iface也不会生效)，==默认不显示在route table中==。但是administrative distance的值不能超过255，超过255就会被机会路由不可达。==通常被作为fallback or backup==，例如：

```
#ip route dest-prefix dest-prefix-mask next-hop administrative-distance
R1(config)#ip route 0.0.0.0 0.0.0.0 192.168.80.10 
R1(config)#ip route 0.0.0.0 0.0.0.0 192.168.80.20 250
...
R1#show ip route
Gateway of last resort is 192.168.80.10 to network 0.0.0.0
#这里不会显示floating static route
C    192.168.80.0/24 is directly connected, FastEthernet0/0
S*   0.0.0.0/0 [1/0] via 192.168.80.10

#但是可以从配置文件中看到
R1#show running-config
ip route 0.0.0.0 0.0.0.0 192.168.80.10
ip route 0.0.0.0 0.0.0.0 192.168.80.20 250

#删除0.0.0.0/0 192.168.80.10，可以发现使用了floating static route
R1(config)#no ip route 0.0.0.0 0.0.0.0 192.168.80.10
R1(config)#end
R1#sh ip route
Gateway of last resort is 192.168.80.20 to network 0.0.0.0
C    192.168.80.0/24 is directly connected, FastEthernet0/0
S*   0.0.0.0/0 [250/0] via 192.168.80.20

```

## recursive static route

参考：

https://www.cnblogs.com/geekHao/p/12269613.html

https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/iproute_pi/configuration/15-s/iri-15-s-book/iri-rec-stat-route.html

https://www.cisco.com/c/en/us/support/docs/dial-access/floating-static-route/118263-technote-nexthop-00.html

recursive static route递归静态路由，指的是一条route entry中的next hop出现在了另外一条route entry中，即需要多次查询route table，会消耗额外的资源。==在实际中需要避免这种配置==，例如

r1 route table

```
#subnetted就表示出现了递归路由
C    192.168.80.0/24 is directly connected, FastEthernet0/0
     10.0.0.0/24 is subnetted, 1 subnets
S       10.10.1.0 [1/0] via 192.168.80.4
```

这里f0/0(192.168.80.3)直连192.168.80.0/24。假设现在要访问10.10.10.2/24，先到达192.168.80.4，发现还先需要到达f0/0。所以查询了两次route table，可以通过如下方式配置递归路由，即递归路由

```
ip route 10.10.1.0 255.255.255.0 f0/0 192.168.80.4
```

