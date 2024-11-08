# custom alias

```
alias ls='lsd'
alias ll='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'
alias pacman='pacman --color always'
alias grep='grep --color=auto -n'
alias fgrep='fgrep --color=auto -n'
alias egrep='egrep --color=auto -n'
alias diff='diff --color=auto'
alias radeontop='radeontop -c'
alias dmesg='dmesg --color=always'
alias cp='rsync --progress -azvh'
# it will remove the source files without the parent directories
# alias mv='rsync --remove-source-files -azvh'
alias mv='mv -v'
alias rm='rm -vi'
alias ip='ip -c=always'
alias split='split --verbose'
alias vbox='VirtualBox %U'

alias -s {json,yaml,yml,html}=code
alias -s md=typora
alias -s txt=subl
```

