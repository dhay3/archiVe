# Manjaro/Windows 升级后 Grub menu 消失 



当升级 Mnajaro/Windows 系统后可能会出现

直接引导启动 Windows 或者 Manjaro，不显示 Grub menu 的问题，可以按照如下方式进行修复

1. 先检查 Bios Dirver 的引导顺序(Windows 升级常常会把 windows boot mananger 置位首位)，将 Grub 所在的磁盘优先级调高。通常第一步操作完成后就会正常，但是也有特殊情况

2. 如果不能正常引导 Grub menu，需要准备一个 live USB(推荐用 ventoy 做一个多启动盘)

3. 进入 live USB，执行如下命令

   ```
   #自动挂载 chroot 环境
   root # manjaro-chroot -a
   #重新安装 grub(这里是 EFI 环境下的)
   root # grub-install --target=x86_64-efi --efi-directory=/boot/efi --bootloader-id=manjaro --recheck
   #更新 grub
   root # grub-update
   root # exit
   root # reboot
   ```

   执行完后正常会显示 Grub menu，但是特殊情况下可能会导致 dual/triple boot 下其他系统不能被正常识别

4. 如果不能正常识别，重启正常进系统修改 `/etc/default/grub`

   ```
   GRUB_TIMEOUT_STYLE=menu
   ```

   然后执行 `update-grub` 更新 grub，即可

**references**

1. https://wiki.manjaro.org/index.php/GRUB/Restore_the_GRUB_Bootloader#Reinstall_GRUB
2. https://forum.manjaro.org/t/restoring-grub-after-installing-windows/110194
3. https://forum.manjaro.org/t/grub-menu-not-showing-on-boot-boots-into-default-kernel-instead/13410/2