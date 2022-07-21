# KVM “安装”

ref
[https://wiki.archlinux.org/title/KVM](https://wiki.archlinux.org/title/KVM)

## 前置条件
### 硬件

> 如果不支持可以看一下BIOS是不是设置了禁止虚拟化，现在一般的CPU的支持VT

A VT capable Intel processor, or an SVM capable AMD processor

可以通过如下方式校验

```
cpl in ~ λ lscpu| grep -i virtual
Address sizes:                   48 bits physical, 48 bits virtual
Virtualization:                  AMD-V
```

如果是intel CPU virtualization 的值必须是 VT-x，如果是amd CPU virtualization 的值必须是AMD-V。如果显示的是 full 或者 para 表示当前的机器已经是虚拟机了



或者通过如下命令查看

```
grep -E --color=auto 'vmx|svm|0xc0f' /proc/cpuinfo
```

如果没有任何输出表示不支持 hardware  virtualization

### 内核

需要有 KVM 的 module

```
cpl in ~ λ zgrep VIRTIO /proc/config.gz
CONFIG_BLK_MQ_VIRTIO=y
CONFIG_VIRTIO_VSOCKETS=m
CONFIG_VIRTIO_VSOCKETS_COMMON=m
CONFIG_BT_VIRTIO=m
...
```

只有值是`y`或者`m`的时候表示模块可用。然后使用下面命令查看模块是否已经载入

```
lsmod | grep kvm
```

如果显示为空，表示模块没有载入，需要使用modprobe 手动载入模块(只会生效一次)

```
modprobe kvm;modprobe kvm_intel
modprobe kvm;modprobe kvm_amd
```

如果 modprobe 不能载入 `kvm_intel` 或者 `kvm_amd` 但是可以载入 `kvm` 同时 `lscpu` 显示支持 VT，大概率可能是 BIOS 没有允许开启。可以通过`dmesg`来导致错误的原因

```
dmesg | grep kvm
```

