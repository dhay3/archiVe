# Shell timeout

timeout用于杀掉任务在指定时间内还在运行的命令，==如果命令已经结束timeout就会失效==，默认发送SIGTERM

syntax：`timeout [option] <duration> <command>`

duration可以是以s，m，h来代表分秒时，默认秒。

```
root in ~ λ timeout  1s top
```

## 参数

- `--preserver-status`

  ==保留命令退出的状态==

  ```
  root in ~ λ timeout 1 top;echo $?
  124
  root in ~ λ timeout --preserve-status 1 top;echo $?
  0
  ```

- `-k <duration>`

  如果在duration之后还没终止，发送SIGKILL，通常作用于SIGTERM失效

  ```
  root in ~ λ timeout -k 1 1 top
  ```

- `-s <POSIXSIG>`

  发送指定POSIXSIG

  ```
   ┌─────( root)─────(/opt) 
   └> $ timeout -s 9  1 top
  ```

  