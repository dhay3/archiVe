# `Metasploit 快速入门(一)

参考:

https://mp.weixin.qq.com/s?__biz=MjM5MTYxNjQxOA==&mid=2652850523&idx=1&sn=38faa318148cf1c0bfa75ca84d9b114b&chksm=bd5935968a2ebc80fd07d10a83c72a465fe8beed36c05e8fdf3f0345ad9b6c8b9971f448b500&scene=21#wechat_redirect

[TOC]

## 框架和相关术语简介

**Metasploit Framework**；这是一个免费的、开源的渗透测试框架。它是使用Ruby语言编写的。它拥有世界上最大渗透测试攻击数据库。

**Vunlerability**：允许攻击者入侵或危害系统安全性的弱点称为漏洞，漏洞可能存在于操作心痛，应用软件设置网络协议中。

**Exploit**：攻击代码或程序，它允许攻击者利用易受攻击的系统并危害其安全。每个漏洞都有相对应的漏洞利用程序。

**Payload**：攻击载荷。它主要用于建立攻击者和受害机直接的连接。

**Module**：模块是一个完整的构件，每个模块执行特定的任务，并通过几个模块组成一个单元运行。这种架构的好处是可以容易的将自己写的利用程序和工具集成到架构中

Metasploit框架具有模块的体系结构

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_12-09-16.png"/>

Metasploit提供两种不同的UI，msfconsole和WebUI，本文中主要使用msfconsole。因为msfconsole对metasploit支持最好，可以使用所有功能。

msf使用postgresql，所以需要开启服务

```shell
systemctl start postgresql && systemctl enable postgresql
```

==ttl中输入msfconsole启动msf==，输入db_status检查数据库连接情况

> 使用help 查看所有命令
>
> 命令后带上 -h 参数查看命令具体信息

## 创建工作区

msf中有工作区的概念，可以用来隔离不同的渗透测试任务，从而避免混肴不同的测试。

默认工作区是default，输入workspace查看

```shell
msf5 > workspace
* default
msf5 > 
```

输入workspace -h 查看命令帮助

```
msf > workspace -h

Usage:

    workspace                  List workspaces

    workspace -v               List workspaces verbosely

    workspace [name]           Switch workspace

    workspace -a [name] ...    Add workspace(s)

    workspace -d [name] ...    Delete workspace(s)

    workspace -D               Delete all workspaces

    workspace -r <old> <new>   Rename workspace

    workspace -h               Show this help information

msf >
```

## 使用数据库

1. 准备工作

   输入`db_import`，查看支持的文件类型

2. 怎么做

   `nmap -Pn -A -oX report 192.168.177.139`，使用nmap将扫描结果以xml形式存储，可以直接在msf控制台中使用

3. 导入到数据库

   `db_import /root/report`

## hosts

数据库中有了数据，就可以使用hosts命令来显示==当前工作区中存储的所有主机==

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_12-40-58.png"/>

## services

services命令显示目标主机上可用的服务

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_12-45-39.png" style="zoom:80%;" />

## search

search命令用来搜索相关的模块, 没有带参数表示匹配所有字段

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_09-12-57.png"/>

```
search mysql type:exploit #查找mysql的攻击模块
search cve:2008 name:linux #查找linux cve2008库中漏洞
search path:shell type:payload platform:windows#查找平台为windows路径中包含shell的攻击载荷
search
```

## show

- show info

  显示当前模块的信息

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_09-27-23.png"/>

- show options

  显示当前模块所有可以设置的参数, 会回显

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_09-23-26.png"/>

- show missing 

  显示当前为设置的必要参数

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_09-24-45.png"/>

- show targets 

  显示漏洞针对的主机版本

- show payload/exploit/auxiliary

  同search,但是没有search好用

## connect

连接目标主机, 发送请求

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_09-29-45.png"/>

连接baidu.com 443端口发送get请求

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_09-32-10.png"/>

`connect -s -z baidu.com 443`只连接不返回响应报文

## creds

显示当前数据库中所有的信息登入凭证信息

具体查看 `creds -h`

## 数据库操作

通过`help database`查看

- `db_connect`

  连接postgresql

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_12-19-47.png"/>

- `db_export `

  将数据库中所有信息导以指定格式导出指定位置, 只能是xml和pwdump格式

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_12-13-27.png"/>

- `db_import <filename>`

  将xml文件导入msf当前工作空间中

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_12-16-19.png"/>

> 可以将nmap扫描的结果存入到msfdb中
>
> nmap baidu.com -oX /opt/test.xml
>
> db_import /opt/test.xml

## 其他命令

### use

使用具体的模块, 一般配合search使用`use <name|index>`, 可是使用索引或是模块名

### back

退出具体模块, 如果在模块中使用exit将退出msf

### set

设置参数的具体内容

> 使用unset解除当前绑定的设置

### run

执行当前模块， 等价于`exploit`，

==`run -j`以daemon方式在当前工作空间中运行，通过`sessions`命令可以查看==

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_11-33-39.png"/>

退出目标机终端`ctrl+c`, 不保留session

### sessions

`session -K` 关闭所有会话

`session -k` 关闭指定session ID或范围的会话

`session -i` 与指定session交互

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_11-38-18.png"/>

`session -v` 以详细信息显示

### msfdb

`msfdb reinit`删除数据库并初始化

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_12-22-22.png"/>

### jobs

查看后台运行的模块

