# MySQL - mysqldump

## 0x01 Overview

> *The `mysqldump` client utility performs logical backups.*

`mysqldump` 是 MySQL 的一个命令行工具，用于 ==logical backup==(a backup that reproduces table structure and data, without copying the actual data files. For example the, the `mysqldump` command produces a logical backup, because its output contains statement such as `CREATE TABLE` and `INSERT` that can re-create the data.)

`mysqldump` 除了可以导出 SQL 的 DDL DML, 还可以以 CSV、XML 的格式导出

## 0x02 Priviledges

如果想要通过 `mysqldump` 来导出 logical backups ，对应的用户需要有相应的权限

> 对应的用户是否有权限可以通过 `SELECT user, host, Grant_priv, Super_priv, Create_tmp_table_priv, Lock_tables_priv, Create_view_priv, Show_view_priv, Create_routine_priv, Alter_routine_priv, Execute_priv, Event_priv, Trigger_priv FROM mysql.user;` 来查看
>
> 这里只使用了 `DESC mysql.user` 来查看对应权限的列名

| Action                                                       | Priviledges      |
| ------------------------------------------------------------ | ---------------- |
| Dumped tables                                                | Select_priv      |
| Dumped views                                                 | Show_view_priv   |
| Dumped triggers                                              | Trigger_priv     |
| `--signle-transaction` is not used                           | Lock_tables_priv |
| `--no-tablespaces` is not used                               | Process_priv     |
| `--single-transaction` and `--gtid_mode=ON` and `--set-gtid=purged=ON|AUTO` | Reload_priv      |

如果用户不具备对应的权限，需要 DBA 授予用户对应的权限

```
GRANT SELECT,SHOW VIEW ON *.* to 'test'@'localhost';
```

## 0x03 Performance

`mysqldump` 是

## 0x04 Syntax

```
Usage: mysqldump [OPTIONS] database [tables]
OR     mysqldump [OPTIONS] --databases [OPTIONS] DB1 [DB2 DB3...]
OR     mysqldump [OPTIONS] --all-databases [OPTIONS]
```

### 0x041 Optional args

> 只列举常见参数，具体参数查看官方文档

#### 0x0411 Connection args

- `--bind-address=#`

  On a computer having multiple network interfaces, use this option to select which interface to use for connecting to the MySQL server

  指定本机使用的接口（一般用于多接口主机连接 MySQL Server）

- `-C | --compress`

  Compress all information sent between the client and the server if possible

  压缩 客户端 和 服务端 之前的消息传输

- `--compression-algorithms=<zlib | uncompressed | zstd>`

  The permitted compression algorithms for connections to the server.

  指定使用的压缩算法

  默认使用 `uncompressed`

- `--zstd-compression-level=<level>`

  指定 zstd 的压缩程度，1 - 22

  只有 `--compression-algorithms=zstd` 是才生效

- `-h<host> | --host=<host>`

  指定连接的主机

  默认 localhost

- `-P<port>| --port=<port>`

  指定连接的端口

  默认 3306

- `--protocol=<tcp | socket | pipe | memory>`

  指定连接的协议

- `-u<user> | --username=<user>`

  指定连接的用户

- `-p[password] | --password`

  指定连接使用的密码

#### 0x0412 DDL args

- `--add-drop-database`

  在每一个 `CREATE DATBASE` 语句前使用 `DROP DATABASE ... IF EXSIT`

- `--add-drop-table`

  在每一个 `CREATE TABSE` 语句前使用 `DROP TABLE ... IF EXSIT`

  默认开启

- `--add-drop-trigger`

  在每一个 `CREATE TRIGGER` 语句迁使用 `DROP TRIGGER ... IF EXSIT`

- `-n | --no-create-db`

  导出的时候没有 `CREATE DB` DDL

- `-t | --no-create-info`

  导出的时候没有 `CREATE TABLE` DDL

- `-d | --no-data`

  只导出表结构

- `--replace`

  使用 `REPLACE` 替代 `INSERT`

#### 0x0412 Debug args

- `-i | --comments`

  以注释的形式导出 program version, server version, host

  ```
  -- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
  --
  -- Host: localhost    Database: test
  -- ------------------------------------------------------
  -- Server version       8.0.29
  ```

  默认开启

- `--debug & --debug-check & --debug-info`

  输出 debug 日志，只有 MySQL 在构建时使用 WITH_DEBUG 才生效 

- `--dump-date`

  导出文件会在末尾有一个时间戳

  ```
  -- Dump completed on 2024-02-01 16:04:27
  ```

  默认开启

- `-v | --verbose`

  详细输出

#### 0x0413 Out args

- `--print-defaults`

  `mysqldump` 会使用的默认参数，可以通过 `my.cnf` 来配置

- `--no-defaults`

  不读取任何 `my.cnf` 中的配置

- `--defualts-file=#`

  只读取指定的 `my.cnf`

- `--add-locks`

- `-A | --all-databases`

  导出所有的 DB

- `-i | --comments`

  导出所有的备注

  默认开启

- 

- 

  







**references**

[^1]:https://dev.mysql.com/doc/refman/8.0/en/mysqldump.html
[^2]:https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_lock-tables
[^3]:https://stackoverflow.com/questions/29539838/replace-versus-insert-in-sql