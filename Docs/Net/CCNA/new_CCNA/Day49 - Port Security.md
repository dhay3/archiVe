# Day49 - Port Security

## Port Security

Port Security 是思科交换机上的一个安全功能

*It allows you to control which source MAC address(es) are allowed to enter the switchport.*

如果收到含有 unauthorized source MAC 的报文，会导致

*The default action is to place the interface in an ‘err-disabled’ state*

例如当前 PC1 直连 SW1, A.A.A 是被授权 Source MAC，PC1 可以正常访问到 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_13-58.vxz75jbfybk.webp)

现在将 PC1 和 SW1 直连的网线拔下来，改成 PC2 和 SW1 互联

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-00.56c2nplptri8.webp)

因为 B.B.B 不是被授权的地址，当 SW1 通过 G0/1 收到 PC2 发过来的报文就会将 G0/1 置为 err-disable 状态

现在将链路改回原来的状态

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-03.3460lzez24c.webp)

此时 PC1 任然是不可以访问 R1 的，因为 SW1 G0/1 的状态并不会自动改变，而是 err-disable 的状态

关于 Port security authorized MAC address 有如下的几条规则

1. When you enable port security on an itnerface with the default settings, ==one MAC address is allowed==

   - You can configure the allowed MAC address manually
   - If you don’t configure it manually, the switch will allow the first source MAC address that enters the interface

2. You can change the maximum number of MAC addresses allowed

   例如 在 IP phone 的场景在 PC1 和 PH11 都要通过 SW1 访问 R1

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-09.10rkifq6wwvk.webp)

   如果这时还是只能配置一个 MAC address.那明显不合理。所以 Port security 允许的 MAC 地址数量不
   
   是固定为 1 的，当收到报文的源 MAC 是 A.A.A 时 SW1 就会自动学习，将 A.A.A 加入到 authorized 中，同理 C.C.C

## Why Port Security

虽然 port security 能限制 Source MAC 的地址，但是 MAC address spoofing 是非常容易的一件事，所以通常只会使用 port security 来限制端口允许的 MAC address 数

例如

1. the attacker spoofed thousands of fake MAC addresses

2. the DHCP server assigned IP addresses to these fake MAC address, exhausting the DHCP pool

3. the switch’s MAC address table can also become full due to such an attack

   Switch 因为会学习 Source MAC 地址，所以也会出现 MAC address 地址打爆的情况

*Limiting the number of MAC addresses on an interface can protect against those attacks*

所以限制端口 MAC addresses 的数量能很好的防止这种攻击

## Port Security Configuration

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-28.3eau1me63aio.webp)

让端口开启 port security 的功能很简单，值需要使用 `SW1(config-if)#switchport port-security [mac-address <address>]`  即可(添加 mac-address 参数表示手动指定 authorized MAC)

但是需要注意的是，==只有显式的声明 switchport 是 access mode 或者是 trunk mode 才可以(默认开启 DTP 是不支持，所以会报错)==

> 默认只允许第一个收到报文中的 Source MAC 作为 authorized MAC

如果想要查看端口是否开启了 port security 的功能，可以使用 `show port-security interface <interface-id>`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-41.115qy6lg6mts.webp)

1. `Port Security： Enabled`

   表示端口开启了 port security 功能

2. `Port Status： Secure-up`

   表示端口开启了 port security 功能，并且是 up/up 的

3. `Violation Mode: Shutdown`

   即收到 unauthorized MAC 时的操作，关闭端口，即 err-disable

4. `Aging Time： 0 mins`

   表示 authorized MAC address 不会过期，默认为 0

   可以通过 `switchport port-security aging time <mins>` 来手动配置

5. `Aging Type: Absolute`

   一共有几个值

   - Absolute

     After the secure(authorized) MAC address is learned, the aging timer starts and the MAC is removed after the timer expires, even if the switch continues receiving frames from that source MAC address

   - Inactivity

     After the secure MAC address is learned, the aging timer starts but is reset every time a frame from that source MAC address is received on the interface

     意味着如果一直收到对应的报文，就永不会过期

     可以通过 `switchport port-security aging type {absolute | inactivity}` 来配置

   > 注意的一点是如果是通过 `switchport port-security mac-address <address>` 手动指定 authorized MAC 的话，也是永远不会过期的(即只有自动学习的 MAC 才会过期)。但是可以通过`switchport port-security aging static` 来声明让手动配置的 authorized MAC 也会按照 Timer 过期

6. `Maximum MAC Addresses: 1`

   表示 authorized MAC 最大的数量

7. `Total MAC Addresses: 0`

   表示当前收到的 MAC 数量

8. `Configured MAC Address: 0`

   手动配置的 MAC 地址数量

8. `Last Source Address:Vlan:0:0:0`

   表示最近一次收到报文的 Source MAC

9. `Security Violation Count: 0`

​	表示收到 unauthorized MAC 的次数

> 如果是默认的 shutdown violation mode，最大值只能是 1，因为进入 err-disable 就收不到其他的报文了

假设现在 PC1 ping R1，使用 `show port-security interface <interface-id>` 来查看

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-43.m02762i91yo.webp)

这里可以看到 Total MAC Addresses 和 Last Source Address 对应的值都发生改变了

现在将 PC2 和 SW1 直连，PC2 ping R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-50.2xyew5e4d5kw.webp)

可以看到 Port Status 从 Secure-up 变成 Secure-shutdown 了，并且使用 `show interfaces status` 查看端口时，可以看见是 err-disabled

同时 Tatal MAC Addresses 重置为 0(因为端口被关闭了)，且 Last source Address 值变为 PC2 MAC，Security Violation Count 加 1

除了 `show port-security interface <interface-id>` 外来查看 port security，还可以使用 `show port-security` 来查看所有开启 port security 功能的端口列表

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_19-34.3jmg1ezipd4w.webp)

### Re-enabling an interface

现在需要让端口正常，该如何操作呢？有两种方式

1. manually
2. errDisable recovery	

#### Manually

手动使用 `shutdown` 和 `no shutdown` 命令来开启端口，和让 err-disable 端口正常的操作一样

> 需要注意的是需要先将 PC2 断开连接，否则当 PC2 发送报文时，SW1 将会将 PC2 MAC 作为 authorized MAC

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_14-56.5fsuvo3w30xs.webp)

这是使用 `show port-security interface g0/1` 查看端口的状态，就会发现复原成 PC1 没有发报文到 R1 的状态了

#### ErrDisable Recovery	

可以让 errdisable 的端口在指定的时间内 enable

端口进入 errdisable 状态通常有如下几种原因，可以使用 `SW1#show errdisable recovery` 来查看

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_15-02.3ss6y4p0e5c.webp)

Timer Status 对应的列会显示 errdisable recovery 功能是否启用，默认均为 disable 表示不开启 errdisable recovery 的功能

Timer Interval 表示

*Every 5 minutes(by default), all err-disabled interfaces will be re-enabled if err-disable recovery has been enabled for the cause of the interface’s disablement*

现在开启 port-security errdisable recovery 的功能

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_15-11.2ogxuxc05e0w.webp)

只需要使用 `SW1(config)#errdisabel recovery cause psecure-violation`，为了方便查看 errdisable recovery 的功能是否生效，这里将 Timer Interval 的值改小 为 3 mins `SW1#errdisable recovery interval 180`

如果使用 `SW1#show errdisable recovery` 来查看 errdisable recovery 的功能是否开启，就可以发现 Timer status 对应的字段值为 disable，同时 Timer Interval 变成 180 seconds，当 180 后端口还是处于 errdisable 就会尝试自动 enable 端口(即执行 shutdown/no shutdown)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_15-14.1ljlmnxdrq1s.webp)

> 需要注意的一点是
>
> *Errdisable recovery is useless if you don’t remove the device that caused the interface to enter the err-disable state*

## Violation Modes

Violation mode 是当收到 unauthorized Source MAC 报文时，端口处理的策略，一共 3 种

1. shutdown
2. restrict
3. protect

### Shutdown

- Effectively shuts down the port by placing it in an err-disabled state
- Generates a Syslog and/or SNMP message when the interface is disabled
- The violation counter is set to 1 when the interface is disabled

### Restrict

- The switch discards traffic from unauthorized MAC addresses
- The interfaces is NOT disabled
- Generates a Syslog and/or SNMP message each time an unauthorized MAC is detected.
- The violation counter is incremented by 1 for each unauthorized  frame

只需要使用 `SW1(config-if)#switchport port-security violatioin restrict` 就可以改为 restrict violation mode

这里 PC2 ping R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_15-30.1l6zwo1mz6ww.webp)

因为手动设置 authorized MAC 是 000a.000a.000a，PC2 MAC 不是 authorized MAC，且是 restrict violation mode，所以 SW1 G0/1 port status 状态认为 secure-up 同时 security violation count 计数上涨

### Protect

- The Switches discards traffic from unauthorized MAC addresses
- The interface is NOT disabled
- It does NOT generate Syslog/SNMP messages for unauthorized MAC
- It does NOT increment the violation counter

和 restrict violation 配置的方式一样

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_19-08.1xuscchadusg.webp)

如果从 PC2 ping R1 这里可以注意到 Security Violation Count 值是不会增加的

## Sticky Secure MAC Address

Sticky 是指学到的 MAC address 永远不会过期，可以通过 `SW1(config-if)#switchport port-security mac-address sticky` 来开启 sticky 的功能

等价于在 running-config 中(如果想要保存，就需要使用 write 写到配置中)使用了 `SW1(config-if)#switchport port-security mac-address sticky <learned address>` 

一旦使用了 `switchport port-security mac-address sticky` 会将当前所有的 secured MAC 直接转为 sticky MAC

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_19-44.1u1bxp0k68bk.webp)

## MAC Address Table

不管是 Secure MAC 还是其他的 MAC 都会被加入到 MAC Address Table 中

- Sticky and Static secure MAC addresses will have a type of STATIC
- Dynamically-learned secure MAC addresses will have a type of DYNAMIC

可以通过 `show mac address-table secure` 来查看 secure MAC address 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_19-49.6zmpdpl4cge8.webp)

## Command summary

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_19-52.517qz0c2gwsg.webp)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_20-01.1ztdw90cd3k0.webp)

### 0x01

Configure port security on the following interfaces:

\#SW1 F0/1, F0/2, F0/3#

Violation mode: Shutdown

Maximum addresses: 1

Sticky learning: Disabled

Aging time: 1 hour

因为默认链路开启 DTP，所以需要显式的声明为 access mode

```
SW1(config)#int range f0/1-3
SW1(config-if-range)#switchport mode access
SW1(config-if-range)#switchport port-security
SW1(config-if-range)#switchport port-security violation shutdown 
SW1(config-if-range)#switchport port-security maximum 1
SW1(config-if-range)#switchport port-security aging time 60
```

配置完成后通过 `show port-security int <interface-id>` 来校验

```
SW1(config-if-range)#do show por int f0/1
Port Security              : Enabled
Port Status                : Secure-up
Violation Mode             : Shutdown
Aging Time                 : 60 mins
Aging Type                 : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses      : 1
Total MAC Addresses        : 0
Configured MAC Addresses   : 0
Sticky MAC Addresses       : 0
Last Source Address:Vlan   : 0000.0000.0000:0
Security Violation Count   : 0
```

通过 `show port-security` 来查看所有开启 port security 功能的接口

```
SW1(config-if-range)#do show port
Secure Port MaxSecureAddr CurrentAddr SecurityViolation Security Action
               (Count)       (Count)        (Count)
--------------------------------------------------------------------
        Fa0/1        1          0                 0         Shutdown
        Fa0/2        1          0                 0         Shutdown
        Fa0/3        1          0                 0         Shutdown
----------------------------------------------------------------------
```

#SW2 G0/1#

Violation mode: Restrict

Maximum addresses: 4

Sticky learning: Enabled

```
SW2(config-if-range)#int g0/1
SW2(config-if)#switchport port-security
SW2(config-if)#switchport port-security violation restrict 
SW2(config-if)#switchport port-security mac-address sticky 
SW2(config-if)#switchport port-security maximum 4
SW2(config-if)#do show port int g0/1
Port Security              : Enabled
Port Status                : Secure-up
Violation Mode             : Restrict
Aging Time                 : 0 mins
Aging Type                 : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses      : 4
Total MAC Addresses        : 1
Configured MAC Addresses   : 0
Sticky MAC Addresses       : 0
Last Source Address:Vlan   : 0060.471C.1D19:1
Security Violation Count   : 0
```

### 0x02

Trigger port security violations on SW1 and SW2 (for example by  connecting another PC) and observe the actions taken by each switch.

新挂一台机器到 SW1 即可

```
SW2#show port int g0/1
Port Security              : Enabled
Port Status                : Secure-up
Violation Mode             : Restrict
Aging Time                 : 0 mins
Aging Type                 : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses      : 4
Total MAC Addresses        : 4
Configured MAC Addresses   : 0
Sticky MAC Addresses       : 3
Last Source Address:Vlan   : 0060.471C.1D19:1
Security Violation Count   : 10
```

> 在 packettracer 中不会显示 Syslog,实际 restrict violation 会显示 Syslog

**references**

1. [^https://www.youtube.com/watch?v=sHN3jOJIido&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=95]