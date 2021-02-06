# burpsuite 不能抓取其他虚拟机环回地址

在使用burpsuite抓起其他虚拟机的环回地址时，发现无法抓取环回地址

<img src="..\..\..\imgs\_Kali\burpsuite\Snipaste_2020-09-13_11-15-56.png"/>

根据提示，经过一番思考后发现问题在于URL的路径。当访问`localhost/sqli/index-1.html`时，经过burp所在的虚拟机的appache服务器，==直接将localhost解析成立当前主机的环回地址，而不是sqli-labs所在的主机==。所以将访问路径改为`192.168.80.129/sqli/index-1.html`即可

<img src="..\..\..\imgs\_Kali\burpsuite\Snipaste_2020-09-13_11-21-08.png"/>