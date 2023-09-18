# Day30 - TCP & UDP



## Functions of Layer 4

- Provides transparent transfer of data between end hosts

  *The hosts themselves aren’t aware of the details of the underlying network, the transfer of data is ‘transparent’ to them*

- Provides(or doesn’t provide) various services to applications

  - reliable data transfer
  - error recovery
  - data sequencing
  - flow control

- Provides Layer 4 addressing(port numbers)

  由 IANA(Internet Assigned Number Authority) 提供，分为 3 大类

  1. Well-knwon prot numbers 0 - 1023
  2. Registered port numbers 1024 - 49151 
  3. Ephemerla(短暂的)/private/dynamic port numbers 491252 - 65535

  主要有如下两个功能

  - Identify the Application Layer protocol
  - Provides session multiplexing

## What is a sessioin

*A session is an exchange of data between two or more communicating devices*

> 在 TCP 中相同 4 元组构成一个 Session

例如下图中的

PC1 访问 SRV1 80 端口

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-19_22-32.6t9vhjttgo00.webp)

在 PC1 中所有 Src:50000 Dst:80 的报文，以及在 SRV1 中所有 Src:80 Dst:50000 的报文都在一个 Session 中，可以简单的理解成相同 4 元组(源目 IP 及端口)的所有报文构成一个 Session

## TCP

Transmission Control Protocol(TCP)

有如下几个特性

- TCP is connection-oriented

- TCP provides reliable communication
- TCP provides sequenciing
- TCP provides flow control

### TCP is connection-oriented

中文也叫做面向连接

*Before actually sending data to the destination host, the two hosts communicate to establish a connection. Once the connection is established, the data exchange begins*

Connection-oriented 体现在两点

1. Establishing Connections
2. Terminating Connections

#### Establishing Connections

在 TCP 中通过 Three-Way handshake 来建立连接

例如 PC1 想要访问 SRV1 http

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_13-41.jdqddpnz39c.webp)

1. PC1 will send a TCP segment to SRV1 with the SYN flag set, meaning that bit is set to 1
2. SRV1 will reply by sending a TCP segment to PC1 with the SYN and ACK flags set. So both bits are set to 1
3. Finally, PC1 will send a TCP segment with the ACK bit set

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_13-44.1edq1qwrzkow.webp)

#### Terminating Connections

在 TCP 中通过 Four-Way handshake 来终止连接

例如 PC1 不再需要和 SRV1 连接

1. First, PC1 sends a TCP segment to SRV1 with the FIN flag set
2. SRV1 reponds with an ACK
3. SRV1 then sends its own FIN
4. Finally, PC1 sends an ACK in response to SRV1’s FIN, and the connection is terminated

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_13-56.759cdpy3jkao.webp)



### TCP provides reliable communication

*The destination host must acknowledge that it received each TCP segment(Layer 4 PDU)*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-04.4bqnlhwp8jgg.webp)

*If a segment isn’t acknowledged, it is sent again*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-07.2rcx2d8yfc3k.webp)

### TCP provides sequencing

*Sequence numbers in the TCP header allow destination hosts to put segments in the correct order even if they arrive out of order*

### TCP provides flow control

*The destination host can tell the source host to increase/decrease the rate that data is sent*

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-10.1n06i0y4fq5c.webp)

### TCP Header

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_13-34.4gm0pqyfm0ao.webp)

CCNA 中不需要了解 TCP header 中所有的字段，主要关注如下几个字段

- Source Port/Destination Port

  每个字段 16 bits = 2 bytes，所以最大就是 $2^{16} = 65536$

- Sequence number/Acknowledgement number

  32 bits = 4 bytes，所以最大就是 $2^{32}$

  *These two fields provide sequencing and reliable communication*

- ACK/SYN/FIN

  These three flags are used to establish and terminate connections

- Window Size

  The windows size is used for flow control

## UDP

User Datagram Protocol(UDP)

有如下几个特性

- UDP is not connection-orinted
- UDP does not provide reliable communication
- UDP does not provide sequencing
- UDP does not provide flow control

### UDP is not connection-oriented

*The sending host does not establish a connection with the destination host before sending data. The data is simply sent*

### UDP does not provide reliable communication

*When UDP is used, acknowledgements are not sent for received segments. If a segment is lost, UDP has no mechanism to re-transmit it.Segments are sent ‘best-effort’*

### UDP does not provide sequencing

*There is no sequence number field in the UDP header. If segment arrive out of order, UDP has no mechanism to put them back in order*

### UDP does not provide flow control

*UDP has no mechanism like TCP’s window sizze to control the flow of data*

### UDP header

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-18.797fe0fkrocg.webp)

## TCP vs UDP

- TCP provides more features than UDP, but at the cost of addition overhead

  因为 TCP header 比 UDP header 要大，所以传输的消费也大

- For applications that require reliable communications(for example dowloading a file), TCP is preferred

  下载文件时，没人喜欢缺页

- For applications like real-time voice and video UDP is preferred

  因为这些场景对延时十分敏感, 使用 TCP 因为报文头比较大传输速率会比 UDP 慢。另外也不需要 TCP 重传的功能，即使丢包了，也可以要求对面的人重新复述

- There are some applications that use UDP, but provide reliability etc within the application itself

- Some applications uses both TCP & UDP, depending on the situation

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-20_14-26.35ytbwitry9s.webp)

**references**

[^jeremy’s IT LAB]:https://www.youtube.com/watch?v=LIEACBqlntY&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=57