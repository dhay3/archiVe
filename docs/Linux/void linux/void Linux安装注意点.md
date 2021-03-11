# void Linux安装注意点

https://docs.voidlinux.org/installation/live-images/partitions.html

https://docs.voidlinux.org/installation/live-images/partitions.html#bios-system-notes

> 在编辑虚拟机的高级中查看引导方式

1. keyword 需要选择us

2. Bios引导，第一个分区需要为type Bios boot，容量为1MB(写入Grub)，==不创建fs==

   efi引导，需要格式化引导第一个分区为vfat(fat32)，并挂载到`/boot/efi`

   ![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-09_15-47-20.png)

3. 如果安装vmtools，需要移动到其他目录解压，然后运行`./vmtools-install.pl`

   这里需要安装perl

   ```
   xbps-install -Su
   xbps-install -u xbps
   xbps-install perl
   ```

   