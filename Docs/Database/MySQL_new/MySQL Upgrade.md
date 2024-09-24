---
title: MySQL Upgrade
author: "0x00"
createTime: 2024-06-03
lastModifiedTime: 2024-06-03-09:29
draft: true
tags:
  - Database
---


> [!info]
> 特指从 5.x - 8.x[^1]

```terminal
mysqldump -u root -p --add-drop-table --routines --events --all-databases --force > data-for-upgrade.sql
```
---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[MySQL :: MySQL 8.0 Reference Manual :: 3 Upgrading MySQL](https://dev.mysql.com/doc/refman/8.0/en/upgrading.html)