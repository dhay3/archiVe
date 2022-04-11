# Type 1 hypervisor VS Type 2 hypervisor

参考：

https://searchservervirtualization.techtarget.com/feature/Whats-the-difference-between-Type-1-and-Type-2-hypervisors

https://medium.com/teamresellerclub/type-1-and-type-2-hypervisors-what-makes-them-different-6a1755d6ae2c

## hypervisor

virtualization需要通过hypervisor来管理virtual machine，同时允许hardware对virtual machine提供共享。分别有两种hypervisor，type 1 和 type2

## Type 1 hypervisors 

直接运行在host machine的物理硬件上(bare metal)==不需要依赖OS，所以也不存在OS和hypervisor的兼容性==，也被称为bare metal hypervisor。使用这种方式运行VM，相对安全。因为常见的Type 2 hypervisors方案host machine OS可能存在漏洞。扩展性比Type 2 hypervisors强

常用的软件有**proxmox**，**KVM**，**Vmware EXSI**

> KVM可以让Linux kernel变成bare metal hypervisor，但是使用的OS所以OS也是type 2 hypervisor

## Type 2 hypervisors

Type 2 hypervisors是基于host machine’s preexisting OS的，所有VM的活动都必须通过hos machine

常见的软件有**Virtualbox**，**Vmware Station**，**QEMU**

![2021-06-19_13-13](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/2021-06-19_13-13.4527enria540.png)