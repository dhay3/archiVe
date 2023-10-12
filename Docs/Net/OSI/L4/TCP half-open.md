# TCP half-open

ref:

https://en.wikipedia.org/wiki/TCP_half-open

https://zhuanlan.zhihu.com/p/99152064

https://zhuanlan.zhihu.com/p/144785626

https://www.youtube.com/watch?v=SJq61Rhr6N4

## Digest

TCP 正常连接是，A 向 B 发 SYN，B 向 A 回 ACK-SYN，A 向 B 发 ACK。之后双端发ACK含数据包。但是针对 TCP half-open A不会给 B 回 ACK 包。

在 TCP 3 次握手时，linux 内核会维护两个队列

1. 半连接队列，也称为 syn 队列
2. 全连接队列，也称为 accept 队列

服务端收到客户端发起的 SYN 请求后，内核会把该连接储存到半连接队列，并向客户端响应 SYN + ACK，接着客户端返回 ACK，服务端收到第 3 次握手的 ACK 后，内核会把连接从半连接队列移除，染回创建新的完全的连接，并将其添加到 accept 队列，等待进程调用 accept 函数时把连接取出

![2022-04-18_22-05](https://github.com/dhay3/image-repo/raw/master/20220418/2022-04-18_22-05.tw4nn0kcjw0.webp)

不管半连接队列还是全连接队列，都有最大长度限制，超过限制时，内核会直接丢弃，或返回RST包

### **实战 - TCP 全连接队列溢出**

> 如何知道应用程序的 TCP 全连接队列大小？

在服务端可以使用 `ss` 命令，来查看 TCP 全连接队列的情况：

但需要注意的是 `ss` 命令获取的 `Recv-Q/Send-Q` 在「LISTEN 状态」和「非 LISTEN 状态」所表达的含义是不同的。从下面的内核代码可以看出区别：

![img](https://pic1.zhimg.com/80/v2-fec2cac566f6e99a6029981214e2b328_720w.jpg)

在「LISTEN 状态」时，`Recv-Q/Send-Q` 表示的含义如下：

![img](https://pic4.zhimg.com/80/v2-58616418e6079239df5c6ebf0125e3ab_720w.jpg)

- Recv-Q：当前全连接队列的大小，也就是当前已完成三次握手并等待服务端 `accept()` 的 TCP 连接；
- Send-Q：当前全连接最大队列长度，上面的输出结果说明监听 8088 端口的 TCP 服务，最大全连接长度为 128；

在「非 LISTEN 状态」时，`Recv-Q/Send-Q` 表示的含义如下：

![img](https://pic3.zhimg.com/80/v2-913628e7d6155a02492720b8ac42a6e6_720w.jpg)

- Recv-Q：已收到但未被应用进程读取的字节数；
- Send-Q：已发送但未收到确认的字节数；

> 如何模拟 TCP 全连接队列溢出的场景？

![img](https://pic1.zhimg.com/80/v2-f2a27d294bf2a7437c2672251478f9f0_720w.jpg)

实验环境：

- 客户端和服务端都是 CentOs 6.5 ，Linux 内核版本 2.6.32
- 服务端 IP 192.168.3.200，客户端 IP 192.168.3.100
- 服务端是 Nginx 服务，端口为 8088

这里先介绍下 `wrk` 工具，它是一款简单的 HTTP 压测工具，它能够在单机多核 CPU 的条件下，使用系统自带的高性能 I/O 机制，通过多线程和事件模式，对目标机器产生大量的负载。

本次模拟实验就使用 `wrk` 工具来压力测试服务端，发起大量的请求，一起看看服务端 TCP 全连接队列满了会发生什么？有什么观察指标？

客户端执行 `wrk` 命令对服务端发起压力测试，并发 3 万个连接：

![img](https://pic3.zhimg.com/80/v2-4b681c122f2c6f3ae699b35fb09f462e_720w.jpg)

在服务端可以使用 `ss` 命令，来查看当前 TCP 全连接队列的情况：

![img](https://pic2.zhimg.com/80/v2-8372d7ab038c68ad93bf4258c0cb76f9_720w.jpg)

其间共执行了两次 ss 命令，从上面的输出结果，可以发现当前 TCP 全连接队列上升到了 129 大小，超过了最大 TCP 全连接队列。

**当超过了 TCP 最大全连接队列，服务端则会丢掉后续进来的 TCP 连接**，丢掉的 TCP 连接的个数会被统计起来，我们可以使用 netstat -s 命令来查看：

![img](https://pic4.zhimg.com/80/v2-160844d5ad3587f07c645252ee95a30b_720w.jpg)

上面看到的 41150 times ，表示全连接队列溢出的次数，注意这个是累计值。可以隔几秒钟执行下，如果这个数字一直在增加的话肯定全连接队列偶尔满了。

从上面的模拟结果，可以得知，**当服务端并发处理大量请求时，如果 TCP 全连接队列过小，就容易溢出。发生 TCP 全连接队溢出的时候，后续的请求就会被丢弃，这样就会出现服务端请求数量上不去的现象。**

![img](https://pic2.zhimg.com/80/v2-59396b0f9eb18eca18fff60398558dc1_720w.jpg)

> Linux 有个参数可以指定当 TCP 全连接队列满了会使用什么策略来回应客户端。

实际上，丢弃连接只是 Linux 的默认行为，我们还可以选择向客户端发送 RST 复位报文，告诉客户端连接已经建立失败。

![img](https://pic2.zhimg.com/80/v2-93d281156a8965f549fe7c92d855d0a5_720w.jpg)

tcp_abort_on_overflow 共有两个值分别是 0 和 1，其分别表示：

- 0 ：如果全连接队列满了，那么 server 扔掉 client 发过来的 ack ；
- 1 ：如果全连接队列满了，server 发送一个 `reset` 包给 client，表示废掉这个握手过程和这个连接；

如果要想知道客户端连接不上服务端，是不是服务端 TCP 全连接队列满的原因，那么可以把 tcp_abort_on_overflow 设置为 1，这时如果在客户端异常中可以看到很多 `connection reset by peer` 的错误，那么就可以证明是由于服务端 TCP 全连接队列溢出的问题。

通常情况下，应当把 tcp_abort_on_overflow 设置为 0，因为这样更有利于应对突发流量。

举个例子，当 TCP 全连接队列满导致服务器丢掉了 ACK，与此同时，客户端的连接状态却是 ESTABLISHED，进程就在建立好的连接上发送请求。只要服务器没有为请求回复 ACK，请求就会被多次**重发**。如果服务器上的进程只是**短暂的繁忙造成 accept 队列满，那么当 TCP 全连接队列有空位时，再次接收到的请求报文由于含有 ACK，仍然会触发服务器端成功建立连接。**

所以，tcp_abort_on_overflow 设为 0 可以提高连接建立的成功率，只有你非常肯定 TCP 全连接队列会长期溢出时，才能设置为 1 以尽快通知客户端。

> 如何增大 TCP 全连接队列呢？

是的，当发现 TCP 全连接队列发生溢出的时候，我们就需要增大该队列的大小，以便可以应对客户端大量的请求。

**TCP 全连接队列足最大值取决于 somaxconn 和 backlog 之间的最小值，也就是 min(somaxconn, backlog)**。从下面的 Linux 内核代码可以得知：

![img](https://pic4.zhimg.com/80/v2-2a307110210956c8fed02297e735de0f_720w.jpg)

- `somaxconn` 是 Linux 内核的参数，默认值是 128，可以通过 `/proc/sys/net/core/somaxconn` 来设置其值；
- `backlog` 是 `listen(int sockfd, int backlog)` 函数中的 backlog 大小，Nginx 默认值是 511，可以通过修改配置文件设置其长度；

前面模拟测试中，我的测试环境：

- somaxconn 是默认值 128；
- Nginx 的 backlog 是默认值 511

所以测试环境的 TCP 全连接队列最大值为 min(128, 511)，也就是 `128`，可以执行 `ss` 命令查看：

![img](https://pic4.zhimg.com/80/v2-74af6eeedc37e20ba2db029feb9b3493_720w.jpg)

现在我们重新压测，把 TCP 全连接队列**搞大**，把 `somaxconn` 设置成 5000：

![img](https://pic3.zhimg.com/80/v2-888624bdb5b1c721372b51cc8b6730fa_720w.jpg)

接着把 Nginx 的 backlog 也同样设置成 5000：

![img](https://pic4.zhimg.com/80/v2-d760548fdf95b4e58f1bdeb22dadd077_720w.jpg)

最后要重启 Nginx 服务，因为只有重新调用 `listen()` 函数 TCP 全连接队列才会重新初始化。

重启完后 Nginx 服务后，服务端执行 ss 命令，查看 TCP 全连接队列大小：

![img](https://pic1.zhimg.com/80/v2-cd2799bbe267253c8fc3f02cedb11018_720w.jpg)

从执行结果，可以发现 TCP 全连接最大值为 5000。

> 增大 TCP 全连接队列后，继续压测

客户端同样以 3 万个连接并发发送请求给服务端：

![img](https://pic3.zhimg.com/80/v2-4b681c122f2c6f3ae699b35fb09f462e_720w.jpg)

服务端执行 `ss` 命令，查看 TCP 全连接队列使用情况：

![img](https://pic2.zhimg.com/80/v2-b32974e7b266aea5fe67b8621c4ff5a9_720w.jpg)

从上面的执行结果，可以发现全连接队列使用增长的很快，但是一直都没有超过最大值，所以就不会溢出，那么 `netstat -s` 就不会有 TCP 全连接队列溢出个数的显示：

![img](https://pic2.zhimg.com/80/v2-65af6fb03e6614cadbfd1580f97bcf8d_720w.jpg)

说明 TCP 全连接队列最大值从 128 增大到 5000 后，服务端抗住了 3 万连接并发请求，也没有发生全连接队列溢出的现象了。

**如果持续不断地有连接因为 TCP 全连接队列溢出被丢弃，就应该调大 backlog 以及 somaxconn 参数。**

------

### **实战 - TCP 半连接队列溢出**

> 如何查看 TCP 半连接队列长度？

很遗憾，TCP 半连接队列长度的长度，没有像全连接队列那样可以用 ss 命令查看。

但是我们可以抓住 TCP 半连接的特点，就是服务端处于 `SYN_RECV` 状态的 TCP 连接，就是在 TCP 半连接队列。

于是，我们可以使用如下命令计算当前 TCP 半连接队列长度：

![img](https://pic3.zhimg.com/80/v2-c29ff3fc73841971b75111e15e83245a_720w.jpg)

> 如何模拟 TCP 半连接队列溢出场景？

模拟 TCP 半连接溢出场景不难，实际上就是对服务端一直发送 TCP SYN 包，但是不回第三次握手 ACK，这样就会使得服务端有大量的处于 `SYN_RECV` 状态的 TCP 连接。

这其实也就是所谓的 SYN 洪泛、SYN 攻击、DDos 攻击。

![img](https://pic3.zhimg.com/80/v2-985def635cc2fe122badc6e1b7eb28d6_720w.jpg)

实验环境：

- 客户端和服务端都是 CentOs 6.5 ，Linux 内核版本 2.6.32
- 服务端 IP 192.168.3.200，客户端 IP 192.168.3.100
- 服务端是 Nginx 服务，端口为 8088

注意：本次模拟实验是没有开启 tcp_syncookies，关于 tcp_syncookies 的作用，后续会说明。

本次实验使用 `hping3` 工具模拟 SYN 攻击：

![img](https://pic2.zhimg.com/80/v2-f00a472a51b8b82c9ad77671b3821699_720w.jpg)

当服务端受到 SYN 攻击后，连接服务端 ssh 就会断开了，无法再连上。只能在服务端主机上执行查看当前 TCP 半连接队列大小：

![img](https://pic2.zhimg.com/80/v2-f3e5cf7e9a79f3fe3d7b1c26c3874c19_720w.jpg)

同时，还可以通过 netstat -s 观察半连接队列溢出的情况：

![img](https://pic3.zhimg.com/80/v2-5e7ad06d807dc5a24dad5ce63cc80e6e_720w.jpg)

上面输出的数值是**累计值**，表示共有多少个 TCP 连接因为半连接队列溢出而被丢弃。**隔几秒执行几次，如果有上升的趋势，说明当前存在半连接队列溢出的现象**。

> 大部分人都说 tcp_max_syn_backlog 是指定半连接队列的大小，是真的吗？

很遗憾，半连接队列的大小并不单单只跟 `tcp_max_syn_backlog` 有关系。

上面模拟 SYN 攻击场景时，服务端的 tcp_max_syn_backlog 的默认值如下：

![img](https://pic3.zhimg.com/80/v2-b89f6fe7fe90ef4b80199893110a01f2_720w.jpg)

但是在测试的时候发现，服务端最多只有 256 个半连接队列，而不是 512，所以**半连接队列的最大长度不一定由 tcp_max_syn_backlog 值决定的**。

> 接下来，走进 Linux 内核的源码，来分析 TCP 半连接队列的最大值是如何决定的。

TCP 第一次握手（收到 SYN 包）的 Linux 内核代码如下，其中缩减了大量的代码，只需要重点关注 TCP 半连接队列溢出的处理逻辑：

![img](https://pic2.zhimg.com/80/v2-19f054b6abe898450ce9f12ef4286055_720w.jpg)

从源码中，我可以得出共有三个条件因队列长度的关系而被丢弃的：

![img](https://pic1.zhimg.com/80/v2-2326c6df59e30888e668cd3727b6ec60_720w.jpg)

1. **如果半连接队列满了，并且没有开启 tcp_syncookies，则会丢弃；**
2. **若全连接队列满了，且没有重传 SYN+ACK 包的连接请求多于 1 个，则会丢弃；**
3. **如果没有开启 tcp_syncookies，并且 max_syn_backlog 减去 当前半连接队列长度小于 (max_syn_backlog >> 2)，则会丢弃；**

关于 tcp_syncookies 的设置，后面在详细说明，可以先给大家说一下，开启 tcp_syncookies 是缓解 SYN 攻击其中一个手段。

接下来，我们继续跟一下检测半连接队列是否满的函数 inet_csk_reqsk_queue_is_full 和 检测全连接队列是否满的函数 sk_acceptq_is_full ：

![img](https://pic2.zhimg.com/80/v2-e50219dd1053121c4abf625467f8ae3d_720w.jpg)

从上面源码，可以得知：

- **全**连接队列的最大值是 `sk_max_ack_backlog` 变量，sk_max_ack_backlog 实际上是在 listen() 源码里指定的，也就是 **min(somaxconn, backlog)**；
- **半**连接队列的最大值是 `max_qlen_log` 变量，max_qlen_log 是在哪指定的呢？现在暂时还不知道，我们继续跟进；

我们继续跟进代码，看一下是哪里初始化了半连接队列的最大值 max_qlen_log：

![img](https://pic4.zhimg.com/80/v2-bf74e43ac16537120e1185cfb3335eeb_720w.jpg)

从上面的代码中，我们可以算出 max_qlen_log 是 8，于是代入到 检测半连接队列是否满的函数 reqsk_queue_is_full ：

![img](https://pic3.zhimg.com/80/v2-b3938f1c160607cbc7bddfee87a4170e_720w.jpg)

也就是 `qlen >> 8` 什么时候为 1 就代表半连接队列满了。这计算这不难，很明显是当 qlen 为 256 时，`256 >> 8 = 1`。

至此，总算知道为什么上面模拟测试 SYN 攻击的时候，服务端处于 `SYN_RECV` 连接最大只有 256 个。

可见，**半连接队列最大值不是单单由 max_syn_backlog 决定，还跟 somaxconn 和 backlog 有关系。**

在 Linux 2.6.32 内核版本，它们之间的关系，总体可以概况为：

![img](https://pic1.zhimg.com/80/v2-00246c3957e12fd149b82d40745bcb04_720w.jpg)

- 当 max_syn_backlog > min(somaxconn, backlog) 时， 半连接队列最大值 max_qlen_log = min(somaxconn, backlog) * 2;
- 当 max_syn_backlog < min(somaxconn, backlog) 时， 半连接队列最大值 max_qlen_log = max_syn_backlog * 2;

> 半连接队列最大值 max_qlen_log 就表示服务端处于 SYN_REVC 状态的最大个数吗？

依然很遗憾，并不是。

max_qlen_log 是**理论**半连接队列最大值，并不一定代表服务端处于 SYN_REVC 状态的最大个数。

在前面我们在分析 TCP 第一次握手（收到 SYN 包）时会被丢弃的三种条件：

1. 如果半连接队列满了，并且没有开启 tcp_syncookies，则会丢弃；
2. 若全连接队列满了，且没有重传 SYN+ACK 包的连接请求多于 1 个，则会丢弃；
3. **如果没有开启 tcp_syncookies，并且 max_syn_backlog 减去 当前半连接队列长度小于 (max_syn_backlog >> 2)，则会丢弃；**

假设条件 1 当前半连接队列的长度 「没有超过」理论的半连接队列最大值 max_qlen_log，那么如果条件 3 成立，则依然会丢弃 SYN 包，也就会使得服务端处于 SYN_REVC 状态的最大个数不会是理论值 max_qlen_log。

似乎很难理解，我们继续接着做实验，实验见真知。

服务端环境如下：

![img](https://pic4.zhimg.com/80/v2-30417b0ee52430c91f14b575909c6017_720w.jpg)

配置完后，服务端要重启 Nginx，因为全连接队列最大和半连接队列最大值是在 listen() 函数初始化。

根据前面的源码分析，我们可以计算出半连接队列 max_qlen_log 的最大值为 256：

![img](https://pic2.zhimg.com/80/v2-071fd4dd0a79ee492f739d7e31c99849_720w.jpg)

客户端执行 hping3 发起 SYN 攻击：

![img](https://pic3.zhimg.com/80/v2-a4c65480ae2c2e687b2e4096187c4dda_720w.jpg)

服务端执行如下命令，查看处于 SYN_RECV 状态的最大个数：

![img](https://pic2.zhimg.com/80/v2-e66779afc68b9c1f0839c62e20317429_720w.jpg)

可以发现，服务端处于 SYN_RECV 状态的最大个数并不是 max_qlen_log 变量的值。

这就是前面所说的原因：**如果当前半连接队列的长度 「没有超过」理论半连接队列最大值 max_qlen_log，那么如果条件 3 成立，则依然会丢弃 SYN 包，也就会使得服务端处于 SYN_REVC 状态的最大个数不会是理论值 max_qlen_log。**

我们来分析一波条件 3 :

![img](https://pic1.zhimg.com/80/v2-dbfbf46876b8421912e2a1decc433910_720w.jpg)

从上面的分析，可以得知如果触发「当前半连接队列长度 > 192」条件，TCP 第一次握手的 SYN 包是会被丢弃的。

在前面我们测试的结果，服务端处于 SYN_RECV 状态的最大个数是 193，正好是触发了条件 3，所以处于 SYN_RECV 状态的个数还没到「理论半连接队列最大值 256」，就已经把 SYN 包丢弃了。

所以，服务端处于 SYN_RECV 状态的最大个数分为如下两种情况：

- 如果「当前半连接队列」**没超过**「理论半连接队列最大值」，但是**超过** max_syn_backlog - (max_syn_backlog >> 2)，那么处于 SYN_RECV 状态的最大个数就是 max_syn_backlog - (max_syn_backlog >> 2)；
- 如果「当前半连接队列」**超过**「理论半连接队列最大值」，那么处于 SYN_RECV 状态的最大个数就是「理论半连接队列最大值」；

> 每个 Linux 内核版本「理论」半连接最大值计算方式会不同。

在上面我们是针对 Linux 2.6.32 版本分析的「理论」半连接最大值的算法，可能每个版本有些不同。

比如在 Linux 5.0.0 的时候，「理论」半连接最大值就是全连接队列最大值，但依然还是有队列溢出的三个条件：

![img](https://pic3.zhimg.com/80/v2-48b219c4ae4366c31191e1c0ab10b99e_720w.jpg)

> 如果 SYN 半连接队列已满，只能丢弃连接吗？

并不是这样，**开启 syncookies 功能就可以在不使用 SYN 半连接队列的情况下成功建立连接**，在前面我们源码分析也可以看到这点，当开启了 syncookies 功能就不会丢弃连接。

syncookies 是这么做的：服务器根据当前状态计算出一个值，放在己方发出的 SYN+ACK 报文中发出，当客户端返回 ACK 报文时，取出该值验证，如果合法，就认为连接建立成功，如下图所示。

![img](https://pic4.zhimg.com/80/v2-dbad175c877ba1e4e05a3f808c70d9ef_720w.jpg)

syncookies 参数主要有以下三个值：

- 0 值，表示关闭该功能；
- 1 值，表示仅当 SYN 半连接队列放不下时，再启用它；
- 2 值，表示无条件开启功能；

那么在应对 SYN 攻击时，只需要设置为 1 即可：

![img](https://pic4.zhimg.com/80/v2-28441ed209099db6190708135b886213_720w.jpg)

> 如何防御 SYN 攻击？

这里给出几种防御 SYN 攻击的方法：

- 增大半连接队列；
- 开启 tcp_syncookies 功能
- 减少 SYN+ACK 重传次数

*方式一：增大半连接队列*

在前面源码和实验中，得知**要想增大半连接队列，我们得知不能只单纯增大 tcp_max_syn_backlog 的值，还需一同增大 somaxconn 和 backlog，也就是增大全连接队列**。否则，只单纯增大 tcp_max_syn_backlog 是无效的。

增大 tcp_max_syn_backlog 和 somaxconn 的方法是修改 Linux 内核参数：

![img](https://pic4.zhimg.com/80/v2-9057529e4cbcb9c0cd1ab4c36d7a55c7_720w.jpg)

增大 backlog 的方式，每个 Web 服务都不同，比如 Nginx 增大 backlog 的方法如下：

![img](https://pic1.zhimg.com/80/v2-923937d30a59d4ad7fe723ae985ba8cc_720w.jpg)

最后，改变了如上这些参数后，要重启 Nginx 服务，因为半连接队列和全连接队列都是在 listen() 初始化的。

*方式二：开启 tcp_syncookies 功能*

开启 tcp_syncookies 功能的方式也很简单，修改 Linux 内核参数：

![img](https://pic4.zhimg.com/80/v2-28441ed209099db6190708135b886213_720w.jpg)

*方式三：减少 SYN+ACK 重传次数*

当服务端受到 SYN 攻击时，就会有大量处于 SYN_REVC 状态的 TCP 连接，处于这个状态的 TCP 会重传 SYN+ACK ，当重传超过次数达到上限后，就会断开连接。

那么针对 SYN 攻击的场景，我们可以减少 SYN+ACK 的重传次数，以加快处于 SYN_REVC 状态的 TCP 连接断开。

![img](https://pic2.zhimg.com/80/v2-f293fe9e265893fbf50b085942cdc65d_720w.jpg)

------

### **巨人的肩膀**

[1] 系统性能调优必知必会.陶辉.极客时间.

[2] [https://www.cnblogs.com/zengkefu/p/5606696.html](https://link.zhihu.com/?target=https%3A//www.cnblogs.com/zengkefu/p/5606696.html)

[3] [https://blog.cloudflare.com/syn-packet-handling-in-the-wild/](https://link.zhihu.com/?target=https%3A//blog.cloudflare.com/syn-packet-handling-in-the-wild/)

------

**小林是专为大家图解的工具人，Goodbye，我们下次见！**

------

「图解网络」文章受到了很多读者的喜爱与支持，为了方便大家阅读，**小林把自己原创的图解网络系列文章整理成了pdf，内容涵盖计算机网络的重点知识，比如 HTTP、TCP、UDP、IP等等，pdf 共「300 页 + 9W字 + 30 张图」。**

并且根据不同读者的阅读习惯，我整了两种风格的图解网络pdf，分别是「亮白版本」和「暗黑版本」。

![img](https://pic3.zhimg.com/80/v2-49f2da44e76ed34c648e77b22f80712e_720w.jpg)亮白版本

![img](https://pic1.zhimg.com/80/v2-60a8748d907d4ebbbf763283bfdb1984_720w.jpg)暗黑版本

百度网盘下载链接如下：

链接：[https://pan.baidu.com/s/1dRNDPW_WjV6vb3liiLPqHA](https://link.zhihu.com/?target=https%3A//pan.baidu.com/s/1dRNDPW_WjV6vb3liiLPqHA)

提取码：t95u

编辑于 2020-10-11 09:55