# Day38 - DNS

## The Purpose of DNS

因为单独记忆 IP 比较困难，而 DNS 就是为了方便人们记忆才出现的

例如如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_14-48.6fycs19l651c.webp)

在 windows 中可以使用 `ipconfig /all` 来查看主机使用的 DNS Servers

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_14-49.3w0l3d2eu6rk.webp)

通过 `nslookup <domain>` 来查看域名使用的 IP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_14-50.3he73s3z67i8.webp)

- `server` 表示使用的 DNS 服务器
- `non-authoritative answer` 表示域名查询的结果

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_14-53.mjk1sygd7io.webp)

> 注意在这个链路中，R1 上并不需要配置任何和 DNS 相关的信息，R1 只启路由转发的功能
>
> 但是在思科的设备中 Router 同样可以启 DNS 的服务

## Wireshark Capture

使用 `nslookup youtube.com` 报文交互如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_14-57.1rt6jo2qcpa8.webp)

*DNS ‘A’ record = Used to map names to IPv4 addresses*

*DNS ‘AAAA’ record = Used to map names to IPv6 addresses*

DNS 虽然可以使用 TCP/UDP 来查询，但是标准的 DNS 一般都使用 UDP，port 53

> 通常 TCP 只会在 query 报文大于 512 bytes 的情况下使用

## DNS Cache

*Devices will save the DNS server’s responses to a local DNS cache. This means they don’t have to query the server every single time they want to access a particular destination*

> 即再次查询域名时，无需发送对应 DNS 的报文

在 Windows 上可以使用 `ipconfig /displaydns` 来查看 DNS cache 相关的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-05.1ok3w5hdbheo.webp)

如果想要清空 DNS cache 可以使用 `ipconfig /flushdns`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-07.1mm3lwko54ow.webp)

清空 cache 后，如果需要查询域名，会再次发送对应 DNS 的报文

## Host file

在早期没有 DNS 的时候，设备上都会有一个 Hosts 文件，将域名和 IP 绑定

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-11.5z1jvfe5a3gg.webp)

通过该文件，会直接使用本机上该文件中对应的内容(设备不会发送 DNS 相关的报文)

## DNS in Cisco IOS

For host in network to use DNS, you don’t need to configure DNS on the routers. They will simply forward the DNS messages like any other packets

However, a Cisco router can be configured as a DNS server, although it’s rare

如果需要让 Cisco router 变为 DNS server，需要使用 `ip dns server`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-15.69ih00mdb2io.webp)

如果设备上有对应过来 DNS query 报文中查询的部分就会回送 response，通过如果几个命令来配置 Router 上对应的条目

1. `ip host <hostname> <IP>`

   configure a list of hostname/IP address mapping

2. `ip name-server <IP>`

   configure a DNS server that R1 will query if the requested record isn’t in its host table

   > 可以想象成默认路由

   需要使用 `ip domain lookup` 来开启查询 name server 的功能，否则 `ip name-server <IP>` 命令无效

   > 大多数 Router 上都默认会使用 `ip domain lookup`，所以也可以无效显式声明

上述配置完成后，现在 PC1 想要 ping PC2

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_14-48.6fycs19l651c.webp)

首先需要将 PC1 的 DNS server 地址指向 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-24.3j6ziqdpsnuo.webp)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-28.2p3vrw01sgjk.webp)

1. 当 PC1 ping PC2 使用 host 时，会发送 DNS query 到 DNS server 即 R1

   What’s the IP address of PC2

2. R1 在收到 PC1 发送的 DNS query 后，查看自己的 DNS 配置中，有一条匹配 PC2 的记录，回送 DNS reponse 

   It’s 192.168.0.102

3. PC1 在得到 R1 回送的 DNS reponse 后，在发送 ICMP request

如果现在需要 ping youtube.com

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-35.6r80n0o3q6w.webp)

1. 当 PC1 ping youtube.com 时，会发送 DNS query 到 DNS server 即 R1

   What’s the IP address of youtube.com

2. R1 在收到 PC1 发送的 DNS query 后，查看自己的 DNS 配置中，并没有明确对应 youtube.com 的条目，所以会按照 `ip name-server <IP>` 配置 8.8.8.8 作为 DNS server 去查询

3. 8.8.8.8 在收到 R1 发送过来的 DNS query，查询到有对应 youtube.com 的条目，回送 DNS reponse
4. R1 转发 8.8.8.8 发送过来的 DNS response 到 PC1
5. PC1 可以正常访问 youtube.com

如果需要查看 Router 上的 DNS 条目(==包括 cached DNS==)，可以使用 `R1#show hosts`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_15-53.1obfyj2bykzk.webp)

Flags 中的 temp 表示 cached DNS，并不是本地配置的，而是通过其他 DNS server 学过来的

现在把 R1 上和 DNS 相关的配置都删除，使用如下配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_16-05.4k1hrvzjyn7k.webp)

如果这时 PC1 仍把 R1 作为 DNS server，R1 并不会回送 DNS reponse 到 PC1

> 对比之前的命令少了 `ip dns server` 所以 R1 并不能作为 DNS server

在 Router 中还可以使用 `R1(config)#ip domain name <domain-name>` 来为没有 标准 domain name 的 hosts 自动添加 domain-name，类似 `resolv.conf` 中的 search 指令

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_17-55.3ox9obiq8i2o.webp)

## Commands Summary 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_17-56.4bzer7nk1zw.webp)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-04_18-24.3m5tivkpj0xs.webp)

### 0x01

Configure a default route to the Internet on R1

```
R1(config)#ip route 0.0.0.0 0.0.0.0 203.0.113.2
R1(config)#do show ip route
Gateway of last resort is 203.0.113.2 to network 0.0.0.0

     192.168.0.0/24 is variably subnetted, 2 subnets, 2 masks
C       192.168.0.0/24 is directly connected, GigabitEthernet0/1
L       192.168.0.254/32 is directly connected, GigabitEthernet0/1
     203.0.113.0/24 is variably subnetted, 2 subnets, 2 masks
C       203.0.113.0/30 is directly connected, GigabitEthernet0/0
L       203.0.113.1/32 is directly connected, GigabitEthernet0/0
S*   0.0.0.0/0 [1/0] via 203.0.113.2
```

### 0x02

Cnfigure PC1,PC2 and PC3 to use 1.1.1.1 as their DNS server

Config -> Global -> settings -> DNS server

### 0x03

Configure R1 to use 1.1.1.1 as its DNS server.

```
R1(config)#ip name-server 1.1.1.1
R1(config)#ip domain lookup 
```

Configure host entries on R1 for R1, PC1, PC2, and PC3.

```
R1(config)#ip host PC1 192.168.0.1
R1(config)#ip host PC2 192.168.0.2
R1(config)#ip host PC3 192.168.0.3
R1(config)#ip host R1 192.168.0.254

R1(config)#do show hosts
Default Domain is not set
Name/address lookup uses domain service
Name servers are 1.1.1.1

Codes: UN - unknown, EX - expired, OK - OK, ?? - revalidate
       temp - temporary, perm - permanent
       NA - Not Applicable None - Not defined

Host                      Port  Flags      Age Type   Address(es)
PC1                       None  (perm, OK)  0   IP      192.168.0.1
PC2                       None  (perm, OK)  0   IP      192.168.0.2
PC3                       None  (perm, OK)  0   IP      192.168.0.3
R1                        None  (perm, OK)  0   IP      192.168.0.254
```

Ping PC1 by name from R1

```
R1(config)#do ping PC1

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.0.1, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 0/0/1 ms
```

### 0x04

From PC1 ping youtube.com by name. Analyze the message being sent

> 注意这里虽然可以 ping youtube.com,但是不能 ping PC2/3 因为，PC1 使用的 DNS server 为 1.1.1.1，过 R1 只是转发对应的 DNS query

**references**

1. [^jeremy’s IT Lab]:https://www.youtube.com/watch?v=4C6eeQes4cs&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=73