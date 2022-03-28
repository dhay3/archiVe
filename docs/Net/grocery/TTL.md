ref:
[https://www.cloudflare.com/zh-cn/learning/cdn/glossary/time-to-live-ttl/](https://www.cloudflare.com/zh-cn/learning/cdn/glossary/time-to-live-ttl/)
[https://www.techtarget.com/searchnetworking/definition/time-to-live](https://www.techtarget.com/searchnetworking/definition/time-to-live)
[https://datatracker.ietf.org/doc/rfc3443/](https://datatracker.ietf.org/doc/rfc3443/)

https://www.youtube.com/watch?v=56sNDuYlhPk

## applicatons
在应用层面 TTL 用来管理 cache（对应http报文头中的`Expires`,`Cache-Control`），常备用在 CDN，DNS 用来表示内容存在的时间
## network
在网络（IP 数据包）中 TTL 的含义区别于应用中的
一个数据包从网络中传输，如果没有设置限制，就会在网络中 pass from router to router indefinitely。为了解决这个问题，所以才产生了 Time to Live (ttl) 也被称为 hop limit，该值由数据包的发送者设置（1 - 255），不同的OS会有不同的默认值，
每经过路由的一跳就会减1，如果ttl的值到了0还未到达目的IP，数据包就会被discarded并回送ICMP echo( Type 11 Time Exceed，ICMP reply 不需要对应的ICMP request)

Each packet has a place where it stores a numerical value determining how much longer it should continue to move through the network. Every time a router receives a packet, it subtracts one from the TTL count and then passes it onto the next location in the network.
If at any point the TTL count is equal to zero after the subtraction, the router will discard the packet and send an ICMP message back to the originating host.

### Example 

ping 和 traceroute 都使用了 TTL 这里在R1上使用 `traceroute 192.168.81.2`，cisco 命令默认使用 UDP

链路 R1 -> R2 -> R3，只配置了静态路由

```
R1#show run int fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.80.1 255.255.255.0
 duplex auto
 speed auto
end
```

```
R2#show run int fa1/0
Building configuration...

*Mar  1 00:30:39.611: %SYS-5-CONFIG_I: Configured from console by console
Current configuration : 97 bytes
!
interface FastEthernet1/0
 ip address 192.168.80.2 255.255.255.0
 duplex auto
 speed auto
end

R2#show run int fa0/1
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/1
 ip address 192.168.81.1 255.255.255.0
 duplex auto
 speed auto
end
```

```
R3#show run int fa0/0
Building configuration...

Current configuration : 97 bytes
!
interface FastEthernet0/0
 ip address 192.168.81.2 255.255.255.0
 duplex auto
 speed auto
end
```

[trace.pcap](/home/cpl/note/appendix)

![2022-03-28_20-48](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220328/2022-03-28_20-48.2mfdy7s8us40.webp)

traceroute 探测一次会发 3 个包，前 3 个包的 ttl 值为 1，到达 192.168.80.1 时 ttl - 1 值为 0 回送给源 ICMP type 11 (ttl 255 表示还未到达目的端，不可达)，第二次探测的 3 个包的 ttl 值会设置为 2，但是到达了目的了，所以就没有第三次探测了，同时回送给源 ICMP type 11。如果 ttl 的值到达了 30 就会终止

