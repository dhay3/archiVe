# windows server 2008/win7 远程控制

[TOC]

## 划分同一局域网

客户端



<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_22-05-32.png" />

服务端

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_22-06-06.png" />

==这里将网段划分的尽量小，通信会快一点==

## 远程桌面协议/RDP

remote desktop protocal

### 一. 将用户添加到远程用户组

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_22-04-14.png" style="zoom:80%;" />

### 二. 连接服务端

运行窗口`mstsc`，输入服务器IP，账户密码

## telnet

### 一. 配置服务端Telnet

使用telnet无需配置远程服务，即可访问服务器。点击服务器管理器，添加功能

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_22-11-12.png"/>

`servieces.msc`进入服务管理器，找到Telnet将服务启动类型设置为自动，并将服务启动，添加用户到Telnet Clients

### 二. 客户端接入服务端

`telnet socket  ||  telnet IP`，==默认端口23==，可以扫描端口采用中间人攻击

输入用户名，密码；进入服务器Dos窗口

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_22-16-10.png" style="zoom:80%;" />



