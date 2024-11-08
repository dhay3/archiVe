# Linux 重置root密码

参考：

https://my.oschina.net/u/3731306/blog/3004813

https://blog.51cto.com/asd9577/1931442

> 使用systemd无需进入rescue.target

1. 在启动界面键入`e`，修改grub

<img src="..\..\..\imgs\_VirtualMachine\_Linux\Snipaste_2020-10-31_09-40-12.png"/>

2. 将`ro` 修改为`rw init=/sysroot/bin/sh`

<img src="..\..\..\imgs\_VirtualMachine\_Linux\Snipaste_2020-10-31_09-43-07.png"/>

3. 运行`chroot /sysroot/ `命令来进入我们真正的系统，修改root密码

<img src="..\..\..\imgs\_VirtualMachine\_Linux\Snipaste_2020-10-31_09-47-11.png"/>

4. ==`touch /.autorelable`==，退出shell，重启。

<img src="..\..\..\imgs\_VirtualMachine\_Linux\Snipaste_2020-10-31_09-52-30.png"/>

   
