# 安装GRUB

参考：

https://www.gnu.org/software/grub/manual/grub/grub.html#Obtaining-and-Building-GRUB

安装GRUB做为boot loader，需要先安装【GRUB system utilities】，可以二进制安装或通过包管理器安装。

## grub-install

安装完成后通过如下命令安装，gurb-install实际上只是一个脚本。具体的操作通过grub-mkimage完成。

```
grub-install <install_device>
#等价，boot-directory表示GRUB 安装的镜像位置为/dir/grub
grub-install --boot-directory=<dir> <install_device>
```

### 参数

- --version

  查看GRUB的版本

  ```
  root in /var/www/html λ grub-install --version
  grub-install (GRUB) 2.02-2ubuntu8.20
  ```

- --boot-directory

  指定GRUB image安装的位置，如果需要【多个不同的引导分区】需要指定

  ```
  #将引导GRUB安装到/mnt/boot/grub
  grub-install --boot-directory=/mnt/boot /dev/sdb
  ```

## legacy BIOS VS EFI

https://www.gnu.org/software/grub/manual/grub/grub.html#BIOS-installation

https://www.gnu.org/software/grub/manual/grub/grub.html#Installing-GRUB-using-grub_002dinstall

- Legacy BIOS需要手动指定GRUB安装的硬盘

- EFI如果创建了一个EFI System partition并挂载到`/boot/efi`，使用`grub-install`不需要指定硬盘，默认直接将GRUB安装到`/boot/efi`。否则需要指定EFI System partition的挂载点

  ```
  grub-install
  #指定EFI system partition
  grub-install --efi-directory=/mnt/efi
  ```

  











