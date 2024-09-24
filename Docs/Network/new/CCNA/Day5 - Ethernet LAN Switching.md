# Day5 - Ethernet LAN Switching

## LAN

Local Area Networks(LAN)

switch 并不会划分 LAN，但是路由器会划分 LAN

![](https://github.com/dhay3/image-repo/raw/master/20230518/2023-05-18_21-18.3npugjmmmy0w.webp)

上述 1234 分别代表不同的 4 个 LAN

## Ethernet Frame

> 这里只考虑 Ethernet

![](https://github.com/dhay3/image-repo/raw/master/20230518/2023-05-18_21-23.66w1mll3264g.webp)

![](https://github.com/dhay3/image-repo/raw/master/20230518/2023-05-18_21-48.5jrq8roburuo.webp)

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
| Type or Length | 2bytes(12bits)  | A value of 1500 or less in this field indicates the LENGTH of the encapsulated packet(in bytes).<br>A value of 1536 or greater in this field indicates the Type of the encapsulated packet, and the length is determined via other method.<br>![](https://github.com/dhay3/image-repo/raw/master/20230518/2023-05-18_21-36.64d5v64f24xs.webp)<br>例如最后字段的值是 0x0800 即十进制 2048 就表示该 frame 是 IPv4 报文<br>如果最后字段的值是 0x86DD 即十进制 34525 就表示该 frame 是 IPv6 报文 |

> Ethernet 2 层帧的报文头为 22 bytes，==但是 Preamble 和 SFD 通常并不算作 2 层帧的报文头==，即 2 层帧的报文头只有 14 bytes[^Ethernet Header]
>
> 整个 2 层帧 (Header + payload + trailer) 最小的长度为 64bytes，即 payload 最小的长度为 46bytes (64 - 12 - 4)，如果 payload 小于 46 就会使用 padding(padding 全为 0) 来填充 payload
>
> 例如 含有数据的 payload 为 36 bytes，那么就会填充 10 bytes padding 到 46bytes

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

![](https://github.com/dhay3/image-repo/raw/master/20230518/2023-05-18_21-56.31iijuh1gqdc.webp)

是一个 12 位的 16 进制数，每 4 位(16 bits)一组。所以上图中

- AAAA.AA 对应 OUI
- 00.0001/00.0003/00.0002 对应设备的唯一标识

假设现在 PC1 .0001 需要发包到 PC2 .0002 这种就叫做 Unicast(单播) 

*a frame destined for a single target(PC2 in this case)*

首先 PC1 会发送一个 unicast frame - Dest:.0002 Src:.0001

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_10-38.5mvl0znguohs.webp)

到达 SW1 后, 因为没有对应的 Source MAC address 的信息,所以 SW1 会将从那个端口==学来的 Source MAC address== 记录在自己的 MAC address table 中。下图中，即表示 .0001 Source MAC address 是从 F0/1 端口(学)过来的。这种也被称为  Dynamically learned MAC address/Dynamic MAC address

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_10-39.4rk9pjqx39xc.webp)

但是目前 SW1 并不知道 Dest MAC address .0002 的“位置”，所以往 SW1 除从报文来的端口(这里即 F0/1)外的所有的端口发送一个 Unknown Unicast frame(Flood 中文也叫做泛洪)。报文到达 PC3 时，因为 MAC address 不匹配 Dest MAC address 所以会被丢弃

> 因为对应的报文中有详细的 Destination MAC address，不是广播地址，这也被称为 Unicast
>
> 但是不知道对应的 MAC 需要通过那个端口转发，所以被称为 unknown Unicast

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_10-45.73c0wpbi4mww.webp)

因为 PC2 MAC address 匹配 Dest MAC address .0002，所以会回包。当回包报文到达 SW1 时，因为没有 .0002 Source MAC address 的信息，同样会记录 MAC address .0002 是从那个端口( 即 F0/2 )学过来的。因为 Dest MAC address 在 MAC table 中有对应的记录，所以不会 Flood 会直接转发(Forward，也被称为 Known Unicast frame)到 F0/1

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_11-51.m2itc3iuun4.webp)

> - MAC address table 会记录没有的 Source MAC address 从那个的端口学来的
> - Dynamic MAC address 在 Cisco 的设备中有效的生命周期为 5 mins。如果在这个时间段内没有收到任何关联 Dynamic MAC address 的报文，就会自动从 MAC adress table 中删除，如果有收到报文就会更新开始计算的时间点

### Another Example

假设现在 PC1 需要发包到 PC3

PC1 MAC .0001 PC2 MAC .0003

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_12-12.1lpluro747xc.webp)

1. Src MAC address .0001 Dest MAC address .0003 到达 SW1
2. 因为 SW1 没有对应 .0001 Source MAC address 的信息，所以 SW1 记录 MAC address .0001, Interface F0/1。并往 F0/2 和 F0/3 发送 Unknown Unicast Frame Src MAC address .0001 Dest MAC address .0003
3. PC2 收到 SW1 发送的 Unknown Unicast Frame 因为 MAC 不匹配，所以丢弃
4. 因为 SW2 没有对应 .0001 Source MAC address 的信息，SW2 记录 MAC address .0001, Interface F0/3。并往 F0/1 和 F0/2 发送 Unknown Unicast Frame Src MAC address .0001 Dest MAC address .0003
5. PC4 收到 SW2 发送的 Unknown Unicast Frame 因为 MAC 不匹配，所以丢弃
6. PC3 收到 SW2 发送的 Unknown Unicast Frame 因为 MAC 匹配，所以回包。Src MAC address .0003 Dest MAC address .0001
7. Src MAC address .0003 Dest MAC address .0001 到达 SW2
8. 因为 SW2 没有对应 .0003 Source MAC address 的信息，所以 SW2 记录 MAC address .0003, Interface F0/1。因为 Dest MAC address .0001 在 MAC address table 中，所以直接转发到 F0/3
9. 因为 SW1 没有对应 .0003 Source MAC address 的信息，所以 SW1 记录 MAC address .0003, Interface F0/3。因为 Dest MAC address .0001 在 MAC address table 中，所以直接转发到 F0/1
10. PC1 收到报文后，整个流程结束

最后 SW1 MAC address table 中内容如下

| MAC   | Interface |
| ----- | --------- |
| .0001 | F0/1      |
| .0003 | F0/3      |

SW2 MAC address table 中内容如下

| MAC   | Interface |
| ----- | --------- |
| .0001 | F0/3      |
| .0003 | F0/1      |

### MAC address command

- `show mac address-table`

  查看当前设备的 MAC address table

  ```
  SW1#show mac-address-table 
            Mac Address Table
  -------------------------------------------
  
  Vlan    Mac Address       Type        Ports
  ----    -----------       --------    -----
  
     1    0004.9a6e.d870    DYNAMIC     Gig0/1
     1    00d0.d3ad.9cab    DYNAMIC     Fa0/1
  ```

  其中的 Type 表示 MAC address 是 Dynamically 学来的

- `clear mac address-table dynamic`

  删除所有的 dynamic Type MAC address

  ```
  clear mac address-table dynamic 
  ```

- `clear mac address-table dynamic address <MAC ADDRESS>`

  删除指定的 dynamic Type MAC address

  ```
  clear mac address-table dynamic 00d0.d3ad.9cab
  ```

- `clear mac address-table dynamic address interface <Interface>`

  删除从指定 Interface 学来的 dynamic address

  ```
  clear mac address-table dynamic address interface fa0/1
  ```

## ARP

Address Resolution Protocol 简称 ARP

*ARP is used to discover the Layer2 address(MAC address) of a known Layer3 address(IP address)*

一个完整的 ARP 包括两部分

- ARP request
- ARP reply

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_16-40.5vqpew48enb4.webp)

假设有如下情况

PC1 IP: 192.168.1.1 MAC: 0C2F.B011.9D00

PC2 IP: 192.168.1.2 MAC: 0C2F.B084.6200

PC3 IP: 192.168.1.3 MAC: 0C2F.B06A.3900

PC4 IP: 192.168.1.4 MAC: 0C2F.B01F.0A00

现在 PC1 想通过 IP address 发送数据到 PC3 (192.168.1.1 -> 192.168.1.3)。因为 IP address 是一个逻辑含义的地址，PC1 想要把数据发送到 PC3 还需要 MAC address，但是现在并不知道 PC3 的 MAC address。所以 PC1 首先会发送一个 ARP request

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_16-52.1oldz0u2e6u8.webp)

ARP Request 中包含如下信息

- Src IP: 192.168.1.1
- Dst IP: 192.168.1.3
- Src MAC: 0C2F.B011.9D00
- Dst MAC: FFFF.FFFF.FFFF

其中 FFFF.FFFF.FFFF 是 broadcast MAC address(二层广播地址)

1. 当 ARP Request 到达 SW1 时，会记录 MAC .9D00 Interface G0/0 到自己的 MAC address table 中(也被称为 dynamic MAC address)。因为 Dst MAC .FFFF 是 broadcast MAC address 所以会做 broadcast，将 ARP Request 发送到除接受端口外的其他端口，即 G0/1 和 G0/2 (和 Unknown unicast 一样，但是在 ARP 中叫做 boardcast)

2. 因为 PC2 的 IP address 不匹配 192.168.1.3，所以丢弃 ARP Request
3. 当 ARP Request 到达 SW2 时，和 SW1 一样，会将 MAC .9D00 Interface G0/2 记录到自己的 MAC address table。同样会将 ARP Request 发送到除接受端口外的其他端口，即 G0/0 和 G0/1
4. 因为 PC4 的 IP address 不匹配 192.168.1.4, 所以丢弃 ARP Request
5. 因为 PC3 的 IP address 匹配 192.168.1.3，所以会回送 ARP Reply 包含
   - Src IP：192.168.1.3
   - Dst IP: 192.168.1.1
   - Src MAC: 0C2F.B06A.3900
   - Dst MAC: 0C2F.B011.9D00

6. 当 ARP Reply 到达 SW2 时，会记录 MAC .3900 Interface G0/0 到自己的 MAC address table 中，因为 Dst MAC address .9D00 在 MAC address table 中，所以转发到 G0/2

7. 当 ARP Reply 到达 SW1 时，会记录 MAC .3900 Interface G0/2 到自己的 MAC address table 中，因为 Dst MAC address .9D00 在 MAC address table 中。所以转发到 G0/0

8. 当 PC1 收到 ARP Reply 时，会将 192.168.1.3 MAC .3900 记录到自己的 ARP table 中

   例如在 Windows 长这样

   ![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_17-20.3dp67hyw64g0.webp)

> 不管是 PC 还是 Router 都有自己的 ARP table/MAC address table，但是 Switch 没有 ARP table，只有 MAC address table。因为 Switch 只是一个工作在 2 层的设备(不考虑 3 层 Switch)

如果需要查看 Cisco Router 中的 ARP table，可以使用 `show arp`

如果需要删除 Cisco Router 中的 ARP table，可以使用

## Ping

Ping 是一个发送 ICMP 报文的工具，用于测试 3 层可达性以及 RTT

![](https://github.com/dhay3/image-repo/raw/master/20230519/2023-05-19_17-25.38v6ps167874.webp)

其中

- period `.` 表示不通，因为没有对应的 ARP 条目，所以第一个报文需要做 ARP
- `!` 表示通

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=u2n762WG0Vo&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=10
[^Ethernet Header]:https://notes.networklessons.com/ethernet-header