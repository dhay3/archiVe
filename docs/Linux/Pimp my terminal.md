# Pimp my terminal

转载：https://drasite.com/blog/Pimp%20my%20terminal

> 本文以Kali Linux & Bash 为基础
>
> ```
> [root@chz themes]# cat /etc/os-release 
> PRETTY_NAME="Kali GNU/Linux Rolling"
> NAME="Kali GNU/Linux"
> ID=kali
> VERSION="2020.3"
> VERSION_ID="2020.3"
> VERSION_CODENAME="kali-rolling"
> ID_LIKE=debian
> ANSI_COLOR="1;31"
> HOME_URL="https://www.kali.org/"
> SUPPORT_URL="https://forums.kali.org/"
> BUG_REPORT_URL="https://bugs.kali.org/"
> 
> ```
>
> 

## Pathced Fonts

https://github.com/ryanoasis/nerd-fonts

https://github.com/source-foundry/Hack

这里使用Hack字体，下载即可自动识别。

如果想要修改系统字体Settings -> appearance

如果想要预定义的glyphs，参考

https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts/Hack#ligatures

```
E:\IDMdownLoad\Compressed>scp -r Hack-v3.003-ttf root@8.135.0.171:/usr/share/fonts/
Hack-Bold.ttf                                                                         100%  310KB   1.1MB/s   00:00
Hack-BoldItalic.ttf                                                                   100%  315KB   1.0MB/s   00:00
Hack-Italic.ttf                                                                       100%  309KB   1.5MB/s   00:00
Hack-Regular.ttf                                                                      100%  302KB   1.5MB/s   00:00
```

将下载后文件放在`/usr/share/fonts/hack/`下即可

## Coloring

### color scheme

https://github.com/Mayccoll/Gogh

> 由于网络问题，推荐使用离线安装
>
> 通用方式==使用`ps -h -o comm -p $PPID`查看当前使用的terminal==
>
> 如果无效请参考：
>
> https://stackoverflow.com/questions/16808231/how-do-i-set-default-terminal-to-terminator
>
> then reboot :)

在使用前查看使用的terminal，==注意gogh不支持qterminal==

```
[root@chz alternatives]# ll /etc/alternatives | grep terminal
lrwxrwxrwx 1 root root  18 Aug 31 09:05 x-terminal-emulator -> /usr/bin/qterminal
```

1. 安装gnome-terminal

   ```
   apt install gnome-terminal*
   ```

2. 配置gnome-terminal为默认terminal

   ```
   [root@chz alternatives]# update-alternatives --config x-terminal-emulator
   There are 3 choices for the alternative x-terminal-emulator (providing /usr/bin/x-terminal-emulator).
   
     Selection    Path                             Priority   Status
   ------------------------------------------------------------
     0            /usr/bin/gnome-terminal.wrapper   40        auto mode
     1            /usr/bin/gnome-terminal.wrapper   40        manual mode
   * 2            /usr/bin/mate-terminal.wrapper    30        manual mode
     3            /usr/bin/qterminal                40        manual mode
   
   Press <enter> to keep the current choice[*], or type selection number: 0
   update-alternatives: using /usr/bin/gnome-terminal.wrapper to provide /usr/bin/x-terminal-emulator (x-terminal-emulator) in manual mode
   ```

   上述`gnome-terminal.wrapper`已经处于auto mode 所以无需修改

**安装可能出现问题：**

https://github.com/Mayccoll/Gogh/issues/203

==顺带关闭gnome terminal 的 sound bell==

<img src="..\..\imgs\_Linux\Snipaste_2020-12-12_15-07-11.png"/>

> 如果需要添加字形(glyphs)，请使用 [Font Awesome ➶](https://github.com/FortAwesome/Font-Awesome), [Devicons ➶](https://vorillaz.github.io/devicons/), [Octicons ➶](https://github.com/primer/octicons), and [others](https://github.com/ryanoasis/nerd-fonts#glyph-sets).

### aliases

`/etc/bashrc || ~/.bashrc`添加

```
alias ls='ls --color=auto'
alias grep='grep --color=auto'
alias fgrep='fgrep --color=auto'
alias egrep='egrep --color=auto'
alias diff='diff --color=auto'
```

### man pages

同样的可以在`/etc/bashrc || ~/.bashrc`文件中添加

```
export LESS_TERMCAP_mb=$'\e[1;32m'
export LESS_TERMCAP_md=$'\e[1;32m'
export LESS_TERMCAP_me=$'\e[0m'
export LESS_TERMCAP_se=$'\e[0m'
export LESS_TERMCAP_so=$'\e[01;33m'
export LESS_TERMCAP_ue=$'\e[0m'
export LESS_TERMCAP_us=$'\e[1;4;31m'
```

## LSD

<img src="..\..\imgs\_Linux\Snipaste_2020-12-12_16-05-35.png"/>

https://github.com/Peltoche/lsd

> 由于没有apt的安装方式，我们选择Form Binaries，选择对应的版本

我们将下载文件放在`/usr/local/bin/`下，方便调用

```
[root@chz Desktop]# cd lsd-0.18.0-x86_64-unknown-linux-gnu && mv lsd /usr/local/bin/
```

修改`/etc/bash.bashrc || ~/.bashrc`

```
alias ls='lsd'
alias l='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'
```

## YTop

https://github.com/cjbassi/ytop

<img src="..\..\imgs\_Linux\Snipaste_2020-12-12_17-03-47.png"/>

和LSD一样需要放在`/usr/local/bin`下

## bat

https://github.com/sharkdp/bat

如果使用`apt`安装需要注意的是`bat`安装的名字为`batcat`。我们创建一个软连接，然后配置别名。

==注意的是查看manual page 还是通过`man bat`==来查看

```
[root@chz bin]# pwd
/usr/local/bin
[root@chz bin]# ln -s /usr/bin/batcat bat
[root@chz bin]# ls
 bat   lsd   ytop
[root@chz bin]# ll
lrwxrwxrwx root root  15 B  Sat Dec 12 03:24:03 2020  bat ⇒ /usr/bin/batcat
.rwxrw-rw- root root 2.4 MB Sat Aug 29 09:53:57 2020  lsd
.rwxr-xr-x root root 5.1 MB Sat Dec 12 03:08:44 2020  ytop
```

配置文件查看

https://github.com/sharkdp/bat#configuration-file

## Prompt

### starship

https://github.com/starship/starship

https://starship.rs/config/#prompt

https://starship.rs/config/#character

https://starship.rs/config/#username

> 网络问题，推荐使用离线下载

### manual

> precodition patched font

https://www.nerdfonts.com/cheat-sheet

<img src="..\..\imgs\_Linux\Snipaste_2020-12-16_13-02-45.png"/>

```
OS_ICON=

PS1="\n \[\033[0;34m\]╭─\[\033[0;31m\]\[\033[0;37m\]\[\033[41m\] $OS_ICON \u\[\033[0m\]\[\033[0;31m\]\[\033[44m\]\[\033[0;34m\]\[\033[44m\]\[\033[0;30m\]\[\033[44m\] \w \[\033[0m\]\[\033[0;34m\] \n \[\033[0;34m\]╰ \[\033[1;36m\]\$ \[\033[0m\]"
```

## Bash Completion Plugin

https://github.com/scop/bash-completion

