# Kali配置

参考:

Kali-Linux-Revealed-1st-edition

https://www.kali.org/docs/

## Kali 2020.1 root用户解决

参考:https://blog.csdn.net/qq_44612786/article/details/104220727?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~all~first_rank_v2~rank_v25-2-104220727.nonecase&utm_term=kali%E6%80%8E%E4%B9%88%E8%BF%9B%E5%85%A5root%E7%94%A8%E6%88%B7

> 2020.1版默认没有使用用root用户

```shell
sudo passswd root
```

## windows 10 卧底模式

```shell
kali-undercover #切换为win10卧底模式, 重新在终端输入即可还原
```

## kali更换apt源

https://developer.aliyun.com/mirror/

将原`etc/apt/sources.list`先备份一份, 然后将其修改为aliyun镜像

```shell
#deb https://mirrors.aliyun.com/kali kali-rolling main non-free contrib
#deb-src https://mirrors.aliyun.com/kali kali-rolling main non-free contrib
```

<img src="..\..\imgs\_Kali\Snipaste_2020-08-31_23-28-11.png"/>

然后更新apt源

```shell
apt update
```

## kali配置静态IP

`cd /etc/network`备份一份interfaces, 然后修改该文件

<img src="..\..\imgs\_Kali\Snipaste_2020-08-31_23-52-35.png"/>

Kali默认自动配置网络, `lo inet loopback`将locahost接口映射到loopback即127.0.0.1，inet表示IPv4。这里也可以使用dhcp。

在这里我使用vmnet8 NAT模式做为交换机, 具体参数需要自行更换，注意这里配置的网关的host-id不能为1

```mysql
auto eth0
iface eth0 inet static # IP类型与接口
address 192.168.80.3 # IP地址
netmask 255.255.255.0 # 子网掩码
broadcast 192.168.80.255 # 广播地址
network 192.168.80.0 # 网段
gateway 192.168.80.2 # 网关
```

`/etc/resolv.conf`修改DNS解析器，这里推荐使用

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_00-19-57.png" style="zoom:80%;" />

==重启网络==`systemctl restart networking`

## kali配置ssh

kali默认关闭ssh远程服务, 需要手动开启。这里同样需要备份一份`ssh_config`

开启公钥认证和允许root用户登入

<img src="..\..\imgs\_Kali\Snipaste_2020-09-01_01-17-05.png"/>

启动ssh服务并查看是否启动成功和端口, 也可以通过`netstat -lnpt |grep ssh`显示当前处于listening的传输层是tcp协议的端口, 来查看ssh的端口

```shell
systemctl start ssh && systemctl status ssh
```

开启ssh服务并设为开机自动启动

```shell
systetmctl start ssh && /lib/systemd/systemd-sysv-install enable ssh

```

## Postgeresql

```shell
systemctl start postgresql && systemctl enable postgresql
```

