# Linux Systemd localectl

## Digest

syntax：`localectl [optinos] {command}`

Control the system locale and keyboard layout settings

实际修改机器上的`/etc/locale.conf`和`/etc/vconsole.conf`

## Postional args

- `status`

  显示当前locale 和 keyboard mapping，缺省值

  ```
  cpl in /etc/systemd λ localectl status
     System Locale: LANG=en_US.UTF-8
                    LANGUAGE=en_US.UTF-8:en_US:en_US:en_US
         VC Keymap: us
        X11 Layout: us
         X11 Model: pc86
  ```

- `list-keymaps`

  显示所有可用的keymap

- `list-locales`

  显示所有可用的locale

- `set-locale LOCALE,set-locale VAR=LOCALE`

  设置 locale，如果只是指定locale表示设置全局变量`LANG=locale`，具体查看man page

- `set-keymap MAP`

  set the system keyboard mapping for the conosle and X11

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