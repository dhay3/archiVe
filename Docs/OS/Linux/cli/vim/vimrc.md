# .vimrc

```
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Maintainer: 
"    Cyberpelican
" Version:
"    0.1
" 
" Common references
" github.com/amix/vimrc/blob/master/vimrcs/basic.vim
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => General Options
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""  
" common global options
set nocompatible
set encoding=utf-8
set confirm
set fileformats=unix,dos,mac

" disable annoying sound on error
set noerrorbells
set novisualbell
set t_vb=
set tm=500

" indention options
set autoindent
set smartindent
set smarttab
set tabstop=2
set expandtab
set shiftwidth=2
set shiftround

" search options
set hlsearch
set magic
set ignorecase
set incsearch
set smartcase
set showmatch
set matchtime=5

" command options
set cmdheight=1
set wildmenu

" diff options
" set diff
" set diffopt=filler,vertical

" display options
set title
set ruler
set number
set laststatus=2
set statusline=\ %{HasPaste()}%F%m%r%h\ %w\ \ CWD:\ %r%{getcwd()}%h\ \ \ Line:\ %l\ \ Column:\ %c
set scrolloff=7
syntax on
" colorscheme torte

" miscellaneous options
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
" => Vundle
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""  
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
Plugin 'VundleVim/Vundle.vim'
call vundle#end()


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Maps
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""  
" normal mode maps
map <F2> GoDate: <Esc>:read !date<CR>kJ
" map <space> /

" insert mode maps
imap <F2> <CR>Date: <Esc>:read !date<CR>kJa<CR>


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" => Plugins
" vimawesome.com
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""  
Plugin 'scrooloose/nerdtree'
Plugin 'vim-airline/vim-airline'
Plugin 'valloric/youcompleteme'
Plugin 'scrooloose/syntastic'
Plugin 'altercation/vim-colors-solarized'

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
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""  
let g:solarized_termcolors = 256
set background=dark
colorscheme solarized

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
```

==注意几点==

1. youcompleteme 需要对应语言的编译环境，所以需要提前安装好

2. syntastic 需要安装对应的 lint 才能生效

   https://github.com/vim-syntastic/syntastic/blob/master/doc/syntastic-checkers.txt