# arch linux 安装注意点

https://wiki.archlinux.org/index.php/Installation_guide

> 如果出现grub问题重新`pacstrap /mnt base linux linux-firmware`

## bios引导

https://zhuanlan.zhihu.com/p/112541071

1. 无需为GRUB创建额外的1MB分区
2. 挂载点是`/mnt`
3. `pacman -S vim`安装vim，`pacman -S grub`安装grub boot loader
4. grub配置文件

## uefi引导

1. 需要对efi system格式化`mkfs.fat -F32 /dev/sda1`

2. 安装`grub`和`efibootmgr`

3. `grub-install --target=x86_64-efi --efi-directory=esp --bootloader-id=GRUB`

   这里的esp添加挂载点，而不是分区

## lvm

1. `pacstrap /mnt lvm2`

2. 修改

   ```
   /etc/mkinitcpio.conf
   HOOKS=(base udev ... block lvm2 filesystems)
   ```

   https://wiki.archlinux.org/index.php/Install_Arch_Linux_on_LVM#Adding_mkinitcpio_hooks

3. 修改

   ```
   /etc/default/grub
   GRUB_PRELOAD_MODULES="... lvm"
   ```

   https://wiki.archlinux.org/index.php/GRUB

4. ==如果根分区使用了lvm==，需要添加额外的参数

   修改

   ```
   GRUB_CMDLINE_LINUX_DEFAULT="root=/dev/vg01/lv01"
   ```

   https://wiki.archlinux.org/index.php/Kernel_parameters

