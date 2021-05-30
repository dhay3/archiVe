# Ettercap dos_attack

一句代码

> ettercap -T -p dos_attack

ettercap 会提示你输入目标机的IP和一个为使用的IP

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-18_22-58-06.png"/>

1. 选中dos_attack插件，输入victim IP

   <img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-18_23-00-52.png" style="zoom:80%;" />

2. 输入一个unuse IP，这里可不在同一网段

   <img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-18_23-05-02.png" style="zoom:80%;" />

3. 开启一台虚拟机，将IP设置为192.168.80.129，我们可以通过wireshark发现攻击机会一直向目标机发包

<img src="..\..\..\imgs\_Kali\ettercap\GIF.gif"/>

