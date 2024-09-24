```
#!/usr/bin/env zsh
#set -x
#######################################################################
# Pre-scripts
#######################################################################
# Start Tmux as the default Shell for user execlude dolphin and jetbrain
if [[ -x "$(command -v tmux)" ]] && [[ -n "${DISPLAY}" ]] && [[ -z "${TMUX}" ]]; then
    if [[ ! "$(readlink -f /proc/${PPID}/exe)" =~ "dolphin" ]] &&
        [[ ! "$(readlink -f /proc/${PPID}/exe)" =~ "jetbrain" ]]; then
        exec tmux
    fi
fi

# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/home/cc/anaconda3/bin/conda' 'shell.zsh' 'hook' 2>/dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/home/cc/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/home/cc/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="/home/cc/anaconda3/bin:$PATH"
    fi
fi
# <<< conda initialize <<<

#######################################################################
# Environments
#######################################################################
# Theme
ZSH_THEME="half-life"
#ZSH_THEME="bira"

# History Settings
HISTFILE=~/.zhistory
HISTSIZE=4096
SAVEHIST=4096

# Enviroment Virables
export ZSH="/home/${USER}/.oh-my-zsh"
# It will cause less to subl in editor mod which is unexpected
# But if use bat instead of less it's not matter cause bat do not support editor mod right now
#export VISUAL="/usr/bin/subl"
export VISUAL="/usr/bin/vim"
export EDITOR="/usr/bin/vim"
export UPDATE_ZSH_DAYS=30
export LANG=en_US.UTF-8
export FZF_BASE=/usr/share/fzf
# Set firefox as the default browser for web-search plugin
export BROWSER="firefox"

#######################################################################
# Plugins
#######################################################################
# >>>> Built-in plugins (start)
plugins=(
    colored-man-pages
    command-not-found
    extract
    fancy-ctrl-z
    fzf
    fzf-tab
    gh
    git
    sudo
    themes
    web-search
    conda-zsh-completion
)
# <<<< Built-in plugins (end)

# >>>> Extra plugins (start)
eval "$(thefuck --alias)"
eval "$(zoxide init zsh)"
# fasd is archived now
#eval "$(fasd --init auto)"
# Use control + g to activate navi
eval "$(navi widget zsh)"
fpath+=/usr/share/zsh/plugins/zsh-completions/src
[ -f /usr/share/zsh/plugins/zsh-autopair/autopair.zsh ] && source /usr/share/zsh/plugins/zsh-autopair/autopair.zsh
[ -f /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh ] && source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
# Use up-arrow or down-arrow to show candidate suggestions
# Use right-arrow to accept the suggestion
[ -f /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh ] && source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh
[ -f $ZSH/oh-my-zsh.sh ] && source $ZSH/oh-my-zsh.sh
autopair-init
# <<<< Extra plugins (end)

# >>>> Vagrant command completion (start)
fpath+=/opt/vagrant/embedded/gems/gems/vagrant-2.4.0/contrib/zsh
compinit
# <<<<  Vagrant command completion (end)

#######################################################################
# Custom scripts
#######################################################################
#fastfetch
[[ -x "$(command -v fastfetch)" ]] && fastfetch --disable-linewrap

# Print a new line after command excuted
precmd() {
    print ""
}

# Use gpg-agent instead of ssh-agent
unset SSH_AGENT_PID
if [ "${gnupg_SSH_AUTH_SOCK_by:-0}" -ne $$ ]; then
    export SSH_AUTH_SOCK="$(gpgconf --list-dirs agent-ssh-socket)"
fi
export GPG_TTY=${TTY:-"$(tty)"}
gpg-connect-agent updatestartuptty /bye >/dev/null

#######################################################################
# Bindings
# ^ means ctrl
# ^[ means escape
#######################################################################
# Set cursor to the beginning of a line
bindkey -M main '^A' beginning-of-line
# Set cursor to the end of a line
bindkey -M main '^E' end-of-line

# Set cursor backward one word
# It's conflict with Tmux prefix, bind Tmux prefix to Ctrl + X
bindkey -M main '^B' backward-word
# Set cursor Forward one word
bindkey -M main '^F' forward-word

# Delete words before the cursor
bindkey -M main '^U' backward-kill-line
# Delete words after the cursor
bindkey -M main '^K' kill-line

# Delete one word before the cursor
bindkey -M main '^W' backward-kill-word
# Delete one word after the cursor
bindkey -M main '^D' kill-word

# Search history backword one line
bindkey -M main '^P' history-search-backward
# Search history forward one line
bindkey -M main '^N' history-search-forward

# Search history via fzf
bindkey -M main '^H' fzf-history-widget
# Search current files via fzf
bindkey -M main '^Q' fzf-file-widget

bindkey -M main '^Y' yank
bindkey -M main '^L' clear-screen

#######################################################################
# Aliases
#######################################################################
# >>>> Built-in aliases (start)
alias c='clear'
alias ls='lsd'
alias l='ll'
alias ll='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'
alias ln='ln -v'
#alias top='btop'
alias cat='bat -pp'
alias less='bat'
alias more='bat'
#alias man='man -P less'
#alias grep='rg'
#alias find='fd'
# Recursive copy will create a dirctory name of the source, it should be trailing slash on the source to copy the contents of the directoy
#alias cp='rsync --progress -azvh'
alias cp='cp -v'
alias mkdir='mkdir -v'
alias mk='mkdir -v'
alias mv='mv -v'
# It is better do not use trash-put
alias rm='trash-put -v'
#alias rm='rm -v'
alias du='dust'
#alias du='du -h'
alias df='duf'
#alias df='df -h'
#alias ps='procs'
alias free='free -h'
alias ip='ip -c=always'
alias diff='diff --color=auto'
alias dmesg='dmesg --color=always'
alias split='split --verbose'
alias pacman='pacman --color always'
# <<<< Built-in aliases (end)

# >>>> Extra aliases (start)
alias n='navi'
alias nc='ncat'
alias jq='jq -C'
alias lynx='lynx -display_charset=utf-8'
alias fzf='fzf --reverse'
alias vbox='VirtualBox %U'
alias vag='vagrant'
alias geeq='geeqie'
alias wirek='wireshark'
alias typo='typora'
alias rag='ranger'
alias wbs='web_search duckduckgo'
alias ytd='yt-dlp'
alias rdm='remotedesktopmanager'
alias xfreerdp='xfreerdp /cert:tofu /fonts /bpp:64 /video /dynamic-resolution /scale:140 /scale-desktop:125'
# Alias for logout KDE plasma with cancel menu
alias logout="qdbus org.kde.LogoutPrompt /LogoutPrompt org.kde.LogoutPrompt.promptLogout"
# <<<< Extra aliases (end)

# >>>> File assosicated aliases (start)
alias -s {json,yaml,yml,txt}=subl
alias -s {mp4,webm,avi}=vlc
alias -s {png,jpeg,jpg}=geeqie
alias -s {xlsx,cvs}=et
alias -s html=firefox
alias -s docs=wps
alias -s ppt=wpp
alias -s pdf=okular
alias -s md=typora
# <<<< File assosicated aliases (end)
```