# Touchpad

## 0x01 Overview

在介绍 Touchpad 的使用前，需要了解一下 Linux 是如何接受 input device 信号的。

## 0x02 libinput

在 Linux 中通过 libinput 来处理因用户对 input devices（mice,keyboards,touchpads,touchscreens，不包括 joystick） 的动作而产生的信号

主要有如下一些动作

1. click
2. gestrue

### Click

![_images/software-buttons-visualized.svg](https://wayland.freedesktop.org/libinput/doc/latest/_images/software-buttons-visualized.svg)

Left, right and middle button events can be triggered as follows:

- if a finger is in the main area or the left button area, a click generates left button events.
- if a finger is in the right area, a click generates right button events.
- if a finger is in the middle area, a click generates middle button events.

更加推荐打开 tab-to-click, 使用 tab 来触发 click 的动作

### Gestures

Gestures 分为几种

1. pinch
2. swipe
3. hold

#### Pinch gesture

Pinch gestures are executed when two or more fingers are located on the touchpad and are either changing the relative distance to each other (pinching) or are changing the relative angle (rotate). 

2 根以上手指相对的移动或者旋转

![_images/pinch-gestures.svg](https://wayland.freedesktop.org/libinput/doc/latest/_images/pinch-gestures.svg)

#### Swipe gesture

Swipe gestures are executed when three or more fingers are moved synchronously in the same direction

3 根以上手指向同一方向移动称为 Swipe

![_images/swipe-gestures.svg](https://wayland.freedesktop.org/libinput/doc/latest/_images/swipe-gestures.svg)

#### Hold gesture

A hold gesture is one where the user places one or more fingers on the device without significant movement.

手指不做任何动作

主要被用在

1. where a two-finger scrolling starts kinetic scrolling in the caller, a subsequent hold gesture can be used to stop that kinetic scroll motion, and

   向下快速滚动时，使用 hold gesture 可以立马停止滚动

2. hold-to-trigger interactions where the interaction could be a click, a context menu, or some other context-specific interaction.

   必须要 hold 才能触发，比如选中文本

## 0x03 libinput-gestures

> libinput-gestures 配置繁琐，如果不是使用 wayland 更推荐使用 touchegg 和 touche 傻瓜式配置

libinput-gestures 是一个客制化 gestures 的工具

注意：

==libinput-gestures 并不能 override KDE 默认的 gestures[^8]==

所以不要尝试在 KDE 原有的 gestures 上自定义 gestures，会导致不必要的预期出现

未来可能在 KDE6 中可以自定义配置 gestures 到时候也就不需要 libinput-gestures

### Installation

1. 安装 libinput-gestures (不建议直接从 AUR 中下载)

2. 需要 python >= 3.6

3. 需要 libinput >= 1.0

4. 用户需要在 input group 中

   ```
   sudo gpasswd -a $USER input
   ```

5. 重启

6. 安装 command 依赖

   > ydotool 需要启动 ydotoold 才可以使用，ydotoold 可以通过 `systemctl --user enable --now ydotool` 来启动

   ```
   #for xorg
   xdotool
   #for xorg and wayland
   ydotool
   #for desktop selection under xorg and wayland
   wmctrl
   ```

7. 安装 gestures 图形化配置工具

8. 配置开启自动启动 libinput-gestures

   ```
   libinput-gestures-setup autostart start
   ```

### libinput-gestures.conf

libinput-gestures 默认会读取 `/etc/libinput-gestures.conf` 中的内容，用于 mapping gestures to touchpad

如果想要设置自己的 mappings，可以在 `~/.config/libinput-gestures.conf` 修改，或者使用图形化工具 gestures 来配置。如果配置修改需要使用 `libinput-gestures-setup restart` 来重载配置

guesture line 由 3 或者是 4 部分组成

```
action motion [finger_count] command
```

#### action motion

其中 action motion 可以 

- swipe up
- swipe down
- swipe left
- swipe right
- swipe left_up
- swipe left_down
- swipe right_down
- pinch in
- pinch out
- pich clockwise
- pinch anticlockwise
- hold on

#### finger_count

finger_count 是一个可选的数字，表示特定的 command 会在指定的 finger_count 下的 action motion 生效

如果 finger_count 为空表示对任何数量的 finger_count 都生效

#### command

表示实际执行的指令，==可以是任意的命令==

1. _internal(for wayland and xorg) 调用 wmctrl,只有在 Desktop selection 中有效

   ```
   _internal ws_up
   _internal ws_down
   ```

2. xdotool (not for wayland)

3. ydotool (for wayland and xorg)

   ```
   ydotool key 125:1 31:1 125:0 31:0
   ydotool key 125:1 32:1 125:0 32:0
   ```

4. 如果使用 KDE 还可以使用 qdbus (推荐，如果使用  ydotool 或者是 xdotool key 会导致在 freerdp 中直接输出对应的字符)

   可以使用 `qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.shortcutNames` 获取所有的 shortcuts

   然后使用 `qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "ShortcutName"` 来调用对应的 shortcut

   或者使用 qdbusviewer 来查看所有可以使用的接口

#### ~/.config/libinput-gestures.conf

> 具体配置是否启用，按照实际选择
>
> 例如 wayland 中 swipe left/right 3, swipe up/down 4 冲突
>
> `gestures` 目前不支持图形化配置 hold on，但是可以通过 `lib-gestures` 的配置文件来配置

```
# Gestures
gesture swipe up 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Show Desktop"
gesture swipe down 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Overview"
#D: gesture swipe left 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Switch to Next Desktop"
#D: gesture swipe right 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Switch to Previous Desktop"
#D: gesture swipe left_up 3 qdbus org.kde.ActivityManager /ActivityManager/Activities org.kde.ActivityManager.Activities.PreviousActivity
#D: gesture swipe left_down 3 qdbus org.kde.ActivityManager /ActivityManager/Activities org.kde.ActivityManager.Activities.NextActivity
#D: gesture swipe right_down 3 qdbus org.kde.ActivityManager /ActivityManager/Activities org.kde.ActivityManager.Activities.NextActivity
#D: gesture swipe right_up 3 qdbus org.kde.ActivityManager /ActivityManager/Activities org.kde.ActivityManager.Activities.PreviousActivity
#D: gesture swipe up 4 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "ShowDesktopGrid"
#D: gesture swipe down 4 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "ExposeAll"
#D: gesture swipe left 4 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Switch to Next Desktop"
#D: gesture swipe right 4 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Switch to Previous Desktop"
gesture pinch out 3 qdbus org.kkcalcde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Window Maximize"
gesture pinch in 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "Window Maximize"
gesture pinch clockwise 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "view_zoom_in"
gesture pinch anticlockwise 3 qdbus org.kde.kglobalaccel /component/kwin org.kde.kglobalaccel.Component.invokeShortcut "view_zoom_out"
#D: gesture swipe left 4 _internal ws_up
#D: gesture swipe right 4 _internal ws_down
gesture hold on 3 qdbus org.kde.lattedock /Latte org.kde.LatteDock.activateLauncherMenu
gesture hold on 4 qdbus org.kde.yakuake /yakuake/window org.kde.yakuake.toggleWindowState
```

## 0x04 touchegg

> 不支持 wayland

文档详细[^4]



**references**

[^1]:https://github.com/bulletmark/libinput-gestures?tab=readme-ov-file
[^2]:https://wayland.freedesktop.org/libinput/doc/latest/what-is-libinput.html
[^3]:https://wiki.archlinux.org/title/libinput
[^4]:https://github.com/JoseExposito/touchegg?tab=readme-ov-file
[^5]:https://www.reddit.com/r/kde/comments/168dsjh/wayland_change_touchpad_gestures/
[^6]:https://github.com/bulletmark/libinput-gestures/issues/342
[^7]:https://www.reddit.com/r/archlinux/comments/142kbrk/ydotoold_background_process/
[^8]:https://github.com/bulletmark/libinput-gestures/issues/355
[^9]:https://bugs.kde.org/show_bug.cgi?id=402857
[^10]:https://www.reddit.com/r/kde/comments/7pftu0/terminal_commands_to_access_kwin_effects/
[^11]:https://www.reddit.com/r/kde/comments/n9urdn/kwin_qdbus_commands_defining_touch_gesture_for/