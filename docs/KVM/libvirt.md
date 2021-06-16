# libvirt

参考：

https://wiki.archlinux.org/title/Libvirt

libvirt针对providers/hypervisors提供简单的管理方式，例如kvm/qemu，virtualbox hybervisors。libvirt有如下几个特性：

- **VM management**：Various domain lifecycle operations such as start, stop, pause, save, restore, and migrate. Hotplug operations for many device types including disk and network interfaces, memory, and CPUs.

- **Remote machine support**: All libvirt functionality is accessible on any machine running the libvirt daemon, including remote machines. A variety of network transports are supported for connecting remotely, with the simplest being SSH, which requires no extra explicit configuration.
- **Storage management**: Any host running the libvirt daemon can be used to manage various types of storage: create file images of various formats (qcow2, vmdk, raw, ...), mount NFS shares, enumerate existing LVM volume groups, create new LVM volume groups and logical volumes, partition raw disk devices, mount iSCSI shares, and much more.
- **Network interface management**: Any host running the libvirt daemon can be used to manage physical and logical network interfaces. Enumerate existing interfaces, as well as configure (and create) interfaces, bridges, vlans, and bond devices.
- **Virtual NAT and Route based networking**: Any host running the libvirt daemon can manage and create virtual networks. Libvirt virtual networks use firewall rules to act as a router, providing VMs transparent access to the host machines network.

## 安装

libvirt使用server/client模式，server只需要安装在宿主机上。`pacman -S libvirt`安装server，client是提供给用户的interface用来管理虚拟机，可以使用的client参考https://wiki.archlinux.org/title/Libvirt#Client

==通常会使用virsh(CLI)，virtual manager(GUI)==

