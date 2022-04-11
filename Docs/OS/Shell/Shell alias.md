# Shell alias

> ==如果不是login shell就不会在子Shell或脚本中生效==

为一个命令设置一个别名，当tty或shell关闭时就会失效，可以通过配置设置为永久生效

一般来说，都会把常用的别名写在`~/.bashrc`的末尾。

pattern：`alias NAME=DEFINITION`

> 如果DEFINITION中有空格需要使用

```
[root@chz Desktop]# alias wm='whereis mysql'
[root@chz Desktop]# wm
mysql: /usr/lib64/mysql /usr/share/mysql
```

显示所有的别名

```
[root@cyberpelican opt]# alias
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

解除别名

```
unalias ll
```

