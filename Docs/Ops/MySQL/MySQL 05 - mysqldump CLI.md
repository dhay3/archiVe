# MySQL 05 - mysqldump CLI

## 0x01 Overview

> *The `mysqldump` client utility performs logical backups.*

`mysqldump` 是 MySQL 的一个命令行工具，用于 ==logical backup==(a backup that reproduces table structure and data, without copying the actual data files. For example the, the `mysqldump` command produces a logical backup, because its output contains statement such as `CREATE TABLE` and `INSERT` that can re-create the data.)

简单的理解就是 `mysqldump` 可以将数据库转为 SQL，除了可以导出 SQL 的 DDL DML, 还可以以 CSV、XML 的格式导出

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
| `--single-transaction` and `--gtid_mode=ON` and `--set-gtid=purged=ON|AUTO `is used | Reload_priv      |
| `--tab` is used                                              | File_priv        |

如果用户不具备对应的权限，需要 DBA 授予用户对应的权限

```
GRANT SELECT,SHOW VIEW ON *.* to 'test'@'localhost';
```

## 0x03 Syntax

如果是单库推荐使用第一种，生成的 SQL 中不包含 `CREATE DATABASE` 和 `USE DATABASE` (不受源目数据库名限制)

```
Usage: mysqldump [OPTIONS] <database> [tables]
OR     mysqldump [OPTIONS] --databases [OPTIONS] DB1 [DB2 DB3...]
OR     mysqldump [OPTIONS] --all-databases [OPTIONS]
```

> 只列举常见参数，具体参数查看官方文档

### Connection Options

- `--bind-address=#`

  On a computer having multiple network interfaces, use this option to select which interface to use for connecting to the MySQL server

  指定连接 MySQL Server 使用的本机接口（一般用于多接口主机无法自动判断路由的情况下使用）

- `-S<name> | --socket=<name>`

  使用指定的 socket 文件用于连接 MySQL server

  默认 `/var/lib/mysql/mysql.socket`

  ==如果服务器上有多实例就需要使用该参数==

- `-h[host] | --host[=host]`

  指定连接使用的 MySQL Server 的地址

  默认 localhost

- `-P[port]| --port[=port]`

  指定连接使用的 MySQL Server 的端口

  默认 3306

- `--protocol=<tcp | socket | pipe | memory>`

  指定连接使用的 MySQL Server 的协议

- `-u[user] | --username[=user]`

  指定连接使用的 MySQL Server 的用户名

  默认 `${USER}`

- `-p[password] | --password[=password]`

  指定连接使用的 MySQL Server 的密码，如果 password 缺省，会 prompt 让用户输入
  
  ==注意 `-p` 和 `[password]` 中间不能有空格==

- `-C | --compress`

  Compres all information sent between the client and the server if possible

  压缩 客户端 和 服务端 之前的消息传输

  8.0.18 MySQL 以上 deprecated

- `--compression-algorithms=<zlib | uncompressed | zstd>`

  The permitted compression algorithms for connections to the server.

  指定连接 服务端 时使用的压缩算法，和 `--compress` 一起使用时才有意义

  默认使用 `uncompressed`

- `--zstd-compression-level=<level>`

  指定 zstd 的压缩程度，1 - 22

  只有 `--compression-algorithms=zstd` 时才生效，和 `--compress` 一起使用时才有意义

### Debug Options

- `--compact`

  等价于同时使用 `--skip-add-drop-table`, `--skip-add-locks`, `--skip-comments`, `skip-disable-keys`， `--skip-set-charset`

- `-i | --comments`

  以注释的形式导出 program version, server version, host, dump date

  ```
  -- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
  --
  -- Host: localhost    Database: test
  -- ------------------------------------------------------
  -- Server version       8.0.29
  ```

  默认开启，可以使用 `--skip-comments` 关闭


- `--dump-date`

  导出文件会在末尾有一个时间戳

  ```
  -- Dump completed on 2024-02-01 16:04:27
  ```

  默认开启，可以使用 `--skip-dump-date` 关闭


- `-v | --verbose`

  详细输出

- `--debug & --debug-check & --debug-info`

  输出 debug 日志，只有 MySQL 在构建时使用 WITH_DEBUG 才生效 

- `--log-error=<name>`

  如果 dump 的过程中出向 erros 或者是 warnings，会将对应的信息记录到 name 文件中

  如果和 `-v` 一起使用，对应的信息将不会输出到 stdout 和 stderr 中

### Display Options

- `-r | --result-file=<file-name>`

  将内容输出到文件，通常用在 Windows 上防止 `\n` 转换成 `\r\n` (在 windows 上使用 `\r\n` 作为换行符，如果不使用该参数会额外输出空行)

- `-X | --xml`

  以 xml 的格式导出，格式如下
  
  ```
  <?xml version="1.0"?>
  <mysqldump xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <database name="world">
  <table_structure name="City">
  <field Field="ID" Type="int(11)" Null="NO" Key="PRI" Extra="auto_increment" />
  <field Field="Name" Type="char(35)" Null="NO" Key="" Default="" Extra="" />
  <field Field="CountryCode" Type="char(3)" Null="NO" Key="" Default="" Extra="" />
  <field Field="District" Type="char(20)" Null="NO" Key="" Default="" Extra="" />
  <field Field="Population" Type="int(11)" Null="NO" Key="" Default="0" Extra="" />
  <key Table="City" Non_unique="0" Key_name="PRIMARY" Seq_in_index="1" Column_name="ID"
  Collation="A" Cardinality="4079" Null="" Index_type="BTREE" Comment="" />
  <options Name="City" Engine="MyISAM" Version="10" Row_format="Fixed" Rows="4079"
  Avg_row_length="67" Data_length="273293" Max_data_length="18858823439613951"
  Index_length="43008" Data_free="0" Auto_increment="4080"
  Create_time="2007-03-31 01:47:01" Update_time="2007-03-31 01:47:02"
  Collation="latin1_swedish_ci" Create_options="" Comment="" />
  </table_structure>
  <table_data name="City">
  <row>
  <field name="ID">1</field>
  <field name="Name">Kabul</field>
  <field name="CountryCode">AFG</field>
  <field name="District">Kabol</field>
  <field name="Population">1780000</field>
  </row>
  ```

- `--fields-terminated-by=`

### Filter Options

- `-A | --all-databases`

  导出所有的数据库，等价与 `--databases <name of all the databases>` 

- `-B | --databases <database...>`

  导出指定的数据库(使用该参数可以导出多个数据库)。导出的内容中每个数据库前会使用`CREATE DATABASE ... IF NOT EXSIT` 和 `USE <DATABSE>` (同理 `--all-databases`)

  ```
  CREATE DATABASE /*!32312 IF NOT EXISTS*/ `test2` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
  
  USE `test2`;
  ```

  ==注意这点和 `mysql [options] <database>` 导出的不同，结果中不包含 `CREATE DATABSE ... IF NOT EXSIT` 或者是 `USE <DATABASE>`（因为没有使用 `--databases` 的格式 `<database>` 后面跟的是 `<tablename>` 只能关联一个数据库，导入数据库时需要用户手动指定）==

  如果原目数据库名不同，使用该参数导出的文件，会在目的 MySQL server 上创建原库名。需要谨慎使用该参数

- `--tables <table...>`

  导出指定的表

  ```
  mysqldump -v -u root -p test2 --tables table1
  ```

  数据库名只能在该参数前

  ```
  #错误
  mysqldump -u root -p --tables table1 test2 -v
  #正确
  mysqldump -u root -p test2 --tables table1 -v
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

- `-n | --no-create-db`

  导出的时候不包含 `CREATE DB` DDL，和 `--databases` 或者 `--all-databases` 一起使用时才有意义

- `-d | --no-data`

  不导出 `INSERT` DML(如果只想导出表结构的话)。可以将该参数记为 minus data

  ==在表结构修改，但是数据没有变的场景下非常有用==

- `-t | --no-create-info`。可以将该参数记为 minus tables

  导出的时候不包含 `CREATE TABLE` DDL(如果只想导出数据的话)。可以理解为只导出表条目

  ==在增量同步数据场景下非常有用==

- `--where<='condition'> | -w'<condition>'`

  只导出对应条件下的表条目。==会把表结构也一起导出来，如果只想要对应条目 INSERT 就需要和 `-d` 一起使用

  ```
  $ mysqldump -v -u root test2 -d -w'col1="t1v1"'
  INSERT INTO `table1` VALUES ('t1v2');
  ```

  在知道对应的条目是那个表的情况下，最好和 `--tables` 一起使用 或者是使用 `mysqldump [database] [tables]`

- `-E | --events`

  导出 envents

- `--triggers`

  导出 triggers

  可以使用 `--skip-triggers` 来关闭

- `-R | --routines`

  导出 `CREATE PROCEDURE` 和 `CREATE FUNCTION` DDL

### DDL Options

> 诸如 `CREATE`, `DROP`, `ALTER`, `TUNCATE`, `COMMENT`, `RENAME` 对表结构进行定义的都为 DDL

- `--add-drop-database`

  在每一个 `CREATE DATBASE` 语句前使用 `DROP DATABASE ... IF EXSIT`

  ```
  CREATE DATABASE /*!32312 IF NOT EXISTS*/ `test2` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
  ```

  只有在使用 `--databases` 或者是 `--all-databases` 情况下生效，因为 `mysqldump [database] [tables]` 并不会生成 `CREATE DABATBASE` DDL 

  ==全库同步的场景下，可能会包含 `DROP DATABASE mysql`，谨慎使用该参数==

- `--add-drop-table`

  在每一个 `CREATE TABLE` 语句前使用 `DROP TABLE ... IF EXSIT`

  ```
  DROP TABLE IF EXISTS `table2`;
  ```

  默认开启，可以使用 `--skip-add-drop-table` 来关闭

- `--add-drop-trigger`

  在每一个 `CREATE TRIGGER` 语句迁使用 `DROP TRIGGER ... IF EXSIT`

- `-T=<dir> | --tab=<dir>`

  DDL 和 DML 会单独生成 dir/table_name.sql(只包含 DDL) 和 dir/table_name.txt(只包含 DML)

  ```
  mysqldump -v -t test2 table1 --tab=.
  ```

  需要用户需要 File_Priv 权限，以及服务器不能使用 `--secure-file-priv`

- `-Q | --quote-names`

  database name,table name,column name 会包裹在 back bracket 中 ` (在 back bracket 中可以包含 MySQL 的关键字)

  默认开启，可以使用 `--skip-quote-names` 关闭

### DML Options

> 诸如 `INSERT`, `UPDATE`, `DELETE`, `LOCK` 对表数据进行操作的都为 DML

- `-c | --complete-insert`

  使用 `INSERT` 时会使用 `INSERT INTO <table> (col_names...) VALUES (values...)` 的格式，而不是默认的 `INSERT INTO <table> VALUES (values...)` 

  ```
  $ mysqldump -v -u root test2 table2 -n
  INSERT INTO `table2` VALUES ('t2v1'),('t2v2');
  
  $ mysqldump -v -u root test2 table2 -c
  INSERT INTO `table2` (`col1`) VALUES ('t2v1'),('t2v2');
  ```

  ==在增量同步的场景中，如果 原表 和 目的表 结构 不同，就需要使用该参数==

- `--insert-ignore`

  使用 `INSERT IGNORE` 替代 `INSERT`

  ```
  $ mysqldump -v -u root -e test2 -t --insert-ignore --compact
  INSERT  IGNORE INTO `table2` VALUES ('t2v1'),('t2v2');
  ```

  即如果表中有对应的数据相同的 primary key 或者 unique key ，就不会插入对应的条目

  注意如果和 `--replace` 一起使用，`mysqldump` 会正常输出 SQL，但是并不存在 `REPLACE IGNORE` 这中语法，导入时会报错。直接理解为和 `--replace` 和 `--insert-ignore` 不能一起使用

  ```
  $ mysqldump -v -u root test2 --replace  --insert-ignore | mysql -v -u root test2
  ...
  --------------
  REPLACE  IGNORE INTO `table2` VALUES ('t2v1'),('t2v2')
  --------------
  
  ERROR 1064 (42000) at line 57: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'IGNORE INTO `table2` VALUES ('t2v1'),('t2v2')' at line 1
  ```

- `-e | --extended-insert`

  使用 `INSERT INTO <table> VALUES (values...),(values)` 的格式，可以减少 dump 出来的文件大小，以及执行 SQL 时的速度

  ```
  $ mysqldump -v -u root -t   --compact test2 table1
  INSERT INTO `table1` VALUES ('t1v1'),('t1v2');
  ```

  默认开启，可以使用 `--skip-extended-insert` 关闭

  ```
  $ mysqldump -v -u root -t  --skip-extended-insert --compact test2 table1
  INSERT INTO `table1` VALUES ('t1v1');
  INSERT INTO `table1` VALUES ('t1v2');
  ```

- `--replace`

  使用 `REPLACE` 替代 `INSERT` 即如果表中有对应的数据，就先删除然后再插入，反之直接插入

  ```
  $ mysqldump -v -t test2 table1 --replace
  REPLACE INTO `table2` VALUES ('t2v1'),('t2v2');
  ```

  同样也适用 `--complete-insert` 

  ```
  $ mysqldump -v -t test2 table1 --replace -c
  REPLACE INTO `table1` (`col1`) VALUES ('t1v1'),('t1v2');
  ```

  虽然和 `--insert-ignore` 一起使用不会报错，但是 MySQL 不支持 `REPLACE IGNORE` 这种语法

  `--replace` 通常和 `--table-info` 一起使用，更新表条目而不更新表结构 

### Transaction Options

> `mysqldump` 默认参数中

- `-l | --lock-tables`

  导表前会对所有需要导的表上锁

  默认开启

- `-x | --lock-all-tables`

  导表的过程中所有数据库的所有表会上 READ LOCK，会默认关闭 `--single-transaction` 和 `--lock-tables`

- `--add-locks`

  执行 `INSERT` 前使用 `LOCK TABLE <table> WRITE` 数据插入完成后使用 `UNLOCK TABLE`

  ```
  LOCK TABLES `table1` WRITE;
  INSERT INTO `table1` VALUES ('t1v1'),('t1v2');
  UNLOCK TABLES;
  ```

  默认开启，可以使用 `--skip-add-locks` 来关闭

- `--single-transaction`

  导数据前会使用 `START TRANSACTION ` 只针对导出数据这一过程，生成的 SQL 并不会包含 `START TRANSACTION`

### Miscellaneous Opargstions

- `-V | --version`

  查看 `mysqldump` 详细的版本信息

- `-h | -?`

  查看帮助信息

## 0x04 Default values

`mysqldump` 中的每一个参数都可以被写入 [MySQL 02 - Option files(my.cnf)](MySQL%2002%20-%20Option%20files(my.cnf).md)

例如在 `/etc/my.cnf` 中声明 `user` 和 `password` 就无需使用 `--user` 或者是 `--password`

```
[mysqldump]
user=root
password=test1234
```

默认会按照顺序从如下文件中读取配置

1. `/etc/my.cnf`
2. `/etc/mysql/my.cnf`
3. `/usr/etc/my.cnf`
4. `~/.my.cnf`

如果配置文件中没有 `mysqldump` 相关的部分就会使用如下默认配置

```
all-databases                     FALSE
all-tablespaces                   FALSE
no-tablespaces                    FALSE
add-drop-database                 FALSE
add-drop-table                    TRUE
add-locks                         TRUE
allow-keywords                    FALSE
apply-slave-statements            FALSE
character-sets-dir                (No default value)
comments                          TRUE
compatible                        (No default value)
compact                           FALSE
complete-insert                   FALSE
compress                          FALSE
create-options                    TRUE
databases                         FALSE
debug-check                       FALSE
debug-info                        FALSE
default-character-set             utf8
delayed-insert                    FALSE
delete-master-logs                FALSE
disable-keys                      TRUE
dump-slave                        0
events                            FALSE
extended-insert                   TRUE
fields-terminated-by              (No default value)
fields-enclosed-by                (No default value)
fields-optionally-enclosed-by     (No default value)
fields-escaped-by                 (No default value)
flush-logs                        FALSE
flush-privileges                  FALSE
force                             FALSE
gtid                              FALSE
hex-blob                          FALSE
host                              (No default value)
include-master-host-port          FALSE
insert-ignore                     FALSE
lines-terminated-by               (No default value)
lock-all-tables                   FALSE
lock-tables                       TRUE
log-error                         (No default value)
log-queries                       TRUE
master-data                       0
max-allowed-packet                16777216
net-buffer-length                 1046528
no-autocommit                     FALSE
no-create-db                      FALSE
no-create-info                    FALSE
no-data                           FALSE
no-data-med                       TRUE
order-by-primary                  FALSE
port                              3306
quick                             TRUE
quote-names                       TRUE
replace                           FALSE
routines                          FALSE
set-charset                       TRUE
single-transaction                FALSE
dump-date                         TRUE
socket                            /var/lib/mysql/mysql.sock
ssl                               FALSE
ssl-ca                            (No default value)
ssl-capath                        (No default value)
ssl-cert                          (No default value)
ssl-cipher                        (No default value)
ssl-key                           (No default value)
ssl-crl                           (No default value)
ssl-crlpath                       (No default value)
ssl-verify-server-cert            FALSE
tab                               (No default value)
triggers                          TRUE
tz-utc                            TRUE
user                              (No default value)
verbose                           FALSE
where                             (No default value)
plugin-dir                        (No default value)
default-auth                      (No default value)
```

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

4. 只更新表数据

   ```
   mysqldump -u root -p test2 --replace -n
   ```

5. 导出库中的某一张表

   ```
   mysqldump -v -u root -p test2 table1 -r test2.table1.sql
   ```

6. 导出库中的某几张表

   ```
   mysqldump -v -u root -p test2 table1 table2 -r test2.sql
   ```

7. 只导出表结构

   ```
   mysqldump -v -u root -p -d test2 table1 table2 -r test2.sql
   ```

8. 只导出表条目

   ```
   mysqldump -v -u root -p -t test2 table1 table2 -r test2.sql
   ```

9. 导出库并同步(最好先备份再同步)

   ```
   mysqldump -v -u root -p test2 | mysql -h 10.0.1.205 -u root -p'test1234' test2
   ```

10. 导出指定表以外的所有表

    ```
    mysqldump -v -u root -p --ignore-table test2.table1 test2 -r table1.sql
    ```

11. 导出查询结果

    ```
    mysqldump -v -u root -p test2 -e "select * from table1" > table1.result
    ```

12. 导出结果为 excel

    ```
    mysqldump -v -u root -p test2 -e "select * from table1" | tr '\t' ',' > table.csv
    ```

    需要注意的一点是，在 windows 上如果直接使用 excel 打开对应的 csv 文件，因为编码可能会显示乱码，可以在 notepad 里先以 UTF-8 的格式打开，然后直接复制到 csv 文件即可

## 0x06 How to restore dumps

1. `mysql -v -u root -p test2 < test2.sql`

2. `mysql -v -u root -p -e "source /path/to/dump.sql" test2`

   等价于先登入 mysql 然后执行 `use test2;source /path/to/dump`

**references**

[^1]:https://dev.mysql.com/doc/refman/8.0/en/mysqldump.html
[^2]:https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_lock-tables
[^3]:https://stackoverflow.com/questions/29539838/replace-versus-insert-in-sql
[^4]:https://www.geeksforgeeks.org/sql-ddl-dql-dml-dcl-tcl-commands/
[^5]:https://stackoverflow.com/questions/12040816/dump-all-tables-in-csv-format-using-mysqldump#25427665
