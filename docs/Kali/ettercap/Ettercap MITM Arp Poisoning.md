# Ettercap MITM Arp Poisoning

参考：

https://pentestmag.com/ettercap-tutorial-for-windows/

一句代码

这里arp欺骗攻击机所在的局域网

> ettercap -T -M arp  /// ///

1. 两台目标机IP地址分别为`192.168.80.128`和`192.168.80.129`；攻击机kali IP`192.168.80.200`

2. 伪造攻击机的MAC地址`macchanger -r eth0`随机生成一个MAC地址，我们通过`macchanger -s eth0`产看etho interface的MAC地址

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_18-16-20.png"/>

3. 启动图形界面`ettercap -G`，设置本地的eth0 interface为监听IP，打勾即可启动ettercap。

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_18-23-01.png" style="zoom:80%;" />

​	先扫描LAN中的host，然后分别将host-id 128，129 分别添加为target1  和 target2

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_18-23-56.png"/>

然后选择`MITM ARP poisoning`，选择sniffing remote connect

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_18-25-59.png"/>

当用192.168.80.128 ping 192.168.80.129时，目标机的arp cache就会将192.168.80.129的MAC地址替换为攻击的MAC地址（==两台目标机所有的traffic都会被ettercap嗅探到==）。我们可以通过wireshark 来看一下，这里的HitronTe就是MAC地址的前两个字节

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_18-35-16.png"/>

我们使用`arp -a`来查看目标机的arp cache，就可以发现出现两个MAC地址相同的IP

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_19-18-43.png"/>

> ettercap 同时能嗅探账号和密码
