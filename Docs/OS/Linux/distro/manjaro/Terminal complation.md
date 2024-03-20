# Terminal complation

## Shell

### Install ohmyzsh

```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

### Plugins 

| Name                                                         | Description                   |
| ------------------------------------------------------------ | ----------------------------- |
| [zsh-completions](https://github.com/zsh-users/zsh-completions) | Shell 补全                    |
| [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting) | Shell 高亮                    |
| [zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions) | Shell 推荐                    |
| [zoxide](https://github.com/ajeetdsouza/zoxide)              | 目录跳转和 z/autojump 类似    |
| [fzf](https://github.com/junegunn/fzf)                       | 文件搜索                      |
| [fuck](https://github.com/nvbn/thefuck)                      | 错误提示                      |
| [zsh-autopair](https://github.com/hlissner/zsh-autopair)     | 括号自动补全                  |
| [colored-man-pages](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/colored-man-pages) | man page 高亮                 |
| [extract](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/extract) | 快解压                        |
| [gh](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/gh) | gh cli 命令补全               |
| [git](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/git) | git 命令的 aliases            |
| [sudo](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/sudo) | 双 <kbd>ECS</kbd> 替代 `sudo` |

#### Installation

```shell
yay -Sy fzf thefuck zsh-autopair zsh-autosuggestions zsh-syntax-highlightin zsh-completions
```

### Modern Commands

| Name                                                    | Description                                            |
| ------------------------------------------------------- | ------------------------------------------------------ |
| [lsd](https://github.com/lsd-rs/lsd)                    | 替代 ls                                                |
| [btop](https://github.com/aristocratos/btop)            | 替代 top                                               |
| [bat](https://github.com/sharkdp/bat)                   | 替代 cat                                               |
| [trash-cli](https://github.com/andreafrancia/trash-cli) | trash can 交互和 rm 互补（需要在 anaconda 环境下安装） |
| [dust](https://github.com/bootandy/dust)                | 替代 du                                                |
| [duf](https://github.com/muesli/duf)                    | 替代 df                                                |
| [ripgrep](https://github.com/BurntSushi/ripgrep)        | 替代 grep                                              |
| [procs](https://github.com/dalance/procs)               | 替代 pf                                                |
| [tldr](https://github.com/tldr-pages/tldr)              | 和 man 互补在线帮助信息                                |
| [navi](https://github.com/denisidoro/navi)              | 和 man/tldr 互补                                       |
| [fd](https://github.com/sharkdp/fd)                     | 替代 find                                              |

```
yay -Sy lsd btop bat tldr ranger dust duf ripgrep procs navi fd
./Anaconda3-*.zsh
python install trash-cli
```

### ~/.zshrc

```Shell
# Start Tmux as the default Shell for user execlude dolphin and jetbrain
if [[ -x "$(command -v tmux)" ]] && [[ -n "${DISPLAY}" ]] && [[ -z "${TMUX}" ]]; then
    if [[ ! "$(readlink -f /proc/${PPID}/exe)" =~ "dolphin" ]] && [[ ! "$(readlink -f /proc/${PPID}/exe)" =~ "jetbrain" ]]; then
        exec tmux
    fi
fi

# Theme
ZSH_THEME="half-life"

# History Settings
HISTFILE=~/.zhistory
HISTSIZE=4096
SAVEHIST=4096

# Enviroment Virables
export ZSH="/home/cc/.oh-my-zsh"
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
eval "$(thefuck --alias)"
eval "$(zoxide init zsh)"
# fasd is archived now
#eval "$(fasd --init auto)"
# Use control + g to activate navi
eval "$(navi widget zsh)"

# Plugins
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
)

# Extra
fpath+=/usr/share/zsh/plugins/zsh-completions/src
[ -f /usr/share/zsh/plugins/zsh-autopair/autopair.zsh ] && source /usr/share/zsh/plugins/zsh-autopair/autopair.zsh
[ -f /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh ] && source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
# Use up-arrow or down-arrow to show candidate suggestions
# Use right-arrow to accept the suggestion
[ -f /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh ] && source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh
[ -f $ZSH/oh-my-zsh.sh ] && source $ZSH/oh-my-zsh.sh
autopair-init

# >>>> Vagrant command completion (start)
fpath=(/opt/vagrant/embedded/gems/gems/vagrant-2.4.0/contrib/zsh $fpath)
compinit
# <<<<  Vagrant command completion (end)

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
unset __conda_setup
# <<< conda initialize <<<

#fastfetch
if [[ -x "$(command -v fastfetch)" ]]; then
    fastfetch --logo Arya \
        --set Colors="" \
        --set Cursor="" \
        --set Locale="" \
        --set DE="" \
        --set WM="" \
        --set WMTheme="" \
        --set Theme="" \
        --set Icons="" \
        --set Font="" \
        --set Terminal="" \
        --set TerminalFont=""
fi

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

# Custom zsh bindings
# ^ means ctrl
# ^[ means escape

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

bindkey -M main '^H' fzf-history-widget
bindkey -M main '^Q' fzf-file-widget

bindkey -M main '^Y' yank
bindkey -M main '^L' clear-screen

# Aliases
alias c='clear'
alias n='navi'
alias l='ll'
alias ls='lsd'
alias ll='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'
alias ln='ln -v'
# Recursive copy will create a dirctory name of the source, it should be trailing slash on the source to copy the contents of the directoy
#alias cp='rsync --progress -azvh'
alias cp='cp -v'
alias mkdir='mkdir -v'
alias mk='mkdir'
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
alias nc='ncat'
alias ip='ip -c=always'
alias jq='jq -C'
#alias top='btop'
alias cat='bat -pp'
alias less='bat'
alias more='bat'
#alias man='man -P less'
#alias grep='rg'
#alias find='fd'
alias vbox='VirtualBox %U'
alias rdm='remotedesktopmanager'
alias xfreerdp='xfreerdp /cert:tofu /fonts /bpp:64 /dynamic-resolution /scale:140 /scale-desktop:125'
# Alias for logout KDE plasma with cancel menu
alias logout="qdbus org.kde.ksmserver /KSMServer logout 1 0 1"
alias lynx='lynx -display_charset=utf-8'
alias fzf='fzf --reverse'
alias diff='diff --color=auto'
alias dmesg='dmesg --color=always'
alias split='split --verbose'
alias pacman='pacman --color always'
alias -s {json,yaml,yml,html}=subl
alias -s md=typora
alias -s txt=subl
```

## Vim

1.  Install [vundle](https://github.com/VundleVim/Vundle.vim)

   ```
   git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
   ```

2. create a `.vimrc` (check below)

3. Install Plugins in vim

   ```
   :PluginInstall
   ```

4. compile youcompleteme

   ```
   pacman -S cmake base-devel nmp nodejs
   python3 ~/.vim/bundle/youcompleteme/install.py --all
   ```

### ~/.vimrc

```
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Maintainer:
"    Cyberpelican
" Version:
"    0.3
"
" Common references
" github.com/amix/vimrc/blob/master/vimrcs/basic.vim
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => General Options
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Common Global Options
set nocompatible
set encoding=utf-8
set confirm
set fileformats=unix,dos,mac
set backspace=indent,eol,start

" Disable Annoying Sound On Error
set noerrorbells
set novisualbell
set t_vb=
set tm=500

" Indention Options
set autoindent
set smartindent
set smarttab
set tabstop=2
set expandtab
set shiftwidth=2
set shiftround

" Search Options
set hlsearch
set magic
set ignorecase
set incsearch
set smartcase
set showmatch
set matchtime=5

" Command Options
set cmdheight=1
set wildmenu

" Diff Options
" set diff
" set diffopt=filler,vertical

" Display Options
set title
set ruler
set number
set laststatus=2
set statusline=\ %{HasPaste()}%F%m%r%h\ %w\ \ CWD:\ %r%{getcwd()}%h\ \ \ Line:\ %l\ \ Column:\ %c
set scrolloff=7
syntax on
" colorscheme torte

" Miscellaneous Options
set lazyredraw
set autoread
set noautowrite
set spell
set wrap
set nocursorline
set shell=sh
set regexpengine=0
filetype plugin on


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Maps
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" normal mode maps
map <F2> GoDate: <Esc>:read !date<CR>kJ
" map <space> /

" insert mode maps
imap <F2> <CR>Date: <Esc>:read !date<CR>kJa<CR>


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Autocmds
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" return to last edit position when opening files
autocmd BufReadPost * if line("'\"") > 1 && line("'\"") <= line("$") | exe "normal! g'\"" | endif

" delete trailing white space on save
if has("autocmd")
    autocmd BufWritePre *.txt,*.js,*.java,*.py,*.go,*.yml,*.yaml,*.json,*.sh :call CleanExtraSpaces()
endif

" NERDTree
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * NERDTree | if argc() > 0 || exists("s:std_in") | wincmd p | endif
autocmd BufEnter * if winnr('$') == 1 && exists('b:NERDTree') && b:NERDTree.isTabTree() | quit | endif

" code autocompletion
" autocmd FileType css  set omnifunc=csscomplete#CompleteCSS
" autocmd FileType js   set omnifunc=javascriptcomplete#CompleteJS
" autocmd FileType html set omnifunc=htmlcomplete#CompleteTags
" autocmd FileType php  set omnifunc=phpcomplete#CompletePHP


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Code
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
function! HasPaste()
    if &paste
        return 'PASTE MODE  '
    endif
    return ''
endfunction

fun! CleanExtraSpaces()
    let save_cursor = getpos(".")
    let old_query = getreg('/')
    silent! %s/\s\+$//e
    call setpos('.', save_cursor)
    call setreg('/', old_query)
endfun


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Vundle
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
Plugin 'VundleVim/Vundle.vim'
call vundle#end()


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Plugins
" vimawesome.com
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plugin 'scrooloose/nerdtree'
Plugin 'ryanoasis/vim-devicons'
Plugin 'vim-airline/vim-airline'
Plugin 'vim-airline/vim-airline-themes'
Plugin 'valloric/youcompleteme'
Plugin 'scrooloose/syntastic'
Plugin 'altercation/vim-colors-solarized'
Plugin 'junegunn/fzf'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => NERDTree
" vimawesome.com/plugin/nerdtree-red
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" let g:NERDTreeDirArrowExpandable = '▸'
" let g:NERDTreeDirArrowCollapsible = '▾'

let g:NERDTreeShowHidden=1
let g:webDevIconsOS='Manjaro'


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Airline
" vimawesome.com/plugin/vim-airline-superman
" vimawesome.com/plugin/vim-airline-themes
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
let g:airline#extensions#tabline#enabled = 1
let g:airline_theme='violet'


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Syntastic
" https://vimawesome.com/plugin/syntastic
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*

let g:syntastic_always_populate_loc_list = 1
let g:syntastic_auto_loc_list = 1
let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 0
let g:syntastic_sh_checkers = ['sh']
let g:syntasitc_c_checkers = ['gcc']
let g:syntasitc_cpp_checkers=['cpplint']
let g:syntasitc_python_checkers = ['pylint']
let g:syntasitc_go_checkers = ['golint']
let g:syntasitc_java_checkers = ['javac']
let g:syntasitc_html_checkers = ['eslint']
let g:syntastic_javascript_checkers = ['jslint']
let g:syntasic_json_checkers = ['jsonlint']
let g:syntastic_css_checkers = ['csslint']
let g:syntastic_ruby_checkers = ['rubylint']
let g:syntastic_sql_checkers = ['sqllint']
let g:syntastic_xml_checkers = ['xmllint']
let g:syntastic_yaml_checkers = ['yamllint']
let g:syntasitc_dockerfile_checkers = ['dockerfile_lint']
let g:syntastic_lua_checkers = ['']
let g:syntastic_perl_checkers = ['']


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Solarized
" vimawesome.com/plugin/vim-colors-solarized-ours
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
let g:solarized_termcolors = 256
set background=dark
colorscheme solarized
```

## Tmux

### ~/.tmux.conf

```
# Plugins
# Use prefix + I to install plugins
set-option -g @plugin 'tmux-plugins/tpm'
set-option -g @plugin 'tmux-plugins/tmux-sensible'
# Use prefix + Ctrl + s to save tmux status
# Use prefix + Ctrl + r to restore tmux status
set-option -g @plugin 'tmux-plugins/tmux-resurrect'
set-option -g @plugin 'fabioluciano/tmux-tokyo-night'

run '~/.tmux/plugins/tpm/tpm'

# Global Opotions
set-option -gq prefix C-x
set-option -gq prefix2 None
set-option -gq default-terminal "screen-256color"
set-option -gq mode-keys vi
set-option -gq status on
set-option -gq monitor-bell off
set-option -gq renumber-windows on
set-option -gq set-clipboard on
set-option -gq set-titles on
set-option -gq mouse on
set-option -gq pane-border-lines simple
set-option -gp pane-border-indicators both
set-option -gq pane-active-border-style 'bg=default fg=#FF1493'
set-option -gq base-index 1
set-option -gq pane-base-index 1
#xclip for xorg
#set-option -gq copy-command 'xclip -i'
#wl-copy for wayland
set-option -gq copy-command 'wl-copy -n'


# Bindings
#unbind-key -a

bind-key -T prefix C-b send-prefix
bind-key -T prefix h list-keys -N
bind-key -T prefix H list-keys
bind-key -T prefix s source-file ~/.tmux.conf
bind-key -T prefix ? display-message
bind-key -T prefix : command-prompt
bind-key -T prefix q display-panes
bind-key -T prefix [ previous-window
bind-key -T prefix ] next-window
bind-key -T prefix \{ swap-pane -U
bind-key -T prefix \} swap-pane -D
bind-key -T prefix ( switch-client -p
bind-key -T prefix ) switch-client -n
bind-key -T prefix < display-menu -T "#[align=centre]#{window_index}:#{window_name}" -t = -x W -y W "#{?#{>:#{session_windows},1},,-}Swap Left" l { swap-window -t :-1 } "#{?#{>:#{session_windows},1},,-}Swap Right" r { swap-window -t :+1 } "#{?pane_marked_set,,-}Swap Marked" s { swap-window } '' Kill k { kill-window } Respawn x { respawn-window -k } "#{?pane_marked,Unmark,Mark}" m { select-pane -m } Rename r { command-prompt -F -I "#W" { rename-window -t "#{window_id}" "%%" } } '' "New After" C { new-window -a } "New At End" c { new-window }
bind-key -T prefix > display-menu -T "#[align=centre]#{pane_index} (#{pane_id})" -x P -y P "#{?#{m/r:(copy|view)-mode,#{pane_mode}},Go To Top,}" < { send-keys -X history-top } "#{?#{m/r:(copy|view)-mode,#{pane_mode}},Go To Bottom,}" > { send-keys -X history-bottom } '' "#{?mouse_word,Search For #[underscore]#{=/9/...:mouse_word},}" C-r { if-shell -F "#{?#{m/r:(copy|view)-mode,#{pane_mode}},0,1}" "copy-mode -t=" ; send-keys -X -t = search-backward "#{q:mouse_word}" } "#{?mouse_word,Type #[underscore]#{=/9/...:mouse_word},}" C-y { copy-mode -q ; send-keys -l "#{q:mouse_word}" } "#{?mouse_word,Copy #[underscore]#{=/9/...:mouse_word},}" c { copy-mode -q ; set-buffer "#{q:mouse_word}" } "#{?mouse_line,Copy Line,}" l { copy-mode -q ; set-buffer "#{q:mouse_line}" } '' "Horizontal Split" h { split-window -h } "Vertical Split" v { split-window -v } '' "#{?#{>:#{window_panes},1},,-}Swap Up" u { swap-pane -U } "#{?#{>:#{window_panes},1},,-}Swap Down" d { swap-pane -D } "#{?pane_marked_set,,-}Swap Marked" s { swap-pane } '' Kill X { kill-pane } Respawn R { respawn-pane -k } "#{?pane_marked,Unmark,Mark}" m { select-pane -m } "#{?#{>:#{window_panes},1},,-}#{?window_zoomed_flag,Unzoom,Zoom}" z { resize-pane -Z }
bind-key -T prefix y copy-mode
bind-key -T prefix p paste-buffer -p
bind-key -T prefix x respawn-window
bind-key -T prefix X respawn-pane
bind-key -T prefix l choose-tree -Zw
bind-key -T prefix L choose-tree -Zs
bind-key -T prefix c new-window
bind-key -T prefix C command-prompt { new-window -n "%%" }
bind-key -T prefix C-a command-prompt { new-window -an "%%" }
bind-key -T prefix C-b command-prompt { new-window -bn "%%" }
bind-key -T prefix r command-prompt -I "#W" { rename-window "%%" }
bind-key -T prefix R command-prompt -I "#S" { rename-session "%%" }
bind-key -T prefix k confirm-before -p "kill-window #W? (y/n)" kill-window
bind-key -T prefix K confirm-before -p "kill-session #S? (y/n)" kill-session
bind-key -T prefix C-k confirm-before -p "kill-server? (y/n)" kill-server
bind-key -T prefix w split-window -h -c "#{pane_current_path}"
bind-key -T prefix W split-window -v -c "#{pane_current_path}"
bind-key -T prefix m command-prompt -T target { move-window -t "%%" }
bind-key -T prefix M command-prompt -T target { move-pane -t "%%" }
bind-key -T prefix z resize-pane -Z
bind-key -T prefix 0 select-window -t :=0
bind-key -T prefix 1 select-window -t :=1
bind-key -T prefix 2 select-window -t :=2
bind-key -T prefix 3 select-window -t :=3
bind-key -T prefix 4 select-window -t :=4
bind-key -T prefix 5 select-window -t :=5
bind-key -T prefix 6 select-window -t :=6
bind-key -T prefix 7 select-window -t :=7
bind-key -T prefix 8 select-window -t :=8
bind-key -T prefix 9 select-window -t :=9

#MouseDown1 == rightclick
#MouseDown3 == Leftclick
bind-key -T root WheelUpPane if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode -e }
bind-key -T root MouseDown1Status select-window -t =
bind-key -T root MouseDown3Status display-menu -T "#[align=centre]#{window_index}:#{window_name}" -t = -x W -y W "#{?#{>:#{session_windows},1},,-}Swap Left" l { swap-window -t :-1 } "#{?#{>:#{session_windows},1},,-}Swap Right" r { swap-window -t :+1 } "#{?pane_marked_set,,-}Swap Marked" s { swap-window } '' Kill k { kill-window } Respawn x { respawn-window -k } "#{?pane_marked,Unmark,Mark}" m { select-pane -m } Rename r { command-prompt -F -I "#W" { rename-window -t "#{window_id}" "%%" } } '' "New After" C { new-window -a } "New At End" c { new-window }
bind-key -T root MouseDown1Pane select-pane -t = \; send-keys -M
bind-key -T root MouseDrag1Pane select-pane -t = \; if-shell -F "#{||:#{pane_in_mode},#{mouse_any_flag}}" { send-keys -M } { copy-mode }
bind-key -T root MouseDrag1Border resize-pane -M
#xorg
#bind-key -T root DoubleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
#bind-key -T root TripleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
#wayland
bind-key -T root DoubleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T root TripleClick1Pane select-pane -t = \; copy-mode -M \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"
bind-key -r -T prefix Up select-pane -U
bind-key -r -T prefix Down select-pane -D
bind-key -r -T prefix Left select-pane -L
bind-key -r -T prefix Right select-pane -R
bind-key -r -T prefix M-Up resize-pane -U 5 
bind-key -r -T prefix M-Down resize-pane -D 5
bind-key -r -T prefix M-Left resize-pane -L 5
bind-key -r -T prefix M-Right resize-pane -R 5

bind-key -T copy-mode C-c send-keys -X cancel
bind-key -T copy-mode / command-prompt -T search -p "(search down)" { send-keys -X search-forward "%%" }
bind-key -T copy-mode n send-keys -X search-again
bind-key -T copy-mode N send-keys -X search-reverse
bind-key -T copy-mode v send-keys -X begin-selection
bind-key -T copy-mode q send-keys -X cancel
bind-key -T copy-mode Escape send-keys -X clear-selection
bind-key -T copy-mode Space send-keys -X begin-selection
#xorg
#bind-key -T copy-mode Enter send-keys -X copy-pipe-and-cancel "xclip -selection clipboard -in"
#wayland
bind-key -T copy-mode Enter send-keys -X copy-pipe-and-cancel "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T copy-mode MouseUp1Pane send-keys -X clear-selection
bind-key -T copy-mode MouseDrag1Pane select-pane \; send-keys -X begin-selection
#xorg
#bind-key -T copy-mode MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in" 
#bind-key -T copy-mode DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in" 
#bind-key -T copy-mode TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
#bind-key -T copy-mode MouseDragEnd1Pane send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
#wayland
bind-key -T copy-mode MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p" 
bind-key -T copy-mode DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p" 
bind-key -T copy-mode TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T copy-mode MouseDragEnd1Pane send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T copy-mode WheelUpPane select-pane \; send-keys -X -N 1 scroll-up
bind-key -T copy-mode WheelDownPane select-pane \; send-keys -X -N 1 scroll-down

bind-key -T copy-mode-vi C-c send-keys -X cancel
bind-key -T copy-mode-vi / command-prompt -T search -p "(search down)" { send-keys -X search-forward "%%" }
bind-key -T copy-mode-vi n send-keys -X search-again
bind-key -T copy-mode-vi N send-keys -X search-reverse
bind-key -T copy-mode-vi v send-keys -X begin-selection
bind-key -T copy-mode-vi q send-keys -X cancel
bind-key -T copy-mode-vi Escape send-keys -X clear-selection
bind-key -T copy-mode-vi Space send-keys -X begin-selection
#xorg
#bind-key -T copy-mode-vi Enter send-keys -X copy-pipe-and-cancel "xclip -selection clipboard -in"
#wayland
bind-key -T copy-mode-vi Enter send-keys -X copy-pipe-and-cancel "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T copy-mode-vi MouseUp1Pane send-keys -X clear-selection
bind-key -T copy-mode-vi MouseDrag1Pane select-pane \; send-keys -X begin-selection
#xorg
#bind-key -T copy-mode-vi MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in" 
#bind-key -T copy-mode-vi DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in" 
#bind-key -T copy-mode-vi TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
#bind-key -T copy-mode-vi MouseDragEnd1Pane send-keys -X copy-pipe-no-clear "xclip -selection clipboard -in"
#wayland
bind-key -T copy-mode-vi MouseDragEnd1Pane select-pane \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p" 
bind-key -T copy-mode-vi DoubleClick1Pane select-pane \; send-keys -X select-word \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p" 
bind-key -T copy-mode-vi TripleClick1Pane select-pane \; send-keys -X select-line \; send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T copy-mode-vi MouseDragEnd1Pane send-keys -X copy-pipe-no-clear "wl-copy && wl-paste -n | wl-copy -p"
bind-key -T copy-mode-vi WheelUpPane select-pane \; send-keys -X -N 1 scroll-up
bind-key -T copy-mode-vi WheelDownPane select-pane \; send-keys -X -N 1 scroll-down
```

**references**

[^1]:https://github.com/ohmyzsh/ohmyzsh
[^2]:https://safjan.com/top-popular-zsh-plugins-on-github-2023/
[^3]:https://forum.manjaro.org/t/cmake-was-unable-to-find-a-build-program-corresponding-to-unix-makefiles/55525
[^4]:https://ostechnix.com/linux-commands-alternatives/
[^5]:https://github.com/ibraheemdev/modern-unix
