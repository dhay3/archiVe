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

Use Clash verge rev instead

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

| Name                     | Description                                    |
| ------------------------ | ---------------------------------------------- |
| virtualbox               | 虚拟机前端                                     |
| brave                    | 替代 chrome                                    |
| zen browser              | 替代 firefox                                   |
| tor browser              | 一键 tor                                       |
| timeshift                | 备份                                           |
| clash verge rev          | 代理客户端                                     |
| sweeper                  | cache 清理                                     |
| bitwarden                | keepass                                        |
| qbittorrent              | bt 下载                                        |
| obsidian                 | 知识库                                         |
| typora                   | markdown 编辑器                                |
| ventoy                   | usb 多启动盘                                   |
| balena-etcher            | iso 烧录                                       |
| beless hex editor        | binary 编辑器                                  |
| sublime-text-4           | 文本编辑器                                     |
| telegram-desktop         | 聊天                                           |
| discord                  | 聊天                                           |
| gparted                  | gparted GUI (==卸载 partitionmanager==)        |
| gns3                     | 网络设备模拟器                                 |
| packettracer             | CISCO 模拟器                                   |
| remote-desktop-manager   | 终端连接                                       |
| wireshark-qt             | 抓包                                           |
| yakuake                  | 便捷终端                                       |
| flameshot                | 截图                                           |
| vlc                      | 视频                                           |
| spotube                  | 听歌 替代 spotify                              |
| netease cloud music gtk4 | 听歌                                           |
| betterbird               | 邮箱                                           |
| localsend-bin            | 传输（替代 kde connect），需要开启防火墙       |
| idea                     | java 编辑器                                    |
| pycharm                  | python 编辑器                                  |
| goland                   | go 编辑器                                      |
| datagrip                 | 数据库 编辑器                                  |
| webstorm                 | 前端 编辑器                                    |
| clion                    | C 编辑器                                       |
| visual-studio-code-bin   | 万能编辑器                                     |
| postman                  | api 调试                                       |
| gimp                     | PS                                             |
| wechat-universal-bwrap   | 微信                                           |
| docker-desktop           | docker 桌面管理                                |
| burpsuite                | 7 层报文抓包                                   |
| fluent-reader            | Rss 阅读(在 IOS 上可以配合 inoreader 一起使用) |
| kdiskmark                | hard drive benchmark tool                      |
| github-desktop           | github 桌面管理                                |
| nessus                   | 漏洞扫描                                       |
| johnny                   | 爆破 GUI                                       |
| kvantum                  | KDE applications 美化                          |
| WPS                      | office 套件                                    |
| Pling Store              | KDE 插件商店                                   |
| gestures                 | libinput-gestures 图形化配置                   |
| DVWA                     | Web 漏洞测试                                   |
| Obs                      | 录屏                                           |
| kdenlive                 | 视频编辑                                       |
| kdiff3                   | diff GUI                                       |
| geeqie                   | 图片浏览                                       |
| octopi                   | 包管理                                         |
| zenmap                   | 端口扫描                                       |
| vmware remote console    | vcenter 控制台                                 |
| eyedropper               | color picker                                   |
| upscayl                  | 图片修复/steghide                              | 


###  Function

| Name                                                     | Description                                            |
| -------------------------------------------------------- | ------------------------------------------------------ |
| zsh                                                      | Shell                                                  |
| kitty                                                    | GPU 终端                                               |
| tmux                                                     | 终端多开                                               |
| zfs                                                      | filesystem                                             |
| transhell                                                | TUI 翻译                                               |
| bind                                                     | 包含 dig                                               |
| yazi                                                     | TUI 文件浏览                                           |
| asciinema                                                | 终端录屏                                               |
| mpv                                                      | TUI 视频浏览                                           |
| spotify-tui/spotifyd                                     | spotify  tui 需要 premium 账号                         |
| yt-dlp                                                   | YouTube 下载工具                                       |
| vim/nvim                                                 | TUI 编辑                                               |
| vagrant                                                  | 虚拟机管理                                             |
| git/git-lfs                                              | 版本控制                                               |
| gh                                                       | github 管理                                            |
| flatpak                                                  | 3 方下载                                               |
| nmap                                                     | 网络探测                                               |
| docker                                                   | 容器                                                   |
| manjaro-downgrade                                        | package 版本降级工具,不要使用 downgrade AUR 会有问题的 |
| uxplay                                                   | apple airplay 投屏                                     |
| strace                                                   | binary system calls debug                              |
| ltrace                                                   | binary library calls debug                             |
| gdb                                                      | binary debug                                           |
| whois                                                    | 域名收集                                               |
| binwalk                                                  | 逆向                                                   |
| ddrescue                                                 | 数据恢复                                               |
| iat                                                      | iso 格式转换                                           |
| john                                                     | 爆破                                                   |
| hashcat                                                  | 爆破                                                   |
| ocs-url                                                  | kde store 下载                                         |
| gping                                                    | ping TUI                                               |
| sysstat                                                  | 包含 sar                                               |
| manjaro-asian-input-support-fcitx5/fcitx5-chinese-addons | 中文输入法(支持双拼)                                   |
| rime/rime-ice-double-pinyin                              | 跨平台双拼                                             |
| aria2                                                    | 下载替代 wget                                          |
| libinput-gestures                                        | 自定义 touchpad(不能覆盖 KDE built-in gestures)        |
| nerd-fonts                                               | hack nerd font                                         |
| woff2                                                    | WPS 字体                                               |
| aur-auto-vote-git                                        | AUR 自动投票                                           |
| metasploit                                               | payload 木马病毒                                       |
| sherlock                                                 | 社工                                                   |
| wifiphisher                                              | WIFI 钓鱼                                              |
| setoolkits                                               | 社工 钓鱼                                              |
| ettercap                                                 | DNS 投毒/DOS                                           |
| macchanger                                               | BIA 修改                                               |
| steghide                                                 | 文件隐写                                               |
| aircrack-ng                                              | WIFI 破解                                              |
| hydra                                                    | 爆破工具                                               |
| beef-xss                                                 | XSS                                                    |
| sqlmap                                                   | SQL 注入                                               |
| anacoda                                                  | Python env 管理                                        |
| nix                                                      | cross platform 包管理器                                |
| ipython                                                  | 命令行 Python                                          |
| iotop                                                    | io 监测                                                |
| btop                                                     | top 替代                                               |
| nexttrace                                                | traceroute 替代                                        |
| duf                                                      | df 替代                                                |
| procs                                                    | ps 替代                                                |
| ripgrep                                                  | grep 替代                                              |
| fd                                                       | find 替代                                              |
| dust                                                     | du 替代                                                |
| fastfectch                                               | neofectch 替代                                         |
| fuck                                                     | command 提示                                           |
| fzf                                                      | zsh 插件                                               |
| pandoc                                                   | 格式转换                                               |
| figlet                                                   | toliet  替代                                           |
| gum                                                      | funny script                                           |
| stow                                                     | 批量 symlink 配置文件管理                              |
| gitmoji                                                  | git commit                                             |
| lazygit                                                  | 懒人 git                                               | 
| cowsay/lolcat                                            | garbage                                                |

**references**

[^1]:https://www.fosslinux.com/46741/things-to-do-after-installing-manjaro.htm
[^2]: https://wiki.archlinux.org/title/PC_speaker
[^3]:https://coda.world/manjaro-optimization-and-setting
[^4]:https://segmentfault.com/a/1190000039901064
[^5]:https://b3logfile.com/pdf/article/1615812052038.pdf
[^6]:https://averagelinuxuser.com/manjaro-xfce-after-install/