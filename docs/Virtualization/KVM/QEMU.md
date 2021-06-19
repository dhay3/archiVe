# QEMU

qemu和vmware，virtualbox一样是一款虚拟化模拟器，和他们不同的是qemu运行虚拟机必须通过命令行或是脚本。但是可以使用`libvirt`来管理qemu创建的虚拟机

## 安装

有两种选择，qemu(带有GUI版本的)，qemu-headless(无GUI版本)，两者不兼容。我们可以通过`pacman -S qemu`来安装，安装后的full-system emulation以`qemu-system-<target_architecture>`来命名，例如`qemu-system-x86_64`

## 虚拟机

运行qemu需要ISO

