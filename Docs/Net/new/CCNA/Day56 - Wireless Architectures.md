# Day56 - Wireless Architectures

## 802.11 Frame

802.11 报文，会根据 802.11 的版本，有不同的字段，但是大体上字段如下

> 对比 Ethernet 报文，802.11 报文更复杂

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_15-07.2w1hln0m8ue0.webp)

-  Frame Control

  Provides information such as the message type and subtype

- Duration/ID

  Depending on the message type, this field can indicate

  1. the time(in microseconds) the channel will be dedicated for transmission of the frame
  2. and identifier for the association(connection)

- Addresses

  Up to four addresses can be present in an 802.11 frame. Which addresses are present, and their order, depends on the message type

  1. Destination Address(DA)

     Final recipient of the frame

  2. Source Address(SA)

     Original sender of the frame

  3. Receiver Address(RA)

     Immediate recipient of the frame

  4. Transmitter Address(TA)

     Immediate sender of the frame

- Sequence Control

  Used to reassembel fragments and eliminate duplicate frames

- QoS Control

  Used in QoS to prioritize certain traffic

- HT(High Throughput) Control

  Added in 802.11n to enable High Throughput operations

  1. 802.11n is also known as ‘High Throughput’ (HT) Wi-Fi
  2. 802.11ac is also known as ‘Very High Throughput’ (VHT) Wi-Fi

- FCS(Frame Check Sequence)

  Same as in an Ethernet frame, used to check  for errors

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_15-20.4a8i0f4spug0.webp)

## AP deployment



**references**

1. ^https://www.youtube.com/watch?v=uX1h0F6wpBY&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=110