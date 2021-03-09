# void Linux安装注意点

https://docs.voidlinux.org/installation/live-images/partitions.html

> keyword 需要选择us

对于EFI驱动的系统，必须要有100MB fat32 挂载点在`/boot/efi`的主引导分区。根分区

，交换分区参考手册

![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-08_14-49-02.png)

最后分区配置成如下形式

![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-08_14-59-49.png)

整体配置

![](D:\asset\note\imgs\_Linux\Snipaste_2021-03-08_15-02-25.png)