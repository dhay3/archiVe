# ISO

## Live USB

在介绍 Live USB 之前需要先了解一下 Bootable USB

当 USB 插入电脑后能被引导的 USB，就是 Bootable USB 。 需要 USB 的引导序号序号比电脑上的硬盘序号高才能被引导

而 Live USB 是 Bootable USB 的真子集，能被引导的同时，还提供可以直接使用的 OS

例如 

如果你使用 rufus 将 Windows 10 写入到 USB，那么这个 USB 就是一个 Bootable USB，只提供安装引导，但是不能直接使用 Windows

如果你写入的 ISO 是 ubuntu live ，那么这个 USb 就是一个 Live USB，可以提供安装以及直接使用 Ubuntu 的功能

## 常用的 ISO

记录一些使用的 iso

```
.
├── Linux
│   ├── alpine-standard-3.18.4-x86_64.iso
│   ├── archlinux-2021.03.01-x86_64.iso
│   ├── backbox-8-desktop-amd64.iso
│   ├── dfly-x86_64-6.4.0_REL.iso
│   ├── finnix-124.iso
│   ├── FreeBSD-13.2-RELEASE-amd64-dvd1.iso
│   ├── kali-linux-2023.1-installer-amd64.iso
│   ├── linuxmint-21.2-cinnamon-64bit.iso
│   ├── manjaro-kde-21.0.5-210519-linux510.iso
│   ├── manjaro-kde-23.0.4-231015-linux65.iso
│   ├── openwrt5.4-gdq-20230218-x86-64-generic-squashfs-combined.img
│   ├── ubuntu-22.04.2-live-server-amd64.iso
│   └── void-live-x86_64-musl-20210218-enlightenment.iso
├── Miscellaneous
│   ├── clonezilla-live-2.8.0-27-amd64.iso
│   ├── FD13LIVE.iso
│   ├── gparted-live-1.5.0-1-amd64.iso
│   └── proxmox-ve_7.3-1.iso
└── Windows
    ├── cn_windows_10_consumer_editions_version_1909_updated_jan_2020_x64_dvd_47161f17.iso
    ├── cn_windows_7_professional_with_sp1_x64_dvd_u_677031.iso
    ├── sc_winxp_pro_with_sp2.iso
    ├── WePE64_V2.2.iso
    └── Win11_22H2_EnglishInternational_x64v2.iso
```

**references**

[^1]:https://www.quora.com/What-difference-between-live-usb-and-bootable-usb-Or-Are-these-words-has-the-same-meaning
