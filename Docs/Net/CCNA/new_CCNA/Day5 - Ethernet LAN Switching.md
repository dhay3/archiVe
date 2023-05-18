# Day5 - Ethernet LAN Switching

## LAN

Local Area Networks(LAN)

switch 并不会划分 LAN，但是路由器会划分 LAN

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_21-18.3npugjmmmy0w.webp)

上述 1234 分别代表不同的 4 个 LAN

## Ethernet Frame

> 这里只考虑 Ethernet

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_21-23.66w1mll3264g.webp)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_21-48.5jrq8roburuo.webp)

一个二层帧由 3 部分组成

- Eth header
- Packet
- Eth trailer

### Eth header

Eth header 通常包含 4 部分

- Preamble
- SFD (start frame delimiter)
- Destination
- Source
- Type or lenght

| Name           | Length          | description                                                  |
| -------------- | --------------- | ------------------------------------------------------------ |
| Preamble       | 7bytes(56 bits) | Allows devices to synchronize their receiver clocks          |
| SFD            | 1bytes(8 bits)  | Marks the end of the preamble. and the beginning of the rest of the frame |
| Destination    | 6bytes(48bits)  | Destination MAC(media access control), indicate the devices receiving the frame |
| Source         | 6bytes(48bits)  | Source MAC(media access control), indicate the devices sending the frame |
| Type or Length | 2bytes(12bits)  | A value of 1500 or less in this field indicates the LENGTH of the encapsulated packet(in bytes).<br>A value of 1536 or greater in this field indicates the Type of the encapsulated packet, and the length is determined via other method.<br>![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_21-36.64d5v64f24xs.webp)<br>例如最后字段的值是 0x0800 即十进制 2048 就表示该 frame 是 IPv4 报文<br>如果最后字段的值是 0x86DD 即十进制 34525 就表示该 frame 是 IPv6 报文 |

所以 Ethernet 2 层帧的报文头通常为 22 bytes，但是也有的标准把 Preamble 和 SFD 作为 1 层的报文头，即 2 层帧只有 12 bytes[^Ethernet Header]

### Eth trailer

Eth trailer 只包含一部分 FCS (Frame check sequence)

- 4bytes(32 bits) in length
- Detects corrupted data by running a ‘CRC’(Cyclic Redundancy Check) algorithm over the received data

## MAC address

- 6 bytes(48 bits) physical address assigned to the device when it is made

- A.K.A. ‘Burned-In address’(BIA)

  因为是烧录在设备上的地址

- Is globally unique

  全球地址唯一，也有 localling-unique MAC address 在 LAN 中唯一

- The first 3 bytes are the OUI(Organizationally Unique Identifier), which is assigned to the company makeing the device
- The last 3 bytes are unique to the devices itself
- Written as 12 hexadecimal characters

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_21-56.31iijuh1gqdc.webp)

是一个 12 位的 16 进制数，每 4 位(16 bits)一组。所以上图中

- AAAA.AA 对应 OUI
- 00.0001/00.0003/00.0002 对应设备的唯一标识

假设现在 PC1 需要发包到 PC2 这种就叫做 Unicast(单播) 

*a frame destined for a single target(PC2 in this case)*

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=u2n762WG0Vo&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=10
[^Ethernet Header]:https://notes.networklessons.com/ethernet-header