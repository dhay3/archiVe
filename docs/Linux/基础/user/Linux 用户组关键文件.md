# Linux 用户组关键文件

> 使用`:`做为分隔符

## /etc/passwd

https://linux.die.net/man/5/passwd

存储用户信息，密码信息存储在`/etc/shadow`中

1. name：登入账户名
2. password：加密的密码，可以使用`x`或是`*`。表示用`/etc/shadows`文件中的密码替换
3. UID：用户的UID，root使用0
4. GID：用户的GID，查看`/etc/group`文件
5. GECOS：comment filed，创建用户时的评论信息，通过`-c`参数创建
6. directory：用户的家目录
7. shell：用户使用的login shell，默认使用`/bin/sh`

```
root:x:0:0:root:/root:/usr/bin/zsh
zs:x:1001:1001::/home/zs:/bin/sh
```

## /etc/shadow

存储用户密码的加密文件。包含九个域

1. login name：登入的账户名

2. encryted password：加密的登入密码，如果使用`!`或`*`表示改账户不能使用密码登入。但是能被系统登入或其他一些途径

   > 日期单位均采用天，从1979.1.1开始

3. date of last password change：密码变更的时间，0表示下次登入需要变更密码

4. minium password age：允许变更密码的最小时间，0表示没有限制

5. maxium password age：==超过指定时间密码可能任有效(参考password inacitivity period)==，但是会被要求更改密码，如果改域为空表示没有限制

6. password warnging period：距离maxium password age，之前提醒

7. password inactivity period：maxium password age 之后，密码有效的时间。==之后密码将永久失效==。如果改域为空，表示没有限制

8. account expiration date：账户失效的时间，==如果改值为空表示没有限制，如果值为0表示立即失效==

9. reserved field：保留域

```
root:$6$dzV6kR7Yttv1cz9O$Hfwumcplh7nQtIft8rgr5rejNLdxtKggneKy3BRTP4ocMo6lw.kfvlmJVMoZUEqPnhVj5h0mK7ya3uR/:18505:0:99999:7:::
sshd:*:18586:0:99999:7:::
```

## /etc/group

用户组信息

1. group name：组名
2. password：组密码。如果为空表示没有
3. GID：组ID
4. user_list：组中的用户，使用逗号隔开

```
root:x:0:
daemon:x:1:
root:x:0:
daemon:x:1:
```

## /etc/default/useradd

使用useradd命令时默认值配置文件
