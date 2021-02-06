# Ettercap DNS Spoofing

参考：

https://pentestmag.com/ettercap-tutorial-for-windows/

一句代码

这里的remote表示嗅探远程网关

> ettercap -Tp -P dns_spoof -M arp remote /192.168.80.129//     /192.168.80.2//

目标机192.168.80.129，gw 192.168.80.2，攻击机 192.168.80.200

我们将baidu.com 解析到 taobao.com

1. dig 搜索taobao.com的服务器地址

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_21-22-56.png"/>

   这我们选用140.205.94.189

2. 修改配置ettercap配置文件

   > 注意新版的etter.dns 不在 /usr/shar/ettercap/etter.dns， 我们通过locate etter.dns来查看
   >
   > ettercap不支持NS记录，修改配置文件后需要重启ettercap

   在最后将baidu.com的A记录指向140.205.94.189

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_21-25-43.png"/>

3. 设置ettercap

   选中dns_spoof模块

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_20-04-14.png"/>

   添加目标，target1 目标机，target2 gw。因为解析dns是通过gw发送请求

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_21-28-45.png"/>

   all set 执行ettercap

4. 登入目标机，访问baidu.com

   > 如果不生效，请将目标机的dns缓冲清空，ipconfig /flushdns

   发现firefox提示有安全问题，由于baidu.com使用https协议

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_21-29-43.png"/>

   但是发现其实已经将域名解析到taobao的服务器了

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_21-30-47.png"/>

5. 替换https为http，我们这里使用360doc.com

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_23-37-59.png"/>

6. 目标机访问360doc.com，就会访问攻击机的/var/www/html，这样也就实行了dns spoof

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-17_23-39-39.png"/>

   

   

