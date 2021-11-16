# mhwd

reference:

https://linuxhint.com/manjaro_hardware_detection_tool/

## 概述

mhwd是一款可以修改和检测manjaro OS的工具

## mhwd

> 可以使用`mhwd -h`来查看帮助命令

### show

- `-li | --listhardware`

  查看连接OS的所有设备（通过PCI或USB连接的设备）

  ```
  cpl in ~ λ mhwd -lh
  > PCI devices:
  --------------------------------------------------------------------------------
                            TYPE            BUS   CLASS  VENDOR  DEVICE   CONFIGS
  --------------------------------------------------------------------------------
              Display controller   0000:03:00.0    0300    1002    1638         3
                          Bridge   0000:00:08.0    0600    1022    1632         0
                          Bridge   0000:00:18.3    0600    1022    166d         0
                          Bridge   0000:00:18.1    0600    1022    166b         0
  
  
  > USB devices:
  --------------------------------------------------------------------------------
                            TYPE            BUS   CLASS  VENDOR  DEVICE   CONFIGS
  --------------------------------------------------------------------------------
                        Keyboard      3-1.3:1.0   10800    04d9    0230         0
                             Hub        3-0:1.0   10a00    1d6b    0002         0
                           Mouse      3-1.2:1.1   10503    046d    c539         0
             Unclassified device        1-3:1.4    0000    13d3    5419         0
                Bluetooth Device        3-4:1.0   11500    8087    0029         0
  ```

  

  可以结合`-d`参数查看具体的信息，也可以使用`--usb`或`--pci`只查看连接的usb设备或pci设备

- `-la | --listall`

  列出所有设备的驱动信息

  ```
  cpl in ~ λ mhwd -la
  > All PCI configs:
  --------------------------------------------------------------------------------
                    NAME               VERSION          FREEDRIVER           TYPE
  --------------------------------------------------------------------------------
       network-rt3562sta            2013.12.07                true            PCI
     network-broadcom-wl            2018.10.07                true            PCI
         network-slmodem            2013.12.07                true            PCI
           network-r8168            2016.04.20                true            PCI
              video-vesa            2017.03.12                true            PCI
            video-sisusb            2020.01.18                true            PCI
            video-voodoo            2017.03.12                true            PCI
        video-openchrome            2020.01.18                true            PCI
  ```

### install

- `-a | --auto <usb|pci> <classid>`

  为某种设备自动安装驱动

  例如：`mhwd -a pci free/nonfree 003`

  表示为classid 为003的设备（mhwd -lh）安装driver

  ```
  cpl in ~ λ mhwd -a pci free 0300
  > Skipping already installed config 'video-linux' for device: 0000:03:00.0 (0300:1002:1638) Display controller ATI Technologies Inc Cezanne
  ```

- `-i | --install <usb|pci> <config>`

  安装指定驱动的配置，可以结合`-f`来强制重新安装驱动

  ```
  sudo mhwd -f -i pci video-nvidia
  ```

## mhwd-kernel

用于管理系统的kernel

- `-li | --listinstalled`

  列出系统已经安装的kernel

  ```
  cpl in ~ λ mhwd-kernel -li
  Currently running: 5.14.10-1-MANJARO (linux514)
  The following kernels are installed in your system:
     * linux513
     * linux514
  ```

- `-i | --install`

  安装kernel，可以在末尾添加rmc（remove current kernel）在安装中删除当前使用的kernel

  ```
  cpl in ~ λ mhwd-kernel -i <kernel_name> rmc
  ```

  

