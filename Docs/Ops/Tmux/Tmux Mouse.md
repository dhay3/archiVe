# Tmux Mouse

在 Tmux 中也支持鼠标的操作

> 需要使用 `set -g mouse on` 开启

- Pressing the left button on a pane will make that pane the active pane.

  右键点击 panel，可以激活 panel

- Pressing the left button on a window name on the status line will make that the current window.

  左键点击 status line 中的 window，可以激活 window

- Dragging with the left button on a pane border resizes the pane.

  左键点击 panel boarder 可以调整 panel 大小

- Dragging with the left button inside a pane selects text; the selected text is copied when the mouse is released.

  选中文本等于使用 <kbd>ctrl</kbd> + <kbd>b</kbd> + <kbd>[</kbd> 进入复制模式并选中

- Pressing the right button on a pane opens a menu with various commands. When the mouse button is released, the selected command is run with the pane as target. Each menu item also has a key shortcut shown in brackets.

  右键会出现一个菜单供用户悬着



**references**

[^1]:https://github.com/tmux/tmux/wiki/Getting-Started#using-the-mouse