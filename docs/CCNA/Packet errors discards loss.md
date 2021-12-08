# Packet errors discards loss

reference：

https://www.auvik.com/franklyit/blog/packet-errors-packet-discards-packet-loss/

## Packet errors

Packet errors 包括两种

- Transimission errors

  包在传输的过程中 damaged

- Format errors

  接受端收到的包格式与预期不一致

如果网络设备是通过bad cable，bad port，broken fiber cable，or dirty fiber connector 连接网络的，包就很容易损坏（包通过无线连接网络是非常容易损坏）。如果 错包了，==TCP会重传（Transmission Control Protocol）==，因为TCP是确保无Packet errors的。但是UDP即使Packet errors也会继续传输。

## Packet discards

Packet discards通常由于包的格式不正确，或者接受端没有足够的内存去处理packets。

虽然Packet discards是不可避免的，但是过度丢包通常指向几个问题：

- network device配置错误。例如P2P的VLAN；如果一个配置了VLAN9，另一个配置了VLAN10，从VLAN9的traffic就会被丢弃
- 端口没有足够的bandwidth
- network device没有足够的CPU或memory usage

## Packet loss

Packet loss 和 Packet discard 类似，但是Packet loss 发生在包到达 destination 之前，而Packet Disard 发生在包到达 destination 之后。例如下载一部电影原本需要30secs，但是现在需要45secs

