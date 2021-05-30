# Vargant centOS7安装

[TOC]

## Install

https://www.vagrantup.com/

文档

https://learn.hashicorp.com/collections/vagrant/getting-started

## Initialize

默认从https://app.vagrantup.com/下载, 国内需要换镜像

http://mirrors.ustc.edu.cn/centos-cloud/centos/6/vagrant/x86_64/images/

As Follow:

==需要在power shell 下运行==， centos7指的是vagrant box中的id

`vagrant init centos7 https://mirrors.ustc.edu.cn/centos-cloud/centos/7/vagrant/x86_64/images/CentOS-7.box`

在本地创建一个文件, 通过`vagrant init`创建一个box

<img src="..\..\imgs\_vagrant\Snipaste_2020-08-14_18-27-29.png" />

命令会自动创建一个Vagrantfile

```sql
Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/bionic64" //对应app.vagrantup.com上的镜像名
  config.vm.box_url = "https://vagrantcloud.com/hashicorp/bionic64"//镜像地址
  config.vm.box_version = "1.0.282"//镜像版本
end
```

相同的`vagrant box add <boxname> <url>`也可以添加(这里可以添加本地镜像), 但是需要手动修改配置

---

## Boot

通过`vagrant up`命令启动虚拟机

<img src="..\..\imgs\_vagrant\Snipaste_2020-08-14_01-56-21.png"/>

虚拟机启动后就会在virtual box中显示

==注意这里默认镜像安装在C盘, 这里需要修改路径==

<img src="..\..\imgs\_vagrant\Snipaste_2020-09-03_11-00-22.png"/>

## SSH

输入`vagrant ssh `进入centOS TTL

<img src="..\..\imgs\_vagrant\Snipaste_2020-08-19_12-47-54.png"/>

通过`logout` 或是 `exit`退出终端

### 切换用户

`su root`

账号密码默认都是`vagrant`

### Xshell登入vagrant

<img src="..\..\imgs\_vagrant\Snipaste_2020-08-19_15-45-05.png"/>

## 网络

参考:

https://www.jianshu.com/p/a1bc23bc7892

vagrant镜像默认使用NAT端口转发, 即使用同一IP, 端口需要转发, 

```
config.vm.network :forwarded_port, guest: 80, host: 8080
```

将虚拟机（被称作guest）的端口80映射为宿主机的端口8080。



为了方便开发需要修改Vagrantfile

```
config.vm.network :private_network, ip: "192.168.1.104"
```

你可以从宿主机自由访问虚拟机，但LAN网络中的其他人不需要也无法访问虚拟机。

值得注意的是，ip地址“192.168.1.104”不是随便指定的。
 首先你可以不指定，这表示虚机启动时会DHCP到一个可用的IP地址（例如：192.168.33.101），这是vagrant通过virtualbox私有网络的DHCP机制获得的。
 如果你要自行指定明确的IP地址，要保证该地址是在恰当的网段中，例如192.168.33.71。

多台虚拟机在私有网络模式下也可以互相访问，只要设置为相同的网段就可以。

本质上说，这是使用provider的HostOnly模式。

如果在windows终端中没有显示网卡地址

<img src="..\..\imgs\_vagrant\Snipaste_2020-08-19_13-51-32.png" style="zoom:50%;" />

## 常见指令

- vagrant box list

  列出所有的box

- vagrant box remove [box name]

  删除指定的box

- vagrant reload

  重新加载Vagrantfile配置

- vagrant ssh-config

  显示ssh连接信息



