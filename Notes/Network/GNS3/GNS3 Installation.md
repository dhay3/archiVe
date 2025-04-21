# GNS3 Instanllation 	

> ==注意为了正常运行GNS3，需要保证 GUI 和 VM 版本对应，否则导入IOS时可能会失效==

## GNS3 GUI

这里只记录Linux的安装方式

GNS GUI 默认支持 Debian 和 Ubuntu 的，如果是 Arch 的可以使用 ARU 安装

如果是通过 AUR 安装的，可能会有依赖冲突的。需要手动安装依赖，大部分可能是 python 的 lib，使用pip手动安装即可

## GNS3 Server

### local GNS3

GNS server 安装在本地宿主机上，不太推荐

### local GNS3 VM

GNS server 按照在本地虚拟机上

https://docs.gns3.com/docs/getting-started/installation/download-gns3-vm

下载镜像直接双击即可，如果出现如下错误

```
Failed to save the settings.

Empty or null host only interface name is not valid.

Result Code: NS_ERROR_FAILURE (0x80004005)
Component: NetworkAdapterWrap
Interface: INetworkAdapter {e9a0c183-7071-4894-93d6-dcbec010fa91}

```

主需要为宿主机添加一张虚拟的NIC adaptor 即可 (File -> Tools -> Network Manager)

https://github.com/GNS3/gns3-vm/issues/102

按照提供的信息在 GNS3 edit - preferences 中配置 server 信息

> 默认账号：gns3
>
> 默认密码：gns3

### remote GNS3 VM 

GNS3 server 安装在 remote host VM 上

> 出现[No module named 'apt_pkg' error](https://askubuntu.com/questions/1069087/modulenotfounderror-no-module-named-apt-pkg-error) 

这里我采用[remote GNS3 VM](https://docs.gns3.com/docs/getting-started/installation/remote-server/)方式，==这种方式project会存储在remote server opt目录以uuid为文件名==

执行如下方法，会在remote server上安装并运行openVPN，作为服务端

```
cd /tmp
curl https://raw.githubusercontent.com/GNS3/gns3-server/master/scripts/remote-install.sh > gns3-remote-install.sh
bash gns3-remote-install.sh --with-openvpn --with-iou --with-i386-repository
```

同时在会在remote server 的$HOME目录下生成客户端的openVPN配置文件，名为`client.opvn`

校验服务端服务是否开启

```
#openvpn默认使用udp1194作为CS之间的通信通道
root in ~ λ netstat -lnpu | grep 1194
udp        0      0 0.0.0.0:1194            0.0.0.0:*                           17386/openvpn  

#创建一个tun
16: tun1194: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 100
    link/none 
    inet 172.16.253.1 peer 172.16.253.2/32 scope global tun1194
```

在客户端运行`client.opvn`，并校验连接

```
#可以使用--daemon以守护进程的方式运行，日志将会被转移到syslog中
cpl in ~/Desktop λ sudo openvpn client.opvn
...
#GNS3默认创建172.16.253.1 3080
nc -nvz 172.16.253.1 3080
```

如果连接不上确保防火墙正常关闭，使用HTTP连接。

按照提供的信息在 edit - preferences 中配置 server 信息

## GNS3 wizard

https://docs.gns3.com/docs/getting-started/setup-wizard-local-server

通过 wizard 可以设置 GNS3 server 的安装方式

help -> setup wizardd