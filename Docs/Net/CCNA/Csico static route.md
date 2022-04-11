# Csico static route

参考：

https://www.cisco.com/c/en/us/support/docs/dial-access/floating-static-route/118263-technote-nexthop-00.html

https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus3000/sw/unicast/503_u1_2/nexus3000_unicast_config_gd_503_u1_2/l3_route.html#wp1076779

## Directly Connected Static Routes

iface配置IP后就会出现directly connect，会在`show ip route`中显示

```
C    192.168.80.0/24 is directly connected, FastEthernet0/0
```

router会认为destination是直接和output interface连接的，destination会被认为是next hop address。next-hop 可以是interface，只针对P2P的interface;但是如果是broadcast interfaces，next-hop 只能是IPv4 address。

例如：访问192.168.80.1 匹配192.168.80.0/24，本端通过fa0/0，next-hop是192.168.80.1，即直联(==不用额外配置路由也可以访问==)

## static route to iface without next hop

参考：https://serverfault.com/questions/226864/specify-the-next-hop-or-the-interface-for-static-routes

如果使用静态路由但是没有指定next hop只有在iface启用的情况下才会生效。例如：

```
#ip route dest-prefix dest-prefix-mask iface
ip route 0.0.0.0 0.0.0.0 f0/0
```

router会将dest address认为和iface(==本端的==)直连(所以会执行ARP)。如果大量使用这种静态路由，就会造成CPU打高和ARP cache堆积(影响RAM的分配)。官方不推荐这种配置方法。

## static route with next hop

static route 最常用的方式是指定一个next hop来进行路由，例如：

```
#ip route dest-prefix dest-prefix-mask  next-hop
ip route 0.0.0.0 0.0.0.0 192.168.80.1
```

## floating static route 

浮动路由

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

## Example

![2021-12-08_22-23](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211208/2021-12-08_22-23.1itbzvxllp0g.png)

fa0/0 - fa1/0 fa0/1 - fa0/0	

配置如下，由于配置了接口IP就会有直连路由，所以R1R2两两互通，R2R3两两互通。

```
R1#show run int fa0/0
interface FastEthernet0/0
 ip address 192.168.80.1 255.255.255.0
 duplex auto
 speed auto
end

R1#show ip route
C    192.168.80.0/24 is directly connected, FastEthernet0/0
```

```
R2#show run int fa0/1
interface FastEthernet0/1
 ip address 192.168.81.1 255.255.255.0
 duplex auto
 speed auto
end

R2#show run int fa1/0
interface FastEthernet1/0
 ip address 192.168.80.2 255.255.255.0
 duplex auto
 speed auto
end

R2#show ip route
C    192.168.81.0/24 is directly connected, FastEthernet0/1
C    192.168.80.0/24 is directly connected, FastEthernet1/0
```

```
R3#show run int fa0/0
interface FastEthernet0/0
 ip address 192.168.81.2 255.255.255.0
 duplex auto
 speed auto
end

R3#show ip route
C    192.168.81.0/24 is directly connected, FastEthernet0/0
```

R1上配置路由

```
R1(config)#ip route 192.168.81.0 255.255.255.0 192.168.80.2

R1# show ip route
S    192.168.81.0/24 [1/0] via 192.168.80.2
C    192.168.80.0/24 is directly connected, FastEthernet0/0
```

现在可以ping通R2:fa0/1，但是现在是ping不通R3:fa0/0的，因为==没有回包的路由==

```
R1#ping 192.168.81.1

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.81.1, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 60/60/64 ms
R1#ping 192.168.81.2

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.81.2, timeout is 2 seconds:
.....
Success rate is 0 percent (0/5)
```

在R3上配置回包的路由

```
R3(config)#ip route 192.168.80.0 255.255.255.0 192.168.81.1
```

这样，R1就可以ping通R3了，同样的R3也可以ping通R1

