# Post Manjaro Installtion Guide 

## Time Sync

```
timedatectl set-ntp true
```

## Set Best Metric Mirros

> 千万不要把 Arch 的源加到 pacman.conf

set the fasttest mirrolist force TLS

```
sudo pacman-mirrors --proto https -a -t 3 -f
```

## Enable AUR/flatpak/snap

```
pacman -Sy yay flatpak snapd
```

## Passwall

~~Rip Clash verge & Clash for Windows~~

> Steps to have working tun mode with appimage format.
> cd "AppImage located folder"
> ./clash-verge_*.AppImage ---appimage-extract.
> cd squashfs-root
>
> Under apprun-hooks folder, delete the line with APPDIR= in linuxdepoly-plugin-gtk.sh
>
> sudo ./AppRun and give permissions to Clash/Meta under core settings.
> ./AppRun enjoy working tun mode, you can now set your desktop exec location to this file.

```
./clash-verge_1.3.8_amd64.AppImage --appimage-extract
cd squashfs-root/apprun-hooks
sed -i /"export APPDIR="/s/^/#/ linuxdeploy-plugin-gtk.sh
#需要先退出 clash verge
sudo ../AppRun
```

然后解锁 Clash core, 后续使用运行 AppRun

当然为了方便使用可以创建一个 desktop file

```
[Desktop Entry]
Exec=/home/0x00/appimages/squashfs-root/AppRun
MimeType=text/plain;
Name=Clash Verge
Type=Application
Icon=/home/0x00/appimages/squashfs-root/clash-verge.png
Terminal=false
Categories=Network
```

## Update System

```
sudo pacman -Syyu
```

## Install Drivers

如果安装驱动后不能正常引导 GUI 参考 [No GUI Desktop Starting Up](https://forum.manjaro.org/t/no-gui-desktop-starting-up/65689)

## Disable Beeping

```
cat << EOF | sudo tee /etc/modprobe.d/nobeep.conf
blacklist pcspkr
blacklist snd_pcsp
EOF
```

## Enable SSD TRIM

```
sudo systemctl start fstrim.timer && sudo systemctl enable fstrim.timer
```

## Kernel Tunnable  

```
cat << EOF | sudo tee /etc/sysctl.d/sysctl-custom.conf
vm.swappiness=0
net.core.default_qdisc=cake
net.core.somaxconn = 8192
net.core.netdev_max_backlog = 10240
net.core.rmem_default=262144
net.core.wmem_default=262144
net.core.rmem_max=67108864
net.core.wmem_max=67108864
net.ipv4.conf.all.rp_filter=1
net.ipv4.ip_forward=1
net.ipv4.ip_default_ttl=64
net.ipv4.ip_local_port_range=30000 60000
net.ipv4.icmp_echo_ignore_all=1
net.ipv4.tcp_max_syn_backlog = 10240
net.ipv4.tcp_max_tw_buckets=5000
net.ipv4.tcp_rmem=4096 102400 67108864
net.ipv4.tcp_wmem=4096 102400 67108864
net.ipv4.tcp_congestion_control=bbr
net.ipv4.tcp_syncookies=1
net.ipv4.tcp_syn_retries=3
net.ipv4.tcp_synack_retries=3
net.ipv4.tcp_fin_timeout=30
net.ipv4.tcp_keepalive_time=3600
net.ipv4.tcp_keepalive_intvl=60
net.ipv4.tcp_keepalive_probes=10
EOF
#shinrk Swappiness
cat /proc/sys/vm/swappiness
echo "vm.swappiness=40" | sudo tee -a sysctl-custom.conf
sysctl -p /etc/sysctl.d/network-custom.conf
```

## Enable Firewall

```
sudo pacman -S ufw
#GUI client for ufw
sudo pacman -S gufw
sudo systemctl start ufw && sudo systemctl enable ufw
```

## Enhance Sudo

add this line into `/etc/sudoers`

```
#below spec will override
@includedir /etc/sudoers.d

#boolean options
Defaults        !always_set_home 
Defaults        compress_io 
Defaults        fast_glob 
Defaults        mail_no_host
Defaults        pwfeedback
Defaults        targetpw
Defaults        sudoedit_follow
Defaults		!env_reset

#integer options
Defaults        passwd_timeout=1
Defaults        passwd_tries=2
Defaults        timestamp_timeout=0

#string options
Defaults        passprompt="password for [%h@%p]: "
Defaults        editor="/usr/bin/vim"

root ALL=(ALL:ALL) ALL
%wheel ALL=(ALL:ALL) NOPASSWD: ALL
```

## Terminal Complation

[]()

## Sogoupinyin

> use `fcitx-diagnose` for troubleshooting 

```
yay -S fcitx-sogoupinyin kcm-fcitx fcitx-configtool 
```



## Remove Orphaned Packages

## KDE Tweak

### shortcuts

### Desktop

latte-dock



### latte-dock



## Custom Grub

## Import SSH/GPG Key

## Backup

## Appendix

### Basic GUI Software

| Name                   | Description                                  |
| ---------------------- | -------------------------------------------- |
| virtualbox             | 虚拟机前端工具                               |
| google-chrome          | 浏览器                                       |
| timeshift              | 备份工具                                     |
| typora                 | markdown 编辑器                              |
| sublime-text-4         | 文本编辑器                                   |
| telegram-desktop       | 聊天工具                                     |
| gparted                | gparted GUI 工具(==卸载 partitionmanager==)  |
| termius                | 终端连接工具                                 |
| remmina                | aur/remmina-plugin-rdesktop (rdp)            |
| wireshark-qt           | 抓包工具                                     |
| yakuake                | 便捷终端工具                                 |
| flameshot              | 截图工具                                     |
| mailspring             | 邮箱工具                                     |
| localsend-bin          | 传输工具（替代 kde connect），需要开启防火墙 |
| fcitx-sogoupinyin      | 自然码双拼                                   |
| idea                   | java 编辑器                                  |
| pycharm                | python 编辑器                                |
| goland                 | go 编辑器                                    |
| datagrip               | 数据库 编辑器                                |
| webstorm               | 前端 编辑器                                  |
| clion                  | C 编辑器                                     |
| visual-studio-code-bin | 万能编辑器                                   |
| gimp                   | PS 工具                                      |
| wechat                 |                                              |
| docker-desktop         |                                              |
| picgo                  | 图床工具                                     |
|                        |                                              |
|                        |                                              |
|                        |                                              |

### Basic Function Software

| Name      | Description          |
| --------- | -------------------- |
| bind      | 包含 dig             |
| tmux      | 终端多开工具         |
| ranger    | TUI 文件浏览         |
| vim       | TUI 编辑工具         |
| vagrant   | 虚拟机管理工具       |
| git       | 版本控制工具         |
| gh        | github 管理工具      |
| flatpak   | 3 方下载工具         |
| snapd     | 3 方下载工具         |
| nmap      | 网络探测工具         |
| docker    | 镜像工具             |
| downgrade | package 版本降级工具 |
| uxplay    | apple airplay 投屏   |
| strace    |                      |
| whois     |                      |
|           |                      |
|           |                      |
|           |                      |
|           |                      |

:cn:

## 

**references**

[^1]:https://www.fosslinux.com/46741/things-to-do-after-installing-manjaro.htm
[^2]: https://wiki.archlinux.org/title/PC_speaker
[^3]:https://coda.world/manjaro-optimization-and-setting
[^4]:https://segmentfault.com/a/1190000039901064
[^5]:https://b3logfile.com/pdf/article/1615812052038.pdf
[^6]:https://averagelinuxuser.com/manjaro-xfce-after-install/