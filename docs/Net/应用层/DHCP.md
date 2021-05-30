# DHCP

互联网广泛使用的动态主机配置协议==DHCP (Dynamic Host Configuration Protocol)==提供了即插即用连网 (plug-and-play networking) 的机制。

这种机制允许一台计算机加入新的网络和获取IP 地址而不用手工参与。

- 需要 IP 地址的主机在启动时就向 DHCP 服务器广播发送发现报文（DHCPDISCOVER），这时该主机就成为 DHCP 客户。
- 本地网络上所有主机都能收到此广播报文，但只有 DHCP 服务器才回答此广播报文。
- DHCP 服务器==先在其数据库中查找该计算机的配置信息。==若找到，则返回找到的信息。若找不到，则从服务器的 IP 地址池(address pool)中取一个地址分配给该计算机。DHCP 服务器的回答报文叫做提供报文（DHCPOFFER）。 

<img src="..\..\..\..\imgs\_Net\计算机网络\Snipaste_2020-09-05_23-23-06.png"/>

1. DHCP服务器被动代开UDP端口67，等待客户端发来的报文

2. DHCP客户从UDP端口68发送DHCP发现保卫呢。

3. 凡收到DHCP发现报文的DHCP服务器都能发出DHCP提供报文，因此DHCP客户可能收到多个DHCP提供报文

4. DHCP客户从几个DHCP服务器中选择其中的一个，并向所选择的DHCP服务器发送DHCP请求报文

5. DHCP服务器发送确认报文DHCPPACK，进入已绑定状态，并可开始使用得到的临时IP地址了。

   DHCP客户现在要根据服务器提供的租用期T设置两个计时器T1和T2，他们的超时时间分别为0.5T和0.875T。当超时时间到达就要请求更新租用期。

6. 租用期过了一半（T1时间到了），DHCP发送请求报文DHCPREQUEST要求更新租用期

7. DHCP服务器若同意，则发回确认报文DHCPPACK。DHCP客户得到了新的租用期，重新设置计时器。

   DHCP服务器若不同意，则发回否认报文DHCPPACK。这是DHCP客户必须立即停止原来的IP地址，而必须重新申请IP地址（回到第2步）

8. 若DHCP服务器不响应第6步的请求报文DHCPREQUEST，则在租用期过了87.5%时，DHCP客户必须重新发送请求报文DHCPREQUEST(重复第6步)，然后又继续后面的步骤

9. DHCP客户可随时提前终止服务器提供的租用期，这时只需向DHCP服务器发送释放报文DHCPRELEASE即可
