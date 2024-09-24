---
author: "0x00"
createTime: 2024-06-04
lastModifiedTime: 2024-06-04-10:35
draft: true
---

# MySQL 03 - Option Files(my.cnf)

## 0x01 Overview[^1]

*Most MySQL programs can read startup options from option files (sometimes called configuration files). Option files provide a convenient way to specify commonly used options so that they need not be entered on the command line each time you run a program.*

大多数 MySQL 关联的应用(例如 mysqld,mysqldump,mysqladmin)都可以将默认需要使用的 options 写入到 option files(也被记作 configuration file)中。这样在每次调用对应的应用时，就不需要显示的指定对应的 options

要想知道应用支持那些 options，可以使用 `--help` (但是除 mysqld 外，需要额外使用 `--verbose`)

例如

```
$ mysqldump --help
...
The following options may be given as the first argument:
...
Variables (--variable-name=value)
...
```

支持的 options 在第二个省略号部分

## 0x02 Option Files

### Read Order

option files 有读取优先级[^1]，可以使用 `--help` 查看，例如

```
$ mysqldump --help
...
Default options are read from the following files in the given order:
/etc/my.cnf /etc/mysql/my.cnf /usr/local/mysql/etc/my.cnf ~/.my.cnf
The following groups are read: mysqldump client
...
```

默认会按照如下顺序读取

1. `/etc/my.cnf`
2. `/etc/mysql/my.cnf`
3. `/usr/local/mysql/etc/my.cnf`
4. `~/.my.cnf`

如果自定义了 option files 且路径和上述不同，就需要使用 `--defaults-file` 来指定 option files，否则会按照默认顺序中的路径读取 option files

### Option File Syntax

> Leading and trailing spaces are automatically deleted from option names and values.        

- `# comment | ; comment`

  注释

- `[group]`

  指定应用的配置集，`[group]` 和 `[group]` 之间 或者 `[group]` 到 EOF 是 group 应用默认使用的参数

  group 的值可以和 应用名相同，也可以不同(具体看文档)

  例如

  `mysqld` 会同时从 `[mysqld]` 和 `[server]` 中读取配置

- `opt_name`

  等价于应用使用了 `--opt_name`

- `opt_name=value`

  等价于应用使用了 `--opt_name=value`

- `!include <file>`

  引入其他 option file 到当前文件中

- `!includedir <dir>`

  引入其他文件中的 option files 到当前文件中

## 0x03 Examples

```
[mysqld]
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock
log-error=/var/log/mysqld.log
pid-file=/var/run/mysqld/mysqld.pid

#使用 mysql cli 时，如果需要使用 root 账号就可以直接免密
[mysql]
user=root
password=t))r

#使用 mysqldump cli 时，如果需要使用 root 账号就可以直接免密
[mysqldump]
user=root
password=t))r
```

## 0x04 Cautions

1. 如果一些参数的值不合理，MySQL 会将其置为合理

   例如 如下配置

   ```
   [mysqld]
   port=999999
   ```

   实际 mysqld 会监听 65535 端口

2. 如果一个 option 在 option files 中出现多次，就会使用最后一次出现的 option。但是除了 `mysqld` 的 `--user` option 外

## 0x05 Options VS Variables

variables[^2] 和 options 都会影响 MySQL 操作，但是 variables 和 options 不同



options 可以配置到 option files 中，但是 variables 不能。因为部分 variables 和 options 的变量名相同(功能上相同)，所以可以通过这部分 options 设置 variables

例如 

`--port` option 可以写入到 option files, 且有对应的 port variable，所以可以使用 `SHOW VARIABLES like 'port'` 来查看对应的值

`--password` option 可以写入到 option files，但是并没有 password variable，也就不能使用 `SHOW variables like 'password'` 来查看对应的值

## 0x06 When are Option Files been Used



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:https://dev.mysql.com/doc/refman/8.4/en/option-files.html
[^2]:https://dev.mysql.com/doc/refman/8.4/en/server-system-variables.html
[^3]:https://dev.mysql.com/doc/refman/8.4/en/option-file-options.html
