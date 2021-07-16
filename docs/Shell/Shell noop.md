# Shell noop

noop（:）等价与true命令

## 用法

1. when the shell syntax requires a command but you have nothing to do

   ```
   while keep_waiting;do
   	:
   done
   ```

2. dump a file’s data

   ```
   cpl in /tmp λ tail -4 ~/.zsh_history
   : 1626398420:0;tail -2 ~/.zsh_history
   : 1626398424:0;tail -4 ~/.zsh_history
   : 1626398432:0;clear
   : 1626398435:0;tail -4 ~/.zsh_history
   cpl in /tmp λ : > ~/.zsh_history 
   cpl in /tmp λ cat ~/.zsh_history 
   : 1626398469:0;cat ~/.zsh_history
   ```

   