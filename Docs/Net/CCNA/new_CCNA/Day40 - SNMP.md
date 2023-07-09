# Day40 - SNMP

## SNMP

Simple Network Management Protocol(SNMP) 是一个标准的 IEEE 协议，通过 SNMP 可以监控设备的状态，配置的变化等

在 SNMP 中有两种类型的角色

1. Managed Devices

   These are the deivces being managed using SNMP

   For example, network devices like routers and switches

2. Network Management Station(NMS)

   The device/devices managing the managed devices

   This is the SNMP server

例如下面拓扑，SRV1 是 NMS，R1 SW1 均为 Managed Devices

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_13-37.5pniwj6v61vk.webp)

如果使用了 SNMP 会出现如下结果

1. Managed devices can notify the NMS of events

   假设 SW1 G0/1 端口 down 或者 up，SW1 会通过 SNMP 把对应的信息发送到 NMS

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_13-43.1ke5b45c7qf4.webp)

2. The NMS can ask the managed devices for information about their current status

   例如 NMS 会主动问 R1 的 CPU 使用率

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_13-42.7j2hlcdgxam8.webp)

3. The NMS can tell the managed devices to change aspects of their configuration

   假设 R1 G0/1 接口地址为 203.0.113.1, NMS 可以要求 R1 G0/1 接口地址变为 203.0.113.1

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_13-44.2d0kmyghce9s.webp)

## SNMP Components

上面的拓扑可以抽象成如下几个组件

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_13-46.3ivgsfvlpojk.webp)

> SNMP Agent uses UDP 161 port
>
> SNMP Manager uses UDP 162 port 

### NMS

1. The SNMP Manager is the software on the NMS that interacts with managed devices. It receives notifications, sends requests for information, send configuration changes, etc

   可以把 SNMP Manager 看成部署在主机上的后端应用

2. The SNMP Application provides an interface for the network admin to interact with. It display alerts statistics, charts

   可以把 SNMP Application 看成部署在主机上的前端应用，通过这个前端应用可以各个 SNMP Managed devices 的信息

### Managed devices

1. The SNMP Agent is the SNMP software running on the managed devices that interact with the SNMP Manager on the NMS

   It sends notifications to/receives massages from the NMS

2. The Management Information Base(MIB) is the structure that contains the variables that are managed by SNMP

   Each variable is indentified with an Object ID(OID)

   Example variables: Interface status, traffic throughput, CPU usage, temperature

#### SNMP OIDs

*SNMP Objects IDs are organized in a hierarchical structure*

SNMP OID 是以层级表示的，例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_13-59.5xz8ze52f1fk.webp)

前一部分数字是后一部分数字的父级(类似 json path)，整个 OID 代表 sysName

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-01.6yqsw9ncxrls.webp)

如果 SRV1 向 SW1 需求 OID .1.3.6.1.2.1.1.5 的值，SW1 就会回送 OID 对应的值 sysName SW1 到 SRV1

> 如果需要查看其他的 OID，可以到 oid-info.com 

## SNMP Versions

SNMP 起源于 1988 年，有很多不同版本的 SNMP，目前使用的最广泛的有 3 个版本

1. SNMPv1

   the original version of SNMP

2. SNMPv2c

   *Allows the NMS to retrieve large amounts of information in a single request, so it is more efficent*

   ‘c’ refers to the ‘community strings’ used as password in SNMPv1, removed from SNMPv2, and then added back for SNMPv2c

3. SNMPv3

   A much more secure version of SNMP that supports strong encryption and authentication Whenever possiable, this version should be used

   安全性更高

## SNMP Messages

NMS 或者 managed devics 会发送如下 NMS messages

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-10.1b610s2bzwow.webp)

Read/Write classes 是从 NMS 发送到 managed devices 的

Notification classes 是从 managed devices 发送到 NMS 的

Response 是针对 message 的回送，除 Trap 外

### Read

Read class messages 中可以包含如下几种 messages

1. Get

   A request sent from the manager to the agent to retrive the value of a variable(OID), or multiple vairables. The agent will send a Response message with the current value of each vairable

   例如 SRV1 发送 Get 问 SW1 G0/1 端口的状态 

   ![https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-16.6qpt6tsc0xs0.webp](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-16.6qpt6tsc0xs0.webp)

2. GetNext

   A request sent from the manager to the agent to discover the available variables in the MIB

   > it says ‘tell me the next OID’, so it can be used to discover what other OIDs are available

   即获取下一个 OID，类似JAVA interation 中的 next()函数

3. GetBulk

   A more efficient version of the GetNext messages(introduced in SNMPv2)

   可以一次性获取多个 OID

### Write

Write class messages 中只包含一种 message

Set

A request sent from the manager to the agent to change the value of one or more variables. The agent will send a Response message with the new values

例如 SRV1 告诉 SW1 把 host name 改成 SW10 就需要通过 Set messages

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-24.4a7t73va5br4.webp)

### Notification

Notification class messages 中包含 2 种 message

1. Trap

   A notification sent from the agent to the manager. The manager does not send a Response message to acknownledge that it received the Trap, so these messages are ‘unreliable’

   例如当 SW1 G0/0 端口 down 了就会发送 Trap message 到 SRV1

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-32.1omkidz47qbk.webp)

2. Inform

   A notification message that is acknowledged with a Response message

   *Originally used for communications between managers, but later updates allow agents to send inform messages to managers,too*

### Response	

Notification class messages 只包含 1 种 message 就是 Response，除了发送 Trap message 外不会回送 Response，其他类型的 messages 都会收到 Response

## SNMPv2c Configuartion

> SNMP 配置并不在 CCNA 考试的范围内，这里只介绍 managed devices 的配置

有如下拓扑，PC1 是 NMS,R1 是 managed device

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-39.3gkd25elxxa8.webp)

使用如下配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-43.2a1cgdiuw4w0.webp)

1. `snmp-server community Jeremy1 ro`

   表示如果 NMS 使用 Jeremy1 passord，只有 read only 的权限，只能使用 Read class messages

2. `snmp-server community Jeremy2 rw`

   表示如果 NMS 使用 Jeremy2 password，可以有 read/write 的权限

3. `snmp-server host 192.168.1.1 version 2c Jeremy1`

   用于指定 NMS 地址，以及需要使用的 password

4. `snmp-server enable traps snmp linkdown linkup`

   linkdown/linkup 会以 trap 的方式发送到 NMS

5. `snmp-server enable traps config`

   配置发生改变会以 trap 的方式发送到 NMS

例如现在 R1 G0/1 linkdown 就会发送 trap message

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-49.4q3hqzni9ocg.webp)

报文详情如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230630/2023-07-06_14-50.6t6wtttijzs.webp)

报文的 variable-bindings 对应 OIDs

例如 1.3.6.1.6.3.1.1.5.3 就对应 linkdown 的 OID

也可以看到使用的 SNMP 版本是 SNMPv2c

community 使用 Jeremy1 

> 虽然 commuity string 表示的是 password，但是在 SNMPv1 和 SNMPv2 中都是以明文表示的，所以是不安全的。而 SNMPv3 报文是加密的，==所以在实际生产的环境中应该要使用 SNMPv3==

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-11.2l295nvmk03k.webp)

### 0x01

Configure the following SNMP communities on R1

read-only:Cisco1

read/write:Cisco2

```
R1(config)#snmp-server community Cisco1 ro
R1(config)#snmp-server community Cisco2 rw
```

在 packettracer 中只能配置这两条命令，所以这里仅仅使用默认的 SNMPv1

### 0x02

Use SNMP ‘Get’ messages via the MIB browser on PC1 to check the following

Desktop -> MIB 然后选择 Advanced 按照下面截图配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-21.o3kziy8smts.webp)

Read Community 对应 ro Cisco1

Write Community 对应 rw Cisco2

1. How long has R1 been running(system uptime)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-24.69udwr0z20ow.webp)

2. What is the currently configured hostname on R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-25.2ek0b0r6nbwg.webp)

3. How many interfaces does R1 have

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-26.3z07gdtusagw.webp)

4. What are those interfaces

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-27.3gr5n714e3eo.webp)

如果需要看 interfaceName，需要使用 ifDescr

### 0x03

Use an SNMP ‘Set’ message from PC1 to change the hostname of R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-06_15-29.4tj2cq28dqf4.webp)

这里的 datatype 必须和需要设置的值的 type 相同，点击 OK 后还需要点击 GO 才会生效

**references**

1. [^https://www.youtube.com/watch?v=HXu0Ifj0oWU&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=77]