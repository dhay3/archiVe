# IPtables-extensions limit Match

ref

https://www.linuxtopia.org/Linux_Firewall_iptables/x2682.html#LIMITMATCH

https://juejin.cn/post/7127958301685317669

https://www.cnblogs.com/crazymakercircle/p/15187184.html

https://en.wikipedia.org/wiki/Token_bucket

https://segmentfault.com/a/1190000015967922

## Digest

使用 token bucket algo (令牌桶算法) 来按照速率匹配报文，可以结合 LOG target 来记录报文

因为可以匹配速率，也就可以实现削峰填谷流量控制，以及防护简单的 Dos 攻击

## bucket algo

在谈令牌桶算法之前需要了解一下漏桶算法

漏桶算法的原理是：

假设现在有一个固定容量的漏桶，能进水也能出水（出水速率恒定）。

![2022-11-18_17-23](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221117/2022-11-18_17-23.btdpon7pwlk.webp)

那么会有以下几种情况( 均以恒定速度考虑 )

1. 如果进水流速大于出水流速，漏桶就会被装满，多的水就会撒出去
2. 如果进水流速等于或者小于出水流速，漏桶就不会被装满，所有的水都会从出水口流出

我们可以抽象成网络用语

进水流速就是进到网卡的报文速度，出水流速就是允许从网卡转发到系统协议栈的报文速度

## Token bucket algo

在解释令牌桶算法前，需要明白令牌桶算法为什么比漏桶算法好

有一个大学食堂，在同一时间提供 3 个窗口服务，假设现在进来很多干饭人，食堂还以 3 个窗口服务的速度。这样就有很多人得不到服务而排队到了食堂外。假设现在这些都同学非常紧急需要立马干饭，如果还以这样的速度服务，显然是不合理的。

如果我们按照干饭人需要干饭的速度增加窗口，这样就能很好的解决这个问题。但是呢由于食堂窗口是固定的不能增加也不能减少。这时需要怎么处理呢？那就是提高食堂大妈打饭的速度

在上述这儿个例子中，我们可以抽象成网络用语

食堂的窗口就是令牌桶，需要干饭人的速率就是请求的报文，食堂大妈打饭的速度就是令牌桶产生的速度。这样我就可以通过大妈打饭的速度来控制流量(但是显示中大妈并不是代码)。

![2022-11-18_17-48](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221117/2022-11-18_17-48.30kqpn19wcn4.webp)

![2022-11-18_17-48_1](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221117/2022-11-18_17-48_1.28owgblow6io.webp)

通常令牌桶算法分为几部分

1. 以一定的速率往令牌桶里放令牌，如果令牌桶满了令牌就会被丢弃
2. 进来的请求需要获取一张令牌才能被处理，如果令牌桶空了就丢弃请求

### Token bucket in limit

> 需要注意的是如果添加了规则，会理解按照指定的速度往令牌桶里放令牌

To further explain the limit match, it is basically a token bucket filter. Consider having a leaky bucket where the bucket leaks X packets per time-unit. X is defined depending on how many matching packets we get, so if we get 3 packets, the bucket leaks 3 packets per that time-unit. The **--limit** option tells us how many packets to refill the bucket with per time-unit（==往令牌桶放入令牌的速率==）, while the **--limit-burst** option tells us how big the bucket is in the first place(==令牌桶大小==). So, setting **--limit 3/minute --limit-burst 5**, and then receiving 5 matches will empty the bucket. After 20 seconds, the bucket is refilled with another token, and so on until the **--limit-burst** is reached again or until they get used.

例如

```
-m limit --limit 5/second --limit-burst 10
```

- `--limit-burst 10`

  令牌桶最大能存储 10 个令牌

- `--limit 5/second`

  每秒往令牌桶放入 5 个令牌，即每 0.2 秒放入 1 个令牌

## Optinal args

- `--limit rate[/second|/minute|/hour|/day] `

  Maximum  average  matching  rate:  specified  as a number, with an optional /second, /minute, /hour, or /day suffix; the default is 3/hour. 

  放入令牌的速率，默认 3 times per hour。==这里还需要注意的一点是，只支持 integer==

- `--limit-burst number `

  Maximum initial number of packets to match: this number gets recharged by one every time the limit specified above is not  reached,  up  to this number; the default is 5.

  令牌桶的初始大小，默认 5 times

## Exmaples

> 下面的例子均在令牌桶满的情况下实验

例如 192.168.1.1 访问 192.168.3.1 

192.168.1.1 以 1 秒间隔发包，192.168.3.1 在收到的报文数小于 3 时正常接受，当超过这个数值时只允许接受每分钟 30 个报文，即每 2 秒接收一个报文(产生一个令牌)

![2022-11-15_15-14](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221115/2022-11-15_15-14.56e538jrk8w0.webp)

192.168.3.1 设置规则

```
#每2秒
[root@labos-2 /]# iptables -t filter -A INPUT -s 192.168.1.1 -p icmp -m limit --limit 30/m --limit-burst 3 -j ACCEPT
[root@labos-2 /]# iptables -t filter -A INPUT -j DROP
```

192.168.1.1 ping 192.168.3.1

这里可以看到前 5 个报文正常

```
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
64 bytes from 192.168.3.1: icmp_seq=1 ttl=62 time=35.7 ms
64 bytes from 192.168.3.1: icmp_seq=2 ttl=62 time=36.5 ms
64 bytes from 192.168.3.1: icmp_seq=3 ttl=62 time=37.1 ms
64 bytes from 192.168.3.1: icmp_seq=4 ttl=62 time=37.7 ms
64 bytes from 192.168.3.1: icmp_seq=5 ttl=62 time=28.4 ms
64 bytes from 192.168.3.1: icmp_seq=7 ttl=62 time=32.9 ms
64 bytes from 192.168.3.1: icmp_seq=9 ttl=62 time=42.8 ms
64 bytes from 192.168.3.1: icmp_seq=11 ttl=62 time=24.4 ms
```

实际应该是前 6 个报文正常，之后 78 以此递归

```
1:
(3 + 0 - 1) % 3 = 2
2:
(2 + 1 - 1) % 3 = 2
3:
(2 + 0 - 1) % 3 = 1
4：
(1 + 1 - 1) % 3 = 1
5:
(1 + 0 - 1) % 3 = 0
6:
(0 + 1 - 1) % 3 = 0
7:
(0 + 0 - 1) % 3 = -1 drop
8:
(0 + 1 - 1) % 3 = 0
```

