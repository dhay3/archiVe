# Linux 修改系统语言

locale是Linux中的语言环境设置，控制系统在用户登录之前和登录之后显示的语言。

## locale(1)

当没有任何参数是`locale`显示当前系统设置的locale

```
cpl in ~ λ locale
LANG=en_US.UTF-8
LC_CTYPE="en_US.UTF-8"
LC_NUMERIC="en_US.UTF-8"
LC_TIME="en_US.UTF-8"
LC_COLLATE="en_US.UTF-8"
LC_MONETARY="en_US.UTF-8"
LC_MESSAGES="en_US.UTF-8"
LC_PAPER="en_US.UTF-8"
LC_NAME="en_US.UTF-8"
LC_ADDRESS="en_US.UTF-8"
LC_TELEPHONE="en_US.UTF-8"
LC_MEASUREMENT="en_US.UTF-8"
LC_IDENTIFICATION="en_US.UTF-8"
LC_ALL=en_US.UTF-8
```

使用`-a`参数打印出系统当前可以使用的locale

```
cpl in ~ λ locale -a
C
en_US.utf8
POSIX
zh_CN.utf8
```

`C`是电脑识别的locale，具体[参考](Linux 环境变量.md)

还可以参考`locale(5)`和`locale(7)`

## locale.conf

`/etc/loacle.conf`在系统的early boot by systemd阶段读取，可以被用户的个人配置覆写。可以被kernale command line options覆写，在`/etc/locale.conf`中可以使用的环境变量具体可以查看`man locale.conf`

> note
>
> `LC_ALL`不能配置在`/etc/locale.conf`中

## localectl

`localectl`可以在系统运行的过程中==修改locale和keymap(键盘输入布局)设置==

- status

  显示当前系统的locale和keyboard settings，当`localectl`没有任何参数时就等价`localectl status`

  ```
  cpl in ~ λ localectl 
     System Locale: LANG=en_US.UTF-8
                    LANGUAGE=en_US.UTF-8:en_US:en_US:en_US
         VC Keymap: us
        X11 Layout: us
         X11 Model: pc86
  ```

- `list-locales | set locales <locale>`

  显示系统可用的locale

- `list-keymaps | set-keymap <keymaps>`

  显示系统key用的keymaps