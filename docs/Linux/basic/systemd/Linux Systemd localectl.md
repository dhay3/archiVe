# Linux Systemd localectl

### localectl

查看和设置本地化(语言)

- `localectl`

  等价于`localectl status`，查看当前系统的本地化设置

  ```
  [root@chz systemd]# localectl 
     System Locale: LANG=en_US.UTF-8
         VC Keymap: us
        X11 Layout: us
  ```

- `localectl set-locale | set-keymap`

  设置显示语言或输入语言，结合`--list-locales`和`--list-keymap`一起使用

  ```
  [root@cyberpelican systemd]# localectl set-locale LANG=zh_CN.UTF-8
  [root@cyberpelican systemd]# localectl
     System Locale: LANG=zh_CN.UTF-8
         VC Keymap: us
        X11 Layout: us
        
  [root@cyberpelican systemd]# localectl set-keymap zh_CN
  [root@cyberpelican systemd]# localectl
     System Locale: LANG=zh_CN.UTF-8
         VC Keymap: zh_CN
        X11 Layout: us
  ```
