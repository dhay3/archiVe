# 01 - MySQL EBNF

> SQL keywords are not case-sensitive and may be written in any lettercase. This manual uses uppercase

- `[ ]`

  Square brackets (“`[`” and    “`]`”) indicate optional words or clauses

  ```
  DROP TABLE [IF EXISTS] tbl_name
  ```

  例如清空表的时候可以选择使用 `IF EXISTS`

- `|`

  When a syntax element consists of a number of alternatives, the alternatives are separated by vertical bars    (“`|`”).

  ```
  TRIM([[BOTH | LEADING | TRAILING] [remstr] FROM] str)
  ```

  例如 `TRIM` 函数如果有 3 个入参第一个入参可以是 BOTH

- `{ }`

  

**references**

1. [^1]:https://dev.mysql.com/doc/refman/8.3/en/manual-info.html