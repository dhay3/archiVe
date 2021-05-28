# KVM

参考：

https://www.linux-kvm.org/page/Main_Page

Kernel-based Virtual Machine(KVM)，是一种针对Linux_x86架构的虚拟化方案。可以在Linux上运行Linux 或者是 Windows镜像(类似的有vmware和virtualbox)。在2.6.20之后被加入到Linux，由QEMU作为GUI。

##  前置条件

1. CPU 支持虚拟化

2. 配置KVM模块

   查看KVM模块

   ```
   lsmod | grep kvm
   ```

   如果没有，使用如命令加载KVM模块，只会一次生效

   ```
   #intel cpu
   modprobe kvm
   modprobe kvm_intel
   #amd cpu
   modprobe kvm
   modprobe kvm_amd
   ```

   永久配置

   ```
   cp /etc/modules /etc/modules.bak;sudo tee /etc/modules << EOF
   kvm
   kvm_adm
   EOF
   ```

3. 配置网络

   为了使虚拟机和宿主机通信

   ```
   sudo /etc/netplan/xxx <<EOF
   # Let NetworkManager manage all devices on this system
   network:
     version: 2
     renderer: NetworkManager
     bridges:
       br0:
         addresses: [192.168.80.100/24]
         gateway4: 192.168.80.1
         mtu: 1500
         nameservers:
           addresses: [8.8.8.8]
         dhcp4: no
         dhcp6: no
   EOF;netplan apply /etc/netplan/xxx
   ```

4. 安装kvm管理组件

   ```
   sudo apt install qemu-kvm qemu-kvm-tools libvirt virt-manager virt-install
   ```

5. 启动libvirtd

   ```
   systemctl start libvirtd.service
   ```

   服务启动后，会多出两张NIC，分别为`virbr0-nic`和`virbr0`两张NIC，virbr0是一个NAT网桥，virbr0-nic就桥接到virbr0上

