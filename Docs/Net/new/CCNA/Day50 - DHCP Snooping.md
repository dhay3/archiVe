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

同时 PC1 会回送 DECLINE 到 R1 表示不会采取 R1 的 offer，因为 Attacker 的 offer 先到，在完成 DORA 后，PC1 获得 Attacker 提供的信息，IP 为 172.16.1.10，默认网关为 172.16.1.2(即为 Attacker 的地址)

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
         >
         > 所以如果 attacker 做 dhcp flood 是不能被有效识别的

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

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_20-13.1a4ey14unxa8.png)

### Rate limiting

DHCP snooping can limit the rate at which DHCP messages are allowed to enter an interface. If the rate of DHCP messages crosses the configured limit, the interface is err-disabled(Like with port security, the interface can be manually re-enabled, or automatically re-enabled with errdisable recovery)

> 上面方式并不能有效的处理 DHCP flooding(因为 attacker  可以伪造 Source MAc 和 CHADDR)，而通过 rate limiting 就可以在一定程度上缓解 flooding

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_20-35.3xrp8zc50jcw.webp)

开启 rate limiting 非常简单，只需要对想要限流的端口使用 `SW1(config-if)#ip dhcp snooping limit rate <num>` 表示每秒允许接受多少个 DHCP 报文(不仅仅是 Discover)。如果超过设定的值，端口就会进入 errdisable 状态，同时输出 syslog

如果想要将因为 rate limiting 而进入 errdisable 的端口重新启用，和 port security 一样可以通过 `SW1(config)#errdisable recovery cause dhcp-rate-limit` 来自动启用

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_20-43.5i00tlhgsow0.webp)

上图表示因为 dhcp rate-limit 而导致端口进入 errdisable 的，会在 300 秒后将端口执行 `shutdown` 和 `no shutdown` 命令

### no ip dhcp snooping information option

information option 也被称为 DHCP relay agent information option(option 82)

*It provides additional information about which DHCP relay agent received the client’s message, on which interface, in which VLAN, etc*

DHCP relay agents 可以对收到的 client’s DHCP messages 增加 Option82

但是如果开启了 DHCP snooping, 思科交换机也会对收到的 client’s DHCP messages 增加 Option82，即使交换机(3 层交换机)并不是一个 DHCP relay agent

而思科的交换机默认会丢弃从 untrusted ports 过来的 Option82 DHCP messages

![ ](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_20-58.77rj1aafbk0.webp)

所以在此拓扑图中想要 PC1 - 3 DHCP 正常，就需要使用 `SW1(config-if)#no ip dhcp snooping information option` 来关闭交换机指定添加 DHCP Option82 的功能

但是即使对 SW1 使用了 `no ip dhcp snooping information option` R1 还是不能正常收到 PC 发送过来的 Discover，因为 SW2 会为从 SW1 过来的 DHCP messages 加上 Option82

当 R1 收到后仍然会丢弃这些报文，因为并不是从 DHCP relay agent 来的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_21-32.6aupwmgvgsxs.webp)

所以 SW2 同样也需要使用 `no ip dhcp snooping information option`

## Command summary

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_21-39.6mh3cqzwftvk.webp)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230727/2023-07-27_21-46.1q9uf8n461vk.webp)

### 0x01

Configure R1 as a DHCP server

Exclude 192.1681.1 - 192.168.1.9 from the pool

Default gateway: R1

```
R1(config)#ip dhcp excluded-address 192.168.1.1 192.168.1.9
R1(config)#ip dhcp pool POOL 
R1(dhcp-config)#network 192.168.1.0 255.255.255.0
R1(dhcp-config)#default-router 192.168.1.1
```

还需要将 PC 的分配模式修改成 DHCP，然后通过 `show ip dhcp binding` 来查看是否生效

```
R1(config)#do show ip dhcp bind
IP address       Client-ID/              Lease expiration        Type
                 Hardware address
192.168.1.10     0001.6432.B922           --                     Automatic
```

### 0x02

Configure DHCP snooping on SW1 and SW2

Configure the uplink interfaces as trusted ports

```
SW2(config)#ip dhcp snooping 
SW2(config)#ip dhcp snooping vlan  1
SW2(config)#int g0/1
SW2(config-if)#ip dhcp snooping trust 

SW1(config)#ip dhcp snooping 
SW1(config)#ip dhcp snooping vlan  1
SW1(config)#int g0/2
SW1(config-if)#ip dhcp snooping trust 
```

配置完成后将 PC2 分配方式配置成 DHCP 来校验

### 0x03

Use ipconfig /renew on PC2 to get an IP address

这里会失败，因为没有启用 option82，还需要使用如下命令

```
SW2(config)#no ip dhcp snooping information option 
SW1(config)#no ip dhcp snooping information option 
```

使用 `show ip dhcp snooping binding` 来校验

```
SW2(config)#do show ip dhcp snooping binding
MacAddress          IpAddress        Lease(sec)  Type           VLAN  Interface
------------------  ---------------  ----------  -------------  ----  -----------------
00:60:70:DB:6A:23   192.168.1.11     86400       dhcp-snooping  1     FastEthernet0/2
Total number of bindings: 1

SW1(config)#do show ip dhcp snoop bind
MacAddress          IpAddress        Lease(sec)  Type           VLAN  Interface
------------------  ---------------  ----------  -------------  ----  -----------------
00:60:70:DB:6A:23   192.168.1.11     86400       dhcp-snooping  1     GigabitEthernet0/1
Total number of bindings: 1
```

PC2

```
C:\>ipconfig /renew

   IP Address......................: 192.168.1.11
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 192.168.1.1
   DNS Server......................: 0.0.0.0
```

**references**

1. [^https://www.youtube.com/watch?v=qYYeg2kz1yE&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=97]	