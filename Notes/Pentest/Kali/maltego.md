# maltego

[TOC]

参考:

https://docs.maltego.com/support/solutions/articles/

首先进入界面调整字体大小

```mermaid
graph LR
a(application menu button)--> b(option)-->c(display)-->d(detail view) & e(machines log) & f(other componets)
```

## Terms

- Entity

  实体做为图中的结点，实体可以是DNS name，Person name, Phone number 等等

  Entity有三个属性

  Type指的是Entity是什么类型

  Value指的Entity最主要的信息

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-33-33.png" style="zoom:80%;" />

  Properties额外的信息

- Transform

  ==会自动通过DNS获取IP, 然后将IP反向解析成域名==

  递推出结点信息的抽象名词

  <img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-10-14.png" style="zoom:80%;" />

- Machine

  可以将多种Transform结合, 查询结果

- Hub Item

  可以从三方服务器上安装一些其他的Transform

## 应用界面



<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-19-49.png" style="zoom:80%;" />

通过`ctrl + T`新建一个tab, 要打开Entity details界面, Entity图标, 除Value以外

<img src="..\..\..\imgs\_Kali\Snipaste_2020-09-01_19-37-14.png"/>

添加了attchments的Entity会有一个回形针样式的图标, 添加了notes会有一个便签的标志

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-40-01.png" style="zoom:50%;" />

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-41-47.png" style="zoom:80%;" />



## 菜单

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-50-12.png" style="zoom:80%;" />

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_19-51-28.png" style="zoom:80%;" />

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_20-00-05.png" style="zoom:80%;" />

## 视图栏

可以改变拓扑图的类型, 双机Link可以添加label

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_20-02-38.png"/>

