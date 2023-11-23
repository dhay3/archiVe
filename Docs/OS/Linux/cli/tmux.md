# Tmux

## Tmers

https://github.com/tmux/tmux/wiki/Getting-Started#sessions-windows-and-panes
$$
terminal \in panel \in window \in session
$$


## Commands

调用 `tmux` 默认会创建一个 Tmux session

- `new-window`

  等价于 `neww`，新创建一个窗口

- `new-session`

  等价于 `new`，新创建一个 Tmux session

- `list-keys`

  等价于 `lsk`，查看所有的 bindings

- `q`

  退出当前页面

- `attach`

  https://github.com/tmux/tmux/wiki/Getting-Started#attaching-and-detaching

- `detach`

  https://github.com/tmux/tmux/wiki/Getting-Started#attaching-and-detaching

- `list-sessions`

  等价于 `ls`，查看 Tmux 当前所有的 sessionss

- `kill-sessions`

  kill 掉指定 session

## Status line

![img](https://github.com/tmux/tmux/wiki/images/tmux_status_line_diagram.png)

## Cheatsheet

> When the prefix key is pressed, tmux waits for another key press and that determines what tmux command is executed. Keys like this are shown here with a space between them: `C-b c` means first the prefix key `C-b` is pressed, then it is released and then the `c` key is pressed. 

Tmux 的指令需要通过 Prefix（默认 <kbd>ctrl</kbd> + <kbd>b</kbd>）和 Keys 来调用

1. Press <kbd>ctrl</kbd> + <kbd>b</kbd> first
2. Then press key to activate Tmux commands

> 这里使用默认 prefix

### Miscellaneous Related

| Name | Description             |
| ---- | ----------------------- |
| ？   | 查看 Tmux bindings      |
| :    | 类似于 `vim` 中输入 `:` |
| d    | detach                  |

### Window/Panel Related

| Name                                   | Description                                                  |
| -------------------------------------- | ------------------------------------------------------------ |
| %                                      | splits panel horizontally                                    |
| "                                      | splits panel vertically                                      |
| RightArrow/LeftArrow/UpArrow/DownArrow | 切换 panel                                                   |
| q                                      | 显示 panel 的序号，可以和数字使用直接切换快速切换到制定 panel, 例如 q 2 |
| 0..9                                   | 切换 window                                                  |
| c                                      | 创建一个新的 window                                          |
| x                                      | 关闭当前 panel                                               |
| space                                  | 调整 panel 的位置，为了方便复制                              |
| z                                      |                                                              |

## ~/.tmux.conf

```
set -g default-terminal "screen-256color"
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'fabioluciano/tmux-tokyo-night'

run '~/.tmux/plugins/tpm/tpm'
```

## Trouble shooting

- ['tmux-256color': unknown terminal type.](https://unix.stackexchange.com/questions/1045/getting-256-colors-to-work-in-tmux)

**references**

[^1]:https://github.com/tmux/tmux/wiki/Getting-Started
[^2]:https://superuser.com/questions/266725/tmux-ctrlb-not-working
[^3]:https://github.com/fabioluciano/tmux-tokyo-night#available-configurations
[^4]:https://github.com/tmux-plugins/tpm