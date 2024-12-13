# Linux 自定义命令

## 临时生效

pattern 

alias identifier=command

```
[root@chz Desktop]# alias hello='echo "hello world"'
[root@chz Desktop]# hello
hello world
```

如果没有给定参数，就会列出当前所有的

```
[root@chz Desktop]# alias
alias cp='cp -i'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'
alias grep='grep --color=auto'
alias l.='ls -d .* --color=auto'
alias ll='ls -l --color=auto'
alias ls='ls --color=auto'
alias mv='mv -i'
alias rm='rm -i'
alias which='alias | /usr/bin/which --tty-only --read-alias --show-dot --show-tilde'
```

## 永久生效

> Debian 系统在 `/etc/bash.bashrc`

修改`~/.bashrc`对当前用户生效, 修改`/bashrc`对所有用户生效

```
..               .bash_profile.bak  .cshrc     .esd_auth             .mozilla   .tcshrc    .xauthVoCfdB
anaconda-ks.cfg  .bashrc            .dbus      .ICEauthority         Music      Templates
.bash_history    .bashrc.bak        Desktop    initial-setup-ks.cfg  Pictures   Videos
.bash_logout     .cache             Documents  .local                .pki       .viminfo
[root@chz ~]# vim .bashrc
[root@chz ~]# source .bashrc
[root@chz ~]# test
alias test
[root@chz ~]# 
```

Debian中添加如下命令行，可以高亮显示

```
alias yum='yum --color=always'
alias ls='ls --color'
alias ll='ls -lh --color'
grep='grep --color'
```

