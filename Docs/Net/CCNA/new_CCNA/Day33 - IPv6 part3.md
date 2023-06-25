# Day33 - IPv6 part3

## IPv6 Header

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-25_20-01.4ovyzn98kr5s.webp)

IPv4 header 长度可以在 20 - 60 bytes 之间，而 IPv6 header 的长度固定在 40 bytes，所以在 IPv6 header 中就没有 header length 只一个字段用于表示 header 的长度。因为报文头都是固定的，对 router 而言处理效率就比 IPv4 高

- Version

  4 bits

  Indicates the version of IP that is used

  Fixed value of 6(0b0110) to indicate IPv6

- Traffic Class

  8 bits

  Used for Qos(Quality of Service), to indicate high priority traffic

  For example IP phone traffic, live video calls,etc, will have a Traffic Class value which gives them priority over other traffic

- Flow Label

  20 bits

  Used to identify specific traffic ‘flows’(communications between as a specific source and destination)

- Payload Length

  16 bits

  Indicates the length of the payload(the encapsulated Layer 4 segment) in bytes. The length of the IPv6 header itself isn’t included, because it’s always 40 bytes

- Next Header

  8 bits

  Indicates the type of the ‘next header’(header of the encapsulated sgement), for example TCP or UDP

  Same function as the IPv4 header’s ‘Protocol’ filed

- Hop limit

  8 bits

  The value in this filed is decremented by 1 by each router that forwards it. If it reaches 0, the packet is discarded

  Same function as the IPv4 header’s ‘TTL’ field

- Source Address/Destination Address

  128 bits each

  These fields contain the IPv6 addresses of the packet’s source and the packet’s intended destination

[jeremy’s IT Lab]:https://www.youtube.com/watch?v=rwkHfsWQwy8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=63