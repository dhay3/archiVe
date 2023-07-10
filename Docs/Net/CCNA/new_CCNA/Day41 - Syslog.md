# Day41 - Syslog

## Syslog

Syslog 是一个用于收集日志的标准的协议

*On network devices, Syslog can be used to log events such as changes in interfaces status(up <=> down), changes in OSPF neighbor status(up<=>down), System restart,etc*

*The messages can be displayed in the CLI,saved in the device’s RAM, or sent to an external Syslog server*

例如对端口使用 `no shutdown` 设备就会立马显示两条 Syslog messages

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-42.wmox357lvk0.webp)

在实际的生产中，Syslog 和 SNMP 形成互补

## Syslog message format

Syslog 会按照 `seq:time stamp: %facility-serverity-MNEMONIC:descprtion` 来显示

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-46.2mgki1gzizb4.webp)

例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_16-51.6p8fyk5xubk0.webp)

> seq 和 time stamp 不一定会出现在 syslog messages，需要看 Syslog 配置

severity 一共有 8 个值，值越大程度越紧急,你可以使用 level 也可以使用 keyword

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-49.5frh94ire0e8.webp)

## Syslog Logging Location

Syslog messages 会出现在几个位置

1. **Console line**: Syslog messages will be displayed in the CLI when connected to the device via the console port. By default, all messages(level 0 -level 7) are displayed

2. **VTY lines**: Syslog messages will be displayed in the CLI when connected to the device via Telnet/SSH(coming in a later video). Displayed by default

   VTY = virtual teletype 虚拟终端，例如 Putty Xshell 等

3. **Buffer**: Syslog messages will be saved to RAM. By default, all messages(level 0 - level 7) are displayed

   可以通过 `show logging` 来查看在 Buffer 中的日志

4. **External server**: You can configure the device to send Syslog messages  to an external server

   Syslog servers will listen for messages on **UDP port 514**

只有 console line 和 buffer 是默认的 

## Syslog Configuration

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_17-00.110mmhf5vhxs.webp)

1. `R1(config#)logging console <severity>`

   用于配置输出小于指定等级日志到 console line

   2. `R1(config)#logging monitor <severity>`

      用于配置输出小于指定等级日志到 vty

      使用这条命令还不能将日志输出到 vty，在 vty 连接设备后还需要使用 `R1#terminal monitor` 日志才能正常输出到 vty(在 PC 上使用)

      只要连接设备，每次都需要使用这个命令，即只针对 session 有效

3. `R1(config)#logging buffered [size] <severity>`

   用于配置输出小于指定等级日志到 buffer，size 可以手动设置单位 byte

4. `R1(config)#logging <external server ip>`

   `R1(config)#logging host <external server ip>`

   用于配置输出日志到 external server 的地址，两条命令功能一样

   `R1(config)#logging trap debugging`

   用于配置输出小于指定等级日志到 external server

### logging synchronous

默认如果你在 CLI 中输命令的时候，syslog messages 可能会输出在中间，例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_17-49.1ng175dmc28w.webp)

如果需要避免这种情况出现，需要使用 logging synchronous，使用如下命令

```
R1(config)#line console 0
R1(config-line)#logging synchronous
```

> 在 Telnet/SSH 中详解

使用上面命令后就会有如下效果

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_17-52.yrvytfdgi2o.webp)

### service timestamps/service sequence-numbers

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_17-56.3wa8h2xpjxfk.webp)

- `service timestamps log <datetime|uptime>`

  设置 syslog messages 是否显示时间戳，一般使用 datetime

- `service sequence-numbers`

  设置 syslog messages 是否显示 seq number，一般不用也没什么关系

## Command summary

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_17-58.3dpnxlk3oc1s.webp)

## Syslog VS SNMP

Syslog 和 SNMP 虽然都用于监控设备的信息，但是两者这间功能和目的不同

- Syslog is used for message logging

  Events that occur within the system are categorized based on facility/severity and logged

  Used for system management, analysis, and troubleshooting

  Messages are sent from the devices to the sever. The server can’t active pull information from the devices(like SNMP Get) or modify variable(like SNMP Set)

- SNMP is used to retrieve and organize information about the SNMP managed devices

  IP addresses, current interface status, temperature, CPU usage, etc

  SNMP servers can use Get to query the clients and Set to modify variables on the clients

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_18-11.35s4ahh5y08w.webp)

### 0x01

Connect to R1’s console port using PC2

> console 是物理接口，只要插上线就可以连，不需要配置 类似 22 这种逻辑端口

Desktop -> Terminal 点 OK 即可

Shut down the G0/0 interface

```
R1(config)#int g0/0
R1(config-if)#shutdown

R1(config-if)#
%LINK-5-CHANGED: Interface GigabitEthernet0/0, changed state to administratively down

%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet0/0, changed state to down
```

After you receive a syslog message, re-enable the interface

```
R1(config-if)#no shutdown

R1(config-if)#
%LINK-5-CHANGED: Interface GigabitEthernet0/0, changed state to up

%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet0/0, changed state to up
```

等级 5 对应 notice

Enable timestamps for logging message

```
R1(config)#service timestamps log  datetime msec 
```

### 0x02

Telnet from PC1 to R1’s G0/0 interface

```
C:\>telnet 192.168.1.1
Trying 192.168.1.1 ...Open


User Access Verification

Username: jeremy
Password: 
R1>en
Password: 
R1#
```

> 这里配置默认使用 telnet 23 端口，所以无需指定端口

Enable the unused G0/1 interface

```
R1#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
R1(config)#int g0/1
R1(config-if)#no shutdown
```

这里并不会将 Syslog message 输出到 VTY，因为是远程连接的，默认需要使用 `terminal monitor` 来允许 Syslog message 输出到 VTY

Enable logging to the VTY lines for the ==current session==

```
R1(config-if)#do terminal monitor
```

> 在 packet tracer 中没有 `logging monitor` 命令，但是默认开启将 Sys

### 0x03

Enable logging to the buffer, and configure the buffer size to 8192 bytes

默认不会开启 buffer logging，可以使用 `show logging` 来查看

```
R1(config)#do show logging
Syslog logging: enabled (0 messages dropped, 0 messages rate-limited,
          0 flushes, 0 overruns, xml disabled, filtering disabled)
No Active Message Discriminator.
No Inactive Message Discriminator.
    Console logging: level debugging, 3 messages logged, xml disabled,
          filtering disabled
    Monitor logging: level debugging, 3 messages logged, xml disabled,
          filtering disabled
    Buffer logging:  disabled, xml disabled,
          filtering disabled
```

如果需要开启 logging buffer 需要使用如下命令

```
R1(config)#logging buffered 8192
```

默认 buffer logging 为 debug

```
R1(config)#do show logging
Syslog logging: enabled (0 messages dropped, 0 messages rate-limited,
          0 flushes, 0 overruns, xml disabled, filtering disabled)
No Active Message Discriminator.
No Inactive Message Discriminator.
    Console logging: level debugging, 3 messages logged, xml disabled,
          filtering disabled
    Monitor logging: level debugging, 3 messages logged, xml disabled,
          filtering disabled
    Buffer logging:  level debugging, 0 messages logged, xml disabled,
          filtering disabled
    Logging Exception size (4096 bytes)
    Count and timestamp logging messages: disabled
    Persistent logging: disabled
No active filter modules.
ESM: 0 messages dropped
    Trap logging: level informational, 3 message lines logged
Log Buffer (8192 bytes):
```

`Trap logging: level informational` 可以得出发送到 remote server 的日志等级为 informational 以上的

### 0x04

Enable logging to the syslog server SRV1 with a level of debugging

```
R1(config)#logging 192.168.1.100
R1(config)#logging trap debugging 
```

同样的也可以使用 `show logging ` 来查看配置是否准确

```
R1(config)#do show logging
ESM: 0 messages dropped
    Trap logging: level debugging, 3 message lines logged
        Logging to 192.168.1.100  (udp port 514,  audit disabled,
             authentication disabled, encryption disabled, link up),
             0 message lines logged,
             0 message lines rate-limited,
             0 message lines dropped-by-MD,
             xml disabled, sequence number disabled
             filtering disabled
Log Buffer (8192 bytes):
```

**references**

1. [^https://www.youtube.com/watch?v=RaQPSKQ4J5A&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=79]

