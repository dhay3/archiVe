# OSI MODEL & TCP/IP SUITE

## What is a networking model

> 需要注意得一点是，networking protocols 是逻辑含义上的规则而不是物理意义上的标准

![](https://github.com/dhay3/image-repo/raw/master/20230516/2023-05-17_16-04.2y85tenknhmo.webp)

设想一下如果没有统一标准的网络模型和协议，每个厂商都用自己的网络标准去生产设备，那么不同厂商之间的设备就不能互相通信

## OSI model

Open Systems Interconnection model (OSI model) 是一个由 International Organization for stanardization (ISO) 定义的逻辑模型，用于对网络协议进行归类，主要分为 7 层

![](https://github.com/dhay3/image-repo/raw/master/20230516/2023-05-17_16-13.16on9011n37k.webp)

- Application Layer = Layer 7
- Presentation Layer = Layer 6
- Session Layer = Layer 5
- Transport Layer = Layer 4
- Network Layer = Layer 3
- Data Link Layer = Layer 2
- Physical Layer = Layer 1

### Application Layer

- This Layer is closest to the end user
- Interacts with software applications, for example your web browser
- HTTP and HTTPS are Layer 7 protocols

### Presentation Layer

- Translate the data in the application layer to different format to be sent over the network

- encryption of data as it is sent, and decryption of data as it is recieved.

### Session Layer

- Controls dialogues(sessions) between communicating hosts
- Establishes, manages, and terminates connections between the local application and the remote application

### Transport Layer

- Segments and reassembles data for communications between end hosts
- Breaks large pieces of data into smaller segments which can be more easily sent over the network and are less likely to cause transmission problems if errors occur.

### Network Layer

- Provides connectivity between end hosts on different networks (ie. outside of the LAN)
- Provides logical addressing (IP adresses)
- Provides path selection between source and destination
- Routers operate at Layer 3 

### Data Link Layer

- Provides node-to-node connectivity and data transfer (for example PC to switch, switch to router, router to router)
- Defines how data is formatted for transmission over a physical medium (for example, copper UTP cables)
- Detects and (possibly) corrects Physical Layer errors
- Uses Layer 2 addressing, separate from Layer 3
- Switches operate at Layer 2

### Physical Layer

- Defines Physical characteristics of the medium used to transfer data between devices. For example, voltage level, maximum transmission distances, physical connectors, cable specifications, etc.
- Convert digital bits to eletrical or radia signals
- Cables,PIN,layouts,etc are related to the Physical Layer  

## TCP/IP Suite

和 OSI model 类似，但是只有 4 层

![](https://github.com/dhay3/image-repo/raw/master/20230516/2023-05-17_17-04.52frgbir6g00.webp)

![](https://github.com/dhay3/image-repo/raw/master/20230516/2023-05-17_17-06.lbg0rgu4yy8.webp)

## PDU

Protocol Data Unit(PDU) is a complete message (message 中文叫报文，包含 header) of the protocol at that layer

![](https://github.com/dhay3/image-repo/raw/master/20230516/2023-05-17_16-45.79jqh55rgghs.webp)

- L4 中 PDU 被称为 Segment，作为 L3 的 data
- L3 中 PDU 被称为 Packet, 作为 L2 的 data
- L2 中 PDU 被称为 Frame (L1 没有 Header，只有 Bit)
- L1 中 PDU 被称为 Bit

## Encapsulation

中文也叫做封包, 从上层来的 data 会加上本层的 hearder

## De-Encapsulation

中文也叫做解封包，从下层来的 message 会去掉本层的 header



**references**

[^jeremy’s IT Lab]: https://www.youtube.com/watch?v=t-ai8JzhHuY&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=6
[^wiki]: https://www.google.com.hk/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwizxNWm-vv-AhWp-zgGHa82BE0QFnoECBIQAQ&url=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FProtocol_data_unit&usg=AOvVaw1xFrzOAAGswdmyP092g5tO
