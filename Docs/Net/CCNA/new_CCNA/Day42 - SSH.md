# Day42 - SSH

## Console line

默认通过 console port 登录思科设备(console line)执行命令并不需要密码，但是你也可以手动为 console line 配置密码

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_10-48.3dwr5h5074e8.webp)

1. `R1(config)#line console 0`

   Only a single console

   意味着同时只能有一个设备通过 console port 连接设备来修改配置，多台设备不能通过 console ports 连接同时修改配置

2. `R1(config-line)#password ccna`

   配置 console line 需要使用的密码

3. `R1(config-line)#loging`

   开启使用 console line 登录密码

你也可以手动指定用户名，只有指定的用户使用了指定的密码才能可以使用 console line

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_10-54.wd1jdie5tn4.webp)

1. `R1(config)#username <uname> secrect <password>`

   指定需要使用的用户名和密码

2. `R1(config)#line console 0`

   和上面的含义一样

3. `R1(config-line)#login local`

   开启使用 console line 登录使用用户名和密码

> 注意这里黄框中是配置了 `login local` 后的配置，之前配置的 `password ccna` 不会生效

4. `exec-timeout <mins> <secs>`

   登录用户如果在时间内没有输入任何的指定就会断开 session

## Layer 2 Switch - Management IP

3 层的设备可以配置 IP，我们可以使用这个 IP 来远程登录 3 层设备

*Layer 2 switches don’t perform packet ruting and don’t build a routing table. They aren’t IP routing aware.*

*However, you can assign an IP address to an SVI to allow remote connections to the CLI of the switch(using Telnet or SSH)*

2 层设备不会对 3 层报文头进行 encapsulation 或者是 de-encapsulation，不具备路由转发的功能，但是可以为 2 层设备 SVI 配置 IP 地址用于远程连接

以如下拓扑为例

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_11-18.hoah62rphnk.webp)

PC2 是网络管理员的主机，需要通过 PC2 对整个网络拓扑中的设备进行配置

交换机需要使用如下配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_11-21.c3tgj77f86o.webp)

1. `SW1(config)#interface valn<vlan-id>`

   `SW1(config-if)#ip address <address> <mask>`

   `SW1(config-if)#on shutdown`

   配置交换机的 SVI

   > 2 层交换机同样也有 SVI，不仅限于 3 层交换机

2. `SW1(config)#ip default-gateway <gw address>`

   如果 PC2 想要直接访问 SW1, 因为不在一个 LAN 中，入向可以到 SW1，但是回向 SW1 不能到 PC2，所以需要配置指向 R1 的默认路由

   SW2 无需使用该配置

   > 这里并不会开启 SW1 路由的功能也不会插入路由表

## Telnet

Telnet(Teletypee Network)is a protocol used to remotely access the CLI of a remote host

Telnet 报文均以明文显示，所以安全系数很低

例如如下报文 10.0.0.2 telnet 10.0.0.1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_11-34.erx39aw5148.webp)

10.0.0.1 回送 password 要求 10.0.0.2 输入密码，10.0.0.2 输入的密码以明文显示

> Telnet server 默认监听 23 端口

### Telnet configuration

如果需要让设备开启允许 telnet 登录的功能，需要使用如下命令

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_11-55.3ysff8zq7oow.webp)

1. `SW1(config)#enable secret <password>`

   指定使用 exec mode 的密码，在必须配置该命令，否则不能进入 exec mode

2. `SW1(config)#username <username> secret <password>`

   可选，登录 SW1 需要的用户名和密码

3. `SW1(config)#access-list <number> permit host <source addr>`

   可选，指定能登录到设备的用户 ACL

   一般需要配置，同样也有 implicit deny 的规则

4. `SW1(config)#line vty 0 15`

   同一时间允许连接的最大 VTY 数量，一共 16 个

5. `SW1(config-line)#login local`

   可选，开启使用用户和密码登录设备

6. `SW1(config-line)#exec-timeout <mins> <secs>`

   可选，用户在指定时间内没有输入任何指令会直接断开 session

7. `SW1(config-line)#transport input telnet`

   允许 PC VTY 通过什么协议连接设备，这里使用 telnet

8. `SW1(config-line)#access-class <num> in`

   可选，将 ACL 应用在 VTY 0 15 上

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_12-06.3y0k9sxv2934.webp)

这里 R2 ping SW1 SVI 同样是放行的，只有在使用 telnet 时(dport 为 23)时 R2 不能访问 SW1 对应 `transport input telnet` 逻辑

在有些老一点的 IOS 上使用 `line vty 0 15` 来配置 VTY 会被拆分成 2 部分，`vty 0 4` 和 `vty 5 15`，因为只支持 vty 0 4

## SSH

SSH(Secure Shell) 是用于替代不安全的登录协议(例如 telnet)而被开发的

*Provides security features such as data encryption and authentication*

SSH 主要有两种协议 SSHv1 和 SSHv2, 如果设备同时支持 SSHv1 和 SSHv2 会以 version 1.99 显示

SSH 报文摘要如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-12.2swga29srqrk.webp)

Encrypted Packet 对应加密的 message 部分

### SSH Configuration

配置 SSH 需要包含如下几步

1. Configure host name

   必须配置 hostname 否则不能生成 RSA keys

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-56.40z3n4bmv98g.webp)

2. Configure DNS domain name

   必须配置 domain name 否则不能生成 RSA keys

   ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-56.40z3n4bmv98g.webp)

> 实际上可以直接为 key pair 指定名字，这个不在 CCNA 考试范围内，所以不介绍

3. Generate RSA key pair

4. Configure enable password, username/password

5. Enable SSHv2(only)

   isn’t necessary, but recommanded

6. Configure VTY lines
7. Connect from PC using `ssh username@ip-addr`

### Generate RSA keys

在配置 SSH 前可以使用 `SW1#show version` 来查看对应设备的 IOS 是否支持 SSH

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-15_1.60fuwi42bnk.webp)

如果 IOS 支持 SSH，一般 IOS images name 中会包含 K9。例如这里设备对应的 IOS 名为 vios_12-ADVENTERPRISEK9-M

除此外还可以使用 `SW1#show ip ssh` 来查看当前设备是否支持 SSH

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-15.tdipyrp9jk0.webp)

- `SSH Disabled - version 1.99`

  表示当前设备支持 SSHv1 和 SSHv2，但是并没有启用

在确认设备支持 SSH 后，需要生成公私钥密对

> 如果要使用 SSH 必须生成 RSA keys，这个逻辑和认证不同，并不是不需要公钥认证就不需要创建 SSH key pair

如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-32.38cdog7bhc74.webp)

- `SW(config)#ip domain name jeremysitelab.com`

  这里配置 domain name 是因为在生成密钥时，会以 FQDN 来命名

  > FQDN = Full Qualified Domain Name(host name + domain name)

  这里可以在 `The name for the keys will be` 看到

- `SW1(config)#crypto key generate rsa`

  生成 rsa 密钥对，也可以使用 `crypto key generate rsa modulus <length>` 来生成 rsa 密钥对

  > 注意 SSHv2 最小的 bit length 为 768

在配置完 SSH 后，可以发现 cosole 会输出一行对应的 syslog message，同时使用 `show ip ssh` 可以看到 `SSH Enabled - version 1.99` 表示正在使用 SSH

### VTY Line SSH Configuration

在 enable SSH(生成 RSA 密钥对) 后，还需要应用 SSH，和配置 telnet 登录类似

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_13-42.29r3t09yhq9s.webp)

1. `SW1(config)#enable secret <password>`

   进入 exec mode 使用的密码，必须配置，否则不能进入 exec mode

2. `SW1(config)#username <username> secret <password>`

   登录设备使用的用户名和密码

3. `SW1(config)#access-list <number> permit host <source addr>`

   可选，指定能登录到设备的用户 ACL

   一般需要配置，同样也有 implicit deny 的规则

4. `SW1(config)#ip ssh version [1|2]`

   指定使用的 SSH 版本，推荐使用 SSHv2 因为更安全

5. `SW1(config)#line vty 0 15`

   同一时间允许连接的最大 VTY 数量，一共 16 个

6. `SW1(config-line)#login local`

   可选，开启使用用户和密码登录设备

7. `SW1(config-line)#exec-timeout <mins> <secs>`

   可选，用户在指定时间内没有输入任何指令会直接断开 session

8. `SW1(config-line)#transport input ssh`

   允许 PC VTY 通过什么协议连接设备，这里使用 ssh

9. `SW1(config-line)#access-class <num> in`

   可选，将 ACL 应用在 VTY 0 15 上

## password vs secret

在使用 `enable` 命令补全时，可以看到有两种

```
SW2(config)#enable ?
  password  Assign the privileged level password
  secret    Assign the privileged level secret
```

或者是 `username <name>` 补全时

```
SW2(config)#username jeremy ?
  password   Specify the password for the user
  privilege  Set user privilege level
  secret     Specify the secret for the user
  <cr>
```

password 和 secret 有什么区别呢？

**如果使用 password 密码会以明文的形式存储在配置中，如果使用 secret 密码会以加密的形式存储在配置中，可以使用 `show running-config` 来校验**

> 所以推荐使用 `secret` 替代 password

## Command Summary

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_14-04.3lffjkswfxj4.webp)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_14-13.b1t7bontcs0.webp)

### 0x01

Connect Laptop 1 to SW2’s console port to perform the following configurations

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230706/2023-07-10_14-22.1l6b3e64dexs.webp)

拓扑连接后选中 laptop 中的 desktop terminal 连接 SW1

```
Switch(config)#hostname SW2
SW2(config)#enable secret ccna
SW2(config)#username jeremy password ccna
SW2(config)#int vlan1
SW2(config-if)#ip add 192.168.2.253 255.255.255.0
SW2(config-if)#no shutdown
SW2(config-if)#exit
SW2(config)#ip default-gateway 192.168.2.254
```

> 注意大部分 SVI 接口默认都是 shutdown 的，所以需要使用 `no shutdown` 来开启接口
>
> 这里如果只需要 Laptop 访问 SW1, 就不需要使用 `ip default-gateway`，如果需要 SW2 访问 laptop 就需要使用 `ip default-gateway`

### 0x02

Configure the following console line security settings on SW2

Authentication: Local user

Exec timeout: 5 minutes

```
SW2(config)#line console 0
SW2(config-line)#login local
SW2(config-line)#exec-timeout 5 0
```

因为目前 Laptop1 是通过 console 口和 SW2 互联的，所以退出当前终端就可以校验配置是否正确

```
SW2#exit
User Access Verification

Username: jeremy
Password: 

SW2>en
Password: 
SW2#
```

### 0x03

Configure SW2 for remote access via SSH

Domain name: jeremysitlab.com

```
SW2(config)#ip domain-name jeremysitlab.com
```

RSA key size：2048 bits

```
SW2(config)#crypto key generate rsa
The name for the keys will be: SW2.jeremysitlab.com
Choose the size of the key modulus in the range of 360 to 4096 for your
  General Purpose Keys. Choosing a key modulus greater than 512 may take
  a few minutes.

How many bits in the modulus [512]: 2048
% Generating 2048 bit RSA keys, keys will be non-exportable...[OK]
```

Authentication: Local user

```
SW2(config)#line vty 0 15
SW2(config-line)#login local
```

Exec timeout: 5 minutes

```
SW2(config-line)#exec-timeout 5 0
```

Protocols: SSH only

```
SW2(config-line)#transport input ssh 
```

+limit access to PC1 only

```
SW2(config-line)#exit
SW2(config)#access-list 1 permit host 192.168.1.1
SW2(config)#line vty 0 15
SW2(config-line)#access-class 1 in
```

从 R2 测试

```
R2#ping 192.168.2.253

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.2.253, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 0/0/0 ms

R2#ssh -l jeremy 192.168.2.253

% Connection refused by remote host
```

这里可以观察到 ICMP 是正常的，但是 SSH 是不能连接到 SW1

> 在 Cisco 的设备上不能通过 `ssh username@address` 的方式登录到设备

从 PC1 测试

```
C:\>ping 192.168.2.253

Pinging 192.168.2.253 with 32 bytes of data:

Request timed out.
Reply from 192.168.2.253: bytes=32 time<1ms TTL=253
Reply from 192.168.2.253: bytes=32 time<1ms TTL=253
Reply from 192.168.2.253: bytes=32 time<1ms TTL=253

C:\>SSH -l jeremy 192.168.2.253

Password: 



SW2>en
Password: 
SW2#
```

**references**

1. [^https://www.youtube.com/watch?v=AvgYqI2qSD4&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=81]
2. [^https://community.cisco.com/t5/networking-knowledge-base/understanding-the-differences-between-the-cisco-password-secret/ta-p/3163238]