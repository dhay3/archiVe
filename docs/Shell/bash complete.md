# bash complete

https://blog.csdn.net/weixin_30781107/article/details/95461444

http://blog.lujun9972.win/blog/2018/03/13/%E5%A6%82%E4%BD%95%E7%BC%96%E5%86%99bash-completion-script/index.html

https://www.gnu.org/software/bash/manual/bash.html#Shell-Variables

https://www.gnu.org/software/bash/manual/bash.html#Programmable-Completion

https://github.com/scop/bash-completion/

> man bash /^\s+Completing
>
> help complete
>
> zsh 不支持使用`chsc -s /bin/bash` 切换

## 概述

某一些脚本没有命令行提示，如ossutil。用起来相当的不顺手，我们可以通过bash complete功能来时实现补全功能

pattern：`complete [option] command`

使用两次TAB补全，如果没有匹配的就不会补全

```
[root@cyberpelican \]# complete -W "him her my" a
[root@cyberpelican \]# a 
her  him  my   

```

还有一个compgen命令，不同的是可以通过`--`表示对指定的参数补全

> 这里`-- `(double dash and space)表示==所有的options==的结束，后面的参数做为文件名或其他参数
>
> https://devmanual.gentoo.org/tasks-reference/completion/index.html
>
> https://unix.stackexchange.com/questions/11376/what-does-double-dash-mean

```
[root@cyberpelican \]# compgen -W "apple bat cat"  -- appl 
apple
```

==complete后面必须跟name，但是compgen不需要(输出所有)==

```
[root@cyberpelican \]# complete -W "him her my" 
complete: usage: complete [-abcdefgjksuv] [-pr] [-DEI] [-o option] [-A action] [-G globpat] [-W wordlist]  [-F function] [-C command] [-X filterpat] [-P prefix] [-S suffix] [name ...]
[root@cyberpelican \]# compgen -W "him her my" 
him
her
my
```

## bash内置变量

> 只能被用于bash complete中的函数
>
> `$1`表示命令(command)，`$2`表示被补全的参数，`$3`表示补全后的参数，==函数内生效==

- `COMP_WORDS`

  数组，当前命令行中所有输入的参数以IFS分隔

- `COMPREPLY`

  数组，候选的补全结果

- `COMP_CWORD`

  整数，当前光标指向单词在COMP_WORDS中的索引。下标同样从0开始。

- `COMP_WORDBREAKS`

  字符串，单词之间的分隔符

- `COMP_LINE`

  字符串，当前命令行

```
#!/usr/bin/env bash
function _foo() {
  echo -e "\n"
  #打印出指定参数中所有的值
  declare -p COMP_WORDS
  declare -p COMP_CWORD
  declare -p COMP_LINE
  declare -p COMP_WORDBREAKS
}
#将函数_foo绑定到foo，交给complete process处理
complete -F _foo foo

---
#多一个IFS，COMP_CWORD就会不同
[root@chz opt]# foo a b c declare -a COMP_WORDS='([0]="foo" [1]="a" [2]="b" [3]="c" [4]="")'
declare -- COMP_CWORD="4"
declare -- COMP_LINE="foo a b c "
declare -- COMP_WORDBREAKS=" 	
\"'><=;|&(:"
^C
[root@chz opt]# foo a b cdeclare -a COMP_WORDS='([0]="foo" [1]="a" [2]="b" [3]="c")'
declare -- COMP_CWORD="3"
declare -- COMP_LINE="foo a b c"
declare -- COMP_WORDBREAKS=" 	
\"'><=;|&(:"
^C

#这里会等待使用ctrl + c结束
```

## 常用参数

> compgen于complete参数相同

- `-F function`

  将函数绑定到一个命令。当执行完毕后，`$1`表示命令(command)，`$2`表示被补全的参数，`$3`表示补全后的参数，==函数内生效==

  ```
  complete -F _foo foo
  ```

- `-W wordlist`

  补全候选词，使用IFS分隔

  ```
  [root@cyberpelican \]# complete -W "a b c" sc.sh
  [root@cyberpelican \]# sc.sh a 
  a  b  c  
  ```

- `-A action`

  用于指定改如何生成补全信息

  1. directory

     指定将要补全的信息是目录，等价于`-d`参数

     ```
     [root@cyberpelican \]# cat sc.sh 
       File: sc.sh
       #!/usr/bin/env bash
       _foo() {
           COMPREPLY=()
           cur="${COMP_WORDS[COMP_CWORD]}"
           COMPREPLY=($(compgen -d -- "$cur"))
       }
       complete  -F _foo foot
       
     [root@cyberpelican \]# foo 
     CVE-2021-3156-main  shell_test  
     ```
  
  2. file
  
     指定将要补全的信息是文件，等价于`-f`参数
  
     ```
     [root@cyberpelican \]# foo 
     compress_log          compress_log.sh       sc.sh                 ssh_payload.sh        testfile              
     compress_log_cron.sh  CVE-2021-3156-main    shell_test            ssh_payload.sh~.gz    
     [root@cyberpelican \]# cat sc.sh 
       File: sc.sh
       #!/usr/bin/env bash
       _foo() {
           COMPREPLY=()
           cur="${COMP_WORDS[COMP_CWORD]}"
           COMPREPLY=($(compgen -f -- "$cur"))
       }
       complete  -F _foo foot
     ```
  
  3. service
  
     指定将要补全的信息是服务，等价于`-s`参数
  
     ```
     [root@cyberpelican \]# foo 
     Display all 284 possibilities? (y or n)
     acr-nema          cmip-man          gds-db            klogin   
     ```
  
  4. signal
  
     指定将要补全的信息是posix signal
  
     ```
     [root@cyberpelican \]# foo 
     DEBUG        SIGCHLD      SIGJUNK(32)  SIGRTMAX     SIGRTMAX-2   SIGRTMAX-9   SIGRTMIN+14  SIGRTMIN+7   SIGTERM      SIGUSR2
     ERR          SIGCONT      SIGJUNK(33)
     [root@cyberpelican \]# cat sc.sh 
       File: sc.sh
       #!/usr/bin/env bash
       _foo() {
           COMPREPLY=()
           cur="${COMP_WORDS[COMP_CWORD]}"
           COMPREPLY=($(compgen -A signal -- "$cur"))
       }
       complete  -F _foo foot
     
     ```
  
- `-o comp-option`

  可以提供额外的补全信息

  > 常用
>
  > 与`-W`使用会冲突
>
  > `complete  -o filenames -o nospace -o bashdefault -F _foo foot`

  1. filename
  
     告诉readline应该按照文件名处理补全，会在补全中加slash
  
  2. nospace
  
     如果存在多个前缀相同的目录或文件，告诉readline不要在补全后添加空格，默认会添加空格
  
  3. bashdefault
  
     使用其余默认的readline设置

## 例子

```
#!/usr/bin/env bash
_foo(){
  local pre cur
  COMPREPLY=()
  pre=${COMP_WORDS[COMP_CWORD-1]}
  cur=${COMP_WORDS[COMP_CWORD]}
  opts="-h --help -f --file -o --output"
  if [ "$cur"=-* ]; then
  #这里括号表示初始化数组，而不是命令组
  #这里需要-- ,表示只对$cur做补全
      COMPREPLY=( $(compgen -W "$opts" -- "$cur") )
  fi
}
complete -F _foo foot
---
[root@cyberpelican \]# . sc.sh 
[root@cyberpelican \]# foo -
-f        --file    -h        --help    -o        --output  

```

## lnstall

同时需要注意一下安装目录，bash-completion会从三个地方加载自动补全脚本：

- /usr/share/bash-completion/completions/：有软件自行管理的自动补全脚本，最好不要改动
- /etc/bash_completion.d/：该目录下可自行添加所需要的自动补全脚本（全局生效）
- $HOME/.bash_completion：该文件也是自动补全脚本，不过只对当前用户成效，直接将编写的脚本往该文件追加即可。

注：由于这些补全命令都是使用shell语言编写的，所以可以直接添加到`/etc/bashrc`或`$HOME/.bashrc`文件中也是可以的，并不一定只能安装在上述几个地方中。