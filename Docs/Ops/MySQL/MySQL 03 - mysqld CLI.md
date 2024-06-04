---
title: MySQL 03 - mysqld CLI
author: "0x00"
createTime: 2024-06-04
lastModifiedTime: 2024-06-04-11:19
draft: true
---

# MySQL 03 - mysqld CLI

## 0x01 Overview

`mysqld` 也被认为是 MySQL server[^1]，是 MySQL 中最核心的组件

## 0x02 Defualt Configuration[^2]

`mysqld` 在启动时，如果没有额外的配置或者参数，会使用默认的配置

具体使用的配置可以通过 `mysqld --help --verbose | grep -A 9999 'Variables'` 查看

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

也可以在 `mysql` interactive mode 中查看

```
mysql> show variables;
```

但是你也可以在 `mysqld` 启动前，使用 options 或者是 [option files](MySQL%2002%20-%20Option%20files(my.cnf).md) 来修改这些配置。或者使用 `set <variables>` 来临时修改对应的配置

## 0x03 Configuration Validation



## 0x02 Syntax

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
[^3]:https://dev.mysql.com/doc/refman/8.4/en/server-option-variable-reference.html
[^4]:https://dev.mysql.com/doc/refman/8.4/en/mysqld-server.html
