ref
[https://wiki.archlinux.org/title/KVM](https://wiki.archlinux.org/title/KVM)
[https://www.linux-kvm.org/page/FAQ#General_KVM_information](https://www.linux-kvm.org/page/FAQ#General_KVM_information)
## Digest
kernel-based virutal machine (KVM) 是一种针对 Linux 的开源的虚拟化技术。KVM 在 kernel 2.6.20 之后已经编译到内核中 
## Terms

- VT

  vitualization technology 虚拟化技术

- HVM

  hardware Virtual machine 通常表示 X86 架构的 VT 扩展

- Intel VT/ AMD-V

  CPU 支持 VT 的扩展

## FQA
### KVM VS Xen
Xen 是一个外部的hypervisor，而 KVM 是Linux 内部的 hypervisor。这也意味着 KVM 比 Xen 使用更便捷
### KVM VS VMware
vmware 是商业软件，而 KVM 是 GPL 开源软件
### KVM VS QEMU
QEMU 使用 emulation，而 KVM 使用 processor extensions（HVM）来虚拟化

