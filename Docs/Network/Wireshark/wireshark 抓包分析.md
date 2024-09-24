# wireshark 抓包分析

> length包括数据包首部

## 0x001

![](/home/cpl/note/imgs/_Wireshark/2021-07-23_13-43.png)

- 74：A到B SYN ，seq随机生成x

- 78：B到A SYN ACK，seq随机生成y，ack = x + 1

- 79：A到B ACK，seq = x + 1，ack = y + 1

- 97：A到B PSH，由于len(78)=0，所以ack等于seq(78) + 1；由于len(74)=0，所以seq等于seq(74)+1

- 107：B到A ACK，由于len(97)=7，所以ack等于seq(97) + 7；由于len(78)=0，所以seq等于seq(78)+1

- 108：http请求

- 109：由于len(97)=7，所以seq等于seq(97)+7；ack=seq(107) + data_length(316)

  370 = 20(tcp首部) + 20(ip首部) + 14(mac首部) + data(316)



























