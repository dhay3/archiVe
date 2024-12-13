# desktop file

参考：

https://www.maketecheasier.com/create-desktop-file-linux/

https://wiki.archlinux.org/title/desktop_entries

https://develop.kde.org/docs/desktop-file/

编写`.desktop`文件可以让应用是显示在application menu中。通常在`~/.local/share/applications`或`/usr/share/applications`

```
[Desktop Entry]

# The type as listed above
Type=Application

# The version of the desktop entry specification to which this file complies
Version=1.0

# The name of the application
Name=jMemorize

# A comment which can/will be used as a tooltip
Comment=Flash card based learning tool

# The path to the folder in which the executable is run
Path=/opt/jmemorise

# The executable of the application, possibly with arguments.
Exec=jmemorize

# The name of the icon that will be used to display this entry
Icon=jmemorize

# Describes whether this application needs to be run in a terminal or not
Terminal=false

# Describes the categories in which this entry should be shown
Categories=Education;Languages;Java;
```

注意需要修改文件的所有权例如`chown cpl:cpl typora.desktop`，且需要运行`update-desktop-database`

## 例子

typora desktop file

```
cpl in /usr/share/applications λ cat typora.desktop 
[Desktop Entry]
Name=Typora
Exec=/usr/bin/Typora
Icon=/usr/share/typora/resources/assets/app.ico
Terminal=false
Type=Application
Categories=Office
```

clion/goland desktop file

```
cpl in /usr/share/applications λ cat typora.desktop 
[Desktop Entry]
Name=CLion
Exec=/usr/local/bin/clion
Icon=/sharing/apps/clion-2022.2/bin/clion.svg
Terminal=false
Type=Application
Categories=Deve

cpl in /usr/share/applications λ cat goland.desktop 
[Desktop Entry]
Exec=/sharing/apps/GoLand-2021.1.2/bin/goland.sh
MimeType=text/plain;
Name=goland
Type=Application
Icon=/sharing/apps/GoLand-2021.1.2/bin/goland.svg
Terminal=false
Categories=Development
```

