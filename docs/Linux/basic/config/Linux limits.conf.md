# Linux limits.conf

https://blog.csdn.net/qq_41378597/article/details/103706237

> 注意ubuntu系统不支持通配符，重启后生效

`/etc/security/limits.conf`只针对用户级别的login shell生效。是ulimit的持久配置。

每一行都由`<domain><type><item><value>`组成。

## domain

> 如果表示对root用户的限制，domain必须指明是root

- 可以是用户名
- 使用`@group`表示组名
- 可以使用通配符

## type

- soft：软限制
- hard：强制限制

## item

- core

  core dump生成的文件的最大值

- data

  最大数据段长度

- fsize

  最大的文件大小

- memlock

  内存中被锁定的大小

- nofile

  能打开的最大文件句柄

- nproc

  能打开的最大进程数

- maxlogins

  用于指组登入的用户个数

  ```
  @student      hard       maxlogins       4
  ```

## 示例

```
         *               soft    core            0
           root            hard    core            100000
           *               hard    nofile          512
           @student        hard    nproc           20
           @faculty        soft    nproc           20
           @faculty        hard    nproc           50
           ftp             hard    nproc           0
           @student        -       maxlogins       4
           :123            hard    cpu             5000
           @500:           soft    cpu             10000
           600:700         hard    locks           10
```











