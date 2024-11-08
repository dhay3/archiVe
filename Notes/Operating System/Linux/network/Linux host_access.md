# Linux host_access

## 概述

`host_access`用于控制指定host能否访问本主机

>    •      Access will be granted when a (daemon,client) pair matches an entry in the /etc/hosts.allow file.
>
>    •      Otherwise, access will be denied when a (daemon,client) pair matches an entry in the /etc/hosts.deny file.
>
>    •      Otherwise, access will be granted.
>
> 默认使用白名单机制，一般只需配置`hosts.deny`文件即可

## Access Controll Rule

使用rule匹配规则

pattern：`daemon_list:client_list[:shell_command]`

- daemon_list：daemon process，server port numbers or wildcard

- client_list：host names，host address，pattern or wildcard

  如果有多个client_list使用逗号分隔

## patterns

1. 如果以`.`开头，表示匹配结尾。`.tue.nl`能匹配`wzv.win.tue.nl`
2. 如果以`.`结尾，表示匹配开头。`131.155.`能匹配`131.155.x.x`
3. `n.n.n.n/m.m.m.m`，表示匹配`net/mask`
4. 支持CIDR形式
5. 可以使用`*`和`?`通配符，不能与1，2，3规则一起使用

## wildcards

通配符

1. ALL，匹配所有
2. LOCAL，配置不包含dot的所有host
3. UNKNOW，匹配

可以使用EXCEPT表示除外

## 案例

- `/etc/hosts.deny`

```
ALL: ALL
ALL: .foobar.edu EXCEPT terminalserver.foobar.edu #除terminalserver.foobar.edu
ALL: some.host.name, .some.domain
in.tftpd: ALL: (/usr/sbin/safe_finger -l @%h | \ #backslash表示拼接命令
               /usr/bin/mail -s %d-%h root) &
```

