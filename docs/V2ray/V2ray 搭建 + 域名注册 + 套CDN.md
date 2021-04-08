# V2ray 搭建 + 域名注册 + 套CDN

参考链接:

https://www.onlyling.com/archives/464

https://blog.sprov.xyz/2019/03/11/cdn-v2ray-safe-proxy/

[https://github.com/Alvin9999/new-pac/wiki/%E8%87%AA%E5%BB%BAv2ray%E6%9C%8D%E5%8A%A1%E5%99%A8%E6%95%99%E7%A8%8B](https://github.com/Alvin9999/new-pac/wiki/自建v2ray服务器教程)

[TOC]

## #前置知识

- DNS

  Domain Name Service 

  访问` google.com`通过DNS将输入的文字解析成真是的服务器`ip`

- GFW

  Greate Fire Wall == 防火墙

  发送数据包会经过防火墙

  GFW的封锁手段

  - DNS污染: 访问`google.com`防火墙会将域名解析成错误的`ip`

  - 过滤关键字
  - 端口阻断
  - `IP`地址批量封锁

- VPN

  Virtual Private Network

  在公用网络上建立专用网络, 进行加密通讯 (直接对流量加密, 不安全)

- 软件翻墙 

  并不是实际意义上的全局代理, 只会接管经过经过如软件代理出去的流量

   Shawdowsocks

  通过本地软件加密流量, 然后发送至远程服务器再进行转发

  - SSR
  - V2Ray
  - Trojan

- 硬件翻墙

  强制代理所有流量, 达到正真的全局代理

  路由器

  

- VPS

  A **virtual private server** (**VPS**) is a virtual machine sold as a service by an Internet hosting service

提供虚拟网络服务

## #购买服务器

https://my.vultr.com/

**在购买之前测试一下速度**

选择`support`, 按照如下测试, 不同的服务商速度不同

<img src="../../imgs/_V2ray/2.png"/>

购买服务器

<img src="../../imgs/_V2ray/1.png" style="zoom:80%;" />

选第一个, 然后根据测速出来的实际情况, 选择服务器, 和服务

<img src="../../imgs/_V2ray/3.png" style="zoom:80%;" />

**注意：2.5美元套餐只提供ipv6 ip，一般的电脑用不了，所以建议选择3.5美元及以上的套餐。**

选好之后其他的不用管, 直接`Deploye Now`, 部署需要花费大概2-3分钟, 然后点击

<img src="../../imgs/_V2ray/4.png"/>

复制`ip`  , 用本机`ping`一下服务器的`ip`

```c
ping -t 复制的ip地址
```

如果不能`ping`通说明开到一个被墙的`ip`, 需要重新部署一个, 这里先不要把被墙的服务器删除, 先新开一个服务器, 成功`ping`通后再删除, 因为有可能开到的还是自己这个被墙的`ip`

<img src="../../imgs/_V2ray/5.png" style="zoom:80%;" />

<img src="../../imgs/_V2ray/6.png"/>

## #搭建服务器

然后使用`XShell`远程登入到服务器, 输入用户名`root`, 密码

安装`v2ray`脚本, 该脚本自动安装`BBR`加速

```
bash <(curl -s -L https://233v2.com/v2ray.sh)
```

选1 , 一路默认即可

<img src="../../imgs/_V2ray\35.png" style="zoom:80%;" />



一切配置完后输入 `v2ray url` 获取`vmess url`

通过客户端`v2rayN`导入链接

<img src="../../imgs/_V2ray\34.png" style="zoom:80%;" />

再下载插件`switchyomega`, 进入控制面板

<img src="../../imgs/_V2ray/13.png" style="zoom:80%;" />

使用`proxy`模式即可网上冲浪

或是使用`v2rayN` 如下配置,  然后将`switchyomega`置为系统代理

<img src="../../imgs/_V2ray\36.png" style="zoom:80%;" />

- PAC (proxy auto-config) 代理

  一个自动代理配置脚本, 它能决定网络流量走默认通道还是代理通道, 控制的流量类包括:  HTTP, HTTPS和

  FTP

  也就是说如果网站不需要翻墙就走默认的代理, 如果网站需要翻墙才会代理出去

  一般使用这种

- Http代理

  默认所有流量都通过代理

**还需要说的一点是**

如果开启了`v2ray`的全局代理或是PAC无需使用`proxyswitcher`

可以输入`v2ray`来修改配置

<img src="../../imgs/_V2ray\37.png" style="zoom:80%;" />

## #域名申请

http://www.freenom.com/en/index.html

免费域名, 可以免费使用1年

<img src="../../imgs/_V2ray\18.png" style="zoom:80%;" />

这里直接使用谷歌登入,  进入`Edit Account Detail`

==ip地址一定要一致==, 可以通过该网址来查询: https://whatismyipaddress.com/

<img src="../../imgs/_V2ray\19.png" style="zoom:80%;" />

然后选择`services` ----> `Register a New Domain`

填写你想要的二级域名, 然后`check avaliability`

<img src="../../imgs/_V2ray\20.png" style="zoom:80%;" />

进入购物车后, 可以选中使用域名的时间

<img src="../../imgs/_V2ray\21.png" style="zoom:80%;" />

==再需要声明的一点是, ip地址一定要一致==

<img src="../../imgs/_V2ray\23.png" style="zoom:80%;" />

成功界面如下

<img src="../../imgs/_V2ray\24.png" style="zoom:80%;" />

## #CDN

https://dash.cloudflare.com/

`CDN` (content delivery network), 内容分发网络

因为`vps`和你之间是直接连接的, 中间必定经过防火墙, 只要你的特征越来越明显, 防火墙觉得合适了就会阻断你的`IP`, 从而导致无法连接`vps`

使用`CDN`可以做反向代理, 防火墙并不知道你的`vps`实际`IP`, 这样就可以有效的防止你的`IP`被墙, 甚至能让被墙的`IP`复活. 但是有可能会增加延迟

这里使用`cloudflare`提供的免费CDN服务

输入你的域名

<img src="../../imgs/_V2ray\25.png" style="zoom:80%;" />

选择0美元/月计划, 然后`cloudflare`会让你绑定DNS

<img src="../../imgs/_V2ray\26.png" style="zoom:80%;" />

我们重新会到`freenom`, 选择`service`--->`My Domain` , 点击`Manage Domain`

<img src="../../imgs/_V2ray\27.png" style="zoom:80%;" />

然后选择`Management Tools`--->`Nameservers`--->`Use custom nameservers (enter below)`

填入`cloudflare`提供的服务器地址, 保存修改

<img src="../../imgs/_V2ray\28.png" style="zoom:80%;" />

重回到`cloudflare`, 点击 检查名称服务器, 然后按照如下内容填写自己对应的信息

要先关闭代理, 

<img src="../../imgs/_V2ray\29.png" style="zoom:80%;" />

SSL/TLS 选择完全

<img src="../../imgs/_V2ray\30.png" style="zoom:80%;" />

## #服务器套CDN

安装, 卸载脚本

```
bash <(curl -s -L https://233v2.com/v2ray.sh)
```

<img src="../../imgs/_V2ray\31.png" style="zoom:80%;" />

选4 `webSocket + TLS`

<img src="../../imgs/_V2ray\32.png" style="zoom:80%;" />

填入域名

<img src="../../imgs/_V2ray\33.png" style="zoom:80%;" />

之后与使用`TPC`伪装的一样导入`vmess url`到`v2rayN`中即可

我们可以通过如下方法来切换服务器

<img src="../../imgs/_V2ray\38.png" style="zoom:80%;" />

最后别忘了 重新开启`cloudflare`代理

<img src="../../imgs/_V2ray\39.png" style="zoom:80%;" />

