# Day37 - NTP



## The importance of time

在思科的设备上可以使用 `show clock` 来查看设备的时间，默认使用 UTC 格式

```
R1#show clock
*23:7:55.803 UTC Sun Feb 28 1993
```

如果使用了 `show clock detail` 还可以查看设备时间参考的源信息(依据)

software clock 基于 hardware clock

```
R1#show clock detail
*23:8:27.450 UTC Sun Feb 28 1993
Time source is hardware calendar
```

*The hardware calendar is the default time source*

其中 \* = time is not consider authoritative 

即设备并不确认时间是否准确

*The internal hardware clock of a device will drift over time, so it is not the ideal time source*

**确保时间的正确性最主要的原因是需要保证设备上的日志是准确的**

可以使用 `show logging` 来查看设备上的日志

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-26.2c7bjene8av.webp)

## Manual Time Configuration

可以通过 `R1#clock set` 来手动配置时间

> 注意这里并不需要进入 global config mode
>
> 因为时间是动态的，通过 `write` 命令来写配置不合乎逻辑

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-32.1qopo83vn35s.webp)

*Although the hardware calendar(built-in clock) is the default time-source, the hardware clock and software clock are separate and be configured separately*

这里配置的是 software clock

可以通过 `R1#calendar set` 来配置 hardware clock 时间

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-37.s5ls6sa7awg.webp)

配置完之后还需要使用 `clock update-calendar` 或者 `clock read-calendar` 来让时间生效，否则会在重启后还原

### clok update-calendar

让 hardware clock 和 software clock 同步，hardware clock 被更新

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-45.6xb3xctzs29s.webp)

### clock read-calendar

让 software clock 和 hardware clock 同步, software clock 被更新

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-48.6mnse5q6ocjk.webp)

### Configuring the Time Zone

可以使用 `R1(config)#clock timezone <zone name> <offset...>` 来配置 timezone，需要在 global config mode 中配置

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_18-53.3vhvll30niv4.webp)

> name of time zone 可以是任意的字符，无需和实际存在的 TimeZone 关联

如果在配置 time zone 后，修改 software 或者是 hardware clock，时间不会加上 time zone 设置的 offset

> 因为 NTP 使用 UTC 格式，所以 time zone 必须要准确否则 NTP 同步的时间会有问题

## NTP

Network Time Protocol(NTP)

如果网络拓扑中设备比较少，使用手动配置时间的方式可能可以接受，但是设备非常多，通过手动配置时间的方式会消耗大量的时间，显然是非常不合理的，所以需要使用 NTP 来自动同步

- NTP allows automatic syncing of time over a network

- NET clients request the time from NTP servers

- A device can be an NTP server and an NTP client at the same time

- NTP allows accuracy of time within ~1 millisecond if the NTP server is the same LAN, or within ~50 milliseconds if connecting to the NTP server over a WAN/the Internet

- Some NTP servers are ‘better’ than others. The ‘distance’ of an NTP server from the original reference clock is called stratum(阶层)

  和 metric 逻辑类似，但是距离越远，stratum 值越高

- NTP uses UDP 123 port to communicate 

### References Clockcs

*A reference clock is usually a very accurate time device like an atomic clock or a GPS clock*

*Reference clocks are stratum 0 within the NTP  hirerchy*

*NTP servers directly connected to reference clocks are stratum 1*

> reference clock 是最精确的设备，stratum 值为 0，和 reference clock 直联的设备 stratum 值为 1

### NTP Hierarchy

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-13.55zde3747uo0.webp)

- stratum 1 的 NTP 服务器和 stratum 0 的 reference clock 做时间同步

- stratum 2 的 NTP 服务器和 stratum 1 的服务器做时间同步

- stratum 3 的 NTP 服务器和 stratum 2 的服务器做时间同步

- stratum 15 是最大的值，如果 NTP 服务器 stratum 超过该值，会被认为不可靠

- 相同 stratum 的服务器(peering)之间也可以做时间同步，这种也被称为 **symmetric active mode**

  思科的设备同样支持 symmetric active mode

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-23.4v8pc68xxeo0.webp)

stratum 1 的设备也被称为 primary servers，stratum 值大于 1 的被称为 secondary servers

## NTP Configuration

例如如下拓扑，需要让 R1/2/3 和 time.google.com 做 NTP 同步

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-25.1wsojifftdwg.webp)

需要先查看 time.google.com 的 DNS 记录

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-30.6rqtzxtwl3eo.webp)

然后使用 `R1(config)#ntp server <A record>` 来配置需要主动同步的 NTP 服务器

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-31.5pp3ma2edrls.webp)

最好配置多个服务器地址，如果一台服务器同步有问题，也可以和其他的服务器同步(NTP 会自动选择)

如果需要固定优先使用某台服务器，可以使用 `R1(config)#ntp server <A record> prefer`

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-33.5wc65s2zgc8w.webp)

配置完后(这里使用没有配置 prefer 的)，可以使用 `show ntp associations` 来查看 NTP 的信息

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-40.5kev90lebtds.webp)

- address

  同步的 NTP 服务器地址，有几种状态如下

  - \* sys.peer

    表示当前设备选中主动同步的服务器，例如上述是 216.239.35.0

  - \+ candidate

    表示当前设备主动同步的服务器是候选的，即备选服务器

  - \~ configured

    表示当前的服务器是手动配置的

  - \- outlyer/ x falseticker

    表示当前服务器不会被选为主动同步的服务器，即不可靠服务器

- ref clock

  例子中的 time.google.com 对应的 NTP 服务器的 reference clock 是 .GOOG.

  > 这里的 reference clock 并不是上面指的，而是 address 对应的 referece server

- st

  例子中的 216.239.35.0 和 reference clock 互联所以 stratum 值为 1

除了使用 `show ntp association` 来查看 NTP 信息外，还可以使用 `show ntp status` 来查看当前选中服务的信息

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_19-59.1knp57vyo0g0.webp)

- stratum 2

  在 NTP 中既可以是 client 也可以是 server，和 NTP 服务器同步后会自动变为 NTP 服务器，216.239.35.12 stratum 为 1，所以  R1 stratum 为 2

- reference is 216.239.35.12

  当前使用主动同步的 NTP 服务器

检查一下系统的时间

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_20-12.6ev5zih5whhc.webp)

> 使用 NTP 必须配置准确的 time zone

这里需要使用 `ntp update-calendar` 来让 hardware clock 同步 NTP 的配置，否则不会生效

**当操作系统关机或者是开启是，software clock 并不会立即启动，所以有必要让 hardware clock 也同步更新**

这里想要 R2 来同步 R1(R1 作为 NTP server)

需要使用 `R1(config)#ntp source loopback0` 来指定发送 NTP messages 的 interface (也可以理解层直接让设备变成 NTP server)

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_20-36.3y4p5e0h3vb4.webp)

> 这里使用 loopback 因为，不受物理端口限制。这里已经提前使用 OSPF 配置了 R1/R2/R3 之间的路由了

配置 R2

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_20-41.2e4jsclnsy2o.webp)

这里 reference clock 是 10.1.1.1 同步的 NTP server 即 reference；而 stratum 为 2，即 R1 到 .GOOG.；如果使用 `show ntp status` 可以看到 R3 到 .GOOG. stratum 为 3 

配置 R3

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-19.388u2i7z0qio.webp)

这里为 R3 分别添加 R1 和 R2 的 loopback 地址，看 R3 会优先选择谁，使用 `show ntp associations` 可以明确的看到会选择 R1

*Servers with lower stratum levels are preferred*

因为 R1 的 stratum 值比 R2 的小，所以会优先选择 R1

### Configuring NTP Server mode

有如下拓扑

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-26_1.7fru82osz728.webp)

如果设备没办法和其他的 NTP server 同步，但是又想要让当前的组网所有设备 NTP 同步，有什么方法呢？

需要让一台机器称为 NTP master，其他所有机器可以参考这台机器做 NTP 同步。使用 `ntp master [stratum]` 来让 R1 成为 NTP master

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-26.2weaesygilts.webp)

这里可以看到 `show ntp asso` 显示 127.127.1.1 stratum 值为 7，所以 R1 到 127.127.1.1 的 stratum 值为 8，所以即 R1 stratum 值为 8(==默认为 8==)，可以使用 `show ntp status` 来校验

这里并不能手动执行 loopback interface 或者其他地址作为 ntp server address，只能使用 127.127.1.1

> ==loopback interface 和 loopback address 是完全不同的两个概念==
>
> loopback interface 是虚拟的接口可以使用这个接口对其他设备 advertise messages，例如 OSPF,NTP 等等
>
> 而 loopback address 是设备内部的地址，不能被其他的设备直接访问

另外也可以看到 R1 stratum 值为 8，因为 R1 需要到 127.127.1.1

配置 R2/3

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-36.6k2g0hur0a2o.webp)

R2/3 均使用 R1 loopback interface 对应的地址

### Configuring NTP symmetric active mode

使用 Configuring NTP Server mode 中的拓扑和配置

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-26_1.7fru82osz728.webp)

让 R2 和 R3 之间成为 symmetric active mode，只需要使用 `ntp peer <peer ip address>` 即可

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-40.19uddunitb28.webp)

### Configuring NTP authentication

NTP authentication 是一个可选的配置，需要 clients 和 server 配置的 authentication keys 都相同

*It allows NTP clients to ensure they only sync to the intended servers*

配置需要按照如下步骤

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-57.5h1juqehrnk.webp)

1. `ntp athenticate`

   开启 NTP authencation 的功能

2. `ntp athentication-key <key-number> md5 <key>`

   key-number 用于标识密钥类似与 ID，key 表示对应的密码

3. `ntp trusted-key <key-number>`

   第几个 kye-number 对应的密钥收信任

4. `ntp server <ip-address> key <key-number>`

   使用第几个 key-number 对应的密钥和 NTP server 验证，在 server 无需使用该命令，只在 clients 上使用

使用 Configuring NTP symmetric active mode 中的拓扑和配置

R1 作为 NTP server

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_21-26_1.7fru82osz728.webp)

R1/2/3 配置如下

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_22-00.6w87rnw0plds.webp)

> 这里还使用了 `ntp peer <peer ip address> key <key number>` 来为 symmetric active mode 配置 NTP authentication

## Summary

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-07-03_22-05.713hcb02idq8.webp)

## LAB

### 0x01

Configure the software clock on R1,R2 and R3 to 12:00:00 Dec 30 2020(UTC)

```
R1#clock set 12:00:00 Dec 30 2020
R2#clock set 12:00:00 Dec 30 2020
R3#clock set 12:00:00 Dec 30 2020
```

使用 `show clock detail` 来校验

```
R1#show clock det
20:4:18.224 AS Wed Dec 30 2020
Time source is user configuration
```

### 0x02

Configure the time zone of R1,R2 and R3 to match your own

```
R1(config)#clock timezone AS 8
R2(config)#clock timezone AS 8
R3(config)#clock timezone AS 8
```

### 0x03

Configure R1 to synchronize to NTP server 1.1.1.1 over the Internet

> 配置前先看一下时间

```
R1(config)#ntp server 1.1.1.1
```

What stratum is 1.1.1.1? 

```
R1(config)#do show ntp as

address         ref clock       st   when     poll    reach  delay          offset            disp
 ~1.1.1.1       .INIT.          16   -        64      0      0.00           0.00              0.00
```

这时并没有选中 1.1.1.1 作为同步的 NTP servser，因为 NTP 同步需要时间

```
R1(config)#do show ntp as

address         ref clock       st   when     poll    reach  delay          offset            disp
*~1.1.1.1       .INIT.          0    12       16      377    0.00           0.00              0.12
```

What stratum is R1?

```
R1(config)#do show ntp st
Clock is synchronized, stratum 1, reference is 1.1.1.1
```

确认一下 R1 是否已经使用 NTP

```
R1(config)#do show clock detail
19:12:59.186 AS Wed Dec 30 2020
Time source is NTP
```

### 0x04 

Configure R1 as a stratum 8 NTP master

```
R1(config)#ntp master 8
R1(config)#do show ntp as

address         ref clock       st   when     poll    reach  delay          offset            disp
*~1.1.1.1       .INIT.          0    24       64      375    0.00           0.00              0.95
 ~127.127.1.1   .LOCL.          7    8        64      377    0.00           0.00              0.48
 * sys.peer, # selected, + candidate, - outlyer, x falseticker, ~ configured
```

Synchronize R2 and R3 to R1 with authentication

```
R1(config)#ntp authenticate 
R1(config)#ntp authentication-key 1 md5 password
R1(config)#ntp trusted-key 1

R2(config)#ntp authenticate 
R2(config)#ntp authentication-key 1 md5 password
R2(config)#ntp trusted-key 1
R2(config)#ntp server 192.168.12.1 key 1

R3(config)#ntp authenticate 
R3(config)#ntp authentication-key 1 md5 password
R3(config)#ntp trusted-key 1
R3(config)#ntp server 192.168.13.1 key 1
```

> 在 packet tracer 中不支持 `ntp source <interface>`

使用 `show ntp association` 和 `show  clock detail` 来查看 NTP 配置是否正确

```
R2(config)#do show ntp as
address         ref clock       st   when     poll    reach  delay          offset            disp
*~192.168.12.1  1.1.1.1         1    9        32      377    0.00           0.00              0.12

R3(config)#do show ntp as
address         ref clock       st   when     poll    reach  delay          offset            disp
*~192.168.13.1  1.1.1.1         1    4        16      1      0.00           0.00              0.00

R2(config)#do show clock detail
11:11:18.688 UTC Wed Dec 30 2020
Time source is NTP

R3(config)#do show clock detail
11:11:18.688 UTC Wed Dec 30 2020
Time source is NTP
```

配置完成后还需要使用 `ntp update-calendar` 来让 hardware clock 同步

```
R1(config)#ntp update-calendar 
R2(config)#ntp update-calendar 
R3(config)#ntp update-calendar 
```

**references**

1. [^jeremy’s IP Lab]:https://www.youtube.com/watch?v=qGJaJx7OfUo&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=71