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
# Enviroment Virables
ZSH_THEME="half-life"
export ZSH="/home/0x00/.oh-my-zsh"
export EDITOR="/usr/bin/vim"
export UPDATE_ZSH_DAYS=30
export LANG=en_US.UTF-8
export FZF_BASE=/usr/share/fzf
eval "$(thefuck --alias)"
eval "$(zoxide init zsh)"
# Use control + g to activate navi
eval "$(navi widget zsh)"

# Built-in
plugins=(
    colored-man-pages
    extract
    fzf
    gh
    git
    sudo
)

# Extra
fpath+=/usr/share/zsh/plugins/zsh-completions/src
[ -f /usr/share/zsh/plugins/zsh-autopair/autopair.zsh ] && source /usr/share/zsh/plugins/zsh-autopair/autopair.zsh
[ -f /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh ] && source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
[ -f /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh ] && source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh
[ -f $ZSH/oh-my-zsh.sh ] && source $ZSH/oh-my-zsh.sh
autopair-init

# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/home/0x00/anaconda3/bin/conda' 'shell.zsh' 'hook' 2>/dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/home/0x00/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/home/0x00/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="/home/0x00/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

#fastfetch
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

# Bindings
bindkey '^U' backward-kill-line
bindkey '^P' history-search-backward
bindkey '^N' history-search-forward
bindkey '\e[A' history-search-backward
bindkey '\e[B' history-search-forward

#aliases
alias c='clear'
alias n='navi'
alias ls='lsd'
alias ll='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'
alias cp='rsync --progress -azvh'
alias mv='mv -v'
alias rm='trash-put -v'
alias du='dust'
alias df='duf'
alias ps='procs'
alias nc='ncat'
alias ip='ip -c=always'
alias top='btop'
alias cat='bat -pp'
alias less='bat'
alias more='bat'
alias grep='rg'
alias find='fd'
alias vbox='VirtualBox %U'
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
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'fabioluciano/tmux-tokyo-night'

run '~/.tmux/plugins/tpm/tpm'
```

**references**

[^1]:https://github.com/ohmyzsh/ohmyzsh
[^2]:https://safjan.com/top-popular-zsh-plugins-on-github-2023/
[^3]:https://forum.manjaro.org/t/cmake-was-unable-to-find-a-build-program-corresponding-to-unix-makefiles/55525
[^4]:https://ostechnix.com/linux-commands-alternatives/
[^5]:https://github.com/ibraheemdev/modern-unix
