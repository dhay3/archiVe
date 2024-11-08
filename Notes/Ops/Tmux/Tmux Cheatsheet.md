# Tmux Cheatsheet



> When the prefix key is pressed, tmux waits for another key press and that determines what tmux command is executed. Keys like this are shown here with a space between them: `C-b c` means first the prefix key `C-b` is pressed, then it is released and then the `c` key is pressed. 

Tmux 的指令需要通过 Prefix（默认 <kbd>ctrl</kbd> + <kbd>b</kbd>）和 Keys 来调用

1. Press <kbd>ctrl</kbd> + <kbd>b</kbd> first
2. Then press key to activate Tmux commands

> 这里使用默认 prefix

### Miscellaneous Related

| Name | Description                                               |
| ---- | --------------------------------------------------------- |
| ？   | 查看 Tmux bindings                                        |
| :    | 类似于 `vim` 中输入 `:` 可以输入类似 `tmux ls` 之类的指令 |
| d    | detach                                                    |

### Window/Panel Related

| Name                                   | Description                                                  |
| -------------------------------------- | ------------------------------------------------------------ |
| %                                      | splits window (into panels) horizontally                     |
| "                                      | splits window (into panels) vertically                       |
| .                                      | rename current session                                       |
| $                                      | rename current session                                       |
| c                                      | create a new window                                          |
| ,                                      | rename current window                                        |
| &                                      | kill current window                                          |
| 0..9                                   | 切换 window                                                  |
| s                                      | choose a session from a list                                 |
| w                                      | choose a window from a list                                  |
| q                                      | 显示 panel 的序号，可以和数字使用直接切换快速切换到制定 panel, 例如 q 2 |
| RightArrow/LeftArrow/UpArrow/DownArrow | 切换 panel                                                   |
| x                                      | kill the active panel                                        |
| space                                  | select layout for panels                                     |
| z                                      | zoom the active panel to full screen                         |

### Copy/paste related

| Name | Description                        |
| ---- | ---------------------------------- |
| [    | enter copy mode                    |
| ]    | paste the most recent paste buffer |

> tmux 默认只能复制到 tmux 的实例中，不能复制到系统的剪切板中
>
> 如果需要通过 `shift + ctrl + v` 的方式复制到系统的剪切板，需要设置 `set -g mode-mouse off`

**references**

[^1]: https://github.com/tmux/tmux/wiki/Getting-Started
[^2]: https://superuser.com/questions/266725/tmux-ctrlb-not-working
[^3]:https://dev.to/iggredible/the-easy-way-to-copy-text-in-tmux-319g
[^4]:https://stackoverflow.com/questions/17445100/getting-back-old-copy-paste-behaviour-in-tmux-with-mouse