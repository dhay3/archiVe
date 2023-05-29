# Day12 - The Life of a Packet

以如下拓扑为例

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230524/2023-05-25_16-05.6gesqr387aww.webp)

橙色框内表示 MAC 地址，例如 PC1 MAC 1111，R1 Gi0/2 aaaa

现在 PC1 192.168.1.1 需要访问 PC2 192.168.4.1

> 这里只考虑 PC1 - SW1 - R1 - R2 - R4 - SW4 - PC4 的链路，走 R3 同样也可以

1. 因为 Dst IP 是 192.168.4.1 和 192.168.1.1 不在一个段，且在 ARP table 中没有对应的条目，所以需要通过 Gateway 或者是 next-hop 发送出去

2. 这时会使用 ARP，对应的 ARP request 报文为

   Src IP: 192.168.1.1

   Dst IP: 192.168.1.254

   Dst MAC: ffff.ffff.ffff

   Src MAC: 1111

3. 当 SW1 收到 ARP request 后因为 Dst MAC 是 ffff.ffff.ffff 广播地址。所以会将 ARP request，广播到除接收端口外的其他端口,例子中为 Gi0/0。并将 PC1 Gi0/0 对应的 MAC 记录到自己的 MAC address table 中

4. 当 R1 收到 ARP request，de-encapsulation 2 层帧头，因为 Dst IP 和自己的 Gi0/2 端口配置的 IP 匹配。所以会回送 ARP reply

   Src IP: 192.168.1.254

   Dst IP: 192.168.1.1

   Dst MAC: 1111

   Src MAC: aaaa

   并将 Src IP 和 Src MAC 记录到自己的 ARP table 中

5. 当 SW1 收到 ARP reply，会将 R1 Gi0/2 对应的 MAC 记录到自己的 MAC address table 中，并对 ARP reply 做 unicast

6. PC1 收到 ARP reply，则发送的报文为

   Src IP: 192.168.1.1

   Dst IP: 192.168.4.1

   Dst MAC: aaaa

   Src MAC: 1111

   > 注意报文的原始 3 层 IP address 是不会改变的，只有 2 层 MAC address 会改成对应接口 MAC

   并将 Dst MAC 和 Dst IP 记录到自己的 ARP table 中

7. R1 收到 PC1 发过来的报文，查看自己的 routing table

   有一条如下路由

   ```
   		R1 Routing Table
   Destination				Next Hop
   192.168.4.0/24		192.168.12.2
   ```

   192.168.4.1 匹配 192.168.4.0/24 的路由，通过 Gi0/0 端口转发到 192.168.12.2

8. 但是 R1 现在同样也不知道 R2 的 MAC address，所以需要通过 ARP request，报文为

   Src IP：192.168.12.1

   Dst IP: 192.168.12.2

   Dst MAC: ffff.ffff.ffff

   Src MAC: bbbb

   > 注意这里 ARP request 只会发送到 192.168.12.0/24 这个段对应的端口，因为 ARP 是在 LAN 中的(广播域)

9. 当 R2 收到 ARP request 时 de-encapsulation 2 层帧头，因为 Dst IP 192.168.12.2 匹配 Gi0/0 配置的 IP，所以回 ARP reply

   Src IP: 192.168.12.2

   Dst IP: 192.168.12.1

   Dst MAC: bbbb

   Src MAC: cccc

​		并将 Src IP 和 Src MAC 记录到自己的 ARP table 中

10. R1 收到 ARP reply，并将 Dst MAC 和 Dst IP 记录到自己的 ARP table 中，并发送报文

    Src IP: 192.168.12.1

    Dst IP: 192.168.4.1

    Dst MAC: cccc

    Src MAC: bbbb

11. R2 收到报文后，查看自己的 routing table

    有一条如下路由

    ```
    		R1 Routing Table
    Destination				Next Hop
    192.168.4.0/24		192.168.24.4
    ```

    192.168.4.1 匹配 192.168.4.0/24 的路由，通过 Gi0/1 端口转发到 192.168.24.4

12. 依次类推...



**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=aHwAm8GYbn8&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=19
[^ARP]:https://help.ui.com/hc/en-us/articles/115005984408-Intro-to-Networking-Address-Resolution-Protocol-ARP-