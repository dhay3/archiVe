# OpenWrt Zsh Installation

ref

https://stackoverflow.com/questions/51366101/git-remote-https-is-not-a-git-command

```
#安装 zsh
opkg install zsh
#安装 git
opkg install git
#安装 git-http, 如果不安装会出现 git: 'remote-https' is not a git command.
opkg install git-http
#安装 ohmyzsh
sh -c "$(wget -O- https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
#这里需要注意一下 zsh 的路径，如果 /etc/passwd shell 指向错误 可能会导致不能通过 ssh 登录 tty
```

