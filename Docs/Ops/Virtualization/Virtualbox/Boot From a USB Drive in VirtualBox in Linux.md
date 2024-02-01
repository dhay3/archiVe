# Boot From a USB Drive in VirtualBox in Linux

想要在 Virtualbox 中引导 USB boot stick 需要按照下图设置

![](https://raw.githubusercontent.com/dhay3/image-repo/master/20240201/2023-01-31_16-00.2b0p0d3pjclc.webp)

同时需要开启 USB passthrough

![](https://raw.githubusercontent.com/dhay3/image-repo/master/20240201/2023-01-31_16-00.1mxf97pghwqo.webp)

启动虚拟机后还需在菜单栏中的 Devices 选中对应的 USB stick(选中后需要通常需要重启，因为虚拟机引导太快了)。然后快速按下 ECS，在 Boot Manager 中选中 USB stick 即可

![](https://raw.githubusercontent.com/dhay3/image-repo/master/20240201/2024-02-01_22-47.3vpuq6eoaq00.webp)



**references**

[^1]:https://forums.virtualbox.org/viewtopic.php?t=110705