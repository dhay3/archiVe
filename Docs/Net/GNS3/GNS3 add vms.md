# GNS3 add linux vms

> 如果 GNS3 server 是通过 VM 的方式安装的，在使用 qemu, virtualbox, vmware, docker 的方式同样需要在 GNS3 server 上安装对应的软件(VM上默认安装，但是需要导入对应的 ISO 或者pull container)

GNS3 有几种方式来添加主机

## Qemu VMs

通过 Qemu（kvm）来添加 vms

## VirtualBox VMs

通过 vbox 添加 vms

## VMware VMs

通过 vmware workstation 添加 vms

## Docker containers

通过 docker 来添加

https://docs.gns3.com/docs/emulators/create-a-docker-container-for-gns3/

按照文档操作就行，可以不需要先在 VM 上 pull container

可以直接指定 image，会自动 pull latest image。

需要这注意的是，通过这种方式启动的容器的网络环境是 none 模式。

这里推荐使用 Dockerfile 打包，这是使用 centos 做镜像源，打包一些常用的工具(可以按需选择，==但是需要注意的一点是如果容器需要使用`ipvsadm`，需要在宿主机（即 GNS-VM）上也要安装`ipvsadm`==)

```
FROM centos$latest
LABEL "author"="kikochz"
LABEL "description"="A image for networking lab"
RUN cd /etc/yum.repos.d/
RUN sed -i 's/mirrorlist=/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN yum -y install traceroute 
RUN yum -y install tcpdump 
RUN yum -y install mtr
RUN yum -y install nc
RUN yum -y install curl
RUN yum -y install nmap
RUN yum -y install telnet
RUN yum -y install iptables
RUN yum -y install ipvsadm
RUN yum -y install nginx
#RUN    yum -y install openssh-server && yum install openssh
#RUN ssh-keygen -A &&  /usr/sbin/sshd

CMD ["/bin/bash"]
```

AUR GNS3 可能有 bug，添加的 VMS 可能不会在面板中显示只需要重启以下 GNS3 就可以了

配置完成之后可以通过 termius telnet 来连接容器，但是需要注意的是 termius telnet 回车是`\n\r`所以在 UNIX-like OS 上每次执行完命令的时候会多输出一行 （因为不能识别`\r`），可能会导致一些命令出现异常，例如 `yum install` 如果不指定 `-y` 就会出现 `Is this ok [y/N]: Operation aborted.` 的错误

## Permanent networking conf

ref

https://github.com/muhamadfaiz/Create-Sub-Interface-in-Debian

GNS3 支持对 iface 配置做永久，可以右击实例，然后在 network configuration 中按照 debian 的格式配置

```
auto eth0
iface eth0 inet static
address 192.168.4.1
netmask 255.255.255.0
gateway 192.168.4.2
# sub interface
auto eth0:0
iface eth0:0 inet static
address 192.168.4.3
netmask 255.255.255.0
gateway 192.168.4.2
# new NIC
auto eth1 
iface eth1 inet static
address 192.168.5.1
netmask 255.255.255.0
gateway 192.168.4.2
```

