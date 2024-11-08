# TCP 多路复用和多路分解

参考：

https://my.oschina.net/u/731676/blog/268106

https://en.wikipedia.org/wiki/Multiplexer

- 多路复用（multiplexing）：将主机不同socket中收集的数据加上运输层的首部从而生成报文，然后发送到网络层的过程。叫做多路复用

  ```mermaid
  graph LR
  a(socket data segment01)--> 运输层-->网络层
  a=b(socket data segment02)--> 运输层
  ```

  

- 多路分解（demultiplexing）：将运输层报文段中的数据交付到正确的socket的过程叫多路分解。

  ```mermaid
  graph RL
  packet -->运输层--> a(socket data segment01)
  运输层--> b(socket data segment01)
  ```

  

