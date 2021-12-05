# BGP

reference:

https://www.jannet.hk/border-gateway-protocol-bgp-zh-hans/

## 概述

BGP全写是Border Gateway Protocol最新的版本是BGP-4，即Version4，亦是最常学习和应用的版本。BGP通常应用与比较大型的网络结构中，用作交换不同AS之间的路由信息，例如ISP与ISP 之间的路由交换。BGP的复杂性在于建立Peers上的一些规限，以及有大量可以影响路由结构的Attribute，要学好BGP，就必须知道配置Attribute的方法

## EBGP和IBGP

学习任何routing protocol，第一步需要了解如何组成neighbors，在BGP中被称为Peers

Peers利用TCP 179 port沟通，分为external BGP（EBGP）和internal BGP（EBGP）两种

如果两个router在相同的AS之间组成Peers，就会称为IBGP peers

如果两个router在不同的AS之间组成Peers，就会称为EBGP peers

## 配置BGP

做如下一个简单的例子

![2021-11-30_22-23](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211130/2021-11-30_22-23.8bu7di38hvw.png)



```
R1#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
R1(config)#int f0/0  
R1(config-if)#ip address 192.168.12.1 255.255.255.0
R1(config-if)#exit    
R1(config)#router bgp 65500
R1(config-router)#
*Nov 30 14:58:08.211: %BGP-4-NORTRID: BGP could not pick a router-id. Please configure manually.
R1(config-router)#neighbor 192.168.12.2 remote

R1#sh run int f0/0       
Building configuration...

Current configuration : 85 bytes
!
interface FastEthernet0/0
 ip address 192.168.12.1 255.255.255.0
 duplex full
end

R1#sh run | se bgp 
router bgp 65500
 bgp log-neighbor-changes
 neighbor 192.168.12.2 remote-as 65501
```

```
R2#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
R2(config)#int f0/0  
R2(config-if)#ip address 192.168.12.2 255.255.255.0
R2(config)#int f0/1
R2(config-if)#ip address 192.168.23.2 255.255.255.0
R2(config-if)#exit    
R2(config)#router bgp 65500
R2(config-router)#
*Nov 30 14:58:08.211: %BGP-4-NORTRID: BGP could not pick a router-id. Please configure manually.
R1(config-router)#neighbor 192.168.12.2 remote
R2(config)#router bgp 65501
R2(config-router)#neighbor 192.168.23.3 remote-as 6550 
*Nov 30 15:07:41.967: %BGP-5-ADJCHANGE: neighbor 192.168.12.1 Up 
R2(config-router)#neighbor 192.168.23.3 remote-as 65501

R2#sh run int f0/0
Building configuration...

Current configuration : 85 bytes
!
interface FastEthernet0/0
 ip address 192.168.12.2 255.255.255.0
 duplex full
end

R2#sh run int f2/0
Building configuration...

Current configuration : 85 bytes
!
interface FastEthernet2/0
 ip address 192.168.23.2 255.255.255.0
 duplex full
end

R2#sh run | se bgp   
router bgp 65501
 bgp log-neighbor-changes
 neighbor 192.168.12.1 remote-as 65500
 neighbor 192.168.23.3 remote-as 65501
```

```
R3#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
R3(config)#int f2/0
R3(config-if)#ip address 192.168.23.3 255.255.255.0
R3(config-if)#no shutdown 
R3(config-if)#exit
*Nov 30 15:13:41.543: %LINK-3-UPDOWN: Interface FastEthernet2/0, changed state to up
*Nov 30 15:13:42.543: %LINEPROTO-5-UPDOWN: Line protocol on Interface FastEthernet2/0, changed state to up
R3(config)#router
R3(config)#router bgp 65501
R3(config-router)#neighbor 192.168.23.2 remote-as 65501

R3#sh run int f2/0
Building configuration...

Current configuration : 85 bytes
!
interface FastEthernet2/0
 ip address 192.168.23.3 255.255.255.0
 duplex full
end

R3#sh run | se bgp
router bgp 65501
 bgp log-neighbor-changes
 neighbor 192.168.23.2 remote-as 65501
```

要知道是否建立成功，可以使用`show ip bgp summary`指令

```

```

## show ip bgp summary

要知道peers是否建立成功，可以使用`show ip bgp summary`

```
R2#show ip bgp summary
BGP router identifier 192.168.23.2, local AS number 65501
BGP table version is 1, main routing table version 1

Neighbor        V    AS MsgRcvd MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
192.168.12.1    4 65500      21      21        1    0    0 00:18:57        0
192.168.23.3    4 65501      19      19        1    0    0 00:16:27        0
```

每一栏的意思如下：

- Neighbor

  就是peer router interface 的 IP address

- V

  BGP的版本

- AS

  peer router的AS number。如果AS相同表示建立了IBGP，如果不同表示建立了EBGP

- MsgRcvd/MsgSent

  MsgRcvd 从 Peer收到的包，MsgSent 发送到Peer的包

- TblVer

  传送给这个Peer的BGP database

- InQ

  从Peer收到但为处理的BGP信息，如果这个数字很大的话，即是很多讯息在排队等待处理，代表peer CPU很忙

- OutQ

  等待发送到Peer的BGP信息，如果这个数字很大的话，可能是本端CPU很忙或bandwidth不足

- UP/Down

  这个Connection的维持上线或下线有多长时间了

- State/PfxRcd

  如果显示的是一个数字（计算是0），代表从这个Peer收到的BGP route 数量，即是Peer已成功建立。==但是如果显示的active的话，表示Peer没有建立成功==

## 建立Peers的状态

一般来说，如果设定的没有问题的话，BGP Peers就会成为*Established*的状态。但实际上，Peers在进入Established之前会经过几个状态

- IDLE

  Router正在搜寻Routing Table，找一条能够连接Neighbor路径（==不会使用Default Route==）

- CONNECT

  Router 已经找到连接Neighbor的路径，并且完成了TCP 3-way handshake

- OPEN SENT

  已经发送了BGP的OPEN封包，告诉对方希望建立Peers

- OPEN CONFIRM

  收到了Neighbor 回传封包，对方赞成建立Peers

- ESTABLISHED

  两个Neighbor已经成功建立了Peers

- ACTIVE

  Router仍然处于主动传送封包的状态，收不到对方回传，如果持续见到此状态的话，代表Peers并未成功建立

可以使用`debug ip bgp`，然后执行`clear ip bgp *`来让BGP Peers重新建立起来

```
R3#debug ip bgp
BGP debugging is on for address family: IPv4 Unicast
R3#clear ip bgp *
R3#
*Mar  1 02:11:24.015: BGPNSF state: 192.168.23.2 went from nsf_not_active to nsf_not_active
*Mar  1 02:11:24.019: BGP: 192.168.23.2 went from Established to Idle
*Mar  1 02:11:24.019: %BGP-5-ADJCHANGE: neighbor 192.168.23.2 Down User reset
*Mar  1 02:11:24.023: BGP: 192.168.23.2 closing
*Mar  1 02:11:24.027: BGP: 192.168.23.2 went from Idle to Active
*Mar  1 02:11:24.039: BGP: 192.168.23.2 open active, local address 192.168.23.3
*Mar  1 02:11:24.111: BGP: 192.168.23.2 went from Active to OpenSent
*Mar  1 02:11:24.111: BGP: 192.168.23.2 sending OPEN, version 4, my as: 65501, holdtime 180 seconds

*Mar  1 02:11:24.111: BGP: 192.168.23.2 send message type 1, length (incl. header) 45
*Mar  1 02:11:24.167: BGP: 192.168.23.2 rcv message type 1, length (excl. header) 26
*Mar  1 02:11:24.167: BGP: 192.168.23.2 rcv OPEN, version 4, holdtime 180 seconds
*Mar  1 02:11:24.167: BGP: 192.168.23.2 rcv OPEN w/ OPTION parameter len: 16
*Mar  1 02:11:24.167: BGP: 192.168.23.2 rcvd OPEN w/ optional parameter type 2 (Capability) len 6
*Mar  1 02:11:24.167: BGP: 192.168.23.2 OPEN has CAPABILITY code: 1, length 4
*Mar  1 02:11:24.167: BGP: 192.168.23.2 OPEN has MP_EXT CAP for afi/safi: 1/1
*Mar  1 02:11:24.167: BGP: 192.168.23.2 rcvd OPEN w/ optional parameter type 2 (Capability) len 2
*Mar  1 02:11:24.167: BGP: 192.168.23.2 OPEN has CAPABILITY code: 128, length 0
*Mar  1 02:11:24.167: BGP: 192.168.23.2 OPEN has ROUTE-REFRESH capability(old) for all address-families
*Mar  1 02:11:24.167: BGP: 192.168.23.2 rcvd OPEN w/ optional parameter type 2 (Capability) len 2
*Mar  1 02:11:24.167: BGP: 192.168.23.2 OPEN has CAPABILITY code: 2, length 0
*Mar  1 02:11:24.167: BGP: 192.168.23.2 OPEN has ROUTE-REFRESH capability(new) for all address-families
BGP: 192.168.23.2 rcvd OPEN w/ remote AS 65501
*Mar  1 02:11:24.167: BGP: 192.168.23.2 went from OpenSent to OpenConfirm
*Mar  1 02:11:24.167: BGP: 192.168.23.2 went from OpenConfirm to Established
*Mar  1 02:11:24.167: %BGP-5-ADJCHANGE: neighbor 192.168.23.2 Up
```

## 用Loopback 来建立IBGP Peers

在一个AS当中，除了BGP之外，一般会使用IGP（例如：OSPF，EIGRP）来做动态的路由交换，在这种情况下我们会使用Loopback Interface作为IBGP的neighbor address。

因为Loopback Interface永远都是UP的，而且Neighbor之间可以通过IGP来寻找到达Loopback的路径，这比起使用Interface IP 来作neighbor address来得灵活一点，也减少了因为Interface down而令BGP table不稳定的概率

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211201/Snipaste_2021-08-11_20-17-25.5fbfg5re3r40.png)

