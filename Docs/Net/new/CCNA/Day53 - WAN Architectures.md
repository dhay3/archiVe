# Day53 - WAN Architectures

## WAN

Wide Area Network(WAN) is a network that extends over a large geographic area, are used to connect geographically separate LANs

Although the Internet itself can be considered a WAN, the term WAN is typically used to refer to an enterprise’s private connections that connect their offices, data centers, and other sites together

> Over public/shared networks like the Internet, VPNs(Virtual Private Networks) can be used to create private WAN connections

例如下图就是一个 WAN

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_13-37.2sr4v77et3w0.webp)

其中 Leased Line 也被称为专线

- A leased line is a dedicated physical link, typically connecting two sites
- Lease lines use serial connections(PPP or HDLC encapsulation)

> 在 WAN 中 Star topology 往往被称为 Hub-and-Spoke topology

但是现在 WAN 通常并不会使用 Leased Line 直联，而会是通过 Service Provider，构建私网

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_13-40.2132n0rgprwg.webp)

但是也可以通过公网，使用 IPsec VPN Tunnel 对报文加密，来实现类似私网的效果

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_13-43.3dkxse9gas20.webp)

## MPLS

另外一种常见的 WAN 架构就是 Multi Protocol Label Switching(MPLS)

*Similar to the Internet, service providers’ MPLS networks are shared infrastructure because many customer enterprises connect to and share the same infrastructure to make WAN connections*

the label switching in the name of MPLS allows VPNs to be created over the MPLS infrastructure through the use of labels

MPLS 就是通过 labels 来划分流量的，在 MPLS 中有几种角色

1. CE router = Customer Edge router
2. PE router = Provider Edge router
3. P router = Provider core router

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_13-53.5serysxic600.webp)

- When the ==PE routers== receive frames from the CE routers, they add a label to the frame
- These labels are used to make forwarding decisions ==within the service provider network, not the destination IP==
- The CE routers do not use MPLS, it is only used by the PE/P routers

When using a Layer 3 MPLS VPN, the CE and PE routers peer using OSPF, for example, to share routing information

例如 下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-05.5jjgholok6o0.webp)

Office A’s CE will peer with one PE, and Office B’s CE will peer with the other PE

Office A’s CE will learn about Office B’s routes via this OSPF peering, and Office B’s CE will learn about Office A’s routes too



When using a Layer 2 MPLS VPN, the CE and PE routers do not form peerings

The service provider network is entirely transparent to the CE routers. In effect, it is like the two CE routers are directly connected(Their WAN interfaces will be in the same subnet)

If a routing protocol is used, the two CE routers will peer directly with each other

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-13.48k53vyut540.webp)

> 可以将 Service Provider 部分想象成一个交换机，CE 之间逻辑上的直联

## DSL

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-19.dkascwrz61k.webp)

## CATV

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-19.13dp90aef99c.webp)

## Redundant Internet Connections

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-23.4ci846gbm400.webp)

## Internet VPNs

Private WAN services such as leased lines and MPLS provide security becasue each customer’s traffic is separated by using dedicated physical connections(leased line) or by MPLS tags

When using the Internet as a WAN to connect sites together, there is no built-in security by default. To privde secure communications over the Internet, VPNs(Virtual Private Networks) are used

### Site-to-Site VPNs

基于 IPsec，通常也被为 IPsec VPNs

Site-to-Site 指的是只能是两台设备之间才能构成 Tunnel，不存在多台设备直接构成 Tunnel

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-30.1rfeaep7oulc.webp)

过程主要如下

1. The sending device combines the original packet and session key(encryption key) and runs them through an encryption formula
2. The sending device encapsulates the encrypted packet with a VPN header and a new IP header
3. The sending device sends the new packet to the device on the other side of the tunnel
4. The receiving device decrypts the data to get the original packet, and then forwards the original packet to its destination

标准的 IPsec 有一些限制

- IPsec doesn’t support broadcast and multicast traffic, only unicast.

  这就意味着 IP routing protocol 例如 OSPF 不能使用 IPsec，因为 OSPF 需要借助 multicast

  但是可以通过 GRE over IPsec 来解决

- Configuring a full mesh of tunnels between many sites is a labor-intensive task

  可以通过 Cisco DMVPN 解决

### Remote-access VPNs

基于 TLS

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_15-01.5g8bedcbbjc0.webp)

例如下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_15-04.3xby3ib47oq0.webp)

## Site-to-Site Vs Remote-Access VPN

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_15-05.1paoyq3ss400.webp)

## GRE Over IPsec

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-52.2o1viq864n40.webp)

例如

报文首先会加上 GRE header 和 新的 IP header

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-53.ow1eoaeojk0.webp)

然后按照 IPsec 规则对报文做 encryption，同时加上 VPN header 和 新的 IP header

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-54.5zhsr0oacbg0.webp)

## DMVPN

*DMVPN(Dynamic Multipoint VPN) is a Cisco-developed solution that allows routers to dynamically create a full mesh of IPsec tunnels without having to manually configure every single tunnel*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_14-58.8d0rdjb7vt0.webp)

## LAB

> GRE 并不在 CCNA 考试的范围内

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_15-11.2g2y9317b9gk.webp)

### 0x01

Configure a GRE tunnel to connect R1 and R2.

```
#创建一个 GRE tunnel
R1(config)#interface tunnel 0
R1(config-if)#tunnel source g0/0/0
R1(config-if)#tunnel destination 200.0.0.2
#指定 tunnel 0 的 IP，用于封装新的 IP header
R1(config-if)#ip add 192.168.1.1 255.255.255.252

R1(config)#interface tunnel 0
R1(config-if)#tunnel source g0/0/0
R1(config-if)#tunnel destination 100.0.0.2
R1(config-if)#ip add 192.168.1.2 255.255.255.252
```

这里使用 `show ip int br` 来查看 tunnel 0 端口是否正常

```
R2(config-if)#do show ip int br
Interface              IP-Address      OK? Method Status                Protocol 
GigabitEthernet0/0     10.0.2.1        YES NVRAM  up                    up 
GigabitEthernet0/1     unassigned      YES NVRAM  administratively down down 
GigabitEthernet0/2     unassigned      YES NVRAM  administratively down down 
GigabitEthernet0/0/0   200.0.0.2       YES manual up                    up 
Tunnel0                192.168.1.2     YES manual up                    down 
Vlan1                  unassigned      YES unset  administratively down down
```

这里可以发现是 up/down 的，因为 R2 还没有到 100.0.0.0/30 的路由(同理 R1)，所以不能正常建立 GRE tunnel，也就不能正常 ping 通 192.168.1.1 或者 192.168.1.2

```
     10.0.0.0/8 is variably subnetted, 2 subnets, 2 masks
C       10.0.2.0/24 is directly connected, GigabitEthernet0/0
L       10.0.2.1/32 is directly connected, GigabitEthernet0/0
     200.0.0.0/24 is variably subnetted, 2 subnets, 2 masks
C       200.0.0.0/30 is directly connected, GigabitEthernet0/0/0
L       200.0.0.2/32 is directly connected, GigabitEthernet0/0/0
```

配置默认路由即可

```
R1(config-if)#ip route 0.0.0.0 0.0.0.0 100.0.0.1
R2(config-if)#ip route 0.0.0.0 0.0.0.0 200.0.0.1
```

配置完成后就可以 ping 通 192.168.1.1 或者是 192.168.1.2 了

```
R2(config)#do ping 192.168.1.1

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.1.1, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 0/0/0 ms

R1(config)#do ping 192.168.1.2

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.1.2, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 0/0/0 ms
```

> 可以观察一下 PDU

### 0x02

Configure OSPF on the tunnel interfaces of R1 and R2，to allow PC1 and PC2 to communicate.

这时 PC1 还是不能 ping 通 PC2 的，因为 R1 并没有 PC2 的路由，所以需要导入对应的路由，这里使用 OSPF

```
R1(config)#router ospf 1
R1(config-router)#network 192.168.1.0 0.0.0.3 area 0
R1(config-router)#network 10.0.1.0 0.0.0.255 area 0
R1(config-router)#passive-interface g0/0

R2(config)#router ospf 1
R2(config-router)#network 192.168.1.0 0.0.0.3 area 0
R2(config-router)#network 10.0.2.0 0.0.0.255 area 0
R2(config-router)#passive-interface g0/0
```

从 PC1 ping PC2 来校验

```
C:\>ping 10.0.2.100

Pinging 10.0.2.100 with 32 bytes of data:

Request timed out.
Reply from 10.0.2.100: bytes=32 time<1ms TTL=126
Reply from 10.0.2.100: bytes=32 time<1ms TTL=126
Reply from 10.0.2.100: bytes=32 time<1ms TTL=126

Ping statistics for 10.0.2.100:
    Packets: Sent = 4, Received = 3, Lost = 1 (25% loss),
Approximate round trip times in milli-seconds:
    Minimum = 0ms, Maximum = 0ms, Average = 0ms
```

**references**

1. ^https://www.youtube.com/watch?v=BW3fQgdf4-w&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=103