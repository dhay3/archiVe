# Metasploit快速入门（二）

参考：

https://mp.weixin.qq.com/s?__biz=MjM5MTYxNjQxOA==&mid=2652850556&idx=1&sn=bbfae36b3cbb012fc498ab3aa20501f3&chksm=bd5935b18a2ebca785209112971dcbde12a9718c94f0a63d6c5922aef573472ec143a2b18f1c&scene=21#wechat_redirect

[TOC]

## 概述

信息收集是渗透测试中首先要做的重要事项之一，目的是尽可能多的查找关于目标的信息，我们掌握的信息越多，渗透成功的机会越大。在信息收集阶段，我们主要任务是收集关于目标机器的一切信息，比如`IP`地址，开放的服务，开放的端口。这些信息在渗透测试过程中启到了至关重要的作用。为了实现这一目的，我们将在本章学习各种扫描技术、如`SMB`扫描、`SSH`服务扫描，`FTP`扫描、`SNMP`枚举、`HTTP`扫描以及`WinRM`扫描和暴力破解。

收集信息的方式主要有三种：

1. **被动信息收集**：==这种方式是指在不物理连接或访问目标的时候（不通过发送数据包到目标）==，获取目标的相关信息，这意味着我们需要使用其他信息来源获得目标信息。比如查询`whois`信息。假设我们的目标是一个在线的Web服务，那么通过`whois`查询可以获得它的`ip`地址，域名信息，子域信息，服务器位置信息等。

2. **主动信息收集**：这种方式是只与目标建立逻辑连接获取信息，这种方式可以进一步为我们提供目标信息，让我们的目标的安全性进一步理解。在端口扫描中，使用最常用的主动扫描技术，探测目标开放的端口和服务

3. **社会工程学**：这种方式类似于被动信息收集，主要是针对人为错误，信息以打印输出、电话交谈、电子邮件等形式泄露。使用这种方法的技术有很多，收集信息的方式也不尽相同，因此，社会工程学本身就是一个技术范畴。

   社会工程的受害者被诱骗发布他们没有意识到会被用来攻击企业网络的信息。例如，企业中的员工可能会被骗向假装是她信任的人透露员工的身份号码。尽管该员工编号对员工来说似乎没有价值，这使得他在一开始就更容易泄露信息，但社会工程师可以将该员工编号与收集到的其他信息一起使用，以便更快的找到进入企业网络的方法。

## msf被动信息收集

> 进入模块后使用back退出模块，如果使用exit会退出msf

### 准备工作

我们将从公司域名开始收集信息，获取公司有关信息，收集域名，检测蜜罐，收集电子邮件地址

#### DNS记录扫描和枚举

DNS扫描和枚举模块可用于从给定的DNS服务器收集有关域名的信息，执行给中DNS查询（如域传送，反向查询，SRV记录）

1. 程序位于`auxiliary`/ɔːɡˈzɪlɪəri/模块中，进入`msfconsole`后，我们可以使用`use`命令调用我们想要的模块，我们要使用的`auxiliary/gather/enum_dns`模块。使用`use  auxiliary/gather/enum_dns` 进入模块，输入`info`可以查看模块的信息，包括作者，描述，基本配置信息等, ==通过`show options`显示具体参数==

   <img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_13-16-15.png" style="zoom:80%;" />

2. 设置具体参数，并运行模块，可以使用`show options`查看设置的参数。与`dig @<server> name any`类似

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_13-21-49.png"/>

   更多

   `dns`扫描和枚举模块也可以用于主动信息收集，通过爆破的方式，设置`ENUM_BRT`为`true`，可以通过字典暴力枚举子域名和主机名。`WORDLIST`选项可以设置字典文件。

#### CropWatch公司名称信息收集

收集公司信息也是必不可少的，我们可以使用`CorpWatch`公司名称信息搜索模块:`auxiliary/gather/corpwatch_lookup_name`，通过该模块可以收集公司的名称，地址，部门和行业信息。该模块与`CorpWatch API`连接，以获取给定公司名称的公开可用信息。

切换到`auxiliary/gather/corpwatch_lookup_name`模块，设置好公司名字，设置信息显示的数量

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_16-17-16.png"/>

Tip：此网站被Q，需要配置代理才能使用这个服务

#### 搜索引擎子域名收集器

收集子域名是寻找新目标的好办法，我们可以使用搜索引擎子域名收集模块

模块名：`auxiliary/gather/searchengine_subdomains_collector`

从`Yahoo`和`Bing`收集域名的子域信息

切换到这个模块，设置好要要查询的域名，然后运行

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_16-28-37.png"/>

#### Censys搜索

`Censys`是一个互联网设备搜索引擎，`Censys`每日通过`ZMap`和`ZGrab`扫描互联网上的主机和网站，持续监控互联网上所有可访问的服务器和设备。

我们可以使用`Censys`搜索模块，通过`Censys REST API`进行信息查询。可以检索超过100W的网站和设备信息。

Tip：如果需要使用`Censys`搜索模块，需要去https://censys.io注册获得API和密钥

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_16-29-17.png"/>

收集到了非常多的IP信息和端口信息

#### Shodan 搜索引擎

`Shodan`搜索引擎是一个付费的互联网设备搜索引擎，`Shodan`运行你搜索网站的`Banners`信息，设备的元数据，比如设备的位置，主机名，操作系统等。

Tip：同样要使用`Shodan`搜索模块，需要先去`Shodan`官网（ https://www.shodan.io）注册获取API Key。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_16-29-35.png"/>

通过`Shodan`搜索模块可以找到更多目标的信息，比如 IP 地址，开放的端口，位置信息等。

#### Shodan 蜜罐检查

检测目标是否为蜜罐，避免浪费时间或因为试图攻击蜜罐而被封锁。使用`Shodan Honeyscore Client`模块，可以利用`Shodan`搜索引擎检测目标是否为蜜罐。结果返回为`0`到`1`的评级分数，如果是`1`，则是一个蜜罐。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_16-29-58.png"/>

#### 邮箱信息收集

收集邮箱信息是渗透测试中常见的部分，它可以让我们了解互联网上目标的痕迹，以便用于后续的暴力攻击以及网络钓鱼等活动

我们可以使用`auxiliary/gather/search_email_collector`模块，该模块是利用搜索引擎获取与目标有关的电子邮件信息

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_19-28-38.png"/>

从输出信息来看，可以看到该模块利用`Google`、`Bing`和`Yohoo`搜索目标有关的电子邮件地址。

## 主动信息收集

> 更推荐使用nmap

通常来说，通过扫描进行主动信息收集，从这一步开始，我们将直接与目标进行逻辑连接。

端口扫描是一个有趣的信息收集过程，它涉及对目标系统更深入的搜索，但是由于主动端口扫描涉及对目标系统直接访问，可能会被防火墙和入侵检测系统检测到。

###### TCP 端口扫描

让我们从`TCP`端口扫描模块开始，看看我们能获取目标的哪些信息？

我们要使用的模块是`use auxiliary/scanner/portscan/tcp`

Tip：我们将利用此模块扫描渗透测试实验环境的网络，请遵守当地法律法规，请勿直接扫描互联网设备。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_19-35-05.png"/>

Tip：扫描器模块一般使用`RHOSTS`，表示扫描整个网络，而不是`RHOST`（单机, ==支持CIDR==）

当我们使用`Metasploit`模块的时候，可以使用`show options`查看所有可配置的选项，使用`show missing`查看必须要配置但尚未配置的参数。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_20-02-08.png"/>

###### TCP SYN 扫描

相对普通的`TCP`扫描来说，`SYN`扫描速度更快，因为它不会完成`TCP`三次握手，而且可以在一定程度上躲避防火墙和入侵检测系统的检测。

使用的模块是`auxiliary/scanner/portscan/syn`，使用该模块，可以指定端口范围，默认1-10000

> db_nmap方式
>
> 使用`db_nmap`的好处在于可以将结果直接存储到`Metasploit`数据库中，而不再需要`db_import`进行导入。
>
> 操作与nmap相同
>
> 具体请查看：https://www.cnblogs.com/kikochz/p/13619449.html

## 基于ARP的主机发现

通过`ARP`请求可以枚举本地网络中的存活主机，为我们提供了一种简单而快速识别目标方法。

当攻击者和目标机器处于同一个局域网时，可以通过执行`ARP`扫描发现主机

使用`ARP`扫描模块（`auxiliary/scanner/discovery/arp_sweep`），设置目标地址范围和并发线程，然后运行。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_21-16-12.png"/>

如果启动了数据库， 结果将存储在msf数据库中，可以使用`hosts`显示已经发现的主机

## UDP服务识别

`UDP`服务扫描模块运行我们检测模板系统的`UDP`服务。由于`UDP`是一个无连接协议（不面向连接），所以探测比`TCP`困难。使用`UDP`服务探测模块可以帮助我们找到一些有用的信息。

选择`auxiliary/scanner/discovery/udp_sweep`模块，设置目标范围，然后运行扫描即可

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_21-19-44.png"/>

## SMB扫描和枚举

多年来，`SMB`协议（一种在 Microsoft Windows系统中使用网络文件共享的协议）已被证明是最容易被攻击的协议之一，它允许攻击者枚举目标文件和用户，甚至远程代码执行。

使用无需身份验证的`SMB`共享枚举模块，可以帮助我们收集一些有价值的信息，比如共享名称，操作系统版本等。

模块名：`auxiliary/scanner/smb/smb_enumshares`

- `SMB`共享枚举模块在后续的攻击阶段也非常有用，通过提供凭据，可以轻松的枚举共享和文件列表

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_21-41-59.png"/>

`Metasploit`还提供其他的一些`SMB`扫描模块，让我们看看其他模块的用法。

- `SMB`版本检测模块可以检测`SMB`的版本

  `smb_version`

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_21-42-34.png"/>

- 用户枚举模块可以通过`SAM RPC`服务枚举哪些用户存在

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_21-46-21.png"/>

- `SMB`登录检测模块可以测试`SMB`登录,需要设置密码和用户名

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_21-48-07.png"/>

> 其他的模块，都在 `auxiliary/scanner/smb/`中，可以敲 `TAB`键查看，你可以一个个学习，这里就不一一举例讲解。

## FTP扫描

使用`auxiliary/scanner/ftp/ftp_version`模块，设置好扫描范围和线程，就可以运行扫描了。==如果扫描过慢请设置threads==

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-08_23-53-51.png"/>

与之前的扫描一样，扫描结果会保存到数据库中，可以使用`services`命令查看已经检测的服务信息。

## HTTP扫描

超文本传输协议（`HTTP`）是一个应用层协议，它是万维网通信的基础。它被众多的应用程序使用，从物联网（IoT）设备到移动应用程序。它也是搜索漏洞的好地方。

- 准备工作

  `HTTP SSL`证书检测模块可以检测`Web`服务器的证书。

  `Robots.txt`内容检测模块可以搜索`robots.txt`文件并分析里面的内容。

  如果服务端允许未授权的`PUT`请求方法，则可以将任意的`Web`页面插入到网站目录中，从而导致执行破坏性的代码或者往服务器填充垃圾数据，从而造成拒绝服务攻击。

  `Jenkins-CI HTTP`扫描模块可以枚举未授权的`Jenkins-CI`服务。

- 使用模块：`use auxiliary/scanner/http/cert`检测目标的`HTTP SSL`证书

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_00-39-14.png"/>

- 使用模块：`use auxiliary/scanner/http/http_put`检测目标是否开启PUT，相同的还有DELETE

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_00-45-54.png"/>

## SHH扫描

`SSH`是一个广泛使用的远程登录程序。它使用强大的加密提供身份认证和保证机密性。在本节中，我们将通过`SSH`版本扫描模块，确定目标使用的`SSH`版本，确定是否为易受攻击的`SSH`版本，如果是，我们可以利用它。

在之前的扫描中，我们发现目标机器开放了`TCP` `22`端口，这也是`SSH`的默认端口，我们用`SSH`版本探测模块来获取目标系统上运行的`SSH`版本信息。

1、模块名称：`auxiliary/scanner/ssh/ssh_version`

```
msf5 > use auxiliary/scanner/ssh/ssh_version
msf5 auxiliary(scanner/ssh/ssh_version) > set RHOSTS 192.168.177.144
RHOSTS => 192.168.177.144
msf5 auxiliary(scanner/ssh/ssh_version) > run

[+] 192.168.177.144:22    - SSH server version: SSH-2.0-OpenSSH_7.1 ( service.version=7.1 service.vendor=OpenBSD service.family=OpenSSH service.product=OpenSSH service.cpe23=cpe:/a:openbsd:openssh:7.1 service.protocol=ssh fingerprint_db=ssh.banner )
[*] 192.168.177.144:22    - Scanned 1 of 1 hosts (100% complete)
[*] Auxiliary module execution completed
msf5 auxiliary(scanner/ssh/ssh_version) >
```

当然这里的`RHOSTS` 选项也可以指定为网络地址，从而扫描整个网段。

获取版本信息之后，我们就可以搜索该版本的漏洞。

2、测试常用口令登录`SSH`，可以使用`SSH`登录测试模块, ==使用自己提供的密码字典或是msf自带的字典==`cd /usr/share/metasploit-framework/data/worklists`

```
msf5 > use auxiliary/scanner/ssh/ssh_login
msf5 auxiliary(scanner/ssh/ssh_login) > set RHOSTS 192.168.177.144
RHOSTS => 192.168.177.144
msf5 auxiliary(scanner/ssh/ssh_login) > set USERNAME user
USERNAME => user
msf5 auxiliary(scanner/ssh/ssh_login) > set PASS_FILE /root/password.lst
PASS_FILE => /root/password.lst
msf5 auxiliary(scanner/ssh/ssh_login) > run

[*] Scanned 1 of 1 hosts (100% complete)
[*] Auxiliary module execution completed
```

3、如果登录成功了，可以用`sessions` 查看会话和与目标进行会话交互

<img src="https://mmbiz.qpic.cn/mmbiz_png/3RhuVysG9LcBK2lYNichibhWRoB60R5m7DBsClknyF1Lg1ffe376eWLkb19aXZcIbiagqaSkic7jRm53iaEUlEqNFeA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1"/>





