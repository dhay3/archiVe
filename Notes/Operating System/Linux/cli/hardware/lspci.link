# lspci

参考：

https://diego.assencio.com/?index=649b7a71b35fc7ad41e03b6d0e825f07

https://www.cnblogs.com/machangwei-8/p/10403495.html

lspci用于显示所有使用PCI接口的设备，例如：

```
cpl in ~ λ lspci 
00:00.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir Root Complex
00:00.2 IOMMU: Advanced Micro Devices, Inc. [AMD] Renoir IOMMU
00:01.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe Dummy Host Bridge
00:02.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe Dummy Host Bridge
00:02.2 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe GPP Bridge
00:02.4 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe GPP Bridge
00:08.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe Dummy Host Bridge
00:08.1 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir Internal PCIe GPP Bridge to Bus
00:14.0 SMBus: Advanced Micro Devices, Inc. [AMD] FCH SMBus Controller (rev 51)
01:00.0 Network controller: Intel Corporation Wi-Fi 6 AX200 (rev 1a)
02:00.0 Non-Volatile memory controller: Sandisk Corp WD Black 2019/PC SN750 NVMe SSD
03:00.0 VGA compatible controller: Advanced Micro Devices, Inc. [AMD/ATI] Device 1638 (rev c5)
03:00.1 Audio device: Advanced Micro Devices, Inc. [AMD/ATI] Device 1637
03:00.2 Encryption controller: Advanced Micro Devices, Inc. [AMD] Family 17h (Models 10h-1fh) Platform Security Processor
03:00.3 USB controller: Advanced Micro Devices, Inc. [AMD] Renoir USB 3.1
```

单独拿出一列来看：

```
00:00.0 Host bridge: Intel Corporation Haswell-ULT DRAM Controller (rev 0b)
```

第一列的第一个00表示bus number，第二个00表示device number 00，第三个0表示function number。设备种类是Host bridge（也被叫做north bridge），设备商(vendor)是intel，设备名是Haswell-ULT DRAM Controller，版本为11(0b十进制)

可以看出电脑上有4个PCI bus(0,1,2,3)。在单个系统上插入多个总线是通过bridge来完成的，==bridge是一种用来连接bus的特殊的PCI外设==

可以使用`-v`查看所有设备详细信息，可以使用`-s`来指定具体的pci设备

```
cpl in ~ λ lspci -v -s 03:00.1
03:00.1 Audio device: Advanced Micro Devices, Inc. [AMD/ATI] Device 1637
        Subsystem: Lenovo Device 3817
        Flags: bus master, fast devsel, latency 0, IRQ 90, IOMMU group 4
        Memory at fd3c8000 (32-bit, non-prefetchable) [size=16K]
        Capabilities: <access denied>
        Kernel driver in use: snd_hda_intel
        Kernel modules: snd_hda_intel
```

可以使用`-mmv`输出机器可以读懂的内容

```
cpl in ~ λ lspci -mmvs 03:00.1
Slot:   03:00.1
Class:  Audio device
Vendor: Advanced Micro Devices, Inc. [AMD/ATI]
Device: Device 1637
SVendor:        Lenovo
SDevice:        Device 3817
IOMMUGroup:     4
```