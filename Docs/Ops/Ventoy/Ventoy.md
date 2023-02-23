# Ventoy

ref

https://www.ventoy.net/en/index.html

## Digest

在制作 Bootable USB drive 时，通常需要将 U 盘格式化后，然后用写盘工具写入镜像。如果不做其他操作默认会写入整个 USB，这样即浪费空间管理也不方便

Ventoy 就是一个用于解决这个问题的开源工具，除此外 ventoy 还支持客制化配置，可玩性极高

## Installation

https://www.ventoy.net/en/doc_start.html

## Openwrt

ventoy 默认不支持引导 Openwrt 的操作系统，如果需要引导参考

https://www.ventoy.net/en/doc_start.html

## injection

ventoy 和 微PE 一样可以将需要的文件注入到 live OS 中，但是 ventoy 是通过脚本的方式注入，具体可以参考

https://www.ventoy.net/cn/plugin_injection.html

如果是 Linux 的必须要按照如下方式注入。建议使用 home 目录 或者 不创建目录表示根目录，使用其他目录可能会失败

https://www.ventoy.net/cn/doc_live_injection.html

## DIY

https://www.gnome-look.org/p/1569525

在原主题上可以修改一些图片和值，来适配 ventoy 

![unnamed](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230222/unnamed.5xekla5me874.webp)

目录层级如下

```
.
├── injection
├── ISO
│   ├── Linux
│   └── Windows
├── System Volume Information
└── ventoy
    ├── live-injection-1.0
    │   └── live-injection-1.0
    │       ├── internal
    │       │   └── distro
    │       │       ├── arch
    │       │       ├── debian
    │       │       ├── manjaro
    │       │       └── rhel7
    │       └── sysroot
    │           └── home
    └── thems
        ├── reaper_1920x1080
        │   └── icons
        ├── reaper_2240x1400
        │   └── icons
        ├── redskull_1920x1080
        │   └── icons
        ├── redskull_2240x1400
        │   └── icons
        ├── wannacry_1920x1080
        │   └── icons
        └── wannacry_2240x1400
            └── icons
```

### themes.txt

```
# DedSec GRUB theme (1440p)

desktop-image: "background.png"
title-text: ""
terminal-font: "Hack Bold 22"
terminal-left: "20%"
terminal-top: "35%"
terminal-width: "60%"
terminal-height: "40%"
terminal-box: "menu_bkg_*.png"

+ boot_menu {
    menu_pixmap_style = "boot_menu_*.png"
    left = 20%
    width = 60%
    top = 30%
    height = 40%
    item_font = "Norwester Regular 28"
    item_color = "#919090"
    selected_item_font = "Norwester Regular 30"
    selected_item_color = "#ffffff"
    icon_width = 48
    icon_height = 48
    item_icon_space = 24
    item_height = 56
    item_padding = 8
    item_spacing = 16
    selected_item_pixmap_style = "select_*.png"
    scrollbar = true
    scrollbar_width = 10
    scrollbar_thumb = "slider_*.png"
}

+ progress_bar {
    id = "__timeout__"
    left = 25%
    width = 50%
    top = 75%
    height = 20
    text = ""
    text_color = "#ffffff"
    font = "Norwester Regular 24"
    bar_style = "progress_bar_*.png"
    highlight_style = "progress_highlight_*.png"
}

+ hbox{ 
    left = 30%
    top = 98%
    width = 10%
    height = 25
    + label {text = "@VTOY_HOTKEY_TIP@" color = "red" align = "left"} 
}

+ hbox{ 
    left = 90%
    top = 2 
    width = 10%
    height = 25
    + label {text = "@VTOY_MEM_DISK@" color = "red" align = "left"} 
}
```

### ventoy.json

```
{
    "control":
    [
        {
            "VTOY_MENU_LANGUAGE": "en_US"
        },
        {
            "VTOY_DEFAULT_MENU_MODE": "1"
        },
        {
            "VTOY_TREE_VIEW_MENU_STYLE": "0"
        },
        {
            "VTOY_FILT_DOT_UNDERSCORE_FILE": "1"
        },
        {
            "VTOY_SORT_CASE_SENSITIVE": "0"
        },
        {
            "VTOY_MAX_SEARCH_LEVEL": "max"
        },
        {
            "VTOY_DEFAULT_SEARCH_ROOT": "/ISO"
        },
        {
            "VTOY_DEFAULT_KBD_LAYOUT": "QWERTY_USA"
        },
        {
            "VTOY_WIN11_BYPASS_CHECK": "1"
        },
        {
            "VTOY_WIN11_BYPASS_NRO": "1"
        },
        {
            "VTOY_LINUX_REMOUNT": "0"
        },
        {
            "VTOY_SECONDARY_BOOT_MENU": "1"
        }
    ],
    "theme":
    {
        "default_file": "0",
        "file":
        [
            "/ventoy/themes/reaper_1920x1080/theme.txt",
            "/ventoy/themes/reaper_2240x1400/theme.txt",
            "/ventoy/themes/redskull_1920x1080/theme.txt",
            "/ventoy/themes/redskull_2240x1400/theme.txt",
            "/ventoy/themes/wannacry_1920x1080/theme.txt",
            "/ventoy/themes/wannacry_2240x1400/theme.txt"
        ],
        "resolution_fit": 1,
        "gfxmode": "max",
        "display_mode": "GUI",
        "ventoy_left": "2%",
        "ventoy_top": "98%",
        "ventoy_color": "red"
    },
    "menu_tip":
    {
        "left": "2%",
        "top": "96%",
        "color": "red",
        "tips":
        [
            {
                "dir": "/ISO/Windows",
                "tip": "This directory contains winxp/7/10 ISO"
            },
            {
                "dir": "/ISO/Linux",
                "tip": "This directory contains arch/manjaro/utuntu/kali/centos/finnix/void/openwrt ISO"
            }
        ]
    },
    "menu_alias":
    [
        {
            "image": "/ISO/Linux/archlinux-2021.03.01-x86_64.iso",
            "alias": "archlinux"
        },
        {
            "image": "/ISO/Linux/CentOS-7-x86_64-DVD-1908.iso",
            "alias": "centos7"
        },
        {
            "image": "/ISO/Linux/finnix-124.iso",
            "alias": "finnix"
        },
        {
            "image": "/ISO/Linux/kali-linux-2020.3-installer-amd64.iso",
            "alias": "kali202003"
        },
        {
            "image": "/ISO/Linux/manjaro-kde-21.0.5-210519-linux510.iso",
            "alias": "manjaro"
        },
        {
            "image": "/ISO/Linux/openwrt5.4-gdq-20230218-x86-64-generic-squashfs-combined.img",
            "alias": "openwrt5.4-gdq"
        },
        {
            "image": "/ISO/Linux/openwrt-22.03.2-x86-64-generic-squashfs-combined.img",
            "alias": "openwrt-official"
        },
        {
            "image": "/ISO/Linux/ubuntu-20.04.2.0-desktop-amd64.iso",
            "alias": "ubuntu"
        },
        {
            "image": "/ISO/Linux/void-live-x86_64-musl-20210218-enlightenment.iso",
            "alias": "void20210218"
        },
        {
            "image": "/ISO/Windows/sc_winxp_pro_with_sp2.iso",
            "alias": "winxp-pro"
        },
        {
            "image": "/ISO/Windows/cn_windows_10_consumer_editions_version_1909_updated_jan_2020_x64_dvd_47161f17.iso",
            "alias": "win10"
        },
        {
            "image": "/ISO/Windows/cn_windows_7_professional_with_sp1_x64_dvd_u_677031.iso",
            "alias": "win7"
        },
        {
            "image": "/ISO/WePE64_V2.2.iso",
            "alias": "WePE"
        }
    ],
    "menu_class":
    [
        {
            "key": "arch",
            "class": "arch"
        },
        {
            "key": "ubuntu",
            "class": "ubuntu"
        },
        {
            "key": "CentOS",
            "class": "centos"
        },
        {
            "key": "manjaro",
            "class": "manjaro"
        },
        {
            "key": "void",
            "class": "void"
        },
        {
            "key": "kali",
            "class": "kali"
        },
        {
            "key": "openwrt",
            "class": "openwrt"
        },
        {
            "key": "WePE",
            "class": "WePE"
        },
        {
            "key": "finnix",
            "class": "finnix"
        },
        {
            "key": "win",
            "class": "win"
        }
    ],
    "injection":
    [
        {
            "image": "/ISO/Linux/finnix-124.iso",
            "archive": "/injection/live_injection.tar.gz"
        },
        {
            "image": "/ISO/WePE64_V2.2.iso",
            "archive": "/injection/physdiskwrite-0.5.3.zip"
        }
    ]
}
```

