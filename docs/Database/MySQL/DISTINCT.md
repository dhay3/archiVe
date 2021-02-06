# DISTINCT

转载:
https://www.cnblogs.com/lbxBlog/p/9383174.html

## 基本使用

`distinct`一般是用来去除查询结果中的重复记录的，而且这个语句在`select`、`insert`、`delete`和`update`中只可以在`select`中使用，具体的语法如下：

```sql
select distinct expression[,expression...] from tables [where conditions];
```

- 1

这里的expressions可以是多个字段。本文的所有操作都是针对如下示例表的：

```sql
CREATE TABLE `NewTable` (
    `id`  int(11) NOT NULL AUTO_INCREMENT ,
    `name`  varchar(30) NULL DEFAULT NULL ,
    `country`  varchar(50) NULL DEFAULT NULL ,
    `province`  varchar(30) NULL DEFAULT NULL ,
    `city`  varchar(30) NULL DEFAULT NULL ,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB
;
```

![这里写图片描述](https://img-blog.csdn.net/20170622231135994?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbG15ODYyNjM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 1.1 只对一列操作

这种操作是最常见和简单的，如下：

```sql
select distinct country from person
```

- 1

结果如下： 
![这里写图片描述](https://img-blog.csdn.net/20170622231159269?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbG15ODYyNjM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 1.2 对多列进行操作

```sql
select distinct country, province from person
```

- 1

结果如下： 
![这里写图片描述](https://img-blog.csdn.net/20170622231212527?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbG15ODYyNjM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

从上例中可以发现，当`distinct`应用到多个字段的时候，其应用的范围是其后面的所有字段，而不只是紧挨着它的一个字段，而且`distinct`只能放到所有字段的前面，如下语句是错误的：

```sql
SELECT country, distinct province from person; // 该语句是错误的
```

- 1

抛出错误如下：

> [Err] 1064 - You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near ‘DISTINCT province from person’ at line 1

### 1.3 针对`NULL`的处理

从1.1和1.2中都可以看出，`distinct`对`NULL`是不进行过滤的，即返回的结果中是包含`NULL`值的。

### 1.4 与`ALL`不能同时使用

默认情况下，查询时返回所有的结果，此时使用的就是`all`语句，这是与`distinct`相对应的，如下：

```sql
select all country, province from person
```

- 1

结果如下： 
![这里写图片描述](https://img-blog.csdn.net/20170622231237714?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbG15ODYyNjM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 1.5 与`distinctrow`同义

```sql
select distinctrow expression[,expression...] from tables [where conditions];
```

- 1

这个语句与distinct的作用是相同的。

### 1.6 对`*`的处理

`*`代表整列，使用distinct对`*`操作

```
sql select DISTINCT * from person 
```

相当于

```sql
select DISTINCT id, `name`, country, province, city from person;
```