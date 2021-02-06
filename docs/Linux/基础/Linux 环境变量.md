# Linux 环境变量

参考：

https://www.jianshu.com/p/ac2bc0ad3d74

> 所有的变量都可以通过$来获取，==env显示所有环境变量==

## built-in

内建系统环境变量

- HOME

  当前用户的家目录

- HOSTNAME

  主机名

- LANG

  语言环境

- MAIL

  当前用户的邮件目录

- USER

  当前登入的用户

- SHELL

  当前的SHELL

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



