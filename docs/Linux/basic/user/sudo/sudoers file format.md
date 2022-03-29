# sudoers file format

参考：

https://www.linux.com/topic/networking/how-wrestle-control-sudo-sudoers/

https://linux.die.net/man/5/sudoers

https://www.sudo.ws/man/1.8.18/sudoers.man.html

sudoer默认会去读取，`/etc/sudoers`和`/etc/sudo.d`下的文件

`/etc/sudoers`使用Backus-Naur Form(EBNF)编写，以如下格式定义

```
symbol ::= definition | alternate1 | alternate2 ...
```

由两部分组成，aliases (basically variables) and user specifications (which specify who may run what).

## wildcard

支持`?`,`*`,`+`通配符

```
 #operator组可以在任意IP下
 %operator ALL = /bin/cat /var/log/messages*
```

==`ALL`是保留关键字，在不同位置表示匹配所有==

```
ubuntu ALL=(ALL:ALL) ALL
从左自右表示所有IP，所有用户，所有命令
```

`！`同样表示取非，如下表示所有用户但是除了root用户

```
ALL,!root
```

## Aliases

每一行由如下格式定义：

```
Alias_Type NAME = item1, item2, ...
```

有Aliases_Type，分别为 User_Alias, Runas_Alias, Host_Alias and Cmnd_Alias.

NAME必须大写，如果Aliases_Type相同，可以通过colon(:)联合

```
Alias_Type NAME = item1, item2, item3 : NAME = item4, item5
```

### User_Alias

由user names，uids(`#`开头)，group names(`%`开头)，gid(`%#`开头)等组成。在之后可以被调用。Runas_Alias和User_Alias差不多

```
User_Alias USER1 = tz,#1,%#1000 : USER2 = ubuntu
```

### Host_Alias

由host name，ip_addr，network/netmask等组成

```
Host_Alias HOST = 192.168.80.200,192.168.80.0/24
```

### Cmnd_Alias

由command names，directories等组成，如果是directories，用户可以使用这个目录下的所有文件，但是不能使用子目录下的文件。

```
Cmnd_Alias POWEROFF=/sbin/poweroff,/sbin/shutdown,/sbin/reboot
```

## User specification

https://www.huaweicloud.com/articles/31f8e1c3207c105437293e7bdb72249a.html

user specification 表示指定用户可以在指定Host上可以通过sudo运行什么命令，由如下结构组成。如果没有`(as_whom:as_whichgroup)`默认root

```
who which_host = (as_whom:as_whichgroup) what_cmd
```

例如：

```
tz  ALL=(root) /bin/ls
```

表示ubuntu 可以所有host上可以扮演root用户使用`ls`命令，但是不允许使用其他命令

```
#可以通过sudo -l 来查看当前用户可以调用的所有命令
tz@win2k:/home/ubuntu$ sudo -l
[sudo] password for tz:
Matching Defaults entries for tz on win2k:
    env_reset, mail_badpass, secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin\:/snap/bin

User tz may run the following commands on win2k:
    (root) /bin/ls

tz@win2k:/home/ubuntu$ pwd
/home/ubuntu
tz@win2k:/home/ubuntu$ sudo pwd
Sorry, user tz is not allowed to execute '/bin/pwd' as root on win2k.
tz@win2k:/home/ubuntu$ sudo ls
5017872_cyberpelican.link_nginx.zip  dtstack_agent.sh  hello-world_29.assert  hello-world_29.snap
```

也可以和Cmnd_Alias一起使用

```
Cmnd_Alias  SHELLS = /bin/sh,/bin/ash,/bin/bsh,/bin/bash,/bin/csh, \
            /bin/ksh,/bin/rsh,/bin/tcsh,/bin/zsh

%tbops  ALL=(ALL) ALL,SUPERMIT,!SHELLS,SUDOEXEC
```

表示可以使用ALL，但是除SHELLS Cmnd_Alias外

### Tag_Spec

用于为command设置标签

- ==NOPASSWD and PASSWD==

  ```
  ray     rushmore = NOPASSWD: /bin/kill, PASSWD: /bin/ls, /usr/bin/lprm
  ```

  ray在rushmore上调用`sudo kill`不需要密码，但是`ls`和`lprm`需要

具体查看`tag_spec`部分

## sudoers options

sudoers options以如下形式定义：

```
Defaults        always_set_home
Defaults        env_reset
```

同样可以使用`！`来表示关闭option

```
Defaults 		!pwfeedback
```

### boolean flags

该类别无值，显示表明就是on，可以使用`Defaults !options`来取消默认值

- always_set_home：设置HOME环境变量，默认off

- authenticate：每次调用命令前需要输入密码认证，但是可以被PASSWD和NOPASSWD覆写，默认on

- compress_io：使用zlib对I/O日志压缩，如果支出zlib默认为on

- env_editor：visudo使用的编译器，需要设置SUDO_EDITOR，VISUAL，EDITOR环境变量。默认为on

- env_reset：sudo只保留(TERM，PATH，HOME，MAIL，SHELL，LOGNAM，SUDO_*，USER)作为环境变量，如果设置的secure_path则使用该值作为PATH，==默认启用==。可以使用`printenv variblae`来查看，==使用`echo $varible`只是读取到当前shell中变量，而不是sudo扮演用户的变量==

  ```
  Defaults	！env_rest
  ➜  sudoers.d sudo printenv http_proxy
  socks5://127.0.0.1:1089
  ➜  sudoers.d sudo printenv USER     
  root
  ```

  https://superuser.com/questions/232231/how-do-i-make-sudo-preserve-my-environment-variables

- fast_glob：快速模式扩展。==相对路劲中不能使用模式扩展，例如`cd ../etc/`，可以防止路径遍历==，默认关闭

- insults：用户密码输入错误时提示

- mail_always：使用sudo命令时，发送邮件

- mail_badpass：使用sudo密码校验错误时，发送邮件

- mail_no_host：用户在非允许的host上运行sudo，发送邮件

- path_info：如果没有PATH环境变量，使用这个指定的

- pwfeedback：输入密码时显示`*`，而不是空白

- requiretty：sudo只能在login session被使用

- rootpw：密码校验时，通过root的密码校验而不是当前用户的密码

- runas_default：sudo默认调用的用户

- shell_noargs：如果设置了该值，使用sudo时必须使用`-s`指定使用的shell

- targetpw：如果要求输入的密码是当前的用户，可以不输入密码；如果是通过`-u`指定的用户，需要输入密码

  ```
  #在ubuntu调用sudo时需要输入root的密码而不是自己的密码
  ubuntu  ALL=(root)  /bin/su
  #在cowrie调用sudo时不需要输入密码, 因为ALL中包含cowrie用户
  cowrie ALL=(ALL:ALL) ALL
  ```

- sudoedit_follow

  sudoedit可以修改连接文件对应的我文件

### integers

该类别需要指定一个整数

- command_timeout：命令允许运行的最长时间
- passwd_tries：尝试密码输入的最大次数

- passwd_timeout：sudo password prompt timeout，mins，0表示没有时限，默认5mins
- ==timestamp_timeout==：elapse before **sudo** will ask for a passwd again，默认4mins，0表示每次使用sudo时都必须提供密码
- umask ：umask使用的值

### string

该类别需要指定一个自负床

- badpass_message：登入错误后显示的错误信息
- editor：sudoedit使用的编辑器，==只有在没有设置EDITOR的环境变量时才生效==，可以通过设置环境变量
- env_keep：env_reset开启时的变量
- env_reset：env_reset开启时需要移除的变量

- passprompt：提示的prompt格式，默认`[sudo] password for %p`
- runas_default：默认扮演的角色，默认使用root
- sudoers_locale：sudo使用的locale设置，默认`C`
- env_file：包含sudo使用的环境变量，自会添加不在env_keep和env_reset中的变量
- logfile：sudo log存储的位置，默认由syslog管理
- secure_path：$PATH的值

## sudo and environment

sudo默认使用`env_reset`即最小的环境变量集，例如在当前用户设置了`http_proxy`，但是使用了`sudo pacman -Sy`是读取不到`http_proxy`的。但是可以通过`evn_keep`可以从invoking user’s environment中继承过来。如果关闭了`env_reset`，并在`env_delete`中没有指明，那么会从invoking user’s enviroment中继承所有的环境变量

## 例子

```
#boolean options
Defaults        !always_set_home 
Defaults        compress_io 
Defaults        fast_glob 
Defaults        mail_no_host
Defaults        pwfeedback
Defaults        targetpw
Defaults        sudoedit_follow
Defaults		!reset_env

#integer options
Defaults        passwd_timeout=1
Defaults        passwd_tries=2
Defaults        timestamp_timeout=0

#string options
Defaults        passprompt="password for [%h@%p]"
Defaults        editor="/usr/bin/vim"
Defaults        secure_path="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin"

#cpl    ALL=(ALL:ALL) NOPASSWD:ALL
cpl     ALL=(ALL:ALL) ALL
```
