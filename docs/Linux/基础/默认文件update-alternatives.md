# update-alternatives 默认文件

## 概述

用于修改，移除，和创建Debian默认的软件。

> 系统的默认alternatives 存储在`/var/lib/dpkg/alternatives/`
>
> 可选的alternatives 存储在`/etc/alternatives/`

### terms

```
[root@chz alternatives]# ll /var/lib/dpkg/alternatives/ | grep terminal
-rw-r--r-- 1 root root   202 Aug 31 09:05 x-terminal-emulator
...
[root@chz alternatives]# pwd
/etc/alternatives

[root@chz alternatives]# ll | grep terminal
lrwxrwxrwx 1 root root  18 Aug 31 09:05 x-terminal-emulator -> /usr/bin/qterminal

[root@chz alternatives]# update-alternatives --config x-terminal-emulator
There are 2 choices for the alternative x-terminal-emulator (providing /usr/bin/x-terminal-emulator).

  Selection    Path                            Priority   Status
------------------------------------------------------------
* 0            /usr/bin/qterminal               40        auto mode
  1            /usr/bin/mate-terminal.wrapper   30        manual mode
  2            /usr/bin/qterminal               40        manual mode

Press <enter> to keep the current choice[*], or type selection number: 

```

- genenric name 

  通用的名字，通过alternatives system 指向真正的软件。存储在`/var/lib/dpkg/alternatives`，上述` x-terminal-emulator`

- alternative name

  在alternative directory下的符号链接。上述`x-terminal-emulator`

- alternative path | alternative

  alternative 对应的路径。上述`/usr/bin/qterminal`

- alternative directory

  存放所有的alternative 链接，默认`/etc/alternatives`

- administrative directory

  存放默认的alternative，默认`/var/lib/dpkg/alternatives`

- link group

  一组由联系的链接。上述`--config `显示的内容，每一个选项都是一个link group

- master link

  决定 link group 中其他的link 怎么配置

- slave link

  被link group 决定的link set

- automatic mode

  如果一个link group 被设为 automatic mode，alternative system会自动选择优先级最高的alternative

- manual mode

  处于manual mode ，==不会对系统产生影响==

> 如果想要修改默认配置，需要通过`--auto` 参数将link group 置为automatic mode 

## 参数

> 推荐统一使用 `--query`来查询

- `--install <link><name><path><priority>`

  安装一组alternatives，指定`/usr/bin/python3`
  
  ```
  root in /usr/bin λ update-alternatives --install /usr/bin/python3  python3 /usr/bin/python3.6 1
  update-alternatives: using /usr/bin/python3.6 to provide /usr/bin/python3 (python3) in auto mode
  root in /usr/bin λ update-alternatives --install /usr/bin/python3  python3 /usr/bin/python3.7 2
  update-alternatives: using /usr/bin/python3.7 to provide /usr/bin/python3 (python3) in auto mode
  ```
  
- `--display <alternative name>`

  ```
  [root@chz alternatives]# update-alternatives --display x-terminal-emulator
  x-terminal-emulator - auto mode
    link best version is /usr/bin/gnome-terminal.wrapper
    link currently points to /usr/bin/gnome-terminal.wrapper
    link x-terminal-emulator is /usr/bin/x-terminal-emulator
    slave x-terminal-emulator.1.gz is /usr/share/man/man1/x-terminal-emulator.1.gz
  /usr/bin/gnome-terminal.wrapper - priority 40
    slave x-terminal-emulator.1.gz: /usr/share/man/man1/gnome-terminal.1.gz
  /usr/bin/mate-terminal.wrapper - priority 30
    slave x-terminal-emulator.1.gz: /usr/share/man/man1/mate-terminal.1.gz
  /usr/bin/qterminal - priority 40
  ```

- `--list <alternative name>`

  展示指定alternative name 使用的 link group

  ```
  [root@chz alternatives]# update-alternatives --list x-terminal-emulator
  /usr/bin/mate-terminal.wrapper
  /usr/bin/qterminal
  ```

- `--query <alternative name>`

  显示alternative name 中link group 的具体内容

  ```
  [root@chz alternatives]# update-alternatives --query x-terminal-emulator
  Name: x-terminal-emulator
  Link: /usr/bin/x-terminal-emulator
  Slaves:
   x-terminal-emulator.1.gz /usr/share/man/man1/x-terminal-emulator.1.gz
  Status: auto
  Best: /usr/bin/qterminal
  Value: /usr/bin/qterminal
  
  Alternative: /usr/bin/mate-terminal.wrapper
  Priority: 30
  Slaves:
   x-terminal-emulator.1.gz /usr/share/man/man1/mate-terminal.1.gz
  
  Alternative: /usr/bin/qterminal
  Priority: 40
  Slaves:
  ```

  1. Name：alternative name
  2. Link：generic name of alternative
  3. status：auto or manual
  4. value：selected alternative
  5. alternative：path of this block's alternative

- `--config <alternative name>`

  配置可选的alternatives

  ```
  [root@chz alternatives]# update-alternatives --config x-terminal-emulator
  There are 2 choices for the alternative x-terminal-emulator (providing /usr/bin/x-terminal-emulator).
  
    Selection    Path                            Priority   Status
  ------------------------------------------------------------
    0            /usr/bin/qterminal               40        auto mode
  * 1            /usr/bin/mate-terminal.wrapper   30        manual mode
    2            /usr/bin/qterminal               40        manual mode
  
  Press <enter> to keep the current choice[*], or type selection number: 
  ```

- `--auto <alternative name>`

  将指定的alternative的 mode转automatic

  ```
  [root@chz alternatives]# update-alternatives --auto x-terminal-emulator 
  ```

## 替换默认terminal

> 如果无效请参考：
>
> https://stackoverflow.com/questions/16808231/how-do-i-set-default-terminal-to-terminator
>
> then reboot :smile:

1. 下载`apt install gnome-terminal*`

2. 配置默认terminal 为 gnome-terminal

   ```
   [root@chz alternatives]# update-alternatives --config x-terminal-emulator
   There are 3 choices for the alternative x-terminal-emulator (providing /usr/bin/x-terminal-emulator).
   
     Selection    Path                             Priority   Status
   ------------------------------------------------------------
     0            /usr/bin/gnome-terminal.wrapper   40        auto mode
     1            /usr/bin/gnome-terminal.wrapper   40        manual mode
   * 2            /usr/bin/mate-terminal.wrapper    30        manual mode
     3            /usr/bin/qterminal                40        manual mode
   
   Press <enter> to keep the current choice[*], or type selection number: 0
   update-alternatives: using /usr/bin/gnome-terminal.wrapper to provide /usr/bin/x-terminal-emulator (x-terminal-emulator) in manual mode
   ```

   上述`gnome-terminal.wrapper`已经处于auto mode 所以无需修改



