# Day39 - DHCP

*DHCP is used to allow hosts  to automatically/dynamically learn various aspects of their network configuration, such as IP address, subnet mask, default gateway, DNS server, etc, without manual/static configuration*

- Typically used for ‘client devices’ such as workstations(PCs),phones,etc

  一般只有 end host 会使用 DHCP

- Devices such as routers, servers, etc,are usually manually configured
- In small networks(such as home networks) the router typically acts as the DHCP server for hosts in the LAN
- In larger networks, the DHCP server is usually a Windows/Linux server

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_18-46.4j9puir40xy.webp)

Windows 中的 Obtain an IP addres automatically 和 Obtain DNS server address automatically 就是通过 DHCP 来获取的，也可以使用 `ipconfig /all` 来查看 DHCP 的功能是否启用

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_18-49.2n218vkl83y8.webp)

如果是手动配置 IP 的，这里会显示 No

- 192.168.0.167(Preferred)

  This PC was previously assigned this IP address by the DHCP server, so it asked to receive the same address again this time

- Lease Obtained

  IP 租期开启的时间

- Lease Expires

  IP 租期结束的时间，在此时间点后 IP 会被重新分配

  > 但是也可以配置成永久的 IP，但是不推荐，可以想象一下如果 IP 固定分配，IP 数量肯定是不够的

- Default Gateway

  DHCP Server

  DNS Server

  均使用 192.168.0.1 通常出现在家庭网络中，即 Gateway/DNS server/DHCP  Sever一体

## ipconfig /release

如果需要手动提前释放 IP，可以使用 `ipconfig /release`

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_19-02.5oyt6l1e58n4.webp)

逻辑上类似下图

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_19-02_1.ko3xeen0lqo.webp)

下图为 192.168.0.167 告诉 DHCP server 192.168.0.1 需要释放通过 DHCP 获取的 IP 对应的报文

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_19-11.6yomwuz1d2ps.webp)

- DHCP 默认使用 UDP 67 作为 sport，UDP 68 作为 dport
- client IP address 表示 client 会被分配的 IP
- Option 用于表示该 DHCP 报文的用途

## ipconfig /renew

如果需要重新获取 DHCP 分配的 IP，可以使用 `ipconfig /renew`

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_19-43.95ywh5hhte8.webp)

这个过程包含 4 部分

### DHCP Discover

用于发现 DHCP 服务器

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_19-44.25cdqqa7sa80.webp)

报文如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_21-06.1h59qtczha00.webp)

这里可以看到 2 层帧报文 Destination MAC 是 ff:ff:ff:ff:ff:ff，这是 2 层广播地址，会被 Switch 广播到所有互联的设备

> 因为现在 PC 还没有 IP，所以当然要使用 2 层地址；另外也不确定网络中有没有对应的 DHCP 服务器，所以需要使用 2 层广播地址，被所有的 switch 转发

3 层包报文 Source address 0.0.0.0，表示当前 PC 还没有任何 IP 地址，Destination address 是 255.255.255.255

> 目的地址为 255.255.255.255，表示在同一广播域中的所有主机都会收到这个包
>
> 通常 255.255.255.255 对应的 MAC 也为 ff:ff：ff:ff:ff:ff，而 ff:ff:ff:ff:ff:ff 不一定要求是 255.255.255.255，例如 ARP

同时在 Dynamic Host Configuration Protocol(Discover) 部分内，可以发现 Client IP address 是 0.0.0.0，Client MAC address 对应 PC 的 MAC

主要观察一下 Option Requested IP Address，因为之前 PC 的地址是 192.168.0.167，所以在发送 DHCP Discover 时会重新要求 DHCP 服务器分配这个 IP

如果是第一次使用 DHCP，这个值会不一样，告诉 DHCP 服务器 PC IP 由 DHCP 服务器来决定

### DHCP Offer

用于向 PC 反馈建议使用的 IP

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_21-34.448ddukbgt34.webp)

报文如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_21-38.23u4lz88wdz4.webp)

Router 发现自己是 DHCP 后，就会发送一个 unicast 报文，目的 MAC 是 PC(通过 DHCP Discover 学习到的)

3 层地址同样也是 unicast 地址，目的 IP 为 DHCP discover 中 PC 要求的地址

- Option 54 是 DHCP server 的地址
- Option 1 是 PC 需要使用 IP 对应的 subnet mask
- Option 6 Domain Name server 是 PC 需要使用的 DNS server 地址
- Option 3 Router 是 PC 需要使用的 gateway

> 这里需要注意的一点是 Boot flags 0x0000 (Unicast) 表示当前报文是一个 unicast 报文，因为在 DHCP discover 中 PC 告诉 DHCP server 想要 192.168.0.167(即如果 DHCP discover 中 Requested IP address 显式声明使用具体某个 IP，那么 Boot flags 对应的值也会是 0x0000 表示 Unicast)
>
> 因为这时 PC 并没有具体的 IP，有一些 Client 不支持 Unicast，就会在将 Boot flags 的值设置为 broadcast，这样对应的 Destination MAC 地址就会是 ff:ff:ff:ff:ff:ff，Destination IP 为 255.255.255.255 

### DHCP Request

PC 向 DHCP server 要求使用具体某个 IP

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_21-53.25ffmmbhhzwg.webp)

在一个广播域中可能有多台 DHCP 服务器，所有的 DHCP 服务器在收到 DHCP discover 时，都会回送 DHCP offer，PC 究竟使用了那个 DHCP 服务器发送的 DHCP offer 呢？

为了解决这个问题，==PC 会遵循规则发送 DHCP Request 到 PC 第一个收到 DHCP offer 的服务器==

报文如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_22-07.1sbzl1z2y8dc.webp)

23 层地址和 DHCP discover 中的还是一样，因为 PC 这时没有采用任何一个 DHCP offer 中提供的 IP，因为是广播地址，所以广播域中的所有 DHCP server 都会收到 DHCP request

在 DHCP server 收到 DHCP request 时会检查 Option 54 DHCP Server Identifier，如果和自己配置，就表示 PC 采用了当前设备提供的 IP

### DHCP ACK

表示 DHCP server 确认 PC 可以使用 DHCP offer 中提供的 IP,一旦 PC 收到这个报文，就会自动配置 IP

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_22-11.6oyhmhknhcw.webp)

报文如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_22-13.23we4zrrqf1c.webp)

这里可以看到和 DHCP offer 中一样，也会使用 unicast，但是实际同样通过 boot flag 也可以做 broadcast

### DHCP DORA

上面的 4 个过程也被称为 DORA

方向以及是否是关播或单播参考下图

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_22-15.6vau43ffylj4.webp)

> 如果是 release DHCP IP，是 unicast

## DHCP Relay

在一些场景下可能会在不同的 LAN 中都设置 DHCP server，但是在大型的网络拓扑中一般会使用 centralized DHCP server(和 PC 不在一个 LAN 中)

因为 DHCP client 通过 broadcast 发送报文，centralized DHCP server 是收不到对应的报文的(==broadcast messages don’t leave the local subnet==)

如果需要解决这个问题，需要让 router 变为 DHCP relay agent, 这样 router 就会将 client 发送过来的 broadcast DHCP messages 转为 unicast DHCP messages 发送给 centralized DHCP server

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-04_22-26.24vxryvkqpgg.webp)

## DHCP server Configuration

假设拓扑如下，需要将 R1 配置成 DHCP server

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-14.2k9zia5t7vk.webp)

需要使用如下命令

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-23.36x9z6xa8c8w.webp)

- `ip dhcp excluded-address <start address> <end address>`

  不能被 DHCP 分配为地址的 IP, 一般被用于 Router 接口

- `ip dhcp pool <pool name>`

  DHCP 池的名字

- `network <network portion> <mask>`

  宣告可用于分配 DHCP 的地址范围，当然不包括 excluded-address

  > 这里和 OSPF/EIGRP 中的 network 命令不同，这里是 mask 不是 wildcard mask

- `dns-server <server address>`

  DHCP 宣告主机需要使用的 DNS 地址

- `domain-name <domain-name>`

  分配的主机可以用 hostname.domain-name 的形式表示

- `default-router <gateway address>`

  DHCP 宣告主机需要使用的 gateway 地址

- `lease <days> <hours> <mins>`

  DHCP 重新为主机分配 IP 的租期，也可以使用 `lease infinite` 表示永不过期，但是不推荐

Router 配置完成后，可以使用 `R1#show ip dhcp binding` 来查看 IP 分配的情况

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-26.5grbixeb9ts0.webp)

PC1 `ipconfig /all` 信息如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-27.51asuvuigq2o.webp)

## DHCP Relay Agent Configuration

假设拓扑如下，R1 需要配置成 DHCP relay agent

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-29.5ik3xvo0t5kw.webp)

配置如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-33.3twro5xy39og.webp)

- `ip helper-address <dhcp server address>`

  只需要在和 PC1 互联的接口上使用这个命令，指定 DHCP server 地址

## DHCP Client Configuration

在思科的设备中 router 也可以是 DHCP client，即接口 IP 由 DHCP server 自动分配，不常见也不推荐

假设拓扑如下，需要让 R2 g0/1 IP 自动由 DHCP server 分配

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-37.3l4w9de7a35s.webp)

> R2 g0/0 任然是需要手动配置的

配置如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-40.13l6n9rkvt40.webp)

- `ip address dhcp`

  声明当前接口使用 DHCP 分配 IP

## Command Summary

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-41.3mdp3mjdxv40.webp)

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-05_16-43.4varryzp1ibk.webp)

### 0x01

 Configure the following DHCP pools on R2:

- POOL1: 192.168.1.0/24 (reserve .1 to .10)

  DNS 8.8.8.8

  Domain: jeremysitlab.com

  Default Gateway: R1

  ```
  R2(config)#ip dhcp excluded-address 192.168.1.1 192.168.1.10
  R2(config)#ip dhcp pool 192.168.1.0
  R2(dhcp-config)#network 192.168.1.0 255.255.255.0
  R2(dhcp-config)#dns-server 8.8.8.8
  R2(dhcp-config)#domain-name jeremysitlab.com
  R2(dhcp-config)#default-router 192.168.1.1
  ```

  > 注意 `ip dhcp excluded-address` 是在 global config mode 中使用的，并不需要指定那个 dhcp pool

- POOL2: 192.168.2.0/24 (reserve .1 to .10)

  DNS 8.8.8.8

  Domain: jeremysitlab.com

  Default Gateway: R2

  ```
  R2(config)#ip dhcp excluded-address 192.168.2.1 192.168.2.10
  R2(config)#ip dhcp pool 192.168.2.0
  R2(dhcp-config)#network 192.168.2.0 255.255.255.0
  R2(dhcp-config)#dns-server 8.8.8.8
  R2(dhcp-config)#domain-name jeremysitlab.com
  R2(dhcp-config)#default-router 192.168.2.1
  ```

- POOL3: 203.0.113.0/30 (reserve .1)

  ```
  R2(config)#ip dhcp excluded-address 203.0.113.1
  R2(config)#ip dhcp pool 203.0.113.0
  R2(config)#network 203.0.113.0 255.255.255.252
  ```

配置完后，PC2 就都可以通过 DHCP 来获取 IP 了

> 但是 PC1 并不能获取 DHCP 分配的 IP，因为 PC1 的 DHCP discover 并不会到 R2，所以需要将  R1 g0/0 配置成 DHCP relay

PC2

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::2E0:B0FF:FE87:76E5
   IPv6 Address....................: ::
   IPv4 Address....................: 0.0.0.0
   Subnet Mask.....................: 0.0.0.0
   Default Gateway.................: ::
                                     0.0.0.0

Bluetooth Connection:

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: ::
   IPv6 Address....................: ::
   IPv4 Address....................: 0.0.0.0
   Subnet Mask.....................: 0.0.0.0
   Default Gateway.................: ::
                                     0.0.0.0

C:\>ipconfig /renew

   IP Address......................: 192.168.2.11
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 192.168.2.1
   DNS Server......................: 8.8.8.8
```

还需要使用 `ipconfig /all` 来确认其他的一些信息是不是准确的

```
C:\>ipconfig /all

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: jeremysitlab.com
   Physical Address................: 00E0.B087.76E5
   Link-local IPv6 Address.........: FE80::2E0:B0FF:FE87:76E5
   IPv6 Address....................: ::
   IPv4 Address....................: 192.168.2.11
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: ::
                                     192.168.2.1
   DHCP Servers....................: 192.168.2.1
   DHCPv6 IAID.....................: 
   DHCPv6 Client DUID..............: 00-01-00-01-5A-4B-7A-84-00-E0-B0-87-76-E5
   DNS Servers.....................: ::
                                     8.8.8.8

```

### 0x02

Configure R1’s G0/0 interface as a DHCP client

```
R1(config)#int g0/0
R1(config-if)#ip address dhcp 
R1(config-if)#no shutdown
```

what IP address did it configure

> 这里需要为 203.0.113.0 pool 声明 network，否则 R1 没办法通过 DHCP 分配地址，因为 192.168.1.0/192.168.2.0 都办法和 203.0.113.1 互联

```
R1(config-if)#do show ip int g0/0
GigabitEthernet0/0 is up, line protocol is up (connected)
  Internet address is 203.0.113.2/30
  Broadcast address is 255.255.255.255
```

### 0x03

Configure R1 as a DHCP relay agent for the 192.168.1.0/24 subnet

```
R1(config-if)#int g0/1
R1(config-if)#ip helper-address 203.0.113.1

R1(config-if)#do show ip int g0/1
GigabitEthernet0/1 is up, line protocol is up (connected)
  Internet address is 192.168.1.1/24
  Broadcast address is 255.255.255.255
  Address determined by setup command
  MTU is 1500 bytes
  Helper address is 203.0.113.1
```

### 0x04

Use the CLI of PC1 and PC2 to make them request an IP address from their DHCP server

PC1

```
C:\>ipconfig

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: FE80::230:F2FF:FE38:8690
   IPv6 Address....................: ::
   Autoconfiguration IPv4 Address..: 169.254.134.147
   Subnet Mask.....................: 255.255.0.0
   Default Gateway.................: ::
                                     0.0.0.0

Bluetooth Connection:

   Connection-specific DNS Suffix..: 
   Link-local IPv6 Address.........: ::
   IPv6 Address....................: ::
   IPv4 Address....................: 0.0.0.0
   Subnet Mask.....................: 0.0.0.0
   Default Gateway.................: ::
                                     0.0.0.0
C:\>ipconfig /renew

   IP Address......................: 192.168.1.12
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 192.168.1.1
   DNS Server......................: 8.8.8.8

C:\>ipconfig /all

FastEthernet0 Connection:(default port)

   Connection-specific DNS Suffix..: jeremysitlab.com
   Physical Address................: 0030.F238.8690
   Link-local IPv6 Address.........: FE80::230:F2FF:FE38:8690
   IPv6 Address....................: ::
   IPv4 Address....................: 192.168.1.12
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: ::
                                     192.168.1.1
   DHCP Servers....................: 203.0.113.1
   DHCPv6 IAID.....................: 
   DHCPv6 Client DUID..............: 00-01-00-01-94-AE-8E-6C-00-30-F2-38-86-90
   DNS Servers.....................: ::
                                     8.8.8.8
```

PC2

```
C:\>ipconfig /renew

   IP Address......................: 192.168.2.11
   Subnet Mask.....................: 255.255.255.0
   Default Gateway.................: 192.168.2.1
   DNS Server......................: 8.8.8.8
```

所有的配置完后可以在 R2 上对应分配的 IP 信息

```
R2#show ip dhcp binding 
IP address       Client-ID/              Lease expiration        Type
                 Hardware address
192.168.1.12     0030.F238.8690           --                     Automatic
192.168.2.11     00E0.B087.76E5           --                     Automatic
203.0.113.2      0001.63B0.5601           --                     Automatic
```

这里因为没有配置 lease ，所以默认为空表示永不过期

**references**

1. [^jeremy’s IT Lab]:https://www.youtube.com/watch?v=hzkleGAC2_Y&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=75