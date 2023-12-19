# Tmux CLI

> 调用 `tmux` 默认会创建一个 Tmux session
>

- `new-session`

  等价于 `new`，新创建一个 Tmux session

  默认 session name 为 0，每创建一个 session 递增

  可以使用 `-s` 指定 session name

  ```
  $ tmux new -smysession
  ```

  除此外还可以使用 `-n` 指定第一个 window name 并运行指定的命令

  ```
  tmux new -nmytopwindow top
  ```

- `rename-session`

  等价于 `rename` 重名 session

  ```
  tmux rename -t0 mysession
  ```

- `rename-window`

  等价与 `renamew` 重名 window

  ```
  tmux renamw -t0 mywindow
  ```

- `new-window`

  需要在 session 内执行

  等价于 `neww`，新创建一个窗口

  可以在指令后面跟上命令，创建窗口的同时执行命令

  ```
  tmux neww top
  ```

- `list-keys`

  等价于 `lsk`，查看所有的 bindings

- `list-sessions`

  等价于 `ls`，查看 Tmux 当前所有的 sessionss

- `list-windows`

  等价于 `lsw`, 查看当前 session 中的所有 windows

- `kill-session`

  kill 当前 session，可以使用 `-t` 指定 session

  ```
  tmux kill-session -t 0
  ```

- `kill-window`

  kill 当前 window，可以使用 `-t` 指定 window

  ```
  tmux kill-window -t 0
  ```

- `kill-panel`

  kill 当前 panel

- `kill-server`

  退出所有的 sessions，即 kill 掉 tmux 进程

- `split-window`

  等价与 `splitw` 将当前 window 分割成 2 个 panel

  ```
  #horizontal split
  split-window -h
  #vertical split
  split-window -v
  ```

- `attach`

  Attaches to an existing session.

  https://github.com/tmux/tmux/wiki/Getting-Started#attaching-and-detaching

- `detach`

  Detaching from tmux means that the client exits and detaches from the outside terminal, returning to the shell and leaving the tmux session and any programs inside it running in the background. 

  https://github.com/tmux/tmux/wiki/Getting-Started#attaching-and-detaching