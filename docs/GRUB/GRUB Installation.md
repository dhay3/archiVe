
参考：

[https://www.gnu.org/software/grub/manual/grub/grub.html#Obtaining-and-Building-GRUB](https://www.gnu.org/software/grub/manual/grub/grub.html#Obtaining-and-Building-GRUB)
[https://www.gnu.org/software/grub/manual/grub/grub.html#Installation](https://www.gnu.org/software/grub/manual/grub/grub.html#Installation)
## Installation

1. 安装GRUB做为boot loader，需要先安装【GRUB system utilities】，可以在Unix下通过二进制安装或通过包管理器安装。
```bash
vagrant@saltmaster:~$ apt install grub2
```
该安装包下包含常用的`update-grub`、`grub-mkconfig`等工具

2. 指定安装 boot loader 的盘符，无需附加盘符
```bash
# grub-install /dev/sda
```
GRUB默认会将image放在`/boot`目录下，如果需要指定目录可以使用`--boot-directory`
如果是 EFI 系统，必须挂在 EFI system partition，默认挂载在`/boot/efi`，可以通过`--boot-direcotory`来指定挂载的目录
```bash
# grub-install --efi-directory=/mnt/efi
```
这里是单磁盘，所以没有指定install_device
## grub-install

安装完成后通过如下命令安装，gurb-install实际上只是一个脚本。具体的操作通过grub-mkimage完成。

```
grub-install <install_device>
#等价，boot-directory表示GRUB 安装的镜像位置为/dir/grub
grub-install --boot-directory=<dir> <install_device>
```

### 参数

-  --version
查看GRUB的版本  
```
root in /var/www/html λ grub-install --version
grub-install (GRUB) 2.02-2ubuntu8.20
```

-  --boot-directory
指定GRUB image安装的位置，如果需要【多个不同的引导分区】需要指定  
```
#将引导GRUB安装到/mnt/boot/grub
grub-install --boot-directory=/mnt/boot /dev/sdb
```

## legacy BIOS VS EFI

[https://www.gnu.org/software/grub/manual/grub/grub.html#BIOS-installation](https://www.gnu.org/software/grub/manual/grub/grub.html#BIOS-installation)

[https://www.gnu.org/software/grub/manual/grub/grub.html#Installing-GRUB-using-grub_002dinstall](https://www.gnu.org/software/grub/manual/grub/grub.html#Installing-GRUB-using-grub_002dinstall)

-  Legacy BIOS需要手动指定GRUB安装的硬盘 
-  EFI如果创建了一个EFI System partition并挂载到`/boot/efi`，使用`grub-install`不需要指定硬盘，默认直接将GRUB安装到`/boot/efi`。否则需要指定EFI System partition的挂载点  
```
grub-install
#指定EFI system partition
grub-install --efi-directory=/mnt/efi
```
