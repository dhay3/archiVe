# Dig

参考：

http://www.ruanyifeng.com/blog/2018/05/root-domain.html

http://www.ruanyifeng.com/blog/2016/06/dns.html

[TOC]

## 概述

如果没有参数，默认查询根域

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_16-01-04.png" style="zoom:80%;" />

这里显示根域有13个NS记录，对应13台DNS服务器的域名，然后向这13台DNS服务器查询

虽然只需要返回一个IP地址，但是DNS的查询过程非常复杂，分成多个步骤。

工具软件`dig`可以显示整个查询过程。

```shell
$ dig math.stackexchange.com
```

上面的命令会输出六段信息。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061501.png"/>

第一段是查询参数和统计。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061502.png"/>

第二段是查询内容。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061503.png"/>

上面结果表示，查询域名`math.stackexchange.com`的`A`记录，`A`是address的缩写。

第三段是DNS服务器的答复。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061504.png"/>

上面结果显示，`math.stackexchange.com`有四个`A`记录，即四个IP地址。`600`是TTL值（Time to live 的缩写），表示缓存时间，即600秒之内不用重新查询。

第四段显示`stackexchange.com`的NS记录（Name Server的缩写），即哪些服务器负责管理`stackexchange.com`的DNS记录。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061505.png"/>

上面结果显示`stackexchange.com`共有四条NS记录，即四个域名服务器，向其中任一台查询就能知道`math.stackexchange.com`的IP地址是什么。

第五段是上面四个域名服务器的IP地址，这是随着前一段一起返回的。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061506.png"/>

第六段是DNS服务器的一些传输信息。

<img src="http://www.ruanyifeng.com/blogimg/asset/2016/bg2016061514.png"/>

上面结果显示，本机的DNS服务器是`192.168.1.253`，查询端口是53（DNS服务器的默认端口），以及回应长度是305字节。

如果不想看到这么多内容，可以使用`+short`参数。

```shell
$ dig +short math.stackexchange.com

151.101.129.69
151.101.65.69
151.101.193.69
151.101.1.69
```

上面命令只返回`math.stackexchange.com`对应的4个IP地址（即`A`记录）。

## syntax

`dig @server name type`

- server

  如果没有给server参数，Dig默认使用`/etc/resolve.conf`中的DNS服务器解析域名。

- name 

  想要被查询的域名

- type

  查询的类型，如果没有默认查询A记录。可以是ANY, A, NS , MX等。==如果手动添加type需要使用指定DNS server==

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_17-26-01.png" style="zoom:80%;" />

## 常用命令

- `dig -x host`

  反向解析

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_17-36-12.png" style="zoom:80%;" />

- `dig +short host`

  只显示最后的结果，不显示中间查询的内容

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_16-44-38.png" style="zoom:80%;" />

- `dig +trace host`

  显示所有路由的过程

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_16-52-04.png" style="zoom:80%;" />

- `dig -ttluints host`

  以可读的形式显示缓存时间， 未加该参数默认以sec显示

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_16-53-19.png" style="zoom:80%;" />

- `dig -yaml host`

  以yaml形式显示结果

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_16-57-06.png" style="zoom:67%;" />

- `whois baidu.com`

  查看域名注册信息
