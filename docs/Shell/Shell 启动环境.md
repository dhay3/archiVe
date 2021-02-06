# Shell 启动环境

> ==所有Shell共享env变量==
>
> 使用`echo $0`判断当前的shell是不是 login shell
>
> ```
> root in /\ λ echo $0
> zsh       #non-login shell
> 
> root in /etc/security/limits.d λ echo $0
> -zsh  #login shell, - 代表这是一个login shell
> ```
>
> ==登入non-login shell不会执行login shell，所以在login-shell中设置的环境变量不会对non-login生效==

## Session

用户每次使用 Shell，都会开启一个与 Shell 的 Session（对话）。

Session 有两种类型：登录 Session 和非登录 Session，也可以叫做 login shell 和 non-login shell。

## login shell

login shell 是用户登入系统以后，系统为用户开启的原始session，通常需要用户输入用户名和密码进行登入

启动的初始化脚本依次如下。

1. `/etc/profile`：所有用户的全局配置脚本。
2. `/etc/profile.d`目录里面所有`.sh`文件
3. `~/.bash_profile`：用户的个人配置脚本。如果该脚本存在，则执行完就不再往下执行。
4. `~/.bash_login`：如果`~/.bash_profile`没找到，则尝试执行这个脚本（C shell 的初始化脚本）。如果该脚本存在，则执行完就不再往下执行。
5. `~/.profile`：如果`~/.bash_profile`和`~/.bash_login`都没找到，则尝试读取这个脚本（Bourne shell 和 Korn shell 的初始化脚本）。

Linux 发行版更新的时候，会更新`/etc`里面的文件，比如`/etc/profile`，因此不要直接修改这个文件。如果想修改所有用户的登陆环境，就在`/etc/profile.d`目录里面新建`.sh`脚本。

==如果想修改你个人的登录环境，一般是写在`~/.bash_profile`里面。==

下面是一个典型的`.bash_profile`文件。

```
# .bash_profile

# Get the aliases and functions
if [ -f ~/.bashrc ]; then
        . ~/.bashrc
fi

# User specific environment and startup programs

PATH=$PATH:$HOME/bin

export PATH
```

我们可以使用`bash --login`来强制执行login shell需要执行的脚本（执行login shell的流程）

## non-login shell

non-login shell是用户进入系统以后，==手动新建的 Session，这时不会执行login shell==

non-login shell的初始化脚本如下

1. `/etc/bash.bashrc`：对全体用户有效。（==注意不同distribution内容不同==）
2. `~/.bashrc`：仅对当前用户有效。

每次新建一个 Bash 窗口，就相当于新建一个非登录 Session，所以`~/.bashrc`每次都会执行。===注意，执行脚本相当于新建一个非互动的 Bash 环境，但是这种情况不会调用`~/.bashrc`。==

## bash_logout

`~/.bash_logout`脚本在每次退出 Session 时执行，通常用来做一些清理工作和记录工作，比如删除临时文件，记录用户在本次 Session 花费的时间。

如果没有退出时要执行的命令，这个文件也可以不存在。
