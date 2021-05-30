# Metaspoiltable 3安装

git地址https://github.com/rapid7/metasploitable3

[TOC]

## 在线安装

### Requirements

- [Packer](https://www.packer.io/intro/getting-started/install.html) 
- [Vagrant](https://www.vagrantup.com/docs/installation/)
- [Vagrant Reload Plugin](https://github.com/aidanns/vagrant-reload#installation)
- [VirtualBox](https://www.virtualbox.org/wiki/Downloads), libvirt/qemu-kvm, or vmware (paid license required)
- Internet connection

packer如果通过choco安装,在代理不稳定的情况极易安装失败。这里建议下载[binary](https://www.packer.io/downloads.html)

==除此之外还需要配置环境变量==, 当在控制台中出现下图所示, 即表示安装成功

<img src="..\..\..\imgs\_Kali\metasploit\metasploit\Snipaste_2020-09-03_12-07-01.png" style="zoom:80%;" />

然后创建一个git仓库`git clone https://github.com/rapid7/metasploitable3.git`

### build

在当前仓库下运行`.\build.ps1 windows2008`, 具体查看github, 如果出现如下错误信息

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-03_12-27-01.png"/>

参考:

https://www.fujieace.com/computer-practical/windows-powershell-ps1.html

以管理员身份运行powershell`Get-ExecutionPolicy -List`

<img src="..\..\..\imgs\_Kali\metasploit\964.png"/>

`set-executionpolicy remotesigned`

<img src="..\..\..\imgs\_Kali\metasploit\965.png"/>

然后执行脚本就没有问题了。操作完成后，为了安全然后删除执行策略`set-executionpolicy Undefined`

> 由于packer的版本迭代非常快 ,在执行脚本之前还需要修改一下镜像Json文件
>
> 参考：https://www.packer.io/docs/commands/fix

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-03_13-24-40.png"/>

进入dos窗口`packer fix windows_2008_r2.json`, 然后运行脚本。还有一点由于使用的时官网镜像，下载会非常慢。需要替换成本地地址

<img src="..\..\imgs\_Kali\Snipaste_2020-09-03_13-29-37.png" style="zoom:80%;" />

同时还需要替换build.sh中的一些内容

<img src="D:\asset\imgs\_Kali\Snipaste_2020-09-03_13-35-38.png" style="zoom:80%;" />

## 本地安装

> ==由于步骤繁琐，且有墙大的GFW。还是使用本地镜像的方式，安装metasploitable3==

参考：

https://www.freebuf.com/sectool/122626.html

https://jeza-chen.com/2018/09/21/MetaSploit3_Setup/

由于使用本地镜像，那也不用git clone

1. 安装Vagrant Reload Plugin

```shell
vagrant plugin install vagrant-reload
```

2. 添加metasploitable3到box仓库

```shell
vagrant box add rapid7/metasploitable3-win2k8 metasploitable3-win2k8.box
```

3. 创建vagrant file

```shell
vagrant init rapid7/metasploitable3-win2k8
```

4. 启动vm

```shell
vagrant up
```

进入下图界面即安装成功

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-06_16-14-50.png"/>

vagrant 默认密码: `vagrant`

administrator默认密码:
