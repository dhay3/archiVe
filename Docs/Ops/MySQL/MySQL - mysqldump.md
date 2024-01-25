# MySQL - mysqldump

## 0x01 Overview

> *The `mysqldump` client utility performs logical backups.*

`mysqldump` 是 MySQL 的一个命令行工具，用于 logical backup(a backup that reproduces table structure and data, without copying the actual data files. For example the, the `mysqldump` command produces a logical backup, because its output contains statement such as `CREATE TABLE` and `INSERT` that can re-create the data.)

`mysqldump` 除了可以导出 SQL 的 DDL DML, 还可以以 CSV、XML 的格式导出

## 0x02 Priviledges

如果想要通过 `mysqldump` 来导出 logical backups 对应的用户需要有相应的权限

| Action                             | Priviledges      |
| ---------------------------------- | ---------------- |
| Dumped tables                      | Select_priv      |
| Dumped views                       | Show_view_priv   |
| Dumped triggers                    | Trigger_priv     |
| `--signle-transaction` is not used | Lock_tables_priv |
| `--no-tablespaces` is not used     | Process_priv     |
|                                    |                  |
|                                    |                  |

**references**

[^1]:https://dev.mysql.com/doc/refman/8.0/en/mysqldump.html