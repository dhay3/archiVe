# su

==`su`没有带有任何参数是以root的身份运行interactive shell==，su默认不会修改运行interactive shell时的目录。`su`使用PAM作为authentication，并且是对unprivileged用户，对于privlileged用户(root)可以使用`runuser username `来切换用户

- -c command

  不分配tty，直接将执行的命令结果返回

  ```
  root in /home/ubuntu λ su -c "top" ubuntu
  ```

- -s shell

  使用指定的shell，全路径

  ```
  root in /home/ubuntu λ su ubuntu -s /bin/zsh
  ```

- `- | -l`

  登入后到用户的家目录

  ```
  sudo - ubuntu
  ```

## Config files

> `man login.defs`

`su`默认读取`/etc/defaults/su`和`/etc/login.defs`中的配置文件。具体[参考](../config/login.defs.md)

