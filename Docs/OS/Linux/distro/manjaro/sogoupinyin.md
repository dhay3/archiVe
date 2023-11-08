# Sogoupinyin

1. Download sogoupinyin and dependencies for AUR

   ```
   yay -Sy fcitx-qt4 fcitx-qt5 fcitx-qt6 fcitx-im kcm-fcitx fcitx-configtool gtk2 gtk3 
   ```
   
2. create a config file `~/.xprofile`

   ```
   export GTK_IM_MODULE=fcitx
   export QT_IM_MODULE=fcitx
   export XMODIFIERS=@im=fcitx
   ```

3. If not work use `fcitx-diagnose` to debug

## Jetbrain

Jetbrain 全家桶和 fcitx5 不兼容会出现输入框偏移（官网现在已经热修复）可以按照如下步骤修复

1. 双击 <kbd>shift</kbd> 调出菜单栏，输入 choose boot Java runtime

2. New 选择 Add Custom Rumtime(JRE 必须要和编译器使用的版本匹配否则会导致编译器不能打开，如果出现不能打开的情况请删除`~/.config/JetBrains/<product><version>/<product>.jdk` 后重启编译器)

3. 2021.1.3 之前的版本用这个 [JRE](https://github.com/RikudouPatrickstar/JetBrainsRuntime-for-Linux-x64/releases/download/202110301849/jbr-linux-x64-202110301849.zip) 不要使用官方给的 JRE

4. 替换掉 JRE 重启即可正常使用 fcitx

**references**

[^1]:https://wiki.archlinux.org/title/Fcitx
[^2]:https://youtrack.jetbrains.com/issue/JBR-2460/Wrong-position-of-input-window-and-no-input-preview-with-fcitx-and-ubuntu-13.04
[^3]:https://blog.csdn.net/lxyoucan/article/details/123289253
[^4]:https://stackoverflow.com/questions/72067909/intellij-idea-doesnt-start-on-ubuntu-after-change-the-boot-java-runtime

