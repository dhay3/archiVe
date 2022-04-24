# DNS记录类型

[TOC]

参考：

http://www.ruanyifeng.com/blog/2016/06/dns.html

https://itbilu.com/other/relate/EyxzdVl3.html

## 概述

> A记录优先于CNAME记录，一个主机地址同时拥有A和CNAME，CNAME不会生效
>
> @直接解析主域名

==正向解析通过域名解析IP==

- A（Address）地址记录

  指定域名指向的IP地址，也可以设置子域指向同一个IP

  > 如果使用CDN，然后使用dig命令获取到的是CDN的IP，而不是真实机的IP
  
  例如：

<img src="..\..\..\..\Images\_Net\计算机网络\Snipaste_2020-08-30_16-58-17.png"/>

第一条访问`chz.cyberpelican.space`跳转到102.168.80.201服务器

第二条使用@解析访问`cyberpelican.space`跳转到36.152.44.95

第三条访问`www.cyberpelica.space`跳转到36.152.44.95

访问相同主机名，但是IP地址不同，用作负载均衡

其中的TTL指的是time to live，即在DNS服务器中缓存的时间

> A与AAAA的区别
>
> A指向的IP地址为IPv4
>
> AAAA指向的地址为IPv6

- NS（Name Server）域名服务器记录

  返回保存下一级域名信息的服务器地址。该记录只能设置为域名，不能设置为IP地址。用来指定该域名由哪个DNS服务器来进行解析。

- MX（Mail Exchange）邮件记录

  指向电子邮件的服务地址

- CNAME（Canonical Name）规范记录名

  > 一般设置CNAME就无需设置A记录

  指向一个域名，即当前查询的域名是一个域名的跳转

  例如：

  `facebook.github.io` CNAME `github.map.fastly.net`
  
  当用户查询`facebook.github.io`时，实际返回的是`github.map.fastly.net`的IP地址。这样的好处是，变更服务器IP地址时，只要修改`github.map.fastly.net`这个域名就可以了，用户的`facebook.github.io`域名不用修改。

==逆向解析通过IP解析域名==

- PTR（Pointer Record）逆向查询记录

  IP地址对应域名







