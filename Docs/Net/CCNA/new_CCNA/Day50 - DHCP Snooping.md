# Day50 - DHCP Snooping

## DHCP Snooping

*DHCP snooping is a security feature of switches that is used to filter DHCP messages received on **untrusted** ports*

- DHCP snooping 只会过滤 DHCP 相关的报文，不会过滤非 DHCP 相关的报文

- 默认所有的端口都为 untrusted

  通常 uplink ports 都会配置成 trusted ports，而 downlink ports 仍保持为 untrusted

  例如

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_20-57.381co7mwg9fk.webp)

  橙色的部分为 downlink 因为针对每台设备而言离 end hosts 更近，每个过来的 DHCP 报文都会被检查是否合理

  蓝色的部分为 uplink 因为针对每台设备而言离 trusted 的设备更近(这里 DHCP server 和 relay 显然不会被用户用来攻击 DHCP 服务)，每个过来的 DHCP 报文不会被检查，而是直接转发

  例如

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_21-04.uaenhx11klc.webp)

  1. PC1 发送 DHCP Discover 到 SW1 downlink 端口就会校验报文是否合法
  2. 如果通过，由 SW1 uplink 端口转发到 SW2 downlink 端口
  3. SW2 downlink 端口同样会校验报文是否合法
  4. 如果通过，由 SW2 uplink 端口转发到 R1
  5. R1 回 DHCP offer 到 PC1，中间无需再校验报文是否合法(因为报文从 uplink 端口过来，是 trusted)

  如果 DHCP 报文判断不合法，就会在 downlink port 被丢弃

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_21-09.7h5uxn2e89vk.webp)

## DHCP attacks

针对 DHCP 的攻击有很多种类型

1. DHCP starvation 也被称为 DHCP exhaustion

### DHCP starvation

An attacker uses spoofed MAC addresses to flood DHCP Discover messages. The target server’s DHCP pool becomes full, resulting in a denial-of-service to other devices

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_21-12.25ihiiwo8vs0.webp)

在 DHCP 中存储实际的 source MAC address 是 CHADDR 字段

> 为什么 DHCP 中需要一个 CHADDR 字段来存储 source MAC address 呢？
>
> 如果是 remote DHCP server 中间过 DHCP relay 转发，server 收到 relay 发过来的 DHCP 报文 source MAC 和 实际发送 DHCP 报文的设备是不一样,为了保持原始设备的 source MAC 所以需要 CHADDR 字段来存储

### DHCP Poisoning

和 ARP poisoning 一样都是 man-in-the-middle 的实现

A spurious DHCP server replies to clients’ DHCP Discover messages and assigns them IP addresses, but makes the clients use the spurious server’s IP as the default gateway

因为 DHCP client 通常只会接受第一个收到的 DHCP offer，在这之后收到的 offer 都会被丢弃

- This will cause the client to send traffic to the attacker instead of the legitimate default gateway
- The attacker can then examine/modify the traffic before forwarding it to the legitimate default gateways

例如如下拓扑

PC1 想要通过 DHCP 获取 IP 地址，R1 是实际的 DHCP 服务器，在同一个 LAN 中有一个攻击者伪装成 DHCP 服务器

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_22-01.39wzvu2gln28.webp)

首先 PC1 会发送 DHCP discover，因为是 3 层目的是 255.255.255.255 所以会被广播到整个 LAN，所以 Attacker 和 DHCP server 都能收到 PC1 发送过来的 DHCP Discover

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_22-21.2fn00c4mlu4g.webp)

假设现在 Attacker 回送的 offer 先比 DHCP server 回送的 offer 找到 PC1，那么 PC1 就会使用 Attacker 提供的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_22-23.54jryt13ecu8.webp)

同时 PC1 会回送 DECLINE 到 R1 表示不会采取 R1 的 offer，因为 Attacker 的 offer 先到，在完成 DORA 后，PC1 获得 Attacker 提供的信息，IP 为 172.16.1.10，默认官网为 172.16.1.2(即为 Attacker 的地址)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_22-25.2xq54yqcr9a8.webp)

所以如果 PC1 想要发送流量到公网，就会先到网关即 Attacker，这样 Attacker 就可以针对报文做监听或者是修改报文的操作

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_22-28.4be8yr0vszy8.webp)

## DHCP Messages

DHCP snooping 会区分是从 DHCP server 来的报文还是从 DHCP client 来的报文

- 如果 DHCP server messages 从 untrusted port 过来，会被直接丢弃，如果是 trusted port 过来，会被转发

  DHCP server messages 主要包括 3 种

  1. OFFER
  2. ACK
  3. NAK = Opposite of ACK, used to decline a client’s REQUEST

- 如果 DHCP client messages 一定是从 untrusted port 过来的，会校验是否合法

  DHCP client messages 主要包括 3 种

  1. DISCOVER
  2. REQUEST
  3. RELEASE = Used to tell the server that the client no longer needs its IP address

## DHCP Snooping Operation

1. If a DHCP message is received on a trusted port, forward it as normal without inspection

2. If a DHCP message is received on an untrusted port, inspect it and act as follow

   1. If it is a DHCP server message, discard it

   2. If it is a DHCP client message, perform the following checks

      1. Discover/Request messages, check if the frame’s source MAC address and the DHCP message’s CHADDR fields match

         Match = forward, mismatch = discard

         > Source MAC 和 CHADDR 都可以被 spoofing，所以这并不是完美的

      2. Release/Decline messages, check if the packet’s source IP address and the receiving interface match the entry in the DHCP snooping binding table

         Match = forward, mismatch = discard

         > 当一个 client 成功地从 DHCP server 租借了一个 IP 地址，会在 DHCP snooping binding table 中添加一条记录

## DHCP snooping configuration

例如配置成如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_20-57.381co7mwg9fk.webp)

需要使用如下命令

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_23-27.1m93phijlc4g.webp)

- `SW2(config)#ip dhcp snooping`

  声明开启 dhcp snooping 功能

- `SW2(config)#ip dhcp snooping vlan <vlan-id>`

  声明 dhcp snooping 在那个 vlan 中启用

- `SW2(config-if)#ip dhcp snooping trust`

  将指定的端口变为 trusted port，如果没有声明该命令默认端口为 untrusted

可以使用 `show ip dhcp snooping binding` 来查看 dhcp snooping binding table

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_23-23.2jnn6gysgww0.webp)

这里可以看到 192.168.100.10 - 12 都是成功租借的地址，所以记录在 dhcp snooping binding table 中

*Release/Decline message will be checked to make sure their IP address/interface ID match the entry in the DHCP snooping table*

假设现在 Attacker 伪装 192.168.100.10 发送了 release dhcp message 告诉 dhcp server 需要释放地址，因为是 untrusted port 过来的所以会校验 IP 地址以及报文的入接口是否匹配 dhcp snooping binding table 中的值，如果匹配就转发，如果不匹配就丢弃

## Rate limiting



**references**

1. [^https://www.youtube.com/watch?v=qYYeg2kz1yE&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=97]