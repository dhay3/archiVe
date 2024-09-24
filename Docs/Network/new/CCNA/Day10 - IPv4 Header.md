# Day10 - IPv4 Header

![](https://github.com/dhay3/image-repo/raw/master/20230524/2023-05-24_13-45.6dgun3aapk3k.webp)

## Version

Lenghth: 4bits

Identifies the version of IP Used

表示 4 层 packet 使用的 IP version

- IPv4 = 4 (0100)

- IPv6 = (0110)

## IHL

Length: 4bits

The Internet Header Length

用于标识 IPv4 header 的长度

Indentifies the length of the header in 4-byte increments

如果值为 5 即 5 x 4 bytes = 20 bytes

- Minimum value is 5(= 20 bytes)，即 IPv4 最小的 Header length 为 20 bytes
- Maximum value is 15(= 15 x 4 = 60 bytes)，即 IPv4 最大的 Header length 为 60 bytes

## DSCP

Length: 6bits

Differentiatied Services Code Point (DSCP)

用于 Qos(Quality of Service)，按照 DSCP 值对流量进行排序处理

## ECN

Length: 2bits

Explicit Congestion Notification (ECN)

Provides end-to-end (between two endpoints) notification of network congestion **without dropping**

需要两端以及底层而网络设备都支持才会启用该功能

## Total Length

Length: 16 bits

Indiacates the total length of the packet ( L3 header + L4 segment )

即 4 层 packet 的大小

- Minimum value of 20 (= IPv4 header with no encapuslated data)

  最小 20 bytes，即不包含封装的 3 层 segment

- Maximum value of 65535(= $2^{16} + 2^{15} +\ ...\ + 2^1 + 2^0$ )

## Indentification

Length: 16bits

- If a packet is fragmented due to being too large, this feild is used to identify which packet the fragment belongs to. All fragments of the same packet will have their own IPv4 header with the smae value in this field

  一般用于标识分包的 packet，分包的 packets 具有相同的 indentification

- Packets are fragmented if larger than the MTU (Maximum Transmission Unit). The MTU is usually 1500 bytes

  分包发生在 packets 大于 MTU 的情况下，一般 MTU 为 1500 bytes

- Fragments are reassembled by the receiving host

  分包会在接受端聚合成原始的 packet

## Flags

Length: 3 bits

Used to control/identify fragments

用于控制标识分包

- Bit 0: reserved, always set to 0
- Bit 1: Don’t Fragment(DF bit), used to indicate a packet that should not be fragmented
- Bit 2: More Fragments (MF bit), set to 1 if there are more fragments in the packet, set to 0 for the last fragment

## Fragment offset

Length: 13 bits

表示当前分包的 packet 在原始 packet 中的偏移位

- Used to indicate the positio of the fragment within the original, unfragmented IP packet
- Allows fragmented packets to be reassembled even if the fragments arrive out of order

## Time to Live

Time to Live (TTL)

Length: 8 bits

- A router wil drop a packet with a TTL of 0
- Used to prevent infinite loops
- Originally designed to indeicate the packet’s maximum lifetime in seconds
- In practice, indicates a ‘hop count’: each time the packet arrives at a router, the router decreases the TTL by 1

## Protocol

Length: 8bits

Indicates the protocol of the encapsulated L4 PDU

用于标识封装的 segment 使用的协议

- value of 6: TCP
- value of 17: UDP
- value of 1: ICMP

## Header checksum

Length: 16bits

A calculated checksum used to check for erros in the IPv4 header

用于计算和校验 IPv4 header 是否有效

When a router receives a packet, it calculates the checksum of the header and compares it to the one in this field of the header. If they do not match, the router drops the packet

## Source/Desination IP address

Length: 32bits(each)

- Source IP address = IPv4 address of the sender of the packet
- Destination IP addres = IPv4 address of the inteded receviver ot the packet

## Options fields

Length: 0 - 320bits

The extension for IPv4 header



**referencs**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=aQB22y4liXA&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=18