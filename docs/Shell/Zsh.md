# Zsh

参考:

https://linux.cn/article-10047-1.html

https://github.com/ohmyzsh/ohmyzsh

https://www.howtogeek.com/669835/how-to-change-your-default-shell-on-linux-with-chsh/

> 本基于Kali Linux

## Zsh特性

1. 智能补全，比bash快太多了

   使用双`tab`键可以，可以通过键盘上下左右来移动选择

   <img src="D:\asset\note\imgs\_Shell\Snipaste_2020-12-27_14-12-36.png" style="zoom:80%;" />

   ==输入部分路径同样可以补全==

   ```
   ╭─root@cyberpelican ~ 
   ╰─cd /v/w/h        
   TAB
   ╭─root@cyberpelican ~ 
   ╰─cd /var/www/html/
   ```

2. 错误提示码

   在输入错误的命令后，zsh会在下一条命令输出错误码127

   <img src="D:\asset\note\imgs\_Shell\Snipaste_2020-12-27_14-17-14.png" style="zoom:80%;" />

3. 支持插件

4. 智能别名

   ```
   alias -s {yml,yaml}=vim
   ```

   输入文件名中带有yml，yaml自动使用yaml代开编辑

   全局别名，可以在任意位置展开。

   ```
   alais -g G=' | grep -i'
   ...
   
   $ ls -l G do
   drwxr-xr-x.  5 rgerardi rgerardi 4096 Aug  7 14:08 Documents
   drwxr-xr-x.  6 rgerardi rgerardi 4096 Aug 24 14:51 Downloads
   ```

## 安装

> 如果出现chsh: PAM: Authentication failure，查看`/etc/passwd`中设置的shell是否错误。

1. `apt install zsh`

2. `chsh -s /usr/bin/zsh`

## 配置文件

http://zsh.sourceforge.net/Doc/Release/Files.html#Files

所有的个人文件在全局配置之后加载

- `/etc/zsh/zshenv`

  优先从该配置文件中读取，内容不能被覆盖

- `/etc/zsh/zprofile`

  如果是login shell从该文件读取

- `/etc/zsh/zshrc`

  如果是交互的shell(==不是指nologin-shell==)，从该文件读取

- `/etc/zsh/zlogout`

  如果退出login-shell，从该文件中读取



## 常用插件

使用插件需要安装`oh-my-zsh`

https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins

所有的插件都可以从`$ZSH/plugins/`中发现。使用插件需要下载相应的软件

- sublilme

- yum

- sudo，键入两次`esc`在命令之前添加`sudo`

- web-search，以`command-cli`的形式使用搜索引擎

- urltools，url编码与反编码

- systemadmin，优化和简便一些命令

  https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/systemadmin

