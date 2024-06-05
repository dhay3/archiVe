---
author: "0x00"
createTime: 2024-06-04
lastModifiedTime: 2024-06-04-11:19
draft: true
---

# MySQL 05 - mysqld CLI

## 0x01 Overview

`mysqld` 也被认为是 MySQL server[^1]，是 MySQL 中最核心的组件

## 0x02 Defualt Configuration[^2]

`mysqld` 在启动时，如果没有配置 [option files](MySQL%2003%20-%20Option%20Files(my.cnf).md) 或者指定 command line options，会使用默认的 operating parameters(optoins + varables)

variables 部分的默认值都可以通过 `mysqld --help --verbose | grep -A 9999 'Variables'` 查看

```
mysqld --help --verbose | grep -A 99999  'Variables'
Variables (--variable-name=value)
and boolean options {FALSE|TRUE}                             Value (after reading options)
------------------------------------------------------------ -------------
abort-slave-event-count                                      0
activate-all-roles-on-login                                  FALSE
admin-address                                                (No default value)
admin-port                                                   33062
admin-ssl                                                    TRUE
...
```

也可以在 interactive mode 中查看

```
mysql> show variables;
```

variables 部分会包括部分 optoins，但是一些 options 不会显示，例如 `user`, `password`

除了 [option files](MySQL%2003%20-%20Option%20Files(my.cnf).md) 和 command line options，也可以在 interactive mode 中使用 `SET <variables>=<values>` 来临时修改对应的配置

## 0x03 Configuration Validation[^3]

`mysqld` 支持一个特殊的参数用于校验 options 以及使用的 option files 是否正常



## 0x04 Syntax

```
Usage: mysqld [OPTIONS]
```

### Options

所有在 command line 或者是 option files 中可以使用的配置项，都可以参考 [Server Command Options][^4]



#### basic options

- `-b dir_name | --basedir=dir_name`

  MySQL 安装的目录

### variables

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:https://dev.mysql.com/doc/refman/8.4/en/mysqld.html
[^2]:https://dev.mysql.com/doc/refman/8.4/en/server-configuration-defaults.html
[^3]:https://dev.mysql.com/doc/refman/8.4/en/server-configuration-validation.html
[^4]:https://dev.mysql.com/doc/refman/8.4/en/server-option-variable-reference.html
[^5]:https://dev.mysql.com/doc/refman/8.4/en/mysqld-server.html
