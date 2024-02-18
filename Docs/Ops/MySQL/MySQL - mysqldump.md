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
Usage: mysqldump [OPTIONS] <database> [tables]
OR     mysqldump [OPTIONS] --databases [OPTIONS] DB1 [DB2 DB3...]
OR     mysqldump [OPTIONS] --all-databases [OPTIONS]
```

> 只列举常见参数，具体参数查看官方文档

### 0x041 Connection args

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

### 0x042 Debug args

- `-i | --comments`

  以注释的形式导出 program version, server version, host, dump date

  ```
  -- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
  --
  -- Host: localhost    Database: test
  -- ------------------------------------------------------
  -- Server version       8.0.29
  ```

  默认开启


- `--skip-comments`

  对 `--comments` 取反

- `--dump-date`

  导出文件会在末尾有一个时间戳

  ```
  -- Dump completed on 2024-02-01 16:04:27
  ```

  默认开启


- `-v | --verbose`

  详细输出

- `--debug & --debug-check & --debug-info`

  输出 debug 日志，只有 MySQL 在构建时使用 WITH_DEBUG 才生效 

### 0x043 Display args

- `-r | --result-file=<file-name>`

  将内容输出到文件，通常用在 Windows 上防止 `\n` 转换成 `\r\n`

- `-X | --xml`

  以 xml 的格式导出

### 0x044 DDL args

- `-A | --all-databases`

  导出所有的数据库，等价与 `--databases <name of all the databases>`

- `-B | --databases <database...>`

  导出指定的数据库(使用该参数可以导出多个数据库)。导出的内容中每个数据库前会使用`CREATE DATABASE ... IF NOT EXSIT` 和 `USE <DATABSE>` 

  ```
  CREATE DATABASE /*!32312 IF NOT EXISTS*/ `test2` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
  
  USE `test2`;
  ```

  ==注意这点和 `mysql [options] <database>` 导出的不同，结果中不包含 `CREATE DATABSE ... IF NOT EXSIT` 或者是 `USE <DATABASE>`（因为没有使用 `--databases` 的格式 `<database>` 后面跟的是 `<tablename>` 只能关联一个数据库，导入数据库时需要用户手动指定）==

- `--tables <table...>`

  导出指定的表(该参数后面的所有值都表示表名)

  ```
  mysqldump -v -u root -p test2 --tables table1
  ```

- `-n | --no-create-db`

  导出的时候没有 `CREATE DB` DDL

- `-t | --no-create-info`

  导出的时候没有 `CREATE TABLE` DDL (包含 `INSERT`)

  ```
  mysqldump -v -u root -p -t test2
  ```

- `--ignore-table=<db_name>.<tbl_name>`

  不导出指定的表,必须指定库名

  ```
  mysql -u root -p -e 'use test2;show tables;'
  +-----------------+
  | Tables_in_test2 |
  +-----------------+
  | table1          |
  | table2          |
  +-----------------+
  
  mysqldump -v -u root -p --ignore-table test2.table1 test2
  ```

- `-d | --no-data`

  只导出表结构 DDL

  ```
  mysqldump -v -u root -p -d test2
  ```

- `-Q | --quote-names`

  database name,table name,column name 会包裹在 back bracket 中 ` (在 back bracket 中可以包含 MySQL 的关键字)

- `--skip-quote-names`

  对 `--quote-names` 取反

- `-c | --complete-insert`

  使用 `INSERT` 时会使用 `INSERT INTO <table> (col_names...) VALUES (values...)` 的格式，而不是默认的 `INSERT INTO <table> VALUES (values...)` 

- `--add-drop-database`

  在每一个 `CREATE DATBASE` 语句前使用 `DROP DATABASE ... IF EXSIT`

- `--add-drop-table`

  在每一个 `CREATE TABSE` 语句前使用 `DROP TABLE ... IF EXSIT`

  默认开启

- `--add-drop-trigger`

  在每一个 `CREATE TRIGGER` 语句迁使用 `DROP TRIGGER ... IF EXSIT`

- `--replace`

  使用 `REPLACE` 替代 `INSERT`

- `-R | --routines`

  导出 `CREATE PROCEDURE` 和 `CREATE FUNCTION` DDL

- `--triggers`

  导出 triggers

### 0x045 Transaction args

- `--add-locks`

  导出的 DDL 执行 `INSERT` 前使用 `LOCK TABLE <table> WRITE` 数据插入完成后使用 `UNLOCK TABLE`

  默认开启

- `-x | --lock-all-tables`

  在导出过程中，被导出的数据库的所有表使用 READ LOCK

  默认开启

## 0x05 Examples

1. 整库备份

   ```
   mysqldump -u root -p -v test2 -r test2.sql
   #等价, 但是推荐第 1 种
   mysqldump -u root -p -v test2 > test2.sql
   ```

2. 多库备份

   ```
   mysqldump -u root -p -v --databases test test2 -r all_test.sql
   ```

3. 所有库备份

   ```
   mysqldump -u root -p --all-databases -r all_dbs.sql
   ```

4. 导出库中的某一张表

   ```
   mysqldump -v -u root -p test2 table1 -r test2.table1.sql
   ```

5. 导出库中的某几张表

   ```
   mysqldump -v -u root -p test2 table1 table2 -r test2.sql
   ```

6. 导出库并同步(最好先备份再同步)

   ```
   mysqldump -v -u root -p test2 | mysql -h 10.0.1.205 -u root -p'test1234' test2
   ```

7. 导出指定表以外的所有表

   ```
   mysqldump -v -u root -p --ignore-table test2.table1 test2 -r table1.sql
   ```

## 0x06 How to restore dumps

1. `mysql -v -u root -p test2 < test2.sql`

2. `mysql -v -u root -p -e "source /path/to/dump.sql" test2`

   等价于先登入 mysql 然后执行 `use test2;source /path/to/dump`

**references**

[^1]:https://dev.mysql.com/doc/refman/8.0/en/mysqldump.html
[^2]:https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_lock-tables
[^3]:https://stackoverflow.com/questions/29539838/replace-versus-insert-in-sql