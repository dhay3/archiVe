# Virtualization Types

ref
[https://stackoverflow.com/questions/21462581/what-is-the-difference-between-full-para-and-hardware-assisted-virtualization](https://stackoverflow.com/questions/21462581/what-is-the-difference-between-full-para-and-hardware-assisted-virtualization)

> 以下都是针对 guest OS 而言的

## Paravirtualization
paravirtualization is virtualization in which the guest operating system(this  one being virtualized) is aware that is a guest and accordingly has drivers that, instead of issuing hardware commands, simply issue commands directly to the host operating system. This also includes memory and thread management as well, which usually require unavailable privileged instructions in the processor
可以理解成半虚拟，guest OS 执行的命令实际是在 host OS 上执行的
## Full Virtualization
completely simulates the underlying hardware
is virtualization in which the guest operating system is unaware that is is in a virtualized enviroment, and therefore harware is virtualized by the host operating system so that the guest can issue commands to what it thinks is actual hardware, but really are just simulated hardware devices created by the host
顾名思义 全虚拟，guest OS 不会认为自己是虚拟的，但是实际是全虚拟的包括 hardware

