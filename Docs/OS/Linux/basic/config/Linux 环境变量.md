# Linux 环境变量



> ==env显示所有环境变量==，check bash manual page shell variables for details

## built-in

内建系统环境变量

- PATH

  可以直接调用的命令所在目录。==会在shell打开时扫描目录（会执行set 和 source）==

- HOME

  当前用户的家目录

- HOSTNAME

  主机名

- MAIL

  当前用户的邮件目录

- USER

  当前登入的用户

- SHELL

  当前的SHELL

- LANG

  对没有设置值的`LC_*`变量的默认值

- ==LC_ALL==

  对LANG和所有`LC_*`变量覆写

  [LC_ALL=C](https://unix.stackexchange.com/questions/87745/what-does-lc-all-c-do)

- [LANGUAGE](https://wiki.archlinux.org/title/Locale#LANGUAGE:_fallback_locales)

  程序使用`gettext`时候才会考虑这个环境变量，只有当`LC_ALL`的值不是C的时候才生效。==如果fallback language有问题可以通过该变量设置==

- LC_TIME

  时间格式的locale catagory

## 自定义环境变量

### 临时

只在当前的Shell中生效，不会对其他Shell生效，关闭与Shell绑定的tty后就失效

```
[root@chz Desktop]# export name=chz
[root@chz Desktop]# echo $name
chz
```

### 持久

对所有用户，所有Shell生效

> 修改所有配置文件前先备份
>
> 使用source /etc/profile 让环境变量生效，==但是只会对使用了该命令的Shell生效，想要对所有的Shell生效需要reboot==

- 自定义**系统环境变量**，对所有用户生效，在`/etc/profile`中添加`export key=value`

  ```
  unset i
  unset -f pathmunge
  export abc='hello world'
  [root@chz Desktop]# source /etc/profile
  [root@chz Desktop]# echo $abc
  hello world
  [root@chz Desktop]# 
  ```

- 自定义**用户环境变量**，只对当前用户生效，不会对其他用生效

  ```
  # User specific environment and startup programs
  
  PATH=$PATH:$HOME/bin
  
  export PATH
  export golang='best'
  [root@chz ~]# source .bash_profile
  [root@chz ~]# echo $golang
  best
  [root@chz ~]# 
  ```



