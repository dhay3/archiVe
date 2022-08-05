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

这里推荐使用 Dockerfile 打包，这是使用 centos 做镜像源，打包一些常用的工具

```
FROM centos$latest
LABEL "author"="kikochz"
LABEL "description"="a image for networking lab"
RUN cd /etc/yum.repos.d/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN yum -y install traceroute
RUN yum -y install tcpdump
RUN yum -y install mtr
RUN yum -y install nc
RUN yum -y install curl
RUN yum -y iptables

CMD ["/bin/bash"]
```

AUR GNS3 可能有 bug，我一直不能在 GNS3 中加对同一个 image 打包的不同 image 
